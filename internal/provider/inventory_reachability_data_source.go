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
	"encoding/json"
	"fmt"
	"strconv"
	"terraform-provider-ndfc/internal/provider/datasources/datasource_inventory_reachability"
	"terraform-provider-ndfc/internal/provider/ndfc"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &InventoryReachabilityDataSource{}

func NewInventoryReachabilityDataSource() datasource.DataSource {
	return &InventoryReachabilityDataSource{}
}

type InventoryReachabilityDataSource struct {
	client *ndfc.NDFC
}

type InventoryReachabilityDataSourceModel struct {
	Id                           types.String `tfsdk:"id"`
	FabricName                   types.String `tfsdk:"fabric_name"`
	SeedIp                       types.String `tfsdk:"seed_ip"`
	AuthProtocol                 types.String `tfsdk:"auth_protocol"`
	UserName                     types.String `tfsdk:"username"`
	Password                     types.String `tfsdk:"password"`
	MaxHops                      types.Int64  `tfsdk:"max_hops"`
	PreserveConfig               types.Bool   `tfsdk:"preserve_config"`
	SetAndUseDiscoveryCredForLan types.Bool   `tfsdk:"set_as_individual_device_write_credential"`
	ReachabilityDetails          types.Set    `tfsdk:"reachability_details"`
}

type ReachabilityDetailsModel struct {
	IpAddress     types.String `tfsdk:"ip_address"`
	Hostname      types.String `tfsdk:"hostname"`
	DeviceIndex   types.String `tfsdk:"device_index"`
	Model         types.String `tfsdk:"model"`
	Version       types.String `tfsdk:"version"`
	SerialNumber  types.String `tfsdk:"serial_number"`
	VdcId         types.String `tfsdk:"vdc_id"`
	VdcMac        types.String `tfsdk:"vdc_mac"`
	Reachable     types.Bool   `tfsdk:"reachable"`
	Selectable    types.Bool   `tfsdk:"selectable"`
	Authenticated types.Bool   `tfsdk:"authenticated"`
	Valid         types.Bool   `tfsdk:"valid"`
	Known         types.Bool   `tfsdk:"known"`
	LastChanged   types.String `tfsdk:"last_changed"`
	StatusReason  types.String `tfsdk:"status_reason"`
}

func (d *InventoryReachabilityDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_reachability"
}

func (d *InventoryReachabilityDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datasource_inventory_reachability.InventoryReachabilityDataSourceSchema(ctx)
}

func (d *InventoryReachabilityDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*ndfc.NDFC)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *ndfc.NDFC, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *InventoryReachabilityDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data InventoryReachabilityDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	payload := map[string]interface{}{
		"cdpSecondTimeout":    5,
		"discoveryCredForLan": data.SetAndUseDiscoveryCredForLan.ValueBool(),
		"maxHops":             strconv.FormatInt(data.MaxHops.ValueInt64(), 10),
		"password":            data.Password.ValueString(),
		"preserveConfig":      data.PreserveConfig.ValueBool(),
		"seedIP":              data.SeedIp.ValueString(),
		"snmpV3AuthProtocol":  snmpAuthenticationProtocol[data.AuthProtocol.ValueString()],
		"username":            data.UserName.ValueString(),
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		resp.Diagnostics.AddError("Provider Error", fmt.Sprintf("Unable to create json payload, got error: %s", err))
	}

	reachable := d.client.TestReachability(ctx, &resp.Diagnostics, data.FabricName.ValueString(), jsonPayload)
	if resp.Diagnostics.HasError() {
		return
	}

	reachableDevices := make([]ReachabilityDetailsModel, 0)
	for _, device := range reachable.Array() {
		reachableDevice := ReachabilityDetailsModel{
			IpAddress:     types.StringValue(device.Get("ipaddr").String()),
			Hostname:      types.StringValue(device.Get("sysName").String()),
			DeviceIndex:   types.StringValue(device.Get("deviceIndex").String()),
			Model:         types.StringValue(device.Get("platform").String()),
			Version:       types.StringValue(device.Get("version").String()),
			SerialNumber:  types.StringValue(device.Get("serialNumber").String()),
			VdcId:         types.StringValue(device.Get("vdcId").String()),
			VdcMac:        types.StringValue(device.Get("vdcMac").String()),
			Reachable:     types.BoolValue(device.Get("reachable").Bool()),
			Selectable:    types.BoolValue(device.Get("selectable").Bool()),
			Authenticated: types.BoolValue(device.Get("auth").Bool()),
			Valid:         types.BoolValue(device.Get("valid").Bool()),
			Known:         types.BoolValue(device.Get("known").Bool()),
			LastChanged:   types.StringValue(device.Get("lastChanged").String()),
			StatusReason:  types.StringValue(device.Get("statusReason").String()),
		}
		reachableDevices = append(reachableDevices, reachableDevice)
	}

	if len(reachableDevices) > 0 {
		reachableDevicesSet, _ := types.SetValueFrom(ctx, data.ReachabilityDetails.ElementType(ctx), reachableDevices)
		data.ReachabilityDetails = reachableDevicesSet
	}

	data.Id = types.StringValue(fmt.Sprintf("fabric-%s_seed_ip-%s_max_hops-%s", data.FabricName.ValueString(), data.SeedIp.ValueString(), data.MaxHops))

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
