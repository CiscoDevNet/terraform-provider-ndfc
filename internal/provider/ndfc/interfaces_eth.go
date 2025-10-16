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
	"slices"
	"strings"
	"terraform-provider-ndfc/internal/provider/ndfc/api"
	"terraform-provider-ndfc/internal/provider/resources/resource_interface_common"
	"terraform-provider-ndfc/internal/provider/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type NDFCEthernetInterface struct {
	NDFCInterfaceCommon
}

var ethIntfPolicyMap = map[string]bool{
	"int_trunk_host":  true,
	"int_routed_host": true,
	"int_access_host": true,
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

func (i *NDFCEthernetInterface) GetInterface(ctx context.Context, diags *diag.Diagnostics, serial string,
	policy string) []resource_interface_common.NDFCInterfacesValue {
	// NDFC automatically changes/removed policies on eth ports when they are associated with port channels
	// so retrieve everything without filtering for policy
	return i.getEthInterfaces(ctx, diags, serial, policy)
}

func (i *NDFCEthernetInterface) DeleteInterface(ctx context.Context, dg *diag.Diagnostics, id string,
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

// getEthInterfaces retrieves all Ethernet interfaces for a given switch serial number.
// It queries NDFC without filtering by policy template because NDFC automatically modifies
// or removes policies on Ethernet ports when they are associated with port channels.
// The function returns only interfaces with names starting with "eth", "Eth", or "ETH".
func (i *NDFCEthernetInterface) getEthInterfaces(ctx context.Context, diags *diag.Diagnostics, serial string, policy string) []resource_interface_common.NDFCInterfacesValue {
	ifList := make([]resource_interface_common.NDFCInterfacesPayload, 0)
	ifApi := api.NewInterfaceAPI(i.lock, i.client)
	ifApi.SwitchSerial = serial
	ifApi.Policy = ""
	res, err := ifApi.Get()
	if err != nil {
		tflog.Error(ctx, "Error getting interfaces")
		diags.AddError("Error getting interfaces", err.Error())
		return nil
	}
	log.Printf("Response=%s", string(res))
	err = json.Unmarshal((res), &ifList)
	if err != nil {
		diags.AddError("Error unmarshalling data", err.Error())
		return nil
	}
	if len(ifList) == 0 {
		diags.AddWarning("No interfaces found", "")
		return nil
	}
	retList := make([]resource_interface_common.NDFCInterfacesValue, 0)
	for _, ifGroup := range ifList {
		// We are only interested in the ports with configured policy
		// or those changed due to membership in a port-channel
		pcMember := ""
		if ifGroup.Policy != policy {
			if strings.HasPrefix(ifGroup.Policy, "int_port_channel_trunk_member") {
				pcMember = ifGroup.Policy
			} else {
				tflog.Debug(ctx, fmt.Sprintf("Skipping Policies %s", ifGroup.Policy))
				continue
			}
		}
		for index := range ifGroup.Interfaces {
			// Check if ifName is Ethernet, ethernet, eth etc
			intf := &ifGroup.Interfaces[index]
			log.Printf("Interface Name: %s", intf.InterfaceName)
			if strings.HasPrefix(strings.ToLower(intf.InterfaceName), "eth") {
				intf.PortChannelPolicy = pcMember
				retList = append(retList, *intf)
			}
		}
	}
	return retList
}

// Ethernet need special handling to take care of port channel movements

func (ei *NDFCEthernetInterface) ModifyInterface(ctx context.Context, dg *diag.Diagnostics,
	inData *resource_interface_common.NDFCInterfaceCommonModel) {

	tflog.Debug(ctx, "Modifying Ethernet interfaces")
	pcMap := make(map[string]*resource_interface_common.NDFCInterfacesPayload)
	if len(inData.Interfaces) > 0 {
		pcPresent := false
		updateIntf := new(resource_interface_common.NDFCInterfacesPayload)
		updateIntf.Policy = inData.Policy
		for i := range inData.Interfaces {
			if inData.Interfaces[i].PortChannelPolicy != "" {
				log.Printf("PortChannelPolicy: %s set on port %s", inData.Interfaces[i].PortChannelPolicy, inData.Interfaces[i].InterfaceName)
				pcPresent = true
				// This is an interface that is part of a port channel
				payloadIntf, ok := pcMap[inData.Interfaces[i].PortChannelPolicy]
				if !ok {
					payloadIntf = new(resource_interface_common.NDFCInterfacesPayload)
					payloadIntf.Policy = inData.Interfaces[i].PortChannelPolicy
					payloadIntf.Interfaces = make([]resource_interface_common.NDFCInterfacesValue, 0)
					payloadIntf.Interfaces = append(payloadIntf.Interfaces, inData.Interfaces[i])
				} else {
					payloadIntf.Interfaces = append(payloadIntf.Interfaces, inData.Interfaces[i])
				}
				//delete the entry from inData
				intf := inData.Interfaces[i]
				intf.FilterThisValue = true
				inData.Interfaces[i] = intf
				pcMap[inData.Interfaces[i].PortChannelPolicy] = payloadIntf
			} else {
				log.Printf("PortChannelPolicy: %s not set on port %s", inData.Interfaces[i].PortChannelPolicy, inData.Interfaces[i].InterfaceName)
			}
			updateIntf.Interfaces = append(updateIntf.Interfaces, inData.Interfaces[i])
		}
		if pcPresent {
			for _, payloadIntf := range pcMap {
				ei.modifyInterface(ctx, dg, payloadIntf)
				if dg.HasError() {
					tflog.Error(ctx, "Error updating interfaces")
					return
				}
			}
			// Remove the filtered interfaces from the update list
			// Its tricky to remove entries from a slice
			for i := 0; i < len(updateIntf.Interfaces); {
				if updateIntf.Interfaces[i].FilterThisValue {
					updateIntf.Interfaces = slices.Delete(updateIntf.Interfaces, i, i+1)
				} else {
					i++
				}
			}
		}
		// process remaining if any
		if len(updateIntf.Interfaces) > 0 {
			ei.modifyInterface(ctx, dg, updateIntf)
			if dg.HasError() {
				tflog.Error(ctx, "Error updating interfaces")
				return
			}
		}
	} else {
		tflog.Debug(ctx, "No interfaces to modify")
	}
}

func (ei *NDFCEthernetInterface) ModifyAttributesForTerraform(ctx context.Context, dg *diag.Diagnostics,
	toModel *resource_interface_common.NDFCInterfacesValue,
	plan *resource_interface_common.NDFCInterfacesValue) {
	/*
	   Making a major assumption here

	   The movement of eth ports to port-channel is only detected in subsequent operations
	   It cannot be detected in the run where the change happened, as typically the change happens
	   in a port-channel resource, after eth resource completes its execution

	   So this code works only if a plan is triggered that fetches the objects from NDFC
	   and then a subsequent apply is run. This would cause the GET to fetch the objects from NDFC
	   and the port-channel info is updated here in this funtion.

	   All Reads (GET) after the port-channel movement should reflect the info

	   As a subsequent apply with any modifications to the eth object would've updated the
	   state with the port-channel info, the modification is expected to apply the new policy

	   Side cases, terraform plan with --refresh=false will not refresh the state from NDFC

	   If user runs terraform plan with --refresh=false saves it and then run terraform apply
	   this flow maynot work properly as the state will not have the port-channel info
	   In this case toModel.PortChannelPolicy is empty like a normal eth port.
	   That would cause the Eth ports  to be overwritten with previously applied policy
	*/

	if toModel.PortChannelPolicy != "" {
		log.Printf("Port %s is moved to a port channel %s - So many parameters that were originally set are not in the GET response", toModel.InterfaceName, toModel.NvPairs.PortChannelName)
		dg.AddWarning("Attributes changes ignored", fmt.Sprintf("Port %s is now member of a port channel %s - attributes that cannot be modifed in the new policy are ignored", toModel.InterfaceName, toModel.NvPairs.PortChannelName))
		// Look for changed attributes
		nvTmp := toModel.NvPairs
		// Overwrite stuff from plan
		toModel.NvPairs = plan.NvPairs
		// Copy computed stuff and other attributes that are changeable in port-chanel policy
		toModel.NvPairs.PortChannelName = nvTmp.PortChannelName

		if nvTmp.InterfaceDescription != "" {
			toModel.NvPairs.InterfaceDescription = nvTmp.InterfaceDescription
		}
		if nvTmp.Mtu != "" {
			toModel.NvPairs.Mtu = nvTmp.Mtu
		}
		if nvTmp.PortTypeFast != "" {
			toModel.NvPairs.PortTypeFast = nvTmp.PortTypeFast
		}
		if nvTmp.BpduGuard != "" {
			toModel.NvPairs.BpduGuard = nvTmp.BpduGuard
		}
		if nvTmp.AccessVlan != nil {
			toModel.NvPairs.AccessVlan = nvTmp.AccessVlan
		}
		if nvTmp.AllowedVlans != "" {
			toModel.NvPairs.AllowedVlans = nvTmp.AllowedVlans
		}
		if nvTmp.FreeformConfig != "" {
			toModel.NvPairs.FreeformConfig = nvTmp.FreeformConfig
		}
		if nvTmp.NativeVlan != nil {
			toModel.NvPairs.NativeVlan = nvTmp.NativeVlan
		}
		if nvTmp.Ptp != "" {
			toModel.NvPairs.Ptp = nvTmp.Ptp
		}
	}
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
