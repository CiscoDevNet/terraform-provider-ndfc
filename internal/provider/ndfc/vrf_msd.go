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
	"terraform-provider-ndfc/internal/provider/resources/resource_msd_vrfs_parent"
	"terraform-provider-ndfc/internal/provider/resources/resource_vrf_bulk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// MSD VRFs Parent Resource Methods (without attachments)

func (c NDFC) RscGetMsdVrfsParent(ctx context.Context, dg *diag.Diagnostics, ID string) *resource_msd_vrfs_parent.MsdVrfsParentModel {
	var filterMap map[string]bool
	tflog.Debug(ctx, fmt.Sprintf("RscGetMsdVrfsParent entry ID %s", ID))

	filterMap = make(map[string]bool)
	fabricName, vrfs := c.CreateFilterMap(ID, &filterMap)
	log.Printf("FilterMap: %v vrfs %v", filterMap, vrfs)
	if fabricName == "" {
		dg.AddError("ID format error", "ID is incorrect")
		return nil
	}
	vrfObj := api.NewVrfAPI(fabricName, c.GetLock(ResourceVrfBulk), &c.apiClient)
	res, err := vrfObj.Get()
	if err != nil {
		dg.AddError("VRF Get Failed", err.Error())
		return nil
	}
	tflog.Debug(ctx, fmt.Sprintf("RscGetMsdVrfsParent: result %s", string(res)))
	ndVrfs := resource_msd_vrfs_parent.NDFCMsdVrfsParentModel{}
	vrfPayload := resource_msd_vrfs_parent.NDFCMsdVrfPayload{}
	err = json.Unmarshal(res, &vrfPayload.Vrfs)
	if err != nil {
		dg.AddError("msd_vrfs_parent: unmarshal failed ", err.Error())
		return nil
	} else {
		tflog.Debug(ctx, "msd_vrfs_parent: Unmarshal OK")
	}
	ndVrfs.FillVrfsFromPayload(&vrfPayload)
	ndVrfs.FabricName = fabricName

	if len(filterMap) > 0 {
		log.Printf("Filtering is configured")
		filteredVrfs := make(map[string]resource_msd_vrfs_parent.NDFCVrfsValue)
		for k, v := range ndVrfs.Vrfs {
			if _, found := filterMap[k]; found {
				filteredVrfs[k] = v
			} else {
				log.Printf("Filtering out VRF %s", k)
			}
		}
		ndVrfs.Vrfs = filteredVrfs
	}

	data := new(resource_msd_vrfs_parent.MsdVrfsParentModel)
	*dg = data.SetModelData(&ndVrfs)
	if dg.HasError() {
		tflog.Error(ctx, "SetModelData failed for msd_vrfs_parent")
		return nil
	}

	if ID != "" {
		data.Id = types.StringValue(ID)
	}

	return data
}

func (c NDFC) RscImportMsdVrfsParent(ctx context.Context, dg *diag.Diagnostics, ID string) *resource_msd_vrfs_parent.MsdVrfsParentModel {
	tflog.Info(ctx, fmt.Sprintf("RscImportMsdVrfsParent entry ID %s", ID))
	return c.RscGetMsdVrfsParent(ctx, dg, ID)
}

func (c NDFC) RscCreateMsdVrfsParent(ctx context.Context, dg *diag.Diagnostics, vrfParent *resource_msd_vrfs_parent.MsdVrfsParentModel) *resource_msd_vrfs_parent.MsdVrfsParentModel {
	tflog.Debug(ctx, "RscCreateMsdVrfsParent entry")
	vrf := vrfParent.GetModelData()
	if vrf == nil {
		tflog.Error(ctx, "Data conversion from model failed")
		dg.AddError("Data conversion from model failed", "GetModelData returned empty")
		return nil
	}

	fabricName := vrf.FabricName

	for k, v := range vrf.Vrfs {
		v.FabricName = fabricName
		vrf.Vrfs[k] = v
	}

	// Form ID
	ID := c.MsdVrfsParentCreateID(vrf)

	tflog.Debug(ctx, fmt.Sprintf("RscCreateMsdVrfsParent ID %s", ID))
	// Check if VRFs already exist
	retVrfs, err := c.VrfBulkIsPresent(ctx, ID)
	if err != nil {
		tflog.Error(ctx, "Error while getting VRFs ", map[string]interface{}{"Err": err})
		dg.AddError("VRF Read Failed", err.Error())
		return nil
	}

	var errs []string
	for i := range retVrfs {
		errs = append(errs, fmt.Sprintf("VRF %s is already configured on %s", retVrfs[i], fabricName))
	}
	if len(errs) > 0 {
		tflog.Error(ctx, "VRFs exist", map[string]interface{}{"Err": errs})
		dg.AddError("VRFs exist", strings.Join(errs, ","))
		return nil
	}

	// Create VRFs (without attachments)
	ndfcVrfPayload := vrf.FillVrfPayloadFromModel()
	data, err := json.Marshal(ndfcVrfPayload.Vrfs)
	if err != nil {
		tflog.Error(ctx, "Json Marshal failure")
		dg.AddError("Cannot marshal VRF payload", err.Error())
		return nil
	}

	vrfObj := api.NewVrfAPI(fabricName, c.GetLock(ResourceVrfBulk), &c.apiClient)
	_, err = vrfObj.Post(data)
	if err != nil {
		tflog.Error(ctx, "Cannot create VRF", map[string]interface{}{"Err": err})
		dg.AddError("Cannot create VRF ", err.Error())
		return nil
	}
	tflog.Info(ctx, fmt.Sprintf("Create MSD VRFs Parent success ID %s", ID))

	// Read back to verify
	outVrf := c.RscGetMsdVrfsParent(ctx, dg, ID)
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

func (c NDFC) RscDeleteMsdVrfsParent(ctx context.Context, dg *diag.Diagnostics, ID string, vrfParent *resource_msd_vrfs_parent.MsdVrfsParentModel) {
	tflog.Info(ctx, fmt.Sprintf("MSD VRFs Parent Delete request %s", ID))
	vrfs, err := c.VrfBulkIsPresent(ctx, ID)
	if err != nil {
		tflog.Error(ctx, "Failed to Read existing VRFs")
		dg.AddError("VRF Read Failed", err.Error())
		return
	}

	results := c.RscBulkSplitID(ID)
	fabricName := results["fabric"][0]
	vrfsFromId := results["rsc"]

	if len(vrfs) != len(vrfsFromId) {
		tflog.Warn(ctx, "Some VRFs are missing in NDFC")
	}

	// Delete VRFs
	vrfObj := api.NewVrfAPI(fabricName, c.GetLock(ResourceVrfBulk), &c.apiClient)
	vrfObj.SetDeleteList(vrfs)
	_, err = vrfObj.Delete()
	if err != nil {
		tflog.Error(ctx, "VRF delete failed")
		dg.AddError("VRF Delete Failed", err.Error())
		return
	}
	tflog.Info(ctx, "MSD VRFs Parent Delete complete")
}

func (c NDFC) RscUpdateMsdVrfsParent(ctx context.Context, dg *diag.Diagnostics, ID string,
	planData, stateData, configData *resource_msd_vrfs_parent.MsdVrfsParentModel) {

	tflog.Info(ctx, fmt.Sprintf("RscUpdateMsdVrfsParent entry ID %s", ID))

	plan := planData.GetModelData()
	state := stateData.GetModelData()
	config := configData.GetModelData()

	if plan == nil || state == nil || config == nil {
		tflog.Error(ctx, "Data conversion from model failed")
		dg.AddError("Data conversion from model failed", "GetModelData returned empty")
		return
	}

	// Extract fabric name from ID (format: fabricName/[vrf1,vrf2,vrf3])
	// For MSD parent resource, this is the parent fabric name where VRFs are defined
	results := c.RscBulkSplitID(ID)
	fabricName := results["fabric"][0]
	if fabricName == "" {
		dg.AddError("ID format error", "Cannot extract fabric name from ID")
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Using fabric name %s for VRF operations", fabricName))

	// Detect VRFs to add, update, or delete
	vrfsToAdd := make(map[string]resource_msd_vrfs_parent.NDFCVrfsValue)
	vrfsToUpdate := make(map[string]resource_msd_vrfs_parent.NDFCVrfsValue)
	vrfsToDelete := []string{}

	// Find VRFs to add or update
	for vrfName, planVrf := range plan.Vrfs {
		// Ensure FabricName is set to the MSD parent fabric name for all VRF operations
		planVrf.FabricName = fabricName
		if _, exists := state.Vrfs[vrfName]; exists {
			vrfsToUpdate[vrfName] = planVrf
		} else {
			vrfsToAdd[vrfName] = planVrf
		}
	}

	// Find VRFs to delete
	for vrfName := range state.Vrfs {
		if _, exists := plan.Vrfs[vrfName]; !exists {
			vrfsToDelete = append(vrfsToDelete, vrfName)
		}
	}

	// Note: VRFs cannot be deleted in ndfc_msd_vrfs_parent as this resource doesn't manage attachments.
	// VRFs can only be deleted via ndfc_vrfs resource which handles attachments and detachment.
	// Log warning if VRFs are being removed from configuration.
	if len(vrfsToDelete) > 0 {
		tflog.Warn(ctx, fmt.Sprintf("VRFs cannot be deleted via ndfc_msd_vrfs_parent resource: %v. Use ndfc_vrfs resource to detach and delete VRFs with attachments.", vrfsToDelete))
	}

	vrfObj := api.NewVrfAPI(fabricName, c.GetLock(ResourceVrfBulk), &c.apiClient)

	// Add new VRFs
	if len(vrfsToAdd) > 0 {
		addPayload := &resource_msd_vrfs_parent.NDFCMsdVrfsParentModel{
			FabricName: fabricName, // Set fabric name on the parent model
			Vrfs:       vrfsToAdd,
		}
		ndfcPayload := addPayload.FillVrfPayloadFromModel()
		data, err := json.Marshal(ndfcPayload.Vrfs)
		if err != nil {
			dg.AddError("Cannot marshal VRF payload", err.Error())
			return
		}
		tflog.Debug(ctx, fmt.Sprintf("Adding VRFs with payload: %s", string(data)))
		_, err = vrfObj.Post(data)
		if err != nil {
			tflog.Error(ctx, "Cannot add VRFs", map[string]interface{}{"Err": err})
			dg.AddError("Cannot add VRFs ", err.Error())
			return
		}
	}

	// Update existing VRFs
	if len(vrfsToUpdate) > 0 {
		// First, get current VRFs from NDFC to preserve immutable fields (like VRF ID)
		currentVrfs, err := c.vrfBulkGet(ctx, fabricName)
		if err != nil {
			tflog.Error(ctx, "Failed to get current VRFs for update", map[string]interface{}{"Err": err})
			dg.AddError("VRF Read Failed", err.Error())
			return
		}

		for vrfName, vrfValue := range vrfsToUpdate {
			vrfObj.PutVrf = vrfName

			// Preserve immutable fields from current state (like vrf_id)
			if currentVrf, exists := currentVrfs.Vrfs[vrfName]; exists {
				vrfValue.VrfId = currentVrf.VrfId
			}

			updatePayload := &resource_msd_vrfs_parent.NDFCMsdVrfsParentModel{
				FabricName: fabricName, // Set fabric name on the parent model
				Vrfs: map[string]resource_msd_vrfs_parent.NDFCVrfsValue{
					vrfName: vrfValue,
				},
			}
			ndfcPayload := updatePayload.FillVrfPayloadFromModel()
			if len(ndfcPayload.Vrfs) > 0 {
				data, err := json.Marshal(ndfcPayload.Vrfs[0])
				if err != nil {
					dg.AddError("Cannot marshal VRF payload", err.Error())
					return
				}
				tflog.Debug(ctx, fmt.Sprintf("Updating VRF %s with payload: %s", vrfName, string(data)))
				_, err = vrfObj.Put(data)
				if err != nil {
					tflog.Error(ctx, fmt.Sprintf("Cannot update VRF %s", vrfName), map[string]interface{}{"Err": err})
					dg.AddError("Cannot update VRFs", err.Error())
					return
				}
			}
		}
	}

	tflog.Info(ctx, "MSD VRFs Parent Update complete")

	// Read back the state after update to get all current values from NDFC
	newID := c.MsdVrfsParentCreateID(plan)
	updatedData := c.RscGetMsdVrfsParent(ctx, dg, newID)
	if updatedData != nil {
		*planData = *updatedData
	}

}

// ID = fabricName/[vrf1,vrf2,vrf3]

func (c NDFC) MsdVrfsParentCreateID(ndVrfs *resource_msd_vrfs_parent.NDFCMsdVrfsParentModel) string {
	if ndVrfs != nil {
		fName := ndVrfs.FabricName
		vrf_list := ndVrfs.GetVrfNames()
		return fmt.Sprintf("%s/[%s]", fName, strings.Join(vrf_list, ","))
	}
	return ""
}

// restoreParentPropertiesInState restores parent-managed properties from config into the VRF model after reading from NDFC
// This preserves parent properties from Terraform config while keeping child-specific properties from NDFC
func (c NDFC) restoreParentPropertiesInState(ctx context.Context, vrfModel *resource_vrf_bulk.VrfBulkModel,
	parentProps map[string]resource_vrf_bulk.ParentManagedVrfProperties) {
	if parentProps != nil {
		model := vrfModel.GetModelData()
		if model != nil {
			model.RestoreParentProperties(parentProps)
			vrfModel.SetModelData(model)
			tflog.Debug(ctx, "Restored parent-managed properties from config into state for MSD child fabric")
		}
	}
}
