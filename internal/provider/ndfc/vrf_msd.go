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
	"strings"
	rva "terraform-provider-ndfc/internal/provider/resources/resource_vrf_attachments"
	"terraform-provider-ndfc/internal/provider/resources/resource_vrf_bulk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// RscCreateBulkVrfMsd handles VRF creation for both MSD parent and MSD child fabrics
// The isMsdParent and isMsdChild flags determine the behavior:
//
// MSD Parent fabric VRF flow (isMsdParent=true):
// 1. VRFs are created and owned by parent fabric
// 2. VRFs that exist in NDFC are updated with new config (VRF ID preserved)
// 3. VRFs that don't exist in NDFC are created
// 4. Only parent-managed properties are allowed, setting child-specific properties will error
// 5. Attachments span multiple child fabrics
//
// MSD Child fabric VRF flow (isMsdChild=true):
// 1. VRFs are created and owned by parent fabric (via parent resource)
// 2. Update child-specific properties using child fabric name in API URL
// 3. Only child properties (TRM, netflow, advertise routes, etc.) are updated
// 4. Parent properties must match parent fabric configuration
// 5. Attachments are created on child fabric
func (c NDFC) RscCreateBulkVrfMsd(ctx context.Context, dg *diag.Diagnostics, vrfBulk *resource_vrf_bulk.VrfBulkModel, isMsdParent bool, isMsdChild bool, parentFabric string) *resource_vrf_bulk.VrfBulkModel {
	fabricName := vrfBulk.FabricName.ValueString()

	if isMsdParent {
		tflog.Debug(ctx, fmt.Sprintf("RscCreateBulkVrfMsd entry for MSD Parent fabric %s", fabricName))
	} else if isMsdChild {
		tflog.Debug(ctx, fmt.Sprintf("RscCreateBulkVrfMsd entry for MSD Child fabric %s (parent: %s)", fabricName, parentFabric))
	}

	vrf := vrfBulk.GetModelData()
	if vrf == nil {
		tflog.Error(ctx, "Data conversion from model failed")
		dg.AddError("Data conversion from model failed", "GetModelData returned empty")
		return nil
	}

	exists := c.FabricExists(ctx, vrf.FabricName)
	if !exists {
		errMsg := fmt.Sprintf("Fabric %s does not exist in NDFC", vrf.FabricName)
		tflog.Error(ctx, errMsg)
		dg.AddError("Fabric not found", errMsg)
		return nil
	}

	// Form ID
	ID := c.VrfBulkCreateID(vrf)

	// Validate attachment fabric field based on fabric type
	c.validateAttachmentFabricField(ctx, dg, vrf, isMsdParent, isMsdChild)
	if dg.HasError() {
		return nil
	}

	// Handle VRF operations based on fabric type
	if isMsdParent {
		// Handle MSD parent VRF operations (create/update VRFs)
		err := c.handleMsdParentVrfOperations(ctx, dg, vrf)
		if err != nil {
			return nil
		}
		tflog.Info(ctx, fmt.Sprintf("MSD Parent: VRF operations completed successfully ID %s", ID))
	} else if isMsdChild {
		// Handle MSD child VRF operations (update child-specific properties)
		err := c.handleMsdChildVrfOperations(ctx, dg, vrf, parentFabric)
		if err != nil {
			return nil
		}
		tflog.Info(ctx, fmt.Sprintf("MSD Child: VRF operations completed successfully ID %s", ID))
	}

	// Part 2: Create Attachments if any
	keyMap := c.vrfAttachmentSerialRemap(ctx, vrf)
	va := vrf.FillAttachPayloadFromModel(false)

	if len(va.VrfAttachments) > 0 {
		var err error
		if isMsdParent {
			// For MSD parent fabrics, attachments span multiple child fabrics
			err = c.RscManageVrfAttachmentsForMsdParent(ctx, dg, vrf, vrf.FabricName, false)
		} else if isMsdChild {
			// MSD child fabric: all attachments go to the same fabric
			err = c.RscCreateVrfAttachments(ctx, dg, va)
		}

		if err != nil {
			tflog.Error(ctx, "VRF Attachments create failed")
			tflog.Error(ctx, "Rolling back the configurations...delete VRFs")
			c.RscDeleteBulkVrf(ctx, dg, ID, vrfBulk)
			return nil
		}
		// Check and deploy
		c.RscDeployVrfAttachments(ctx, dg, vrf)
	}

	outVrf := c.RscGetBulkVrf(ctx, dg, ID, &va.DepMap, &keyMap)
	if outVrf == nil {
		tflog.Error(ctx, "Failed to verify: Reading from NDFC after create failed")
		dg.AddError("Failed to verify", "Reading from NDFC after create failed")
		return nil
	}
	if ID != "" {
		outVrf.Id = types.StringValue(ID)
	}
	return outVrf
}

// RscUpdateBulkVrfMsd handles VRF update for both MSD parent and MSD child fabrics
// The isMsdParent and isMsdChild flags determine the behavior:
//
// MSD Parent fabric VRF update flow (isMsdParent=true):
// 1. VRFs are owned by parent fabric
// 2. VRFs that exist are updated, new VRFs are created
// 3. Only parent-managed properties are allowed
// 4. Attachments span multiple child fabrics
//
// MSD Child fabric VRF update flow (isMsdChild=true):
// 1. VRFs are owned by parent fabric
// 2. Only child-specific properties can be updated
// 3. Parent properties must match parent fabric configuration
// 4. Attachments are managed on child fabric
// 5. VRF delete is a no-op (parent manages VRF lifecycle)
func (c NDFC) RscUpdateBulkVrfMsd(ctx context.Context,
	dg *diag.Diagnostics, ID string,
	vrfBulkPlan *resource_vrf_bulk.VrfBulkModel,
	vrfState *resource_vrf_bulk.VrfBulkModel, vrfConfig *resource_vrf_bulk.VrfBulkModel,
	isMsdParent bool, isMsdChild bool, parentFabric string) {

	fabricName := vrfBulkPlan.FabricName.ValueString()

	if isMsdParent {
		tflog.Debug(ctx, fmt.Sprintf("RscUpdateBulkVrfMsd entry for MSD Parent fabric %s", fabricName))
	} else if isMsdChild {
		tflog.Debug(ctx, fmt.Sprintf("RscUpdateBulkVrfMsd entry for MSD Child fabric %s (parent: %s)", fabricName, parentFabric))
	}

	c.validateVrfsUpdate(ctx, dg, vrfBulkPlan, vrfState)
	if dg.HasError() {
		return
	}

	// Get plan model for validation
	planModel := vrfBulkPlan.GetModelData()
	if planModel == nil {
		tflog.Error(ctx, "Data conversion from model failed")
		dg.AddError("Data conversion from model failed", "GetModelData returned empty")
		return
	}

	// Validate attachment fabric field based on fabric type
	c.validateAttachmentFabricField(ctx, dg, planModel, isMsdParent, isMsdChild)
	if dg.HasError() {
		return
	}

	// For MSD parent, validate that only parent properties are set
	if isMsdParent {
		c.ValidateMsdParentVrfParameters(ctx, dg, planModel)
		if dg.HasError() {
			return
		}
		tflog.Debug(ctx, fmt.Sprintf("Fabric %s is MSD parent - validated parent properties only", fabricName))
	}

	// Get diff actions
	actions := c.vrfBulkGetDiff(ctx, vrfBulkPlan, vrfState, vrfConfig)

	delVrfs := actions["del"].(*resource_vrf_bulk.NDFCVrfBulkModel)
	putVrfs := actions["put"].(*resource_vrf_bulk.NDFCVrfBulkModel)
	newVrfs := actions["add"].(*resource_vrf_bulk.NDFCVrfBulkModel)
	plan := actions["plan"].(*resource_vrf_bulk.NDFCVrfBulkModel)
	state := actions["state"].(*resource_vrf_bulk.NDFCVrfBulkModel)

	ndfcVRFs, err := c.vrfBulkGet(ctx, fabricName)
	if err != nil {
		tflog.Error(ctx, "Failed to Read existing VRFs")
		dg.AddError("VRF Read Failed", err.Error())
		return
	}

	// Step 1 - Check if VRFs to update are present in NDFC
	for vrf := range putVrfs.Vrfs {
		_, ok := ndfcVRFs.Vrfs[vrf]
		if !ok {
			errString := fmt.Sprintf("VRF %s to update is missing in NDFC", vrf)
			tflog.Error(ctx, errString)
			dg.AddError("In place update failed", errString)
			return
		}
	}

	// Step 2 - Handle VRFs to delete based on fabric type
	if isMsdParent {
		// MSD Parent: VRFs to delete are available
		for vrf := range delVrfs.Vrfs {
			_, ok := ndfcVRFs.Vrfs[vrf]
			if !ok {
				delete(delVrfs.Vrfs, vrf)
				tflog.Error(ctx, fmt.Sprintf("VRF to DELETE %s is missing in NDFC", vrf))
			}
		}
	} else if isMsdChild {
		// MSD Child: VRF delete is skipped (parent manages VRF lifecycle)
		if len(delVrfs.Vrfs) > 0 {
			tflog.Info(ctx, fmt.Sprintf("MSD Child fabric: Skipping VRF delete for %v - parent manages VRF lifecycle", delVrfs.GetVrfNames()))
		}
	}

	// Step 3 - Check if VRFs to be created are present in NDFC
	for vrf := range newVrfs.Vrfs {
		ndfcVrf, ok := ndfcVRFs.Vrfs[vrf]
		if ok {
			if _, ok := delVrfs.Vrfs[vrf]; ok {
				tflog.Debug(ctx, "VRF marked for delete", map[string]interface{}{"vrfName": vrf})
				continue
			}
			if isMsdParent {
				tflog.Info(ctx, fmt.Sprintf("MSD Parent: VRF %s already exists in NDFC, moving to update list", vrf))
			} else {
				tflog.Info(ctx, fmt.Sprintf("MSD Child: VRF %s already exists in NDFC (from parent), moving to update list", vrf))
			}
			vrfEntry := newVrfs.Vrfs[vrf]
			vrfEntry.VrfId = ndfcVrf.VrfId
			putVrfs.Vrfs[vrf] = vrfEntry
			delete(newVrfs.Vrfs, vrf)
		}
	}

	// For MSD child, validate parent properties for all VRFs to be updated
	if isMsdChild && len(putVrfs.Vrfs) > 0 {
		tflog.Info(ctx, "MSD Child Update: Validating parent properties match NDFC")
		c.validateVrfsMsdParentPropertiesInChild(ctx, dg, putVrfs, ndfcVRFs)
		if dg.HasError() {
			return
		}
	}

	// Begin update - No rollback from here on
	if isMsdParent {
		// MSD Parent: Handle delete operations
		if len(delVrfs.Vrfs) > 0 {
			// MSD parent: attachments span multiple child fabrics
			err := c.RscManageVrfAttachmentsForMsdParent(ctx, dg, delVrfs, fabricName, true)
			if err != nil {
				tflog.Error(ctx, "MSD Parent: VRF Attachments delete failed")
				dg.AddError("VRF Attachments delete failed", err.Error())
				return
			}
			tflog.Info(ctx, fmt.Sprintf("MSD Parent: Deleting VRFs %v, as part of Bulk Update", delVrfs.GetVrfNames()))
			err = c.vrfBulkDelete(ctx, fabricName, delVrfs.GetVrfNames())
			if err != nil {
				dg.AddError("Bulk Delete failed", err.Error())
			}
		} else {
			tflog.Info(ctx, "Nothing to delete")
		}
	} else if isMsdChild {
		// MSD Child: Skip attachment and VRF delete
		tflog.Info(ctx, fmt.Sprintf("MSD Child fabric: Skipping attachment and VRF delete for child fabric %s", fabricName))
	}

	// Modify VRFs
	if len(putVrfs.Vrfs) > 0 {
		if isMsdParent {
			tflog.Info(ctx, "MSD Parent: Modifying VRFs, as part of Bulk Update")
		} else {
			tflog.Info(ctx, "MSD Child: Modifying VRFs, as part of Bulk Update")
		}
		c.vrfBulkUpdate(ctx, dg, putVrfs)
		if dg.HasError() {
			tflog.Info(ctx, fmt.Sprintf("Modifying VRFs Failed %v", dg.Errors()))
			return
		}
	} else {
		tflog.Info(ctx, "Nothing to modify")
	}

	// Handle new VRFs based on fabric type
	if len(newVrfs.Vrfs) > 0 {
		if isMsdParent {
			// MSD Parent: Can create new VRFs
			tflog.Info(ctx, "MSD Parent: Adding VRFs, as part of Bulk Update")
			newVrfPayload := newVrfs.FillVrfPayloadFromModel(nil)
			err := c.vrfCreateBulk(ctx, fabricName, newVrfPayload)
			if err != nil {
				dg.AddError("VRF create failed", err.Error())
				return
			}
		} else if isMsdChild {
			// MSD Child: VRF creation is not allowed (must be done from parent)
			errString := fmt.Sprintf("MSD Child: Cannot create VRFs %v - VRFs must be created from parent fabric", newVrfs.GetVrfNames())
			tflog.Error(ctx, errString)
			dg.AddError("VRF creation not allowed on MSD child", errString)
			return
		}
	}

	// Handle attachment updates based on fabric type
	if isMsdParent {
		// MSD Parent: Handle attachments spanning multiple child fabrics
		c.RscUpdateVrfAttachments(ctx, dg, plan, state)
		if dg.HasError() {
			tflog.Error(ctx, "Error during update")
			return
		}
	} else if isMsdChild {
		// MSD Child: Skip VRF attachment updates - parent manages all attachments
		tflog.Info(ctx, "MSD Child: Skipping VRF attachment updates - parent manages all attachments")
	}

	keyMap := c.vrfAttachmentSerialRemap(ctx, plan)
	newID := c.VrfBulkCreateID(plan)
	depMap := FillDeployMap(plan)

	*(vrfBulkPlan) = *(c.RscGetBulkVrf(ctx, dg, newID, &depMap, &keyMap))
}

// RscDeleteBulkVrfMsd handles VRF deletion for both MSD parent and MSD child fabrics
// The isMsdParent and isMsdChild flags determine the behavior:
//
// MSD Parent fabric VRF delete flow (isMsdParent=true):
// 1. Delete attachments from all child fabrics
// 2. Delete VRFs from parent fabric
//
// MSD Child fabric VRF delete flow (isMsdChild=true):
// 1. Delete is a no-op - parent fabric manages VRF lifecycle
// 2. VRFs and attachments are deleted when parent resource is destroyed
func (c NDFC) RscDeleteBulkVrfMsd(ctx context.Context, dg *diag.Diagnostics, ID string, vrfBulk *resource_vrf_bulk.VrfBulkModel, isMsdParent bool, isMsdChild bool, parentFabric string) {
	fabricName := ""
	results := c.RscBulkSplitID(ID)
	if len(results["fabric"]) > 0 {
		fabricName = results["fabric"][0]
	}

	if isMsdParent {
		tflog.Info(ctx, fmt.Sprintf("RscDeleteBulkVrfMsd: MSD Parent fabric %s delete request %s", fabricName, ID))
	} else if isMsdChild {
		// For MSD child fabrics, delete is a no-op
		// The parent resource manages deletion of VRFs and attachments
		tflog.Info(ctx, fmt.Sprintf("RscDeleteBulkVrfMsd: MSD Child fabric %s (parent: %s) delete request %s", fabricName, parentFabric, ID))
		tflog.Info(ctx, fmt.Sprintf("MSD Child fabric detected. Delete is no-op. Parent fabric %s manages VRF deletion.", parentFabric))
		return
	}

	// MSD Parent: Proceed with deletion
	vrfs, err := c.VrfBulkIsPresent(ctx, ID)
	if err != nil {
		tflog.Error(ctx, "Failed to Read existing VRFs")
		dg.AddError("VRF Read Failed", err.Error())
		return
	}

	vrfsFromId := results["rsc"]
	if len(vrfs) != len(vrfsFromId) {
		errString := fmt.Sprintf("Mismatch in VRF data: fabric %s Read %v, from ID %v", fabricName, vrfs, vrfsFromId)
		tflog.Error(ctx, errString)
		dg.AddWarning("Mismatch in VRF data", errString)
		vrfsFromId = vrfs
	}

	// Delete attachments from child fabrics
	delVrf := vrfBulk.GetModelData()

	tflog.Info(ctx, "MSD Parent: Deleting attachments from child fabrics")
	err = c.RscManageVrfAttachmentsForMsdParent(ctx, dg, delVrf, fabricName, true)
	if err != nil {
		tflog.Error(ctx, "MSD Parent: VRF Attachments delete failed")
		dg.AddError("VRF Attachments delete failed", err.Error())
		return
	}

	// Delete the VRFs
	err = c.vrfBulkDelete(ctx, fabricName, vrfsFromId)
	if err != nil {
		errString := fmt.Sprintf("VRF delete failed on fabric %s vrfs %v", fabricName, vrfsFromId)
		tflog.Error(ctx, "VRF delete failed")
		dg.AddError(errString, err.Error())
		return
	}

	tflog.Info(ctx, fmt.Sprintf("MSD Parent: VRF Bulk Delete Success %s", ID))
}

// handleMsdParentVrfOperations handles VRF creation/update for MSD parent fabric
// MSD Parent fabric VRF flow:
// 1. VRFs are created and owned by parent fabric
// 2. VRFs that exist in NDFC are updated with new config (VRF ID preserved)
// 3. VRFs that don't exist in NDFC are created
// 4. Only parent-managed properties are allowed
func (c NDFC) handleMsdParentVrfOperations(ctx context.Context, dg *diag.Diagnostics, vrf *resource_vrf_bulk.NDFCVrfBulkModel) error {
	tflog.Info(ctx, fmt.Sprintf("MSD Parent fabric detected: %s", vrf.FabricName))

	// Validate that only parent properties are set
	c.ValidateMsdParentVrfParameters(ctx, dg, vrf)
	if dg.HasError() {
		return fmt.Errorf("MSD parent validation failed")
	}

	tflog.Debug(ctx, fmt.Sprintf("Fabric %s is MSD parent - validated parent properties only", vrf.FabricName))

	// Get existing VRFs from NDFC
	ndfcVRFs, err := c.vrfBulkGet(ctx, vrf.FabricName)
	if err != nil {
		tflog.Error(ctx, "Error while getting VRFs for MSD Parent fabric", map[string]interface{}{"Err": err})
		dg.AddError("VRF Read Failed", err.Error())
		return err
	}

	// Separate VRFs into those that exist (need update) and those that don't (need create)
	existingVrfs := resource_vrf_bulk.NDFCVrfBulkModel{
		FabricName: vrf.FabricName,
		Vrfs:       make(map[string]resource_vrf_bulk.NDFCVrfsValue),
	}
	newVrfs := resource_vrf_bulk.NDFCVrfBulkModel{
		FabricName: vrf.FabricName,
		Vrfs:       make(map[string]resource_vrf_bulk.NDFCVrfsValue),
	}

	for vrfName, vrfEntry := range vrf.Vrfs {
		if ndfcVrfEntry, found := ndfcVRFs.Vrfs[vrfName]; found {
			// VRF exists - add to update list and preserve VRF ID
			vrfEntry.VrfId = ndfcVrfEntry.VrfId
			existingVrfs.Vrfs[vrfName] = vrfEntry
			tflog.Debug(ctx, fmt.Sprintf("MSD Parent: VRF %s already exists, will update it", vrfName))
		} else {
			// VRF doesn't exist - add to create list
			newVrfs.Vrfs[vrfName] = vrfEntry
			tflog.Debug(ctx, fmt.Sprintf("MSD Parent: VRF %s doesn't exist, will create it", vrfName))
		}
	}

	// Update existing VRFs
	if len(existingVrfs.Vrfs) > 0 {
		tflog.Info(ctx, fmt.Sprintf("MSD Parent: Updating %d existing VRFs", len(existingVrfs.Vrfs)))
		c.vrfBulkUpdate(ctx, dg, &existingVrfs)
		if dg.HasError() {
			tflog.Error(ctx, "MSD Parent: Cannot update existing VRFs")
			return fmt.Errorf("MSD parent VRF update failed")
		}
		tflog.Info(ctx, "MSD Parent: Update existing VRFs success")
	}

	// Create new VRFs
	if len(newVrfs.Vrfs) > 0 {
		tflog.Info(ctx, fmt.Sprintf("MSD Parent: Creating %d new VRFs", len(newVrfs.Vrfs)))
		ndfcVrfBulkPayload := newVrfs.FillVrfPayloadFromModel(nil)

		err = c.vrfCreateBulk(ctx, vrf.FabricName, ndfcVrfBulkPayload)
		if err != nil {
			tflog.Error(ctx, "MSD Parent: Cannot create VRF", map[string]interface{}{"Err": err})
			dg.AddError("Cannot create VRF ", err.Error())
			return err
		}
		tflog.Info(ctx, "MSD Parent: Create new VRFs success")
	}

	tflog.Info(ctx, "MSD Parent: VRF operations completed successfully")
	return nil
}

// handleMsdChildVrfOperations handles VRF update for MSD child fabric
// MSD Child fabric VRF flow:
// 1. VRFs are created and owned by parent fabric (via parent resource)
// 2. Update child-specific properties using child fabric name in API URL
// 3. Only child properties (TRM, netflow, advertise routes, etc.) are updated
// 4. Parent properties must match parent fabric configuration
func (c NDFC) handleMsdChildVrfOperations(ctx context.Context, dg *diag.Diagnostics, vrf *resource_vrf_bulk.NDFCVrfBulkModel, parentFabric string) error {
	tflog.Info(ctx, fmt.Sprintf("MSD Child fabric detected: %s (parent: %s)", vrf.FabricName, parentFabric))

	// Query child fabric to get existing VRFs
	ndfcVRFs, err := c.vrfBulkGet(ctx, vrf.FabricName)
	if err != nil {
		tflog.Error(ctx, "Error while getting existing VRFs for MSD Child fabric", map[string]interface{}{"Err": err})
		dg.AddError("VRF Read Failed", err.Error())
		return err
	}

	// Check all vrfs in config are present before we allow in-place update
	var missingVrfs []string
	for vrfName := range vrf.Vrfs {
		if _, found := ndfcVRFs.Vrfs[vrfName]; !found {
			missingVrfs = append(missingVrfs, vrfName)
		} else {
			vrfEntry := vrf.Vrfs[vrfName]
			vrfEntry.VrfId = ndfcVRFs.Vrfs[vrfName].VrfId
			vrf.Vrfs[vrfName] = vrfEntry
		}
	}
	if len(missingVrfs) > 0 {
		errString := fmt.Sprintf("Missing VRF(s) %v in NDFC for child fabric update %s", missingVrfs, parentFabric)
		tflog.Error(ctx, errString)
		dg.AddError("VRFs missing", errString)
		return fmt.Errorf("%s", errString)
	}

	// Validate parent properties in child config match what NDFC returns.
	// The child fabric GET already returns parent properties from the parent fabric,
	// so we validate against ndfcVRFs (no need for separate parent fabric GET).
	tflog.Info(ctx, "MSD Child: Validating parent properties match NDFC")
	c.validateVrfsMsdParentPropertiesInChild(ctx, dg, vrf, ndfcVRFs)
	if dg.HasError() {
		return fmt.Errorf("MSD child parent property validation failed")
	}

	// Update VRFs in parent fabric using child fabric name for API URL
	c.vrfBulkUpdate(ctx, dg, vrf)
	if dg.HasError() {
		tflog.Info(ctx, fmt.Sprintf("Modifying MSD child VRFs Failed %v", dg.Errors()))
		return fmt.Errorf("MSD child VRF update failed")
	}

	// // Copy all parent properties from NDFC to config
	// // This ensures parent properties are properly set in state
	// for vrfName := range vrf.Vrfs {
	// 	if parentEntry, found := ndfcVRFs.Vrfs[vrfName]; found {
	// 		vrfEntry := vrf.Vrfs[vrfName]
	// 		// Copy all parent-managed properties from parent fabric
	// 		vrfEntry.VrfId = parentEntry.VrfId
	// 		vrfEntry.VrfTemplate = parentEntry.VrfTemplate
	// 		vrfEntry.VrfExtensionTemplate = parentEntry.VrfExtensionTemplate
	// 		vrfEntry.VrfTemplateConfig.VlanId = parentEntry.VrfTemplateConfig.VlanId
	// 		vrfEntry.VrfTemplateConfig.VlanName = parentEntry.VrfTemplateConfig.VlanName
	// 		vrfEntry.VrfTemplateConfig.InterfaceDescription = parentEntry.VrfTemplateConfig.InterfaceDescription
	// 		vrfEntry.VrfTemplateConfig.VrfDescription = parentEntry.VrfTemplateConfig.VrfDescription
	// 		vrfEntry.VrfTemplateConfig.Mtu = parentEntry.VrfTemplateConfig.Mtu
	// 		vrfEntry.VrfTemplateConfig.LoopbackRoutingTag = parentEntry.VrfTemplateConfig.LoopbackRoutingTag
	// 		vrfEntry.VrfTemplateConfig.RouteTargetImport = parentEntry.VrfTemplateConfig.RouteTargetImport
	// 		vrfEntry.VrfTemplateConfig.RouteTargetExport = parentEntry.VrfTemplateConfig.RouteTargetExport
	// 		vrfEntry.VrfTemplateConfig.RouteTargetImportEvpn = parentEntry.VrfTemplateConfig.RouteTargetImportEvpn
	// 		vrfEntry.VrfTemplateConfig.RouteTargetExportEvpn = parentEntry.VrfTemplateConfig.RouteTargetExportEvpn
	// 		vrfEntry.VrfTemplateConfig.RpLoopbackId = parentEntry.VrfTemplateConfig.RpLoopbackId
	// 		vrfEntry.VrfTemplateConfig.MaxBgpPaths = parentEntry.VrfTemplateConfig.MaxBgpPaths
	// 		vrfEntry.VrfTemplateConfig.MaxIbgpPaths = parentEntry.VrfTemplateConfig.MaxIbgpPaths
	// 		vrfEntry.VrfTemplateConfig.Ipv6LinkLocal = parentEntry.VrfTemplateConfig.Ipv6LinkLocal
	// 		vrfEntry.VrfTemplateConfig.DisableRtAuto = parentEntry.VrfTemplateConfig.DisableRtAuto
	// 		// Child-specific properties remain from child config
	// 		vrf.Vrfs[vrfName] = vrfEntry
	// 	}
	//}

	tflog.Info(ctx, "MSD Child: VRF update completed successfully")
	return nil
}

// validateAttachmentFabricField validates the fabric field in attachments based on fabric type:
// - MSD Parent: fabric field REQUIRED (specifies which child fabric the attachment belongs to)
// - MSD Child: fabric field should NOT be set (computed from resource's fabric_name)
// - Regular: fabric field should NOT be set (computed from resource's fabric_name)
func (c NDFC) validateAttachmentFabricField(ctx context.Context, dg *diag.Diagnostics, vrf *resource_vrf_bulk.NDFCVrfBulkModel, isMsdParent bool, isMsdChild bool) {
	for vrfName, vrfEntry := range vrf.Vrfs {
		for serialNum, attachment := range vrfEntry.AttachList {
			if isMsdParent {
				// MSD parent: fabric field is REQUIRED
				if attachment.Fabric == "" {
					errMsg := fmt.Sprintf("VRF '%s' attachment '%s': MSD parent fabric requires 'fabric' field to be set in attach_list to specify which child fabric the attachment belongs to",
						vrfName, serialNum)
					tflog.Error(ctx, errMsg)
					dg.AddError("Missing fabric field in attachment", errMsg)
				}
			} else {
				// MSD child or Regular fabric: fabric field should NOT be set
				if attachment.Fabric != "" {
					var fabricType string
					if isMsdChild {
						fabricType = "MSD child"
					} else {
						fabricType = "Regular (non-MSD)"
					}
					errMsg := fmt.Sprintf("VRF '%s' attachment '%s': %s fabric should not have 'fabric' field set in attach_list. The fabric is computed from the resource's fabric_name. Remove the 'fabric' field from the attachment configuration",
						vrfName, serialNum, fabricType)
					tflog.Error(ctx, errMsg)
					dg.AddError("Invalid fabric field in attachment", errMsg)
				}
			}
		}
	}
}

// validateVrfsMsdParentPropertiesInChild validates that parent properties in MSD child config match the parent fabric
// For MSD child fabrics, parent-managed properties must match what's configured in the parent fabric
func (c NDFC) validateVrfsMsdParentPropertiesInChild(ctx context.Context, dg *diag.Diagnostics, childVrf *resource_vrf_bulk.NDFCVrfBulkModel, parentVrfs *resource_vrf_bulk.NDFCVrfBulkModel) {
	for vrfName, childEntry := range childVrf.Vrfs {
		parentEntry, found := parentVrfs.Vrfs[vrfName]
		if !found {
			continue // Already validated in caller that VRF exists
		}

		var mismatchedProps []string

		// Compare parent-managed properties (not child-specific properties)
		childCfg := childEntry.VrfTemplateConfig
		parentCfg := parentEntry.VrfTemplateConfig

		// VRF ID (critical parent property)
		if childEntry.VrfId != nil && *childEntry.VrfId != 0 && parentEntry.VrfId != nil && *childEntry.VrfId != *parentEntry.VrfId {
			mismatchedProps = append(mismatchedProps, fmt.Sprintf("vrf_id (child: %d, parent: %d)", *childEntry.VrfId, *parentEntry.VrfId))
		}

		// VRF Template
		if childEntry.VrfTemplate != "" && childEntry.VrfTemplate != parentEntry.VrfTemplate {
			mismatchedProps = append(mismatchedProps, fmt.Sprintf("vrf_template (child: %s, parent: %s)", childEntry.VrfTemplate, parentEntry.VrfTemplate))
		}

		// VRF Extension Template
		if childEntry.VrfExtensionTemplate != "" && childEntry.VrfExtensionTemplate != parentEntry.VrfExtensionTemplate {
			mismatchedProps = append(mismatchedProps, fmt.Sprintf("vrf_extension_template (child: %s, parent: %s)", childEntry.VrfExtensionTemplate, parentEntry.VrfExtensionTemplate))
		}

		// VRF Vlan Name
		if childCfg.VlanName != "" && childCfg.VlanName != parentCfg.VlanName {
			mismatchedProps = append(mismatchedProps, fmt.Sprintf("vrf_vlan_name (child: %s, parent: %s)", childCfg.VlanName, parentCfg.VlanName))
		}

		// VRF Interface Description
		if childCfg.InterfaceDescription != "" && childCfg.InterfaceDescription != parentCfg.InterfaceDescription {
			mismatchedProps = append(mismatchedProps, fmt.Sprintf("vrf_intf_description (child: %s, parent: %s)", childCfg.InterfaceDescription, parentCfg.InterfaceDescription))
		}

		// VRF Description
		if childCfg.VrfDescription != "" && childCfg.VrfDescription != parentCfg.VrfDescription {
			mismatchedProps = append(mismatchedProps, fmt.Sprintf("vrf_description (child: %s, parent: %s)", childCfg.VrfDescription, parentCfg.VrfDescription))
		}

		// MTU
		if childCfg.Mtu != nil && *childCfg.Mtu != 0 && parentCfg.Mtu != nil && *childCfg.Mtu != *parentCfg.Mtu {
			mismatchedProps = append(mismatchedProps, fmt.Sprintf("mtu (child: %d, parent: %d)", *childCfg.Mtu, *parentCfg.Mtu))
		}

		// Loopback Routing Tag
		if childCfg.LoopbackRoutingTag != nil && *childCfg.LoopbackRoutingTag != 0 && parentCfg.LoopbackRoutingTag != nil && *childCfg.LoopbackRoutingTag != *parentCfg.LoopbackRoutingTag {
			mismatchedProps = append(mismatchedProps, fmt.Sprintf("tag (child: %d, parent: %d)", *childCfg.LoopbackRoutingTag, *parentCfg.LoopbackRoutingTag))
		}

		// Route Target Import/Export
		if childCfg.RouteTargetImport != "" && childCfg.RouteTargetImport != parentCfg.RouteTargetImport {
			mismatchedProps = append(mismatchedProps, fmt.Sprintf("route_target_import (child: %s, parent: %s)", childCfg.RouteTargetImport, parentCfg.RouteTargetImport))
		}
		if childCfg.RouteTargetExport != "" && childCfg.RouteTargetExport != parentCfg.RouteTargetExport {
			mismatchedProps = append(mismatchedProps, fmt.Sprintf("route_target_export (child: %s, parent: %s)", childCfg.RouteTargetExport, parentCfg.RouteTargetExport))
		}
		if childCfg.RouteTargetImportEvpn != "" && childCfg.RouteTargetImportEvpn != parentCfg.RouteTargetImportEvpn {
			mismatchedProps = append(mismatchedProps, fmt.Sprintf("route_target_import_evpn (child: %s, parent: %s)", childCfg.RouteTargetImportEvpn, parentCfg.RouteTargetImportEvpn))
		}
		if childCfg.RouteTargetExportEvpn != "" && childCfg.RouteTargetExportEvpn != parentCfg.RouteTargetExportEvpn {
			mismatchedProps = append(mismatchedProps, fmt.Sprintf("route_target_export_evpn (child: %s, parent: %s)", childCfg.RouteTargetExportEvpn, parentCfg.RouteTargetExportEvpn))
		}

		// RP Loopback ID
		if childCfg.RpLoopbackId != nil && !childCfg.RpLoopbackId.IsEmpty() && parentCfg.RpLoopbackId != nil && !parentCfg.RpLoopbackId.IsEmpty() {
			if *childCfg.RpLoopbackId != *parentCfg.RpLoopbackId {
				mismatchedProps = append(mismatchedProps, fmt.Sprintf("rp_loopback_id (child: %d, parent: %d)", *childCfg.RpLoopbackId, *parentCfg.RpLoopbackId))
			}
		}

		// VLAN ID
		if childCfg.VlanId != nil && !childCfg.VlanId.IsEmpty() && parentCfg.VlanId != nil && !parentCfg.VlanId.IsEmpty() {
			if *childCfg.VlanId != *parentCfg.VlanId {
				mismatchedProps = append(mismatchedProps, fmt.Sprintf("vlan_id (child: %d, parent: %d)", *childCfg.VlanId, *parentCfg.VlanId))
			}
		}

		// Max BGP/IBGP Paths
		if childCfg.MaxBgpPaths != nil && *childCfg.MaxBgpPaths != 0 && parentCfg.MaxBgpPaths != nil && *childCfg.MaxBgpPaths != *parentCfg.MaxBgpPaths {
			mismatchedProps = append(mismatchedProps, fmt.Sprintf("max_bgp_paths (child: %d, parent: %d)", *childCfg.MaxBgpPaths, *parentCfg.MaxBgpPaths))
		}
		if childCfg.MaxIbgpPaths != nil && *childCfg.MaxIbgpPaths != 0 && parentCfg.MaxIbgpPaths != nil && *childCfg.MaxIbgpPaths != *parentCfg.MaxIbgpPaths {
			mismatchedProps = append(mismatchedProps, fmt.Sprintf("max_ibgp_paths (child: %d, parent: %d)", *childCfg.MaxIbgpPaths, *parentCfg.MaxIbgpPaths))
		}

		// IPv6 Link Local Flag
		if childCfg.Ipv6LinkLocal != "" && childCfg.Ipv6LinkLocal != parentCfg.Ipv6LinkLocal {
			mismatchedProps = append(mismatchedProps, fmt.Sprintf("ipv6_link_local_flag (child: %s, parent: %s)", childCfg.Ipv6LinkLocal, parentCfg.Ipv6LinkLocal))
		}

		// Disable RT Auto
		if childCfg.DisableRtAuto != "" && childCfg.DisableRtAuto != parentCfg.DisableRtAuto {
			mismatchedProps = append(mismatchedProps, fmt.Sprintf("disable_rt_auto (child: %s, parent: %s)", childCfg.DisableRtAuto, parentCfg.DisableRtAuto))
		}

		// If there are mismatches, report error
		if len(mismatchedProps) > 0 {
			errMsg := fmt.Sprintf("VRF '%s': MSD child fabric configuration has parent properties that don't match the parent fabric. "+
				"Parent-managed properties must be identical to the parent fabric configuration. "+
				"Mismatched properties: %s",
				vrfName, strings.Join(mismatchedProps, ", "))
			tflog.Error(ctx, errMsg)
			dg.AddError("Mismatch of config between parent and child attributes", errMsg)
		}
	}
}

// ValidateMsdParentVrfParameters validates that child-specific properties are not set in MSD parent
func (c NDFC) ValidateMsdParentVrfParameters(ctx context.Context, dg *diag.Diagnostics, vrf *resource_vrf_bulk.NDFCVrfBulkModel) {
	// Validate VRF parameters for MSD parent fabric
	// This function validates that child-specific properties are not set in MSD parent
	// Attachment fabric field validation is handled by validateAttachmentFabricField
	isMsdParent, _ := c.CheckIsFabricMsdType(ctx, dg, vrf.FabricName)
	if !isMsdParent {
		return // Validation only for MSD parent fabrics
	}

	for _, entry := range vrf.Vrfs {
		switch {
		case len(entry.VrfTemplateConfig.AdvertiseHostRoutes) != 0:
			tflog.Error(ctx, "VRF operation failed, AdvertiseHostRoutes is managed per child fabric and should not be set in MSD parent")
			dg.AddError("VRF operation failed", "AdvertiseHostRoutes is managed per child fabric and should not be set in MSD parent")
		case len(entry.VrfTemplateConfig.AdvertiseDefaultRoute) != 0:
			tflog.Error(ctx, "VRF operation failed, AdvertiseDefaultRoute is managed per child fabric and should not be set in MSD parent")
			dg.AddError("VRF operation failed", "AdvertiseDefaultRoute is managed per child fabric and should not be set in MSD parent")
		case len(entry.VrfTemplateConfig.ConfigureStaticDefaultRoute) != 0:
			tflog.Error(ctx, "VRF operation failed, ConfigureStaticDefaultRoute is managed per child fabric and should not be set in MSD parent")
			dg.AddError("VRF operation failed", "ConfigureStaticDefaultRoute is managed per child fabric and should not be set in MSD parent")
		case len(entry.VrfTemplateConfig.BgpPassword) != 0:
			tflog.Error(ctx, "VRF operation failed, BgpPassword is managed per child fabric and should not be set in MSD parent")
			dg.AddError("VRF operation failed", "BgpPassword is managed per child fabric and should not be set in MSD parent")
		case len(entry.VrfTemplateConfig.BgpPasswordType) != 0:
			tflog.Error(ctx, "VRF operation failed, BgpPasswordType is managed per child fabric and should not be set in MSD parent")
			dg.AddError("VRF operation failed", "BgpPasswordType is managed per child fabric and should not be set in MSD parent")
		case len(entry.VrfTemplateConfig.Netflow) != 0:
			tflog.Error(ctx, "VRF operation failed, Netflow is managed per child fabric and should not be set in MSD parent")
			dg.AddError("VRF operation failed", "Netflow is managed per child fabric and should not be set in MSD parent")
		case len(entry.VrfTemplateConfig.NetflowMonitor) != 0:
			tflog.Error(ctx, "VRF operation failed, NetflowMonitor is managed per child fabric and should not be set in MSD parent")
			dg.AddError("VRF operation failed", "NetflowMonitor is managed per child fabric and should not be set in MSD parent")
		case len(entry.VrfTemplateConfig.Trm) != 0:
			tflog.Error(ctx, "VRF operation failed, TRM is managed per child fabric and should not be set in MSD parent")
			dg.AddError("VRF operation failed", "TRM is managed per child fabric and should not be set in MSD parent")
		case len(entry.VrfTemplateConfig.TrmBgwMsite) != 0:
			tflog.Error(ctx, "VRF operation failed, TRM BGP Multi-Site is managed per child fabric and should not be set in MSD parent")
			dg.AddError("VRF operation failed", "TRM BGP Multi-Site is managed per child fabric and should not be set in MSD parent")
		case len(entry.VrfTemplateConfig.NoRp) != 0:
			tflog.Error(ctx, "VRF operation failed, NoRp is managed per child fabric and should not be set in MSD parent")
			dg.AddError("VRF operation failed", "NoRp is managed per child fabric and should not be set in MSD parent")
		case len(entry.VrfTemplateConfig.RpAddress) != 0:
			tflog.Error(ctx, "VRF operation failed, RpAddress is managed per child fabric and should not be set in MSD parent")
			dg.AddError("VRF operation failed", "RpAddress is managed per child fabric and should not be set in MSD parent")
		case entry.VrfTemplateConfig.RpLoopbackId != nil:
			tflog.Error(ctx, "VRF operation failed, RpLoopbackId is managed per child fabric and should not be set in MSD parent")
			dg.AddError("VRF operation failed", "RpLoopbackId is managed per child fabric and should not be set in MSD parent")
		case len(entry.VrfTemplateConfig.UnderlayMulticastAddress) != 0:
			tflog.Error(ctx, "VRF operation failed, UnderlayMulticastAddress is managed per child fabric and should not be set in MSD parent")
			dg.AddError("VRF operation failed", "UnderlayMulticastAddress is managed per child fabric and should not be set in MSD parent")
		case len(entry.VrfTemplateConfig.OverlayMulticastGroups) != 0:
			tflog.Error(ctx, "VRF operation failed, OverlayMulticastGroups is managed per child fabric and should not be set in MSD parent")
			dg.AddError("VRF operation failed", "OverlayMulticastGroups is managed per child fabric and should not be set in MSD parent")
		case len(entry.VrfTemplateConfig.RouteTargetImportMvpn) != 0:
			tflog.Error(ctx, "VRF operation failed, RouteTargetImportMvpn is managed per child fabric and should not be set in MSD parent")
			dg.AddError("VRF operation failed", "RouteTargetImportMvpn is managed per child fabric and should not be set in MSD parent")
		case len(entry.VrfTemplateConfig.RouteTargetExportMvpn) != 0:
			tflog.Error(ctx, "VRF operation failed, RouteTargetExportMvpn is managed per child fabric and should not be set in MSD parent")
			dg.AddError("VRF operation failed", "RouteTargetExportMvpn is managed per child fabric and should not be set in MSD parent")
		}
	}
}

// groupAttachmentsByChildFabric organizes VRF attachments by their target child fabric.
// Each attachment specifies a fabric field indicating which child fabric it belongs to.
// Returns a map of childFabricName -> list of VRF attachments for that fabric.
func groupAttachmentsByChildFabric(ctx context.Context, dg *diag.Diagnostics, vrfAttachments []rva.NDFCVrfAttachmentsPayload) (map[string][]rva.NDFCVrfAttachmentsPayload, error) {
	attachmentsByFabric := make(map[string][]rva.NDFCVrfAttachmentsPayload)

	for _, vrfAttachment := range vrfAttachments {
		for _, attach := range vrfAttachment.AttachList {
			childFabric := attach.Fabric
			if childFabric == "" {
				errMsg := fmt.Sprintf("Attachment for VRF %s on device %s is missing the required fabric field",
					vrfAttachment.VrfName, attach.SerialNumber)
				tflog.Error(ctx, errMsg)
				dg.AddError("Missing fabric field in attachment", errMsg)
				return nil, fmt.Errorf("%s", errMsg)
			}

			// Find or create VRF attachment entry for this child fabric
			found := false
			for i, existingVrfAttach := range attachmentsByFabric[childFabric] {
				if existingVrfAttach.VrfName == vrfAttachment.VrfName {
					// Add this attachment to existing VRF
					attachmentsByFabric[childFabric][i].AttachList = append(
						attachmentsByFabric[childFabric][i].AttachList, attach)
					found = true
					break
				}
			}

			if !found {
				// Create new VRF attachment entry for this fabric
				newVrfAttach := rva.NDFCVrfAttachmentsPayload{
					VrfName:    vrfAttachment.VrfName,
					AttachList: []rva.NDFCAttachListValue{attach},
				}
				attachmentsByFabric[childFabric] = append(attachmentsByFabric[childFabric], newVrfAttach)
			}
		}
	}

	return attachmentsByFabric, nil
}

// RscManageVrfAttachmentsForMsdParent handles VRF attachment create/delete for MSD parent fabrics.
// For MSD parent fabrics, attachments span multiple child fabrics. This function:
// 1. Groups attachments by child fabric (based on the fabric field in each attachment)
// 2. Posts bulk attachments to each child fabric in a single API call per fabric
// 3. For delete operations, immediately deploys the detachment on each child fabric
//
// Note: The NDFC API is fabric-specific (/fabrics/{fabricName}/vrfs/attachments), so we must
// post to each child fabric separately. However, for each fabric we post ALL VRF attachments
// in a single bulk operation, not one-by-one.
//
// Parameters:
//   - vrfModel: The VRF bulk model containing attachment data
//   - parentFabricName: The MSD parent fabric name (for logging/context)
//   - isDelete: true for delete operation, false for create
func (c NDFC) RscManageVrfAttachmentsForMsdParent(ctx context.Context, dg *diag.Diagnostics, vrfModel *resource_vrf_bulk.NDFCVrfBulkModel, parentFabricName string, isDelete bool) error {
	operation := "Create"
	if isDelete {
		operation = "Delete"
	}
	tflog.Debug(ctx, fmt.Sprintf("RscManageVrfAttachmentsForMsdParent: Entering %s operation", operation))
	tflog.Info(ctx, fmt.Sprintf("MSD Parent fabric '%s': Processing bulk attachments grouped by child fabric for %s", parentFabricName, operation))

	// Remap serial numbers before building attachment payload
	_ = c.vrfAttachmentSerialRemap(ctx, vrfModel)

	// Get the attachment payload (isDelete flag sets deployment to false for detach)
	// Returns a flat list of VRF attachments with fabric field set per attachment
	attachmentPayload := vrfModel.FillAttachPayloadFromModel(isDelete)
	if len(attachmentPayload.VrfAttachments) == 0 {
		tflog.Info(ctx, fmt.Sprintf("RscManageVrfAttachmentsForMsdParent: No attachments to %s", operation))
		return nil
	}

	// Group attachments by their target child fabric for bulk posting
	// Each child fabric will receive one bulk API call with all its VRF attachments
	attachmentsByFabric, err := groupAttachmentsByChildFabric(ctx, dg, attachmentPayload.VrfAttachments)
	if err != nil {
		return err
	}

	// Post bulk attachments to each child fabric
	// Each iteration makes ONE API call per fabric with ALL VRF attachments for that fabric
	var failedFabrics []string

	for childFabric, vrfAttachments := range attachmentsByFabric {
		tflog.Info(ctx, fmt.Sprintf("MSD Parent: Posting bulk payload with %d VRF attachment(s) to child fabric '%s'",
			len(vrfAttachments), childFabric))

		// Marshal bulk attachments for this child fabric into a single payload
		payloadData, err := json.Marshal(vrfAttachments)
		if err != nil {
			errMsg := fmt.Sprintf("Failed to marshal VRF attachments for child fabric '%s': %s",
				childFabric, err.Error())
			tflog.Error(ctx, errMsg)
			dg.AddError("Marshalling failed", errMsg)

			if isDelete {
				// For delete operations, try to continue with other fabrics
				failedFabrics = append(failedFabrics, childFabric)
				continue
			}
			return fmt.Errorf("%s", errMsg)
		}

		// Post bulk payload to the child fabric's endpoint in a single API call
		err = c.vrfAttachmentsPost(ctx, childFabric, payloadData)
		if err != nil {
			errMsg := fmt.Sprintf("Failed to %s bulk attachments on child fabric '%s': %s",
				operation, childFabric, err.Error())
			tflog.Error(ctx, errMsg)
			dg.AddError(fmt.Sprintf("Bulk attachment %s failed", operation), errMsg)

			if isDelete {
				// For delete operations, try to continue with other fabrics
				failedFabrics = append(failedFabrics, childFabric)
				continue
			}
			return fmt.Errorf("%s", errMsg)
		}

		tflog.Info(ctx, fmt.Sprintf("MSD Parent: Successfully %sd bulk attachments on child fabric '%s'",
			operation, childFabric))

		// For delete operations, immediately deploy the detachment on this child fabric
		if isDelete {
			tflog.Info(ctx, fmt.Sprintf("MSD Parent: Deploying detachment on child fabric '%s'", childFabric))

			// Create deployment payload for this child fabric
			deploymentPayload := &rva.NDFCVrfAttachmentsPayloads{
				FabricName:     childFabric,
				VrfAttachments: vrfAttachments,
				DepMap:         map[string][]string{"global": {"all"}},
			}

			c.RscDeployVrfAttachments(ctx, dg, deploymentPayload)
			if dg.HasError() {
				tflog.Error(ctx, fmt.Sprintf("MSD Parent: Failed to deploy detachment on child fabric '%s'", childFabric))
				failedFabrics = append(failedFabrics, childFabric)
			}
		}
	}

	// Report any failures
	if len(failedFabrics) > 0 {
		errMsg := fmt.Sprintf("Attachment operations failed on %d fabric(s): %v",
			len(failedFabrics), failedFabrics)
		tflog.Error(ctx, errMsg)
		return fmt.Errorf("%s", errMsg)
	}

	tflog.Info(ctx, fmt.Sprintf("MSD Parent: Successfully completed %s operation on all child fabrics", operation))
	return nil
}
