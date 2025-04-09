// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package api

import (
	"strings"
	"sync"

	"github.com/netascode/go-nd"
)

type RestAPI struct {
	NDFCAPICommon
	mutex    *sync.Mutex
	Url      string
	DeleteQp string
}

func (c *RestAPI) GetLock() *sync.Mutex {
	return c.mutex
}

func (c *RestAPI) GetUrl() string {
	return c.Url
}

func (c *RestAPI) PostUrl() string {
	return c.Url
}

func (c *RestAPI) PutUrl() string {
	return c.Url
}

func (c *RestAPI) DeleteUrl() string {
	return c.Url
}

func (c *RestAPI) GetDeleteQP() []string {
	if c.DeleteQp == "" {
		return nil
	}
	/* split format key=value1&key=value2 */
	qp := strings.Split(c.DeleteQp, "&")

	return qp
}

func (c *RestAPI) RscName() string {
	return "rest_api"
}

func NewRestAPI(lock *sync.Mutex, c *nd.Client) *RestAPI {
	papi := new(RestAPI)
	papi.mutex = lock
	papi.client = c
	papi.NDFCAPI = papi
	return papi
}
