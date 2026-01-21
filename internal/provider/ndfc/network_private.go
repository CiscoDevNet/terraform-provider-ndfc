// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package ndfc

import (
	"context"
	"fmt"
	"net"
	rna "terraform-provider-ndfc/internal/provider/resources/resource_network_attachments"
	"terraform-provider-ndfc/internal/provider/resources/resource_networks"
)

func (c NDFC) netAttachmentSerialRemap(ctx context.Context, nw *resource_networks.NDFCNetworksModel) map[string]string {
	keyMap := make(map[string]string)
	for nwName, nwEntry := range nw.Networks {
		for attKey, attachEntry := range nwEntry.Attachments {
			// check if key is IP or Serial
			if net.ParseIP(attKey) != nil {
				// IP
				serial := c.GetSerialFromIP(ctx, nw.FabricName, attKey)
				if serial == "" {
					panic(fmt.Sprintf("Failed to get serial for IP %s in fabric %s", attKey, nw.FabricName))
				}
				attachEntry.SerialNumber = serial
			} else {
				// Serial
				attachEntry.SerialNumber = attKey
			}
			nwEntry.Attachments[attKey] = attachEntry
			kk := nwName + ":" + attachEntry.SerialNumber
			keyMap[kk] = attKey
		}
		nw.Networks[nwName] = nwEntry
	}
	return keyMap
}

func (c NDFC) netAttachmentSerialRemapFromPayload(ctx context.Context, payload *rna.NDFCNetworkAttachments) {
	for i := range payload.NetworkAttachments {
		nwEntry := &payload.NetworkAttachments[i]
		for j := range nwEntry.Attachments {
			// check if key is IP or Serial
			if net.ParseIP(nwEntry.Attachments[j].SerialNumber) != nil {
				// IP
				serial := c.GetSerialFromIP(ctx, payload.FabricName, nwEntry.Attachments[j].SerialNumber)
				if serial == "" {
					panic(fmt.Sprintf("Failed to get serial for IP %s in fabric %s", nwEntry.Attachments[j].SerialNumber, payload.FabricName))
				}
				nwEntry.Attachments[j].SerialNumber = serial
			}
		}
	}
}
