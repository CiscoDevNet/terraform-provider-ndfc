// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_vrf_attachments

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Inside TEST_HELPER_STATE_CHECK

func (c *NDFCVrfAttachmentsModel) HelperStateCheck(RscName string, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{

		resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_name").String(), c.FabricName),
	}
	for key, value := range c.VrfAttachments {
		attrNewPath := attrPath.AtMapKey(key)
		ret = append(ret, value.HelperStateCheck(RscName, attrNewPath)...)
	}
	return ret
}

// Inside TEST_HELPER_STATE_CHECK
func (c *NDFCAttachListValue) HelperStateCheck(RscName string, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{

		resource.TestCheckResourceAttr(RscName, attrPath.AtName("switch_name").String(), c.SwitchName),

		resource.TestCheckResourceAttr(RscName, attrPath.AtName("attach_state").String(), c.AttachState),

		resource.TestCheckResourceAttr(RscName, attrPath.AtName("freeform_config").String(), c.FreeformConfig),

		resource.TestCheckResourceAttr(RscName, attrPath.AtName("loopback_ipv4").String(), c.InstanceValues.LoopbackIpv4),
		resource.TestCheckResourceAttr(RscName, attrPath.AtName("loopback_ipv6").String(), c.InstanceValues.LoopbackIpv6),
	}
	if c.Vlan != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vlan").String(), strconv.Itoa(int(*c.Vlan))))
	}
	if c.InstanceValues.LoopbackId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("loopback_id").String(), strconv.Itoa(int(*c.InstanceValues.LoopbackId))))
	}
	return ret
}

// Inside TEST_HELPER_STATE_CHECK
func (c *NDFCVrfAttachmentsValue) HelperStateCheck(RscName string, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{

		resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_name").String(), c.VrfName),
	}
	for key, value := range c.AttachList {
		attrNewPath := attrPath.AtMapKey(key)
		ret = append(ret, value.HelperStateCheck(RscName, attrNewPath)...)
	}
	return ret
}
