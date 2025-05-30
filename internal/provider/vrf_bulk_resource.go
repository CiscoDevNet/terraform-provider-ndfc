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
	"terraform-provider-ndfc/internal/provider/resources/resource_vrf_bulk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = (*vrfBulkResource)(nil)
var _ resource.ResourceWithImportState = (*vrfBulkResource)(nil)

func NewVrfBulkResource() resource.Resource {
	return &vrfBulkResource{}
}

type vrfBulkResource struct {
	client *ndfc.NDFC
}

func (r *vrfBulkResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + ndfc.ResourceVrfBulk
}

func (r *vrfBulkResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_vrf_bulk.VrfBulkResourceSchema(ctx)
	//resp.Schema.Attributes["new_attribute"] = schema.BoolAttribute{Optional: true}
}

func (d *vrfBulkResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	tflog.Info(ctx, "vrf_bulk Configure")
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

func (r *vrfBulkResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var in resource_vrf_bulk.VrfBulkModel
	// Read Terraform plan data into the model
	log.Printf("[TRACE] Create VRF Bulk")
	resp.Diagnostics.Append(req.Plan.Get(ctx, &in)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Create API call logic
	vrfDone := r.client.RscCreateBulkVrf(ctx, &resp.Diagnostics, &in)
	if vrfDone == nil {
		tflog.Error(ctx, "Create Bulk VRF Failed")
		return
	}
	// Example data value setting
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, vrfDone)...)
}

func (r *vrfBulkResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data resource_vrf_bulk.VrfBulkModel

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
	dataVrf := data.GetModelData()
	deployMap := make(map[string][]string)

	if dataVrf.DeployAllAttachments {
		log.Printf("[DEBUG] DeployAllAttachments flag set")
		deployMap["global"] = append(deployMap["global"], "all")
	}
	for vrfName, v := range dataVrf.Vrfs {
		v.VrfName = vrfName
		if v.DeployAttachments {
			//first element is the vrf itself - means deploy enabled at vrf level
			deployMap[v.VrfName] = append(deployMap[v.VrfName], v.VrfName)
		}
		for serial, s := range v.AttachList {
			s.SerialNumber = serial
			if s.DeployThisAttachment {
				deployMap[v.VrfName] = append(deployMap[v.VrfName], s.SerialNumber)
			}
			v.AttachList[serial] = s
		}
		dataVrf.Vrfs[vrfName] = v
	}

	tflog.Info(ctx, fmt.Sprintf("Incoming ID %s", unique_id))
	dd := r.client.RscGetBulkVrf(ctx, &resp.Diagnostics, unique_id, &deployMap)
	if dd == nil {
		tflog.Error(ctx, "Read Bulk VRF Failed")
		resp.Diagnostics.AddWarning("Read Failure", "No configuration found in NDFC")
		//resp.Diagnostics.AddError("Read Failure", "No data received from NDFC")

	}
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, dd)...)
}

func (r *vrfBulkResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	var planData resource_vrf_bulk.VrfBulkModel
	var stateData resource_vrf_bulk.VrfBulkModel
	var configData resource_vrf_bulk.VrfBulkModel

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

	// Create API call logic
	r.client.RscUpdateBulkVrf(ctx, &resp.Diagnostics, unique_id, &planData, &stateData, &configData)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Create Bulk VRF Failed")
		return
	}
	//

	// Save updated data into Terraform state

	resp.Diagnostics.Append(resp.State.Set(ctx, &planData)...)
}

func (r *vrfBulkResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data resource_vrf_bulk.VrfBulkModel

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

	r.client.RscDeleteBulkVrf(ctx, &resp.Diagnostics, data.Id.ValueString(), &data)

	// Delete API call logic
}

/* DONOT use the deployment flags at different level at the same time */
/* use only one at a time */
/* eg. if global deploy_all_attachments is set, then do not use deploy_attachments or deploy_this_attachment at vrf/attachment level */
func (r vrfBulkResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data resource_vrf_bulk.VrfBulkModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	globalDeploy := false
	if !data.DeployAllAttachments.IsNull() || !data.DeployAllAttachments.IsUnknown() {
		tflog.Debug(ctx, "DeployAllAttachments is set")
		globalDeploy = data.DeployAllAttachments.ValueBool()
	}
	if !data.Vrfs.IsNull() && !data.Vrfs.IsUnknown() {
		elements1 := make(map[string]resource_vrf_bulk.VrfsValue, len(data.Vrfs.Elements()))
		dg := data.Vrfs.ElementsAs(ctx, &elements1, false)
		if dg.HasError() {
			resp.Diagnostics.AddError("Error in reading Input TF", fmt.Sprintf("%v", dg.Errors()))
			return
		}
		seen := make(map[string]bool)
		for vrf, v := range elements1 {
			vrfDeploy := false
			if got, ok := seen[vrf]; ok && got {
				tflog.Error(ctx, fmt.Sprintf("Duplicate VRF %s", vrf))
				resp.Diagnostics.AddError("Vrfs", fmt.Sprintf("Duplicate entry for VRF %s", vrf))
				return
			}
			seen[vrf] = true
			if !v.DeployAttachments.IsNull() && !v.DeployAttachments.IsUnknown() {
				tflog.Debug(ctx, fmt.Sprintf("DeployAttachments is set for vrf %s", vrf))
				vrfDeploy = v.DeployAttachments.ValueBool()
				if globalDeploy && vrfDeploy {
					tflog.Error(ctx, "Conflicting Deployment Flags - Global & vrf level")
					resp.Diagnostics.AddAttributeError(
						path.Root("deploy_all_attachments"),
						"Conflicting Deployment Flags",
						"Use only one deployment flag at a time",
					)
					resp.Diagnostics.AddAttributeError(
						path.Root("vrfs").AtMapKey(vrf).AtName("deploy_attachments"),
						"Conflicting Deployment Flags",
						"Use only one deployment flag at a time",
					)
					return
				}
			}
			if !v.AttachList.IsNull() && !v.AttachList.IsUnknown() {
				elements2 := make(map[string]resource_vrf_bulk.AttachListValue, len(v.AttachList.Elements()))
				dg := v.AttachList.ElementsAs(ctx, &elements2, false)
				if dg.HasError() {
					resp.Diagnostics.AddError("Error in reading AttachList Input", fmt.Sprintf("%v", dg.Errors()))
					return
				}
				seen := make(map[string]bool)
				for serial, s := range elements2 {
					if got, ok := seen[serial]; ok && got {
						tflog.Error(ctx, fmt.Sprintf("Duplicate Serial %s", serial))
						resp.Diagnostics.AddError("AttachList", fmt.Sprintf("Duplicate entry for Attachment %s/%s", vrf, serial))
						return
					}
					seen[serial] = true
					if !s.DeployThisAttachment.IsNull() && !s.DeployThisAttachment.IsUnknown() {
						attachDeploy := s.DeployThisAttachment.ValueBool()
						tflog.Debug(ctx, fmt.Sprintf("DeployThisAttachment is set for %s/%s", vrf, serial))
						if globalDeploy && attachDeploy {
							tflog.Error(ctx, "Conflicting Deployment Flags - Global & attachment level")
							resp.Diagnostics.AddAttributeError(
								path.Root("deploy_all_attachments"),
								"Conflicting Deployment Flags",
								"Use only one deployment flag at a time",
							)
							resp.Diagnostics.AddAttributeError(
								path.Root("vrfs").AtMapKey(vrf).AtName("attach_list").AtMapKey(serial).AtName("deploy_this_attachment"),
								"Conflicting Deployment Flags",
								"Use only one deployment flag at a time",
							)
							return
						}
						if vrfDeploy && attachDeploy {
							tflog.Error(ctx, "Conflicting Deployment Flags - vrf & attachment level")
							resp.Diagnostics.AddAttributeError(
								path.Root("vrfs").AtMapKey(vrf).AtName("deploy_attachments"),
								"Conflicting Deployment Flags",
								"Use only one deployment flag at a time",
							)
							resp.Diagnostics.AddAttributeError(
								path.Root("vrfs").AtMapKey(vrf).AtName("attach_list").AtMapKey(serial).AtName("deploy_this_attachment"),
								"Conflicting Deployment Flags",
								"Use only one deployment flag at a time",
							)
							return
						}
					}
				}
			}
		}
	}
}

func (r *vrfBulkResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	//
	unique_id := req.ID
	tflog.Info(ctx, fmt.Sprintf("Incoming ID %s", unique_id))
	dd := r.client.RscImportBulkVrf(ctx, &resp.Diagnostics, unique_id)
	if dd == nil {
		tflog.Error(ctx, "Read Bulk VRF Failed")
		resp.Diagnostics.AddWarning("Read Failure", "No configuration found in NDFC")
	}
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, dd)...)
}
