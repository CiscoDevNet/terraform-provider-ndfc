
resource "ndfc_interface_loopback" "test_resource_interface_loopback_1" {
  policy        = "int_loopback"
  deploy        = true
  serial_number = "9DBYO6WQJ46"
  interfaces = {
    "Loopback100" = {
      interface_name        = "loopback100"
      admin_state           = false
      interface_description = "This is a loopback interface used for XYZ"
      vrf                   = "default"
      ipv4_address          = "192.168.20.1"
      ipv6_address          = "2002:db8::1"
      route_map_tag         = "100"
    }
  }

}