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
	"terraform-provider-ndfc/internal/provider/ndfc/api"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

const ResourceConfigDeploy = "config_deploy"

func (c NDFC) SaveConfiguration(ctx context.Context, diags *diag.Diagnostics, fabricName string) {
	saveApi := api.NewConfigDeploymentAPI(c.GetLock(ResourceConfigDeploy), &c.apiClient)
	saveApi.FabricName = fabricName
	saveApi.Deploy = false
	_, err := saveApi.Post([]byte{})
	if err != nil {
		diags.AddError("Deploy failed", err.Error())
		time.Sleep(3 * time.Second)
		_, err := saveApi.Get()
		// TODO determine which error should be returned
		diags.AddError("Deploy Errors:", err.Error())
	}
}

func (c NDFC) DeployConfiguration(ctx context.Context, diags *diag.Diagnostics, fabricName string, serialNumbers []string) {
	var err error
	deployApi := api.NewConfigDeploymentAPI(c.GetLock(ResourceConfigDeploy), &c.apiClient)
	deployApi.FabricName = fabricName
	deployApi.Deploy = true
	if len(serialNumbers) > 0 {
		deployApi.SerialNumbers = serialNumbers
	}
	_, err = deployApi.Post([]byte{})
	if err != nil {
		diags.AddError("Deploy failed", err.Error())
		time.Sleep(3 * time.Second)
		_, err := deployApi.Get()
		// TODO determine which error should be returned
		diags.AddError("Deploy Errors:", err.Error())
	}
}
