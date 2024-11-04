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
	"strings"
	"sync"
	"terraform-provider-ndfc/tfutils/go-nd"
)

type NetAttachAPI struct {
	NDFCAPICommon
	mutex      *sync.Mutex
	fabricName string
	GetnwList  []string
	Payload    string
}

const UrlNetAttachGetBulk = "/lan-fabric/rest/top-down/v2/fabrics/%s/networks/attachments?network-names=%s"
const UrlNetAttachBulk = "/lan-fabric/rest/top-down/v2/fabrics/%s/networks/attachments"
const UrlNetAttachGet = "lan-fabric/rest/top-down/v2/fabrics/%s/networks/attachments"
const UrlNetAttachUpdate = "/lan-fabric/rest/top-down/v2/fabrics/%s/networks/%s"

func (c *NetAttachAPI) GetLock() *sync.Mutex {
	return c.mutex
}

func (c *NetAttachAPI) GetUrl() string {
	url := fmt.Sprintf(UrlNetAttachGetBulk, c.fabricName, strings.Join(c.GetnwList, ","))
	return url
}

func (c *NetAttachAPI) PostUrl() string {
	return fmt.Sprintf(UrlNetAttachBulk, c.fabricName)
}

func (c *NetAttachAPI) PutUrl() string {
	return fmt.Sprintf(UrlNetAttachBulk, c.fabricName)
}

func (c *NetAttachAPI) DeleteUrl() string {
	return ""
}

func (c NetAttachAPI) GetDeleteQP() []string {
	return []string{}
}

func (c *NetAttachAPI) SetDeleteList(qp []string) {

}

func NewNetAttachAPI(fabricName string, lock *sync.Mutex, client *nd.Client) *NetAttachAPI {
	api := new(NetAttachAPI)
	api.fabricName = fabricName
	api.mutex = lock
	api.NDFCAPI = api
	api.client = client
	return api
}
