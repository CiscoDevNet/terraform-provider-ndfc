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
	"errors"
	"fmt"
	"log"
	"net"
	rva "terraform-provider-ndfc/internal/provider/resources/resource_vrf_attachments"
	"terraform-provider-ndfc/internal/provider/resources/resource_vrf_bulk"
	. "terraform-provider-ndfc/internal/provider/types"

	api "terraform-provider-ndfc/internal/provider/ndfc/api"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
)

const (
	KeyTypeIP     = uint16(1)
	KeyTypeSerial = uint16(2)
)

func (c NDFC) vrfCreateBulk(ctx context.Context, fabricName string, vrfsPayload *resource_vrf_bulk.NDFCBulkVrfPayload) error {
	tflog.Info(ctx, fmt.Sprintf("Beginning Bulk VRF create in fabric %s", fabricName))
	data, err := json.Marshal(vrfsPayload.Vrfs)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Json Marshal failure %s", err.Error()))
		return err
	}
	log.Println("Data to be posted", string(data))

	vrfObj := api.NewVrfAPI(fabricName, c.GetLock(ResourceVrfBulk), &c.apiClient)
	res, err := vrfObj.Post(data)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Error POST:  %s", err.Error()))
		okList, err1 := c.processBulkResponse(ctx, res)
		err2 := c.vrfBulkDelete(ctx, fabricName, okList)
		return errors.Join(err, err1, err2)
	}

	tflog.Info(ctx, fmt.Sprintf("vrfCreateBulk: Success res : %v", res.String()))
	return nil
}

func (c NDFC) vrfBulkDelete(ctx context.Context, fabricName string, vrfList []string) error {
	if len(vrfList) == 0 {
		return nil
	}
	tflog.Info(ctx, fmt.Sprintf("Attempting to delete VRFs fabric_name=%s, vrfs = %v", fabricName, vrfList))

	vrfObj := api.NewVrfAPI(fabricName, c.GetLock(ResourceVrfBulk), &c.apiClient)
	vrfObj.SetDeleteList(vrfList)
	res, err := vrfObj.Delete()
	if err != nil {

		_, err1 := c.processBulkResponse(ctx, res)
		return err1
	}
	tflog.Info(ctx, fmt.Sprintf("Deleting VRFs OK fabric_name=%s, vrfs = %v", fabricName, vrfList))
	return nil

}

func (c NDFC) vrfBulkGet(ctx context.Context, fabricName string) (*resource_vrf_bulk.NDFCVrfBulkModel, error) {

	vrfObj := api.NewVrfAPI(fabricName, c.GetLock(ResourceVrfBulk), &c.apiClient)
	res, err := vrfObj.Get()
	if err != nil {
		return nil, err
	}
	tflog.Debug(ctx, fmt.Sprintf("vrfBulkGet: result %s", string(res)))
	ndVrfs := resource_vrf_bulk.NDFCVrfBulkModel{}
	vrfPayloads := resource_vrf_bulk.NDFCBulkVrfPayload{}
	err = json.Unmarshal(res, &vrfPayloads.Vrfs)
	if err != nil {
		return nil, err
	} else {
		tflog.Debug(ctx, "vrfBulkGet: Unmarshal OK")
	}
	ndVrfs.FillVrfsFromPayload(&vrfPayloads)
	return &ndVrfs, nil
}

/*
func (c NDFC) vrfBulkCreateCheck(ctx context.Context, ID string) error {

	retVrfs, err := c.vrfBulkIsPresent(ctx, ID)
	if err != nil {
		tflog.Error(ctx, "Error while getting VRFs ", map[string]interface{}{"Err": err})
		return err
	}

	var errs []error
	for i := range retVrfs {
		errs = append(errs, fmt.Errorf("VRF %s is already configured on %s", retVrfs[i], vrfs.FabricName.ValueString()))
	}

	if len(errs) > 0 {
		tflog.Error(ctx, "VRFs exist", map[string]interface{}{"Err": errs})
		return errors.Join(errs...)
	}

	tflog.Info(ctx, "VRFs were not found in NDFC, ok to create")
	return nil
}
*/
/*
Failure Response format
{"failureList":

	     [{"name":"Murali_Bulk_01","message":"VRF\tMurali_Bulk_01\talready exists","status":"Failed"},
		  {"name":"Murali_Bulk_02","message":"VRF\tMurali_Bulk_02\talready exists","status":"Failed"}
		  ],
		  "successList":[]
		  {"successList":
		  [{"name":"Murali_Bulk_01","message":"VRF is successfully created.","status":"Success"},
		   {"name":"Murali_Bulk_02","message":"VRF is successfully created.","status":"Success"}]}
	}
*/
func (c NDFC) processBulkResponse(ctx context.Context, res gjson.Result) ([]string, error) {
	tflog.Error(ctx, res.String())
	flist := res.Get("failureList")
	var failed []map[string]string
	var errs []error
	e := json.Unmarshal([]byte(flist.Raw), &failed)
	if e != nil {
		log.Println("Error unmarshalling response from NDFC")
		errs = append(errs, e)
	}
	for _, v := range failed {
		errs = append(errs, fmt.Errorf("VRF=%s Status=%s Message=%s", v["name"], v["status"], v["message"]))
	}
	var arr []string
	if res.Get("successList.#.name").Exists() {
		slist := res.Get("successList.#.name").Value().([]interface{})
		for _, v := range slist {
			arr = append(arr, v.(string))
		}
	}
	return arr, errors.Join(errs...)
}

/*
	func (c NDFC) vrfBulkSplitID(ID string) (string, []string) {
		idSplit := strings.Split(ID, "/")
		fabricName := idSplit[0]
		var vrfs []string
		if len(idSplit) > 1 {
			vrfs_string := idSplit[1]
			log.Println(vrfs_string)
			if vrfs_string[0] == '{' && vrfs_string[len(vrfs_string)-1] == '}' {
				vrfs = strings.Split(vrfs_string[1:len(vrfs_string)-1], ",")
			}
		}
		return fabricName, vrfs
	}
*/
func (c NDFC) vrfBulkGetDiff(ctx context.Context,
	vPlan *resource_vrf_bulk.VrfBulkModel,
	vState *resource_vrf_bulk.VrfBulkModel, _ *resource_vrf_bulk.VrfBulkModel) map[string]interface{} {

	actions := make(map[string]interface{})
	vrfState := vState.GetModelData()
	vrfConfig := vPlan.GetModelData()
	//vrfConfig := vConfig.GetModelData()

	//var delVrfs []string
	var deployVrfs []string

	putVRFs := new(resource_vrf_bulk.NDFCVrfBulkModel)
	putVRFs.Vrfs = make(map[string]resource_vrf_bulk.NDFCVrfsValue)
	putVRFs.FabricName = vPlan.FabricName.ValueString()

	newVRFs := new(resource_vrf_bulk.NDFCVrfBulkModel)
	newVRFs.Vrfs = make(map[string]resource_vrf_bulk.NDFCVrfsValue)
	newVRFs.FabricName = vPlan.FabricName.ValueString()

	delVrfs := new(resource_vrf_bulk.NDFCVrfBulkModel)
	delVrfs.Vrfs = make(map[string]resource_vrf_bulk.NDFCVrfsValue)
	delVrfs.FabricName = vPlan.FabricName.ValueString()

	for sVrfName, sVrf := range vrfState.Vrfs {
		if vrf, ok := vrfConfig.Vrfs[sVrfName]; ok {

			vrf.FilterThisValue = true
			vrf.FabricName = newVRFs.FabricName
			vrf.VrfName = sVrfName
			cf := false                               //ignored here - should be taken care in attachments
			updateAction := vrf.CreatePlan(sVrf, &cf) //vrfState.Vrfs[i].DeepEqual(*vrf)
			if updateAction == ActionNone {
				//Case 1: Both VRFs are equal - no change to the VRF entry
				tflog.Info(ctx, fmt.Sprintf("%s not changed", sVrfName))

			} else if updateAction == RequiresReplace {
				//Case 2: attribute that cannot be modified in-place has changed - DELETE and Create
				tflog.Info(ctx, fmt.Sprintf("%s Needs to be replaced - Delete and Add |%s|", sVrfName, vrf.VrfName))

				// Updating action field in plan so that attachments are taken care
				updateVrfUpdateActionFlag(&vrf, NewEntry)
				//use the object in state for delete
				delVrfs.Vrfs[vrf.VrfName] = sVrf
				newVRFs.Vrfs[vrf.VrfName] = vrf
			} else if updateAction == ControlFlagUpdate {
				deployVrfs = append(deployVrfs, vrf.VrfName)

			} else {
				//Case 3: attributes have changed - Do update
				// mark the attachments for deploy - if config is true
				updateVrfUpdateActionFlag(&vrf, Deploy)
				putVRFs.Vrfs[vrf.VrfName] = vrf
				tflog.Info(ctx, fmt.Sprintf("%s has changed", vrf.VrfName))
			}
			//put back updates
			vrfConfig.Vrfs[sVrfName] = vrf
		} else {
			//case 4: VRF is missing in plan data - Delete it
			tflog.Info(ctx, fmt.Sprintf("%s Missing in Plan - Needs deletion", sVrfName))
			delVrfs.Vrfs[sVrfName] = sVrf
		}
	}
	//case 5: Deal with New VRFs in plan - Add
	for k, v := range vrfConfig.Vrfs {
		if !v.FilterThisValue {
			v.FabricName = newVRFs.FabricName
			newVRFs.Vrfs[k] = v
		}
	}
	actions["add"] = newVRFs
	actions["put"] = putVRFs
	actions["plan"] = vrfConfig
	actions["state"] = vrfState
	actions["del"] = delVrfs
	actions["deploy"] = deployVrfs

	return actions
}

func (c NDFC) vrfBulkUpdate(ctx context.Context, dg *diag.Diagnostics, ndVRFs *resource_vrf_bulk.NDFCVrfBulkModel) {
	// PUT for each vrf
	payload := ndVRFs.FillVrfPayloadFromModel(nil)

	vrfObj := api.NewVrfAPI(ndVRFs.FabricName, c.GetLock(ResourceVrfBulk), &c.apiClient)

	for i := range payload.Vrfs {
		data, err := json.Marshal(payload.Vrfs[i])
		if err != nil {
			dg.AddError("Marshal Failed", fmt.Sprintf("VRF %s Marshall error %v", payload.Vrfs[i].VrfName, err))
			return
		}
		tflog.Info(ctx, fmt.Sprintf("Update VRF %s", payload.Vrfs[i].VrfName))
		vrfObj.PutVrf = payload.Vrfs[i].VrfName
		res, err := vrfObj.Put(data)
		if err != nil {
			dg.AddError(fmt.Sprintf("VRF %s, Update failed", payload.Vrfs[i].VrfName), fmt.Sprintf("Error %v, response %s", err, res.String()))
			return
		}
		tflog.Info(ctx, fmt.Sprintf("Update VRF %s Successful. Message %s", payload.Vrfs[i].VrfName, res.String()))
	}
}

func (c NDFC) validateVrfsUpdate(ctx context.Context, dg *diag.Diagnostics, vrfBulkPlan *resource_vrf_bulk.VrfBulkModel, vrfState *resource_vrf_bulk.VrfBulkModel) {
	plan := vrfBulkPlan.GetModelData()
	state := vrfState.GetModelData()

	if state.DeployAllAttachments && !plan.DeployAllAttachments {
		dg.AddError("Deploy flag cannot be changed from true to false", "deploy_all_attachments cannot be changed from true to false")
		return
	}
	if plan.DeployAllAttachments {
		return
	}

	// Check if any deploy flags has changed from true to false
	for vrf, v := range plan.Vrfs {
		sVrf, ok := state.Vrfs[vrf]
		if ok {
			if !v.DeployAttachments && sVrf.DeployAttachments {
				dg.AddError(fmt.Sprintf("VRF %s, Deploy flag changed from true to false", v.VrfName), "Deploy flag cannot be changed from true to false")
				return
			}
		} else {
			log.Printf("VRF %s not found in state", v.VrfName)
			continue
		}
		if v.DeployAttachments {
			continue
		}

		for serial, attach := range v.AttachList {
			log.Printf("validateVrfsUpdate: Attachment %s Plan Deploy flag %t", serial, attach.DeployThisAttachment)
			if sAttach, ok := sVrf.AttachList[serial]; ok {
				log.Printf("validateVrfsUpdate: Attachment %s State Deploy flag %t", serial, sAttach.DeployThisAttachment)
				if !attach.DeployThisAttachment && sAttach.DeployThisAttachment {
					dg.AddError(fmt.Sprintf("VRF %s, Attachment %s Deploy flag changed from true to false", vrf, serial), "Deploy flag cannot be changed from true to false")
					return
				}
			} else {
				log.Printf("Attachment %s not found in state", serial)
			}
		}
	}
}

func updateVrfUpdateActionFlag(planVrf *resource_vrf_bulk.NDFCVrfsValue, action uint16) {
	for serial, attach := range planVrf.AttachList {
		attach.UpdateAction = ActionNone
		attach.UpdateAction |= action
		planVrf.AttachList[serial] = attach
	}
}

// VrfSwitchDetailResponse represents the response structure from the VRF switch details API
type VrfSwitchDetailResponse struct {
	VrfName           string                  `json:"vrfName"`
	TemplateName      string                  `json:"templateName"`
	SwitchDetailsList []SwitchDetailsListItem `json:"switchDetailsList"`
}

type SwitchDetailsListItem struct {
	SwitchName       string `json:"switchName"`
	Vlan             int    `json:"vlan"`
	SerialNumber     string `json:"serialNumber"`
	FreeformConfig   string `json:"freeformConfig"`
	InstanceValues   string `json:"instanceValues"`
	ExtensionValues  string `json:"extensionValues"`
	IsLanAttached    bool   `json:"islanAttached"`
	LanAttachedState string `json:"lanAttachedState"`
}

func (c NDFC) fillMissingParams(ctx context.Context, ndVRFs *resource_vrf_bulk.NDFCVrfBulkModel) error {
	// Collect all VRF names and serial numbers
	vrfNames := make([]string, 0)
	serialNumbersSet := make(map[string]bool)

	for vrfName, vrfEntry := range ndVRFs.Vrfs {
		vrfNames = append(vrfNames, vrfName)
		for serial := range vrfEntry.AttachList {
			if vrfEntry.AttachList[serial].FilterThisValue {
				continue
			}
			serialNumbersSet[vrfEntry.AttachList[serial].SwitchSerialNo] = true
		}
	}

	if len(vrfNames) == 0 || len(serialNumbersSet) == 0 {
		tflog.Info(ctx, "fillMissingParams: No VRFs or attachments to process")
		return nil
	}

	tflog.Info(ctx, fmt.Sprintf("fillMissingParams: Fetching details for VRFs=%v, SerialCount=%d", vrfNames, len(serialNumbersSet)))
	// Add counters
	totalUpdated := 0
	totalSkipped := 0
	totalErrors := 0

	// Make individual API calls per serial number to avoid switch type mismatch errors
	// (API requires switches to be of same type Border/Leaf)
	for serial := range serialNumbersSet {
		tflog.Debug(ctx, fmt.Sprintf("fillMissingParams: Fetching details for Serial=%s", serial))

		vrfObj := api.NewVrfAPI(ndVRFs.FabricName, c.GetLock(ResourceVrfBulk), &c.apiClient)
		res, err := vrfObj.GetVrfSwitchDetails(vrfNames, []string{serial})
		if err != nil {
			tflog.Warn(ctx, fmt.Sprintf("fillMissingParams: Error fetching switch details for Serial=%s: %v", serial, err))
			// Continue with other serials even if one fails
			totalErrors++
			continue
		}

		tflog.Debug(ctx, fmt.Sprintf("fillMissingParams: API Response for Serial=%s: %s", serial, string(res)))

		// Parse response
		var switchDetails []VrfSwitchDetailResponse
		err = json.Unmarshal(res, &switchDetails)
		if err != nil {
			tflog.Warn(ctx, fmt.Sprintf("fillMissingParams: Error unmarshalling response for Serial=%s: %v", serial, err))
			// Continue with other serials even if one fails
			totalErrors++
			continue
		}

		// Map freeformConfig to attachments
		for _, vrfDetail := range switchDetails {
			vrfEntry, ok := ndVRFs.Vrfs[vrfDetail.VrfName]
			if !ok {
				continue
			}

			for _, switchDetail := range vrfDetail.SwitchDetailsList {
				if switchDetail.SerialNumber != serial {
					continue
				}

				attachEntry, ok := vrfEntry.AttachList[switchDetail.SerialNumber]
				if ok {
					// Update freeformConfig if not empty in the response
					if switchDetail.FreeformConfig != "" {
						attachEntry.FreeformConfig = switchDetail.FreeformConfig
						vrfEntry.AttachList[switchDetail.SerialNumber] = attachEntry
						tflog.Debug(ctx, fmt.Sprintf("fillMissingParams: Updated freeformConfig for VRF=%s, Serial=%s",
							vrfDetail.VrfName, switchDetail.SerialNumber))
						totalUpdated++
					} else {
						totalSkipped++
					}
				}
			}

			// Put the updated vrfEntry back
			ndVRFs.Vrfs[vrfDetail.VrfName] = vrfEntry
		}
	}

	tflog.Info(ctx, fmt.Sprintf("fillMissingParams: Complete - Updated=%d, Skipped=%d, Errors=%d",
		totalUpdated, totalSkipped, totalErrors))
	return nil
}

func (c NDFC) vrfAttachmentSerialRemap(ctx context.Context, vrf *resource_vrf_bulk.NDFCVrfBulkModel) map[string]string {
	keyMap := make(map[string]string)
	for vrfName, vrfEntry := range vrf.Vrfs {
		for attKey, attachEntry := range vrfEntry.AttachList {
			// check if key is IP or Serial
			if net.ParseIP(attKey) != nil {
				// IP
				serial := c.GetSerialFromIP(ctx, vrf.FabricName, attKey)
				if serial == "" {
					panic(fmt.Sprintf("Failed to get serial for IP %s in fabric %s", attKey, vrf.FabricName))
				}
				attachEntry.SerialNumber = serial
			} else {
				// Serial
				attachEntry.SerialNumber = attKey
			}
			vrfEntry.AttachList[attKey] = attachEntry
			kk := vrfName + ":" + attachEntry.SerialNumber
			keyMap[kk] = attKey
		}
		vrf.Vrfs[vrfName] = vrfEntry
	}
	return keyMap
}

func (c NDFC) vrfAttachmentSerialRemapFromPayload(ctx context.Context, payload *rva.NDFCVrfAttachmentsPayloads) {
	for v := range payload.VrfAttachments {
		vrfEntry := &payload.VrfAttachments[v]
		for i := range vrfEntry.AttachList {
			// check if key is IP or Serial
			if net.ParseIP(vrfEntry.AttachList[i].SerialNumber) != nil {
				// IP
				serial := c.GetSerialFromIP(ctx, payload.FabricName, vrfEntry.AttachList[i].SerialNumber)
				if serial == "" {
					panic(fmt.Sprintf("Failed to get serial for IP %s in fabric %s", vrfEntry.AttachList[i].SerialNumber, payload.FabricName))
				}
				vrfEntry.AttachList[i].SerialNumber = serial
			}
		}
	}
}

func vrfAttachmentKeyMap(ctx context.Context, vrf *resource_vrf_bulk.NDFCVrfBulkModel) map[string]uint16 {
	keyMap := make(map[string]uint16)
	for vrfName, vrfEntry := range vrf.Vrfs {
		for attKey := range vrfEntry.AttachList {
			key := vrfName + ":" + attKey
			if net.ParseIP(attKey) != nil {
				keyMap[key] = KeyTypeIP
			} else {
				keyMap[key] = KeyTypeSerial
			}
		}
	}
	return keyMap
}
