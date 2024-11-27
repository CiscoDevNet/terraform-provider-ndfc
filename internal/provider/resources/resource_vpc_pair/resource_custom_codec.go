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
	SerialNumber         string `json:"serialNumber"`
	RecommendationReason string `json:"recommendationReason"`
	LogicalName          string `json:"logicalName"`
	UseVirtualPeerlink   bool   `json:"useVirtualPeerlink"`
	Recommended          bool   `json:"recommended"`
}

type NDFCSwitchesByFabric struct {
	SerialNumber string `json:"serialNumber"`
	PeerSerialNumber string `json:"peerSerialNumber"`
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
	m.PeerOneSwitchDetails.FabricName = customModel.PeerOneSwitchDetails.FabricName
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