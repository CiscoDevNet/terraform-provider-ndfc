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

func FabricVxlanEvpnModelHelperStateCheck(RscName string, c resource_fabric_common.NDFCFabricCommonModel, attrPath path.Path) []resource.TestCheckFunc {
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
	if c.AdvertisePipBgp != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("advertise_pip_bgp").String(), c.AdvertisePipBgp))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("advertise_pip_bgp").String(), "false"))
	}
	if c.AdvertisePipOnBorder != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("advertise_pip_on_border").String(), c.AdvertisePipOnBorder))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("advertise_pip_on_border").String(), "true"))
	}
	if c.AnycastBgwAdvertisePip != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("anycast_bgw_advertise_pip").String(), c.AnycastBgwAdvertisePip))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("anycast_bgw_advertise_pip").String(), "false"))
	}
	if c.AnycastGwMac != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("anycast_gw_mac").String(), c.AnycastGwMac))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("anycast_gw_mac").String(), "2020.0000.00aa"))
	}
	if c.AnycastLbId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("anycast_lb_id").String(), strconv.Itoa(int(*c.AnycastLbId))))
	}
	if c.AnycastRpIpRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("anycast_rp_ip_range").String(), c.AnycastRpIpRange))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("anycast_rp_ip_range").String(), "10.254.254.0/24"))
	}
	if c.AutoSymmetricDefaultVrf != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("auto_symmetric_default_vrf").String(), c.AutoSymmetricDefaultVrf))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("auto_symmetric_default_vrf").String(), "false"))
	}
	if c.AutoSymmetricVrfLite != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("auto_symmetric_vrf_lite").String(), c.AutoSymmetricVrfLite))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("auto_symmetric_vrf_lite").String(), "false"))
	}
	if c.AutoUniqueVrfLiteIpPrefix != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("auto_unique_vrf_lite_ip_prefix").String(), c.AutoUniqueVrfLiteIpPrefix))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("auto_unique_vrf_lite_ip_prefix").String(), "false"))
	}
	if c.AutoVrfliteIfcDefaultVrf != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("auto_vrflite_ifc_default_vrf").String(), c.AutoVrfliteIfcDefaultVrf))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("auto_vrflite_ifc_default_vrf").String(), "false"))
	}
	if c.Banner != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("banner").String(), c.Banner))
	}
	if c.BfdAuthEnable != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bfd_auth_enable").String(), c.BfdAuthEnable))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bfd_auth_enable").String(), "false"))
	}
	if c.BfdAuthKey != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bfd_auth_key").String(), c.BfdAuthKey))
	}
	if c.BfdAuthKeyId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bfd_auth_key_id").String(), strconv.Itoa(int(*c.BfdAuthKeyId))))
	}
	if c.BfdEnable != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bfd_enable").String(), c.BfdEnable))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bfd_enable").String(), "false"))
	}
	if c.BfdIbgpEnable != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bfd_ibgp_enable").String(), c.BfdIbgpEnable))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bfd_ibgp_enable").String(), "false"))
	}
	if c.BfdIsisEnable != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bfd_isis_enable").String(), c.BfdIsisEnable))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bfd_isis_enable").String(), "false"))
	}
	if c.BfdOspfEnable != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bfd_ospf_enable").String(), c.BfdOspfEnable))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bfd_ospf_enable").String(), "false"))
	}
	if c.BfdPimEnable != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bfd_pim_enable").String(), c.BfdPimEnable))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bfd_pim_enable").String(), "false"))
	}
	if c.BgpAs != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bgp_as").String(), c.BgpAs))
	}
	if c.BgpAuthEnable != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bgp_auth_enable").String(), c.BgpAuthEnable))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bgp_auth_enable").String(), "false"))
	}
	if c.BgpAuthKey != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bgp_auth_key").String(), c.BgpAuthKey))
	}
	if c.BgpAuthKeyType != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bgp_auth_key_type").String(), strconv.Itoa(int(*c.BgpAuthKeyType))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bgp_auth_key_type").String(), "3"))
	}
	if c.BgpLbId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bgp_lb_id").String(), strconv.Itoa(int(*c.BgpLbId))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bgp_lb_id").String(), "0"))
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
	if c.BrownfieldNetworkNameFormat != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("brownfield_network_name_format").String(), c.BrownfieldNetworkNameFormat))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("brownfield_network_name_format").String(), "Auto_Net_VNI$$VNI$$_VLAN$$VLAN_ID$$"))
	}
	if c.BrownfieldSkipOverlayNetworkAttachments != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("brownfield_skip_overlay_network_attachments").String(), c.BrownfieldSkipOverlayNetworkAttachments))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("brownfield_skip_overlay_network_attachments").String(), "false"))
	}
	if c.CdpEnable != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("cdp_enable").String(), c.CdpEnable))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("cdp_enable").String(), "false"))
	}
	if c.CoppPolicy != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("copp_policy").String(), c.CoppPolicy))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("copp_policy").String(), "strict"))
	}
	if c.DciSubnetRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dci_subnet_range").String(), c.DciSubnetRange))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dci_subnet_range").String(), "10.33.0.0/16"))
	}
	if c.DciSubnetTargetMask != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dci_subnet_target_mask").String(), strconv.Itoa(int(*c.DciSubnetTargetMask))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dci_subnet_target_mask").String(), "30"))
	}
	if c.DefaultQueuingPolicyCloudscale != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("default_queuing_policy_cloudscale").String(), c.DefaultQueuingPolicyCloudscale))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("default_queuing_policy_cloudscale").String(), "queuing_policy_default_8q_cloudscale"))
	}
	if c.DefaultQueuingPolicyOther != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("default_queuing_policy_other").String(), c.DefaultQueuingPolicyOther))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("default_queuing_policy_other").String(), "queuing_policy_default_other"))
	}
	if c.DefaultQueuingPolicyRSeries != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("default_queuing_policy_r_series").String(), c.DefaultQueuingPolicyRSeries))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("default_queuing_policy_r_series").String(), "queuing_policy_default_r_series"))
	}
	if c.DefaultVrfRedisBgpRmap != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("default_vrf_redis_bgp_rmap").String(), c.DefaultVrfRedisBgpRmap))
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
	if c.EnableDefaultQueuingPolicy != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_default_queuing_policy").String(), c.EnableDefaultQueuingPolicy))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_default_queuing_policy").String(), "false"))
	}
	if c.EnableFabricVpcDomainId != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_fabric_vpc_domain_id").String(), c.EnableFabricVpcDomainId))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_fabric_vpc_domain_id").String(), "false"))
	}
	if c.EnableMacsec != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_macsec").String(), c.EnableMacsec))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_macsec").String(), "false"))
	}
	if c.EnableNetflow != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_netflow").String(), c.EnableNetflow))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_netflow").String(), "false"))
	}
	if c.EnableNgoam != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_ngoam").String(), c.EnableNgoam))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_ngoam").String(), "true"))
	}
	if c.EnableNxapi != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_nxapi").String(), c.EnableNxapi))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_nxapi").String(), "true"))
	}
	if c.EnableNxapiHttp != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_nxapi_http").String(), c.EnableNxapiHttp))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_nxapi_http").String(), "true"))
	}
	if c.EnablePbr != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_pbr").String(), c.EnablePbr))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_pbr").String(), "false"))
	}
	if c.EnablePvlan != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_pvlan").String(), c.EnablePvlan))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_pvlan").String(), "false"))
	}
	if c.EnableTenantDhcp != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_tenant_dhcp").String(), c.EnableTenantDhcp))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_tenant_dhcp").String(), "true"))
	}
	if c.EnableTrm != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_trm").String(), c.EnableTrm))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_trm").String(), "false"))
	}
	if c.EnableVpcPeerLinkNativeVlan != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_vpc_peer_link_native_vlan").String(), c.EnableVpcPeerLinkNativeVlan))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_vpc_peer_link_native_vlan").String(), "false"))
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
	if c.ExtraConfTor != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("extra_conf_tor").String(), c.ExtraConfTor))
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
	if c.FabricVpcDomainId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_vpc_domain_id").String(), strconv.Itoa(int(*c.FabricVpcDomainId))))
	}
	if c.FabricVpcQos != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_vpc_qos").String(), c.FabricVpcQos))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_vpc_qos").String(), "false"))
	}
	if c.FabricVpcQosPolicyName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_vpc_qos_policy_name").String(), c.FabricVpcQosPolicyName))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_vpc_qos_policy_name").String(), "spine_qos_for_fabric_vpc_peering"))
	}
	if c.FeaturePtp != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("feature_ptp").String(), c.FeaturePtp))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("feature_ptp").String(), "false"))
	}
	if c.GrfieldDebugFlag != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("grfield_debug_flag").String(), c.GrfieldDebugFlag))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("grfield_debug_flag").String(), "Disable"))
	}
	if c.HdTime != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("hd_time").String(), strconv.Itoa(int(*c.HdTime))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("hd_time").String(), "180"))
	}
	if c.HostIntfAdminState != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("host_intf_admin_state").String(), c.HostIntfAdminState))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("host_intf_admin_state").String(), "true"))
	}
	if c.IbgpPeerTemplate != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ibgp_peer_template").String(), c.IbgpPeerTemplate))
	}
	if c.IbgpPeerTemplateLeaf != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ibgp_peer_template_leaf").String(), c.IbgpPeerTemplateLeaf))
	}
	if c.InbandDhcpServers != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("inband_dhcp_servers").String(), c.InbandDhcpServers))
	}
	if c.InbandMgmt != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("inband_mgmt").String(), c.InbandMgmt))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("inband_mgmt").String(), "false"))
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
	if c.IsisOverloadElapseTime != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("isis_overload_elapse_time").String(), strconv.Itoa(int(*c.IsisOverloadElapseTime))))
	}
	if c.IsisOverloadEnable != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("isis_overload_enable").String(), c.IsisOverloadEnable))
	}
	if c.IsisP2pEnable != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("isis_p2p_enable").String(), c.IsisP2pEnable))
	}
	if c.L2HostIntfMtu != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l2_host_intf_mtu").String(), strconv.Itoa(int(*c.L2HostIntfMtu))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l2_host_intf_mtu").String(), "9216"))
	}
	if c.L2SegmentIdRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l2_segment_id_range").String(), c.L2SegmentIdRange))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l2_segment_id_range").String(), "30000-49000"))
	}
	if c.L3vniMcastGroup != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l3vni_mcast_group").String(), c.L3vniMcastGroup))
	}
	if c.L3PartitionIdRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l3_partition_id_range").String(), c.L3PartitionIdRange))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l3_partition_id_range").String(), "50000-59000"))
	}
	if c.LinkStateRouting != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("link_state_routing").String(), c.LinkStateRouting))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("link_state_routing").String(), "ospf"))
	}
	if c.LinkStateRoutingTag != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("link_state_routing_tag").String(), c.LinkStateRoutingTag))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("link_state_routing_tag").String(), "UNDERLAY"))
	}
	if c.Loopback0Ipv6Range != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("loopback0_ipv6_range").String(), c.Loopback0Ipv6Range))
	}
	if c.Loopback0IpRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("loopback0_ip_range").String(), c.Loopback0IpRange))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("loopback0_ip_range").String(), "10.2.0.0/22"))
	}
	if c.Loopback1Ipv6Range != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("loopback1_ipv6_range").String(), c.Loopback1Ipv6Range))
	}
	if c.Loopback1IpRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("loopback1_ip_range").String(), c.Loopback1IpRange))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("loopback1_ip_range").String(), "10.3.0.0/22"))
	}
	if c.MacsecAlgorithm != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("macsec_algorithm").String(), c.MacsecAlgorithm))
	}
	if c.MacsecCipherSuite != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("macsec_cipher_suite").String(), c.MacsecCipherSuite))
	}
	if c.MacsecFallbackAlgorithm != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("macsec_fallback_algorithm").String(), c.MacsecFallbackAlgorithm))
	}
	if c.MacsecFallbackKeyString != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("macsec_fallback_key_string").String(), c.MacsecFallbackKeyString))
	}
	if c.MacsecKeyString != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("macsec_key_string").String(), c.MacsecKeyString))
	}
	if c.MacsecReportTimer != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("macsec_report_timer").String(), strconv.Itoa(int(*c.MacsecReportTimer))))
	}
	if c.MgmtGw != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mgmt_gw").String(), c.MgmtGw))
	}
	if c.MgmtPrefix != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mgmt_prefix").String(), strconv.Itoa(int(*c.MgmtPrefix))))
	}
	if c.MgmtV6prefix != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mgmt_v6prefix").String(), strconv.Itoa(int(*c.MgmtV6prefix))))
	}
	if c.MplsHandoff != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mpls_handoff").String(), c.MplsHandoff))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mpls_handoff").String(), "false"))
	}
	if c.MplsLbId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mpls_lb_id").String(), strconv.Itoa(int(*c.MplsLbId))))
	}
	if c.MplsLoopbackIpRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mpls_loopback_ip_range").String(), c.MplsLoopbackIpRange))
	}
	if c.MstInstanceRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mst_instance_range").String(), c.MstInstanceRange))
	}
	if c.MulticastGroupSubnet != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("multicast_group_subnet").String(), c.MulticastGroupSubnet))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("multicast_group_subnet").String(), "239.1.1.0/25"))
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
	if c.NetworkVlanRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("network_vlan_range").String(), c.NetworkVlanRange))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("network_vlan_range").String(), "2300-2999"))
	}
	if c.NtpServerIpList != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ntp_server_ip_list").String(), c.NtpServerIpList))
	}
	if c.NtpServerVrf != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ntp_server_vrf").String(), c.NtpServerVrf))
	}
	if c.NveLbId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("nve_lb_id").String(), strconv.Itoa(int(*c.NveLbId))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("nve_lb_id").String(), "1"))
	}
	if c.NxapiHttpsPort != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("nxapi_https_port").String(), strconv.Itoa(int(*c.NxapiHttpsPort))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("nxapi_https_port").String(), "443"))
	}
	if c.NxapiHttpPort != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("nxapi_http_port").String(), strconv.Itoa(int(*c.NxapiHttpPort))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("nxapi_http_port").String(), "80"))
	}
	if c.ObjectTrackingNumberRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("object_tracking_number_range").String(), c.ObjectTrackingNumberRange))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("object_tracking_number_range").String(), "100-299"))
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
	if c.OverlayMode != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("overlay_mode").String(), c.OverlayMode))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("overlay_mode").String(), "cli"))
	}
	if c.PerVrfLoopbackAutoProvision != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("per_vrf_loopback_auto_provision").String(), c.PerVrfLoopbackAutoProvision))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("per_vrf_loopback_auto_provision").String(), "false"))
	}
	if c.PerVrfLoopbackIpRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("per_vrf_loopback_ip_range").String(), c.PerVrfLoopbackIpRange))
	}
	if c.PhantomRpLbId1 != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("phantom_rp_lb_id1").String(), strconv.Itoa(int(*c.PhantomRpLbId1))))
	}
	if c.PhantomRpLbId2 != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("phantom_rp_lb_id2").String(), strconv.Itoa(int(*c.PhantomRpLbId2))))
	}
	if c.PhantomRpLbId3 != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("phantom_rp_lb_id3").String(), strconv.Itoa(int(*c.PhantomRpLbId3))))
	}
	if c.PhantomRpLbId4 != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("phantom_rp_lb_id4").String(), strconv.Itoa(int(*c.PhantomRpLbId4))))
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
	if c.ReplicationMode != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("replication_mode").String(), c.ReplicationMode))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("replication_mode").String(), "Multicast"))
	}
	if c.RouterIdRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("router_id_range").String(), c.RouterIdRange))
	}
	if c.RouteMapSequenceNumberRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("route_map_sequence_number_range").String(), c.RouteMapSequenceNumberRange))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("route_map_sequence_number_range").String(), "1-65534"))
	}
	if c.RpCount != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("rp_count").String(), strconv.Itoa(int(*c.RpCount))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("rp_count").String(), "2"))
	}
	if c.RpLbId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("rp_lb_id").String(), strconv.Itoa(int(*c.RpLbId))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("rp_lb_id").String(), "254"))
	}
	if c.RpMode != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("rp_mode").String(), c.RpMode))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("rp_mode").String(), "asm"))
	}
	if c.RrCount != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("rr_count").String(), strconv.Itoa(int(*c.RrCount))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("rr_count").String(), "2"))
	}
	if c.SeedSwitchCoreInterfaces != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("seed_switch_core_interfaces").String(), c.SeedSwitchCoreInterfaces))
	}
	if c.ServiceNetworkVlanRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("service_network_vlan_range").String(), c.ServiceNetworkVlanRange))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("service_network_vlan_range").String(), "3000-3199"))
	}
	if c.SiteId != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("site_id").String(), c.SiteId))
	}
	if c.SlaIdRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("sla_id_range").String(), c.SlaIdRange))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("sla_id_range").String(), "10000-19999"))
	}
	if c.SnmpServerHostTrap != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("snmp_server_host_trap").String(), c.SnmpServerHostTrap))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("snmp_server_host_trap").String(), "true"))
	}
	if c.SpineSwitchCoreInterfaces != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("spine_switch_core_interfaces").String(), c.SpineSwitchCoreInterfaces))
	}
	if c.StaticUnderlayIpAlloc != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("static_underlay_ip_alloc").String(), c.StaticUnderlayIpAlloc))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("static_underlay_ip_alloc").String(), "false"))
	}
	if c.StpBridgePriority != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("stp_bridge_priority").String(), strconv.Itoa(int(*c.StpBridgePriority))))
	}
	if c.StpRootOption != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("stp_root_option").String(), c.StpRootOption))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("stp_root_option").String(), "unmanaged"))
	}
	if c.StpVlanRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("stp_vlan_range").String(), c.StpVlanRange))
	}
	if c.StrictCcMode != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("strict_cc_mode").String(), c.StrictCcMode))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("strict_cc_mode").String(), "false"))
	}
	if c.SubinterfaceRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("subinterface_range").String(), c.SubinterfaceRange))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("subinterface_range").String(), "2-511"))
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
	if c.TcamAllocation != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("tcam_allocation").String(), c.TcamAllocation))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("tcam_allocation").String(), "true"))
	}
	if c.UnderlayIsV6 != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("underlay_is_v6").String(), c.UnderlayIsV6))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("underlay_is_v6").String(), "false"))
	}
	if c.UnnumBootstrapLbId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("unnum_bootstrap_lb_id").String(), strconv.Itoa(int(*c.UnnumBootstrapLbId))))
	}
	if c.UnnumDhcpEnd != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("unnum_dhcp_end").String(), c.UnnumDhcpEnd))
	}
	if c.UnnumDhcpStart != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("unnum_dhcp_start").String(), c.UnnumDhcpStart))
	}
	if c.UseLinkLocal != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("use_link_local").String(), c.UseLinkLocal))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("use_link_local").String(), "true"))
	}
	if c.V6SubnetRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("v6_subnet_range").String(), c.V6SubnetRange))
	}
	if c.V6SubnetTargetMask != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("v6_subnet_target_mask").String(), strconv.Itoa(int(*c.V6SubnetTargetMask))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("v6_subnet_target_mask").String(), "126"))
	}
	if c.VpcAutoRecoveryTime != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vpc_auto_recovery_time").String(), strconv.Itoa(int(*c.VpcAutoRecoveryTime))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vpc_auto_recovery_time").String(), "360"))
	}
	if c.VpcDelayRestore != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vpc_delay_restore").String(), strconv.Itoa(int(*c.VpcDelayRestore))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vpc_delay_restore").String(), "150"))
	}
	if c.VpcDomainIdRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vpc_domain_id_range").String(), c.VpcDomainIdRange))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vpc_domain_id_range").String(), "1-1000"))
	}
	if c.VpcEnableIpv6NdSync != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vpc_enable_ipv6_nd_sync").String(), c.VpcEnableIpv6NdSync))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vpc_enable_ipv6_nd_sync").String(), "true"))
	}
	if c.VpcPeerKeepAliveOption != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vpc_peer_keep_alive_option").String(), c.VpcPeerKeepAliveOption))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vpc_peer_keep_alive_option").String(), "management"))
	}
	if c.VpcPeerLinkPo != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vpc_peer_link_po").String(), strconv.Itoa(int(*c.VpcPeerLinkPo))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vpc_peer_link_po").String(), "500"))
	}
	if c.VpcPeerLinkVlan != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vpc_peer_link_vlan").String(), strconv.Itoa(int(*c.VpcPeerLinkVlan))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vpc_peer_link_vlan").String(), "3600"))
	}
	if c.VrfLiteAutoconfig != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_lite_autoconfig").String(), c.VrfLiteAutoconfig))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_lite_autoconfig").String(), "Manual"))
	}
	if c.VrfVlanRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_vlan_range").String(), c.VrfVlanRange))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_vlan_range").String(), "2000-2299"))
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
	if c.EnableRealtimeBackup != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_realtime_backup").String(), c.EnableRealtimeBackup))
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
