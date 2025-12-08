// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package api

import (
	"encoding/json"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/netascode/go-nd"
)

type NDFCAPI interface {
	GetLock() *sync.Mutex
	//ProcessResponse(ctx context.Context, res gjson.Result) ([]string, error)
	GetUrl() string
	PostUrl() string
	PutUrl() string
	DeleteUrl() string
	GetDeleteQP() []string
	RscName() string
}

type NDFCAPICommon struct {
	NDFCAPI
	LockedForDeploy bool
	DeployStartTime time.Time
	DeployEndTime   time.Time
	client          *nd.Client
}

var fnGlobalDeployTryLock func(string) bool
var fnRscAcquireLock func(string)
var fnRscReleaseLock func(string)

/*
	func (c NDFCAPICommon) GetLock() *sync.Mutex {
		panic("Not implemented")
	}

	func (c NDFCAPICommon) ProcessResponse(ctx context.Context, res gjson.Result) ([]string, error) {
		panic("Not implemented")
	}

	func (c NDFCAPICommon) GetUrl() string {
		panic("Not implemented")
	}

	func (c NDFCAPICommon) PostUrl() string {
		panic("Not implemented")
	}

	func (c NDFCAPICommon) PutUrl() string {
		panic("Not implemented")
	}

	func (c NDFCAPICommon) DeleteUrl() string {
		panic("Not implemented")
	}
*/
func (c NDFCAPICommon) Get() ([]byte, error) {
	lock := c.NDFCAPI.GetLock()
	lock.Lock()
	url := c.NDFCAPI.GetUrl()
	log.Printf("Get URL: %s\n", url)
	if c.client == nil {
		log.Printf("************Client is nil********************")
	}
	res, err := c.client.GetRawJson(url)
	if err != nil {
		lock.Unlock()
		return nil, err
	}

	lock.Unlock()
	log.Printf("Finished GET: %s %v\n", c.NDFCAPI.GetUrl(), lock)
	return res, nil
}

func (c NDFCAPICommon) Post(payload []byte) (nd.Res, error) {

	url := c.NDFCAPI.PostUrl()
	if strings.Contains(url, "deploy") {
		panic("Deploy URL detected in Post call. Use DeployPost method for deployments")
		//log.Fatal("Deploy URL in Post. Call DeployPost instead")
	}
	// Acquire deploy read lock if not already locked
	// This blocks if a deployment in in progress
	// This is for all Create/Update Post operations
	fnRscAcquireLock(c.NDFCAPI.RscName())
	defer fnRscReleaseLock(c.NDFCAPI.RscName())
	log.Printf("Post URL: %s\n", c.NDFCAPI.PostUrl())
	lock := c.NDFCAPI.GetLock()
	lock.Lock()
	defer lock.Unlock()
	log.Printf("Post URL acquired lock: %s\n", c.NDFCAPI.PostUrl())
	var res nd.Res
	var err error
	if !json.Valid(payload) {
		res, err = c.client.Post(c.NDFCAPI.PostUrl(), string(payload), nd.RemoveContentType)
	} else {
		res, err = c.client.Post(c.NDFCAPI.PostUrl(), string(payload))
	}
	if err != nil {
		// Ensure response body is populated for error reporting
		if res.Str == "" && res.String() != "" {
			res.Str = res.String()
		}
		if res.Str == "" && res.Raw != "" {
			res.Str = res.Raw
		}
		log.Printf("[ERR] Failed to make POST request: %s |%v|", err, res)
		return res, err
	}
	log.Printf("[DEBUG] POST request was ok")
	return res, nil
}

func (c NDFCAPICommon) Put(payload []byte) (nd.Res, error) {
	// Acquire deploy r lock
	fnRscAcquireLock(c.NDFCAPI.RscName())
	defer fnRscReleaseLock(c.NDFCAPI.RscName())
	log.Printf("Put URL: %s\n", c.NDFCAPI.PutUrl())
	lock := c.NDFCAPI.GetLock()
	lock.Lock()
	defer lock.Unlock()
	log.Printf("Put URL acquired lock: %s\n", c.NDFCAPI.PutUrl())
	var res nd.Res
	var err error
	res, err = c.client.Put(c.NDFCAPI.PutUrl(), string(payload))
	if err != nil {
		// Ensure response body is populated for error reporting
		if res.Str == "" && res.String() != "" {
			res.Str = res.String()
		}
		if res.Str == "" && res.Raw != "" {
			res.Str = res.Raw
		}
		log.Printf("[ERR] Failed to make PUT request: %s |%v|", err, res)
		return res, err
	}
	log.Printf("[DEBUG] PUT request was ok")
	return res, nil
}

func (c NDFCAPICommon) Delete() (nd.Res, error) {
	fnRscAcquireLock(c.NDFCAPI.RscName())
	defer fnRscReleaseLock(c.NDFCAPI.RscName())

	c.NDFCAPI.GetLock().Lock()
	defer c.NDFCAPI.GetLock().Unlock()
	qp := c.NDFCAPI.GetDeleteQP()
	log.Printf("Delete URL: %s\n", c.NDFCAPI.DeleteUrl())
	log.Printf("Delete Query Parameters: %v\n", qp)
	var res nd.Res
	var err error
	if qp != nil {
		res, err = c.client.Delete(c.NDFCAPI.DeleteUrl(), "", func(req *nd.Req) {
			q := req.HttpReq.URL.Query()
			for _, s := range qp {
				keys := strings.Split(s, "=")
				q.Add(keys[0], keys[1])

			}
			req.HttpReq.URL.RawQuery = q.Encode()
		})
	} else {
		res, err = c.client.Delete(c.NDFCAPI.DeleteUrl(), "")
	}
	if err != nil {
		// Ensure response body is populated for error reporting
		if res.Str == "" && res.String() != "" {
			res.Str = res.String()
		}
		if res.Str == "" && res.Raw != "" {
			res.Str = res.Raw
		}
		log.Printf("[ERR] Failed to make DELETE request: %s |%v|", err, res)
		return res, err
	}
	log.Printf("[DEBUG] DELETE request was ok")
	return res, nil
}

func (c NDFCAPICommon) DeleteWithPayload(payload []byte) (nd.Res, error) {
	fnRscAcquireLock(c.NDFCAPI.RscName())
	defer fnRscReleaseLock(c.NDFCAPI.RscName())

	c.NDFCAPI.GetLock().Lock()
	defer c.NDFCAPI.GetLock().Unlock()
	log.Printf("DeleteWithPayload URL: %s\n", c.NDFCAPI.DeleteUrl())
	var res nd.Res
	var err error
	res, err = c.client.Delete(c.NDFCAPI.DeleteUrl(), string(payload))
	if err != nil {
		// Ensure response body is populated for error reporting
		if res.Str == "" && res.String() != "" {
			res.Str = res.String()
		}
		if res.Str == "" && res.Raw != "" {
			res.Str = res.Raw
		}
		log.Printf("[ERR] Failed to make DELETE request with payload: %s |%v|", err, res)
		return res, err
	}
	log.Printf("[DEBUG] DELETE request with payload was ok")
	return res, nil
}

func (c *NDFCAPICommon) SetDeployLocked() {
	c.LockedForDeploy = true
}

func (c NDFCAPICommon) DeployPost(payload []byte) (nd.Res, error) {
	// Global write lock must be acquired before deploy lock
	// Check
	if fnGlobalDeployTryLock(c.NDFCAPI.RscName()) {
		//Try lock successful - means lock is available
		panic("Deploy write Lock not taken by caller. GlobalDeployLock must be taken before calling DeployPost")
	} else {
		log.Printf("Deploy write lock is already acquired for %s", c.NDFCAPI.RscName())
	}
	log.Printf("Deploy Post URL: %s\n", c.NDFCAPI.PostUrl())
	lock := c.NDFCAPI.GetLock()
	lock.Lock()
	defer lock.Unlock()
	log.Printf("Deploy Post URL acquired lock: %s\n", c.NDFCAPI.PostUrl())
	var res nd.Res
	var err error
	res, err = c.client.Post(c.NDFCAPI.PostUrl(), string(payload))
	if err != nil {
		// Ensure response body is populated for error reporting
		if res.Str == "" && res.String() != "" {
			res.Str = res.String()
		}
		if res.Str == "" && res.Raw != "" {
			res.Str = res.Raw
		}
		log.Printf("[ERR] Failed to make DEPLOY POST request: %s |%v|", err, res)
		return res, err
	}
	log.Printf("[DEBUG] DEPLOY POST request was ok")
	return res, nil
}

func SetLockFns(tryFn func(string) bool, locks []func(string)) {
	fnGlobalDeployTryLock = tryFn
	fnRscAcquireLock = locks[0]
	fnRscReleaseLock = locks[1]
}
