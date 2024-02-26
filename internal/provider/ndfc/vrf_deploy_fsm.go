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
)

const (
	EventStartDeployment = "StartDeployment"
	EventDeployed        = "Deploy"
	EventTimeout         = "Timeout"
	EventFailed          = "Failure"
	EventRetry           = "Retry"
	EventComplete        = "Completed"
	EventPoll            = "Poll"
	EventWait            = "Wait"
)

type DeployFSM struct {
	fsm             *fsm.FSM
	Deployment      *NDFCVrfDeployment
	Events          []fsm.EventDesc
	CallBacks       fsm.Callbacks
	ndfc            *NDFC
	dg              *diag.Diagnostics
	checkCount      int
	MaxCheckCount   int
	PollTimer       time.Duration
	TrustFactor     int
	RxCompleteCount int
	FailureRetry    int
	failCount       int
	eventChannel    chan string
}

func (fsm *DeployFSM) StartDeploymentHandler(ctx context.Context, e *fsm.Event) {
	tflog.Debug(ctx, "StartDeploymentHandler: Entering")
	// Start the deployment
	err := fsm.ndfc.DeployBulk(ctx, fsm.dg, fsm.Deployment)
	if err != nil {
		tflog.Error(ctx, "StartDeploymentHandler: Error in deployment", map[string]interface{}{"error": err.Error()})
		//fsm.fsm.Event(ctx, EventFailed)
	}
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

func (fsm *DeployFSM) RetryDeploymentHandler(ctx context.Context, e *fsm.Event) {
	tflog.Debug(ctx, "DeployFSM: RetryDeploymentHandler: Retrying Deployment", map[string]interface{}{"event": e.Event})
	// ReDo the deployment for failed items
	if len(fsm.Deployment.RetryList) > 0 {
		tflog.Debug(ctx, "DeployFSM: RetryDeploymentHandler: Retrylist is non-empty", map[string]interface{}{"Retrylist": fsm.Deployment.RetryList})
		fsm.Deployment.retryFlag = true
	}

	fsm.postEvent(ctx, EventStartDeployment)
	//fsm.fsm.Event(ctx, EventStartDeployment)

}

func (fsm *DeployFSM) StateInfoHandler(ctx context.Context, e *fsm.Event) {
	tflog.Debug(ctx, "DeployFSM: State transition ", map[string]interface{}{"from": e.Src,
		"to": e.Dst, "event": e.Event})
}

func (fsm *DeployFSM) EventInfoHandler(ctx context.Context, e *fsm.Event) {
	tflog.Debug(ctx, "DeployFSM: Event Occured ", map[string]interface{}{"event": e.Event})

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
	ret := fsm.ndfc.stateChecker(ctx, fsm.dg, fsm.Deployment)
	if ret == EventComplete {
		if fsm.RxCompleteCount < fsm.TrustFactor {
			fsm.RxCompleteCount++
			ret = EventWait
		}
	} else if ret == EventFailed {
		fsm.failCount++
		// NDFC states are not trustworthy
		// Its important to retry
		if fsm.failCount < fsm.FailureRetry {
			tflog.Debug(ctx, "DeployFSM: CheckStateHandler: Failure - retrying", map[string]interface{}{"Status": ret})
			ret = EventRetry
		}
	}
	tflog.Debug(ctx, "DeployFSM: CheckStateHandler: StateChecker completed", map[string]interface{}{"Status": ret})
	fsm.postEvent(ctx, ret)
	//fsm.fsm.Event(ctx, ret)
}

func (fsm *DeployFSM) wait(ctx context.Context) {
	pollTimer := time.NewTimer(fsm.PollTimer * time.Second)
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
	fsm.eventChannel <- event
}

func (depfsm *DeployFSM) Init() {
	depfsm.CallBacks = fsm.Callbacks{
		EventTimeout:            depfsm.TimeoutHandler,
		StateFailed:             depfsm.FailedHandler,
		StateComplete:           depfsm.CompleteHandler,
		"enter_" + StatePending: depfsm.RetryDeploymentHandler,
		"leave_" + StatePending: depfsm.StartDeploymentHandler,
		StateInPro:              depfsm.InProHandler,
		StateCheck:              depfsm.CheckStateHandler,
	}

	depfsm.Events = []fsm.EventDesc{
		{Name: EventStartDeployment, Src: []string{StatePending}, Dst: StateInPro},
		{Name: EventRetry, Src: []string{StateCheck}, Dst: StatePending},
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
	noAttachments := len(depfsm.Deployment.ExpectedState)
	depfsm.MaxCheckCount = noAttachments * 10
	if depfsm.MaxCheckCount < 60 {
		depfsm.MaxCheckCount = 60
	}
	tflog.Debug(context.Background(), "DeployFSM: MaxCheckCount", map[string]interface{}{"MaxCheckCount": depfsm.MaxCheckCount})
}

func (c *NDFC) CreateFSM(ctx context.Context, dg *diag.Diagnostics, d *NDFCVrfDeployment) *DeployFSM {
	fsm := &DeployFSM{
		Deployment:   d,
		ndfc:         c,
		dg:           dg,
		PollTimer:    5,
		TrustFactor:  5,
		failCount:    0,
		FailureRetry: 3,
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
			fsm.fsm.Event(ctx, event)
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
