# =============================================================================
# Shared VRF Configuration Variables
# These variables are shared between msd_fabric parent and child VRFs
# =============================================================================

variable "vrf_configs" {
  description = "Shared VRF configuration for both msd_fabric parent and child VRFs"
  type = map(object({
    vlan_id                       = number
    disable_rt_auto               = optional(bool, false)
    ipv6_link_local               = optional(bool, true)
    loopback_routing_tag          = optional(number, 12345)
    max_bgp_paths                 = optional(number, 1)
    max_ibgp_paths                = optional(number, 2)
    mtu                           = optional(number, 9216)
    redistribute_direct_route_map = optional(string, "FABRIC-RMAP-REDIST-SUBNET")
    vrf_extension_template        = optional(string, "Default_VRF_Extension_Universal")
    vrf_template                  = optional(string, "Default_VRF_Universal")
    interface_description         = optional(string)
    vlan_name                     = optional(string)
    vrf_description               = optional(string)
    # ndfc_vrfs specific attributes (not available in msd_fabric parent)
    advertise_default_route = optional(bool, true)
    advertise_host_routes   = optional(bool, true)
  }))
  default = {
    "vrf_01" = {
      vlan_id                       = 234
      disable_rt_auto               = false
      ipv6_link_local               = true
      loopback_routing_tag          = 12345
      max_bgp_paths                 = 1
      max_ibgp_paths                = 2
      mtu                           = 9216
      redistribute_direct_route_map = "FABRIC-RMAP-REDIST-SUBNET"
      vrf_extension_template        = "Default_VRF_Extension_Universal"
      vrf_template                  = "Default_VRF_Universal"
      advertise_default_route       = true
      advertise_host_routes         = true
    },
    "vrf_02" = {
      vlan_id                       = 24
      disable_rt_auto               = false
      ipv6_link_local               = true
      loopback_routing_tag          = 12345
      max_bgp_paths                 = 1
      max_ibgp_paths                = 2
      mtu                           = 9211
      redistribute_direct_route_map = "FABRIC-RMAP-REDIST-SUBNET"
      vrf_extension_template        = "Default_VRF_Extension_Universal"
      vrf_template                  = "Default_VRF_Universal"
      advertise_default_route       = true
      advertise_host_routes         = true
    }
  }
}

# =============================================================================
# Shared Network Configuration Variables
# These variables are shared between msd_fabric parent and child Networks.
# =============================================================================

variable "network_configs" {
  description = "Shared Network configuration for both msd_fabric parent and child Networks"
  type = map(object({
    # Parent properties: is_l2_only, vrf_name, net_id, vlan_id, vlan_name, int_desc,
    # gw_ip_address, gw_ipv6_address, secondary_ip_address, mtu_l3intf, route_tag,
    # arp_supress, route_target_both
    vrf_name              = string
    network_id            = number
    vlan_id               = number
    vlan_name             = optional(string)
    layer2_only           = optional(bool, false)
    interface_description = optional(string)
    gateway_ipv4_address  = optional(string)
    gateway_ipv6_address  = optional(string)
    secondary_gateway_1   = optional(string)
    secondary_gateway_2   = optional(string)
    secondary_gateway_3   = optional(string)
    secondary_gateway_4   = optional(string)
    mtu                   = optional(number, 9216)
    routing_tag           = optional(number)
    arp_suppression       = optional(bool)
    route_target_both     = optional(bool)
    # Child properties: dhcp_loopback_id, dhcp_servers, multicast_group_address,
    # trm_enable, netflow_enable, vlan_netflow_monitor, l3gw_on_border, igmp_version
    dhcp_loopback_id     = optional(number)
    dhcp_servers         = optional(list(object({ address = string, vrf = string })))
    multicast_group      = optional(string)
    trm                  = optional(bool)
    netflow              = optional(bool)
    vlan_netflow_monitor = optional(string)
    l3_gateway_on_border = optional(bool)
    igmp_version         = optional(number)
  }))
  default = {
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
      # Child properties
      trm                  = true
      multicast_group      = "239.1.1.1"
      igmp_version         = 3
      l3_gateway_on_border = true
    },
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
      # Child properties
      trm             = false
      multicast_group = "239.1.1.2"
      igmp_version    = 3
    }
  }
}
