// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package resource_interface_loopback

import (
	"context"
	"log"
	"strconv"
	"terraform-provider-ndfc/internal/provider/resources/resource_interface_common"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (v *InterfaceLoopbackModel) SetModelData(jsonData *resource_interface_common.NDFCInterfaceCommonModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.Policy != "" {
		v.Policy = types.StringValue(jsonData.Policy)
	} else {
		v.Policy = types.StringNull()
	}

	v.Deploy = types.BoolValue(jsonData.Deploy)
	if jsonData.SerialNumber != "" {
		v.SerialNumber = types.StringValue(jsonData.SerialNumber)
	} else {
		v.SerialNumber = types.StringNull()
	}

	if len(jsonData.Interfaces) == 0 {
		log.Printf("v.Interfaces is empty")
		v.Interfaces = types.MapNull(InterfacesValue{}.Type(context.Background()))
	} else {
		mapData := make(map[string]InterfacesValue)
		for key, item := range jsonData.Interfaces {
			if item.FilterThisValue {
				//Skip this entry - this parameter allows filtering
				continue
			}

			data := new(InterfacesValue)
			err = data.SetValue(&item)
			if err != nil {
				log.Printf("Error in InterfacesValue.SetValue")
				return err
			}
			data.state = attr.ValueStateKnown
			mapData[key] = *data
		}
		v.Interfaces, err = types.MapValueFrom(context.Background(), InterfacesValue{}.Type(context.Background()), mapData)
		if err != nil {
			log.Printf("Error in converting map[string]InterfacesValue to  Map")

		}
	}

	return err
}

func (v *InterfacesValue) SetValue(jsonData *resource_interface_common.NDFCInterfacesValue) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.SerialNumber != "" {
		v.SerialNumber = types.StringValue(jsonData.SerialNumber)
	} else {
		v.SerialNumber = types.StringNull()
	}

	if jsonData.InterfaceName != "" {
		v.InterfaceName = types.StringValue(jsonData.InterfaceName)
	} else {
		v.InterfaceName = types.StringNull()
	}

	if jsonData.NvPairs.AdminState != "" {
		x, _ := strconv.ParseBool(jsonData.NvPairs.AdminState)
		v.AdminState = types.BoolValue(x)
	} else {
		v.AdminState = types.BoolNull()
	}

	if jsonData.NvPairs.FreeformConfig != "" {
		v.FreeformConfig = types.StringValue(jsonData.NvPairs.FreeformConfig)
	} else {
		v.FreeformConfig = types.StringNull()
	}

	if jsonData.NvPairs.InterfaceDescription != "" {
		v.InterfaceDescription = types.StringValue(jsonData.NvPairs.InterfaceDescription)
	} else {
		v.InterfaceDescription = types.StringNull()
	}

	if jsonData.NvPairs.Vrf != "" {
		v.Vrf = types.StringValue(jsonData.NvPairs.Vrf)
	} else {
		v.Vrf = types.StringNull()
	}

	if jsonData.NvPairs.Ipv4Address != "" {
		v.Ipv4Address = types.StringValue(jsonData.NvPairs.Ipv4Address)
	} else {
		v.Ipv4Address = types.StringNull()
	}

	if jsonData.NvPairs.Ipv6Address != "" {
		v.Ipv6Address = types.StringValue(jsonData.NvPairs.Ipv6Address)
	} else {
		v.Ipv6Address = types.StringNull()
	}

	if jsonData.NvPairs.RouteMapTag != "" {
		v.RouteMapTag = types.StringValue(jsonData.NvPairs.RouteMapTag)
	} else {
		v.RouteMapTag = types.StringNull()
	}

	if jsonData.DeploymentStatus != "" {
		v.DeploymentStatus = types.StringValue(jsonData.DeploymentStatus)
	} else {
		v.DeploymentStatus = types.StringNull()
	}

	return err
}

func (v InterfaceLoopbackModel) GetModelData() *resource_interface_common.NDFCInterfaceCommonModel {
	var data = new(resource_interface_common.NDFCInterfaceCommonModel)

	//MARSHAL_BODY

	if !v.Policy.IsNull() && !v.Policy.IsUnknown() {
		data.Policy = v.Policy.ValueString()
	} else {
		data.Policy = ""
	}

	if !v.Deploy.IsNull() && !v.Deploy.IsUnknown() {
		data.Deploy = v.Deploy.ValueBool()
	}

	if !v.SerialNumber.IsNull() && !v.SerialNumber.IsUnknown() {
		data.SerialNumber = v.SerialNumber.ValueString()
	} else {
		data.SerialNumber = ""
	}

	if !v.Interfaces.IsNull() && !v.Interfaces.IsUnknown() {
		elements1 := make(map[string]InterfacesValue, len(v.Interfaces.Elements()))

		data.Interfaces = make(map[string]resource_interface_common.NDFCInterfacesValue)
		diag := v.Interfaces.ElementsAs(context.Background(), &elements1, false)
		if diag != nil {
			panic(diag)
		}
		for k1, ele1 := range elements1 {
			data1 := new(resource_interface_common.NDFCInterfacesValue)
			// filter_this_value | Bool| []| true
			// serial_number | String| []| false
			if !ele1.SerialNumber.IsNull() && !ele1.SerialNumber.IsUnknown() {

				data1.SerialNumber = ele1.SerialNumber.ValueString()
			} else {
				data1.SerialNumber = ""
			}

			// interface_name | String| []| false
			if !ele1.InterfaceName.IsNull() && !ele1.InterfaceName.IsUnknown() {

				data1.InterfaceName = ele1.InterfaceName.ValueString()
			} else {
				data1.InterfaceName = ""
			}

			// interface_type | String| []| true
			// interface_name | String| [nvPairs]| true
			// admin_state | Bool| [nvPairs]| false
			if !ele1.AdminState.IsNull() && !ele1.AdminState.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.AdminState = strconv.FormatBool(ele1.AdminState.ValueBool())
			} else {
				data1.NvPairs.AdminState = ""
			}

			// freeform_config | String| [nvPairs]| false
			if !ele1.FreeformConfig.IsNull() && !ele1.FreeformConfig.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.FreeformConfig = ele1.FreeformConfig.ValueString()
			} else {
				data1.NvPairs.FreeformConfig = ""
			}

			// interface_description | String| [nvPairs]| false
			if !ele1.InterfaceDescription.IsNull() && !ele1.InterfaceDescription.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.InterfaceDescription = ele1.InterfaceDescription.ValueString()
			} else {
				data1.NvPairs.InterfaceDescription = ""
			}

			// vrf | String| [nvPairs]| false
			if !ele1.Vrf.IsNull() && !ele1.Vrf.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.Vrf = ele1.Vrf.ValueString()
			} else {
				data1.NvPairs.Vrf = ""
			}

			// ipv4_address | String| [nvPairs]| false
			if !ele1.Ipv4Address.IsNull() && !ele1.Ipv4Address.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.Ipv4Address = ele1.Ipv4Address.ValueString()
			} else {
				data1.NvPairs.Ipv4Address = ""
			}

			// ipv6_address | String| [nvPairs]| false
			if !ele1.Ipv6Address.IsNull() && !ele1.Ipv6Address.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.Ipv6Address = ele1.Ipv6Address.ValueString()
			} else {
				data1.NvPairs.Ipv6Address = ""
			}

			// route_map_tag | String| [nvPairs]| false
			if !ele1.RouteMapTag.IsNull() && !ele1.RouteMapTag.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.RouteMapTag = ele1.RouteMapTag.ValueString()
			} else {
				data1.NvPairs.RouteMapTag = ""
			}

			// deployment_status | String| []| false
			data.Interfaces[k1] = *data1

		}
	}

	return data
}
