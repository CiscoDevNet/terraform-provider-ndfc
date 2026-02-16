# =============================================================================
# MSD Parent Fabric
# =============================================================================

variable "msd_fabric_name" {
  description = "Name of the MSD parent fabric"
  type        = string
}

variable "msd_fabric_config" {
  description = "All ndfc_fabric_vxlan_msd attributes (except fabric_name and child_fabrics)"
  type        = any
}

# =============================================================================
# Child Fabrics
# =============================================================================

variable "child_fabrics" {
  description = "Map of child fabric name to its ndfc_fabric_vxlan_evpn configuration. Key = fabric name."
  type        = map(any)
}

# =============================================================================
# Switch-to-Child Mapping
# =============================================================================

variable "switch_fabric_map" {
  description = "Flat map of switch serial number to child fabric name"
  type        = map(string)
}

# =============================================================================
# VRF Configuration
# =============================================================================

variable "vrf_configs" {
  description = "VRF configurations applied to MSD parent and all child fabrics"
  type = map(object({
    # Common attributes (parent + child)
    vlan_id                        = optional(number)
    vrf_id                         = optional(number)
    disable_rt_auto                = optional(bool)
    ipv6_link_local                = optional(bool)
    loopback_routing_tag           = optional(number)
    max_bgp_paths                  = optional(number)
    max_ibgp_paths                 = optional(number)
    mtu                            = optional(number)
    redistribute_direct_route_map  = optional(string)
    vrf_extension_template         = optional(string)
    vrf_template                   = optional(string)
    interface_description          = optional(string)
    vlan_name                      = optional(string)
    vrf_description                = optional(string)
    bgp_password                   = optional(string)
    bgp_password_type              = optional(string)
    configure_static_default_route = optional(bool)
    mvpn_inter_as                  = optional(bool)
    netflow                        = optional(bool)
    netflow_monitor                = optional(string)
    no_rp                          = optional(bool)
    overlay_multicast_groups       = optional(string)
    route_target_export            = optional(string)
    route_target_export_cloud_evpn = optional(string)
    route_target_export_evpn       = optional(string)
    route_target_export_mvpn       = optional(string)
    route_target_import            = optional(string)
    route_target_import_cloud_evpn = optional(string)
    route_target_import_evpn       = optional(string)
    route_target_import_mvpn       = optional(string)
    rp_address                     = optional(string)
    rp_external                    = optional(bool)
    rp_loopback_id                 = optional(number)
    trm                            = optional(bool)
    trm_bgw_msite                  = optional(bool)
    underlay_multicast_address     = optional(string)

    # Child-only attributes (excluded from parent, included in child)
    advertise_default_route = optional(bool)
    advertise_host_routes   = optional(bool)
  }))
  default = {}
}

# =============================================================================
# Network Configuration
# =============================================================================

variable "network_configs" {
  description = "Network configurations applied to MSD parent and all child fabrics"
  type = map(object({
    # Common attributes (parent + child)
    vrf_name                   = string
    network_id                 = number
    vlan_id                    = optional(number)
    vlan_name                  = optional(string)
    display_name               = optional(string)
    network_template           = optional(string)
    network_extension_template = optional(string)
    network_type               = optional(string)
    primary_network_id         = optional(number)
    gateway_ipv4_address       = optional(string)
    gateway_ipv6_address       = optional(string)
    layer2_only                = optional(bool)
    interface_description      = optional(string)
    mtu                        = optional(number)
    secondary_gateway_1        = optional(string)
    secondary_gateway_2        = optional(string)
    secondary_gateway_3        = optional(string)
    secondary_gateway_4        = optional(string)
    routing_tag                = optional(number)
    arp_suppression            = optional(bool)
    route_target_both          = optional(bool)
    ingress_replication        = optional(bool)

    # Child-only attributes (excluded from parent, included in child)
    dhcp_relay_loopback_id = optional(number)
    dhcp_relay_servers     = optional(list(object({ address = string, vrf = string })))
    multicast_group        = optional(string)
    trm                    = optional(bool)
    netflow                = optional(bool)
    svi_netflow_monitor    = optional(string)
    vlan_netflow_monitor   = optional(string)
    l3_gatway_border       = optional(bool)
    igmp_version           = optional(string)
  }))
  default = {}
}

