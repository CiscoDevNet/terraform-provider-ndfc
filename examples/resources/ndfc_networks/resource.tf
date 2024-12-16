
resource "ndfc_networks" "test_resource_networks_1" {
  fabric_name            = "CML"
  deploy_all_attachments = true
  networks = {
    "NET1" = {
      display_name               = "NET1"
      network_id                 = 30001
      network_template           = "Default_Network_Universal"
      network_extension_template = "Default_Network_Extension_Universal"
      vrf_name                   = "VRF1"
      primary_network_id         = 30000
      network_type               = "Normal"
      gateway_ipv4_address       = "192.0.2.1/24"
      gateway_ipv6_address       = "2001:db8::1/64"
      vlan_id                    = 1600
      vlan_name                  = "VLAN2000"
      layer2_only                = false
      interface_description      = "My int description"
      mtu                        = 9200
      secondary_gateway_1        = "192.168.2.1/24"
      secondary_gateway_2        = "192.168.3.1/24"
      secondary_gateway_3        = "192.168.4.1/24"
      secondary_gateway_4        = "192.168.5.1/24"
      arp_suppression            = false
      ingress_replication        = false
      multicast_group            = "233.1.1.1"
      dhcp_relay_loopback_id     = 134
      routing_tag                = 11111
      trm                        = true
      route_target_both          = true
      netflow                    = false
      svi_netflow_monitor        = "MON1"
      vlan_netflow_monitor       = "MON1"
      l3_gatway_border           = true
      igmp_version               = "3"
      deploy_attachments         = false
      attachments = {
        "FDO245206N5" = {
          vlan                   = 1600
          deploy_this_attachment = false
          switch_ports           = ["Ethernet1/1", "Ethernet1/2"]
        }
      }

    }
  }

}