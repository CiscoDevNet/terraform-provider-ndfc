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

	"github.com/netascode/go-nd"
)

type ConfigDeploymentAPI struct {
	NDFCAPICommon
	mutex         *sync.Mutex
	SerialNumbers []string
	Deploy        bool
	Preview       bool
	FabricName    string
}

// For additional functions

const UrlGlobalConfigDeploy = "/lan-fabric/rest/control/fabrics/%s/config-deploy?forceShowRun=false"
const UrlSwitchConfigDeploy = "/lan-fabric/rest/control/fabrics/%s/config-deploy/%s?forceShowRun=false"
const UrlSaveConfig = "/lan-fabric/rest/control/fabrics/%s/config-save"
const UrlGetFabricErrors = "/lan-fabric/rest/control/fabrics/%s/errors"
const UrlGetGlobalConfigPreview = "/lan-fabric/rest/control/fabrics/%s/config-preview/%s"
const UrlGetConfigPreview = "/lan-fabric/rest/control/fabrics/%s/config-preview"

func (c *ConfigDeploymentAPI) GetLock() *sync.Mutex {
	return c.mutex
}

func (c *ConfigDeploymentAPI) GetUrl() string {
	if c.Preview {
		if len(c.SerialNumbers) > 0 {
			return fmt.Sprintf(UrlGetGlobalConfigPreview, c.FabricName, strings.Join(c.SerialNumbers, ","))
		}
		return fmt.Sprintf(UrlGetConfigPreview, c.FabricName)
	} else {
		return fmt.Sprintf(UrlGetFabricErrors, c.FabricName)
	}
}

func (c *ConfigDeploymentAPI) PostUrl() string {
	if c.Deploy {
		if len(c.SerialNumbers) > 0 {
			return fmt.Sprintf(UrlSwitchConfigDeploy, c.FabricName, strings.Join(c.SerialNumbers, ","))
		} else {
			return fmt.Sprintf(UrlGlobalConfigDeploy, c.FabricName)
		}
	} else {
		return fmt.Sprintf(UrlSaveConfig, c.FabricName)
	}
}

func (c *ConfigDeploymentAPI) PutUrl() string {
	panic("PUT Not supported")
}

func (c *ConfigDeploymentAPI) DeleteUrl() string {
	panic("DELETE Not supported")
}

func (c ConfigDeploymentAPI) GetDeleteQP() []string {
	panic("Not supported")
}

func (c *ConfigDeploymentAPI) SetDeleteList(qp []string) {
	panic("Not supported")
}

func (c *ConfigDeploymentAPI) RscName() string {
	return "config-deploy"
}

func NewConfigDeploymentAPI(lock *sync.Mutex, client *nd.Client) *ConfigDeploymentAPI {
	api := new(ConfigDeploymentAPI)
	api.mutex = lock
	api.NDFCAPI = api
	api.client = client
	return api
}
