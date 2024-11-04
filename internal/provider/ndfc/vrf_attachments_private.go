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
	"fmt"
	"log"
	"strings"
	rva "terraform-provider-ndfc/internal/provider/resources/resource_vrf_attachments"
	"terraform-provider-ndfc/internal/provider/resources/resource_vrf_bulk"
	. "terraform-provider-ndfc/internal/provider/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
)

const UrlVrfAttachmentsCreate = "/lan-fabric/rest/top-down/v2/fabrics/%s/vrfs/attachments"
const UrlVrfAttachmentsGet = "/lan-fabric/rest/top-down/fabrics/%s/vrfs/attachments"
const UrlQP = "?%s=%s"

func (c NDFC) vrfAttachmentsGet(ctx context.Context, fabricName string, vrfs []string) ([]byte, error) {

	url := fmt.Sprintf(UrlVrfAttachmentsGet, fabricName)
	if len(vrfs) > 0 {
		qp := fmt.Sprintf(UrlQP, "vrf-names", strings.Join(vrfs, ","))
		url += qp
	}
	tflog.Debug(ctx, fmt.Sprintf("vrfAttachmentsGet: url %s", url))
	c.GetLock(ResourceVrfBulk).Lock()
	defer c.GetLock(ResourceVrfBulk).Unlock()
	res, err := c.apiClient.GetRawJson(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c NDFC) processAttachResponse(res gjson.Result) error {
	err := error(nil)
	res.ForEach(func(k, v gjson.Result) bool {
		if !strings.Contains(v.String(), "SUCCESS") && !strings.Contains(v.String(), "already in detached state") {
			err = fmt.Errorf("failed to configure attachments, got error: %s, %s", k.String(), v.String())
		}
		return true
	})
	return err
}

func (c NDFC) vrfAttachmentsPost(ctx context.Context, fabricName string, data []byte) error {

	log.Println("Data to be posted", string(data))
	c.GetLock(ResourceVrfBulk).Lock()
	defer c.GetLock(ResourceVrfBulk).Unlock()
	res, err := c.apiClient.Post(fmt.Sprintf(UrlVrfAttachmentsCreate, fabricName), string(data))
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Error POST:  %v", res))
		err1 := fmt.Errorf("%s:%v", err.Error(), res)
		return err1
	}
	err = c.processAttachResponse(res)
	if err != nil {
		return err
	}
	tflog.Info(ctx, fmt.Sprintf("vrfAttachmentsCreate: Success res : %v", res.Str))
	return nil
}

func (c NDFC) getVrfAttachments(ctx context.Context, dg *diag.Diagnostics,
	fabricName string, vrfs []string) ([]byte, error) {

	tflog.Debug(ctx, fmt.Sprintf("getVrfAttachments: Entering Id %s/{%v}", fabricName, vrfs))
	// Get the VRF Attachments
	res, err := c.vrfAttachmentsGet(ctx, fabricName, vrfs)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("getVrfAttachments: Error getting VRF Attachments %s", err.Error()))
		return nil, err
	}
	tflog.Debug(ctx, fmt.Sprintf("getVrfAttachments: data read from NDFC: %s", string(res)))
	return res, nil
}

// Attachment comparison action
const (
	NoChange    uint16 = 0x0000
	Deploy      uint16 = 0x0001
	UnDeploy    uint16 = 0x0002
	Update      uint16 = 0x0004
	NewEntry    uint16 = 0x0008
	DeployAll   uint16 = 0x0010
	UnDeployAll uint16 = 0x0020
	Detach      uint16 = 0x0040
)

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
func (c NDFC) updateVRFAttachmentAction(ctx context.Context, plan *resource_vrf_bulk.NDFCVrfsValue,
	state *resource_vrf_bulk.NDFCVrfsValue) uint16 {

	actionFlag := NoChange

	if plan.DeployAttachments && !state.DeployAttachments {
		tflog.Debug(ctx, fmt.Sprintf("compareAttachments: VRF %s, deployment flag changed from false to true", plan.VrfName))
		actionFlag |= DeployAll
	} else if !plan.DeployAttachments && state.DeployAttachments {
		tflog.Debug(ctx, fmt.Sprintf("compareAttachments: VRF %s, deployment flag changed from true to false", plan.VrfName))
		actionFlag |= UnDeployAll
	} else {
		tflog.Debug(ctx, fmt.Sprintf("compareAttachments: VRF %s, deployment flag unchanged", plan.VrfName))
	}

	for serial, planAttach := range plan.AttachList {
		tflog.Debug(ctx, fmt.Sprintf("compareAttachments: Attachment %s/%s in plan", plan.VrfName, serial))
		planAttach.SerialNumber = serial
		planAttach.FabricName = plan.FabricName
		planAttach.VrfName = plan.VrfName
		planAttach.UpdateAction = NoChange
		controlFlag := NoChange

		//look for attachment in state
		stateAttachment, found := state.AttachList[serial]
		if !found {
			//New attachment in plan
			controlFlag |= NewEntry
			tflog.Debug(ctx, fmt.Sprintf("compareAttachments: New attachment %s/%s in plan",
				plan.VrfName, serial))
			planAttach.SerialNumber = serial
			planAttach.Deployment = "true"
			if planAttach.DeployThisAttachment || plan.DeployAttachments {
				tflog.Debug(ctx, fmt.Sprintf("compareAttachments: Attachment %s/%s in plan needs deploy",
					plan.VrfName, serial))
				controlFlag |= Deploy
			}
		} else {
			stateAttachment.FilterThisValue = true
			//Existing attachment in plan
			tflog.Debug(ctx, fmt.Sprintf("compareAttachments: Existing attachment %s/%s in plan",
				plan.VrfName, serial))
			//Check if parameters are different
			retVal := planAttach.DeepEqual(stateAttachment)
			log.Printf("compareAttachments: Attachment %s/%s  - DeepEqual %d", plan.VrfName, serial, retVal)
			if retVal == ValuesDeeplyEqual {
				tflog.Debug(ctx, fmt.Sprintf("compareAttachments: attachment %s/%s - unchanged",
					plan.VrfName, serial))

			} else if retVal == ControlFlagUpdate {
				//Control Flag Update
				tflog.Debug(ctx, fmt.Sprintf("compareAttachments: Deploy flag changed in attachment %s/%s in plan",
					plan.VrfName, serial))
				if stateAttachment.DeployThisAttachment && !planAttach.DeployThisAttachment {
					//undeploy needed
					tflog.Debug(ctx, fmt.Sprintf("compareAttachments: Attachment %s/%s in plan needs un-deploy",
						plan.VrfName, serial))
					controlFlag |= UnDeploy
				} else if !stateAttachment.DeployThisAttachment && planAttach.DeployThisAttachment {
					//deploy needed
					tflog.Debug(ctx, fmt.Sprintf("compareAttachments: Attachment %s/%s in plan needs deploy",
						plan.VrfName, serial))
					controlFlag |= Deploy
				}
			} else {
				//Modified attachment in plan list
				tflog.Debug(ctx, fmt.Sprintf("compareAttachments: Modified attachment %s/%s in plan",
					plan.VrfName, serial))
				controlFlag |= Update
				planAttach.Deployment = "true"
				if stateAttachment.DeployThisAttachment && !planAttach.DeployThisAttachment {
					//undeploy needed
					tflog.Debug(ctx, fmt.Sprintf("compareAttachments: Attachment %s/%s in plan needs un-deploy",
						plan.VrfName, serial))
					controlFlag |= UnDeploy
				} else if planAttach.DeployThisAttachment || plan.DeployAttachments {
					//deploy needed
					tflog.Debug(ctx, fmt.Sprintf("compareAttachments: Attachment %s/%s in plan needs deploy",
						plan.VrfName, serial))
					controlFlag |= Deploy
				}
			}
			//put modified entry back
			state.AttachList[serial] = stateAttachment
		}
		//put modified entry back
		planAttach.UpdateAction |= controlFlag
		plan.AttachList[serial] = planAttach
		actionFlag |= controlFlag
	}
	tflog.Debug(ctx, fmt.Sprintf("compareAttachments: Actions %b", actionFlag))
	return actionFlag
}

func (c NDFC) diffVrfAttachments(ctx context.Context, planData *resource_vrf_bulk.NDFCVrfBulkModel,
	stateData *resource_vrf_bulk.NDFCVrfBulkModel) map[string]*rva.NDFCVrfAttachmentsPayloads {

	action := make(map[string]*rva.NDFCVrfAttachmentsPayloads)

	tflog.Debug(ctx, "diffVrfAttachments: Entering")
	//ID, _ := c.VrfAttachmentsCreateID(planData)
	vaUpdate := new(resource_vrf_bulk.NDFCVrfBulkModel)
	vaUpdate.Vrfs = make(map[string]resource_vrf_bulk.NDFCVrfsValue)
	vaUpdate.FabricName = planData.FabricName

	vaUpdatePayload := new(rva.NDFCVrfAttachmentsPayloads)
	vaUpdatePayload.FabricName = planData.FabricName

	vaDeployPayload := new(rva.NDFCVrfAttachmentsPayloads)
	vaDeployPayload.FabricName = planData.FabricName

	vaUnDeployPayload := new(rva.NDFCVrfAttachmentsPayloads)

	vaUnDeployPayload.FabricName = planData.FabricName

	if stateData.DeployAllAttachments && !planData.DeployAllAttachments {
		//Global undeploy
		tflog.Debug(ctx, "diffVrfAttachments: Global undeploy needed")
		vaUnDeployPayload.GlobalUndeploy = true
	} else if !stateData.DeployAllAttachments && planData.DeployAllAttachments {
		//Global deploy
		tflog.Debug(ctx, "diffVrfAttachments: Global deploy needed")
		vaDeployPayload.GlobalDeploy = true
	} else {
		tflog.Debug(ctx, "diffVrfAttachments: Global deploy flag unchanged")
		vaUnDeployPayload.GlobalDeploy = false
		vaUnDeployPayload.GlobalUndeploy = false
		vaDeployPayload.GlobalDeploy = false
		vaDeployPayload.GlobalUndeploy = false
	}
	for vrf, planVrf := range planData.Vrfs {
		planVrf.FabricName = planData.FabricName
		planVrf.VrfName = vrf
		//Look for vrf in state
		if sVrf, ok := stateData.Vrfs[vrf]; ok {
			tflog.Debug(ctx, fmt.Sprintf("diffVrfAttachments: Existing VRF %s in plan", vrf))
			//found - now compare attachments
			//Check if there are new attachments
			action := c.updateVRFAttachmentAction(ctx, &planVrf, &sVrf)
			planData.Vrfs[vrf] = planVrf
			stateData.Vrfs[vrf] = sVrf

			// DeployAll set in the VRF, put all attachments in VRF for deploy
			if action&DeployAll > 0 {
				//Get all - no filtering
				vaDeployPayload.AddEntry(vrf, planVrf.GetAttachmentValues(0, ""))
			}

			// UndeployAll set in the VRF, put all attachments in VRF for undeploy
			if action&UnDeployAll > 0 {
				//Get all - no filtering
				vaUnDeployPayload.AddEntry(vrf, planVrf.GetAttachmentValues(0, ""))
				vaUpdatePayload.AddEntry(vrf, planVrf.GetAttachmentValues(0, "false"))
			}

			// Deploy is set for some attachments in the VRF, put those attachments for deploy
			if action&Deploy > 0 {
				vaDeployPayload.AddEntry(vrf, planVrf.GetAttachmentValues(Deploy, ""))
			}

			if action&UnDeploy > 0 {
				vaUnDeployPayload.AddEntry(vrf, planVrf.GetAttachmentValues(UnDeploy, "true"))
				//Undeployments need a detach first - put the attachment in update list to detach first
				vaUpdatePayload.AddEntry(vrf, planVrf.GetAttachmentValues(UnDeploy, "false"))
			}

			if action&Update > 0 {
				vaUpdatePayload.AddEntry(vrf, planVrf.GetAttachmentValues(Update, "true"))
			}

			if action&NewEntry > 0 {
				vaUpdatePayload.AddEntry(vrf, planVrf.GetAttachmentValues(NewEntry, "true"))
			}

		} else {
			//New VRF entry in plan
			tflog.Debug(ctx, fmt.Sprintf("diffVrfAttachments: New VRF %s in plan, not in state - must be new ", vrf))
			//New VRF entry in plan - Get all attachments without any filtering
			vaUpdatePayload.AddEntry(vrf, planVrf.GetAttachmentValues(0, "true"))
			if planVrf.DeployAttachments {
				vaDeployPayload.AddEntry(vrf, planVrf.GetAttachmentValues(0, "true"))
			} else {
				//Check each entry for deploy flag and mark the bitmask
				for serial, attachEntry := range planVrf.AttachList {
					if attachEntry.DeployThisAttachment {
						attachEntry.UpdateAction |= Deploy
						planVrf.AttachList[serial] = attachEntry
					}
				}
				vaDeployPayload.AddEntry(vrf, planVrf.GetAttachmentValues(Deploy, "true"))
			}

		}
		planData.Vrfs[vrf] = planVrf
	}

	for vrf, vrfEntry := range stateData.Vrfs {
		vrfEntry.FabricName = stateData.FabricName
		vrfEntry.VrfName = vrf

		for serial, attachEntry := range vrfEntry.AttachList {
			if attachEntry.FilterThisValue {
				//seen in plan
				continue
			}
			//attachment not seen in plan - needs to be detached
			tflog.Debug(ctx, fmt.Sprintf("diffVrfAttachments: To be Detached attachment %s/%s",
				vrf,
				serial))
			attachEntry.UpdateAction |= (Detach | Deploy)
			vrfEntry.AttachList[serial] = attachEntry
		}
		vaUpdatePayload.AddEntry(vrf, vrfEntry.GetAttachmentValues(Detach, "false"))
		vaDeployPayload.AddEntry(vrf, vrfEntry.GetAttachmentValues(Detach, "false"))
	}

	action["update"] = vaUpdatePayload
	action["deploy"] = vaDeployPayload
	action["undeploy"] = vaUnDeployPayload
	return action
}
