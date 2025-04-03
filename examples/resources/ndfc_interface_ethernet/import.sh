
# Format of ID used for import:
# if_policy:serial_number[comma seperated list of interfaces]
# if_policy:serial_number
terraform import ndfc_interface_ethernet.test_resource_interface_ethernet int_trunk_host:FDO245206N5[Ethernet1/1,Ethernet1/2],9990IQNFEZ6[Ethernet1/3,Ethernet1/2]
terraform import ndfc_interface_ethernet.test_resource_interface_ethernet int_access_host:FDO245206N5

