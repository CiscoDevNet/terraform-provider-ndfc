// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package resource_links

import (
	"context"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCLinksModel struct {
	LinkUuid              string            `json:"-"`
	SourceFabric          string            `json:"sourceFabric,omitempty"`
	DestinationFabric     string            `json:"destinationFabric,omitempty"`
	SourceDevice          string            `json:"sourceDevice,omitempty"`
	DestinationDevice     string            `json:"destinationDevice,omitempty"`
	SourceInterface       string            `json:"sourceInterface,omitempty"`
	DestinationInterface  string            `json:"destinationInterface,omitempty"`
	TemplateName          string            `json:"templateName,omitempty"`
	LinkParameters        map[string]string `json:"nvPairs,omitempty"`
	LinkParamsComputed    map[string]string `json:"-"`
	SourceSwitchName      string            `json:"sourceSwitchName,omitempty"`
	DestinationSwitchName string            `json:"destinationSwitchName,omitempty"`
	LinkDbid              string            `json:"link_dbid,omitempty"`
	IsPresent             string            `json:"is_present,omitempty"`
	LinkType              string            `json:"link_type,omitempty"`
	IsDiscovered          string            `json:"is_discovered,omitempty"`
}

func (v *LinksModel) SetModelData(jsonData *NDFCLinksModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.LinkUuid != "" {
		v.LinkUuid = types.StringValue(jsonData.LinkUuid)
	} else {
		v.LinkUuid = types.StringNull()
	}

	if jsonData.SourceFabric != "" {
		v.SourceFabric = types.StringValue(jsonData.SourceFabric)
	} else {
		v.SourceFabric = types.StringNull()
	}

	if jsonData.DestinationFabric != "" {
		v.DestinationFabric = types.StringValue(jsonData.DestinationFabric)
	} else {
		v.DestinationFabric = types.StringNull()
	}

	if jsonData.SourceDevice != "" {
		v.SourceDevice = types.StringValue(jsonData.SourceDevice)
	} else {
		v.SourceDevice = types.StringNull()
	}

	if jsonData.DestinationDevice != "" {
		v.DestinationDevice = types.StringValue(jsonData.DestinationDevice)
	} else {
		v.DestinationDevice = types.StringNull()
	}

	if jsonData.SourceInterface != "" {
		v.SourceInterface = types.StringValue(jsonData.SourceInterface)
	} else {
		v.SourceInterface = types.StringNull()
	}

	if jsonData.DestinationInterface != "" {
		v.DestinationInterface = types.StringValue(jsonData.DestinationInterface)
	} else {
		v.DestinationInterface = types.StringNull()
	}

	if jsonData.TemplateName != "" {
		v.TemplateName = types.StringValue(jsonData.TemplateName)
	} else {
		v.TemplateName = types.StringNull()
	}

	if len(jsonData.LinkParameters) == 0 {
		log.Printf("v.LinkParameters is empty")
		v.LinkParameters = types.MapNull(types.StringType)
	} else {
		mapData := make(map[string]attr.Value)
		for key, item := range jsonData.LinkParameters {
			mapData[key] = types.StringValue(item)
		}
		v.LinkParameters, err = types.MapValue(types.StringType, mapData)
		if err != nil {
			log.Printf("Error in converting map[string]string to  Map")
			return err
		}
	}

	if len(jsonData.LinkParamsComputed) == 0 {
		log.Printf("v.LinkParamsComputed is empty")
		v.LinkParamsComputed = types.MapNull(types.StringType)
	} else {
		mapData := make(map[string]attr.Value)
		for key, item := range jsonData.LinkParamsComputed {
			mapData[key] = types.StringValue(item)
		}
		v.LinkParamsComputed, err = types.MapValue(types.StringType, mapData)
		if err != nil {
			log.Printf("Error in converting map[string]string to  Map")
			return err
		}
	}
	if jsonData.SourceSwitchName != "" {
		v.SourceSwitchName = types.StringValue(jsonData.SourceSwitchName)
	} else {
		v.SourceSwitchName = types.StringNull()
	}

	if jsonData.DestinationSwitchName != "" {
		v.DestinationSwitchName = types.StringValue(jsonData.DestinationSwitchName)
	} else {
		v.DestinationSwitchName = types.StringNull()
	}

	if jsonData.IsPresent != "" {
		x, _ := strconv.ParseBool(jsonData.IsPresent)
		v.IsPresent = types.BoolValue(x)
	} else {
		v.IsPresent = types.BoolNull()
	}

	if jsonData.LinkType != "" {
		v.LinkType = types.StringValue(jsonData.LinkType)
	} else {
		v.LinkType = types.StringNull()
	}

	if jsonData.IsDiscovered != "" {
		x, _ := strconv.ParseBool(jsonData.IsDiscovered)
		v.IsDiscovered = types.BoolValue(x)
	} else {
		v.IsDiscovered = types.BoolNull()
	}

	return err
}

func (v LinksModel) GetModelData() *NDFCLinksModel {
	var data = new(NDFCLinksModel)

	//MARSHAL_BODY

	if !v.SourceFabric.IsNull() && !v.SourceFabric.IsUnknown() {
		data.SourceFabric = v.SourceFabric.ValueString()
	} else {
		data.SourceFabric = ""
	}

	if !v.DestinationFabric.IsNull() && !v.DestinationFabric.IsUnknown() {
		data.DestinationFabric = v.DestinationFabric.ValueString()
	} else {
		data.DestinationFabric = ""
	}

	if !v.SourceDevice.IsNull() && !v.SourceDevice.IsUnknown() {
		data.SourceDevice = v.SourceDevice.ValueString()
	} else {
		data.SourceDevice = ""
	}

	if !v.DestinationDevice.IsNull() && !v.DestinationDevice.IsUnknown() {
		data.DestinationDevice = v.DestinationDevice.ValueString()
	} else {
		data.DestinationDevice = ""
	}

	if !v.SourceInterface.IsNull() && !v.SourceInterface.IsUnknown() {
		data.SourceInterface = v.SourceInterface.ValueString()
	} else {
		data.SourceInterface = ""
	}

	if !v.DestinationInterface.IsNull() && !v.DestinationInterface.IsUnknown() {
		data.DestinationInterface = v.DestinationInterface.ValueString()
	} else {
		data.DestinationInterface = ""
	}

	if !v.TemplateName.IsNull() && !v.TemplateName.IsUnknown() {
		data.TemplateName = v.TemplateName.ValueString()
	} else {
		data.TemplateName = ""
	}

	if !v.LinkParameters.IsNull() && !v.LinkParameters.IsUnknown() {
		mapStringData := make(map[string]string)
		dg := v.LinkParameters.ElementsAs(context.Background(), &mapStringData, false)
		if dg.HasError() {
			panic(dg.Errors())
		}
		data.LinkParameters = make(map[string]string)
		for key, value := range mapStringData {
			data.LinkParameters[key] = value
		}
	}

	return data
}
