terraform {
  required_providers {
    ndfc = {
      source = "registry.terraform.io/cisco/ndfc"
    }
  }
}

provider "ndfc" {
  username = "test"
  password = "test"
  host     = "https://10.78.210.161"
  insecure = true
}

data "ndfc_networks" "test_evpn_vxlan_deployment_nw1" {
  fabric_name = "test_evpn_vxlan"
}


data "ndfc_vrf_bulk" "test_evpn_vxlan_deployment_vrf1" {
  fabric_name = "test_evpn_vxlan"
}

