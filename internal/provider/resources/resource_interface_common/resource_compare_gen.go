// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package resource_interface_common

import (
	"log"
	. "terraform-provider-ndfc/internal/provider/types"
)

func (v NDFCPVlanMappingListValue) DeepEqual(c NDFCPVlanMappingListValue) int {
	cf := false
	if v.SVlan != c.SVlan {
		log.Printf("v.SVlan=%v, c.SVlan=%v", v.SVlan, c.SVlan)
		return RequiresUpdate
	}
	if v.PVlan != c.PVlan {
		log.Printf("v.PVlan=%v, c.PVlan=%v", v.PVlan, c.PVlan)
		return RequiresUpdate
	}

	if cf {
		return ControlFlagUpdate
	}
	return ValuesDeeplyEqual
}

func (v NDFCPVlanAssocListValue) DeepEqual(c NDFCPVlanAssocListValue) int {
	cf := false
	if v.SVlan != c.SVlan {
		log.Printf("v.SVlan=%v, c.SVlan=%v", v.SVlan, c.SVlan)
		return RequiresUpdate
	}
	if v.PVlan != c.PVlan {
		log.Printf("v.PVlan=%v, c.PVlan=%v", v.PVlan, c.PVlan)
		return RequiresUpdate
	}

	if cf {
		return ControlFlagUpdate
	}
	return ValuesDeeplyEqual
}

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
	if v.NvPairs.Peer1PoFreeformConfig != c.NvPairs.Peer1PoFreeformConfig {
		log.Printf("v.NvPairs.Peer1PoFreeformConfig=%s, c.NvPairs.Peer1PoFreeformConfig=%s", v.NvPairs.Peer1PoFreeformConfig, c.NvPairs.Peer1PoFreeformConfig)
		return RequiresUpdate
	}
	if v.NvPairs.Peer2PoFreeformConfig != c.NvPairs.Peer2PoFreeformConfig {
		log.Printf("v.NvPairs.Peer2PoFreeformConfig=%s, c.NvPairs.Peer2PoFreeformConfig=%s", v.NvPairs.Peer2PoFreeformConfig, c.NvPairs.Peer2PoFreeformConfig)
		return RequiresUpdate
	}
	if v.NvPairs.Peer1PoDescription != c.NvPairs.Peer1PoDescription {
		log.Printf("v.NvPairs.Peer1PoDescription=%s, c.NvPairs.Peer1PoDescription=%s", v.NvPairs.Peer1PoDescription, c.NvPairs.Peer1PoDescription)
		return RequiresUpdate
	}
	if v.NvPairs.Peer2PoDescription != c.NvPairs.Peer2PoDescription {
		log.Printf("v.NvPairs.Peer2PoDescription=%s, c.NvPairs.Peer2PoDescription=%s", v.NvPairs.Peer2PoDescription, c.NvPairs.Peer2PoDescription)
		return RequiresUpdate
	}
	if v.NvPairs.Peer1AllowedVlans != c.NvPairs.Peer1AllowedVlans {
		log.Printf("v.NvPairs.Peer1AllowedVlans=%s, c.NvPairs.Peer1AllowedVlans=%s", v.NvPairs.Peer1AllowedVlans, c.NvPairs.Peer1AllowedVlans)
		return RequiresUpdate
	}
	if v.NvPairs.Peer2AllowedVlans != c.NvPairs.Peer2AllowedVlans {
		log.Printf("v.NvPairs.Peer2AllowedVlans=%s, c.NvPairs.Peer2AllowedVlans=%s", v.NvPairs.Peer2AllowedVlans, c.NvPairs.Peer2AllowedVlans)
		return RequiresUpdate
	}

	if !v.NvPairs.Peer1NativeVlan.IsEmpty() && !c.NvPairs.Peer1NativeVlan.IsEmpty() {
		if *v.NvPairs.Peer1NativeVlan != *c.NvPairs.Peer1NativeVlan {
			log.Printf("v.NvPairs.Peer1NativeVlan=%v, c.NvPairs.Peer1NativeVlan=%v", *v.NvPairs.Peer1NativeVlan, *c.NvPairs.Peer1NativeVlan)
			return RequiresUpdate
		}
	} else {
		if !v.NvPairs.Peer1NativeVlan.IsEmpty() {
			log.Printf("v.NvPairs.Peer1NativeVlan=%v", *v.NvPairs.Peer1NativeVlan)
			return RequiresUpdate
		} else if !c.NvPairs.Peer1NativeVlan.IsEmpty() {
			log.Printf("c.NvPairs.Peer1NativeVlan=%v", *c.NvPairs.Peer1NativeVlan)
			return RequiresUpdate
		}
	}

	if !v.NvPairs.Peer2NativeVlan.IsEmpty() && !c.NvPairs.Peer2NativeVlan.IsEmpty() {
		if *v.NvPairs.Peer2NativeVlan != *c.NvPairs.Peer2NativeVlan {
			log.Printf("v.NvPairs.Peer2NativeVlan=%v, c.NvPairs.Peer2NativeVlan=%v", *v.NvPairs.Peer2NativeVlan, *c.NvPairs.Peer2NativeVlan)
			return RequiresUpdate
		}
	} else {
		if !v.NvPairs.Peer2NativeVlan.IsEmpty() {
			log.Printf("v.NvPairs.Peer2NativeVlan=%v", *v.NvPairs.Peer2NativeVlan)
			return RequiresUpdate
		} else if !c.NvPairs.Peer2NativeVlan.IsEmpty() {
			log.Printf("c.NvPairs.Peer2NativeVlan=%v", *c.NvPairs.Peer2NativeVlan)
			return RequiresUpdate
		}
	}
	if v.NvPairs.Peer1MemberInterfaces != c.NvPairs.Peer1MemberInterfaces {
		log.Printf("v.NvPairs.Peer1MemberInterfaces=%s, c.NvPairs.Peer1MemberInterfaces=%s", v.NvPairs.Peer1MemberInterfaces, c.NvPairs.Peer1MemberInterfaces)
		return RequiresUpdate
	}
	if v.NvPairs.Peer2MemberInterfaces != c.NvPairs.Peer2MemberInterfaces {
		log.Printf("v.NvPairs.Peer2MemberInterfaces=%s, c.NvPairs.Peer2MemberInterfaces=%s", v.NvPairs.Peer2MemberInterfaces, c.NvPairs.Peer2MemberInterfaces)
		return RequiresUpdate
	}

	if v.NvPairs.Peer1PortChannelId != nil && c.NvPairs.Peer1PortChannelId != nil {
		if *v.NvPairs.Peer1PortChannelId != *c.NvPairs.Peer1PortChannelId {
			log.Printf("v.NvPairs.Peer1PortChannelId=%v, c.NvPairs.Peer1PortChannelId=%v", *v.NvPairs.Peer1PortChannelId, *c.NvPairs.Peer1PortChannelId)
			return RequiresReplace
		}
	} else {
		if v.NvPairs.Peer1PortChannelId != nil {
			log.Printf("v.NvPairs.Peer1PortChannelId=%v", *v.NvPairs.Peer1PortChannelId)
			return RequiresReplace

		} else if c.NvPairs.Peer1PortChannelId != nil {
			log.Printf("c.NvPairs.Peer1PortChannelId=%v", *c.NvPairs.Peer1PortChannelId)
			return RequiresReplace
		}
	}

	if !v.NvPairs.Peer2PortChannelId.IsEmpty() && !c.NvPairs.Peer2PortChannelId.IsEmpty() {
		if *v.NvPairs.Peer2PortChannelId != *c.NvPairs.Peer2PortChannelId {
			log.Printf("v.NvPairs.Peer2PortChannelId=%v, c.NvPairs.Peer2PortChannelId=%v", *v.NvPairs.Peer2PortChannelId, *c.NvPairs.Peer2PortChannelId)
			return RequiresReplace
		}
	} else {
		if !v.NvPairs.Peer2PortChannelId.IsEmpty() {
			log.Printf("v.NvPairs.Peer2PortChannelId=%v", *v.NvPairs.Peer2PortChannelId)
			return RequiresReplace
		} else if !c.NvPairs.Peer2PortChannelId.IsEmpty() {
			log.Printf("c.NvPairs.Peer2PortChannelId=%v", *c.NvPairs.Peer2PortChannelId)
			return RequiresReplace
		}
	}
	if v.NvPairs.EnablePimSparse != c.NvPairs.EnablePimSparse {
		log.Printf("v.NvPairs.EnablePimSparse=%s, c.NvPairs.EnablePimSparse=%s", v.NvPairs.EnablePimSparse, c.NvPairs.EnablePimSparse)
		return RequiresUpdate
	}
	if v.NvPairs.PimDrPriority != c.NvPairs.PimDrPriority {
		log.Printf("v.NvPairs.PimDrPriority=%s, c.NvPairs.PimDrPriority=%s", v.NvPairs.PimDrPriority, c.NvPairs.PimDrPriority)
		return RequiresUpdate
	}
	if v.NvPairs.EnablePfc != c.NvPairs.EnablePfc {
		log.Printf("v.NvPairs.EnablePfc=%s, c.NvPairs.EnablePfc=%s", v.NvPairs.EnablePfc, c.NvPairs.EnablePfc)
		return RequiresUpdate
	}
	if v.NvPairs.EnableQos != c.NvPairs.EnableQos {
		log.Printf("v.NvPairs.EnableQos=%s, c.NvPairs.EnableQos=%s", v.NvPairs.EnableQos, c.NvPairs.EnableQos)
		return RequiresUpdate
	}
	if v.NvPairs.QosPolicy != c.NvPairs.QosPolicy {
		log.Printf("v.NvPairs.QosPolicy=%s, c.NvPairs.QosPolicy=%s", v.NvPairs.QosPolicy, c.NvPairs.QosPolicy)
		return RequiresUpdate
	}
	if v.NvPairs.QueuingPolicy != c.NvPairs.QueuingPolicy {
		log.Printf("v.NvPairs.QueuingPolicy=%s, c.NvPairs.QueuingPolicy=%s", v.NvPairs.QueuingPolicy, c.NvPairs.QueuingPolicy)
		return RequiresUpdate
	}
	if v.NvPairs.LinkStateRoutingProtocol != "" {
		if v.NvPairs.LinkStateRoutingProtocol != c.NvPairs.LinkStateRoutingProtocol {
			log.Printf("v.NvPairs.LinkStateRoutingProtocol=%s, c.NvPairs.LinkStateRoutingProtocol=%s", v.NvPairs.LinkStateRoutingProtocol, c.NvPairs.LinkStateRoutingProtocol)
			return RequiresUpdate
		}
	} else {
		log.Printf("Skipping - v.NvPairs.LinkStateRoutingProtocol=%s, c.NvPairs.LinkStateRoutingProtocol=%s", v.NvPairs.LinkStateRoutingProtocol, c.NvPairs.LinkStateRoutingProtocol)
	}
	if v.NvPairs.LinkStateRoutingTag != "" {
		if v.NvPairs.LinkStateRoutingTag != c.NvPairs.LinkStateRoutingTag {
			log.Printf("v.NvPairs.LinkStateRoutingTag=%s, c.NvPairs.LinkStateRoutingTag=%s", v.NvPairs.LinkStateRoutingTag, c.NvPairs.LinkStateRoutingTag)
			return RequiresUpdate
		}
	} else {
		log.Printf("Skipping - v.NvPairs.LinkStateRoutingTag=%s, c.NvPairs.LinkStateRoutingTag=%s", v.NvPairs.LinkStateRoutingTag, c.NvPairs.LinkStateRoutingTag)
	}
	if v.NvPairs.Ipv6Addr != c.NvPairs.Ipv6Addr {
		log.Printf("v.NvPairs.Ipv6Addr=%s, c.NvPairs.Ipv6Addr=%s", v.NvPairs.Ipv6Addr, c.NvPairs.Ipv6Addr)
		return RequiresUpdate
	}
	if v.NvPairs.Ipv6PrefixLength != c.NvPairs.Ipv6PrefixLength {
		log.Printf("v.NvPairs.Ipv6PrefixLength=%s, c.NvPairs.Ipv6PrefixLength=%s", v.NvPairs.Ipv6PrefixLength, c.NvPairs.Ipv6PrefixLength)
		return RequiresUpdate
	}
	if v.NvPairs.CdpEnable != c.NvPairs.CdpEnable {
		log.Printf("v.NvPairs.CdpEnable=%s, c.NvPairs.CdpEnable=%s", v.NvPairs.CdpEnable, c.NvPairs.CdpEnable)
		return RequiresUpdate
	}
	if v.NvPairs.PortDuplexMode != c.NvPairs.PortDuplexMode {
		log.Printf("v.NvPairs.PortDuplexMode=%s, c.NvPairs.PortDuplexMode=%s", v.NvPairs.PortDuplexMode, c.NvPairs.PortDuplexMode)
		return RequiresUpdate
	}
	if v.NvPairs.EnableMonitor != c.NvPairs.EnableMonitor {
		log.Printf("v.NvPairs.EnableMonitor=%s, c.NvPairs.EnableMonitor=%s", v.NvPairs.EnableMonitor, c.NvPairs.EnableMonitor)
		return RequiresUpdate
	}
	if v.NvPairs.PvlanMode != c.NvPairs.PvlanMode {
		log.Printf("v.NvPairs.PvlanMode=%s, c.NvPairs.PvlanMode=%s", v.NvPairs.PvlanMode, c.NvPairs.PvlanMode)
		return RequiresUpdate
	}
	if v.NvPairs.PvlanAllowedVlans != c.NvPairs.PvlanAllowedVlans {
		log.Printf("v.NvPairs.PvlanAllowedVlans=%s, c.NvPairs.PvlanAllowedVlans=%s", v.NvPairs.PvlanAllowedVlans, c.NvPairs.PvlanAllowedVlans)
		return RequiresUpdate
	}
	if v.NvPairs.PvlanNativeVlan != c.NvPairs.PvlanNativeVlan {
		log.Printf("v.NvPairs.PvlanNativeVlan=%s, c.NvPairs.PvlanNativeVlan=%s", v.NvPairs.PvlanNativeVlan, c.NvPairs.PvlanNativeVlan)
		return RequiresUpdate
	}
	if v.NvPairs.IgForFex != c.NvPairs.IgForFex {
		log.Printf("v.NvPairs.IgForFex=%s, c.NvPairs.IgForFex=%s", v.NvPairs.IgForFex, c.NvPairs.IgForFex)
		return RequiresUpdate
	}
	if v.NvPairs.AutoNegotiate != c.NvPairs.AutoNegotiate {
		log.Printf("v.NvPairs.AutoNegotiate=%s, c.NvPairs.AutoNegotiate=%s", v.NvPairs.AutoNegotiate, c.NvPairs.AutoNegotiate)
		return RequiresUpdate
	}

	if v.NvPairs.PathCost != nil && c.NvPairs.PathCost != nil {
		if *v.NvPairs.PathCost != *c.NvPairs.PathCost {
			log.Printf("v.NvPairs.PathCost=%v, c.NvPairs.PathCost=%v", *v.NvPairs.PathCost, *c.NvPairs.PathCost)
			return RequiresUpdate
		}
	} else {
		if v.NvPairs.PathCost != nil {
			log.Printf("v.NvPairs.PathCost=%v", *v.NvPairs.PathCost)
			return RequiresUpdate
		} else if c.NvPairs.PathCost != nil {
			log.Printf("c.NvPairs.PathCost=%v", *c.NvPairs.PathCost)
			return RequiresUpdate
		}
	}

	if v.NvPairs.GuardMode != c.NvPairs.GuardMode {
		log.Printf("v.NvPairs.GuardMode=%s, c.NvPairs.GuardMode=%s", v.NvPairs.GuardMode, c.NvPairs.GuardMode)
		return RequiresUpdate
	}
	if v.NvPairs.Ttag != c.NvPairs.Ttag {
		log.Printf("v.NvPairs.Ttag=%s, c.NvPairs.Ttag=%s", v.NvPairs.Ttag, c.NvPairs.Ttag)
		return RequiresUpdate
	}

	if len(v.NvPairs.PVlanMappingList) != len(c.NvPairs.PVlanMappingList) {
		log.Printf("len(v.NvPairs.PVlanMappingList)=%d, len(c.NvPairs.PVlanMappingList)=%d", len(v.NvPairs.PVlanMappingList), len(c.NvPairs.PVlanMappingList))
		return RequiresUpdate
	}
	for i := range v.NvPairs.PVlanMappingList {
		retVal := v.NvPairs.PVlanMappingList[i].DeepEqual(c.NvPairs.PVlanMappingList[i])
		if retVal != ValuesDeeplyEqual {
			return retVal
		}
	}

	if len(v.NvPairs.PVlanAssocList) != len(c.NvPairs.PVlanAssocList) {
		log.Printf("len(v.NvPairs.PVlanAssocList)=%d, len(c.NvPairs.PVlanAssocList)=%d", len(v.NvPairs.PVlanAssocList), len(c.NvPairs.PVlanAssocList))
		return RequiresUpdate
	}
	for i := range v.NvPairs.PVlanAssocList {
		retVal := v.NvPairs.PVlanAssocList[i].DeepEqual(c.NvPairs.PVlanAssocList[i])
		if retVal != ValuesDeeplyEqual {
			return retVal
		}
	}

	if cf {
		return ControlFlagUpdate
	}
	return ValuesDeeplyEqual
}

func (v *NDFCPVlanMappingListValue) CreatePlan(c NDFCPVlanMappingListValue, cf *bool) int {
	action := ActionNone

	if v.SVlan != c.SVlan {
		log.Printf("SVlan-Update: v.SVlan=%v, c.SVlan=%v", v.SVlan, c.SVlan)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.PVlan != c.PVlan {
		log.Printf("PVlan-Update: v.PVlan=%v, c.PVlan=%v", v.PVlan, c.PVlan)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	return action
}

func (v *NDFCPVlanAssocListValue) CreatePlan(c NDFCPVlanAssocListValue, cf *bool) int {
	action := ActionNone

	if v.SVlan != c.SVlan {
		log.Printf("SVlan-Update: v.SVlan=%v, c.SVlan=%v", v.SVlan, c.SVlan)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.PVlan != c.PVlan {
		log.Printf("PVlan-Update: v.PVlan=%v, c.PVlan=%v", v.PVlan, c.PVlan)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	return action
}

func (v *NDFCInterfacesValue) CreatePlan(c NDFCInterfacesValue, cf *bool) int {
	action := ActionNone

	if v.SerialNumber != c.SerialNumber {
		log.Printf("SerialNumber-Update: v.SerialNumber=%v, c.SerialNumber=%v", v.SerialNumber, c.SerialNumber)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.InterfaceName != c.InterfaceName {
		log.Printf("InterfaceName-Update: v.InterfaceName=%v, c.InterfaceName=%v", v.InterfaceName, c.InterfaceName)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresReplace
		}
	}

	if v.NvPairs.AdminState != "" {
		if v.NvPairs.AdminState != c.NvPairs.AdminState {
			log.Printf("Update: v.NvPairs.AdminState=%v, c.NvPairs.AdminState=%v", v.NvPairs.AdminState, c.NvPairs.AdminState)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.AdminState=%v, c.NvPairs.AdminState=%v", v.NvPairs.AdminState, c.NvPairs.AdminState)
		v.NvPairs.AdminState = c.NvPairs.AdminState
	}

	if v.NvPairs.FreeformConfig != c.NvPairs.FreeformConfig {
		log.Printf("Update: v.NvPairs.FreeformConfig=%v, c.NvPairs.FreeformConfig=%v", v.NvPairs.FreeformConfig, c.NvPairs.FreeformConfig)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.InterfaceDescription != c.NvPairs.InterfaceDescription {
		log.Printf("Update: v.NvPairs.InterfaceDescription=%v, c.NvPairs.InterfaceDescription=%v", v.NvPairs.InterfaceDescription, c.NvPairs.InterfaceDescription)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.Vrf != c.NvPairs.Vrf {
		log.Printf("Update: v.NvPairs.Vrf=%v, c.NvPairs.Vrf=%v", v.NvPairs.Vrf, c.NvPairs.Vrf)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.Ipv4Address != c.NvPairs.Ipv4Address {
		log.Printf("Update: v.NvPairs.Ipv4Address=%v, c.NvPairs.Ipv4Address=%v", v.NvPairs.Ipv4Address, c.NvPairs.Ipv4Address)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.Ipv6Address != c.NvPairs.Ipv6Address {
		log.Printf("Update: v.NvPairs.Ipv6Address=%v, c.NvPairs.Ipv6Address=%v", v.NvPairs.Ipv6Address, c.NvPairs.Ipv6Address)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.RouteMapTag != c.NvPairs.RouteMapTag {
		log.Printf("Update: v.NvPairs.RouteMapTag=%v, c.NvPairs.RouteMapTag=%v", v.NvPairs.RouteMapTag, c.NvPairs.RouteMapTag)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.BpduGuard != c.NvPairs.BpduGuard {
		log.Printf("Update: v.NvPairs.BpduGuard=%v, c.NvPairs.BpduGuard=%v", v.NvPairs.BpduGuard, c.NvPairs.BpduGuard)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.PortTypeFast != "" {
		if v.NvPairs.PortTypeFast != c.NvPairs.PortTypeFast {
			log.Printf("Update: v.NvPairs.PortTypeFast=%v, c.NvPairs.PortTypeFast=%v", v.NvPairs.PortTypeFast, c.NvPairs.PortTypeFast)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.PortTypeFast=%v, c.NvPairs.PortTypeFast=%v", v.NvPairs.PortTypeFast, c.NvPairs.PortTypeFast)
		v.NvPairs.PortTypeFast = c.NvPairs.PortTypeFast
	}

	if v.NvPairs.Mtu != c.NvPairs.Mtu {
		log.Printf("Update: v.NvPairs.Mtu=%v, c.NvPairs.Mtu=%v", v.NvPairs.Mtu, c.NvPairs.Mtu)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.Speed != c.NvPairs.Speed {
		log.Printf("Update: v.NvPairs.Speed=%v, c.NvPairs.Speed=%v", v.NvPairs.Speed, c.NvPairs.Speed)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
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
	}

	if v.NvPairs.OrphanPort != "" {
		if v.NvPairs.OrphanPort != c.NvPairs.OrphanPort {
			log.Printf("Update: v.NvPairs.OrphanPort=%v, c.NvPairs.OrphanPort=%v", v.NvPairs.OrphanPort, c.NvPairs.OrphanPort)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.OrphanPort=%v, c.NvPairs.OrphanPort=%v", v.NvPairs.OrphanPort, c.NvPairs.OrphanPort)
		v.NvPairs.OrphanPort = c.NvPairs.OrphanPort
	}

	if v.NvPairs.Ptp != "" {
		if v.NvPairs.Ptp != c.NvPairs.Ptp {
			log.Printf("Update: v.NvPairs.Ptp=%v, c.NvPairs.Ptp=%v", v.NvPairs.Ptp, c.NvPairs.Ptp)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.Ptp=%v, c.NvPairs.Ptp=%v", v.NvPairs.Ptp, c.NvPairs.Ptp)
		v.NvPairs.Ptp = c.NvPairs.Ptp
	}

	if v.NvPairs.Netflow != "" {
		if v.NvPairs.Netflow != c.NvPairs.Netflow {
			log.Printf("Update: v.NvPairs.Netflow=%v, c.NvPairs.Netflow=%v", v.NvPairs.Netflow, c.NvPairs.Netflow)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.Netflow=%v, c.NvPairs.Netflow=%v", v.NvPairs.Netflow, c.NvPairs.Netflow)
		v.NvPairs.Netflow = c.NvPairs.Netflow
	}

	if v.NvPairs.NetflowMonitor != c.NvPairs.NetflowMonitor {
		log.Printf("Update: v.NvPairs.NetflowMonitor=%v, c.NvPairs.NetflowMonitor=%v", v.NvPairs.NetflowMonitor, c.NvPairs.NetflowMonitor)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.NetflowSampler != c.NvPairs.NetflowSampler {
		log.Printf("Update: v.NvPairs.NetflowSampler=%v, c.NvPairs.NetflowSampler=%v", v.NvPairs.NetflowSampler, c.NvPairs.NetflowSampler)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.AllowedVlans != c.NvPairs.AllowedVlans {
		log.Printf("Update: v.NvPairs.AllowedVlans=%v, c.NvPairs.AllowedVlans=%v", v.NvPairs.AllowedVlans, c.NvPairs.AllowedVlans)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
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
	}

	if v.NvPairs.Ipv4PrefixLength != c.NvPairs.Ipv4PrefixLength {
		log.Printf("Update: v.NvPairs.Ipv4PrefixLength=%v, c.NvPairs.Ipv4PrefixLength=%v", v.NvPairs.Ipv4PrefixLength, c.NvPairs.Ipv4PrefixLength)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.RoutingTag != c.NvPairs.RoutingTag {
		log.Printf("Update: v.NvPairs.RoutingTag=%v, c.NvPairs.RoutingTag=%v", v.NvPairs.RoutingTag, c.NvPairs.RoutingTag)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.DisableIpRedirects != "" {
		if v.NvPairs.DisableIpRedirects != c.NvPairs.DisableIpRedirects {
			log.Printf("Update: v.NvPairs.DisableIpRedirects=%v, c.NvPairs.DisableIpRedirects=%v", v.NvPairs.DisableIpRedirects, c.NvPairs.DisableIpRedirects)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.DisableIpRedirects=%v, c.NvPairs.DisableIpRedirects=%v", v.NvPairs.DisableIpRedirects, c.NvPairs.DisableIpRedirects)
		v.NvPairs.DisableIpRedirects = c.NvPairs.DisableIpRedirects
	}

	if v.NvPairs.EnableHsrp != "" {
		if v.NvPairs.EnableHsrp != c.NvPairs.EnableHsrp {
			log.Printf("Update: v.NvPairs.EnableHsrp=%v, c.NvPairs.EnableHsrp=%v", v.NvPairs.EnableHsrp, c.NvPairs.EnableHsrp)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.EnableHsrp=%v, c.NvPairs.EnableHsrp=%v", v.NvPairs.EnableHsrp, c.NvPairs.EnableHsrp)
		v.NvPairs.EnableHsrp = c.NvPairs.EnableHsrp
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
	}

	if v.NvPairs.HsrpVip != c.NvPairs.HsrpVip {
		log.Printf("Update: v.NvPairs.HsrpVip=%v, c.NvPairs.HsrpVip=%v", v.NvPairs.HsrpVip, c.NvPairs.HsrpVip)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
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
	}

	if v.NvPairs.HsrpVersion != c.NvPairs.HsrpVersion {
		log.Printf("Update: v.NvPairs.HsrpVersion=%v, c.NvPairs.HsrpVersion=%v", v.NvPairs.HsrpVersion, c.NvPairs.HsrpVersion)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.Preempt != "" {
		if v.NvPairs.Preempt != c.NvPairs.Preempt {
			log.Printf("Update: v.NvPairs.Preempt=%v, c.NvPairs.Preempt=%v", v.NvPairs.Preempt, c.NvPairs.Preempt)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.Preempt=%v, c.NvPairs.Preempt=%v", v.NvPairs.Preempt, c.NvPairs.Preempt)
		v.NvPairs.Preempt = c.NvPairs.Preempt
	}

	if v.NvPairs.Mac != c.NvPairs.Mac {
		log.Printf("Update: v.NvPairs.Mac=%v, c.NvPairs.Mac=%v", v.NvPairs.Mac, c.NvPairs.Mac)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.DhcpServerAddr1 != c.NvPairs.DhcpServerAddr1 {
		log.Printf("Update: v.NvPairs.DhcpServerAddr1=%v, c.NvPairs.DhcpServerAddr1=%v", v.NvPairs.DhcpServerAddr1, c.NvPairs.DhcpServerAddr1)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.DhcpServerAddr2 != c.NvPairs.DhcpServerAddr2 {
		log.Printf("Update: v.NvPairs.DhcpServerAddr2=%v, c.NvPairs.DhcpServerAddr2=%v", v.NvPairs.DhcpServerAddr2, c.NvPairs.DhcpServerAddr2)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.DhcpServerAddr3 != c.NvPairs.DhcpServerAddr3 {
		log.Printf("Update: v.NvPairs.DhcpServerAddr3=%v, c.NvPairs.DhcpServerAddr3=%v", v.NvPairs.DhcpServerAddr3, c.NvPairs.DhcpServerAddr3)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.VrfDhcp1 != c.NvPairs.VrfDhcp1 {
		log.Printf("Update: v.NvPairs.VrfDhcp1=%v, c.NvPairs.VrfDhcp1=%v", v.NvPairs.VrfDhcp1, c.NvPairs.VrfDhcp1)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.VrfDhcp2 != c.NvPairs.VrfDhcp2 {
		log.Printf("Update: v.NvPairs.VrfDhcp2=%v, c.NvPairs.VrfDhcp2=%v", v.NvPairs.VrfDhcp2, c.NvPairs.VrfDhcp2)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.VrfDhcp3 != c.NvPairs.VrfDhcp3 {
		log.Printf("Update: v.NvPairs.VrfDhcp3=%v, c.NvPairs.VrfDhcp3=%v", v.NvPairs.VrfDhcp3, c.NvPairs.VrfDhcp3)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.AdvertiseSubnetInUnderlay != "" {
		if v.NvPairs.AdvertiseSubnetInUnderlay != c.NvPairs.AdvertiseSubnetInUnderlay {
			log.Printf("Update: v.NvPairs.AdvertiseSubnetInUnderlay=%v, c.NvPairs.AdvertiseSubnetInUnderlay=%v", v.NvPairs.AdvertiseSubnetInUnderlay, c.NvPairs.AdvertiseSubnetInUnderlay)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.AdvertiseSubnetInUnderlay=%v, c.NvPairs.AdvertiseSubnetInUnderlay=%v", v.NvPairs.AdvertiseSubnetInUnderlay, c.NvPairs.AdvertiseSubnetInUnderlay)
		v.NvPairs.AdvertiseSubnetInUnderlay = c.NvPairs.AdvertiseSubnetInUnderlay
	}

	if v.NvPairs.CopyPoDescription != "" {
		if v.NvPairs.CopyPoDescription != c.NvPairs.CopyPoDescription {
			log.Printf("Update: v.NvPairs.CopyPoDescription=%v, c.NvPairs.CopyPoDescription=%v", v.NvPairs.CopyPoDescription, c.NvPairs.CopyPoDescription)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.CopyPoDescription=%v, c.NvPairs.CopyPoDescription=%v", v.NvPairs.CopyPoDescription, c.NvPairs.CopyPoDescription)
		v.NvPairs.CopyPoDescription = c.NvPairs.CopyPoDescription
	}

	if v.NvPairs.PortchannelMode != c.NvPairs.PortchannelMode {
		log.Printf("Update: v.NvPairs.PortchannelMode=%v, c.NvPairs.PortchannelMode=%v", v.NvPairs.PortchannelMode, c.NvPairs.PortchannelMode)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.MemberInterfaces != c.NvPairs.MemberInterfaces {
		log.Printf("Update: v.NvPairs.MemberInterfaces=%v, c.NvPairs.MemberInterfaces=%v", v.NvPairs.MemberInterfaces, c.NvPairs.MemberInterfaces)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.Peer1PoFreeformConfig != c.NvPairs.Peer1PoFreeformConfig {
		log.Printf("Update: v.NvPairs.Peer1PoFreeformConfig=%v, c.NvPairs.Peer1PoFreeformConfig=%v", v.NvPairs.Peer1PoFreeformConfig, c.NvPairs.Peer1PoFreeformConfig)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.Peer2PoFreeformConfig != c.NvPairs.Peer2PoFreeformConfig {
		log.Printf("Update: v.NvPairs.Peer2PoFreeformConfig=%v, c.NvPairs.Peer2PoFreeformConfig=%v", v.NvPairs.Peer2PoFreeformConfig, c.NvPairs.Peer2PoFreeformConfig)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.Peer1PoDescription != c.NvPairs.Peer1PoDescription {
		log.Printf("Update: v.NvPairs.Peer1PoDescription=%v, c.NvPairs.Peer1PoDescription=%v", v.NvPairs.Peer1PoDescription, c.NvPairs.Peer1PoDescription)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.Peer2PoDescription != c.NvPairs.Peer2PoDescription {
		log.Printf("Update: v.NvPairs.Peer2PoDescription=%v, c.NvPairs.Peer2PoDescription=%v", v.NvPairs.Peer2PoDescription, c.NvPairs.Peer2PoDescription)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.Peer1AllowedVlans != c.NvPairs.Peer1AllowedVlans {
		log.Printf("Update: v.NvPairs.Peer1AllowedVlans=%v, c.NvPairs.Peer1AllowedVlans=%v", v.NvPairs.Peer1AllowedVlans, c.NvPairs.Peer1AllowedVlans)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.Peer2AllowedVlans != c.NvPairs.Peer2AllowedVlans {
		log.Printf("Update: v.NvPairs.Peer2AllowedVlans=%v, c.NvPairs.Peer2AllowedVlans=%v", v.NvPairs.Peer2AllowedVlans, c.NvPairs.Peer2AllowedVlans)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if !v.NvPairs.Peer1NativeVlan.IsEmpty() && !c.NvPairs.Peer1NativeVlan.IsEmpty() {
		if *v.NvPairs.Peer1NativeVlan != *c.NvPairs.Peer1NativeVlan {
			log.Printf("Update: v.NvPairs.Peer1NativeVlan=%v, c.NvPairs.Peer1NativeVlan=%v", *v.NvPairs.Peer1NativeVlan, *c.NvPairs.Peer1NativeVlan)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else if !v.NvPairs.Peer1NativeVlan.IsEmpty() {
		log.Printf("Update: v.NvPairs.Peer1NativeVlan=%v", *v.NvPairs.Peer1NativeVlan)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if !c.NvPairs.Peer1NativeVlan.IsEmpty() {
		log.Printf("Copy from State: c.NvPairs.Peer1NativeVlan=%v", *c.NvPairs.Peer1NativeVlan)
		v.NvPairs.Peer1NativeVlan = new(Int64Custom)
		*v.NvPairs.Peer1NativeVlan = *c.NvPairs.Peer1NativeVlan
	}

	if !v.NvPairs.Peer2NativeVlan.IsEmpty() && !c.NvPairs.Peer2NativeVlan.IsEmpty() {
		if *v.NvPairs.Peer2NativeVlan != *c.NvPairs.Peer2NativeVlan {
			log.Printf("Update: v.NvPairs.Peer2NativeVlan=%v, c.NvPairs.Peer2NativeVlan=%v", *v.NvPairs.Peer2NativeVlan, *c.NvPairs.Peer2NativeVlan)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else if !v.NvPairs.Peer2NativeVlan.IsEmpty() {
		log.Printf("Update: v.NvPairs.Peer2NativeVlan=%v", *v.NvPairs.Peer2NativeVlan)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if !c.NvPairs.Peer2NativeVlan.IsEmpty() {
		log.Printf("Copy from State: c.NvPairs.Peer2NativeVlan=%v", *c.NvPairs.Peer2NativeVlan)
		v.NvPairs.Peer2NativeVlan = new(Int64Custom)
		*v.NvPairs.Peer2NativeVlan = *c.NvPairs.Peer2NativeVlan
	}

	if v.NvPairs.Peer1MemberInterfaces != c.NvPairs.Peer1MemberInterfaces {
		log.Printf("Update: v.NvPairs.Peer1MemberInterfaces=%v, c.NvPairs.Peer1MemberInterfaces=%v", v.NvPairs.Peer1MemberInterfaces, c.NvPairs.Peer1MemberInterfaces)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.Peer2MemberInterfaces != c.NvPairs.Peer2MemberInterfaces {
		log.Printf("Update: v.NvPairs.Peer2MemberInterfaces=%v, c.NvPairs.Peer2MemberInterfaces=%v", v.NvPairs.Peer2MemberInterfaces, c.NvPairs.Peer2MemberInterfaces)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.Peer1PortChannelId != nil && c.NvPairs.Peer1PortChannelId != nil {
		if *v.NvPairs.Peer1PortChannelId != *c.NvPairs.Peer1PortChannelId {
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresReplace
			}
			log.Printf("Update: v.NvPairs.Peer1PortChannelId=%v, c.NvPairs.Peer1PortChannelId=%v", *v.NvPairs.Peer1PortChannelId, *c.NvPairs.Peer1PortChannelId)
		}
	} else if v.NvPairs.Peer1PortChannelId != nil {
		log.Printf("Update: v.NvPairs.Peer1PortChannelId=%v, c.NvPairs.Peer1PortChannelId=nil", *v.NvPairs.Peer1PortChannelId)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresReplace
		}
	} else if c.NvPairs.Peer1PortChannelId != nil {
		v.NvPairs.Peer1PortChannelId = new(int64)
		log.Printf("Copy from state: v.NvPairs.Peer1PortChannelId=nil, c.NvPairs.Peer1PortChannelId=%v", *c.NvPairs.Peer1PortChannelId)
		*v.NvPairs.Peer1PortChannelId = *c.NvPairs.Peer1PortChannelId
	}

	if !v.NvPairs.Peer2PortChannelId.IsEmpty() && !c.NvPairs.Peer2PortChannelId.IsEmpty() {
		if *v.NvPairs.Peer2PortChannelId != *c.NvPairs.Peer2PortChannelId {
			log.Printf("Update: v.NvPairs.Peer2PortChannelId=%v, c.NvPairs.Peer2PortChannelId=%v", *v.NvPairs.Peer2PortChannelId, *c.NvPairs.Peer2PortChannelId)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresReplace
			}
		}
	} else if !v.NvPairs.Peer2PortChannelId.IsEmpty() {
		log.Printf("Update: v.NvPairs.Peer2PortChannelId=%v", *v.NvPairs.Peer2PortChannelId)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresReplace
		}
	} else if !c.NvPairs.Peer2PortChannelId.IsEmpty() {
		log.Printf("Copy from State: c.NvPairs.Peer2PortChannelId=%v", *c.NvPairs.Peer2PortChannelId)
		v.NvPairs.Peer2PortChannelId = new(Int64Custom)
		*v.NvPairs.Peer2PortChannelId = *c.NvPairs.Peer2PortChannelId
	}

	if v.NvPairs.EnablePimSparse != "" {
		if v.NvPairs.EnablePimSparse != c.NvPairs.EnablePimSparse {
			log.Printf("Update: v.NvPairs.EnablePimSparse=%v, c.NvPairs.EnablePimSparse=%v", v.NvPairs.EnablePimSparse, c.NvPairs.EnablePimSparse)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.EnablePimSparse=%v, c.NvPairs.EnablePimSparse=%v", v.NvPairs.EnablePimSparse, c.NvPairs.EnablePimSparse)
		v.NvPairs.EnablePimSparse = c.NvPairs.EnablePimSparse
	}

	if v.NvPairs.PimDrPriority != c.NvPairs.PimDrPriority {
		log.Printf("Update: v.NvPairs.PimDrPriority=%v, c.NvPairs.PimDrPriority=%v", v.NvPairs.PimDrPriority, c.NvPairs.PimDrPriority)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.EnablePfc != "" {
		if v.NvPairs.EnablePfc != c.NvPairs.EnablePfc {
			log.Printf("Update: v.NvPairs.EnablePfc=%v, c.NvPairs.EnablePfc=%v", v.NvPairs.EnablePfc, c.NvPairs.EnablePfc)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.EnablePfc=%v, c.NvPairs.EnablePfc=%v", v.NvPairs.EnablePfc, c.NvPairs.EnablePfc)
		v.NvPairs.EnablePfc = c.NvPairs.EnablePfc
	}

	if v.NvPairs.EnableQos != "" {
		if v.NvPairs.EnableQos != c.NvPairs.EnableQos {
			log.Printf("Update: v.NvPairs.EnableQos=%v, c.NvPairs.EnableQos=%v", v.NvPairs.EnableQos, c.NvPairs.EnableQos)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.EnableQos=%v, c.NvPairs.EnableQos=%v", v.NvPairs.EnableQos, c.NvPairs.EnableQos)
		v.NvPairs.EnableQos = c.NvPairs.EnableQos
	}

	if v.NvPairs.QosPolicy != c.NvPairs.QosPolicy {
		log.Printf("Update: v.NvPairs.QosPolicy=%v, c.NvPairs.QosPolicy=%v", v.NvPairs.QosPolicy, c.NvPairs.QosPolicy)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.QueuingPolicy != c.NvPairs.QueuingPolicy {
		log.Printf("Update: v.NvPairs.QueuingPolicy=%v, c.NvPairs.QueuingPolicy=%v", v.NvPairs.QueuingPolicy, c.NvPairs.QueuingPolicy)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.LinkStateRoutingProtocol != c.NvPairs.LinkStateRoutingProtocol {
		log.Printf("Update: v.NvPairs.LinkStateRoutingProtocol=%v, c.NvPairs.LinkStateRoutingProtocol=%v", v.NvPairs.LinkStateRoutingProtocol, c.NvPairs.LinkStateRoutingProtocol)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.LinkStateRoutingTag != c.NvPairs.LinkStateRoutingTag {
		log.Printf("Update: v.NvPairs.LinkStateRoutingTag=%v, c.NvPairs.LinkStateRoutingTag=%v", v.NvPairs.LinkStateRoutingTag, c.NvPairs.LinkStateRoutingTag)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.Ipv6Addr != c.NvPairs.Ipv6Addr {
		log.Printf("Update: v.NvPairs.Ipv6Addr=%v, c.NvPairs.Ipv6Addr=%v", v.NvPairs.Ipv6Addr, c.NvPairs.Ipv6Addr)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.Ipv6PrefixLength != c.NvPairs.Ipv6PrefixLength {
		log.Printf("Update: v.NvPairs.Ipv6PrefixLength=%v, c.NvPairs.Ipv6PrefixLength=%v", v.NvPairs.Ipv6PrefixLength, c.NvPairs.Ipv6PrefixLength)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.CdpEnable != "" {
		if v.NvPairs.CdpEnable != c.NvPairs.CdpEnable {
			log.Printf("Update: v.NvPairs.CdpEnable=%v, c.NvPairs.CdpEnable=%v", v.NvPairs.CdpEnable, c.NvPairs.CdpEnable)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.CdpEnable=%v, c.NvPairs.CdpEnable=%v", v.NvPairs.CdpEnable, c.NvPairs.CdpEnable)
		v.NvPairs.CdpEnable = c.NvPairs.CdpEnable
	}

	if v.NvPairs.PortDuplexMode != c.NvPairs.PortDuplexMode {
		log.Printf("Update: v.NvPairs.PortDuplexMode=%v, c.NvPairs.PortDuplexMode=%v", v.NvPairs.PortDuplexMode, c.NvPairs.PortDuplexMode)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.EnableMonitor != "" {
		if v.NvPairs.EnableMonitor != c.NvPairs.EnableMonitor {
			log.Printf("Update: v.NvPairs.EnableMonitor=%v, c.NvPairs.EnableMonitor=%v", v.NvPairs.EnableMonitor, c.NvPairs.EnableMonitor)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.EnableMonitor=%v, c.NvPairs.EnableMonitor=%v", v.NvPairs.EnableMonitor, c.NvPairs.EnableMonitor)
		v.NvPairs.EnableMonitor = c.NvPairs.EnableMonitor
	}

	if v.NvPairs.PvlanMode != c.NvPairs.PvlanMode {
		log.Printf("Update: v.NvPairs.PvlanMode=%v, c.NvPairs.PvlanMode=%v", v.NvPairs.PvlanMode, c.NvPairs.PvlanMode)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.PvlanAllowedVlans != c.NvPairs.PvlanAllowedVlans {
		log.Printf("Update: v.NvPairs.PvlanAllowedVlans=%v, c.NvPairs.PvlanAllowedVlans=%v", v.NvPairs.PvlanAllowedVlans, c.NvPairs.PvlanAllowedVlans)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.PvlanNativeVlan != c.NvPairs.PvlanNativeVlan {
		log.Printf("Update: v.NvPairs.PvlanNativeVlan=%v, c.NvPairs.PvlanNativeVlan=%v", v.NvPairs.PvlanNativeVlan, c.NvPairs.PvlanNativeVlan)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.IgForFex != "" {
		if v.NvPairs.IgForFex != c.NvPairs.IgForFex {
			log.Printf("Update: v.NvPairs.IgForFex=%v, c.NvPairs.IgForFex=%v", v.NvPairs.IgForFex, c.NvPairs.IgForFex)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.IgForFex=%v, c.NvPairs.IgForFex=%v", v.NvPairs.IgForFex, c.NvPairs.IgForFex)
		v.NvPairs.IgForFex = c.NvPairs.IgForFex
	}

	if v.NvPairs.AutoNegotiate != c.NvPairs.AutoNegotiate {
		log.Printf("Update: v.NvPairs.AutoNegotiate=%v, c.NvPairs.AutoNegotiate=%v", v.NvPairs.AutoNegotiate, c.NvPairs.AutoNegotiate)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.PathCost != nil && c.NvPairs.PathCost != nil {
		if *v.NvPairs.PathCost != *c.NvPairs.PathCost {
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
			log.Printf("Update: v.NvPairs.PathCost=%v, c.NvPairs.PathCost=%v", *v.NvPairs.PathCost, *c.NvPairs.PathCost)
		}
	} else if v.NvPairs.PathCost != nil {
		log.Printf("Update: v.NvPairs.PathCost=%v, c.NvPairs.PathCost=nil", *v.NvPairs.PathCost)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if c.NvPairs.PathCost != nil {
		v.NvPairs.PathCost = new(int64)
		log.Printf("Copy from state: v.NvPairs.PathCost=nil, c.NvPairs.PathCost=%v", *c.NvPairs.PathCost)
		*v.NvPairs.PathCost = *c.NvPairs.PathCost
	}

	if v.NvPairs.GuardMode != c.NvPairs.GuardMode {
		log.Printf("Update: v.NvPairs.GuardMode=%v, c.NvPairs.GuardMode=%v", v.NvPairs.GuardMode, c.NvPairs.GuardMode)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NvPairs.Ttag != "" {
		if v.NvPairs.Ttag != c.NvPairs.Ttag {
			log.Printf("Update: v.NvPairs.Ttag=%v, c.NvPairs.Ttag=%v", v.NvPairs.Ttag, c.NvPairs.Ttag)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NvPairs.Ttag=%v, c.NvPairs.Ttag=%v", v.NvPairs.Ttag, c.NvPairs.Ttag)
		v.NvPairs.Ttag = c.NvPairs.Ttag
	}

	if len(v.NvPairs.PVlanMappingList) != len(c.NvPairs.PVlanMappingList) {
		log.Printf("PVlanMappingList-Update: len(v.NvPairs.PVlanMappingList)=%d, len(c.NvPairs.PVlanMappingList)=%d", len(v.NvPairs.PVlanMappingList), len(c.NvPairs.PVlanMappingList))
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}
	for i := range v.NvPairs.PVlanMappingList {
		retVal := v.NvPairs.PVlanMappingList[i].CreatePlan(c.NvPairs.PVlanMappingList[i], cf)
		if retVal != ActionNone {
			if action == ActionNone || action == RequiresUpdate {
				action = retVal
			}
		}
	}

	if len(v.NvPairs.PVlanAssocList) != len(c.NvPairs.PVlanAssocList) {
		log.Printf("PVlanAssocList-Update: len(v.NvPairs.PVlanAssocList)=%d, len(c.NvPairs.PVlanAssocList)=%d", len(v.NvPairs.PVlanAssocList), len(c.NvPairs.PVlanAssocList))
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}
	for i := range v.NvPairs.PVlanAssocList {
		retVal := v.NvPairs.PVlanAssocList[i].CreatePlan(c.NvPairs.PVlanAssocList[i], cf)
		if retVal != ActionNone {
			if action == ActionNone || action == RequiresUpdate {
				action = retVal
			}
		}
	}
	if len(v.CustomPolicyParameters) != len(c.CustomPolicyParameters) {
		log.Printf("Update: len(v.CustomPolicyParameters)=%d, len(c.CustomPolicyParameters)=%d", len(v.CustomPolicyParameters), len(c.CustomPolicyParameters))
		return RequiresUpdate
	}
	for kk, vv := range v.CustomPolicyParameters {
		cc, ok := c.CustomPolicyParameters[kk]
		if !ok {
			log.Printf("Update: v.CustomPolicyParameters[%s]=%s, c.CustomPolicyParameters[%s]=nil", kk, vv, kk)
			return RequiresUpdate
		}
		if vv != cc {
			log.Printf("Update: v.CustomPolicyParameters[%s]=%s, c.CustomPolicyParameters[%s]=%s", kk, vv, kk, cc)
			return RequiresUpdate
		}
	}

	return action
}
