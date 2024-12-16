
resource "ndfc_interface_vpc" "test_resource_interface_vpc_1" {
  policy        = "int_vpc_trunk_host"
  deploy        = true
  serial_number = "9TQYTJSZ1VJ~9Q34PHYLDB5"
  interfaces = {
    "vPC1" = {
      interface_name           = "vPC1"
      admin_state              = false
      bpdu_guard               = "true"
      port_type_fast           = false
      mtu                      = "default"
      speed                    = "Auto"
      netflow                  = false
      netflow_monitor          = "MON1"
      netflow_sampler          = "SAMPLER1"
      copy_po_description      = false
      portchannel_mode         = "on"
      peer1_po_freeform_config = "delay 200"
      peer2_po_freeform_config = "delay 200"
      peer1_po_description     = "My interface description"
      peer2_po_description     = "My interface description"
      peer1_allowed_vlans      = "10-20"
      peer2_allowed_vlans      = "10-20"
      peer1_native_vlan        = 1
      peer2_native_vlan        = 1
      peer1_member_interfaces  = "eth1/15"
      peer2_member_interfaces  = "eth1/15"
      peer1_port_channel_id    = 120
      peer2_port_channel_id    = 120
    }
  }

}