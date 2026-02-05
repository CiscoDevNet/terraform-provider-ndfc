
resource "ndfc_configuration_deploy" "test_resource_configuration_deploy_1" {
  fabric_name    = "CML"
  serial_numbers = ["ALL"]
  recalculate    = true
  always_execute = false
  deploy         = true
}