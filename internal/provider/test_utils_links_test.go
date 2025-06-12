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
	"terraform-provider-ndfc/internal/provider/resources/resource_links"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func LinksModelHelperStateCheck(RscName string, c resource_links.NDFCLinksModel, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.LinkUuid != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("link_uuid").String(), c.LinkUuid))
	}
	if c.SourceFabric != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("source_fabric").String(), c.SourceFabric))
	}
	if c.DestinationFabric != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("destination_fabric").String(), c.DestinationFabric))
	}
	if c.SourceDevice != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("source_device").String(), c.SourceDevice))
	}
	if c.DestinationDevice != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("destination_device").String(), c.DestinationDevice))
	}
	if c.SourceInterface != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("source_interface").String(), c.SourceInterface))
	}
	if c.DestinationInterface != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("destination_interface").String(), c.DestinationInterface))
	}
	if c.TemplateName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("template_name").String(), c.TemplateName))
	}
	if c.SourceSwitchName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("source_switch_name").String(), c.SourceSwitchName))
	}
	if c.DestinationSwitchName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("destination_switch_name").String(), c.DestinationSwitchName))
	}

	if c.IsPresent != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("is_present").String(), c.IsPresent))
	}
	if c.LinkType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("link_type").String(), c.LinkType))
	}
	if c.IsDiscovered != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("is_discovered").String(), c.IsDiscovered))
	}
	return ret
}
