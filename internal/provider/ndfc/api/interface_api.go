// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package api

import (
	"log"
	"sync"

	"github.com/netascode/go-nd"
)

type InterfaceAPI struct {
	NDFCAPICommon
	SwitchSerial string
	Policy       string
	IfTypes      string
	PortMode     string
	Excludes     string
	mutex        *sync.Mutex
	deployFlag   bool
	APIFunction  int
}

const UrlInterface = "/lan-fabric/rest/interface"
const UrlInterfaceDeploy = "/lan-fabric/rest/interface/deploy"
const UrlGetInterfaceDetailed = "/lan-fabric/rest/interface/detail/filter"
const UrlInterfaceGlobal = "/lan-fabric/rest/globalInterface"

const (
	GetInterface = iota
	GetInterfaceDetailed
	PuttInterfaceEthernet
	PostInterfaceCreate
	PostInterfaceDeploy
)

func (c *InterfaceAPI) GetLock() *sync.Mutex {
	return c.mutex
}

func (c *InterfaceAPI) GetUrl() string {
	switch c.APIFunction {
	case GetInterfaceDetailed:
		url := UrlGetInterfaceDetailed
		if c.SwitchSerial == "" {
			log.Fatalf("SwitchSerial is required for GetInterfaceDetailed")
			return ""
		}
		url += "?serialNumber=" + c.SwitchSerial
		if c.PortMode != "" {
			url += "&portModes=" + c.PortMode
		}
		if c.Excludes != "" {
			url += "&excludes=" + c.Excludes
		}
		if c.IfTypes != "" {
			url += "&ifTypes=" + c.IfTypes
		}
		return url
	case GetInterface:
		fallthrough
	default:
		if c.SwitchSerial != "" && c.Policy != "" {
			return (UrlInterface + "?serialNumber=" + c.SwitchSerial + "&templateName=" + c.Policy)
		}
		if c.SwitchSerial != "" {
			return (UrlInterface + "?serialNumber=" + c.SwitchSerial)
		}
		return UrlInterface
	}
}

func (c *InterfaceAPI) PostUrl() string {
	switch c.APIFunction {
	case PostInterfaceCreate:
		return UrlInterfaceGlobal
	case PostInterfaceDeploy:
		return UrlInterfaceDeploy
	default:
		return UrlInterface
	}
}

func (c *InterfaceAPI) PutUrl() string {
	return UrlInterface
}

func (c *InterfaceAPI) DeleteUrl() string {
	return UrlInterface
}

func (c *InterfaceAPI) GetDeleteQP() []string {
	return nil
}

func (c *InterfaceAPI) SetAPI(function int) {
	c.APIFunction = function
}

func (c *InterfaceAPI) RscName() string {
	return "interface"
}

func NewInterfaceAPI(lock *sync.Mutex, client *nd.Client) *InterfaceAPI {
	api := new(InterfaceAPI)
	api.mutex = lock
	api.NDFCAPI = api
	api.client = client
	return api
}
