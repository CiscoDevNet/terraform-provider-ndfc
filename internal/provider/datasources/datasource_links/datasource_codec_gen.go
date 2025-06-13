// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package datasource_links

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCLinksModel struct {
	LinkUuid          string          `json:"link_uuid,omitempty"`
	SourceFabric      string          `json:"sourceFabric,omitempty"`
	DestinationFabric string          `json:"destinationFabric,omitempty"`
	SourceDevice      string          `json:"sourceDevice,omitempty"`
	DestinationDevice string          `json:"destinationDevice,omitempty"`
	Links             NDFCLinksValues `json:"links,omitempty"`
}

type NDFCLinksValues []NDFCLinksValue

type NDFCLinksValue struct {
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

	if len(jsonData.Links) == 0 {
		log.Printf("v.Links is empty")
		v.Links = types.ListNull(LinksValue{}.Type(context.Background()))
	} else {
		log.Printf("v.Links contains %d elements", len(jsonData.Links))
		listData := make([]LinksValue, 0)
		for _, item := range jsonData.Links {
			data := new(LinksValue)
			err = data.SetValue(&item)
			if err != nil {
				log.Printf("Error in LinksValue.SetValue")
				return err
			}
			data.state = attr.ValueStateKnown
			listData = append(listData, *data)
		}
		v.Links, err = types.ListValueFrom(context.Background(), LinksValue{}.Type(context.Background()), listData)
		if err != nil {
			log.Printf("Error in converting []LinksValue to  List")
			return err
		}
	}

	return err
}

func (v *LinksValue) SetValue(jsonData *NDFCLinksValue) diag.Diagnostics {

	var err diag.Diagnostics
	err = nil

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

	return data
}
