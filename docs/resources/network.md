---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ndfc_network Resource - terraform-provider-ndfc"
subcategory: "Fabric"
description: |-
  This resource can manage a Network.
---

# ndfc_network (Resource)

This resource can manage a Network.

## Example Usage

```terraform
resource "ndfc_network" "example" {
  fabric_name                = "CML"
  network_name               = "NET1"
  display_name               = "NET1"
  network_template           = "Default_Network_Universal"
  network_extension_template = "Default_Network_Extension_Universal"
  vrf_name                   = "VRF1"
  gateway_ipv4_address       = "192.0.2.1/24"
  vlan_id                    = 1500
  gateway_ipv6_address       = "2001:db8::1/64,2001:db9::1/64"
  layer2_only                = false
  arp_suppression            = false
  ingress_replication        = false
  multicast_group            = "233.1.1.1"
  dhcp_relay_servers = [
    {
      address = "2.3.4.5"
      vrf     = "VRF1"
    }
  ]
  dhcp_relay_loopback_id = 134
  vlan_name              = "VLANXXX"
  interface_description  = "My int description"
  mtu                    = 9200
  loopback_routing_tag   = 11111
  trm                    = true
  secondary_gateway_1    = "192.168.2.1/24"
  secondary_gateway_2    = "192.168.3.1/24"
  secondary_gateway_3    = "192.168.4.1/24"
  secondary_gateway_4    = "192.168.5.1/24"
  route_target_both      = true
  netflow                = false
  svi_netflow_monitor    = "MON1"
  vlan_netflow_monitor   = "MON1"
  l3_gatway_border       = true
  attachments = [
    {
      serial_number       = "9DBYO6WQJ46"
      attach_switch_ports = "Ethernet1/10,Ethernet1/11"
      vlan_id             = 2010
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_name` (String) The name of the network
- `vrf_name` (String) The name of the vrf

### Optional

- `arp_suppression` (Boolean) ARP suppression is only supported if SVI is present when Layer-2-Only is not enabled. NX-OS Specific
- `attachments` (Attributes Set) A list of attachments (see [below for nested schema](#nestedatt--attachments))
- `dhcp_relay_loopback_id` (Number) Loopback ID for DHCP Relay interface
  - Range: `0`-`1023`
- `dhcp_relay_servers` (Attributes List) List of DHCP relay servers (see [below for nested schema](#nestedatt--dhcp_relay_servers))
- `display_name` (String) Customized name of the network. By default, it will be same as the network name
- `fabric_name` (String) The name of the fabric
- `gateway_ipv4_address` (String) Gateway IPv4 address, for example `192.0.2.1/24`
- `gateway_ipv6_address` (String) Gateway IPv6 addresses, for example `2001:db8::1/64,2001:db9::1/64`
- `ingress_replication` (Boolean) Ingress replication flag
  - Default value: `false`
- `interface_description` (String) Interface description
- `l3_gatway_border` (Boolean) Enable L3 Gateway on Border
  - Default value: `false`
- `layer2_only` (Boolean) Layer-2 only flag
  - Default value: `false`
- `loopback_routing_tag` (Number) Loopback routing tag
  - Range: `0`-`4294967295`
  - Default value: `12345`
- `mtu` (Number) Interface MTU
  - Range: `68`-`9216`
  - Default value: `9216`
- `multicast_group` (String) Multicast group address
- `netflow` (Boolean) Netflow is supported only if it is enabled on fabric. For NX-OS only
  - Default value: `false`
- `network_extension_template` (String) The name of the network extension template. Applicable to Switch(es) with role Border
  - Default value: `Default_Network_Extension_Universal`
- `network_id` (Number) VNI ID of the network
  - Range: `1`-`16777214`
- `network_template` (String) The name of the network template
  - Default value: `Default_Network_Universal`
- `route_target_both` (Boolean) L2 VNI Route-Target Both Enable
  - Default value: `false`
- `secondary_gateway_1` (String) Secondary gateway 1
- `secondary_gateway_2` (String) Secondary gateway 2
- `secondary_gateway_3` (String) Secondary gateway 3
- `secondary_gateway_4` (String) Secondary gateway 4
- `svi_netflow_monitor` (String) Applicable only if 'Layer 2 Only' is not enabled. Provide monitor name defined in fabric setting for Layer 3 Record. For NX-OS only
- `trm` (Boolean) Enable Tenant Routed Multicast
- `vlan_id` (Number) VLAN ID
  - Range: `2`-`4094`
- `vlan_name` (String) VLAN name
- `vlan_netflow_monitor` (String) Provide monitor name defined in fabric setting for Layer 3 Record. For NX-OS only

### Read-Only

- `id` (String) The id of the object

<a id="nestedatt--attachments"></a>
### Nested Schema for `attachments`

Required:

- `serial_number` (String) Serial number of switch to attach

Optional:

- `attach_switch_ports` (String) Comma separated list of attached switchports
- `detach_switch_ports` (String) Comma separated list of detached switchports
- `freeform_config` (String) This field covers any configuration not included in overlay templates which is needed as part of this VRF attachment
- `vlan_id` (Number) Override VLAN ID. `-1` to use VLAN ID defined at VRF level
  - Range: `-1`-`4092`
  - Default value: `-1`


<a id="nestedatt--dhcp_relay_servers"></a>
### Nested Schema for `dhcp_relay_servers`

Optional:

- `address` (String) Server IP V4 Address
- `vrf` (String) If management vrf, enter 'management'. If default/global vrf, enter 'default'.

## Import

Import is supported using the following syntax:

```shell
terraform import ndfc_network.example "CML:NET1"
```
