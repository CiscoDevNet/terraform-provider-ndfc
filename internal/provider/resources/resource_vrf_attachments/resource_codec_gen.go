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
	. "terraform-provider-ndfc/internal/provider/types"
)

type NDFCVrfAttachmentsModel struct {
	FabricName           string                             `json:"fabric,omitempty"`
	DeployAllAttachments bool                               `json:"-"`
	VrfAttachments       map[string]NDFCVrfAttachmentsValue `json:"attachments,omitempty"`
}

type NDFCVrfAttachmentsValue struct {
	Id                   *int64                         `json:"-"`
	FilterThisValue      bool                           `json:"-"`
	VrfName              string                         `json:"vrfName,omitempty"`
	DeployAllAttachments bool                           `json:"-"`
	AttachList           map[string]NDFCAttachListValue `json:"lanAttachList,omitempty"`
}

type NDFCAttachListValue struct {
	FilterThisValue      bool                    `json:"-"`
	Id                   *int64                  `json:"-"`
	FabricName           string                  `json:"fabric,omitempty"`
	VrfName              string                  `json:"vrfName,omitempty"`
	SerialNumber         string                  `json:"serialNumber,omitempty"`
	SwitchSerialNo       string                  `json:"switchSerialNo,omitempty"`
	SwitchName           string                  `json:"switchName,omitempty"`
	Vlan                 *Int64Custom            `json:"vlan,omitempty"`
	VlanId               *Int64Custom            `json:"vlanId,omitempty"`
	Deployment           string                  `json:"deployment,omitempty"`
	AttachState          string                  `json:"lanAttachState,omitempty"`
	Attached             *bool                   `json:"isLanAttached,omitempty"`
	FreeformConfig       string                  `json:"freeformconfig,omitempty"`
	DeployThisAttachment bool                    `json:"-"`
	UpdateAction         uint16                  `json:"-"`
	InstanceValues       NDFCInstanceValuesValue `json:"instanceValues,omitempty"`
}

type NDFCInstanceValuesValue struct {
	LoopbackId   *Int64Custom `json:"loopbackId,omitempty"`
	LoopbackIpv4 string       `json:"loopbackIpAddress,omitempty"`
	LoopbackIpv6 string       `json:"loopbackIpv6Address,omitempty"`
}
