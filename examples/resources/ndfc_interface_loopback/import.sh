
# Format of ID used for import:
# if_policy:serial_number[comma seperated list of interfaces]
# if_policy:serial_number
terraform import ndfc_interface_loopback.test_resource_interface_loopback int_loopback:FDO245206N5[Loopback1,Loopback2],9990IQNFEZ6[Loopback0,Loopback1]
terraform import ndfc_interface_loopback.test_resource_interface_loopback int_loopback:FDO245206N5

