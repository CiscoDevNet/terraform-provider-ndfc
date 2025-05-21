# Terraform Provider NDFC

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.23

## Installing The Provider

1. Clone the repository
2. Enter the repository directory
3. Build the provider using the Go Install command:

```shell
go install
```

The provider should be available in `$GOPATH/bin`

## Using the provider

For using the locally compiled provider:
Refer <https://developer.hashicorp.com/terraform/cli/config/config-file>
Section: Development Overrides for Provider Developers
To add dev overrides to use a the plugin under development


## Provider plugin Documentation

[Provider](docs/index.md)
[Resources](docs/resources)
[Datasources](docs/data-sources)

## Sample Workflow

Following is a sample integrated config that creates a new fabric and adds a switch into it.
Save this to a `config.tf` file, modify the paramters according to the NDFC in use

```
terraform {
  required_providers {
    ndfc = {
      source = "registry.terraform.io/cisco/ndfc"
    }
  }
}

provider "ndfc" {
  username = "admin"
  password = "test"
  url     = "https://my-ndfc"
  insecure = true
}

resource "ndfc_fabric_vxlan_evpn" "my_fabric_1" {
  fabric_name                                 = "my_fabric_name"
  bgp_as                                      = "65000"
  deploy                                      = false
}
resource "ndfc_inventory_devices" "test_resource_inventory_devices_1" {
  fabric_name                               = "my_fabric_name"
  auth_protocol                             = "md5"
  username                                  = "admin"
  password                                  = "admin_password"
  max_hops                                  = 0
  set_as_individual_device_write_credential = false
  preserve_config                           = false
  save                                      = true
  deploy                                    = false
  retries                                   = 300
  retry_wait_timeout                        = 20
  devices = {
    "10.1.1.1" = {
      role                    = "spine"
      discovery_type          = "discover"
      discovery_auth_protocol = "md5"
    }
  }
}

# resource to do a final recalculate and deploy
resource "ndfc_configuration_deploy" "test_resource_configuration_deploy_1" {
  fabric_name              = "my_fabric_name"
  serial_numbers           = ["ALL"]
  config_save              = true
  trigger_deploy_on_update = false
  # Adding depends on ensures that features are configured in the right order in NDFC
  depends_on = [
    ndfc_inventory_devices.test_resource_inventory_devices_1, 
    ndfc_fabric_vxlan_evpn.my_fabric_1
  ]
}
```

### Execute following to configure NDFC

```shell
terraform plan
terraform apply
```

### Execute following to cleanup

```shell
terraform destroy
```
