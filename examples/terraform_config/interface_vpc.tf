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

resource "ndfc_interface_vpc" "test_evpn_vxlan_vpc_intf1" {
  policy        = "int_vpc_trunk_host"
  deploy        = "true"
 
  interfaces = {
    "intf1" : {
      serial_number = "9TQYTJSZ1VJ~9Q34PHYLDB5"
      interface_name        = "vPC1"
      mtu                   = "default"
      peer1_port_channel_id = 2
      admin_state           = "true"
      mirror_config = "true"
      copy_po_description = "true"
      peer1_allowed_vlans         = "100-300"
      peer1_native_vlan           = 99
      speed                 = "Auto"
      portchannel_mode     = "passive"
      peer1_member_interfaces     = "Ethernet1/5"
      peer2_member_interfaces     = "Ethernet1/7"
    },
     "intf2" : {
      serial_number = "9TQYTJSZ1VJ~9Q34PHYLDB5"
      interface_name        = "vPC2"
      mtu                   = "jumbo"
      peer1_port_channel_id = 5
      admin_state           = "true"
      mirror_config = "true"
      copy_po_description = "true"
      peer1_allowed_vlans         = "100-300"
      peer1_native_vlan           = 99
      speed                 = "Auto"
      portchannel_mode     = "active"
      peer1_member_interfaces     = "Ethernet1/6"
      peer2_member_interfaces     = "Ethernet1/6"
    },
     "intf3" : {
      serial_number = "9TQYTJSZ1VJ~9Q34PHYLDB5"
      interface_name        = "vPC10"
      mtu                   = "jumbo"
      peer1_port_channel_id = 10
      admin_state           = "true"
      copy_po_description = "true"
      peer1_allowed_vlans         = "100-2000"
      peer1_native_vlan           = 99
      speed                 = "Auto"
      portchannel_mode     = "passive"
      peer1_member_interfaces     = "Ethernet1/10"
      peer2_member_interfaces     = "Ethernet1/12"
    }
  }
}
