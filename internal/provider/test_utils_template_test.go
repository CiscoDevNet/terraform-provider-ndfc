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
	"terraform-provider-ndfc/internal/provider/resources/resource_template"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TemplateModelHelperStateCheck(RscName string, c resource_template.NDFCTemplateModel, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.InstanceId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("instance_id").String(), strconv.Itoa(int(*c.InstanceId))))
	}
	if c.TemplateName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("template_name").String(), c.TemplateName))
	}
	if c.Description != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("description").String(), c.Description))
	}
	if c.SupportedPlatforms != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("supported_platforms").String(), c.SupportedPlatforms))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("supported_platforms").String(), "All"))
	}
	if c.FileName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("file_name").String(), c.FileName))
	}
	if c.TemplateType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("template_type").String(), c.TemplateType))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("template_type").String(), "POLICY"))
	}
	if c.TemplateContent != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("template_content").String(), c.TemplateContent))
	}
	if c.ContentType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("content_type").String(), c.ContentType))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("content_type").String(), "TEMPLATE_CLI"))
	}
	if c.TemplateSubType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("template_sub_type").String(), c.TemplateSubType))
	}
	return ret
}
