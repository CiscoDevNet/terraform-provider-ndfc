resource "ndfc_fabric_vxlan_msd" "msd_fabric" {
  depends_on    = [ndfc_fabric_vxlan.child_fabric1, ndfc_fabric_vxlan.child_fabric2]
  fabric_name   = "msd_fabric"
  deploy        = false
  child_fabrics = ["child_fabric1", "child_fabric2"]
}
