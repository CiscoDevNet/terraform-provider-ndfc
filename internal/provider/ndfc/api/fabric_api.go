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

// FabricAPI is the API client for the vpc pair resource
const (
	UrlFabricGetAll         = "/lan-fabric/rest/control/fabrics"
	urlPerFabric            = UrlFabricGetAll + "/%s"
	urlFabricTemplate       = urlPerFabric + "/%s"
	UrlSwitchesByFabric     = "/lan-fabric/rest/control/fabrics/%s/inventory/switchesByFabric"
	urlFabricNameFromSerial = "/lan-fabric/rest/control/switches/%s/fabric-name"
)

type FabricAPI struct {
	NDFCAPICommon
	mutex               *sync.Mutex
	FabricName          string
	GetSwitchesInFabric bool
	Serialnumber        string
	FabricType          string
}

func (c *FabricAPI) GetLock() *sync.Mutex {
	return c.mutex
}

func (c *FabricAPI) GetUrl() string {
	if c.Serialnumber != "" {
		return fmt.Sprintf(urlFabricNameFromSerial, c.Serialnumber)
	} else if c.FabricName != "" {
		if c.GetSwitchesInFabric {
			return fmt.Sprintf(UrlSwitchesByFabric, c.FabricName)
		} else {
			return fmt.Sprintf(urlPerFabric, c.FabricName)
		}
	} else {
		return UrlFabricGetAll
	}
}

func (c *FabricAPI) PostUrl() string {
	return fmt.Sprintf(urlFabricTemplate, c.FabricName, c.FabricType)
}

func (c *FabricAPI) PutUrl() string {
	return fmt.Sprintf(urlFabricTemplate, c.FabricName, c.FabricType)
}

func (c *FabricAPI) DeleteUrl() string {
	return fmt.Sprintf(urlPerFabric, c.FabricName)
}
func (c *FabricAPI) GetDeleteQP() []string {
	return nil
}

func (c *FabricAPI) RscName() string {
	return "fabric"
}

func NewFabricAPI(lock *sync.Mutex, c *nd.Client) *FabricAPI {
	papi := new(FabricAPI)
	papi.mutex = lock
	papi.client = c
	papi.NDFCAPI = papi
	return papi
}
