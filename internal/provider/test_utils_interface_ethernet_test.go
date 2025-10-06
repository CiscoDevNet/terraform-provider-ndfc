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

func InterfaceEthernetModelHelperStateCheck(RscName string, c resource_interface_common.NDFCInterfaceCommonModel, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.Policy != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("policy").String(), c.Policy))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("policy").String(), "int_trunk_host"))
	}
	if c.PolicyType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("policy_type").String(), c.PolicyType))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("policy_type").String(), "system"))
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
		ret = append(ret, InterfaceEthernetInterfacesValueHelperStateCheck(RscName, value, attrNewPath)...)
	}
	return ret
}

func InterfacesPVlanMappingListValueHelperStateCheck(RscName string, c resource_interface_common.NDFCPVlanMappingListValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.SVlan != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("s_vlan").String(), c.SVlan))
	}
	if c.PVlan != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("p_vlan").String(), c.PVlan))
	}
	return ret
}

func InterfacesPVlanAssocListValueHelperStateCheck(RscName string, c resource_interface_common.NDFCPVlanAssocListValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.SVlan != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("s_vlan").String(), c.SVlan))
	}
	if c.PVlan != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("p_vlan").String(), c.PVlan))
	}
	return ret
}

func InterfaceEthernetInterfacesValueHelperStateCheck(RscName string, c resource_interface_common.NDFCInterfacesValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.SerialNumber != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("serial_number").String(), c.SerialNumber))
	}
	if c.InterfaceName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_name").String(), c.InterfaceName))
	}

	if c.NvPairs.FreeformConfig != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("freeform_config").String(), c.NvPairs.FreeformConfig))
	}
	if c.NvPairs.AdminState != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("admin_state").String(), c.NvPairs.AdminState))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("admin_state").String(), "true"))
	}
	if c.NvPairs.InterfaceDescription != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_description").String(), c.NvPairs.InterfaceDescription))
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
	if c.NvPairs.AccessVlan != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("access_vlan").String(), strconv.Itoa(int(*c.NvPairs.AccessVlan))))
	}
	if c.NvPairs.OrphanPort != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("orphan_port").String(), c.NvPairs.OrphanPort))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("orphan_port").String(), "false"))
	}
	if c.NvPairs.Ptp != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ptp").String(), c.NvPairs.Ptp))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ptp").String(), "false"))
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
	if c.NvPairs.AllowedVlans != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("allowed_vlans").String(), c.NvPairs.AllowedVlans))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("allowed_vlans").String(), "none"))
	}
	if c.NvPairs.NativeVlan != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("native_vlan").String(), strconv.Itoa(int(*c.NvPairs.NativeVlan))))
	}
	if c.NvPairs.Vrf != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf").String(), c.NvPairs.Vrf))
	}
	if c.NvPairs.Ipv4Address != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ipv4_address").String(), c.NvPairs.Ipv4Address))
	}
	if c.NvPairs.Ipv4PrefixLength != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ipv4_prefix_length").String(), c.NvPairs.Ipv4PrefixLength))
	}
	if c.NvPairs.RoutingTag != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("routing_tag").String(), c.NvPairs.RoutingTag))
	}
	if c.NvPairs.DisableIpRedirects != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("disable_ip_redirects").String(), c.NvPairs.DisableIpRedirects))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("disable_ip_redirects").String(), "false"))
	}
	if c.NvPairs.EnablePimSparse != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_pim_sparse").String(), c.NvPairs.EnablePimSparse))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_pim_sparse").String(), "false"))
	}
	if c.NvPairs.PimDrPriority != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("pim_dr_priority").String(), c.NvPairs.PimDrPriority))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("pim_dr_priority").String(), "1"))
	}
	if c.NvPairs.EnablePfc != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_pfc").String(), c.NvPairs.EnablePfc))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_pfc").String(), "false"))
	}
	if c.NvPairs.EnableQos != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_qos").String(), c.NvPairs.EnableQos))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_qos").String(), "false"))
	}
	if c.NvPairs.QosPolicy != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("qos_policy").String(), c.NvPairs.QosPolicy))
	}
	if c.NvPairs.QueuingPolicy != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("queuing_policy").String(), c.NvPairs.QueuingPolicy))
	}
	if c.NvPairs.LinkStateRoutingProtocol != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("link_state_routing_protocol").String(), c.NvPairs.LinkStateRoutingProtocol))
	}
	if c.NvPairs.LinkStateRoutingTag != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("link_state_routing_tag").String(), c.NvPairs.LinkStateRoutingTag))
	}
	if c.NvPairs.Ipv6Addr != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ipv6_addr").String(), c.NvPairs.Ipv6Addr))
	}
	if c.NvPairs.Ipv6PrefixLength != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ipv6_prefix_length").String(), c.NvPairs.Ipv6PrefixLength))
	}
	if c.NvPairs.CdpEnable != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("cdp_enable").String(), c.NvPairs.CdpEnable))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("cdp_enable").String(), "false"))
	}
	if c.NvPairs.PortDuplexMode != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("port_duplex_mode").String(), c.NvPairs.PortDuplexMode))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("port_duplex_mode").String(), "auto"))
	}
	if c.NvPairs.EnableMonitor != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_monitor").String(), c.NvPairs.EnableMonitor))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_monitor").String(), "false"))
	}
	if c.NvPairs.PvlanMode != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("pvlan_mode").String(), c.NvPairs.PvlanMode))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("pvlan_mode").String(), "host"))
	}
	if c.NvPairs.PvlanAllowedVlans != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("pvlan_allowed_vlans").String(), c.NvPairs.PvlanAllowedVlans))
	}
	if c.NvPairs.PvlanNativeVlan != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("pvlan_native_vlan").String(), c.NvPairs.PvlanNativeVlan))
	}
	if c.NvPairs.IgForFex != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ig_for_fex").String(), c.NvPairs.IgForFex))
	}
	if c.NvPairs.AutoNegotiate != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("auto_negotiate").String(), c.NvPairs.AutoNegotiate))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("auto_negotiate").String(), "on"))
	}
	if c.NvPairs.PathCost != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("path_cost").String(), strconv.Itoa(int(*c.NvPairs.PathCost))))
	}
	if c.NvPairs.GuardMode != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("guard_mode").String(), c.NvPairs.GuardMode))
	}
	if c.NvPairs.Ttag != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ttag").String(), c.NvPairs.Ttag))
	}
	if c.DeploymentStatus != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("deployment_status").String(), c.DeploymentStatus))
	}
	return ret
}
