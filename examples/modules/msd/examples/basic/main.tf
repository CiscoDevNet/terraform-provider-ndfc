module "msd_dc" {
  source = "../../"

  # ── MSD Parent ──
  msd_fabric_name = "MSD_PROD"
  msd_fabric_config = {
    deploy                 = false
    anycast_gw_mac         = "2020.0000.00aa"
    bgw_routing_tag        = 54321
    border_gwy_connections = "Manual"
    dci_subnet_range       = "10.10.1.0/24"
    dci_subnet_target_mask = 30
    delay_restore          = 300
    l2_segment_id_range    = "30000-49000"
    l3_partition_id_range  = "50000-59000"
    loopback100_ip_range   = "10.10.0.0/24"
    ms_loopback_id         = 100
    default_network        = "Default_Network_Universal"
    default_vrf            = "Default_VRF_Universal"
  }

  # ── Child Fabrics ──
  child_fabrics = {
    "CHILD_DC1" = {
      bgp_as             = "65001"
      deploy             = false
      anycast_gw_mac     = "2020.0000.00aa"
      loopback0_ip_range = "10.2.0.0/22"
      loopback1_ip_range = "10.3.0.0/22"
      subnet_range       = "10.4.0.0/16"
      subnet_target_mask = 30
      replication_mode   = "Multicast"
    }
    "CHILD_DC2" = {
      bgp_as             = "65002"
      deploy             = false
      anycast_gw_mac     = "2020.0000.00aa"
      loopback0_ip_range = "10.5.0.0/22"
      loopback1_ip_range = "10.6.0.0/22"
      subnet_range       = "10.7.0.0/16"
      subnet_target_mask = 30
      replication_mode   = "Multicast"
    }
    "CHILD_DC3" = {
      bgp_as             = "65003"
      deploy             = false
      anycast_gw_mac     = "2020.0000.00aa"
      loopback0_ip_range = "10.8.0.0/22"
      loopback1_ip_range = "10.9.0.0/22"
      subnet_range       = "10.10.0.0/16"
      subnet_target_mask = 30
      replication_mode   = "Multicast"
    }
  }

  # ── Switch → Child mapping (flat map) ──
  switch_fabric_map = {
    "9PJFRPZZ26Z" = "CHILD_DC1"
    "9G4KUS84LFI" = "CHILD_DC1"
    "92Z298RRNPO" = "CHILD_DC2"
    "9ZHWL1HRWKR" = "CHILD_DC2"
    "8ABCD1234EF" = "CHILD_DC3"
    "8WXYZ5678GH" = "CHILD_DC3"
  }

  # ── VRFs (applied to parent + all children) ──
  vrf_configs = {
    "vrf_prod" = {
      vlan_id                       = 234
      mtu                           = 9216
      ipv6_link_local               = true
      loopback_routing_tag          = 12345
      max_bgp_paths                 = 1
      max_ibgp_paths                = 2
      redistribute_direct_route_map = "FABRIC-RMAP-REDIST-SUBNET"
      vrf_template                  = "Default_VRF_Universal"
      vrf_extension_template        = "Default_VRF_Extension_Universal"
      # child-only
      advertise_default_route = true
      advertise_host_routes   = true
    }
    "vrf_dev" = {
      vlan_id                 = 235
      mtu                     = 9211
      advertise_default_route = false
      advertise_host_routes   = true
    }
  }

  # ── Networks (applied to parent + all children) ──
  network_configs = {
    "nw_prod" = {
      vrf_name             = "vrf_prod"
      network_id           = 31001
      vlan_id              = 301
      gateway_ipv4_address = "192.168.1.1/24"
      gateway_ipv6_address = "2001::2/64"
      mtu                  = 9100
      routing_tag          = 12345
      # child-only
      trm              = true
      multicast_group  = "239.1.1.1"
      igmp_version     = "3"
      l3_gatway_border = true
    }
    "nw_dev" = {
      vrf_name             = "vrf_dev"
      network_id           = 31002
      vlan_id              = 302
      gateway_ipv4_address = "192.168.2.1/24"
      mtu                  = 9100
      routing_tag          = 12345
    }
  }

  # ── Optional deploy ──
  deploy = false
}
