
resource "ndfc_msd_vrfs_parent" "test_resource_msd_vrfs_parent_1" {
  fabric_name = "CML"
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
      rp_external                    = true
      mvpn_inter_as                  = false
      disable_rt_auto                = true
      route_target_import            = "1:1"
      route_target_export            = "1:1"
      route_target_import_evpn       = "1:1"
      route_target_export_evpn       = "1:1"
      route_target_import_cloud_evpn = "1:1"
      route_target_export_cloud_evpn = "1:1"
    }
  }

}