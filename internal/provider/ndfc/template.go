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
	"terraform-provider-ndfc/internal/provider/ndfc/api"
	"terraform-provider-ndfc/internal/provider/resources/resource_template"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const ResourceTemplate = "template"

func (c *NDFC) RscCreateTemplate(ctx context.Context, dg *diag.Diagnostics, in *resource_template.TemplateModel) {

	inData := in.GetModelData()

	payload, err := json.Marshal(inData)
	if err != nil {
		dg.AddError("Error marshalling template model", err.Error())
		return
	}
	// Create API call
	tmplApi := api.NewTemplateAPI(c.GetLock(ResourceTemplate), &c.apiClient)
	tmplApi.SetValidation(false)

	res, err := tmplApi.Post(payload)
	if err != nil {
		dg.AddError("Error creating template", err.Error())
		return
	}
	tflog.Debug(ctx, fmt.Sprintf("Template created. Response %v", res))
	c.RscGetTemplate(ctx, dg, inData.TemplateName, in)
}

func (c *NDFC) RscGetTemplate(ctx context.Context, dg *diag.Diagnostics, tmplName string, outModel *resource_template.TemplateModel) {
	// Create API call

	tmplApi := api.NewTemplateAPI(c.GetLock(ResourceTemplate), &c.apiClient)
	tmplApi.SetTemplateName(tmplName)
	res, err := tmplApi.Get()
	if err != nil {
		dg.AddError("Error getting template", err.Error())
		return
	}
	log.Println("Response: ", string(res))

	var out resource_template.NDFCTemplateModel
	if err := json.Unmarshal(res, &out); err != nil {
		dg.AddError("Error unmarshalling template model", err.Error())
		return
	}

	outModel.SetModelData(&out)
}

func (c *NDFC) RscValidateTemplateContent(ctx context.Context, dg *diag.Diagnostics, content string) {
	// Create API call
	tmplApi := api.NewTemplateAPI(c.GetLock(ResourceTemplate), &c.apiClient)
	tmplApi.SetValidation(true)
	res, err := tmplApi.Post([]byte(content))
	if err != nil {
		dg.AddError("Template validation failed", fmt.Sprintf("%v: %v", err.Error(), res))
		return
	}
	tflog.Debug(ctx, fmt.Sprintf("Validation successful:%v", res))
}

func (c *NDFC) RscUpdateTemplate(ctx context.Context, dg *diag.Diagnostics,
	planData, stateData, configData *resource_template.TemplateModel) {

	inData := planData.GetModelData()
	tmplApi := api.NewTemplateAPI(c.GetLock(ResourceTemplate), &c.apiClient)
	tmplApi.SetValidation(false)
	tmplApi.SetTemplateName(inData.TemplateName)
	payload, err := json.Marshal(inData)
	if err != nil {
		dg.AddError("Error marshalling template model", err.Error())
		return
	}
	// PUT call
	res, err := tmplApi.Put(payload)
	if err != nil {
		dg.AddError("Error updating template", err.Error())
		return
	}
	tflog.Debug(ctx, fmt.Sprintf("Template updated. Response %v", res))
	c.RscGetTemplate(ctx, dg, inData.TemplateName, planData)
}

func (c *NDFC) RscDeleteTemplate(ctx context.Context, dg *diag.Diagnostics,
	in *resource_template.TemplateModel) {
	inData := in.GetModelData()
	tmplApi := api.NewTemplateAPI(c.GetLock(ResourceTemplate), &c.apiClient)
	tmplApi.SetTemplateName(inData.TemplateName)
	// DELETE call
	res, err := tmplApi.Delete()
	if err != nil {
		dg.AddError("Error deleting template", err.Error())
		return
	}
	tflog.Debug(ctx, fmt.Sprintf("Template deleted. Response %v", res))
}

func (c *NDFC) RscImportTemplate(ctx context.Context, dg *diag.Diagnostics, id string, in *resource_template.TemplateModel) {
	c.RscGetTemplate(ctx, dg, id, in)
}
