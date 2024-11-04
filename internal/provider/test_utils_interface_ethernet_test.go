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
	if c.DeploymentStatus != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("deployment_status").String(), c.DeploymentStatus))
	}
	return ret
}
