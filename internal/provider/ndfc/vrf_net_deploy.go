// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package ndfc

//This file contains VRF resource specific implementation of NDFCDeployment and NDFCDeployRsc interfaces

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"
	"terraform-provider-ndfc/internal/provider/ndfc/api"
	rva "terraform-provider-ndfc/internal/provider/resources/resource_vrf_attachments"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Implement DeployRsc interface

type NDFCVrfNetworkDeployRsc struct {
	NDFCDeployRsc //common parameters
	RsList        []string
	Attachment    []string
	PostCount     int
}

type NDFCVrfNetworkDeployment struct {
	NDFCDeployment
	RsType           string
	RscByAttachments map[string][]DeployRsc
}

func NewVrfNetworkDeployment(c *NDFC, fabricName string, rsType string) *NDFCVrfNetworkDeployment {
	o := new(NDFCVrfNetworkDeployment)
	o.NDFCDeployment.Deployment = o
	o.FabricName = fabricName
	o.ctrlr = c
	o.DeployRscDB = make(map[string]map[string]DeployRsc)
	o.DeployMap = make(map[string]DeployRsc)
	o.RscByAttachments = make(map[string][]DeployRsc)
	o.RsType = rsType
	return o
}

func (c *NDFCVrfNetworkDeployment) GetLock() *sync.Mutex {
	if c.RsType == ResourceNetworks {
		return c.ctrlr.GetLock(ResourceNetworks)
	}

	if c.RsType == "vrfs" { // ResourceVrfBulk
		return c.ctrlr.GetLock(ResourceVrfBulk)
	}
	return nil
}

func (c *NDFCVrfNetworkDeployRsc) GetPostPayload() string {
	deploy_post_payload := "\"" + c.Attachment[0] + "\":\""
	deploy_post_payload += strings.Join(c.RsList, ",")
	deploy_post_payload += "\""
	return deploy_post_payload
}

func (d *NDFCVrfNetworkDeployment) deployPayload(depRsc *NDFCVrfNetworkDeployRsc) string {
	deploy_post_payload := "{"
	deploy_post_payload += depRsc.GetPostPayload()
	deploy_post_payload += "}"
	deploy_post_payload = strings.ReplaceAll(deploy_post_payload, ",}", "}")
	return deploy_post_payload
}

func (d *NDFCVrfNetworkDeployment) deployPayloadBulk(serial string) string {
	deploy_post_payload := "{"
	deploy_post_payload += "\"" + serial + "\":\""
	for _, rsc := range d.RscByAttachments[serial] {
		depRsc := rsc.(*NDFCVrfNetworkDeployRsc)
		deploy_post_payload += strings.Join(depRsc.RsList, ",")
		deploy_post_payload += ","
	}
	deploy_post_payload = strings.TrimSuffix(deploy_post_payload, ",")
	deploy_post_payload += "\""
	deploy_post_payload += "}"
	return deploy_post_payload
}

// call back from FSM
func (d *NDFCVrfNetworkDeployment) Deploy(ctx context.Context, dg *diag.Diagnostics, deployRscList []DeployRsc, bulk bool) {
	tflog.Debug(ctx, fmt.Sprintf("Deploying %s", d.RsType), map[string]interface{}{"deployRscList": deployRscList})
	retryRscList := make([]DeployRsc, 0)
	//bulkRetry := make(map[string]bool)

	if bulk {
		// Bulk Deploy
		// deployRsc is nil - use RscByAttachments
		for serial := range d.RscByAttachments {
			deployment_ok := false
			depAPI := api.NewDeploymentAPI("", d.GetLock(), &d.ctrlr.apiClient, d.RsType)
			//depAPI.SetDeployLocked()
			deploy_post_payload := d.deployPayloadBulk(serial)
			tflog.Info(ctx, fmt.Sprintf("Deploy:%s Deploying Attachments %s", d.RsType, deploy_post_payload))
			for i := 0; i < 3; i++ {
				res, err := depAPI.DeployPost([]byte(deploy_post_payload))
				if err != nil {
					tflog.Error(ctx, fmt.Sprintf("Deploy: Error in bulk deployment %s", err.Error()))
					continue
				}
				deployment_ok = true
				tflog.Info(ctx, fmt.Sprintf("Deploy, bulk: Success res : %v", res.Str))
				break
			}
			if !deployment_ok {
				tflog.Error(ctx, fmt.Sprintf("Deploy: Max retry count reached for %s", serial))
				dg.AddError("POST failure", fmt.Sprintf("Max retry count reached for %s", serial))
				for _, rsc := range d.RscByAttachments[serial] {
					depRsc := rsc.(*NDFCVrfNetworkDeployRsc)
					depRsc.SetCurrentState(NDFCStateFailed)
					depRsc.failCnt = 3
				}
			}
		}
		return
	}
	if len(deployRscList) != 0 {
		for _, v := range deployRscList {
			depRsc := v.(*NDFCVrfNetworkDeployRsc)
			if depRsc.PostCount < 3 {
				deploy_post_payload := d.deployPayload(depRsc)
				tflog.Info(ctx, fmt.Sprintf("Deploy: %s: Deploying Attachments %s", d.RsType, deploy_post_payload))
				depAPI := api.NewDeploymentAPI("", d.GetLock(), &d.ctrlr.apiClient, d.RsType)
				//depAPI.SetDeployLocked()
				res, err := depAPI.DeployPost([]byte(deploy_post_payload))
				if err != nil {
					//dg.AddError("Error in deploying attachments", err.Error())
					tflog.Error(ctx, fmt.Sprintf("Deploy: %s: Error in deploying attachments %s", d.RsType, err.Error()))
					depRsc.PostCount++
					retryRscList = append(retryRscList, v)
				}
				tflog.Info(ctx, fmt.Sprintf("Deploy: %s: Success res : %v", d.RsType, res.Str))
			} else {
				depRsc.SetCurrentState(NDFCStateFailed)
				depRsc.failCnt = depRsc.PostCount
				tflog.Error(ctx, fmt.Sprintf("Deploy: %s: Max retry count reached for %s", d.RsType, depRsc.Key))
				dg.AddError("Deploy failure during POST", fmt.Sprintf("Max retry count reached for %s", depRsc.Key))
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

func (d *NDFCVrfNetworkDeployment) getRsList(deploy []DeployRsc) []string {
	ret := make(map[string]bool)
	for _, v := range deploy {
		depRsc := v.(*NDFCVrfNetworkDeployRsc)
		for _, name := range depRsc.RsList {
			ret[name] = true
		}
	}
	retList := make([]string, 0)
	for k := range ret {
		retList = append(retList, k)
	}
	return retList
}

func (d *NDFCVrfNetworkDeployment) updateState(deployRsc []DeployRsc, state string) {
	for _, v := range deployRsc {
		depRsc := v.(*NDFCVrfNetworkDeployRsc)
		depRsc.SetCurrentState(state)
	}
}

type NDFCDeployStatus struct {
	Key    string
	Status string
}

func (d *NDFCVrfNetworkDeployment) getDeployStatus(ctx context.Context, rsList []string) []NDFCDeployStatus {
	ret := make([]NDFCDeployStatus, 0)
	if d.RsType == ResourceNetworks {
		ndNws, err := d.ctrlr.netAttachmentsGet(ctx, d.FabricName, rsList)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("StateChecker: Error getting Network Attachments %s", err.Error()))
			return nil
		}
		//tflog.Debug(ctx, fmt.Sprintf("Status Checker:  res : %v", string(res)))

		for i := range ndNws.NetworkAttachments {
			for j := range ndNws.NetworkAttachments[i].Attachments {
				/*
					if !*(ndNws.NetworkAttachments[i].Attachments[j].Attached) {
						log.Println("StateChecker: Skipping unattached entry", ndNws.NetworkAttachments[i].Attachments[j].SwitchSerialNo)
						continue
					}
				*/
				rs := NDFCDeployStatus{}
				rs.Key = ndNws.NetworkAttachments[i].NetworkName + "/" + ndNws.NetworkAttachments[i].Attachments[j].SwitchSerialNo
				rs.Status = ndNws.NetworkAttachments[i].Attachments[j].AttachState
				ret = append(ret, rs)
			}
		}
	}
	if d.RsType == "vrfs" {
		res, err := d.ctrlr.vrfAttachmentsGet(ctx, d.FabricName, rsList)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("StateChecker: Error getting VRF Attachments %s", err.Error()))
			return nil
		}
		tflog.Debug(ctx, fmt.Sprintf("Status Checker:  res : %v", string(res)))

		ndVrfs := rva.NDFCVrfAttachmentsPayloads{}
		err = json.Unmarshal(res, &ndVrfs.VrfAttachments)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("StateChecker: Error unmarshalling  VRF Attachments %s", err.Error()))
			log.Printf("StateChecker: Error unmarshalling VRF Attachments %s", string(res))
			return nil
		}

		for i := range ndVrfs.VrfAttachments {
			for j := range ndVrfs.VrfAttachments[i].AttachList {
				rs := NDFCDeployStatus{}
				rs.Key = ndVrfs.VrfAttachments[i].VrfName + "/" + ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo
				rs.Status = ndVrfs.VrfAttachments[i].AttachList[j].AttachState
				ret = append(ret, rs)
			}
		}
	}
	return ret
}

func (d *NDFCVrfNetworkDeployment) CheckState(ctx context.Context, dg *diag.Diagnostics, deployRscList []DeployRsc) string {

	//d.constructExpectedState(deployRscList)
	rsList := d.getRsList(deployRscList)
	tflog.Debug(ctx, fmt.Sprintf("Checking %s", d.RsType), map[string]interface{}{d.RsType: rsList})

	depStates := d.getDeployStatus(ctx, rsList)
	if depStates == nil {
		tflog.Error(ctx, "StateChecker: Error getting deploy status")
		d.updateState(deployRscList, NDFCStateFailed)
		return EventFailed
	}

	failed := 0
	inPro := 0
	retry := 0
	completed := 0

	for _, dep := range depStates {
		r := d.GetDeployRsc(dep.Key)
		if r == nil {
			log.Printf("StateChecker:%s: No deployRsc found for: %s Skip", d.RsType, dep.Key)
			//not interested in this entry - skip
			continue
		}
		rsc := r.(*NDFCVrfNetworkDeployRsc)
		rsc.checkCount++
		tflog.Debug(ctx, fmt.Sprintf("StateChecker:%s:  %s State %s expected %s",
			d.RsType, dep.Key, dep.Status, rsc.State.ExpectedState))
		if dep.Status != rsc.State.ExpectedState {
			switch dep.Status {
			case NDFCStateFailed:
				tflog.Error(ctx, fmt.Sprintf("StateChecker:%s:  %s deployment failed",
					d.RsType, dep.Key))
				failed++
				rsc.SetCurrentState(NDFCStateFailed)
				//d.RetryList[ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo] = append(d.RetryList[ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo], ndVrfs.VrfAttachments[i].VrfName)
				//return EventFailed
			case NDFCStateDeployed:
				// Deployed coming here would be a case of deletion where expected is NA
				// This is a special case - NDFC keeps detached entries in "DEPLOYED" state for a while
				// before moving to NA
				// Keep such entries in hanging list and check again at the end
				if rsc.State.ExpectedState == NDFCStateNA {
					rsc.SetCurrentState(NDFCHanging)
				} else {
					rsc.SetCurrentState(NDFCStateInPro)
					inPro++
				}
			case NDFCStateInPro:
				inPro++
				tflog.Info(ctx, fmt.Sprintf("StateChecker:%s: %s is in progress",
					d.RsType, dep.Key))
				rsc.SetCurrentState(NDFCStateInPro)
				//return EventWait
			case NDFCStatePending:
				tflog.Info(ctx, fmt.Sprintf("StateChecker:%s: %s is pending",
					d.RsType, dep.Key))
				retry++
				rsc.SetCurrentState(NDFCStatePending)
				//d.RetryList[ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo] = append(d.RetryList[ndVrfs.VrfAttachments[i].AttachList[j].SwitchSerialNo], ndVrfs.VrfAttachments[i].VrfName)
			case NDFCStateOutOfSync:
				if rsc.LastState == NDFCStateOutOfSync && rsc.OOSyncCnt > 3 {
					tflog.Error(ctx, fmt.Sprintf("StateChecker:%s: %s is out-of-sync for too long retry",
						d.RsType, dep.Key))
					retry++
					rsc.SetCurrentState(NDFCStateOutOfSync)
					rsc.OOSyncCnt = 0
				} else {
					rsc.OOSyncCnt++
					tflog.Info(ctx, fmt.Sprintf("StateChecker:%s: %s is out-of-sync for %d Waiting",
						d.RsType, dep.Key, rsc.OOSyncCnt))
					inPro++
					rsc.SetCurrentState(NDFCStateInPro)
				}
			}
		} else {
			rsc.SetCurrentState(rsc.State.ExpectedState)
			rsc.checkTick++
			completed++
			//TODO optimize by removing these from check - after the tolerance level has met
			tflog.Info(ctx, fmt.Sprintf("StateChecker:%s: %s has reached expected state %s:%s",
				d.RsType, dep.Key, dep.Status, rsc.State.ExpectedState))
		}
		rsc.LastState = dep.Status
	}

	for _, v := range deployRscList {
		rsc := v.(*NDFCVrfNetworkDeployRsc)
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

func NewNDFCVrfNetworkDeployRsc(key string) *NDFCVrfNetworkDeployRsc {
	o := new(NDFCVrfNetworkDeployRsc)
	o.RsList = make([]string, 0)
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

func (d *NDFCVrfNetworkDeployment) updateDeploymentDB(serial, name, state string) {
	var depRsc *NDFCVrfNetworkDeployRsc
	key := name + "/" + serial
	dep := d.GetDeployRsc(key)
	if dep != nil {
		depRsc = dep.(*NDFCVrfNetworkDeployRsc)
	} else {
		depRsc = NewNDFCVrfNetworkDeployRsc(key)
		d.AddDeploymentRsc(key, depRsc)
	}
	log.Println("updateDeploymentDB: ", key, state)
	depRsc.RsList = append(depRsc.RsList, name)
	depRsc.Attachment = append(depRsc.Attachment, serial)
	updateExpectedState(&depRsc.State, state)
	rsList, ok := d.RscByAttachments[serial]
	if !ok {
		rsList = make([]DeployRsc, 0)
	}
	rsList = append(rsList, depRsc)
	d.RscByAttachments[serial] = rsList
}
