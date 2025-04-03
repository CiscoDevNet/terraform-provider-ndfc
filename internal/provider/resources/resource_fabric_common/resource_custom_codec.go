// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_fabric_common

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

type FabricModel interface {
	GetModelData() *NDFCFabricCommonModel
	SetModelData(*NDFCFabricCommonModel) diag.Diagnostics
}
type NdfcFabricPayload struct {
	FabricName        string                `json:"fabricName,omitempty"`
	FabricType        string                `json:"templateName,omitempty"`
	NdfcFabricNvPairs NDFCFabricCommonModel `json:"nvPairs,omitempty"`
}
type NdfcFabricNamePayload struct {
	FabricName string `json:"fabricName,omitempty"`
}
type CustomNdfcFabricNamePayload NdfcFabricNamePayload
type CustomNdfcFabricPayload NdfcFabricPayload

func (m *NdfcFabricNamePayload) UnmarshalJSON(data []byte) error {
	var customModel CustomNdfcFabricNamePayload
	err := json.Unmarshal(data, &customModel)
	if err != nil {
		return err
	}
	m.FabricName = customModel.FabricName
	return nil
}
func (m *NdfcFabricPayload) UnmarshalJSON(data []byte) error {
	var customModel CustomNdfcFabricPayload
	err := json.Unmarshal(data, &customModel)
	if err != nil {
		return err
	}
	m.NdfcFabricNvPairs = customModel.NdfcFabricNvPairs
	m.FabricName = customModel.FabricName
	m.FabricType = customModel.FabricType
	return nil
}
