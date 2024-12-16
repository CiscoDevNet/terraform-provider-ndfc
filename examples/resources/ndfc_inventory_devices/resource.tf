
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
  devices = {
    "" = {
      role                    = "spine"
      discovery_type          = "pre_provision"
      discovery_username      = "username"
      discovery_password      = "password"
      discovery_auth_protocol = "sha"
      serial_number           = "FGE20360RRZ"
      model                   = "N9K-9000v"
      version                 = "9.2(1)"
      hostname                = "s1-leaf-101"
      image_policy            = "TODO_CHOOSE_EXAMPLE"
      gateway                 = "10.61.124.1/24"
      modules_model           = ["N9K-X9736C-EX", "N9K-X9732C-FX"]
      breakout                = "interface breakout module 1 port 1,11,19 map 10g-4x; interface breakout module 1 port 7 map 25g-4x"
      port_mode               = "hardware profile portmode 48x25G + 2x100G + 4x40G"
    }
  }

}