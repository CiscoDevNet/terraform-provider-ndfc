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

func (i *NDFCVPCInterface) DeleteInterface(ctx context.Context, dg *diag.Diagnostics, id string,
	inData *resource_interface_common.NDFCInterfaceCommonModel) {

	tflog.Debug(ctx, fmt.Sprintf("NDFCVPCInterface: Deleting interfaces policy %s", inData.Policy))
	if len(inData.Interfaces) <= 0 {
		tflog.Debug(ctx, "No interfaces to delete")
		return
	}
	// DELETE and Deploy uses similar payload
	intfPayload := resource_interface_common.NDFCInterfacesDeploy{}
	/*
	 * For vPC interface, the serial number is of the format switch1~switch2
	 * when a vPC interface is created, the serial number send in payload is switch1~switch2 as given in config
	 * however NDFC can accept the config, but maintains the serial number in reverse order i.e switch2~switch1
	 * In GET responses this reverse order is sent back by NDFC.
	 * This causes inconsistency in TF and to avoid this, the order used in config is maintained in state.
	 * As the APIs accept only the order maintained in NDFC, using the serial_number from state may fail
	 * Hence, use the GET to retrieve the interfaces and take serial from the GET payload
	 */
	ifMap := ifIdToMap(id)
	serialMap := make(map[string]bool)
	if inData.SerialNumber == "" {
		if len(ifMap) <= 0 {
			tflog.Debug(ctx, "ID is empty - we may have to still delete any entries that are in inData")
			// find all the unique serial numbers if global serial is not set
			for _, intfToDel := range inData.Interfaces {
				serialMap[intfToDel.SerialNumber] = true
			}
		} else {
			for switchSerial := range ifMap {
				serialMap[switchSerial] = true
			}
		}
	} else {
		serialMap[inData.SerialNumber] = true
	}

	// Do for each switch
	for switchSerial := range serialMap {
		interfaces := i.GetInterface(ctx, dg, switchSerial, inData.Policy)
		// All the interfaces from the switch with the policy are returned
		for _, intf := range interfaces {
			// filter out interfaces that are in the resource
			// Delete only that are part of the resource
			for _, intfToDel := range inData.Interfaces {
				if intf.InterfaceName == intfToDel.InterfaceName && intfToDel.SerialNumber == switchSerial {
					intfPayload = append(intfPayload, resource_interface_common.NDFCInterfaceDeploy{IfName: intf.InterfaceName,
						SerialNumber: intf.SerialNumber})
					tflog.Debug(ctx, fmt.Sprintf("NDFCVPCInterface: Deleting interface: %s:%s", intf.SerialNumber, intf.InterfaceName))
				}
			}
		}
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
