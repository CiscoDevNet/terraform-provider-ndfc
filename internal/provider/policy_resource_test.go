// filepath: /Users/muralidm/code/go/src/terraform-provider-ndfc/internal/provider/policy_resource_test.go
// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"terraform-provider-ndfc/internal/provider/ndfc"
	"terraform-provider-ndfc/internal/provider/resources/resource_policy"
	helper "terraform-provider-ndfc/internal/provider/testing"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const policyTestRscName = "ndfc_policy.policy_test"

func TestAccPolicyResource(t *testing.T) {

	x := &map[string]string{
		"RscType":  ndfc.ResourcePolicy,
		"RscName":  "policy_test",
		"User":     helper.GetConfig("policy").NDFC.User,
		"Password": helper.GetConfig("policy").NDFC.Password,
		"Host":     helper.GetConfig("policy").NDFC.URL,
		"Insecure": helper.GetConfig("policy").NDFC.Insecure,
	}

	tf_config := new(string)
	*tf_config = `provider "ndfc" {
		host     = "https://"
		username = "admin"
		password = "admin!@#"
		domain   = "example.com"
		insecure = true
		}
		resource ndfc_vrfs "net_test" {
			fabric_name = "dummy"
		}`

	stepCount := new(int)
	*stepCount = 0

	// Create a new instance of the NDFC client
	policyResource := new(resource_policy.NDFCPolicyModel)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "policy") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Create a new resource with 20 interfaces
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					/*
						resource "ndfc_policy" "test_resource_policy_1" {
							is_policy_group      = false
							deploy               = true
							entity_name          = "Switch"
							entity_type          = "SWITCH"
							description          = "Policy for switch"
							template_name        = "TelemetryDst_EF"
							source               = "CLI"
							priority             = 500
							device_serial_number = "FDO245206N5"
							policy_parameters = {
							  DSTGRP = "501"
							  IPADDR = "5.5.5.6"
							  PORT   = "57900"
							  VRF    = "management"
							}
							deleted = false
						  }
					*/
					policyResource.Deploy = true
					policyResource.EntityName = "Switch"
					policyResource.EntityType = "SWITCH"
					policyResource.Description = "Policy for switch"
					policyResource.TemplateName = "TelemetryDst_EF"
					policyResource.Source = "CLI"
					policyResource.Priority = new(int64)
					*policyResource.Priority = 500
					policyResource.DeviceSerialNumber = helper.GetConfig("policy").NDFC.Switches[0]
					policyResource.PolicyParameters = map[string]string{
						"DSTGRP": "501",
						"IPADDR": "5.5.5.6",
						"PORT":   "57900",
						"VRF":    "management",
					}
					policyResource.Deleted = new(bool)
					*policyResource.Deleted = false
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{policyResource}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(PolicyModelHelperStateCheck("ndfc_policy.policy_test", *policyResource, path.Empty())...),
			},
		}})
}

func TestAccPolicyResourceVrfLite(t *testing.T) {

	x := &map[string]string{
		"RscType":  ndfc.ResourcePolicy,
		"RscName":  "policy_test",
		"User":     helper.GetConfig("policy").NDFC.User,
		"Password": helper.GetConfig("policy").NDFC.Password,
		"Host":     helper.GetConfig("policy").NDFC.URL,
		"Insecure": helper.GetConfig("policy").NDFC.Insecure,
	}

	tf_config := new(string)
	*tf_config = `provider "ndfc" {
		host     = "https://"
		username = "admin"
		password = "admin!@#"
		domain   = "example.com"
		insecure = true
		}
		resource ndfc_vrfs "net_test" {
			fabric_name = "dummy"
		}`

	stepCount := new(int)
	*stepCount = 0

	// Create a new instance of the NDFC client
	policyResource := new(resource_policy.NDFCPolicyModel)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "policy") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Create a new resource with 20 interfaces
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					/*
											{
						  "nvPairs": {
						    "IF_NAME": "eth1/20",
						    "DOT1Q_ID": "2001",
						    "PEER_VRF_NAME": "default",
						    "IP_MASK": "10.1.1.1/24",
						    "MTU": "9216",
						    "NEIGHBOR_IP": "10.1.1.2",
						    "NEIGHBOR_ASN": "29500",
						    "IPV6_MASK": "",
						    "IPV6_NEIGHBOR": "",
						    "TRM_ENABLE": "",
						    "asn": "",
						    "bgpPassword": "",
						    "bgpPasswordKeyType": "",
						    "SERIAL_NUMBER": "",
						    "SOURCE": "",
						    "POLICY_ID": ""
						  },
						  "entityName": "SWITCH",
						  "entityType": "SWITCH",
						  "source": "",
						  "priority": 500,
						  "description": "VRF LITE Policy",
						  "templateName": "Ext_VRF_Lite_Jython",
						  "serialNumber": "9Q34PHYLDB5"
						}

					*/
					policyResource.Deploy = true
					policyResource.EntityName = "SWITCH"
					policyResource.EntityType = "SWITCH"
					policyResource.Description = "VRF LITE Policy"
					policyResource.TemplateName = "Ext_VRF_Lite_Jython"
					policyResource.Source = "CLI"
					policyResource.Priority = new(int64)
					*policyResource.Priority = 500
					policyResource.DeviceSerialNumber = helper.GetConfig("policy").NDFC.Switches[0]
					policyResource.PolicyParameters = map[string]string{
						"IF_NAME":       "vlan100",
						"DOT1Q_ID":      "2001",
						"PEER_VRF_NAME": "default",
						"IP_MASK":       "10.1.1.1/24",
						"MTU":           "9216",
						"NEIGHBOR_IP":   "10.1.1.2",
						"NEIGHBOR_ASN":  "29500",
					}
					policyResource.Deleted = new(bool)
					*policyResource.Deleted = false
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{policyResource}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(PolicyModelHelperStateCheck("ndfc_policy.policy_test", *policyResource, path.Empty())...),
			},
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					*policyResource.Priority = 550
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{policyResource}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(PolicyModelHelperStateCheck("ndfc_policy.policy_test", *policyResource, path.Empty())...),
			},
		}})
}

func TestAccPolicyResourceVrfLiteRouted(t *testing.T) {

	x := &map[string]string{
		"RscType":  ndfc.ResourcePolicy,
		"RscName":  "policy_test",
		"User":     helper.GetConfig("policy").NDFC.User,
		"Password": helper.GetConfig("policy").NDFC.Password,
		"Host":     helper.GetConfig("policy").NDFC.URL,
		"Insecure": helper.GetConfig("policy").NDFC.Insecure,
	}

	tf_config := new(string)
	*tf_config = `provider "ndfc" {
		host     = "https://"
		username = "admin"
		password = "admin!@#"
		domain   = "example.com"
		insecure = true
		}
		resource ndfc_vrfs "net_test" {
			fabric_name = "dummy"
		}`

	stepCount := new(int)
	*stepCount = 0

	// Create a new instance of the NDFC client
	policyResource := new(resource_policy.NDFCPolicyModel)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "policy") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Create a new resource with 20 interfaces
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					/*
							{
									"nvPairs": {
										"IF_NAME": "ethernet1/20",
										"VRF_NAME": "default",
										"IP_MASK": "20.1.1.1/24",
										"NEIGHBOR_IP": "20.1.1.12",
										"IPv6_MASK": "",
										"NEIGHBOR_IPv6": "",
										"NEIGHBOR_ASN": "86000",
										"asn": "",
										"bgpPassword": "",
										"bgpPasswordKeyType": "",
										"MTU": "9216",
										"SERIAL_NUMBER": "",
										"SOURCE": "",
										"POLICY_ID": "",
										"ROUTING_TAG": "",
										"ROUTE_MAP_IN": "",
										"ROUTE_MAP_OUT": "",
										"IPV6_ROUTE_MAP_IN": "",
										"IPV6_ROUTE_MAP_OUT": "",
										"DESC": "",
										"CONF": "",
										"ADMIN_STATE": "true"
									},
									"entityName": "SWITCH",
									"entityType": "SWITCH",
									"source": "",
									"priority": 500,
									"description": "",
									"templateName": "Ext_VRF_Lite_Routed",
									"serialNumber": "9TQYTJSZ1VJ"
						}
					*/
					policyResource.Deploy = true
					policyResource.EntityName = "SWITCH"
					policyResource.EntityType = "SWITCH"
					policyResource.Description = "VRF LITE Routed Policy"
					policyResource.TemplateName = "Ext_VRF_Lite_Routed"
					policyResource.Source = "CLI"
					policyResource.Priority = new(int64)
					*policyResource.Priority = 500
					policyResource.DeviceSerialNumber = helper.GetConfig("policy").NDFC.Switches[1]
					policyResource.PolicyParameters = map[string]string{
						"IF_NAME":            "ethernet1/22",
						"VRF_NAME":           "default",
						"IP_MASK":            "20.1.1.1/24",
						"NEIGHBOR_IP":        "20.1.1.10",
						"NEIGHBOR_ASN":       "86000",
						"MTU":                "9216",
						"ADMIN_STATE":        "true",
						"IPv6_MASK":          "",
						"ROUTING_TAG":        "",
						"DESC":               "",
						"CONF":               "",
						"NEIGHBOR_IPv6":      "",
						"ROUTE_MAP_IN":       "",
						"ROUTE_MAP_OUT":      "",
						"IPV6_ROUTE_MAP_IN":  "",
						"IPV6_ROUTE_MAP_OUT": "",
					}
					policyResource.Deleted = new(bool)
					*policyResource.Deleted = false
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{policyResource}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(PolicyModelHelperStateCheck("ndfc_policy.policy_test", *policyResource, path.Empty())...),
			},
		}})
}
