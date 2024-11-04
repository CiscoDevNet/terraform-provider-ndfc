// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
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
	"terraform-provider-ndfc/internal/provider/resources/resource_vpc_pair"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func VpcPairModelHelperStateCheck(RscName string, c resource_vpc_pair.NDFCVpcPairModel, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.UseVirtualPeerlink != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("use_virtual_peerlink").String(), strconv.FormatBool(*c.UseVirtualPeerlink)))
	}
	if c.TemplateName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("template_name").String(), c.TemplateName))
	}
	if c.NvPairs.DomainId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("domain_id").String(), strconv.Itoa(int(*c.NvPairs.DomainId))))
	}
	if c.NvPairs.Peer1KeepAliveLocalIp != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("peer1_keep_alive_local_ip").String(), c.NvPairs.Peer1KeepAliveLocalIp))
	}
	if c.NvPairs.Peer2KeepAliveLocalIp != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("peer2_keep_alive_local_ip").String(), c.NvPairs.Peer2KeepAliveLocalIp))
	}
	if c.NvPairs.KeepAliveVrf != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("keep_alive_vrf").String(), c.NvPairs.KeepAliveVrf))
	}
	if c.NvPairs.KeepAliveHoldTimeout != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("keep_alive_hold_timeout").String(), strconv.Itoa(int(*c.NvPairs.KeepAliveHoldTimeout))))
	}
	if c.NvPairs.IsVpcPlus != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("is_vpc_plus").String(), c.NvPairs.IsVpcPlus))
	}
	if c.NvPairs.FabricpathSwitchId != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabricpath_switch_id").String(), c.NvPairs.FabricpathSwitchId))
	}
	if c.NvPairs.Peer1SourceLoopback != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("peer1_source_loopback").String(), c.NvPairs.Peer1SourceLoopback))
	}
	if c.NvPairs.Peer2SourceLoopback != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("peer2_source_loopback").String(), c.NvPairs.Peer2SourceLoopback))
	}
	if c.NvPairs.Peer1PrimaryIp != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("peer1_primary_ip").String(), c.NvPairs.Peer1PrimaryIp))
	}
	if c.NvPairs.Peer2PrimaryIp != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("peer2_primary_ip").String(), c.NvPairs.Peer2PrimaryIp))
	}
	if c.NvPairs.LoopbackSecondaryIp != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("loopback_secondary_ip").String(), c.NvPairs.LoopbackSecondaryIp))
	}
	if c.NvPairs.Peer1DomainConf != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("peer1_domain_conf").String(), c.NvPairs.Peer1DomainConf))
	}
	if c.NvPairs.ClearPolicy != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("clear_policy").String(), c.NvPairs.ClearPolicy))
	}
	if c.NvPairs.FabricName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_name").String(), c.NvPairs.FabricName))
	}
	if c.NvPairs.Peer1Pcid != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("peer1_pcid").String(), c.NvPairs.Peer1Pcid))
	}
	if c.NvPairs.Peer2Pcid != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("peer2_pcid").String(), c.NvPairs.Peer2Pcid))
	}
	if c.NvPairs.Peer1MemberInterfaces != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("peer1_member_interfaces").String(), c.NvPairs.Peer1MemberInterfaces))
	}
	if c.NvPairs.Peer2MemberInterfaces != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("peer2_member_interfaces").String(), c.NvPairs.Peer2MemberInterfaces))
	}
	if c.NvPairs.PcMode != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("pc_mode").String(), c.NvPairs.PcMode))
	}
	if c.NvPairs.Peer1PoDesc != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("peer1_po_desc").String(), c.NvPairs.Peer1PoDesc))
	}
	if c.NvPairs.Peer2PoDesc != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("peer2_po_desc").String(), c.NvPairs.Peer2PoDesc))
	}
	if c.NvPairs.AdminState != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("admin_state").String(), c.NvPairs.AdminState))
	}
	if c.NvPairs.AllowedVlans != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("allowed_vlans").String(), c.NvPairs.AllowedVlans))
	}
	if c.NvPairs.Peer1PoConf != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("peer1_po_conf").String(), c.NvPairs.Peer1PoConf))
	}
	if c.NvPairs.Peer2PoConf != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("peer2_po_conf").String(), c.NvPairs.Peer2PoConf))
	}
	if c.NvPairs.IsVteps != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("is_vteps").String(), c.NvPairs.IsVteps))
	}
	if c.NvPairs.NveInterface != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("nve_interface").String(), c.NvPairs.NveInterface))
	}
	return ret
}
