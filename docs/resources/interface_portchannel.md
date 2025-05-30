---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ndfc_interface_portchannel Resource - terraform-provider-ndfc"
subcategory: ""
description: |-
  Resource to configure port-channel interfaces on a switch
---

# ndfc_interface_portchannel (Resource)

Resource to configure port-channel interfaces on a switch

## Example Usage

```terraform
resource "ndfc_interface_portchannel" "test_resource_interface_portchannel_1" {
  policy        = "int_port_channel_trunk_host"
  deploy        = true
  serial_number = "9DBYO6WQJ46"
  interfaces = {
    "Port-channel100" = {
      interface_name        = "port-channel100"
      admin_state           = false
      interface_description = "My interface description"
      bpdu_guard            = "true"
      port_type_fast        = false
      mtu                   = "default"
      speed                 = "Auto"
      orphan_port           = false
      netflow               = false
      netflow_monitor       = "MON1"
      netflow_sampler       = "SAMPLER1"
      allowed_vlans         = "10-20"
      native_vlan           = 1
      copy_po_description   = false
      portchannel_mode      = "on"
      member_interfaces     = "eth1/5-6"
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
- `policy` (String) "Name of the policy. 
  Supported policies: 
    * `int_port_channel_trunk_host`
    * `int_port_channel_access_host`
    * `int_port_channel_dot1q_tunnel_host`
    * `int_port_channel_pvlan_host`
    * `int_l3_port_channel`
    * `int_monitor_port_channel`"
- `serial_number` (String) Serial number of switch to configure. This field cannot be specified if `serial_number` inside `interfaces` block is specified`

### Read-Only

- `id` (String) Unique identifier for the interface

<a id="nestedatt--interfaces"></a>
### Nested Schema for `interfaces`

Required:

- `interface_name` (String) Name of the Interface. Example: `port-channel1`

Optional:

- `admin_state` (Boolean) Enable or disable the interface
- `allowed_vlans` (String) Allowed vlans for the ethernet interface. Allowed values are `none`, `all` or VLAN ranges (1-200,500-2000,3000)
- `bpdu_guard` (String) Enable spanning-tree bpduguard: true='enable', false='disable', no='return to default settings'
- `copy_po_description` (Boolean) Netflow is supported only if it is enabled on fabric
- `freeform_config` (String) Additional CLI for the interface
- `interface_description` (String) Interface description
- `member_interfaces` (String) Member interfaces of the port channel. Allowed formats are "eth1/1-10" or "eth1/1,eth1/2,eth1/3"
- `mtu` (String) MTU for the interface
- `native_vlan` (Number) Set native VLAN for the interface
- `netflow` (Boolean) Netflow is supported only if it is enabled on fabric
- `netflow_monitor` (String) Provide the Layer 2 Monitor Name
- `netflow_sampler` (String) Netflow sampler name, applicable to N7K only
- `orphan_port` (Boolean) If enabled, configure the interface as a vPC orphan port to be suspended by the secondary peer in vPC failures
- `port_type_fast` (Boolean) Enable spanning-tree edge port behavior
- `portchannel_mode` (String) Port-channel mode. Allowed values are `on`, `active`, `passive`
- `serial_number` (String) Serial number of switch to configure. This field cannot be specified if `serial_number` is already mentioned outside
- `speed` (String) Interface speed

Read-Only:

- `deployment_status` (String) Status of the deployment

## Import

Import is supported using the following syntax:

```shell
# Format of ID used for import:
# if_policy:serial_number[comma seperated list of interfaces]
# if_policy:serial_number
terraform import ndfc_interface_portchannel.test_resource_interface_portchannel int_port_channel_trunk_host:FDO245206N5[Port-channel1,Port-channel2],9990IQNFEZ6[Port-channel0,Port-channel1]
terraform import ndfc_interface_portchannel.test_resource_interface_portchannel int_port_channel_trunk_host:FDO245206N5
```
