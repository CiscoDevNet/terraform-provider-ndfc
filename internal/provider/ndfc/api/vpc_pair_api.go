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
	"sync"
	"terraform-provider-ndfc/tfutils/go-nd"
)

// VpcPairAPI is the API client for the vpc pair resource
const urlVpcPair = "/lan-fabric/rest/vpcpair"
const urlVpcPairGet = urlVpcPair + "?serialNumber=%s"
const urlVpcPairRecmd = urlVpcPair + "/recommendation?serialNumber=%s&useVirtualPeerlink=%t"
const urlswitchesByFabric = "/lan-fabric/rest/control/fabrics/%s/inventory/switchesByFabric"

type VpcPairAPI struct {
	NDFCAPICommon
	mutex           *sync.Mutex
	CheckStatus     map[string]bool
	FabricName      string
	VirtualPeerLink bool
	VpcPairID       string
}

func (c *VpcPairAPI) GetLock() *sync.Mutex {
	return c.mutex
}

func (c *VpcPairAPI) GetUrl() string {
	if c.CheckStatus["recommendations"] {
		return fmt.Sprintf(urlVpcPairRecmd, c.VpcPairID, c.VirtualPeerLink)
	} else if c.CheckStatus["switchesByFabric"] {
		return fmt.Sprintf(urlswitchesByFabric, c.FabricName)
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
