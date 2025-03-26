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
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/looplab/fsm"
)

const (
	StatePending  = "Pending"
	StateInPro    = "In Progress"
	StateComplete = "COMPLETE"
	StateFailed   = "Failed"
	StateCheck    = "Checking"
)

const (
	NDFCStatePending   = "PENDING"
	NDFCStateNA        = "NA"
	NDFCStateInPro     = "IN PROGRESS"
	NDFCStateDeployed  = "DEPLOYED"
	NDFCStateOutOfSync = "OUT-OF-SYNC"
	NDFCStateFailed    = "FAILED"
	NDFCHanging        = "HANGING"
)

const (
	EventStartDeployment = "StartDeployment"
	EventDeploy          = "Deploy"
	EventTimeout         = "Timeout"
	EventFailed          = "Failure"
	EventRetry           = "Retry"
	EventComplete        = "Completed"
	EventPoll            = "Poll"
	EventWait            = "Wait"
)

type DeployFSM struct {
	fsm        *fsm.FSM
	Deployment *NDFCDeployment
	Events     []fsm.EventDesc
	CallBacks  fsm.Callbacks
	//ndfc              *NDFC
	dg                *diag.Diagnostics
	checkCount        int
	MaxCheckCount     int
	PollTimer         time.Duration
	MaxParallelDeploy int
	TrustFactor       int
	RxCompleteCount   int
	FailureRetry      int
	failCount         int
	eventChannel      chan string
}

func (fsm *DeployFSM) DeploymentHandler(ctx context.Context, e *fsm.Event) {
	tflog.Debug(ctx, "StartDeploymentHandler: Entering")
	// Start the deployment
	// Move maxParallel Rsc from pending to in pro
	// call deploy with the in pro list
	fsm.Deployment.Deploy(ctx, fsm.dg, nil, false)
}

func (fsm *DeployFSM) DeployHandler(ctx context.Context, e *fsm.Event) {
	tflog.Debug(ctx, "DeployFSM: DeployHandler: Entering")
}

func (fsm *DeployFSM) TimeoutHandler(ctx context.Context, e *fsm.Event) {
	tflog.Debug(ctx, "DeployFSM: TimeoutHandler: Entering")
	tflog.Error(ctx, "Deployment Timeout")
	fsm.dg.AddError("Deployment Timeout", "Deployment Timeout")
}

func (fsm *DeployFSM) FailedHandler(ctx context.Context, e *fsm.Event) {
	tflog.Debug(ctx, "DeployFSM: FailedHandler: Entering")
	tflog.Error(ctx, "Deployment Failed")
	fsm.dg.AddError("Deployment Failed", "Deployment Failed")
}

func (fsm *DeployFSM) CompleteHandler(ctx context.Context, e *fsm.Event) {
	tflog.Debug(ctx, "DeployFSM: CompleteHandler: Entering")
}

func (fsm *DeployFSM) PollTimerHandler(ctx context.Context, e *fsm.Event) {
	tflog.Debug(ctx, "DeployFSM: PollTimerHandler: Entering", map[string]interface{}{"event": e.Event, "state": fsm.fsm.Current()})
}

func (fsm *DeployFSM) BatchDeploymentHandler(ctx context.Context, e *fsm.Event) {
	tflog.Debug(ctx, "DeployFSM: RetryDeploymentHandler: Retrying Deployment", map[string]interface{}{"event": e.Event})
	fsm.postEvent(ctx, EventStartDeployment)
	//fsm.fsm.Event(ctx, EventStartDeployment)

}

func (fsm *DeployFSM) StateInfoHandler(ctx context.Context, e *fsm.Event) {
	tflog.Debug(ctx, "DeployFSM: State transition ", map[string]interface{}{"from": e.Src,
		"to": e.Dst, "event": e.Event})
}

func (fsm *DeployFSM) EventInfoHandler(ctx context.Context, e *fsm.Event) {
	tflog.Debug(ctx, "DeployFSM: Event Occurred ", map[string]interface{}{"event": e.Event})

}

func (fsm *DeployFSM) InProHandler(ctx context.Context, e *fsm.Event) {
	tflog.Debug(ctx, "DeployFSM: InProHandler ", map[string]interface{}{"event": e.Event})
	//Deployment started
	// Start a poll timer and Wait

}

func (fsm *DeployFSM) CheckStateHandler(ctx context.Context, e *fsm.Event) {
	fsm.checkCount++
	tflog.Debug(ctx, "DeployFSM: CheckStateHandler ", map[string]interface{}{"CheckCount": fsm.checkCount, "event": e.Event, "MaxChecks": fsm.MaxCheckCount})
	if fsm.checkCount > fsm.MaxCheckCount {
		tflog.Error(ctx, "DeployFSM: CheckStateHandler: Max Check Count Exceeded")
		fsm.postEvent(ctx, EventTimeout)
		//fsm.fsm.Event(ctx, EventTimeout)
		return
	}
	//tflog.Debug(ctx, "DeployFSM: CheckStateHandler", map[string]interface{}{"event": e.Event})

	ret := fsm.Deployment.CheckState(ctx, fsm.dg, nil)
	tflog.Debug(ctx, "DeployFSM: CheckStateHandler: StateChecker completed", map[string]interface{}{"Status": ret})
	fsm.postEvent(ctx, ret)
	//fsm.fsm.Event(ctx, ret)
}

func (fsm *DeployFSM) wait(ctx context.Context) {
	pollTimer := time.NewTimer(fsm.PollTimer)
	select {
	case <-ctx.Done():
		tflog.Error(ctx, "DeployFSM: Context Was done exit")
		fsm.postEvent(ctx, EventFailed)
		//fsm.fsm.Event(ctx, EventFailed)
		return
	case <-pollTimer.C:
		tflog.Debug(ctx, "DeployFSM: PollTimer Expired")
		fsm.postEvent(ctx, EventPoll)
		//fsm.fsm.Event(ctx, EventPoll)
		break
	}
}

func (fsm *DeployFSM) postEvent(ctx context.Context, event string) {
	tflog.Debug(ctx, "DeployFSM: Posting Event", map[string]interface{}{"event": event})
	fsm.eventChannel <- event
}

func (depfsm *DeployFSM) Init() {
	depfsm.CallBacks = fsm.Callbacks{
		EventTimeout:            depfsm.TimeoutHandler,
		StateFailed:             depfsm.FailedHandler,
		StateComplete:           depfsm.CompleteHandler,
		"enter_" + StatePending: depfsm.BatchDeploymentHandler,
		"leave_" + StatePending: depfsm.DeploymentHandler,
		StateInPro:              depfsm.InProHandler,
		StateCheck:              depfsm.CheckStateHandler,
	}

	depfsm.Events = []fsm.EventDesc{
		{Name: EventStartDeployment, Src: []string{StatePending}, Dst: StateInPro},
		{Name: EventDeploy, Src: []string{StateCheck}, Dst: StatePending},
		{Name: EventPoll, Src: []string{StateInPro}, Dst: StateCheck},
		{Name: EventWait, Src: []string{StateCheck}, Dst: StateInPro},
		{Name: EventTimeout, Src: []string{StateCheck}, Dst: StateFailed},
		{Name: EventComplete, Src: []string{StateCheck}, Dst: StateComplete},
		{Name: EventFailed, Src: []string{StateCheck}, Dst: StateFailed},
	}
	depfsm.fsm = fsm.NewFSM(StatePending, depfsm.Events, depfsm.CallBacks)
	depfsm.checkCount = 0
	depfsm.RxCompleteCount = 0
	depfsm.eventChannel = make(chan string, 10)
	noAttachments := len(depfsm.Deployment.DeployRscDB[DeployPending])
	depfsm.MaxCheckCount = noAttachments * 10
	if depfsm.MaxCheckCount < 60 {
		depfsm.MaxCheckCount = 60
	}
	tflog.Debug(context.Background(), "DeployFSM: MaxCheckCount", map[string]interface{}{"MaxCheckCount": depfsm.MaxCheckCount})
}

func NewDeployFSM(ctx context.Context, dg *diag.Diagnostics, d *NDFCDeployment) *DeployFSM {
	pollTimer := time.Second * time.Duration(d.ctrlr.DeployPollTimer)
	fsm := &DeployFSM{
		Deployment: d,
		//ndfc:         c,
		dg:                dg,
		PollTimer:         pollTimer,
		TrustFactor:       d.ctrlr.DeployTrustFactor,
		failCount:         0,
		FailureRetry:      d.ctrlr.FailureRetry,
		MaxParallelDeploy: d.ctrlr.MaxParallelDeploy,
	}

	fsm.Init()
	return fsm
}

func (fsm *DeployFSM) Run(ctx context.Context) {
	tflog.Debug(ctx, "DeployFSM: Starting FSM")

	fsm.postEvent(ctx, EventStartDeployment)

	for loop := true; loop; {
		select {
		case <-ctx.Done():
			tflog.Error(ctx, "DeployFSM: Context Was done exit")
			loop = false
		case event := <-fsm.eventChannel:
			err := fsm.fsm.Event(ctx, event)
			if err != nil {
				tflog.Error(ctx, "DeployFSM: Event error", map[string]interface{}{"event": event, "err": err})
			}
			continue
		default:
		}
		switch fsm.fsm.Current() {
		case StateComplete:
			tflog.Debug(ctx, "DeployFSM: Deployment Completed")
			loop = false
		case StateFailed:
			tflog.Debug(ctx, "DeployFSM: Deployment Failed")
			loop = false
		case StateInPro:
			tflog.Debug(ctx, "DeployFSM: Deployment In Progress")
			fsm.wait(ctx)
		default:
			tflog.Debug(ctx, fmt.Sprintf("DeployFSM: Current state %s", fsm.fsm.Current()))
			log.Panicf("DeployFSM: Invalid State %s", fsm.fsm.Current())

		}
	}

}
