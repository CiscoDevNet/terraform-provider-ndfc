// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
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
	"terraform-provider-ndfc/internal/provider/resources/resource_links"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// The resource will be registered in the provider's Resources() method

var _ resource.Resource = (*linksResource)(nil)

func NewLinksResource() resource.Resource {
	return &linksResource{}
}

type linksResource struct {
	client *ndfc.NDFC
}

func (r *linksResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_links"
}

func (r *linksResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_links.LinksResourceSchema(ctx)
}

func (r *linksResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
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
}

func (r *linksResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan resource_links.LinksModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating link")

	// Create the link using the NDFC client
	r.client.RscCreateLinks(ctx, resp, &plan)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Created link with ID: %s", plan.LinkUuid.ValueString()))

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *linksResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_links.LinksModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if state.LinkUuid.IsNull() || state.LinkUuid.IsUnknown() {
		resp.Diagnostics.AddError("Link ID cannot be empty", "Link ID should be present")
		resp.State.RemoveResource(ctx)
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Reading link with ID: %s", state.LinkUuid.ValueString()))

	// Read the link using the NDFC client
	r.client.RscReadLinks(ctx, &resp.Diagnostics, &state, nil)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Read link with ID: %s", state.LinkUuid.ValueString()))

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *linksResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var planData, stateData, configData resource_links.LinksModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &planData)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)
	resp.Diagnostics.Append(req.Config.Get(ctx, &configData)...)

	if resp.Diagnostics.HasError() {
		return
	}

	unique_id := stateData.LinkUuid.ValueString()
	tflog.Debug(ctx, fmt.Sprintf("Updating link with ID: %s", unique_id))
	if unique_id == "" {
		resp.Diagnostics.AddError("ID cannot be empty for update", "Id is mandatory - State may be corrupted")
		return
	}

	// Update the link using the NDFC client
	r.client.RscUpdateLinks(ctx, &resp.Diagnostics, unique_id, &planData, &stateData, &configData)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Update link Failed")
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Updated link with ID: %s", planData.LinkUuid.ValueString()))

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &planData)...)
}

func (r *linksResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_links.LinksModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if state.LinkUuid.IsNull() || state.LinkUuid.IsUnknown() {
		resp.Diagnostics.AddError("Delete: Link ID cannot be empty", "Link ID should be present")
		resp.State.RemoveResource(ctx)
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Deleting link with ID: %s", state.LinkUuid.ValueString()))

	// Delete the link using the NDFC client
	r.client.RscDeleteLinks(ctx, &resp.Diagnostics, &state)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Delete link Failed")
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Deleted link with ID: %s", state.LinkUuid.ValueString()))

	// Remove resource from state
	resp.State.RemoveResource(ctx)
}

func (r *linksResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Implement import functionality if needed
}
