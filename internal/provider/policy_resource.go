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
	"terraform-provider-ndfc/internal/provider/resources/resource_policy"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = (*policyResource)(nil)
var _ resource.ResourceWithImportState = (*policyResource)(nil)

func NewPolicyResource() resource.Resource {
	return &policyResource{}
}

type policyResource struct {
	client *ndfc.NDFC
}

func (r *policyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + ndfc.ResourcePolicy
}

func (r *policyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_policy.PolicyResourceSchema(ctx)
	//resp.Schema.Attributes["new_attribute"] = schema.BoolAttribute{Optional: true}
}

func (d *policyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	tflog.Info(ctx, "Policy Configure")
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

func (r *policyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var in resource_policy.PolicyModel
	// Read Terraform plan data into the model
	log.Printf("[TRACE] Create Policy")
	resp.Diagnostics.Append(req.Plan.Get(ctx, &in)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.client.RscCreatePolicy(ctx, &resp.Diagnostics, &in)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Create Bulk Policy Failed")
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, in)...)
}

func (r *policyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data resource_policy.PolicyModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	dg1 := diag.Diagnostics{}
	r.client.RscReadPolicy(ctx, &dg1, &data)
	if dg1.HasError() {
		tflog.Error(ctx, "Read Policy Failed")
		resp.Diagnostics.AddWarning("Read Failure", "No configuration found in NDFC")
		//resp.Diagnostics.AddError("Read Failure", "No data received from NDFC")
		//resp.Diagnostics.Append(resp.State.Set(ctx, nil)...)
		tflog.Error(ctx, "Read Policy Failed - removing resource from state")
		resp.State.RemoveResource(ctx)
		return
	}
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r *policyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	var planData resource_policy.PolicyModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &planData)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create API call logic
	r.client.RscUpdatePolicy(ctx, &resp.Diagnostics, &planData)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Update Policy Failed")
		return
	}
	//
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &planData)...)
}

func (r *policyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data resource_policy.PolicyModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	r.client.RscDeletePolicy(ctx, &resp.Diagnostics, &data)

	// Delete API call logic
}

func (r policyResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data resource_policy.PolicyModel
	tflog.Debug(ctx, "ValidateConfig called")
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *policyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	//
	unique_id := req.ID
	tflog.Info(ctx, fmt.Sprintf("Incoming ID %s", unique_id))
	data := new(resource_policy.PolicyModel)

	r.client.RscImportPolicy(ctx, &resp.Diagnostics, unique_id, data)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Import  Policy Failed")
		resp.Diagnostics.AddWarning("Import Failure", "No configuration found in NDFC")
	}
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}
