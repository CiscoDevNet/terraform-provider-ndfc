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

type TemplateAPI struct {
	NDFCAPICommon
	TemplateName string
	Validation   bool
	mutex        *sync.Mutex
}

const UrlTemplate = "/configtemplate/rest/config/templates/template"
const UrlGetTemlate = "/configtemplate/rest/config/templates/%s"
const UrlValidateTemplate = "/configtemplate/rest/config/templates/validate"

func (c *TemplateAPI) GetLock() *sync.Mutex {
	return c.mutex
}

func (c *TemplateAPI) GetUrl() string {
	if c.TemplateName != "" {
		return fmt.Sprintf(UrlGetTemlate, c.TemplateName)
	}
	return UrlTemplate
}

func (c *TemplateAPI) PostUrl() string {
	if c.Validation {
		return UrlValidateTemplate
	}
	return UrlTemplate
}

func (c *TemplateAPI) PutUrl() string {
	if c.TemplateName != "" {
		return fmt.Sprintf(UrlGetTemlate, c.TemplateName)
	}
	return UrlTemplate
}

func (c *TemplateAPI) DeleteUrl() string {
	if c.TemplateName != "" {
		return fmt.Sprintf(UrlGetTemlate, c.TemplateName)
	}
	return UrlTemplate
}

func (c *TemplateAPI) GetDeleteQP() []string {
	return nil
}

func (c *TemplateAPI) SetValidation(flag bool) {
	c.Validation = flag
}

func (c *TemplateAPI) SetTemplateName(name string) {
	c.TemplateName = name
}

func (c *TemplateAPI) RscName() string {
	return "template"
}


func NewTemplateAPI(lock *sync.Mutex, client *nd.Client) *TemplateAPI {
	tApi := new(TemplateAPI)
	tApi.mutex = lock
	tApi.NDFCAPI = tApi
	tApi.client = client
	return tApi
}
