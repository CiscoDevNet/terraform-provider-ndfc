# NDFC MSD Workflow Module

Terraform module to provision a complete NDFC MSD (Multi-Site Domain) fabric workflow:
MSD parent fabric, N child VXLAN EVPN fabrics, VRFs, and networks.

## What This Module Creates

1. **Child fabrics** — `ndfc_fabric_vxlan_evpn` (variable number)
2. **MSD parent fabric** — `ndfc_fabric_vxlan_msd` referencing child fabrics
3. **VRFs at MSD parent** — `ndfc_vrfs` with attachments carrying `fabric` field
4. **VRFs at each child** — `ndfc_vrfs` per child with child-specific attributes
5. **Networks at MSD parent** — `ndfc_networks` with attachments carrying `fabric` field
6. **Networks at each child** — `ndfc_networks` per child with child-specific attributes
7. **Deploy** (optional) — `ndfc_configuration_deploy`

## Usage

```hcl
module "msd_dc" {
  source = "./examples/modules/msd"

  msd_fabric_name = "MSD_PROD"
  msd_fabric_config = {
    deploy             = false
    anycast_gw_mac     = "2020.0000.00aa"
    dci_subnet_range   = "10.10.1.0/24"
    default_network    = "Default_Network_Universal"
    default_vrf        = "Default_VRF_Universal"
  }

  child_fabrics = {
    "CHILD_DC1" = { bgp_as = "65001", deploy = false, ... }
    "CHILD_DC2" = { bgp_as = "65002", deploy = false, ... }
  }

  switch_fabric_map = {
    "SERIAL1" = "CHILD_DC1"
    "SERIAL2" = "CHILD_DC2"
  }

  vrf_configs     = { ... }
  network_configs = { ... }
  deploy          = false
}
```

See [examples/basic/main.tf](examples/basic/main.tf) for a complete working example.

## Requirements

| Name      | Version  |
|-----------|----------|
| terraform | >= 1.5.0 |
| ndfc      | >= 0.1.0 |

## Inputs

| Name | Description | Type | Required | Default |
|------|-------------|------|----------|---------|
| `msd_fabric_name` | Name of the MSD parent fabric | `string` | yes | — |
| `msd_fabric_config` | MSD parent fabric attributes (all `ndfc_fabric_vxlan_msd` optional attrs) | `any` | yes | — |
| `child_fabrics` | Map of child fabric name → `ndfc_fabric_vxlan_evpn` config | `map(any)` | yes | — |
| `switch_fabric_map` | Flat map of switch serial number → child fabric name | `map(string)` | yes | — |
| `vrf_configs` | VRF configurations (applied to parent + all children) | `map(object({...}))` | no | `{}` |
| `network_configs` | Network configurations (applied to parent + all children) | `map(object({...}))` | no | `{}` |
| `deploy` | Whether to trigger `ndfc_configuration_deploy` after creation | `bool` | no | `false` |

## Outputs

| Name | Description |
|------|-------------|
| `msd_fabric_name` | MSD parent fabric name |
| `msd_fabric_id` | Terraform ID of the MSD parent fabric |
| `child_fabric_names` | List of child fabric names |
| `child_fabric_ids` | Map of child fabric name → Terraform ID |
| `switches_by_child` | Switches grouped by child fabric name |

## Design Notes

### Variable number of children

Child fabrics use `for_each` over the input map. VRF and network child resources also use `for_each` over the same key set.

### Switch-to-child mapping

A flat `map(string)` of `serial → child_name`. The module computes `switches_by_child` to group them:
- **Parent attachments**: every switch, with `fabric` field
- **Child attachments**: only that child's switches, without `fabric` field

### Parent vs child attribute split

Based on NDFC MSD workflow documentation:

**VRF child-only**: `advertise_default_route`, `advertise_host_routes`

**Network child-only**: `dhcp_relay_loopback_id`, `dhcp_relay_servers`, `multicast_group`, `trm`, `netflow`, `svi_netflow_monitor`, `vlan_netflow_monitor`, `l3_gatway_border`, `igmp_version`

### Dependency chain

```
child fabrics → MSD parent → parent VRFs → child VRFs → parent networks → child networks → deploy
```

### Fabric config pass-through

Child fabric and MSD parent configs use `type = any` with `try(value, null)` to avoid duplicating the 230+ attribute schema. The provider handles defaults and validation.
