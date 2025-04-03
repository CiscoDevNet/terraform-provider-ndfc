
resource "ndfc_inventory_devices" "test_resource_inventory_devices_1" {
  fabric_name                               = "CML"
  auth_protocol                             = "md5"
  username                                  = "admin"
  password                                  = "admin_password"
  max_hops                                  = 0
  set_as_individual_device_write_credential = false
  preserve_config                           = false
  save                                      = true
  deploy                                    = true
  retries                                   = 300
  retry_wait_timeout                        = 20
  devices = {
    "10.1.1.1" = {
      role                    = "spine"
      discovery_type          = "discover"
      discovery_auth_protocol = "md5"
    }
  }

}