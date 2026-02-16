resource "ndfc_vrfs" "msd_parent_vrfs_attributes" {
  fabric_name = ndfc_fabric_vxlan_msd.msd_fabric.fabric_name # msd_fabric is the parent fabric created earlier
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
      interface_description         = vrf_config.interface_description
      vlan_name                     = vrf_config.vlan_name
      vrf_description               = vrf_config.vrf_description
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

resource "ndfc_vrfs" "child1_specific_vrfs_attributes" {
  depends_on  = [ndfc_vrfs.msd_parent_vrfs_attributes]
  fabric_name = tolist(ndfc_fabric_vxlan_msd.msd_fabric.child_fabrics)[0]
  vrfs = {
    for vrf_name, vrf_config in var.vrf_configs : vrf_name => {
      disable_rt_auto               = vrf_config.disable_rt_auto
      ipv6_link_local               = vrf_config.ipv6_link_local
      loopback_routing_tag          = vrf_config.loopback_routing_tag
      max_bgp_paths                 = vrf_config.max_bgp_paths
      max_ibgp_paths                = vrf_config.max_ibgp_paths
      mtu                           = vrf_config.mtu
      redistribute_direct_route_map = vrf_config.redistribute_direct_route_map
      vrf_extension_template        = vrf_config.vrf_extension_template
      vrf_template                  = vrf_config.vrf_template
      interface_description         = vrf_config.interface_description
      vlan_name                     = vrf_config.vlan_name
      vrf_description               = vrf_config.vrf_description
      advertise_default_route       = vrf_config.advertise_default_route
      advertise_host_routes         = vrf_config.advertise_host_routes
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

resource "ndfc_vrfs" "child2_specific_vrfs_attributes" {
  depends_on  = [ndfc_vrfs.msd_parent_vrfs_attributes]
  fabric_name = tolist(ndfc_fabric_vxlan_msd.msd_fabric.child_fabrics)[1]
  vrfs = {
    for vrf_name, vrf_config in var.vrf_configs : vrf_name => {
      disable_rt_auto               = vrf_config.disable_rt_auto
      ipv6_link_local               = vrf_config.ipv6_link_local
      loopback_routing_tag          = vrf_config.loopback_routing_tag
      max_bgp_paths                 = vrf_config.max_bgp_paths
      max_ibgp_paths                = vrf_config.max_ibgp_paths
      mtu                           = vrf_config.mtu
      redistribute_direct_route_map = vrf_config.redistribute_direct_route_map
      vrf_extension_template        = vrf_config.vrf_extension_template
      vrf_template                  = vrf_config.vrf_template
      interface_description         = vrf_config.interface_description
      vlan_name                     = vrf_config.vlan_name
      vrf_description               = vrf_config.vrf_description
      advertise_default_route       = vrf_config.advertise_default_route
      advertise_host_routes         = vrf_config.advertise_host_routes
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
