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
	"terraform-provider-ndfc/internal/provider/ndfc"
	"terraform-provider-ndfc/internal/provider/resources/resource_fabric_ipfm"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = (*fabricIpfmResource)(nil)
var _ resource.ResourceWithImportState = (*fabricIpfmResource)(nil)

func NewFabricIpfmResource() resource.Resource {
	return &fabricIpfmResource{}
}

type fabricIpfmResource struct {
	client *ndfc.NDFC
}

func (r *fabricIpfmResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + ndfc.ResourceIpfmFabric
}

func (r *fabricIpfmResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_fabric_ipfm.FabricIpfmResourceSchema(ctx)
}

func (d *fabricIpfmResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	tflog.Info(ctx, "Fabric Configure")
	client, ok := req.ProviderData.(*ndfc.NDFC)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected resource  Configure Type",
			fmt.Sprintf("Expected *nd.NDFC, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}
	d.client = client
}

func (r *fabricIpfmResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resource_fabric_ipfm.FabricIpfmModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if r.client == nil {
		panic("Client is nil")
	}
	// Create API call logic
	deploy := data.Deploy.ValueBool()
	r.client.RscCreateFabric(ctx, &resp.Diagnostics, &data, ndfc.ResourceIpfmFabricType)
	data.Deploy = types.BoolValue(deploy)
	data.Id = data.FabricName
	tflog.Debug(ctx, "data.Id = "+data.Id.ValueString())
	if deploy {
		if resp.Diagnostics.HasError() || resp.Diagnostics.WarningsCount() > 0 {
			data.DeploymentStatus = types.StringValue("Deployment pending")
		} else {
			data.DeploymentStatus = types.StringValue("Deployment successful")
		}
	} else {
		data.DeploymentStatus = types.StringValue("Deployment pending")
	}
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Create Fabric Failed")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r *fabricIpfmResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data resource_fabric_ipfm.FabricIpfmModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "data.Id = "+data.Id.ValueString())
	if data.Id.IsNull() || data.Id.IsUnknown() {
		resp.Diagnostics.AddError("Id cannot be empty", "Id should be present")
		resp.State.RemoveResource(ctx)
		return
	}
	unique_id := data.Id.ValueString()
	tflog.Info(ctx, fmt.Sprintf("Incoming ID %s", unique_id))
	deploy := data.Deploy.ValueBool()
	r.client.RscReadFabric(ctx, &resp.Diagnostics, &data, ndfc.ResourceIpfmFabricType)
	data.Deploy = types.BoolValue(deploy)
	data.Id = data.FabricName
	tflog.Debug(ctx, "data.FabricName = "+data.FabricName.ValueString())
	if data.FabricName.IsNull() || data.FabricName.IsUnknown() {
		// make diags error empty because fabric is not present in NDFC,
		// it needs to be recreated.
		resp.Diagnostics = diag.Diagnostics{}
		// This will clear the state for current fabric, making it eligible for creation
		resp.State.RemoveResource(ctx)
	} else {
		if resp.Diagnostics.HasError() {
			return
		}
		// Save updated data into Terraform state
		resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
	}

}

func (r *fabricIpfmResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	var planData resource_fabric_ipfm.FabricIpfmModel
	var stateData resource_fabric_ipfm.FabricIpfmModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &planData)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

	if resp.Diagnostics.HasError() {
		return
	}
	// Create API call logic
	deploy := planData.Deploy.ValueBool()
	r.client.RscUpdateFabric(ctx, &resp.Diagnostics, &planData, ndfc.ResourceIpfmFabricType)
	planData.Deploy = types.BoolValue(deploy)
	planData.Id = planData.FabricName
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Update Fabric Failed")
		return
	}
	unique_id := planData.Id.ValueString()
	tflog.Info(ctx, fmt.Sprintf("Update Fabric Success %s", unique_id))
	if deploy {
		if resp.Diagnostics.HasError() || resp.Diagnostics.WarningsCount() > 0 {
			planData.DeploymentStatus = types.StringValue("Deployment pending")
		} else {
			planData.DeploymentStatus = types.StringValue("Deployment successful")
		}
	} else {
		planData.DeploymentStatus = types.StringValue("Deployment pending")
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &planData)...)
}

func (r *fabricIpfmResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data resource_fabric_ipfm.FabricIpfmModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Id.IsNull() || data.Id.IsUnknown() {
		resp.Diagnostics.AddError("Delete: Id cannot be empty", "Id should be present")
		resp.State.RemoveResource(ctx)
		return
	}
	r.client.RscDeleteFabric(ctx, &resp.Diagnostics, &data, ndfc.ResourceIpfmFabricType)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Delete Fabric Failed")
		return
	}
	data.Id = types.StringNull()
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r *fabricIpfmResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data resource_fabric_ipfm.FabricIpfmModel
	tflog.Info(ctx, fmt.Sprintf("Import Fabric Incoming ID %s", req.ID))
	if req.ID == "" {
		resp.Diagnostics.AddError("ID cannot be empty for import", "Id is mandatory")
		return
	}
	data.FabricName = types.StringValue(req.ID)
	r.client.RscImportFabric(ctx, &resp.Diagnostics, &data, ndfc.ResourceIpfmFabricType)
	if resp.Diagnostics.HasError() {
		return
	}
	data.Id = types.StringValue(req.ID)
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)

}
