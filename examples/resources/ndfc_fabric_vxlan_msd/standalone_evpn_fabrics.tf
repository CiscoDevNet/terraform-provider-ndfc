# =============================================================================
# BEFORE: Two standalone VXLAN EVPN fabrics with their own VRFs and Networks
# =============================================================================

# Standalone Fabric 1
resource "ndfc_fabric_vxlan_evpn" "child_fabric1" {
  fabric_name    = "child_fabric1"
  bgp_as         = 65001
  anycast_gw_mac = "2020.0000.00aa"
  deploy         = true
}

# Standalone Fabric 2
resource "ndfc_fabric_vxlan_evpn" "child_fabric2" {
  fabric_name    = "child_fabric2"
  bgp_as         = 65002
  anycast_gw_mac = "2020.0000.00aa"
  deploy         = true
}

# Inventory for child_fabric1
resource "ndfc_inventory_devices" "child_fabric1_inventory" {
  depends_on  = [ndfc_fabric_vxlan_evpn.child_fabric1]
  fabric_name = ndfc_fabric_vxlan_evpn.child_fabric1.fabric_name
  # ... device configuration ...
}

# Inventory for child_fabric2
resource "ndfc_inventory_devices" "child_fabric2_inventory" {
  depends_on  = [ndfc_fabric_vxlan_evpn.child_fabric2]
  fabric_name = ndfc_fabric_vxlan_evpn.child_fabric2.fabric_name
  # ... device configuration ...
}

# VRFs for child_fabric1 - depends on its own fabric resource
resource "ndfc_vrfs" "child_fabric1_vrfs" {
  depends_on  = [ndfc_inventory_devices.child_fabric1_inventory]
  fabric_name = ndfc_fabric_vxlan_evpn.child_fabric1.fabric_name
  vrfs = {
    for vrf_name, vrf_config in var.vrf_configs : vrf_name => {
      vlan_id                       = vrf_config.vlan_id
      disable_rt_auto               = vrf_config.disable_rt_auto
      ipv6_link_local               = vrf_config.ipv6_link_local
      loopback_routing_tag          = vrf_config.loopback_routing_tag
      max_bgp_paths                 = vrf_config.max_bgp_paths
      max_ibgp_paths                = vrf_config.max_ibgp_paths
      mtu                           = vrf_config.mtu
      redistribute_direct_route_map = vrf_config.redistribute_direct_route_map
      vrf_extension_template        = vrf_config.vrf_extension_template
      vrf_template                  = vrf_config.vrf_template
      advertise_default_route       = vrf_config.advertise_default_route
      advertise_host_routes         = vrf_config.advertise_host_routes
      attach_list = {
        "9PJFRPZZ26Z" = { deploy_this_attachment = true }
        "9G4KUS84LFI" = { deploy_this_attachment = true }
      }
    }
  }
}

# VRFs for child_fabric2 - depends on its own fabric resource
resource "ndfc_vrfs" "child_fabric2_vrfs" {
  depends_on  = [ndfc_inventory_devices.child_fabric2_inventory]
  fabric_name = ndfc_fabric_vxlan_evpn.child_fabric2.fabric_name
  vrfs = {
    for vrf_name, vrf_config in var.vrf_configs : vrf_name => {
      vlan_id                       = vrf_config.vlan_id
      disable_rt_auto               = vrf_config.disable_rt_auto
      ipv6_link_local               = vrf_config.ipv6_link_local
      loopback_routing_tag          = vrf_config.loopback_routing_tag
      max_bgp_paths                 = vrf_config.max_bgp_paths
      max_ibgp_paths                = vrf_config.max_ibgp_paths
      mtu                           = vrf_config.mtu
      redistribute_direct_route_map = vrf_config.redistribute_direct_route_map
      vrf_extension_template        = vrf_config.vrf_extension_template
      vrf_template                  = vrf_config.vrf_template
      advertise_default_route       = vrf_config.advertise_default_route
      advertise_host_routes         = vrf_config.advertise_host_routes
      attach_list = {
        "92Z298RRNPO" = { deploy_this_attachment = true }
        "9ZHWL1HRWKR" = { deploy_this_attachment = true }
      }
    }
  }
}

# Networks for child_fabric1 - depends on VRF resource
resource "ndfc_networks" "child_fabric1_networks" {
  depends_on  = [ndfc_vrfs.child_fabric1_vrfs]
  fabric_name = ndfc_fabric_vxlan_evpn.child_fabric1.fabric_name
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
      routing_tag           = network_config.routing_tag
      trm                   = network_config.trm
      multicast_group       = network_config.multicast_group
      igmp_version          = network_config.igmp_version
      l3_gateway_on_border  = network_config.l3_gateway_on_border
      attach_list = {
        "9PJFRPZZ26Z" = { deploy_this_attachment = true }
        "9G4KUS84LFI" = { deploy_this_attachment = true }
      }
    }
  }
}

# Networks for child_fabric2 - depends on VRF resource
resource "ndfc_networks" "child_fabric2_networks" {
  depends_on  = [ndfc_vrfs.child_fabric2_vrfs]
  fabric_name = ndfc_fabric_vxlan_evpn.child_fabric2.fabric_name
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
      routing_tag           = network_config.routing_tag
      trm                   = network_config.trm
      multicast_group       = network_config.multicast_group
      igmp_version          = network_config.igmp_version
      l3_gateway_on_border  = network_config.l3_gateway_on_border
      attach_list = {
        "92Z298RRNPO" = { deploy_this_attachment = true }
        "9ZHWL1HRWKR" = { deploy_this_attachment = true }
      }
    }
  }
}
