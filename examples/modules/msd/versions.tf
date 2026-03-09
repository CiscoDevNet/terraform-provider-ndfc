terraform {
  required_version = ">= 1.5.0"

  required_providers {
    ndfc = {
      source  = "CiscoDevNet/ndfc"
      version = ">= 0.1.0"
    }
  }
}
