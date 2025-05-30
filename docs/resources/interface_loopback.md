---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ndfc_interface_loopback Resource - terraform-provider-ndfc"
subcategory: ""
description: |-
  Resource to configure loopback interfaces on a switch
---

# ndfc_interface_loopback (Resource)

Resource to configure loopback interfaces on a switch

## Example Usage

```terraform
resource "ndfc_interface_loopback" "test_resource_interface_loopback_1" {
  policy        = "int_loopback"
  deploy        = true
  serial_number = "9DBYO6WQJ46"
  interfaces = {
    "Loopback100" = {
      interface_name        = "loopback100"
      admin_state           = false
      interface_description = "This is a loopback interface used for XYZ"
      vrf                   = "default"
      ipv4_address          = "192.168.20.1"
      ipv6_address          = "2002:db8::1"
      route_map_tag         = "100"
    }
  }

}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `interfaces` (Attributes Map) interfaces to configure (see [below for nested schema](#nestedatt--interfaces))

### Optional

- `deploy` (Boolean) Deploy the configuration
- `policy` (String) Name of the policy. Supported policies:: `int_loopback`, `int_multisite_loopback`
- `serial_number` (String) Serial number of switch to configure. This field cannot be specified if `serial_number` inside `interfaces` block is specified`

### Read-Only

- `id` (String) Unique identifier for the interface

<a id="nestedatt--interfaces"></a>
### Nested Schema for `interfaces`

Required:

- `interface_name` (String) Name of the Interface. Example: `loopback1`

Optional:

- `admin_state` (Boolean) Administratively enable or disable the interface
- `freeform_config` (String) Additional CLI commands to be executed for the interface
- `interface_description` (String) Interface description
- `ipv4_address` (String) IPv4 address
- `ipv6_address` (String) IPv6 address
- `route_map_tag` (String) Route map tag
- `serial_number` (String) Serial number of switch to configure. This field cannot be specified if `serial_number` outside `interfaces` block is specified
- `vrf` (String) VRF name

Read-Only:

- `deployment_status` (String) Status of the deployment

## Import

Import is supported using the following syntax:

```shell
# Format of ID used for import:
# if_policy:serial_number[comma seperated list of interfaces]
# if_policy:serial_number
terraform import ndfc_interface_loopback.test_resource_interface_loopback int_loopback:FDO245206N5[Loopback1,Loopback2],9990IQNFEZ6[Loopback0,Loopback1]
terraform import ndfc_interface_loopback.test_resource_interface_loopback int_loopback:FDO245206N5
```
