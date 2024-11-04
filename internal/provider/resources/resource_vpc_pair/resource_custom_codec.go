// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_vpc_pair

import (
	"encoding/json"
)

type NDFCVpcPairRecommendations struct {
	LanId                int     `json:"lanId"`
	EthSwitchId          int     `json:"ethSwitchId"`
	LogicalName          string  `json:"logicalName"`
	IpAddress            string  `json:"ipAddress"`
	VdcName              *string `json:"vdcName"` // Use a pointer to handle null values
	VdcId                int     `json:"vdcId"`
	SerialNumber         string  `json:"serialNumber"`
	NxosVersion          string  `json:"nxosVersion"`
	FabricName           *string `json:"fabricName"` // Use a pointer to handle null values
	RecommendationReason string  `json:"recommendationReason"`
	BlockSelection       bool    `json:"blockSelection"`
	Uuid                 *string `json:"uuid"`         // Use a pointer to handle null values
	PlatformType         *string `json:"platformType"` // Use a pointer to handle null values
	UseVirtualPeerlink   bool    `json:"useVirtualPeerlink"`
	CurrentPeer          bool    `json:"currentPeer"`
	Recommended          bool    `json:"recommended"`
}

type CustomNDFCVpcPairModel NDFCVpcPairModel

func (m *NDFCVpcPairModel) UnmarshalJSON(data []byte) error {
	var customModel CustomNDFCVpcPairModel
	if err := json.Unmarshal(data, &customModel); err != nil {
		return err
	}
	m.SerialNumbers = []string{customModel.PeerOneId, customModel.PeerTwoId}
	m.UseVirtualPeerlink = customModel.UseVirtualPeerlink
	m.PeerOneId = customModel.PeerOneId
	m.PeerTwoId = customModel.PeerTwoId
	return nil
}
func (m *NDFCVpcPairModel) MarshalJSON() ([]byte, error) {
	var customModel CustomNDFCVpcPairModel
	customModel.PeerOneId = m.SerialNumbers[0]
	customModel.PeerTwoId = m.SerialNumbers[1]
	customModel.UseVirtualPeerlink = m.UseVirtualPeerlink
	customModel.SerialNumbers = m.SerialNumbers
	return json.Marshal(customModel)
}
