resource "ndfc_configuration_deploy" "this" {
  count = var.deploy ? 1 : 0

  depends_on = [
    ndfc_networks.parent,
    ndfc_networks.child,
  ]

  fabric_name    = ndfc_fabric_vxlan_msd.parent.fabric_name
  serial_numbers = ["ALL"]
  recalculate    = true
  deploy         = true
}
