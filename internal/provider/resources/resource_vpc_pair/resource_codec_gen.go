// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package resource_vpc_pair

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCVpcPairModel struct {
	SerialNumbers        []string                      `json:"-"`
	PeerOneId            string                        `json:"peerOneId,omitempty"`
	PeerTwoId            string                        `json:"peerTwoId,omitempty"`
	UseVirtualPeerlink   *bool                         `json:"useVirtualPeerlink,omitempty"`
	Deploy               bool                          `json:"-"`
	PeerOneSwitchDetails NDFCPeerOneSwitchDetailsValue `json:"peerOneSwitchDetails,omitempty"`
}

type NDFCPeerOneSwitchDetailsValue struct {
	FabricName string `json:"fabricName,omitempty"`
}

func (v *VpcPairModel) SetModelData(jsonData *NDFCVpcPairModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if len(jsonData.SerialNumbers) == 0 {
		log.Printf("v.SerialNumbers is empty")
		v.SerialNumbers, err = types.SetValue(types.StringType, []attr.Value{})
		if err != nil {
			log.Printf("Error in converting []string to  List %v", err)
			return err
		}
	} else {
		listData := make([]attr.Value, len(jsonData.SerialNumbers))
		for i, item := range jsonData.SerialNumbers {
			listData[i] = types.StringValue(item)
		}
		v.SerialNumbers, err = types.SetValue(types.StringType, listData)
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}

	if jsonData.UseVirtualPeerlink != nil {
		v.UseVirtualPeerlink = types.BoolValue(*jsonData.UseVirtualPeerlink)

	} else {
		v.UseVirtualPeerlink = types.BoolNull()
	}

	v.Deploy = types.BoolValue(jsonData.Deploy)

	return err
}

func (v VpcPairModel) GetModelData() *NDFCVpcPairModel {
	var data = new(NDFCVpcPairModel)

	//MARSHAL_BODY

	if !v.SerialNumbers.IsNull() && !v.SerialNumbers.IsUnknown() {
		listStringData := make([]string, len(v.SerialNumbers.Elements()))
		dg := v.SerialNumbers.ElementsAs(context.Background(), &listStringData, false)
		if dg.HasError() {
			panic(dg.Errors())
		}
		data.SerialNumbers = make([]string, len(listStringData))
		copy(data.SerialNumbers, listStringData)
	}

	if !v.UseVirtualPeerlink.IsNull() && !v.UseVirtualPeerlink.IsUnknown() {
		data.UseVirtualPeerlink = new(bool)
		*data.UseVirtualPeerlink = v.UseVirtualPeerlink.ValueBool()
	} else {
		data.UseVirtualPeerlink = nil
	}

	if !v.Deploy.IsNull() && !v.Deploy.IsUnknown() {
		data.Deploy = v.Deploy.ValueBool()
	}

	return data
}
