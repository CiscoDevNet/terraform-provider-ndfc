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

// PolicyAPI is the API client for the policy resource
const UrlPolicyCreate = "/lan-fabric/rest/control/policies"
const UrlPolicy = "/lan-fabric/rest/control/policies/%s"
const UrlPolicyDeploy = "/lan-fabric/rest/control/policies/deploy"

const UrlPolicyGroup = "/lan-fabric/rest/control/policies/policygroup/%s"
const UrlPolicyGroupCreate = "/lan-fabric/rest/control/policies/policygroup/create?serialNumbers=%s"
const UrlPolicyGroupUpdate = "/lan-fabric/rest/control/policies/policygroup/%s?serialNumbers=%s&mark-delete-and-update=true"
const UrlPolicyGroupDelete = "/lan-fabric/rest/control/policies/policygroup/policyIds"

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
	if c.PolicyGroup {
		return fmt.Sprintf(UrlPolicyGroup, c.PolicyID)
	}
	return fmt.Sprintf(UrlPolicy, c.PolicyID)
}

func (c *PolicyAPI) PostUrl() string {
	if c.Deploy {
		if c.PolicyGroup {
			if len(c.DeploySwitches) != 0 {
				return fmt.Sprintf("%s?serialNumber=%s", UrlPolicyDeploy, strings.Join(c.DeploySwitches, ","))
			} else {
				panic("Switches cannot be empty")
			}
		}
		return UrlPolicyDeploy
	}
	if c.PolicyGroup {
		if len(c.DeploySwitches) != 0 {
			return fmt.Sprintf(UrlPolicyGroupCreate, strings.Join(c.DeploySwitches, ","))
		} else {
			panic("Switches cannot be empty")
		}
	}
	return UrlPolicyCreate
}

func (c *PolicyAPI) PutUrl() string {
	if c.PolicyGroup {
		if len(c.DeploySwitches) != 0 {
			return fmt.Sprintf(UrlPolicyGroupUpdate, c.PolicyID, strings.Join(c.DeploySwitches, ","))
		} else {
			panic("Switches cannot be empty")
		}
	}
	return fmt.Sprintf(UrlPolicy, c.PolicyID)
}

func (c *PolicyAPI) DeleteUrl() string {
	if c.PolicyGroup {
		return UrlPolicyGroupDelete
	}
	return fmt.Sprintf(UrlPolicy, c.PolicyID)
}

func (c *PolicyAPI) GetDeleteQP() []string {
	if c.PolicyGroup {
		if len(c.DeploySwitches) != 0 {
			return []string{"policyIds=" + c.PolicyID, "serialNumbers=" + strings.Join(c.DeploySwitches, ",")}
		} else {
			panic("Switches cannot be empty")
		}
	}
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
