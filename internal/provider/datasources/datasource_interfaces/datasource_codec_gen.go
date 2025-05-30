// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package datasource_interfaces

import (
	"context"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCInterfacesModel struct {
	SerialNumber   string               `json:"-"`
	InterfaceTypes string               `json:"-"`
	PortModes      string               `json:"-"`
	Excludes       string               `json:"-"`
	Interfaces     NDFCInterfacesValues `json:"interfaces,omitempty"`
}

type NDFCInterfacesValues []NDFCInterfacesValue

type NDFCInterfacesValue struct {
	InterfaceName    string `json:"ifName,omitempty"`
	InterfaceType    string `json:"ifType,omitempty"`
	InterfaceIndex   *int64 `json:"ifIndex,omitempty"`
	IsPhysical       string `json:"isPhysical,omitempty"`
	Mode             string `json:"mode,omitempty"`
	NativeVlanId     *int64 `json:"nativeVlanId,omitempty"`
	OperStatus       string `json:"operStatusStr,omitempty"`
	OperStatusCause  string `json:"operStatusCause,omitempty"`
	PolicyName       string `json:"policyName,omitempty"`
	Speed            string `json:"speedStr,omitempty"`
	SwitchName       string `json:"sysName,omitempty"`
	Vrf              string `json:"vrf,omitempty"`
	AdminStatus      string `json:"adminStatusStr,omitempty"`
	AllowedVlans     string `json:"allowedVLANs,omitempty"`
	FabricName       string `json:"fabricName,omitempty"`
	Ipv4Address      string `json:"ipAddress,omitempty"`
	Description      string `json:"alias,omitempty"`
	DeploymentStatus string `json:"complianceStatus,omitempty"`
	SwitchDbid       *int64 `json:"switchDbId,omitempty"`
}

func (v *InterfacesModel) SetModelData(jsonData *NDFCInterfacesModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.SerialNumber != "" {
		v.SerialNumber = types.StringValue(jsonData.SerialNumber)
	} else {
		v.SerialNumber = types.StringNull()
	}

	if jsonData.InterfaceTypes != "" {
		v.InterfaceTypes = types.StringValue(jsonData.InterfaceTypes)
	} else {
		v.InterfaceTypes = types.StringNull()
	}

	if jsonData.PortModes != "" {
		v.PortModes = types.StringValue(jsonData.PortModes)
	} else {
		v.PortModes = types.StringNull()
	}

	if jsonData.Excludes != "" {
		v.Excludes = types.StringValue(jsonData.Excludes)
	} else {
		v.Excludes = types.StringNull()
	}

	if len(jsonData.Interfaces) == 0 {
		log.Printf("v.Interfaces is empty")
		v.Interfaces = types.ListNull(InterfacesValue{}.Type(context.Background()))
	} else {
		log.Printf("v.Interfaces contains %d elements", len(jsonData.Interfaces))
		listData := make([]InterfacesValue, 0)
		for _, item := range jsonData.Interfaces {
			data := new(InterfacesValue)
			err = data.SetValue(&item)
			if err != nil {
				log.Printf("Error in InterfacesValue.SetValue")
				return err
			}
			data.state = attr.ValueStateKnown
			listData = append(listData, *data)
		}
		v.Interfaces, err = types.ListValueFrom(context.Background(), InterfacesValue{}.Type(context.Background()), listData)
		if err != nil {
			log.Printf("Error in converting []InterfacesValue to  List")
			return err
		}
	}

	return err
}

func (v *InterfacesValue) SetValue(jsonData *NDFCInterfacesValue) diag.Diagnostics {

	var err diag.Diagnostics
	err = nil

	if jsonData.InterfaceName != "" {
		v.InterfaceName = types.StringValue(jsonData.InterfaceName)
	} else {
		v.InterfaceName = types.StringNull()
	}

	if jsonData.InterfaceType != "" {
		v.InterfaceType = types.StringValue(jsonData.InterfaceType)
	} else {
		v.InterfaceType = types.StringNull()
	}

	if jsonData.InterfaceIndex != nil {
		v.InterfaceIndex = types.Int64Value(*jsonData.InterfaceIndex)

	} else {
		v.InterfaceIndex = types.Int64Null()
	}

	if jsonData.IsPhysical != "" {
		x, _ := strconv.ParseBool(jsonData.IsPhysical)
		v.IsPhysical = types.BoolValue(x)
	} else {
		v.IsPhysical = types.BoolNull()
	}

	if jsonData.Mode != "" {
		v.Mode = types.StringValue(jsonData.Mode)
	} else {
		v.Mode = types.StringNull()
	}

	if jsonData.NativeVlanId != nil {
		v.NativeVlanId = types.Int64Value(*jsonData.NativeVlanId)

	} else {
		v.NativeVlanId = types.Int64Null()
	}

	if jsonData.OperStatus != "" {
		v.OperStatus = types.StringValue(jsonData.OperStatus)
	} else {
		v.OperStatus = types.StringNull()
	}

	if jsonData.OperStatusCause != "" {
		v.OperStatusCause = types.StringValue(jsonData.OperStatusCause)
	} else {
		v.OperStatusCause = types.StringNull()
	}

	if jsonData.PolicyName != "" {
		v.PolicyName = types.StringValue(jsonData.PolicyName)
	} else {
		v.PolicyName = types.StringNull()
	}

	if jsonData.Speed != "" {
		v.Speed = types.StringValue(jsonData.Speed)
	} else {
		v.Speed = types.StringNull()
	}

	if jsonData.SwitchName != "" {
		v.SwitchName = types.StringValue(jsonData.SwitchName)
	} else {
		v.SwitchName = types.StringNull()
	}

	if jsonData.Vrf != "" {
		v.Vrf = types.StringValue(jsonData.Vrf)
	} else {
		v.Vrf = types.StringNull()
	}

	if jsonData.AdminStatus != "" {
		v.AdminStatus = types.StringValue(jsonData.AdminStatus)
	} else {
		v.AdminStatus = types.StringNull()
	}

	if jsonData.AllowedVlans != "" {
		v.AllowedVlans = types.StringValue(jsonData.AllowedVlans)
	} else {
		v.AllowedVlans = types.StringNull()
	}

	if jsonData.FabricName != "" {
		v.FabricName = types.StringValue(jsonData.FabricName)
	} else {
		v.FabricName = types.StringNull()
	}

	if jsonData.Ipv4Address != "" {
		v.Ipv4Address = types.StringValue(jsonData.Ipv4Address)
	} else {
		v.Ipv4Address = types.StringNull()
	}

	if jsonData.Description != "" {
		v.Description = types.StringValue(jsonData.Description)
	} else {
		v.Description = types.StringNull()
	}

	if jsonData.DeploymentStatus != "" {
		v.DeploymentStatus = types.StringValue(jsonData.DeploymentStatus)
	} else {
		v.DeploymentStatus = types.StringNull()
	}

	if jsonData.SwitchDbid != nil {
		v.SwitchDbid = types.Int64Value(*jsonData.SwitchDbid)

	} else {
		v.SwitchDbid = types.Int64Null()
	}

	return err
}

func (v InterfacesModel) GetModelData() *NDFCInterfacesModel {
	var data = new(NDFCInterfacesModel)

	//MARSHAL_BODY

	if !v.SerialNumber.IsNull() && !v.SerialNumber.IsUnknown() {
		data.SerialNumber = v.SerialNumber.ValueString()
	} else {
		data.SerialNumber = ""
	}

	if !v.InterfaceTypes.IsNull() && !v.InterfaceTypes.IsUnknown() {
		data.InterfaceTypes = v.InterfaceTypes.ValueString()
	} else {
		data.InterfaceTypes = ""
	}

	if !v.PortModes.IsNull() && !v.PortModes.IsUnknown() {
		data.PortModes = v.PortModes.ValueString()
	} else {
		data.PortModes = ""
	}

	if !v.Excludes.IsNull() && !v.Excludes.IsUnknown() {
		data.Excludes = v.Excludes.ValueString()
	} else {
		data.Excludes = ""
	}

	//MARSHALL_LIST

	if !v.Interfaces.IsNull() && !v.Interfaces.IsUnknown() {
		elements := make([]InterfacesValue, len(v.Interfaces.Elements()))
		data.Interfaces = make([]NDFCInterfacesValue, len(v.Interfaces.Elements()))

		diag := v.Interfaces.ElementsAs(context.Background(), &elements, false)
		if diag != nil {
			panic(diag)
		}
		for i1, ele1 := range elements {

			if !ele1.Mode.IsNull() && !ele1.Mode.IsUnknown() {

				data.Interfaces[i1].Mode = ele1.Mode.ValueString()
			} else {
				data.Interfaces[i1].Mode = ""
			}

		}
	}

	return data
}
