// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package resource_inventory_devices

import (
	"log"
	. "terraform-provider-ndfc/internal/provider/types"
)

func (v NDFCDevicesValue) DeepEqual(c NDFCDevicesValue) int {
	cf := false
	if v.Role != c.Role {
		log.Printf("v.Role=%v, c.Role=%v", v.Role, c.Role)
		return RequiresUpdate
	}
	if v.DiscoveryType != c.DiscoveryType {
		log.Printf("v.DiscoveryType=%v, c.DiscoveryType=%v", v.DiscoveryType, c.DiscoveryType)
		return RequiresUpdate
	}
	if v.DiscoveryUsername != c.DiscoveryUsername {
		log.Printf("v.DiscoveryUsername=%v, c.DiscoveryUsername=%v", v.DiscoveryUsername, c.DiscoveryUsername)
		return RequiresUpdate
	}
	if v.DiscoveryPassword != c.DiscoveryPassword {
		log.Printf("v.DiscoveryPassword=%v, c.DiscoveryPassword=%v", v.DiscoveryPassword, c.DiscoveryPassword)
		return RequiresUpdate
	}
	if v.DiscoveryAuthProtocol != c.DiscoveryAuthProtocol {
		log.Printf("v.DiscoveryAuthProtocol=%v, c.DiscoveryAuthProtocol=%v", v.DiscoveryAuthProtocol, c.DiscoveryAuthProtocol)
		return RequiresUpdate
	}
	if v.SerialNumber != c.SerialNumber {
		log.Printf("v.SerialNumber=%v, c.SerialNumber=%v", v.SerialNumber, c.SerialNumber)
		return RequiresUpdate
	}
	if v.Model != c.Model {
		log.Printf("v.Model=%v, c.Model=%v", v.Model, c.Model)
		return RequiresUpdate
	}
	if v.Version != c.Version {
		log.Printf("v.Version=%v, c.Version=%v", v.Version, c.Version)
		return RequiresUpdate
	}
	if v.Hostname != c.Hostname {
		log.Printf("v.Hostname=%v, c.Hostname=%v", v.Hostname, c.Hostname)
		return RequiresUpdate
	}
	if v.ImagePolicy != c.ImagePolicy {
		log.Printf("v.ImagePolicy=%v, c.ImagePolicy=%v", v.ImagePolicy, c.ImagePolicy)
		return RequiresUpdate
	}
	if v.Gateway != c.Gateway {
		log.Printf("v.Gateway=%v, c.Gateway=%v", v.Gateway, c.Gateway)
		return RequiresUpdate
	}

	if len(v.ModulesModel) != len(c.ModulesModel) {
		log.Printf("len(v.ModulesModel)=%d, len(c.ModulesModel)=%d", len(v.ModulesModel), len(c.ModulesModel))
		return PortListUpdate
	}
	for i := range v.ModulesModel {
		if v.ModulesModel[i] != c.ModulesModel[i] {
			log.Printf("v.ModulesModel[%d]=%s, c.ModulesModel[%d]=%s", i, v.ModulesModel[i], i, c.ModulesModel[i])
			return PortListUpdate
		}
	}
	if v.Breakout != c.Breakout {
		log.Printf("v.Breakout=%v, c.Breakout=%v", v.Breakout, c.Breakout)
		return RequiresUpdate
	}
	if v.PortMode != c.PortMode {
		log.Printf("v.PortMode=%v, c.PortMode=%v", v.PortMode, c.PortMode)
		return RequiresUpdate
	}

	if cf {
		return ControlFlagUpdate
	}
	return ValuesDeeplyEqual
}

func (v *NDFCDevicesValue) CreatePlan(c NDFCDevicesValue, cf *bool) int {
	action := ActionNone

	if v.Role != c.Role {
		log.Printf("Role-Update: v.Role=%v, c.Role=%v", v.Role, c.Role)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.DiscoveryType != c.DiscoveryType {
		log.Printf("DiscoveryType-Update: v.DiscoveryType=%v, c.DiscoveryType=%v", v.DiscoveryType, c.DiscoveryType)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.DiscoveryUsername != c.DiscoveryUsername {
		log.Printf("DiscoveryUsername-Update: v.DiscoveryUsername=%v, c.DiscoveryUsername=%v", v.DiscoveryUsername, c.DiscoveryUsername)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.DiscoveryPassword != c.DiscoveryPassword {
		log.Printf("DiscoveryPassword-Update: v.DiscoveryPassword=%v, c.DiscoveryPassword=%v", v.DiscoveryPassword, c.DiscoveryPassword)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.DiscoveryAuthProtocol != c.DiscoveryAuthProtocol {
		log.Printf("DiscoveryAuthProtocol-Update: v.DiscoveryAuthProtocol=%v, c.DiscoveryAuthProtocol=%v", v.DiscoveryAuthProtocol, c.DiscoveryAuthProtocol)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.SerialNumber != c.SerialNumber {
		log.Printf("SerialNumber-Update: v.SerialNumber=%v, c.SerialNumber=%v", v.SerialNumber, c.SerialNumber)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.Model != c.Model {
		log.Printf("Model-Update: v.Model=%v, c.Model=%v", v.Model, c.Model)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.Version != c.Version {
		log.Printf("Version-Update: v.Version=%v, c.Version=%v", v.Version, c.Version)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.Hostname != c.Hostname {
		log.Printf("Hostname-Update: v.Hostname=%v, c.Hostname=%v", v.Hostname, c.Hostname)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.ImagePolicy != c.ImagePolicy {
		log.Printf("ImagePolicy-Update: v.ImagePolicy=%v, c.ImagePolicy=%v", v.ImagePolicy, c.ImagePolicy)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.Gateway != c.Gateway {
		log.Printf("Gateway-Update: v.Gateway=%v, c.Gateway=%v", v.Gateway, c.Gateway)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if len(v.ModulesModel) != len(c.ModulesModel) {
		log.Printf("Update: len(v.ModulesModel)=%d, len(c.ModulesModel)=%d", len(v.ModulesModel), len(c.ModulesModel))
		return RequiresUpdate
	}
	for i := range v.ModulesModel {
		if v.ModulesModel[i] != c.ModulesModel[i] {
			log.Printf("Update: v.ModulesModel[%d]=%s, c.ModulesModel[%d]=%s", i, v.ModulesModel[i], i, c.ModulesModel[i])
			return RequiresUpdate
		}
	}

	if v.Breakout != c.Breakout {
		log.Printf("Breakout-Update: v.Breakout=%v, c.Breakout=%v", v.Breakout, c.Breakout)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.PortMode != c.PortMode {
		log.Printf("PortMode-Update: v.PortMode=%v, c.PortMode=%v", v.PortMode, c.PortMode)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	return action
}
