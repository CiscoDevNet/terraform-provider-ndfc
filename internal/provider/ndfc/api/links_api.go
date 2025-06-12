// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
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

	"github.com/netascode/go-nd"
)

// LinksAPI provides API functionality for NDFC links
type LinksAPI struct {
	NDFCAPICommon
	LinkUUID string
	mutex    *sync.Mutex
}

// NewLinksAPI creates a new LinksAPI instance
func NewLinksAPI(client *nd.Client, linkUUID string) *LinksAPI {
	api := &LinksAPI{
		LinkUUID: linkUUID,
		mutex:    &sync.Mutex{},
	}
	api.client = client
	// Set the API interface to self for polymorphic behavior
	api.NDFCAPI = api
	return api
}

// URL constants for the NDFC links API
const UrlLinks = "/lan-fabric/rest/control/links"
const UrlLinkDetail = "/lan-fabric/rest/control/links/%s"

// GetLock returns the mutex for this API
func (c *LinksAPI) GetLock() *sync.Mutex {
	return c.mutex
}

// GetUrl returns the URL for getting link details
func (c *LinksAPI) GetUrl() string {
	if c.LinkUUID != "" {
		return fmt.Sprintf(UrlLinkDetail, c.LinkUUID)
	}
	return UrlLinks
}

// PostUrl returns the URL for creating links
func (c *LinksAPI) PostUrl() string {
	return UrlLinks
}

// PutUrl returns the URL for updating links
func (c *LinksAPI) PutUrl() string {
	if c.LinkUUID == "" {
		panic("LinkUUID is required for PutUrl")
	}
	return fmt.Sprintf(UrlLinkDetail, c.LinkUUID)
}

// DeleteUrl returns the URL for deleting links
func (c *LinksAPI) DeleteUrl() string {
	if c.LinkUUID == "" {
		panic("LinkUUID is required for DeleteUrl")
	}
	return fmt.Sprintf(UrlLinkDetail, c.LinkUUID)
}

// GetDeleteQP returns query parameters for delete operations
func (c *LinksAPI) GetDeleteQP() []string {
	return []string{}
}

// RscName returns the resource name
func (c *LinksAPI) RscName() string {
	return "links"
}
