# =============================================================================
# AFTER: child_fabric2 removed from MSD, now operating as standalone fabric
# =============================================================================

# MSD fabric now only manages child_fabric1
resource "ndfc_fabric_vxlan_msd" "msd_fabric" {
  depends_on    = [ndfc_inventory_devices.child_fabric1_inventory]
  fabric_name   = "msd_fabric"
  deploy        = false
  child_fabrics = ["child_fabric1"]
}

# Parent VRF resource - child_fabric2 attachments REMOVED
resource "ndfc_vrfs" "msd_parent_vrfs" {
  depends_on  = [ndfc_fabric_vxlan_msd.msd_fabric]
  fabric_name = ndfc_fabric_vxlan_msd.msd_fabric.fabric_name
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
      attach_list = {
        "9PJFRPZZ26Z" = {
          deploy_this_attachment = true
          fabric                 = tolist(ndfc_fabric_vxlan_msd.msd_fabric.child_fabrics)[0]
        }
        "9G4KUS84LFI" = {
          deploy_this_attachment = true
          fabric                 = tolist(ndfc_fabric_vxlan_msd.msd_fabric.child_fabrics)[0]
        }
      }
    }
  }
}

# child_fabric1 VRFs - still depends on parent MSD VRF resource
resource "ndfc_vrfs" "child_fabric1_vrfs" {
  depends_on  = [ndfc_vrfs.msd_parent_vrfs]
  fabric_name = tolist(ndfc_fabric_vxlan_msd.msd_fabric.child_fabrics)[0]
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

# child_fabric2 VRFs - NOW STANDALONE (depends on its own inventory, not MSD)
resource "ndfc_vrfs" "child_fabric2_vrfs" {
  depends_on  = [ndfc_inventory_devices.child_fabric2_inventory]
  fabric_name = "child_fabric2"
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

# Parent Network resource - child_fabric2 attachments REMOVED
resource "ndfc_networks" "msd_parent_networks" {
  depends_on  = [ndfc_vrfs.msd_parent_vrfs, ndfc_vrfs.child_fabric1_vrfs]
  fabric_name = ndfc_fabric_vxlan_msd.msd_fabric.fabric_name
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
      attach_list = {
        "9PJFRPZZ26Z" = {
          deploy_this_attachment = true
          fabric                 = tolist(ndfc_fabric_vxlan_msd.msd_fabric.child_fabrics)[0]
        }
        "9G4KUS84LFI" = {
          deploy_this_attachment = true
          fabric                 = tolist(ndfc_fabric_vxlan_msd.msd_fabric.child_fabrics)[0]
        }
      }
    }
  }
}

# child_fabric1 Networks - still depends on parent MSD Network resource
resource "ndfc_networks" "child_fabric1_networks" {
  depends_on  = [ndfc_networks.msd_parent_networks]
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

# child_fabric2 Networks - NOW STANDALONE (depends on its own VRF resource)
resource "ndfc_networks" "child_fabric2_networks" {
  depends_on  = [ndfc_vrfs.child_fabric2_vrfs]
  fabric_name = "child_fabric2"
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
