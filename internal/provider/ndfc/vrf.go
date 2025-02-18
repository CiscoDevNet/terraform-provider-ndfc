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
	"regexp"
	"strings"
	"terraform-provider-ndfc/internal/provider/datasources/datasource_vrf_bulk"
	"terraform-provider-ndfc/internal/provider/ndfc/api"
	"terraform-provider-ndfc/internal/provider/resources/resource_vrf_bulk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const ResourceVrfBulk = "vrfs"

func (c NDFC) DSGetBulkVrf(ctx context.Context, dg *diag.Diagnostics, fabricName string) *datasource_vrf_bulk.VrfBulkModel {
	log.Printf("DSGetBulkVrf entry fabirc %s", fabricName)

	vrfObj := api.NewVrfAPI(fabricName, c.GetLock(ResourceVrfBulk), &c.apiClient)
	res, err := vrfObj.Get()
	if err != nil {
		dg.AddError("VRF Get Failed", err.Error())
		return nil
	}

	ndVrfs := datasource_vrf_bulk.NDFCVrfBulkModel{}
	err = json.Unmarshal(res, &ndVrfs.Vrfs)
	if err != nil {
		dg.AddError("datasource_vrf_bulk: unmarshal failed ", err.Error())
		return nil
	} else {
		tflog.Debug(ctx, "datasource_vrf_bulk: Unmarshal OK")
	}
	if len(ndVrfs.Vrfs) > 0 {
		ndVrfs.FabricName = fabricName
	} else {
		ndVrfs.FabricName = ""
		dg.AddWarning("No VRFs found", fmt.Sprintf("No VRFs configured in fabric %s", fabricName))
		return nil
	}

	err = c.DsGetVrfAttachments(ctx, dg, &ndVrfs)
	if err != nil {
		dg.AddError("VRF Attachments read failed", err.Error())
	}

	data := new(datasource_vrf_bulk.VrfBulkModel)
	d := data.SetModelData(&ndVrfs)
	if d != nil {
		*dg = d
		return nil
	} else {
		tflog.Debug(ctx, "datasource_vrf_bulk: SetModelData OK")
	}
	return data
}

// ID = fabricName/[vrf1{attachment1,attachment2,attachment3},vrf2{attachment1,attachment2,attachment3},vrf3{attachment1,attachment2,attachment3}]

func (c NDFC) VrfBulkCreateID(ndVrfs *resource_vrf_bulk.NDFCVrfBulkModel) string {
	uniqueID := "%s/%v"
	if ndVrfs != nil {
		fName := ndVrfs.FabricName
		vrf_list := ndVrfs.GetVrfNames()
		for i := range vrf_list {
			serials := ndVrfs.GetAttachmentNames(vrf_list[i])
			if len(serials) > 0 {
				vrf_list[i] = fmt.Sprintf("%s{%s}", vrf_list[i], strings.Join(serials, ","))
			}
		}
		vrf_list_str := "[" + strings.Join(vrf_list, ",") + "]"
		return fmt.Sprintf(uniqueID, fName, vrf_list_str)
	}
	return ""
}

func (c NDFC) RscGetBulkVrf(ctx context.Context, dg *diag.Diagnostics, ID string, depMap *map[string][]string) *resource_vrf_bulk.VrfBulkModel {
	var filterMap map[string]bool
	tflog.Debug(ctx, fmt.Sprintf("RscGetBulkVrf entry fabirc %s", ID))

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
	tflog.Debug(ctx, fmt.Sprintf("RscGetBulkVrf: result %s", string(res)))
	ndVrfs := resource_vrf_bulk.NDFCVrfBulkModel{}
	vrfPayloads := resource_vrf_bulk.NDFCBulkVrfPayload{}
	err = json.Unmarshal(res, &vrfPayloads.Vrfs)
	if err != nil {
		dg.AddError("resource_vrf_bulk: unmarshal failed ", err.Error())
		return nil
	} else {
		tflog.Debug(ctx, "resource_vrf_bulk: Unmarshal OK")
	}
	ndVrfs.FillVrfsFromPayload(&vrfPayloads)

	if len(filterMap) > 0 {
		log.Printf("Filtering is configured")
		//Set value filter - skip the vrfs that are not in ID
		for i, vrfEntry := range ndVrfs.Vrfs {
			if _, found := (filterMap)[vrfEntry.VrfName]; !found {
				log.Printf("Filtering out VRF %s", vrfEntry.VrfName)
				vrfEntry.FilterThisValue = true
				ndVrfs.Vrfs[i] = vrfEntry
			}
		}
	}

	if _, ok := (*depMap)["global"]; ok {
		//This cannot be validated - as we don't know if all deployments were ok
		log.Printf("[DEBUG]: Setting DeployAllAttachments flag")
		ndVrfs.DeployAllAttachments = true
	}

	//Get Attachments
	err = c.RscGetVrfAttachments(ctx, dg, &ndVrfs)
	if err == nil {
		for i, vrfEntry := range ndVrfs.Vrfs {
			if vrfEntry.FilterThisValue {
				continue
			}
			vrfLevelDep := false

			if !ndVrfs.DeployAllAttachments {
				if vl, vlOk := (*depMap)[i]; vlOk {
					//first element is vrf name if vrf level deploy is set
					if vl[0] == i {
						vrfEntry.DeployAttachments = (vrfEntry.VrfStatus == "DEPLOYED")
						log.Printf("Setting VRF level dep flag for %s to %v", i, vrfEntry.DeployAttachments)
						vrfLevelDep = true
					}
				}
			}
			for j, attachEntry := range vrfEntry.AttachList {
				if attachEntry.FilterThisValue {
					continue
				}
				log.Printf("Attachment %s added to VRF %s", j, i)
				if !ndVrfs.DeployAllAttachments && !vrfLevelDep {
					if attachEntry.AttachState == "DEPLOYED" {
						attachEntry.DeployThisAttachment = true
					}
					log.Printf("[DEBUG] Set Attachment level deploy flag %s/%s:%v", i, j, attachEntry.DeployThisAttachment)
				}
				//put modified entry back
				vrfEntry.AttachList[j] = attachEntry
			}
			//put modified entry back
			ndVrfs.Vrfs[i] = vrfEntry

		}
	} else {
		dg.AddError("VRF Attachments read failed", err.Error())
	}
	vModel := new(resource_vrf_bulk.VrfBulkModel)
	vModel.Id = types.StringValue(ID)
	d := vModel.SetModelData(&ndVrfs)
	if d != nil {
		dg.Append(d.Errors()...)
	}
	return vModel
}

func (c NDFC) RscImportBulkVrf(ctx context.Context, dg *diag.Diagnostics, ID string) *resource_vrf_bulk.VrfBulkModel {
	var filterMap map[string]bool
	tflog.Debug(ctx, fmt.Sprintf("RscImportBulkVrf entry fabric %s", ID))
	re := regexp.MustCompile(`^[\w-]+\/\[(?:[\w-]+\,?)+\]$`)
	if !re.Match([]byte(ID)) {
		dg.AddError("ID format error", "use fabricName/[vrf1,vrf2...] format")
		return nil
	}
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
	tflog.Debug(ctx, fmt.Sprintf("RscImportBulkVrf: result %s", string(res)))
	ndVrfs := resource_vrf_bulk.NDFCVrfBulkModel{}

	vrfPayloads := resource_vrf_bulk.NDFCBulkVrfPayload{}
	err = json.Unmarshal(res, &vrfPayloads.Vrfs)
	if err != nil {
		dg.AddError("resource_vrf_bulk: unmarshal failed ", err.Error())
		return nil
	} else {
		tflog.Debug(ctx, "resource_vrf_bulk: Unmarshal OK")
	}
	ndVrfs.FillVrfsFromPayload(&vrfPayloads)
	inCount := 0
	if len(filterMap) > 0 {
		log.Printf("Filtering is configured")
		//Set value filter - skip the vrfs that are not in ID
		for i, vrfEntry := range ndVrfs.Vrfs {
			if _, found := (filterMap)[vrfEntry.VrfName]; !found {
				log.Printf("Filtering out VRF %s", vrfEntry.VrfName)
				vrfEntry.FilterThisValue = true
				ndVrfs.Vrfs[i] = vrfEntry
			} else {
				inCount++
			}
		}
	}

	if inCount == 0 {
		dg.AddError("VRFs not found", "No VRFs found")
		return nil
	}

	if len(ndVrfs.Vrfs) == 0 {
		dg.AddError("VRFs not found", "No VRFs found")
		return nil
	}
	//Get Attachments
	err = c.RscGetVrfAttachments(ctx, dg, &ndVrfs)
	if err == nil {
		for i, vrfEntry := range ndVrfs.Vrfs {
			if vrfEntry.FilterThisValue {
				continue
			}
			for j, attachEntry := range vrfEntry.AttachList {
				if attachEntry.FilterThisValue {
					continue
				}
				log.Printf("Attachment %s added to VRF %s", j, i)
				if attachEntry.AttachState == "DEPLOYED" {
					attachEntry.DeployThisAttachment = true
				}
				log.Printf("[DEBUG] Set Attachment level deploy flag %s/%s:%v", i, j, attachEntry.DeployThisAttachment)
				//put modified entry back

				vrfEntry.AttachList[j] = attachEntry
			}
			//put modified entry back
			ndVrfs.Vrfs[i] = vrfEntry
		}
	} else {
		dg.AddError("VRF Attachments read failed", err.Error())
	}

	vModel := new(resource_vrf_bulk.VrfBulkModel)
	vModel.Id = types.StringValue(ID)
	d := vModel.SetModelData(&ndVrfs)
	if d != nil {
		dg.Append(d.Errors()...)
	}
	return vModel
}

/*
Usecase: Bulk VRF Creation

Case 1: No VRFs from list exist on NDFC already

	action: create everything

Case 2: Some VRFs already exists
  - user created some manually
  - vrfs were part of another resource (bulk/single)
    Action : Don't proceed - infra tainted.
    User can manually cleanup or update the resource to exclude existing

Case 3: During Creation, some failed - probably due to a param problem or some other issues

	VRFs partially created
	Action:
	Delete what we created - rollback
	return failure - reporting the issues

case 4: Nothing was created

	return failure, report the errors coming from NDFC
*/
func (c NDFC) RscCreateBulkVrf(ctx context.Context, dg *diag.Diagnostics, vrfBulk *resource_vrf_bulk.VrfBulkModel) *resource_vrf_bulk.VrfBulkModel {
	tflog.Debug(ctx, fmt.Sprintf("RscCreateBulkVrf entry fabirc %s", vrfBulk.FabricName.ValueString()))
	vrf := vrfBulk.GetModelData()
	if vrf == nil {
		tflog.Error(ctx, "Data conversion from model failed")
		dg.AddError("Data conversion from model failed", "GetModelData returned empty")
		return nil
	}
	//form ID
	ID := c.VrfBulkCreateID(vrf)
	//create

	depMap := make(map[string][]string)

	retVrfs, err := c.VrfBulkIsPresent(ctx, ID)

	if err != nil {
		tflog.Error(ctx, "Error while getting VRFs ", map[string]interface{}{"Err": err})
		dg.AddError("VRF Read Failed", err.Error())
		return nil
	}

	var errs []string
	for i := range retVrfs {
		errs = append(errs, fmt.Sprintf("VRF %s is already configured on %s", retVrfs[i], vrf.FabricName))
	}
	if len(errs) > 0 {
		tflog.Error(ctx, "VRFs exist", map[string]interface{}{"Err": errs})
		dg.AddError("VRFs exist", strings.Join(errs, ","))
		return nil
	}

	//Part 1: Create VRFs
	ndfcVrfBulkPayload := vrf.FillVrfPayloadFromModel(&depMap)
	//ndfcVrfBulkPayload.Vrfs = vrf.GetVrfValues()

	err = c.vrfCreateBulk(ctx, vrf.FabricName, ndfcVrfBulkPayload)
	if err != nil {
		tflog.Error(ctx, "Cannot create VRF", map[string]interface{}{"Err": err})
		dg.AddError("Cannot create VRF ", err.Error())
		return nil
	}
	tflog.Info(ctx, fmt.Sprintf("Create Bulk VRF success ID %s", ID))
	//Part 2: Create Attachments if any

	if vrf.DeployAllAttachments {
		log.Printf("[DEBUG] DeployAllAttachments flag set")
		depMap["global"] = append(depMap["global"], "all")
	}
	// fill the attachment entries
	va := vrf.FillAttachPayloadFromModel(false)

	if len(va.VrfAttachments) > 0 {
		err := c.RscCreateVrfAttachments(ctx, dg, va)
		if err != nil {
			tflog.Error(ctx, "VRF Attachments create failed")
			tflog.Error(ctx, "Rolling back the configurations...delete VRFs")
			c.RscDeleteBulkVrf(ctx, dg, ID, vrfBulk)
			return nil
		}
		//Check and deploy
		c.RscDeployVrfAttachments(ctx, dg, vrf)
	}
	outVrf := c.RscGetBulkVrf(ctx, dg, ID, &depMap)
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

func (c NDFC) RscDeleteBulkVrf(ctx context.Context, dg *diag.Diagnostics, ID string, vrfBulk *resource_vrf_bulk.VrfBulkModel) {
	tflog.Info(ctx, fmt.Sprintf("VRF Bulk Delete request %s", ID))
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
		errString := fmt.Sprintf("Mismatch in VRF data: fabric %s Read %v, from ID %v", fabricName, vrfs, vrfsFromId)
		tflog.Error(ctx, errString)
		dg.AddWarning(("Mismatch in VRF data"), errString)
		vrfsFromId = vrfs
	}

	delVrf := vrfBulk.GetModelData()

	err = c.RscDeleteVrfAttachments(ctx, dg, delVrf)
	if err != nil {
		tflog.Error(ctx, "VRF Attachments delete failed")
		dg.AddError("VRF Attachments delete failed", err.Error())
		return
	}

	err = c.vrfBulkDelete(ctx, fabricName, vrfsFromId)
	if err != nil {
		errString := fmt.Sprintf("VRF delete failed on fabric %s vrfs %v", fabricName, vrfsFromId)
		tflog.Error(ctx, "VRF delete failed")
		dg.AddError(errString, err.Error())
		return
	}

	tflog.Info(ctx, fmt.Sprintf("VRF Bulk Delete Success %s", ID))

}

/*
1. Compare state and plan
2. Do bulk update for those VRFs changed
*/
func (c NDFC) RscUpdateBulkVrf(ctx context.Context,
	dg *diag.Diagnostics, ID string,
	vrfBulkPlan *resource_vrf_bulk.VrfBulkModel,
	vrfState *resource_vrf_bulk.VrfBulkModel, vrfConfig *resource_vrf_bulk.VrfBulkModel) {

	actions := c.vrfBulkGetDiff(ctx, dg, vrfBulkPlan, vrfState, vrfConfig)

	// Validate the Diff
	//Get the current VRFs from NDFC

	delVrfs := actions["del"].(*resource_vrf_bulk.NDFCVrfBulkModel)
	//deployVrfs := actions["deploy"].([]string)

	putVrfs := actions["put"].(*resource_vrf_bulk.NDFCVrfBulkModel)
	newVrfs := actions["add"].(*resource_vrf_bulk.NDFCVrfBulkModel)
	plan := actions["plan"].(*resource_vrf_bulk.NDFCVrfBulkModel)
	state := actions["state"].(*resource_vrf_bulk.NDFCVrfBulkModel)

	ndfcVRFs, err := c.vrfBulkGet(ctx, vrfBulkPlan.FabricName.ValueString())
	// Step 1 - Check if VRFs to update are present in NDFC
	if err != nil {
		tflog.Error(ctx, "Failed to Read existing VRFs")
		dg.AddError("VRF Read Failed", err.Error())
		return
	}
	for vrf := range putVrfs.Vrfs {
		_, ok := ndfcVRFs.Vrfs[vrf]
		if !ok {
			errString := fmt.Sprintf("VRF %s to update is missing in NDFC", vrf)
			tflog.Error(ctx, errString)
			dg.AddError("In place update failed", errString)
			return
		}
	}
	/*
		updateVA := resource_vrf_attachments.NDFCVrfAttachmentsModel{}
		updateVA.FabricName = vrfBulkPlan.FabricName.ValueString()
		updateVA.VrfAttachments = make(resource_vrf_attachments.NDFCVrfAttachmentsValues, 0)
	*/
	// Step 2 - VRFs to delete are available
	for vrf := range delVrfs.Vrfs {
		_, ok := ndfcVRFs.Vrfs[vrf]
		if !ok {
			// is this error a big deal? VRFs to be deleted missing - so ignore??
			delete(delVrfs.Vrfs, vrf)
			tflog.Error(ctx, fmt.Sprintf("VRF to DELETE %s is missing in NDFC", vrf))
		}
	}
	// Step 3 - Check if VRFs to be created (create-delete) are not present in NDFC

	for vrf := range newVrfs.Vrfs {
		_, ok := ndfcVRFs.Vrfs[vrf]
		if ok {
			//VRF present in NDFC
			//Check if the VRF is marked for deletion as part of Delete-Create diff

			if _, ok := delVrfs.Vrfs[vrf]; ok {
				tflog.Debug(ctx, "VRF marked for delete", map[string]interface{}{"vrfName": vrf})
				continue
			}
			errString := fmt.Sprintf("VRF %s to Create is Already Present in NDFC. Update cancelled", vrf)
			tflog.Error(ctx, errString)
			dg.AddError("In place update failed", errString)
			return
		}
	}

	//Begin update - No rollback from here on

	if len(delVrfs.Vrfs) > 0 {
		// Check and delete attachments
		err := c.RscDeleteVrfAttachments(ctx, dg, delVrfs)
		if err != nil {
			tflog.Error(ctx, "VRF Attachments delete failed")
			dg.AddError("VRF Attachments delete failed", err.Error())
			return
		}
		tflog.Info(ctx, fmt.Sprintf("Deleting VRFs %v, as part of Bulk Update", delVrfs.GetVrfNames()))
		err = c.vrfBulkDelete(ctx, vrfBulkPlan.FabricName.ValueString(), delVrfs.GetVrfNames())
		if err != nil {
			dg.AddError("Bulk Delete failed", err.Error())
		}
	} else {
		tflog.Info(ctx, "Nothing to delete")
	}
	if len(putVrfs.Vrfs) > 0 {
		tflog.Info(ctx, "Modifying VRFs , as part of Bulk Update")
		c.vrfBulkUpdate(ctx, dg, putVrfs)
		if dg.HasError() {
			tflog.Info(ctx, fmt.Sprintf("Modifying VRFs Failed %v", dg.Errors()))
			return
		}
	} else {
		tflog.Info(ctx, "Nothing to modify")
	}
	// TODO: All deployed attachments of modified VRFs must be re-deployed

	if len(newVrfs.Vrfs) > 0 {
		tflog.Info(ctx, "Adding VRFs , as part of Bulk Update")
		newVrfPayload := newVrfs.FillVrfPayloadFromModel(nil)
		err := c.vrfCreateBulk(ctx, vrfBulkPlan.FabricName.ValueString(), newVrfPayload)
		if err != nil {
			dg.AddError("VRF create failed", err.Error())
			return
		}
	}

	//Deal with attachments
	//updateVA.VrfAttachments = make(resource_vrf_attachments.NDFCVrfAttachmentsValues, 0)
	//copyVrfAttachments(plan, &updateVA)
	//stateVA := resource_vrf_attachments.NDFCVrfAttachmentsModel{}
	//copyVrfAttachments(state, &stateVA)
	c.RscUpdateVrfAttachments(ctx, dg, plan, state)
	if dg.HasError() {
		tflog.Error(ctx, "Error during update")
		return
	}
	newID := c.VrfBulkCreateID(plan)
	depMap := make(map[string][]string)
	if plan.DeployAllAttachments {
		depMap["global"] = append(depMap["global"], "all")
	}

	for i := range plan.Vrfs {
		if plan.Vrfs[i].DeployAttachments {
			depMap[plan.Vrfs[i].VrfName] = append(depMap[plan.Vrfs[i].VrfName], plan.Vrfs[i].VrfName)
		}
		for j := range plan.Vrfs[i].AttachList {
			if plan.Vrfs[i].AttachList[j].DeployThisAttachment {
				depMap[plan.Vrfs[i].VrfName] = append(depMap[plan.Vrfs[i].VrfName], plan.Vrfs[i].AttachList[j].SerialNumber)
			}
		}
	}
	*(vrfBulkPlan) = *(c.RscGetBulkVrf(ctx, dg, newID, &depMap))
}

func (c NDFC) VrfBulkIsPresent(ctx context.Context, ID string) ([]string, error) {
	var retVrfs []string
	filterMap := make(map[string]bool)

	fabricName, _ := c.CreateFilterMap(ID, &filterMap)

	vrfObj := api.NewVrfAPI(fabricName, c.GetLock(ResourceVrfBulk), &c.apiClient)
	res, err := vrfObj.Get()
	if err != nil {
		tflog.Error(ctx, "Error while getting VRFs ", map[string]interface{}{"Err": err})
		return nil, err
	}

	vrfPayload := resource_vrf_bulk.NDFCBulkVrfPayload{}
	err = json.Unmarshal(res, &vrfPayload.Vrfs)
	if err != nil {
		tflog.Error(ctx, "resource_vrf_bulk: unmarshal failed ", map[string]interface{}{"Err": err})
		return nil, err
	}

	for i := range vrfPayload.Vrfs {
		ok, found := filterMap[vrfPayload.Vrfs[i].VrfName]
		if ok && found {
			retVrfs = append(retVrfs, vrfPayload.Vrfs[i].VrfName)
		}
	}
	return retVrfs, nil

}

/*
func (c NDFC) vrfGet(fabric, vrfName string) *resource_vrf_bulk.NDFCVrfsValue {
	vrfObj := api.NewVrfAPI(fabric, c.GetLock(ResourceVrfBulk), &c.apiClient)
	res, err := vrfObj.GetSingleVRF(fabric, vrfName)
	if err != nil {
		return nil
	}
	vrf := resource_vrf_bulk.NDFCVrfsValue{}
	err = json.Unmarshal(res, &vrf)
	if err != nil {
		return nil
	}
	return &vrf
}
*/
