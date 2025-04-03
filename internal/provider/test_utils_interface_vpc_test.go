// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package provider

import (
	"strconv"
	"terraform-provider-ndfc/internal/provider/resources/resource_interface_common"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func InterfaceVpcModelHelperStateCheck(RscName string, c resource_interface_common.NDFCInterfaceCommonModel, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.Policy != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("policy").String(), c.Policy))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("policy").String(), "int_vpc_trunk_host"))
	}
	if c.Deploy {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("deploy").String(), "true"))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("deploy").String(), "false"))
	}
	if c.SerialNumber != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("serial_number").String(), c.SerialNumber))
	}
	for key, value := range c.Interfaces {
		attrNewPath := attrPath.AtName("interfaces").AtName(key)
		ret = append(ret, InterfaceVpcInterfacesValueHelperStateCheck(RscName, value, attrNewPath)...)
	}
	return ret
}

func InterfaceVpcInterfacesValueHelperStateCheck(RscName string, c resource_interface_common.NDFCInterfacesValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.SerialNumber != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("serial_number").String(), c.SerialNumber))
	}
	if c.InterfaceName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_name").String(), c.InterfaceName))
	}

	if c.NvPairs.AdminState != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("admin_state").String(), c.NvPairs.AdminState))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("admin_state").String(), "true"))
	}
	if c.NvPairs.BpduGuard != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bpdu_guard").String(), c.NvPairs.BpduGuard))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bpdu_guard").String(), "true"))
	}
	if c.NvPairs.PortTypeFast != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("port_type_fast").String(), c.NvPairs.PortTypeFast))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("port_type_fast").String(), "true"))
	}
	if c.NvPairs.Mtu != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mtu").String(), c.NvPairs.Mtu))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mtu").String(), "jumbo"))
	}
	if c.NvPairs.Speed != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("speed").String(), c.NvPairs.Speed))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("speed").String(), "Auto"))
	}
	if c.NvPairs.Netflow != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("netflow").String(), c.NvPairs.Netflow))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("netflow").String(), "false"))
	}
	if c.NvPairs.NetflowMonitor != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("netflow_monitor").String(), c.NvPairs.NetflowMonitor))
	}
	if c.NvPairs.NetflowSampler != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("netflow_sampler").String(), c.NvPairs.NetflowSampler))
	}
	if c.NvPairs.CopyPoDescription != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("copy_po_description").String(), c.NvPairs.CopyPoDescription))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("copy_po_description").String(), "false"))
	}
	if c.NvPairs.PortchannelMode != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("portchannel_mode").String(), c.NvPairs.PortchannelMode))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("portchannel_mode").String(), "on"))
	}

	if c.NvPairs.Peer1PoFreeformConfig != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("peer1_po_freeform_config").String(), c.NvPairs.Peer1PoFreeformConfig))
	}
	if c.NvPairs.Peer2PoFreeformConfig != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("peer2_po_freeform_config").String(), c.NvPairs.Peer2PoFreeformConfig))
	}
	if c.NvPairs.Peer1PoDescription != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("peer1_po_description").String(), c.NvPairs.Peer1PoDescription))
	}
	if c.NvPairs.Peer2PoDescription != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("peer2_po_description").String(), c.NvPairs.Peer2PoDescription))
	}
	if c.NvPairs.Peer1AllowedVlans != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("peer1_allowed_vlans").String(), c.NvPairs.Peer1AllowedVlans))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("peer1_allowed_vlans").String(), "none"))
	}
	if c.NvPairs.Peer2AllowedVlans != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("peer2_allowed_vlans").String(), c.NvPairs.Peer2AllowedVlans))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("peer2_allowed_vlans").String(), "none"))
	}
	if c.NvPairs.Peer1NativeVlan != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("peer1_native_vlan").String(), strconv.Itoa(int(*c.NvPairs.Peer1NativeVlan))))
	}
	if c.NvPairs.Peer2NativeVlan != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("peer2_native_vlan").String(), strconv.Itoa(int(*c.NvPairs.Peer2NativeVlan))))
	}
	if c.NvPairs.Peer1MemberInterfaces != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("peer1_member_interfaces").String(), c.NvPairs.Peer1MemberInterfaces))
	}
	if c.NvPairs.Peer2MemberInterfaces != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("peer2_member_interfaces").String(), c.NvPairs.Peer2MemberInterfaces))
	}
	if c.NvPairs.Peer1PortChannelId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("peer1_port_channel_id").String(), strconv.Itoa(int(*c.NvPairs.Peer1PortChannelId))))
	}
	if c.NvPairs.Peer2PortChannelId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("peer2_port_channel_id").String(), strconv.Itoa(int(*c.NvPairs.Peer2PortChannelId))))
	}
	if c.DeploymentStatus != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("deployment_status").String(), c.DeploymentStatus))
	}
	return ret
}
