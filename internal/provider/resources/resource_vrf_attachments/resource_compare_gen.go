// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package resource_vrf_attachments

import (
	"log"
	. "terraform-provider-ndfc/internal/provider/types"
)

func (v NDFCAttachListValue) DeepEqual(c NDFCAttachListValue) int {
	cf := false

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

	if !v.InstanceValues.LoopbackId.IsEmpty() && !c.InstanceValues.LoopbackId.IsEmpty() {
		if *v.InstanceValues.LoopbackId != *c.InstanceValues.LoopbackId {
			log.Printf("v.InstanceValues.LoopbackId=%v, c.InstanceValues.LoopbackId=%v", *v.InstanceValues.LoopbackId, *c.InstanceValues.LoopbackId)
			return RequiresUpdate
		}
	} else {
		if !v.InstanceValues.LoopbackId.IsEmpty() {
			log.Printf("v.InstanceValues.LoopbackId=%v", *v.InstanceValues.LoopbackId)
			return RequiresUpdate
		} else if !c.InstanceValues.LoopbackId.IsEmpty() {
			log.Printf("c.InstanceValues.LoopbackId=%v", *c.InstanceValues.LoopbackId)
			return RequiresUpdate
		}
	}
	if v.InstanceValues.LoopbackIpv4 != c.InstanceValues.LoopbackIpv4 {
		log.Printf("v.InstanceValues.LoopbackIpv4=%s, c.InstanceValues.LoopbackIpv4=%s", v.InstanceValues.LoopbackIpv4, c.InstanceValues.LoopbackIpv4)
		return RequiresUpdate
	}
	if v.InstanceValues.LoopbackIpv6 != c.InstanceValues.LoopbackIpv6 {
		log.Printf("v.InstanceValues.LoopbackIpv6=%s, c.InstanceValues.LoopbackIpv6=%s", v.InstanceValues.LoopbackIpv6, c.InstanceValues.LoopbackIpv6)
		return RequiresUpdate
	}

	if cf {
		return ControlFlagUpdate
	}
	return ValuesDeeplyEqual
}

func (v NDFCVrfAttachmentsValue) DeepEqual(c NDFCVrfAttachmentsValue) int {
	cf := false
	if v.VrfName != c.VrfName {
		log.Printf("v.VrfName=%v, c.VrfName=%v", v.VrfName, c.VrfName)
		return RequiresReplace
	}
	if v.DeployAllAttachments != c.DeployAllAttachments {
		log.Printf("v.DeployAllAttachments=%v, c.DeployAllAttachments=%v", v.DeployAllAttachments, c.DeployAllAttachments)
		cf = true
	}

	if cf {
		return ControlFlagUpdate
	}
	return ValuesDeeplyEqual
}

func (v *NDFCAttachListValue) CreatePlan(c NDFCAttachListValue, cf *bool) int {
	action := ActionNone

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

	if !v.InstanceValues.LoopbackId.IsEmpty() && !c.InstanceValues.LoopbackId.IsEmpty() {
		if *v.InstanceValues.LoopbackId != *c.InstanceValues.LoopbackId {
			log.Printf("Update: v.InstanceValues.LoopbackId=%v, c.InstanceValues.LoopbackId=%v", *v.InstanceValues.LoopbackId, *c.InstanceValues.LoopbackId)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else if !v.InstanceValues.LoopbackId.IsEmpty() {
		log.Printf("Update: v.InstanceValues.LoopbackId=%v", *v.InstanceValues.LoopbackId)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if !c.InstanceValues.LoopbackId.IsEmpty() {
		log.Printf("Copy from State: c.InstanceValues.LoopbackId=%v", *c.InstanceValues.LoopbackId)
		v.InstanceValues.LoopbackId = new(Int64Custom)
		*v.InstanceValues.LoopbackId = *c.InstanceValues.LoopbackId
	}

	if v.InstanceValues.LoopbackIpv4 != "" {
		if v.InstanceValues.LoopbackIpv4 != c.InstanceValues.LoopbackIpv4 {
			log.Printf("Update: v.InstanceValues.LoopbackIpv4=%v, c.InstanceValues.LoopbackIpv4=%v", v.InstanceValues.LoopbackIpv4, c.InstanceValues.LoopbackIpv4)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.InstanceValues.LoopbackIpv4=%v, c.InstanceValues.LoopbackIpv4=%v", v.InstanceValues.LoopbackIpv4, c.InstanceValues.LoopbackIpv4)
		v.InstanceValues.LoopbackIpv4 = c.InstanceValues.LoopbackIpv4
	}

	if v.InstanceValues.LoopbackIpv6 != "" {
		if v.InstanceValues.LoopbackIpv6 != c.InstanceValues.LoopbackIpv6 {
			log.Printf("Update: v.InstanceValues.LoopbackIpv6=%v, c.InstanceValues.LoopbackIpv6=%v", v.InstanceValues.LoopbackIpv6, c.InstanceValues.LoopbackIpv6)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.InstanceValues.LoopbackIpv6=%v, c.InstanceValues.LoopbackIpv6=%v", v.InstanceValues.LoopbackIpv6, c.InstanceValues.LoopbackIpv6)
		v.InstanceValues.LoopbackIpv6 = c.InstanceValues.LoopbackIpv6
	}

	return action
}

func (v *NDFCVrfAttachmentsValue) CreatePlan(c NDFCVrfAttachmentsValue, cf *bool) int {
	action := ActionNone

	if v.VrfName != "" {

		if v.VrfName != c.VrfName {
			log.Printf("Update: v.VrfName=%v, c.VrfName=%v", v.VrfName, c.VrfName)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresReplace
			}
		}

	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.VrfName=%v, c.VrfName=%v", v.VrfName, c.VrfName)
		v.VrfName = c.VrfName
	}

	if v.DeployAllAttachments != c.DeployAllAttachments {
		log.Printf("Update: v.DeployAllAttachments=%v, c.DeployAllAttachments=%v", v.DeployAllAttachments, c.DeployAllAttachments)
		*cf = true
	}

	return action
}
