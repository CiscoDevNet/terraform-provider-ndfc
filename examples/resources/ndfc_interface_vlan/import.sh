
# Format of ID used for import:
# if_policy:serial_number[comma seperated list of interfaces]
# if_policy:serial_number
terraform import ndfc_interface_vlan.test_resource_interface_vlan int_vlan:FDO245206N5[Vlan1000,Vlan1001],9990IQNFEZ6[Vlan1000,Vlan1001]
terraform import ndfc_interface_vlan.test_resource_interface_vlan int_vlan:FDO245206N5

