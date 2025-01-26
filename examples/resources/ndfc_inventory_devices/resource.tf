
resource "ndfc_inventory_devices" "test_resource_inventory_devices_1" {
  fabric_name                               = "CML"
  auth_protocol                             = "sha"
  username                                  = "admin"
  password                                  = "admin_password"
  seed_ip                                   = "10.0.0.1"
  max_hops                                  = 10
  set_as_individual_device_write_credential = true
  preserve_config                           = true
  save                                      = true
  deploy                                    = true
  retries                                   = 500
  retry_wait_timeout                        = 20
}