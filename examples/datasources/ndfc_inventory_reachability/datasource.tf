
data "ndfc_inventory_reachability" "test_resource_inventory_reachability_1" {

  fabric_name = "CML"

  auth_protocol = "sha"

  username = "admin"

  password = "admin_password"

  seed_ip = "10.0.0.1"

  max_hops = "10"

  set_as_individual_device_write_credential = "true"

  preserve_config = "true"


}
