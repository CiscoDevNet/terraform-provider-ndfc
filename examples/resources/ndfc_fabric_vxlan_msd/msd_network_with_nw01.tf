# =============================================================================
# BEFORE: MSD fabric with nw_01 network in parent and child resources
# =============================================================================

# Parent Network resource with nw_01
resource "ndfc_networks" "msd_parent_networks" {
  depends_on  = [ndfc_vrfs.msd_parent_vrfs, ndfc_vrfs.child_fabric1_vrfs, ndfc_vrfs.child_fabric2_vrfs]
  fabric_name = ndfc_fabric_vxlan_msd.msd_fabric.fabric_name
  networks = {
    "nw_01" = {
      vrf_name              = "vrf_01"
      network_id            = 31001
      gateway_ipv4_address  = "192.168.1.1/24"
      gateway_ipv6_address  = "2001::2/64"
      vlan_id               = 301
      vlan_name             = "nw_01"
      layer2_only           = false
      interface_description = "nw_01 SVI"
      mtu                   = 9100
      routing_tag           = 12345
      attach_list = {
        "9PJFRPZZ26Z" = {
          deploy_this_attachment = true
          fabric                 = "child_fabric1"
        }
        "9G4KUS84LFI" = {
          deploy_this_attachment = true
          fabric                 = "child_fabric1"
        }
        "92Z298RRNPO" = {
          deploy_this_attachment = true
          fabric                 = "child_fabric2"
        }
        "9ZHWL1HRWKR" = {
          deploy_this_attachment = true
          fabric                 = "child_fabric2"
        }
      }
    }
    "nw_02" = {
      vrf_name              = "vrf_02"
      network_id            = 31002
      gateway_ipv4_address  = "192.168.2.1/24"
      gateway_ipv6_address  = "2002::1/64"
      vlan_id               = 32
      vlan_name             = "nw_02"
      layer2_only           = false
      interface_description = "nw_02 SVI"
      mtu                   = 9100
      routing_tag           = 12345
      attach_list = {
        "9PJFRPZZ26Z" = {
          deploy_this_attachment = true
          fabric                 = "child_fabric1"
        }
        "9G4KUS84LFI" = {
          deploy_this_attachment = true
          fabric                 = "child_fabric1"
        }
        "92Z298RRNPO" = {
          deploy_this_attachment = true
          fabric                 = "child_fabric2"
        }
        "9ZHWL1HRWKR" = {
          deploy_this_attachment = true
          fabric                 = "child_fabric2"
        }
      }
    }
  }
}

# child_fabric1 Networks with nw_01
resource "ndfc_networks" "child_fabric1_networks" {
  depends_on  = [ndfc_networks.msd_parent_networks]
  fabric_name = "child_fabric1"
  networks = {
    "nw_01" = {
      vrf_name              = "vrf_01"
      network_id            = 31001
      gateway_ipv4_address  = "192.168.1.1/24"
      gateway_ipv6_address  = "2001::2/64"
      vlan_id               = 301
      vlan_name             = "nw_01"
      layer2_only           = false
      interface_description = "nw_01 SVI"
      mtu                   = 9100
      routing_tag           = 12345
      trm                   = true
      multicast_group       = "239.1.1.1"
      igmp_version          = 3
      l3_gateway_on_border  = true
      attach_list = {
        "9PJFRPZZ26Z" = { deploy_this_attachment = true }
        "9G4KUS84LFI" = { deploy_this_attachment = true }
      }
    }
    "nw_02" = {
      vrf_name              = "vrf_02"
      network_id            = 31002
      gateway_ipv4_address  = "192.168.2.1/24"
      gateway_ipv6_address  = "2002::1/64"
      vlan_id               = 32
      vlan_name             = "nw_02"
      layer2_only           = false
      interface_description = "nw_02 SVI"
      mtu                   = 9100
      routing_tag           = 12345
      trm                   = false
      multicast_group       = "239.1.1.2"
      igmp_version          = 3
      attach_list = {
        "9PJFRPZZ26Z" = { deploy_this_attachment = true }
        "9G4KUS84LFI" = { deploy_this_attachment = true }
      }
    }
  }
}

# child_fabric2 Networks with nw_01
resource "ndfc_networks" "child_fabric2_networks" {
  depends_on  = [ndfc_networks.msd_parent_networks]
  fabric_name = "child_fabric2"
  networks = {
    "nw_01" = {
      vrf_name              = "vrf_01"
      network_id            = 31001
      gateway_ipv4_address  = "192.168.1.1/24"
      gateway_ipv6_address  = "2001::2/64"
      vlan_id               = 301
      vlan_name             = "nw_01"
      layer2_only           = false
      interface_description = "nw_01 SVI"
      mtu                   = 9100
      routing_tag           = 12345
      trm                   = true
      multicast_group       = "239.1.1.1"
      igmp_version          = 3
      l3_gateway_on_border  = true
      attach_list = {
        "92Z298RRNPO" = { deploy_this_attachment = true }
        "9ZHWL1HRWKR" = { deploy_this_attachment = true }
      }
    }
    "nw_02" = {
      vrf_name              = "vrf_02"
      network_id            = 31002
      gateway_ipv4_address  = "192.168.2.1/24"
      gateway_ipv6_address  = "2002::1/64"
      vlan_id               = 32
      vlan_name             = "nw_02"
      layer2_only           = false
      interface_description = "nw_02 SVI"
      mtu                   = 9100
      routing_tag           = 12345
      trm                   = false
      multicast_group       = "239.1.1.2"
      igmp_version          = 3
      attach_list = {
        "92Z298RRNPO" = { deploy_this_attachment = true }
        "9ZHWL1HRWKR" = { deploy_this_attachment = true }
      }
    }
  }
}
