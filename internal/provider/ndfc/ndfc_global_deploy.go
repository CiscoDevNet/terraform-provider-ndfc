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

func (c NDFC) SaveConfiguration(ctx context.Context, diags *diag.Diagnostics, fabricName string) {
	saveApi := api.NewConfigDeploymentAPI(c.GetLock(ResourceConfigDeploy), &c.apiClient)
	saveApi.FabricName = fabricName
	saveApi.Deploy = false
	_, err := saveApi.Post([]byte{})
	if err != nil {
		diags.AddError("Config Save failed", err.Error())
		//saveApi.Preview = false
		//time.Sleep(3 * time.Second)
		//res, _ := saveApi.Get()
		// TODO determine which error should be returned
		///diags.AddError("Deploy Errors:", string(res))
	}
}

func (c NDFC) DeployConfiguration(ctx context.Context, diags *diag.Diagnostics, fabricName string, serialNumbers []string) {
	var err error
	deployApi := api.NewConfigDeploymentAPI(c.GetLock(ResourceConfigDeploy), &c.apiClient)
	deployApi.FabricName = fabricName
	deployApi.Deploy = true

	for retry := 0; retry < 3; retry++ {
		serialNumbers = c.CheckDeployStatus(ctx, diags, fabricName, serialNumbers)
		if diags.HasError() {
			return
		}
		if len(serialNumbers) == 0 {
			return
		} else {
			deployApi.SerialNumbers = serialNumbers
		}
		_, err = deployApi.Post([]byte{})
		if err != nil {
			diags.AddError("Deploy failed", err.Error())
			time.Sleep(3 * time.Second)
			deployApi.Preview = false
			res, _ := deployApi.Get()
			// TODO determine which error should be returned
			diags.AddError("Deploy Errors:", string(res))
			return
		}
	}
	diags.AddError("Deploy failed", "Switches are still out of sync")
}

func (c NDFC) CheckDeployStatus(ctx context.Context, diags *diag.Diagnostics, fabricName string, serialNumbers []string) []string {
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
	tempList := serialNumbers
	if len(serialNumbers) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("Switches out of sync: %v", serialNumbers))
		tflog.Debug(ctx, fmt.Sprintf("Switches in sync: %v", response.SerialNumMap))
		for index, serialNumber := range serialNumbers {
			if response.SerialNumMap[serialNumber].Status == inSync {
				// Delete serial number from outOfSync
				tempList = append(serialNumbers[:index], serialNumbers[index+1:]...)
			} else if response.SerialNumMap[serialNumber].Status == failed {
				diags.AddError("Deploy failed", fmt.Sprintf("Deploy failed for serial number %s", serialNumber))
			}
		}
	} else {
		for serialNumber, entry := range response.SerialNumMap {
			if entry.Status == outOfSync {
				// Add serial number from outOfSync
				tempList = append(tempList, serialNumber)
			} else if entry.Status == failed {
				diags.AddError("Deploy failed", fmt.Sprintf("Deploy failed for serial number %s", serialNumber))
				return nil
			}
		}
	}
	tflog.Debug(ctx, fmt.Sprintf("Switches out of sync: %v", tempList))
	return tempList
}
