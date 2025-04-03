
resource "ndfc_configuration_deploy" "test_resource_configuration_deploy_1" {
  fabric_name              = "CML"
  serial_numbers           = ["ALL"]
  config_save              = true
  trigger_deploy_on_update = false
}