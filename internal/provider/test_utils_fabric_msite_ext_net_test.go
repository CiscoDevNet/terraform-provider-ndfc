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

func FabricMsiteExtNetModelHelperStateCheck(RscName string, c resource_fabric_common.NDFCFabricCommonModel, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.FabricName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_name").String(), c.FabricName))
	}
	if c.AaaRemoteIpEnabled != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("aaa_remote_ip_enabled").String(), c.AaaRemoteIpEnabled))
	}
	if c.AaaServerConf != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("aaa_server_conf").String(), c.AaaServerConf))
	}
	if c.AllowNxc != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("allow_nxc").String(), c.AllowNxc))
	}
	if c.AllowNxcPrev != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("allow_nxc_prev").String(), c.AllowNxcPrev))
	}
	if c.BgpAs != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bgp_as").String(), c.BgpAs))
	}
	if c.BootstrapConf != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bootstrap_conf").String(), c.BootstrapConf))
	}
	if c.BootstrapConfXe != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bootstrap_conf_xe").String(), c.BootstrapConfXe))
	}
	if c.BootstrapEnable != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bootstrap_enable").String(), c.BootstrapEnable))
	}
	if c.BootstrapMultisubnet != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bootstrap_multisubnet").String(), c.BootstrapMultisubnet))
	}
	if c.BootstrapMultisubnetInternal != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bootstrap_multisubnet_internal").String(), c.BootstrapMultisubnetInternal))
	}
	if c.CdpEnable != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("cdp_enable").String(), c.CdpEnable))
	}
	if c.DciSubnetRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dci_subnet_range").String(), c.DciSubnetRange))
	}
	if c.DciSubnetTargetMask != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dci_subnet_target_mask").String(), strconv.Itoa(int(*c.DciSubnetTargetMask))))
	}
	if c.DeploymentFreeze != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("deployment_freeze").String(), c.DeploymentFreeze))
	}
	if c.DhcpEnable != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dhcp_enable").String(), c.DhcpEnable))
	}
	if c.DhcpEnd != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dhcp_end").String(), c.DhcpEnd))
	}
	if c.DhcpEndInternal != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dhcp_end_internal").String(), c.DhcpEndInternal))
	}
	if c.DhcpIpv6Enable != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dhcp_ipv6_enable").String(), c.DhcpIpv6Enable))
	}
	if c.DhcpIpv6EnableInternal != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dhcp_ipv6_enable_internal").String(), c.DhcpIpv6EnableInternal))
	}
	if c.DhcpStart != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dhcp_start").String(), c.DhcpStart))
	}
	if c.DhcpStartInternal != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dhcp_start_internal").String(), c.DhcpStartInternal))
	}
	if c.DomainName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("domain_name").String(), c.DomainName))
	}
	if c.DomainNameInternal != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("domain_name_internal").String(), c.DomainNameInternal))
	}
	if c.EnableAaa != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_aaa").String(), c.EnableAaa))
	}
	if c.EnableNetflow != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_netflow").String(), c.EnableNetflow))
	}
	if c.EnableNetflowPrev != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_netflow_prev").String(), c.EnableNetflowPrev))
	}
	if c.EnableNxapi != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_nxapi").String(), c.EnableNxapi))
	}
	if c.EnableNxapiHttp != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_nxapi_http").String(), c.EnableNxapiHttp))
	}
	if c.EnableRtIntfStats != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_rt_intf_stats").String(), c.EnableRtIntfStats))
	}
	if c.ExtFabricType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ext_fabric_type").String(), c.ExtFabricType))
	}
	if c.FabricFreeform != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_freeform").String(), c.FabricFreeform))
	}
	if c.FabricType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_type").String(), c.FabricType))
	}
	if c.FeaturePtp != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("feature_ptp").String(), c.FeaturePtp))
	}
	if c.FeaturePtpInternal != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("feature_ptp_internal").String(), c.FeaturePtpInternal))
	}
	if c.Ff != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ff").String(), c.Ff))
	}
	if c.InbandEnable != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("inband_enable").String(), c.InbandEnable))
	}
	if c.InbandEnablePrev != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("inband_enable_prev").String(), c.InbandEnablePrev))
	}
	if c.InbandMgmt != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("inband_mgmt").String(), c.InbandMgmt))
	}
	if c.InbandMgmtPrev != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("inband_mgmt_prev").String(), c.InbandMgmtPrev))
	}
	if c.IntfStatLoadInterval != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("intf_stat_load_interval").String(), strconv.Itoa(int(*c.IntfStatLoadInterval))))
	}
	if c.IsReadOnly != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("is_read_only").String(), c.IsReadOnly))
	}
	if c.Loopback0IpRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("loopback0_ip_range").String(), c.Loopback0IpRange))
	}
	if c.MgmtGw != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mgmt_gw").String(), c.MgmtGw))
	}
	if c.MgmtGwInternal != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mgmt_gw_internal").String(), c.MgmtGwInternal))
	}
	if c.MgmtPrefix != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mgmt_prefix").String(), strconv.Itoa(int(*c.MgmtPrefix))))
	}
	if c.MgmtPrefixInternal != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mgmt_prefix_internal").String(), strconv.Itoa(int(*c.MgmtPrefixInternal))))
	}
	if c.MgmtV6prefix != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mgmt_v6prefix").String(), strconv.Itoa(int(*c.MgmtV6prefix))))
	}
	if c.MgmtV6prefixInternal != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mgmt_v6prefix_internal").String(), strconv.Itoa(int(*c.MgmtV6prefixInternal))))
	}
	if c.MplsHandoff != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mpls_handoff").String(), c.MplsHandoff))
	}
	if c.MplsLbId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mpls_lb_id").String(), strconv.Itoa(int(*c.MplsLbId))))
	}
	if c.MplsLoopbackIpRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mpls_loopback_ip_range").String(), c.MplsLoopbackIpRange))
	}
	if c.MsoConnectivityDeployed != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mso_connectivity_deployed").String(), c.MsoConnectivityDeployed))
	}
	if c.MsoControlerId != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mso_controler_id").String(), c.MsoControlerId))
	}
	if c.MsoSiteGroupName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mso_site_group_name").String(), c.MsoSiteGroupName))
	}
	if c.MsoSiteId != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mso_site_id").String(), c.MsoSiteId))
	}
	if c.NetflowExporterList != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("netflow_exporter_list").String(), c.NetflowExporterList))
	}
	if c.NetflowMonitorList != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("netflow_monitor_list").String(), c.NetflowMonitorList))
	}
	if c.NetflowRecordList != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("netflow_record_list").String(), c.NetflowRecordList))
	}
	if c.NetflowSamplerList != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("netflow_sampler_list").String(), c.NetflowSamplerList))
	}
	if c.NxapiHttpsPort != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("nxapi_https_port").String(), strconv.Itoa(int(*c.NxapiHttpsPort))))
	}
	if c.NxapiHttpPort != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("nxapi_http_port").String(), strconv.Itoa(int(*c.NxapiHttpPort))))
	}
	if c.NxcDestVrf != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("nxc_dest_vrf").String(), c.NxcDestVrf))
	}
	if c.NxcProxyPort != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("nxc_proxy_port").String(), strconv.Itoa(int(*c.NxcProxyPort))))
	}
	if c.NxcProxyServer != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("nxc_proxy_server").String(), c.NxcProxyServer))
	}
	if c.NxcSrcIntf != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("nxc_src_intf").String(), c.NxcSrcIntf))
	}
	if c.OverwriteGlobalNxc != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("overwrite_global_nxc").String(), c.OverwriteGlobalNxc))
	}
	if c.PmEnable != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("pm_enable").String(), c.PmEnable))
	}
	if c.PmEnablePrev != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("pm_enable_prev").String(), c.PmEnablePrev))
	}
	if c.PnpEnable != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("pnp_enable").String(), c.PnpEnable))
	}
	if c.PnpEnableInternal != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("pnp_enable_internal").String(), c.PnpEnableInternal))
	}
	if c.PowerRedundancyMode != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("power_redundancy_mode").String(), c.PowerRedundancyMode))
	}
	if c.PremsoParentFabric != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("premso_parent_fabric").String(), c.PremsoParentFabric))
	}
	if c.PtpDomainId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ptp_domain_id").String(), strconv.Itoa(int(*c.PtpDomainId))))
	}
	if c.PtpLbId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ptp_lb_id").String(), strconv.Itoa(int(*c.PtpLbId))))
	}
	if c.SnmpServerHostTrap != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("snmp_server_host_trap").String(), c.SnmpServerHostTrap))
	}
	if c.SubinterfaceRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("subinterface_range").String(), c.SubinterfaceRange))
	}
	if c.EnableRealTimeBackup != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_real_time_backup").String(), c.EnableRealTimeBackup))
	}
	if c.EnableScheduledBackup != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_scheduled_backup").String(), c.EnableScheduledBackup))
	}
	if c.ScheduledTime != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("scheduled_time").String(), c.ScheduledTime))
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
