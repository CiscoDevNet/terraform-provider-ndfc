// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package resource_msd_vrfs_parent

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
	if v.VrfTemplateConfig.Trm != "" {
		if v.VrfTemplateConfig.Trm != c.VrfTemplateConfig.Trm {
			log.Printf("v.VrfTemplateConfig.Trm=%s, c.VrfTemplateConfig.Trm=%s", v.VrfTemplateConfig.Trm, c.VrfTemplateConfig.Trm)
			return RequiresUpdate
		}
	} else {
		log.Printf("Skipping - v.VrfTemplateConfig.Trm=%s, c.VrfTemplateConfig.Trm=%s", v.VrfTemplateConfig.Trm, c.VrfTemplateConfig.Trm)
	}
	if v.VrfTemplateConfig.RpExternal != "" {
		if v.VrfTemplateConfig.RpExternal != c.VrfTemplateConfig.RpExternal {
			log.Printf("v.VrfTemplateConfig.RpExternal=%s, c.VrfTemplateConfig.RpExternal=%s", v.VrfTemplateConfig.RpExternal, c.VrfTemplateConfig.RpExternal)
			return RequiresUpdate
		}
	} else {
		log.Printf("Skipping - v.VrfTemplateConfig.RpExternal=%s, c.VrfTemplateConfig.RpExternal=%s", v.VrfTemplateConfig.RpExternal, c.VrfTemplateConfig.RpExternal)
	}
	if v.VrfTemplateConfig.MvpnInterAs != "" {
		if v.VrfTemplateConfig.MvpnInterAs != c.VrfTemplateConfig.MvpnInterAs {
			log.Printf("v.VrfTemplateConfig.MvpnInterAs=%s, c.VrfTemplateConfig.MvpnInterAs=%s", v.VrfTemplateConfig.MvpnInterAs, c.VrfTemplateConfig.MvpnInterAs)
			return RequiresUpdate
		}
	} else {
		log.Printf("Skipping - v.VrfTemplateConfig.MvpnInterAs=%s, c.VrfTemplateConfig.MvpnInterAs=%s", v.VrfTemplateConfig.MvpnInterAs, c.VrfTemplateConfig.MvpnInterAs)
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
	if v.VrfTemplateConfig.RouteTargetImportCloudEvpn != c.VrfTemplateConfig.RouteTargetImportCloudEvpn {
		log.Printf("v.VrfTemplateConfig.RouteTargetImportCloudEvpn=%s, c.VrfTemplateConfig.RouteTargetImportCloudEvpn=%s", v.VrfTemplateConfig.RouteTargetImportCloudEvpn, c.VrfTemplateConfig.RouteTargetImportCloudEvpn)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.RouteTargetExportCloudEvpn != c.VrfTemplateConfig.RouteTargetExportCloudEvpn {
		log.Printf("v.VrfTemplateConfig.RouteTargetExportCloudEvpn=%s, c.VrfTemplateConfig.RouteTargetExportCloudEvpn=%s", v.VrfTemplateConfig.RouteTargetExportCloudEvpn, c.VrfTemplateConfig.RouteTargetExportCloudEvpn)
		return RequiresUpdate
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

	if v.VrfTemplateConfig.RpExternal != c.VrfTemplateConfig.RpExternal {
		log.Printf("Update: v.VrfTemplateConfig.RpExternal=%v, c.VrfTemplateConfig.RpExternal=%v", v.VrfTemplateConfig.RpExternal, c.VrfTemplateConfig.RpExternal)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.VrfTemplateConfig.MvpnInterAs != c.VrfTemplateConfig.MvpnInterAs {
		log.Printf("Update: v.VrfTemplateConfig.MvpnInterAs=%v, c.VrfTemplateConfig.MvpnInterAs=%v", v.VrfTemplateConfig.MvpnInterAs, c.VrfTemplateConfig.MvpnInterAs)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
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

	return action
}
