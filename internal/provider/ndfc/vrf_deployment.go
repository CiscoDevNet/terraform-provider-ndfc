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
	RetryList     map[string][]string
	retryFlag     bool
}

func NewNDFCVrfDeployment() *NDFCVrfDeployment {
	return &NDFCVrfDeployment{
		DeployMap:     make(map[string][]string),
		DeployVrfList: make([]string, 0),
		ExpectedState: make(map[string]DeploymentState),
		RetryList:     make(map[string][]string),
	}
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
	d := NewNDFCVrfDeployment()
	d.FabricName = va.FabricName
	deployAllVrf := false

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
				updateExpectedState(&d.ExpectedState, vrfName,
					serial, attachEntry.Deployment)
				d.DeployMap[serial] = append(d.DeployMap[serial], vrfName)
			}
		}
	}
	if len(d.DeployMap) == 0 {
		tflog.Info(ctx, "RscDeployAttachments: No attachments to deploy")
		//dg.AddWarning("Deployment not done", "No attachments to deploy")
		return
	}
	c.DeployFSM(ctx, dg, d)
}

// Called from Update
func (c NDFC) DeployFromPayload(ctx context.Context, dg *diag.Diagnostics, payload *rva.NDFCVrfAttachmentsPayloads) {
	deployment := NewNDFCVrfDeployment()
	deployment.FabricName = payload.FabricName

	for _, vrfEntry := range payload.VrfAttachments {
		for _, attachEntry := range vrfEntry.AttachList {
			deployment.DeployMap[attachEntry.SerialNumber] = append(deployment.DeployMap[attachEntry.SerialNumber], vrfEntry.VrfName)
			updateExpectedState(&deployment.ExpectedState, vrfEntry.VrfName, attachEntry.SerialNumber, attachEntry.Deployment)
		}
	}

	if len(deployment.DeployMap) == 0 {
		tflog.Info(ctx, "vrfAttachmentsDeploy: No attachments to deploy")
		//dg.AddWarning("Deployment not done", "No attachments to deploy")
		return
	}
	c.DeployFSM(ctx, dg, deployment)
}

func (c NDFC) DeployFSM(ctx context.Context, dg *diag.Diagnostics, deployment *NDFCVrfDeployment) {

	log.Printf("============================Starting Deployment===================================")
	tflog.Debug(ctx, fmt.Sprintf("vrfAttachmentsDeploy: Starting FSM %v", deployment.DeployMap))
	deployFsm := c.CreateFSM(ctx, dg, deployment)
	deployFsm.Run(ctx)
	tflog.Debug(ctx, fmt.Sprintf("vrfAttachmentsDeploy: Finishing FSM %v", deployment.DeployMap))
}

// TODO See if this needs chunking
func (c NDFC) DeployBulk(ctx context.Context, dg *diag.Diagnostics, deployment *NDFCVrfDeployment) error {
	var deployMap *map[string][]string
	tflog.Debug(ctx, "deploy: Entering", map[string]interface{}{"deployVrfList": deployment.DeployVrfList, "deploy_map": deployment.DeployMap})
	if deployment.retryFlag {
		//Retry set - use only the failed/out-of-sync ones in deploy
		tflog.Debug(ctx, "deploy: Retrying", map[string]interface{}{"retryList": deployment.RetryList})
		deployMap = &deployment.RetryList
		deployment.retryFlag = false
	} else {
		tflog.Debug(ctx, "deploy: Normal deploy", map[string]interface{}{"deployMap": deployment.DeployMap})
		deployMap = &deployment.DeployMap

	}
	if len(*(deployMap)) != 0 {
		deploy_post_payload := c.deployPayloadBuilder(deployMap)
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
	if len(deployment.DeployVrfList) != 0 {
		payload := map[string]string{"vrfNames": strings.Join(deployment.DeployVrfList, ",")}
		data, err := json.Marshal(payload)
		if err != nil {
			dg.AddError("Error in deploying attachments", err.Error())
			return err
		}
		tflog.Info(ctx, fmt.Sprintf("vrfAttachmentsDeploy: Deploying VRFs %s", string(data)))
		c.GetLock(ResourceVrfBulk).Lock()
		defer c.GetLock(ResourceVrfBulk).Unlock()
		res, err := c.apiClient.Post(fmt.Sprintf(UrlVrfDeployment, deployment.FabricName), string(data))
		if err != nil {
			dg.AddError("Error in deploying attachments", err.Error())
			return err
		}
		tflog.Info(ctx, fmt.Sprintf("vrfAttachmentsDeploy: Success res : %v", res.Str))
	}
	//Reset the retry list
	deployment.RetryList = nil
	deployment.RetryList = make(map[string][]string)

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

func (c NDFC) stateChecker(ctx context.Context, dg *diag.Diagnostics, d *NDFCVrfDeployment) string {
	//Check the state of the VRFs
	vrfs := make([]string, 0)
	for _, v := range d.DeployMap {
		vrfs = append(vrfs, v...)
	}
	vrfs = append(vrfs, d.DeployVrfList...)
	//Get the state of the VRFs
	res, err := c.vrfAttachmentsGet(ctx, d.FabricName, vrfs)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("StateChecker: Error getting VRF Attachments %s", err.Error()))
		return EventFailed
	}
	tflog.Debug(ctx, fmt.Sprintf("Status Checker:  res : %v", string(res)))

	ndVrfs := rva.NDFCVrfAttachmentsPayloads{}
	err = json.Unmarshal(res, &ndVrfs.VrfAttachments)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("StateChecker: Error unmarshalling VRF Attachments %s", err.Error()))
		log.Printf("StateChecker: Error unmarshalling VRF Attachments %s", string(res))
		return EventFailed
	}
	retry := 0
	inPro := 0
	failed := 0
	completed := 0

	for i := range ndVrfs.VrfAttachments {
		for j := range ndVrfs.VrfAttachments[i].AttachList {
			key := ndVrfs.VrfAttachments[i].VrfName + "/" + ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo
			expected, found := d.ExpectedState[key]
			if !found {
				continue
			}
			expected.Seen = true
			d.ExpectedState[key] = expected

			tflog.Debug(ctx, fmt.Sprintf("StateChecker: VRF %s Attachment %s State %s expected %s",
				ndVrfs.VrfAttachments[i].VrfName,
				ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo,
				ndVrfs.VrfAttachments[i].AttachList[j].AttachState, expected.State))
			if ndVrfs.VrfAttachments[i].AttachList[j].AttachState != expected.State {
				switch ndVrfs.VrfAttachments[i].AttachList[j].AttachState {
				case NDFCStateFailed:
					tflog.Error(ctx, fmt.Sprintf("StateChecker: VRF %s/%s deployment failed",
						ndVrfs.VrfAttachments[i].VrfName, ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo))
					//dg.AddError("VRF Deployment Failed", fmt.Sprintf("VRF %s/%s deployment failed",
					//ndVrfs.VrfAttachments[i].VrfName, ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo))
					failed++
					d.RetryList[ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo] = append(d.RetryList[ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo], ndVrfs.VrfAttachments[i].VrfName)
					//return EventFailed
				case NDFCStateDeployed:
					fallthrough
				case NDFCStateInPro:
					inPro++
					tflog.Info(ctx, fmt.Sprintf("StateChecker: VRF %s/%s is in progress",
						ndVrfs.VrfAttachments[i].VrfName, ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo))
					//return EventWait
				case NDFCStatePending:
					tflog.Info(ctx, fmt.Sprintf("StateChecker: VRF %s/%s is pending",
						ndVrfs.VrfAttachments[i].VrfName, ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo))
					retry++
					d.RetryList[ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo] = append(d.RetryList[ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo], ndVrfs.VrfAttachments[i].VrfName)
				case NDFCStateOutOfSync:
					retry++
					tflog.Info(ctx, fmt.Sprintf("StateChecker: VRF %s/%s is out-of-sync",
						ndVrfs.VrfAttachments[i].VrfName, ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo))
					d.RetryList[ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo] = append(d.RetryList[ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo], ndVrfs.VrfAttachments[i].VrfName)
					//return EventRetry
				}
			} else {
				completed++
				//TODO optimize by removing these from check - after the tolerance level has met
				tflog.Info(ctx, fmt.Sprintf("StateChecker:  %s/%s has reached expected state %s:%s", ndVrfs.VrfAttachments[i].VrfName,
					ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo, ndVrfs.VrfAttachments[i].AttachList[j].AttachState, expected.State))
			}
		}
	}
	for k, v := range d.ExpectedState {
		if !v.Seen {
			tflog.Error(ctx, fmt.Sprintf("StateChecker:  %s has not been seen in attachments", k))
			dg.AddError("Attachment missing", fmt.Sprintf("VRF:Attachment %s not visible in NDFC", k))
			return EventFailed
		}
	}
	tflog.Debug(ctx, "StateChecker: Status", map[string]interface{}{"retry": retry, "inPro": inPro, "failed": failed, "completed": completed})

	if failed > 0 {
		return EventFailed
	}
	if retry > 0 {
		return EventRetry
	}
	if inPro > 0 {
		return EventWait
	}

	tflog.Info(ctx, "StateChecker: All VRFs deployed successfully")
	return EventComplete
	//Check the state of the VRFs

}
