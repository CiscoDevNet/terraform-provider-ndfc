
resource "ndfc_vpc_pair" "test_resource_vpc_pair_1" {
  serial_numbers       = ["FGE20360RRZ", "FGE20360RRY"]
  use_virtual_peerlink = false
  deploy               = true
}