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
	"strings"
	"terraform-provider-ndfc/tfutils/go-nd"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/tidwall/gjson"
)

func (c NDFC) GetFabricInventory(ctx context.Context, diags *diag.Diagnostics, fabricName string) gjson.Result {
	res, err := c.apiClient.Get(fmt.Sprintf("/lan-fabric/rest/control/fabrics/%s/inventory/switchesByFabric", fabricName))
	if err != nil {
		diags.AddError("Retrieval Inventory Failed", err.Error())
	}
	return res
}

func (c NDFC) GetFabricInventoryPoap(ctx context.Context, diags *diag.Diagnostics, fabricName string) gjson.Result {
	res, err := c.apiClient.Get(fmt.Sprintf("/lan-fabric/rest/control/fabrics/%s/inventory/poap", fabricName))
	if err != nil {
		diags.AddError("Retrieval POAP Failed", err.Error())
	}
	return res
}

func (c NDFC) IsLanCredentialSet(ctx context.Context, diags *diag.Diagnostics) gjson.Result {
	res, err := c.apiClient.Get("/lan-fabric/rest/lanConfig/isLanCredentialsSet")
	if err != nil {
		diags.AddError("Retrieval of Lan Credentials Set Failed", err.Error())
	}
	return res
}

func (c NDFC) GetLanCredentialSet(ctx context.Context, diags *diag.Diagnostics) gjson.Result {
	res, err := c.apiClient.Get("/lan-fabric/rest/lanConfig/getLanSwitchCredentialsWithType")
	if err != nil {
		diags.AddError("Retrieval of Lan Credentials Set Failed", err.Error())
	}
	return res
}

func (c NDFC) SetLanCredentialSet(ctx context.Context, diags *diag.Diagnostics, lanCredentials string) {
	_, err := c.apiClient.Post(
		"/lan-fabric/rest/lanConfig/saveSwitchCredentials",
		lanCredentials,
		func(req *nd.Req) {
			req.HttpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		},
	)
	if err != nil {
		diags.AddError("Setting of Lan Credentials Failed", err.Error())
	}
}

func (c NDFC) TestReachability(ctx context.Context, diags *diag.Diagnostics, fabricName string, deviceData []byte) gjson.Result {
	res, err := c.apiClient.Post(fmt.Sprintf("/lan-fabric/rest/control/fabrics/%s/inventory/test-reachability", fabricName), string(deviceData))
	if err != nil {
		diags.AddError("Test Reachability Failed", err.Error())
	}
	return res
}

func (c NDFC) SetFabricInventoryDiscover(ctx context.Context, diags *diag.Diagnostics, fabricName string, setAndUseDiscoveryCredForLan bool, deviceData []byte) {
	_, err := c.apiClient.Post(
		fmt.Sprintf("/lan-fabric/rest/control/fabrics/%s/inventory/discover?setAndUseDiscoveryCredForLan=%t", fabricName, setAndUseDiscoveryCredForLan),
		string(deviceData),
	)
	if err != nil {
		diags.AddError("Discover Devices Failed", err.Error())
	}
}

func (c NDFC) SetFabricInventoryRediscover(ctx context.Context, diags *diag.Diagnostics, switchDbIDs []byte) {
	_, err := c.apiClient.Post("/lan-discovery/rediscoverSwitch", string(switchDbIDs))
	if err != nil {
		diags.AddError("Re-Discover Devices Failed", err.Error())
	}
}

func (c NDFC) SetFabricInventoryPoap(ctx context.Context, diags *diag.Diagnostics, fabricName string, poapData []byte) gjson.Result {
	res, err := c.apiClient.Post(fmt.Sprintf("/lan-fabric/rest/control/fabrics/%s/inventory/poap", fabricName), string(poapData))
	if err != nil {
		diags.AddError("POAP Devices Failed", err.Error())
	}
	return res
}

func (c NDFC) SetFabricInventorySerialNumber(ctx context.Context, diags *diag.Diagnostics, fabricName, preprovisionSerial, serialNumber string) gjson.Result {
	res, err := c.apiClient.Post(fmt.Sprintf("/lan-fabric/rest/control/fabrics/%s/swapSN/%s/%s", fabricName, preprovisionSerial, serialNumber), "")
	if err != nil {
		diags.AddError("Update of Serial Number Failed", err.Error())
	}
	return res
}

func (c NDFC) SetFabricInventoryRma(ctx context.Context, diags *diag.Diagnostics, fabricName string, rmaData []byte) gjson.Result {
	res, err := c.apiClient.Post(fmt.Sprintf("/lan-fabric/rest/control/fabrics/%s/rma", fabricName), string(rmaData))
	if err != nil {
		diags.AddError("RMA Devices Failed", err.Error())
	}
	return res
}

func (c NDFC) SetMaintananceMode(ctx context.Context, diags *diag.Diagnostics, fabricName, serialNumber string) {
	_, err := c.apiClient.Post(fmt.Sprintf("/lan-fabric/rest/control/fabrics/%s/switches/%s/maintenance-mode", fabricName, serialNumber), "")
	if err != nil {
		diags.AddError("Set Maintanance Mode Failed", err.Error())
	}
}

func (c NDFC) UpdateRole(ctx context.Context, diags *diag.Diagnostics, switchDbID, role string) {
	_, err := c.apiClient.Put(fmt.Sprintf("/rest/topology/role/%s?newRole=%s", switchDbID, role), "")
	if err != nil {
		diags.AddError("Update of Role Failed", err.Error())
	}
}

func (c NDFC) UpdateSerialNumber(ctx context.Context, diags *diag.Diagnostics, fabricName, oldSerial, newSerial string) {
	_, err := c.apiClient.Post(fmt.Sprintf("/lan-fabric/rest/control/fabrics/%s/swapSN/%s/%s", fabricName, oldSerial, newSerial), "")
	if err != nil {
		diags.AddError("Update of Serial Number Failed", err.Error())
	}
}

func (c NDFC) DeleteFabricInventoryDevices(ctx context.Context, diags *diag.Diagnostics, fabricName string, uuidList []string) {
	_, err := c.apiClient.Delete(fmt.Sprintf("/lan-fabric/rest/control/fabrics/%s/switches/UUID/%s", fabricName, strings.Join(uuidList, ",")), "")
	if err != nil {
		diags.AddError("Delete Devices Failed", err.Error())
	}
}

func (c NDFC) GetDeviceRole(ctx context.Context, diags *diag.Diagnostics, serialNum string) gjson.Result {
	res, err := c.apiClient.Get(fmt.Sprintf("/rest/control/switches/roles?serialNumber=%s", serialNum))
	if err != nil {
		diags.AddError("Get Device Role Failed", err.Error())
	}
	return res
}
