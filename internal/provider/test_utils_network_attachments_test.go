// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
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
	"terraform-provider-ndfc/internal/provider/resources/resource_network_attachments"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func NetworkAttachmentsModelHelperStateCheck(RscName string, c resource_network_attachments.NDFCNetworkAttachmentsModel, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	for key, value := range c.NetworkAttachments {
		attrNewPath := attrPath.AtName("network_attachments").AtName(key)
		ret = append(ret, NetworkAttachmentsValueHelperStateCheck(RscName, value, attrNewPath)...)
	}
	return ret
}

func AttachmentsValueHelperStateCheck(RscName string, c resource_network_attachments.NDFCAttachmentsValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.SwitchName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("switch_name").String(), c.SwitchName))
	}
	if c.DisplayName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("display_name").String(), c.DisplayName))
	}
	if c.Vlan != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vlan").String(), strconv.Itoa(int(*c.Vlan))))
	}

	if c.AttachState != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("attach_state").String(), c.AttachState))
	}
	if c.Attached != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("attached").String(), strconv.FormatBool(*c.Attached)))
	}
	if c.FreeformConfig != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("freeform_config").String(), c.FreeformConfig))
	}
	if c.DeployThisAttachment {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("deploy_this_attachment").String(), "true"))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("deploy_this_attachment").String(), "false"))
	}

	if c.InstanceValues != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("instance_values").String(), c.InstanceValues))
	}

	return ret
}

func NetworkAttachmentsValueHelperStateCheck(RscName string, c resource_network_attachments.NDFCNetworkAttachmentsValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.NetworkName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("network_name").String(), c.NetworkName))
	}
	for key, value := range c.Attachments {
		attrNewPath := attrPath.AtName("attachments").AtName(key)
		ret = append(ret, AttachmentsValueHelperStateCheck(RscName, value, attrNewPath)...)
	}
	return ret
}
