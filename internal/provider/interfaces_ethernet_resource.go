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
	"terraform-provider-ndfc/internal/provider/resources/resource_interface_ethernet"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = (*interfaceEthernetResource)(nil)

var _ resource.ResourceWithModifyPlan = (*interfaceEthernetResource)(nil)
var _ resource.ResourceWithImportState = (*interfaceEthernetResource)(nil)

func NewInterfaceEthernetResource() resource.Resource {
	return &interfaceEthernetResource{}
}

type interfaceEthernetResource struct {
	client *ndfc.NDFC
}

func (r *interfaceEthernetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + ndfc.ResourceEthernetInterface
}

func (r *interfaceEthernetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_interface_ethernet.InterfaceEthernetResourceSchema(ctx)
}

func (d *interfaceEthernetResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	tflog.Info(ctx, "interfaces Configure")
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

func (r *interfaceEthernetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var in resource_interface_ethernet.InterfaceEthernetModel
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &in)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if r.client == nil {
		panic("Client is nil")
	}
	// Create API call logic
	r.client.RscCreateInterfaces(ctx, resp, &in)
}

func (r *interfaceEthernetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data resource_interface_ethernet.InterfaceEthernetModel

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
	r.client.RscGetInterfaces(ctx, &resp.Diagnostics, &data)
	if resp.Diagnostics.HasError() {
		return
	}
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)

}

func (r *interfaceEthernetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	var planData resource_interface_ethernet.InterfaceEthernetModel
	var stateData resource_interface_ethernet.InterfaceEthernetModel
	var configData resource_interface_ethernet.InterfaceEthernetModel

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
	r.client.RscUpdateInterfaces(ctx, &resp.Diagnostics, unique_id, &planData, &stateData, &configData)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Update interface Failed")
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &planData)...)
}

func (r *interfaceEthernetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data resource_interface_ethernet.InterfaceEthernetModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Id.IsNull() || data.Id.IsUnknown() {
		resp.Diagnostics.AddError("Delete: Id cannot be empty", "Id should be present")
		resp.State.RemoveResource(ctx)
		return
	}
	r.client.RscDeleteInterfaces(ctx, &resp.Diagnostics, &data)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Delete interface Failed")
		return
	}
	data.Id = types.StringNull()
	data.Policy = types.StringNull()
	data.Interfaces = types.MapNull(resource_interface_ethernet.InterfacesValue{}.Type(ctx))
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
	// Delete API call logic
}

/* DONOT use the deployment flags at different level at the same time */
/* use only one at a time */
/* eg. if global deploy_all_attachments is set, then do not use deploy_attachments or deploy_this_attachment at network/attachment level */
func (r interfaceEthernetResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data resource_interface_ethernet.InterfaceEthernetModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	inData := data.GetModelData()

	ifMap := make(map[string]string)
	for k, v := range inData.Interfaces {
		mapKey := v.SerialNumber + ":" + v.InterfaceName
		if dupKey, ok := ifMap[mapKey]; ok {
			resp.Diagnostics.AddError("Duplicate Interface", fmt.Sprintf("Duplicate Interface %s at keys  %s and %s", mapKey, k, dupKey))
		} else {
			ifMap[mapKey] = k
		}
	}
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *interfaceEthernetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data resource_interface_ethernet.InterfaceEthernetModel

	tflog.Info(ctx, fmt.Sprintf("Import Ethernet Intf Incoming ID %s", req.ID))
	if req.ID == "" {
		resp.Diagnostics.AddError("ID cannot be empty for import", "Id is mandatory - State may be corrupted")
		return
	}
	data.Id = types.StringValue(req.ID)
	r.client.RscImportInterfaces(ctx, &resp.Diagnostics, &data)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)

}

/* ModifyPlan is a very tricky override
 * This is implemented to avoid default values being taken from schema attributes, for custom policy cases
 * The plan data modifications is final and TF does not add anything to it
 * The plan written here MUST match with the infra after the operation
 * Note: The incoming config is not filled with default values. TF doesn't fill if ModifyPlan is implemented
 * So it is our responsibility to set default values for "system" policy cases
 * Even computed variables, once set here cannot change after the op.
 * Mark computed with unknown if there is modification in the resource or it is expected to change
 * Using state and plan to determine if the resource is being created or updated
 */
func (r interfaceEthernetResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	tflog.Info(ctx, "interface_ethernet.ModfyPlan: ")
	var configData resource_interface_ethernet.InterfaceEthernetModel
	var stateData resource_interface_ethernet.InterfaceEthernetModel

	if req.Plan.Raw.IsNull() {
		tflog.Info(ctx, "interface_ethernet.ModfyPlan: Plan is empty - Destroy case")
		return
	}

	resp.Diagnostics.Append(req.Config.Get(ctx, &configData)...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "interface_ethernet.ModfyPlan: config Get Failed")
		return
	}

	state_present := false
	if !req.State.Raw.IsNull() {
		state_present = true
		log.Printf("[DEBUG] interface_ethernet.ModfyPlan: being called for update")
		resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)
		if resp.Diagnostics.HasError() {
			tflog.Error(ctx, "interface_ethernet.ModfyPlan: state Get Failed")
			return
		}
	}

	elementState := make(map[string]resource_interface_ethernet.InterfacesValue, len(stateData.Interfaces.Elements()))
	if state_present {
		// Check if the state is empty, then it is a create case
		if !stateData.Interfaces.IsNull() && !stateData.Interfaces.IsUnknown() {
			diag := stateData.Interfaces.ElementsAs(context.Background(), &elementState, false)
			if diag != nil {
				tflog.Error(ctx, "interface_ethernet.ModfyPlan:  ElementsAs Failed")
				resp.Diagnostics.Append(diag...)
				return
			}
		}
	}
	// Set default values at the Model level
	configData.SetDefaultValues()
	elements1 := make(map[string]resource_interface_ethernet.InterfacesValue, len(configData.Interfaces.Elements()))
	dg := configData.Interfaces.ElementsAs(context.Background(), &elements1, false)
	if dg != nil {
		tflog.Error(ctx, "interface_ethernet.ModfyPlan:  ElementsAs Failed")
		resp.Diagnostics.Append(dg...)
		return
	}
	log.Printf("[DEBUG] interface_ethernet.ModfyPlan:  - Number of entries %d", len(configData.Interfaces.Elements()))
	for k, v := range elements1 {
		if configData.PolicyType.ValueString() == "system" {
			log.Printf("[DEBUG] interface_ethernet.ModfyPlan:  - System Policy")
			// CustomPolicyParameters is expected to be empty for system template
			v.CustomPolicyParameters = types.MapNull(types.StringType)
			// Set default values at intf level
			v.SetDefaultValues()
		} else {
			log.Printf("[DEBUG] interface_ethernet.ModfyPlan:  - Custom Policy")
		}
		if state_present {
			if ifEntry, ok := elementState[k]; ok {
				if ifEntry.DeploymentStatus.IsNull() || ifEntry.DeploymentStatus.IsUnknown() {
					log.Printf("[DEBUG] interface_ethernet.ModfyPlan:  - DeploymentStatus in state is empty")
					v.DeploymentStatus = types.StringUnknown()
				} else {
					log.Printf("[DEBUG] interface_ethernet.ModfyPlan:  - DeploymentStatus in state %s", elementState[k].DeploymentStatus.ValueString())
					v.DeploymentStatus = ifEntry.DeploymentStatus
				}
			}

		} else {
			v.DeploymentStatus = types.StringUnknown()
		}
		elements1[k] = v
		log.Printf("[DEBUG] interface_ethernet.ModfyPlan:  - Setting plan %s=%s", k, v.InterfaceName.ValueString())
	}
	if state_present {
		// Write interfaces back to configData before comparing so that defaults are factored in
		configData.Interfaces, dg = types.MapValueFrom(ctx, resource_interface_ethernet.InterfacesValue{}.Type(context.Background()), elements1)
		if dg != nil {
			tflog.Error(ctx, "interface_ethernet.ModfyPlan:  MapValueFrom Failed")
			resp.Diagnostics.Append(dg...)
			return
		}
		cfgDataModel := configData.GetModelData()
		stateDataModel := stateData.GetModelData()
		rscModified := false
		// Look for modifications - if so computed attributes have to be set to unknown
		// If not, when they are set from NDFC TF would complain that they are modified
		updates := r.client.IntfRscModified(ctx, &resp.Diagnostics, cfgDataModel, stateDataModel)
		for kk := range updates {
			if updates[kk] {
				log.Printf("[DEBUG] %s is modified - Setting computed value to unknown", kk)
				if intf, ok := elements1[kk]; ok {
					intf.DeploymentStatus = types.StringUnknown()
					elements1[kk] = intf
					rscModified = true
				} else {
					rscModified = true
				}
			}
		}
		// Set Id to unknown if resource is modified
		// Otherwise set it to state value if state value is present
		// If nothing is present, set it to unknown
		if rscModified {
			log.Printf("[DEBUG] interface_ethernet.ModfyPlan:  - Resource is modified")
			configData.Id = types.StringUnknown()
		} else if state_present && !stateData.Id.IsNull() && !stateData.Id.IsUnknown() {
			configData.Id = stateData.Id
		}
	} else {
		configData.Id = types.StringUnknown()
	}
	// A second time write of interfaces back to configData is needed as computed attributes may be set to unknown
	configData.Interfaces, dg = types.MapValueFrom(ctx, resource_interface_ethernet.InterfacesValue{}.Type(context.Background()), elements1)
	if dg != nil {
		tflog.Error(ctx, "interface_ethernet.ModfyPlan:  MapValueFrom Failed")
		resp.Diagnostics.Append(dg...)
		return
	}
	dd := resp.Plan.Set(ctx, &configData)
	if dd.HasError() {
		tflog.Error(ctx, "interface_ethernet.ModfyPlan:  Set Failed")
		log.Printf("[DEBUG] interface_ethernet.ModfyPlan:  - Set failed %v", dd.Errors())
		resp.Diagnostics.Append(dd...)
		return
	}
	log.Printf("[DEBUG] interface_ethernet.ModfyPlan:  - Plan Set Dg: %v", resp.Diagnostics)
}
