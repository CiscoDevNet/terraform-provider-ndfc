package api

import (
	"fmt"
	"sync"
	"terraform-provider-ndfc/tfutils/go-nd"
)

type DeploymentAPI struct {
	NDFCAPICommon
	mutex         *sync.Mutex
	fabricName    string
	AttachPayload string
	RsType        string
}

// For additional functions

const UrlVrfNetAttachmentsDeploy = "/lan-fabric/rest/top-down/v2/%s/deploy"
const UrlVrfNetworkDeployment = "/lan-fabric/rest/top-down/v2/fabrics/%s/%s/deployments"

func (c *DeploymentAPI) GetLock() *sync.Mutex {
	return c.mutex
}

func (c *DeploymentAPI) GetUrl() string {
	panic("Get Not supported")
}

func (c *DeploymentAPI) PostUrl() string {
	return fmt.Sprintf(UrlVrfNetAttachmentsDeploy, c.RsType)
}

func (c *DeploymentAPI) PutUrl() string {
	panic("PUT Not supported")
}

func (c *DeploymentAPI) DeleteUrl() string {
	panic("DELETE Not supported")
}

func (c DeploymentAPI) GetDeleteQP() []string {
	panic("Not supported")
}

func (c *DeploymentAPI) SetDeleteList(qp []string) {
	panic("Not supported")
}

func NewDeploymentAPI(fabricName string, lock *sync.Mutex, client *nd.Client, rsType string) *DeploymentAPI {
	api := new(DeploymentAPI)
	api.fabricName = fabricName
	api.mutex = lock
	api.NDFCAPI = api
	api.client = client
	api.RsType = rsType
	return api
}
