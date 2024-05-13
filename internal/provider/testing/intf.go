package testing

import (
	"fmt"
	"math/rand"
	"terraform-provider-ndfc/internal/provider/resources/resource_interface_common"
)

func GenerateIntfResource(intfObj **resource_interface_common.NDFCInterfaceCommonModel, ifStart, ifCount int,
	ifType string, deployNeeded bool, serials []string, globalSerial bool) {

	intf := new(resource_interface_common.NDFCInterfaceCommonModel)
	intf.Interfaces = make(map[string]resource_interface_common.NDFCInterfacesValue)

	ifPrefix := ""
	switch ifType {
	case "ethernet":
		intf.Policy = "int_trunk_host"
		ifPrefix = "Ethernet1/"
	case "loopback":
		intf.Policy = "int_loopback"
		ifPrefix = "loopback"
	}
	if globalSerial {
		intf.SerialNumber = serials[0]
	} else {
		intf.SerialNumber = ""
	}

	intf.Deploy = deployNeeded

	for i := 0; i < ifCount; i++ {
		ifTmp := new(resource_interface_common.NDFCInterfacesValue)
		key := ""
		intfName := fmt.Sprintf("%s%d", ifPrefix, ifStart+i+1)

		if !globalSerial {
			ifTmp.SerialNumber = serials[i%len(serials)]
			key = ifTmp.SerialNumber + ":" + intfName
		} else {
			key = intf.SerialNumber + ":" + intfName
		}
		ifTmp.InterfaceName = intfName
		ifTmp.NvPairs.AdminState = "true"
		ifTmp.NvPairs.FreeformConfig = ""
		ifTmp.NvPairs.InterfaceDescription = "Interface " + key
		if ifType == "ethernet" {
			ifTmp.NvPairs.Speed = "Auto"
			ifTmp.NvPairs.Mtu = "jumbo"
			ifTmp.NvPairs.Netflow = "false"
			ifTmp.NvPairs.BpduGuard = "true"
			ifTmp.NvPairs.AccessVlan = new(int64)
			*ifTmp.NvPairs.AccessVlan = int64(rand.Intn(4090) + 1)
			ifTmp.NvPairs.AllowedVlans = "10-2000"
		} else if ifType == "loopback" {
			ifTmp.NvPairs.Ipv4Address = fmt.Sprintf("192.168.%d.10", i%256)
			ifTmp.NvPairs.Vrf = "default"

		}
		intf.Interfaces[key] = *ifTmp

	}
	*intfObj = intf
}
