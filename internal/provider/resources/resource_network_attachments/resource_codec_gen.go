// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package resource_network_attachments

import (
	. "terraform-provider-ndfc/internal/provider/types"
)

type NDFCNetworkAttachmentsModel struct {
	NetworkAttachments map[string]NDFCNetworkAttachmentsValue `json:"networkAttachments,omitempty"`
}

type NDFCNetworkAttachmentsValue struct {
	NetworkName string                          `json:"networkName,omitempty"`
	Attachments map[string]NDFCAttachmentsValue `json:"lanAttachList,omitempty"`
}

type NDFCAttachmentsValue struct {
	FilterThisValue      bool         `json:"-"`
	Id                   *int64       `json:"-"`
	FabricName           string       `json:"fabric,omitempty"`
	NetworkName          string       `json:"networkName,omitempty"`
	SerialNumber         string       `json:"serialNumber,omitempty"`
	SwitchSerialNo       string       `json:"switchSerialNo,omitempty"`
	SwitchName           string       `json:"switchName,omitempty"`
	DisplayName          string       `json:"displayName,omitempty"`
	Vlan                 *Int64Custom `json:"vlan,omitempty"`
	VlanId               *Int64Custom `json:"vlanId,omitempty"`
	Deployment           string       `json:"deployment,omitempty"`
	AttachState          string       `json:"lanAttachState,omitempty"`
	Attached             *bool        `json:"isLanAttached,omitempty"`
	FreeformConfig       string       `json:"freeformconfig,omitempty"`
	DeployThisAttachment bool         `json:"-"`
	SwitchPorts          CSVString    `json:"switchPorts,omitempty"`
	DetachSwitchPorts    CSVString    `json:"detachSwitchPorts,omitempty"`
	PortNames            string       `json:"portNames,omitempty"`
	TorPorts             CSVString    `json:"torPorts,omitempty"`
	InstanceValues       string       `json:"instanceValues,omitempty"`
	UpdateAction         uint16       `json:"-"`
}
