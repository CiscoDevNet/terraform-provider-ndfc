# Serial is mandatory
# Optional parameters
# interface_types  - filter by interface types
#  * Port Channel -- INTERFACE_PORT_CHANNEL
#  * virtual Port Channel (vPC) --- INTERFACE_VPC
#  * Loopback --- INTERFACE_LOOPBACK
#  * Straight-through (ST) FEX --- STRAIGHT_TROUGH_FEX
#  * Active-Active (AA) FEX --- AA_FEX
#  * Subinterface --- SUBINTERFACE
#  * Ethernet --- INTERFACE_ETHERNET
#  * Switch Virtual Interface (SVI) --- INTERFACE_VLAN
# port_modes - filter by modes
# excludes - filter out using exclude parameter - see NDFC documentation

                     
data "ndfc_interfaces" "test_evpn_vxlan_deployment_leaf01" {
  serial_number = "9FE076D8EJL"
  interface_types = "INTERFACE_LOOPBACK"
}

data "ndfc_interfaces" "test_evpn_vxlan_deployment_leaf02" {
  serial_number = "9TQYTJSZ1VJ"
  interface_types = "INTERFACE_LOOPBACK"
}

data "ndfc_interfaces" "test_evpn_vxlan_deployment_spine01" {
    serial_number = "9QBCTIN0FMY"
    interface_types = "INTERFACE_LOOPBACK"
  
}

data "ndfc_interfaces" "test_evpn_vxlan_deployment_leaf01_eth" {
  serial_number = "9FE076D8EJL"
  interface_types = "INTERFACE_ETHERNET"
}

data "ndfc_interfaces" "test_evpn_vxlan_deployment_leaf02_eth" {
  serial_number = "9TQYTJSZ1VJ"
  interface_types = "INTERFACE_ETHERNET"
}

data "ndfc_interfaces" "test_evpn_vxlan_deployment_spine01_eth" {
    serial_number = "9QBCTIN0FMY"
    interface_types = "INTERFACE_ETHERNET"
  
}