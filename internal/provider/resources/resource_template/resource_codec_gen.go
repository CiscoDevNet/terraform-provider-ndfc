// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package resource_template

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCTemplateModel struct {
	InstanceId         *int64   `json:"instanceClassId,omitempty"`
	TemplateName       string   `json:"templatename,omitempty"`
	Description        string   `json:"description,omitempty"`
	Tags               []string `json:"tags,omitempty"`
	SupportedPlatforms string   `json:"supportedPlatforms,omitempty"`
	FileName           string   `json:"fileName,omitempty"`
	TemplateType       string   `json:"templateType,omitempty"`
	TemplateContent    string   `json:"newContent,omitempty"`
	ContentType        string   `json:"contentType,omitempty"`
	TemplateSubType    string   `json:"templateSubType,omitempty"`
}

func (v *TemplateModel) SetModelData(jsonData *NDFCTemplateModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.InstanceId != nil {
		v.InstanceId = types.Int64Value(*jsonData.InstanceId)

	} else {
		v.InstanceId = types.Int64Null()
	}

	if jsonData.TemplateName != "" {
		v.TemplateName = types.StringValue(jsonData.TemplateName)
	} else {
		v.TemplateName = types.StringNull()
	}

	if jsonData.Description != "" {
		v.Description = types.StringValue(jsonData.Description)
	} else {
		v.Description = types.StringNull()
	}

	if len(jsonData.Tags) == 0 {
		log.Printf("v.Tags is empty")
		v.Tags, err = types.SetValue(types.StringType, []attr.Value{})
		if err != nil {
			log.Printf("Error in converting []string to  List %v", err)
			return err
		}
	} else {
		listData := make([]attr.Value, len(jsonData.Tags))
		for i, item := range jsonData.Tags {
			listData[i] = types.StringValue(item)
		}
		v.Tags, err = types.SetValue(types.StringType, listData)
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}
	if jsonData.SupportedPlatforms != "" {
		v.SupportedPlatforms = types.StringValue(jsonData.SupportedPlatforms)
	} else {
		v.SupportedPlatforms = types.StringNull()
	}

	if jsonData.FileName != "" {
		v.FileName = types.StringValue(jsonData.FileName)
	} else {
		v.FileName = types.StringNull()
	}

	if jsonData.TemplateType != "" {
		v.TemplateType = types.StringValue(jsonData.TemplateType)
	} else {
		v.TemplateType = types.StringNull()
	}

	if jsonData.TemplateContent != "" {
		v.TemplateContent = types.StringValue(jsonData.TemplateContent)
	} else {
		v.TemplateContent = types.StringNull()
	}

	if jsonData.ContentType != "" {
		v.ContentType = types.StringValue(jsonData.ContentType)
	} else {
		v.ContentType = types.StringNull()
	}

	if jsonData.TemplateSubType != "" {
		v.TemplateSubType = types.StringValue(jsonData.TemplateSubType)
	} else {
		v.TemplateSubType = types.StringNull()
	}

	return err
}

func (v TemplateModel) GetModelData() *NDFCTemplateModel {
	var data = new(NDFCTemplateModel)

	//MARSHAL_BODY

	if !v.TemplateName.IsNull() && !v.TemplateName.IsUnknown() {
		data.TemplateName = v.TemplateName.ValueString()
	} else {
		data.TemplateName = ""
	}

	if !v.Description.IsNull() && !v.Description.IsUnknown() {
		data.Description = v.Description.ValueString()
	} else {
		data.Description = ""
	}

	if !v.Tags.IsNull() && !v.Tags.IsUnknown() {
		listStringData := make([]string, len(v.Tags.Elements()))
		dg := v.Tags.ElementsAs(context.Background(), &listStringData, false)
		if dg.HasError() {
			panic(dg.Errors())
		}
		data.Tags = make([]string, len(listStringData))
		copy(data.Tags, listStringData)
	}

	if !v.SupportedPlatforms.IsNull() && !v.SupportedPlatforms.IsUnknown() {
		data.SupportedPlatforms = v.SupportedPlatforms.ValueString()
	} else {
		data.SupportedPlatforms = ""
	}

	if !v.TemplateType.IsNull() && !v.TemplateType.IsUnknown() {
		data.TemplateType = v.TemplateType.ValueString()
	} else {
		data.TemplateType = ""
	}

	if !v.TemplateContent.IsNull() && !v.TemplateContent.IsUnknown() {
		data.TemplateContent = v.TemplateContent.ValueString()
	} else {
		data.TemplateContent = ""
	}

	if !v.ContentType.IsNull() && !v.ContentType.IsUnknown() {
		data.ContentType = v.ContentType.ValueString()
	} else {
		data.ContentType = ""
	}

	if !v.TemplateSubType.IsNull() && !v.TemplateSubType.IsUnknown() {
		data.TemplateSubType = v.TemplateSubType.ValueString()
	} else {
		data.TemplateSubType = ""
	}

	return data
}
