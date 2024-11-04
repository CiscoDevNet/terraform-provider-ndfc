// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package ndfc

import (
	"log"
	"sync"
	"terraform-provider-ndfc/tfutils/go-nd"
	"time"
)

type NDFC struct {
	url                   string
	apiClient             nd.Client
	DeployPollTimer       int
	DeployTrustFactor     int
	MaxParallelDeploy     int
	FailureRetry          int
	rscMutex              map[string]*sync.Mutex
	WaitForDeployComplete bool
}

var instance *NDFC

func NewNDFCClient(host string, user string, pass string, domain string, insecure bool, timeout int64) (*NDFC, error) {
	log.Printf("New NDFC client")
	ndfc := new(NDFC)
	ndfc.url = "/appcenter/cisco/ndfc/api/v1"
	ndfc.DeployPollTimer = 5
	ndfc.DeployTrustFactor = 5
	ndfc.FailureRetry = 3
	ndfc.MaxParallelDeploy = 0
	ndfc.WaitForDeployComplete = true
	var err error
	ndfc.apiClient, err = nd.NewClient(host, ndfc.url, user, pass, domain, insecure, nd.MaxRetries(500), nd.RequestTimeout(time.Duration(timeout)))
	if err != nil {
		return nil, err
	}
	ndfc.rscMutex = make(map[string]*sync.Mutex)

	log.Printf("[DEBUG] Authentication during creation of NewNDFCClient")
	err = ndfc.apiClient.Authenticate()
	if err != nil {
		log.Printf("[DEBUG] Authentication failed during creation of NewNDFCClient")
		return nil, err
	}
	log.Printf("[DEBUG] Authentication succesful during creation of NewNDFCClient with token: %s", ndfc.apiClient.Token)

	instance = ndfc
	return ndfc, nil
}

func NewResource(rscName string) {
	log.Printf("New Resource %s\n", rscName)
	if _, ok := instance.rscMutex[rscName]; !ok {
		instance.rscMutex[rscName] = new(sync.Mutex)
	}
}

func (c NDFC) GetLock(rscName string) *sync.Mutex {
	return (instance.rscMutex[rscName])
}

func GetLock(rscName string) *sync.Mutex {
	return instance.GetLock(rscName)
}
