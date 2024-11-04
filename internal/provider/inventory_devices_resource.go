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
	"net/url"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"terraform-provider-ndfc/internal/provider/ndfc"
	"terraform-provider-ndfc/internal/provider/resources/resource_inventory_devices"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
)

var _ resource.Resource = &InventoryDevicesResource{}
var _ resource.ResourceWithImportState = &InventoryDevicesResource{}

var loggingInventoryDevices = "Inventory Devices Resource"

var retries = int64(300)
var retryWaitTimeout = int64(5)

var defaultAuthProtocol = "md5"
var defaultSetAndUseDiscoveryCredForLan = false
var defaultPreserveConfig = false
var defaultSave = false
var defaultDeploy = false
var defaultDiscoveryType = "discover"

// var defaultRole = "leaf"

func NewInventoryDevicesResource() resource.Resource {
	return &InventoryDevicesResource{}
}

type InventoryDevicesResource struct {
	client *ndfc.NDFC
}

type InventoryDevicesModel struct {
	Id                           types.String `tfsdk:"id"`
	FabricName                   types.String `tfsdk:"fabric_name"`
	SeedIp                       types.String `tfsdk:"seed_ip"`
	AuthProtocol                 types.String `tfsdk:"auth_protocol"`
	UserName                     types.String `tfsdk:"username"`
	Password                     types.String `tfsdk:"password"`
	MaxHops                      types.Int64  `tfsdk:"max_hops"`
	PreserveConfig               types.Bool   `tfsdk:"preserve_config"`
	SetAndUseDiscoveryCredForLan types.Bool   `tfsdk:"set_as_individual_device_write_credential"`
	Save                         types.Bool   `tfsdk:"save"`
	Deploy                       types.Bool   `tfsdk:"deploy"`
	Retries                      types.Int64  `tfsdk:"retries"`
	RetryWaitTimeout             types.Int64  `tfsdk:"retry_wait_timeout"`
	Devices                      types.Map    `tfsdk:"devices"`
}

type DevicesValue struct {
	Role                  types.String `tfsdk:"role"`
	DiscoveryType         types.String `tfsdk:"discovery_type"`
	DiscoveryUsername     types.String `tfsdk:"discovery_username"`
	DiscoveryPassword     types.String `tfsdk:"discovery_password"`
	DiscoveryAuthProtocol types.String `tfsdk:"discovery_auth_protocol"`
	SerialNumber          types.String `tfsdk:"serial_number"`
	Model                 types.String `tfsdk:"model"`
	Version               types.String `tfsdk:"version"`
	ImagePolicy           types.String `tfsdk:"image_policy"`
	ModulesModel          types.Set    `tfsdk:"modules_model"`
	Breakout              types.String `tfsdk:"breakout"`
	PortMode              types.String `tfsdk:"port_mode"`
	Hostname              types.String `tfsdk:"hostname"`
	Gateway               types.String `tfsdk:"gateway"`
	Uuid                  types.String `tfsdk:"uuid"`
	SwitchDbId            types.String `tfsdk:"switch_db_id"`
	DeviceIndex           types.String `tfsdk:"device_index"`
	VdcId                 types.String `tfsdk:"vdc_id"`
	VdcMac                types.String `tfsdk:"vdc_mac"`
	Mode                  types.String `tfsdk:"mode"`
	ConfigStatus          types.String `tfsdk:"config_status"`
	OperStatus            types.String `tfsdk:"oper_status"`
	DiscoveryStatus       types.String `tfsdk:"discovery_status"`
	Managable             types.Bool   `tfsdk:"managable"`
}

var snmpAuthenticationProtocol = map[string]int{"md5": 0, "sha": 1, "md5_des": 2, "md5_aes": 3, "sha_des": 4, "sha_aes": 5}

func (r *InventoryDevicesResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	tflog.Debug(ctx, fmt.Sprintf("Start of %s Metadata", loggingInventoryDevices))
	resp.TypeName = req.ProviderTypeName + "_inventory_devices"
	tflog.Debug(ctx, fmt.Sprintf("End of %s Metadata", loggingInventoryDevices))
}

func (r InventoryDevicesResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	tflog.Debug(ctx, fmt.Sprintf("Start of %s ValidateConfig", loggingInventoryDevices))
	var data InventoryDevicesModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var devices map[string]DevicesValue
	data.Devices.ElementsAs(ctx, &devices, false)

	for ipAddress, device := range devices {
		if device.DiscoveryType.ValueString() == "discover" || device.DiscoveryType.ValueString() == "" {
			if !device.SerialNumber.IsUnknown() && !device.SerialNumber.IsNull() {
				resp.Diagnostics.AddError("Schema Validation Error", fmt.Sprintf("device with ip address '%s' has set discovery_type 'discover' which does not allow providing a serial number", ipAddress))
			}
			if !device.Model.IsUnknown() && !device.Model.IsNull() {
				resp.Diagnostics.AddError("Schema Validation Error", fmt.Sprintf("device with ip address '%s' has set discovery_type 'discover' which does not allow providing a model", ipAddress))
			}
			if !device.Version.IsUnknown() && !device.Version.IsNull() {
				resp.Diagnostics.AddError("Schema Validation Error", fmt.Sprintf("device with ip address '%s' has set discovery_type 'discover' which does not allow providing a version", ipAddress))
			}
			if !device.Hostname.IsUnknown() && !device.Hostname.IsNull() {
				resp.Diagnostics.AddError("Schema Validation Error", fmt.Sprintf("device with ip address '%s' has set discovery_type 'discover' which does not allow providing a hostname", ipAddress))
			}
		} else if device.DiscoveryType.ValueString() == "pre_provision" {
			if device.SerialNumber.IsUnknown() || device.SerialNumber.IsNull() {
				resp.Diagnostics.AddError("Schema Validation Error", fmt.Sprintf("device with ip address '%s' has set discovery_type 'pre_provision' which requires a serial number", ipAddress))
			}
			if device.Model.IsUnknown() || device.Model.IsNull() {
				resp.Diagnostics.AddError("Schema Validation Error", fmt.Sprintf("device with ip address '%s' has set discovery_type 'pre_provision' which requires a model", ipAddress))
			}
			if device.Version.IsUnknown() || device.Version.IsNull() {
				resp.Diagnostics.AddError("Schema Validation Error", fmt.Sprintf("device with ip address '%s' has set discovery_type 'pre_provision' which requires a version", ipAddress))
			}
			if device.Hostname.IsUnknown() || device.Hostname.IsNull() {
				resp.Diagnostics.AddError("Schema Validation Error", fmt.Sprintf("device with ip address '%s' has set discovery_type 'pre_provision' which requires a hostname", ipAddress))
			}
			if device.Gateway.IsUnknown() || device.Gateway.IsNull() {
				resp.Diagnostics.AddError("Schema Validation Error", fmt.Sprintf("device with ip address '%s' has set discovery_type 'pre_provision' which requires a gateway", ipAddress))
			}
			if device.ModulesModel.IsUnknown() || device.ModulesModel.IsNull() {
				resp.Diagnostics.AddError("Schema Validation Error", fmt.Sprintf("device with ip address '%s' has set discovery_type 'pre_provision' which requires a modules_model", ipAddress))
			}
		} else if device.DiscoveryType.ValueString() == "bootstrap" {
			if device.SerialNumber.IsUnknown() || device.SerialNumber.IsNull() {
				resp.Diagnostics.AddError("Schema Validation Error", fmt.Sprintf("device with ip address '%s' has set discovery_type 'bootstrap' which requires a serial number", ipAddress))
			}
			if device.Hostname.IsUnknown() || device.Hostname.IsNull() {
				resp.Diagnostics.AddError("Schema Validation Error", fmt.Sprintf("device with ip address '%s' has set discovery_type 'bootstrap' which requires a hostname", ipAddress))
			}
			if !device.Model.IsUnknown() && !device.Model.IsNull() {
				resp.Diagnostics.AddError("Schema Validation Error", fmt.Sprintf("device with ip address '%s' has set discovery_type 'bootstrap' which does not allow providing a model", ipAddress))
			}
			if !device.Version.IsUnknown() && !device.Version.IsNull() {
				resp.Diagnostics.AddError("Schema Validation Error", fmt.Sprintf("device with ip address '%s' has set discovery_type 'bootstrap' which does not allow providing a version", ipAddress))
			}
		} else if device.DiscoveryType.ValueString() == "rma" {
			if device.SerialNumber.IsUnknown() || device.SerialNumber.IsNull() {
				resp.Diagnostics.AddError("Schema Validation Error", fmt.Sprintf("device with ip address '%s' has set discovery_type 'rma' which requires a serial number", ipAddress))
			}
			if device.Hostname.IsUnknown() || device.Hostname.IsNull() {
				resp.Diagnostics.AddError("Schema Validation Error", fmt.Sprintf("device with ip address '%s' has set discovery_type 'rma' which requires a hostname", ipAddress))
			}
			if !device.Model.IsUnknown() && !device.Model.IsNull() {
				resp.Diagnostics.AddError("Schema Validation Error", fmt.Sprintf("device with ip address '%s' has set discovery_type 'rma' which does not allow providing a model", ipAddress))
			}
			if !device.Role.IsUnknown() && !device.Role.IsNull() {
				resp.Diagnostics.AddError("Schema Validation Error", fmt.Sprintf("device with ip address '%s' has set discovery_type 'rma' which does not allow providing a role", ipAddress))
			}
			if !device.Version.IsUnknown() && !device.Version.IsNull() {
				resp.Diagnostics.AddError("Schema Validation Error", fmt.Sprintf("device with ip address '%s' has set discovery_type 'rma' which does not allow providing a version", ipAddress))
			}
		}
	}
	tflog.Debug(ctx, fmt.Sprintf("End of %s ValidateConfig", loggingInventoryDevices))
}

func (r *InventoryDevicesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	tflog.Debug(ctx, fmt.Sprintf("Start of %s Schema", loggingInventoryDevices))
	resp.Schema = resource_inventory_devices.InventoryDevicesResourceSchema(ctx)
	tflog.Debug(ctx, fmt.Sprintf("End of %s Schema", loggingInventoryDevices))
}

func (r *InventoryDevicesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Debug(ctx, fmt.Sprintf("Start of %s Configure", loggingInventoryDevices))
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
	tflog.Debug(ctx, fmt.Sprintf("End of %s Configure", loggingInventoryDevices))
}

func (r *InventoryDevicesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, fmt.Sprintf("Start of %s Create", loggingInventoryDevices))
	var planData, stateData InventoryDevicesModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &planData)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &stateData)...)
	if resp.Diagnostics.HasError() {
		return
	}

	planData.Id = basetypes.NewStringValue(planData.FabricName.ValueString())

	getInventoryData(ctx, r.client, &resp.Diagnostics, &stateData)
	if resp.Diagnostics.HasError() {
		return
	}

	var stateDevices map[string]DevicesValue
	stateData.Devices.ElementsAs(ctx, &stateDevices, false)
	if resp.Diagnostics.HasError() {
		return
	} else if len(stateDevices) > 0 {
		resp.Diagnostics.AddError("Create Validation Error", "Devices already exist in the inventory, please import the state or delete the existing devices")
		return
	}

	addDevicesToInventory(ctx, r.client, &resp.Diagnostics, &planData, &stateData)
	if resp.Diagnostics.HasError() {
		return
	}

	deployAndSave(ctx, r.client, &resp.Diagnostics, &planData)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &planData)...)
	tflog.Debug(ctx, fmt.Sprintf("End of %s Create", loggingInventoryDevices))
}

func (r *InventoryDevicesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, fmt.Sprintf("Start of %s Read", loggingInventoryDevices))
	var stateData InventoryDevicesModel

	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if stateData.FabricName.IsNull() || stateData.FabricName.IsUnknown() {
		stateData.FabricName = stateData.Id
	}
	if stateData.Retries.IsNull() || stateData.Retries.IsUnknown() {
		stateData.Retries = basetypes.NewInt64Value(retries)
	}
	if stateData.RetryWaitTimeout.IsNull() || stateData.RetryWaitTimeout.IsUnknown() {
		stateData.RetryWaitTimeout = basetypes.NewInt64Value(retryWaitTimeout)
	}

	getInventoryData(ctx, r.client, &resp.Diagnostics, &stateData)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &stateData)...)
	tflog.Debug(ctx, fmt.Sprintf("End of %s Read", loggingInventoryDevices))
}

func (r *InventoryDevicesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, fmt.Sprintf("Start of %s Update", loggingInventoryDevices))
	var createData, planData, stateData InventoryDevicesModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &createData)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &planData)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

	if resp.Diagnostics.HasError() {
		return
	}

	planData.Id = basetypes.NewStringValue(planData.FabricName.ValueString())

	var planDevices, stateDevices map[string]DevicesValue
	planData.Devices.ElementsAs(ctx, &planDevices, false)
	stateData.Devices.ElementsAs(ctx, &stateDevices, false)
	if resp.Diagnostics.HasError() {
		return
	}

	uuidList := []string{}
	for ipAddress, stateDevice := range stateDevices {
		if _, ok := planDevices[ipAddress]; !ok {
			uuidList = append(uuidList, stateDevice.Uuid.ValueString())
		}
	}
	if len(uuidList) > 0 {
		r.client.DeleteFabricInventoryDevices(ctx, &resp.Diagnostics, planData.FabricName.ValueString(), uuidList)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	createDevices, updateDevices := map[string]DevicesValue{}, map[string]DevicesValue{}
	for ipAddress, planDevice := range planDevices {
		if _, ok := stateDevices[ipAddress]; ok {
			if planDevice.DiscoveryType.ValueString() == "rma" {
				createDevices[ipAddress] = planDevice
			} else {
				updateDevices[ipAddress] = planDevice
			}
		} else {
			createDevices[ipAddress] = planDevice
		}
	}

	if len(createDevices) > 0 {
		devicesMap, _ := types.MapValueFrom(ctx, createData.Devices.ElementType(ctx), createDevices)
		createData.Devices = devicesMap
		if resp.Diagnostics.HasError() {
			return
		}

		addDevicesToInventory(ctx, r.client, &resp.Diagnostics, &createData, &stateData)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	if len(updateDevices) > 0 {
		for ipAddress, updateDevice := range updateDevices {
			stateDevice := stateDevices[ipAddress]

			if stateDevice.Role.ValueString() != updateDevice.Role.ValueString() {
				r.client.UpdateRole(ctx, &resp.Diagnostics, stateDevice.SwitchDbId.ValueString(), updateDevice.Role.ValueString())
				if resp.Diagnostics.HasError() {
					return
				}
			}

			if stateDevice.SerialNumber.ValueString() != updateDevice.SerialNumber.ValueString() && !updateDevice.SerialNumber.IsUnknown() && !updateDevice.SerialNumber.IsNull() {
				if updateDevice.DiscoveryType.ValueString() != "pre_provision" || stateDevice.DiscoveryType.ValueString() == "pre_provision" {
					resp.Diagnostics.AddError("Provider Error", fmt.Sprintf("Serial Number updates are only allowed for pre_provision devices, got: %s", updateDevice.DiscoveryType.ValueString()))
					return
				}
				r.client.UpdateSerialNumber(ctx, &resp.Diagnostics, planData.FabricName.ValueString(), stateDevice.SerialNumber.ValueString(), updateDevice.SerialNumber.ValueString())
				if resp.Diagnostics.HasError() {
					return
				}
			}
		}

	}

	deployAndSave(ctx, r.client, &resp.Diagnostics, &planData)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &planData)...)
	tflog.Debug(ctx, fmt.Sprintf("End of %s Update", loggingInventoryDevices))
}

func (r *InventoryDevicesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, fmt.Sprintf("Start of %s Delete", loggingInventoryDevices))
	var data InventoryDevicesModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	devices := map[string]DevicesValue{}
	data.Devices.ElementsAs(ctx, &devices, false)
	if resp.Diagnostics.HasError() {
		return
	}

	uuidList := []string{}
	for _, device := range devices {
		uuidList = append(uuidList, device.Uuid.ValueString())
	}

	sort.Strings(uuidList)

	if len(uuidList) > 0 {
		r.client.DeleteFabricInventoryDevices(ctx, &resp.Diagnostics, data.FabricName.ValueString(), uuidList)
		if resp.Diagnostics.HasError() {
			return
		}
	}
	tflog.Debug(ctx, fmt.Sprintf("End of %s Delete", loggingInventoryDevices))
}

func (r *InventoryDevicesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Debug(ctx, fmt.Sprintf("Start of %s ImportState", loggingInventoryDevices))
	auth_protocol := os.Getenv("NDFC_INVENTORY_AUTH_PROTOCOL")
	if auth_protocol == "" {
		auth_protocol = defaultAuthProtocol
	} else if _, ok := snmpAuthenticationProtocol[auth_protocol]; !ok {
		resp.Diagnostics.AddError("Invalid input", fmt.Sprintf("Invalid value for auth_protocol: %s, please set the NDFC_INVENTORY_AUTH_PROTOCOL environment variable to one of %v", auth_protocol, reflect.ValueOf(snmpAuthenticationProtocol).MapKeys()))
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("auth_protocol"), auth_protocol)...)

	username := os.Getenv("NDFC_INVENTORY_USERNAME")
	if username == "" {
		resp.Diagnostics.AddError("Missing input", "A username must be provided during import, please set the NDFC_INVENTORY_USERNAME environment variable")
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("username"), username)...)

	password := os.Getenv("NDFC_INVENTORY_PASSWORD")
	if password == "" {
		resp.Diagnostics.AddError("Missing input", "A password must be provided during import, please set the NDFC_INVENTORY_PASSWORD environment variable")
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("password"), password)...)

	setAndUseDiscoveryCredForLan := getBoolValueFromEnv(ctx, &resp.Diagnostics, "NDFC_INVENTORY_SET_AS_INDIVIDUAL_DEVICE_WRITE_CREDENTIAL", defaultSetAndUseDiscoveryCredForLan)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("set_as_individual_device_write_credential"), setAndUseDiscoveryCredForLan)...)

	preserveConfig := getBoolValueFromEnv(ctx, &resp.Diagnostics, "NDFC_INVENTORY_PRESERVE_CONFIG", defaultPreserveConfig)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("preserve_config"), preserveConfig)...)

	save := getBoolValueFromEnv(ctx, &resp.Diagnostics, "NDFC_INVENTORY_SAVE", defaultSave)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("save"), save)...)

	deploy := getBoolValueFromEnv(ctx, &resp.Diagnostics, "NDFC_INVENTORY_DEPLOY", defaultDeploy)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("deploy"), deploy)...)

	data := InventoryDevicesModel{}
	data.FabricName = basetypes.NewStringValue(req.ID)
	data.Devices = types.MapNull(resource_inventory_devices.DevicesValue{}.Type(ctx))
	getInventoryData(ctx, r.client, &resp.Diagnostics, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("devices"), data.Devices)...)
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	tflog.Debug(ctx, fmt.Sprintf("End of %s ImportState", loggingInventoryDevices))
}

func getBoolValueFromEnv(ctx context.Context, diags *diag.Diagnostics, envKey string, defaultValue bool) bool {
	envValue := os.Getenv(envKey)
	if envValue == "" {
		return defaultValue
	} else {
		value, err := strconv.ParseBool(envValue)
		if err != nil {
			diags.AddError(
				fmt.Sprintf("Invalid input '%s'", envValue),
				fmt.Sprintf("A boolean value must be provided for the %s environment variable", envKey),
			)
		} else {
			return value
		}
	}
	return defaultValue
}

func getInventoryData(ctx context.Context, client *ndfc.NDFC, diags *diag.Diagnostics, data *InventoryDevicesModel) {

	tflog.Debug(ctx, fmt.Sprintf("Start of %s getInventoryData", loggingInventoryDevices))
	result := client.GetFabricInventory(ctx, diags, data.FabricName.ValueString())
	if diags.HasError() {
		return
	}
	devicesValues := map[string]DevicesValue{}
	devices := map[string]DevicesValue{}
	*diags = data.Devices.ElementsAs(ctx, &devices, false)
	if diags.HasError() {
		return
	}

	for _, device := range result.Array() {

		/*
			WORKAROUND: handle different types used by endpoints in API.
			   - ../test-reachability returns vdcMac as a string
			   - ../switchesByFabric returns vdcMac as a null when string is empty
		*/
		deviceUpdate := DevicesValue{}
		if device.Get("vdcMac").Value() != nil {
			vdcMac := device.Get("vdcMac").String()
			deviceUpdate.VdcMac = basetypes.NewStringValue(vdcMac)
		} else {
			deviceUpdate.VdcMac = basetypes.NewStringNull()
		}
		deviceUpdate = DevicesValue{
			Hostname:        basetypes.NewStringValue(device.Get("logicalName").String()),
			Uuid:            basetypes.NewStringValue(device.Get("swUUID").String()),
			DeviceIndex:     basetypes.NewStringValue(fmt.Sprintf("%s(%s)", device.Get("logicalName"), device.Get("serialNumber"))),
			Model:           basetypes.NewStringValue(device.Get("model").String()),
			Version:         basetypes.NewStringValue(device.Get("release").String()),
			SerialNumber:    basetypes.NewStringValue(device.Get("serialNumber").String()),
			VdcId:           basetypes.NewStringValue(device.Get("vdcId").String()),
			Role:            basetypes.NewStringValue(strings.ReplaceAll(device.Get("switchRole").String(), " ", "_")),
			SwitchDbId:      basetypes.NewStringValue(device.Get("switchDbID").String()),
			Mode:            basetypes.NewStringValue(device.Get("mode").String()),
			ConfigStatus:    basetypes.NewStringValue(device.Get("ccStatus").String()),
			OperStatus:      basetypes.NewStringValue(device.Get("operStatus").String()),
			DiscoveryStatus: basetypes.NewStringValue(device.Get("status").String()),
			Managable:       basetypes.NewBoolValue(device.Get("managable").Bool()),
		}

		if stateDevice, ok := devices[device.Get("ipAddress").String()]; ok {
			if !stateDevice.DiscoveryType.IsNull() && !stateDevice.DiscoveryType.IsUnknown() {
				deviceUpdate.DiscoveryType = stateDevice.DiscoveryType
			} else {
				deviceUpdate.DiscoveryType = basetypes.NewStringValue(defaultDiscoveryType)
			}
			if !stateDevice.DiscoveryAuthProtocol.IsNull() && !stateDevice.DiscoveryAuthProtocol.IsUnknown() {
				deviceUpdate.DiscoveryAuthProtocol = stateDevice.DiscoveryAuthProtocol
			} else {
				deviceUpdate.DiscoveryAuthProtocol = basetypes.NewStringNull()
			}
			if !stateDevice.DiscoveryUsername.IsNull() && !stateDevice.DiscoveryUsername.IsUnknown() {
				deviceUpdate.DiscoveryUsername = stateDevice.DiscoveryUsername
			} else {
				deviceUpdate.DiscoveryUsername = basetypes.NewStringNull()
			}
			if !stateDevice.DiscoveryPassword.IsNull() && !stateDevice.DiscoveryPassword.IsUnknown() {
				deviceUpdate.DiscoveryPassword = stateDevice.DiscoveryPassword
			} else {
				deviceUpdate.DiscoveryPassword = basetypes.NewStringNull()
			}
			if !stateDevice.ImagePolicy.IsNull() && !stateDevice.ImagePolicy.IsUnknown() {
				deviceUpdate.ImagePolicy = stateDevice.ImagePolicy
			} else {
				deviceUpdate.ImagePolicy = basetypes.NewStringNull()
			}
			if !stateDevice.Gateway.IsNull() && !stateDevice.Gateway.IsUnknown() {
				deviceUpdate.Gateway = stateDevice.Gateway
			} else {
				deviceUpdate.Gateway = basetypes.NewStringNull()
			}
			if !stateDevice.ModulesModel.IsNull() && !stateDevice.ModulesModel.IsUnknown() {
				deviceUpdate.ModulesModel = stateDevice.ModulesModel
			} else {
				deviceUpdate.ModulesModel = basetypes.NewSetNull(basetypes.StringType{})
			}
			if !stateDevice.Breakout.IsNull() && !stateDevice.Breakout.IsUnknown() {
				deviceUpdate.Breakout = stateDevice.Breakout
			} else {
				deviceUpdate.Breakout = basetypes.NewStringNull()
			}
			if !stateDevice.PortMode.IsNull() && !stateDevice.PortMode.IsUnknown() {
				deviceUpdate.PortMode = stateDevice.PortMode
			} else {
				deviceUpdate.PortMode = basetypes.NewStringNull()
			}
		} else {
			deviceUpdate.DiscoveryType = basetypes.NewStringValue(defaultDiscoveryType)
			deviceUpdate.DiscoveryAuthProtocol = basetypes.NewStringNull()
			deviceUpdate.DiscoveryUsername = basetypes.NewStringNull()
			deviceUpdate.DiscoveryPassword = basetypes.NewStringNull()
			deviceUpdate.ImagePolicy = basetypes.NewStringNull()
			deviceUpdate.Gateway = basetypes.NewStringNull()
			deviceUpdate.ModulesModel = basetypes.NewSetNull(basetypes.StringType{})
			deviceUpdate.Breakout = basetypes.NewStringNull()
			deviceUpdate.PortMode = basetypes.NewStringNull()
		}
		devicesValues[device.Get("ipAddress").String()] = deviceUpdate
	}

	if len(devicesValues) > 0 {
		devicesMap, diags := types.MapValueFrom(ctx, data.Devices.ElementType(ctx), devicesValues)
		if diags.HasError() {
			return
		}
		data.Devices = devicesMap
	} else {
		data.Devices = types.MapNull(data.Devices.ElementType(ctx))
	}
	tflog.Debug(ctx, fmt.Sprintf("End of %s getInventoryData", loggingInventoryDevices))
}

func addDevicesToInventory(ctx context.Context, client *ndfc.NDFC, diags *diag.Diagnostics, data, stateData *InventoryDevicesModel) {
	tflog.Debug(ctx, fmt.Sprintf("Start of %s addDevicesToInventory", loggingInventoryDevices))
	discoverDevices := map[string]DevicesValue{}
	poapDevices := map[string]DevicesValue{}
	rmaDevices := map[string]DevicesValue{}

	var devices, stateDevices map[string]DevicesValue
	data.Devices.ElementsAs(ctx, &devices, false)
	stateData.Devices.ElementsAs(ctx, &stateDevices, false)
	if diags.HasError() {
		return
	}

	for ipAddress, device := range devices {
		switch device.DiscoveryType.ValueString() {
		case "discover":
			discoverDevices[ipAddress] = device
		case "bootstrap":
			poapDevices[ipAddress] = device
		case "pre_provision":
			poapDevices[ipAddress] = device
		case "rma":
			if _, ok := stateDevices[ipAddress]; !ok {
				diags.AddError("Configuration Error", fmt.Sprintf("Device with ip address '%s' has set discovery_type 'rma' which is not supported for new devices", ipAddress))
			} else if stateDevices[ipAddress].Mode.ValueString() != "Maintenance" {
				diags.AddError("Configuration Error", fmt.Sprintf("Device with ip address '%s' has set mode to '%s' instead of 'Maintenance'", ipAddress, stateDevices[ipAddress].Mode.ValueString()))
			} else if stateDevices[ipAddress].SerialNumber.ValueString() != device.SerialNumber.ValueString() {
				rmaDevices[ipAddress] = device
			}
		}
	}

	if len(discoverDevices) > 0 {
		discover(ctx, client, diags, data, discoverDevices)
		if diags.HasError() {
			return
		}
	}

	if len(poapDevices) > 0 {
		poap(ctx, client, diags, data, poapDevices)
		if diags.HasError() {
			return
		}
	}

	if len(rmaDevices) > 0 {
		rma(ctx, client, diags, data, rmaDevices, stateDevices)
		if diags.HasError() {
			return
		}
		// Validate that the serial number is updated in NDFC else terraform will error with state inconsistency
		validateFabric(ctx, client, diags, data, "serial")
		if diags.HasError() {
			return
		}
	}

	isLanCredentialsSet := client.IsLanCredentialSet(ctx, diags)
	if diags.HasError() {
		return
	}

	if !isLanCredentialsSet.Bool() {

		lanCredentialsSet := client.GetLanCredentialSet(ctx, diags)
		if diags.HasError() {
			return
		}

		for _, lanCredential := range lanCredentialsSet.Array() {

			if device, ok := devices[lanCredential.Get("ipAddress").String()]; ok && lanCredential.Get("credType").String() == "" {

				username := data.UserName.ValueString()
				if !device.DiscoveryUsername.IsNull() && !device.DiscoveryUsername.IsUnknown() {
					username = device.DiscoveryUsername.ValueString()
				}

				password := data.Password.ValueString()
				if !device.DiscoveryPassword.IsNull() && !device.DiscoveryPassword.IsUnknown() {
					password = device.DiscoveryPassword.ValueString()
				}

				payload := url.Values{}
				payload.Add("switchIds", lanCredential.Get("switchDbID").String())
				payload.Add("username", username)
				payload.Add("password", password)

				client.SetLanCredentialSet(ctx, diags, payload.Encode())

			}
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("End of %s addDevicesToInventory", loggingInventoryDevices))
}

func rma(ctx context.Context, client *ndfc.NDFC, diags *diag.Diagnostics, data *InventoryDevicesModel, devices, stateDevices map[string]DevicesValue) {
	tflog.Debug(ctx, fmt.Sprintf("Start of %s rma", loggingInventoryDevices))
	bootstrap := map[string]map[string]interface{}{}

	fabricInventoryPoap := client.GetFabricInventoryPoap(ctx, diags, data.FabricName.ValueString())
	if diags.HasError() {
		return
	}

	for _, device := range fabricInventoryPoap.Array() {
		bootstrap[device.Get("serialNumber").String()] = map[string]interface{}{
			"data":      device.Get("data").String(),
			"model":     device.Get("model").String(),
			"publicKey": device.Get("publicKey").String(),
			"version":   device.Get("version").String(),
		}
	}

	for ipAddress, device := range devices {

		devicePayload := map[string]interface{}{
			"oldSerialNumber": stateDevices[ipAddress].SerialNumber.ValueString(),
			"newSerialNumber": device.SerialNumber.ValueString(),
			"hostname":        device.Hostname.ValueString(),
			"ipAddress":       ipAddress,
			"password":        data.Password.ValueString(),
			"imagePolicy":     device.ImagePolicy.ValueString(),
			"data":            bootstrap[device.SerialNumber.ValueString()]["data"],
			"model":           bootstrap[device.SerialNumber.ValueString()]["model"],
			"publicKey":       bootstrap[device.SerialNumber.ValueString()]["publicKey"],
			"version":         bootstrap[device.SerialNumber.ValueString()]["version"],
		}

		if !device.DiscoveryUsername.IsNull() && !device.DiscoveryUsername.IsUnknown() {
			devicePayload["discoveryUsername"] = device.DiscoveryUsername.ValueString()
		}
		if !device.DiscoveryPassword.IsNull() && !device.DiscoveryPassword.IsUnknown() {
			devicePayload["discoveryPassword"] = device.DiscoveryPassword.ValueString()
		}
		if !device.DiscoveryAuthProtocol.IsNull() && !device.DiscoveryAuthProtocol.IsUnknown() {
			devicePayload["discoveryAuthProtocol"] = snmpAuthenticationProtocol[device.DiscoveryAuthProtocol.ValueString()]
		} else {
			devicePayload["discoveryAuthProtocol"] = snmpAuthenticationProtocol[data.AuthProtocol.ValueString()]
		}

		jsonPayload, err := json.Marshal(devicePayload)
		if err != nil {
			diags.AddError("Provider Error", fmt.Sprintf("Unable to create json payload, got error: %s", err))
			return
		}

		client.SetFabricInventoryRma(ctx, diags, data.FabricName.ValueString(), jsonPayload)

	}
	tflog.Debug(ctx, fmt.Sprintf("End of %s rma", loggingInventoryDevices))
}

func poap(ctx context.Context, client *ndfc.NDFC, diags *diag.Diagnostics, data *InventoryDevicesModel, devices map[string]DevicesValue) {
	tflog.Debug(ctx, fmt.Sprintf("Start of %s poap", loggingInventoryDevices))
	var payload []map[string]interface{}
	bootstrap := map[string]map[string]interface{}{}

	fabricInventoryPoap := client.GetFabricInventoryPoap(ctx, diags, data.FabricName.ValueString())
	if diags.HasError() {
		return
	}

	for _, device := range fabricInventoryPoap.Array() {
		bootstrap[device.Get("serialNumber").String()] = map[string]interface{}{
			"data":         device.Get("data").String(),
			"fingerprint":  device.Get("fingerprint").String(),
			"model":        device.Get("model").String(),
			"publicKey":    device.Get("publicKey").String(),
			"reAdd":        device.Get("reAdd").Bool(),
			"serialNumber": device.Get("serialNumber").String(),
			"version":      device.Get("version").String(),
		}
	}

	for ipAddress, device := range devices {

		var devicePayload map[string]interface{}

		if device.DiscoveryType.ValueString() == "pre_provision" {

			modulesModel := []string{}
			device.ModulesModel.ElementsAs(ctx, &modulesModel, false)

			dataPayload := map[string]interface{}{
				"modulesModel": modulesModel,
				"gateway":      device.Gateway.ValueString(),
				"breakout":     device.Breakout.ValueString(),
				"portMode":     device.PortMode.ValueString(),
			}

			jsonDataPayload, err := json.Marshal(dataPayload)
			if err != nil {
				diags.AddError("Provider Error", fmt.Sprintf("Unable to create json payload, got error: %s", err))
				return
			}

			devicePayload = map[string]interface{}{
				"serialNumber": device.SerialNumber.ValueString(),
				"model":        device.Model.ValueString(),
				"version":      device.Version.ValueString(),
				"hostname":     device.Hostname.ValueString(),
				"ipAddress":    ipAddress,
				"imagePolicy":  device.ImagePolicy.ValueString(),
				"role":         device.Role.ValueString(),
				"data":         string(jsonDataPayload),
				"password":     data.Password.ValueString(),
			}

			if !device.DiscoveryUsername.IsNull() && !device.DiscoveryUsername.IsUnknown() {
				devicePayload["discoveryUsername"] = device.DiscoveryUsername.ValueString()
			}
			if !device.DiscoveryPassword.IsNull() && !device.DiscoveryPassword.IsUnknown() {
				devicePayload["discoveryPassword"] = device.DiscoveryPassword.ValueString()
			}

		} else if device.DiscoveryType.ValueString() == "bootstrap" {

			if bootstrapDevice, ok := bootstrap[device.SerialNumber.ValueString()]; ok {
				devicePayload = map[string]interface{}{
					"data":         bootstrapDevice["data"].(string),
					"fingerprint":  bootstrapDevice["fingerprint"].(string),
					"hostname":     device.Hostname.ValueString(),
					"imagePolicy":  device.ImagePolicy.ValueString(),
					"ipAddress":    ipAddress,
					"model":        bootstrapDevice["model"].(string),
					"password":     data.Password.ValueString(),
					"publicKey":    bootstrapDevice["publicKey"].(string),
					"reAdd":        bootstrapDevice["reAdd"].(bool),
					"role":         device.Role.ValueString(),
					"serialNumber": device.SerialNumber.ValueString(),
					"version":      bootstrapDevice["version"].(string),
				}
			} else {
				diags.AddError("Provider Error", fmt.Sprintf("Device with serial number '%s' not found in bootstrap devices", device.SerialNumber.ValueString()))
				return
			}
		}

		if !device.DiscoveryAuthProtocol.IsNull() && !device.DiscoveryAuthProtocol.IsUnknown() {
			devicePayload["discoveryAuthProtocol"] = snmpAuthenticationProtocol[device.DiscoveryAuthProtocol.ValueString()]
		} else {
			devicePayload["discoveryAuthProtocol"] = snmpAuthenticationProtocol[data.AuthProtocol.ValueString()]
		}

		payload = append(payload, devicePayload)

	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		diags.AddError("Provider Error", fmt.Sprintf("Unable to create json payload, got error: %s", err))
		return
	}

	client.SetFabricInventoryPoap(ctx, diags, data.FabricName.ValueString(), jsonPayload)
	tflog.Debug(ctx, fmt.Sprintf("End of %s poap", loggingInventoryDevices))
}

func discover(ctx context.Context, client *ndfc.NDFC, diags *diag.Diagnostics, data *InventoryDevicesModel, devices map[string]DevicesValue) {
	tflog.Debug(ctx, fmt.Sprintf("Start of %s discover", loggingInventoryDevices))
	reachableDevices := map[string]gjson.Result{}
	var payload map[string]interface{}
	if diags.HasError() {
		return
	}

	if !data.MaxHops.IsNull() && !data.MaxHops.IsUnknown() && !data.SeedIp.IsNull() && !data.SeedIp.IsUnknown() {
		payload = map[string]interface{}{
			"cdpSecondTimeout":    5,
			"discoveryCredForLan": data.SetAndUseDiscoveryCredForLan.ValueBool(),
			"password":            data.Password.ValueString(),
			"maxHops":             strconv.FormatInt(data.MaxHops.ValueInt64(), 10),
			"preserveConfig":      data.PreserveConfig.ValueBool(),
			"snmpV3AuthProtocol":  snmpAuthenticationProtocol[data.AuthProtocol.ValueString()],
			"username":            data.UserName.ValueString(),
			"seedIP":              data.SeedIp.ValueString(),
		}

		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			diags.AddError("Provider Error", fmt.Sprintf("Unable to create json payload, got error: %s", err))
			return
		}

		reachables := client.TestReachability(ctx, diags, data.FabricName.ValueString(), jsonPayload)
		if diags.HasError() {
			return
		}

		for _, reachableDevice := range reachables.Array() {
			reachableDevices[reachableDevice.Get("ipaddr").String()] = reachableDevice
		}

	} else {
		payload = map[string]interface{}{
			"cdpSecondTimeout":    5,
			"discoveryCredForLan": data.SetAndUseDiscoveryCredForLan.ValueBool(),
			"password":            data.Password.ValueString(),
			"preserveConfig":      data.PreserveConfig.ValueBool(),
			"snmpV3AuthProtocol":  snmpAuthenticationProtocol[data.AuthProtocol.ValueString()],
			"username":            data.UserName.ValueString(),
		}
	}

	payload["maxHops"] = 0
	ipAdresses := []string{}
	switches := []map[string]interface{}{}
	DevicesValues := map[string]DevicesValue{}
	for ipAddress, device := range devices {
		ipAdresses = append(ipAdresses, ipAddress)
		if _, ok := reachableDevices[ipAddress]; !ok {
			payload["seedIP"] = ipAddress

			jsonPayload, err := json.Marshal(payload)
			if err != nil {
				diags.AddError("Provider Error", fmt.Sprintf("Unable to create json payload, got error: %s", err))
				return
			}

			reachable := client.TestReachability(ctx, diags, data.FabricName.ValueString(), jsonPayload)
			if diags.HasError() {
				return
			}

			if len(reachable.Array()) != 1 {
				diags.AddError("Too many results in return", fmt.Sprintf("Expected 1 device, got %d: %v", len(reachable.Array()), reachable))
				return
			}
			reachableDevices[ipAddress] = reachable.Array()[0]
		}
		reachableDevice := reachableDevices[ipAddress]

		if reachableDevice.Get("selectable").Bool() && reachableDevice.Get("reachable").Bool() {
			switchPayload := map[string]interface{}{
				"ipaddr":       reachableDevice.Get("ipaddr").String(),
				"sysName":      reachableDevice.Get("sysName").String(),
				"deviceIndex":  reachableDevice.Get("deviceIndex").String(),
				"platform":     reachableDevice.Get("platform").String(),
				"version":      reachableDevice.Get("version").String(),
				"serialNumber": reachableDevice.Get("serialNumber").String(),
				"vdcId":        reachableDevice.Get("vdcId").String(),
				"vdcMac":       reachableDevice.Get("vdcMac").String(),
			}
			switches = append(switches, switchPayload)
		} else {
			diags.AddError("Device not reachable", fmt.Sprintf("Device with ip address '%s' is not reachable", ipAddress))
			return
		}

		discoveredDevice := DevicesValue{
			Hostname:      basetypes.NewStringValue(reachableDevice.Get("sysName").String()),
			DeviceIndex:   basetypes.NewStringValue(reachableDevice.Get("deviceIndex").String()),
			Model:         basetypes.NewStringValue(reachableDevice.Get("platform").String()),
			Version:       basetypes.NewStringValue(reachableDevice.Get("version").String()),
			SerialNumber:  basetypes.NewStringValue(reachableDevice.Get("serialNumber").String()),
			VdcId:         basetypes.NewStringValue(reachableDevice.Get("vdcId").String()),
			VdcMac:        basetypes.NewStringValue(reachableDevice.Get("vdcMac").String()),
			Role:          basetypes.NewStringValue(device.Role.ValueString()),
			DiscoveryType: basetypes.NewStringValue(device.DiscoveryType.ValueString()),
		}

		if !device.DiscoveryUsername.IsNull() && !device.DiscoveryUsername.IsUnknown() {
			discoveredDevice.DiscoveryUsername = device.DiscoveryUsername
		}
		if !device.DiscoveryAuthProtocol.IsNull() && !device.DiscoveryAuthProtocol.IsUnknown() {
			discoveredDevice.DiscoveryAuthProtocol = device.DiscoveryAuthProtocol
		}
		if !device.DiscoveryPassword.IsNull() && !device.DiscoveryPassword.IsUnknown() {
			discoveredDevice.DiscoveryPassword = device.DiscoveryPassword
		}
		if !device.ImagePolicy.IsNull() && !device.ImagePolicy.IsUnknown() {
			discoveredDevice.ImagePolicy = device.ImagePolicy
		}
		if !device.Gateway.IsNull() && !device.Gateway.IsUnknown() {
			discoveredDevice.Gateway = device.Gateway
		}
		if !device.ModulesModel.IsNull() && !device.ModulesModel.IsUnknown() {
			discoveredDevice.ModulesModel = device.ModulesModel
		}
		if !device.Breakout.IsNull() && !device.Breakout.IsUnknown() {
			discoveredDevice.Breakout = device.Breakout
		}
		if !device.PortMode.IsNull() && !device.PortMode.IsUnknown() {
			discoveredDevice.PortMode = device.PortMode
		}

		DevicesValues[ipAddress] = device
	}

	payload["switches"] = switches
	payload["seedIP"] = strings.Join(ipAdresses, ",")

	jsonSwitchPayload, err := json.Marshal(payload)
	if err != nil {
		diags.AddError("Provider Error", fmt.Sprintf("Unable to create json payload, got error: %s", err))
		return
	}

	client.SetFabricInventoryDiscover(
		ctx, diags, data.FabricName.ValueString(), data.SetAndUseDiscoveryCredForLan.ValueBool(), jsonSwitchPayload,
	)
	if diags.HasError() {
		return
	}

	fabricInventory := client.GetFabricInventory(ctx, diags, data.FabricName.ValueString())
	if diags.HasError() {
		return
	}

	for _, fabricDevice := range fabricInventory.Array() {
		if device, ok := DevicesValues[fabricDevice.Get("ipAddress").String()]; ok {
			client.UpdateRole(ctx, diags, fabricDevice.Get("switchDbID").String(), strings.ReplaceAll(device.Role.ValueString(), "_", "%20"))
			if diags.HasError() {
				return
			}
		}
	}
	tflog.Debug(ctx, fmt.Sprintf("End of %s discover", loggingInventoryDevices))
}

func deployAndSave(ctx context.Context, client *ndfc.NDFC, diags *diag.Diagnostics, data *InventoryDevicesModel) {
	tflog.Debug(ctx, fmt.Sprintf("Start of %s deployAndSave", loggingInventoryDevices))
	var devices map[string]DevicesValue
	data.Devices.ElementsAs(ctx, &devices, false)
	if diags.HasError() {
		return
	}

	validateFabric(ctx, client, diags, data, "managed")
	if diags.HasError() {
		return
	}

	if data.Save.ValueBool() || data.Deploy.ValueBool() && len(devices) > 0 {

		rediscoverDevices(ctx, client, diags, data, []string{})
		if diags.HasError() {
			return
		}

		if data.Save.ValueBool() {
			save(ctx, client, diags, data)
			if diags.HasError() {
				return
			}
		}

		if data.Deploy.ValueBool() {
			deploy(ctx, client, diags, data)
			if diags.HasError() {
				return
			}
		}
	}
	tflog.Debug(ctx, fmt.Sprintf("End of %s deployAndSave", loggingInventoryDevices))
}

func rediscoverDevices(ctx context.Context, client *ndfc.NDFC, diags *diag.Diagnostics, data *InventoryDevicesModel, switchDbIDs []string) {
	tflog.Debug(ctx, fmt.Sprintf("Start of %s rediscoverDevices", loggingInventoryDevices))
	if len(switchDbIDs) == 0 {
		var devices map[string]DevicesValue
		data.Devices.ElementsAs(ctx, &devices, false)
		if diags.HasError() {
			return
		}
		for _, device := range devices {
			switchDbIDs = append(switchDbIDs, device.SwitchDbId.ValueString())
		}
	}

	payload, err := json.Marshal(switchDbIDs)
	if err != nil {
		diags.AddError("Provider Error", fmt.Sprintf("Unable to create json payload, got error: %s", err))
		return
	}

	client.SetFabricInventoryRediscover(ctx, diags, payload)
	if diags.HasError() {
		return
	}
	tflog.Debug(ctx, fmt.Sprintf("End of %s rediscoverDevices", loggingInventoryDevices))
}

func save(ctx context.Context, client *ndfc.NDFC, diags *diag.Diagnostics, data *InventoryDevicesModel) {
	tflog.Debug(ctx, fmt.Sprintf("Start of %s save", loggingInventoryDevices))
	validateFabric(ctx, client, diags, data, "discovered")
	if diags.HasError() {
		return
	}

	if data.Save.ValueBool() {
		client.SaveConfiguration(ctx, diags, data.FabricName.ValueString())
		if diags.HasError() {
			return
		}
		validateFabric(ctx, client, diags, data, "discovered")
	}
	tflog.Debug(ctx, fmt.Sprintf("End of %s save", loggingInventoryDevices))
}

func deploy(ctx context.Context, client *ndfc.NDFC, diags *diag.Diagnostics, data *InventoryDevicesModel) {
	tflog.Debug(ctx, fmt.Sprintf("Start of %s deploy", loggingInventoryDevices))
	validateFabric(ctx, client, diags, data, "discovered")
	if diags.HasError() {
		return
	}

	if data.Deploy.ValueBool() {
		var devices map[string]DevicesValue
		data.Devices.ElementsAs(ctx, &devices, false)
		if diags.HasError() {
			return
		}

		serialNumbers := []string{}
		for _, device := range devices {
			if device.ConfigStatus.ValueString() != "In-Sync" && device.DiscoveryType.ValueString() != "pre_provision" && device.Mode.ValueString() == "Normal" {
				serialNumbers = append(serialNumbers, device.SerialNumber.ValueString())
			}
		}

		sort.Strings(serialNumbers)

		client.DeployConfiguration(ctx, diags, data.FabricName.ValueString(), serialNumbers)
		if diags.HasError() {
			return
		}

		validateFabric(ctx, client, diags, data, "configured")
	}
	tflog.Debug(ctx, fmt.Sprintf("End of %s deploy", loggingInventoryDevices))
}

func validateFabric(ctx context.Context, client *ndfc.NDFC, diags *diag.Diagnostics, data *InventoryDevicesModel, state string) {
	tflog.Debug(ctx, fmt.Sprintf("Start of %s validateFabric", loggingInventoryDevices))
	// Hardcoded delay to prevent validation before devices are ready
	time.Sleep(30 * time.Second)

	configDevices := map[string]DevicesValue{}
	data.Devices.ElementsAs(ctx, &configDevices, false)

	attempt := int64(1)
	var allNormal bool
	for attempt <= data.Retries.ValueInt64() {
		getInventoryData(ctx, client, diags, data)
		if diags.HasError() {
			return
		}

		allNormal = true
		devices := map[string]DevicesValue{}
		data.Devices.ElementsAs(ctx, &devices, false)

		for ipAddress, device := range devices {

			mode := device.Mode.ValueString()
			manageable := device.Managable.ValueBool()
			configStatus := device.ConfigStatus.ValueString()
			discoverType := device.DiscoveryType.ValueString()
			discoveryStatus := device.DiscoveryStatus.ValueString()

			if state == "managed" && !manageable && discoverType != "pre_provision" {
				allNormal = false
				break
			} else if state == "discovered" && (!manageable || discoveryStatus != "ok" || (mode != "Normal" && mode != "Maintenance")) && discoverType != "pre_provision" {
				if manageable {
					rediscoverDevices(ctx, client, diags, data, []string{device.SwitchDbId.ValueString()})
				}
				allNormal = false
				break
			} else if state == "configured" && !manageable && configStatus != "In-Sync" && discoverType != "pre_provision" {
				allNormal = false
				break
			} else if state == "serial" {
				if configDevice, ok := configDevices[ipAddress]; ok && configDevice.SerialNumber.ValueString() != device.SerialNumber.ValueString() {
					allNormal = false
					break
				}
			}
		}

		if !allNormal && attempt == data.Retries.ValueInt64() {
			diags.AddError(
				"Fabric Mode Failed",
				fmt.Sprintf(
					"Not all devices in expected mode after %d attempts of %d seconds. Please check the fabric manually, and/or retry with a higher retry_wait_timeout and more retries.",
					data.Retries.ValueInt64(),
					data.RetryWaitTimeout.ValueInt64(),
				),
			)
		} else if allNormal {
			break
		}
		time.Sleep(time.Duration(data.RetryWaitTimeout.ValueInt64()) * time.Second)
		attempt++
	}
	tflog.Debug(ctx, fmt.Sprintf("End of %s validateFabric", loggingInventoryDevices))
}
