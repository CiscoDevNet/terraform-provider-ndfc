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

// PolicyAPI is the API client for the policy resource
const UrlPolicyCreate = "/lan-fabric/rest/control/policies"
const UrlPolicy = "/lan-fabric/rest/control/policies/%s"
const UrlPolicyDeploy = "/lan-fabric/rest/control/policies/deploy"

const UrlPolicyGroup = "/lan-fabric/rest/control/policies/policygroup/%s"
const UrlPolicyGroupCreate = "lan-fabric/rest/control/policies/policygroup/create"

type PolicyAPI struct {
	NDFCAPICommon
	mutex          *sync.Mutex
	Deploy         bool
	PolicyGroup    bool
	DeploySwitches []string
	PolicyID       string
}

func (c *PolicyAPI) GetLock() *sync.Mutex {
	return c.mutex
}

func (c *PolicyAPI) GetUrl() string {
	return fmt.Sprintf(UrlPolicy, c.PolicyID)
}

func (c *PolicyAPI) PostUrl() string {
	if c.Deploy {
		return UrlPolicyDeploy
	}
	return UrlPolicyCreate
}

func (c *PolicyAPI) PutUrl() string {
	return fmt.Sprintf(UrlPolicy, c.PolicyID)
}

func (c *PolicyAPI) DeleteUrl() string {
	return fmt.Sprintf(UrlPolicy, c.PolicyID)
}

func (c *PolicyAPI) GetDeleteQP() []string {
	return nil
}

func (c *PolicyAPI) RscName() string {
	return "policy"
}


func NewPolicyAPI(lock *sync.Mutex, c *nd.Client) *PolicyAPI {
	papi := new(PolicyAPI)
	papi.mutex = lock
	papi.client = c
	papi.NDFCAPI = papi
	return papi
}
