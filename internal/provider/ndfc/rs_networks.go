// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package ndfc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"terraform-provider-ndfc/internal/provider/ndfc/api"
	"terraform-provider-ndfc/internal/provider/resources/resource_networks"
	. "terraform-provider-ndfc/internal/provider/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func (c NDFC) networksCreate(ctx context.Context, fabricName string, rscModel *resource_networks.NDFCNetworksModel) error {
	tflog.Info(ctx, fmt.Sprintf("Beginning Bulk Networks create in fabric %s", fabricName))
	payload := rscModel.FillNetworksPayloadFromModel()

	data, err := json.Marshal(payload.Networks)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Json Marshal failure %s", err.Error()))
		return err
	}
	log.Println("Data to be posted", string(data))

	rsObj := api.NewNetworksAPI(fabricName, c.GetLock(ResourceNetworks), &c.apiClient)
	res, err := rsObj.Post(data)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Error POST:  %s", err.Error()))
		okList, err1 := c.processBulkResponse(ctx, res)
		err2 := c.networksDelete(ctx, fabricName, okList)
		return errors.Join(err, err1, err2)
	}

	tflog.Info(ctx, fmt.Sprintf("resource create: Success res : %v", res.String()))
	return nil
}

func (c NDFC) networksDelete(ctx context.Context, fabricName string, rsList []string) error {
	if len(rsList) == 0 {
		return nil
	}
	tflog.Info(ctx, fmt.Sprintf("Attempting to delete resource fabric_name=%s, resources = %v", fabricName, rsList))
	rsObj := api.NewNetworksAPI(fabricName, c.GetLock(ResourceNetworks), &c.apiClient)
	rsObj.SetDeleteList(rsList)
	res, err := rsObj.Delete()
	if err != nil {
		_, err1 := c.processBulkResponse(ctx, res)
		return err1
	}
	tflog.Info(ctx, fmt.Sprintf("Deleting resources OK fabric_name=%s, resource names = %v", fabricName, rsList))
	return nil

}

func (c NDFC) networksGet(ctx context.Context, fabricName string) (*resource_networks.NDFCNetworksModel, error) {

	rsObj := api.NewNetworksAPI(fabricName, c.GetLock(ResourceNetworks), &c.apiClient)
	res, err := rsObj.Get()
	if err != nil {
		return nil, err
	}
	tflog.Debug(ctx, fmt.Sprintf("RscGetNetworks: result %s", string(res)))
	rsModel := resource_networks.NDFCNetworksModel{}
	payloads := resource_networks.NDFCNetworksPayload{}
	err = json.Unmarshal(res, &payloads.Networks)
	if err != nil {
		return nil, err
	} else {
		tflog.Debug(ctx, "resource_networks: Unmarshal OK")
	}
	rsModel.FillNetworksFromPayload(&payloads)
	return &rsModel, nil
}

func (c NDFC) networksIsPresent(ctx context.Context, ID string) ([]string, error) {
	var ret []string
	filterMap := make(map[string]bool)

	fabricName, _ := c.CreateFilterMap(ID, &filterMap)
	rsModel, err := c.networksGet(ctx, fabricName)
	if err != nil {
		tflog.Error(ctx, "Error while getting resource ", map[string]interface{}{"Err": err})
		return nil, err
	}
	for k, _ := range rsModel.Networks {
		ok, found := filterMap[k]
		if ok && found {
			ret = append(ret, k)
		}
	}
	return ret, nil
}

func (c NDFC) networksUpdate(ctx context.Context, dg *diag.Diagnostics, updateRsc *resource_networks.NDFCNetworksModel, retry int) {
	// PUT for each object
	payload := updateRsc.FillNetworksPayloadFromModel()

	if retry > 5 {
		dg.AddError("Update Failed", "Retry count exceeded")
		return
	}
	retryIndices := make([]int, 0)
	rsObj := api.NewNetworksAPI(updateRsc.FabricName, c.GetLock(ResourceNetworks), &c.apiClient)

	for i := range payload.Networks {
		data, err := json.Marshal(payload.Networks[i])
		if err != nil {
			dg.AddError("Marshal Failed", fmt.Sprintf("Resource %s Marshall error %v", payload.Networks[i].NetworkName, err))
			return
		}
		tflog.Info(ctx, fmt.Sprintf("Update resource %s", payload.Networks[i].NetworkName))
		rsObj.PutNetworkName = payload.Networks[i].NetworkName
		res, err := rsObj.Put(data)
		if err != nil {
			log.Printf("Error PUT: %v |%s|", res, res.String())
			dg.AddError(fmt.Sprintf("Resource %s, Update failed", payload.Networks[i].NetworkName), fmt.Sprintf("Error %v, response %s", err, res.String()))
			return
		}
		// Get and check if they Mismatch
		// There is a bug in NDFC where PUT sometimes resets some values to DefaultValue
		// This is mostly seen with dhcpServers payload
		// Re-do ing PUT solves the problem in most cases
		// Hence do a GET nd verify if all params got updated
		// if not add them to a list and re-do PUT until its correct or retry count is exceeded
		rsObj.GetNetworkName = payload.Networks[i].NetworkName
		rs, err := rsObj.Get()
		if err != nil {
			tflog.Error(ctx, "Read resource after PUT Failed")
			dg.AddError(fmt.Sprintf("Resource %s, Get failed", payload.Networks[i].NetworkName), fmt.Sprintf("Error %v, response %s", err, string(rs)))
			return
		}
		rsNewValue := resource_networks.NDFCNetworksValue{}
		err = json.Unmarshal(rs, &rsNewValue)
		if err != nil {
			tflog.Error(ctx, "Unmarshal Failed, networks GET followed by PUT")
			dg.AddError(fmt.Sprintf("Resource %s, Unmarshal failed", payload.Networks[i].NetworkName), fmt.Sprintf("Error %v, response %s", err, res.String()))
			return
		}
		changes := payload.Networks[i].DeepEqual(rsNewValue)
		if changes != ValuesDeeplyEqual && changes != ControlFlagUpdate {
			tflog.Error(ctx, "Mismatch in data retrieved after PUT - add to retry list")
			retryIndices = append(retryIndices, i)
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Update resource %s Successfull. Message %s", payload.Networks[i].NetworkName, res.String()))
	}
	if len(retryIndices) > 0 {
		tflog.Info(ctx, "Retrying network update due to mismatch", map[string]interface{}{"Err": "Mismatch in network data retrieved after PUT"})
		redoRsc := new(resource_networks.NDFCNetworksModel)
		redoRsc.FabricName = updateRsc.FabricName
		redoRsc.Networks = make(map[string]resource_networks.NDFCNetworksValue)
		for _, i := range retryIndices {
			redoRsc.Networks[payload.Networks[i].NetworkName] = payload.Networks[i]
		}
		c.networksUpdate(ctx, dg, redoRsc, retry+1)
	}
}

func (c NDFC) networksGetDiff(ctx context.Context, dg *diag.Diagnostics,
	vPlan *resource_networks.NetworksModel,
	vState *resource_networks.NetworksModel, vConfig *resource_networks.NetworksModel) map[string]interface{} {

	actions := make(map[string]interface{})
	rsState := vState.GetModelData()
	rsConfig := vPlan.GetModelData()

	//rsConfig := vConfig.GetModelData()
	//var delRscs []string
	var deployRscs []string

	putRscs := new(resource_networks.NDFCNetworksModel)
	putRscs.Networks = make(map[string]resource_networks.NDFCNetworksValue)
	putRscs.FabricName = vPlan.FabricName.ValueString()

	newRscs := new(resource_networks.NDFCNetworksModel)
	newRscs.Networks = make(map[string]resource_networks.NDFCNetworksValue)
	newRscs.FabricName = vPlan.FabricName.ValueString()

	delRscs := new(resource_networks.NDFCNetworksModel)
	delRscs.Networks = make(map[string]resource_networks.NDFCNetworksValue)
	delRscs.FabricName = vPlan.FabricName.ValueString()

	for sRsName, sRsc := range rsState.Networks {
		if rsPlan, ok := rsConfig.Networks[sRsName]; ok {
			rsPlan.FilterThisValue = true
			rsPlan.FabricName = newRscs.FabricName
			rsPlan.NetworkName = sRsName
			cf := false
			updateAction := rsPlan.CreatePlan(sRsc, &cf) //dummy cf
			if updateAction == ActionNone {
				//Case 1: Both networks entries are equal - no change
				tflog.Info(ctx, fmt.Sprintf("%s not changed", sRsName))

			} else if updateAction == RequiresReplace {
				//Case 2: attribute that cannot be modified in-place has changed - DELETE and Create
				tflog.Info(ctx, fmt.Sprintf("%s Needs to be replaced - Delete and Add", sRsName))
				//use the object in state for delete
				delRscs.Networks[rsPlan.NetworkName] = sRsc
				newRscs.Networks[rsPlan.NetworkName] = rsPlan
			} else if updateAction == ControlFlagUpdate {
				deployRscs = append(deployRscs, rsPlan.NetworkName)

			} else {
				//Case 3: attributes have changed - Do update
				putRscs.Networks[rsPlan.NetworkName] = rsPlan
				tflog.Info(ctx, fmt.Sprintf("%s has changed", rsPlan.NetworkName))
			}
			//put back updates
			rsConfig.Networks[sRsName] = rsPlan
		} else {
			//case 4: Rsc is missing in plan data - Delete it
			tflog.Info(ctx, fmt.Sprintf("%s Missing in Plan - Needs deletion", sRsName))
			delRscs.Networks[sRsName] = sRsc
		}
	}
	//case 5: Deal with New Rscs in plan - Add
	for k, v := range rsConfig.Networks {
		if !v.FilterThisValue {
			v.FabricName = newRscs.FabricName
			newRscs.Networks[k] = v
		}
	}
	actions["add"] = newRscs
	actions["put"] = putRscs
	actions["plan"] = rsConfig
	actions["state"] = rsState
	actions["del"] = delRscs
	actions["deploy"] = deployRscs

	return actions
}

// NetworkSwitchDetailResponse represents the response structure from the network switch details API
type NetworkSwitchDetailResponse struct {
	NetworkName       string                         `json:"networkName"`
	TemplateName      string                         `json:"templateName"`
	SwitchDetailsList []NetworkSwitchDetailsListItem `json:"switchDetailsList"`
}

type NetworkSwitchDetailsListItem struct {
	SwitchName       string `json:"switchName"`
	Vlan             int    `json:"vlan"`
	SerialNumber     string `json:"serialNumber"`
	FreeformConfig   string `json:"freeformConfig"`
	InstanceValues   string `json:"instanceValues"`
	ExtensionValues  string `json:"extensionValues"`
	IsLanAttached    bool   `json:"islanAttached"`
	LanAttachedState string `json:"lanAttachedState"`
}

// fillNetworkAttachmentMissingParams fetches additional attachment details like freeformConfig
func (c NDFC) fillNetworkAttachmentMissingParams(ctx context.Context, ndNetworks *resource_networks.NDFCNetworksModel, keyMap *map[string]string) error {
	// Collect all network names and serial numbers
	networkNames := make([]string, 0)
	serialNumbersSet := make(map[string]bool)

	for networkName, networkEntry := range ndNetworks.Networks {
		networkNames = append(networkNames, networkName)
		for attachKey, attachEntry := range networkEntry.Attachments {
			if attachEntry.FilterThisValue {
				continue
			}
			// Use SerialNumber field if set, otherwise use key
			serial := attachEntry.SwitchSerialNo
			if serial == "" {
				serial = attachKey
			}
			serialNumbersSet[serial] = true
		}
	}

	if len(networkNames) == 0 || len(serialNumbersSet) == 0 {
		tflog.Info(ctx, "fillNetworkAttachmentMissingParams: No networks or attachments to process")
		return nil
	}

	tflog.Info(ctx, fmt.Sprintf("fillNetworkAttachmentMissingParams: Fetching details for Networks=%v, SerialCount=%d", networkNames, len(serialNumbersSet)))
	// Add counters
	totalUpdated := 0
	totalSkipped := 0
	totalErrors := 0

	// Make individual API calls per serial number to avoid switch type mismatch errors
	// (API requires switches to be of same type Border/Leaf)
	for serial := range serialNumbersSet {
		tflog.Debug(ctx, fmt.Sprintf("fillNetworkAttachmentMissingParams: Fetching details for Serial=%s", serial))

		networkObj := api.NewNetworksAPI(ndNetworks.FabricName, c.GetLock(ResourceNetworks), &c.apiClient)
		res, err := networkObj.GetNetworkSwitchDetails(networkNames, []string{serial})
		if err != nil {
			tflog.Warn(ctx, fmt.Sprintf("fillNetworkAttachmentMissingParams: Error fetching switch details for Serial=%s: %v", serial, err))
			// Continue with other serials even if one fails
			totalErrors++
			continue
		}

		tflog.Debug(ctx, fmt.Sprintf("fillNetworkAttachmentMissingParams: API Response for Serial=%s: %s", serial, string(res)))

		// Parse response
		var switchDetails []NetworkSwitchDetailResponse
		err = json.Unmarshal(res, &switchDetails)
		if err != nil {
			tflog.Warn(ctx, fmt.Sprintf("fillNetworkAttachmentMissingParams: Error unmarshalling response for Serial=%s: %v", serial, err))
			// Continue with other serials even if one fails
			totalErrors++
			continue
		}

		// Map freeformConfig to attachments
		for _, networkDetail := range switchDetails {
			networkEntry, ok := ndNetworks.Networks[networkDetail.NetworkName]
			if !ok {
				continue
			}

			for _, switchDetail := range networkDetail.SwitchDetailsList {
				if switchDetail.SerialNumber != serial {
					continue
				}

				// Find attachment key - could be IP or serial
				attachKey := switchDetail.SerialNumber
				if keyMap != nil {
					kk := networkDetail.NetworkName + ":" + switchDetail.SerialNumber
					if key, found := (*keyMap)[kk]; found {
						attachKey = key
					}
				}

				attachEntry, ok := networkEntry.Attachments[attachKey]
				if ok {
					// Update freeformConfig if not empty in the response
					if switchDetail.FreeformConfig != "" {
						attachEntry.FreeformConfig = switchDetail.FreeformConfig
						networkEntry.Attachments[attachKey] = attachEntry
						tflog.Debug(ctx, fmt.Sprintf("fillNetworkAttachmentMissingParams: Updated freeformConfig for Network=%s, Serial=%s",
							networkDetail.NetworkName, switchDetail.SerialNumber))
						totalUpdated++
					} else {
						totalSkipped++
					}
				}
			}

			// Put the updated networkEntry back
			ndNetworks.Networks[networkDetail.NetworkName] = networkEntry
		}
	}

	tflog.Info(ctx, fmt.Sprintf("fillNetworkAttachmentMissingParams: Complete - Updated=%d, Skipped=%d, Errors=%d",
		totalUpdated, totalSkipped, totalErrors))
	return nil
}
