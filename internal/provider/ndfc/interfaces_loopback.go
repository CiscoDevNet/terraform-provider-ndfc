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
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"terraform-provider-ndfc/internal/provider/resources/resource_interface_common"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type NDFCLoopbackInterface struct {
	NDFCInterfaceCommon
}

const ResourceLoopbackInterface = "interface_loopback"

func (i *NDFCLoopbackInterface) CreateInterface(ctx context.Context, diags *diag.Diagnostics, inData *resource_interface_common.NDFCInterfaceCommonModel) {
	intfPayload := resource_interface_common.NDFCInterfacesPayload{}
	if len(inData.Interfaces) <= 0 {
		tflog.Debug(ctx, "No interfaces to create")
		return
	}
	intfPayload.Policy = inData.Policy
	for i, intf := range inData.Interfaces {
		inData.Interfaces[i] = intf
		intfPayload.Interfaces = append(intfPayload.Interfaces, intf)
	}
	i.createInterface(ctx, diags, &intfPayload)
}

func (i *NDFCLoopbackInterface) DeleteInterface(ctx context.Context, dg *diag.Diagnostics,
	inData *resource_interface_common.NDFCInterfaceCommonModel) {

	tflog.Debug(ctx, "Deleting interfaces")
	if len(inData.Interfaces) <= 0 {
		tflog.Debug(ctx, "No interfaces to delete")
		return
	}
	// DELETE and Deploy uses similar payload
	intfPayload := resource_interface_common.NDFCInterfacesDeploy{}

	for _, intf := range inData.Interfaces {
		intfPayload = append(intfPayload, resource_interface_common.NDFCInterfaceDeploy{IfName: intf.InterfaceName,
			SerialNumber: intf.SerialNumber})
		tflog.Debug(ctx, fmt.Sprintf("Deleting interface: %s:%s", intf.SerialNumber, intf.InterfaceName))
	}

	i.deleteInterface(ctx, dg, &intfPayload)
	if dg.HasError() {
		tflog.Error(ctx, "Error deleting interfaces")
		return
	}
	i.deployInterface(ctx, dg, intfPayload)
}

// NDFC BUG: Certain fields in NDFC interface must be present even if empty for processing to work correctly
// For such fields omitempty is not set in struct and json marshall would include them in payload
// The tags that are not needed for loopback but used for some other ifType due to this can be removed here before sending to API
func (i *NDFCLoopbackInterface) GetPayload(ctx context.Context, diags *diag.Diagnostics,
	intfPayload *resource_interface_common.NDFCInterfacesPayload) ([]byte, error) {

	tflog.Debug(ctx, "GetPayload - overrided call for NDFCEthernetInterface")

	rawData, err := json.Marshal(intfPayload)
	if err != nil {
		return nil, err
	}
	removeAttributes := []string{
		",\"ROUTING_TAG\":\"\"", // VLAN interface needs empty "ROUTING_TAG" field in payload
	}
	log.Printf("Removing attributes from payload before sending to API")
	log.Printf("Data %s", string(rawData))
	for _, attr := range removeAttributes {
		rawData = []byte(strings.Replace(string(rawData), attr, "", -1))
	}
	log.Printf("Data after removing attributes %s", string(rawData))
	return rawData, nil
}
