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
	"terraform-provider-ndfc/internal/provider/resources/resource_fabric_common"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func FabricVxlanMsdModelHelperStateCheck(RscName string, c resource_fabric_common.NDFCFabricCommonModel, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.FabricName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_name").String(), c.FabricName))
	}
	if c.AnycastGwMac != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("anycast_gw_mac").String(), c.AnycastGwMac))
	}
	if c.BgpRpAsn != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bgp_rp_asn").String(), c.BgpRpAsn))
	}
	if c.BgwRoutingTag != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bgw_routing_tag").String(), strconv.Itoa(int(*c.BgwRoutingTag))))
	}
	if c.BgwRoutingTagPrev != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bgw_routing_tag_prev").String(), c.BgwRoutingTagPrev))
	}
	if c.BorderGwyConnections != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("border_gwy_connections").String(), c.BorderGwyConnections))
	}
	if c.CloudsecAlgorithm != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("cloudsec_algorithm").String(), c.CloudsecAlgorithm))
	}
	if c.CloudsecAutoconfig != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("cloudsec_autoconfig").String(), c.CloudsecAutoconfig))
	}
	if c.CloudsecEnforcement != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("cloudsec_enforcement").String(), c.CloudsecEnforcement))
	}
	if c.CloudsecKeyString != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("cloudsec_key_string").String(), c.CloudsecKeyString))
	}
	if c.CloudsecReportTimer != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("cloudsec_report_timer").String(), strconv.Itoa(int(*c.CloudsecReportTimer))))
	}
	if c.DciSubnetRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dci_subnet_range").String(), c.DciSubnetRange))
	}
	if c.DciSubnetTargetMask != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dci_subnet_target_mask").String(), strconv.Itoa(int(*c.DciSubnetTargetMask))))
	}
	if c.DcnmId != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dcnm_id").String(), c.DcnmId))
	}
	if c.DelayRestore != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("delay_restore").String(), strconv.Itoa(int(*c.DelayRestore))))
	}
	if c.EnableBgpBfd != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_bgp_bfd").String(), c.EnableBgpBfd))
	}
	if c.EnableBgpLogNeighborChange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_bgp_log_neighbor_change").String(), c.EnableBgpLogNeighborChange))
	}
	if c.EnableBgpSendComm != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_bgp_send_comm").String(), c.EnableBgpSendComm))
	}
	if c.EnablePvlan != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_pvlan").String(), c.EnablePvlan))
	}
	if c.EnablePvlanPrev != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_pvlan_prev").String(), c.EnablePvlanPrev))
	}
	if c.EnableRsRedistDirect != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_rs_redist_direct").String(), c.EnableRsRedistDirect))
	}
	if c.EnableSgt != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_sgt").String(), c.EnableSgt))
	}
	if c.EnableSgtPrev != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_sgt_prev").String(), c.EnableSgtPrev))
	}
	if c.EnableTrmTrmv6 != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_trm_trmv6").String(), c.EnableTrmTrmv6))
	}
	if c.EnableTrmTrmv6Prev != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_trm_trmv6_prev").String(), c.EnableTrmTrmv6Prev))
	}
	if c.ExtFabricType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ext_fabric_type").String(), c.ExtFabricType))
	}
	if c.FabricType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_type").String(), c.FabricType))
	}
	if c.Ff != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ff").String(), c.Ff))
	}
	if c.L2SegmentIdRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l2_segment_id_range").String(), c.L2SegmentIdRange))
	}
	if c.L3PartitionIdRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l3_partition_id_range").String(), c.L3PartitionIdRange))
	}
	if c.Loopback100Ipv6Range != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("loopback100_ipv6_range").String(), c.Loopback100Ipv6Range))
	}
	if c.Loopback100IpRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("loopback100_ip_range").String(), c.Loopback100IpRange))
	}
	if c.MsoControlerId != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mso_controler_id").String(), c.MsoControlerId))
	}
	if c.MsoSiteGroupName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mso_site_group_name").String(), c.MsoSiteGroupName))
	}
	if c.MsIfcBgpAuthKeyType != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ms_ifc_bgp_auth_key_type").String(), strconv.Itoa(int(*c.MsIfcBgpAuthKeyType))))
	}
	if c.MsIfcBgpAuthKeyTypePrev != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ms_ifc_bgp_auth_key_type_prev").String(), strconv.Itoa(int(*c.MsIfcBgpAuthKeyTypePrev))))
	}
	if c.MsIfcBgpPassword != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ms_ifc_bgp_password").String(), c.MsIfcBgpPassword))
	}
	if c.MsIfcBgpPasswordEnable != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ms_ifc_bgp_password_enable").String(), c.MsIfcBgpPasswordEnable))
	}
	if c.MsIfcBgpPasswordEnablePrev != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ms_ifc_bgp_password_enable_prev").String(), c.MsIfcBgpPasswordEnablePrev))
	}
	if c.MsIfcBgpPasswordPrev != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ms_ifc_bgp_password_prev").String(), c.MsIfcBgpPasswordPrev))
	}
	if c.MsLoopbackId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ms_loopback_id").String(), strconv.Itoa(int(*c.MsLoopbackId))))
	}
	if c.MsUnderlayAutoconfig != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ms_underlay_autoconfig").String(), c.MsUnderlayAutoconfig))
	}
	if c.ParentOnemanageFabric != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("parent_onemanage_fabric").String(), c.ParentOnemanageFabric))
	}
	if c.PremsoParentFabric != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("premso_parent_fabric").String(), c.PremsoParentFabric))
	}
	if c.RpServerIp != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("rp_server_ip").String(), c.RpServerIp))
	}
	if c.RsRoutingTag != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("rs_routing_tag").String(), strconv.Itoa(int(*c.RsRoutingTag))))
	}
	if c.SgtIdRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("sgt_id_range").String(), c.SgtIdRange))
	}
	if c.SgtIdRangePrev != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("sgt_id_range_prev").String(), c.SgtIdRangePrev))
	}
	if c.SgtNamePrefix != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("sgt_name_prefix").String(), c.SgtNamePrefix))
	}
	if c.SgtNamePrefixPrev != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("sgt_name_prefix_prev").String(), c.SgtNamePrefixPrev))
	}
	if c.SgtOperStatus != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("sgt_oper_status").String(), c.SgtOperStatus))
	}
	if c.SgtPreprovision != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("sgt_preprovision").String(), c.SgtPreprovision))
	}
	if c.SgtPreprovisionPrev != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("sgt_preprovision_prev").String(), c.SgtPreprovisionPrev))
	}
	if c.SgtPreprovRecalcStatus != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("sgt_preprov_recalc_status").String(), c.SgtPreprovRecalcStatus))
	}
	if c.SgtRecalcStatus != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("sgt_recalc_status").String(), c.SgtRecalcStatus))
	}
	if c.TorAutoDeploy != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("tor_auto_deploy").String(), c.TorAutoDeploy))
	}
	if c.V6DciSubnetRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("v6_dci_subnet_range").String(), c.V6DciSubnetRange))
	}
	if c.V6DciSubnetTargetMask != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("v6_dci_subnet_target_mask").String(), strconv.Itoa(int(*c.V6DciSubnetTargetMask))))
	}
	if c.VxlanUnderlayIsV6 != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vxlan_underlay_is_v6").String(), c.VxlanUnderlayIsV6))
	}
	if c.DefaultNetwork != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("default_network").String(), c.DefaultNetwork))
	}
	if c.DefaultPvlanSecNetwork != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("default_pvlan_sec_network").String(), c.DefaultPvlanSecNetwork))
	}
	if c.DefaultVrf != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("default_vrf").String(), c.DefaultVrf))
	}
	if c.EnableScheduledBackup != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_scheduled_backup").String(), c.EnableScheduledBackup))
	}
	if c.NetworkExtensionTemplate != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("network_extension_template").String(), c.NetworkExtensionTemplate))
	}
	if c.ScheduledTime != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("scheduled_time").String(), c.ScheduledTime))
	}
	if c.VrfExtensionTemplate != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_extension_template").String(), c.VrfExtensionTemplate))
	}
	if c.Deploy {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("deploy").String(), "true"))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("deploy").String(), "false"))
	}
	if c.DeploymentStatus != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("deployment_status").String(), c.DeploymentStatus))
	}
	return ret
}
