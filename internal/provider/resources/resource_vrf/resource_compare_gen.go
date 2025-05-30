// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_vrf

import "log"

const (
	ValuesDeeplyEqual = iota
	RequiresReplace
	RequiresUpdate
)

func (v NDFCAttachmentsValue) DeepEqual(c NDFCAttachmentsValue) int {
	if v.SerialNumber != c.SerialNumber {
		log.Printf("v.SerialNumber=%v, c.SerialNumber=%v", v.SerialNumber, c.SerialNumber)
		return RequiresUpdate
	}

	if v.VlanId != nil && c.VlanId != nil {
		if *v.VlanId != *c.VlanId {
			log.Printf("v.VlanId=%v, c.VlanId=%v", *v.VlanId, *c.VlanId)
			return RequiresUpdate
		}
	} else {
		if v.VlanId != nil {
			log.Printf("v.VlanId=%v", *v.VlanId)
			return RequiresUpdate
		} else if c.VlanId != nil {
			log.Printf("c.VlanId=%v", *c.VlanId)
			return RequiresUpdate
		}
	}
	if v.FreeformConfig != c.FreeformConfig {
		log.Printf("v.FreeformConfig=%v, c.FreeformConfig=%v", v.FreeformConfig, c.FreeformConfig)
		return RequiresUpdate
	}

	if v.LoopbackId != nil && c.LoopbackId != nil {
		if *v.LoopbackId != *c.LoopbackId {
			log.Printf("v.LoopbackId=%v, c.LoopbackId=%v", *v.LoopbackId, *c.LoopbackId)
			return RequiresUpdate
		}
	} else {
		if v.LoopbackId != nil {
			log.Printf("v.LoopbackId=%v", *v.LoopbackId)
			return RequiresUpdate
		} else if c.LoopbackId != nil {
			log.Printf("c.LoopbackId=%v", *c.LoopbackId)
			return RequiresUpdate
		}
	}
	if v.LoopbackIpv4 != c.LoopbackIpv4 {
		log.Printf("v.LoopbackIpv4=%v, c.LoopbackIpv4=%v", v.LoopbackIpv4, c.LoopbackIpv4)
		return RequiresUpdate
	}
	if v.LoopbackIpv6 != c.LoopbackIpv6 {
		log.Printf("v.LoopbackIpv6=%v, c.LoopbackIpv6=%v", v.LoopbackIpv6, c.LoopbackIpv6)
		return RequiresUpdate
	}

	return ValuesDeeplyEqual
}
