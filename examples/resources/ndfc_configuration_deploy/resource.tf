
resource "ndfc_configuration_deploy" "test_resource_configuration_deploy_1" {
  fabric_name    = "CML"
  serial_numbers = ["FGE20360RRZ", "FGE20360RRY"]
  config_save    = true
}