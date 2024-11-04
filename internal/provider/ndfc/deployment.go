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

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const (
	DeployPending   = "pending"
	DeployInpro     = "inpro"
	DeployRetry     = "retry"
	DeployTrustWait = "trustwait"
	DeployComplete  = "complete"
	DeployFailed    = "failed"
	DeployHanging   = "hanging"
)

type DeploymentState struct {
	Rsc           DeployRsc
	ExpectedState string
	CurrentState  string
	Seen          bool
}

type Deployment interface {
	Deploy(ctx context.Context, dg *diag.Diagnostics, deployRscList []DeployRsc, bulk bool)
	CheckState(ctx context.Context, dg *diag.Diagnostics, deployRscList []DeployRsc) string
}
type NDFCDeployment struct {
	Deployment
	ctrlr       *NDFC
	FabricName  string
	DeployRscDB map[string]map[string]DeployRsc
	DeployMap   map[string]DeployRsc
	retryFlag   bool
}

/*
Call coming from FSM

	rsc is nil, bulk is ignored.
	Parameters retained for interface compatibility
*/
func (n *NDFCDeployment) Deploy(ctx context.Context, dg *diag.Diagnostics, rsc []DeployRsc, bulk bool) {
	//rsc is nil at this level - using the same API for interface compatibility
	//Call the override method
	log.Println("Deploy: ", rsc, n)
	rscList := make([]DeployRsc, 0)
	depCount := 0
	inProCount := len(n.DeployRscDB[DeployInpro])
	if n.ctrlr.MaxParallelDeploy == 0 {
		//Deploy Everything
		depCount = len(n.DeployRscDB[DeployPending])
		// Do a bulk deploy at the start
		// further retries only for resources re-added to pending list
		if n.retryFlag {
			bulk = false
		} else {
			bulk = true
		}
	} else {
		depCount = (n.ctrlr.MaxParallelDeploy - inProCount)
		bulk = false
	}
	i := 0
	if depCount <= 0 {
		tflog.Debug(ctx, "Deployments already in progress ", map[string]interface{}{
			"in-progress":  inProCount,
			"max-parallel": n.ctrlr.MaxParallelDeploy,
			"pending":      len(n.DeployRscDB[DeployPending])})
		return
	}
	//var depList map[string]DeployRsc

	if len(n.DeployRscDB[DeployPending]) != 0 {
		for _, v := range n.DeployRscDB[DeployPending] {
			rscList = append(rscList, v)
			i++
			if i >= depCount {
				break
			}
		}
	}
	for _, v := range rscList {
		n.MoveList(DeployPending, DeployInpro, v)
	}

	if len(rscList) > 0 {
		n.Deployment.Deploy(ctx, dg, rscList, bulk)
	}
	//Check if there was a deploy failure
	for _, v := range rscList {
		if v.GetCurrentState() == NDFCStateFailed {
			tflog.Error(ctx, "Deploy: Resource failed to deploy due to error in deploy operation. ", map[string]interface{}{
				"resource": v.GetKey(),
				"state":    v.GetCurrentState(),
			})

			n.MoveList(DeployInpro, DeployFailed, v)
		}
	}
}

func (n *NDFCDeployment) CheckState(ctx context.Context, dg *diag.Diagnostics, rscs []DeployRsc) string {
	rscs = make([]DeployRsc, 0)
	//Check the states of resources in inpro list
	for _, v := range n.DeployRscDB[DeployInpro] {
		rscs = append(rscs, v)
	}
	n.Deployment.CheckState(ctx, dg, rscs)
	for _, v := range rscs {
		if v.GetCurrentState() == v.GetExpectedState() {
			if v.GetCheckTick() < n.ctrlr.DeployTrustFactor {
				tflog.Debug(ctx, "CheckState: Resource reached expected state. Moving to trustwait", map[string]interface{}{
					"resource": v.GetKey(),
					"state":    v.GetCurrentState(),
				})
				n.MoveList(DeployInpro, DeployTrustWait, v)
			} else {
				tflog.Debug(ctx, "CheckState: Resource in expected state until TrustFactor. Moving to complete", map[string]interface{}{
					"resource":     v.GetKey(),
					"state":        v.GetCurrentState(),
					"trust-checks": v.GetCheckTick(),
				})
				n.MoveList(DeployInpro, DeployComplete, v)
			}
		} else {
			if v.GetCurrentState() == NDFCStateFailed {
				if v.GetFailureCount() < n.ctrlr.FailureRetry {
					tflog.Warn(ctx, "CheckState: Resource failed. Could be transient; Move to Pending list for retry", map[string]interface{}{
						"resource":    v.GetKey(),
						"state":       v.GetCurrentState(),
						"retry-count": v.GetFailureCount(),
						"max-retry":   n.ctrlr.FailureRetry,
						"new-state":   NDFCStatePending,
					})
					n.retryFlag = true
					v.SetCurrentState(NDFCStatePending)
					n.MoveList(DeployInpro, DeployPending, v)
				} else {
					tflog.Error(ctx, "CheckState: Resource failed. Max retry reached. Moving to failed list", map[string]interface{}{
						"resource":    v.GetKey(),
						"state":       v.GetCurrentState(),
						"retry-count": v.GetFailureCount(),
						"max-retry":   n.ctrlr.FailureRetry,
					})

					n.MoveList(DeployInpro, DeployFailed, v)
				}
			} else if v.GetCurrentState() == NDFCStatePending {
				tflog.Debug(ctx, "CheckState: Resource moved to a transient state, moving to pending list for retry", map[string]interface{}{
					"resource": v.GetKey(),
					"state":    v.GetCurrentState(),
				})

				n.retryFlag = true
				n.MoveList(DeployInpro, DeployPending, v)
			} else if v.GetCurrentState() == NDFCHanging {
				tflog.Error(ctx, "CheckState: Resource in transient deployed state. Moving to hanging list for checking later", map[string]interface{}{
					"resource": v.GetKey(),
					"state":    v.GetCurrentState(),
				})
				n.MoveList(DeployInpro, DeployHanging, v)
			}
		}
	}
	maxDeploy := n.ctrlr.MaxParallelDeploy
	if n.ctrlr.MaxParallelDeploy == 0 {
		maxDeploy = len(n.DeployMap)
	}
	if len(n.DeployRscDB[DeployInpro]) < maxDeploy {
		tflog.Debug(ctx, "CheckState: Inpro list got room for next set", map[string]interface{}{
			"max-parallel": maxDeploy,
			"inpro":        len(n.DeployRscDB[DeployInpro]),
		})

		if len(n.DeployRscDB[DeployPending]) > 0 {
			tflog.Debug(ctx, "CheckState: Pending list has resources, triggering deploy", map[string]interface{}{
				"pending": len(n.DeployRscDB[DeployPending]),
			})
			return EventDeploy
		}
		tflog.Debug(ctx, "CheckState: No resources in pending list", map[string]interface{}{
			"pending": len(n.DeployRscDB[DeployPending]),
		})

		// Nothing in pending list
		//Move everything from trustwait to In Pro for ensuring status
		tflog.Debug(ctx, "CheckState: Moving resources from trustwait to Inpro", map[string]interface{}{
			"trustwait": len(n.DeployRscDB[DeployTrustWait]),
		})
		for _, v := range n.DeployRscDB[DeployTrustWait] {
			log.Printf("Moving %s from trustwait to inpro", v.GetKey())
			n.MoveList(DeployTrustWait, DeployInpro, v)
		}
		if len(n.DeployRscDB[DeployInpro]) > 0 {
			tflog.Debug(ctx, "CheckState: Inpro list has resources - triggering wait", map[string]interface{}{
				"inpro": len(n.DeployRscDB[DeployInpro]),
			})
			return EventWait
		}
		// Everything processed - time to check hanging list
		if len(n.DeployRscDB[DeployHanging]) > 0 {
			tflog.Debug(ctx, "CheckState: Hanging list has resources - triggering wait", map[string]interface{}{
				"hanging": len(n.DeployRscDB[DeployHanging]),
			})
			for _, v := range n.DeployRscDB[DeployHanging] {
				n.MoveList(DeployHanging, DeployInpro, v)
			}
			return EventWait
		}
		//Nothing in hanging list
		//Check if there are any failed resources
		tflog.Debug(ctx, "All deploy lists  processed - ending deployment cycles", map[string]interface{}{
			"inpro":     len(n.DeployRscDB[DeployInpro]),
			"pending":   len(n.DeployRscDB[DeployPending]),
			"trustwait": len(n.DeployRscDB[DeployTrustWait]),
			"failed":    len(n.DeployRscDB[DeployFailed]),
		})

		if len(n.DeployRscDB[DeployFailed]) > 0 {
			tflog.Error(ctx, "CheckState: Failed list has resources; trigger failure", map[string]interface{}{
				"failed": len(n.DeployRscDB[DeployFailed]),
			})
			dg.AddError("Deployment failed", "Some resources failed to deploy")
			for _, v := range n.DeployRscDB[DeployFailed] {
				dg.AddError("Resource Deployment failed", v.GetKey())
			}
			return EventFailed
		}
		//Nothing in failed list
		tflog.Info(ctx, "CheckState: All resources deployed successfully", map[string]interface{}{
			"complete": len(n.DeployRscDB[DeployComplete]),
		})
		return EventComplete

	}
	tflog.Debug(ctx, "CheckState: No room for additional resources in inpro list - trigger wait", map[string]interface{}{
		"max-parallel": maxDeploy,
		"inpro":        len(n.DeployRscDB[DeployInpro]),
	})
	return EventWait
}

func (n *NDFCDeployment) DeployFSM(ctx context.Context, dg *diag.Diagnostics) {
	log.Printf("DeployFSM")
	log.Printf("============================Starting Deployment===================================")
	tflog.Debug(ctx, fmt.Sprintf("vrfAttachmentsDeploy: Starting FSM %v", n.DeployMap))
	deployFsm := NewDeployFSM(ctx, dg, n)
	deployFsm.Run(ctx)
	tflog.Debug(ctx, fmt.Sprintf("vrfAttachmentsDeploy: Finishing FSM %v", n.DeployMap))

}

func (dep *NDFCDeployment) AddRscToList(listName string, rsc DeployRsc) {
	targetMap, ok := dep.DeployRscDB[listName]
	if !ok {
		targetMap = make(map[string]DeployRsc)
	}
	targetMap[rsc.GetKey()] = rsc
	dep.DeployRscDB[listName] = targetMap
}

func (dep *NDFCDeployment) MoveList(from, to string, rsc DeployRsc) {
	log.Printf("Move %s: %s=>%s", rsc.GetKey(), from, to)
	if _, ok := dep.DeployRscDB[from]; !ok {
		log.Println("MoveList: No resources in list", from)
		panic("MoveList: No resources in list" + from)
	}
	delete(dep.DeployRscDB[from], rsc.GetKey())
	dep.AddRscToList(to, rsc)
}
func (dep *NDFCDeployment) GetDeployRsc(key string) DeployRsc {
	return dep.DeployMap[key]
}

func (dep *NDFCDeployment) AddDeploymentRsc(key string, rsc DeployRsc) {
	dep.DeployMap[key] = rsc
	dep.AddRscToList(DeployPending, rsc)
}

func (dep NDFCDeployment) GetDeployPendingCount() int {
	return len(dep.DeployRscDB[DeployPending])
}
