package api

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"terraform-provider-ndfc/tfutils/go-nd"
)

type VrfAPI struct {
	NDFCAPICommon
	fabricName string
	mutex      *sync.Mutex
	PutVrf     string
	Payload    string
	DelList    []string
}

const UrlVrfGetBulk = "/lan-fabric/rest/top-down/v2/fabrics/%s/vrfs"
const UrlVrfCreateBulk = "/lan-fabric/rest/top-down/v2/bulk-create/vrfs"
const UrlVrfDeleteBulk = "/lan-fabric/rest/top-down/v2/fabrics/%s/bulk-delete/vrfs"
const UrlVrfGet = "/lan-fabric/rest/top-down/v2/fabrics/%s/vrf/%s"
const UrlVrfUpdate = "/lan-fabric/rest/top-down/v2/fabrics/%s/vrfs/%s"

func (c *VrfAPI) GetLock() *sync.Mutex {
	log.Printf("GetLock - VrfAPI %v", c.mutex)
	return c.mutex
}

func (c *VrfAPI) GetUrl() string {
	log.Printf("GetUrl - VrfAPI")
	return fmt.Sprintf(UrlVrfGetBulk, c.fabricName)
}

func (c *VrfAPI) PostUrl() string {
	log.Printf("PostUrl - VrfAPI")
	return UrlVrfCreateBulk
}

func (c *VrfAPI) PutUrl() string {
	log.Printf("PutUrl - VrfAPI")
	return fmt.Sprintf(UrlVrfUpdate, c.fabricName, c.PutVrf)
}

func (c *VrfAPI) DeleteUrl() string {
	log.Printf("DeleteUrl - VrfAPI")
	return fmt.Sprintf(UrlVrfDeleteBulk, c.fabricName)
}

func (c *VrfAPI) SetDeleteList(qp []string) {
	log.Printf("SetDeleteList - VrfAPI")
	c.DelList = make([]string, len(qp))
	copy(c.DelList, qp)
}

func (c *VrfAPI) GetDeleteQP() []string {
	log.Printf("GetDeleteQP - VrfAPI")
	return []string{"vrf-names=" + strings.Join(c.DelList, ",")}
}

func NewVrfAPI(fabricName string, lock *sync.Mutex, client *nd.Client) *VrfAPI {
	api := new(VrfAPI)
	api.fabricName = fabricName
	api.mutex = lock
	api.NDFCAPI = api
	api.client = client
	return api
}
