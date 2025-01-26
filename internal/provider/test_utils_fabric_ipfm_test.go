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
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("aaa_remote_ip_enabled").String(), "false"))
	}
	if c.AaaServerConf != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("aaa_server_conf").String(), c.AaaServerConf))
	}
	if c.AsmGroupRanges != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("asm_group_ranges").String(), c.AsmGroupRanges))
	}
	if c.BootstrapConf != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bootstrap_conf").String(), c.BootstrapConf))
	}
	if c.BootstrapEnable != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bootstrap_enable").String(), c.BootstrapEnable))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bootstrap_enable").String(), "false"))
	}
	if c.BootstrapMultisubnet != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bootstrap_multisubnet").String(), c.BootstrapMultisubnet))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bootstrap_multisubnet").String(), "#Scope_Start_IP, Scope_End_IP, Scope_Default_Gateway, Scope_Subnet_Prefix"))
	}
	if c.CdpEnable != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("cdp_enable").String(), c.CdpEnable))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("cdp_enable").String(), "false"))
	}
	if c.DhcpEnable != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dhcp_enable").String(), c.DhcpEnable))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dhcp_enable").String(), "false"))
	}
	if c.DhcpEnd != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dhcp_end").String(), c.DhcpEnd))
	}
	if c.DhcpIpv6Enable != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dhcp_ipv6_enable").String(), c.DhcpIpv6Enable))
	}
	if c.DhcpStart != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dhcp_start").String(), c.DhcpStart))
	}
	if c.DnsServerIpList != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dns_server_ip_list").String(), c.DnsServerIpList))
	}
	if c.DnsServerVrf != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dns_server_vrf").String(), c.DnsServerVrf))
	}
	if c.EnableAaa != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_aaa").String(), c.EnableAaa))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_aaa").String(), "false"))
	}
	if c.EnableAsm != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_asm").String(), c.EnableAsm))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_asm").String(), "false"))
	}
	if c.EnableNbmPassive != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_nbm_passive").String(), c.EnableNbmPassive))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_nbm_passive").String(), "false"))
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
	if c.FabricInterfaceType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_interface_type").String(), c.FabricInterfaceType))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_interface_type").String(), "p2p"))
	}
	if c.FabricMtu != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_mtu").String(), strconv.Itoa(int(*c.FabricMtu))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_mtu").String(), "9216"))
	}
	if c.FeaturePtp != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("feature_ptp").String(), c.FeaturePtp))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("feature_ptp").String(), "false"))
	}
	if c.IsisAuthEnable != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("isis_auth_enable").String(), c.IsisAuthEnable))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("isis_auth_enable").String(), "false"))
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
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("isis_level").String(), "level-2"))
	}
	if c.IsisP2pEnable != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("isis_p2p_enable").String(), c.IsisP2pEnable))
	}
	if c.L2HostIntfMtu != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l2_host_intf_mtu").String(), strconv.Itoa(int(*c.L2HostIntfMtu))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l2_host_intf_mtu").String(), "9216"))
	}
	if c.LinkStateRouting != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("link_state_routing").String(), c.LinkStateRouting))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("link_state_routing").String(), "ospf"))
	}
	if c.LinkStateRoutingTag != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("link_state_routing_tag").String(), c.LinkStateRoutingTag))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("link_state_routing_tag").String(), "1"))
	}
	if c.Loopback0IpRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("loopback0_ip_range").String(), c.Loopback0IpRange))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("loopback0_ip_range").String(), "10.2.0.0/22"))
	}
	if c.MgmtGw != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mgmt_gw").String(), c.MgmtGw))
	}
	if c.MgmtPrefix != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mgmt_prefix").String(), strconv.Itoa(int(*c.MgmtPrefix))))
	}
	if c.NtpServerIpList != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ntp_server_ip_list").String(), c.NtpServerIpList))
	}
	if c.NtpServerVrf != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ntp_server_vrf").String(), c.NtpServerVrf))
	}
	if c.NxapiVrf != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("nxapi_vrf").String(), c.NxapiVrf))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("nxapi_vrf").String(), "management"))
	}
	if c.OspfAreaId != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ospf_area_id").String(), c.OspfAreaId))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ospf_area_id").String(), "0.0.0.0"))
	}
	if c.OspfAuthEnable != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ospf_auth_enable").String(), c.OspfAuthEnable))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ospf_auth_enable").String(), "false"))
	}
	if c.OspfAuthKey != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ospf_auth_key").String(), c.OspfAuthKey))
	}
	if c.OspfAuthKeyId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ospf_auth_key_id").String(), strconv.Itoa(int(*c.OspfAuthKeyId))))
	}
	if c.PimHelloAuthEnable != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("pim_hello_auth_enable").String(), c.PimHelloAuthEnable))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("pim_hello_auth_enable").String(), "false"))
	}
	if c.PimHelloAuthKey != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("pim_hello_auth_key").String(), c.PimHelloAuthKey))
	}
	if c.PmEnable != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("pm_enable").String(), c.PmEnable))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("pm_enable").String(), "false"))
	}
	if c.PowerRedundancyMode != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("power_redundancy_mode").String(), c.PowerRedundancyMode))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("power_redundancy_mode").String(), "ps-redundant"))
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
	if c.RoutingLbId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("routing_lb_id").String(), strconv.Itoa(int(*c.RoutingLbId))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("routing_lb_id").String(), "0"))
	}
	if c.RpIpRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("rp_ip_range").String(), c.RpIpRange))
	}
	if c.RpLbId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("rp_lb_id").String(), strconv.Itoa(int(*c.RpLbId))))
	}
	if c.SnmpServerHostTrap != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("snmp_server_host_trap").String(), c.SnmpServerHostTrap))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("snmp_server_host_trap").String(), "true"))
	}
	if c.StaticUnderlayIpAlloc != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("static_underlay_ip_alloc").String(), c.StaticUnderlayIpAlloc))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("static_underlay_ip_alloc").String(), "false"))
	}
	if c.SubnetRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("subnet_range").String(), c.SubnetRange))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("subnet_range").String(), "10.4.0.0/16"))
	}
	if c.SubnetTargetMask != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("subnet_target_mask").String(), strconv.Itoa(int(*c.SubnetTargetMask))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("subnet_target_mask").String(), "30"))
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
