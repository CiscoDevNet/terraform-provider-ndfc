// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package api

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/netascode/go-nd"
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
const UrlVrfGet = "/lan-fabric/rest/top-down/v2/fabrics/%s/vrfs/%s"
const UrlVrfUpdate = "/lan-fabric/rest/top-down/v2/fabrics/%s/vrfs/%s"
const UrlVrfAttachmentsGet = "/lan-fabric/rest/top-down/fabrics/%s/vrfs/attachments"

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

func (c *VrfAPI) RscName() string {
	return "vrf"
}

/*
	func (c *VrfAPI) GetSingleVRF(fabricName, vrfName string) ([]byte, error) {
		log.Printf("GetSingleVRF - VrfAPI")
		return c.client.GetRawJson(fmt.Sprintf(UrlVrfGet, fabricName, vrfName))
	}

	func (c *VrfAPI) GetVrfAttachments(fabricName, vrfName string) ([]byte, error) {
		log.Printf("GetVrfAttachment - VrfAPI")
		url := fmt.Sprintf(UrlVrfAttachmentsGet, fabricName)
		url += "?vrf-names=" + vrfName
		return c.client.GetRawJson(url)
	}
*/
func NewVrfAPI(fabricName string, lock *sync.Mutex, client *nd.Client) *VrfAPI {
	api := new(VrfAPI)
	api.fabricName = fabricName
	api.mutex = lock
	api.NDFCAPI = api
	api.client = client
	return api
}
