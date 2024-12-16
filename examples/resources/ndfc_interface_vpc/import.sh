
# Format of ID used for import:
# if_policy:serial_number[comma seperated list of interfaces]
# if_policy:serial_number
terraform import ndfc_interface_vpc.test_resource_interface_vpc int_vpc_trunk_host:FDO245206N5[vPC1,vPC2],9990IQNFEZ6[vPC0,vPC1]
terraform import ndfc_interface_vpc.test_resource_interface_vpc int_vpc_trunk_host:FDO245206N5

