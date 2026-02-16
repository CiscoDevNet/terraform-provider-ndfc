# =============================================================================
# VRFs at MSD Parent Level
# Parent gets common attributes only. Child-only attrs excluded.
# Attachments include ALL switches with 'fabric' field.
# =============================================================================

resource "ndfc_vrfs" "parent" {
  fabric_name = ndfc_fabric_vxlan_msd.parent.fabric_name

  vrfs = {
    for vrf_name, c in var.vrf_configs : vrf_name => {
      # Common attributes only (child-only attributes excluded at parent level)
      vlan_id                        = c.vlan_id
      vrf_id                         = c.vrf_id
      disable_rt_auto                = c.disable_rt_auto
      ipv6_link_local                = c.ipv6_link_local
      loopback_routing_tag           = c.loopback_routing_tag
      max_bgp_paths                  = c.max_bgp_paths
      max_ibgp_paths                 = c.max_ibgp_paths
      mtu                            = c.mtu
      redistribute_direct_route_map  = c.redistribute_direct_route_map
      vrf_extension_template         = c.vrf_extension_template
      vrf_template                   = c.vrf_template
      interface_description          = c.interface_description
      vlan_name                      = c.vlan_name
      vrf_description                = c.vrf_description
      bgp_password                   = c.bgp_password
      bgp_password_type              = c.bgp_password_type
      configure_static_default_route = c.configure_static_default_route
      mvpn_inter_as                  = c.mvpn_inter_as
      netflow                        = c.netflow
      netflow_monitor                = c.netflow_monitor
      no_rp                          = c.no_rp
      overlay_multicast_groups       = c.overlay_multicast_groups
      route_target_export            = c.route_target_export
      route_target_export_cloud_evpn = c.route_target_export_cloud_evpn
      route_target_export_evpn       = c.route_target_export_evpn
      route_target_export_mvpn       = c.route_target_export_mvpn
      route_target_import            = c.route_target_import
      route_target_import_cloud_evpn = c.route_target_import_cloud_evpn
      route_target_import_evpn       = c.route_target_import_evpn
      route_target_import_mvpn       = c.route_target_import_mvpn
      rp_address                     = c.rp_address
      rp_external                    = c.rp_external
      rp_loopback_id                 = c.rp_loopback_id
      trm                            = c.trm
      trm_bgw_msite                  = c.trm_bgw_msite
      underlay_multicast_address     = c.underlay_multicast_address

      # NOTE: advertise_default_route, advertise_host_routes excluded (child-only)

      deploy_attachments = true

      # Parent attachments: ALL switches with fabric field
      attach_list = local.parent_vrf_attach
    }
  }
}

# =============================================================================
# VRFs at Each Child Fabric Level
# Gets common + child-only attributes.
# Attachments include only that child's switches, WITHOUT 'fabric' field.
# =============================================================================

resource "ndfc_vrfs" "child" {
  for_each = var.child_fabrics

  depends_on  = [ndfc_vrfs.parent]
  fabric_name = each.key

  vrfs = {
    for vrf_name, c in var.vrf_configs : vrf_name => {
      # Common attributes (same as parent)
      vlan_id                        = c.vlan_id
      vrf_id                         = c.vrf_id
      disable_rt_auto                = c.disable_rt_auto
      ipv6_link_local                = c.ipv6_link_local
      loopback_routing_tag           = c.loopback_routing_tag
      max_bgp_paths                  = c.max_bgp_paths
      max_ibgp_paths                 = c.max_ibgp_paths
      mtu                            = c.mtu
      redistribute_direct_route_map  = c.redistribute_direct_route_map
      vrf_extension_template         = c.vrf_extension_template
      vrf_template                   = c.vrf_template
      interface_description          = c.interface_description
      vlan_name                      = c.vlan_name
      vrf_description                = c.vrf_description
      bgp_password                   = c.bgp_password
      bgp_password_type              = c.bgp_password_type
      configure_static_default_route = c.configure_static_default_route
      mvpn_inter_as                  = c.mvpn_inter_as
      netflow                        = c.netflow
      netflow_monitor                = c.netflow_monitor
      no_rp                          = c.no_rp
      overlay_multicast_groups       = c.overlay_multicast_groups
      route_target_export            = c.route_target_export
      route_target_export_cloud_evpn = c.route_target_export_cloud_evpn
      route_target_export_evpn       = c.route_target_export_evpn
      route_target_export_mvpn       = c.route_target_export_mvpn
      route_target_import            = c.route_target_import
      route_target_import_cloud_evpn = c.route_target_import_cloud_evpn
      route_target_import_evpn       = c.route_target_import_evpn
      route_target_import_mvpn       = c.route_target_import_mvpn
      rp_address                     = c.rp_address
      rp_external                    = c.rp_external
      rp_loopback_id                 = c.rp_loopback_id
      trm                            = c.trm
      trm_bgw_msite                  = c.trm_bgw_msite
      underlay_multicast_address     = c.underlay_multicast_address

      # Child-only attributes
      advertise_default_route = c.advertise_default_route
      advertise_host_routes   = c.advertise_host_routes

      deploy_attachments = true

      # Only this child's switches, no fabric field
      attach_list = {
        for serial in local.switches_by_child[each.key] : serial => {
        }
      }
    }
  }
}
