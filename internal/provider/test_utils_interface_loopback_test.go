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
	"terraform-provider-ndfc/internal/provider/resources/resource_interface_common"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func InterfaceLoopbackModelHelperStateCheck(RscName string, c resource_interface_common.NDFCInterfaceCommonModel, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.Policy != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("policy").String(), c.Policy))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("policy").String(), "int_loopback"))
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
		ret = append(ret, InterfaceLoopbackInterfacesValueHelperStateCheck(RscName, value, attrNewPath)...)
	}
	return ret
}

func InterfaceLoopbackInterfacesValueHelperStateCheck(RscName string, c resource_interface_common.NDFCInterfacesValue, attrPath path.Path) []resource.TestCheckFunc {
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
	if c.NvPairs.FreeformConfig != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("freeform_config").String(), c.NvPairs.FreeformConfig))
	}
	if c.NvPairs.InterfaceDescription != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_description").String(), c.NvPairs.InterfaceDescription))
	}
	if c.NvPairs.Vrf != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf").String(), c.NvPairs.Vrf))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf").String(), "default"))
	}
	if c.NvPairs.Ipv4Address != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ipv4_address").String(), c.NvPairs.Ipv4Address))
	}
	if c.NvPairs.Ipv6Address != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ipv6_address").String(), c.NvPairs.Ipv6Address))
	}
	if c.NvPairs.RouteMapTag != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("route_map_tag").String(), c.NvPairs.RouteMapTag))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("route_map_tag").String(), "12345"))
	}
	if c.DeploymentStatus != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("deployment_status").String(), c.DeploymentStatus))
	}
	return ret
}
