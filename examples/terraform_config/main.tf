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
#  password = "ins3965!"
#  host      = "https://10.195.225.193"
  host      = "https://10.78.210.161"
  insecure = true
}
