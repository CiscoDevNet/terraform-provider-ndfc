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
	"terraform-provider-ndfc/internal/provider/resources/resource_vrf_bulk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func VrfBulkModelHelperStateCheck(RscName string, c resource_vrf_bulk.NDFCVrfBulkModel, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.FabricName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_name").String(), c.FabricName))
	}
	if c.DeployAllAttachments {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("deploy_all_attachments").String(), "true"))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("deploy_all_attachments").String(), "false"))
	}
	for key, value := range c.Vrfs {
		attrNewPath := attrPath.AtName("vrfs").AtName(key)
		ret = append(ret, VrfsValueHelperStateCheck(RscName, value, attrNewPath)...)
	}
	return ret
}

func VrfsValueHelperStateCheck(RscName string, c resource_vrf_bulk.NDFCVrfsValue, attrPath path.Path) []resource.TestCheckFunc {
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
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("trm").String(), "false"))
	}
	if c.VrfTemplateConfig.NoRp != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("no_rp").String(), c.VrfTemplateConfig.NoRp))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("no_rp").String(), "false"))
	}
	if c.VrfTemplateConfig.RpExternal != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("rp_external").String(), c.VrfTemplateConfig.RpExternal))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("rp_external").String(), "false"))
	}
	if c.VrfTemplateConfig.RpAddress != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("rp_address").String(), c.VrfTemplateConfig.RpAddress))
	}
	if c.VrfTemplateConfig.RpLoopbackId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("rp_loopback_id").String(), strconv.Itoa(int(*c.VrfTemplateConfig.RpLoopbackId))))
	}
	if c.VrfTemplateConfig.UnderlayMulticastAddress != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("underlay_multicast_address").String(), c.VrfTemplateConfig.UnderlayMulticastAddress))
	}
	if c.VrfTemplateConfig.OverlayMulticastGroups != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("overlay_multicast_groups").String(), c.VrfTemplateConfig.OverlayMulticastGroups))
	}
	if c.VrfTemplateConfig.MvpnInterAs != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mvpn_inter_as").String(), c.VrfTemplateConfig.MvpnInterAs))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mvpn_inter_as").String(), "false"))
	}
	if c.VrfTemplateConfig.TrmBgwMsite != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("trm_bgw_msite").String(), c.VrfTemplateConfig.TrmBgwMsite))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("trm_bgw_msite").String(), "false"))
	}
	if c.VrfTemplateConfig.AdvertiseHostRoutes != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("advertise_host_routes").String(), c.VrfTemplateConfig.AdvertiseHostRoutes))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("advertise_host_routes").String(), "false"))
	}
	if c.VrfTemplateConfig.AdvertiseDefaultRoute != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("advertise_default_route").String(), c.VrfTemplateConfig.AdvertiseDefaultRoute))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("advertise_default_route").String(), "true"))
	}
	if c.VrfTemplateConfig.ConfigureStaticDefaultRoute != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("configure_static_default_route").String(), c.VrfTemplateConfig.ConfigureStaticDefaultRoute))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("configure_static_default_route").String(), "true"))
	}
	if c.VrfTemplateConfig.BgpPassword != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bgp_password").String(), c.VrfTemplateConfig.BgpPassword))
	}
	if c.VrfTemplateConfig.BgpPasswordType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bgp_password_type").String(), c.VrfTemplateConfig.BgpPasswordType))
	}
	if c.VrfTemplateConfig.Netflow != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("netflow").String(), c.VrfTemplateConfig.Netflow))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("netflow").String(), "false"))
	}
	if c.VrfTemplateConfig.NetflowMonitor != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("netflow_monitor").String(), c.VrfTemplateConfig.NetflowMonitor))
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
	if c.VrfTemplateConfig.RouteTargetImportMvpn != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("route_target_import_mvpn").String(), c.VrfTemplateConfig.RouteTargetImportMvpn))
	}
	if c.VrfTemplateConfig.RouteTargetExportMvpn != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("route_target_export_mvpn").String(), c.VrfTemplateConfig.RouteTargetExportMvpn))
	}
	if c.VrfTemplateConfig.RouteTargetImportCloudEvpn != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("route_target_import_cloud_evpn").String(), c.VrfTemplateConfig.RouteTargetImportCloudEvpn))
	}
	if c.VrfTemplateConfig.RouteTargetExportCloudEvpn != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("route_target_export_cloud_evpn").String(), c.VrfTemplateConfig.RouteTargetExportCloudEvpn))
	}
	if c.DeployAttachments {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("deploy_attachments").String(), "true"))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("deploy_attachments").String(), "false"))
	}
	for key, value := range c.AttachList {
		attrNewPath := attrPath.AtName("attach_list").AtName(key)
		ret = append(ret, AttachListValueHelperStateCheck(RscName, value, attrNewPath)...)
	}
	return ret
}
