---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ndfc_interface_loopback Resource - terraform-provider-ndfc"
subcategory: "Interface"
description: |-
  This resource can manage a Interface Loopback.
---

# ndfc_interface_loopback (Resource)

This resource can manage a Interface Loopback.

## Example Usage

```terraform
resource "ndfc_interface_loopback" "example" {
  serial_number         = "9DBYO6WQJ46"
  interface_name        = "loopback123"
  policy                = "int_loopback"
  vrf                   = "VRF1"
  ipv4_address          = "5.6.7.8"
  ipv6_address          = "2001::10"
  route_map_tag         = "12346"
  interface_description = "My interface description"
  freeform_config       = "logging event port link-status"
  admin_state           = false
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `admin_state` (Boolean) Enable or disable the interface
  - Default value: `true`
- `freeform_config` (String) Additional CLI for the interface
- `interface_description` (String) Interface description
- `interface_name` (String) Name of the Interface. Example: `loopback123`
- `ipv4_address` (String) For VxLAN fabrics, configure an IPv4 address if underlay is V4 and VRF is default, otherwise add the config to freeform if underlay is V6.  For non-VxLAN fabrics or non-default VRF, loopback interfaces can have both IPv4 and IPv6 addresses.
- `ipv6_address` (String) For VxLAN fabrics, configure an IPv6 address if underlay is V6 and VRF is default, otherwise add the config to freeform if underlay is V4.  For non-VxLAN fabrics or non-default VRF, loopback interfaces can have both IPv4 and IPv6 addresses.
- `policy` (String) Name of the policy. Examples: `int_loopback`, `int_multisite_loopback`, `int_freeform`
  - Default value: `int_loopback`
- `route_map_tag` (String) Route-Map tag associated with interface IP
  - Default value: `12345`
- `serial_number` (String) Serial number of switch to configure
- `vrf` (String) Interface VRF name, default VRF if not specified

### Read-Only

- `id` (String) The id of the object

## Import

Import is supported using the following syntax:

```shell
terraform import ndfc_interface_loopback.example "9DBYO6WQJ46:loopback123"
```
