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

func FabricIpfmModelHelperStateCheck(RscName string, c resource_fabric_common.NDFCFabricCommonModel, attrPath path.Path) []resource.TestCheckFunc {
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
	if c.ActiveMigration != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("active_migration").String(), c.ActiveMigration))
	}
	if c.AgentIntf != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("agent_intf").String(), c.AgentIntf))
	}
	if c.AsmGroupRanges != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("asm_group_ranges").String(), c.AsmGroupRanges))
	}
	if c.BootstrapConf != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bootstrap_conf").String(), c.BootstrapConf))
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
	if c.BrfieldDebugFlag != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("brfield_debug_flag").String(), c.BrfieldDebugFlag))
	}
	if c.CdpEnable != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("cdp_enable").String(), c.CdpEnable))
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
	if c.DnsServerIpList != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dns_server_ip_list").String(), c.DnsServerIpList))
	}
	if c.DnsServerVrf != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dns_server_vrf").String(), c.DnsServerVrf))
	}
	if c.EnableAaa != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_aaa").String(), c.EnableAaa))
	}
	if c.EnableAgent != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_agent").String(), c.EnableAgent))
	}
	if c.EnableAsm != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_asm").String(), c.EnableAsm))
	}
	if c.EnableNbmPassive != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_nbm_passive").String(), c.EnableNbmPassive))
	}
	if c.EnableNbmPassivePrev != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_nbm_passive_prev").String(), c.EnableNbmPassivePrev))
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
	if c.ExtraConfIntraLinks != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("extra_conf_intra_links").String(), c.ExtraConfIntraLinks))
	}
	if c.ExtraConfLeaf != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("extra_conf_leaf").String(), c.ExtraConfLeaf))
	}
	if c.ExtraConfSpine != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("extra_conf_spine").String(), c.ExtraConfSpine))
	}
	if c.ExtFabricType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ext_fabric_type").String(), c.ExtFabricType))
	}
	if c.FabricInterfaceType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_interface_type").String(), c.FabricInterfaceType))
	}
	if c.FabricMtu != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_mtu").String(), strconv.Itoa(int(*c.FabricMtu))))
	}
	if c.FabricMtuPrev != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_mtu_prev").String(), strconv.Itoa(int(*c.FabricMtuPrev))))
	}
	if c.FabricTechnology != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_technology").String(), c.FabricTechnology))
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
	if c.GrfieldDebugFlag != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("grfield_debug_flag").String(), c.GrfieldDebugFlag))
	}
	if c.InterfaceEthernetDefaultPolicy != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_ethernet_default_policy").String(), c.InterfaceEthernetDefaultPolicy))
	}
	if c.InterfaceLoopbackDefaultPolicy != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_loopback_default_policy").String(), c.InterfaceLoopbackDefaultPolicy))
	}
	if c.InterfacePortChannelDefaultPolicy != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_port_channel_default_policy").String(), c.InterfacePortChannelDefaultPolicy))
	}
	if c.InterfaceVlanDefaultPolicy != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_vlan_default_policy").String(), c.InterfaceVlanDefaultPolicy))
	}
	if c.IntfStatLoadInterval != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("intf_stat_load_interval").String(), strconv.Itoa(int(*c.IntfStatLoadInterval))))
	}
	if c.IsisAuthEnable != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("isis_auth_enable").String(), c.IsisAuthEnable))
	}
	if c.IsisAuthKey != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("isis_auth_key").String(), c.IsisAuthKey))
	}
	if c.IsisAuthKeychainKeyId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("isis_auth_keychain_key_id").String(), strconv.Itoa(int(*c.IsisAuthKeychainKeyId))))
	}
	if c.IsisAuthKeychainName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("isis_auth_keychain_name").String(), c.IsisAuthKeychainName))
	}
	if c.IsisLevel != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("isis_level").String(), c.IsisLevel))
	}
	if c.IsisP2pEnable != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("isis_p2p_enable").String(), c.IsisP2pEnable))
	}
	if c.L2HostIntfMtu != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l2_host_intf_mtu").String(), strconv.Itoa(int(*c.L2HostIntfMtu))))
	}
	if c.L2HostIntfMtuPrev != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l2_host_intf_mtu_prev").String(), strconv.Itoa(int(*c.L2HostIntfMtuPrev))))
	}
	if c.LinkStateRouting != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("link_state_routing").String(), c.LinkStateRouting))
	}
	if c.LinkStateRoutingTag != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("link_state_routing_tag").String(), c.LinkStateRoutingTag))
	}
	if c.LinkStateRoutingTagPrev != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("link_state_routing_tag_prev").String(), c.LinkStateRoutingTagPrev))
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
	if c.NtpServerIpList != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ntp_server_ip_list").String(), c.NtpServerIpList))
	}
	if c.NtpServerVrf != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ntp_server_vrf").String(), c.NtpServerVrf))
	}
	if c.NxapiHttpsPort != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("nxapi_https_port").String(), strconv.Itoa(int(*c.NxapiHttpsPort))))
	}
	if c.NxapiHttpPort != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("nxapi_http_port").String(), strconv.Itoa(int(*c.NxapiHttpPort))))
	}
	if c.NxapiVrf != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("nxapi_vrf").String(), c.NxapiVrf))
	}
	if c.OspfAreaId != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ospf_area_id").String(), c.OspfAreaId))
	}
	if c.OspfAuthEnable != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ospf_auth_enable").String(), c.OspfAuthEnable))
	}
	if c.OspfAuthKey != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ospf_auth_key").String(), c.OspfAuthKey))
	}
	if c.OspfAuthKeyId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ospf_auth_key_id").String(), strconv.Itoa(int(*c.OspfAuthKeyId))))
	}
	if c.PimHelloAuthEnable != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("pim_hello_auth_enable").String(), c.PimHelloAuthEnable))
	}
	if c.PimHelloAuthKey != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("pim_hello_auth_key").String(), c.PimHelloAuthKey))
	}
	if c.PmEnable != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("pm_enable").String(), c.PmEnable))
	}
	if c.PmEnablePrev != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("pm_enable_prev").String(), c.PmEnablePrev))
	}
	if c.PowerRedundancyMode != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("power_redundancy_mode").String(), c.PowerRedundancyMode))
	}
	if c.PtpDomainId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ptp_domain_id").String(), strconv.Itoa(int(*c.PtpDomainId))))
	}
	if c.PtpLbId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ptp_lb_id").String(), strconv.Itoa(int(*c.PtpLbId))))
	}
	if c.PtpProfile != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ptp_profile").String(), c.PtpProfile))
	}
	if c.ReplicationMode != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("replication_mode").String(), c.ReplicationMode))
	}
	if c.RoutingLbId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("routing_lb_id").String(), strconv.Itoa(int(*c.RoutingLbId))))
	}
	if c.RpIpRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("rp_ip_range").String(), c.RpIpRange))
	}
	if c.RpIpRangeInternal != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("rp_ip_range_internal").String(), c.RpIpRangeInternal))
	}
	if c.RpLbId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("rp_lb_id").String(), strconv.Itoa(int(*c.RpLbId))))
	}
	if c.SnmpServerHostTrap != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("snmp_server_host_trap").String(), c.SnmpServerHostTrap))
	}
	if c.SpineCount != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("spine_count").String(), strconv.Itoa(int(*c.SpineCount))))
	}
	if c.StaticUnderlayIpAlloc != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("static_underlay_ip_alloc").String(), c.StaticUnderlayIpAlloc))
	}
	if c.SubnetRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("subnet_range").String(), c.SubnetRange))
	}
	if c.SubnetTargetMask != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("subnet_target_mask").String(), strconv.Itoa(int(*c.SubnetTargetMask))))
	}
	if c.SyslogServerIpList != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("syslog_server_ip_list").String(), c.SyslogServerIpList))
	}
	if c.SyslogServerVrf != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("syslog_server_vrf").String(), c.SyslogServerVrf))
	}
	if c.SyslogSev != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("syslog_sev").String(), c.SyslogSev))
	}
	if c.UpgradeFromVersion != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("upgrade_from_version").String(), c.UpgradeFromVersion))
	}
	if c.AbstractDhcp != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("abstract_dhcp").String(), c.AbstractDhcp))
	}
	if c.AbstractExtraConfigBootstrap != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("abstract_extra_config_bootstrap").String(), c.AbstractExtraConfigBootstrap))
	}
	if c.AbstractExtraConfigLeaf != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("abstract_extra_config_leaf").String(), c.AbstractExtraConfigLeaf))
	}
	if c.AbstractExtraConfigSpine != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("abstract_extra_config_spine").String(), c.AbstractExtraConfigSpine))
	}
	if c.AbstractIsis != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("abstract_isis").String(), c.AbstractIsis))
	}
	if c.AbstractIsisInterface != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("abstract_isis_interface").String(), c.AbstractIsisInterface))
	}
	if c.AbstractLoopbackInterface != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("abstract_loopback_interface").String(), c.AbstractLoopbackInterface))
	}
	if c.AbstractOspf != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("abstract_ospf").String(), c.AbstractOspf))
	}
	if c.AbstractOspfInterface != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("abstract_ospf_interface").String(), c.AbstractOspfInterface))
	}
	if c.AbstractPimInterface != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("abstract_pim_interface").String(), c.AbstractPimInterface))
	}
	if c.AbstractRoutedHost != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("abstract_routed_host").String(), c.AbstractRoutedHost))
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
