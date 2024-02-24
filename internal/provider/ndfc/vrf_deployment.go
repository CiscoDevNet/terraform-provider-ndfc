package ndfc

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	rva "terraform-provider-ndfc/internal/provider/resources/resource_vrf_attachments"
	"terraform-provider-ndfc/internal/provider/resources/resource_vrf_bulk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type DeploymentState struct {
	State string
	Seen  bool
}

type NDFCVrfDeployment struct {
	FabricName    string
	DeployVrfList []string
	DeployMap     map[string][]string
	ExpectedState map[string]DeploymentState
}

func updateExpectedState(out *map[string]DeploymentState, vrfName string, serial string, state string) {
	key := vrfName + "/" + serial
	depState := DeploymentState{}
	if state == "false" {
		depState.State = NDFCStateNA

	} else {
		depState.State = NDFCStateDeployed
	}
	depState.Seen = false
	(*out)[key] = depState
}

// Called from Create/Delete Flows
func (c NDFC) RscDeployAttachments(ctx context.Context, dg *diag.Diagnostics, va *resource_vrf_bulk.NDFCVrfBulkModel) {
	//mapkey: serial entry list of VRFs
	deploy_map := make(map[string][]string)
	deployVrfList := make([]string, 0)
	deployAllVrf := false
	expected_state := make(map[string]DeploymentState)

	if va.DeployAllAttachments {
		tflog.Info(ctx, "RscDeployAttachments: Deploying all attachments")
		deployAllVrf = true
	}
	//Deploy all attachments
	for vrfName, vrfEntry := range va.Vrfs {
		for serial, attachEntry := range vrfEntry.AttachList {
			if attachEntry.Deployment == "false" {
				tflog.Info(ctx, fmt.Sprintf("RscDeployAttachments: Deploying Attachment %s/%s due to detach", vrfName, serial))
			}
			if attachEntry.Deployment == "false" || deployAllVrf ||
				vrfEntry.DeployAttachments || attachEntry.DeployThisAttachment {
				tflog.Info(ctx, fmt.Sprintf("RscDeployAttachments: Deploying Attachment %s/%s", vrfName, serial))
				updateExpectedState(&expected_state, vrfName,
					serial, attachEntry.Deployment)
				deploy_map[serial] = append(deploy_map[serial], vrfName)

			}
		}
	}

	if len(deploy_map) == 0 && len(deployVrfList) == 0 {
		tflog.Info(ctx, "RscDeployAttachments: No attachments to deploy")
		//dg.AddWarning("Deployment not done", "No attachments to deploy")
		return
	}
	c.DeployFSM(ctx, dg, va.FabricName, deploy_map, deployVrfList, expected_state)
	return
}

// Called from Update
func (c NDFC) DeployFromPayload(ctx context.Context, dg *diag.Diagnostics, payload *rva.NDFCVrfAttachmentsPayloads) {
	deployment := NDFCVrfDeployment{}
	deployment.FabricName = payload.FabricName
	deployment.DeployMap = make(map[string][]string)
	deployment.DeployVrfList = make([]string, 0)
	deployment.ExpectedState = make(map[string]DeploymentState)

	for _, vrfEntry := range payload.VrfAttachments {
		for _, attachEntry := range vrfEntry.AttachList {
			deployment.DeployMap[attachEntry.SerialNumber] = append(deployment.DeployMap[attachEntry.SerialNumber], vrfEntry.VrfName)
			updateExpectedState(&deployment.ExpectedState, vrfEntry.VrfName, attachEntry.SerialNumber, attachEntry.Deployment)
		}
	}

	if len(deployment.DeployMap) == 0 && len(deployment.DeployVrfList) == 0 {
		tflog.Info(ctx, "vrfAttachmentsDeploy: No attachments to deploy")
		//dg.AddWarning("Deployment not done", "No attachments to deploy")
		return
	}
	c.DeployFSM(ctx, dg, payload.FabricName, deployment.DeployMap, deployment.DeployVrfList, deployment.ExpectedState)
	return
}

func (c NDFC) DeployFSM(ctx context.Context, dg *diag.Diagnostics, fabricName string,
	deploy_map map[string][]string, deployVrfList []string, expected_state map[string]DeploymentState) {

	log.Printf("============================Starting Deployment===================================")

	tflog.Debug(ctx, fmt.Sprintf("vrfAttachmentsDeploy: Starting FSM %v", deploy_map))
	deployFsm := c.createFSM(ctx, dg, fabricName, deploy_map, deployVrfList, expected_state)
	deployFsm.Run(ctx)
	tflog.Debug(ctx, fmt.Sprintf("vrfAttachmentsDeploy: Finishing FSM %v", deploy_map))

	return
}

func (c NDFC) deploy(ctx context.Context, dg *diag.Diagnostics, fabricName string, deployVrfList *[]string, deploy_map *map[string][]string) error {

	tflog.Debug(ctx, "deploy: Entering", map[string]interface{}{"deployVrfList": *deployVrfList, "deploy_map": deploy_map})
	if len(*deploy_map) != 0 {
		deploy_post_payload := c.deployPayloadBuilder(deploy_map)
		tflog.Info(ctx, fmt.Sprintf("vrfAttachmentsDeploy: Deploying Attachments %s", deploy_post_payload))
		c.GetLock(ResourceVrfBulk).Lock()
		defer c.GetLock(ResourceVrfBulk).Unlock()
		res, err := c.apiClient.Post(UrlVrfAttachmentsDeploy, deploy_post_payload)
		if err != nil {
			dg.AddError("Error in deploying attachments", err.Error())
			return err
		}
		tflog.Info(ctx, fmt.Sprintf("vrfAttachmentsDeploy: Success res : %v", res.Str))

	}
	if len(*deployVrfList) != 0 {
		payload := map[string]string{"vrfNames": strings.Join(*deployVrfList, ",")}
		data, err := json.Marshal(payload)
		if err != nil {
			dg.AddError("Error in deploying attachments", err.Error())
			return err
		}
		tflog.Info(ctx, fmt.Sprintf("vrfAttachmentsDeploy: Deploying VRFs %s", string(data)))
		c.GetLock(ResourceVrfBulk).Lock()
		defer c.GetLock(ResourceVrfBulk).Unlock()
		res, err := c.apiClient.Post(fmt.Sprintf(UrlVrfDeployment, fabricName), string(data))
		if err != nil {
			dg.AddError("Error in deploying attachments", err.Error())
			return err
		}
		tflog.Info(ctx, fmt.Sprintf("vrfAttachmentsDeploy: Success res : %v", res.Str))
	}
	return nil

}

func (c NDFC) deployPayloadBuilder(deploy_map *map[string][]string) string {
	deploy_post_payload := "{"
	for k, v := range *deploy_map {
		deploy_post_payload += "\"" + k + "\":\""
		deploy_post_payload += strings.Join(v, ",")
		deploy_post_payload += "\","
	}
	deploy_post_payload += "}"
	deploy_post_payload = strings.ReplaceAll(deploy_post_payload, ",}", "}")
	return deploy_post_payload
}

func (c NDFC) stateChecker(ctx context.Context, dg *diag.Diagnostics, fabricName string,
	deployVrfList []string, deploy_map map[string][]string, expected_state map[string]DeploymentState) string {
	//Check the state of the VRFs
	vrfs := make([]string, 0)
	for _, v := range deploy_map {
		vrfs = append(vrfs, v...)
	}
	vrfs = append(vrfs, deployVrfList...)
	//Get the state of the VRFs
	res, err := c.vrfAttachmentsGet(ctx, fabricName, vrfs)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("getVrfAttachments: Error getting VRF Attachments %s", err.Error()))
		return EventFailed
	}
	tflog.Debug(ctx, fmt.Sprintf("Status Checker:  res : %v", string(res)))

	ndVrfs := rva.NDFCVrfAttachmentsPayloads{}
	err = json.Unmarshal(res, &ndVrfs.VrfAttachments)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("getVrfAttachments: Error unmarshalling VRF Attachments %s", err.Error()))
		log.Printf("getVrfAttachments: Error unmarshalling VRF Attachments %s", string(res))
		return EventFailed
	}

	for i := range ndVrfs.VrfAttachments {
		for j := range ndVrfs.VrfAttachments[i].AttachList {
			key := ndVrfs.VrfAttachments[i].VrfName + "/" + ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo
			expected, found := expected_state[key]
			if !found {
				continue
			}
			expected.Seen = true
			expected_state[key] = expected

			tflog.Debug(ctx, fmt.Sprintf("vrfAttachmentsDeploy: VRF %s Attachment %s State %s expected %s",
				ndVrfs.VrfAttachments[i].VrfName,
				ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo,
				ndVrfs.VrfAttachments[i].AttachList[j].AttachState, expected.State))
			if ndVrfs.VrfAttachments[i].AttachList[j].AttachState != expected.State {
				switch ndVrfs.VrfAttachments[i].AttachList[j].AttachState {
				case NDFCStateFailed:
					tflog.Error(ctx, fmt.Sprintf("vrfAttachmentsDeploy: VRF %s/%s deployment failed",
						ndVrfs.VrfAttachments[i].VrfName, ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo))
					dg.AddError("VRF Deployment Failed", fmt.Sprintf("VRF %s/%s deployment failed",
						ndVrfs.VrfAttachments[i].VrfName, ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo))
					return EventFailed
				case NDFCStateDeployed:
					fallthrough
				case NDFCStateInPro:
					tflog.Info(ctx, fmt.Sprintf("vrfAttachmentsDeploy: VRF %s/%s is in progress",
						ndVrfs.VrfAttachments[i].VrfName, ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo))
					return EventWait
				case NDFCStatePending:
					tflog.Info(ctx, fmt.Sprintf("vrfAttachmentsDeploy: VRF %s/%s is pending",
						ndVrfs.VrfAttachments[i].VrfName, ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo))
					fallthrough
				case NDFCStateOutOfSync:
					return EventRetry
				}
			} else {
				//TODO optimize by removing these from check - after the tolerance level has met
				tflog.Info(ctx, fmt.Sprintf("vrfAttachmentsDeploy:  %s/%s has reached expected state %s:%s", ndVrfs.VrfAttachments[i].VrfName,
					ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo, ndVrfs.VrfAttachments[i].AttachList[j].AttachState, expected.State))
			}
		}
	}
	for k, v := range expected_state {
		if !v.Seen {
			tflog.Error(ctx, fmt.Sprintf("vrfAttachmentsDeploy:  %s has not been seen in attachments", k))
			dg.AddError("Attachment missing", fmt.Sprintf("VRF:Attachment %s not visible in NDFC", k))
			return EventFailed
		}
	}

	tflog.Info(ctx, "vrfAttachmentsDeploy: All VRFs deployed successfully")
	return EventComplete
	//Check the state of the VRFs

}
