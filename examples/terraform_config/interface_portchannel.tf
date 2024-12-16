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

resource "ndfc_interface_portchannel" "test_evpn_vxlan_pc_intf1" {
  policy        = "int_port_channel_trunk_host"
  deploy        = "true"
  serial_number = "9FE076D8EJL"

  interfaces = {
    "intf1" : {
      interface_name        = "port-channel605"
      mtu                   = "default"
      interface_description = "Port-channel605 interface"
      admin_state           = "true"
      allowed_vlans         = "100-300"
      native_vlan           = 99
      speed                 = "Auto"
      portchannel_mode      = "active"
      member_interfaces     = "Ethernet1/25,Ethernet1/26"
    },

    "intf2" : {
      interface_name        = "port-channel602"
      mtu                   = "jumbo"
      interface_description = "Port-channel602 interface"
      admin_state           = "false"
      allowed_vlans         = "101-299"
      native_vlan           = 199
      speed                 = "Auto"
      portchannel_mode      = "passive"
      member_interfaces     = "Ethernet1/31,Ethernet1/32"
    },
    "intf3" : {
      interface_name        = "port-channel601"
      mtu                   = "jumbo"
      interface_description = "Port-channel601 interface"
      admin_state           = "false"
      allowed_vlans         = "101-299"
      native_vlan           = 199
      speed                 = "Auto"
      portchannel_mode      = "passive"
      member_interfaces     = "Ethernet1/21,Ethernet1/22,Ethernet1/23-24"
    },
  }
}
