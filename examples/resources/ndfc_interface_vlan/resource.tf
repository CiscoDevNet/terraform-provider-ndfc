
resource "ndfc_interface_vlan" "test_resource_interface_vlan_1" {
  policy        = "int_vlan"
  deploy        = true
  serial_number = "9DBYO6WQJ46"
  interfaces = {
    "Vlan1000" = {
      interface_name               = "Vlan1000"
      admin_state                  = false
      interface_description        = "My interface description"
      mtu                          = "1518"
      netflow                      = false
      netflow_monitor              = "MON1"
      netflow_sampler              = "SAMPLER1"
      vrf                          = "default"
      ipv4_address                 = "10.1.1.1"
      ipv4_prefix_length           = "24"
      routing_tag                  = "123"
      disable_ip_redirects         = false
      enable_hsrp                  = false
      preempt                      = false
      mac                          = "00:0a:00:00:00:98"
      dhcp_server_addr1            = "20.10.1.2"
      dhcp_server_addr2            = "29.1.2.1"
      dhcp_server_addr3            = "31.2.3.1"
      vrf_dhcp1                    = "mobile_net"
      vrf_dhcp2                    = "triple_play"
      vrf_dhcp3                    = "iptv"
      advertise_subnet_in_underlay = true
    }
  }

}