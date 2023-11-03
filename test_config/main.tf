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
  host      = "https://10.104.251.69"
}

data "ndfc_vrfs" "vrf_example" {
  fabric_name = "test_evpn_vxlan"
}

output "vrf_example" {
  value = data.ndfc_vrfs.vrf_example
}

