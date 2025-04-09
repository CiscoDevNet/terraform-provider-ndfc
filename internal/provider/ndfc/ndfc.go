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
	"os"
	"sync"
	"terraform-provider-ndfc/internal/provider/ndfc/api"
	"time"

	"github.com/netascode/go-nd"
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
	deployMutex           *sync.RWMutex
}

const ResourceRestAPI = "rest_api"
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
	ndfc.deployMutex = new(sync.RWMutex)
	api.SetLockFns(GlobalDeployTrylock, []func(string){AcquireResourceLock, ReleaseResourceLock})
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
	log.Printf("[DEBUG] Authentication successful during creation of NewNDFCClient with token: %s", ndfc.apiClient.Token)

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

/*
When global switch level deployment or resource level deplotment is started, if other resources are being configured
at the same time, NDFC moves to out-of-sync state sometimes
To avoid this, all resource create/update must stop until deployment is complete
To achieve this, a readers-writer lock is used
All resources before Create/Update operaton must acquire a read lock
Global/Resource level Deployment operation must acquire a write lock
This ensures that other resources can run in parallel while write lock(depoyment not happening) is available
Deployment (writer lock) waits for all resource create/update (readers lock) to exit before starting deployment
Once deployment is complete, it releases the write lock
*/
func GlobalDeployLock(who string) {
	log.Printf("**********Waiting for GlobalDeployLock Lock for %s deployment by pid %d***************", who, os.Getpid())
	instance.deployMutex.Lock()
	log.Printf("**********GlobalDeployLock Locked for %s deployment by pid %d***************", who, os.Getpid())
}

func GlobalDeployTrylock(who string) bool {
	log.Printf("**********Checking  GlobalDeployLock  for %s deployment by pid %d***************", who, os.Getpid())
	return instance.deployMutex.TryLock()
}

func GlobalDeployUnlock(who string) {
	instance.deployMutex.Unlock()
	log.Printf("++++++++++GlobalDeployUnlock unlocked after %s deployment complete pid %d+++++++++++++++", who, os.Getpid())
}

func AcquireResourceLock(rscName string) {
	log.Printf("***************Waiting for AcquireResourceLock for %s by pid %d***********", rscName, os.Getpid())
	instance.deployMutex.RLock()
	log.Printf("***************AcquireResourceLock for %s***********pid %d", rscName, os.Getpid())
}

func ReleaseResourceLock(rscName string) {
	instance.deployMutex.RUnlock()
	log.Printf("+++++++++++++++ReleaseResourceLock for %s+++++++++++pid %d", rscName, os.Getpid())
}

func getNDFCClient() *NDFC {
	return instance
}
