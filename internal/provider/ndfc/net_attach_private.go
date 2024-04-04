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

func (c NDFC) netAttachmentsDetach(ctx context.Context, rscModel *resource_networks.NDFCNetworksModel) error {
	tflog.Info(ctx, fmt.Sprintf("Beginning Bulk NetworkAttachments delete in fabric %s", rscModel.FabricName))
	payload := rna.NDFCNetworkAttachments{}
	rscModel.FillAttachmentsPayloadFromModel(&payload, resource_networks.NwAttachmentDetach)

	data, err := json.Marshal(payload.NetworkAttachments)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Json Marshal failure %s", err.Error()))
		return err
	}
	return c.netAttachmentsPostPayload(ctx, rscModel.FabricName, data)
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
	state *resource_networks.NDFCNetworksValue) uint16 {

	actionFlag := NoChange
	/*
		if plan.DeployAttachments && !state.DeployAttachments {
			tflog.Debug(ctx, fmt.Sprintf("compareAttachments: VRF %s, deployment flag changed from false to true", plan.VrfName))
			actionFlag |= DeployAll
		} else if !plan.DeployAttachments && state.DeployAttachments {
			tflog.Debug(ctx, fmt.Sprintf("compareAttachments: VRF %s, deployment flag changed from true to false", plan.VrfName))
			actionFlag |= UnDeployAll
		} else {
			tflog.Debug(ctx, fmt.Sprintf("compareAttachments: VRF %s, deployment flag unchanged", plan.VrfName))
		}
	*/
	for serial, planAttach := range plan.Attachments {
		tflog.Debug(ctx, fmt.Sprintf("compareAttachments: Attachment %s/%s in plan", plan.NetworkName, serial))
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
			tflog.Debug(ctx, fmt.Sprintf("compareAttachments: New attachment %s/%s in plan",
				plan.NetworkName, serial))
			planAttach.Deployment = "true"
			/* TBD Deployment implementation
			if planAttach.DeployThisAttachment {
				tflog.Debug(ctx, fmt.Sprintf("compareAttachments: Attachment %s/%s in plan needs deploy",
					plan.NetworkName, serial))
				controlFlag |= Deploy
			}
			*/
		} else {
			stateAttachment.FilterThisValue = true
			//Existing attachment in plan
			tflog.Debug(ctx, fmt.Sprintf("compareAttachments: Existing attachment %s/%s in plan",
				plan.NetworkName, serial))
			//Check if parameters are different
			retVal := planAttach.CreatePlan(stateAttachment)
			log.Printf("compareAttachments: Attachment %s/%s  - CreatePlan returned %d", plan.NetworkName, serial, retVal)
			if retVal == ActionNone {
				tflog.Debug(ctx, fmt.Sprintf("compareAttachments: attachment %s/%s - unchanged",
					plan.NetworkName, serial))

			} else if retVal == ControlFlagUpdate {
				//Control Flag Update
				/* TBD Deployment implementation
				tflog.Debug(ctx, fmt.Sprintf("compareAttachments: Deploy flag changed in attachment %s/%s in plan",
					plan.NetworkName, serial))
				if stateAttachment.DeployThisAttachment && !planAttach.DeployThisAttachment {
					//undeploy needed
					tflog.Debug(ctx, fmt.Sprintf("compareAttachments: Attachment %s/%s in plan needs un-deploy",
						plan.NetworkName, serial))
					controlFlag |= UnDeploy
				} else if !stateAttachment.DeployThisAttachment && planAttach.DeployThisAttachment {
					//deploy needed
					tflog.Debug(ctx, fmt.Sprintf("compareAttachments: Attachment %s/%s in plan needs deploy",
						plan.NetworkName, serial))
					controlFlag |= Deploy
				}
				*/
			} else {

				//Modified attachment in plan list
				tflog.Debug(ctx, fmt.Sprintf("compareAttachments: attachment %s/%s modified in plan",
					plan.NetworkName, serial))
				controlFlag |= Update
				planAttach.Deployment = "true"
				/*
					if stateAttachment.DeployThisAttachment && !planAttach.DeployThisAttachment {
						//undeploy needed
						tflog.Debug(ctx, fmt.Sprintf("compareAttachments: Attachment %s/%s in plan needs un-deploy",
							plan.NetworkName, serial))
						controlFlag |= UnDeploy
					} else if !stateAttachment.DeployThisAttachment && planAttach.DeployThisAttachment {
						//deploy needed
						tflog.Debug(ctx, fmt.Sprintf("compareAttachments: Attachment %s/%s in plan needs deploy",
							plan.NetworkName, serial))
						controlFlag |= Deploy
					}
				*/
				// Look inside port lists to see if something has changed
				planAttach.SwitchPorts, planAttach.DetachSwitchPorts = portListDiff(planAttach.SwitchPorts, stateAttachment.SwitchPorts)
				//TODO Tor Ports
			}
			//put modified entry back
			state.Attachments[serial] = stateAttachment
		}
		//put modified entry back
		planAttach.UpdateAction |= controlFlag
		plan.Attachments[serial] = planAttach
		actionFlag |= controlFlag
	}
	tflog.Debug(ctx, fmt.Sprintf("compareAttachments: Actions %b", actionFlag))
	return actionFlag
}

func (c NDFC) networkAttachmentsGetDiff(ctx context.Context, dg *diag.Diagnostics,
	vPlan *resource_networks.NDFCNetworksModel,
	vState *resource_networks.NDFCNetworksModel) map[string]interface{} {

	actions := make(map[string]interface{})
	naUpdate := new(rna.NDFCNetworkAttachments)
	naUpdate.FabricName = vPlan.FabricName

	for nw, planNw := range vPlan.Networks {
		planNw.NetworkName = nw
		planNw.FabricName = vPlan.FabricName

		if stateNw, ok := vState.Networks[nw]; ok {
			tflog.Debug(ctx, fmt.Sprintf("Existing Network %s in plan", nw))
			// Check if there is a change
			action := c.checkNwAttachmentsAction(ctx, &planNw, &stateNw)
			vPlan.Networks[nw] = planNw
			vState.Networks[nw] = stateNw

			if action&Update > 0 {
				naUpdate.AddEntry(nw, planNw.GetAttachmentValues(Update, "true"))
			}

			if action&NewEntry > 0 {
				naUpdate.AddEntry(nw, planNw.GetAttachmentValues(NewEntry, "true"))
			}
			// TODO: Implement Deploy/UnDeploy

		} else {
			tflog.Debug(ctx, fmt.Sprintf("New Network %s in plan", nw))
			naUpdate.AddEntry(nw, planNw.GetAttachmentValues(NewEntry, "true"))
			// TODO: Implement Deploy/UnDeploy
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
			attachEntry.UpdateAction |= Detach
			nwEntry.Attachments[serial] = attachEntry
		}
		naUpdate.AddEntry(nw, nwEntry.GetAttachmentValues(Detach, "false"))
	}
	actions["update"] = naUpdate
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
	attachEntry.SwitchPorts = append(attachEntry.SwitchPorts, strings.Split(attachEntry.PortNames, ",")...)

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
