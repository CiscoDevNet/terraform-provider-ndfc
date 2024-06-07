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

	if !v.NvPairs.AccessVlan.IsEmpty() && !c.NvPairs.AccessVlan.IsEmpty() {
		if *v.NvPairs.AccessVlan != *c.NvPairs.AccessVlan {
			log.Printf("v.NvPairs.AccessVlan=%v, c.NvPairs.AccessVlan=%v", *v.NvPairs.AccessVlan, *c.NvPairs.AccessVlan)
			return RequiresUpdate
		}
	} else {
		if !v.NvPairs.AccessVlan.IsEmpty() {
			log.Printf("v.NvPairs.AccessVlan=%v", *v.NvPairs.AccessVlan)
			return RequiresUpdate
		} else if !c.NvPairs.AccessVlan.IsEmpty() {
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

	if !v.NvPairs.NativeVlan.IsEmpty() && !c.NvPairs.NativeVlan.IsEmpty() {
		if *v.NvPairs.NativeVlan != *c.NvPairs.NativeVlan {
			log.Printf("v.NvPairs.NativeVlan=%v, c.NvPairs.NativeVlan=%v", *v.NvPairs.NativeVlan, *c.NvPairs.NativeVlan)
			return RequiresUpdate
		}
	} else {
		if !v.NvPairs.NativeVlan.IsEmpty() {
			log.Printf("v.NvPairs.NativeVlan=%v", *v.NvPairs.NativeVlan)
			return RequiresUpdate
		} else if !c.NvPairs.NativeVlan.IsEmpty() {
			log.Printf("c.NvPairs.NativeVlan=%v", *c.NvPairs.NativeVlan)
			return RequiresUpdate
		}
	}
	if v.NvPairs.Ipv4PrefixLength != c.NvPairs.Ipv4PrefixLength {
		log.Printf("v.NvPairs.Ipv4PrefixLength=%s, c.NvPairs.Ipv4PrefixLength=%s", v.NvPairs.Ipv4PrefixLength, c.NvPairs.Ipv4PrefixLength)
		return RequiresUpdate
	}
	if v.NvPairs.RoutingTag != c.NvPairs.RoutingTag {
		log.Printf("v.NvPairs.RoutingTag=%s, c.NvPairs.RoutingTag=%s", v.NvPairs.RoutingTag, c.NvPairs.RoutingTag)
		return RequiresUpdate
	}
	if v.NvPairs.DisableIpRedirects != c.NvPairs.DisableIpRedirects {
		log.Printf("v.NvPairs.DisableIpRedirects=%s, c.NvPairs.DisableIpRedirects=%s", v.NvPairs.DisableIpRedirects, c.NvPairs.DisableIpRedirects)
		return RequiresUpdate
	}
	if v.NvPairs.EnableHsrp != c.NvPairs.EnableHsrp {
		log.Printf("v.NvPairs.EnableHsrp=%s, c.NvPairs.EnableHsrp=%s", v.NvPairs.EnableHsrp, c.NvPairs.EnableHsrp)
		return RequiresUpdate
	}

	if !v.NvPairs.HsrpGroup.IsEmpty() && !c.NvPairs.HsrpGroup.IsEmpty() {
		if *v.NvPairs.HsrpGroup != *c.NvPairs.HsrpGroup {
			log.Printf("v.NvPairs.HsrpGroup=%v, c.NvPairs.HsrpGroup=%v", *v.NvPairs.HsrpGroup, *c.NvPairs.HsrpGroup)
			return RequiresUpdate
		}
	} else {
		if !v.NvPairs.HsrpGroup.IsEmpty() {
			log.Printf("v.NvPairs.HsrpGroup=%v", *v.NvPairs.HsrpGroup)
			return RequiresUpdate
		} else if !c.NvPairs.HsrpGroup.IsEmpty() {
			log.Printf("c.NvPairs.HsrpGroup=%v", *c.NvPairs.HsrpGroup)
			return RequiresUpdate
		}
	}
	if v.NvPairs.HsrpVip != c.NvPairs.HsrpVip {
		log.Printf("v.NvPairs.HsrpVip=%s, c.NvPairs.HsrpVip=%s", v.NvPairs.HsrpVip, c.NvPairs.HsrpVip)
		return RequiresUpdate
	}

	if !v.NvPairs.HsrpPriority.IsEmpty() && !c.NvPairs.HsrpPriority.IsEmpty() {
		if *v.NvPairs.HsrpPriority != *c.NvPairs.HsrpPriority {
			log.Printf("v.NvPairs.HsrpPriority=%v, c.NvPairs.HsrpPriority=%v", *v.NvPairs.HsrpPriority, *c.NvPairs.HsrpPriority)
			return RequiresUpdate
		}
	} else {
		if !v.NvPairs.HsrpPriority.IsEmpty() {
			log.Printf("v.NvPairs.HsrpPriority=%v", *v.NvPairs.HsrpPriority)
			return RequiresUpdate
		} else if !c.NvPairs.HsrpPriority.IsEmpty() {
			log.Printf("c.NvPairs.HsrpPriority=%v", *c.NvPairs.HsrpPriority)
			return RequiresUpdate
		}
	}
	if v.NvPairs.HsrpVersion != c.NvPairs.HsrpVersion {
		log.Printf("v.NvPairs.HsrpVersion=%s, c.NvPairs.HsrpVersion=%s", v.NvPairs.HsrpVersion, c.NvPairs.HsrpVersion)
		return RequiresUpdate
	}
	if v.NvPairs.Preempt != c.NvPairs.Preempt {
		log.Printf("v.NvPairs.Preempt=%s, c.NvPairs.Preempt=%s", v.NvPairs.Preempt, c.NvPairs.Preempt)
		return RequiresUpdate
	}
	if v.NvPairs.Mac != c.NvPairs.Mac {
		log.Printf("v.NvPairs.Mac=%s, c.NvPairs.Mac=%s", v.NvPairs.Mac, c.NvPairs.Mac)
		return RequiresUpdate
	}
	if v.NvPairs.DhcpServerAddr1 != c.NvPairs.DhcpServerAddr1 {
		log.Printf("v.NvPairs.DhcpServerAddr1=%s, c.NvPairs.DhcpServerAddr1=%s", v.NvPairs.DhcpServerAddr1, c.NvPairs.DhcpServerAddr1)
		return RequiresUpdate
	}
	if v.NvPairs.DhcpServerAddr2 != c.NvPairs.DhcpServerAddr2 {
		log.Printf("v.NvPairs.DhcpServerAddr2=%s, c.NvPairs.DhcpServerAddr2=%s", v.NvPairs.DhcpServerAddr2, c.NvPairs.DhcpServerAddr2)
		return RequiresUpdate
	}
	if v.NvPairs.DhcpServerAddr3 != c.NvPairs.DhcpServerAddr3 {
		log.Printf("v.NvPairs.DhcpServerAddr3=%s, c.NvPairs.DhcpServerAddr3=%s", v.NvPairs.DhcpServerAddr3, c.NvPairs.DhcpServerAddr3)
		return RequiresUpdate
	}
	if v.NvPairs.VrfDhcp1 != c.NvPairs.VrfDhcp1 {
		log.Printf("v.NvPairs.VrfDhcp1=%s, c.NvPairs.VrfDhcp1=%s", v.NvPairs.VrfDhcp1, c.NvPairs.VrfDhcp1)
		return RequiresUpdate
	}
	if v.NvPairs.VrfDhcp2 != c.NvPairs.VrfDhcp2 {
		log.Printf("v.NvPairs.VrfDhcp2=%s, c.NvPairs.VrfDhcp2=%s", v.NvPairs.VrfDhcp2, c.NvPairs.VrfDhcp2)
		return RequiresUpdate
	}
	if v.NvPairs.VrfDhcp3 != c.NvPairs.VrfDhcp3 {
		log.Printf("v.NvPairs.VrfDhcp3=%s, c.NvPairs.VrfDhcp3=%s", v.NvPairs.VrfDhcp3, c.NvPairs.VrfDhcp3)
		return RequiresUpdate
	}
	if v.NvPairs.AdvertiseSubnetInUnderlay != c.NvPairs.AdvertiseSubnetInUnderlay {
		log.Printf("v.NvPairs.AdvertiseSubnetInUnderlay=%s, c.NvPairs.AdvertiseSubnetInUnderlay=%s", v.NvPairs.AdvertiseSubnetInUnderlay, c.NvPairs.AdvertiseSubnetInUnderlay)
		return RequiresUpdate
	}
	if v.NvPairs.CopyPoDescription != c.NvPairs.CopyPoDescription {
		log.Printf("v.NvPairs.CopyPoDescription=%s, c.NvPairs.CopyPoDescription=%s", v.NvPairs.CopyPoDescription, c.NvPairs.CopyPoDescription)
		return RequiresUpdate
	}
	if v.NvPairs.PortchannelMode != c.NvPairs.PortchannelMode {
		log.Printf("v.NvPairs.PortchannelMode=%s, c.NvPairs.PortchannelMode=%s", v.NvPairs.PortchannelMode, c.NvPairs.PortchannelMode)
		return RequiresUpdate
	}
	if v.NvPairs.MemberInterfaces != c.NvPairs.MemberInterfaces {
		log.Printf("v.NvPairs.MemberInterfaces=%s, c.NvPairs.MemberInterfaces=%s", v.NvPairs.MemberInterfaces, c.NvPairs.MemberInterfaces)
		return RequiresUpdate
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

	if !v.NvPairs.AccessVlan.IsEmpty() && !c.NvPairs.AccessVlan.IsEmpty() {
		if *v.NvPairs.AccessVlan != *c.NvPairs.AccessVlan {
			log.Printf("Update: v.NvPairs.AccessVlan=%v, c.NvPairs.AccessVlan=%v", *v.NvPairs.AccessVlan, *c.NvPairs.AccessVlan)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else if !v.NvPairs.AccessVlan.IsEmpty() {
		log.Printf("Update: v.NvPairs.AccessVlan=%v", *v.NvPairs.AccessVlan)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if !c.NvPairs.AccessVlan.IsEmpty() {
		log.Printf("Copy from State: c.NvPairs.AccessVlan=%v", *c.NvPairs.AccessVlan)
		v.NvPairs.AccessVlan = new(Int64Custom)
		*v.NvPairs.AccessVlan = *c.NvPairs.AccessVlan
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
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

	if !v.NvPairs.NativeVlan.IsEmpty() && !c.NvPairs.NativeVlan.IsEmpty() {
		if *v.NvPairs.NativeVlan != *c.NvPairs.NativeVlan {
			log.Printf("Update: v.NvPairs.NativeVlan=%v, c.NvPairs.NativeVlan=%v", *v.NvPairs.NativeVlan, *c.NvPairs.NativeVlan)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else if !v.NvPairs.NativeVlan.IsEmpty() {
		log.Printf("Update: v.NvPairs.NativeVlan=%v", *v.NvPairs.NativeVlan)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if !c.NvPairs.NativeVlan.IsEmpty() {
		log.Printf("Copy from State: c.NvPairs.NativeVlan=%v", *c.NvPairs.NativeVlan)
		v.NvPairs.NativeVlan = new(Int64Custom)
		*v.NvPairs.NativeVlan = *c.NvPairs.NativeVlan
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.Ipv4PrefixLength != "" {
		if v.NvPairs.Ipv4PrefixLength != c.NvPairs.Ipv4PrefixLength {
			log.Printf("Update: v.NvPairs.Ipv4PrefixLength=%v, c.NvPairs.Ipv4PrefixLength=%v", v.NvPairs.Ipv4PrefixLength, c.NvPairs.Ipv4PrefixLength)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.Ipv4PrefixLength=%v, c.NvPairs.Ipv4PrefixLength=%v", v.NvPairs.Ipv4PrefixLength, c.NvPairs.Ipv4PrefixLength)
		v.NvPairs.Ipv4PrefixLength = c.NvPairs.Ipv4PrefixLength
	}

	if v.NvPairs.RoutingTag != "" {
		if v.NvPairs.RoutingTag != c.NvPairs.RoutingTag {
			log.Printf("Update: v.NvPairs.RoutingTag=%v, c.NvPairs.RoutingTag=%v", v.NvPairs.RoutingTag, c.NvPairs.RoutingTag)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.RoutingTag=%v, c.NvPairs.RoutingTag=%v", v.NvPairs.RoutingTag, c.NvPairs.RoutingTag)
		v.NvPairs.RoutingTag = c.NvPairs.RoutingTag
	}

	if v.NvPairs.DisableIpRedirects != c.NvPairs.DisableIpRedirects {
		log.Printf("Update: v.NvPairs.DisableIpRedirects=%v, c.NvPairs.DisableIpRedirects=%v", v.NvPairs.DisableIpRedirects, c.NvPairs.DisableIpRedirects)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.EnableHsrp != c.NvPairs.EnableHsrp {
		log.Printf("Update: v.NvPairs.EnableHsrp=%v, c.NvPairs.EnableHsrp=%v", v.NvPairs.EnableHsrp, c.NvPairs.EnableHsrp)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if !v.NvPairs.HsrpGroup.IsEmpty() && !c.NvPairs.HsrpGroup.IsEmpty() {
		if *v.NvPairs.HsrpGroup != *c.NvPairs.HsrpGroup {
			log.Printf("Update: v.NvPairs.HsrpGroup=%v, c.NvPairs.HsrpGroup=%v", *v.NvPairs.HsrpGroup, *c.NvPairs.HsrpGroup)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else if !v.NvPairs.HsrpGroup.IsEmpty() {
		log.Printf("Update: v.NvPairs.HsrpGroup=%v", *v.NvPairs.HsrpGroup)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if !c.NvPairs.HsrpGroup.IsEmpty() {
		log.Printf("Copy from State: c.NvPairs.HsrpGroup=%v", *c.NvPairs.HsrpGroup)
		v.NvPairs.HsrpGroup = new(Int64Custom)
		*v.NvPairs.HsrpGroup = *c.NvPairs.HsrpGroup
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.HsrpVip != "" {
		if v.NvPairs.HsrpVip != c.NvPairs.HsrpVip {
			log.Printf("Update: v.NvPairs.HsrpVip=%v, c.NvPairs.HsrpVip=%v", v.NvPairs.HsrpVip, c.NvPairs.HsrpVip)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.HsrpVip=%v, c.NvPairs.HsrpVip=%v", v.NvPairs.HsrpVip, c.NvPairs.HsrpVip)
		v.NvPairs.HsrpVip = c.NvPairs.HsrpVip
	}

	if !v.NvPairs.HsrpPriority.IsEmpty() && !c.NvPairs.HsrpPriority.IsEmpty() {
		if *v.NvPairs.HsrpPriority != *c.NvPairs.HsrpPriority {
			log.Printf("Update: v.NvPairs.HsrpPriority=%v, c.NvPairs.HsrpPriority=%v", *v.NvPairs.HsrpPriority, *c.NvPairs.HsrpPriority)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else if !v.NvPairs.HsrpPriority.IsEmpty() {
		log.Printf("Update: v.NvPairs.HsrpPriority=%v", *v.NvPairs.HsrpPriority)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if !c.NvPairs.HsrpPriority.IsEmpty() {
		log.Printf("Copy from State: c.NvPairs.HsrpPriority=%v", *c.NvPairs.HsrpPriority)
		v.NvPairs.HsrpPriority = new(Int64Custom)
		*v.NvPairs.HsrpPriority = *c.NvPairs.HsrpPriority
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.HsrpVersion != "" {
		if v.NvPairs.HsrpVersion != c.NvPairs.HsrpVersion {
			log.Printf("Update: v.NvPairs.HsrpVersion=%v, c.NvPairs.HsrpVersion=%v", v.NvPairs.HsrpVersion, c.NvPairs.HsrpVersion)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.HsrpVersion=%v, c.NvPairs.HsrpVersion=%v", v.NvPairs.HsrpVersion, c.NvPairs.HsrpVersion)
		v.NvPairs.HsrpVersion = c.NvPairs.HsrpVersion
	}

	if v.NvPairs.Preempt != c.NvPairs.Preempt {
		log.Printf("Update: v.NvPairs.Preempt=%v, c.NvPairs.Preempt=%v", v.NvPairs.Preempt, c.NvPairs.Preempt)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.Mac != "" {
		if v.NvPairs.Mac != c.NvPairs.Mac {
			log.Printf("Update: v.NvPairs.Mac=%v, c.NvPairs.Mac=%v", v.NvPairs.Mac, c.NvPairs.Mac)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.Mac=%v, c.NvPairs.Mac=%v", v.NvPairs.Mac, c.NvPairs.Mac)
		v.NvPairs.Mac = c.NvPairs.Mac
	}

	if v.NvPairs.DhcpServerAddr1 != "" {
		if v.NvPairs.DhcpServerAddr1 != c.NvPairs.DhcpServerAddr1 {
			log.Printf("Update: v.NvPairs.DhcpServerAddr1=%v, c.NvPairs.DhcpServerAddr1=%v", v.NvPairs.DhcpServerAddr1, c.NvPairs.DhcpServerAddr1)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.DhcpServerAddr1=%v, c.NvPairs.DhcpServerAddr1=%v", v.NvPairs.DhcpServerAddr1, c.NvPairs.DhcpServerAddr1)
		v.NvPairs.DhcpServerAddr1 = c.NvPairs.DhcpServerAddr1
	}

	if v.NvPairs.DhcpServerAddr2 != "" {
		if v.NvPairs.DhcpServerAddr2 != c.NvPairs.DhcpServerAddr2 {
			log.Printf("Update: v.NvPairs.DhcpServerAddr2=%v, c.NvPairs.DhcpServerAddr2=%v", v.NvPairs.DhcpServerAddr2, c.NvPairs.DhcpServerAddr2)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.DhcpServerAddr2=%v, c.NvPairs.DhcpServerAddr2=%v", v.NvPairs.DhcpServerAddr2, c.NvPairs.DhcpServerAddr2)
		v.NvPairs.DhcpServerAddr2 = c.NvPairs.DhcpServerAddr2
	}

	if v.NvPairs.DhcpServerAddr3 != "" {
		if v.NvPairs.DhcpServerAddr3 != c.NvPairs.DhcpServerAddr3 {
			log.Printf("Update: v.NvPairs.DhcpServerAddr3=%v, c.NvPairs.DhcpServerAddr3=%v", v.NvPairs.DhcpServerAddr3, c.NvPairs.DhcpServerAddr3)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.DhcpServerAddr3=%v, c.NvPairs.DhcpServerAddr3=%v", v.NvPairs.DhcpServerAddr3, c.NvPairs.DhcpServerAddr3)
		v.NvPairs.DhcpServerAddr3 = c.NvPairs.DhcpServerAddr3
	}

	if v.NvPairs.VrfDhcp1 != "" {
		if v.NvPairs.VrfDhcp1 != c.NvPairs.VrfDhcp1 {
			log.Printf("Update: v.NvPairs.VrfDhcp1=%v, c.NvPairs.VrfDhcp1=%v", v.NvPairs.VrfDhcp1, c.NvPairs.VrfDhcp1)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.VrfDhcp1=%v, c.NvPairs.VrfDhcp1=%v", v.NvPairs.VrfDhcp1, c.NvPairs.VrfDhcp1)
		v.NvPairs.VrfDhcp1 = c.NvPairs.VrfDhcp1
	}

	if v.NvPairs.VrfDhcp2 != "" {
		if v.NvPairs.VrfDhcp2 != c.NvPairs.VrfDhcp2 {
			log.Printf("Update: v.NvPairs.VrfDhcp2=%v, c.NvPairs.VrfDhcp2=%v", v.NvPairs.VrfDhcp2, c.NvPairs.VrfDhcp2)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.VrfDhcp2=%v, c.NvPairs.VrfDhcp2=%v", v.NvPairs.VrfDhcp2, c.NvPairs.VrfDhcp2)
		v.NvPairs.VrfDhcp2 = c.NvPairs.VrfDhcp2
	}

	if v.NvPairs.VrfDhcp3 != "" {
		if v.NvPairs.VrfDhcp3 != c.NvPairs.VrfDhcp3 {
			log.Printf("Update: v.NvPairs.VrfDhcp3=%v, c.NvPairs.VrfDhcp3=%v", v.NvPairs.VrfDhcp3, c.NvPairs.VrfDhcp3)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.VrfDhcp3=%v, c.NvPairs.VrfDhcp3=%v", v.NvPairs.VrfDhcp3, c.NvPairs.VrfDhcp3)
		v.NvPairs.VrfDhcp3 = c.NvPairs.VrfDhcp3
	}

	if v.NvPairs.AdvertiseSubnetInUnderlay != c.NvPairs.AdvertiseSubnetInUnderlay {
		log.Printf("Update: v.NvPairs.AdvertiseSubnetInUnderlay=%v, c.NvPairs.AdvertiseSubnetInUnderlay=%v", v.NvPairs.AdvertiseSubnetInUnderlay, c.NvPairs.AdvertiseSubnetInUnderlay)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.CopyPoDescription != c.NvPairs.CopyPoDescription {
		log.Printf("Update: v.NvPairs.CopyPoDescription=%v, c.NvPairs.CopyPoDescription=%v", v.NvPairs.CopyPoDescription, c.NvPairs.CopyPoDescription)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.PortchannelMode != "" {
		if v.NvPairs.PortchannelMode != c.NvPairs.PortchannelMode {
			log.Printf("Update: v.NvPairs.PortchannelMode=%v, c.NvPairs.PortchannelMode=%v", v.NvPairs.PortchannelMode, c.NvPairs.PortchannelMode)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.PortchannelMode=%v, c.NvPairs.PortchannelMode=%v", v.NvPairs.PortchannelMode, c.NvPairs.PortchannelMode)
		v.NvPairs.PortchannelMode = c.NvPairs.PortchannelMode
	}

	if v.NvPairs.MemberInterfaces != "" {
		if v.NvPairs.MemberInterfaces != c.NvPairs.MemberInterfaces {
			log.Printf("Update: v.NvPairs.MemberInterfaces=%v, c.NvPairs.MemberInterfaces=%v", v.NvPairs.MemberInterfaces, c.NvPairs.MemberInterfaces)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.MemberInterfaces=%v, c.NvPairs.MemberInterfaces=%v", v.NvPairs.MemberInterfaces, c.NvPairs.MemberInterfaces)
		v.NvPairs.MemberInterfaces = c.NvPairs.MemberInterfaces
	}

	return action
}
