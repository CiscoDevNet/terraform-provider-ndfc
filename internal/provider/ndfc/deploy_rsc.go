package ndfc

type DeployRsc interface {
	GetKey() string
	GetPostPayload() string
	GetGetQP() string
	GetExpectedState() string
	GetCurrentState() string
	GetFailureCount() int
	GetCheckTick() int

	//SetCurrentState(string)
}

type NDFCDeployRsc struct {
	DeployRsc
	Key          string
	PostPayload  string
	GetQP        string
	State        DeploymentState
	checkTick    int
	failCnt      int
	checkCount   int
}

func (n NDFCDeployRsc) GetKey() string {
	return n.Key
}

func (n NDFCDeployRsc) GetPostPayload() string {
	if n.PostPayload != "" {
		return n.PostPayload
	}
	n.PostPayload = n.DeployRsc.GetPostPayload()
	return n.PostPayload
}

// http get query parameter
func (n NDFCDeployRsc) GetGetQP() string {
	return n.GetQP
}

func (n NDFCDeployRsc) GetExpectedState() string {
	return n.State.ExpectedState
}

func (n *NDFCDeployRsc) SetCurrentState(state string) {
	if state == NDFCStateFailed {
		n.failCnt++
	}
	n.State.CurrentState = state
}

func (n NDFCDeployRsc) GetCurrentState() string {
	return n.State.CurrentState
}

func (n NDFCDeployRsc) GetFailureCount() int {
	return n.failCnt
}

func (n NDFCDeployRsc) GetCheckTick() int {
	return n.checkTick
}



