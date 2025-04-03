// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package resource_vrf_bulk

import (
	"log"
	. "terraform-provider-ndfc/internal/provider/types"
)

func (v NDFCVrfsValue) DeepEqual(c NDFCVrfsValue) int {
	cf := false
	if v.VrfTemplate != c.VrfTemplate {
		log.Printf("v.VrfTemplate=%v, c.VrfTemplate=%v", v.VrfTemplate, c.VrfTemplate)
		return RequiresUpdate
	}
	if v.VrfExtensionTemplate != c.VrfExtensionTemplate {
		log.Printf("v.VrfExtensionTemplate=%v, c.VrfExtensionTemplate=%v", v.VrfExtensionTemplate, c.VrfExtensionTemplate)
		return RequiresUpdate
	}

	if v.VrfId != nil && c.VrfId != nil {
		if *v.VrfId != *c.VrfId {
			log.Printf("v.VrfId=%v, c.VrfId=%v", *v.VrfId, *c.VrfId)
			return RequiresReplace
		}
	} else {
		if v.VrfId != nil {
			log.Printf("v.VrfId=%v", *v.VrfId)
			return RequiresReplace
		} else if c.VrfId != nil {
			log.Printf("c.VrfId=%v", *c.VrfId)
			return RequiresReplace
		}
	}

	if !v.VrfTemplateConfig.VlanId.IsEmpty() && !c.VrfTemplateConfig.VlanId.IsEmpty() {
		if *v.VrfTemplateConfig.VlanId != *c.VrfTemplateConfig.VlanId {
			log.Printf("v.VrfTemplateConfig.VlanId=%v, c.VrfTemplateConfig.VlanId=%v", *v.VrfTemplateConfig.VlanId, *c.VrfTemplateConfig.VlanId)
			return RequiresUpdate
		}
	} else {
		if !v.VrfTemplateConfig.VlanId.IsEmpty() {
			log.Printf("v.VrfTemplateConfig.VlanId=%v", *v.VrfTemplateConfig.VlanId)
			return RequiresUpdate
		} else if !c.VrfTemplateConfig.VlanId.IsEmpty() {
			log.Printf("c.VrfTemplateConfig.VlanId=%v", *c.VrfTemplateConfig.VlanId)
			return RequiresUpdate
		}
	}
	if v.VrfTemplateConfig.VlanName != c.VrfTemplateConfig.VlanName {
		log.Printf("v.VrfTemplateConfig.VlanName=%s, c.VrfTemplateConfig.VlanName=%s", v.VrfTemplateConfig.VlanName, c.VrfTemplateConfig.VlanName)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.InterfaceDescription != c.VrfTemplateConfig.InterfaceDescription {
		log.Printf("v.VrfTemplateConfig.InterfaceDescription=%s, c.VrfTemplateConfig.InterfaceDescription=%s", v.VrfTemplateConfig.InterfaceDescription, c.VrfTemplateConfig.InterfaceDescription)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.VrfDescription != c.VrfTemplateConfig.VrfDescription {
		log.Printf("v.VrfTemplateConfig.VrfDescription=%s, c.VrfTemplateConfig.VrfDescription=%s", v.VrfTemplateConfig.VrfDescription, c.VrfTemplateConfig.VrfDescription)
		return RequiresUpdate
	}

	if v.VrfTemplateConfig.Mtu != nil && c.VrfTemplateConfig.Mtu != nil {
		if *v.VrfTemplateConfig.Mtu != *c.VrfTemplateConfig.Mtu {
			log.Printf("v.VrfTemplateConfig.Mtu=%v, c.VrfTemplateConfig.Mtu=%v", *v.VrfTemplateConfig.Mtu, *c.VrfTemplateConfig.Mtu)
			return RequiresUpdate
		}
	} else {
		if v.VrfTemplateConfig.Mtu != nil {
			log.Printf("v.VrfTemplateConfig.Mtu=%v", *v.VrfTemplateConfig.Mtu)
			return RequiresUpdate
		} else if c.VrfTemplateConfig.Mtu != nil {
			log.Printf("c.VrfTemplateConfig.Mtu=%v", *c.VrfTemplateConfig.Mtu)
			return RequiresUpdate
		}
	}

	if v.VrfTemplateConfig.LoopbackRoutingTag != nil && c.VrfTemplateConfig.LoopbackRoutingTag != nil {
		if *v.VrfTemplateConfig.LoopbackRoutingTag != *c.VrfTemplateConfig.LoopbackRoutingTag {
			log.Printf("v.VrfTemplateConfig.LoopbackRoutingTag=%v, c.VrfTemplateConfig.LoopbackRoutingTag=%v", *v.VrfTemplateConfig.LoopbackRoutingTag, *c.VrfTemplateConfig.LoopbackRoutingTag)
			return RequiresUpdate
		}
	} else {
		if v.VrfTemplateConfig.LoopbackRoutingTag != nil {
			log.Printf("v.VrfTemplateConfig.LoopbackRoutingTag=%v", *v.VrfTemplateConfig.LoopbackRoutingTag)
			return RequiresUpdate
		} else if c.VrfTemplateConfig.LoopbackRoutingTag != nil {
			log.Printf("c.VrfTemplateConfig.LoopbackRoutingTag=%v", *c.VrfTemplateConfig.LoopbackRoutingTag)
			return RequiresUpdate
		}
	}

	if v.VrfTemplateConfig.RedistributeDirectRouteMap != c.VrfTemplateConfig.RedistributeDirectRouteMap {
		log.Printf("v.VrfTemplateConfig.RedistributeDirectRouteMap=%s, c.VrfTemplateConfig.RedistributeDirectRouteMap=%s", v.VrfTemplateConfig.RedistributeDirectRouteMap, c.VrfTemplateConfig.RedistributeDirectRouteMap)
		return RequiresUpdate
	}

	if v.VrfTemplateConfig.MaxBgpPaths != nil && c.VrfTemplateConfig.MaxBgpPaths != nil {
		if *v.VrfTemplateConfig.MaxBgpPaths != *c.VrfTemplateConfig.MaxBgpPaths {
			log.Printf("v.VrfTemplateConfig.MaxBgpPaths=%v, c.VrfTemplateConfig.MaxBgpPaths=%v", *v.VrfTemplateConfig.MaxBgpPaths, *c.VrfTemplateConfig.MaxBgpPaths)
			return RequiresUpdate
		}
	} else {
		if v.VrfTemplateConfig.MaxBgpPaths != nil {
			log.Printf("v.VrfTemplateConfig.MaxBgpPaths=%v", *v.VrfTemplateConfig.MaxBgpPaths)
			return RequiresUpdate
		} else if c.VrfTemplateConfig.MaxBgpPaths != nil {
			log.Printf("c.VrfTemplateConfig.MaxBgpPaths=%v", *c.VrfTemplateConfig.MaxBgpPaths)
			return RequiresUpdate
		}
	}

	if v.VrfTemplateConfig.MaxIbgpPaths != nil && c.VrfTemplateConfig.MaxIbgpPaths != nil {
		if *v.VrfTemplateConfig.MaxIbgpPaths != *c.VrfTemplateConfig.MaxIbgpPaths {
			log.Printf("v.VrfTemplateConfig.MaxIbgpPaths=%v, c.VrfTemplateConfig.MaxIbgpPaths=%v", *v.VrfTemplateConfig.MaxIbgpPaths, *c.VrfTemplateConfig.MaxIbgpPaths)
			return RequiresUpdate
		}
	} else {
		if v.VrfTemplateConfig.MaxIbgpPaths != nil {
			log.Printf("v.VrfTemplateConfig.MaxIbgpPaths=%v", *v.VrfTemplateConfig.MaxIbgpPaths)
			return RequiresUpdate
		} else if c.VrfTemplateConfig.MaxIbgpPaths != nil {
			log.Printf("c.VrfTemplateConfig.MaxIbgpPaths=%v", *c.VrfTemplateConfig.MaxIbgpPaths)
			return RequiresUpdate
		}
	}

	if v.VrfTemplateConfig.Ipv6LinkLocal != c.VrfTemplateConfig.Ipv6LinkLocal {
		log.Printf("v.VrfTemplateConfig.Ipv6LinkLocal=%s, c.VrfTemplateConfig.Ipv6LinkLocal=%s", v.VrfTemplateConfig.Ipv6LinkLocal, c.VrfTemplateConfig.Ipv6LinkLocal)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.Trm != c.VrfTemplateConfig.Trm {
		log.Printf("v.VrfTemplateConfig.Trm=%s, c.VrfTemplateConfig.Trm=%s", v.VrfTemplateConfig.Trm, c.VrfTemplateConfig.Trm)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.NoRp != c.VrfTemplateConfig.NoRp {
		log.Printf("v.VrfTemplateConfig.NoRp=%s, c.VrfTemplateConfig.NoRp=%s", v.VrfTemplateConfig.NoRp, c.VrfTemplateConfig.NoRp)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.RpExternal != c.VrfTemplateConfig.RpExternal {
		log.Printf("v.VrfTemplateConfig.RpExternal=%s, c.VrfTemplateConfig.RpExternal=%s", v.VrfTemplateConfig.RpExternal, c.VrfTemplateConfig.RpExternal)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.RpAddress != c.VrfTemplateConfig.RpAddress {
		log.Printf("v.VrfTemplateConfig.RpAddress=%s, c.VrfTemplateConfig.RpAddress=%s", v.VrfTemplateConfig.RpAddress, c.VrfTemplateConfig.RpAddress)
		return RequiresUpdate
	}

	if !v.VrfTemplateConfig.RpLoopbackId.IsEmpty() && !c.VrfTemplateConfig.RpLoopbackId.IsEmpty() {
		if *v.VrfTemplateConfig.RpLoopbackId != *c.VrfTemplateConfig.RpLoopbackId {
			log.Printf("v.VrfTemplateConfig.RpLoopbackId=%v, c.VrfTemplateConfig.RpLoopbackId=%v", *v.VrfTemplateConfig.RpLoopbackId, *c.VrfTemplateConfig.RpLoopbackId)
			return RequiresUpdate
		}
	} else {
		if !v.VrfTemplateConfig.RpLoopbackId.IsEmpty() {
			log.Printf("v.VrfTemplateConfig.RpLoopbackId=%v", *v.VrfTemplateConfig.RpLoopbackId)
			return RequiresUpdate
		} else if !c.VrfTemplateConfig.RpLoopbackId.IsEmpty() {
			log.Printf("c.VrfTemplateConfig.RpLoopbackId=%v", *c.VrfTemplateConfig.RpLoopbackId)
			return RequiresUpdate
		}
	}
	if v.VrfTemplateConfig.UnderlayMulticastAddress != c.VrfTemplateConfig.UnderlayMulticastAddress {
		log.Printf("v.VrfTemplateConfig.UnderlayMulticastAddress=%s, c.VrfTemplateConfig.UnderlayMulticastAddress=%s", v.VrfTemplateConfig.UnderlayMulticastAddress, c.VrfTemplateConfig.UnderlayMulticastAddress)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.OverlayMulticastGroups != c.VrfTemplateConfig.OverlayMulticastGroups {
		log.Printf("v.VrfTemplateConfig.OverlayMulticastGroups=%s, c.VrfTemplateConfig.OverlayMulticastGroups=%s", v.VrfTemplateConfig.OverlayMulticastGroups, c.VrfTemplateConfig.OverlayMulticastGroups)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.MvpnInterAs != c.VrfTemplateConfig.MvpnInterAs {
		log.Printf("v.VrfTemplateConfig.MvpnInterAs=%s, c.VrfTemplateConfig.MvpnInterAs=%s", v.VrfTemplateConfig.MvpnInterAs, c.VrfTemplateConfig.MvpnInterAs)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.TrmBgwMsite != c.VrfTemplateConfig.TrmBgwMsite {
		log.Printf("v.VrfTemplateConfig.TrmBgwMsite=%s, c.VrfTemplateConfig.TrmBgwMsite=%s", v.VrfTemplateConfig.TrmBgwMsite, c.VrfTemplateConfig.TrmBgwMsite)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.AdvertiseHostRoutes != c.VrfTemplateConfig.AdvertiseHostRoutes {
		log.Printf("v.VrfTemplateConfig.AdvertiseHostRoutes=%s, c.VrfTemplateConfig.AdvertiseHostRoutes=%s", v.VrfTemplateConfig.AdvertiseHostRoutes, c.VrfTemplateConfig.AdvertiseHostRoutes)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.AdvertiseDefaultRoute != c.VrfTemplateConfig.AdvertiseDefaultRoute {
		log.Printf("v.VrfTemplateConfig.AdvertiseDefaultRoute=%s, c.VrfTemplateConfig.AdvertiseDefaultRoute=%s", v.VrfTemplateConfig.AdvertiseDefaultRoute, c.VrfTemplateConfig.AdvertiseDefaultRoute)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.ConfigureStaticDefaultRoute != c.VrfTemplateConfig.ConfigureStaticDefaultRoute {
		log.Printf("v.VrfTemplateConfig.ConfigureStaticDefaultRoute=%s, c.VrfTemplateConfig.ConfigureStaticDefaultRoute=%s", v.VrfTemplateConfig.ConfigureStaticDefaultRoute, c.VrfTemplateConfig.ConfigureStaticDefaultRoute)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.BgpPassword != c.VrfTemplateConfig.BgpPassword {
		log.Printf("v.VrfTemplateConfig.BgpPassword=%s, c.VrfTemplateConfig.BgpPassword=%s", v.VrfTemplateConfig.BgpPassword, c.VrfTemplateConfig.BgpPassword)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.BgpPasswordType != c.VrfTemplateConfig.BgpPasswordType {
		log.Printf("v.VrfTemplateConfig.BgpPasswordType=%s, c.VrfTemplateConfig.BgpPasswordType=%s", v.VrfTemplateConfig.BgpPasswordType, c.VrfTemplateConfig.BgpPasswordType)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.Netflow != c.VrfTemplateConfig.Netflow {
		log.Printf("v.VrfTemplateConfig.Netflow=%s, c.VrfTemplateConfig.Netflow=%s", v.VrfTemplateConfig.Netflow, c.VrfTemplateConfig.Netflow)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.NetflowMonitor != c.VrfTemplateConfig.NetflowMonitor {
		log.Printf("v.VrfTemplateConfig.NetflowMonitor=%s, c.VrfTemplateConfig.NetflowMonitor=%s", v.VrfTemplateConfig.NetflowMonitor, c.VrfTemplateConfig.NetflowMonitor)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.DisableRtAuto != c.VrfTemplateConfig.DisableRtAuto {
		log.Printf("v.VrfTemplateConfig.DisableRtAuto=%s, c.VrfTemplateConfig.DisableRtAuto=%s", v.VrfTemplateConfig.DisableRtAuto, c.VrfTemplateConfig.DisableRtAuto)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.RouteTargetImport != c.VrfTemplateConfig.RouteTargetImport {
		log.Printf("v.VrfTemplateConfig.RouteTargetImport=%s, c.VrfTemplateConfig.RouteTargetImport=%s", v.VrfTemplateConfig.RouteTargetImport, c.VrfTemplateConfig.RouteTargetImport)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.RouteTargetExport != c.VrfTemplateConfig.RouteTargetExport {
		log.Printf("v.VrfTemplateConfig.RouteTargetExport=%s, c.VrfTemplateConfig.RouteTargetExport=%s", v.VrfTemplateConfig.RouteTargetExport, c.VrfTemplateConfig.RouteTargetExport)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.RouteTargetImportEvpn != c.VrfTemplateConfig.RouteTargetImportEvpn {
		log.Printf("v.VrfTemplateConfig.RouteTargetImportEvpn=%s, c.VrfTemplateConfig.RouteTargetImportEvpn=%s", v.VrfTemplateConfig.RouteTargetImportEvpn, c.VrfTemplateConfig.RouteTargetImportEvpn)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.RouteTargetExportEvpn != c.VrfTemplateConfig.RouteTargetExportEvpn {
		log.Printf("v.VrfTemplateConfig.RouteTargetExportEvpn=%s, c.VrfTemplateConfig.RouteTargetExportEvpn=%s", v.VrfTemplateConfig.RouteTargetExportEvpn, c.VrfTemplateConfig.RouteTargetExportEvpn)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.RouteTargetImportMvpn != c.VrfTemplateConfig.RouteTargetImportMvpn {
		log.Printf("v.VrfTemplateConfig.RouteTargetImportMvpn=%s, c.VrfTemplateConfig.RouteTargetImportMvpn=%s", v.VrfTemplateConfig.RouteTargetImportMvpn, c.VrfTemplateConfig.RouteTargetImportMvpn)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.RouteTargetExportMvpn != c.VrfTemplateConfig.RouteTargetExportMvpn {
		log.Printf("v.VrfTemplateConfig.RouteTargetExportMvpn=%s, c.VrfTemplateConfig.RouteTargetExportMvpn=%s", v.VrfTemplateConfig.RouteTargetExportMvpn, c.VrfTemplateConfig.RouteTargetExportMvpn)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.RouteTargetImportCloudEvpn != c.VrfTemplateConfig.RouteTargetImportCloudEvpn {
		log.Printf("v.VrfTemplateConfig.RouteTargetImportCloudEvpn=%s, c.VrfTemplateConfig.RouteTargetImportCloudEvpn=%s", v.VrfTemplateConfig.RouteTargetImportCloudEvpn, c.VrfTemplateConfig.RouteTargetImportCloudEvpn)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.RouteTargetExportCloudEvpn != c.VrfTemplateConfig.RouteTargetExportCloudEvpn {
		log.Printf("v.VrfTemplateConfig.RouteTargetExportCloudEvpn=%s, c.VrfTemplateConfig.RouteTargetExportCloudEvpn=%s", v.VrfTemplateConfig.RouteTargetExportCloudEvpn, c.VrfTemplateConfig.RouteTargetExportCloudEvpn)
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

func (v *NDFCVrfsValue) CreatePlan(c NDFCVrfsValue, cf *bool) int {
	action := ActionNone

	if v.VrfTemplate != "" {

		if v.VrfTemplate != c.VrfTemplate {
			log.Printf("Update: v.VrfTemplate=%v, c.VrfTemplate=%v", v.VrfTemplate, c.VrfTemplate)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}

	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.VrfTemplate=%v, c.VrfTemplate=%v", v.VrfTemplate, c.VrfTemplate)
		v.VrfTemplate = c.VrfTemplate
	}

	if v.VrfExtensionTemplate != "" {

		if v.VrfExtensionTemplate != c.VrfExtensionTemplate {
			log.Printf("Update: v.VrfExtensionTemplate=%v, c.VrfExtensionTemplate=%v", v.VrfExtensionTemplate, c.VrfExtensionTemplate)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}

	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.VrfExtensionTemplate=%v, c.VrfExtensionTemplate=%v", v.VrfExtensionTemplate, c.VrfExtensionTemplate)
		v.VrfExtensionTemplate = c.VrfExtensionTemplate
	}

	if v.VrfId != nil && c.VrfId != nil {
		if *v.VrfId != *c.VrfId {
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresReplace
			}
			log.Printf("Update:: v.VrfId=%v, c.VrfId=%v", *v.VrfId, *c.VrfId)
		}
	} else if v.VrfId != nil {
		log.Printf("Update: v.VrfId=%v, c.VrfId=nil", *v.VrfId)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresReplace
		}
	} else if c.VrfId != nil {
		v.VrfId = new(int64)
		log.Printf("Copy from state: v.VrfId=nil, c.VrfId=%v", *c.VrfId)
		*v.VrfId = *c.VrfId
	}
	if !v.VrfTemplateConfig.VlanId.IsEmpty() && !c.VrfTemplateConfig.VlanId.IsEmpty() {
		if *v.VrfTemplateConfig.VlanId != *c.VrfTemplateConfig.VlanId {
			log.Printf("Update: v.VrfTemplateConfig.VlanId=%v, c.VrfTemplateConfig.VlanId=%v", *v.VrfTemplateConfig.VlanId, *c.VrfTemplateConfig.VlanId)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else if !v.VrfTemplateConfig.VlanId.IsEmpty() {
		log.Printf("Update: v.VrfTemplateConfig.VlanId=%v", *v.VrfTemplateConfig.VlanId)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if !c.VrfTemplateConfig.VlanId.IsEmpty() {
		log.Printf("Copy from State: c.VrfTemplateConfig.VlanId=%v", *c.VrfTemplateConfig.VlanId)
		v.VrfTemplateConfig.VlanId = new(Int64Custom)
		*v.VrfTemplateConfig.VlanId = *c.VrfTemplateConfig.VlanId
	}

	if v.VrfTemplateConfig.VlanName != "" {
		if v.VrfTemplateConfig.VlanName != c.VrfTemplateConfig.VlanName {
			log.Printf("Update: v.VrfTemplateConfig.VlanName=%v, c.VrfTemplateConfig.VlanName=%v", v.VrfTemplateConfig.VlanName, c.VrfTemplateConfig.VlanName)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.VrfTemplateConfig.VlanName=%v, c.VrfTemplateConfig.VlanName=%v", v.VrfTemplateConfig.VlanName, c.VrfTemplateConfig.VlanName)
		v.VrfTemplateConfig.VlanName = c.VrfTemplateConfig.VlanName
	}

	if v.VrfTemplateConfig.InterfaceDescription != "" {
		if v.VrfTemplateConfig.InterfaceDescription != c.VrfTemplateConfig.InterfaceDescription {
			log.Printf("Update: v.VrfTemplateConfig.InterfaceDescription=%v, c.VrfTemplateConfig.InterfaceDescription=%v", v.VrfTemplateConfig.InterfaceDescription, c.VrfTemplateConfig.InterfaceDescription)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.VrfTemplateConfig.InterfaceDescription=%v, c.VrfTemplateConfig.InterfaceDescription=%v", v.VrfTemplateConfig.InterfaceDescription, c.VrfTemplateConfig.InterfaceDescription)
		v.VrfTemplateConfig.InterfaceDescription = c.VrfTemplateConfig.InterfaceDescription
	}

	if v.VrfTemplateConfig.VrfDescription != "" {
		if v.VrfTemplateConfig.VrfDescription != c.VrfTemplateConfig.VrfDescription {
			log.Printf("Update: v.VrfTemplateConfig.VrfDescription=%v, c.VrfTemplateConfig.VrfDescription=%v", v.VrfTemplateConfig.VrfDescription, c.VrfTemplateConfig.VrfDescription)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.VrfTemplateConfig.VrfDescription=%v, c.VrfTemplateConfig.VrfDescription=%v", v.VrfTemplateConfig.VrfDescription, c.VrfTemplateConfig.VrfDescription)
		v.VrfTemplateConfig.VrfDescription = c.VrfTemplateConfig.VrfDescription
	}

	if v.VrfTemplateConfig.Mtu != nil && c.VrfTemplateConfig.Mtu != nil {
		if *v.VrfTemplateConfig.Mtu != *c.VrfTemplateConfig.Mtu {
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
			log.Printf("Update: v.VrfTemplateConfig.Mtu=%v, c.VrfTemplateConfig.Mtu=%v", *v.VrfTemplateConfig.Mtu, *c.VrfTemplateConfig.Mtu)
		}
	} else if v.VrfTemplateConfig.Mtu != nil {
		log.Printf("Update: v.VrfTemplateConfig.Mtu=%v, c.VrfTemplateConfig.Mtu=nil", *v.VrfTemplateConfig.Mtu)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if c.VrfTemplateConfig.Mtu != nil {
		v.VrfTemplateConfig.Mtu = new(int64)
		log.Printf("Copy from state: v.VrfTemplateConfig.Mtu=nil, c.VrfTemplateConfig.Mtu=%v", *c.VrfTemplateConfig.Mtu)
		*v.VrfTemplateConfig.Mtu = *c.VrfTemplateConfig.Mtu
	}

	if v.VrfTemplateConfig.LoopbackRoutingTag != nil && c.VrfTemplateConfig.LoopbackRoutingTag != nil {
		if *v.VrfTemplateConfig.LoopbackRoutingTag != *c.VrfTemplateConfig.LoopbackRoutingTag {
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
			log.Printf("Update: v.VrfTemplateConfig.LoopbackRoutingTag=%v, c.VrfTemplateConfig.LoopbackRoutingTag=%v", *v.VrfTemplateConfig.LoopbackRoutingTag, *c.VrfTemplateConfig.LoopbackRoutingTag)
		}
	} else if v.VrfTemplateConfig.LoopbackRoutingTag != nil {
		log.Printf("Update: v.VrfTemplateConfig.LoopbackRoutingTag=%v, c.VrfTemplateConfig.LoopbackRoutingTag=nil", *v.VrfTemplateConfig.LoopbackRoutingTag)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if c.VrfTemplateConfig.LoopbackRoutingTag != nil {
		v.VrfTemplateConfig.LoopbackRoutingTag = new(int64)
		log.Printf("Copy from state: v.VrfTemplateConfig.LoopbackRoutingTag=nil, c.VrfTemplateConfig.LoopbackRoutingTag=%v", *c.VrfTemplateConfig.LoopbackRoutingTag)
		*v.VrfTemplateConfig.LoopbackRoutingTag = *c.VrfTemplateConfig.LoopbackRoutingTag
	}

	if v.VrfTemplateConfig.RedistributeDirectRouteMap != "" {
		if v.VrfTemplateConfig.RedistributeDirectRouteMap != c.VrfTemplateConfig.RedistributeDirectRouteMap {
			log.Printf("Update: v.VrfTemplateConfig.RedistributeDirectRouteMap=%v, c.VrfTemplateConfig.RedistributeDirectRouteMap=%v", v.VrfTemplateConfig.RedistributeDirectRouteMap, c.VrfTemplateConfig.RedistributeDirectRouteMap)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.VrfTemplateConfig.RedistributeDirectRouteMap=%v, c.VrfTemplateConfig.RedistributeDirectRouteMap=%v", v.VrfTemplateConfig.RedistributeDirectRouteMap, c.VrfTemplateConfig.RedistributeDirectRouteMap)
		v.VrfTemplateConfig.RedistributeDirectRouteMap = c.VrfTemplateConfig.RedistributeDirectRouteMap
	}

	if v.VrfTemplateConfig.MaxBgpPaths != nil && c.VrfTemplateConfig.MaxBgpPaths != nil {
		if *v.VrfTemplateConfig.MaxBgpPaths != *c.VrfTemplateConfig.MaxBgpPaths {
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
			log.Printf("Update: v.VrfTemplateConfig.MaxBgpPaths=%v, c.VrfTemplateConfig.MaxBgpPaths=%v", *v.VrfTemplateConfig.MaxBgpPaths, *c.VrfTemplateConfig.MaxBgpPaths)
		}
	} else if v.VrfTemplateConfig.MaxBgpPaths != nil {
		log.Printf("Update: v.VrfTemplateConfig.MaxBgpPaths=%v, c.VrfTemplateConfig.MaxBgpPaths=nil", *v.VrfTemplateConfig.MaxBgpPaths)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if c.VrfTemplateConfig.MaxBgpPaths != nil {
		v.VrfTemplateConfig.MaxBgpPaths = new(int64)
		log.Printf("Copy from state: v.VrfTemplateConfig.MaxBgpPaths=nil, c.VrfTemplateConfig.MaxBgpPaths=%v", *c.VrfTemplateConfig.MaxBgpPaths)
		*v.VrfTemplateConfig.MaxBgpPaths = *c.VrfTemplateConfig.MaxBgpPaths
	}

	if v.VrfTemplateConfig.MaxIbgpPaths != nil && c.VrfTemplateConfig.MaxIbgpPaths != nil {
		if *v.VrfTemplateConfig.MaxIbgpPaths != *c.VrfTemplateConfig.MaxIbgpPaths {
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
			log.Printf("Update: v.VrfTemplateConfig.MaxIbgpPaths=%v, c.VrfTemplateConfig.MaxIbgpPaths=%v", *v.VrfTemplateConfig.MaxIbgpPaths, *c.VrfTemplateConfig.MaxIbgpPaths)
		}
	} else if v.VrfTemplateConfig.MaxIbgpPaths != nil {
		log.Printf("Update: v.VrfTemplateConfig.MaxIbgpPaths=%v, c.VrfTemplateConfig.MaxIbgpPaths=nil", *v.VrfTemplateConfig.MaxIbgpPaths)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if c.VrfTemplateConfig.MaxIbgpPaths != nil {
		v.VrfTemplateConfig.MaxIbgpPaths = new(int64)
		log.Printf("Copy from state: v.VrfTemplateConfig.MaxIbgpPaths=nil, c.VrfTemplateConfig.MaxIbgpPaths=%v", *c.VrfTemplateConfig.MaxIbgpPaths)
		*v.VrfTemplateConfig.MaxIbgpPaths = *c.VrfTemplateConfig.MaxIbgpPaths
	}

	if v.VrfTemplateConfig.Ipv6LinkLocal != c.VrfTemplateConfig.Ipv6LinkLocal {
		log.Printf("Update: v.VrfTemplateConfig.Ipv6LinkLocal=%v, c.VrfTemplateConfig.Ipv6LinkLocal=%v", v.VrfTemplateConfig.Ipv6LinkLocal, c.VrfTemplateConfig.Ipv6LinkLocal)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.VrfTemplateConfig.Trm != c.VrfTemplateConfig.Trm {
		log.Printf("Update: v.VrfTemplateConfig.Trm=%v, c.VrfTemplateConfig.Trm=%v", v.VrfTemplateConfig.Trm, c.VrfTemplateConfig.Trm)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.VrfTemplateConfig.NoRp != c.VrfTemplateConfig.NoRp {
		log.Printf("Update: v.VrfTemplateConfig.NoRp=%v, c.VrfTemplateConfig.NoRp=%v", v.VrfTemplateConfig.NoRp, c.VrfTemplateConfig.NoRp)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.VrfTemplateConfig.RpExternal != c.VrfTemplateConfig.RpExternal {
		log.Printf("Update: v.VrfTemplateConfig.RpExternal=%v, c.VrfTemplateConfig.RpExternal=%v", v.VrfTemplateConfig.RpExternal, c.VrfTemplateConfig.RpExternal)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.VrfTemplateConfig.RpAddress != "" {
		if v.VrfTemplateConfig.RpAddress != c.VrfTemplateConfig.RpAddress {
			log.Printf("Update: v.VrfTemplateConfig.RpAddress=%v, c.VrfTemplateConfig.RpAddress=%v", v.VrfTemplateConfig.RpAddress, c.VrfTemplateConfig.RpAddress)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.VrfTemplateConfig.RpAddress=%v, c.VrfTemplateConfig.RpAddress=%v", v.VrfTemplateConfig.RpAddress, c.VrfTemplateConfig.RpAddress)
		v.VrfTemplateConfig.RpAddress = c.VrfTemplateConfig.RpAddress
	}

	if !v.VrfTemplateConfig.RpLoopbackId.IsEmpty() && !c.VrfTemplateConfig.RpLoopbackId.IsEmpty() {
		if *v.VrfTemplateConfig.RpLoopbackId != *c.VrfTemplateConfig.RpLoopbackId {
			log.Printf("Update: v.VrfTemplateConfig.RpLoopbackId=%v, c.VrfTemplateConfig.RpLoopbackId=%v", *v.VrfTemplateConfig.RpLoopbackId, *c.VrfTemplateConfig.RpLoopbackId)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else if !v.VrfTemplateConfig.RpLoopbackId.IsEmpty() {
		log.Printf("Update: v.VrfTemplateConfig.RpLoopbackId=%v", *v.VrfTemplateConfig.RpLoopbackId)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if !c.VrfTemplateConfig.RpLoopbackId.IsEmpty() {
		log.Printf("Copy from State: c.VrfTemplateConfig.RpLoopbackId=%v", *c.VrfTemplateConfig.RpLoopbackId)
		v.VrfTemplateConfig.RpLoopbackId = new(Int64Custom)
		*v.VrfTemplateConfig.RpLoopbackId = *c.VrfTemplateConfig.RpLoopbackId
	}

	if v.VrfTemplateConfig.UnderlayMulticastAddress != "" {
		if v.VrfTemplateConfig.UnderlayMulticastAddress != c.VrfTemplateConfig.UnderlayMulticastAddress {
			log.Printf("Update: v.VrfTemplateConfig.UnderlayMulticastAddress=%v, c.VrfTemplateConfig.UnderlayMulticastAddress=%v", v.VrfTemplateConfig.UnderlayMulticastAddress, c.VrfTemplateConfig.UnderlayMulticastAddress)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.VrfTemplateConfig.UnderlayMulticastAddress=%v, c.VrfTemplateConfig.UnderlayMulticastAddress=%v", v.VrfTemplateConfig.UnderlayMulticastAddress, c.VrfTemplateConfig.UnderlayMulticastAddress)
		v.VrfTemplateConfig.UnderlayMulticastAddress = c.VrfTemplateConfig.UnderlayMulticastAddress
	}

	if v.VrfTemplateConfig.OverlayMulticastGroups != "" {
		if v.VrfTemplateConfig.OverlayMulticastGroups != c.VrfTemplateConfig.OverlayMulticastGroups {
			log.Printf("Update: v.VrfTemplateConfig.OverlayMulticastGroups=%v, c.VrfTemplateConfig.OverlayMulticastGroups=%v", v.VrfTemplateConfig.OverlayMulticastGroups, c.VrfTemplateConfig.OverlayMulticastGroups)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.VrfTemplateConfig.OverlayMulticastGroups=%v, c.VrfTemplateConfig.OverlayMulticastGroups=%v", v.VrfTemplateConfig.OverlayMulticastGroups, c.VrfTemplateConfig.OverlayMulticastGroups)
		v.VrfTemplateConfig.OverlayMulticastGroups = c.VrfTemplateConfig.OverlayMulticastGroups
	}

	if v.VrfTemplateConfig.MvpnInterAs != c.VrfTemplateConfig.MvpnInterAs {
		log.Printf("Update: v.VrfTemplateConfig.MvpnInterAs=%v, c.VrfTemplateConfig.MvpnInterAs=%v", v.VrfTemplateConfig.MvpnInterAs, c.VrfTemplateConfig.MvpnInterAs)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.VrfTemplateConfig.TrmBgwMsite != c.VrfTemplateConfig.TrmBgwMsite {
		log.Printf("Update: v.VrfTemplateConfig.TrmBgwMsite=%v, c.VrfTemplateConfig.TrmBgwMsite=%v", v.VrfTemplateConfig.TrmBgwMsite, c.VrfTemplateConfig.TrmBgwMsite)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.VrfTemplateConfig.AdvertiseHostRoutes != c.VrfTemplateConfig.AdvertiseHostRoutes {
		log.Printf("Update: v.VrfTemplateConfig.AdvertiseHostRoutes=%v, c.VrfTemplateConfig.AdvertiseHostRoutes=%v", v.VrfTemplateConfig.AdvertiseHostRoutes, c.VrfTemplateConfig.AdvertiseHostRoutes)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.VrfTemplateConfig.AdvertiseDefaultRoute != c.VrfTemplateConfig.AdvertiseDefaultRoute {
		log.Printf("Update: v.VrfTemplateConfig.AdvertiseDefaultRoute=%v, c.VrfTemplateConfig.AdvertiseDefaultRoute=%v", v.VrfTemplateConfig.AdvertiseDefaultRoute, c.VrfTemplateConfig.AdvertiseDefaultRoute)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.VrfTemplateConfig.ConfigureStaticDefaultRoute != c.VrfTemplateConfig.ConfigureStaticDefaultRoute {
		log.Printf("Update: v.VrfTemplateConfig.ConfigureStaticDefaultRoute=%v, c.VrfTemplateConfig.ConfigureStaticDefaultRoute=%v", v.VrfTemplateConfig.ConfigureStaticDefaultRoute, c.VrfTemplateConfig.ConfigureStaticDefaultRoute)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.VrfTemplateConfig.BgpPassword != "" {
		if v.VrfTemplateConfig.BgpPassword != c.VrfTemplateConfig.BgpPassword {
			log.Printf("Update: v.VrfTemplateConfig.BgpPassword=%v, c.VrfTemplateConfig.BgpPassword=%v", v.VrfTemplateConfig.BgpPassword, c.VrfTemplateConfig.BgpPassword)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.VrfTemplateConfig.BgpPassword=%v, c.VrfTemplateConfig.BgpPassword=%v", v.VrfTemplateConfig.BgpPassword, c.VrfTemplateConfig.BgpPassword)
		v.VrfTemplateConfig.BgpPassword = c.VrfTemplateConfig.BgpPassword
	}

	if v.VrfTemplateConfig.BgpPasswordType != "" {
		if v.VrfTemplateConfig.BgpPasswordType != c.VrfTemplateConfig.BgpPasswordType {
			log.Printf("Update: v.VrfTemplateConfig.BgpPasswordType=%v, c.VrfTemplateConfig.BgpPasswordType=%v", v.VrfTemplateConfig.BgpPasswordType, c.VrfTemplateConfig.BgpPasswordType)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.VrfTemplateConfig.BgpPasswordType=%v, c.VrfTemplateConfig.BgpPasswordType=%v", v.VrfTemplateConfig.BgpPasswordType, c.VrfTemplateConfig.BgpPasswordType)
		v.VrfTemplateConfig.BgpPasswordType = c.VrfTemplateConfig.BgpPasswordType
	}

	if v.VrfTemplateConfig.Netflow != c.VrfTemplateConfig.Netflow {
		log.Printf("Update: v.VrfTemplateConfig.Netflow=%v, c.VrfTemplateConfig.Netflow=%v", v.VrfTemplateConfig.Netflow, c.VrfTemplateConfig.Netflow)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.VrfTemplateConfig.NetflowMonitor != "" {
		if v.VrfTemplateConfig.NetflowMonitor != c.VrfTemplateConfig.NetflowMonitor {
			log.Printf("Update: v.VrfTemplateConfig.NetflowMonitor=%v, c.VrfTemplateConfig.NetflowMonitor=%v", v.VrfTemplateConfig.NetflowMonitor, c.VrfTemplateConfig.NetflowMonitor)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.VrfTemplateConfig.NetflowMonitor=%v, c.VrfTemplateConfig.NetflowMonitor=%v", v.VrfTemplateConfig.NetflowMonitor, c.VrfTemplateConfig.NetflowMonitor)
		v.VrfTemplateConfig.NetflowMonitor = c.VrfTemplateConfig.NetflowMonitor
	}

	if v.VrfTemplateConfig.DisableRtAuto != c.VrfTemplateConfig.DisableRtAuto {
		log.Printf("Update: v.VrfTemplateConfig.DisableRtAuto=%v, c.VrfTemplateConfig.DisableRtAuto=%v", v.VrfTemplateConfig.DisableRtAuto, c.VrfTemplateConfig.DisableRtAuto)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.VrfTemplateConfig.RouteTargetImport != "" {
		if v.VrfTemplateConfig.RouteTargetImport != c.VrfTemplateConfig.RouteTargetImport {
			log.Printf("Update: v.VrfTemplateConfig.RouteTargetImport=%v, c.VrfTemplateConfig.RouteTargetImport=%v", v.VrfTemplateConfig.RouteTargetImport, c.VrfTemplateConfig.RouteTargetImport)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.VrfTemplateConfig.RouteTargetImport=%v, c.VrfTemplateConfig.RouteTargetImport=%v", v.VrfTemplateConfig.RouteTargetImport, c.VrfTemplateConfig.RouteTargetImport)
		v.VrfTemplateConfig.RouteTargetImport = c.VrfTemplateConfig.RouteTargetImport
	}

	if v.VrfTemplateConfig.RouteTargetExport != "" {
		if v.VrfTemplateConfig.RouteTargetExport != c.VrfTemplateConfig.RouteTargetExport {
			log.Printf("Update: v.VrfTemplateConfig.RouteTargetExport=%v, c.VrfTemplateConfig.RouteTargetExport=%v", v.VrfTemplateConfig.RouteTargetExport, c.VrfTemplateConfig.RouteTargetExport)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.VrfTemplateConfig.RouteTargetExport=%v, c.VrfTemplateConfig.RouteTargetExport=%v", v.VrfTemplateConfig.RouteTargetExport, c.VrfTemplateConfig.RouteTargetExport)
		v.VrfTemplateConfig.RouteTargetExport = c.VrfTemplateConfig.RouteTargetExport
	}

	if v.VrfTemplateConfig.RouteTargetImportEvpn != "" {
		if v.VrfTemplateConfig.RouteTargetImportEvpn != c.VrfTemplateConfig.RouteTargetImportEvpn {
			log.Printf("Update: v.VrfTemplateConfig.RouteTargetImportEvpn=%v, c.VrfTemplateConfig.RouteTargetImportEvpn=%v", v.VrfTemplateConfig.RouteTargetImportEvpn, c.VrfTemplateConfig.RouteTargetImportEvpn)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.VrfTemplateConfig.RouteTargetImportEvpn=%v, c.VrfTemplateConfig.RouteTargetImportEvpn=%v", v.VrfTemplateConfig.RouteTargetImportEvpn, c.VrfTemplateConfig.RouteTargetImportEvpn)
		v.VrfTemplateConfig.RouteTargetImportEvpn = c.VrfTemplateConfig.RouteTargetImportEvpn
	}

	if v.VrfTemplateConfig.RouteTargetExportEvpn != "" {
		if v.VrfTemplateConfig.RouteTargetExportEvpn != c.VrfTemplateConfig.RouteTargetExportEvpn {
			log.Printf("Update: v.VrfTemplateConfig.RouteTargetExportEvpn=%v, c.VrfTemplateConfig.RouteTargetExportEvpn=%v", v.VrfTemplateConfig.RouteTargetExportEvpn, c.VrfTemplateConfig.RouteTargetExportEvpn)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.VrfTemplateConfig.RouteTargetExportEvpn=%v, c.VrfTemplateConfig.RouteTargetExportEvpn=%v", v.VrfTemplateConfig.RouteTargetExportEvpn, c.VrfTemplateConfig.RouteTargetExportEvpn)
		v.VrfTemplateConfig.RouteTargetExportEvpn = c.VrfTemplateConfig.RouteTargetExportEvpn
	}

	if v.VrfTemplateConfig.RouteTargetImportMvpn != "" {
		if v.VrfTemplateConfig.RouteTargetImportMvpn != c.VrfTemplateConfig.RouteTargetImportMvpn {
			log.Printf("Update: v.VrfTemplateConfig.RouteTargetImportMvpn=%v, c.VrfTemplateConfig.RouteTargetImportMvpn=%v", v.VrfTemplateConfig.RouteTargetImportMvpn, c.VrfTemplateConfig.RouteTargetImportMvpn)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.VrfTemplateConfig.RouteTargetImportMvpn=%v, c.VrfTemplateConfig.RouteTargetImportMvpn=%v", v.VrfTemplateConfig.RouteTargetImportMvpn, c.VrfTemplateConfig.RouteTargetImportMvpn)
		v.VrfTemplateConfig.RouteTargetImportMvpn = c.VrfTemplateConfig.RouteTargetImportMvpn
	}

	if v.VrfTemplateConfig.RouteTargetExportMvpn != "" {
		if v.VrfTemplateConfig.RouteTargetExportMvpn != c.VrfTemplateConfig.RouteTargetExportMvpn {
			log.Printf("Update: v.VrfTemplateConfig.RouteTargetExportMvpn=%v, c.VrfTemplateConfig.RouteTargetExportMvpn=%v", v.VrfTemplateConfig.RouteTargetExportMvpn, c.VrfTemplateConfig.RouteTargetExportMvpn)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.VrfTemplateConfig.RouteTargetExportMvpn=%v, c.VrfTemplateConfig.RouteTargetExportMvpn=%v", v.VrfTemplateConfig.RouteTargetExportMvpn, c.VrfTemplateConfig.RouteTargetExportMvpn)
		v.VrfTemplateConfig.RouteTargetExportMvpn = c.VrfTemplateConfig.RouteTargetExportMvpn
	}

	if v.VrfTemplateConfig.RouteTargetImportCloudEvpn != "" {
		if v.VrfTemplateConfig.RouteTargetImportCloudEvpn != c.VrfTemplateConfig.RouteTargetImportCloudEvpn {
			log.Printf("Update: v.VrfTemplateConfig.RouteTargetImportCloudEvpn=%v, c.VrfTemplateConfig.RouteTargetImportCloudEvpn=%v", v.VrfTemplateConfig.RouteTargetImportCloudEvpn, c.VrfTemplateConfig.RouteTargetImportCloudEvpn)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.VrfTemplateConfig.RouteTargetImportCloudEvpn=%v, c.VrfTemplateConfig.RouteTargetImportCloudEvpn=%v", v.VrfTemplateConfig.RouteTargetImportCloudEvpn, c.VrfTemplateConfig.RouteTargetImportCloudEvpn)
		v.VrfTemplateConfig.RouteTargetImportCloudEvpn = c.VrfTemplateConfig.RouteTargetImportCloudEvpn
	}

	if v.VrfTemplateConfig.RouteTargetExportCloudEvpn != "" {
		if v.VrfTemplateConfig.RouteTargetExportCloudEvpn != c.VrfTemplateConfig.RouteTargetExportCloudEvpn {
			log.Printf("Update: v.VrfTemplateConfig.RouteTargetExportCloudEvpn=%v, c.VrfTemplateConfig.RouteTargetExportCloudEvpn=%v", v.VrfTemplateConfig.RouteTargetExportCloudEvpn, c.VrfTemplateConfig.RouteTargetExportCloudEvpn)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.VrfTemplateConfig.RouteTargetExportCloudEvpn=%v, c.VrfTemplateConfig.RouteTargetExportCloudEvpn=%v", v.VrfTemplateConfig.RouteTargetExportCloudEvpn, c.VrfTemplateConfig.RouteTargetExportCloudEvpn)
		v.VrfTemplateConfig.RouteTargetExportCloudEvpn = c.VrfTemplateConfig.RouteTargetExportCloudEvpn
	}

	if v.DeployAttachments != c.DeployAttachments {
		log.Printf("Update: v.DeployAttachments=%v, c.DeployAttachments=%v", v.DeployAttachments, c.DeployAttachments)
		*cf = true
	}

	return action
}
