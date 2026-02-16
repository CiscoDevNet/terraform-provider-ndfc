resource "ndfc_networks" "msd_parent_networks_attributes" {
  depends_on  = [ndfc_vrfs.msd_parent_vrfs_attributes, ndfc_vrfs.child1_specific_vrfs_attributes, ndfc_vrfs.child2_specific_vrfs_attributes]
  fabric_name = ndfc_fabric_vxlan_msd.msd_fabric.fabric_name
  networks = {
    for network_name, network_config in var.network_configs : network_name => {
      # Parent properties only
      vrf_name              = network_config.vrf_name
      network_id            = network_config.network_id
      gateway_ipv4_address  = network_config.gateway_ipv4_address
      gateway_ipv6_address  = network_config.gateway_ipv6_address
      vlan_id               = network_config.vlan_id
      vlan_name             = network_config.vlan_name
      layer2_only           = network_config.layer2_only
      interface_description = network_config.interface_description
      mtu                   = network_config.mtu
      secondary_gateway_1   = network_config.secondary_gateway_1
      secondary_gateway_2   = network_config.secondary_gateway_2
      secondary_gateway_3   = network_config.secondary_gateway_3
      secondary_gateway_4   = network_config.secondary_gateway_4
      routing_tag           = network_config.routing_tag
      arp_suppression       = network_config.arp_suppression
      route_target_both     = network_config.route_target_both
      # Attachments with fabric field
      attach_list = {
        "9PJFRPZZ26Z" = {
          deploy_this_attachment = true
          fabric                 = tolist(ndfc_fabric_vxlan_msd.msd_fabric.child_fabrics)[0]
        }
        "9G4KUS84LFI" = {
          deploy_this_attachment = true
          fabric                 = tolist(ndfc_fabric_vxlan_msd.msd_fabric.child_fabrics)[0]
        }
        "92Z298RRNPO" = {
          deploy_this_attachment = true
          fabric                 = tolist(ndfc_fabric_vxlan_msd.msd_fabric.child_fabrics)[1]
        }
        "9ZHWL1HRWKR" = {
          deploy_this_attachment = true
          fabric                 = tolist(ndfc_fabric_vxlan_msd.msd_fabric.child_fabrics)[1]
        }
      }
    }
  }
}

resource "ndfc_networks" "child1_specific_networks_attributes" {
  depends_on  = [ndfc_networks.msd_parent_networks_attributes]
  fabric_name = tolist(ndfc_fabric_vxlan_msd.msd_fabric.child_fabrics)[0]
  networks = {
    for network_name, network_config in var.network_configs : network_name => {
      vrf_name              = network_config.vrf_name
      network_id            = network_config.network_id
      gateway_ipv4_address  = network_config.gateway_ipv4_address
      gateway_ipv6_address  = network_config.gateway_ipv6_address
      vlan_id               = network_config.vlan_id
      vlan_name             = network_config.vlan_name
      layer2_only           = network_config.layer2_only
      interface_description = network_config.interface_description
      mtu                   = network_config.mtu
      secondary_gateway_1   = network_config.secondary_gateway_1
      secondary_gateway_2   = network_config.secondary_gateway_2
      secondary_gateway_3   = network_config.secondary_gateway_3
      secondary_gateway_4   = network_config.secondary_gateway_4
      routing_tag           = network_config.routing_tag
      arp_suppression       = network_config.arp_suppression
      route_target_both     = network_config.route_target_both
      dhcp_loopback_id      = network_config.dhcp_loopback_id
      dhcp_servers          = network_config.dhcp_servers
      multicast_group       = network_config.multicast_group
      trm                   = network_config.trm
      netflow               = network_config.netflow
      vlan_netflow_monitor  = network_config.vlan_netflow_monitor
      l3_gateway_on_border  = network_config.l3_gateway_on_border
      igmp_version          = network_config.igmp_version
      # Attachments (no fabric field for child)
      attach_list = {
        "9PJFRPZZ26Z" = {
          deploy_this_attachment = true
        }
        "9G4KUS84LFI" = {
          deploy_this_attachment = true
        }
      }
    }
  }
}

resource "ndfc_networks" "child2_specific_networks_attributes" {
  depends_on  = [ndfc_networks.msd_parent_networks_attributes]
  fabric_name = tolist(ndfc_fabric_vxlan_msd.msd_fabric.child_fabrics)[1]
  networks = {
    for network_name, network_config in var.network_configs : network_name => {
      vrf_name              = network_config.vrf_name
      network_id            = network_config.network_id
      gateway_ipv4_address  = network_config.gateway_ipv4_address
      gateway_ipv6_address  = network_config.gateway_ipv6_address
      vlan_id               = network_config.vlan_id
      vlan_name             = network_config.vlan_name
      layer2_only           = network_config.layer2_only
      interface_description = network_config.interface_description
      mtu                   = network_config.mtu
      secondary_gateway_1   = network_config.secondary_gateway_1
      secondary_gateway_2   = network_config.secondary_gateway_2
      secondary_gateway_3   = network_config.secondary_gateway_3
      secondary_gateway_4   = network_config.secondary_gateway_4
      routing_tag           = network_config.routing_tag
      arp_suppression       = network_config.arp_suppression
      route_target_both     = network_config.route_target_both
      dhcp_loopback_id      = network_config.dhcp_loopback_id
      dhcp_servers          = network_config.dhcp_servers
      multicast_group       = network_config.multicast_group
      trm                   = network_config.trm
      netflow               = network_config.netflow
      vlan_netflow_monitor  = network_config.vlan_netflow_monitor
      l3_gateway_on_border  = network_config.l3_gateway_on_border
      igmp_version          = network_config.igmp_version
      # Attachments (no fabric field for child)
      attach_list = {
        "92Z298RRNPO" = {
          deploy_this_attachment = true
        }
        "9ZHWL1HRWKR" = {
          deploy_this_attachment = true
        }
      }
    }
  }
}
