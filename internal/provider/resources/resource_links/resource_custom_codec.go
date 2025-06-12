// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_links

import (
	"encoding/json"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// GetID returns the link UUID for the model
func (l *LinksModel) GetID() string {
	return l.LinkUuid.ValueString()
}

// SetID sets the link UUID for the model
func (l *LinksModel) SetID(id string) {
	if id == "" {
		l.LinkUuid = types.StringNull()
		return
	}
	l.LinkUuid = types.StringValue(id)
}

// NDFCLinksGetPayload represents the complete structure of the GET response
// including sw1-info and sw2-info fields that contain additional information
// SwitchInfo represents the switch information structure returned in the links API response
type SwitchInfo struct {
	SWUUIDID       int    `json:"sw-UUID-ID,omitempty"`
	SWModelName    string `json:"sw-model-name,omitempty"`
	EthswDBID      int    `json:"ethsw-DBID,omitempty"`
	FabricName     string `json:"fabric-name,omitempty"`
	IsVDC          string `json:"is-vdc,omitempty"`
	SwitchRole     string `json:"switch-role,omitempty"`
	IfAdminStatus  string `json:"if-admin-status,omitempty"`
	IfName         string `json:"if-name,omitempty"`
	SWSerialNumber string `json:"sw-serial-number,omitempty"`
	IfOpReason     string `json:"if-op-reason,omitempty"`
	IfOpStatus     string `json:"if-op-status,omitempty"`
	FabricID       string `json:"fabric-id,omitempty"`
	VDCID          string `json:"vdc-id,omitempty"`
	SWUUID         string `json:"sw-UUID,omitempty"`
	SWSystemName   string `json:"sw-sys-name,omitempty"`
}

type NDFCLinksGetPayload struct {
	// Embed the standard NDFCLinksModel fields
	NDFCLinksModel

	// Additional fields from the GET response
	LinkUUID      string     `json:"link-uuid,omitempty"`
	IsPresent     bool       `json:"is-present,omitempty"`
	LinkType      string     `json:"link-type,omitempty"`
	IsDiscovered  bool       `json:"is-discovered,omitempty"`
	IsPlanned     bool       `json:"is-planned,omitempty"`
	PolicyID      string     `json:"policyId,omitempty"`
	IsPortChannel bool       `json:"is-port-channel,omitempty"`
	FabricName    string     `json:"fabricName,omitempty"`
	Sw1Info       SwitchInfo `json:"sw1-info,omitempty"`
	Sw2Info       SwitchInfo `json:"sw2-info,omitempty"`

	// Switch info fields will be stored in LinkParameters with appropriate keys
}

// FillMissingFields fills in missing nvPairs values from sw1-info and sw2-info
// This ensures all data is available even if the API doesn't include everything in the nvPairs section
func (s *NDFCLinksGetPayload) FillMissingFields(result *NDFCLinksModel) {

	// Fill in source device and fabric info from LinkParameters
	if result.SourceDevice == "" && s.Sw1Info.SWSerialNumber != "" {
		result.SourceDevice = s.Sw1Info.SWSerialNumber
	}

	if result.SourceFabric == "" && s.Sw1Info.FabricName != "" {
		result.SourceFabric = s.Sw1Info.FabricName
	}

	if result.SourceSwitchName == "" && s.Sw1Info.SWSystemName != "" {
		result.SourceSwitchName = s.Sw1Info.SWSystemName
	}

	if result.SourceInterface == "" && s.Sw1Info.IfName != "" {
		result.SourceInterface = s.Sw1Info.IfName
	}

	// Fill in destination device and fabric info from LinkParameters
	if result.DestinationDevice == "" && s.Sw2Info.SWSerialNumber != "" {
		result.DestinationDevice = s.Sw2Info.SWSerialNumber
	}

	if result.DestinationFabric == "" && s.Sw2Info.FabricName != "" {
		result.DestinationFabric = s.Sw2Info.FabricName
	}

	if result.DestinationSwitchName == "" && s.Sw2Info.SWSystemName != "" {
		result.DestinationSwitchName = s.Sw2Info.SWSystemName
	}

	if result.DestinationInterface == "" && s.Sw2Info.IfName != "" {
		result.DestinationInterface = s.Sw2Info.IfName
	}

}

// MergeLinksData performs a 3-way merge between currentData, planModelData, and stateModelData.
// It returns the merged NDFCLinksModel that should be sent to the API.
func MergeLinksData(currentData, planModelData, stateModelData *NDFCLinksModel) {
	// Start with the current data as the base
	mergedData := currentData

	currentData.SourceFabric = planModelData.SourceFabric
	currentData.DestinationFabric = planModelData.DestinationFabric
	currentData.SourceDevice = planModelData.SourceDevice
	currentData.DestinationDevice = planModelData.DestinationDevice
	currentData.SourceSwitchName = stateModelData.SourceSwitchName
	currentData.DestinationSwitchName = stateModelData.DestinationSwitchName
	currentData.SourceInterface = planModelData.SourceInterface
	currentData.DestinationInterface = planModelData.DestinationInterface

	// Merge LinkParameters
	if planModelData.LinkParameters != nil {
		for key, value := range planModelData.LinkParameters {
			if value != stateModelData.LinkParameters[key] {
				mergedData.LinkParameters[key] = value
			}
		}
	}

	dataJSON, _ := json.MarshalIndent(planModelData, "", "  ")
	log.Printf("[DEBUG] Plan data from API: %s", string(dataJSON))

	dataJSON, _ = json.MarshalIndent(stateModelData, "", "  ")
	log.Printf("[DEBUG] State data from API: %s", string(dataJSON))

}
