// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_interface_common

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

type InterfaceModel interface {
	GetModelData() *NDFCInterfaceCommonModel
	SetModelData(*NDFCInterfaceCommonModel) diag.Diagnostics
	GetInterfaceType() string
	GetID() string
	SetID(string)
}

type InterfaceValue interface {
	SetValue(*NDFCInterfacesValue)
}

type NDFCInterfacesPayload struct {
	Policy     string                `json:"policy,omitempty"`
	Interfaces []NDFCInterfacesValue `json:"interfaces,omitempty"`
}

type NDFCInterfacesDeploy []NDFCInterfaceDeploy

type NDFCInterfaceDeploy struct {
	IfName       string `json:"ifName"`
	SerialNumber string `json:"serialNumber"`
}

type NDFCInterfaceDetails struct {
	InterfaceName string `json:"ifName"`
	DeployStatus  string `json:"complianceStatus"`
}
type CustomInterfacePayload NDFCInterfacesPayload
type CustomInterfaceValue NDFCInterfacesValue
type CustomTemplateValue struct {
	TemplateKV map[string]string `json:"nvPairs,omitempty"`
}

// Override Marshall to merge custom template parameters to nvPairs
// In a normal Marshall, only NvPairs alone with other attributes are marshalled into payload
// To handle custom template parameters, we need to load customPolicyParameters to nvPairs
// This is done by overriding the Marshall method for NDFCInterfacesValue

func (c NDFCInterfacesValue) MarshalJSON() ([]byte, error) {

	/* Step 1: Load the custom struct and convert to JSON */

	custom := CustomInterfaceValue{}
	custom.SerialNumber = c.SerialNumber
	custom.InterfaceName = c.InterfaceName
	custom.InterfaceType = c.InterfaceType
	custom.DeploymentStatus = c.DeploymentStatus
	custom.NvPairs = c.NvPairs

	/* Custom template parameters not present
	 * This is default case; just return normal payload
	 */
	if len(c.CustomPolicyParameters) == 0 {
		return json.Marshal(custom)
	}
	// customData is the JSON representation
	customData, err := json.Marshal(custom)
	if err != nil {
		return nil, err
	}

	/* Step 2: Load the JSON back to map[string]interface{}
	 * This is done so that we can merge the customPolicyParameters to nvPairs
	 */
	customMap := map[string]interface{}{}
	err = json.Unmarshal(customData, &customMap)
	if err != nil {
		return nil, err
	}

	/* Step 3: Load the custom template KV to nvPairs
	 * Clear out any content from NvPairs and load from customPolicyParameters
	 */

	customTmpl := customMap["nvPairs"].(map[string]interface{})
	// Remove all entries from nvPairs
	clear(customTmpl)
	// Load customPolicyParameters to nvPairs
	for k, v := range c.CustomPolicyParameters {
		customTmpl[k] = v
	}
	// Marshall to final JSON payload ready to be sent to NDFC
	return json.Marshal(customMap)
}

func (c *NDFCInterfacesValue) UnmarshalJSON(data []byte) error {
	var custom CustomInterfaceValue
	//log.Printf("NDFCInterfacesValue: UnmarshalJSON: %s", string(data))
	err := json.Unmarshal(data, &custom)
	if err != nil {
		return err
	}

	c.SerialNumber = custom.SerialNumber
	c.InterfaceName = custom.InterfaceName
	c.InterfaceType = custom.InterfaceType

	c.NvPairs = custom.NvPairs

	var cTmpl CustomTemplateValue

	err = json.Unmarshal(data, &cTmpl)
	if err != nil {
		return err
	}
	// Load nvPairs into CustomPolicyParameters
	c.CustomPolicyParameters = make(map[string]string)
	for k, v := range cTmpl.TemplateKV {
		//log.Printf("NDFCInterfacesValue: UnmarshalJSON: CustomPolicyParameters: %s:%s", k, v)
		c.CustomPolicyParameters[k] = v
	}

	return nil
}
