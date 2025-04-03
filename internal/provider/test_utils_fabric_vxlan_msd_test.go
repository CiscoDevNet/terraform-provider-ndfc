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
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("anycast_gw_mac").String(), "2020.0000.00aa"))
	}
	if c.BgpRpAsn != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bgp_rp_asn").String(), c.BgpRpAsn))
	}
	if c.BgwRoutingTag != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bgw_routing_tag").String(), strconv.Itoa(int(*c.BgwRoutingTag))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bgw_routing_tag").String(), "54321"))
	}
	if c.BorderGwyConnections != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("border_gwy_connections").String(), c.BorderGwyConnections))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("border_gwy_connections").String(), "Manual"))
	}
	if c.CloudsecAlgorithm != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("cloudsec_algorithm").String(), c.CloudsecAlgorithm))
	}
	if c.CloudsecAutoconfig != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("cloudsec_autoconfig").String(), c.CloudsecAutoconfig))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("cloudsec_autoconfig").String(), "false"))
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
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dci_subnet_range").String(), "10.10.1.0/24"))
	}
	if c.DciSubnetTargetMask != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dci_subnet_target_mask").String(), strconv.Itoa(int(*c.DciSubnetTargetMask))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dci_subnet_target_mask").String(), "30"))
	}
	if c.DelayRestore != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("delay_restore").String(), strconv.Itoa(int(*c.DelayRestore))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("delay_restore").String(), "300"))
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
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_pvlan").String(), "false"))
	}
	if c.EnableRsRedistDirect != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_rs_redist_direct").String(), c.EnableRsRedistDirect))
	}
	if c.L2SegmentIdRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l2_segment_id_range").String(), c.L2SegmentIdRange))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l2_segment_id_range").String(), "30000-49000"))
	}
	if c.L3PartitionIdRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l3_partition_id_range").String(), c.L3PartitionIdRange))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l3_partition_id_range").String(), "50000-59000"))
	}
	if c.Loopback100IpRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("loopback100_ip_range").String(), c.Loopback100IpRange))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("loopback100_ip_range").String(), "10.10.0.0/24"))
	}
	if c.MsIfcBgpAuthKeyType != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ms_ifc_bgp_auth_key_type").String(), strconv.Itoa(int(*c.MsIfcBgpAuthKeyType))))
	}
	if c.MsIfcBgpPassword != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ms_ifc_bgp_password").String(), c.MsIfcBgpPassword))
	}
	if c.MsIfcBgpPasswordEnable != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ms_ifc_bgp_password_enable").String(), c.MsIfcBgpPasswordEnable))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ms_ifc_bgp_password_enable").String(), "false"))
	}
	if c.MsLoopbackId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ms_loopback_id").String(), strconv.Itoa(int(*c.MsLoopbackId))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ms_loopback_id").String(), "100"))
	}
	if c.MsUnderlayAutoconfig != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ms_underlay_autoconfig").String(), c.MsUnderlayAutoconfig))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ms_underlay_autoconfig").String(), "false"))
	}
	if c.RpServerIp != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("rp_server_ip").String(), c.RpServerIp))
	}
	if c.RsRoutingTag != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("rs_routing_tag").String(), strconv.Itoa(int(*c.RsRoutingTag))))
	}
	if c.TorAutoDeploy != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("tor_auto_deploy").String(), c.TorAutoDeploy))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("tor_auto_deploy").String(), "false"))
	}
	if c.DefaultNetwork != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("default_network").String(), c.DefaultNetwork))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("default_network").String(), "Default_Network_Universal"))
	}
	if c.DefaultPvlanSecNetwork != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("default_pvlan_sec_network").String(), c.DefaultPvlanSecNetwork))
	}
	if c.DefaultVrf != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("default_vrf").String(), c.DefaultVrf))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("default_vrf").String(), "Default_VRF_Universal"))
	}
	if c.EnableScheduledBackup != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_scheduled_backup").String(), c.EnableScheduledBackup))
	}
	if c.NetworkExtensionTemplate != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("network_extension_template").String(), c.NetworkExtensionTemplate))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("network_extension_template").String(), "Default_Network_Extension_Universal"))
	}
	if c.ScheduledTime != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("scheduled_time").String(), c.ScheduledTime))
	}
	if c.VrfExtensionTemplate != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_extension_template").String(), c.VrfExtensionTemplate))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_extension_template").String(), "Default_VRF_Extension_Universal"))
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
