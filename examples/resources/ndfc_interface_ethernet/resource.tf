
resource "ndfc_interface_ethernet" "test_resource_interface_ethernet_1" {
  policy        = "int_access_host"
  policy_type   = "system"
  deploy        = true
  serial_number = "9DBYO6WQJ46"
  interfaces = {
    "Ethernet1/10" = {
      interface_name        = "Ethernet1/10"
      freeform_config       = "delay 200"
      admin_state           = false
      interface_description = "My interface description"
      bpdu_guard            = "true"
      port_type_fast        = false
      mtu                   = "default"
      speed                 = "Auto"
      access_vlan           = 500
      orphan_port           = false
      ptp                   = false
      netflow               = false
      netflow_monitor       = "MON1"
      netflow_sampler       = "SAMPLER1"
      allowed_vlans         = "10-20"
      native_vlan           = 1
    }
  }

}