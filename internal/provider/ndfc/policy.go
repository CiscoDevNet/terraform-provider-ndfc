// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Package ndfc provides functionality for managing NDFC resources
package ndfc

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"

	"terraform-provider-ndfc/internal/provider/ndfc/api"
	"terraform-provider-ndfc/internal/provider/resources/resource_policy"
	custtypes "terraform-provider-ndfc/internal/provider/types"
)

// Error constants for consistent error messages
const (
	ErrGetPolicy          = "failed to get policy"
	ErrCreatePolicy       = "failed to create policy"
	ErrUpdatePolicy       = "failed to update policy"
	ErrDeletePolicy       = "failed to delete policy"
	ErrDeployPolicy       = "failed to deploy policy"
	ErrMarshalPolicy      = "failed to marshal policy data"
	ErrUnmarshalPolicy    = "failed to unmarshal policy data"
	ErrEmptySerialNumbers = "serial numbers cannot be empty for policy group"
	ErrInvalidResponse    = "invalid response from server"
	ErrPolicyNotFound     = "policy not found"
	ErrRollbackFailed     = "failed to rollback policy creation"
)

// Context keys for structured logging
const (
	CtxKeyTransactionID = "transaction_id"
	CtxKeyPolicyID      = "policy_id"
	CtxKeyIsPolicyGroup = "is_policy_group"
	CtxKeyDeployFlag    = "deploy_flag"
	CtxKeyNumSwitches   = "num_switches"
	CtxKeyPayloadSize   = "payload_size"
	CtxKeyError         = "error"
	CtxKeyResponse      = "response"
)

const ResourcePolicy = "policy"

// RscCreatePolicy creates a new policy or policy group in NDFC.
// It handles both the creation and optional deployment of the policy.
// For policy groups, it validates that serial numbers are provided.
// Returns any errors through the provided diagnostics.
func (c *NDFC) RscCreatePolicy(ctx context.Context, dg *diag.Diagnostics, model *resource_policy.PolicyModel) {
	// Start transaction logging
	txID := fmt.Sprintf("tx-create-%d", time.Now().UnixNano())
	tflog.Info(ctx, "Starting policy creation",
		map[string]interface{}{
			CtxKeyTransactionID: txID,
			CtxKeyIsPolicyGroup: model.IsPolicyGroup.ValueBool(),
		})

	// Initialize API client
	policyApi := api.NewPolicyAPI(c.GetLock(ResourcePolicy), &c.apiClient)
	inData := model.GetModelData()

	// Validate policy group configuration
	if inData.IsPolicyGroup {
		if len(inData.SerialNumbers) == 0 {
			errMsg := "serial numbers cannot be empty for policy group"
			tflog.Error(ctx, errMsg,
				map[string]interface{}{
					"transaction_id": txID,
				})
			dg.AddError("Invalid configuration", errMsg)
			return
		}
		policyApi.PolicyGroup = true
		policyApi.DeploySwitches = inData.SerialNumbers
		tflog.Debug(ctx, "Configured policy group",
			map[string]interface{}{
				CtxKeyTransactionID: txID,
				CtxKeyNumSwitches:   len(inData.SerialNumbers),
			})
	} else if len(inData.SerialNumbers) > 1 {
		errMsg := "multiple serial numbers not allowed for non-policy group"
		tflog.Error(ctx, errMsg,
			map[string]interface{}{
				CtxKeyTransactionID: txID,
			})
		dg.AddError("Invalid configuration", errMsg)
		return
	}

	// Prepare request data
	tflog.Debug(ctx, "Marshaling policy data",
		map[string]interface{}{
			CtxKeyTransactionID: txID,
		})

	data, err := json.Marshal(inData)
	if err != nil {
		errMsg := fmt.Sprintf("failed to marshal policy data: %v", err)
		tflog.Error(ctx, errMsg,
			map[string]interface{}{
				CtxKeyTransactionID: txID,
				CtxKeyError:         err.Error(),
			})
		dg.AddError("Failed to marshal policy data", errMsg)
		return
	}

	// Create policy
	tflog.Info(ctx, "Creating policy",
		map[string]interface{}{
			CtxKeyTransactionID: txID,
			CtxKeyPayloadSize:   len(data),
		})

	res, err := policyApi.Post(data)
	if err != nil {
		errMsg := fmt.Sprintf("%v:%s", err, res.String())
		tflog.Error(ctx, errMsg,
			map[string]interface{}{
				CtxKeyTransactionID: txID,
				CtxKeyError:         err.Error(),
				CtxKeyResponse:      res.String(),
			})
		dg.AddError("Failed to create policy", errMsg)
		return
	}

	// Extract policy ID from response
	policyId := getPolicyIdFromResponse(&res, inData.IsPolicyGroup)
	if policyId == "" {
		errMsg := "empty policy ID in create response"
		tflog.Error(ctx, errMsg,
			map[string]interface{}{
				CtxKeyTransactionID: txID,
				CtxKeyResponse:      res.String(),
			})
		dg.AddError("Failed to create policy", "Empty policy ID in response")
		return
	}

	tflog.Info(ctx, "Successfully created policy",
		map[string]interface{}{
			CtxKeyTransactionID: txID,
			CtxKeyPolicyID:      policyId,
		})

	// Handle deployment if requested
	if inData.Deploy {
		tflog.Info(ctx, "Initiating policy deployment",
			map[string]interface{}{
				CtxKeyTransactionID: txID,
				CtxKeyPolicyID:      policyId,
			})

		deployData := resource_policy.PolicyDeploy{
			PolicyId:     policyId,
			SerialNumber: inData.SerialNumbers,
			PolicyGroup:  inData.IsPolicyGroup,
			DeleteFlag:   false,
		}

		c.RscDeployPolicy(ctx, dg, deployData)
		if dg.HasError() {
			tflog.Error(ctx, "Policy created but deployment failed",
				map[string]interface{}{
					CtxKeyTransactionID: txID,
					CtxKeyPolicyID:      policyId,
				})
			// Rollback policy creation
			tflog.Info(ctx, "Initiating rollback of policy creation",
				map[string]interface{}{
					CtxKeyTransactionID: txID,
					CtxKeyPolicyID:      policyId,
				})
			inData.PolicyId = policyId
			delDg := diag.Diagnostics{}
			c.policyDelete(ctx, &delDg, inData)
			if delDg.HasError() {
				tflog.Error(ctx, "Failed to rollback policy creation",
					map[string]interface{}{
						CtxKeyTransactionID: txID,
						CtxKeyPolicyID:      policyId,
					})
				return
			}
			tflog.Error(ctx, fmt.Sprintf("Policy-id %s is deleted", policyId))
			return
		}
	}

	newModel := c.rscGetPolicy(ctx, dg, policyId)
	if newModel == nil {
		tflog.Error(ctx, "Failed to get policy")
		dg.AddError("Failed to get policy", "Failed to get policy")
		return
	}
	if inData.IsPolicyGroup {
		c.rscGetPolicyGroup(ctx, dg, policyId, newModel)
	}
	c.policyTrim(inData, newModel)
	model.SetModelData(newModel)
	model.Deploy = types.BoolValue(inData.Deploy)
}

/*
Policy Group
------------
NDFC is inconsistent in keeping the response format same.
It changes across releases
NDFC 12.1.3b {"successList":[{"name":"[9Q34PHYLDB5]","message":"POLICY-GROUP-1223200","status":"Success"}]}
NDFC 12.2.2 ( ND 3.2.1i) - {"successList":[{"name":"POLICY-GROUP-12942380","message":"Policy group is successfully created","status":"Success"}]}
Policy:
A copy of policy Model sent in POST
*/
func getPolicyIdFromResponse(res *gjson.Result, isPolicyGroup bool) string {
	if isPolicyGroup {
		pdId := res.Get("successList.0.message").String()
		// NDFC Inconsistency - in newer releases the format has been changed
		// Does pdId look anything like a policyGroup ID which is POLICY-GROUP-XXXXXX
		if strings.Contains(pdId, "POLICY-GROUP-") {
			return pdId
		}
		// Check another option
		pdId = res.Get("successList.0.name").String()
		if strings.Contains(pdId, "POLICY-GROUP-") {
			return pdId
		}
		return ""
	}
	return res.Get("policyId").String()
}

// RscReadPolicy reads a policy or policy group from NDFC
func (c *NDFC) RscReadPolicy(ctx context.Context, dg *diag.Diagnostics, model *resource_policy.PolicyModel) {
	txID := fmt.Sprintf("tx-read-%d", time.Now().UnixNano())
	deployFlag := model.Deploy.ValueBool()
	policyID := model.PolicyId.ValueString()

	tflog.Debug(ctx, "Reading policy",
		map[string]interface{}{
			CtxKeyTransactionID: txID,
			CtxKeyPolicyID:      policyID,
			CtxKeyDeployFlag:    deployFlag,
		})

	// Get the base policy data
	data := c.rscGetPolicy(ctx, dg, policyID)
	if data == nil {
		errMsg := fmt.Sprintf("%s: %s", ErrGetPolicy, policyID)
		tflog.Error(ctx, errMsg,
			map[string]interface{}{
				CtxKeyTransactionID: txID,
				CtxKeyPolicyID:      policyID,
			})
		dg.AddError(ErrGetPolicy, fmt.Sprintf("Failed to get policy %s", policyID))
		return
	}

	// Get additional details for policy groups
	if model.IsPolicyGroup.ValueBool() {
		tflog.Debug(ctx, "Fetching policy group details",
			map[string]interface{}{
				CtxKeyTransactionID: txID,
				CtxKeyPolicyID:      policyID,
			})

		c.rscGetPolicyGroup(ctx, dg, policyID, data)
		if dg.HasError() {
			tflog.Error(ctx, "Failed to get policy group details",
				map[string]interface{}{
					CtxKeyTransactionID: txID,
					CtxKeyPolicyID:      policyID,
				})
			return
		}
	}

	// Update the model with the retrieved data
	c.policyTrim(model.GetModelData(), data)
	data.Deploy = deployFlag
	model.SetModelData(data)

	tflog.Debug(ctx, "Successfully read policy",
		map[string]interface{}{
			CtxKeyTransactionID: txID,
			CtxKeyPolicyID:      policyID,
		})
}

// rscGetPolicy retrieves a policy by its ID
func (c *NDFC) rscGetPolicy(ctx context.Context, dg *diag.Diagnostics, pID string) *resource_policy.NDFCPolicyModel {
	txID, ok := ctx.Value(CtxKeyTransactionID).(string)
	if !ok {
		txID = fmt.Sprintf("tx-get-%d", time.Now().UnixNano())
	}

	tflog.Debug(ctx, "Retrieving policy",
		map[string]interface{}{
			CtxKeyTransactionID: txID,
			CtxKeyPolicyID:      pID,
		})

	model := new(resource_policy.NDFCPolicyModel)
	policyApi := api.NewPolicyAPI(c.GetLock(ResourcePolicy), &c.apiClient)
	policyApi.PolicyID = pID

	res, err := policyApi.Get()
	if err != nil {
		errMsg := fmt.Errorf("%s: %w", ErrGetPolicy, err).Error()
		tflog.Error(ctx, errMsg,
			map[string]interface{}{
				CtxKeyTransactionID: txID,
				CtxKeyPolicyID:      pID,
				CtxKeyError:         err.Error(),
			})
		dg.AddError(ErrGetPolicy, fmt.Sprintf("Error getting policy %s: %v", pID, err))
		return nil
	}

	tflog.Debug(ctx, "Received policy data",
		map[string]interface{}{
			CtxKeyTransactionID: txID,
			CtxKeyPolicyID:      pID,
			"response_size":     len(res),
		})

	if !json.Valid(res) {
		errMsg := fmt.Sprintf("%s: invalid JSON response", ErrGetPolicy)
		tflog.Error(ctx, errMsg,
			map[string]interface{}{
				CtxKeyTransactionID: txID,
				CtxKeyPolicyID:      pID,
				"response":          string(res),
			})

		if strings.Contains(string(res), "does not exist") {
			dg.AddError(ErrPolicyNotFound, fmt.Sprintf("Policy %s does not exist", pID))
			return nil
		}
		dg.AddError(ErrInvalidResponse, fmt.Sprintf("Invalid JSON response for policy %s", pID))
		return nil
	}

	if err := json.Unmarshal(res, &model); err != nil {
		errMsg := fmt.Errorf("%s: %w", ErrUnmarshalPolicy, err).Error()
		tflog.Error(ctx, errMsg,
			map[string]interface{}{
				CtxKeyTransactionID: txID,
				CtxKeyPolicyID:      pID,
				CtxKeyError:         err.Error(),
			})
		dg.AddError(ErrUnmarshalPolicy, fmt.Sprintf("Error unmarshaling policy %s: %v", pID, err))
		return nil
	}

	return model
}

func (c *NDFC) RscUpdatePolicy(ctx context.Context, dg *diag.Diagnostics, model *resource_policy.PolicyModel, stateModel *resource_policy.PolicyModel) {
	// Start transaction logging with a unique ID for this operation
	txID := fmt.Sprintf("tx-%d", time.Now().UnixNano())
	policyID := model.PolicyId.ValueString()

	tflog.Info(ctx, "Starting policy update operation",
		map[string]interface{}{
			"transaction_id":  txID,
			"policy_id":       policyID,
			"is_policy_group": model.IsPolicyGroup.ValueBool(),
		})

	// Log policy details before update
	tflog.Debug(ctx, "Policy update details",
		map[string]interface{}{
			"transaction_id": txID,
			"policy_id":      policyID,
			"deploy_flag":    model.Deploy.ValueBool(),
		})

	// NDFC inconsistency: PUT requires id field to be present
	// even though policy-id is present. Id is a redundant parameter here
	// Hence id has been added to schema and used only in PUT
	ID := model.Id.ValueInt64()

	// Prepare policy data for update
	policyData := model.GetModelData()
	policyData.PolicyId = policyID
	policyData.Id = &ID

	if model.IsPolicyGroup.ValueBool() {
		if len(policyData.SerialNumbers) == 0 {
			errMsg := "Serial numbers cannot be empty for policy group"
			tflog.Error(ctx, errMsg,
				map[string]interface{}{
					"transaction_id": txID,
					"policy_id":      policyID,
				})
			dg.AddError("Invalid configuration", errMsg)
			return
		}

		// Detect what changed between state and plan
		switchesChanged := !model.SerialNumbers.Equal(stateModel.SerialNumbers)
		nvPairsChanged := !model.PolicyParameters.Equal(stateModel.PolicyParameters)

		tflog.Info(ctx, "Policy group change detection",
			map[string]interface{}{
				"transaction_id":   txID,
				"policy_id":        policyID,
				"switches_changed": switchesChanged,
				"nv_pairs_changed": nvPairsChanged,
			})

		// Case 1: Switch list changed - GET current policy, PUT it back with new serial numbers in QP
		if switchesChanged {
			c.policyGroupUpdateSwitches(ctx, dg, txID, policyID, policyData)
			if dg.HasError() {
				return
			}

			if policyData.Deploy {
				// Deploy to union of old + new switches so removed switches get the deletion deployed
				stateData := stateModel.GetModelData()
				allSwitches := unionStrings(policyData.SerialNumbers, stateData.SerialNumbers)
				tflog.Debug(ctx, "Deploying to all affected switches (old + new)",
					map[string]interface{}{
						"transaction_id":  txID,
						"policy_id":       policyID,
						"deploy_switches": allSwitches,
					})
				savedSwitches := policyData.SerialNumbers
				policyData.SerialNumbers = allSwitches
				c.policyDeployAfterUpdate(ctx, dg, txID, policyID, policyData, model.IsPolicyGroup.ValueBool())
				policyData.SerialNumbers = savedSwitches
				if dg.HasError() {
					return
				}
			}
		}

		// Case 2: nvPairs changed - PUT with regular policy URL
		if nvPairsChanged {
			c.policyGroupUpdateNvPairs(ctx, dg, txID, policyID, policyData, ID)
			if dg.HasError() {
				return
			}

			if policyData.Deploy {
				c.policyDeployAfterUpdate(ctx, dg, txID, policyID, policyData, model.IsPolicyGroup.ValueBool())
				if dg.HasError() {
					return
				}
			}
		}
	} else {
		// Non-policy-group: single PUT with regular policy URL
		c.policyUpdateSingle(ctx, dg, txID, policyID, policyData)
		if dg.HasError() {
			return
		}

		if policyData.Deploy {
			c.policyDeployAfterUpdate(ctx, dg, txID, policyID, policyData, model.IsPolicyGroup.ValueBool())
			if dg.HasError() {
				return
			}
		}
	}

	// Refresh the policy data after update
	tflog.Debug(ctx, "Refreshing policy data",
		map[string]interface{}{
			"transaction_id": txID,
			"policy_id":      policyID,
		})

	newModel := c.rscGetPolicy(ctx, dg, policyID)
	if newModel == nil {
		errMsg := "Failed to refresh policy data after update"
		tflog.Error(ctx, errMsg,
			map[string]interface{}{
				"transaction_id": txID,
				"policy_id":      policyID,
			})
		dg.AddError("Failed to get policy", errMsg)
		return
	}
	if model.IsPolicyGroup.ValueBool() {
		c.rscGetPolicyGroup(ctx, dg, policyID, newModel)
		newModel.IsPolicyGroup = model.IsPolicyGroup.ValueBool()
		// NDFC policygroup GET may return stale switch list after update;
		// use the planned serial numbers since we just PUT them
		newModel.SerialNumbers = policyData.SerialNumbers
	}

	// Update the model with the refreshed data
	c.policyTrim(policyData, newModel)

	model.SetModelData(newModel)
	model.Deploy = types.BoolValue(policyData.Deploy)

	tflog.Info(ctx, "Successfully completed policy update operation",
		map[string]interface{}{
			CtxKeyTransactionID: txID,
			CtxKeyPolicyID:      policyID,
		})
}

// policyDeployAfterUpdate handles policy deployment after an update operation.
func (c *NDFC) policyDeployAfterUpdate(ctx context.Context, dg *diag.Diagnostics, txID, policyID string, policyData *resource_policy.NDFCPolicyModel, isPolicyGroup bool) {
	tflog.Info(ctx, "Initiating policy deployment",
		map[string]interface{}{
			"transaction_id": txID,
			"policy_id":      policyID,
		})

	deployData := resource_policy.PolicyDeploy{
		PolicyId:     policyID,
		SerialNumber: policyData.SerialNumbers,
		PolicyGroup:  isPolicyGroup,
		DeleteFlag:   true, // updates need full deployment to flush any removals
	}

	c.RscDeployPolicy(ctx, dg, deployData)
	if dg.HasError() {
		tflog.Error(ctx, "Policy update succeeded but deployment failed",
			map[string]interface{}{
				"transaction_id": txID,
				"policy_id":      policyID,
			})
		// cannot rollback to old config as the old data is overwritten in NDFC
		// throw error so that user can correct the config and re-apply
		dg.AddError("Failed to deploy policy", "Policy was updated but deployment failed. Please check the logs and retry.")
		return
	}
	tflog.Info(ctx, "Successfully deployed policy",
		map[string]interface{}{
			"transaction_id": txID,
			"policy_id":      policyID,
		})
}

// policyGroupUpdateSwitches handles the case where the switch list changed for a policy group.
// It GETs the current policy from NDFC and PUTs the same payload back with the new serial numbers in the query parameter.
func (c *NDFC) policyGroupUpdateSwitches(ctx context.Context, dg *diag.Diagnostics, txID, policyID string, policyData *resource_policy.NDFCPolicyModel) {
	tflog.Info(ctx, "Updating policy group switch membership",
		map[string]interface{}{
			"transaction_id": txID,
			"policy_id":      policyID,
			"new_switches":   policyData.SerialNumbers,
		})

	// GET the current policy data from NDFC
	currentPolicy := c.rscGetPolicy(ctx, dg, policyID)
	if currentPolicy == nil {
		dg.AddError("Failed to get policy for switch update",
			fmt.Sprintf("Cannot read current policy %s before updating switches", policyID))
		return
	}

	// Marshal the current policy data (unchanged) for the PUT body
	data, err := json.Marshal(currentPolicy)
	if err != nil {
		dg.AddError("Failed to marshal policy data",
			fmt.Sprintf("Failed to marshal policy data: %v", err))
		return
	}

	// PUT with policygroup URL using new serial numbers in QP (no mark-delete-and-update for switch-only update)
	policyApi := api.NewPolicyAPI(c.GetLock(ResourcePolicy), &c.apiClient)
	policyApi.PolicyID = policyID
	policyApi.PolicyGroup = true
	policyApi.SwitchOnlyUpdate = true
	policyApi.DeploySwitches = policyData.SerialNumbers

	log.Printf("[DEBUG] [%s] Updating policy group switches for %s: PUT Data |%s|", txID, policyID, string(data))

	res, err := policyApi.Put(data)
	if err != nil {
		dg.AddError("Failed to update policy group switches",
			fmt.Sprintf("Failed to update policy group switches: %v - Response: %s", err, res.String()))
		return
	}

	tflog.Info(ctx, "Successfully updated policy group switches",
		map[string]interface{}{
			"transaction_id": txID,
			"policy_id":      policyID,
		})
}

// policyGroupUpdateNvPairs handles the case where nvPairs changed for a policy group.
// It PUTs the updated policy data using the policygroup URL with mark-delete-and-update.
func (c *NDFC) policyGroupUpdateNvPairs(ctx context.Context, dg *diag.Diagnostics, txID, policyID string, policyData *resource_policy.NDFCPolicyModel, ID int64) {
	tflog.Info(ctx, "Updating policy group nvPairs",
		map[string]interface{}{
			"transaction_id": txID,
			"policy_id":      policyID,
		})

	policyData.Id = &ID
	data, err := json.Marshal(policyData)
	if err != nil {
		dg.AddError("Failed to marshal policy data",
			fmt.Sprintf("Failed to marshal policy data: %v", err))
		return
	}

	// Use policygroup URL with current serial numbers for nvPairs update
	policyApi := api.NewPolicyAPI(c.GetLock(ResourcePolicy), &c.apiClient)
	policyApi.PolicyID = policyID
	policyApi.PolicyGroup = true
	policyApi.DeploySwitches = policyData.SerialNumbers

	log.Printf("[DEBUG] [%s] Updating policy group nvPairs for %s: PUT Data |%s|", txID, policyID, string(data))

	res, err := policyApi.Put(data)
	if err != nil {
		dg.AddError("Failed to update policy nvPairs",
			fmt.Sprintf("Failed to update policy nvPairs: %v - Response: %s", err, res.String()))
		return
	}

	tflog.Info(ctx, "Successfully updated policy group nvPairs",
		map[string]interface{}{
			"transaction_id": txID,
			"policy_id":      policyID,
		})
}

// policyUpdateSingle handles the update for a non-policy-group policy (single PUT).
func (c *NDFC) policyUpdateSingle(ctx context.Context, dg *diag.Diagnostics, txID, policyID string, policyData *resource_policy.NDFCPolicyModel) {
	data, err := json.Marshal(policyData)
	if err != nil {
		dg.AddError("Failed to marshal policy data",
			fmt.Sprintf("Failed to marshal policy data: %v", err))
		return
	}

	policyApi := api.NewPolicyAPI(c.GetLock(ResourcePolicy), &c.apiClient)
	policyApi.PolicyID = policyID

	log.Printf("[DEBUG] [%s] Updating policy with ID: %v: PUT Data |%s|", txID, policyID, string(data))

	res, err := policyApi.Put(data)
	if err != nil {
		dg.AddError("Failed to update policy",
			fmt.Sprintf("Failed to update policy: %v - Response: %s", err, res.String()))
		return
	}

	tflog.Info(ctx, "Successfully updated policy",
		map[string]interface{}{
			"transaction_id": txID,
			"policy_id":      policyID,
		})
}

// RscDeletePolicy deletes a policy or policy group from NDFC
func (c *NDFC) RscDeletePolicy(ctx context.Context, dg *diag.Diagnostics, model *resource_policy.PolicyModel) {
	txID := fmt.Sprintf("tx-delete-%d", time.Now().UnixNano())
	policyID := model.PolicyId.ValueString()

	tflog.Info(ctx, "Starting policy deletion",
		map[string]interface{}{
			CtxKeyTransactionID: txID,
			CtxKeyPolicyID:      policyID,
			CtxKeyIsPolicyGroup: model.IsPolicyGroup.ValueBool(),
		})

	// Get serial numbers before we modify the model
	serialNumbers := model.GetModelData().SerialNumbers
	if len(serialNumbers) == 0 {
		errMsg := "no serial numbers found for policy deletion"
		tflog.Error(ctx, errMsg,
			map[string]interface{}{
				CtxKeyTransactionID: txID,
				CtxKeyPolicyID:      policyID,
			})
		dg.AddError("Invalid configuration", errMsg)
		return
	}

	fabricName := c.GetFabricName(ctx, serialNumbers[0])
	tflog.Debug(ctx, "Preparing to delete policy",
		map[string]interface{}{
			CtxKeyTransactionID: txID,
			CtxKeyPolicyID:      policyID,
			"fabric_name":       fabricName,
			"num_switches":      len(serialNumbers),
		})

	ID := model.Id.ValueInt64()

	// Initialize API client
	policyAPI := api.NewPolicyAPI(c.GetLock(ResourcePolicy), &c.apiClient)
	policyAPI.PolicyID = policyID

	// Prepare policy data for deletion
	policyData := model.GetModelData()
	policyData.Deleted = new(bool)
	*policyData.Deleted = true // Set delete flag
	policyData.PolicyId = policyID
	policyData.IsPolicyGroup = model.IsPolicyGroup.ValueBool()
	policyData.Id = &ID

	// Handle policy group specific configurations
	if policyData.IsPolicyGroup {
		policyAPI.PolicyGroup = true
		policyAPI.DeploySwitches = serialNumbers

		tflog.Debug(ctx, "Configuring policy group deletion",
			map[string]interface{}{
				CtxKeyTransactionID: txID,
				CtxKeyPolicyID:      policyID,
				CtxKeyNumSwitches:   len(serialNumbers),
			})
	}

	// Marshal policy data for deletion
	data, err := json.Marshal(policyData)
	if err != nil {
		errMsg := fmt.Errorf("%s: %w", ErrMarshalPolicy, err).Error()
		tflog.Error(ctx, errMsg,
			map[string]interface{}{
				CtxKeyTransactionID: txID,
				CtxKeyPolicyID:      policyID,
				CtxKeyError:         err.Error(),
			})
		dg.AddError(ErrMarshalPolicy, fmt.Sprintf("Error marshaling policy %s: %v", policyID, err))
		return
	}

	tflog.Debug(ctx, "Sending delete request",
		map[string]interface{}{
			CtxKeyTransactionID: txID,
			CtxKeyPolicyID:      policyID,
			"payload_size":      len(data),
		})

	// Execute delete operation (via PUT with deleted flag set )
	res, err := policyAPI.Put(data)
	if err != nil {
		log.Printf("[ERROR] Error deleting policy: %v, %v", err, res.String())
		tflog.Error(ctx, "Failed to mark delete policy - attempting full delete",
			map[string]interface{}{
				CtxKeyTransactionID: txID,
				CtxKeyPolicyID:      policyID,
				CtxKeyError:         err.Error(),
				CtxKeyResponse:      res.String(),
			})

		// As marking deleted failed - we must attempt a full sweep delete using the DELETE API
		// This is to avoid the policy being left stale in NDFC causing subsequent deployments to fail
		c.policyDelete(ctx, dg, policyData)
		tflog.Info(ctx, "Successfully force-deleted policy after mark delete failed",
			map[string]interface{}{
				CtxKeyTransactionID: txID,
				CtxKeyPolicyID:      policyID,
			})
	}

	tflog.Info(ctx, "Successfully marked policy for deletion",
		map[string]interface{}{
			CtxKeyTransactionID: txID,
			CtxKeyPolicyID:      policyID,
		})

	// Handle switch level deployment for the deleted policy
	tflog.Debug(ctx, "Initiating switch deployment for deleted policy",
		map[string]interface{}{
			CtxKeyTransactionID: txID,
			CtxKeyPolicyID:      policyID,
			"fabric_name":       fabricName,
		})

	switchDeployData := resource_policy.PolicyDeploy{
		PolicyId:     policyID,
		SerialNumber: serialNumbers,
		PolicyGroup:  policyData.IsPolicyGroup,
		DeleteFlag:   true,
	}

	// Use a separate diagnostics to track deployment errors without failing the operation
	switchDeployDg := diag.Diagnostics{}
	c.RscDeployPolicy(ctx, &switchDeployDg, switchDeployData)
	if switchDeployDg.HasError() {
		tflog.Warn(ctx, "Switch deployment completed with warnings",
			map[string]interface{}{
				CtxKeyTransactionID: txID,
				CtxKeyPolicyID:      policyID,
			})
	}

	// Verify the policy was actually deleted
	newDbg := diag.Diagnostics{}
	newModel := c.rscGetPolicy(ctx, &newDbg, policyID)
	if newModel != nil {
		// Check if deleted flag is set - if so NDFC will remove the policy - so terraform state must be removed
		if newModel.Deleted != nil && *newModel.Deleted {
			tflog.Debug(ctx, "Policy marked as deleted in NDFC",
				map[string]interface{}{
					CtxKeyTransactionID: txID,
					CtxKeyPolicyID:      policyID,
				})
			return
		}
		// Policy is not marked for delete - this is an error
		errMsg := fmt.Sprintf("Policy %s still exists and is not marked for deletion", policyID)
		tflog.Error(ctx, errMsg,
			map[string]interface{}{
				CtxKeyTransactionID: txID,
				CtxKeyPolicyID:      policyID,
			})
		// Attempt forced deletion
		c.policyDelete(ctx, dg, policyData)
		if dg.HasError() {
			tflog.Error(ctx, errMsg,
				map[string]interface{}{
					CtxKeyTransactionID: txID,
					CtxKeyPolicyID:      policyID,
					CtxKeyError:         err.Error(),
				})
			dg.AddError(ErrDeletePolicy, fmt.Sprintf("Error deleting policy %s: %v", policyID, err))
			return
		}
	}

	tflog.Info(ctx, "Successfully completed policy deletion",
		map[string]interface{}{
			CtxKeyTransactionID: txID,
			CtxKeyPolicyID:      policyID,
		})
}

// RscImportPolicy imports an existing policy into Terraform state
func (c *NDFC) RscImportPolicy(ctx context.Context, dg *diag.Diagnostics, ID string, model *resource_policy.PolicyModel) {
	txID := fmt.Sprintf("tx-import-%d", time.Now().UnixNano())
	tflog.Debug(ctx, fmt.Sprintf("RscImportPolicy: Importing policy ID %s", ID))
	data := c.rscGetPolicy(ctx, dg, ID)
	if data == nil {
		tflog.Error(ctx, "Failed to get policy")
		dg.AddError("Failed to get policy", "Failed to get policy")
		return
	}
	// Check if the imported policy is a policy group
	if strings.Contains(ID, "POLICY-GROUP") {
		data.IsPolicyGroup = true
		c.rscGetPolicyGroup(ctx, dg, ID, data)
		if dg.HasError() {
			tflog.Error(ctx, "Failed to get policy group details",
				map[string]interface{}{
					"transaction_id": txID,
					"policy_id":      ID,
				})
			return
		}
	} else {
		data.IsPolicyGroup = false
	}
	model.SetModelData(data)

}

func (c *NDFC) RscDeployPolicy(ctx context.Context, dg *diag.Diagnostics, deployData resource_policy.PolicyDeploy) {
	tflog.Debug(ctx, fmt.Sprintf("RscDeployPolicy: Deploying policy ID %s", deployData.PolicyId))

	// For delete, do switch level  deployment
	if deployData.DeleteFlag {
		fabricName := c.GetFabricName(ctx, deployData.SerialNumber[0])
		tflog.Debug(ctx, fmt.Sprintf("RscDeployPolicy: Deploying configuration for Fabric: %s, Serial Numbers: %s", fabricName, deployData.SerialNumber))
		c.RecalculateAndDeploy(ctx, dg, fabricName, true, true, deployData.SerialNumber)
		return
	}

	GlobalDeployLock("policy")
	defer GlobalDeployUnlock("policy")
	policyApi := api.NewPolicyAPI(c.GetLock(ResourcePolicy), &c.apiClient)
	policyApi.Deploy = true
	//policyApi.SetDeployLocked()
	if deployData.PolicyGroup {
		policyApi.PolicyGroup = true
		policyApi.DeploySwitches = deployData.SerialNumber
	}
	postData, err := json.Marshal([]string{deployData.PolicyId})
	if err != nil {
		tflog.Error(ctx, "Failed to marshal policy data")
		dg.AddError("Failed to marshal policy data", fmt.Sprintf("Error %v", err))
		return
	}
	startTime := time.Now()
	log.Printf("[DEBUG] Deploying policy with ID: %v: POST Data |%s|", deployData.PolicyId, string(postData))
	res, err := policyApi.DeployPost(postData)
	endTime := time.Now()
	if err != nil {
		tflog.Error(ctx, "Failed to deploy policy")
		c.fillDeployResponse(ctx, dg, &res, startTime, endTime)
		dg.AddError("Failed to deploy policy", fmt.Sprintf("Error %v: %v", err, res.String()))
		return
	}
	log.Printf("[TRACE] Policy deployment ID: %v Response: |%v|", deployData.PolicyId, res)

	if res.Get("0.failedPTIList").Exists() {
		// "failedPTIList" is present in the response
		// Add your code here
		log.Printf("[ERROR] Policy deployment failed for policy ID: %v - %v", deployData.PolicyId, res.Get("0.failedPTIList").String())
		dg.AddError("Failed to deploy policy", fmt.Sprintf("Error %v: %v", err, res.String()))
		c.fillDeployResponse(ctx, dg, &res, startTime, endTime)
		return
	}

}

// Response format:
// [{"failedPTIList":"POLICY-GROUP-1226830","switchSN":"9Q34PHYLDB5"}]
func (c *NDFC) fillDeployResponse(ctx context.Context, dg *diag.Diagnostics, res *gjson.Result, startTime, endTime time.Time) {
	// Parse the response to get failed serials
	failedSwitches := make([]string, 0)
	for i := 0; i < len(res.Array()); i++ {
		failedSwitches = append(failedSwitches, res.Get(fmt.Sprintf("%d.switchSN", i)).String())
	}
	if len(failedSwitches) > 0 {
		fabricName := c.GetFabricName(ctx, failedSwitches[0])
		startTimeFormatted := startTime.UTC().Format("2006-01-02 15:04:05")
		endTimeFormatted := endTime.UTC().Format("2006-01-02 15:04:05")
		log.Printf("[DEBUG] Getting deployment history for fabric: %v, switches: %v, start time: %v, end time: %v", fabricName, failedSwitches, startTimeFormatted, endTimeFormatted)
		resp, err := c.GetDeploymentHistoryWithFilters(ctx, fabricName, failedSwitches, "FAILED", "", startTimeFormatted, "")
		if err != nil {
			dg.AddError("Failed to get deployment history", fmt.Sprintf("Error: %v", err))
			return
		}
		if len(resp) > 0 {
			for _, deployResponse := range resp {
				dg.AddError("Failed to deploy policy", fmt.Sprintf("Switch %v: Status %s, error: %s", deployResponse.SerialNumber, deployResponse.Status, deployResponse.StatusDescription))
			}
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

// unionStrings returns the union of two string slices (no duplicates).
func unionStrings(a, b []string) []string {
	seen := make(map[string]struct{}, len(a)+len(b))
	result := make([]string, 0, len(a)+len(b))
	for _, s := range a {
		if _, ok := seen[s]; !ok {
			seen[s] = struct{}{}
			result = append(result, s)
		}
	}
	for _, s := range b {
		if _, ok := seen[s]; !ok {
			seen[s] = struct{}{}
			result = append(result, s)
		}
	}
	return result
}

func (c NDFC) policyTrim(pdata *resource_policy.NDFCPolicyModel, ndata *resource_policy.NDFCPolicyModel) {
	// Trim the data to remove any fields that are  returned by NDFC and not in policy_parameters
	// This is to avoid inconsistencies between the data in the state file and the data in NDFC
	ndata.ComputedParameters = make(map[string]string)

	for k := range ndata.PolicyParameters {
		if val, ok := pdata.PolicyParameters[k]; !ok {
			log.Printf("[DEBUG] Moving policy parameter: (%s,%s) as it was not in planned config", k, ndata.PolicyParameters[k])
			ndata.ComputedParameters[k] = val
			delete(ndata.PolicyParameters, k)
		}
	}
}

func (c NDFC) policyDelete(ctx context.Context, dg *diag.Diagnostics, policy *resource_policy.NDFCPolicyModel) {
	papi := api.NewPolicyAPI(c.GetLock(ResourcePolicy), &c.apiClient)
	papi.PolicyID = policy.PolicyId
	if policy.IsPolicyGroup {
		papi.PolicyGroup = true
		papi.DeploySwitches = policy.SerialNumbers
	}
	res, err := papi.Delete()
	if err != nil {
		errMsg := fmt.Sprintf("Failed to force delete policy Id %s: %v | %v", policy.PolicyId, err, res.String())
		tflog.Error(ctx, errMsg,
			map[string]interface{}{
				CtxKeyPolicyID: policy.PolicyId,
				CtxKeyError:    err.Error(),
				CtxKeyResponse: res.String(),
			})
		dg.AddError(ErrDeletePolicy, errMsg)
		return
	}
	log.Printf("[TRACE] Force deleted policy Id %s", policy.PolicyId)
}

// rscGetPolicyGroup retrieves policy group details including associated switches
func (c *NDFC) rscGetPolicyGroup(ctx context.Context, dg *diag.Diagnostics, pID string, model *resource_policy.NDFCPolicyModel) {
	txID, ok := ctx.Value(CtxKeyTransactionID).(string)
	if !ok {
		txID = fmt.Sprintf("tx-group-%d", time.Now().UnixNano())
	}

	tflog.Debug(ctx, "Retrieving policy group details",
		map[string]interface{}{
			CtxKeyTransactionID: txID,
			CtxKeyPolicyID:      pID,
		})

	policyAPI := api.NewPolicyAPI(c.GetLock(ResourcePolicy), &c.apiClient)
	policyAPI.PolicyID = pID
	policyAPI.PolicyGroup = true

	res, err := policyAPI.Get()
	if err != nil {
		errMsg := fmt.Errorf("%s: %w", "failed to get policy group", err).Error()
		tflog.Error(ctx, errMsg,
			map[string]interface{}{
				CtxKeyTransactionID: txID,
				CtxKeyPolicyID:      pID,
				CtxKeyError:         err.Error(),
			})
		dg.AddError("Failed to get policy group", fmt.Sprintf("Error getting policy group %s: %v", pID, err))
		return
	}

	if !json.Valid(res) {
		errMsg := fmt.Sprintf("%s: invalid JSON response", "failed to get policy group")
		tflog.Error(ctx, errMsg,
			map[string]interface{}{
				CtxKeyTransactionID: txID,
				CtxKeyPolicyID:      pID,
				"response":          string(res),
			})

		if strings.Contains(string(res), "does not exist") {
			dg.AddError(ErrPolicyNotFound, fmt.Sprintf("Policy group %s does not exist", pID))
			return
		}
		dg.AddError(ErrInvalidResponse, fmt.Sprintf("Invalid JSON response for policy group %s", pID))
		return
	}

	switches := make([]string, 0)
	if err := json.Unmarshal(res, &switches); err != nil {
		errMsg := fmt.Errorf("%s: %w", "failed to unmarshal policy group data", err).Error()
		tflog.Error(ctx, errMsg,
			map[string]interface{}{
				CtxKeyTransactionID: txID,
				CtxKeyPolicyID:      pID,
				CtxKeyError:         err.Error(),
			})
		dg.AddError("Failed to unmarshal policy group data",
			fmt.Sprintf("Error unmarshaling policy group %s: %v", pID, err))
		return
	}

	model.SerialNumbers = custtypes.CSVString(switches)
	model.IsPolicyGroup = true

	tflog.Debug(ctx, "Retrieved policy group details",
		map[string]interface{}{
			CtxKeyTransactionID: txID,
			CtxKeyPolicyID:      pID,
			"num_switches":      len(switches),
		})
}
