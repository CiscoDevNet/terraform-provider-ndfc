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
	"strings"
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

/*
	resource "ndfc_policy" "alias" {
	  is_policy_group      = false
	  deploy               = true
	  entity_name          = "Switch"
	  entity_type          = "SWITCH"
	  description          = "Policy for alias"
	  template_name        = "switch_freeform"
	  source               = "CLI"
	  priority             = 500
	  device_serial_number = "9Q34PHYLDB5"
	  policy_parameters = {
	    "CONF" = <<-EOT
	    cli alias name wr copy run start
	    EOT
	  }
	}

	resource "ndfc_policy" "banner" {
	  is_policy_group      = false
	  deploy               = true
	  entity_name          = "Switch"
	  entity_type          = "SWITCH"
	  description          = "Policy for banner"
	  template_name        = "banner"
	  source               = "CLI"
	  priority             = 500
	  device_serial_number = "9Q34PHYLDB5"
	  policy_parameters = {
	    BANNER = "VXLAN as Code Banner"
	    BANNERDELIMITER = "_"
	    TYPE = "motd"
	  }
	}

	resource "ndfc_policy" "SSH_KEY_2048" {
	  is_policy_group      = false
	  deploy               = true
	  entity_name          = "Switch"
	  entity_type          = "SWITCH"
	  description          = "Policy for SSH_KEY_2048"
	  template_name        = "ssh_key_rsa_force"
	  source               = "CLI"
	  priority             = 500
	  device_serial_number = "9Q34PHYLDB5"
	  policy_parameters = {
	    BITS = 2048
	  }
	}

	resource "ndfc_policy" "iBGP_ISN1_ISN3" {
	  is_policy_group      = false
	  deploy               = true
	  entity_name          = "Switch"
	  entity_type          = "SWITCH"
	  description          = "Policy for iBGP_ISN1_ISN3"
	  template_name        = "vrf_lite_ibgp"
	  source               = "CLI"
	  priority             = 500
	  device_serial_number = "9Q34PHYLDB5"
	  policy_parameters = {
	    vrfName = "default"
	    BGP_ASN = "29500"
	    NEIGHBOR_IP = "100.65.0.3"
	    bgpPassword = "9125d59c18a9b015"
	    bgpPasswordKeyType = 3
	  }
	}

	resource "ndfc_policy" "iBGP_ISN1_ISN2" {
	  is_policy_group      = false
	  deploy               = true
	  entity_name          = "Switch"
	  entity_type          = "SWITCH"
	  description          = "Policy for iBGP_ISN1_ISN2"
	  template_name        = "vrf_lite_ibgp"
	  source               = "CLI"
	  priority             = 500
	  device_serial_number = "9Q34PHYLDB5"
	  policy_parameters = {
	    vrfName = "default"
	    BGP_ASN = "29500"
	    NEIGHBOR_IP = "100.65.0.5"
	    bgpPassword = "9125d59c18a9b015"
	    bgpPasswordKeyType = 3
	  }
	}

	resource "ndfc_policy" "iBGP_ISN1_ISN4" {
	  is_policy_group      = false
	  deploy               = true
	  entity_name          = "Switch"
	  entity_type          = "SWITCH"
	  description          = "Policy for iBGP_ISN1_ISN4"
	  template_name        = "vrf_lite_ibgp"
	  source               = "CLI"
	  priority             = 500
	  device_serial_number = "9Q34PHYLDB5"
	  policy_parameters = {
	    vrfName = "default"
	    BGP_ASN = "29500"
	    NEIGHBOR_IP = "100.65.0.9"
	    bgpPassword = "9125d59c18a9b015"
	    bgpPasswordKeyType = 3
	  }
	}

	resource "ndfc_policy" "iBGP_ISN2_ISN1" {
	  is_policy_group      = false
	  deploy               = true
	  entity_name          = "Switch"
	  entity_type          = "SWITCH"
	  description          = "Policy for iBGP_ISN2_ISN1"
	  template_name        = "vrf_lite_ibgp"
	  source               = "CLI"
	  priority             = 500
	  device_serial_number = "9Q34PHYLDB5"
	  policy_parameters = {
	    vrfName = "default"
	    BGP_ASN = "29500"
	    NEIGHBOR_IP = "100.65.0.4"
	    bgpPassword = "9125d59c18a9b015"
	    bgpPasswordKeyType = 3
	  }
	}

	resource "ndfc_policy" "iBGP_ISN2_ISN3" {
	  is_policy_group      = false
	  deploy               = true
	  entity_name          = "Switch"
	  entity_type          = "SWITCH"
	  description          = "Policy for iBGP_ISN2_ISN3"
	  template_name        = "vrf_lite_ibgp"
	  source               = "CLI"
	  priority             = 500
	  device_serial_number = "9Q34PHYLDB5"
	  policy_parameters = {
	    vrfName = "default"
	    BGP_ASN = "29500"
	    NEIGHBOR_IP = "100.65.0.1"
	    bgpPassword = "9125d59c18a9b015"
	    bgpPasswordKeyType = 3
	  }
	}

	resource "ndfc_policy" "iBGP_ISN2_ISN4" {
	  is_policy_group      = false
	  deploy               = true
	  entity_name          = "Switch"
	  entity_type          = "SWITCH"
	  description          = "Policy for iBGP_ISN2_ISN4"
	  template_name        = "vrf_lite_ibgp"
	  source               = "CLI"
	  priority             = 500
	  device_serial_number = "9Q34PHYLDB5"
	  policy_parameters = {
	    vrfName = "default"
	    BGP_ASN = "29500"
	    NEIGHBOR_IP = "100.65.0.7"
	    bgpPassword = "9125d59c18a9b015"
	    bgpPasswordKeyType = 3
	  }
	}

	resource "ndfc_policy" "iBGP_ISN3_ISN1" {
	  is_policy_group      = false
	  deploy               = true
	  entity_name          = "Switch"
	  entity_type          = "SWITCH"
	  description          = "Policy for iBGP_ISN3_ISN1"
	  template_name        = "vrf_lite_ibgp"
	  source               = "CLI"
	  priority             = 500
	  device_serial_number = "9Q34PHYLDB5"
	  policy_parameters = {
	    vrfName = "default"
	    BGP_ASN = "29500"
	    NEIGHBOR_IP = "100.65.0.2"
	    bgpPassword = "9125d59c18a9b015"
	    bgpPasswordKeyType = 3
	  }
	}

	resource "ndfc_policy" "iBGP_ISN3_ISN2" {
	  is_policy_group      = false
	  deploy               = true
	  entity_name          = "Switch"
	  entity_type          = "SWITCH"
	  description          = "Policy for iBGP_ISN3_ISN2"
	  template_name        = "vrf_lite_ibgp"
	  source               = "CLI"
	  priority             = 500
	  device_serial_number = "9Q34PHYLDB5"
	  policy_parameters = {
	    vrfName = "default"
	    BGP_ASN = "29500"
	    NEIGHBOR_IP = "100.65.0.0"
	    bgpPassword = "9125d59c18a9b015"
	    bgpPasswordKeyType = 3
	  }
	}

	resource "ndfc_policy" "iBGP_ISN3_ISN4" {
	  is_policy_group      = false
	  deploy               = true
	  entity_name          = "Switch"
	  entity_type          = "SWITCH"
	  description          = "Policy for iBGP_ISN3_ISN4"
	  template_name        = "vrf_lite_ibgp"
	  source               = "CLI"
	  priority             = 500
	  device_serial_number = "9Q34PHYLDB5"
	  policy_parameters = {
	    vrfName = "default"
	    BGP_ASN = "29500"
	    NEIGHBOR_IP = "100.65.0.10"
	    bgpPassword = "9125d59c18a9b015"
	    bgpPasswordKeyType = 3
	  }
	}

	resource "ndfc_policy" "eBGP_ISN3_MPLS_PE" {
	  is_policy_group      = false
	  deploy               = true
	  entity_name          = "Switch"
	  entity_type          = "SWITCH"
	  description          = "Policy for eBGP_ISN3_MPLS_PE"
	  template_name        = "External_VRF_Lite_eBGP"
	  source               = "CLI"
	  priority             = 500
	  device_serial_number = "9Q34PHYLDB5"
	  policy_parameters = {
	    vrfName = "default"
	    vrf_description = "To MPLS-PE1"
	    asn = "29500"
	    NEIGHBOR_ASN = 65111
	    NEIGHBOR_IP = "100.65.0.13"
	    bgpPassword = "9125d59c18a9b015"
	    bgpPasswordKeyType = 3
	  }
	}

	resource "ndfc_policy" "iBGP_ISN4_ISN1" {
	  is_policy_group      = false
	  deploy               = true
	  entity_name          = "Switch"
	  entity_type          = "SWITCH"
	  description          = "Policy for iBGP_ISN4_ISN1"
	  template_name        = "vrf_lite_ibgp"
	  source               = "CLI"
	  priority             = 500
	  device_serial_number = "9Q34PHYLDB5"
	  policy_parameters = {
	    vrfName = "default"
	    BGP_ASN = "29500"
	    NEIGHBOR_IP = "100.65.0.6"
	    bgpPassword = "9125d59c18a9b015"
	    bgpPasswordKeyType = 3
	  }
	}

	resource "ndfc_policy" "iBGP_ISN4_ISN2" {
	  is_policy_group      = false
	  deploy               = true
	  entity_name          = "Switch"
	  entity_type          = "SWITCH"
	  description          = "Policy for iBGP_ISN4_ISN2"
	  template_name        = "vrf_lite_ibgp"
	  source               = "CLI"
	  priority             = 500
	  device_serial_number = "9Q34PHYLDB5"
	  policy_parameters = {
	    vrfName = "default"
	    BGP_ASN = "29500"
	    NEIGHBOR_IP = "100.65.0.8"
	    bgpPassword = "9125d59c18a9b015"
	    bgpPasswordKeyType = 3
	  }
	}

	resource "ndfc_policy" "iBGP_ISN4_ISN3" {
	  is_policy_group      = false
	  deploy               = true
	  entity_name          = "Switch"
	  entity_type          = "SWITCH"
	  description          = "Policy for iBGP_ISN4_ISN3"
	  template_name        = "vrf_lite_ibgp"
	  source               = "CLI"
	  priority             = 500
	  device_serial_number = "9Q34PHYLDB5"
	  policy_parameters = {
	    vrfName = "default"
	    BGP_ASN = "29500"
	    NEIGHBOR_IP = "100.65.0.11"
	    bgpPassword = "9125d59c18a9b015"
	    bgpPasswordKeyType = 3
	  }
	}
*/
func TestAccPolicyResourceMultiple(t *testing.T) {

	x := &map[string]string{
		"RscType":  ndfc.ResourcePolicy,
		"RscName":  "",
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
	policyResources := []*resource_policy.NDFCPolicyModel{
		{
			Deploy:             true,
			EntityName:         "Switch",
			EntityType:         "SWITCH",
			Description:        "Policy for alias",
			TemplateName:       "switch_freeform",
			Source:             "CLI",
			Priority:           new(int64),
			DeviceSerialNumber: helper.GetConfig("policy").NDFC.Switches[0],
			PolicyParameters:   map[string]string{"CONF": "cli alias name wr copy run start"},
			Deleted:            new(bool),
		},
		{
			Deploy:             true,
			EntityName:         "Switch",
			EntityType:         "SWITCH",
			Description:        "Policy for banner",
			TemplateName:       "banner",
			Source:             "CLI",
			Priority:           new(int64),
			DeviceSerialNumber: helper.GetConfig("policy").NDFC.Switches[0],
			PolicyParameters:   map[string]string{"BANNER": "VXLAN as Code Banner", "BANNERDELIMITER": "_", "TYPE": "motd"},
			Deleted:            new(bool),
		},
		{
			Deploy:             true,
			EntityName:         "Switch",
			EntityType:         "SWITCH",
			Description:        "Policy for iBGP_ISN1_ISN3",
			TemplateName:       "vrf_lite_ibgp",
			Source:             "CLI",
			Priority:           new(int64),
			DeviceSerialNumber: helper.GetConfig("policy").NDFC.Switches[0],
			PolicyParameters:   map[string]string{"vrfName": "default", "BGP_ASN": "29500", "NEIGHBOR_IP": "100.65.0.3", "bgpPassword": "9125d59c18a9b015", "bgpPasswordKeyType": "3"},
			Deleted:            new(bool),
		},
		{
			Deploy:             true,
			EntityName:         "Switch",
			EntityType:         "SWITCH",
			Description:        "Policy for iBGP_ISN1_ISN2",
			TemplateName:       "vrf_lite_ibgp",
			Source:             "CLI",
			Priority:           new(int64),
			DeviceSerialNumber: helper.GetConfig("policy").NDFC.Switches[0],
			PolicyParameters:   map[string]string{"vrfName": "default", "BGP_ASN": "29500", "NEIGHBOR_IP": "100.65.0.5", "bgpPassword": "9125d59c18a9b015", "bgpPasswordKeyType": "3"},
			Deleted:            new(bool),
		},
		{
			Deploy:             true,
			EntityName:         "Switch",
			EntityType:         "SWITCH",
			Description:        "Policy for iBGP_ISN1_ISN4",
			TemplateName:       "vrf_lite_ibgp",
			Source:             "CLI",
			Priority:           new(int64),
			DeviceSerialNumber: helper.GetConfig("policy").NDFC.Switches[0],
			PolicyParameters:   map[string]string{"vrfName": "default", "BGP_ASN": "29500", "NEIGHBOR_IP": "100.65.0.9", "bgpPassword": "9125d59c18a9b015", "bgpPasswordKeyType": "3"},
			Deleted:            new(bool),
		},
		{
			Deploy:             true,
			EntityName:         "Switch",
			EntityType:         "SWITCH",
			Description:        "Policy for iBGP_ISN2_ISN3",
			TemplateName:       "vrf_lite_ibgp",
			Source:             "CLI",
			Priority:           new(int64),
			DeviceSerialNumber: helper.GetConfig("policy").NDFC.Switches[0],
			PolicyParameters:   map[string]string{"vrfName": "default", "BGP_ASN": "29500", "NEIGHBOR_IP": "100.65.0.1", "bgpPassword": "9125d59c18a9b015", "bgpPasswordKeyType": "3"},
			Deleted:            new(bool),
		},
		{
			Deploy:             true,
			EntityName:         "Switch",
			EntityType:         "SWITCH",
			Description:        "Policy for iBGP_ISN2_ISN1",
			TemplateName:       "vrf_lite_ibgp",
			Source:             "CLI",
			Priority:           new(int64),
			DeviceSerialNumber: helper.GetConfig("policy").NDFC.Switches[0],
			PolicyParameters:   map[string]string{"vrfName": "default", "BGP_ASN": "29500", "NEIGHBOR_IP": "100.65.0.4", "bgpPassword": "9125d59c18a9b015", "bgpPasswordKeyType": "3"},
			Deleted:            new(bool),
		},
		{
			Deploy:             true,
			EntityName:         "Switch",
			EntityType:         "SWITCH",
			Description:        "Policy for iBGP_ISN2_ISN4",
			TemplateName:       "vrf_lite_ibgp",
			Source:             "CLI",
			Priority:           new(int64),
			DeviceSerialNumber: helper.GetConfig("policy").NDFC.Switches[0],
			PolicyParameters:   map[string]string{"vrfName": "default", "BGP_ASN": "29500", "NEIGHBOR_IP": "100.65.0.7", "bgpPassword": "9125d59c18a9b015", "bgpPasswordKeyType": "3"},
			Deleted:            new(bool),
		},
		{
			Deploy:             true,
			EntityName:         "Switch",
			EntityType:         "SWITCH",
			Description:        "Policy for iBGP_ISN3_ISN1",
			TemplateName:       "vrf_lite_ibgp",
			Source:             "CLI",
			Priority:           new(int64),
			DeviceSerialNumber: helper.GetConfig("policy").NDFC.Switches[0],
			PolicyParameters:   map[string]string{"vrfName": "default", "BGP_ASN": "29500", "NEIGHBOR_IP": "100.65.0.2", "bgpPassword": "9125d59c18a9b015", "bgpPasswordKeyType": "3"},
			Deleted:            new(bool),
		},
		{
			Deploy:             true,
			EntityName:         "Switch",
			EntityType:         "SWITCH",
			Description:        "Policy for iBGP_ISN3_ISN2",
			TemplateName:       "vrf_lite_ibgp",
			Source:             "CLI",
			Priority:           new(int64),
			DeviceSerialNumber: helper.GetConfig("policy").NDFC.Switches[0],
			PolicyParameters:   map[string]string{"vrfName": "default", "BGP_ASN": "29500", "NEIGHBOR_IP": "100.65.0.0", "bgpPassword": "9125d59c18a9b015", "bgpPasswordKeyType": "3"},
			Deleted:            new(bool),
		},
		{
			Deploy:             true,
			EntityName:         "Switch",
			EntityType:         "SWITCH",
			Description:        "Policy for iBGP_ISN3_ISN4",
			TemplateName:       "vrf_lite_ibgp",
			Source:             "CLI",
			Priority:           new(int64),
			DeviceSerialNumber: helper.GetConfig("policy").NDFC.Switches[0],
			PolicyParameters:   map[string]string{"vrfName": "default", "BGP_ASN": "29500", "NEIGHBOR_IP": "100.65.0.10", "bgpPassword": "9125d59c18a9b015", "bgpPasswordKeyType": "3"},
			Deleted:            new(bool),
		},
		{
			Deploy:             true,
			EntityName:         "Switch",
			EntityType:         "SWITCH",
			Description:        "Policy for eBGP_ISN3_MPLS_PE",
			TemplateName:       "External_VRF_Lite_eBGP",
			Source:             "CLI",
			Priority:           new(int64),
			DeviceSerialNumber: helper.GetConfig("policy").NDFC.Switches[0],
			PolicyParameters:   map[string]string{"vrfName": "default", "vrf_description": "To MPLS-PE1", "asn": "29500", "NEIGHBOR_ASN": "65111", "NEIGHBOR_IP": "100.65.0.13", "bgpPassword": "9125d59c18a9b015", "bgpPasswordKeyType": "3"},
			Deleted:            new(bool),
		},
		{
			Deploy:             true,
			EntityName:         "Switch",
			EntityType:         "SWITCH",
			Description:        "Policy for iBGP_ISN4_ISN1",
			TemplateName:       "vrf_lite_ibgp",
			Source:             "CLI",
			Priority:           new(int64),
			DeviceSerialNumber: helper.GetConfig("policy").NDFC.Switches[0],
			PolicyParameters:   map[string]string{"vrfName": "default", "BGP_ASN": "29500", "NEIGHBOR_IP": "100.65.0.6", "bgpPassword": "9125d59c18a9b015", "bgpPasswordKeyType": "3"},
			Deleted:            new(bool),
		},
		{
			Deploy:             true,
			EntityName:         "Switch",
			EntityType:         "SWITCH",
			Description:        "Policy for iBGP_ISN4_ISN2",
			TemplateName:       "vrf_lite_ibgp",
			Source:             "CLI",
			Priority:           new(int64),
			DeviceSerialNumber: helper.GetConfig("policy").NDFC.Switches[0],
			PolicyParameters:   map[string]string{"vrfName": "default", "BGP_ASN": "29500", "NEIGHBOR_IP": "100.65.0.8", "bgpPassword": "9125d59c18a9b015", "bgpPasswordKeyType": "3"},
			Deleted:            new(bool),
		},
		{
			Deploy:             true,
			EntityName:         "Switch",
			EntityType:         "SWITCH",
			Description:        "Policy for iBGP_ISN4_ISN3",
			TemplateName:       "vrf_lite_ibgp",
			Source:             "CLI",
			Priority:           new(int64),
			DeviceSerialNumber: helper.GetConfig("policy").NDFC.Switches[0],
			PolicyParameters:   map[string]string{"vrfName": "default", "BGP_ASN": "29500", "NEIGHBOR_IP": "100.65.0.11", "bgpPassword": "9125d59c18a9b015", "bgpPasswordKeyType": "3"},
			Deleted:            new(bool),
		},
	}
	policyRsArr := []interface{}{}
	for _, policyResource := range policyResources {
		*policyResource.Priority = 500
		if (*x)["RscName"] == "" {
			(*x)["RscName"] = strings.Split(policyResource.Description, " ")[2]
		} else {
			(*x)["RscName"] = (*x)["RscName"] + "," + strings.Split(policyResource.Description, " ")[2]
		}
		*policyResource.Deleted = false
		policyRsArr = append(policyRsArr, policyResource)
	}

	resourceNames := strings.Split((*x)["RscName"], ",")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "policy") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GetTFConfigWithSingleResource(tName, *x, policyRsArr, &tf_config)
					return *tf_config
				}(),

				Check: resource.ComposeTestCheckFunc(
					func() []resource.TestCheckFunc {
						ret := []resource.TestCheckFunc{}
						for i, policyResource := range policyResources {
							ret = append(ret, PolicyModelHelperStateCheck("ndfc_policy."+resourceNames[i], *policyResource, path.Empty())...)
						}
						return ret
					}()...),
			},
		},
	})
}
