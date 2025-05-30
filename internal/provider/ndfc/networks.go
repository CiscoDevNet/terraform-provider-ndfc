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
	"terraform-provider-ndfc/internal/provider/datasources/datasource_networks"
	"terraform-provider-ndfc/internal/provider/ndfc/api"
	"terraform-provider-ndfc/internal/provider/resources/resource_networks"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const ResourceNetworks = "networks"

func (c *NDFC) RscCreateNetworks(ctx context.Context, dg *diag.Diagnostics, in *resource_networks.NetworksModel) *resource_networks.NetworksModel {
	// Create API call logic
	tflog.Debug(ctx, fmt.Sprintf("RscCreateNetworks entry fabirc %s", in.FabricName.ValueString()))
	nw := in.GetModelData()
	if nw == nil {
		tflog.Error(ctx, "Data conversion from model failed")
		dg.AddError("Data conversion from model failed", "GetModelData returned empty")
		return nil
	}
	//form ID
	ID := c.RscCreateID(nw, ResourceNetworks)
	//create

	depMap := make(map[string][]string)

	retNets, err := c.networksIsPresent(ctx, ID)

	if err != nil {
		tflog.Error(ctx, "Error while getting Networks ", map[string]interface{}{"Err": err})
		dg.AddError("Network Read Failed", err.Error())
		return nil
	}

	var errs []string
	for i := range retNets {
		errs = append(errs, fmt.Sprintf("Network %s is already configured on %s", retNets[i], nw.FabricName))
	}
	if len(errs) > 0 {
		tflog.Error(ctx, "Networks exist", map[string]interface{}{"Err": errs})
		dg.AddError("Networks exist", strings.Join(errs, ","))
		return nil
	}
	// Check if VRFs referenced in the networks exist and they have same attachments

	err = c.CheckNetworkVrfConfig(ctx, dg, nw)
	if err != nil {
		tflog.Error(ctx, "CheckNetworkVrfConfig failed", map[string]interface{}{"Err": err})
		return nil
	}

	//Part 1: Create Networks

	err = c.networksCreate(ctx, nw.FabricName, nw)
	if err != nil {
		tflog.Error(ctx, "Cannot create Network", map[string]interface{}{"Err": err})
		dg.AddError("Cannot create Network", err.Error())
		return nil
	}
	tflog.Info(ctx, fmt.Sprintf("Create Bulk VRF success ID %s", ID))
	//Part 2: Create Attachments if any

	if nw.DeployAllAttachments {
		depMap["global"] = append(depMap["global"], "all")
	}
	for i, entry := range nw.Networks {
		if entry.DeployAttachments {
			depMap[i] = append(depMap[i], i)
		}
		for j, attachEntry := range entry.Attachments {
			if attachEntry.DeployThisAttachment {
				depMap[i] = append(depMap[i], j)
			}
		}
	}

	err = c.netAttachmentsAttach(ctx, nw)
	if err != nil {
		tflog.Error(ctx, "Network Attachments create failed")
		dg.AddError("Network Attachments create failed", err.Error())
		tflog.Error(ctx, "Rolling back the configurations...delete Networks")
		c.RscDeleteNetworks(ctx, dg, ID, in)
		return nil
	}
	//Part 3: Deploy Attachments if any
	c.RscDeployNetworkAttachments(ctx, dg, nw)

	out := c.RscGetNetworks(ctx, dg, ID, &depMap)
	if out == nil {
		tflog.Error(ctx, "Failed to verify: Reading from NDFC after create failed")
		dg.AddError("Failed to verify", "Reading from NDFC after create failed")
		return nil
	}
	if ID != "" {
		out.Id = types.StringValue(ID)
	}
	return out
}

func (c *NDFC) RscGetNetworks(ctx context.Context, dg *diag.Diagnostics, ID string, depMap *map[string][]string) *resource_networks.NetworksModel {
	// Read API call logic
	var filterMap map[string]bool
	tflog.Debug(ctx, fmt.Sprintf("RscGetNetworks entry fabirc %s", ID))

	filterMap = make(map[string]bool)
	fabricName, networks := c.CreateFilterMap(ID, &filterMap)

	log.Printf("FilterMap: %v networks %v", filterMap, networks)

	if fabricName == "" {
		dg.AddError("ID format error", "ID is incorrect")
		return nil
	}
	ndNets, err := c.networksGet(ctx, fabricName)
	if err != nil {
		dg.AddError("Network Get Failed", err.Error())
		return nil
	}

	if len(ndNets.Networks) == 0 {
		dg.AddWarning("No Networks found", "No Networks found in NDFC")
		return nil
	}
	netCount := 0
	if len(filterMap) > 0 {
		log.Printf("Filtering is configured")
		//Set value filter - skip the nws that are not in ID
		for i, entry := range ndNets.Networks {
			if _, found := (filterMap)[entry.NetworkName]; !found {
				log.Printf("Filtering out Network %s", entry.NetworkName)
				entry.FilterThisValue = true
				ndNets.Networks[i] = entry
			} else {
				netCount++
			}
		}
	}
	if netCount == 0 {
		dg.AddWarning("No Networks found", "No Networks found in NDFC")
		return nil
	}
	globalDep := false

	if _, ok := (*depMap)["global"]; ok {
		globalDep = true
		//This cannot be validated - as we don't know if all deployments were ok
		ndNets.DeployAllAttachments = true
	}

	//Get Attachments

	err = c.RscGetNetworkAttachments(ctx, ndNets)
	if err == nil {
		tflog.Debug(ctx, "Network Attachments read success")

		for i, nwEntry := range ndNets.Networks {
			if nwEntry.FilterThisValue {
				continue
			}
			vrfLevelDep := false

			if vl, vlOk := (*depMap)[i]; vlOk {
				//first element is network name if network level deploy is set
				if vl[0] == i {
					nwEntry.DeployAttachments = (nwEntry.NetworkStatus == "DEPLOYED")
					log.Printf("Setting Network level dep flag for %s to %v", i, nwEntry.DeployAttachments)
					vrfLevelDep = true
				}
			}

			for j, attachEntry := range nwEntry.Attachments {
				if attachEntry.FilterThisValue {
					continue
				}
				attachLevelDep := false
				deps := (*depMap)[i]
				if len(deps) > 0 {
					for _, dep := range deps {
						if dep == j {
							attachLevelDep = true
							log.Printf("Setting Attachment level dep flag is set for  %s/%s", i, j)
						}
					}
				}
				log.Printf("Attachment %s added to Network %s", j, i)
				if !globalDep && !vrfLevelDep {
					if attachEntry.AttachState == "DEPLOYED" && attachLevelDep {
						log.Printf("Network Attachment %s/%s deployed flag set to true", i, j)
						attachEntry.DeployThisAttachment = true
					}
				}

				if portOrder, ok := (*depMap)["SwitchPorts:"+i+"/"+j]; ok {
					log.Printf("SwitchPorts order %v", portOrder)
					processPortListOrder(&attachEntry.SwitchPorts, portOrder)
				}
				if portOrder, ok := (*depMap)["TorPorts:"+i+"/"+j]; ok {
					log.Printf("TorPorts order %v", portOrder)
					processPortListOrder(&attachEntry.TorPorts, portOrder)
				}
				//put modified entry back
				nwEntry.Attachments[j] = attachEntry
			}
			//put modified entry back
			ndNets.Networks[i] = nwEntry
		}

	} else {
		tflog.Error(ctx, "Network Attachments read failed", map[string]interface{}{"Err": err})
		dg.AddError("Network Attachments read failed", err.Error())
	}

	vModel := new(resource_networks.NetworksModel)
	vModel.Id = types.StringValue(ID)
	d := vModel.SetModelData(ndNets)
	if d != nil {
		dg.Append(d.Errors()...)
	}
	return vModel
}

// Import is similar to GET with minor changes. Not combining to keep it simple
func (c *NDFC) RscImportNetworks(ctx context.Context, dg *diag.Diagnostics, ID string) *resource_networks.NetworksModel {
	// Read API call logic
	var filterMap map[string]bool
	tflog.Debug(ctx, fmt.Sprintf("RscImportNetworks entry fabirc %s", ID))

	filterMap = make(map[string]bool)
	fabricName, networks := c.CreateFilterMap(ID, &filterMap)

	log.Printf("FilterMap: %v networks %v", filterMap, networks)

	if fabricName == "" {
		dg.AddError("ID format error", "ID is incorrect")
		return nil
	}

	re := regexp.MustCompile(`^[\w-]+\/\[(?:[\w-]+\,?)+\]$`)
	if !re.Match([]byte(ID)) {
		dg.AddError("ID format error", "use fabricName/[net1,net2...] format")
		return nil
	}

	ndNets, err := c.networksGet(ctx, fabricName)
	if err != nil {
		dg.AddError("Network Get Failed", err.Error())
		return nil
	}

	if len(ndNets.Networks) == 0 {
		dg.AddError("No Networks found", "No Networks found in NDFC")
		return nil
	}
	netCount := 0
	if len(filterMap) > 0 {
		log.Printf("Filtering is configured")
		//Set value filter - skip the nws that are not in ID
		for i, entry := range ndNets.Networks {
			if _, found := (filterMap)[entry.NetworkName]; !found {
				log.Printf("Filtering out Network %s", entry.NetworkName)
				entry.FilterThisValue = true
				ndNets.Networks[i] = entry
			} else {
				netCount++
			}
		}
	}
	if netCount == 0 {
		dg.AddError("No Networks found", "No Networks found in NDFC")
		return nil
	}
	//Get Attachments
	err = c.RscGetNetworkAttachments(ctx, ndNets)
	if err == nil {
		tflog.Debug(ctx, "Network Attachments read success")
		for i, nwEntry := range ndNets.Networks {
			if nwEntry.FilterThisValue {
				continue
			}
			for j, attachEntry := range nwEntry.Attachments {
				if attachEntry.FilterThisValue {
					continue
				}
				log.Printf("Attachment %s added to Network %s", j, i)

				if attachEntry.AttachState == "DEPLOYED" {
					log.Printf("Attachment %s deployed", j)
					attachEntry.DeployThisAttachment = true
				}
				//put modified entry back
				nwEntry.Attachments[j] = attachEntry
			}
			//put modified entry back
			ndNets.Networks[i] = nwEntry
		}
	} else {
		tflog.Error(ctx, "Network Attachments read failed", map[string]interface{}{"Err": err})
		dg.AddError("Network Attachments read failed", err.Error())
	}
	vModel := new(resource_networks.NetworksModel)
	vModel.Id = types.StringValue(ID)
	d := vModel.SetModelData(ndNets)
	if d != nil {
		dg.Append(d.Errors()...)
	}
	return vModel
}

// plan is updated at the end to reflect the current config and returned to TF
func (c *NDFC) RscUpdateNetworks(ctx context.Context, dg *diag.Diagnostics, ID string,
	plan *resource_networks.NetworksModel,
	state *resource_networks.NetworksModel,
	config *resource_networks.NetworksModel) {

	netActions := c.networksGetDiff(ctx, dg, plan, state, config)

	delNw := netActions["del"].(*resource_networks.NDFCNetworksModel)
	putNw := netActions["put"].(*resource_networks.NDFCNetworksModel)
	newNw := netActions["add"].(*resource_networks.NDFCNetworksModel)
	planNw := netActions["plan"].(*resource_networks.NDFCNetworksModel)
	stateNw := netActions["state"].(*resource_networks.NDFCNetworksModel)

	netAttachActions := c.networkAttachmentsGetDiff(ctx, dg, planNw, stateNw, putNw)

	//******** validations begin *****************************************************

	// 1. Check if networks to be added exists
	nws, err := c.networksGet(ctx, plan.FabricName.ValueString())
	if err != nil {
		tflog.Error(ctx, "Network Get Failed", map[string]interface{}{"Err": err})
		dg.AddError("Network Get Failed", err.Error())
		return
	}

	// 2. Check VRF dependencies of networks to be added

	err = c.CheckNetworkVrfConfig(ctx, dg, planNw)
	if err != nil {
		tflog.Error(ctx, "CheckNetworkVrfConfig failed", map[string]interface{}{"Err": err})
		return
	}

	for k := range newNw.Networks {
		if _, ok := nws.Networks[k]; ok {
			// check if the network is marked for deletion
			if _, ok := delNw.Networks[k]; ok {
				tflog.Debug(ctx, "Network marked for deletion", map[string]interface{}{"Network": k})
				continue
			}
			tflog.Error(ctx, "Network already exists", map[string]interface{}{"Network": k})
			dg.AddError("Network already exists", fmt.Sprintf("Network %s already exists", k))
			return
		}
	}

	// 3 Check if networks to be deleted exists

	for k := range delNw.Networks {
		if _, ok := nws.Networks[k]; !ok {
			// Should we error out in this situation or just continue ?
			tflog.Error(ctx, "Network does not exist", map[string]interface{}{"Network": k})
			dg.AddWarning("Network does not exist", fmt.Sprintf("Network %s does not exist", k))
		}
		// deleting a network affects the VRF attachments
		// add a warning so that user is aware
		if len(nws.Networks[k].Attachments) > 0 {
			dg.AddWarning("NDFC VRF config change!", fmt.Sprintf("Network %s has attachments, which may affect attachments in associated VRF %s", k, nws.Networks[k].VrfName))
		}
	}

	// 4 Check if networks to be updated exists

	for k := range putNw.Networks {
		if _, ok := nws.Networks[k]; !ok {
			tflog.Error(ctx, "Network does not exist", map[string]interface{}{"Network": k})
			dg.AddError("Network does not exist", fmt.Sprintf("Network %s does not exist", k))
			return
		}
		// check and add warning if VRF is changed
		// NDFC changes associated VRF attachments when a network is changed
		if stateNw.Networks[k].VrfName != planNw.Networks[k].VrfName {
			dg.AddWarning("Changing VRF association in network", fmt.Sprintf("VRF change  from %s to %s in network %s would impact the VRF attachments in both", stateNw.Networks[k].VrfName, planNw.Networks[k].VrfName, k))
		}
	}
	// ************** validations end ****************************************************
	// ************** update begin *******************************************************

	// 1. Delete items marked for delete as well as re-create

	if len(delNw.Networks) > 0 {
		// Detach Attachments
		err := c.RscDeleteNetAttachments(ctx, dg, delNw)
		if err != nil {
			tflog.Error(ctx, "Network Attachments delete failed", map[string]interface{}{"Err": err})
			return
		}
		// Delete Networks
		err = c.networksDelete(ctx, delNw.FabricName, delNw.GetNetworksNames())
		if err != nil {
			tflog.Error(ctx, "Networks delete failed", map[string]interface{}{"Err": err})
			dg.AddError("Networks delete failed", err.Error())
			return
		}

	}

	//2. Update items marked for update
	if len(putNw.Networks) > 0 {
		c.networksUpdate(ctx, dg, putNw, 0)
		if dg.HasError() {
			//Roll back here is pretty complex - so not attempting
			tflog.Error(ctx, "Networks update failed", map[string]interface{}{"Err": err})
			dg.AddError("Networks update failed", "Rollback not attempted")
			return
		}
	}

	//3. Create new items
	if len(newNw.Networks) > 0 {

		err := c.networksCreate(ctx, newNw.FabricName, newNw)
		if err != nil {
			//On failure - creation api deletes any successful entry
			tflog.Error(ctx, "Networks create failed", map[string]interface{}{"Err": err})
			dg.AddError("Networks create failed", err.Error())
			return
		}
	}

	// Deal with attachments
	c.RscUpdateNetAttachments(ctx, dg, netAttachActions)
	if dg.HasError() {
		tflog.Error(ctx, "Network Attachments update failed")
		return
	}

	newID := c.RscCreateID(plan.GetModelData(), ResourceNetworks)
	tflog.Info(ctx, fmt.Sprintf("New ID after update %s", newID))

	depMap := make(map[string][]string)

	if planNw.DeployAllAttachments {
		depMap["global"] = append(depMap["global"], "all")
	}
	for i, entry := range planNw.Networks {
		if entry.DeployAttachments {
			depMap[i] = append(depMap[i], i)
		}
		for j, attachEntry := range entry.Attachments {
			if attachEntry.DeployThisAttachment {
				depMap[i] = append(depMap[i], j)
			}
		}
	}
	log.Printf("Deploy Map %v", depMap)

	//4. Read the updated data from NDFC
	*plan = *(c.RscGetNetworks(ctx, dg, newID, &depMap))
}

func (c *NDFC) RscDeleteNetworks(ctx context.Context, dg *diag.Diagnostics, ID string, in *resource_networks.NetworksModel) {
	// Delete API call logic

	// Get Network names from ID
	rsList, err := c.networksIsPresent(ctx, ID)
	if err != nil {
		tflog.Error(ctx, "Error while getting Networks ", map[string]interface{}{"Err": err})
		dg.AddError("Network Read Failed", err.Error())
		return
	}

	if len(rsList) == 0 {
		tflog.Info(ctx, "No Networks to delete")
		return
	}

	results := c.RscBulkSplitID(ID)
	fabricName := results["fabric"][0]
	nwFromId := results["rsc"]

	if len(nwFromId) != len(rsList) {
		tflog.Error(ctx, "Mismatch in Network - Some entries are already deleted", map[string]interface{}{
			"ID":       ID,
			"nwFromId": nwFromId,
			"rsList":   rsList,
		})
		dg.AddWarning("Deleting only what is available in NDFC", "ID is not up to date")
		//rsList = nwFromId
	}

	// Detach Attachments
	err = c.RscDeleteNetAttachments(ctx, dg, in.GetModelData())
	if err != nil {
		tflog.Error(ctx, "Network Attachments delete failed", map[string]interface{}{"Err": err})
		dg.AddError("Network Attachments delete failed", err.Error())
		return
	}

	err = c.networksDelete(ctx, fabricName, rsList)
	if err != nil {
		tflog.Error(ctx, "Networks delete failed", map[string]interface{}{"Err": err})
		dg.AddError("Networks delete failed", err.Error())
		return
	}
	tflog.Info(ctx, "Networks delete success")

}

func (c *NDFC) DsGetNetworks(ctx context.Context, dg *diag.Diagnostics, fabricName string) *datasource_networks.NetworksModel {
	// Read API call logic
	tflog.Debug(ctx, fmt.Sprintf("DsGetNetworks entry fabirc %s", fabricName))

	if fabricName == "" {
		dg.AddError("ID format error", "ID is incorrect")
		return nil
	}
	rsObj := api.NewNetworksAPI(fabricName, c.GetLock(ResourceNetworks), &c.apiClient)
	res, err := rsObj.Get()
	if err != nil {
		dg.AddError("Network Get Failed", err.Error())
		return nil
	}
	ndNets := new(datasource_networks.NDFCNetworksModel)

	err = json.Unmarshal(res, &ndNets.Networks)
	if err != nil {
		dg.AddError("Network Unmarshal Failed", err.Error())
		return nil
	}
	if len(ndNets.Networks) == 0 {
		dg.AddWarning("No Networks found", "No Networks found in NDFC")
		return nil
	}

	ndNets.CreateSearchMap()
	//Get Attachments
	var nwAttachments []datasource_networks.NDFCNetworkAttachmentsPayload
	rsObjAttach := api.NewNetAttachAPI(fabricName, c.GetLock(ResourceNetworks), &c.apiClient)

	rsObjAttach.GetnwList = ndNets.GetNetworksNames()
	resAttach, err := rsObjAttach.Get()
	if err != nil {
		dg.AddError("Network Attachments Get Failed", err.Error())
		return nil
	}
	err = json.Unmarshal(resAttach, &nwAttachments)
	if err != nil {
		dg.AddError("Network Attachments Unmarshal Failed", err.Error())
		return nil
	}

	//Fill Attachments

	for _, attach := range nwAttachments {

		if nwEntry, ok := ndNets.NetworksMap[attach.NetworkName]; ok {
			for _, attachEntry := range attach.Attachments {
				//skip implicit attachments
				if *attachEntry.Attached {
					nwEntry.Attachments = append(nwEntry.Attachments, attachEntry)
				}
			}
		}
	}
	vModel := new(datasource_networks.NetworksModel)
	vModel.FabricName = types.StringValue(fabricName)
	d := vModel.SetModelData(ndNets)
	if d != nil {
		dg.Append(d.Errors()...)
	}
	return vModel
}
