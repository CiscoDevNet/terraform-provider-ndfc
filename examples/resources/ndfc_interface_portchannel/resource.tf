
resource "ndfc_interface_portchannel" "test_resource_interface_portchannel_1" {
  policy        = "int_port_channel_trunk_host"
  deploy        = true
  serial_number = "9DBYO6WQJ46"
  interfaces = {
    "Port-channel100" = {
      interface_name        = "port-channel100"
      admin_state           = false
      interface_description = "My interface description"
      bpdu_guard            = "true"
      port_type_fast        = false
      mtu                   = "default"
      speed                 = "Auto"
      orphan_port           = false
      netflow               = false
      netflow_monitor       = "MON1"
      netflow_sampler       = "SAMPLER1"
      allowed_vlans         = "10-20"
      native_vlan           = 1
      copy_po_description   = false
      portchannel_mode      = "on"
      member_interfaces     = "eth1/5-6"
    }
  }

}