
# Format of ID used for import:
# if_policy:serial_number[comma seperated list of interfaces]
# if_policy:serial_number
terraform import ndfc_interface_portchannel.test_resource_interface_portchannel int_port_channel_trunk_host:FDO245206N5[Port-channel1,Port-channel2],9990IQNFEZ6[Port-channel0,Port-channel1]
terraform import ndfc_interface_portchannel.test_resource_interface_portchannel int_port_channel_trunk_host:FDO245206N5

