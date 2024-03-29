// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

//template:begin imports
import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nd"
	"github.com/netascode/terraform-provider-ndfc/internal/provider/helpers"
)

//template:end imports

//template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &InterfaceEthernetResource{}
var _ resource.ResourceWithImportState = &InterfaceEthernetResource{}

func NewInterfaceEthernetResource() resource.Resource {
	return &InterfaceEthernetResource{}
}

type InterfaceEthernetResource struct {
	client      *nd.Client
	updateMutex *sync.Mutex
}

func (r *InterfaceEthernetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_ethernet"
}

func (r *InterfaceEthernetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Interface Ethernet.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"serial_number": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Serial number of switch to configure").String,
				Optional:            true,
			},
			"interface_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the Interface. Example: `Ethernet1/3`").String,
				Optional:            true,
			},
			"policy": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the policy. Examples: `int_trunk_host`, `int_access_host`").AddDefaultValueDescription("int_trunk_host").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("int_trunk_host"),
			},
			"bpdu_guard": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable spanning-tree bpduguard: true='enable', false='disable', no='return to default settings'").AddStringEnumDescription("true", "false", "no").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("true", "false", "no"),
				},
				Default: stringdefault.StaticString("true"),
			},
			"port_type_fast": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable spanning-tree edge port behavior").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"mtu": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("MTU for the interface").AddStringEnumDescription("default", "jumbo").AddDefaultValueDescription("jumbo").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("default", "jumbo"),
				},
				Default: stringdefault.StaticString("jumbo"),
			},
			"speed": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interface speed").AddStringEnumDescription("Auto", "10Mb", "100Mb", "1Gb", "2.5Gb", "5Gb", "10Gb", "25Gb", "40Gb", "50Gb", "100Gb", "200Gb", "400Gb").AddDefaultValueDescription("Auto").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Auto", "10Mb", "100Mb", "1Gb", "2.5Gb", "5Gb", "10Gb", "25Gb", "40Gb", "50Gb", "100Gb", "200Gb", "400Gb"),
				},
				Default: stringdefault.StaticString("Auto"),
			},
			"access_vlan": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Access VLAN ID").AddIntegerRangeDescription(1, 4094).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 4094),
				},
			},
			"interface_description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interface description").String,
				Optional:            true,
			},
			"orphan_port": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("If enabled, configure the interface as a vPC orphan port to be suspended by the secondary peer in vPC failures").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"freeform_config": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Additional CLI for the interface").String,
				Optional:            true,
			},
			"admin_state": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable or disable the interface").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"ptp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable PTP").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"netflow": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Netflow is supported only if it is enabled on fabric").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"netflow_monitor": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Provide the Layer 2 Monitor Name").String,
				Optional:            true,
			},
			"netflow_sampler": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Netflow sampler name, applicable to N7K only").String,
				Optional:            true,
			},
			"allowed_vlans": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Allowed vlans for the ethernet interface. Allowed values are `none`, `all` or VLAN ranges (1-200,500-2000,3000)").AddDefaultValueDescription("none").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("none"),
			},
			"native_vlan": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set native VLAN for the interface").AddIntegerRangeDescription(1, 4094).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 4094),
				},
			},
		},
	}
}

func (r *InterfaceEthernetResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*NdfcProviderData).Client
	r.updateMutex = req.ProviderData.(*NdfcProviderData).UpdateMutex
}

//template:end model

func (r *InterfaceEthernetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceEthernet

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx)

	res, err := r.client.Put(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	// Deploy interface
	r.updateMutex.Lock()
	defer r.updateMutex.Unlock()
	diags = helpers.DeployInterface(ctx, r.client, plan.SerialNumber.ValueString(), plan.InterfaceName.ValueString())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.Id = types.StringValue(plan.SerialNumber.ValueString() + "/" + plan.InterfaceName.ValueString())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *InterfaceEthernetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceEthernet

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	res, err := r.client.Get(fmt.Sprintf("%v?serialNumber=%v&ifName=%v", state.getPath(), state.SerialNumber.ValueString(), state.InterfaceName.ValueString()))
	if err != nil {
		if strings.Contains(err.Error(), "StatusCode 400") || strings.Contains(err.Error(), "StatusCode 500") {
			resp.State.RemoveResource(ctx)
			return
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
			return
		}
	}

	state.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *InterfaceEthernetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan InterfaceEthernet

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	body := plan.toBody(ctx)
	res, err := r.client.Put(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	// Deploy interface
	r.updateMutex.Lock()
	defer r.updateMutex.Unlock()
	diags = helpers.DeployInterface(ctx, r.client, plan.SerialNumber.ValueString(), plan.InterfaceName.ValueString())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *InterfaceEthernetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceEthernet

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

//template:begin import
func (r *InterfaceEthernetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	idParts := strings.Split(req.ID, ":")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: '<serial_number>:<interface_name>'. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("serial_number"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("interface_name"), idParts[1])...)
}

//template:end import
