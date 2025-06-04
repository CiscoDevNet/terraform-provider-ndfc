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
	"terraform-provider-ndfc/internal/provider/resources/resource_vpc_pair"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = (*vpcPairResource)(nil)
var _ resource.ResourceWithImportState = (*vpcPairResource)(nil)

func NewVpcPairResource() resource.Resource {
	return &vpcPairResource{}
}

type vpcPairResource struct {
	client *ndfc.NDFC
}

func (r *vpcPairResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + ndfc.ResourceVpcPair
}

func (r *vpcPairResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_vpc_pair.VpcPairResourceSchema(ctx)
}

func (d *vpcPairResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	tflog.Info(ctx, "vPC Pair Configure")
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

func (r *vpcPairResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var in resource_vpc_pair.VpcPairModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &in)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if r.client == nil {
		panic("Client is nil")
	}
	// Create API call logic
	r.client.RscCreateVpcPair(ctx, &resp.Diagnostics, &in)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Create vPC Pair Failed")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, in)...)
}

func (r *vpcPairResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data resource_vpc_pair.VpcPairModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() || data.Id.IsUnknown() {
		resp.Diagnostics.AddError("Id cannot be empty", "Id should be present")
		resp.State.RemoveResource(ctx)
		return
	}

	unique_id := data.Id.ValueString()
	tflog.Info(ctx, fmt.Sprintf("Incoming ID %s", unique_id))
	err := r.client.RscReadVpcPair(ctx, resp, &data)
	/*Check what error is thrown by the API.
	* The vPC pair API get returns all vPC pairs in the system if it does not match the given ID.
	* We need to check if the vPC pair was deleted by scanning through the list and
	* verifying if the given serial numbers are removed.
	* Now, if it is not present, we return an error as ErrVpcPairNotInFabric.
	* Therefore, in the below case if err is ErrVpcPairNotInFabric, we need to recreate the vPC pair.
	*/
	if err != nil && err.Error() == ndfc.ErrVpcPairNotInFabric {
		// make diags error empty because vPC pair is not present in NDFC,
		// it needs to be recreated.
		tflog.Debug(ctx, "vPC Pair is not present in NDFC")
		resp.Diagnostics = diag.Diagnostics{}
		// This will clear the state for current fabric, making it eligible for creation
		resp.State.RemoveResource(ctx)
	} else if err != nil {
		tflog.Error(ctx, "Read vPC Pair Failed")
		return
	} else {
		// Save updated data into Terraform state
		resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
	}

}

func (r *vpcPairResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	var planData resource_vpc_pair.VpcPairModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &planData)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create API call logic
	r.client.RscUpdateVpcPair(ctx, &resp.Diagnostics, &planData)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Update Vpc Pair Failed")
		return
	}
	unique_id := planData.Id.ValueString()
	tflog.Info(ctx, fmt.Sprintf("Update vPC Pair Success %s", unique_id))

	resp.Diagnostics.Append(resp.State.Set(ctx, &planData)...)
}

func (r *vpcPairResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data resource_vpc_pair.VpcPairModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Id.IsNull() || data.Id.IsUnknown() {
		resp.Diagnostics.AddError("Delete: Id cannot be empty", "Id should be present")
		resp.State.RemoveResource(ctx)
		return
	}
	r.client.RscDeleteVpcPair(ctx, &resp.Diagnostics, &data)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Delete vPC Pair Failed")
		return
	}
	data.Id = types.StringNull()
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r *vpcPairResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data resource_vpc_pair.VpcPairModel
	tflog.Info(ctx, fmt.Sprintf("Import vPC Pair Incoming ID %s", req.ID))
	if req.ID == "" {
		resp.Diagnostics.AddError("ID cannot be empty for import", "Id is mandatory")
		return
	}
	data.Id = types.StringValue(req.ID)
	r.client.RscImportVpcPairs(ctx, &resp.Diagnostics, &data)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)

}
func (r vpcPairResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data resource_vpc_pair.VpcPairModel
	tflog.Debug(ctx, "ValidateConfig called")
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
