// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package api_test

import (
	"sync"
	"testing"

	"github.com/netascode/go-nd"
	"github.com/stretchr/testify/assert"

	"terraform-provider-ndfc/internal/provider/ndfc/api"
)

func TestPolicyAPI_URLs(t *testing.T) {
	// Test cases for Deploy flag with Policy Group
	t.Run("Deploy with Policy Group", func(t *testing.T) {
		client, _ := nd.NewClient("https://ndfc.example.com", "user", "pass", "", "", false)
		papi := api.NewPolicyAPI(&sync.Mutex{}, &client)
		papi.PolicyID = "POLICY-GROUP-123"
		papi.PolicyGroup = true
		papi.Deploy = true
		papi.DeploySwitches = []string{"SAL1234ABCD", "SAL5678EFGH"}

		// Test PostUrl for deploy with policy group
		expectedUrl := "/lan-fabric/rest/control/policies/deploy?serialNumber=SAL1234ABCD,SAL5678EFGH"
		assert.Equal(t, expectedUrl, papi.PostUrl())

		// Test that other URLs are not affected by deploy flag
		expectedGetUrl := "/lan-fabric/rest/control/policies/policygroup/POLICY-GROUP-123"
		assert.Equal(t, expectedGetUrl, papi.GetUrl())

		expectedPutUrl := "/lan-fabric/rest/control/policies/policygroup/POLICY-GROUP-123?serialNumbers=SAL1234ABCD,SAL5678EFGH&mark-delete-and-update=true"
		assert.Equal(t, expectedPutUrl, papi.PutUrl())

		expectedDeleteUrl := "/lan-fabric/rest/control/policies/policygroup/POLICY-GROUP-123"
		assert.Equal(t, expectedDeleteUrl, papi.DeleteUrl())
	})

	// Test cases for Deploy flag with regular Policy
	t.Run("Deploy with regular Policy", func(t *testing.T) {
		client, _ := nd.NewClient("https://ndfc.example.com", "user", "pass", "", "", false)
		papi := api.NewPolicyAPI(&sync.Mutex{}, &client)
		papi.PolicyID = "POLICY-456"
		papi.Deploy = true

		// Test PostUrl for deploy with regular policy
		expectedUrl := "/lan-fabric/rest/control/policies/deploy"
		assert.Equal(t, expectedUrl, papi.PostUrl())

		// Test that other URLs are not affected by deploy flag
		expectedGetUrl := "/lan-fabric/rest/control/policies/POLICY-456"
		assert.Equal(t, expectedGetUrl, papi.GetUrl())

		expectedPutUrl := "/lan-fabric/rest/control/policies/POLICY-456"
		assert.Equal(t, expectedPutUrl, papi.PutUrl())

		expectedDeleteUrl := "/lan-fabric/rest/control/policies/POLICY-456"
		assert.Equal(t, expectedDeleteUrl, papi.DeleteUrl())
	})

	// Test cases for Policy (non-policy group)
	// Test cases for Policy (non-policy group)
	t.Run("Policy URLs", func(t *testing.T) {
		client, _ := nd.NewClient("https://ndfc.example.com", "user", "pass", "", "", false)
		papi := api.NewPolicyAPI(&sync.Mutex{}, &client)
		papi.PolicyID = "POLICY-123"

		// Test GetUrl
		expectedGetUrl := "/lan-fabric/rest/control/policies/POLICY-123"
		assert.Equal(t, expectedGetUrl, papi.GetUrl())

		// Test PostUrl (create)
		expectedPostUrl := "/lan-fabric/rest/control/policies"
		assert.Equal(t, expectedPostUrl, papi.PostUrl())

		// Test PutUrl (update)
		expectedPutUrl := "/lan-fabric/rest/control/policies/POLICY-123"
		assert.Equal(t, expectedPutUrl, papi.PutUrl())

		// Test DeleteUrl
		expectedDeleteUrl := "/lan-fabric/rest/control/policies/POLICY-123"
		assert.Equal(t, expectedDeleteUrl, papi.DeleteUrl())

		// Test GetDeleteQP (should be nil for non-policy group)
		nilQP := papi.GetDeleteQP()
		assert.Nil(t, nilQP)
	})

	// Test cases for Policy Group
	t.Run("Policy Group URLs", func(t *testing.T) {
		client, _ := nd.NewClient("https://ndfc.example.com", "user", "pass", "", "", false)
		papi := api.NewPolicyAPI(&sync.Mutex{}, &client)
		papi.PolicyID = "POLICY-GROUP-123"
		papi.PolicyGroup = true
		papi.DeploySwitches = []string{"SAL1234ABCD", "SAL5678EFGH"}

		// Test GetUrl for policy group
		expectedGetUrl := "/lan-fabric/rest/control/policies/policygroup/POLICY-GROUP-123"
		assert.Equal(t, expectedGetUrl, papi.GetUrl())

		// Test PostUrl for policy group create with switches
		expectedPostUrl := "/lan-fabric/rest/control/policies/policygroup/create?serialNumbers=SAL1234ABCD,SAL5678EFGH"
		assert.Equal(t, expectedPostUrl, papi.PostUrl())

		// Test PutUrl for policy group update with switches
		expectedPutUrl := "/lan-fabric/rest/control/policies/policygroup/POLICY-GROUP-123?serialNumbers=SAL1234ABCD,SAL5678EFGH&mark-delete-and-update=true"
		assert.Equal(t, expectedPutUrl, papi.PutUrl())

		// Test DeleteUrl for policy group
		expectedDeleteUrl := "/lan-fabric/rest/control/policies/policygroup/POLICY-GROUP-123"
		assert.Equal(t, expectedDeleteUrl, papi.DeleteUrl())

		// Test GetDeleteQP for policy group
		expectedQP := []string{"mark-delete-and-update=true, serialNumbers=SAL1234ABCD,SAL5678EFGH"}
		assert.Equal(t, expectedQP, papi.GetDeleteQP())
	})

	// Test cases for Deploy URL
	t.Run("Deploy URL", func(t *testing.T) {
		client, _ := nd.NewClient("https://ndfc.example.com", "user", "pass", "", "", false)
		papi := api.NewPolicyAPI(&sync.Mutex{}, &client)
		papi.Deploy = true

		// Test PostUrl for deploy
		expectedDeployUrl := "/lan-fabric/rest/control/policies/deploy"
		assert.Equal(t, expectedDeployUrl, papi.PostUrl())
	})

	// Test cases for Policy Group without switches - should panic
	t.Run("Policy Group without switches should panic", func(t *testing.T) {
		client, _ := nd.NewClient("https://ndfc.example.com", "user", "pass", "", "", false)
		papi := api.NewPolicyAPI(&sync.Mutex{}, &client)
		papi.PolicyID = "POLICY-GROUP-456"
		papi.PolicyGroup = true
		papi.DeploySwitches = []string{} // Empty switches

		// Test that PostUrl panics when switches are empty for policy group
		assert.Panics(t, func() {
			papi.PostUrl()
		}, "Expected panic when calling PostUrl with empty switches for policy group")

		// Test that PutUrl panics when switches are empty for policy group
		assert.Panics(t, func() {
			papi.PutUrl()
		}, "Expected panic when calling PutUrl with empty switches for policy group")

		// Test that GetDeleteQP panics when switches are empty for policy group
		assert.Panics(t, func() {
			papi.GetDeleteQP()
		}, "Expected panic when calling GetDeleteQP with empty switches for policy group")
	})
}

func TestPolicyAPI_RscName(t *testing.T) {
	client, _ := nd.NewClient("https://ndfc.example.com", "user", "pass", "", "", false)
	papi := api.NewPolicyAPI(&sync.Mutex{}, &client)

	// Test RscName
	expectedRscName := "policy"
	assert.Equal(t, expectedRscName, papi.RscName())
}
