package ndfc

//This file contains VRF resource specific implementation of NDFCDeployment and NDFCDeployRsc interfaces

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

// Implement DeployRsc interface

type NDFCVrfDeployRsc struct {
	NDFCDeployRsc //common parameters
	VrfList       []string
	Attachment    []string
	PostCount     int
}

type NDFCVrfDeployment struct {
	NDFCDeployment
	RscByAttachments map[string][]DeployRsc
}

func NewVrfDeployment(c *NDFC, fabricName string) *NDFCVrfDeployment {
	o := new(NDFCVrfDeployment)
	o.NDFCDeployment.Deployment = o
	o.FabricName = fabricName
	o.ctrlr = c
	o.DeployRscDB = make(map[string]map[string]DeployRsc)
	o.DeployMap = make(map[string]DeployRsc)
	o.RscByAttachments = make(map[string][]DeployRsc)
	return o
}

func (c *NDFCVrfDeployRsc) GetPostPayload() string {
	deploy_post_payload := "\"" + c.Attachment[0] + "\":\""
	deploy_post_payload += strings.Join(c.VrfList, ",")
	deploy_post_payload += "\""
	return deploy_post_payload
}

func (d *NDFCVrfDeployment) deployPayload(depRsc *NDFCVrfDeployRsc) string {
	deploy_post_payload := "{"
	deploy_post_payload += depRsc.GetPostPayload()
	deploy_post_payload += "}"
	deploy_post_payload = strings.ReplaceAll(deploy_post_payload, ",}", "}")
	return deploy_post_payload
}

func (d *NDFCVrfDeployment) deployPayloadBulk(serial string) string {
	deploy_post_payload := "{"
	deploy_post_payload += "\"" + serial + "\":\""
	for _, rsc := range d.RscByAttachments[serial] {
		depRsc := rsc.(*NDFCVrfDeployRsc)
		deploy_post_payload += strings.Join(depRsc.VrfList, ",")
		deploy_post_payload += ","
	}
	deploy_post_payload = strings.TrimSuffix(deploy_post_payload, ",")
	deploy_post_payload += "\""
	deploy_post_payload += "}"
	return deploy_post_payload
}

func (d *NDFCVrfDeployment) Deploy(ctx context.Context, dg *diag.Diagnostics, deployRscList []DeployRsc, bulk bool) {
	tflog.Debug(ctx, "Deploying VRFs", map[string]interface{}{"deployRscList": deployRscList})
	retryRscList := make([]DeployRsc, 0)
	//bulkRetry := make(map[string]bool)

	if bulk {
		// Bulk Deploy
		// deployRsc is nil - use RscByAttachments
		for serial, _ := range d.RscByAttachments {
			deployment_ok := false
			deploy_post_payload := d.deployPayloadBulk(serial)
			tflog.Info(ctx, fmt.Sprintf("vrfAttachmentsDeploy: Deploying Attachments %s", deploy_post_payload))
			d.ctrlr.GetLock(ResourceVrfBulk).Lock()
			for i := 0; i < 3; i++ {
				res, err := d.ctrlr.apiClient.Post(UrlVrfAttachmentsDeploy, deploy_post_payload)
				if err != nil {
					tflog.Error(ctx, fmt.Sprintf("Deploy: Error in bulk deployment %s", err.Error()))
					continue
				}
				deployment_ok = true
				tflog.Info(ctx, fmt.Sprintf("Deploy, bulk: Success res : %v", res.Str))
				break
			}
			d.ctrlr.GetLock(ResourceVrfBulk).Unlock()
			if !deployment_ok {
				tflog.Error(ctx, fmt.Sprintf("Deploy: Max retry count reached for %s", serial))
				dg.AddError("POST failure", fmt.Sprintf("Max retry count reached for %s", serial))
				for _, rsc := range d.RscByAttachments[serial] {
					depRsc := rsc.(*NDFCVrfDeployRsc)
					depRsc.SetCurrentState(NDFCStateFailed)
					depRsc.failCnt = 3
				}
			}
		}
		return
	}
	if len(deployRscList) != 0 {
		for _, v := range deployRscList {
			depRsc := v.(*NDFCVrfDeployRsc)
			if depRsc.PostCount < 3 {
				deploy_post_payload := d.deployPayload(depRsc)
				tflog.Info(ctx, fmt.Sprintf("vrfAttachmentsDeploy: Deploying Attachments %s", deploy_post_payload))
				d.ctrlr.GetLock(ResourceVrfBulk).Lock()
				res, err := d.ctrlr.apiClient.Post(UrlVrfAttachmentsDeploy, deploy_post_payload)
				if err != nil {
					//dg.AddError("Error in deploying attachments", err.Error())
					tflog.Error(ctx, fmt.Sprintf("vrfAttachmentsDeploy: Error in deploying attachments %s", err.Error()))
					depRsc.PostCount++
					retryRscList = append(retryRscList, v)
				}
				d.ctrlr.GetLock(ResourceVrfBulk).Unlock()
				tflog.Info(ctx, fmt.Sprintf("vrfAttachmentsDeploy: Success res : %v", res.Str))
			} else {
				depRsc.SetCurrentState(NDFCStateFailed)
				depRsc.failCnt = depRsc.PostCount
				tflog.Error(ctx, fmt.Sprintf("vrfAttachmentsDeploy: Max retry count reached for %s", depRsc.Key))
				dg.AddError("POST failure", fmt.Sprintf("Max retry count reached for %s", depRsc.Key))
			}

		}
		if len(retryRscList) != 0 {
			d.Deploy(ctx, dg, retryRscList, false)
		}
	} else {
		tflog.Error(ctx, "Deploy list Empty")
		return
	}
}

func (d *NDFCVrfDeployment) getVrfList(deploy []DeployRsc) []string {
	ret := make(map[string]bool)
	for _, v := range deploy {
		depRsc := v.(*NDFCVrfDeployRsc)
		for _, vrfName := range depRsc.VrfList {
			ret[vrfName] = true
		}
	}
	retList := make([]string, 0)
	for k := range ret {
		retList = append(retList, k)
	}
	return retList
}

func (d *NDFCVrfDeployment) updateState(deployRsc []DeployRsc, state string) {
	for _, v := range deployRsc {
		depRsc := v.(*NDFCVrfDeployRsc)
		depRsc.SetCurrentState(state)
	}
}

func (d *NDFCVrfDeployment) CheckState(ctx context.Context, dg *diag.Diagnostics, deployRscList []DeployRsc) string {

	//d.constructExpectedState(deployRscList)
	vrfList := d.getVrfList(deployRscList)
	tflog.Debug(ctx, "Checking VRFs", map[string]interface{}{"vrfs": vrfList})
	res, err := d.ctrlr.vrfAttachmentsGet(ctx, d.FabricName, vrfList)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("StateChecker: Error getting VRF Attachments %s", err.Error()))
		d.updateState(deployRscList, NDFCStateFailed)
		return EventFailed
	}
	tflog.Debug(ctx, fmt.Sprintf("Status Checker:  res : %v", string(res)))

	ndVrfs := rva.NDFCVrfAttachmentsPayloads{}
	err = json.Unmarshal(res, &ndVrfs.VrfAttachments)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("StateChecker: Error unmarshalling VRF Attachments %s", err.Error()))
		log.Printf("StateChecker: Error unmarshalling VRF Attachments %s", string(res))
		d.updateState(deployRscList, NDFCStateFailed)
		return EventFailed
	}

	failed := 0
	inPro := 0
	retry := 0
	completed := 0

	for i := range ndVrfs.VrfAttachments {
		for j := range ndVrfs.VrfAttachments[i].AttachList {
			key := ndVrfs.VrfAttachments[i].VrfName + "/" + ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo
			r := d.GetDeployRsc(key)
			if r == nil {
				log.Println("StateChecker: No deployRsc found for: Skip", key)
				//not interested in this entry - skip
				continue
			}
			rsc := r.(*NDFCVrfDeployRsc)
			rsc.checkCount++
			tflog.Debug(ctx, fmt.Sprintf("StateChecker: VRF %s Attachment %s State %s expected %s",
				ndVrfs.VrfAttachments[i].VrfName,
				ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo,
				ndVrfs.VrfAttachments[i].AttachList[j].AttachState, rsc.State.ExpectedState))

			if ndVrfs.VrfAttachments[i].AttachList[j].AttachState != rsc.State.ExpectedState {
				switch ndVrfs.VrfAttachments[i].AttachList[j].AttachState {
				case NDFCStateFailed:
					tflog.Error(ctx, fmt.Sprintf("StateChecker: VRF %s/%s deployment failed",
						ndVrfs.VrfAttachments[i].VrfName, ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo))
					failed++
					rsc.SetCurrentState(NDFCStateFailed)
					//d.RetryList[ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo] = append(d.RetryList[ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo], ndVrfs.VrfAttachments[i].VrfName)
					//return EventFailed
				case NDFCStateDeployed:
					// Deployed coming here would be a case of deletion where expected is NA
					// This is a special case - NDFC keeps detached entries in "DEPLOYED" state for a while
					// before moving to NA
					// Keep such entries in pending list and check again
					if rsc.State.ExpectedState == NDFCStateNA {
						rsc.SetCurrentState(NDFCHanging)
					} else {
						rsc.SetCurrentState(NDFCStateInPro)
						inPro++
					}
				case NDFCStateInPro:
					inPro++
					tflog.Info(ctx, fmt.Sprintf("StateChecker: VRF %s/%s is in progress",
						ndVrfs.VrfAttachments[i].VrfName, ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo))
					rsc.SetCurrentState(NDFCStateInPro)
					//return EventWait
				case NDFCStatePending:
					tflog.Info(ctx, fmt.Sprintf("StateChecker: VRF %s/%s is pending",
						ndVrfs.VrfAttachments[i].VrfName, ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo))
					retry++
					rsc.SetCurrentState(NDFCStatePending)
					//d.RetryList[ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo] = append(d.RetryList[ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo], ndVrfs.VrfAttachments[i].VrfName)
				case NDFCStateOutOfSync:
					retry++
					tflog.Info(ctx, fmt.Sprintf("StateChecker: VRF %s/%s is out-of-sync",
						ndVrfs.VrfAttachments[i].VrfName, ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo))
					rsc.SetCurrentState(NDFCStateOutOfSync)
					//d.RetryList[ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo] = append(d.RetryList[ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo], ndVrfs.VrfAttachments[i].VrfName)
					//return EventRetry
				}
			} else {
				rsc.SetCurrentState(rsc.State.ExpectedState)
				rsc.checkTick++
				completed++
				//TODO optimize by removing these from check - after the tolerance level has met
				tflog.Info(ctx, fmt.Sprintf("StateChecker:  %s/%s has reached expected state %s:%s", ndVrfs.VrfAttachments[i].VrfName,
					ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo, ndVrfs.VrfAttachments[i].AttachList[j].AttachState, rsc.State.ExpectedState))
			}
		}
	}
	for _, v := range deployRscList {
		rsc := v.(*NDFCVrfDeployRsc)
		if rsc.checkCount == 0 {
			failed++
			tflog.Error(ctx, fmt.Sprintf("StateChecker:  %s has not been seen in attachments", rsc.Key))
			rsc.SetCurrentState(NDFCStateFailed)
		}
	}

	if failed > 0 {
		return EventFailed
	}
	if retry > 0 {
		return EventRetry
	}
	if inPro > 0 {
		return EventWait
	}
	return EventComplete
}

func NewNDFCVrfDeployRsc(key string) *NDFCVrfDeployRsc {
	o := new(NDFCVrfDeployRsc)
	o.VrfList = make([]string, 0)
	o.DeployRsc = o
	o.Attachment = make([]string, 0)
	o.Key = key
	return o
}

func updateExpectedState(out *DeploymentState, state string) {
	if state == "false" {
		out.ExpectedState = NDFCStateNA
	} else {
		out.ExpectedState = NDFCStateDeployed
	}
	out.CurrentState = ""
	out.Seen = false
}

func (d *NDFCVrfDeployment) updateDeploymentDB(serial, vrfName, state string) {
	var depRsc *NDFCVrfDeployRsc
	key := vrfName + "/" + serial
	dep := d.GetDeployRsc(key)
	if dep != nil {
		depRsc = dep.(*NDFCVrfDeployRsc)
	} else {
		depRsc = NewNDFCVrfDeployRsc(key)
		d.AddDeploymentRsc(key, depRsc)
	}
	log.Println("updateDeploymentDB: ", key, state)
	depRsc.VrfList = append(depRsc.VrfList, vrfName)
	depRsc.Attachment = append(depRsc.Attachment, serial)
	updateExpectedState(&depRsc.State, state)
	rsList, ok := d.RscByAttachments[serial]
	if !ok {
		rsList = make([]DeployRsc, 0)
	}
	rsList = append(rsList, depRsc)
	d.RscByAttachments[serial] = rsList
}
