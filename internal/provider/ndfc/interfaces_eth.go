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
	"terraform-provider-ndfc/internal/provider/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type NDFCEthernetInterface struct {
	NDFCInterfaceCommon
}

const ResourceEthernetInterface = "interface_ethernet"

func (i *NDFCEthernetInterface) CreateInterface(ctx context.Context, diags *diag.Diagnostics, inData *resource_interface_common.NDFCInterfaceCommonModel) {
	intfPayload := resource_interface_common.NDFCInterfacesPayload{}
	intfPayload.Policy = inData.Policy
	if len(inData.Interfaces) <= 0 {
		tflog.Debug(ctx, "No interfaces to create")
		return
	}
	for i, intf := range inData.Interfaces {
		inData.Interfaces[i] = intf
		intfPayload.Interfaces = append(intfPayload.Interfaces, intf)
	}
	i.modifyInterface(ctx, diags, &intfPayload)
}

func (i *NDFCEthernetInterface) DeleteInterface(ctx context.Context, dg *diag.Diagnostics,
	inData *resource_interface_common.NDFCInterfaceCommonModel) {
	tflog.Debug(ctx, "Deleting interfaces")
	if len(inData.Interfaces) <= 0 {
		tflog.Debug(ctx, "No interfaces to delete")
		return
	}
	intfPayload := resource_interface_common.NDFCInterfacesPayload{}

	ifDeployPayload := resource_interface_common.NDFCInterfacesDeploy{}
	client := getNDFCClient()
	resp := client.GetDeviceRole(ctx, dg, inData.SerialNumber)
	var role string
	if !resp.Exists() {
		// Avoiding error if the device role is not found and assuming the role is leaf
		role = "leaf"
	} else {
		role = resp.Array()[0].Get("role").String()
	}
	tflog.Debug(ctx, fmt.Sprintf("Device role: %s", role))
	if role == "leaf" {
		intfPayload.Policy = "int_trunk_host"
	} else {
		intfPayload.Policy = "int_routed_host"
	}
	for k, intf := range inData.Interfaces {
		tflog.Debug(ctx, fmt.Sprintf("Deleting interface: %s:%s", intf.SerialNumber, intf.InterfaceName))

		//Empty out some parameters
		intf.NvPairs.InterfaceDescription = "###"
		intf.NvPairs.FreeformConfig = "###"
		intf.NvPairs.NativeVlan = new(types.Int64Custom)
		*intf.NvPairs.NativeVlan = types.Int64Custom(-4096)
		intf.NvPairs.AccessVlan = new(types.Int64Custom)
		*intf.NvPairs.AccessVlan = types.Int64Custom(-4096)
		intf.NvPairs.Netflow = "false"
		intf.NvPairs.NetflowMonitor = "###"
		intf.NvPairs.NetflowSampler = "###"
		intf.NvPairs.InterfaceName = intf.InterfaceName
		if role == "leaf" {
			//Set default values when role is leaf
			intf.NvPairs.Speed = "Auto"
			intf.NvPairs.Mtu = "jumbo"
			intf.NvPairs.FreeformConfig = "no shutdown"
			intf.NvPairs.AllowedVlans = "none"
			intf.NvPairs.BpduGuard = "false"
			intf.NvPairs.PortTypeFast = "true"
		} else {
			intf.NvPairs.Speed = "Auto"
			intf.NvPairs.Mtu = "9216"
			intf.NvPairs.FreeformConfig = "no shutdown"
			intf.NvPairs.Vrf = ""
			intf.NvPairs.Ipv4Address = ""
			intf.NvPairs.Ipv4PrefixLength = ""
			intf.NvPairs.RoutingTag = ""
		}
		inData.Interfaces[k] = intf
		intfPayload.Interfaces = append(intfPayload.Interfaces, intf)
		ifDeployPayload = append(ifDeployPayload, resource_interface_common.NDFCInterfaceDeploy{
			IfName:       intf.InterfaceName,
			SerialNumber: intf.SerialNumber,
		})

	}
	i.modifyInterface(ctx, dg, &intfPayload)
	if dg.HasError() {
		tflog.Error(ctx, "Error deleting interfaces")
		return
	}
	i.deployInterface(ctx, dg, ifDeployPayload)
}

func (i *NDFCEthernetInterface) GetPayload(ctx context.Context, diags *diag.Diagnostics,
	intfPayload *resource_interface_common.NDFCInterfacesPayload) ([]byte, error) {

	tflog.Debug(ctx, "GetPayload - overrided call for NDFCEthernetInterface")

	rawData, err := json.Marshal(intfPayload)
	if err != nil {
		return nil, err
	}
	removeAttributes := []string{
		",\"ROUTING_TAG\":\"\"",
	}

	emptyOutAttributes := []string{
		"\"###\"",   //Empty string
		"\"-4096\"", //Empty VLAN
	}

	log.Printf("Removing attributes from payload before sending to API")
	log.Printf("Data %s", string(rawData))
	for _, attr := range removeAttributes {
		rawData = []byte(strings.Replace(string(rawData), attr, "", -1))
	}

	for _, attr := range emptyOutAttributes {
		rawData = []byte(strings.Replace(string(rawData), attr, "\"\"", -1))
	}

	log.Printf("Data after removing attributes %s", string(rawData))
	return rawData, nil
}

/*
func printModel(model resource_interface_common.InterfaceModel) {
	log.Printf("Model: %v", model)
	inData := model.GetModelData()
	log.Printf("Model Data: %v", inData)
	for i := range inData.Interfaces {
		log.Printf("Interface: |%v|", inData.Interfaces[i])
		log.Printf("Interface Name: |%s|", inData.Interfaces[i].InterfaceName)
		log.Printf("Serial Number: |%s|", inData.Interfaces[i].SerialNumber)
		log.Printf("Freeform Config: |%s|", inData.Interfaces[i].NvPairs.FreeformConfig)
		log.Printf("Speed: |%s|", inData.Interfaces[i].NvPairs.Speed)
		log.Printf("Mtu: |%s|", inData.Interfaces[i].NvPairs.Mtu)
		log.Printf("Port Type Fast: |%s|", inData.Interfaces[i].NvPairs.PortTypeFast)
		log.Printf("Bpdu Guard: |%s|", inData.Interfaces[i].NvPairs.BpduGuard)
		if inData.Interfaces[i].NvPairs.AccessVlan != nil {
			log.Printf("Access Vlan: |%d|", *inData.Interfaces[i].NvPairs.AccessVlan)
		}
		log.Printf("Interface Description: |%s|", inData.Interfaces[i].NvPairs.InterfaceDescription)
		log.Printf("Orphan Port: |%s|", inData.Interfaces[i].NvPairs.OrphanPort)
		log.Printf("AdminState: |%s|", inData.Interfaces[i].NvPairs.AdminState)
		log.Printf("Ptp : |%s|", inData.Interfaces[i].NvPairs.Ptp)
		log.Printf("Netflow : |%s|", inData.Interfaces[i].NvPairs.Netflow)
		log.Printf("NetflowMonitor : |%s|", inData.Interfaces[i].NvPairs.NetflowMonitor)
		log.Printf("NetflowSampler : |%s|", inData.Interfaces[i].NvPairs.NetflowSampler)
		log.Printf("AllowedVlans : |%s|", inData.Interfaces[i].NvPairs.AllowedVlans)
		log.Printf("NativeVlan : |%v|", *inData.Interfaces[i].NvPairs.NativeVlan)
		log.Printf("Interface Type: |%s|", model.GetInterfaceType())
		log.Printf("Policy: |%s|", inData.Policy)
	}
}
*/
