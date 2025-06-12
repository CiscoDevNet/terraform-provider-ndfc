// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
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
	"terraform-provider-ndfc/internal/provider/ndfc/api"
	"terraform-provider-ndfc/internal/provider/ndfc/ndfctemplates"
	"terraform-provider-ndfc/internal/provider/resources/resource_links"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const ResourceLinks = "links"

// RscCreateLinks creates a link in NDFC
func (c *NDFC) RscCreateLinks(ctx context.Context, resp *resource.CreateResponse, in *resource_links.LinksModel) {
	dg := &resp.Diagnostics
	var err error
	tflog.Debug(ctx, "RscCreateLinks: Creating link")

	// Get model data
	modelData := in.GetModelData()

	// Read the template
	ndfcTmpl := c.getTemplateFields(ctx, dg, modelData.TemplateName)
	if ndfcTmpl == nil {
		return
	}

	// Validate the template
	tmplPayload, err := json.Marshal(modelData.LinkParameters)
	if err != nil {
		dg.AddError("Create Error", fmt.Sprintf("Failed to marshal link parameters, got error: %s", err))
		return
	}
	if valid, errors := ndfcTmpl.ValidatePayload(tmplPayload); !valid && len(errors) > 0 {
		dg.AddError("Create Error", fmt.Sprintf("Template validation errors in link parameters, got errors: %s", errors))
		return
	}

	//c.adjustLinksPayload(ctx, dg, modelData, ndfcTmpl)
	// Retrieve switch names before marshaling the model
	if modelData.SourceSwitchName == "" && modelData.SourceDevice != "" {
		sourceSwitchName, err := c.GetDeviceName(ctx, modelData.SourceFabric, modelData.SourceDevice)
		if err != nil {
			log.Printf("Error: Unable to retrieve source switch name: %s", err.Error())
			dg.AddError("Create Error", fmt.Sprintf("Error retrieving switch name for device %s in fabric %s: %s",
				modelData.SourceDevice, modelData.SourceFabric, err.Error()))
			return
		} else {
			modelData.SourceSwitchName = sourceSwitchName
			log.Printf("Set source switch name to: %s", sourceSwitchName)
		}
	}

	if modelData.DestinationSwitchName == "" && modelData.DestinationDevice != "" {
		destSwitchName, err := c.GetDeviceName(ctx, modelData.DestinationFabric, modelData.DestinationDevice)
		if err != nil {
			log.Printf("Error: Unable to retrieve destination switch name: %s", err.Error())
			dg.AddError("Create Error", fmt.Sprintf("Error retrieving switch name for device %s in fabric %s: %s",
				modelData.DestinationDevice, modelData.DestinationFabric, err.Error()))
			return
		} else {
			modelData.DestinationSwitchName = destSwitchName
			log.Printf("Set destination switch name to: %s", destSwitchName)
		}
	}

	// Fill default values
	ndfcTmpl.FillDefaultValues(&modelData.LinkParameters)

	payload, err := json.Marshal(modelData)
	if err != nil {
		dg.AddError("Create Error", fmt.Sprintf("Failed to marshal link payload, got error: %s", err))
		return
	}

	// Get the links API and send the request
	linksAPI := c.getLinksAPI("")
	response, err := linksAPI.Post(payload)
	if err != nil {
		dg.AddError("Create Error", fmt.Sprintf("Failed to create link, got error: %s - %s", err, response.String()))
		return
	}

	// Parse the response
	var result map[string]interface{}
	err = json.Unmarshal([]byte(response.Raw), &result)
	if err != nil {
		dg.AddError("Client Error", fmt.Sprintf("Failed to parse response: %s", err))
		return
	}

	// Set the ID from the response
	linkUUID, ok := result["linkUUID"].(string)
	if !ok {
		dg.AddError("Client Error", "Failed to get link UUID from response")
		return
	}

	in.LinkUuid = types.StringValue(linkUUID)
	tflog.Debug(ctx, fmt.Sprintf("Created link with ID: %s", linkUUID))

	// Get the created link
	c.RscReadLinks(ctx, dg, in, ndfcTmpl)
}

func (c NDFC) adjustLinksPayload(ctx context.Context, dg *diag.Diagnostics,
	in *resource_links.NDFCLinksModel, out *resource_links.NDFCLinksModel,
	ndfcTmpl *ndfctemplates.NDFCTemplate) {

	log.Printf("[DEBUG] Adjusting link payload, templateName in %s, out %s", in.TemplateName, out.TemplateName)
	if ndfcTmpl == nil {
		ndfcTmpl = c.getTemplateFields(ctx, dg, in.TemplateName)
		if ndfcTmpl == nil {
			dg.AddError("Client Error", fmt.Sprintf("Failed to get template fields for template: %s", in.TemplateName))
			return
		}
	}

	// Fill outer values from nvPairs - remove them from nvPairs as they are not present in user config

	srcFabricField := ndfcTmpl.GetFieldWithFlag("IsSourceFabric")
	if srcFabricField != "" {
		out.SourceFabric = out.LinkParameters[srcFabricField]
		if ndfcTmpl.Fields[srcFabricField].IsInternal {
			delete(out.LinkParameters, srcFabricField)
		}
	} else {
		dg.AddError("Client Error", "Source Fabric Missing in payload")
		return
	}

	dstFabricField := ndfcTmpl.GetFieldWithFlag("IsDestinationFabric")
	if dstFabricField != "" {
		out.DestinationFabric = out.LinkParameters[dstFabricField]
		if ndfcTmpl.Fields[dstFabricField].IsInternal {
			delete(out.LinkParameters, dstFabricField)
		}
	} else {
		dg.AddError("Client Error", "Destination Fabric Missing in payload")
		return
	}

	srcDeviceField := ndfcTmpl.GetFieldWithFlag("IsSourceDevice")
	if srcDeviceField != "" {
		out.SourceDevice = out.LinkParameters[srcDeviceField]
		if ndfcTmpl.Fields[srcDeviceField].IsInternal {
			delete(out.LinkParameters, srcDeviceField)
		}
	} else {
		dg.AddError("Client Error", "Source Device Missing in payload")
		return
	}

	dstDeviceField := ndfcTmpl.GetFieldWithFlag("IsDestinationDevice")
	if dstDeviceField != "" {
		out.DestinationDevice = out.LinkParameters[dstDeviceField]
		if ndfcTmpl.Fields[dstDeviceField].IsInternal {
			delete(out.LinkParameters, dstDeviceField)
		}
	} else {
		dg.AddError("Client Error", "Destination Device Missing in payload")
		return
	}

	srcInterfaceField := ndfcTmpl.GetFieldWithFlag("IsSourceInterface")
	if srcInterfaceField != "" {
		out.SourceInterface = out.LinkParameters[srcInterfaceField]
		if ndfcTmpl.Fields[srcInterfaceField].IsInternal {
			delete(out.LinkParameters, srcInterfaceField)
		}
	} else {
		dg.AddError("Client Error", "Source Interface Missing in payload")
		return
	}

	dstInterfaceField := ndfcTmpl.GetFieldWithFlag("IsDestinationInterface")
	if dstInterfaceField != "" {
		out.DestinationInterface = out.LinkParameters[dstInterfaceField]
		if ndfcTmpl.Fields[dstInterfaceField].IsInternal {
			delete(out.LinkParameters, dstInterfaceField)
		}
	} else {
		dg.AddError("Client Error", "Destination Interface Missing in payload")
		return
	}

	srcSwNameField := ndfcTmpl.GetFieldWithFlag("IsSourceSwitchName")
	if srcSwNameField != "" {
		out.SourceSwitchName = out.LinkParameters[srcSwNameField]
		if ndfcTmpl.Fields[srcSwNameField].IsInternal {
			delete(out.LinkParameters, srcSwNameField)
		}
	} else {
		dg.AddError("Client Error", "Source Switch Name Missing in payload")
		return
	}

	dstSwNameField := ndfcTmpl.GetFieldWithFlag("IsDestinationSwitchName")
	if dstSwNameField != "" {
		out.DestinationSwitchName = out.LinkParameters[dstSwNameField]
		if ndfcTmpl.Fields[dstSwNameField].IsInternal {
			delete(out.LinkParameters, dstSwNameField)
		}
	} else {
		dg.AddError("Client Error", "Destination Switch Name Missing in payload")
		return
	}

	// Fill UUID
	out.LinkUuid = out.LinkParameters["LINK_UUID"]

	// Remove all internal Fields

	for key, _ := range out.LinkParameters {
		if ndfcTmpl.IsInternal(key) {
			delete(out.LinkParameters, key)
		}
	}

	// Check if the key/val is present in input, else move them to link_params_computed
	out.LinkParamsComputed = make(map[string]string)
	for key, val := range out.LinkParameters {
		if _, ok := in.LinkParameters[key]; !ok {
			out.LinkParamsComputed[key] = val
			delete(out.LinkParameters, key)
		}
	}

}

func (c *NDFC) readLinks(ctx context.Context, dg *diag.Diagnostics, id string) (*resource_links.NDFCLinksGetPayload, error) {

	linksAPI := c.getLinksAPI(id)
	response, err := linksAPI.Get()
	if err != nil {
		// If error contains "404", the link is not found
		if strings.Contains(err.Error(), "404") {
			// Link not found
			dg.AddError("Client Error", fmt.Sprintf("link with uuid %s not found", id))
			tflog.Debug(ctx, fmt.Sprintf("Link with uuid %s not found", id))
			return nil, fmt.Errorf("link with uuid %s not found", id)
		}
		dg.AddError("Client Error", fmt.Sprintf("Failed to read link, got error: %s", err))
		return nil, err
	}
	tflog.Debug(ctx, fmt.Sprintf("Raw API response for link %s: %s", id, string(response)))
	// Parse the response into a NDFCLinksGetPayload
	modelData := new(resource_links.NDFCLinksGetPayload)
	err = json.Unmarshal(response, modelData)
	if err != nil {
		dg.AddError("Client Error", fmt.Sprintf("Failed to parse response into model: %s", err))
		return nil, err
	}

	// Return the parsed model
	return modelData, nil
}

// RscReadLinks reads a link from NDFC
func (c *NDFC) RscReadLinks(ctx context.Context, dg *diag.Diagnostics,
	in *resource_links.LinksModel, ndfcTmpl *ndfctemplates.NDFCTemplate) {

	linkUUID := in.LinkUuid.ValueString()
	tflog.Debug(ctx, fmt.Sprintf("RscReadLinks: Reading link with ID: %s", linkUUID))

	// Get the links API and send the request

	modelDataPayload, err := c.readLinks(ctx, dg, linkUUID)
	if err != nil {
		dg.AddError("Client Error", fmt.Sprintf("Failed to read link, got error: %s", err))
		return
	}
	modelData := &modelDataPayload.NDFCLinksModel

	inData := in.GetModelData()
	// NDFC links APIs are cumbersome with highly inconsistent payload format for each API
	// POST and GET are way different
	// Need adjustments as nvPairs in POST and GET are way different
	// Also we need to fill out the mandatory fields, which are not available in GET payload
	c.adjustLinksPayload(ctx, dg, inData, modelData, ndfcTmpl)
	if dg.HasError() {
		return
	}

	// Even after this, it is possible that mandatory fields get missed
	// NDFC BUG - When a PUT api is called, fields such as Interface, Device serials etc disappear from nvPair structure
	// Need to look inside sw1-info and sw2-info of the payload to backfill these fields or else terraform flags inconsistency

	// Backfill missing fields
	modelDataPayload.FillMissingFields(modelData)
	// Log the processed data for debugging
	dataJSON, _ := json.MarshalIndent(modelData, "", "  ")
	tflog.Debug(ctx, fmt.Sprintf("Processed link data after filling missing fields: %s", string(dataJSON)))

	// Update the model with the complete response data
	diags := in.SetModelData(modelData)
	if diags.HasError() {
		dg.Append(diags...)
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Read link with ID: %s", in.LinkUuid.ValueString()))
}

// RscUpdateLinks updates a link in NDFC
func (c *NDFC) RscUpdateLinks(ctx context.Context, dg *diag.Diagnostics, id string,
	planData *resource_links.LinksModel, stateData *resource_links.LinksModel,
	configData *resource_links.LinksModel) {
	tflog.Debug(ctx, fmt.Sprintf("RscUpdateLinks: Updating link with ID: %s", id))

	// Step 1: Convert Terraform model data to NDFC models
	planModelData := planData.GetModelData()
	stateModelData := stateData.GetModelData()

	//Computed attributes are not copied in GetModel
	// Here Id and switch names are needed in update payload
	stateModelData.LinkUuid = stateData.LinkUuid.ValueString()
	stateModelData.SourceSwitchName = stateData.SourceSwitchName.ValueString()
	stateModelData.DestinationSwitchName = stateData.DestinationSwitchName.ValueString()

	ndfcTmpl := c.getTemplateFields(ctx, dg, planModelData.TemplateName)
	if ndfcTmpl == nil {
		dg.AddError("Client Error", fmt.Sprintf("Failed to get template fields for template: %s", planModelData.TemplateName))
		return
	}
	currentDataPayload, err := c.readLinks(ctx, dg, id)
	if err != nil {
		dg.AddError("Client Error", fmt.Sprintf("Failed to read link, got error: %s", err))
		return
	}

	currentData := &currentDataPayload.NDFCLinksModel
	currentDataPayload.FillMissingFields(currentData)
	// Log the current data from the API
	dataJSON, _ := json.MarshalIndent(currentData, "", "  ")
	tflog.Debug(ctx, fmt.Sprintf("Current data from API: %s", string(dataJSON)))

	// Step 3: Perform 3-way merge between current data, plan data, and state data
	resource_links.MergeLinksData(currentData, planModelData, stateModelData)

	tmplPayload, err := json.Marshal(currentData.LinkParameters)
	if err != nil {
		dg.AddError("Client Error", fmt.Sprintf("Failed to marshal link parameters, got error: %s", err))
		return
	}
	if valid, errors := ndfcTmpl.ValidatePayload(tmplPayload); !valid && len(errors) > 0 {
		dg.AddError("Client Error", fmt.Sprintf("Template validation errors in link parameters, got errors: %s", errors))
		return
	}

	ndfcTmpl.FillDefaultValues(&currentData.LinkParameters)

	// Log the merged data that will be sent to the API
	mergedJSON, _ := json.MarshalIndent(currentData, "", "  ")
	tflog.Debug(ctx, fmt.Sprintf("Merged data to be sent to update API: %s", string(mergedJSON)))

	// Step 4: Marshal the merged data to JSON and send the PUT request
	payload, err := json.Marshal(currentData)
	if err != nil {
		dg.AddError("Client Error", fmt.Sprintf("Failed to marshal link payload, got error: %s", err))
		return
	}

	// Send the PUT request with the merged payload
	linksAPI := c.getLinksAPI(id)
	res, err := linksAPI.Put(payload)
	if err != nil {
		dg.AddError("Client Error", fmt.Sprintf("Failed to update link, got error: %s - %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Updated link with ID: %s", id))

	// Get the updated link
	c.RscReadLinks(ctx, dg, planData, ndfcTmpl)
}

// RscDeleteLinks deletes a link from NDFC
func (c *NDFC) RscDeleteLinks(ctx context.Context, dg *diag.Diagnostics, in *resource_links.LinksModel) {
	linkUUID := in.LinkUuid.ValueString()
	tflog.Debug(ctx, fmt.Sprintf("RscDeleteLinks: Deleting link with ID: %s", linkUUID))

	// Get the links API and send the request
	linksAPI := c.getLinksAPI(linkUUID)
	res, err := linksAPI.Delete()
	if err != nil {
		dg.AddError("Client Error", fmt.Sprintf("Failed to delete link, got error: %s - %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Deleted link with ID: %s - %s", linkUUID, res.String()))
}

// getLinksAPI returns a LinksAPI instance for the specified link UUID
func (c *NDFC) getLinksAPI(linkUUID string) *api.LinksAPI {
	return api.NewLinksAPI(&c.apiClient, linkUUID)
}

func (c *NDFC) getTemplateFields(ctx context.Context, dg *diag.Diagnostics, templateName string) *ndfctemplates.NDFCTemplate {
	log.Printf("[DEBUG] Attempting to get template %s", templateName)
	tmplContent := c.getTemplateContent(ctx, dg, templateName)
	if tmplContent == nil {
		tflog.Error(ctx, "Failed to get template content")
		dg.AddError("Client Error", "Failed to get template content")
		return nil
	}

	ndfcTmpl := ndfctemplates.NewNDFCTemplate()
	err := ndfcTmpl.ParseTemplate(tmplContent.TemplateContent)
	if err != nil {
		dg.AddError("Client Error", fmt.Sprintf("Failed to parse template, got error: %s", err))
		return nil
	}
	ndfcTmpl.DumpTemplate()
	return ndfcTmpl
}
