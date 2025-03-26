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
	"terraform-provider-ndfc/internal/provider/ndfc/api"
	"terraform-provider-ndfc/internal/provider/resources/resource_configuration_deploy"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const ResourceConfigDeploy = "config_deploy"
const inSync = "In-Sync"
const outOfSync = "Out-of-Sync"
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
	_, err := saveApi.DeployPost([]byte{})
	if err != nil {
		diags.AddError("Config Save failed", err.Error())
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

func (c NDFC) checkDeployStatus(ctx context.Context, diags *diag.Diagnostics, fabricName string, serialNumbers []string) []string {
	previewApi := api.NewConfigDeploymentAPI(c.GetLock(ResourceConfigDeploy), &c.apiClient)
	previewApi.FabricName = fabricName
	previewApi.Preview = true
	var response resource_configuration_deploy.SwitchStatusDB

	payload, err := previewApi.Get()
	if err != nil || len(payload) == 0 {
		diags.AddError("Deploy failed", "Configuration preview failed")
		return nil
	}
	err = json.Unmarshal(payload, &response)
	if err != nil {
		diags.AddError("Deploy failed", err.Error())
	}
	ooList := make([]string, 0)

	if len(serialNumbers) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("Switches out of sync: %v", serialNumbers))
		tflog.Debug(ctx, fmt.Sprintf("Switches in sync: %v", response.SerialNumMap))
		for _, serialNumber := range serialNumbers {
			if response.SerialNumMap[serialNumber].Status == inSync {
				// Delete serial number from outOfSync
				log.Printf("Switch %s is in sync - remove from list", serialNumber)
			} else if response.SerialNumMap[serialNumber].Status == failed {
				diags.AddError("Deploy failed", fmt.Sprintf("Deploy failed for serial number %s", serialNumber))
				tflog.Error(ctx, fmt.Sprintf("Deploy failed for serial number %s", serialNumber))
				return nil

			} else if response.SerialNumMap[serialNumber].Status == outOfSync {
				log.Printf("Switch %s is out of sync", serialNumber)
				// Add serial number from outOfSync
				ooList = append(ooList, serialNumber)
			}
		}
	} else {
		for serialNumber, entry := range response.SerialNumMap {
			if entry.Status == outOfSync {
				// Add serial number from outOfSync
				log.Printf("Switch %s is out of sync", serialNumber)
				ooList = append(ooList, serialNumber)
			} else if entry.Status == failed {
				tflog.Error(ctx, fmt.Sprintf("Deploy failed for serial number %s", serialNumber))
				diags.AddError("Deploy failed", fmt.Sprintf("Deploy failed for serial number %s", serialNumber))
				return nil
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
