// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package resource_networks

import (
	"log"
	. "terraform-provider-ndfc/internal/provider/types"
)

func (v NDFCDhcpRelayServersValue) DeepEqual(c NDFCDhcpRelayServersValue) int {
	cf := false
	if v.Address != c.Address {
		log.Printf("v.Address=%v, c.Address=%v", v.Address, c.Address)
		return RequiresUpdate
	}
	if v.Vrf != c.Vrf {
		log.Printf("v.Vrf=%v, c.Vrf=%v", v.Vrf, c.Vrf)
		return RequiresUpdate
	}

	if cf {
		return ControlFlagUpdate
	}
	return ValuesDeeplyEqual
}

func (v NDFCNetworksValue) DeepEqual(c NDFCNetworksValue) int {
	cf := false
	if v.DisplayName != c.DisplayName {
		log.Printf("v.DisplayName=%v, c.DisplayName=%v", v.DisplayName, c.DisplayName)
		return RequiresUpdate
	}

	if v.NetworkId != nil && c.NetworkId != nil {
		if *v.NetworkId != *c.NetworkId {
			log.Printf("v.NetworkId=%v, c.NetworkId=%v", *v.NetworkId, *c.NetworkId)
			return RequiresReplace
		}
	} else {
		if v.NetworkId != nil {
			log.Printf("v.NetworkId=%v", *v.NetworkId)
			return RequiresReplace
		} else if c.NetworkId != nil {
			log.Printf("c.NetworkId=%v", *c.NetworkId)
			return RequiresReplace
		}
	}

	if v.NetworkTemplate != c.NetworkTemplate {
		log.Printf("v.NetworkTemplate=%v, c.NetworkTemplate=%v", v.NetworkTemplate, c.NetworkTemplate)
		return RequiresUpdate
	}
	if v.NetworkExtensionTemplate != c.NetworkExtensionTemplate {
		log.Printf("v.NetworkExtensionTemplate=%v, c.NetworkExtensionTemplate=%v", v.NetworkExtensionTemplate, c.NetworkExtensionTemplate)
		return RequiresUpdate
	}
	if v.VrfName != c.VrfName {
		log.Printf("v.VrfName=%v, c.VrfName=%v", v.VrfName, c.VrfName)
		return RequiresUpdate
	}

	if !v.PrimaryNetworkId.IsEmpty() && !c.PrimaryNetworkId.IsEmpty() {
		if *v.PrimaryNetworkId != *c.PrimaryNetworkId {
			log.Printf("v.PrimaryNetworkId=%v, c.PrimaryNetworkId=%v", *v.PrimaryNetworkId, *c.PrimaryNetworkId)
			return RequiresUpdate
		}
	} else {
		if !v.PrimaryNetworkId.IsEmpty() {
			log.Printf("v.PrimaryNetworkId=%v", *v.PrimaryNetworkId)
			return RequiresUpdate
		} else if !c.PrimaryNetworkId.IsEmpty() {
			log.Printf("c.PrimaryNetworkId=%v", *c.PrimaryNetworkId)
			return RequiresUpdate
		}
	}
	if v.NetworkType != c.NetworkType {
		log.Printf("v.NetworkType=%v, c.NetworkType=%v", v.NetworkType, c.NetworkType)
		return RequiresUpdate
	}
	if v.NetworkTemplateConfig.GatewayIpv4Address != c.NetworkTemplateConfig.GatewayIpv4Address {
		log.Printf("v.NetworkTemplateConfig.GatewayIpv4Address=%s, c.NetworkTemplateConfig.GatewayIpv4Address=%s", v.NetworkTemplateConfig.GatewayIpv4Address, c.NetworkTemplateConfig.GatewayIpv4Address)
		return RequiresUpdate
	}
	if v.NetworkTemplateConfig.GatewayIpv6Address != c.NetworkTemplateConfig.GatewayIpv6Address {
		log.Printf("v.NetworkTemplateConfig.GatewayIpv6Address=%s, c.NetworkTemplateConfig.GatewayIpv6Address=%s", v.NetworkTemplateConfig.GatewayIpv6Address, c.NetworkTemplateConfig.GatewayIpv6Address)
		return RequiresUpdate
	}

	if !v.NetworkTemplateConfig.VlanId.IsEmpty() && !c.NetworkTemplateConfig.VlanId.IsEmpty() {
		if *v.NetworkTemplateConfig.VlanId != *c.NetworkTemplateConfig.VlanId {
			log.Printf("v.NetworkTemplateConfig.VlanId=%v, c.NetworkTemplateConfig.VlanId=%v", *v.NetworkTemplateConfig.VlanId, *c.NetworkTemplateConfig.VlanId)
			return RequiresUpdate
		}
	} else {
		if !v.NetworkTemplateConfig.VlanId.IsEmpty() {
			log.Printf("v.NetworkTemplateConfig.VlanId=%v", *v.NetworkTemplateConfig.VlanId)
			return RequiresUpdate
		} else if !c.NetworkTemplateConfig.VlanId.IsEmpty() {
			log.Printf("c.NetworkTemplateConfig.VlanId=%v", *c.NetworkTemplateConfig.VlanId)
			return RequiresUpdate
		}
	}
	if v.NetworkTemplateConfig.VlanName != c.NetworkTemplateConfig.VlanName {
		log.Printf("v.NetworkTemplateConfig.VlanName=%s, c.NetworkTemplateConfig.VlanName=%s", v.NetworkTemplateConfig.VlanName, c.NetworkTemplateConfig.VlanName)
		return RequiresUpdate
	}
	if v.NetworkTemplateConfig.Layer2Only != c.NetworkTemplateConfig.Layer2Only {
		log.Printf("v.NetworkTemplateConfig.Layer2Only=%s, c.NetworkTemplateConfig.Layer2Only=%s", v.NetworkTemplateConfig.Layer2Only, c.NetworkTemplateConfig.Layer2Only)
		return RequiresUpdate
	}
	if v.NetworkTemplateConfig.InterfaceDescription != c.NetworkTemplateConfig.InterfaceDescription {
		log.Printf("v.NetworkTemplateConfig.InterfaceDescription=%s, c.NetworkTemplateConfig.InterfaceDescription=%s", v.NetworkTemplateConfig.InterfaceDescription, c.NetworkTemplateConfig.InterfaceDescription)
		return RequiresUpdate
	}

	if !v.NetworkTemplateConfig.Mtu.IsEmpty() && !c.NetworkTemplateConfig.Mtu.IsEmpty() {
		if *v.NetworkTemplateConfig.Mtu != *c.NetworkTemplateConfig.Mtu {
			log.Printf("v.NetworkTemplateConfig.Mtu=%v, c.NetworkTemplateConfig.Mtu=%v", *v.NetworkTemplateConfig.Mtu, *c.NetworkTemplateConfig.Mtu)
			return RequiresUpdate
		}
	} else {
		if !v.NetworkTemplateConfig.Mtu.IsEmpty() {
			log.Printf("v.NetworkTemplateConfig.Mtu=%v", *v.NetworkTemplateConfig.Mtu)
			return RequiresUpdate
		} else if !c.NetworkTemplateConfig.Mtu.IsEmpty() {
			log.Printf("c.NetworkTemplateConfig.Mtu=%v", *c.NetworkTemplateConfig.Mtu)
			return RequiresUpdate
		}
	}
	if v.NetworkTemplateConfig.SecondaryGateway1 != c.NetworkTemplateConfig.SecondaryGateway1 {
		log.Printf("v.NetworkTemplateConfig.SecondaryGateway1=%s, c.NetworkTemplateConfig.SecondaryGateway1=%s", v.NetworkTemplateConfig.SecondaryGateway1, c.NetworkTemplateConfig.SecondaryGateway1)
		return RequiresUpdate
	}
	if v.NetworkTemplateConfig.SecondaryGateway2 != c.NetworkTemplateConfig.SecondaryGateway2 {
		log.Printf("v.NetworkTemplateConfig.SecondaryGateway2=%s, c.NetworkTemplateConfig.SecondaryGateway2=%s", v.NetworkTemplateConfig.SecondaryGateway2, c.NetworkTemplateConfig.SecondaryGateway2)
		return RequiresUpdate
	}
	if v.NetworkTemplateConfig.SecondaryGateway3 != c.NetworkTemplateConfig.SecondaryGateway3 {
		log.Printf("v.NetworkTemplateConfig.SecondaryGateway3=%s, c.NetworkTemplateConfig.SecondaryGateway3=%s", v.NetworkTemplateConfig.SecondaryGateway3, c.NetworkTemplateConfig.SecondaryGateway3)
		return RequiresUpdate
	}
	if v.NetworkTemplateConfig.SecondaryGateway4 != c.NetworkTemplateConfig.SecondaryGateway4 {
		log.Printf("v.NetworkTemplateConfig.SecondaryGateway4=%s, c.NetworkTemplateConfig.SecondaryGateway4=%s", v.NetworkTemplateConfig.SecondaryGateway4, c.NetworkTemplateConfig.SecondaryGateway4)
		return RequiresUpdate
	}
	if v.NetworkTemplateConfig.ArpSuppression != c.NetworkTemplateConfig.ArpSuppression {
		log.Printf("v.NetworkTemplateConfig.ArpSuppression=%s, c.NetworkTemplateConfig.ArpSuppression=%s", v.NetworkTemplateConfig.ArpSuppression, c.NetworkTemplateConfig.ArpSuppression)
		return RequiresUpdate
	}
	if v.NetworkTemplateConfig.IngressReplication != c.NetworkTemplateConfig.IngressReplication {
		log.Printf("v.NetworkTemplateConfig.IngressReplication=%s, c.NetworkTemplateConfig.IngressReplication=%s", v.NetworkTemplateConfig.IngressReplication, c.NetworkTemplateConfig.IngressReplication)
		return RequiresUpdate
	}
	if v.NetworkTemplateConfig.MulticastGroup != c.NetworkTemplateConfig.MulticastGroup {
		log.Printf("v.NetworkTemplateConfig.MulticastGroup=%s, c.NetworkTemplateConfig.MulticastGroup=%s", v.NetworkTemplateConfig.MulticastGroup, c.NetworkTemplateConfig.MulticastGroup)
		return RequiresUpdate
	}

	if len(v.NetworkTemplateConfig.DhcpRelayServers) != len(c.NetworkTemplateConfig.DhcpRelayServers) {
		log.Printf("len(v.NetworkTemplateConfig.DhcpRelayServers)=%d, len(c.NetworkTemplateConfig.DhcpRelayServers)=%d", len(v.NetworkTemplateConfig.DhcpRelayServers), len(c.NetworkTemplateConfig.DhcpRelayServers))
		return RequiresUpdate
	}
	for i := range v.NetworkTemplateConfig.DhcpRelayServers {
		retVal := v.NetworkTemplateConfig.DhcpRelayServers[i].DeepEqual(c.NetworkTemplateConfig.DhcpRelayServers[i])
		if retVal != ValuesDeeplyEqual {
			return retVal
		}
	}
	if !v.NetworkTemplateConfig.DhcpRelayLoopbackId.IsEmpty() && !c.NetworkTemplateConfig.DhcpRelayLoopbackId.IsEmpty() {
		if *v.NetworkTemplateConfig.DhcpRelayLoopbackId != *c.NetworkTemplateConfig.DhcpRelayLoopbackId {
			log.Printf("v.NetworkTemplateConfig.DhcpRelayLoopbackId=%v, c.NetworkTemplateConfig.DhcpRelayLoopbackId=%v", *v.NetworkTemplateConfig.DhcpRelayLoopbackId, *c.NetworkTemplateConfig.DhcpRelayLoopbackId)
			return RequiresUpdate
		}
	} else {
		if !v.NetworkTemplateConfig.DhcpRelayLoopbackId.IsEmpty() {
			log.Printf("v.NetworkTemplateConfig.DhcpRelayLoopbackId=%v", *v.NetworkTemplateConfig.DhcpRelayLoopbackId)
			return RequiresUpdate
		} else if !c.NetworkTemplateConfig.DhcpRelayLoopbackId.IsEmpty() {
			log.Printf("c.NetworkTemplateConfig.DhcpRelayLoopbackId=%v", *c.NetworkTemplateConfig.DhcpRelayLoopbackId)
			return RequiresUpdate
		}
	}

	if !v.NetworkTemplateConfig.RoutingTag.IsEmpty() && !c.NetworkTemplateConfig.RoutingTag.IsEmpty() {
		if *v.NetworkTemplateConfig.RoutingTag != *c.NetworkTemplateConfig.RoutingTag {
			log.Printf("v.NetworkTemplateConfig.RoutingTag=%v, c.NetworkTemplateConfig.RoutingTag=%v", *v.NetworkTemplateConfig.RoutingTag, *c.NetworkTemplateConfig.RoutingTag)
			return RequiresUpdate
		}
	} else {
		if !v.NetworkTemplateConfig.RoutingTag.IsEmpty() {
			log.Printf("v.NetworkTemplateConfig.RoutingTag=%v", *v.NetworkTemplateConfig.RoutingTag)
			return RequiresUpdate
		} else if !c.NetworkTemplateConfig.RoutingTag.IsEmpty() {
			log.Printf("c.NetworkTemplateConfig.RoutingTag=%v", *c.NetworkTemplateConfig.RoutingTag)
			return RequiresUpdate
		}
	}
	if v.NetworkTemplateConfig.Trm != c.NetworkTemplateConfig.Trm {
		log.Printf("v.NetworkTemplateConfig.Trm=%s, c.NetworkTemplateConfig.Trm=%s", v.NetworkTemplateConfig.Trm, c.NetworkTemplateConfig.Trm)
		return RequiresUpdate
	}
	if v.NetworkTemplateConfig.RouteTargetBoth != c.NetworkTemplateConfig.RouteTargetBoth {
		log.Printf("v.NetworkTemplateConfig.RouteTargetBoth=%s, c.NetworkTemplateConfig.RouteTargetBoth=%s", v.NetworkTemplateConfig.RouteTargetBoth, c.NetworkTemplateConfig.RouteTargetBoth)
		return RequiresUpdate
	}
	if v.NetworkTemplateConfig.Netflow != c.NetworkTemplateConfig.Netflow {
		log.Printf("v.NetworkTemplateConfig.Netflow=%s, c.NetworkTemplateConfig.Netflow=%s", v.NetworkTemplateConfig.Netflow, c.NetworkTemplateConfig.Netflow)
		return RequiresUpdate
	}
	if v.NetworkTemplateConfig.SviNetflowMonitor != c.NetworkTemplateConfig.SviNetflowMonitor {
		log.Printf("v.NetworkTemplateConfig.SviNetflowMonitor=%s, c.NetworkTemplateConfig.SviNetflowMonitor=%s", v.NetworkTemplateConfig.SviNetflowMonitor, c.NetworkTemplateConfig.SviNetflowMonitor)
		return RequiresUpdate
	}
	if v.NetworkTemplateConfig.VlanNetflowMonitor != c.NetworkTemplateConfig.VlanNetflowMonitor {
		log.Printf("v.NetworkTemplateConfig.VlanNetflowMonitor=%s, c.NetworkTemplateConfig.VlanNetflowMonitor=%s", v.NetworkTemplateConfig.VlanNetflowMonitor, c.NetworkTemplateConfig.VlanNetflowMonitor)
		return RequiresUpdate
	}
	if v.NetworkTemplateConfig.L3GatwayBorder != c.NetworkTemplateConfig.L3GatwayBorder {
		log.Printf("v.NetworkTemplateConfig.L3GatwayBorder=%s, c.NetworkTemplateConfig.L3GatwayBorder=%s", v.NetworkTemplateConfig.L3GatwayBorder, c.NetworkTemplateConfig.L3GatwayBorder)
		return RequiresUpdate
	}
	if v.NetworkTemplateConfig.IgmpVersion != c.NetworkTemplateConfig.IgmpVersion {
		log.Printf("v.NetworkTemplateConfig.IgmpVersion=%s, c.NetworkTemplateConfig.IgmpVersion=%s", v.NetworkTemplateConfig.IgmpVersion, c.NetworkTemplateConfig.IgmpVersion)
		return RequiresUpdate
	}
	if v.DeployAttachments != c.DeployAttachments {
		log.Printf("v.DeployAttachments=%v, c.DeployAttachments=%v", v.DeployAttachments, c.DeployAttachments)
		cf = true
	}

	if cf {
		return ControlFlagUpdate
	}
	return ValuesDeeplyEqual
}

func (v *NDFCDhcpRelayServersValue) CreatePlan(c NDFCDhcpRelayServersValue, cf *bool) int {
	action := ActionNone

	if v.Address != "" {

		if v.Address != c.Address {
			log.Printf("Update: v.Address=%v, c.Address=%v", v.Address, c.Address)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}

	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.Address=%v, c.Address=%v", v.Address, c.Address)
		v.Address = c.Address
	}

	if v.Vrf != "" {

		if v.Vrf != c.Vrf {
			log.Printf("Update: v.Vrf=%v, c.Vrf=%v", v.Vrf, c.Vrf)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}

	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.Vrf=%v, c.Vrf=%v", v.Vrf, c.Vrf)
		v.Vrf = c.Vrf
	}

	return action
}

func (v *NDFCNetworksValue) CreatePlan(c NDFCNetworksValue, cf *bool) int {
	action := ActionNone

	if v.DisplayName != "" {

		if v.DisplayName != c.DisplayName {
			log.Printf("Update: v.DisplayName=%v, c.DisplayName=%v", v.DisplayName, c.DisplayName)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}

	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.DisplayName=%v, c.DisplayName=%v", v.DisplayName, c.DisplayName)
		v.DisplayName = c.DisplayName
	}

	if v.NetworkId != nil && c.NetworkId != nil {
		if *v.NetworkId != *c.NetworkId {
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresReplace
			}
			log.Printf("Update:: v.NetworkId=%v, c.NetworkId=%v", *v.NetworkId, *c.NetworkId)
		}
	} else if v.NetworkId != nil {
		log.Printf("Update: v.NetworkId=%v, c.NetworkId=nil", *v.NetworkId)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresReplace
		}
	} else if c.NetworkId != nil {
		v.NetworkId = new(int64)
		log.Printf("Copy from state: v.NetworkId=nil, c.NetworkId=%v", *c.NetworkId)
		*v.NetworkId = *c.NetworkId
	}
	if v.NetworkTemplate != "" {

		if v.NetworkTemplate != c.NetworkTemplate {
			log.Printf("Update: v.NetworkTemplate=%v, c.NetworkTemplate=%v", v.NetworkTemplate, c.NetworkTemplate)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}

	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NetworkTemplate=%v, c.NetworkTemplate=%v", v.NetworkTemplate, c.NetworkTemplate)
		v.NetworkTemplate = c.NetworkTemplate
	}

	if v.NetworkExtensionTemplate != "" {

		if v.NetworkExtensionTemplate != c.NetworkExtensionTemplate {
			log.Printf("Update: v.NetworkExtensionTemplate=%v, c.NetworkExtensionTemplate=%v", v.NetworkExtensionTemplate, c.NetworkExtensionTemplate)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}

	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NetworkExtensionTemplate=%v, c.NetworkExtensionTemplate=%v", v.NetworkExtensionTemplate, c.NetworkExtensionTemplate)
		v.NetworkExtensionTemplate = c.NetworkExtensionTemplate
	}

	if v.VrfName != "" {

		if v.VrfName != c.VrfName {
			log.Printf("Update: v.VrfName=%v, c.VrfName=%v", v.VrfName, c.VrfName)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}

	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.VrfName=%v, c.VrfName=%v", v.VrfName, c.VrfName)
		v.VrfName = c.VrfName
	}

	if !v.PrimaryNetworkId.IsEmpty() && !c.PrimaryNetworkId.IsEmpty() {
		if *v.PrimaryNetworkId != *c.PrimaryNetworkId {
			log.Printf("Update: v.PrimaryNetworkId=%v, c.PrimaryNetworkId=%v", *v.PrimaryNetworkId, *c.PrimaryNetworkId)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		if !v.PrimaryNetworkId.IsEmpty() {
			log.Printf("Update: v.PrimaryNetworkId=%v", *v.PrimaryNetworkId)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		} else if !c.PrimaryNetworkId.IsEmpty() {
			log.Printf("Copy from State: c.PrimaryNetworkId=%v", *c.PrimaryNetworkId)
			v.PrimaryNetworkId = new(Int64Custom)
			*v.PrimaryNetworkId = *c.PrimaryNetworkId
		}
	}
	if v.NetworkType != "" {

		if v.NetworkType != c.NetworkType {
			log.Printf("Update: v.NetworkType=%v, c.NetworkType=%v", v.NetworkType, c.NetworkType)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}

	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NetworkType=%v, c.NetworkType=%v", v.NetworkType, c.NetworkType)
		v.NetworkType = c.NetworkType
	}

	if v.NetworkTemplateConfig.GatewayIpv4Address != "" {
		if v.NetworkTemplateConfig.GatewayIpv4Address != c.NetworkTemplateConfig.GatewayIpv4Address {
			log.Printf("Update: v.NetworkTemplateConfig.GatewayIpv4Address=%v, c.NetworkTemplateConfig.GatewayIpv4Address=%v", v.NetworkTemplateConfig.GatewayIpv4Address, c.NetworkTemplateConfig.GatewayIpv4Address)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NetworkTemplateConfig.GatewayIpv4Address=%v, c.NetworkTemplateConfig.GatewayIpv4Address=%v", v.NetworkTemplateConfig.GatewayIpv4Address, c.NetworkTemplateConfig.GatewayIpv4Address)
		v.NetworkTemplateConfig.GatewayIpv4Address = c.NetworkTemplateConfig.GatewayIpv4Address
	}

	if v.NetworkTemplateConfig.GatewayIpv6Address != "" {
		if v.NetworkTemplateConfig.GatewayIpv6Address != c.NetworkTemplateConfig.GatewayIpv6Address {
			log.Printf("Update: v.NetworkTemplateConfig.GatewayIpv6Address=%v, c.NetworkTemplateConfig.GatewayIpv6Address=%v", v.NetworkTemplateConfig.GatewayIpv6Address, c.NetworkTemplateConfig.GatewayIpv6Address)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NetworkTemplateConfig.GatewayIpv6Address=%v, c.NetworkTemplateConfig.GatewayIpv6Address=%v", v.NetworkTemplateConfig.GatewayIpv6Address, c.NetworkTemplateConfig.GatewayIpv6Address)
		v.NetworkTemplateConfig.GatewayIpv6Address = c.NetworkTemplateConfig.GatewayIpv6Address
	}

	if !v.NetworkTemplateConfig.VlanId.IsEmpty() && !c.NetworkTemplateConfig.VlanId.IsEmpty() {
		if *v.NetworkTemplateConfig.VlanId != *c.NetworkTemplateConfig.VlanId {
			log.Printf("Update: v.NetworkTemplateConfig.VlanId=%v, c.NetworkTemplateConfig.VlanId=%v", *v.NetworkTemplateConfig.VlanId, *c.NetworkTemplateConfig.VlanId)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else if !v.NetworkTemplateConfig.VlanId.IsEmpty() {
		log.Printf("Update: v.NetworkTemplateConfig.VlanId=%v", *v.NetworkTemplateConfig.VlanId)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if !c.NetworkTemplateConfig.VlanId.IsEmpty() {
		log.Printf("Copy from State: c.NetworkTemplateConfig.VlanId=%v", *c.NetworkTemplateConfig.VlanId)
		v.NetworkTemplateConfig.VlanId = new(Int64Custom)
		*v.NetworkTemplateConfig.VlanId = *c.NetworkTemplateConfig.VlanId
	}

	if v.NetworkTemplateConfig.VlanName != "" {
		if v.NetworkTemplateConfig.VlanName != c.NetworkTemplateConfig.VlanName {
			log.Printf("Update: v.NetworkTemplateConfig.VlanName=%v, c.NetworkTemplateConfig.VlanName=%v", v.NetworkTemplateConfig.VlanName, c.NetworkTemplateConfig.VlanName)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NetworkTemplateConfig.VlanName=%v, c.NetworkTemplateConfig.VlanName=%v", v.NetworkTemplateConfig.VlanName, c.NetworkTemplateConfig.VlanName)
		v.NetworkTemplateConfig.VlanName = c.NetworkTemplateConfig.VlanName
	}

	if v.NetworkTemplateConfig.Layer2Only != c.NetworkTemplateConfig.Layer2Only {
		log.Printf("Update: v.NetworkTemplateConfig.Layer2Only=%v, c.NetworkTemplateConfig.Layer2Only=%v", v.NetworkTemplateConfig.Layer2Only, c.NetworkTemplateConfig.Layer2Only)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NetworkTemplateConfig.InterfaceDescription != "" {
		if v.NetworkTemplateConfig.InterfaceDescription != c.NetworkTemplateConfig.InterfaceDescription {
			log.Printf("Update: v.NetworkTemplateConfig.InterfaceDescription=%v, c.NetworkTemplateConfig.InterfaceDescription=%v", v.NetworkTemplateConfig.InterfaceDescription, c.NetworkTemplateConfig.InterfaceDescription)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NetworkTemplateConfig.InterfaceDescription=%v, c.NetworkTemplateConfig.InterfaceDescription=%v", v.NetworkTemplateConfig.InterfaceDescription, c.NetworkTemplateConfig.InterfaceDescription)
		v.NetworkTemplateConfig.InterfaceDescription = c.NetworkTemplateConfig.InterfaceDescription
	}

	if !v.NetworkTemplateConfig.Mtu.IsEmpty() && !c.NetworkTemplateConfig.Mtu.IsEmpty() {
		if *v.NetworkTemplateConfig.Mtu != *c.NetworkTemplateConfig.Mtu {
			log.Printf("Update: v.NetworkTemplateConfig.Mtu=%v, c.NetworkTemplateConfig.Mtu=%v", *v.NetworkTemplateConfig.Mtu, *c.NetworkTemplateConfig.Mtu)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else if !v.NetworkTemplateConfig.Mtu.IsEmpty() {
		log.Printf("Update: v.NetworkTemplateConfig.Mtu=%v", *v.NetworkTemplateConfig.Mtu)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if !c.NetworkTemplateConfig.Mtu.IsEmpty() {
		log.Printf("Copy from State: c.NetworkTemplateConfig.Mtu=%v", *c.NetworkTemplateConfig.Mtu)
		v.NetworkTemplateConfig.Mtu = new(Int64Custom)
		*v.NetworkTemplateConfig.Mtu = *c.NetworkTemplateConfig.Mtu
	}

	if v.NetworkTemplateConfig.SecondaryGateway1 != "" {
		if v.NetworkTemplateConfig.SecondaryGateway1 != c.NetworkTemplateConfig.SecondaryGateway1 {
			log.Printf("Update: v.NetworkTemplateConfig.SecondaryGateway1=%v, c.NetworkTemplateConfig.SecondaryGateway1=%v", v.NetworkTemplateConfig.SecondaryGateway1, c.NetworkTemplateConfig.SecondaryGateway1)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NetworkTemplateConfig.SecondaryGateway1=%v, c.NetworkTemplateConfig.SecondaryGateway1=%v", v.NetworkTemplateConfig.SecondaryGateway1, c.NetworkTemplateConfig.SecondaryGateway1)
		v.NetworkTemplateConfig.SecondaryGateway1 = c.NetworkTemplateConfig.SecondaryGateway1
	}

	if v.NetworkTemplateConfig.SecondaryGateway2 != "" {
		if v.NetworkTemplateConfig.SecondaryGateway2 != c.NetworkTemplateConfig.SecondaryGateway2 {
			log.Printf("Update: v.NetworkTemplateConfig.SecondaryGateway2=%v, c.NetworkTemplateConfig.SecondaryGateway2=%v", v.NetworkTemplateConfig.SecondaryGateway2, c.NetworkTemplateConfig.SecondaryGateway2)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NetworkTemplateConfig.SecondaryGateway2=%v, c.NetworkTemplateConfig.SecondaryGateway2=%v", v.NetworkTemplateConfig.SecondaryGateway2, c.NetworkTemplateConfig.SecondaryGateway2)
		v.NetworkTemplateConfig.SecondaryGateway2 = c.NetworkTemplateConfig.SecondaryGateway2
	}

	if v.NetworkTemplateConfig.SecondaryGateway3 != "" {
		if v.NetworkTemplateConfig.SecondaryGateway3 != c.NetworkTemplateConfig.SecondaryGateway3 {
			log.Printf("Update: v.NetworkTemplateConfig.SecondaryGateway3=%v, c.NetworkTemplateConfig.SecondaryGateway3=%v", v.NetworkTemplateConfig.SecondaryGateway3, c.NetworkTemplateConfig.SecondaryGateway3)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NetworkTemplateConfig.SecondaryGateway3=%v, c.NetworkTemplateConfig.SecondaryGateway3=%v", v.NetworkTemplateConfig.SecondaryGateway3, c.NetworkTemplateConfig.SecondaryGateway3)
		v.NetworkTemplateConfig.SecondaryGateway3 = c.NetworkTemplateConfig.SecondaryGateway3
	}

	if v.NetworkTemplateConfig.SecondaryGateway4 != "" {
		if v.NetworkTemplateConfig.SecondaryGateway4 != c.NetworkTemplateConfig.SecondaryGateway4 {
			log.Printf("Update: v.NetworkTemplateConfig.SecondaryGateway4=%v, c.NetworkTemplateConfig.SecondaryGateway4=%v", v.NetworkTemplateConfig.SecondaryGateway4, c.NetworkTemplateConfig.SecondaryGateway4)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NetworkTemplateConfig.SecondaryGateway4=%v, c.NetworkTemplateConfig.SecondaryGateway4=%v", v.NetworkTemplateConfig.SecondaryGateway4, c.NetworkTemplateConfig.SecondaryGateway4)
		v.NetworkTemplateConfig.SecondaryGateway4 = c.NetworkTemplateConfig.SecondaryGateway4
	}

	if v.NetworkTemplateConfig.ArpSuppression != c.NetworkTemplateConfig.ArpSuppression {
		log.Printf("Update: v.NetworkTemplateConfig.ArpSuppression=%v, c.NetworkTemplateConfig.ArpSuppression=%v", v.NetworkTemplateConfig.ArpSuppression, c.NetworkTemplateConfig.ArpSuppression)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NetworkTemplateConfig.IngressReplication != c.NetworkTemplateConfig.IngressReplication {
		log.Printf("Update: v.NetworkTemplateConfig.IngressReplication=%v, c.NetworkTemplateConfig.IngressReplication=%v", v.NetworkTemplateConfig.IngressReplication, c.NetworkTemplateConfig.IngressReplication)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NetworkTemplateConfig.MulticastGroup != "" {
		if v.NetworkTemplateConfig.MulticastGroup != c.NetworkTemplateConfig.MulticastGroup {
			log.Printf("Update: v.NetworkTemplateConfig.MulticastGroup=%v, c.NetworkTemplateConfig.MulticastGroup=%v", v.NetworkTemplateConfig.MulticastGroup, c.NetworkTemplateConfig.MulticastGroup)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NetworkTemplateConfig.MulticastGroup=%v, c.NetworkTemplateConfig.MulticastGroup=%v", v.NetworkTemplateConfig.MulticastGroup, c.NetworkTemplateConfig.MulticastGroup)
		v.NetworkTemplateConfig.MulticastGroup = c.NetworkTemplateConfig.MulticastGroup
	}

	if len(v.NetworkTemplateConfig.DhcpRelayServers) != len(c.NetworkTemplateConfig.DhcpRelayServers) {
		log.Printf("Update: len(v.NetworkTemplateConfig.DhcpRelayServers)=%d, len(c.NetworkTemplateConfig.DhcpRelayServers)=%d", len(v.NetworkTemplateConfig.DhcpRelayServers), len(c.NetworkTemplateConfig.DhcpRelayServers))
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}
	for i := range v.NetworkTemplateConfig.DhcpRelayServers {
		retVal := v.NetworkTemplateConfig.DhcpRelayServers[i].CreatePlan(c.NetworkTemplateConfig.DhcpRelayServers[i], cf)
		if retVal != ActionNone {
			if action == ActionNone || action == RequiresUpdate {
				action = retVal
			}
		}
	}
	if !v.NetworkTemplateConfig.DhcpRelayLoopbackId.IsEmpty() && !c.NetworkTemplateConfig.DhcpRelayLoopbackId.IsEmpty() {
		if *v.NetworkTemplateConfig.DhcpRelayLoopbackId != *c.NetworkTemplateConfig.DhcpRelayLoopbackId {
			log.Printf("Update: v.NetworkTemplateConfig.DhcpRelayLoopbackId=%v, c.NetworkTemplateConfig.DhcpRelayLoopbackId=%v", *v.NetworkTemplateConfig.DhcpRelayLoopbackId, *c.NetworkTemplateConfig.DhcpRelayLoopbackId)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else if !v.NetworkTemplateConfig.DhcpRelayLoopbackId.IsEmpty() {
		log.Printf("Update: v.NetworkTemplateConfig.DhcpRelayLoopbackId=%v", *v.NetworkTemplateConfig.DhcpRelayLoopbackId)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if !c.NetworkTemplateConfig.DhcpRelayLoopbackId.IsEmpty() {
		log.Printf("Copy from State: c.NetworkTemplateConfig.DhcpRelayLoopbackId=%v", *c.NetworkTemplateConfig.DhcpRelayLoopbackId)
		v.NetworkTemplateConfig.DhcpRelayLoopbackId = new(Int64Custom)
		*v.NetworkTemplateConfig.DhcpRelayLoopbackId = *c.NetworkTemplateConfig.DhcpRelayLoopbackId
	}

	if !v.NetworkTemplateConfig.RoutingTag.IsEmpty() && !c.NetworkTemplateConfig.RoutingTag.IsEmpty() {
		if *v.NetworkTemplateConfig.RoutingTag != *c.NetworkTemplateConfig.RoutingTag {
			log.Printf("Update: v.NetworkTemplateConfig.RoutingTag=%v, c.NetworkTemplateConfig.RoutingTag=%v", *v.NetworkTemplateConfig.RoutingTag, *c.NetworkTemplateConfig.RoutingTag)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else if !v.NetworkTemplateConfig.RoutingTag.IsEmpty() {
		log.Printf("Update: v.NetworkTemplateConfig.RoutingTag=%v", *v.NetworkTemplateConfig.RoutingTag)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if !c.NetworkTemplateConfig.RoutingTag.IsEmpty() {
		log.Printf("Copy from State: c.NetworkTemplateConfig.RoutingTag=%v", *c.NetworkTemplateConfig.RoutingTag)
		v.NetworkTemplateConfig.RoutingTag = new(Int64Custom)
		*v.NetworkTemplateConfig.RoutingTag = *c.NetworkTemplateConfig.RoutingTag
	}

	if v.NetworkTemplateConfig.Trm != c.NetworkTemplateConfig.Trm {
		log.Printf("Update: v.NetworkTemplateConfig.Trm=%v, c.NetworkTemplateConfig.Trm=%v", v.NetworkTemplateConfig.Trm, c.NetworkTemplateConfig.Trm)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NetworkTemplateConfig.RouteTargetBoth != c.NetworkTemplateConfig.RouteTargetBoth {
		log.Printf("Update: v.NetworkTemplateConfig.RouteTargetBoth=%v, c.NetworkTemplateConfig.RouteTargetBoth=%v", v.NetworkTemplateConfig.RouteTargetBoth, c.NetworkTemplateConfig.RouteTargetBoth)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NetworkTemplateConfig.Netflow != c.NetworkTemplateConfig.Netflow {
		log.Printf("Update: v.NetworkTemplateConfig.Netflow=%v, c.NetworkTemplateConfig.Netflow=%v", v.NetworkTemplateConfig.Netflow, c.NetworkTemplateConfig.Netflow)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NetworkTemplateConfig.SviNetflowMonitor != "" {
		if v.NetworkTemplateConfig.SviNetflowMonitor != c.NetworkTemplateConfig.SviNetflowMonitor {
			log.Printf("Update: v.NetworkTemplateConfig.SviNetflowMonitor=%v, c.NetworkTemplateConfig.SviNetflowMonitor=%v", v.NetworkTemplateConfig.SviNetflowMonitor, c.NetworkTemplateConfig.SviNetflowMonitor)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NetworkTemplateConfig.SviNetflowMonitor=%v, c.NetworkTemplateConfig.SviNetflowMonitor=%v", v.NetworkTemplateConfig.SviNetflowMonitor, c.NetworkTemplateConfig.SviNetflowMonitor)
		v.NetworkTemplateConfig.SviNetflowMonitor = c.NetworkTemplateConfig.SviNetflowMonitor
	}

	if v.NetworkTemplateConfig.VlanNetflowMonitor != "" {
		if v.NetworkTemplateConfig.VlanNetflowMonitor != c.NetworkTemplateConfig.VlanNetflowMonitor {
			log.Printf("Update: v.NetworkTemplateConfig.VlanNetflowMonitor=%v, c.NetworkTemplateConfig.VlanNetflowMonitor=%v", v.NetworkTemplateConfig.VlanNetflowMonitor, c.NetworkTemplateConfig.VlanNetflowMonitor)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NetworkTemplateConfig.VlanNetflowMonitor=%v, c.NetworkTemplateConfig.VlanNetflowMonitor=%v", v.NetworkTemplateConfig.VlanNetflowMonitor, c.NetworkTemplateConfig.VlanNetflowMonitor)
		v.NetworkTemplateConfig.VlanNetflowMonitor = c.NetworkTemplateConfig.VlanNetflowMonitor
	}

	if v.NetworkTemplateConfig.L3GatwayBorder != c.NetworkTemplateConfig.L3GatwayBorder {
		log.Printf("Update: v.NetworkTemplateConfig.L3GatwayBorder=%v, c.NetworkTemplateConfig.L3GatwayBorder=%v", v.NetworkTemplateConfig.L3GatwayBorder, c.NetworkTemplateConfig.L3GatwayBorder)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.NetworkTemplateConfig.IgmpVersion != "" {
		if v.NetworkTemplateConfig.IgmpVersion != c.NetworkTemplateConfig.IgmpVersion {
			log.Printf("Update: v.NetworkTemplateConfig.IgmpVersion=%v, c.NetworkTemplateConfig.IgmpVersion=%v", v.NetworkTemplateConfig.IgmpVersion, c.NetworkTemplateConfig.IgmpVersion)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NetworkTemplateConfig.IgmpVersion=%v, c.NetworkTemplateConfig.IgmpVersion=%v", v.NetworkTemplateConfig.IgmpVersion, c.NetworkTemplateConfig.IgmpVersion)
		v.NetworkTemplateConfig.IgmpVersion = c.NetworkTemplateConfig.IgmpVersion
	}

	if v.DeployAttachments != c.DeployAttachments {
		log.Printf("Update: v.DeployAttachments=%v, c.DeployAttachments=%v", v.DeployAttachments, c.DeployAttachments)
		*cf = true
	}

	return action
}
