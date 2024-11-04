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
	"strings"

	"terraform-provider-ndfc/internal/provider/ndfc/api"
	"terraform-provider-ndfc/internal/provider/resources/resource_policy"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
)

const ResourcePolicy = "policy"

func (c *NDFC) RscCreatePolicy(ctx context.Context, dg *diag.Diagnostics, model *resource_policy.PolicyModel) {
	policyApi := api.NewPolicyAPI(c.GetLock(ResourcePolicy), &c.apiClient)

	inData := model.GetModelData()
	if inData.IsPolicyGroup {
		tflog.Error(ctx, "Policy group is not supported")
		return
	}

	data, err := json.Marshal(inData)
	if err != nil {
		tflog.Error(ctx, "Failed to marshal policy data")
		dg.AddError("Failed to marshal policy data", fmt.Sprintf("Error %v", err))
		return
	}

	res, err := policyApi.Post(data)
	if err != nil {
		tflog.Error(ctx, "Failed to create policy")
		dg.AddError("Failed to create policy", fmt.Sprintf("Error %v: %v", err, res.String()))
		return
	}
	// Response to create is - the created policy
	// Try unmarshalling it
	policyId := getPolicyIdFromResponse(&res)
	log.Printf("[TRACE] Policy created with ID: %v", policyId)
	//Deployment
	if inData.Deploy {
		c.RscDeployPolicy(ctx, dg, policyId)
		if dg.HasError() {
			tflog.Error(ctx, "Failed to deploy policy")
			//Roll Back
			model.PolicyId = types.StringValue(policyId)
			c.RscDeletePolicy(ctx, dg, model)
			return
		}
	}

	newModel := c.rscGetPolicy(ctx, dg, policyId)
	if newModel == nil {
		tflog.Error(ctx, "Failed to get policy")
		dg.AddError("Failed to get policy", "Failed to get policy")
		return
	}
	model.SetModelData(newModel)
	model.Deploy = types.BoolValue(inData.Deploy)
	model.FabricName = types.StringValue(inData.FabricName)
}

func getPolicyIdFromResponse(res *gjson.Result) string {
	return res.Get("policyId").String()
}

func (c *NDFC) RscReadPolicy(ctx context.Context, dg *diag.Diagnostics, model *resource_policy.PolicyModel) {

	FabricName := model.FabricName.ValueString()
	Deploy := model.Deploy.ValueBool()
	if model.IsPolicyGroup.ValueBool() {
		tflog.Error(ctx, "Policy group is not supported")
		dg.AddError("Policy group is not supported", "Policy group is not supported")
		return
	}

	policyId := model.PolicyId.ValueString()
	data := c.rscGetPolicy(ctx, dg, policyId)
	if data == nil {
		tflog.Error(ctx, "Failed to get policy")
		dg.AddError("Failed to get policy", "Failed to get policy")
		return
	}
	data.Deploy = Deploy
	data.FabricName = FabricName
	model.SetModelData(data)

}
func (c *NDFC) rscGetPolicy(ctx context.Context, dg *diag.Diagnostics, pID string) *resource_policy.NDFCPolicyModel {

	tflog.Debug(ctx, fmt.Sprintf("RscGetPolicy: Getting policy with ID: %v", pID))
	model := new(resource_policy.NDFCPolicyModel)
	policyApi := api.NewPolicyAPI(c.GetLock(ResourcePolicy), &c.apiClient)
	policyApi.PolicyID = pID
	res, err := policyApi.Get()
	if err != nil {
		tflog.Error(ctx, "Failed to get policy")
		return nil
	}
	log.Printf("[TRACE] Policy data: %v", string(res))
	err = json.Unmarshal(res, &model)
	if err != nil {
		tflog.Error(ctx, "Failed to unmarshal policy data")
		dg.AddError("Failed to unmarshal policy data", fmt.Sprintf("Error %v", err))
		return nil
	}
	return model
}

func (c *NDFC) RscUpdatePolicy(ctx context.Context, dg *diag.Diagnostics, model *resource_policy.PolicyModel) {

	policyID := model.PolicyId.ValueString()
	tflog.Debug(ctx, fmt.Sprintf("RscUpdatePolicy: Updating policy ID %s", policyID))

	if model.IsPolicyGroup.ValueBool() {
		tflog.Error(ctx, "Policy group is not supported")
		dg.AddError("Policy group is not supported", "Policy group is not supported")
		return
	}
	// NDFC inconsistency: PUT requires id field to be present
	// even though policy-id is present. Id is a redundant parameter here
	// Hence id has been added to schema and used only in PUT
	ID := model.Id.ValueInt64()

	policyApi := api.NewPolicyAPI(c.GetLock(ResourcePolicy), &c.apiClient)
	policyApi.PolicyID = policyID
	policyData := model.GetModelData()
	policyData.PolicyId = policyID
	policyData.Id = &ID
	data, err := json.Marshal(policyData)
	if err != nil {
		tflog.Error(ctx, "Failed to marshal policy data")
		dg.AddError("Failed to marshal policy data", fmt.Sprintf("Error %v", err))
		return
	}
	log.Printf("[DEBUG] Updating policy with ID: %v: PUT Data |%s|", policyID, string(data))
	res, err := policyApi.Put(data)
	if err != nil {
		tflog.Error(ctx, "Failed to update policy")
		dg.AddError("Failed to update policy", fmt.Sprintf("Error %v: %v", err, res.String()))
		return
	}
	log.Printf("[TRACE] Policy updated with ID: %v", policyID)
	if policyData.Deploy {
		c.RscDeployPolicy(ctx, dg, policyID)
		if dg.HasError() {
			// cannot rollback to old config as the old data is overwritten in NDFC
			// throw error so that user can correct the config and re-apply
			tflog.Error(ctx, "Failed to deploy policy")
			return
		}
	}
	newModel := c.rscGetPolicy(ctx, dg, policyID)
	if newModel == nil {
		tflog.Error(ctx, "Failed to get policy")
		dg.AddError("Failed to get policy", "Failed to get policy")
		return
	}

	model.SetModelData(newModel)
	model.Deploy = types.BoolValue(policyData.Deploy)
	model.FabricName = types.StringValue(policyData.FabricName)
}

func (c *NDFC) RscDeletePolicy(ctx context.Context, dg *diag.Diagnostics, model *resource_policy.PolicyModel) {

	policyID := model.PolicyId.ValueString()
	FabricName := model.FabricName.ValueString()
	tflog.Debug(ctx, fmt.Sprintf("Before Deploying configuration for Fabric: %s", FabricName))
	tflog.Debug(ctx, fmt.Sprintf("RscDeletePolicy: Deleting policy ID %s", policyID))

	if model.IsPolicyGroup.ValueBool() {
		tflog.Error(ctx, "Policy group is not supported")
		dg.AddError("Policy group is not supported", "Policy group is not supported")
		return
	}
	ID := model.Id.ValueInt64()

	policyApi := api.NewPolicyAPI(c.GetLock(ResourcePolicy), &c.apiClient)
	policyApi.PolicyID = policyID
	policyData := model.GetModelData()
	policyData.Deleted = new(bool)
	*policyData.Deleted = true
	policyData.PolicyId = policyID
	policyData.Id = &ID
	data, err := json.Marshal(policyData)
	if err != nil {
		tflog.Error(ctx, "Failed to marshal policy data")
		dg.AddError("Failed to marshal policy data", fmt.Sprintf("Error %v", err))
		return
	}
	log.Printf("[DEBUG] Deleting policy with ID: %v: Marking delete with PUT Data |%s|", policyID, string(data))
	res, err := policyApi.Put(data)
	if err != nil {
		tflog.Error(ctx, "Failed to mark delete policy")
		dg.AddError("Failed to mark delete policy", fmt.Sprintf("Error %v: %v", err, res.String()))
		return
	}
	// Incase of delete, policyID is made as FabricName:SerialNumber for switch deploy
	// This avoids additional parameters to be passed to delete and also policyID is not used in delete
	policyID = FabricName + ":" + model.GetModelData().DeviceSerialNumber
	c.RscDeployPolicy(ctx, dg, policyID)
	if dg.HasError() {
		tflog.Error(ctx, "Failed to switch deploy")
		return
	}

	log.Printf("[TRACE] Policy deleted with ID: %v", policyID)
	newModel := c.rscGetPolicy(ctx, dg, policyID)
	if newModel != nil {
		dg.AddError("Failed to delete policy", fmt.Sprintf("Policy ID %s still exists", policyID))
		return
	}

}

func (c *NDFC) RscImportPolicy(ctx context.Context, dg *diag.Diagnostics, ID string, model *resource_policy.PolicyModel) {
	tflog.Debug(ctx, fmt.Sprintf("RscImportPolicy: Importing policy ID %s", ID))
	data := c.rscGetPolicy(ctx, dg, ID)
	if data == nil {
		tflog.Error(ctx, "Failed to get policy")
		dg.AddError("Failed to get policy", "Failed to get policy")
		return
	}
	model.SetModelData(data)
}

func (c *NDFC) RscDeployPolicy(ctx context.Context, dg *diag.Diagnostics, policyID string) {
	tflog.Debug(ctx, fmt.Sprintf("RscDeployPolicy: Deploying policy ID %s", policyID))

	// For delete, policyID is in FabricName:SerialNumber format
	// Extract FabricName and SerialNumber
	parts := strings.Split(policyID, ":")
	if len(parts) == 2 {
		FabricName := parts[0]
		serialNumber := parts[1]
		switchSerialNumber := []string{serialNumber}

		tflog.Debug(ctx, fmt.Sprintf("Deploying configuration for Fabric: %s, Serial Numbers: %s", FabricName, serialNumber))
		c.DeployConfiguration(ctx, dg, FabricName, switchSerialNumber)
		return

	} else {
		policyApi := api.NewPolicyAPI(c.GetLock(ResourcePolicy), &c.apiClient)
		policyApi.Deploy = true
		postData, err := json.Marshal([]string{policyID})
		if err != nil {
			tflog.Error(ctx, "Failed to marshal policy data")
			dg.AddError("Failed to marshal policy data", fmt.Sprintf("Error %v", err))
			return
		}
		log.Printf("[DEBUG] Deploying policy with ID: %v: POST Data |%s|", policyID, string(postData))
		res, err := policyApi.Post(postData)
		if err != nil {
			tflog.Error(ctx, "Failed to deploy policy")
			dg.AddError("Failed to deploy policy", fmt.Sprintf("Error %v: %v", err, res.String()))
			return
		}
		log.Printf("[TRACE] Policy deployment ID: %v Response: |%v|", policyID, res)

		if res.Get("0.failedPTIList").Exists() {
			// "failedPTIList" is present in the response
			// Add your code here
			log.Printf("[ERROR] Policy deployment failed for policy ID: %v - %v", policyID, res.Get("0.failedPTIList").String())
			dg.AddError("Failed to deploy policy", fmt.Sprintf("Error %v: %v", err, res.String()))
			return
		}
	}
}

/*

Retaining for reference/future use
func (c NDFC) rsUpdatePolicy(ctx context.Context, dg *diag.Diagnostics, pdata *resource_policy.NDFCPolicyModel) {
	// Update the policy
	data, err := json.Marshal(pdata)
	if err != nil {
		tflog.Error(ctx, "Failed to marshal policy data")
		dg.AddError("Failed to marshal policy data", fmt.Sprintf("Error %v", err))
		return
	}
	policyApi := api.NewPolicyAPI(c.GetLock(ResourcePolicy), &c.apiClient)
	policyApi.PolicyID = pdata.PolicyId
	log.Printf("[DEBUG] Updating policy with ID: %v: PUT Data |%s|", pdata.PolicyId, string(data))
	res, err := policyApi.Put(data)
	if err != nil {
		tflog.Error(ctx, "Failed to update policy")
		dg.AddError("Failed to update policy", fmt.Sprintf("Error %v: %v", err, res.String()))
		return
	}

}
*/
