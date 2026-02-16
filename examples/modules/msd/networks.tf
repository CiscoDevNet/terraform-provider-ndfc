# =============================================================================
# Networks at MSD Parent Level
# Parent gets common attributes only. Child-only attrs excluded.
# Attachments include ALL switches with 'fabric' field.
# =============================================================================

resource "ndfc_networks" "parent" {
  depends_on  = [ndfc_vrfs.parent, ndfc_vrfs.child]
  fabric_name = ndfc_fabric_vxlan_msd.parent.fabric_name

  networks = {
    for net_name, c in var.network_configs : net_name => {
      # Common attributes only (child-only attributes excluded)
      vrf_name                   = c.vrf_name
      network_id                 = c.network_id
      vlan_id                    = c.vlan_id
      vlan_name                  = c.vlan_name
      display_name               = c.display_name
      network_template           = c.network_template
      network_extension_template = c.network_extension_template
      network_type               = c.network_type
      primary_network_id         = c.primary_network_id
      gateway_ipv4_address       = c.gateway_ipv4_address
      gateway_ipv6_address       = c.gateway_ipv6_address
      layer2_only                = c.layer2_only
      interface_description      = c.interface_description
      mtu                        = c.mtu
      secondary_gateway_1        = c.secondary_gateway_1
      secondary_gateway_2        = c.secondary_gateway_2
      secondary_gateway_3        = c.secondary_gateway_3
      secondary_gateway_4        = c.secondary_gateway_4
      routing_tag                = c.routing_tag
      arp_suppression            = c.arp_suppression
      route_target_both          = c.route_target_both
      ingress_replication        = c.ingress_replication

      # NOTE: child-only attrs excluded at parent level

      deploy_attachments = true

      # Parent attachments: ALL switches with fabric field
      attachments = local.parent_net_attach
    }
  }
}

# =============================================================================
# Networks at Each Child Fabric Level
# Gets common + child-only attributes.
# Attachments include only that child's switches, WITHOUT 'fabric' field.
# =============================================================================

resource "ndfc_networks" "child" {
  for_each = var.child_fabrics

  depends_on  = [ndfc_networks.parent]
  fabric_name = each.key

  networks = {
    for net_name, c in var.network_configs : net_name => {
      # Common attributes
      vrf_name                   = c.vrf_name
      network_id                 = c.network_id
      vlan_id                    = c.vlan_id
      vlan_name                  = c.vlan_name
      display_name               = c.display_name
      network_template           = c.network_template
      network_extension_template = c.network_extension_template
      network_type               = c.network_type
      primary_network_id         = c.primary_network_id
      gateway_ipv4_address       = c.gateway_ipv4_address
      gateway_ipv6_address       = c.gateway_ipv6_address
      layer2_only                = c.layer2_only
      interface_description      = c.interface_description
      mtu                        = c.mtu
      secondary_gateway_1        = c.secondary_gateway_1
      secondary_gateway_2        = c.secondary_gateway_2
      secondary_gateway_3        = c.secondary_gateway_3
      secondary_gateway_4        = c.secondary_gateway_4
      routing_tag                = c.routing_tag
      arp_suppression            = c.arp_suppression
      route_target_both          = c.route_target_both
      ingress_replication        = c.ingress_replication

      # Child-only attributes
      dhcp_relay_loopback_id = c.dhcp_relay_loopback_id
      dhcp_relay_servers     = c.dhcp_relay_servers
      multicast_group        = c.multicast_group
      trm                    = c.trm
      netflow                = c.netflow
      svi_netflow_monitor    = c.svi_netflow_monitor
      vlan_netflow_monitor   = c.vlan_netflow_monitor
      l3_gatway_border       = c.l3_gatway_border
      igmp_version           = c.igmp_version

      deploy_attachments = true

      # Only this child's switches, no fabric field
      attachments = {
        for serial in local.switches_by_child[each.key] : serial => {
        }
      }
    }
  }
}
