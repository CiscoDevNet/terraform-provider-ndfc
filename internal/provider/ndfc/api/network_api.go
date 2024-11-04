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
	"terraform-provider-ndfc/tfutils/go-nd"
)

type NetworkAPI struct {
	NDFCAPICommon
	mutex          *sync.Mutex
	fabricName     string
	PutNetworkName string
	Payload        string
	DelList        []string
	GetNetworkName string
}

const UrlNetworkGetBulk = "/lan-fabric/rest/top-down/v2/fabrics/%s/networks"
const UrlNetworkCreateBulk = "/lan-fabric/rest/top-down/v2/bulk-create/networks"
const UrlNetworkDeleteBulk = "/lan-fabric/rest/top-down/v2/fabrics/%s/bulk-delete/networks"
const UrlNetworkGet = "/lan-fabric/rest/top-down/v2/fabrics/%s/networks/%s"
const UrlNetworkUpdate = "/lan-fabric/rest/top-down/v2/fabrics/%s/networks/%s"

// For additional functions
const UrlNetworkGenMulticastIP = "/lan-fabric/rest/top-down/v2/fabrics/%s/generateMulticastIp"
const UrlNetworkGenVNI = "/lan-fabric/rest/top-down/v2/fabrics/%s/netinfo"
const UrlGetFreeVlanId = "/lan-fabric/rest/resource-manager/vlan/%s"

func (c *NetworkAPI) GetLock() *sync.Mutex {
	return c.mutex
}

func (c *NetworkAPI) GetUrl() string {
	log.Printf("GetUrl - NetworkAPI |%v|", c.GetNetworkName)
	if c.GetNetworkName == "" {
		return fmt.Sprintf(UrlNetworkGetBulk, c.fabricName)
	}
	url := fmt.Sprintf(UrlNetworkGet, c.fabricName, c.GetNetworkName)
	//Read clear
	c.GetNetworkName = ""
	return url
}

func (c *NetworkAPI) PostUrl() string {
	return UrlNetworkCreateBulk
}

func (c *NetworkAPI) PutUrl() string {
	return fmt.Sprintf(UrlNetworkUpdate, c.fabricName, c.PutNetworkName)
}

func (c *NetworkAPI) DeleteUrl() string {
	return fmt.Sprintf(UrlNetworkDeleteBulk, c.fabricName)
}

/*
 * Response format
 *	{
 *	    "multicast-ip": "239.1.1.2"
 *	}
 */
func (c NetworkAPI) GenMcastIP() string {
	url := fmt.Sprintf(UrlNetworkGenMulticastIP, c.fabricName)
	c.GetLock().Lock()
	defer c.GetLock().Unlock()
	res, err := c.client.Get(url)
	if err != nil {
		return ""
	}
	return res.Get("multicast-ip").String()
}

/*
 * Response format
 * {
    "mcastip": "239.1.1.1",
    "l2vni": 30002,
    "network-prefix": "MyNetwork_"
	}
 *
*/

func (c NetworkAPI) GenVNI() string {
	url := fmt.Sprintf(UrlNetworkGenVNI, c.fabricName)
	c.GetLock().Lock()
	defer c.GetLock().Unlock()
	res, err := c.client.Get(url)
	if err != nil {
		return ""
	}
	return res.Get("l2vni").String()
}

func (c NetworkAPI) GetDeleteQP() []string {
	log.Printf("Returning delete list %v", c.DelList)
	return []string{"network-names=" + strings.Join(c.DelList, ",")}
}

func (c *NetworkAPI) SetDeleteList(qp []string) {
	log.Printf("Setting delete list %v", qp)
	c.DelList = make([]string, len(qp))
	copy(c.DelList, qp)
	log.Printf("Copied delete list %v", c.DelList)

}

func NewNetworksAPI(fabricName string, lock *sync.Mutex, client *nd.Client) *NetworkAPI {
	api := new(NetworkAPI)
	api.fabricName = fabricName
	api.mutex = lock
	api.NDFCAPI = api
	api.client = client
	return api
}
