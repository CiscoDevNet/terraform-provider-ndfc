// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package resource_inventory_devices

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCInventoryDevicesModel struct {
	FabricName                           string                      `json:"fabricName,omitempty"`
	AuthProtocol                         string                      `json:"snmpV3AuthProtocol,string,omitempty"`
	Username                             string                      `json:"username,string,omitempty"`
	Password                             string                      `json:"password,string,omitempty"`
	SeedIp                               string                      `json:"seedIP,omitempty"`
	MaxHops                              *int64                      `json:"maxHops,omitempty"`
	SetAsIndividualDeviceWriteCredential *bool                       `json:"discoveryCredForLan,omitempty"`
	PreserveConfig                       *bool                       `json:"preserveConfig,omitempty"`
	Save                                 *bool                       `json:"save,omitempty"`
	Deploy                               *bool                       `json:"deploy,omitempty"`
	Retries                              *int64                      `json:"retries,omitempty"`
	RetryWaitTimeout                     *int64                      `json:"retryWaitTimeout,omitempty"`
	Devices                              map[string]NDFCDevicesValue `json:"devices,omitempty"`
}

type NDFCDevicesValue struct {
	Role                  string   `json:"role,string,omitempty"`
	DiscoveryType         string   `json:"discovery_type,string,omitempty"`
	DiscoveryUsername     string   `json:"discovery_username,string,omitempty"`
	DiscoveryPassword     string   `json:"discovery_password,string,omitempty"`
	DiscoveryAuthProtocol string   `json:"snmpV3AuthProtocol,string,omitempty"`
	SerialNumber          string   `json:"serial_number,string,omitempty"`
	Model                 string   `json:"model,string,omitempty"`
	Version               string   `json:"version,string,omitempty"`
	Hostname              string   `json:"hostname,string,omitempty"`
	ImagePolicy           string   `json:"image_policy,string,omitempty"`
	Gateway               string   `json:"gateway,string,omitempty"`
	ModulesModel          []string `json:"modules_model,omitempty"`
	Breakout              string   `json:"breakout,string,omitempty"`
	PortMode              string   `json:"port_mode,string,omitempty"`
	Uuid                  string   `json:"uuid,string,omitempty"`
	SwitchDbId            string   `json:"switch_db_id,string,omitempty"`
	DeviceIndex           string   `json:"device_index,string,omitempty"`
	VdcId                 string   `json:"vdc_id,string,omitempty"`
	VdcMac                string   `json:"vdc_mac,string,omitempty"`
	Mode                  string   `json:"mode,string,omitempty"`
	ConfigStatus          string   `json:"config_status,string,omitempty"`
	OperStatus            string   `json:"oper_status,string,omitempty"`
	DiscoveryStatus       string   `json:"discovery_status,string,omitempty"`
	Manageable            *bool    `json:"managable,omitempty"`
}

func (v *InventoryDevicesModel) SetModelData(jsonData *NDFCInventoryDevicesModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.FabricName != "" {
		v.FabricName = types.StringValue(jsonData.FabricName)
	} else {
		v.FabricName = types.StringNull()
	}

	if jsonData.AuthProtocol != "" {
		v.AuthProtocol = types.StringValue(jsonData.AuthProtocol)
	} else {
		v.AuthProtocol = types.StringNull()
	}

	if jsonData.Username != "" {
		v.Username = types.StringValue(jsonData.Username)
	} else {
		v.Username = types.StringNull()
	}

	if jsonData.Password != "" {
		v.Password = types.StringValue(jsonData.Password)
	} else {
		v.Password = types.StringNull()
	}

	if jsonData.SeedIp != "" {
		v.SeedIp = types.StringValue(jsonData.SeedIp)
	} else {
		v.SeedIp = types.StringNull()
	}

	if jsonData.MaxHops != nil {
		v.MaxHops = types.Int64Value(*jsonData.MaxHops)

	} else {
		v.MaxHops = types.Int64Null()
	}

	if jsonData.SetAsIndividualDeviceWriteCredential != nil {
		v.SetAsIndividualDeviceWriteCredential = types.BoolValue(*jsonData.SetAsIndividualDeviceWriteCredential)

	} else {
		v.SetAsIndividualDeviceWriteCredential = types.BoolNull()
	}

	if jsonData.PreserveConfig != nil {
		v.PreserveConfig = types.BoolValue(*jsonData.PreserveConfig)

	} else {
		v.PreserveConfig = types.BoolNull()
	}

	if jsonData.Save != nil {
		v.Save = types.BoolValue(*jsonData.Save)

	} else {
		v.Save = types.BoolNull()
	}

	if jsonData.Deploy != nil {
		v.Deploy = types.BoolValue(*jsonData.Deploy)

	} else {
		v.Deploy = types.BoolNull()
	}

	if jsonData.Retries != nil {
		v.Retries = types.Int64Value(*jsonData.Retries)

	} else {
		v.Retries = types.Int64Null()
	}

	if jsonData.RetryWaitTimeout != nil {
		v.RetryWaitTimeout = types.Int64Value(*jsonData.RetryWaitTimeout)

	} else {
		v.RetryWaitTimeout = types.Int64Null()
	}

	if len(jsonData.Devices) == 0 {
		log.Printf("v.Devices is empty")
		v.Devices = types.MapNull(DevicesValue{}.Type(context.Background()))
	} else {
		mapData := make(map[string]DevicesValue)
		for key, item := range jsonData.Devices {
			data := new(DevicesValue)
			err = data.SetValue(&item)
			if err != nil {
				log.Printf("Error in DevicesValue.SetValue")
				return err
			}
			data.state = attr.ValueStateKnown
			mapData[key] = *data
		}
		v.Devices, err = types.MapValueFrom(context.Background(), DevicesValue{}.Type(context.Background()), mapData)
		if err != nil {
			log.Printf("Error in converting map[string]DevicesValue to  Map")

		}
	}

	return err
}

func (v *DevicesValue) SetValue(jsonData *NDFCDevicesValue) diag.Diagnostics {

	var err diag.Diagnostics
	err = nil

	if jsonData.Role != "" {
		v.Role = types.StringValue(jsonData.Role)
	} else {
		v.Role = types.StringNull()
	}

	if jsonData.DiscoveryType != "" {
		v.DiscoveryType = types.StringValue(jsonData.DiscoveryType)
	} else {
		v.DiscoveryType = types.StringNull()
	}

	if jsonData.DiscoveryUsername != "" {
		v.DiscoveryUsername = types.StringValue(jsonData.DiscoveryUsername)
	} else {
		v.DiscoveryUsername = types.StringNull()
	}

	if jsonData.DiscoveryPassword != "" {
		v.DiscoveryPassword = types.StringValue(jsonData.DiscoveryPassword)
	} else {
		v.DiscoveryPassword = types.StringNull()
	}

	if jsonData.DiscoveryAuthProtocol != "" {
		v.DiscoveryAuthProtocol = types.StringValue(jsonData.DiscoveryAuthProtocol)
	} else {
		v.DiscoveryAuthProtocol = types.StringNull()
	}

	if jsonData.SerialNumber != "" {
		v.SerialNumber = types.StringValue(jsonData.SerialNumber)
	} else {
		v.SerialNumber = types.StringNull()
	}

	if jsonData.Model != "" {
		v.Model = types.StringValue(jsonData.Model)
	} else {
		v.Model = types.StringNull()
	}

	if jsonData.Version != "" {
		v.Version = types.StringValue(jsonData.Version)
	} else {
		v.Version = types.StringNull()
	}

	if jsonData.Hostname != "" {
		v.Hostname = types.StringValue(jsonData.Hostname)
	} else {
		v.Hostname = types.StringNull()
	}

	if jsonData.ImagePolicy != "" {
		v.ImagePolicy = types.StringValue(jsonData.ImagePolicy)
	} else {
		v.ImagePolicy = types.StringNull()
	}

	if jsonData.Gateway != "" {
		v.Gateway = types.StringValue(jsonData.Gateway)
	} else {
		v.Gateway = types.StringNull()
	}

	if len(jsonData.ModulesModel) == 0 {
		log.Printf("v.ModulesModel is empty")
		v.ModulesModel, err = types.SetValue(types.StringType, []attr.Value{})
		if err != nil {
			log.Printf("Error in converting []string to  List %v", err)
			return err
		}
	} else {
		listData := make([]attr.Value, len(jsonData.ModulesModel))
		for i, item := range jsonData.ModulesModel {
			listData[i] = types.StringValue(item)
		}
		v.ModulesModel, err = types.SetValue(types.StringType, listData)
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}
	if jsonData.Breakout != "" {
		v.Breakout = types.StringValue(jsonData.Breakout)
	} else {
		v.Breakout = types.StringNull()
	}

	if jsonData.PortMode != "" {
		v.PortMode = types.StringValue(jsonData.PortMode)
	} else {
		v.PortMode = types.StringNull()
	}

	if jsonData.Uuid != "" {
		v.Uuid = types.StringValue(jsonData.Uuid)
	} else {
		v.Uuid = types.StringNull()
	}

	if jsonData.SwitchDbId != "" {
		v.SwitchDbId = types.StringValue(jsonData.SwitchDbId)
	} else {
		v.SwitchDbId = types.StringNull()
	}

	if jsonData.DeviceIndex != "" {
		v.DeviceIndex = types.StringValue(jsonData.DeviceIndex)
	} else {
		v.DeviceIndex = types.StringNull()
	}

	if jsonData.VdcId != "" {
		v.VdcId = types.StringValue(jsonData.VdcId)
	} else {
		v.VdcId = types.StringNull()
	}

	if jsonData.VdcMac != "" {
		v.VdcMac = types.StringValue(jsonData.VdcMac)
	} else {
		v.VdcMac = types.StringNull()
	}

	if jsonData.Mode != "" {
		v.Mode = types.StringValue(jsonData.Mode)
	} else {
		v.Mode = types.StringNull()
	}

	if jsonData.ConfigStatus != "" {
		v.ConfigStatus = types.StringValue(jsonData.ConfigStatus)
	} else {
		v.ConfigStatus = types.StringNull()
	}

	if jsonData.OperStatus != "" {
		v.OperStatus = types.StringValue(jsonData.OperStatus)
	} else {
		v.OperStatus = types.StringNull()
	}

	if jsonData.DiscoveryStatus != "" {
		v.DiscoveryStatus = types.StringValue(jsonData.DiscoveryStatus)
	} else {
		v.DiscoveryStatus = types.StringNull()
	}

	if jsonData.Manageable != nil {
		v.Manageable = types.BoolValue(*jsonData.Manageable)

	} else {
		v.Manageable = types.BoolNull()
	}

	return err
}

func (v InventoryDevicesModel) GetModelData() *NDFCInventoryDevicesModel {
	var data = new(NDFCInventoryDevicesModel)

	//MARSHAL_BODY

	if !v.FabricName.IsNull() && !v.FabricName.IsUnknown() {
		data.FabricName = v.FabricName.ValueString()
	} else {
		data.FabricName = ""
	}

	if !v.AuthProtocol.IsNull() && !v.AuthProtocol.IsUnknown() {
		data.AuthProtocol = v.AuthProtocol.ValueString()
	} else {
		data.AuthProtocol = ""
	}

	if !v.Username.IsNull() && !v.Username.IsUnknown() {
		data.Username = v.Username.ValueString()
	} else {
		data.Username = ""
	}

	if !v.Password.IsNull() && !v.Password.IsUnknown() {
		data.Password = v.Password.ValueString()
	} else {
		data.Password = ""
	}

	if !v.SeedIp.IsNull() && !v.SeedIp.IsUnknown() {
		data.SeedIp = v.SeedIp.ValueString()
	} else {
		data.SeedIp = ""
	}

	if !v.MaxHops.IsNull() && !v.MaxHops.IsUnknown() {
		data.MaxHops = new(int64)
		*data.MaxHops = v.MaxHops.ValueInt64()

	} else {
		data.MaxHops = nil
	}

	if !v.SetAsIndividualDeviceWriteCredential.IsNull() && !v.SetAsIndividualDeviceWriteCredential.IsUnknown() {
		data.SetAsIndividualDeviceWriteCredential = new(bool)
		*data.SetAsIndividualDeviceWriteCredential = v.SetAsIndividualDeviceWriteCredential.ValueBool()
	} else {
		data.SetAsIndividualDeviceWriteCredential = nil
	}

	if !v.PreserveConfig.IsNull() && !v.PreserveConfig.IsUnknown() {
		data.PreserveConfig = new(bool)
		*data.PreserveConfig = v.PreserveConfig.ValueBool()
	} else {
		data.PreserveConfig = nil
	}

	if !v.Save.IsNull() && !v.Save.IsUnknown() {
		data.Save = new(bool)
		*data.Save = v.Save.ValueBool()
	} else {
		data.Save = nil
	}

	if !v.Deploy.IsNull() && !v.Deploy.IsUnknown() {
		data.Deploy = new(bool)
		*data.Deploy = v.Deploy.ValueBool()
	} else {
		data.Deploy = nil
	}

	if !v.Retries.IsNull() && !v.Retries.IsUnknown() {
		data.Retries = new(int64)
		*data.Retries = v.Retries.ValueInt64()

	} else {
		data.Retries = nil
	}

	if !v.RetryWaitTimeout.IsNull() && !v.RetryWaitTimeout.IsUnknown() {
		data.RetryWaitTimeout = new(int64)
		*data.RetryWaitTimeout = v.RetryWaitTimeout.ValueInt64()

	} else {
		data.RetryWaitTimeout = nil
	}

	if !v.Devices.IsNull() && !v.Devices.IsUnknown() {
		elements1 := make(map[string]DevicesValue, len(v.Devices.Elements()))

		data.Devices = make(map[string]NDFCDevicesValue)

		diag := v.Devices.ElementsAs(context.Background(), &elements1, false)
		if diag != nil {
			panic(diag)
		}
		for k1, ele1 := range elements1 {
			data1 := new(NDFCDevicesValue)

			// role | String| []| false
			if !ele1.Role.IsNull() && !ele1.Role.IsUnknown() {

				data1.Role = ele1.Role.ValueString()
			} else {
				data1.Role = ""
			}

			// discovery_type | String| []| false
			if !ele1.DiscoveryType.IsNull() && !ele1.DiscoveryType.IsUnknown() {

				data1.DiscoveryType = ele1.DiscoveryType.ValueString()
			} else {
				data1.DiscoveryType = ""
			}

			// discovery_username | String| []| false
			if !ele1.DiscoveryUsername.IsNull() && !ele1.DiscoveryUsername.IsUnknown() {

				data1.DiscoveryUsername = ele1.DiscoveryUsername.ValueString()
			} else {
				data1.DiscoveryUsername = ""
			}

			// discovery_password | String| []| false
			if !ele1.DiscoveryPassword.IsNull() && !ele1.DiscoveryPassword.IsUnknown() {

				data1.DiscoveryPassword = ele1.DiscoveryPassword.ValueString()
			} else {
				data1.DiscoveryPassword = ""
			}

			// discovery_auth_protocol | String| []| false
			if !ele1.DiscoveryAuthProtocol.IsNull() && !ele1.DiscoveryAuthProtocol.IsUnknown() {

				data1.DiscoveryAuthProtocol = ele1.DiscoveryAuthProtocol.ValueString()
			} else {
				data1.DiscoveryAuthProtocol = ""
			}

			// serial_number | String| []| false
			if !ele1.SerialNumber.IsNull() && !ele1.SerialNumber.IsUnknown() {

				data1.SerialNumber = ele1.SerialNumber.ValueString()
			} else {
				data1.SerialNumber = ""
			}

			// model | String| []| false
			if !ele1.Model.IsNull() && !ele1.Model.IsUnknown() {

				data1.Model = ele1.Model.ValueString()
			} else {
				data1.Model = ""
			}

			// version | String| []| false
			if !ele1.Version.IsNull() && !ele1.Version.IsUnknown() {

				data1.Version = ele1.Version.ValueString()
			} else {
				data1.Version = ""
			}

			// hostname | String| []| false
			if !ele1.Hostname.IsNull() && !ele1.Hostname.IsUnknown() {

				data1.Hostname = ele1.Hostname.ValueString()
			} else {
				data1.Hostname = ""
			}

			// image_policy | String| []| false
			if !ele1.ImagePolicy.IsNull() && !ele1.ImagePolicy.IsUnknown() {

				data1.ImagePolicy = ele1.ImagePolicy.ValueString()
			} else {
				data1.ImagePolicy = ""
			}

			// gateway | String| []| false
			if !ele1.Gateway.IsNull() && !ele1.Gateway.IsUnknown() {

				data1.Gateway = ele1.Gateway.ValueString()
			} else {
				data1.Gateway = ""
			}

			// modules_model | List:String| []| false
			if !ele1.ModulesModel.IsNull() && !ele1.ModulesModel.IsUnknown() {

				listStringData := make([]string, len(ele1.ModulesModel.Elements()))
				dg := ele1.ModulesModel.ElementsAs(context.Background(), &listStringData, false)
				if dg.HasError() {
					panic(dg.Errors())
				}
				data1.ModulesModel = make([]string, len(listStringData))
				copy(data1.ModulesModel, listStringData)
			}

			// breakout | String| []| false
			if !ele1.Breakout.IsNull() && !ele1.Breakout.IsUnknown() {

				data1.Breakout = ele1.Breakout.ValueString()
			} else {
				data1.Breakout = ""
			}

			// port_mode | String| []| false
			if !ele1.PortMode.IsNull() && !ele1.PortMode.IsUnknown() {

				data1.PortMode = ele1.PortMode.ValueString()
			} else {
				data1.PortMode = ""
			}

			// uuid | String| []| false
			// switch_db_id | String| []| false
			// device_index | String| []| false
			// vdc_id | String| []| false
			// vdc_mac | String| []| false
			// mode | String| []| false
			// config_status | String| []| false
			// oper_status | String| []| false
			// discovery_status | String| []| false
			// manageable | Bool| []| false
			data.Devices[k1] = *data1

		}
	}

	return data
}
