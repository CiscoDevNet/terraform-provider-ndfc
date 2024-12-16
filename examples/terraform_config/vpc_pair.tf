terraform {
  required_providers {
    ndfc = {
      source = "registry.terraform.io/cisco/ndfc"
    }
  }
}

provider "ndfc" {
  username = "admin"
  password = "admin!@#"
  host     = "https://"
  insecure = true
}

resource "ndfc_vpc_pair" "test_vpc_pair" {
  serial_numbers       = ["9HGCZABXAUY", "9N14Y8PVD2Y"]
  use_virtual_peerlink = false
}
