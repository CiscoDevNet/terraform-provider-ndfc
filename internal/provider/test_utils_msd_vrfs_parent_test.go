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
	"terraform-provider-ndfc/internal/provider/resources/resource_msd_vrfs_parent"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func MsdVrfsParentModelHelperStateCheck(RscName string, c resource_msd_vrfs_parent.NDFCMsdVrfsParentModel, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.FabricName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_name").String(), c.FabricName))
	}
	for key, value := range c.Vrfs {
		attrNewPath := attrPath.AtName("vrfs").AtName(key)
		ret = append(ret, VrfsValueHelperStateCheck(RscName, value, attrNewPath)...)
	}
	return ret
}

func VrfsValueHelperStateCheck(RscName string, c resource_msd_vrfs_parent.NDFCVrfsValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.VrfTemplate != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_template").String(), c.VrfTemplate))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_template").String(), "Default_VRF_Universal"))
	}
	if c.VrfExtensionTemplate != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_extension_template").String(), c.VrfExtensionTemplate))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_extension_template").String(), "Default_VRF_Extension_Universal"))
	}
	if c.VrfId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_id").String(), strconv.Itoa(int(*c.VrfId))))
	}
	if c.VrfTemplateConfig.VlanId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vlan_id").String(), strconv.Itoa(int(*c.VrfTemplateConfig.VlanId))))
	}
	if c.VrfTemplateConfig.VlanName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vlan_name").String(), c.VrfTemplateConfig.VlanName))
	}
	if c.VrfTemplateConfig.InterfaceDescription != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_description").String(), c.VrfTemplateConfig.InterfaceDescription))
	}
	if c.VrfTemplateConfig.VrfDescription != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_description").String(), c.VrfTemplateConfig.VrfDescription))
	}
	if c.VrfTemplateConfig.Mtu != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mtu").String(), strconv.Itoa(int(*c.VrfTemplateConfig.Mtu))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mtu").String(), "9216"))
	}
	if c.VrfStatus != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_status").String(), c.VrfStatus))
	}
	if c.VrfTemplateConfig.LoopbackRoutingTag != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("loopback_routing_tag").String(), strconv.Itoa(int(*c.VrfTemplateConfig.LoopbackRoutingTag))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("loopback_routing_tag").String(), "12345"))
	}
	if c.VrfTemplateConfig.RedistributeDirectRouteMap != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("redistribute_direct_route_map").String(), c.VrfTemplateConfig.RedistributeDirectRouteMap))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("redistribute_direct_route_map").String(), "FABRIC-RMAP-REDIST-SUBNET"))
	}
	if c.VrfTemplateConfig.MaxBgpPaths != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("max_bgp_paths").String(), strconv.Itoa(int(*c.VrfTemplateConfig.MaxBgpPaths))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("max_bgp_paths").String(), "1"))
	}
	if c.VrfTemplateConfig.MaxIbgpPaths != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("max_ibgp_paths").String(), strconv.Itoa(int(*c.VrfTemplateConfig.MaxIbgpPaths))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("max_ibgp_paths").String(), "2"))
	}
	if c.VrfTemplateConfig.Ipv6LinkLocal != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ipv6_link_local").String(), c.VrfTemplateConfig.Ipv6LinkLocal))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ipv6_link_local").String(), "true"))
	}
	if c.VrfTemplateConfig.Trm != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("trm").String(), c.VrfTemplateConfig.Trm))
	}
	if c.VrfTemplateConfig.RpExternal != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("rp_external").String(), c.VrfTemplateConfig.RpExternal))
	}
	if c.VrfTemplateConfig.MvpnInterAs != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mvpn_inter_as").String(), c.VrfTemplateConfig.MvpnInterAs))
	}
	if c.VrfTemplateConfig.DisableRtAuto != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("disable_rt_auto").String(), c.VrfTemplateConfig.DisableRtAuto))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("disable_rt_auto").String(), "false"))
	}
	if c.VrfTemplateConfig.RouteTargetImport != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("route_target_import").String(), c.VrfTemplateConfig.RouteTargetImport))
	}
	if c.VrfTemplateConfig.RouteTargetExport != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("route_target_export").String(), c.VrfTemplateConfig.RouteTargetExport))
	}
	if c.VrfTemplateConfig.RouteTargetImportEvpn != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("route_target_import_evpn").String(), c.VrfTemplateConfig.RouteTargetImportEvpn))
	}
	if c.VrfTemplateConfig.RouteTargetExportEvpn != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("route_target_export_evpn").String(), c.VrfTemplateConfig.RouteTargetExportEvpn))
	}
	if c.VrfTemplateConfig.RouteTargetImportCloudEvpn != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("route_target_import_cloud_evpn").String(), c.VrfTemplateConfig.RouteTargetImportCloudEvpn))
	}
	if c.VrfTemplateConfig.RouteTargetExportCloudEvpn != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("route_target_export_cloud_evpn").String(), c.VrfTemplateConfig.RouteTargetExportCloudEvpn))
	}
	return ret
}
