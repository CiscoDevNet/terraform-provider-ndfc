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

resource "ndfc_interface_ethernet" "test_evpn_vxlan_eth_intf1" {
  policy        = "int_trunk_host"
  deploy        = "true"
  serial_number = "9FE076D8EJL"

  interfaces = {
    "intf1" : {
      interface_name        = "Ethernet1/1"
      mtu                   = "jumbo"
      interface_description = "test 111111111 interface"
      admin_state           = "true"
      allowed_vlans         = "100-300"
      native_vlan           = 99
      speed                 = "Auto"
    },

    "intf4" : {
      interface_name        = "Ethernet1/2"
      mtu                   = "jumbo"
      interface_description = "test 4 interface"
      admin_state           = "true"
      allowed_vlans         = "101-299"
      native_vlan           = 199
      speed                 = "Auto"
    },

    "intf6" : {
      interface_name        = "Ethernet1/3"
      mtu                   = "jumbo"
      interface_description = "test 6 interface"
      admin_state           = "true"
      allowed_vlans         = "101-299"
      native_vlan           = 200
      speed                 = "Auto"
    },

    "intf8" : {
      interface_name        = "Ethernet1/4"
      mtu                   = "jumbo"
      interface_description = "test 8 interface"
      admin_state           = "true"
      allowed_vlans         = "101-299"
      native_vlan           = 200
      speed                 = "Auto"
    },
    "intf9" : {
      interface_name        = "Ethernet1/5"
      mtu                   = "jumbo"
      interface_description = "test 9 interface"
      admin_state           = "true"
      allowed_vlans         = "101-299"
      native_vlan           = 200
      speed                 = "Auto"
    },

    "intf12" : {
      interface_name        = "Ethernet1/9"
      mtu                   = "jumbo"
      interface_description = "test 12 interface"
      admin_state           = "true"
      allowed_vlans         = "101-299"
      native_vlan           = 200
      speed                 = "Auto"
    },
    "intf14" : {

      interface_name        = "Ethernet1/19"
      mtu                   = "jumbo"
      interface_description = "test 19 interface"
      admin_state           = "true"
      allowed_vlans         = "20-299"
      native_vlan           = 201
      speed                 = "Auto"
    },

  }
}

resource "ndfc_interface_ethernet" "test_evpn_vxlan_eth_intf2" {
  policy = "int_trunk_host"
  deploy = "true"
  interfaces = {
    "intf1" : {
      serial_number         = "9TQYTJSZ1VJ"
      interface_name        = "Ethernet1/11"
      mtu                   = "jumbo"
      interface_description = "test 1111 interface"
      admin_state           = "true"
      allowed_vlans         = "100-300"
      native_vlan           = 99
      speed                 = "Auto"
    },
    "intf4" : {
      serial_number         = "9TQYTJSZ1VJ"
      interface_name        = "Ethernet1/12"
      mtu                   = "jumbo"
      interface_description = "test 4 interface"
      admin_state           = "true"
      allowed_vlans         = "101-299"
      native_vlan           = 199
      speed                 = "Auto"
    },
    "intf6" : {
      serial_number         = "9TQYTJSZ1VJ"
      interface_name        = "Ethernet1/13"
      mtu                   = "jumbo"
      interface_description = "test 6 interface"
      admin_state           = "true"
      allowed_vlans         = "101-299"
      native_vlan           = 200
      speed                 = "Auto"
    },
    "intf8" : {
      serial_number         = "9TQYTJSZ1VJ"
      interface_name        = "Ethernet1/14"
      mtu                   = "jumbo"
      interface_description = "test 8 interface"
      admin_state           = "true"
      allowed_vlans         = "101-299"
      native_vlan           = 200
      speed                 = "Auto"
    },
    "intf9" : {
      serial_number         = "9TQYTJSZ1VJ"
      interface_name        = "Ethernet1/15"
      mtu                   = "jumbo"
      interface_description = "test 9 interface"
      admin_state           = "true"
      allowed_vlans         = "101-299"
      native_vlan           = 200
      speed                 = "Auto"
    },
    "intf12" : {
      serial_number         = "9TQYTJSZ1VJ"
      interface_name        = "Ethernet1/16"
      mtu                   = "jumbo"
      interface_description = "test 12 interface"
      admin_state           = "true"
      allowed_vlans         = "101-299"
      native_vlan           = 200
      speed                 = "Auto"
    },
    "intf13" : {
      serial_number         = "9TQYTJSZ1VJ"
      interface_name        = "Ethernet1/17"
      mtu                   = "jumbo"
      interface_description = "test 13 interface"
      admin_state           = "true"
      allowed_vlans         = "101-299"
      native_vlan           = 200
      speed                 = "Auto"
    },

  }
}

