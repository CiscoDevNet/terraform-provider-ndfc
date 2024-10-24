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
	"terraform-provider-ndfc/internal/provider/resources/resource_vpc_pair"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func VpcPairModelHelperStateCheck(RscName string, c resource_vpc_pair.NDFCVpcPairModel, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.UseVirtualPeerlink != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("use_virtual_peerlink").String(), strconv.FormatBool(*c.UseVirtualPeerlink)))
	}
	if c.NvPairs.FabricName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_name").String(), c.NvPairs.FabricName))
	}
	if c.Deploy {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("deploy").String(), "true"))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("deploy").String(), "false"))
	}
	return ret
}
