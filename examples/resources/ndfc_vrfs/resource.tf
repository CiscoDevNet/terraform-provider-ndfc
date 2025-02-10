
resource "ndfc_vrfs" "test_resource_vrfs_1" {
  fabric_name            = "CML"
  deploy_all_attachments = true
  vrfs = {
    "VRF1" = {
      vrf_template                   = "Default_VRF_Universal"
      vrf_extension_template         = "Default_VRF_Extension_Universal"
      vrf_id                         = 50000
      vlan_id                        = 1500
      vlan_name                      = "VLAN1500"
      interface_description          = "My int description"
      vrf_description                = "My vrf description"
      mtu                            = 9200
      loopback_routing_tag           = 11111
      redistribute_direct_route_map  = "FABRIC-RMAP-REDIST"
      max_bgp_paths                  = 2
      max_ibgp_paths                 = 3
      ipv6_link_local                = false
      trm                            = true
      no_rp                          = false
      rp_external                    = true
      rp_address                     = "1.2.3.4"
      rp_loopback_id                 = 500
      underlay_multicast_address     = "233.1.1.1"
      overlay_multicast_groups       = "234.0.0.0/8"
      mvpn_inter_as                  = false
      trm_bgw_msite                  = true
      advertise_host_routes          = true
      advertise_default_route        = false
      configure_static_default_route = false
      bgp_password                   = "1234567890ABCDEF"
      bgp_password_type              = "7"
      netflow                        = false
      netflow_monitor                = "MON1"
      disable_rt_auto                = true
      route_target_import            = "1:1"
      route_target_export            = "1:1"
      route_target_import_evpn       = "1:1"
      route_target_export_evpn       = "1:1"
      route_target_import_cloud_evpn = "1:1"
      route_target_export_cloud_evpn = "1:1"
      deploy_attachments             = false
      attach_list = {
        "FDO245206N5" = {
          vlan                   = 1500
          deploy_this_attachment = false
          loopback_id            = 101
          loopback_ipv4          = "1.2.3.4"
          loopback_ipv6          = "2001::1"
        }
      }

    }
  }

}