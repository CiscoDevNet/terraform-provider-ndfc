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
	"terraform-provider-ndfc/internal/provider/resources/resource_fabric_vxlan_msd"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = (*fabricVxlanMsdResource)(nil)
var _ resource.ResourceWithImportState = (*fabricVxlanMsdResource)(nil)

func NewFabricVxlanMsdResource() resource.Resource {
	return &fabricVxlanMsdResource{}
}

type fabricVxlanMsdResource struct {
	client *ndfc.NDFC
}

func (r *fabricVxlanMsdResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + ndfc.ResourceVxlanMsdFabric
}

func (r *fabricVxlanMsdResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_fabric_vxlan_msd.FabricVxlanMsdResourceSchema(ctx)
}

func (d *fabricVxlanMsdResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *fabricVxlanMsdResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resource_fabric_vxlan_msd.FabricVxlanMsdModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if r.client == nil {
		panic("Client is nil")
	}
	// Values are not updated part of NDFC payload get, save them and set them back
	deploy := data.Deploy.ValueBool()
	// There is no fabric deploy for MSD fabric
	data.Deploy = types.BoolValue(false)
	childFabrics := data.ChildFabrics
	r.client.RscCreateFabric(ctx, &resp.Diagnostics, &data)

	// Setting back the values that are not updated by NDFC
	data.ChildFabrics = childFabrics
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
	r.client.AddChildFabricsToMsd(ctx, &resp.Diagnostics, &data)
	if resp.Diagnostics.HasError() {
		// Cleanup - delete the fabric created
		r.client.RscDeleteFabric(ctx, &resp.Diagnostics, data.FabricName.ValueString())
		return
	}
	// Get and set back the child fabrics to make sure they are correctly added in NDFC
	data.ChildFabrics = r.GetChildFabrics(ctx, &resp.Diagnostics, data.FabricName.ValueString())
	if resp.Diagnostics.HasError() {
		// Cleanup - delete the fabric created
		r.client.RscDeleteFabric(ctx, &resp.Diagnostics, data.FabricName.ValueString())
		return
	}
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r *fabricVxlanMsdResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data resource_fabric_vxlan_msd.FabricVxlanMsdModel

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
	id := data.Id.ValueString()
	tflog.Info(ctx, fmt.Sprintf("Incoming ID %s", id))
	deploy := data.Deploy.ValueBool()

	r.client.RscReadFabric(ctx, &resp.Diagnostics, &data, data.FabricName.ValueString())
	data.Deploy = types.BoolValue(deploy)
	data.Id = data.FabricName
	data.ChildFabrics = r.GetChildFabrics(ctx, &resp.Diagnostics, data.FabricName.ValueString())
	if resp.Diagnostics.HasError() {
		return
	}

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

func (r *fabricVxlanMsdResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	var planData resource_fabric_vxlan_msd.FabricVxlanMsdModel
	var stateData resource_fabric_vxlan_msd.FabricVxlanMsdModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &planData)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

	if resp.Diagnostics.HasError() {
		return
	}
	// Create API call logic
	deploy := planData.Deploy.ValueBool()
	// There is no fabric deploy for MSD fabric
	planData.Deploy = types.BoolValue(false)
	childFabrics := planData.ChildFabrics

	r.client.RscUpdateFabric(ctx, &resp.Diagnostics, &planData)
	planData.Deploy = types.BoolValue(deploy)
	planData.Id = planData.FabricName
	planData.ChildFabrics = childFabrics
	if resp.Diagnostics.HasError() {
		return
	}

	if deploy {
		if resp.Diagnostics.HasError() || resp.Diagnostics.WarningsCount() > 0 {
			planData.DeploymentStatus = types.StringValue("Deployment pending")
		} else {
			planData.DeploymentStatus = types.StringValue("Deployment successful")
		}
	} else {
		planData.DeploymentStatus = types.StringValue("Deployment pending")
	}
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Update Fabric Failed")
		return
	}
	id := planData.Id.ValueString()
	tflog.Info(ctx, fmt.Sprintf("Update Fabric Success %s", id))
	r.client.UpdateChildFabricsToMsd(ctx, &resp.Diagnostics, &planData, &stateData)
	// Get and set back the child fabrics to make sure they are correctly added in NDFC
	planData.ChildFabrics = r.GetChildFabrics(ctx, &resp.Diagnostics, planData.FabricName.ValueString())
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &planData)...)
}

func (r *fabricVxlanMsdResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data resource_fabric_vxlan_msd.FabricVxlanMsdModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Id.IsNull() || data.Id.IsUnknown() {
		resp.Diagnostics.AddError("Delete: Id cannot be empty", "Id should be present")
		resp.State.RemoveResource(ctx)
		return
	}
	tflog.Debug(ctx, fmt.Sprintf("Delete data = %+v", data))
	r.client.RemoveChildFabricsFromMsd(ctx, &resp.Diagnostics, &data)
	if resp.Diagnostics.HasError() {
		return
	}
	// Check child fabrics are removed from NDFC
	data.ChildFabrics = r.GetChildFabrics(ctx, &resp.Diagnostics, data.FabricName.ValueString())
	if resp.Diagnostics.HasError() {
		return
	}
	if len(data.ChildFabrics.Elements()) > 0 {
		resp.Diagnostics.AddError("Failed to delete child fabrics", fmt.Sprintf("Child fabrics %v are not deleted", data.ChildFabrics.Elements()))
		return
	}
	r.client.RscDeleteFabric(ctx, &resp.Diagnostics, data.FabricName.ValueString())
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Delete Fabric Failed")
		return
	}
	data.Id = types.StringNull()
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r *fabricVxlanMsdResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data resource_fabric_vxlan_msd.FabricVxlanMsdModel
	tflog.Info(ctx, fmt.Sprintf("Import Fabric Incoming ID %s", req.ID))
	if req.ID == "" {
		resp.Diagnostics.AddError("ID cannot be empty for import", "Id is mandatory")
		return
	}
	data.FabricName = types.StringValue(req.ID)
	r.client.RscImportFabric(ctx, &resp.Diagnostics, &data, req.ID)
	if resp.Diagnostics.HasError() {
		return
	}
	data.ChildFabrics = r.GetChildFabrics(ctx, &resp.Diagnostics, data.FabricName.ValueString())
	if resp.Diagnostics.HasError() {
		return
	}
	data.Id = types.StringValue(req.ID)
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r *fabricVxlanMsdResource) GetChildFabrics(ctx context.Context, dg *diag.Diagnostics, fname string) basetypes.SetValue {
	var ChildFabrics basetypes.SetValue

	childFabrics := r.client.GetMsdChildFabricAssociations(ctx, dg, fname)
	tflog.Debug(ctx, fmt.Sprintf("Child fabrics part of %s = %+v", fname, childFabrics))

	listData := make([]attr.Value, len(childFabrics))
	for i, item := range childFabrics {
		listData[i] = types.StringValue(item)
	}

	if len(listData) == 0 {
		return types.SetNull(types.StringType)
	} else {
		ChildFabrics, *dg = types.SetValue(types.StringType, listData)
		if dg.HasError() {
			return types.SetNull(types.StringType)
		}
		return ChildFabrics
	}
}
