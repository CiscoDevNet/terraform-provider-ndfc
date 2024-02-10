terraform {
  required_providers {
    ndfc = {
      source = "registry.terraform.io/cisco/ndfc"
    }
  }
}

provider "ndfc" {
  username = "admin"
  password = "admin!@#"
#  password = "ins3965!"
#  host      = "https://10.195.225.193"
  host      = "https://10.78.210.161"
  insecure = true
}



resource "ndfc_vrf_bulk" "test_evpn_vxlan_deployments1" {
  fabric_name = "test_evpn_vxlan"
  vrfs = [
    {
      vrf_name                       = "MyVRF_60055"
      vrf_template                   = "Default_VRF_Universal"
      vrf_extension_template         = "Default_VRF_Extension_Universal"
      vrf_id                         = 10022
      vlan_id                        = 1501
      vlan_name                      = "VLAN1501"
      interface_description          = "My int description"
      vrf_description                = "My vrf description"
      mtu                            = 9201
      loopback_routing_tag           = 11111
      redistribute_direct_route_map  = "FABRIC-RMAP-REDIST"
      max_bgp_paths                  = 2
      max_ibgp_paths                 = 3
      ipv6_link_local                = false
      trm                            = true
      no_rp                          = false
      rp_external                    = true
      rp_address                     = "1.2.3.4"
      rp_loopback_id                 = 100
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
    },
    {
      vrf_name = "Murali_Bulk_001"
    },
    {
      vrf_name = "Murali_Bulk_002"
    },
    {
      vrf_name                       = "MyVRF_60066"
      vrf_template                   = "Default_VRF_Universal"
      vrf_extension_template         = "Default_VRF_Extension_Universal"
      vlan_id                        = 1501
      vlan_name                      = "VLAN1501"
      interface_description          = "My int description"
      vrf_description                = "My vrf description"
      mtu                            = 9201
      loopback_routing_tag           = 11111
      redistribute_direct_route_map  = "FABRIC-RMAP-REDIST"
      max_bgp_paths                  = 2
      max_ibgp_paths                 = 3
      ipv6_link_local                = false
      trm                            = true
      no_rp                          = false
      rp_external                    = true
      rp_address                     = "1.2.3.5"
      rp_loopback_id                 = 101
      underlay_multicast_address     = "233.1.1.2"
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
      attach_list = [
      {
         serial_number = "9TQYTJSZ1VJ"
         deploy_this_attachment = true
      },
      {
        serial_number = "9QBCTIN0FMY"
        deploy_this_attachment = true
      }
      ]
    }

  ]
}
