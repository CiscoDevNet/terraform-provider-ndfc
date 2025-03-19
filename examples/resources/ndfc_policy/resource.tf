
resource "ndfc_policy" "test_resource_policy_1" {
  is_policy_group      = false
  deploy               = true
  entity_name          = "Switch"
  entity_type          = "SWITCH"
  description          = "Policy for switch"
  template_name        = "TelemetryDst_EF"
  source               = "CLI"
  priority             = 500
  device_serial_number = "FDO245206N5"
  policy_parameters = {
    DSTGRP = "501"
    IPADDR = "5.5.5.6"
    PORT   = "57900"
    VRF    = "management"
  }
}