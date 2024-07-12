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

resource "ndfc_policy" "test_policy" {
  entity_name          = "Switch"
  entity_type          = "SWITCH"
  description          = "Test Policy from TF"
  template_name        = "TelemetryDst_EF"
  device_serial_number = "9FE076D8EJL"
  deploy = true
  priority = 500
  policy_parameters = {
    DSTGRP = "501"
    IPADDR = "5.5.5.6"
    PORT   = "57900"
    VRF    = "management"
  }
}