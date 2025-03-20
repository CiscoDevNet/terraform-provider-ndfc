// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package provider

import (
	"strconv"
	"terraform-provider-ndfc/internal/provider/resources/resource_inventory_devices"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func InventoryDevicesModelHelperStateCheck(RscName string, c resource_inventory_devices.NDFCInventoryDevicesModel, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.FabricName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_name").String(), c.FabricName))
	}
	if c.AuthProtocol != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("auth_protocol").String(), c.AuthProtocol))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("auth_protocol").String(), "md5"))
	}
	if c.Username != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("username").String(), c.Username))
	}
	if c.Password != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("password").String(), c.Password))
	}
	if c.SeedIp != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("seed_ip").String(), c.SeedIp))
	}
	if c.MaxHops != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("max_hops").String(), strconv.Itoa(int(*c.MaxHops))))
	}
	if c.SetAsIndividualDeviceWriteCredential != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("set_as_individual_device_write_credential").String(), strconv.FormatBool(*c.SetAsIndividualDeviceWriteCredential)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("set_as_individual_device_write_credential").String(), "false"))
	}
	if c.PreserveConfig != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("preserve_config").String(), strconv.FormatBool(*c.PreserveConfig)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("preserve_config").String(), "false"))
	}
	if c.Save != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("save").String(), strconv.FormatBool(*c.Save)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("save").String(), "true"))
	}
	if c.Deploy != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("deploy").String(), strconv.FormatBool(*c.Deploy)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("deploy").String(), "true"))
	}
	if c.Retries != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("retries").String(), strconv.Itoa(int(*c.Retries))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("retries").String(), "300"))
	}
	if c.RetryWaitTimeout != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("retry_wait_timeout").String(), strconv.Itoa(int(*c.RetryWaitTimeout))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("retry_wait_timeout").String(), "5"))
	}
	for key, value := range c.Devices {
		attrNewPath := attrPath.AtName("devices").AtName(key)
		ret = append(ret, DevicesValueHelperStateCheck(RscName, value, attrNewPath)...)
	}
	return ret
}

func DevicesValueHelperStateCheck(RscName string, c resource_inventory_devices.NDFCDevicesValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.Role != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("role").String(), c.Role))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("role").String(), "leaf"))
	}
	if c.DiscoveryType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("discovery_type").String(), c.DiscoveryType))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("discovery_type").String(), "discover"))
	}
	if c.DiscoveryUsername != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("discovery_username").String(), c.DiscoveryUsername))
	}
	if c.DiscoveryPassword != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("discovery_password").String(), c.DiscoveryPassword))
	}
	if c.DiscoveryAuthProtocol != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("discovery_auth_protocol").String(), c.DiscoveryAuthProtocol))
	}
	if c.SerialNumber != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("serial_number").String(), c.SerialNumber))
	}
	if c.Model != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("model").String(), c.Model))
	}
	if c.Version != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("version").String(), c.Version))
	}
	if c.Hostname != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("hostname").String(), c.Hostname))
	}
	if c.ImagePolicy != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("image_policy").String(), c.ImagePolicy))
	}
	if c.Gateway != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("gateway").String(), c.Gateway))
	}
	if c.Breakout != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("breakout").String(), c.Breakout))
	}
	if c.PortMode != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("port_mode").String(), c.PortMode))
	}
	if c.Uuid != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("uuid").String(), c.Uuid))
	}
	if c.SwitchDbId != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("switch_db_id").String(), c.SwitchDbId))
	}
	if c.DeviceIndex != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("device_index").String(), c.DeviceIndex))
	}
	if c.VdcId != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vdc_id").String(), c.VdcId))
	}
	if c.VdcMac != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vdc_mac").String(), c.VdcMac))
	}
	if c.Mode != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mode").String(), c.Mode))
	}
	if c.ConfigStatus != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("config_status").String(), c.ConfigStatus))
	}
	if c.OperStatus != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("oper_status").String(), c.OperStatus))
	}
	if c.DiscoveryStatus != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("discovery_status").String(), c.DiscoveryStatus))
	}
	if c.Managable != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("managable").String(), strconv.FormatBool(*c.Managable)))
	}
	return ret
}
