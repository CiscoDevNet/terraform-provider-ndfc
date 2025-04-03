// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"strings"
	"terraform-provider-ndfc/internal/provider/ndfc"
	"terraform-provider-ndfc/internal/provider/resources/resource_configuration_deploy"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = &ConfigDeployResource{}

var loggingConfigDeploy = "Configuration Deploy Resource"

func NewConfigDeployResource() resource.Resource {
	return &ConfigDeployResource{}
}

type ConfigDeployResource struct {
	client *ndfc.NDFC
}

func (r *ConfigDeployResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	tflog.Debug(ctx, fmt.Sprintf("Start of %s Metadata", loggingConfigDeploy))
	resp.TypeName = req.ProviderTypeName + "_configuration_deploy"
	tflog.Debug(ctx, fmt.Sprintf("End of %s Metadata", loggingConfigDeploy))
}

func (r *ConfigDeployResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	tflog.Debug(ctx, fmt.Sprintf("Start of %s Schema", loggingConfigDeploy))
	resp.Schema = resource_configuration_deploy.ConfigurationDeployResourceSchema(ctx)
	tflog.Debug(ctx, fmt.Sprintf("End of %s Schema", loggingConfigDeploy))
}

func (r *ConfigDeployResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data resource_configuration_deploy.ConfigurationDeployModel
	var serialNumbers []string

	tflog.Debug(ctx, "ValidateConfig called")

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if !data.ConfigSave.ValueBool() && data.SerialNumbers.IsNull() {
		resp.Diagnostics.AddAttributeError(path.Root("config_save"), "config_save and serial_numbers", "At least one of the fields 'config_save=true' or 'serial_numbers' must be set")
		return
	}
	data.SerialNumbers.ElementsAs(ctx, &serialNumbers, false)
	if len(serialNumbers) > 0 {
		for _, serialNumber := range serialNumbers {
			if strings.ToUpper(serialNumber) == "ALL" && len(serialNumbers) > 1 {
				resp.Diagnostics.AddError("Invalid Configuration", "Serial numbers can have values 'ALL' or a list of Serial numbers")
				return
			}
			if serialNumber == "" {
				resp.Diagnostics.AddError("Invalid Configuration", "Serial number cannot be empty")
				return
			}
		}
	}
}

func (r *ConfigDeployResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Debug(ctx, fmt.Sprintf("Start of %s Configure", loggingConfigDeploy))
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*ndfc.NDFC)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *ndfc.NDFC, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
	tflog.Debug(ctx, fmt.Sprintf("End of %s Configure", loggingConfigDeploy))
}

func (r *ConfigDeployResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, fmt.Sprintf("Start of %s Create", loggingConfigDeploy))
	var data resource_configuration_deploy.ConfigurationDeployModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.Deploy(ctx, &resp.Diagnostics, &data)
	data.Id = data.FabricName

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End of %s Create", loggingConfigDeploy))
}

func (r *ConfigDeployResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data resource_configuration_deploy.ConfigurationDeployModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	//  False returned on purpose, so that terraform detects a change in plan and triggers update for this resource during next apply
	data.TriggerDeployOnUpdate = types.BoolValue(false)
	tflog.Debug(ctx, fmt.Sprintf("Start of %s Read", loggingConfigDeploy))
	tflog.Debug(ctx, fmt.Sprintf("End of %s Read data.TriggerDeployOnUpdate %v   ", loggingConfigDeploy, data.TriggerDeployOnUpdate.ValueBool()))
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ConfigDeployResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, fmt.Sprintf("Start of %s Update", loggingConfigDeploy))

	var data resource_configuration_deploy.ConfigurationDeployModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	r.Deploy(ctx, &resp.Diagnostics, &data)
	data.Id = data.FabricName

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End of %s Update", loggingConfigDeploy))
}

func (r *ConfigDeployResource) Deploy(ctx context.Context, dg *diag.Diagnostics, data *resource_configuration_deploy.ConfigurationDeployModel) {

	var serialNumbers []string
	*dg = data.SerialNumbers.ElementsAs(ctx, &serialNumbers, false)
	if dg.HasError() {
		return
	}
	tflog.Debug(ctx, fmt.Sprintf("Serial Numbers: %v", serialNumbers))

	if len(serialNumbers) != 0 {
		FirstIndex := strings.ToUpper(serialNumbers[0])
		if FirstIndex == "ALL" {
			// Deploy configuration for all serial numbers
			r.client.RecalculateAndDeploy(ctx, dg, data.FabricName.ValueString(), data.ConfigSave.ValueBool(), true, nil)
		} else {
			// Deploy configuration for specific serial numbers
			r.client.RecalculateAndDeploy(ctx, dg, data.FabricName.ValueString(), data.ConfigSave.ValueBool(), true, serialNumbers)
		}
	}

}
func (r *ConfigDeployResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, fmt.Sprintf("Start of %s Delete", loggingConfigDeploy))
	tflog.Debug(ctx, fmt.Sprintf("End of %s Delete", loggingConfigDeploy))
}
