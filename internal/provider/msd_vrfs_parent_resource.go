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
	"log"
	"terraform-provider-ndfc/internal/provider/ndfc"
	"terraform-provider-ndfc/internal/provider/resources/resource_msd_vrfs_parent"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = (*msdVrfsParentResource)(nil)
var _ resource.ResourceWithImportState = (*msdVrfsParentResource)(nil)

func NewMsdVrfsParentResource() resource.Resource {
	return &msdVrfsParentResource{}
}

type msdVrfsParentResource struct {
	client *ndfc.NDFC
}

func (r *msdVrfsParentResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_msd_vrfs_parent"
}

func (r *msdVrfsParentResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_msd_vrfs_parent.MsdVrfsParentResourceSchema(ctx)
}

func (d *msdVrfsParentResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	tflog.Info(ctx, "msd_vrfs_parent Configure")
	client, ok := req.ProviderData.(*ndfc.NDFC)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected resource Configure Type",
			fmt.Sprintf("Expected *ndfc.NDFC, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}
	d.client = client
}

func (r *msdVrfsParentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var in resource_msd_vrfs_parent.MsdVrfsParentModel
	// Read Terraform plan data into the model
	log.Printf("[TRACE] Create MSD VRFs Parent")
	resp.Diagnostics.Append(req.Plan.Get(ctx, &in)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Create API call logic
	vrfDone := r.client.RscCreateMsdVrfsParent(ctx, &resp.Diagnostics, &in)
	if vrfDone == nil {
		tflog.Error(ctx, "Create MSD VRFs Parent Failed")
		return
	}
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, vrfDone)...)
}

func (r *msdVrfsParentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data resource_msd_vrfs_parent.MsdVrfsParentModel

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
	// unique_id = fabric_name/[vrf1,vrf2,vrf3...]

	tflog.Info(ctx, fmt.Sprintf("Incoming ID %s", unique_id))
	dd := r.client.RscGetMsdVrfsParent(ctx, &resp.Diagnostics, unique_id)
	if dd == nil {
		tflog.Error(ctx, "Read MSD VRFs Parent Failed")
		resp.Diagnostics.AddWarning("Read Failure", "No configuration found in NDFC")
	}
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, dd)...)
}

func (r *msdVrfsParentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	var planData resource_msd_vrfs_parent.MsdVrfsParentModel
	var stateData resource_msd_vrfs_parent.MsdVrfsParentModel
	var configData resource_msd_vrfs_parent.MsdVrfsParentModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &planData)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)
	resp.Diagnostics.Append(req.Config.Get(ctx, &configData)...)

	if resp.Diagnostics.HasError() {
		return
	}
	unique_id := stateData.Id.ValueString()
	tflog.Info(ctx, fmt.Sprintf("Incoming ID %s", unique_id))
	if unique_id == "" {
		resp.Diagnostics.AddError("ID cannot be empty for update", "Id is mandatory - State may be corrupted")
		return
	}
	// Update API call logic
	r.client.RscUpdateMsdVrfsParent(ctx, &resp.Diagnostics, unique_id, &planData, &stateData, &configData)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Update MSD VRFs Parent Failed")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &planData)...)
}

func (r *msdVrfsParentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data resource_msd_vrfs_parent.MsdVrfsParentModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Id.IsNull() || data.Id.IsUnknown() {
		resp.Diagnostics.AddError("Delete: Id cannot be empty", "Id should be present")
		resp.State.RemoveResource(ctx)
		return
	}

	r.client.RscDeleteMsdVrfsParent(ctx, &resp.Diagnostics, data.Id.ValueString(), &data)
}

func (r msdVrfsParentResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data resource_msd_vrfs_parent.MsdVrfsParentModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if !data.Vrfs.IsNull() && !data.Vrfs.IsUnknown() {
		elements1 := make(map[string]resource_msd_vrfs_parent.VrfsValue, len(data.Vrfs.Elements()))
		dg := data.Vrfs.ElementsAs(ctx, &elements1, false)
		if dg.HasError() {
			resp.Diagnostics.AddError("Error in reading Input TF", fmt.Sprintf("%v", dg.Errors()))
			return
		}
		seen := make(map[string]bool)
		for vrf, _ := range elements1 {
			if got, ok := seen[vrf]; ok && got {
				tflog.Error(ctx, fmt.Sprintf("Duplicate VRF %s", vrf))
				resp.Diagnostics.AddError("Vrfs", fmt.Sprintf("Duplicate entry for VRF %s", vrf))
				return
			}
			seen[vrf] = true
		}
	}
}

func (r *msdVrfsParentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	unique_id := req.ID
	tflog.Info(ctx, fmt.Sprintf("Incoming ID %s", unique_id))
	dd := r.client.RscImportMsdVrfsParent(ctx, &resp.Diagnostics, unique_id)
	if dd == nil {
		tflog.Error(ctx, "Import MSD VRFs Parent Failed")
		resp.Diagnostics.AddWarning("Import Failure", "No configuration found in NDFC")
	}
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, dd)...)
}
