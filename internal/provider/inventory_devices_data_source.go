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
	"strings"
	"terraform-provider-ndfc/internal/provider/datasources/datasource_inventory_devices"
	"terraform-provider-ndfc/internal/provider/ndfc"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ datasource.DataSource = &InventoryDevicesDataSource{}

func NewInventoryDevicesDataSource() datasource.DataSource {
	return &InventoryDevicesDataSource{}
}

type InventoryDevicesDataSource struct {
	client *ndfc.NDFC
}

type InventoryDevicesDataSourceModel struct {
	Id         types.String `tfsdk:"id"`
	FabricName types.String `tfsdk:"fabric_name"`
	Devices    types.Set    `tfsdk:"devices"`
}

type DeviceDataSourceModel struct {
	IpAddress       types.String `tfsdk:"ip_address"`
	Role            types.String `tfsdk:"role"`
	SerialNumber    types.String `tfsdk:"serial_number"`
	Model           types.String `tfsdk:"model"`
	Version         types.String `tfsdk:"version"`
	Hostname        types.String `tfsdk:"hostname"`
	Uuid            types.String `tfsdk:"uuid"`
	SwitchDbId      types.String `tfsdk:"switch_db_id"`
	DeviceIndex     types.String `tfsdk:"device_index"`
	VdcId           types.String `tfsdk:"vdc_id"`
	VdcMac          types.String `tfsdk:"vdc_mac"`
	Mode            types.String `tfsdk:"mode"`
	ConfigStatus    types.String `tfsdk:"config_status"`
	OperStatus      types.String `tfsdk:"oper_status"`
	DiscoveryStatus types.String `tfsdk:"discovery_status"`
	Managable       types.Bool   `tfsdk:"managable"`
}

func (d *InventoryDevicesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_inventory_devices"
}

func (d *InventoryDevicesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datasource_inventory_devices.InventoryDevicesDataSourceSchema(ctx)
}

func (d *InventoryDevicesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *InventoryDevicesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data InventoryDevicesDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	result := d.client.GetFabricInventory(ctx, &resp.Diagnostics, data.FabricName.ValueString())
	if resp.Diagnostics.HasError() {
		return
	}

	devices := []DeviceDataSourceModel{}

	for _, device := range result.Array() {

		deviceUpdate := DeviceDataSourceModel{
			IpAddress:       basetypes.NewStringValue(device.Get("ipAddress").String()),
			Hostname:        basetypes.NewStringValue(device.Get("logicalName").String()),
			Uuid:            basetypes.NewStringValue(device.Get("swUUID").String()),
			DeviceIndex:     basetypes.NewStringValue(fmt.Sprintf("%s(%s)", device.Get("logicalName"), device.Get("serialNumber"))),
			Model:           basetypes.NewStringValue(device.Get("model").String()),
			Version:         basetypes.NewStringValue(device.Get("release").String()),
			SerialNumber:    basetypes.NewStringValue(device.Get("serialNumber").String()),
			VdcId:           basetypes.NewStringValue(device.Get("vdcId").String()),
			VdcMac:          basetypes.NewStringValue(device.Get("vdcMac").String()),
			Role:            basetypes.NewStringValue(strings.ReplaceAll(device.Get("switchRole").String(), " ", "_")),
			SwitchDbId:      basetypes.NewStringValue(device.Get("switchDbID").String()),
			Mode:            basetypes.NewStringValue(device.Get("mode").String()),
			ConfigStatus:    basetypes.NewStringValue(device.Get("ccStatus").String()),
			OperStatus:      basetypes.NewStringValue(device.Get("operStatus").String()),
			DiscoveryStatus: basetypes.NewStringValue(device.Get("status").String()),
			Managable:       basetypes.NewBoolValue(device.Get("managable").Bool()),
		}

		devices = append(devices, deviceUpdate)
	}

	if len(devices) > 0 {
		devicesSet, _ := types.SetValueFrom(ctx, data.Devices.ElementType(ctx), devices)
		data.Devices = devicesSet
	}

	data.Id = data.FabricName

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
