package api

import (
	"fmt"
	"sync"
	"terraform-provider-ndfc/tfutils/go-nd"
)

// VpcPairAPI is the API client for the vpc pair resource
const urlVpcPair = "/lan-fabric/rest/vpcpair"
const urlVpcPairGet = urlVpcPair + "?serialNumber=%s"
const urlVpcPairRecmd = urlVpcPair + "/recommendation?serialNumber=%s&useVirtualPeerlink=%t"

type VpcPairAPI struct {
	NDFCAPICommon
	mutex                *sync.Mutex
	CheckRecommendations bool
	FabricName           string
	VirtualPeerLink      bool
	VpcPairID            string
}

func (c *VpcPairAPI) GetLock() *sync.Mutex {
	return c.mutex
}

func (c *VpcPairAPI) GetUrl() string {
	if c.CheckRecommendations {
		return fmt.Sprintf(urlVpcPairRecmd, c.VpcPairID, c.VirtualPeerLink)
	} else {
		return fmt.Sprintf(urlVpcPairGet, c.VpcPairID)
	}
}

func (c *VpcPairAPI) PostUrl() string {
	return urlVpcPair
}

func (c *VpcPairAPI) PutUrl() string {
	return urlVpcPair
}

func (c *VpcPairAPI) DeleteUrl() string {
	url := urlVpcPair
	url += "?serialNumber=" + c.VpcPairID
	return url
}
func (c *VpcPairAPI) GetDeleteQP() []string {
	return nil
}

func NewVpcPairAPI(lock *sync.Mutex, c *nd.Client) *VpcPairAPI {
	papi := new(VpcPairAPI)
	papi.mutex = lock
	papi.client = c
	papi.NDFCAPI = papi
	return papi
}
