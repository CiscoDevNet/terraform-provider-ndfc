package resource_interface_common

import (
	"log"
	. "terraform-provider-ndfc/internal/provider/types"
)

func (v NDFCInterfacesValue) DeepEqual(c NDFCInterfacesValue) int {
	cf := false
	if v.SerialNumber != c.SerialNumber {
		log.Printf("v.SerialNumber=%v, c.SerialNumber=%v", v.SerialNumber, c.SerialNumber)
		return RequiresUpdate
	}
	if v.InterfaceName != c.InterfaceName {
		log.Printf("v.InterfaceName=%v, c.InterfaceName=%v", v.InterfaceName, c.InterfaceName)
		return RequiresReplace
	}
	if v.NvPairs.AdminState != c.NvPairs.AdminState {
		log.Printf("v.NvPairs.AdminState=%s, c.NvPairs.AdminState=%s", v.NvPairs.AdminState, c.NvPairs.AdminState)
		return RequiresUpdate
	}
	if v.NvPairs.FreeformConfig != c.NvPairs.FreeformConfig {
		log.Printf("v.NvPairs.FreeformConfig=%s, c.NvPairs.FreeformConfig=%s", v.NvPairs.FreeformConfig, c.NvPairs.FreeformConfig)
		return RequiresUpdate
	}
	if v.NvPairs.InterfaceDescription != c.NvPairs.InterfaceDescription {
		log.Printf("v.NvPairs.InterfaceDescription=%s, c.NvPairs.InterfaceDescription=%s", v.NvPairs.InterfaceDescription, c.NvPairs.InterfaceDescription)
		return RequiresUpdate
	}
	if v.NvPairs.Vrf != c.NvPairs.Vrf {
		log.Printf("v.NvPairs.Vrf=%s, c.NvPairs.Vrf=%s", v.NvPairs.Vrf, c.NvPairs.Vrf)
		return RequiresUpdate
	}
	if v.NvPairs.Ipv4Address != c.NvPairs.Ipv4Address {
		log.Printf("v.NvPairs.Ipv4Address=%s, c.NvPairs.Ipv4Address=%s", v.NvPairs.Ipv4Address, c.NvPairs.Ipv4Address)
		return RequiresUpdate
	}
	if v.NvPairs.Ipv6Address != c.NvPairs.Ipv6Address {
		log.Printf("v.NvPairs.Ipv6Address=%s, c.NvPairs.Ipv6Address=%s", v.NvPairs.Ipv6Address, c.NvPairs.Ipv6Address)
		return RequiresUpdate
	}
	if v.NvPairs.RouteMapTag != c.NvPairs.RouteMapTag {
		log.Printf("v.NvPairs.RouteMapTag=%s, c.NvPairs.RouteMapTag=%s", v.NvPairs.RouteMapTag, c.NvPairs.RouteMapTag)
		return RequiresUpdate
	}
	if v.NvPairs.BpduGuard != c.NvPairs.BpduGuard {
		log.Printf("v.NvPairs.BpduGuard=%s, c.NvPairs.BpduGuard=%s", v.NvPairs.BpduGuard, c.NvPairs.BpduGuard)
		return RequiresUpdate
	}
	if v.NvPairs.PortTypeFast != c.NvPairs.PortTypeFast {
		log.Printf("v.NvPairs.PortTypeFast=%s, c.NvPairs.PortTypeFast=%s", v.NvPairs.PortTypeFast, c.NvPairs.PortTypeFast)
		return RequiresUpdate
	}
	if v.NvPairs.Mtu != c.NvPairs.Mtu {
		log.Printf("v.NvPairs.Mtu=%s, c.NvPairs.Mtu=%s", v.NvPairs.Mtu, c.NvPairs.Mtu)
		return RequiresUpdate
	}
	if v.NvPairs.Speed != c.NvPairs.Speed {
		log.Printf("v.NvPairs.Speed=%s, c.NvPairs.Speed=%s", v.NvPairs.Speed, c.NvPairs.Speed)
		return RequiresUpdate
	}

	if v.NvPairs.AccessVlan != nil && c.NvPairs.AccessVlan != nil {
		if *v.NvPairs.AccessVlan != *c.NvPairs.AccessVlan {
			log.Printf("v.NvPairs.AccessVlan=%v, c.NvPairs.AccessVlan=%v", *v.NvPairs.AccessVlan, *c.NvPairs.AccessVlan)
			return RequiresUpdate
		}
	} else {
		if v.NvPairs.AccessVlan != nil {
			log.Printf("v.NvPairs.AccessVlan=%v", *v.NvPairs.AccessVlan)
			return RequiresUpdate
		} else if c.NvPairs.AccessVlan != nil {
			log.Printf("c.NvPairs.AccessVlan=%v", *c.NvPairs.AccessVlan)
			return RequiresUpdate
		}
	}

	if v.NvPairs.OrphanPort != c.NvPairs.OrphanPort {
		log.Printf("v.NvPairs.OrphanPort=%s, c.NvPairs.OrphanPort=%s", v.NvPairs.OrphanPort, c.NvPairs.OrphanPort)
		return RequiresUpdate
	}
	if v.NvPairs.Ptp != c.NvPairs.Ptp {
		log.Printf("v.NvPairs.Ptp=%s, c.NvPairs.Ptp=%s", v.NvPairs.Ptp, c.NvPairs.Ptp)
		return RequiresUpdate
	}
	if v.NvPairs.Netflow != c.NvPairs.Netflow {
		log.Printf("v.NvPairs.Netflow=%s, c.NvPairs.Netflow=%s", v.NvPairs.Netflow, c.NvPairs.Netflow)
		return RequiresUpdate
	}
	if v.NvPairs.NetflowMonitor != c.NvPairs.NetflowMonitor {
		log.Printf("v.NvPairs.NetflowMonitor=%s, c.NvPairs.NetflowMonitor=%s", v.NvPairs.NetflowMonitor, c.NvPairs.NetflowMonitor)
		return RequiresUpdate
	}
	if v.NvPairs.NetflowSampler != c.NvPairs.NetflowSampler {
		log.Printf("v.NvPairs.NetflowSampler=%s, c.NvPairs.NetflowSampler=%s", v.NvPairs.NetflowSampler, c.NvPairs.NetflowSampler)
		return RequiresUpdate
	}
	if v.NvPairs.AllowedVlans != c.NvPairs.AllowedVlans {
		log.Printf("v.NvPairs.AllowedVlans=%s, c.NvPairs.AllowedVlans=%s", v.NvPairs.AllowedVlans, c.NvPairs.AllowedVlans)
		return RequiresUpdate
	}

	if v.NvPairs.NativeVlan != nil && c.NvPairs.NativeVlan != nil {
		if *v.NvPairs.NativeVlan != *c.NvPairs.NativeVlan {
			log.Printf("v.NvPairs.NativeVlan=%v, c.NvPairs.NativeVlan=%v", *v.NvPairs.NativeVlan, *c.NvPairs.NativeVlan)
			return RequiresUpdate
		}
	} else {
		if v.NvPairs.NativeVlan != nil {
			log.Printf("v.NvPairs.NativeVlan=%v", *v.NvPairs.NativeVlan)
			return RequiresUpdate
		} else if c.NvPairs.NativeVlan != nil {
			log.Printf("c.NvPairs.NativeVlan=%v", *c.NvPairs.NativeVlan)
			return RequiresUpdate
		}
	}

	if cf {
		return ControlFlagUpdate
	}
	return ValuesDeeplyEqual
}

func (v *NDFCInterfacesValue) CreatePlan(c NDFCInterfacesValue, cf *bool) int {
	action := ActionNone

	if v.SerialNumber != "" {

		if v.SerialNumber != c.SerialNumber {
			log.Printf("Update: v.SerialNumber=%v, c.SerialNumber=%v", v.SerialNumber, c.SerialNumber)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}

	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.SerialNumber=%v, c.SerialNumber=%v", v.SerialNumber, c.SerialNumber)
		v.SerialNumber = c.SerialNumber
	}

	if v.InterfaceName != "" {

		if v.InterfaceName != c.InterfaceName {
			log.Printf("Update: v.InterfaceName=%v, c.InterfaceName=%v", v.InterfaceName, c.InterfaceName)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresReplace
			}
		}

	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.InterfaceName=%v, c.InterfaceName=%v", v.InterfaceName, c.InterfaceName)
		v.InterfaceName = c.InterfaceName
	}

	if v.NvPairs.AdminState != c.NvPairs.AdminState {
		log.Printf("Update: v.NvPairs.AdminState=%v, c.NvPairs.AdminState=%v", v.NvPairs.AdminState, c.NvPairs.AdminState)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.FreeformConfig != "" {
		if v.NvPairs.FreeformConfig != c.NvPairs.FreeformConfig {
			log.Printf("Update: v.NvPairs.FreeformConfig=%v, c.NvPairs.FreeformConfig=%v", v.NvPairs.FreeformConfig, c.NvPairs.FreeformConfig)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.FreeformConfig=%v, c.NvPairs.FreeformConfig=%v", v.NvPairs.FreeformConfig, c.NvPairs.FreeformConfig)
		v.NvPairs.FreeformConfig = c.NvPairs.FreeformConfig
	}

	if v.NvPairs.InterfaceDescription != "" {
		if v.NvPairs.InterfaceDescription != c.NvPairs.InterfaceDescription {
			log.Printf("Update: v.NvPairs.InterfaceDescription=%v, c.NvPairs.InterfaceDescription=%v", v.NvPairs.InterfaceDescription, c.NvPairs.InterfaceDescription)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.InterfaceDescription=%v, c.NvPairs.InterfaceDescription=%v", v.NvPairs.InterfaceDescription, c.NvPairs.InterfaceDescription)
		v.NvPairs.InterfaceDescription = c.NvPairs.InterfaceDescription
	}

	if v.NvPairs.Vrf != "" {
		if v.NvPairs.Vrf != c.NvPairs.Vrf {
			log.Printf("Update: v.NvPairs.Vrf=%v, c.NvPairs.Vrf=%v", v.NvPairs.Vrf, c.NvPairs.Vrf)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.Vrf=%v, c.NvPairs.Vrf=%v", v.NvPairs.Vrf, c.NvPairs.Vrf)
		v.NvPairs.Vrf = c.NvPairs.Vrf
	}

	if v.NvPairs.Ipv4Address != "" {
		if v.NvPairs.Ipv4Address != c.NvPairs.Ipv4Address {
			log.Printf("Update: v.NvPairs.Ipv4Address=%v, c.NvPairs.Ipv4Address=%v", v.NvPairs.Ipv4Address, c.NvPairs.Ipv4Address)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.Ipv4Address=%v, c.NvPairs.Ipv4Address=%v", v.NvPairs.Ipv4Address, c.NvPairs.Ipv4Address)
		v.NvPairs.Ipv4Address = c.NvPairs.Ipv4Address
	}

	if v.NvPairs.Ipv6Address != "" {
		if v.NvPairs.Ipv6Address != c.NvPairs.Ipv6Address {
			log.Printf("Update: v.NvPairs.Ipv6Address=%v, c.NvPairs.Ipv6Address=%v", v.NvPairs.Ipv6Address, c.NvPairs.Ipv6Address)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.Ipv6Address=%v, c.NvPairs.Ipv6Address=%v", v.NvPairs.Ipv6Address, c.NvPairs.Ipv6Address)
		v.NvPairs.Ipv6Address = c.NvPairs.Ipv6Address
	}

	if v.NvPairs.RouteMapTag != "" {
		if v.NvPairs.RouteMapTag != c.NvPairs.RouteMapTag {
			log.Printf("Update: v.NvPairs.RouteMapTag=%v, c.NvPairs.RouteMapTag=%v", v.NvPairs.RouteMapTag, c.NvPairs.RouteMapTag)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.RouteMapTag=%v, c.NvPairs.RouteMapTag=%v", v.NvPairs.RouteMapTag, c.NvPairs.RouteMapTag)
		v.NvPairs.RouteMapTag = c.NvPairs.RouteMapTag
	}

	if v.NvPairs.BpduGuard != "" {
		if v.NvPairs.BpduGuard != c.NvPairs.BpduGuard {
			log.Printf("Update: v.NvPairs.BpduGuard=%v, c.NvPairs.BpduGuard=%v", v.NvPairs.BpduGuard, c.NvPairs.BpduGuard)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.BpduGuard=%v, c.NvPairs.BpduGuard=%v", v.NvPairs.BpduGuard, c.NvPairs.BpduGuard)
		v.NvPairs.BpduGuard = c.NvPairs.BpduGuard
	}

	if v.NvPairs.PortTypeFast != c.NvPairs.PortTypeFast {
		log.Printf("Update: v.NvPairs.PortTypeFast=%v, c.NvPairs.PortTypeFast=%v", v.NvPairs.PortTypeFast, c.NvPairs.PortTypeFast)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.Mtu != "" {
		if v.NvPairs.Mtu != c.NvPairs.Mtu {
			log.Printf("Update: v.NvPairs.Mtu=%v, c.NvPairs.Mtu=%v", v.NvPairs.Mtu, c.NvPairs.Mtu)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.Mtu=%v, c.NvPairs.Mtu=%v", v.NvPairs.Mtu, c.NvPairs.Mtu)
		v.NvPairs.Mtu = c.NvPairs.Mtu
	}

	if v.NvPairs.Speed != "" {
		if v.NvPairs.Speed != c.NvPairs.Speed {
			log.Printf("Update: v.NvPairs.Speed=%v, c.NvPairs.Speed=%v", v.NvPairs.Speed, c.NvPairs.Speed)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.Speed=%v, c.NvPairs.Speed=%v", v.NvPairs.Speed, c.NvPairs.Speed)
		v.NvPairs.Speed = c.NvPairs.Speed
	}

	if v.NvPairs.AccessVlan != nil && c.NvPairs.AccessVlan != nil {
		if *v.NvPairs.AccessVlan != *c.NvPairs.AccessVlan {
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
			log.Printf("Update: v.NvPairs.AccessVlan=%v, c.NvPairs.AccessVlan=%v", *v.NvPairs.AccessVlan, *c.NvPairs.AccessVlan)
		}
	} else if v.NvPairs.AccessVlan != nil {
		log.Printf("Update: v.NvPairs.AccessVlan=%v, c.NvPairs.AccessVlan=%v", *v.NvPairs.AccessVlan, *c.NvPairs.AccessVlan)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if c.NvPairs.AccessVlan != nil {
		v.NvPairs.AccessVlan = new(int64)
		log.Printf("Copy from state: v.NvPairs.AccessVlan=%v, c.NvPairs.AccessVlan=%v", *v.NvPairs.AccessVlan, *c.NvPairs.AccessVlan)
		*v.NvPairs.AccessVlan = *c.NvPairs.AccessVlan
	}

	if v.NvPairs.OrphanPort != c.NvPairs.OrphanPort {
		log.Printf("Update: v.NvPairs.OrphanPort=%v, c.NvPairs.OrphanPort=%v", v.NvPairs.OrphanPort, c.NvPairs.OrphanPort)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.Ptp != c.NvPairs.Ptp {
		log.Printf("Update: v.NvPairs.Ptp=%v, c.NvPairs.Ptp=%v", v.NvPairs.Ptp, c.NvPairs.Ptp)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.Netflow != c.NvPairs.Netflow {
		log.Printf("Update: v.NvPairs.Netflow=%v, c.NvPairs.Netflow=%v", v.NvPairs.Netflow, c.NvPairs.Netflow)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.NetflowMonitor != "" {
		if v.NvPairs.NetflowMonitor != c.NvPairs.NetflowMonitor {
			log.Printf("Update: v.NvPairs.NetflowMonitor=%v, c.NvPairs.NetflowMonitor=%v", v.NvPairs.NetflowMonitor, c.NvPairs.NetflowMonitor)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.NetflowMonitor=%v, c.NvPairs.NetflowMonitor=%v", v.NvPairs.NetflowMonitor, c.NvPairs.NetflowMonitor)
		v.NvPairs.NetflowMonitor = c.NvPairs.NetflowMonitor
	}

	if v.NvPairs.NetflowSampler != "" {
		if v.NvPairs.NetflowSampler != c.NvPairs.NetflowSampler {
			log.Printf("Update: v.NvPairs.NetflowSampler=%v, c.NvPairs.NetflowSampler=%v", v.NvPairs.NetflowSampler, c.NvPairs.NetflowSampler)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.NetflowSampler=%v, c.NvPairs.NetflowSampler=%v", v.NvPairs.NetflowSampler, c.NvPairs.NetflowSampler)
		v.NvPairs.NetflowSampler = c.NvPairs.NetflowSampler
	}

	if v.NvPairs.AllowedVlans != "" {
		if v.NvPairs.AllowedVlans != c.NvPairs.AllowedVlans {
			log.Printf("Update: v.NvPairs.AllowedVlans=%v, c.NvPairs.AllowedVlans=%v", v.NvPairs.AllowedVlans, c.NvPairs.AllowedVlans)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.AllowedVlans=%v, c.NvPairs.AllowedVlans=%v", v.NvPairs.AllowedVlans, c.NvPairs.AllowedVlans)
		v.NvPairs.AllowedVlans = c.NvPairs.AllowedVlans
	}

	if v.NvPairs.NativeVlan != nil && c.NvPairs.NativeVlan != nil {
		if *v.NvPairs.NativeVlan != *c.NvPairs.NativeVlan {
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
			log.Printf("Update: v.NvPairs.NativeVlan=%v, c.NvPairs.NativeVlan=%v", *v.NvPairs.NativeVlan, *c.NvPairs.NativeVlan)
		}
	} else if v.NvPairs.NativeVlan != nil {
		log.Printf("Update: v.NvPairs.NativeVlan=%v, c.NvPairs.NativeVlan=%v", *v.NvPairs.NativeVlan, *c.NvPairs.NativeVlan)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if c.NvPairs.NativeVlan != nil {
		v.NvPairs.NativeVlan = new(int64)
		log.Printf("Copy from state: v.NvPairs.NativeVlan=%v, c.NvPairs.NativeVlan=%v", *v.NvPairs.NativeVlan, *c.NvPairs.NativeVlan)
		*v.NvPairs.NativeVlan = *c.NvPairs.NativeVlan
	}

	return action
}
