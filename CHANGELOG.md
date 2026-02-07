## 0.3.0
_Pre Release_
### BREAKING CHANGES
- **`ndfc_configuration_deploy` resource**: Renamed attributes for clarity
  - `config_save` renamed to `recalculate`
  - `trigger_deploy_on_update` renamed to `always_execute`
### Added
- **`ndfc_configuration_deploy` resource**: Added `deploy` flag to control deployment separately from recalculation
### Fixed
- Bug fixes and stability improvements

## 0.2.1
_Pre Release_
### Added
- Bug Fixes

## 0.2.0

_Pre Release_
### Added
- New Resources
    - `ndfc_links`
    - `ndfc_rest_api`
- New Data Sources
    - `ndfc_rest_api`
- Policy group support in `ndfc_policy` resource
- Bug Fixes

## 0.1.0 

_Pre Release_
### Added
- Fabric resources and data source
    - `ndfc_fabric_vxlan_evpn`
    - `ndfc_fabric_vxlan_msd`
    - `ndfc_fabric_lan_classic`
    - `ndfc_fabric_msite_ext_net`
    - `ndfc_fabric_ipfm`
- Inventory Management `ndfc_inventory_devices` resource and data source
- VPC Pair `ndfc_vpc_pair` resource
- Interface resources and data source
    - `ndfc_interface_ethernet`
    - `ndfc_interface_portchannel`
    - `ndfc_interface_loopback`
    - `ndfc_interface_vlan` 
    - `ndfc_interface_vpc`
- Policy resource `ndfc_policy`
- Template resource `ndfc_template`
- VRF management `ndfc_vrfs`
- Network management `ndfc_networks`




