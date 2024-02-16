package ndfc

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	rva "terraform-provider-ndfc/internal/provider/resources/resource_vrf_attachments"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func (c NDFC) vrfAttachmentsDeploy(ctx context.Context, dg *diag.Diagnostics, va *rva.NDFCVrfAttachmentsModel) error {
	//mapkey: serial entry list of VRFs
	deploy_map := make(map[string][]string)
	deployVrfList := make([]string, 0)
	deployEach := false
	deployVrf := false
	expected_state := make(map[string]string)

	updateExpectedState := func(out *map[string]string, vrfName string, serial string, state string) {
		key := vrfName + "/" + serial
		if state == "false" {
			(*out)[key] = NDFCStateNA
		} else {
			(*out)[key] = NDFCStateDeployed
		}
	}

	log.Printf("============================Starting Deployment===================================")

	if va.DeployAllAttachments {
		deployVrf = true
		//Deploy all attachments
		for i := range va.VrfAttachments {
			deployVrfList = append(deployVrfList, va.VrfAttachments[i].VrfName)
			for j := range va.VrfAttachments[i].AttachList {
				updateExpectedState(&expected_state, va.VrfAttachments[i].VrfName,
					va.VrfAttachments[i].AttachList[j].SerialNumber, va.VrfAttachments[i].AttachList[j].Deployment)
			}
		}
	} else {
		// Deploy VRFs in the resource if flag is set
		for i := range va.VrfAttachments {
			if va.VrfAttachments[i].DeployAllAttachments {
				log.Printf("vrfAttachmentsDeploy: Deploying all attachments in VRF %s", va.VrfAttachments[i].VrfName)
				deployVrf = true
				deployVrfList = append(deployVrfList, va.VrfAttachments[i].VrfName)
				for j := range va.VrfAttachments[i].AttachList {
					updateExpectedState(&expected_state, va.VrfAttachments[i].VrfName,
						va.VrfAttachments[i].AttachList[j].SerialNumber, va.VrfAttachments[i].AttachList[j].Deployment)
				}
			} else {
				log.Printf("vrfAttachmentsDeploy: Deploying specific attachments in VRF %s", va.VrfAttachments[i].VrfName)
				//Deploy specific attachments in the VRF
				for j := range va.VrfAttachments[i].AttachList {
					if va.VrfAttachments[i].AttachList[j].DeployThisAttachment ||
						va.VrfAttachments[i].AttachList[j].Deployment == "false" {
						tflog.Debug(ctx, "Deployment required",
							map[string]interface{}{"vrf": va.VrfAttachments[i].VrfName,
								"attachment": va.VrfAttachments[i].AttachList[j].SerialNumber})
						deployEach = true
						deploy_map[va.VrfAttachments[i].AttachList[j].SerialNumber] =
							append(deploy_map[va.VrfAttachments[i].AttachList[j].SerialNumber], va.VrfAttachments[i].VrfName)
						updateExpectedState(&expected_state, va.VrfAttachments[i].VrfName, va.VrfAttachments[i].AttachList[j].SerialNumber,
							va.VrfAttachments[i].AttachList[j].Deployment)

					}
				}

			}
		}
	}

	if !deployEach && !deployVrf {
		tflog.Info(ctx, "vrfAttachmentsDeploy: No attachments to deploy")
		//dg.AddWarning("Deployment not done", "No attachments to deploy")
		return nil
	}
	tflog.Debug(ctx, fmt.Sprintf("vrfAttachmentsDeploy: Starting FSM %v", deploy_map))
	deployFsm := c.createFSM(ctx, dg, va.FabricName, deploy_map, deployVrfList, expected_state)
	deployFsm.Run(ctx)
	tflog.Debug(ctx, fmt.Sprintf("vrfAttachmentsDeploy: Finishing FSM %v", deploy_map))

	return nil
}

func (c NDFC) deploy(ctx context.Context, dg *diag.Diagnostics, fabricName string, deployVrfList *[]string, deploy_map *map[string][]string) error {

	tflog.Debug(ctx, "deploy: Entering", map[string]interface{}{"deployVrfList": *deployVrfList, "deploy_map": deploy_map})
	if len(*deploy_map) != 0 {
		deploy_post_payload := c.deployPayloadBuilder(deploy_map)
		tflog.Info(ctx, fmt.Sprintf("vrfAttachmentsDeploy: Deploying Attachments %s", deploy_post_payload))
		c.GetLock(ResourceVrfAttachments).Lock()
		defer c.GetLock(ResourceVrfAttachments).Unlock()
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
		c.GetLock(ResourceVrfAttachments).Lock()
		defer c.GetLock(ResourceVrfAttachments).Unlock()
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
	deployVrfList []string, deploy_map map[string][]string, expected_state map[string]string) string {
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
		return ""
	}
	tflog.Debug(ctx, fmt.Sprintf("Status Checker:  res : %v", string(res)))

	ndVrfs := rva.NDFCVrfAttachmentsModel{}
	err = json.Unmarshal(res, &ndVrfs.VrfAttachments)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("getVrfAttachments: Error unmarshalling VRF Attachments %s", err.Error()))
		log.Printf("getVrfAttachments: Error unmarshalling VRF Attachments %s", string(res))
		return ""
	}

	for i := range ndVrfs.VrfAttachments {
		for j := range ndVrfs.VrfAttachments[i].AttachList {
			expected, found := expected_state[ndVrfs.VrfAttachments[i].VrfName+"/"+ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo]
			if !found {
				continue
			}

			tflog.Debug(ctx, fmt.Sprintf("vrfAttachmentsDeploy: VRF %s Attachment %s State %s expected %s",
				ndVrfs.VrfAttachments[i].VrfName,
				ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo,
				ndVrfs.VrfAttachments[i].AttachList[j].AttachState, expected))
			if ndVrfs.VrfAttachments[i].AttachList[j].AttachState != expected {
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
					ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo, ndVrfs.VrfAttachments[i].AttachList[j].AttachState, expected))
			}
		}
	}
	tflog.Info(ctx, "vrfAttachmentsDeploy: All VRFs deployed successfully")
	return EventComplete
	//Check the state of the VRFs

}
