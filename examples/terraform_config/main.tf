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
  host      = "https://10.104.251.69"
  insecure = true
}

/*
data "ndfc_vrf_bulk" "vrf_in_fabric" {
   fabric_name = "test_evpn_vxlan"
}


data "ndfc_fabric" "all_fabrics" {
   
}


output "vrfs_current" {
  value = data.ndfc_vrf_bulk.vrf_in_fabric
}

output "fabrics_current" {
  value = data.ndfc_fabric.all_fabrics
}
*/
resource "ndfc_vrf_bulk" "test_evpn_vxlan_vrfs1" {
  fabric_name = "test_evpn_vxlan"
  vrfs = [
      {

    vrf_name = "Murali_Bulk_000"
    vrf_id = 56599
    },
    {

    vrf_name = "Murali_Bulk_001"
    vrf_id = 56600
    },
    {
    vrf_name = "Murali_Bulk_002"
    vrf_id = 56601
    }
    
    
  ]
}

resource "ndfc_vrf_bulk" "test_evpn_vxlan_vrfs2" {
  fabric_name = "test_evpn_vxlan"
  vrfs = [
    {
      vrf_name = "Murali_Bulk_003"
      vrf_id = 56602
      vlan_id = 2039
      vlan_name = "03_Bulk"
    },
    {
      vrf_name = "Murali_Bulk_004"
      vrf_id = 56704
      vlan_id = 3092
      vlan_name = "04_Bulk"
    },
    {
      vrf_name = "Murali_Bulk_005"
      vrf_id = 56604
      vlan_id = 3096
      vlan_name = "55_Bulk"
    },
    {
      vrf_name = "Murali_Bulk_220"
      vrf_id = 56220
      vlan_id = 2039
      vlan_name = "220_Bulk"
    },
    {
      vrf_name = "Murali_long_list_1"
      vrf_extension_template         = "Default_VRF_Extension_Universal"
      vrf_id                         = 10001
      vlan_id                        = 1550
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

    }
  ]
}



