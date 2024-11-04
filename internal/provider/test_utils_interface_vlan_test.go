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
	"terraform-provider-ndfc/internal/provider/resources/resource_interface_common"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func InterfaceVlanModelHelperStateCheck(RscName string, c resource_interface_common.NDFCInterfaceCommonModel, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.Policy != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("policy").String(), c.Policy))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("policy").String(), "int_vlan"))
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
		ret = append(ret, InterfaceVlanInterfacesValueHelperStateCheck(RscName, value, attrNewPath)...)
	}
	return ret
}

func InterfaceVlanInterfacesValueHelperStateCheck(RscName string, c resource_interface_common.NDFCInterfacesValue, attrPath path.Path) []resource.TestCheckFunc {
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
	if c.NvPairs.Mtu != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mtu").String(), c.NvPairs.Mtu))
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
	if c.NvPairs.Vrf != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf").String(), c.NvPairs.Vrf))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf").String(), "default"))
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
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("disable_ip_redirects").String(), "true"))
	}
	if c.NvPairs.EnableHsrp != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_hsrp").String(), c.NvPairs.EnableHsrp))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_hsrp").String(), "false"))
	}
	if c.NvPairs.HsrpGroup != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("hsrp_group").String(), strconv.Itoa(int(*c.NvPairs.HsrpGroup))))
	}
	if c.NvPairs.HsrpVip != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("hsrp_vip").String(), c.NvPairs.HsrpVip))
	}
	if c.NvPairs.HsrpPriority != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("hsrp_priority").String(), strconv.Itoa(int(*c.NvPairs.HsrpPriority))))
	}
	if c.NvPairs.HsrpVersion != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("hsrp_version").String(), c.NvPairs.HsrpVersion))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("hsrp_version").String(), "2"))
	}
	if c.NvPairs.Preempt != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("preempt").String(), c.NvPairs.Preempt))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("preempt").String(), "false"))
	}
	if c.NvPairs.Mac != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mac").String(), c.NvPairs.Mac))
	}
	if c.NvPairs.DhcpServerAddr1 != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dhcp_server_addr1").String(), c.NvPairs.DhcpServerAddr1))
	}
	if c.NvPairs.DhcpServerAddr2 != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dhcp_server_addr2").String(), c.NvPairs.DhcpServerAddr2))
	}
	if c.NvPairs.DhcpServerAddr3 != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dhcp_server_addr3").String(), c.NvPairs.DhcpServerAddr3))
	}
	if c.NvPairs.VrfDhcp1 != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_dhcp1").String(), c.NvPairs.VrfDhcp1))
	}
	if c.NvPairs.VrfDhcp2 != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_dhcp2").String(), c.NvPairs.VrfDhcp2))
	}
	if c.NvPairs.VrfDhcp3 != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_dhcp3").String(), c.NvPairs.VrfDhcp3))
	}
	if c.NvPairs.AdvertiseSubnetInUnderlay != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("advertise_subnet_in_underlay").String(), c.NvPairs.AdvertiseSubnetInUnderlay))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("advertise_subnet_in_underlay").String(), "false"))
	}
	if c.DeploymentStatus != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("deployment_status").String(), c.DeploymentStatus))
	}
	return ret
}
