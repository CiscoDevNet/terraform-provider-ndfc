// Code generated;  DO NOT EDIT.

package provider

import (
	"strconv"
	"terraform-provider-ndfc/internal/provider/resources/resource_networks"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func NetworksModelHelperStateCheck(RscName string, c resource_networks.NDFCNetworksModel, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.FabricName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_name").String(), c.FabricName))
	}
	if c.DeployAllAttachments {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("deploy_all_attachments").String(), "true"))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("deploy_all_attachments").String(), "false"))
	}
	for key, value := range c.Networks {
		attrNewPath := attrPath.AtName("networks").AtName(key)
		ret = append(ret, NetworksValueHelperStateCheck(RscName, value, attrNewPath)...)
	}
	return ret
}

func DhcpRelayServersValueHelperStateCheck(RscName string, c resource_networks.NDFCDhcpRelayServersValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.Address != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("address").String(), c.Address))
	}
	if c.Vrf != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf").String(), c.Vrf))
	}
	return ret
}

func NetworksValueHelperStateCheck(RscName string, c resource_networks.NDFCNetworksValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.DisplayName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("display_name").String(), c.DisplayName))
	}
	if c.NetworkId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("network_id").String(), strconv.Itoa(int(*c.NetworkId))))
	}
	if c.NetworkTemplate != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("network_template").String(), c.NetworkTemplate))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("network_template").String(), "Default_Network_Universal"))
	}
	if c.NetworkExtensionTemplate != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("network_extension_template").String(), c.NetworkExtensionTemplate))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("network_extension_template").String(), "Default_Network_Extension_Universal"))
	}
	if c.VrfName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_name").String(), c.VrfName))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_name").String(), "NA"))
	}
	if c.PrimaryNetworkId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("primary_network_id").String(), strconv.Itoa(int(*c.PrimaryNetworkId))))
	}
	if c.NetworkType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("network_type").String(), c.NetworkType))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("network_type").String(), "Normal"))
	}
	if c.NetworkStatus != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("network_status").String(), c.NetworkStatus))
	}
	if c.NetworkTemplateConfig.GatewayIpv4Address != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("gateway_ipv4_address").String(), c.NetworkTemplateConfig.GatewayIpv4Address))
	}
	if c.NetworkTemplateConfig.GatewayIpv6Address != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("gateway_ipv6_address").String(), c.NetworkTemplateConfig.GatewayIpv6Address))
	}
	if c.NetworkTemplateConfig.VlanId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vlan_id").String(), strconv.Itoa(int(*c.NetworkTemplateConfig.VlanId))))
	}
	if c.NetworkTemplateConfig.VlanName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vlan_name").String(), c.NetworkTemplateConfig.VlanName))
	}
	if c.NetworkTemplateConfig.Layer2Only != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("layer2_only").String(), c.NetworkTemplateConfig.Layer2Only))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("layer2_only").String(), "false"))
	}
	if c.NetworkTemplateConfig.InterfaceDescription != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_description").String(), c.NetworkTemplateConfig.InterfaceDescription))
	}
	if c.NetworkTemplateConfig.Mtu != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mtu").String(), strconv.Itoa(int(*c.NetworkTemplateConfig.Mtu))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mtu").String(), "9216"))
	}
	if c.NetworkTemplateConfig.SecondaryGateway1 != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("secondary_gateway_1").String(), c.NetworkTemplateConfig.SecondaryGateway1))
	}
	if c.NetworkTemplateConfig.SecondaryGateway2 != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("secondary_gateway_2").String(), c.NetworkTemplateConfig.SecondaryGateway2))
	}
	if c.NetworkTemplateConfig.SecondaryGateway3 != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("secondary_gateway_3").String(), c.NetworkTemplateConfig.SecondaryGateway3))
	}
	if c.NetworkTemplateConfig.SecondaryGateway4 != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("secondary_gateway_4").String(), c.NetworkTemplateConfig.SecondaryGateway4))
	}
	if c.NetworkTemplateConfig.ArpSuppression != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("arp_suppression").String(), c.NetworkTemplateConfig.ArpSuppression))
	}
	if c.NetworkTemplateConfig.IngressReplication != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ingress_replication").String(), c.NetworkTemplateConfig.IngressReplication))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ingress_replication").String(), "false"))
	}
	if c.NetworkTemplateConfig.MulticastGroup != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("multicast_group").String(), c.NetworkTemplateConfig.MulticastGroup))
	}
	if c.NetworkTemplateConfig.DhcpRelayLoopbackId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dhcp_relay_loopback_id").String(), strconv.Itoa(int(*c.NetworkTemplateConfig.DhcpRelayLoopbackId))))
	}
	if c.NetworkTemplateConfig.RoutingTag != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("routing_tag").String(), strconv.Itoa(int(*c.NetworkTemplateConfig.RoutingTag))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("routing_tag").String(), "12345"))
	}
	if c.NetworkTemplateConfig.Trm != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("trm").String(), c.NetworkTemplateConfig.Trm))
	}
	if c.NetworkTemplateConfig.RouteTargetBoth != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("route_target_both").String(), c.NetworkTemplateConfig.RouteTargetBoth))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("route_target_both").String(), "false"))
	}
	if c.NetworkTemplateConfig.Netflow != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("netflow").String(), c.NetworkTemplateConfig.Netflow))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("netflow").String(), "false"))
	}
	if c.NetworkTemplateConfig.SviNetflowMonitor != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("svi_netflow_monitor").String(), c.NetworkTemplateConfig.SviNetflowMonitor))
	}
	if c.NetworkTemplateConfig.VlanNetflowMonitor != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vlan_netflow_monitor").String(), c.NetworkTemplateConfig.VlanNetflowMonitor))
	}
	if c.NetworkTemplateConfig.L3GatwayBorder != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l3_gatway_border").String(), c.NetworkTemplateConfig.L3GatwayBorder))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l3_gatway_border").String(), "false"))
	}
	if c.NetworkTemplateConfig.IgmpVersion != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("igmp_version").String(), c.NetworkTemplateConfig.IgmpVersion))
	}
	if c.DeployAttachments {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("deploy_attachments").String(), "true"))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("deploy_attachments").String(), "false"))
	}
	for key, value := range c.Attachments {
		attrNewPath := attrPath.AtName("attachments").AtName(key)
		ret = append(ret, AttachmentsValueHelperStateCheck(RscName, value, attrNewPath)...)
	}
	return ret
}
