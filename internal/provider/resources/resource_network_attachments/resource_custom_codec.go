// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_network_attachments

type NDFCNetworkAttachmentsPayload struct {
	NetworkName string                 `json:"networkName,omitempty"`
	Attachments []NDFCAttachmentsValue `json:"lanAttachList,omitempty"`
}

type NDFCNetworkAttachments struct {
	GlobalDeploy       bool
	GlobalUndeploy     bool
	FabricName         string
	NetworkAttachments []NDFCNetworkAttachmentsPayload // Attachment payload for NDFC
	DepMap             map[string][]string             // use for backfilling DeploymentFlag in TF state        // for deployment
}

func (p *NDFCNetworkAttachments) AddEntry(nwName string, attachList []NDFCAttachmentsValue) {
	if len(attachList) == 0 {
		return
	}
	nwAttachEntry := NDFCNetworkAttachmentsPayload{NetworkName: nwName}
	nwAttachEntry.Attachments = attachList
	p.NetworkAttachments = append(p.NetworkAttachments, nwAttachEntry)
}
