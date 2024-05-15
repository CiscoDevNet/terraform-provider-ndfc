terraform {
  required_providers {
    ndfc = {
      source = "registry.terraform.io/cisco/ndfc"
    }
  }
}

provider "ndfc" {
  username = "test"
  password = "test"
  host     = "https://10.78.210.161"
  insecure = true
}

resource "ndfc_interface_loopback" "loopback_test_evpn_vxlan_intf1" {
  policy        = "int_loopback"
  deploy        = "true"
  interfaces = {
    "loopback1" : {
      serial_number = "9FE076D8EJL"
      interface_name        = "loopback12"
      interface_description = "test loopback 12"
      admin_state          = "true"
      ipv4_address         = "10.1.1.10"

    },
    "loopback2" : {
      serial_number = "9FE076D8EJL"
      interface_name        = "loopback13"
      interface_description = "test loopback 13"
      admin_state          = "true"
      ipv4_address         = "10.2.1.1"
    },
    "loopback4" : {
      serial_number = "9FE076D8EJL"
      interface_name        = "Loopback15"
      interface_description = "test loopback 15"
      admin_state          = "true"
      ipv4_address         = "10.4.1.2"
    },
     "loopback5" : {
      serial_number = "9FE076D8EJL"
      interface_name        = "Loopback16"
      interface_description = "test loopback 16"
      admin_state          = "true"
      ipv4_address         = "10.5.1.2"
    }
  }
}
