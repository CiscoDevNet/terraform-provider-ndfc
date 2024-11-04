// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package ndfc

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"terraform-provider-ndfc/internal/provider/resources/resource_interface_common"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const ResourceVPCInterface = "interface_vpc"

type NDFCVPCInterface struct {
	NDFCInterfaceCommon
}

func (i *NDFCVPCInterface) CreateInterface(ctx context.Context, diags *diag.Diagnostics, inData *resource_interface_common.NDFCInterfaceCommonModel) {
	intfPayload := resource_interface_common.NDFCInterfacesPayload{}
	intfPayload.Policy = inData.Policy
	if len(inData.Interfaces) <= 0 {
		tflog.Debug(ctx, "No interfaces to create")
		return
	}
	for _, intf := range inData.Interfaces {
		intfPayload.Interfaces = append(intfPayload.Interfaces, intf)
	}
	i.createInterface(ctx, diags, &intfPayload)
}

func (i *NDFCVPCInterface) DeleteInterface(ctx context.Context, dg *diag.Diagnostics,
	inData *resource_interface_common.NDFCInterfaceCommonModel) {

	tflog.Debug(ctx, "NDFCVPCInterface: Deleting interfaces")
	if len(inData.Interfaces) <= 0 {
		tflog.Debug(ctx, "No interfaces to delete")
		return
	}

	// DELETE and Deploy uses similar payload
	intfPayload := resource_interface_common.NDFCInterfacesDeploy{}

	//ifDeployPayload := resource_interface_common.NDFCInterfacesDeploy{}

	for _, intf := range inData.Interfaces {
		intfPayload = append(intfPayload, resource_interface_common.NDFCInterfaceDeploy{IfName: intf.InterfaceName,
			SerialNumber: intf.SerialNumber})
		tflog.Debug(ctx, fmt.Sprintf("NDFCVPCInterface: Deleting interface: %s:%s", intf.SerialNumber, intf.InterfaceName))
	}

	i.deleteInterface(ctx, dg, &intfPayload)
	if dg.HasError() {
		tflog.Error(ctx, "Error deleting interfaces")
		return
	}
	i.deployInterface(ctx, dg, intfPayload)
}

// Override
func (i *NDFCVPCInterface) ModifyInterface(ctx context.Context, dg *diag.Diagnostics,
	inData *resource_interface_common.NDFCInterfaceCommonModel) {
	if len(inData.Interfaces) > 0 {
		updateIntf := new(resource_interface_common.NDFCInterfacesPayload)
		updateIntf.Policy = inData.Policy
		for _, intf := range inData.Interfaces {
			updateIntf.Interfaces = append(updateIntf.Interfaces, intf)
		}
		i.modifyInterface(ctx, dg, updateIntf)
		if dg.HasError() {
			tflog.Error(ctx, "Error updating interfaces")
			return
		}
	} else {
		tflog.Debug(ctx, "No interfaces to modify")
	}
}

func (i *NDFCVPCInterface) GetPayload(ctx context.Context, diags *diag.Diagnostics,
	intfPayload *resource_interface_common.NDFCInterfacesPayload) ([]byte, error) {

	tflog.Debug(ctx, "GetPayload - overrided call for NDFCVPCInterface")

	rawData, err := json.Marshal(intfPayload)
	if err != nil {
		return nil, err
	}
	removeAttributes := []string{
		",\"ROUTING_TAG\":\"\"", // VLAN interface needs empty "ROUTING_TAG" field in payload
	}
	log.Printf("Removing attributes from payload before sending to API")
	log.Printf("Data %s", string(rawData))
	for _, attr := range removeAttributes {
		rawData = []byte(strings.Replace(string(rawData), attr, "", -1))
	}
	log.Printf("Data after removing attributes %s", string(rawData))
	return rawData, nil
}
func (i *NDFCVPCInterface) GetInterface(ctx context.Context, diags *diag.Diagnostics, serial string,
	policy string) []resource_interface_common.NDFCInterfacesValue {
	// serial contains 2 switches, GET only supports one
	serial = strings.Split(serial, "~")[0]
	ifList := i.getInterfaces(ctx, diags, serial, policy)
	for i := range ifList {
		// sometimes ifName returns all small `vpc<x>`, use the INTF_NAME param
		ifList[i].InterfaceName = ifList[i].NvPairs.InterfaceName
	}
	return ifList
}
