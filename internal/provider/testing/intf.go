package testing

import (
	"fmt"
	"terraform-provider-ndfc/internal/provider/resources/resource_interface_common"
	"terraform-provider-ndfc/internal/provider/types"
)

func GenerateIntfResource(intfObj **resource_interface_common.NDFCInterfaceCommonModel, ifStart, ifCount int,
	ifType string, deployNeeded bool, serials []string, globalSerial bool, append bool) {
	intf := *intfObj
	if !append {
		//intf := new(resource_interface_common.NDFCInterfaceCommonModel)
		intf.Interfaces = make(map[string]resource_interface_common.NDFCInterfacesValue)
	}
	ifPrefix := ""
	switch ifType {
	case "ethernet":
		intf.Policy = "int_trunk_host"
		ifPrefix = "Ethernet1/"
	case "loopback":
		intf.Policy = "int_loopback"
		ifPrefix = "loopback"
	case "vlan":
		intf.Policy = "int_vlan"
		ifPrefix = "vlan"
	}
	if globalSerial {
		intf.SerialNumber = serials[0]
	} else {
		intf.SerialNumber = ""
	}

	intf.Deploy = deployNeeded

	if ifCount < 0 {
		// Delete entries
		for i := 0; i < (ifCount * -1); i++ {
			intfNumber := ifStart + i
			intfName := fmt.Sprintf("%s%d", ifPrefix, intfNumber)
			delete(intf.Interfaces, intfName)
		}
		*intfObj = intf
		return
	}

	for i := 0; i < ifCount; i++ {
		ifTmp := new(resource_interface_common.NDFCInterfacesValue)
		key := ""
		intfNumber := ifStart + i
		intfName := fmt.Sprintf("%s%d", ifPrefix, intfNumber)
		if !globalSerial {
			ifTmp.SerialNumber = serials[i%len(serials)]
		}
		key = intfName
		ifTmp.InterfaceName = intfName
		ifTmp.NvPairs.AdminState = "true"
		ifTmp.NvPairs.FreeformConfig = ""
		ifTmp.NvPairs.InterfaceDescription = "Interface " + key
		if ifType == "ethernet" {
			ifTmp.NvPairs.Speed = "Auto"
			ifTmp.NvPairs.Mtu = "jumbo"
			ifTmp.NvPairs.Netflow = "false"
			ifTmp.NvPairs.BpduGuard = "true"
			ifTmp.NvPairs.AccessVlan = new(types.Int64Custom)
			*ifTmp.NvPairs.AccessVlan = types.Int64Custom(1500 + intfNumber)
			ifTmp.NvPairs.AllowedVlans = "10-2000"
		} else if ifType == "loopback" {
			ifTmp.NvPairs.Ipv4Address = fmt.Sprintf("192.168.%d.10", intfNumber%256)
			ifTmp.NvPairs.Vrf = "default"

		} else if ifType == "vlan" {
			ifTmp.InterfaceType = "vlan"
			ifTmp.NvPairs.Ipv4Address = fmt.Sprintf("192.100.%d.10", intfNumber%256)
			ifTmp.NvPairs.Vrf = "default"
			ifTmp.NvPairs.Ipv4PrefixLength = "24"
			ifTmp.NvPairs.RoutingTag = ""
		}
		intf.Interfaces[key] = *ifTmp

	}
	*intfObj = intf
}

func ModifyInterface(intfObj **resource_interface_common.NDFCInterfaceCommonModel,
	ifStart, ifCount int,
	ifType string, values map[string]interface{}) {

	intf := *intfObj
	ifPrefix := ""
	switch ifType {
	case "ethernet":
		ifPrefix = "Ethernet1/"
	case "loopback":
		ifPrefix = "loopback"
	case "vlan":
		ifPrefix = "vlan"
	}
	for i := 0; i < ifCount; i++ {
		intfNumber := ifStart + i
		intfName := fmt.Sprintf("%s%d", ifPrefix, intfNumber)

		if ifTmp, ok := intf.Interfaces[intfName]; ok {
			if val, ok := values["adminState"]; ok {
				ifTmp.NvPairs.AdminState = val.(string)
			}
			if val, ok := values["freeformConfig"]; ok {
				ifTmp.NvPairs.FreeformConfig = val.(string)
			}
			if val, ok := values["interfaceDescription"]; ok {
				ifTmp.NvPairs.InterfaceDescription = val.(string)
			}
			if val, ok := values["ipv4Address"]; ok {
				ifTmp.NvPairs.Ipv4Address = val.(string)
			}
			if val, ok := values["vrf"]; ok {
				ifTmp.NvPairs.Vrf = val.(string)
			}
			if val, ok := values["ipv4PrefixLength"]; ok {
				ifTmp.NvPairs.Ipv4PrefixLength = val.(string)
			}
			if val, ok := values["routingTag"]; ok {
				ifTmp.NvPairs.RoutingTag = val.(string)
			}
			if val, ok := values["speed"]; ok {
				ifTmp.NvPairs.Speed = val.(string)
			}
			if val, ok := values["mtu"]; ok {
				ifTmp.NvPairs.Mtu = val.(string)
			}
			if val, ok := values["netflow"]; ok {
				ifTmp.NvPairs.Netflow = val.(string)
			}
			if val, ok := values["bpduGuard"]; ok {
				ifTmp.NvPairs.BpduGuard = val.(string)
			}
			if val, ok := values["accessVlan"]; ok {
				if ifTmp.NvPairs.AccessVlan == nil {
					ifTmp.NvPairs.AccessVlan = new(types.Int64Custom)
				}
				*ifTmp.NvPairs.AccessVlan = val.(types.Int64Custom)
			}
			if val, ok := values["allowedVlans"]; ok {
				ifTmp.NvPairs.AllowedVlans = val.(string)
			}
			if val, ok := values["nativeVlan"]; ok {
				if ifTmp.NvPairs.NativeVlan == nil {
					ifTmp.NvPairs.NativeVlan = new(types.Int64Custom)
				}
				*ifTmp.NvPairs.NativeVlan = val.(types.Int64Custom)
			}
			if val, ok := values["hsrpGroup"]; ok {
				if ifTmp.NvPairs.HsrpGroup == nil {
					ifTmp.NvPairs.HsrpGroup = new(types.Int64Custom)
				}
				*ifTmp.NvPairs.HsrpGroup = val.(types.Int64Custom)
			}
			if val, ok := values["hsrpPriority"]; ok {
				if ifTmp.NvPairs.HsrpPriority == nil {
					ifTmp.NvPairs.HsrpPriority = new(types.Int64Custom)
				}
				*ifTmp.NvPairs.HsrpPriority = val.(types.Int64Custom)
			}
			if val, ok := values["preempt"]; ok {
				ifTmp.NvPairs.Preempt = val.(string)
			}
			if val, ok := values["mac"]; ok {
				ifTmp.NvPairs.Mac = val.(string)
			}
			if val, ok := values["dhcpServerAddr1"]; ok {
				ifTmp.NvPairs.DhcpServerAddr1 = val.(string)
			}
			if val, ok := values["dhcpServerAddr2"]; ok {
				ifTmp.NvPairs.DhcpServerAddr2 = val.(string)
			}
			if val, ok := values["dhcpServerAddr3"]; ok {
				ifTmp.NvPairs.DhcpServerAddr3 = val.(string)
			}
			if val, ok := values["vrfDhcp1"]; ok {
				ifTmp.NvPairs.VrfDhcp1 = val.(string)
			}
			if val, ok := values["vrfDhcp2"]; ok {
				ifTmp.NvPairs.VrfDhcp2 = val.(string)
			}
			if val, ok := values["vrfDhcp3"]; ok {
				ifTmp.NvPairs.VrfDhcp3 = val.(string)
			}
			if val, ok := values["advertiseSubnetInUnderlay"]; ok {
				ifTmp.NvPairs.AdvertiseSubnetInUnderlay = val.(string)
			}
			intf.Interfaces[intfName] = ifTmp

		}

	}

}
