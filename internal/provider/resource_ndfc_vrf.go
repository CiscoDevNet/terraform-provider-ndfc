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
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nd"
	"github.com/netascode/terraform-provider-ndfc/internal/provider/helpers"
	"github.com/tidwall/sjson"
)

//template:end imports

//template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &VRFResource{}
var _ resource.ResourceWithImportState = &VRFResource{}

func NewVRFResource() resource.Resource {
	return &VRFResource{}
}

type VRFResource struct {
	client      *nd.Client
	updateMutex *sync.Mutex
}

func (r *VRFResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vrf"
}

func (r *VRFResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a VRF.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"fabric_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The name of the fabric").String,
				Optional:            true,
			},
			"vrf_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The name of the VRF").String,
				Required:            true,
			},
			"vrf_template": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The name of the VRF template").AddDefaultValueDescription("Default_VRF_Universal").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("Default_VRF_Universal"),
			},
			"vrf_extension_template": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The name of the VRF extension template").AddDefaultValueDescription("Default_VRF_Extension_Universal").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("Default_VRF_Extension_Universal"),
			},
			"vrf_id": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("VNI ID of VRF").AddIntegerRangeDescription(1, 16777214).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 16777214),
				},
			},
			"vlan_id": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("VLAN ID").AddIntegerRangeDescription(2, 4094).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(2, 4094),
				},
			},
			"vlan_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("VLAN name").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^[^\?,\\,\s]{1,128}$`), ""),
				},
			},
			"interface_description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interface description").String,
				Optional:            true,
			},
			"vrf_description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("VRF description").String,
				Optional:            true,
			},
			"mtu": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interface MTU").AddIntegerRangeDescription(68, 9216).AddDefaultValueDescription("9216").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(68, 9216),
				},
				Default: int64default.StaticInt64(9216),
			},
			"loopback_routing_tag": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Loopback routing tag").AddIntegerRangeDescription(0, 4294967295).AddDefaultValueDescription("12345").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 4294967295),
				},
				Default: int64default.StaticInt64(12345),
			},
			"redistribute_direct_route_map": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Redistribute direct route map").AddDefaultValueDescription("FABRIC-RMAP-REDIST-SUBNET").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("FABRIC-RMAP-REDIST-SUBNET"),
			},
			"max_bgp_paths": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Maximum BGP paths").AddIntegerRangeDescription(1, 64).AddDefaultValueDescription("1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 64),
				},
				Default: int64default.StaticInt64(1),
			},
			"max_ibgp_paths": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Maximum iBGP paths").AddIntegerRangeDescription(1, 64).AddDefaultValueDescription("2").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 64),
				},
				Default: int64default.StaticInt64(2),
			},
			"ipv6_link_local": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enables IPv6 link-local Option under VRF SVI").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"trm": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Tenant Routed Multicast").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"no_rp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("There is no RP as only SSM is used").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"rp_external": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Is RP external to the fabric").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"rp_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv4 address").String,
				Optional:            true,
			},
			"rp_loopback_id": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("RP loopback ID").AddIntegerRangeDescription(0, 1023).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 1023),
				},
			},
			"underlay_multicast_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv4 Multicast Address. Applicable only when TRM is enabled.").String,
				Optional:            true,
			},
			"overlay_multicast_groups": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Overlay multicast groups").String,
				Optional:            true,
			},
			"mvpn_inter_as": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Use the inter-as keyword for the MVPN address family routes to cross the BGP autonomous system (AS) boundaries, applicable when TRM is enabled. IOS XE Specific").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"trm_bgw_msite": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable TRM on Border Gateway Multisite").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"advertise_host_routes": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Flag to Control Advertisement of /32 and /128 Routes to Edge Routers").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"advertise_default_route": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Flag to Control Advertisement of Default Route Internally").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"configure_static_default_route": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Flag to Control Static Default Route Configuration").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"bgp_password": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("VRF Lite BGP neighbor password (Hex String)").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^[a-fA-F0-9]+$`), ""),
				},
			},
			"bgp_password_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("VRF Lite BGP Key Encryption Type: 3 - 3DES, 7 - Cisco").AddStringEnumDescription("3", "7").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("3", "7"),
				},
			},
			"netflow": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("For netflow on VRF-LITE Sub-interface. Supported only if netflow is enabled on fabric. For NX-OS only").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"netflow_monitor": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Netflow monitor. For NX-OS only").String,
				Optional:            true,
			},
			"disable_rt_auto": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Applicable to IPv4, IPv6 VPN/EVPN/MVPN").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"route_target_import": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("For VPN Routes Import, One or a Comma Separated List").String,
				Optional:            true,
			},
			"route_target_export": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("For VPN Routes Export, One or a Comma Separated List").String,
				Optional:            true,
			},
			"route_target_import_evpn": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("For EVPN Routes Import, One or a Comma Separated List").String,
				Optional:            true,
			},
			"route_target_export_evpn": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("For EVPN Routes Export, One or a Comma Separated List").String,
				Optional:            true,
			},
			"route_target_import_mvpn": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("For MVPN Routes Import, One or a Comma Separated List").String,
				Optional:            true,
			},
			"route_target_export_mvpn": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("For MVPN Routes Export, One or a Comma Separated List").String,
				Optional:            true,
			},
			"route_target_import_cloud_evpn": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("For Cloud EVPN Routes Import, One or a Comma Separated List").String,
				Optional:            true,
			},
			"route_target_export_cloud_evpn": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("For Cloud EVPN Routes Export, One or a Comma Separated List").String,
				Optional:            true,
			},
			"attachments": schema.SetNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("A list of attachments").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"serial_number": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Serial number of switch to attach").String,
							Required:            true,
						},
						"vlan_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Override VLAN ID. `-1` to use VLAN ID defined at VRF level").AddIntegerRangeDescription(-1, 4092).AddDefaultValueDescription("-1").String,
							Optional:            true,
							Computed:            true,
							Validators: []validator.Int64{
								int64validator.Between(-1, 4092),
							},
							Default: int64default.StaticInt64(-1),
						},
						"freeform_config": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("This field covers any configuration not included in overlay templates which is needed as part of this VRF attachment").String,
							Optional:            true,
						},
						"loopback_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Override loopback ID").AddIntegerRangeDescription(0, 1023).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 1023),
							},
						},
						"loopback_ipv4": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Override loopback IPv4 address").String,
							Optional:            true,
						},
						"loopback_ipv6": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Override loopback IPv6 address").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *VRFResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*NdfcProviderData).Client
	r.updateMutex = req.ProviderData.(*NdfcProviderData).UpdateMutex
}

//template:end model

func (r *VRFResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan VRF

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// create vrf
	body := plan.toBody(ctx)
	r.updateMutex.Lock()
	res, err := r.client.Post(plan.getPath(), body)
	r.updateMutex.Unlock()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(plan.FabricName.ValueString() + "/" + plan.VrfName.ValueString())

	if len(plan.Attachments) > 0 {
		// attach
		res, err = r.client.Get(fmt.Sprintf("%vattachments?vrf-names=%v", plan.getPath(), plan.VrfName.ValueString()))
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve VRF attachments, got error: %s, %s", err, res.String()))
			return
		}
		bodyAttachments := plan.toBodyAttachments(ctx, res)
		r.updateMutex.Lock()
		res, err = r.client.Post(plan.getPath()+"attachments", bodyAttachments)
		r.updateMutex.Unlock()
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure VRF attachments, got error: %s, %s", err, res.String()))
			return
		}
		diags = helpers.CheckAttachmentResponse(ctx, res)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		// deploy
		diags = r.Deploy(ctx, plan, "DEPLOYED")
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	res, err = r.client.Get(fmt.Sprintf("%v%v", plan.getPath(), plan.VrfName.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}
	plan.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *VRFResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state VRF

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	res, err := r.client.Get(fmt.Sprintf("%v%v", state.getPath(), state.VrfName.ValueString()))
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

	res, err = r.client.Get(fmt.Sprintf("%vattachments?vrf-names=%v", state.getPath(), state.VrfName.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve VRF attachments, got error: %s, %s", err, res.String()))
		return
	}
	state.fromBodyAttachments(ctx, res, false)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *VRFResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state VRF

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Read state
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	plan.VrfId = state.VrfId
	body := plan.toBody(ctx)
	r.updateMutex.Lock()
	res, err := r.client.Put(fmt.Sprintf("%v%v", plan.getPath(), plan.VrfName.ValueString()), body)
	r.updateMutex.Unlock()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	if len(plan.Attachments) > 0 {
		// attach
		res, err = r.client.Get(fmt.Sprintf("%vattachments?vrf-names=%v", plan.getPath(), plan.VrfName.ValueString()))
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve VRF attachments, got error: %s, %s", err, res.String()))
			return
		}
		bodyAttachments := plan.toBodyAttachments(ctx, res)
		r.updateMutex.Lock()
		res, err = r.client.Post(plan.getPath()+"attachments", bodyAttachments)
		r.updateMutex.Unlock()
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure VRF attachments, got error: %s, %s", err, res.String()))
			return
		}
		diags = helpers.CheckAttachmentResponse(ctx, res)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		// deploy
		diags = r.Deploy(ctx, plan, "DEPLOYED")
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	res, err = r.client.Get(fmt.Sprintf("%v%v", plan.getPath(), plan.VrfName.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}
	plan.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *VRFResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state VRF

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	if len(state.Attachments) > 0 {
		// detach everything
		res, err := r.client.Get(fmt.Sprintf("%vattachments?vrf-names=%v", state.getPath(), state.VrfName.ValueString()))
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve VRF attachments, got error: %s, %s", err, res.String()))
			return
		}
		state.Attachments = make([]VRFAttachments, 0)
		bodyAttachments := state.toBodyAttachments(ctx, res)
		r.updateMutex.Lock()
		res, err = r.client.Post(state.getPath()+"attachments", bodyAttachments)
		r.updateMutex.Unlock()
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure VRF attachments, got error: %s, %s", err, res.String()))
			return
		}
		diags = helpers.CheckAttachmentResponse(ctx, res)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		// deploy
		diags = r.Deploy(ctx, state, "NA")
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		// if there is an ongoing deploy, wait for it to finish
		diags = r.WaitForStatus(ctx, state, "NA")
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	// delete vrf
	res, err := r.client.Delete(fmt.Sprintf("%v%v", state.getPath(), state.VrfName.ValueString()), "")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *VRFResource) Deploy(ctx context.Context, state VRF, expectedStatus string) diag.Diagnostics {
	var diags diag.Diagnostics
	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Deploy", state.Id.ValueString()))

	body := ""
	body, _ = sjson.Set(body, "vrfNames", state.VrfName.ValueString())
	r.updateMutex.Lock()
	defer r.updateMutex.Unlock()
	res, err := r.client.Post(state.getPath()+"deployments", body)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Failed to deploy VRF, got error: %s, %s", err, res.String()))
		return diags
	}

	d := r.WaitForStatus(ctx, state, expectedStatus)
	diags.Append(d...)
	if diags.HasError() {
		return diags
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Deploy finished successfully", state.Id.ValueString()))

	return diags
}

func (r *VRFResource) WaitForStatus(ctx context.Context, state VRF, expectedStatus string) diag.Diagnostics {
	var diags diag.Diagnostics
	status := ""
	for i := 0; i < (helpers.FABRIC_DEPLOY_TIMEOUT / 5); i++ {
		res, err := r.client.Get(state.getPath())
		if err != nil {
			diags.AddError("Client Error", fmt.Sprintf("Failed to retrieve VRFs, got error: %s, %s", err, res.String()))
			return diags
		}
		status = res.Get(`#(vrfName="` + state.VrfName.ValueString() + `").vrfStatus`).String()

		if status == expectedStatus {
			break
		}
		time.Sleep(5 * time.Second)
	}
	if status != expectedStatus {
		diags.AddError("Client Error", fmt.Sprintf("VRF deployment timed out, got status: %s", status))
		return diags
	}
	return diags
}

//template:begin import
func (r *VRFResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	idParts := strings.Split(req.ID, ":")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: '<fabric_name>:<vrf_name>'. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("fabric_name"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("vrf_name"), idParts[1])...)
}

//template:end import
