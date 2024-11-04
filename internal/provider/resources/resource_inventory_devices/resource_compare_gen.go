// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
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

	if v.Role != "" {

		if v.Role != c.Role {
			log.Printf("Update: v.Role=%v, c.Role=%v", v.Role, c.Role)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}

	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.Role=%v, c.Role=%v", v.Role, c.Role)
		v.Role = c.Role
	}

	if v.DiscoveryType != "" {

		if v.DiscoveryType != c.DiscoveryType {
			log.Printf("Update: v.DiscoveryType=%v, c.DiscoveryType=%v", v.DiscoveryType, c.DiscoveryType)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}

	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.DiscoveryType=%v, c.DiscoveryType=%v", v.DiscoveryType, c.DiscoveryType)
		v.DiscoveryType = c.DiscoveryType
	}

	if v.DiscoveryUsername != "" {

		if v.DiscoveryUsername != c.DiscoveryUsername {
			log.Printf("Update: v.DiscoveryUsername=%v, c.DiscoveryUsername=%v", v.DiscoveryUsername, c.DiscoveryUsername)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}

	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.DiscoveryUsername=%v, c.DiscoveryUsername=%v", v.DiscoveryUsername, c.DiscoveryUsername)
		v.DiscoveryUsername = c.DiscoveryUsername
	}

	if v.DiscoveryPassword != "" {

		if v.DiscoveryPassword != c.DiscoveryPassword {
			log.Printf("Update: v.DiscoveryPassword=%v, c.DiscoveryPassword=%v", v.DiscoveryPassword, c.DiscoveryPassword)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}

	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.DiscoveryPassword=%v, c.DiscoveryPassword=%v", v.DiscoveryPassword, c.DiscoveryPassword)
		v.DiscoveryPassword = c.DiscoveryPassword
	}

	if v.DiscoveryAuthProtocol != "" {

		if v.DiscoveryAuthProtocol != c.DiscoveryAuthProtocol {
			log.Printf("Update: v.DiscoveryAuthProtocol=%v, c.DiscoveryAuthProtocol=%v", v.DiscoveryAuthProtocol, c.DiscoveryAuthProtocol)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}

	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.DiscoveryAuthProtocol=%v, c.DiscoveryAuthProtocol=%v", v.DiscoveryAuthProtocol, c.DiscoveryAuthProtocol)
		v.DiscoveryAuthProtocol = c.DiscoveryAuthProtocol
	}

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

	if v.Model != "" {

		if v.Model != c.Model {
			log.Printf("Update: v.Model=%v, c.Model=%v", v.Model, c.Model)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}

	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.Model=%v, c.Model=%v", v.Model, c.Model)
		v.Model = c.Model
	}

	if v.Version != "" {

		if v.Version != c.Version {
			log.Printf("Update: v.Version=%v, c.Version=%v", v.Version, c.Version)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}

	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.Version=%v, c.Version=%v", v.Version, c.Version)
		v.Version = c.Version
	}

	if v.Hostname != "" {

		if v.Hostname != c.Hostname {
			log.Printf("Update: v.Hostname=%v, c.Hostname=%v", v.Hostname, c.Hostname)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}

	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.Hostname=%v, c.Hostname=%v", v.Hostname, c.Hostname)
		v.Hostname = c.Hostname
	}

	if v.ImagePolicy != "" {

		if v.ImagePolicy != c.ImagePolicy {
			log.Printf("Update: v.ImagePolicy=%v, c.ImagePolicy=%v", v.ImagePolicy, c.ImagePolicy)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}

	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.ImagePolicy=%v, c.ImagePolicy=%v", v.ImagePolicy, c.ImagePolicy)
		v.ImagePolicy = c.ImagePolicy
	}

	if v.Gateway != "" {

		if v.Gateway != c.Gateway {
			log.Printf("Update: v.Gateway=%v, c.Gateway=%v", v.Gateway, c.Gateway)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}

	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.Gateway=%v, c.Gateway=%v", v.Gateway, c.Gateway)
		v.Gateway = c.Gateway
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

	if v.Breakout != "" {

		if v.Breakout != c.Breakout {
			log.Printf("Update: v.Breakout=%v, c.Breakout=%v", v.Breakout, c.Breakout)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}

	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.Breakout=%v, c.Breakout=%v", v.Breakout, c.Breakout)
		v.Breakout = c.Breakout
	}

	if v.PortMode != "" {

		if v.PortMode != c.PortMode {
			log.Printf("Update: v.PortMode=%v, c.PortMode=%v", v.PortMode, c.PortMode)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}

	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.PortMode=%v, c.PortMode=%v", v.PortMode, c.PortMode)
		v.PortMode = c.PortMode
	}

	return action
}
