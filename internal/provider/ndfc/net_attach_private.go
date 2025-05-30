// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Generated code Do not EDIT
package ndfc

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"terraform-provider-ndfc/internal/provider/ndfc/api"
	rna "terraform-provider-ndfc/internal/provider/resources/resource_network_attachments"
	"terraform-provider-ndfc/internal/provider/resources/resource_networks"
	"terraform-provider-ndfc/internal/provider/resources/resource_vpc_pair"
	"terraform-provider-ndfc/internal/provider/types"
	. "terraform-provider-ndfc/internal/provider/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func (c NDFC) netAttachmentsAttach(ctx context.Context, rscModel *resource_networks.NDFCNetworksModel) error {
	tflog.Info(ctx, fmt.Sprintf("Beginning Bulk NetworkAttachments create in fabric %s", rscModel.FabricName))
	payload := rna.NDFCNetworkAttachments{}
	rscModel.FillAttachmentsPayloadFromModel(&payload, resource_networks.NwAttachmentAttach)
	if len(payload.NetworkAttachments) == 0 {
		tflog.Info(ctx, "No Attachments to create")
		return nil
	}
	data, err := json.Marshal(payload.NetworkAttachments)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Json Marshal failure %s", err.Error()))
		return err
	}
	return c.netAttachmentsPostPayload(ctx, rscModel.FabricName, data)

}

func (c NDFC) netAttachmentsDetach(ctx context.Context, rscModel *resource_networks.NDFCNetworksModel) (error, int) {
	tflog.Info(ctx, fmt.Sprintf("Beginning Bulk NetworkAttachments delete in fabric %s", rscModel.FabricName))
	payload := rna.NDFCNetworkAttachments{}
	rscModel.FillAttachmentsPayloadFromModel(&payload, resource_networks.NwAttachmentDetach)

	if len(payload.NetworkAttachments) == 0 {
		tflog.Info(ctx, "No Attachments to delete")
		return nil, 0
	}

	data, err := json.Marshal(payload.NetworkAttachments)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Json Marshal failure %s", err.Error()))
		return err, len(payload.NetworkAttachments)
	}
	return c.netAttachmentsPostPayload(ctx, rscModel.FabricName, data), len(payload.NetworkAttachments)
}

func (c NDFC) netAttachmentsPostPayload(ctx context.Context, fabricName string, payload []byte) error {
	tflog.Info(ctx, fmt.Sprintf("Posting  NetworkAttachments  - %s", fabricName))
	log.Println("Data to be posted:", string(payload))

	rsObj := api.NewNetAttachAPI(fabricName, c.GetLock(ResourceNetworks), &c.apiClient)
	res, err := rsObj.Post(payload)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Error POST:  %s", err.Error()))
		return err

	}
	log.Printf("netAttachmentsPostPayload: Response %s", res.Str)
	err1 := c.processAttachResponse(res)
	if err1 != nil {
		tflog.Error(ctx, fmt.Sprintf("Error POST:  %s", err1.Error()))
		return err1
	}
	return nil
}

func (c NDFC) netAttachmentsGet(ctx context.Context, fabricName string, netList []string) (*rna.NDFCNetworkAttachments, error) {

	rsObj := api.NewNetAttachAPI(fabricName, c.GetLock(ResourceNetworks), &c.apiClient)
	rsObj.GetnwList = netList
	res, err := rsObj.Get()
	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, fmt.Sprintf("netAttachmentsGet: result %s", string(res)))
	payloads := rna.NDFCNetworkAttachments{}
	err = json.Unmarshal(res, &payloads.NetworkAttachments)
	if err != nil {
		return nil, err
	} else {
		tflog.Debug(ctx, "resource_network_attachments: Unmarshal OK")
	}
	return &payloads, nil
}

/*

func (c NDFC) networkAttachmentsUpdate(ctx context.Context, dg *diag.Diagnostics, updateRsc *resource_network_attachments.NDFCNetworkAttachmentsModel, retry int) {
	// PUT for each object
	payload := updateRsc.FillAttachmentsPayloadFromModel()

	if retry > 5 {
		dg.AddError("Update Failed", "Retry count exceeded")
		return
	}
	retryIndices := make([]int, 0)
	rsObj := api.NewNetAttachAPI(updateRsc.FabricName, c.GetLock(ResourceNetworks), &c.apiClient)

	for i := range payload.Attachments {
		data, err := json.Marshal(payload.Attachments[i])
		if err != nil {
			dg.AddError("Marshal Failed", fmt.Sprintf("Resource %s Marshall error %v", payload.Attachments[i].SerialNumber, err))
			return
		}
		tflog.Info(ctx, fmt.Sprintf("Update resource %s", payload.Attachments[i].SerialNumber))
		rsObj.PutSerialNumber = payload.Attachments[i].SerialNumber
		res, err := rsObj.Put(data)
		if err != nil {
			dg.AddError(fmt.Sprintf("Resource %s, Update failed", payload.Attachments[i].SerialNumber), fmt.Sprintf("Error %v, response %s", err, res.Str))
			return
		}
		// Get and check if they Mismatch
		// There is a bug in NDFC where PUT sometimes resets some values to DefaultValue
		// This is mostly seen with dhcpServers payload
		// Re-do ing PUT solves the problem in most cases
		// Hence do a GET nd verify if all params got updated
		// if not add them to a list and re-do PUT until its correct or retry count is exceeded
		rsObj.GetSerialNumber = payload.Attachments[i].SerialNumber
		rs, err := rsObj.Get()
		if err != nil {
			tflog.Error(ctx, "Read resource after PUT Failed")
			dg.AddError(fmt.Sprintf("Resource %s, Get failed", payload.Attachments[i].SerialNumber), fmt.Sprintf("Error %v, response %s", err, res.Str))
			return
		}
		rsNewValue := resource_networkAttachments.NDFCNetworkAttachmentsValue{}
		err = json.Unmarshal(rs, &rsNewValue)
		if err != nil {
			tflog.Error(ctx, "Unmarshal Failed, attachments GET followed by PUT")
			dg.AddError(fmt.Sprintf("Resource %s, Unmarshal failed", payload.Attachments[i].SerialNumber), fmt.Sprintf("Error %v, response %s", err, res.Str))
			return
		}
		if ValuesDeeplyEqual != rsNewValue.DeepEqual(payload.Attachments[i]) {
			tflog.Error(ctx, "Mismatch in data retrieved after PUT - add to retry list")
			retryIndices = append(retryIndices, i)
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Update resource %s Successfull. Message %s", payload.Attachments[i].SerialNumber, res.Str))
	}
	if len(retryIndices) > 0 {
		tflog.Info(ctx, "Retrying network update due to mismatch", map[string]interface{}{"Err": "Mismatch in network data retrieved after PUT"})
		redoRsc := new(resource_network_attachments.NDFCNetworkAttachmentsModel)
		redoRsc.FabricName = updateRsc.FabricName
		redoRsc.Attachments = make(map[string]resource_network_attachments.NDFCNetworkAttachmentsValue)
		for _, i := range retryIndices {
			redoRsc.Attachments[payload.Attachments[i].SerialNumber] = payload.Attachments[i]
		}
		c.networkAttachmentsUpdate(ctx, dg, redoRsc, retry+1)
	}
}
*/

/*
This function compares plan and state entry and returns the action to be taken
on the attachment:
1. New attachment in plan - NewEntry
2. New attachment in plan with deploy set - NewEntry + Deploy
3. Existing attachment in plan with modifications - Update
4. Existing attachment in plan with deployment flag change - Update + Deploy/UnDeploy
5. Unchanged attachment in plan  - NoChange
Attachment level values are set in UpdateAction field, which is used to filter out the entries.
*/
func (c NDFC) checkNwAttachmentsAction(ctx context.Context, plan *resource_networks.NDFCNetworksValue,
	state *resource_networks.NDFCNetworksValue, vpcPairMap map[string]string, nwModified bool, globalDeploy bool) uint16 {

	actionFlag := NoChange

	if plan.DeployAttachments && !state.DeployAttachments {
		tflog.Debug(ctx, fmt.Sprintf("checkNwAttachmentsAction: Net %s, deployment flag changed from false to true", plan.NetworkName))
		actionFlag |= DeployAll
	} else if !plan.DeployAttachments && state.DeployAttachments {
		tflog.Debug(ctx, fmt.Sprintf("checkNwAttachmentsAction: Net %s, deployment flag changed from true to false", plan.NetworkName))
		actionFlag |= UnDeployAll
	} else {
		tflog.Debug(ctx, fmt.Sprintf("checkNwAttachmentsAction: Net %s, deployment flag unchanged", plan.NetworkName))
	}
	if (plan.DeployAttachments || globalDeploy) && nwModified {
		tflog.Debug(ctx, fmt.Sprintf("checkNwAttachmentsAction: Net %s modified and network/global level deploy is true", plan.NetworkName))
		actionFlag |= DeployAll
	}
	for serial, planAttach := range plan.Attachments {
		tflog.Debug(ctx, fmt.Sprintf("checkNwAttachmentsAction: Attachment %s/%s in plan", plan.NetworkName, serial))
		planAttach.SerialNumber = serial
		planAttach.FabricName = plan.FabricName
		planAttach.NetworkName = plan.NetworkName
		planAttach.UpdateAction = NoChange
		controlFlag := NoChange

		//look for attachment in state
		stateAttachment, found := state.Attachments[serial]
		if !found {
			//New attachment in plan
			controlFlag |= NewEntry
			tflog.Debug(ctx, fmt.Sprintf("checkNwAttachmentsAction: New attachment %s/%s in plan",
				plan.NetworkName, serial))
			planAttach.Deployment = "true"
			if planAttach.DeployThisAttachment || plan.DeployAttachments || globalDeploy {
				tflog.Debug(ctx, fmt.Sprintf("checkNwAttachmentsAction: Attachment %s/%s in plan needs deploy",
					plan.NetworkName, serial))
				controlFlag |= Deploy
			}
		} else {
			stateAttachment.FilterThisValue = true
			//Existing attachment in plan
			tflog.Debug(ctx, fmt.Sprintf("checkNwAttachmentsAction: Existing attachment %s/%s in plan",
				plan.NetworkName, serial))
			//Check if parameters are different
			cf := false
			retVal := planAttach.CreatePlan(stateAttachment, &cf)
			log.Printf("checkNwAttachmentsAction: Attachment %s/%s  - CreatePlan returned %d", plan.NetworkName, serial, retVal)
			if retVal == RequiresUpdate {

				//Modified attachment in plan list
				tflog.Debug(ctx, fmt.Sprintf("checkNwAttachmentsAction: attachment %s/%s modified in plan",
					plan.NetworkName, serial))
				controlFlag |= Update
				planAttach.Deployment = "true"

				if stateAttachment.DeployThisAttachment && !planAttach.DeployThisAttachment {
					//undeploy needed
					tflog.Debug(ctx, fmt.Sprintf("checkNwAttachmentsAction: Attachment %s/%s in plan needs un-deploy",
						plan.NetworkName, serial))
					controlFlag |= UnDeploy
				} else if planAttach.DeployThisAttachment || plan.DeployAttachments || globalDeploy {
					//deploy needed
					tflog.Debug(ctx, fmt.Sprintf("checkNwAttachmentsAction: Attachment %s/%s in plan is modified and shall be deployed",
						plan.NetworkName, serial))
					controlFlag |= Deploy
				} else {
					tflog.Debug(ctx, fmt.Sprintf("checkNwAttachmentsAction: Attachment %s/%s is changed but deploy flag is false at all levels",
						plan.NetworkName, serial))
				}

				// Look inside port lists to see if something has changed
				planAttach.SwitchPorts, planAttach.DetachSwitchPorts = portListDiff(planAttach.SwitchPorts, stateAttachment.SwitchPorts)
				//TODO Tor Ports
			} else if nwModified {
				// Net was modified - deployment of attachment needed if flag is set
				if planAttach.DeployThisAttachment || plan.DeployAttachments || globalDeploy {
					tflog.Debug(ctx, fmt.Sprintf("checkNwAttachmentsAction: Net %s modified and attachment %s/%s shall be deployed",
						plan.NetworkName, plan.NetworkName, serial))
					controlFlag |= Deploy
				}
			} else {
				tflog.Debug(ctx, fmt.Sprintf("checkNwAttachmentsAction: Attachment %s/%s in plan is unchanged and no deploy needed",
					plan.NetworkName, serial))
			}
			if cf {
				//Control Flag Update
				tflog.Debug(ctx, fmt.Sprintf("checkNwAttachmentsAction: Deploy flag changed in attachment %s/%s in plan",
					plan.NetworkName, serial))
				if stateAttachment.DeployThisAttachment && !planAttach.DeployThisAttachment {
					//undeploy needed
					tflog.Debug(ctx, fmt.Sprintf("checkNwAttachmentsAction: Attachment %s/%s in plan needs un-deploy",
						plan.NetworkName, serial))
					controlFlag |= UnDeploy
				} else if !stateAttachment.DeployThisAttachment && planAttach.DeployThisAttachment {
					//deploy needed
					tflog.Debug(ctx, fmt.Sprintf("checkNwAttachmentsAction: Attachment %s/%s in plan needs deploy",
						plan.NetworkName, serial))
					controlFlag |= Deploy
				}
			}
			//put modified entry back
			state.Attachments[serial] = stateAttachment
		}
		//put modified entry back
		planAttach.UpdateAction |= controlFlag
		plan.Attachments[serial] = planAttach
		actionFlag |= controlFlag
		if controlFlag&Update > 0 {
			if peerSerial := vpcPairMap[serial]; peerSerial != "" {
				tflog.Debug(ctx, fmt.Sprintf("checkNwAttachmentsAction: vPC peer found attachment %s", peerSerial))
				peerAttachments := plan.Attachments[peerSerial]
				peerAttachments.UpdateAction |= controlFlag
				plan.Attachments[peerSerial] = peerAttachments
			}
		}
	}
	tflog.Debug(ctx, fmt.Sprintf("checkNwAttachmentsAction: Actions %b", actionFlag))
	return actionFlag
}

func (c NDFC) networkAttachmentsGetDiff(ctx context.Context, dg *diag.Diagnostics,
	vPlan *resource_networks.NDFCNetworksModel,
	vState *resource_networks.NDFCNetworksModel, vModified *resource_networks.NDFCNetworksModel) map[string]interface{} {

	actions := make(map[string]interface{})
	naUpdate := new(rna.NDFCNetworkAttachments)
	naDeploy := new(rna.NDFCNetworkAttachments)
	naUnDeploy := new(rna.NDFCNetworkAttachments)

	naUpdate.FabricName = vPlan.FabricName
	naDeploy.FabricName = vPlan.FabricName
	naUnDeploy.FabricName = vPlan.FabricName

	if vState.DeployAllAttachments && !vPlan.DeployAllAttachments {
		//Global undeploy
		tflog.Debug(ctx, "networkAttachmentsGetDiff: Global undeploy needed")
		naUnDeploy.GlobalUndeploy = true
	} else if !vState.DeployAllAttachments && vPlan.DeployAllAttachments {
		//Global deploy
		tflog.Debug(ctx, "networkAttachmentsGetDiff: Global deploy needed")
		naDeploy.GlobalDeploy = true
	} else {
		tflog.Debug(ctx, "networkAttachmentsGetDiff: Global deploy flag unchanged")
		naUnDeploy.GlobalDeploy = false
		naUnDeploy.GlobalUndeploy = false
		naDeploy.GlobalDeploy = false
		naDeploy.GlobalUndeploy = false
	}

	vpcPairMap := c.createVpcPairMap(ctx, vPlan.FabricName)
	for nw, planNw := range vPlan.Networks {
		planNw.NetworkName = nw
		planNw.FabricName = vPlan.FabricName
		netModified := false

		if _, ok := vModified.Networks[nw]; ok {
			tflog.Debug(ctx, fmt.Sprintf("Network %s was modified", nw))
			netModified = true
		} else {
			tflog.Debug(ctx, fmt.Sprintf("Network %s was not modified", nw))
		}

		if stateNw, ok := vState.Networks[nw]; ok {
			tflog.Debug(ctx, fmt.Sprintf("Existing Network %s in plan", nw))
			// Check if there is a change
			action := c.checkNwAttachmentsAction(ctx, &planNw, &stateNw, vpcPairMap, netModified, vPlan.DeployAllAttachments)
			vPlan.Networks[nw] = planNw
			vState.Networks[nw] = stateNw

			if action&Update > 0 {
				naUpdate.AddEntry(nw, planNw.GetAttachmentValues(Update, "true"))
			}
			if action&NewEntry > 0 {
				naUpdate.AddEntry(nw, planNw.GetAttachmentValues(NewEntry, "true"))
			}
			if action&DeployAll > 0 {
				//include all attachments
				naDeploy.AddEntry(nw, planNw.GetAttachmentValues(0, "true"))
			}
			if action&UnDeployAll > 0 {
				naUnDeploy.AddEntry(nw, planNw.GetAttachmentValues(0, "false"))
				// undeploy needs a detach first
				naUpdate.AddEntry(nw, planNw.GetAttachmentValues(0, "false"))
			}
			if action&Deploy > 0 {
				naDeploy.AddEntry(nw, planNw.GetAttachmentValues(Deploy, "true"))
			}
			if action&UnDeploy > 0 {
				naUnDeploy.AddEntry(nw, planNw.GetAttachmentValues(UnDeploy, "false"))
			}
		} else {
			tflog.Debug(ctx, fmt.Sprintf("New Network %s in plan", nw))
			// Add everything for update
			naUpdate.AddEntry(nw, planNw.GetAttachmentValues(0, "true"))
			if planNw.DeployAttachments || vPlan.DeployAllAttachments {
				tflog.Debug(ctx, fmt.Sprintf("New Network %s in plan needs deploy as global/network level deploy is set", nw))
				// Add all attachments for deploy
				naDeploy.AddEntry(nw, planNw.GetAttachmentValues(0, "true"))
			} else {

				for serial, planAttach := range planNw.Attachments {
					// Select the attachments that needs deployment
					if planAttach.DeployThisAttachment {
						planAttach.UpdateAction |= Deploy
						tflog.Debug(ctx, fmt.Sprintf("New Network %s:%s in plan needs deploy as attachment level deploy is set", nw, serial))
						planNw.Attachments[serial] = planAttach
					}
				}
				naDeploy.AddEntry(nw, planNw.GetAttachmentValues(Deploy, "true"))
			}
		}
		vPlan.Networks[nw] = planNw
	}
	for nw, nwEntry := range vState.Networks {
		nwEntry.FabricName = vState.FabricName
		nwEntry.NetworkName = nw
		for serial, attachEntry := range nwEntry.Attachments {
			if attachEntry.FilterThisValue {
				//seen in plan
				continue
			}
			// not seen in plan - Detach
			tflog.Debug(ctx, fmt.Sprintf("networkAttachmentsGetDiff: To be Detached attachment %s/%s",
				nw,
				serial))
			attachEntry.Deployment = "false"
			attachEntry.UpdateAction = NoChange
			attachEntry.UpdateAction |= (Detach | Deploy)
			nwEntry.Attachments[serial] = attachEntry
		}
		naUpdate.AddEntry(nw, nwEntry.GetAttachmentValues(Detach, "false"))
		naDeploy.AddEntry(nw, nwEntry.GetAttachmentValues(Deploy, "false"))
	}
	actions["update"] = naUpdate
	actions["plan"] = vPlan
	actions["state"] = vState
	actions["deploy"] = naDeploy
	actions["undeploy"] = naUnDeploy
	return actions
}

func portListDiff(plan types.CSVString, state types.CSVString) (types.CSVString, types.CSVString) {
	var attachPorts, detachPorts types.CSVString
	statePortMap := make(map[string]bool)

	for i := range state {
		statePortMap[state[i]] = false
	}

	for i := range plan {
		attachPorts = append(attachPorts, plan[i])
		if _, ok := statePortMap[plan[i]]; ok {
			// found in state, mark it
			statePortMap[plan[i]] = true
		}
	}
	for i := range statePortMap {
		if !statePortMap[i] {
			//Not found in plan, so it is a detach port
			detachPorts = append(detachPorts, i)
		}
	}
	return attachPorts, detachPorts
}

func processPortList(attachEntry *rna.NDFCAttachmentsValue) {
	//Fill SwitchPorts  and TorPorts from PortNames
	if attachEntry.PortNames == "" {
		attachEntry.SwitchPorts = nil
		attachEntry.TorPorts = nil
	} else {
		attachEntry.SwitchPorts = append(attachEntry.SwitchPorts, strings.Split(attachEntry.PortNames, ",")...)
	}
}

func processPortListOrder(portList *types.CSVString, order []string) {
	//Make sure portList is in same order as order
	for i := range order {
		for j := range *portList {
			if order[i] == (*portList)[j] {
				// swap
				tflog.Debug(context.Background(), fmt.Sprintf("processPortListOrder: Swapping %s and %s", order[i], (*portList)[j]))
				(*portList)[i], (*portList)[j] = (*portList)[j], (*portList)[i]
				break
			}
		}
	}
}

func (c NDFC) createVpcPairMap(ctx context.Context, fabricName string) map[string]string {

	payload, err := c.GetSwitchesInFabric(ctx, fabricName)
	if err != nil || string(payload) == "[]" || payload == nil {
		tflog.Debug(ctx, "createVpcPairMap: Failed to get switchesByFabric")
		return nil
	}
	switchList := []resource_vpc_pair.NDFCSwitchesByFabric{}
	err = json.Unmarshal(payload, &switchList)
	if err != nil {
		tflog.Error(ctx, "createVpcPairMap: Failed to unmarshal vPC Pair data")
		return nil
	}
	log.Printf("SwitchList %v", switchList)
	var vpcPairMap = make(map[string]string)
	for _, entry := range switchList {
		vpcPairMap[entry.SerialNumber] = entry.PeerSerialNumber
	}
	log.Printf("vpcPairMap %v", vpcPairMap)
	return vpcPairMap
}
