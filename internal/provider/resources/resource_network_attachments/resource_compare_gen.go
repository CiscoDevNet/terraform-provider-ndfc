// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package resource_network_attachments

import (
	"log"
	. "terraform-provider-ndfc/internal/provider/types"
)

func (v NDFCAttachmentsValue) DeepEqual(c NDFCAttachmentsValue) int {
	cf := false
	if v.DisplayName != c.DisplayName {
		log.Printf("v.DisplayName=%v, c.DisplayName=%v", v.DisplayName, c.DisplayName)
		return RequiresUpdate
	}

	if !v.Vlan.IsEmpty() && !c.Vlan.IsEmpty() {
		if *v.Vlan != *c.Vlan {
			log.Printf("v.Vlan=%v, c.Vlan=%v", *v.Vlan, *c.Vlan)
			return RequiresUpdate
		}
	} else {
		if !v.Vlan.IsEmpty() {
			log.Printf("v.Vlan=%v", *v.Vlan)
			return RequiresUpdate
		} else if !c.Vlan.IsEmpty() {
			log.Printf("c.Vlan=%v", *c.Vlan)
			return RequiresUpdate
		}
	}
	if v.FreeformConfig != c.FreeformConfig {
		log.Printf("v.FreeformConfig=%v, c.FreeformConfig=%v", v.FreeformConfig, c.FreeformConfig)
		return RequiresUpdate
	}
	if v.DeployThisAttachment != c.DeployThisAttachment {
		log.Printf("v.DeployThisAttachment=%v, c.DeployThisAttachment=%v", v.DeployThisAttachment, c.DeployThisAttachment)
		cf = true
	}

	if len(v.SwitchPorts) != len(c.SwitchPorts) {
		log.Printf("len(v.SwitchPorts)=%d, len(c.SwitchPorts)=%d", len(v.SwitchPorts), len(c.SwitchPorts))
		return PortListUpdate
	}
	for i := range v.SwitchPorts {
		if v.SwitchPorts[i] != c.SwitchPorts[i] {
			log.Printf("v.SwitchPorts[%d]=%s, c.SwitchPorts[%d]=%s", i, v.SwitchPorts[i], i, c.SwitchPorts[i])
			return PortListUpdate
		}
	}

	if len(v.TorPorts) != len(c.TorPorts) {
		log.Printf("len(v.TorPorts)=%d, len(c.TorPorts)=%d", len(v.TorPorts), len(c.TorPorts))
		return PortListUpdate
	}
	for i := range v.TorPorts {
		if v.TorPorts[i] != c.TorPorts[i] {
			log.Printf("v.TorPorts[%d]=%s, c.TorPorts[%d]=%s", i, v.TorPorts[i], i, c.TorPorts[i])
			return PortListUpdate
		}
	}
	if v.InstanceValues != c.InstanceValues {
		log.Printf("v.InstanceValues=%v, c.InstanceValues=%v", v.InstanceValues, c.InstanceValues)
		return RequiresUpdate
	}

	if cf {
		return ControlFlagUpdate
	}
	return ValuesDeeplyEqual
}

func (v NDFCNetworkAttachmentsValue) DeepEqual(c NDFCNetworkAttachmentsValue) int {
	cf := false
	if v.NetworkName != c.NetworkName {
		log.Printf("v.NetworkName=%v, c.NetworkName=%v", v.NetworkName, c.NetworkName)
		return RequiresReplace
	}

	if cf {
		return ControlFlagUpdate
	}
	return ValuesDeeplyEqual
}

func (v *NDFCAttachmentsValue) CreatePlan(c NDFCAttachmentsValue, cf *bool) int {
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

	if !v.Vlan.IsEmpty() && !c.Vlan.IsEmpty() {
		if *v.Vlan != *c.Vlan {
			log.Printf("Update: v.Vlan=%v, c.Vlan=%v", *v.Vlan, *c.Vlan)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		if !v.Vlan.IsEmpty() {
			log.Printf("Update: v.Vlan=%v", *v.Vlan)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		} else if !c.Vlan.IsEmpty() {
			log.Printf("Copy from State: c.Vlan=%v", *c.Vlan)
			v.Vlan = new(Int64Custom)
			*v.Vlan = *c.Vlan
		}
	}
	if v.FreeformConfig != "" {

		if v.FreeformConfig != c.FreeformConfig {
			log.Printf("Update: v.FreeformConfig=%v, c.FreeformConfig=%v", v.FreeformConfig, c.FreeformConfig)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}

	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.FreeformConfig=%v, c.FreeformConfig=%v", v.FreeformConfig, c.FreeformConfig)
		v.FreeformConfig = c.FreeformConfig
	}

	if v.DeployThisAttachment != c.DeployThisAttachment {
		log.Printf("Update: v.DeployThisAttachment=%v, c.DeployThisAttachment=%v", v.DeployThisAttachment, c.DeployThisAttachment)
		*cf = true
	}

	if len(v.SwitchPorts) != len(c.SwitchPorts) {
		log.Printf("Update: len(v.SwitchPorts)=%d, len(c.SwitchPorts)=%d", len(v.SwitchPorts), len(c.SwitchPorts))
		return RequiresUpdate
	}
	for i := range v.SwitchPorts {
		if v.SwitchPorts[i] != c.SwitchPorts[i] {
			log.Printf("Update: v.SwitchPorts[%d]=%s, c.SwitchPorts[%d]=%s", i, v.SwitchPorts[i], i, c.SwitchPorts[i])
			return RequiresUpdate
		}
	}

	if len(v.TorPorts) != len(c.TorPorts) {
		log.Printf("Update: len(v.TorPorts)=%d, len(c.TorPorts)=%d", len(v.TorPorts), len(c.TorPorts))
		return RequiresUpdate
	}
	for i := range v.TorPorts {
		if v.TorPorts[i] != c.TorPorts[i] {
			log.Printf("Update: v.TorPorts[%d]=%s, c.TorPorts[%d]=%s", i, v.TorPorts[i], i, c.TorPorts[i])
			return RequiresUpdate
		}
	}

	if v.InstanceValues != "" {

		if v.InstanceValues != c.InstanceValues {
			log.Printf("Update: v.InstanceValues=%v, c.InstanceValues=%v", v.InstanceValues, c.InstanceValues)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}

	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.InstanceValues=%v, c.InstanceValues=%v", v.InstanceValues, c.InstanceValues)
		v.InstanceValues = c.InstanceValues
	}

	return action
}

func (v *NDFCNetworkAttachmentsValue) CreatePlan(c NDFCNetworkAttachmentsValue, cf *bool) int {
	action := ActionNone

	if v.NetworkName != "" {

		if v.NetworkName != c.NetworkName {
			log.Printf("Update: v.NetworkName=%v, c.NetworkName=%v", v.NetworkName, c.NetworkName)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresReplace
			}
		}

	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.NetworkName=%v, c.NetworkName=%v", v.NetworkName, c.NetworkName)
		v.NetworkName = c.NetworkName
	}

	return action
}
