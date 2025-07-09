// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package ndfc

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"terraform-provider-ndfc/internal/provider/ndfc/api"
	"terraform-provider-ndfc/internal/provider/resources/resource_configuration_deploy"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const ResourceConfigDeploy = "config_deploy"
const inSync = "In-Sync"
const failed = "Failed"
const deployRetry = 5
const deployRetryInterval = 10

func (c NDFC) RecalculateAndDeploy(ctx context.Context, diags *diag.Diagnostics, fabricName string,
	saveConfig bool, deployConfig bool, serialNumbers []string) {
	// Take deploy Write Lock - wait until all resources have unlocked their reads
	GlobalDeployLock("config_deploy")
	defer GlobalDeployUnlock("config_deploy")
	if saveConfig {
		c.saveConfiguration(diags, fabricName)
		if diags.HasError() {
			return
		}
		time.Sleep(10 * time.Second)
	}
	if deployConfig {
		c.deployConfiguration(ctx, diags, fabricName, serialNumbers)
		if diags.HasError() {
			return
		}
	}
}
func (c NDFC) saveConfiguration(diags *diag.Diagnostics, fabricName string) {
	saveApi := api.NewConfigDeploymentAPI(c.GetLock(ResourceConfigDeploy), &c.apiClient)
	saveApi.FabricName = fabricName
	saveApi.Deploy = false
	//saveApi.SetDeployLocked()
	res, err := saveApi.DeployPost([]byte{})
	if err != nil {
		diags.AddError("config save failed", fmt.Sprintf("Error: %s response: %v", err.Error(), res))
		return
	}
}

func (c NDFC) deployConfiguration(ctx context.Context, diags *diag.Diagnostics, fabricName string, serialNumbers []string) {
	var err error
	deployApi := api.NewConfigDeploymentAPI(c.GetLock(ResourceConfigDeploy), &c.apiClient)
	deployApi.FabricName = fabricName
	deployApi.Deploy = true

	for range deployRetry {
		serialNumbers = c.checkDeployStatus(ctx, diags, fabricName, serialNumbers)
		if diags.HasError() {
			return
		}
		if len(serialNumbers) == 0 {
			return
		} else {
			deployApi.SerialNumbers = serialNumbers
		}
		//deployApi.SetDeployLocked()
		_, err = deployApi.DeployPost([]byte{})
		if err != nil {
			diags.AddError("Deploy failed", err.Error())
			time.Sleep(deployRetryInterval * time.Second)
			deployApi.Preview = false
			res, _ := deployApi.Get()
			// TODO determine which error should be returned
			diags.AddError("Deploy Errors:", string(res))
			return
		}
	}
	diags.AddError("Deploy failed", "Switches are still out of sync")
}

func (c *NDFC) checkDeployStatus(ctx context.Context, diags *diag.Diagnostics, fabricName string, serialNumbers []string) []string {
	previewApi := api.NewConfigDeploymentAPI(c.GetLock(ResourceConfigDeploy), &c.apiClient)
	previewApi.FabricName = fabricName
	previewApi.Preview = true
	var response resource_configuration_deploy.SwitchStatusDB

	// Config Preview refreshes the config status of switches in the fabric
	payload, err := previewApi.Get()
	if len(payload) == 0 || string(payload) == "[]" || err != nil {
		diags.AddError("Deploy failed", "Configuration preview failed")
		return nil
	}

	// Get the current config status of switches in the fabric
	payload, err = c.GetSwitchesInFabric(ctx, fabricName)
	if len(payload) == 0 || string(payload) == "[]" || err != nil {
		diags.AddError("Deploy failed", "Failed to get switches in fabric")
		return nil
	}

	err = json.Unmarshal(payload, &response)
	if err != nil {
		diags.AddError("Deploy failed", err.Error())
	}
	ooList := make([]string, 0)

	if len(serialNumbers) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("Switches out of sync: %v", serialNumbers))
		tflog.Debug(ctx, fmt.Sprintf("Current Switch status: %v", response.SerialNumMap))
		for _, serialNumber := range serialNumbers {
			switch response.SerialNumMap[serialNumber].Status {
			case inSync:
				// Switch is in sync
				log.Printf("Switch %s is in sync", serialNumber)
			case failed:
				diags.AddError("Deploy failed", fmt.Sprintf("Deploy failed for serial number %s", serialNumber))
				tflog.Error(ctx, fmt.Sprintf("Deploy failed for serial number %s", serialNumber))
				return nil
			default:
				log.Printf("Switch %s is in status %s", serialNumber, response.SerialNumMap[serialNumber].Status)
				// Switch status is not inSync or failed, could be pending or outOfSync
				ooList = append(ooList, serialNumber)
			}
		}
	} else {
		for serialNumber, entry := range response.SerialNumMap {
			switch entry.Status {
			case failed:
				tflog.Error(ctx, fmt.Sprintf("Deploy failed for serial number %s", serialNumber))
				diags.AddError("Deploy failed", fmt.Sprintf("Deploy failed for serial number %s", serialNumber))
				return nil
			case inSync:
				// Switch is in sync
				log.Printf("Switch %s is in sync", serialNumber)
			default:
				// Switch status is not inSync or failed, could be pending or outOfSync
				log.Printf("Switch %s is in status %s", serialNumber, entry.Status)
				ooList = append(ooList, serialNumber)
			}
		}
	}
	if len(ooList) == 0 {
		tflog.Debug(ctx, "All switches are in sync")
		return nil
	}
	tflog.Debug(ctx, fmt.Sprintf("List of Switches that are out of sync: {%v}", ooList))
	return ooList
}

// GetDeploymentHistory retrieves the deployment history for a specific fabric
// If serialNumber is provided, it will filter the results to that specific switch
// Returns DeployResponses and error if any
func (c *NDFC) GetDeploymentHistory(ctx context.Context, fabricName string, serialNumber []string) (resource_configuration_deploy.DeployResponses, error) {
	// URL for deployment history API

	configDeployAPI := api.NewConfigDeploymentAPI(c.GetLock(ResourceConfigDeploy), &c.apiClient)
	configDeployAPI.History = true
	configDeployAPI.FabricName = fabricName
	configDeployAPI.SerialNumbers = serialNumber

	// Make the API request
	res, err := configDeployAPI.Get()
	if err != nil {
		tflog.Error(ctx, "Failed to retrieve deployment history", map[string]interface{}{
			"error":  err.Error(),
			"fabric": fabricName,
		})
		return nil, fmt.Errorf("failed to retrieve deployment history for fabric %s: %v", fabricName, err)
	}

	// Check for empty response
	if len(res) == 0 || string(res) == "[]" {
		tflog.Warn(ctx, "No deployment history found", map[string]interface{}{
			"fabric": fabricName,
		})
		return resource_configuration_deploy.DeployResponses{}, nil
	}

	// Parse the response
	var deployResponses resource_configuration_deploy.DeployResponses
	err = json.Unmarshal(res, &deployResponses)
	if err != nil {
		tflog.Error(ctx, "Failed to parse deployment history", map[string]interface{}{
			"error":  err.Error(),
			"fabric": fabricName,
		})
		return nil, fmt.Errorf("failed to parse deployment history: %v", err)
	}
	return deployResponses, nil
}

// GetDeploymentHistoryWithFilters provides more detailed filtering options for deployment history
func (c *NDFC) GetDeploymentHistoryWithFilters(ctx context.Context, fabricName string, serialNumber []string, status, user, startTime, endTime string) (resource_configuration_deploy.DeployResponses, error) {
	// Get all deployment history first
	allResponses, err := c.GetDeploymentHistory(ctx, fabricName, serialNumber)
	if err != nil {
		return nil, err
	}

	// No need to filter further if we have no responses or only filtering by serial number
	if len(allResponses) == 0 || (status == "" && user == "" && startTime == "" && endTime == "") {
		return allResponses, nil
	}

	// Parse time bounds if provided
	var startTimeParsed, endTimeParsed time.Time
	var startTimeErr, endTimeErr error

	if startTime != "" {
		startTimeParsed, startTimeErr = time.Parse("2006-01-02 15:04:05", startTime)
		if startTimeErr != nil {
			return nil, fmt.Errorf("invalid start time format (expected 'YYYY-MM-DD HH:MM:SS'): %v", startTimeErr)
		}
	}

	if endTime != "" {
		endTimeParsed, endTimeErr = time.Parse("2006-01-02 15:04:05", endTime)
		if endTimeErr != nil {
			return nil, fmt.Errorf("invalid end time format (expected 'YYYY-MM-DD HH:MM:SS'): %v", endTimeErr)
		}
	}

	// Apply filters
	var filteredResponses resource_configuration_deploy.DeployResponses

	for _, resp := range allResponses {
		// Filter by status
		if status != "" && !strings.EqualFold(resp.Status, status) {
			continue
		}

		// Filter by user
		if user != "" && !strings.EqualFold(resp.User, user) {
			continue
		}

		// Filter by time range
		if startTime != "" || endTime != "" {
			submittedTime, err := resp.GetSubmittedTime()
			if err != nil {
				continue // Skip entries with unparseable time formats
			}

			// Check if before start time
			if startTime != "" && submittedTime.Before(startTimeParsed) {
				continue
			}

			// Check if after end time
			if endTime != "" && submittedTime.After(endTimeParsed) {
				continue
			}
		}

		// Passed all filters
		filteredResponses = append(filteredResponses, resp)
	}

	return filteredResponses, nil
}
