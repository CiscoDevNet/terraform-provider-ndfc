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
	"regexp"
	"testing"

	//"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/hashicorp/terraform-plugin-framework/path"

	"terraform-provider-ndfc/internal/provider/ndfc"
	rna "terraform-provider-ndfc/internal/provider/resources/resource_network_attachments"
	"terraform-provider-ndfc/internal/provider/resources/resource_networks"
	"terraform-provider-ndfc/internal/provider/resources/resource_vrf_bulk"
	"terraform-provider-ndfc/internal/provider/types"

	helper "terraform-provider-ndfc/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

/*
When a test is ran,
Terraform runs plan, apply, refresh, and then final plan for each TestStep in the TestCase.
*/

func TestAccNetworksResourceCRUD(t *testing.T) {

	x := &map[string]string{
		"RscType":  ndfc.ResourceNetworks,
		"RscName":  "network_test",
		"User":     helper.GetConfig("network").NDFC.User,
		"Password": helper.GetConfig("network").NDFC.Password,
		"Host":     helper.GetConfig("network").NDFC.URL,
		"Insecure": helper.GetConfig("network").NDFC.Insecure,
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

	networkRsc := new(resource_networks.NDFCNetworksModel)
	vrfRsc := new(resource_vrf_bulk.NDFCVrfBulkModel)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "network") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		//CheckDestroy:             testAccCheckVrfBulkResourceDestroy(networkRsc),
		Steps: []resource.TestStep{
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateSingleVrfObject(&vrfRsc, helper.GetConfig("network").NDFC.VrfPrefix, helper.GetConfig("network").NDFC.Fabric, 1, false, false, false, helper.GetConfig("network").NDFC.Switches)
					helper.GenerateNetworksObject(&networkRsc, helper.GetConfig("network").NDFC.Fabric,
						10, false, false, false, helper.GetConfig("network").NDFC.VrfPrefix+"1", nil)
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			{
				//Add 10 more Networks
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.IncreaseNetCount(&networkRsc,
						10, false, false, false, helper.GetConfig("network").NDFC.VrfPrefix+"1", nil)
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.DeleteNetworks(&networkRsc,
						11, 20)
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			{

				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					netNewRsc := new(resource_networks.NDFCNetworksModel)
					helper.GenerateNetworksObject(&netNewRsc, helper.GetConfig("network").NDFC.Fabric,
						10, false, false, false, helper.GetConfig("network").NDFC.VrfPrefix+"1", nil)

					(*x)["RscName"] = "vrf_test,network_test,network_test1"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc, netNewRsc}, &tf_config)
					return *tf_config
				}(),
				ExpectError: regexp.MustCompile(".*Networks exist.*"),
			},

			{
				//Modify Few Params in Nets
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					tf_config := new(string)
					helper.ModifyNetworksObject(&networkRsc, 1, map[string]interface{}{
						"vlan_id":                1220,
						"multicast_group":        "239.0.0.100",
						"dhcp_relay_loopback_id": 10,
						"arp_suppression":        "true",
					})
					helper.ModifyNetworksObject(&networkRsc, 10, map[string]interface{}{
						"vlan_id":                1102,
						"multicast_group":        "239.0.0.200",
						"dhcp_relay_loopback_id": 12,
						"arp_suppression":        "true",
						"gateway_ipv4_address":   "192.168.101.1/24",
						"secondary_gateway_1":    "192.168.100.2/24",
					})
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},
		}})
}

func TestAccNetworksResourceAttachmentCRUD(t *testing.T) {

	x := &map[string]string{
		"RscType":  ndfc.ResourceVrfBulk,
		"RscName":  "vrf_test,network_test",
		"User":     helper.GetConfig("network").NDFC.User,
		"Password": helper.GetConfig("network").NDFC.Password,
		"Host":     helper.GetConfig("network").NDFC.URL,
		"Insecure": helper.GetConfig("network").NDFC.Insecure,
	}
	stepCount := new(int)
	*stepCount = 0
	networkRsc := new(resource_networks.NDFCNetworksModel)
	vrfRsc := new(resource_vrf_bulk.NDFCVrfBulkModel)

	resource.Test(t, resource.TestCase{

		PreCheck:                 func() { testAccPreCheck(t, "network") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			//Create VRFs with 2 attachments
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					tf_config := new(string)
					helper.GenerateSingleVrfObject(&vrfRsc, helper.GetConfig("network").NDFC.VrfPrefix, helper.GetConfig("network").NDFC.Fabric, 1, false, false, false, helper.GetConfig("network").NDFC.Switches)
					helper.GenerateNetworksObject(&networkRsc, helper.GetConfig("network").NDFC.Fabric,
						10, false, false, false, helper.GetConfig("network").NDFC.VrfPrefix+"1", []string{helper.GetConfig("network").NDFC.Switches[0], helper.GetConfig("network").NDFC.Switches[1]})
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			{ // Remove both attachments _detach
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					tf_config := new(string)
					//helper.VrfAttachmentsMod(&vrfRsc, 1, 1, nil, "", nil)
					helper.NetAttachmentsMod(&networkRsc, 1, len(networkRsc.Networks), nil, "", nil)
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				// Detach also detaches the entry in VRF in NDFC
				// This causes terraform to report VRF attachment as missing and need re-Add
				// This is a known limitation, due to the way the NDFC  works
				ExpectNonEmptyPlan: true,

				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},
			{
				// Add 2 attachments to all VRFs
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					tf_config := new(string)
					helper.NetAttachmentsMod(&networkRsc, 1, len(networkRsc.Networks), []string{helper.GetConfig("network").NDFC.Switches[0], helper.GetConfig("network").NDFC.Switches[1]}, "", nil)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			{
				// Add 3rd attachment half of them
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					tf_config := new(string)
					helper.NetAttachmentsMod(&networkRsc, 1, len(networkRsc.Networks)/2, helper.GetConfig("network").NDFC.Switches, "", nil)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			{
				// Add 3rd attachment remaining  half
				// Remove 3rd from others
				Config: func() string {
					tf_config := new(string)
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.NetAttachmentsMod(&networkRsc, 1, len(networkRsc.Networks)/2, []string{helper.GetConfig("network").NDFC.Switches[0], helper.GetConfig("network").NDFC.Switches[1]}, "", nil)
					helper.NetAttachmentsMod(&networkRsc, (len(networkRsc.Networks)/2)+1, len(networkRsc.Networks)/2, helper.GetConfig("network").NDFC.Switches, "", nil)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				// Detach and attach also does the same in VRF entry in NDFC
				// This causes terraform to report VRF attachment as missing/added in plan
				// This is a known limitation, due to the way the NDFC works
				ExpectNonEmptyPlan: true,
				Check:              resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			{
				//Modify params
				Config: func() string {
					tf_config := new(string)
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.NetAttachmentsMod(&networkRsc, 1, 1, helper.GetConfig("network").NDFC.Switches, helper.GetConfig("network").NDFC.Switches[2], map[string]interface{}{
						"switch_ports": types.CSVString{"Ethernet1/10", "Ethernet1/12"},
					})
					helper.NetAttachmentsMod(&networkRsc, 10, 10, helper.GetConfig("network").NDFC.Switches, helper.GetConfig("network").NDFC.Switches[2], map[string]interface{}{
						"switch_ports": types.CSVString{"Ethernet1/10", "Ethernet1/12"},
					})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},
			{
				//Modify port list, remove one, add another
				Config: func() string {
					tf_config := new(string)
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.NetAttachmentsMod(&networkRsc, 1, 1, helper.GetConfig("network").NDFC.Switches, helper.GetConfig("network").NDFC.Switches[2], map[string]interface{}{
						"switch_ports": types.CSVString{"Ethernet1/12"},
					})
					helper.NetAttachmentsMod(&networkRsc, 10, 10, helper.GetConfig("network").NDFC.Switches, helper.GetConfig("network").NDFC.Switches[2], map[string]interface{}{
						"switch_ports": types.CSVString{"Ethernet1/12", "Ethernet1/13"},
					})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},
		}})
}

func TestAccNetworksResourceGlobalDeploy(t *testing.T) {
	x := &map[string]string{
		"RscType":  ndfc.ResourceNetworks,
		"RscName":  "network_test",
		"User":     helper.GetConfig("network").NDFC.User,
		"Password": helper.GetConfig("network").NDFC.Password,
		"Host":     helper.GetConfig("network").NDFC.URL,
		"Insecure": helper.GetConfig("network").NDFC.Insecure,
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

	networkRsc := new(resource_networks.NDFCNetworksModel)
	vrfRsc := new(resource_vrf_bulk.NDFCVrfBulkModel)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "network") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				//GlobalDeploy
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateSingleVrfObject(&vrfRsc, helper.GetConfig("network").NDFC.VrfPrefix, helper.GetConfig("network").NDFC.Fabric, 1, false, false, false, helper.GetConfig("network").NDFC.Switches)
					helper.GenerateNetworksObject(&networkRsc, helper.GetConfig("network").NDFC.Fabric,
						10, true, false, false, helper.GetConfig("network").NDFC.VrfPrefix+"1", helper.GetConfig("network").NDFC.Switches)
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				// Network Deploy also deploys VRF, so expect some change in plan
				ExpectNonEmptyPlan: false,
				Check:              resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},
		}})
}

func TestAccNetworksResourceAttachmentDeployNetLevel(t *testing.T) {
	x := &map[string]string{
		"RscType":  ndfc.ResourceNetworks,
		"RscName":  "network_test",
		"User":     helper.GetConfig("network").NDFC.User,
		"Password": helper.GetConfig("network").NDFC.Password,
		"Host":     helper.GetConfig("network").NDFC.URL,
		"Insecure": helper.GetConfig("network").NDFC.Insecure,
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

	networkRsc := new(resource_networks.NDFCNetworksModel)
	vrfRsc := new(resource_vrf_bulk.NDFCVrfBulkModel)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "network") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				//Deploy Net Level
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateSingleVrfObject(&vrfRsc, helper.GetConfig("network").NDFC.VrfPrefix, helper.GetConfig("network").NDFC.Fabric, 1, false, false, true, helper.GetConfig("network").NDFC.Switches)
					helper.GenerateNetworksObject(&networkRsc, helper.GetConfig("network").NDFC.Fabric,
						10, false, true, false, helper.GetConfig("network").NDFC.VrfPrefix+"1", helper.GetConfig("network").NDFC.Switches)
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				// Network Deploy also deploys VRF, so expect some change in plan
				ExpectNonEmptyPlan: false,
				Check:              resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},
		}})
}

func TestAccNetworksResourceAttachmentDeployAttachments(t *testing.T) {
	x := &map[string]string{
		"RscType":  ndfc.ResourceNetworks,
		"RscName":  "network_test",
		"User":     helper.GetConfig("network").NDFC.User,
		"Password": helper.GetConfig("network").NDFC.Password,
		"Host":     helper.GetConfig("network").NDFC.URL,
		"Insecure": helper.GetConfig("network").NDFC.Insecure,
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

	networkRsc := new(resource_networks.NDFCNetworksModel)
	vrfRsc := new(resource_vrf_bulk.NDFCVrfBulkModel)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "network") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				//Deploy Attachments
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateSingleVrfObject(&vrfRsc, helper.GetConfig("network").NDFC.VrfPrefix, helper.GetConfig("network").NDFC.Fabric, 1, false, false, true, helper.GetConfig("network").NDFC.Switches)
					helper.GenerateNetworksObject(&networkRsc, helper.GetConfig("network").NDFC.Fabric,
						10, false, false, true, helper.GetConfig("network").NDFC.VrfPrefix+"1", helper.GetConfig("network").NDFC.Switches)
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				// Network Deploy also deploys VRF, so expect some change in plan
				ExpectNonEmptyPlan: false,
				Check:              resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},
		}})
}

func TestAccNetworksResourceRscUpdateAndGlobalDeploy(t *testing.T) {
	x := &map[string]string{
		"RscType":  ndfc.ResourceNetworks,
		"RscName":  "network_test",
		"User":     helper.GetConfig("network").NDFC.User,
		"Password": helper.GetConfig("network").NDFC.Password,
		"Host":     helper.GetConfig("network").NDFC.URL,
		"Insecure": helper.GetConfig("network").NDFC.Insecure,
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

	networkRsc := new(resource_networks.NDFCNetworksModel)
	vrfRsc := new(resource_vrf_bulk.NDFCVrfBulkModel)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "network") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				//Deploy Attachments
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateSingleVrfObject(&vrfRsc, helper.GetConfig("network").NDFC.VrfPrefix, helper.GetConfig("network").NDFC.Fabric, 1, false, false, true, helper.GetConfig("network").NDFC.Switches)
					helper.GenerateNetworksObject(&networkRsc, helper.GetConfig("network").NDFC.Fabric,
						10, true, false, false, helper.GetConfig("network").NDFC.VrfPrefix+"1", helper.GetConfig("network").NDFC.Switches)
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				// Network Deploy also deploys VRF, so expect some change in plan
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},
			{
				// Modify and Deploy
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.ModifyNetworksObject(&networkRsc, 1, map[string]interface{}{
						"vlan_id": 1220})
					helper.ModifyNetworksObject(&networkRsc, 2, map[string]interface{}{
						"vlan_id": 1221})
					helper.ModifyNetworksObject(&networkRsc, 3, map[string]interface{}{
						"vlan_id": 1222})
					helper.ModifyNetworksObject(&networkRsc, 4, map[string]interface{}{
						"vlan_id": 1223})
					helper.ModifyNetworksObject(&networkRsc, 1, map[string]interface{}{
						"multicast_group": "224.30.1.2"})

					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},
		}})
}

/*
	func TestAccNetwotksResourceGlobalDeployWithChanges(t *testing.T) {
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t, "network") },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					Config: `
					terraform {
						required_providers {
							ndfc = {
							source = "registry.terraform.io/cisco/ndfc"
							}
						}
					}


					resource "ndfc_vrfs" "test_resource_vrf_bulk_1" {
						fabric_name = "CML"
						deploy_all_attachments = true
						vrfs = {
							"VRF1" = {
								deploy_attachments = false
								attach_list = {
									"9FE076D8EJL" = {
										serial_number          = "9FE076D8EJL"
										deploy_this_attachment = false
									}
								}
							}
						}
					}

					resource "ndfc_networks" "test_resource_networks_1" {
					    depends_on = [ndfc_vrfs.test_resource_vrf_bulk_1]
						fabric_name            = "CML"
						deploy_all_attachments = true
						networks = {
							"NET1" = {
								display_name               = "NET1"
								network_id                 = 30001
								network_template           = "Default_Network_Universal"
								network_extension_template = "Default_Network_Extension_Universal"
								vrf_name                   = "VRF1"
								primary_network_id         = 30000
								network_type               = "Normal"
								gateway_ipv4_address       = "192.0.2.1/24"
								gateway_ipv6_address       = "2001:db8::1/64"
								vlan_id                    = 1500
								vlan_name                  = "VLAN2000"
								layer2_only                = false
								interface_description      = "My int description"
								mtu                        = 9200
								secondary_gateway_1        = "192.168.2.1/24"
								secondary_gateway_2        = "192.168.3.1/24"
								secondary_gateway_3        = "192.168.4.1/24"
								secondary_gateway_4        = "192.168.5.1/24"
								arp_suppression            = false
								ingress_replication        = false
								multicast_group            = "233.1.1.1"
								dhcp_relay_loopback_id     = 134
								routing_tag                = 100
								trm                        = true
								route_target_both          = true
								netflow                    = false
								svi_netflow_monitor        = "MON1"
								vlan_netflow_monitor       = "MON1"
								l3_gatway_border           = true
								igmp_version               = "3"
								deploy_attachments         = false
								attachments = {
									"9FE076D8EJL" = {
										deploy_this_attachment = false
									}
								}
							}
						}
					}`,
					Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("ndfc_networks.test_resource_networks_1", "fabric_name", "CML")),
				},
			},
		})

}
*/
func TestAccNetworksResourceAddRemoveNetworks(t *testing.T) {
	x := &map[string]string{
		"RscType":  ndfc.ResourceNetworks,
		"RscName":  "network_test",
		"User":     helper.GetConfig("network").NDFC.User,
		"Password": helper.GetConfig("network").NDFC.Password,
		"Host":     helper.GetConfig("network").NDFC.URL,
		"Insecure": helper.GetConfig("network").NDFC.Insecure,
	}

	tf_config := new(string)
	stepCount := new(int)
	*stepCount = 0
	networkRsc := new(resource_networks.NDFCNetworksModel)
	vrfRsc := new(resource_vrf_bulk.NDFCVrfBulkModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "network") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create 3 networks with 2 attachments each
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateSingleVrfObject(&vrfRsc, helper.GetConfig("network").NDFC.VrfPrefix, helper.GetConfig("network").NDFC.Fabric, 1, false, false, false, helper.GetConfig("network").NDFC.Switches)
					helper.GenerateNetworksObject(&networkRsc, helper.GetConfig("network").NDFC.Fabric,
						3, false, false, false, helper.GetConfig("network").NDFC.VrfPrefix+"1", []string{helper.GetConfig("network").NDFC.Switches[0], helper.GetConfig("network").NDFC.Switches[1]})
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 2: Remove the 3rd network entirely
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.DeleteNetworks(&networkRsc, 3, 3)
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 3: Add the 3rd network back
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Regenerate all 3 networks to add network 3 back
					helper.GenerateNetworksObject(&networkRsc, helper.GetConfig("network").NDFC.Fabric,
						3, false, false, false, helper.GetConfig("network").NDFC.VrfPrefix+"1", []string{helper.GetConfig("network").NDFC.Switches[0], helper.GetConfig("network").NDFC.Switches[1]})
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 4: Remove the 1st network
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.DeleteNetworks(&networkRsc, 1, 1)
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 5: Remove the 2nd network
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.DeleteNetworks(&networkRsc, 2, 2)
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 6: Add both 1st and 2nd networks back
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Regenerate all 3 networks to add networks 1 and 2 back
					helper.GenerateNetworksObject(&networkRsc, helper.GetConfig("network").NDFC.Fabric,
						3, false, false, false, helper.GetConfig("network").NDFC.VrfPrefix+"1", []string{helper.GetConfig("network").NDFC.Switches[0], helper.GetConfig("network").NDFC.Switches[1]})
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 7: Remove networks 2 and 3
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.DeleteNetworks(&networkRsc, 2, 3)
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},
		}},
	)
}

func TestAccNetworksResourceAddRemoveNetworksWithDeploy(t *testing.T) {
	x := &map[string]string{
		"RscType":  ndfc.ResourceNetworks,
		"RscName":  "network_test",
		"User":     helper.GetConfig("network").NDFC.User,
		"Password": helper.GetConfig("network").NDFC.Password,
		"Host":     helper.GetConfig("network").NDFC.URL,
		"Insecure": helper.GetConfig("network").NDFC.Insecure,
	}

	tf_config := new(string)
	stepCount := new(int)
	*stepCount = 0
	networkRsc := new(resource_networks.NDFCNetworksModel)
	vrfRsc := new(resource_vrf_bulk.NDFCVrfBulkModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "network") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create 3 networks with 2 attachments each (with deploy)
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateSingleVrfObject(&vrfRsc, helper.GetConfig("network").NDFC.VrfPrefix, helper.GetConfig("network").NDFC.Fabric, 1, false, false, false, helper.GetConfig("network").NDFC.Switches)
					helper.GenerateNetworksObject(&networkRsc, helper.GetConfig("network").NDFC.Fabric,
						3, false, false, true, helper.GetConfig("network").NDFC.VrfPrefix+"1", []string{helper.GetConfig("network").NDFC.Switches[0], helper.GetConfig("network").NDFC.Switches[1]})
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 2: Remove the 3rd network entirely
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.DeleteNetworks(&networkRsc, 3, 3)
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 3: Add the 3rd network back
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Regenerate all 3 networks to add network 3 back
					helper.GenerateNetworksObject(&networkRsc, helper.GetConfig("network").NDFC.Fabric,
						3, false, false, true, helper.GetConfig("network").NDFC.VrfPrefix+"1", []string{helper.GetConfig("network").NDFC.Switches[0], helper.GetConfig("network").NDFC.Switches[1]})
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 4: Remove the 1st network
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.DeleteNetworks(&networkRsc, 1, 1)
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 5: Remove the 2nd network
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.DeleteNetworks(&networkRsc, 2, 2)
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 6: Add both 1st and 2nd networks back
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Regenerate all 3 networks to add networks 1 and 2 back
					helper.GenerateNetworksObject(&networkRsc, helper.GetConfig("network").NDFC.Fabric,
						3, false, false, true, helper.GetConfig("network").NDFC.VrfPrefix+"1", []string{helper.GetConfig("network").NDFC.Switches[0], helper.GetConfig("network").NDFC.Switches[1]})
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 7: Remove networks 2 and 3
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.DeleteNetworks(&networkRsc, 2, 3)
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},
		}},
	)
}

func TestAccNetworksResourceAddRemoveNetworksWithNetDeployFlag(t *testing.T) {
	x := &map[string]string{
		"RscType":  ndfc.ResourceNetworks,
		"RscName":  "network_test",
		"User":     helper.GetConfig("network").NDFC.User,
		"Password": helper.GetConfig("network").NDFC.Password,
		"Host":     helper.GetConfig("network").NDFC.URL,
		"Insecure": helper.GetConfig("network").NDFC.Insecure,
	}

	tf_config := new(string)
	stepCount := new(int)
	*stepCount = 0
	networkRsc := new(resource_networks.NDFCNetworksModel)
	vrfRsc := new(resource_vrf_bulk.NDFCVrfBulkModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "network") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create 3 networks with 2 attachments each (with network-level deploy)
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateSingleVrfObject(&vrfRsc, helper.GetConfig("network").NDFC.VrfPrefix, helper.GetConfig("network").NDFC.Fabric, 1, false, false, false, helper.GetConfig("network").NDFC.Switches)
					helper.GenerateNetworksObject(&networkRsc, helper.GetConfig("network").NDFC.Fabric,
						3, false, true, false, helper.GetConfig("network").NDFC.VrfPrefix+"1", []string{helper.GetConfig("network").NDFC.Switches[0], helper.GetConfig("network").NDFC.Switches[1]})
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 4: Remove the 1st network
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.DeleteNetworks(&networkRsc, 1, 1)
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 6: Add both 1st and 2nd networks back
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Regenerate all 3 networks to add networks 1 and 2 back
					helper.GenerateNetworksObject(&networkRsc, helper.GetConfig("network").NDFC.Fabric,
						3, false, true, false, helper.GetConfig("network").NDFC.VrfPrefix+"1", []string{helper.GetConfig("network").NDFC.Switches[0], helper.GetConfig("network").NDFC.Switches[1]})
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 7: Remove networks 2 and 3
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.DeleteNetworks(&networkRsc, 2, 3)
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},
		}},
	)
}

func TestAccNetworksResourceAddRemoveNetworksWithGlobalDeployFlag(t *testing.T) {
	x := &map[string]string{
		"RscType":  ndfc.ResourceNetworks,
		"RscName":  "network_test",
		"User":     helper.GetConfig("network").NDFC.User,
		"Password": helper.GetConfig("network").NDFC.Password,
		"Host":     helper.GetConfig("network").NDFC.URL,
		"Insecure": helper.GetConfig("network").NDFC.Insecure,
	}

	tf_config := new(string)
	stepCount := new(int)
	*stepCount = 0
	networkRsc := new(resource_networks.NDFCNetworksModel)
	vrfRsc := new(resource_vrf_bulk.NDFCVrfBulkModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "network") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create 3 networks with 2 attachments each (with global deploy)
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateSingleVrfObject(&vrfRsc, helper.GetConfig("network").NDFC.VrfPrefix, helper.GetConfig("network").NDFC.Fabric, 1, false, false, false, helper.GetConfig("network").NDFC.Switches)
					helper.GenerateNetworksObject(&networkRsc, helper.GetConfig("network").NDFC.Fabric,
						3, true, false, false, helper.GetConfig("network").NDFC.VrfPrefix+"1", []string{helper.GetConfig("network").NDFC.Switches[0], helper.GetConfig("network").NDFC.Switches[1]})
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 4: Remove the 1st network
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.DeleteNetworks(&networkRsc, 1, 1)
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 6: Add both 1st and 2nd networks back
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Regenerate all 3 networks to add networks 1 and 2 back
					helper.GenerateNetworksObject(&networkRsc, helper.GetConfig("network").NDFC.Fabric,
						3, true, false, false, helper.GetConfig("network").NDFC.VrfPrefix+"1", []string{helper.GetConfig("network").NDFC.Switches[0], helper.GetConfig("network").NDFC.Switches[1]})
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 7: Remove networks 2 and 3
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.DeleteNetworks(&networkRsc, 2, 3)
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},
		}},
	)
}

func TestAccNetworksResourceAddRemoveNetworksComboDeploy(t *testing.T) {
	x := &map[string]string{
		"RscType":  ndfc.ResourceNetworks,
		"RscName":  "network_test",
		"User":     helper.GetConfig("network").NDFC.User,
		"Password": helper.GetConfig("network").NDFC.Password,
		"Host":     helper.GetConfig("network").NDFC.URL,
		"Insecure": helper.GetConfig("network").NDFC.Insecure,
	}

	tf_config := new(string)
	stepCount := new(int)
	*stepCount = 0
	networkRsc := new(resource_networks.NDFCNetworksModel)
	vrfRsc := new(resource_vrf_bulk.NDFCVrfBulkModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "network") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create 3 networks with 2 attachments each (deploy OFF)
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateSingleVrfObject(&vrfRsc, helper.GetConfig("network").NDFC.VrfPrefix, helper.GetConfig("network").NDFC.Fabric, 1, false, false, false, helper.GetConfig("network").NDFC.Switches)
					helper.GenerateNetworksObject(&networkRsc, helper.GetConfig("network").NDFC.Fabric,
						3, false, false, false, helper.GetConfig("network").NDFC.VrfPrefix+"1", []string{helper.GetConfig("network").NDFC.Switches[0], helper.GetConfig("network").NDFC.Switches[1]})
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 2: Delete network 3, deploy remaining ones using deploy_this_attachment
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.DeleteNetworks(&networkRsc, 3, 3)
					// Turn on deploy_this_attachment for networks 1 and 2
					for i := 1; i <= 2; i++ {
						nwName := helper.GetConfig("network").NDFC.NetPrefix + fmt.Sprintf("%d", i)
						if nw, ok := networkRsc.Networks[nwName]; ok {
							for serial, attach := range nw.Attachments {
								attach.DeployThisAttachment = true
								nw.Attachments[serial] = attach
							}
							networkRsc.Networks[nwName] = nw
						}
					}
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 3: Add network 3 back with deploy OFF
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Regenerate all 3 networks - networks 1,2 keep deploy=true from step 2
					helper.GenerateNetworksObject(&networkRsc, helper.GetConfig("network").NDFC.Fabric,
						3, false, false, false, helper.GetConfig("network").NDFC.VrfPrefix+"1", []string{helper.GetConfig("network").NDFC.Switches[0], helper.GetConfig("network").NDFC.Switches[1]})
					// Restore deploy_this_attachment=true for networks 1 and 2
					for i := 1; i <= 2; i++ {
						nwName := helper.GetConfig("network").NDFC.NetPrefix + fmt.Sprintf("%d", i)
						if nw, ok := networkRsc.Networks[nwName]; ok {
							for serial, attach := range nw.Attachments {
								attach.DeployThisAttachment = true
								nw.Attachments[serial] = attach
							}
							networkRsc.Networks[nwName] = nw
						}
					}
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 4: Turn on deploy flag for network 3
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Turn on deploy_this_attachment for network 3
					nwName := helper.GetConfig("network").NDFC.NetPrefix + "3"
					if nw, ok := networkRsc.Networks[nwName]; ok {
						for serial, attach := range nw.Attachments {
							attach.DeployThisAttachment = true
							nw.Attachments[serial] = attach
						}
						networkRsc.Networks[nwName] = nw
					}
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 5: Remove attachments from network 3
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.NetAttachmentsMod(&networkRsc, 3, 3, nil, "", nil)
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 6: Add attachments back to network 3 with deploy=true
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Add attachments back
					nwName := helper.GetConfig("network").NDFC.NetPrefix + "3"
					if nw, ok := networkRsc.Networks[nwName]; ok {
						nw.Attachments = make(map[string]rna.NDFCAttachmentsValue)
						for _, serial := range []string{helper.GetConfig("network").NDFC.Switches[0], helper.GetConfig("network").NDFC.Switches[1]} {
							attach := rna.NDFCAttachmentsValue{}
							attach.SerialNumber = serial
							attach.DeployThisAttachment = true
							nw.Attachments[serial] = attach
						}
						networkRsc.Networks[nwName] = nw
					}
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 7: Remove networks 2 and 3
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.DeleteNetworks(&networkRsc, 2, 3)
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},
		}},
	)
}

// TestAccNetworkResourceFreeformConfig tests freeform config on Network attachments
// Tests setting, updating, and removing freeform config
func TestAccNetworksResourceFreeformConfig(t *testing.T) {
	x := &map[string]string{
		"RscType":  ndfc.ResourceNetworks,
		"RscName":  "network_test",
		"User":     helper.GetConfig("network").NDFC.User,
		"Password": helper.GetConfig("network").NDFC.Password,
		"Host":     helper.GetConfig("network").NDFC.URL,
		"Insecure": helper.GetConfig("network").NDFC.Insecure,
	}

	tfConfig := new(string)
	stepCount := new(int)
	*stepCount = 0
	networkRsc := new(resource_networks.NDFCNetworksModel)
	vrfRsc := new(resource_vrf_bulk.NDFCVrfBulkModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "network") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create 3 Networks with 2 attachments each (no freeform config initially)
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateSingleVrfObject(&vrfRsc, helper.GetConfig("network").NDFC.VrfPrefix, helper.GetConfig("network").NDFC.Fabric, 1, false, false, false, helper.GetConfig("network").NDFC.Switches)
					helper.GenerateNetworksObject(
						&networkRsc,                             // network resource model
						helper.GetConfig("network").NDFC.Fabric, // fabric name
						3,                                       // number of networks
						false,                                   // global deploy flag
						false,                                   // network deploy flag
						false,                                   // attachment deploy needed
						helper.GetConfig("network").NDFC.VrfPrefix+"1",                                                       // VRF name
						[]string{helper.GetConfig("network").NDFC.Switches[0], helper.GetConfig("network").NDFC.Switches[1]}, // switch serials
					)
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tfConfig)
					return *tfConfig
				}(),
			},

			// Step 2: Add freeform_config to first attachment of Network 1
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.NetAttachmentsMod(
						&networkRsc, // network resource model
						1,           // start network index
						1,           // end network index
						[]string{helper.GetConfig("network").NDFC.Switches[0], helper.GetConfig("network").NDFC.Switches[1]}, // switch serials to preserve
						helper.GetConfig("network").NDFC.Switches[0],                                                         // serial to modify
						map[string]interface{}{ // parameters to set
							"freeform_config": "interface Vlan100\\n  description Network {networkName} on Switch\\n  mtu 9216\\n  no shutdown",
						},
					)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tfConfig)
					return *tfConfig
				}(),
			},

			// Step 3: Add freeform_config to multiple attachments across different Networks
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Network 1, Switch 2: Add new freeform config
					helper.NetAttachmentsMod(
						&networkRsc, // network resource model
						1,           // start network index
						1,           // end network index
						[]string{helper.GetConfig("network").NDFC.Switches[0], helper.GetConfig("network").NDFC.Switches[1]}, // switch serials to preserve
						helper.GetConfig("network").NDFC.Switches[1],                                                         // serial to modify
						map[string]interface{}{ // parameters to set
							"freeform_config": "interface Vlan101\\n  description Network {networkName} on Switch 2\\n  mtu 9216\\n  no shutdown",
						},
					)
					// Network 2, Switch 1: Add freeform config
					helper.NetAttachmentsMod(
						&networkRsc, // network resource model
						2,           // start network index
						2,           // end network index
						[]string{helper.GetConfig("network").NDFC.Switches[0], helper.GetConfig("network").NDFC.Switches[1]}, // switch serials to preserve
						helper.GetConfig("network").NDFC.Switches[0],                                                         // serial to modify
						map[string]interface{}{ // parameters to set
							"freeform_config": "interface Vlan200\\n  description Network {networkName}\\n  mtu 9000\\n  no shutdown",
						},
					)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tfConfig)
					return *tfConfig
				}(),
			},

			// Step 4: Update existing freeform_config (modify content)
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.NetAttachmentsMod(
						&networkRsc, // network resource model
						1,           // start network index
						1,           // end network index
						[]string{helper.GetConfig("network").NDFC.Switches[0], helper.GetConfig("network").NDFC.Switches[1]}, // switch serials to preserve
						helper.GetConfig("network").NDFC.Switches[0],                                                         // serial to modify
						map[string]interface{}{ // parameters to set
							"freeform_config": "interface Vlan100\\n  description Updated Network {networkName}\\n  mtu 9000\\n  ip address 10.10.10.1/24\\n  no shutdown",
						},
					)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tfConfig)
					return *tfConfig
				}(),
			},

			// Step 5: Remove freeform_config from one attachment (set to empty string)
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.NetAttachmentsMod(
						&networkRsc, // network resource model
						2,           // start network index
						2,           // end network index
						[]string{helper.GetConfig("network").NDFC.Switches[0], helper.GetConfig("network").NDFC.Switches[1]}, // switch serials to preserve
						helper.GetConfig("network").NDFC.Switches[0],                                                         // serial to modify
						map[string]interface{}{ // parameters to set
							"freeform_config": "",
						},
					)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tfConfig)
					return *tfConfig
				}(),
			},

			// Step 6: Add freeform_config along with other attachment parameters
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.NetAttachmentsMod(
						&networkRsc, // network resource model
						3,           // start network index
						3,           // end network index
						[]string{helper.GetConfig("network").NDFC.Switches[0], helper.GetConfig("network").NDFC.Switches[1]}, // switch serials to preserve
						helper.GetConfig("network").NDFC.Switches[0],                                                         // serial to modify
						map[string]interface{}{ // parameters to set
							"freeform_config": "interface Vlan300\\n  description Network {networkName} with VLAN\\n  mtu 9216\\n  no shutdown",
							"switch_ports":    types.CSVString{"Ethernet1/10", "Ethernet1/11"},
						},
					)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tfConfig)
					return *tfConfig
				}(),
			},
		}},
	)
}

// TestAccNetworkResourceParamModifications tests modifying various network parameters
// Tests string, bool, and int parameter modifications with resets
func TestAccNetworksResourceParamModifications(t *testing.T) {
	x := &map[string]string{
		"RscType":  ndfc.ResourceNetworks,
		"RscName":  "network_test",
		"User":     helper.GetConfig("network").NDFC.User,
		"Password": helper.GetConfig("network").NDFC.Password,
		"Host":     helper.GetConfig("network").NDFC.URL,
		"Insecure": helper.GetConfig("network").NDFC.Insecure,
	}

	tfConfig := new(string)
	stepCount := new(int)
	*stepCount = 0
	networkRsc := new(resource_networks.NDFCNetworksModel)
	vrfRsc := new(resource_vrf_bulk.NDFCVrfBulkModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "network") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create 3 Networks with 2 attachments each (no deployment)
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateSingleVrfObject(&vrfRsc, helper.GetConfig("network").NDFC.VrfPrefix, helper.GetConfig("network").NDFC.Fabric, 1, false, false, false, helper.GetConfig("network").NDFC.Switches)
					helper.GenerateNetworksObject(&networkRsc, helper.GetConfig("network").NDFC.Fabric,
						3, false, false, false, helper.GetConfig("network").NDFC.VrfPrefix+"1",
						[]string{helper.GetConfig("network").NDFC.Switches[0], helper.GetConfig("network").NDFC.Switches[1]})
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 2: Modify string parameters (display_name, vlan_name, interface_description)
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.ModifyNetworksObject(&networkRsc, 1, map[string]interface{}{
						"display_name":          "Network1_Display",
						"vlan_name":             "VLAN_NET1",
						"interface_description": "Interface for Network 1",
					})
					helper.ModifyNetworksObject(&networkRsc, 2, map[string]interface{}{
						"display_name":          "Network2_Display",
						"vlan_name":             "VLAN_NET2",
						"interface_description": "Interface for Network 2",
					})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 3: Reset string parameters back to empty
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.ModifyNetworksObject(&networkRsc, 1, map[string]interface{}{
						"display_name":          "",
						"vlan_name":             "",
						"interface_description": "",
					})
					helper.ModifyNetworksObject(&networkRsc, 2, map[string]interface{}{
						"display_name":          "",
						"vlan_name":             "",
						"interface_description": "",
					})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 4: Modify bool parameters (arp_suppression, trm, route_target_both)
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.ModifyNetworksObject(&networkRsc, 1, map[string]interface{}{
						"arp_suppression":   "true",
						"route_target_both": "true",
					})
					helper.ModifyNetworksObject(&networkRsc, 3, map[string]interface{}{
						"arp_suppression":   "true",
						"route_target_both": "true",
					})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 5: Reset bool parameters back to false
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.ModifyNetworksObject(&networkRsc, 1, map[string]interface{}{
						"arp_suppression":   "false",
						"route_target_both": "false",
					})
					helper.ModifyNetworksObject(&networkRsc, 3, map[string]interface{}{
						"arp_suppression":   "false",
						"route_target_both": "false",
					})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 6: Modify int parameters (vlan_id, mtu, dhcp_relay_loopback_id, routing_tag)
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.ModifyNetworksObject(&networkRsc, 1, map[string]interface{}{
						"vlan_id":                1100,
						"mtu":                    9000,
						"dhcp_relay_loopback_id": 10,
						"routing_tag":            54321,
					})
					helper.ModifyNetworksObject(&networkRsc, 2, map[string]interface{}{
						"vlan_id":                1200,
						"mtu":                    1500,
						"dhcp_relay_loopback_id": 20,
					})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 7: Reset int parameters back to defaults
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.ModifyNetworksObject(&networkRsc, 1, map[string]interface{}{
						"mtu":                    9216,
						"dhcp_relay_loopback_id": 0,
						"routing_tag":            12345,
					})
					helper.ModifyNetworksObject(&networkRsc, 2, map[string]interface{}{
						"mtu":                    9216,
						"dhcp_relay_loopback_id": 0,
					})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},
		}},
	)
}

// TestAccNetworksResourceAttachmentParamModifications tests modifying attachment-level parameters
// Tests string, int, and list parameter modifications with resets
func TestAccNetworksResourceAttachmentParamModifications(t *testing.T) {
	x := &map[string]string{
		"RscType":  ndfc.ResourceNetworks,
		"RscName":  "network_test",
		"User":     helper.GetConfig("network").NDFC.User,
		"Password": helper.GetConfig("network").NDFC.Password,
		"Host":     helper.GetConfig("network").NDFC.URL,
		"Insecure": helper.GetConfig("network").NDFC.Insecure,
	}

	tfConfig := new(string)
	stepCount := new(int)
	*stepCount = 0
	networkRsc := new(resource_networks.NDFCNetworksModel)
	vrfRsc := new(resource_vrf_bulk.NDFCVrfBulkModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "network") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create 3 Networks with 2 attachments each (no deployment)
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateSingleVrfObject(&vrfRsc, helper.GetConfig("network").NDFC.VrfPrefix, helper.GetConfig("network").NDFC.Fabric, 1, false, false, false, helper.GetConfig("network").NDFC.Switches)
					helper.GenerateNetworksObject(
						&networkRsc,                             // network resource model
						helper.GetConfig("network").NDFC.Fabric, // fabric name
						3,                                       // number of networks
						false,                                   // global deploy flag
						false,                                   // network deploy flag
						false,                                   // attachment deploy needed
						helper.GetConfig("network").NDFC.VrfPrefix+"1",                                                       // VRF name
						[]string{helper.GetConfig("network").NDFC.Switches[0], helper.GetConfig("network").NDFC.Switches[1]}, // switch serials
					)
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 2: Modify string parameter (freeform_config) on attachments
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.NetAttachmentsMod(
						&networkRsc, // network resource model
						1,           // start network index
						1,           // end network index
						[]string{helper.GetConfig("network").NDFC.Switches[0], helper.GetConfig("network").NDFC.Switches[1]}, // switch serials to preserve
						helper.GetConfig("network").NDFC.Switches[0],                                                         // serial to modify
						map[string]interface{}{ // parameters to set
							"freeform_config": "interface vlan{networkName}\\n  description Test Config 1",
						},
					)
					helper.NetAttachmentsMod(
						&networkRsc, // network resource model
						2,           // start network index
						2,           // end network index
						[]string{helper.GetConfig("network").NDFC.Switches[0], helper.GetConfig("network").NDFC.Switches[1]}, // switch serials to preserve
						helper.GetConfig("network").NDFC.Switches[1],                                                         // serial to modify
						map[string]interface{}{ // parameters to set
							"freeform_config": "interface vlan{networkName}\\n  description Test Config 2",
						},
					)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tfConfig)
					return *tfConfig
				}(),
				// freeform string comparisons fail as it maynot match exactly due to escape chars
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ndfc_networks.network_test", "networks."+helper.GetConfig("network").NDFC.NetPrefix+"1.attachments."+helper.GetConfig("network").NDFC.Switches[0]+".freeform_config"),
					resource.TestCheckResourceAttrSet("ndfc_networks.network_test", "networks."+helper.GetConfig("network").NDFC.NetPrefix+"2.attachments."+helper.GetConfig("network").NDFC.Switches[1]+".freeform_config"),
				),
			},

			// Step 3: Reset string parameter back to empty
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.NetAttachmentsMod(
						&networkRsc, // network resource model
						1,           // start network index
						1,           // end network index
						[]string{helper.GetConfig("network").NDFC.Switches[0], helper.GetConfig("network").NDFC.Switches[1]}, // switch serials to preserve
						helper.GetConfig("network").NDFC.Switches[0],                                                         // serial to modify
						map[string]interface{}{ // parameters to set
							"freeform_config": "",
						},
					)
					helper.NetAttachmentsMod(
						&networkRsc, // network resource model
						2,           // start network index
						2,           // end network index
						[]string{helper.GetConfig("network").NDFC.Switches[0], helper.GetConfig("network").NDFC.Switches[1]}, // switch serials to preserve
						helper.GetConfig("network").NDFC.Switches[1],                                                         // serial to modify
						map[string]interface{}{ // parameters to set
							"freeform_config": "",
						},
					)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 4: Modify int parameter (vlan) on attachments and add switch_ports
			// COMMENTED OUT due to issues
			/*
				{
					Config: func() string {
						*stepCount++
						tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
						helper.NetAttachmentsMod(&networkRsc, 1, 1, nil, helper.GetConfig("network").NDFC.Switches[0], map[string]interface{}{
							"vlan":         2100,
							"switch_ports": types.CSVString{"Ethernet1/10", "Ethernet1/11"},
						})
						helper.NetAttachmentsMod(&networkRsc, 2, 2, nil, helper.GetConfig("network").NDFC.Switches[0], map[string]interface{}{
							"vlan":         2200,
							"switch_ports": types.CSVString{"Ethernet1/12", "Ethernet1/13"},
						})
						helper.NetAttachmentsMod(&networkRsc, 3, 3, nil, helper.GetConfig("network").NDFC.Switches[1], map[string]interface{}{
							"vlan":         2300,
							"switch_ports": types.CSVString{"Ethernet1/14", "Ethernet1/15"},
						})
						helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tfConfig)
						return *tfConfig
					}(),
					Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
				},
			*/

			// // Step 5: Reset vlan to empty (4095 = special value for empty), keep switch_ports
			// {
			// 	Config: func() string {
			// 		*stepCount++
			// 		tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
			// 		helper.NetAttachmentsMod(&networkRsc, 1, 1, nil, helper.GetConfig("network").NDFC.Switches[0], map[string]interface{}{
			// 			"switch_ports": types.CSVString{"Ethernet1/10", "Ethernet1/11"},
			// 		})
			// 		helper.NetAttachmentsMod(&networkRsc, 2, 2, nil, helper.GetConfig("network").NDFC.Switches[0], map[string]interface{}{
			// 			"vlan":         4095,
			// 			"switch_ports": types.CSVString{"Ethernet1/12", "Ethernet1/13"},
			// 		})
			// 		helper.NetAttachmentsMod(&networkRsc, 3, 3, nil, helper.GetConfig("network").NDFC.Switches[1], map[string]interface{}{
			// 			"vlan":         4095,
			// 			"switch_ports": types.CSVString{"Ethernet1/14", "Ethernet1/15"},
			// 		})
			// 		helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tfConfig)
			// 		return *tfConfig
			// 	}(),
			// 	Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			// },

			// Step 6: Modify list parameter (switch_ports) on attachments - set 3 ports
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.NetAttachmentsMod(
						&networkRsc, // network resource model
						1,           // start network index
						1,           // end network index
						[]string{helper.GetConfig("network").NDFC.Switches[0], helper.GetConfig("network").NDFC.Switches[1]}, // switch serials to preserve
						helper.GetConfig("network").NDFC.Switches[0],                                                         // serial to modify
						map[string]interface{}{ // parameters to set
							"switch_ports": types.CSVString{"Ethernet1/10", "Ethernet1/11", "Ethernet1/12"},
						},
					)
					helper.NetAttachmentsMod(
						&networkRsc, // network resource model
						2,           // start network index
						2,           // end network index
						[]string{helper.GetConfig("network").NDFC.Switches[0], helper.GetConfig("network").NDFC.Switches[1]}, // switch serials to preserve
						helper.GetConfig("network").NDFC.Switches[1],                                                         // serial to modify
						map[string]interface{}{ // parameters to set
							"switch_ports": types.CSVString{"Ethernet1/13", "Ethernet1/14", "Ethernet1/15"},
						},
					)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 7: Modify list parameter (switch_ports) - reduce to 1 port
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.NetAttachmentsMod(
						&networkRsc, // network resource model
						1,           // start network index
						1,           // end network index
						[]string{helper.GetConfig("network").NDFC.Switches[0], helper.GetConfig("network").NDFC.Switches[1]}, // switch serials to preserve
						helper.GetConfig("network").NDFC.Switches[0],                                                         // serial to modify
						map[string]interface{}{ // parameters to set
							"switch_ports": types.CSVString{"Ethernet1/10"},
						},
					)
					helper.NetAttachmentsMod(
						&networkRsc, // network resource model
						2,           // start network index
						2,           // end network index
						[]string{helper.GetConfig("network").NDFC.Switches[0], helper.GetConfig("network").NDFC.Switches[1]}, // switch serials to preserve
						helper.GetConfig("network").NDFC.Switches[1],                                                         // serial to modify
						map[string]interface{}{ // parameters to set
							"switch_ports": types.CSVString{"Ethernet1/13"},
						},
					)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},
		}},
	)
}

// TestAccNetworkResourceIPKeyAttachmentCRUD tests basic CRUD operations for attachments using IP keys
func TestAccNetworksResourceIPKeyAttachmentCRUD(t *testing.T) {

	x := &map[string]string{
		"RscType":  ndfc.ResourceVrfBulk,
		"RscName":  "vrf_test,network_test",
		"User":     helper.GetConfig("network").NDFC.User,
		"Password": helper.GetConfig("network").NDFC.Password,
		"Host":     helper.GetConfig("network").NDFC.URL,
		"Insecure": helper.GetConfig("network").NDFC.Insecure,
	}
	stepCount := new(int)
	*stepCount = 0
	networkRsc := new(resource_networks.NDFCNetworksModel)
	vrfRsc := new(resource_vrf_bulk.NDFCVrfBulkModel)

	resource.Test(t, resource.TestCase{

		PreCheck:                 func() { testAccPreCheck(t, "network") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create networks with IP-keyed attachments
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					tf_config := new(string)
					helper.GenerateSingleVrfObject(&vrfRsc, helper.GetConfig("network").NDFC.VrfPrefix, helper.GetConfig("network").NDFC.Fabric, 1, false, false, false, helper.GetConfig("network").NDFC.Switches)
					helper.GenerateNetworksObject(&networkRsc, helper.GetConfig("network").NDFC.Fabric,
						5, false, false, false, helper.GetConfig("network").NDFC.VrfPrefix+"1", []string{helper.GetConfig("network").NDFC.SwitchIP[0], helper.GetConfig("network").NDFC.SwitchIP[1]})
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 2: Remove attachments (detach)
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					tf_config := new(string)
					helper.NetAttachmentsMod(&networkRsc, 1, len(networkRsc.Networks), nil, "", nil)
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				ExpectNonEmptyPlan: true,
				Check:              resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 3: Add 2 attachments back using IP keys
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					tf_config := new(string)
					helper.NetAttachmentsMod(&networkRsc, 1, len(networkRsc.Networks), []string{helper.GetConfig("network").NDFC.SwitchIP[0], helper.GetConfig("network").NDFC.SwitchIP[1]}, "", nil)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 4: Add 3rd attachment using IP key
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					tf_config := new(string)
					helper.NetAttachmentsMod(&networkRsc, 1, len(networkRsc.Networks), helper.GetConfig("network").NDFC.SwitchIP, "", nil)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 5: Modify params on specific attachment using IP key
			{
				Config: func() string {
					tf_config := new(string)
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.NetAttachmentsMod(&networkRsc, 1, 1, helper.GetConfig("network").NDFC.SwitchIP, helper.GetConfig("network").NDFC.SwitchIP[2], map[string]interface{}{
						"switch_ports": types.CSVString{"Ethernet1/10"},
					})

					helper.NetAttachmentsMod(&networkRsc, 3, 3, helper.GetConfig("network").NDFC.SwitchIP, helper.GetConfig("network").NDFC.SwitchIP[2], map[string]interface{}{
						"switch_ports": types.CSVString{"Ethernet1/11"},
					})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},
		}})
}

// TestAccNetworkResourceIPKeyGlobalDeploy tests global deployment with IP-based attachment keys
func TestAccNetworksResourceIPKeyGlobalDeploy(t *testing.T) {

	x := &map[string]string{
		"RscType":  ndfc.ResourceNetworks,
		"RscName":  "network_test",
		"User":     helper.GetConfig("network").NDFC.User,
		"Password": helper.GetConfig("network").NDFC.Password,
		"Host":     helper.GetConfig("network").NDFC.URL,
		"Insecure": helper.GetConfig("network").NDFC.Insecure,
	}
	networkRsc := new(resource_networks.NDFCNetworksModel)
	vrfRsc := new(resource_vrf_bulk.NDFCVrfBulkModel)
	stepCount := new(int)
	*stepCount = 0

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "network") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create networks with global deploy and IP-keyed attachments
			{
				PreConfig: func() {
					t.Logf(logTestStep, t.Name(), *stepCount)
				},
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					tf_config := new(string)
					helper.GenerateSingleVrfObject(&vrfRsc, helper.GetConfig("network").NDFC.VrfPrefix, helper.GetConfig("network").NDFC.Fabric, 1, true, false, false, helper.GetConfig("network").NDFC.Switches)
					helper.GenerateNetworksObject(&networkRsc, helper.GetConfig("network").NDFC.Fabric,
						5, true, false, false, helper.GetConfig("network").NDFC.VrfPrefix+"1", []string{helper.GetConfig("network").NDFC.SwitchIP[0], helper.GetConfig("network").NDFC.SwitchIP[1]})
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},
			// Step 2: Add 3rd Attachment using IP key
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					tf_config := new(string)
					helper.NetAttachmentsMod(&networkRsc, 1, len(networkRsc.Networks), helper.GetConfig("network").NDFC.SwitchIP, "", nil)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},
		},
	})
}

// TestAccNetworkResourceIPKeyNetworkLevelDeploy tests network-level deployment with IP-based keys
func TestAccNetworksResourceIPKeyNetworkLevelDeploy(t *testing.T) {

	x := &map[string]string{
		"RscType":  ndfc.ResourceNetworks,
		"RscName":  "network_test",
		"User":     helper.GetConfig("network").NDFC.User,
		"Password": helper.GetConfig("network").NDFC.Password,
		"Host":     helper.GetConfig("network").NDFC.URL,
		"Insecure": helper.GetConfig("network").NDFC.Insecure,
	}
	networkRsc := new(resource_networks.NDFCNetworksModel)
	vrfRsc := new(resource_vrf_bulk.NDFCVrfBulkModel)
	stepCount := new(int)
	*stepCount = 0

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "network") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create networks with network-level deploy and IP-keyed attachments
			{
				PreConfig: func() {
					t.Logf(logTestStep, t.Name(), *stepCount)
				},
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					tf_config := new(string)
					helper.GenerateSingleVrfObject(&vrfRsc, helper.GetConfig("network").NDFC.VrfPrefix, helper.GetConfig("network").NDFC.Fabric, 1, false, false, true, helper.GetConfig("network").NDFC.Switches)
					helper.GenerateNetworksObject(&networkRsc, helper.GetConfig("network").NDFC.Fabric,
						3, false, true, false, helper.GetConfig("network").NDFC.VrfPrefix+"1", []string{helper.GetConfig("network").NDFC.SwitchIP[0], helper.GetConfig("network").NDFC.SwitchIP[1]})
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},
			// Step 2: Add 3rd attachment using IP key (auto-deployed via network-level flag)
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					tf_config := new(string)
					helper.NetAttachmentsMod(&networkRsc, 1, len(networkRsc.Networks), helper.GetConfig("network").NDFC.SwitchIP, "", nil)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},
		},
	})
}

// TestAccNetworkResourceIPKeyMixedKeys tests attachments with mixed keys (serial + IP) in same network
func TestAccNetworksResourceIPKeyMixedKeys(t *testing.T) {

	x := &map[string]string{
		"RscType":  ndfc.ResourceNetworks,
		"RscName":  "network_test",
		"User":     helper.GetConfig("network").NDFC.User,
		"Password": helper.GetConfig("network").NDFC.Password,
		"Host":     helper.GetConfig("network").NDFC.URL,
		"Insecure": helper.GetConfig("network").NDFC.Insecure,
	}

	tfConfig := new(string)
	stepCount := new(int)
	*stepCount = 0
	networkRsc := new(resource_networks.NDFCNetworksModel)
	vrfRsc := new(resource_vrf_bulk.NDFCVrfBulkModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "network") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create networks with mixed keys - 1st attachment uses serial, 2nd uses IP
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					// Create VRF first
					helper.GenerateSingleVrfObject(&vrfRsc, helper.GetConfig("network").NDFC.VrfPrefix, helper.GetConfig("network").NDFC.Fabric, 1, true, false, false, helper.GetConfig("network").NDFC.Switches)

					// Create networks with mixed attachment keys
					networkRsc.FabricName = helper.GetConfig("network").NDFC.Fabric
					networkRsc.DeployAllAttachments = false
					networkRsc.Networks = make(map[string]resource_networks.NDFCNetworksValue)

					for i := 1; i <= 3; i++ {
						nwName := helper.GetConfig("network").NDFC.NetPrefix + fmt.Sprintf("%d", i)
						nw := resource_networks.NDFCNetworksValue{}
						nw.NetworkName = nwName
						nw.VrfName = helper.GetConfig("network").NDFC.VrfPrefix + "1"
						nwId := int64(30000 + i)
						nw.NetworkId = &nwId
						vlanId := types.Int64Custom(2500 + i)
						nw.NetworkTemplateConfig.VlanId = &vlanId
						nw.DeployAttachments = true
						nw.Attachments = make(map[string]rna.NDFCAttachmentsValue)

						// First attachment uses serial number as key
						attach1 := rna.NDFCAttachmentsValue{}
						nw.Attachments[helper.GetConfig("network").NDFC.Switches[0]] = attach1

						// Second attachment uses IP address as key
						attach2 := rna.NDFCAttachmentsValue{}
						nw.Attachments[helper.GetConfig("network").NDFC.SwitchIP[1]] = attach2

						networkRsc.Networks[nwName] = nw
					}

					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},
			// Step 2: Add 3rd attachment using IP key
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					for nwName, nw := range networkRsc.Networks {
						attach3 := rna.NDFCAttachmentsValue{}
						nw.Attachments[helper.GetConfig("network").NDFC.SwitchIP[2]] = attach3
						networkRsc.Networks[nwName] = nw
					}

					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},
			// Step 3: Remove the serial-based attachment, keep IP-based ones
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					for nwName, nw := range networkRsc.Networks {
						delete(nw.Attachments, helper.GetConfig("network").NDFC.Switches[0])
						networkRsc.Networks[nwName] = nw
					}

					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},
		},
	})
}

// TestAccNetworkResourceIPKeyAddRemoveNetworksWithDeploy tests add/remove networks with deploy using IP-based keys
func TestAccNetworksResourceIPKeyAddRemoveNetworksWithDeploy(t *testing.T) {
	x := &map[string]string{
		"RscType":  ndfc.ResourceNetworks,
		"RscName":  "network_test",
		"User":     helper.GetConfig("network").NDFC.User,
		"Password": helper.GetConfig("network").NDFC.Password,
		"Host":     helper.GetConfig("network").NDFC.URL,
		"Insecure": helper.GetConfig("network").NDFC.Insecure,
	}

	tfConfig := new(string)
	stepCount := new(int)
	*stepCount = 0
	networkRsc := new(resource_networks.NDFCNetworksModel)
	vrfRsc := new(resource_vrf_bulk.NDFCVrfBulkModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "network") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create 3 networks with 2 attachments each using IP keys (with deploy)
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateSingleVrfObject(&vrfRsc, helper.GetConfig("network").NDFC.VrfPrefix, helper.GetConfig("network").NDFC.Fabric, 1, false, false, true, helper.GetConfig("network").NDFC.Switches)
					helper.GenerateNetworksObject(&networkRsc, helper.GetConfig("network").NDFC.Fabric,
						3, false, true, false, helper.GetConfig("network").NDFC.VrfPrefix+"1", []string{helper.GetConfig("network").NDFC.SwitchIP[0], helper.GetConfig("network").NDFC.SwitchIP[1]})
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 2: Remove the 3rd network entirely
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.DeleteNetworks(&networkRsc, 3, 3)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 3: Add the 3rd network back using IP keys
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateNetworksObject(&networkRsc, helper.GetConfig("network").NDFC.Fabric,
						3, false, true, false, helper.GetConfig("network").NDFC.VrfPrefix+"1", []string{helper.GetConfig("network").NDFC.SwitchIP[0], helper.GetConfig("network").NDFC.SwitchIP[1]})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 4: Remove the 1st network
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.DeleteNetworks(&networkRsc, 1, 1)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 5: Add both 1st and 3rd networks back using IP keys
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateNetworksObject(&networkRsc, helper.GetConfig("network").NDFC.Fabric,
						3, false, true, false, helper.GetConfig("network").NDFC.VrfPrefix+"1", []string{helper.GetConfig("network").NDFC.SwitchIP[0], helper.GetConfig("network").NDFC.SwitchIP[1]})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},
		}},
	)
}

// TestAccNetworkResourceIPKeySimultaneousModifications tests complex updates with IP-based keys
func TestAccNetworksResourceIPKeySimultaneousModifications(t *testing.T) {
	x := &map[string]string{
		"RscType":  ndfc.ResourceNetworks,
		"RscName":  "network_test",
		"User":     helper.GetConfig("network").NDFC.User,
		"Password": helper.GetConfig("network").NDFC.Password,
		"Host":     helper.GetConfig("network").NDFC.URL,
		"Insecure": helper.GetConfig("network").NDFC.Insecure,
	}

	tfConfig := new(string)
	stepCount := new(int)
	*stepCount = 0
	networkRsc := new(resource_networks.NDFCNetworksModel)
	vrfRsc := new(resource_vrf_bulk.NDFCVrfBulkModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "network") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create 5 networks with 2 attachments each using IP keys
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateSingleVrfObject(&vrfRsc, helper.GetConfig("network").NDFC.VrfPrefix, helper.GetConfig("network").NDFC.Fabric, 1, false, false, true, helper.GetConfig("network").NDFC.Switches)
					helper.GenerateNetworksObject(&networkRsc, helper.GetConfig("network").NDFC.Fabric,
						5, false, false, false, helper.GetConfig("network").NDFC.VrfPrefix+"1", []string{helper.GetConfig("network").NDFC.SwitchIP[0], helper.GetConfig("network").NDFC.SwitchIP[1]})
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 2: Add 3rd attachment with deploy=false, toggle existing 2 attachments to deploy=true
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Add 3rd attachment and enable deployment on existing attachments for all networks
					for i := 1; i <= 5; i++ {
						nwName := helper.GetConfig("network").NDFC.NetPrefix + fmt.Sprintf("%d", i)
						if nw, ok := networkRsc.Networks[nwName]; ok {
							// Toggle existing 2 attachments to deploy=true
							for ipKey, attach := range nw.Attachments {
								attach.DeployThisAttachment = true
								nw.Attachments[ipKey] = attach
							}
							// Add 3rd attachment with deploy=false using IP key
							thirdAttach := rna.NDFCAttachmentsValue{}
							thirdAttach.DeployThisAttachment = false
							nw.Attachments[helper.GetConfig("network").NDFC.SwitchIP[2]] = thirdAttach
							networkRsc.Networks[nwName] = nw
						}
					}
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 3: Toggle 3rd attachment to deploy=true
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Toggle 3rd attachment to deploy=true for all networks
					for i := 1; i <= 5; i++ {
						nwName := helper.GetConfig("network").NDFC.NetPrefix + fmt.Sprintf("%d", i)
						if nw, ok := networkRsc.Networks[nwName]; ok {
							thirdIP := helper.GetConfig("network").NDFC.SwitchIP[2]
							if attach, ok := nw.Attachments[thirdIP]; ok {
								attach.DeployThisAttachment = true
								nw.Attachments[thirdIP] = attach
								networkRsc.Networks[nwName] = nw
							}
						}
					}
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			// Step 4: Complex simultaneous modifications across multiple networks
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					// 1. Modify switch_ports on attachment 2 (SwitchIP[1]) for networks 1-3
					for i := 1; i <= 3; i++ {
						nwName := helper.GetConfig("network").NDFC.NetPrefix + fmt.Sprintf("%d", i)
						if nw, ok := networkRsc.Networks[nwName]; ok {
							secondIP := helper.GetConfig("network").NDFC.SwitchIP[1]
							if attach, ok := nw.Attachments[secondIP]; ok {
								attach.SwitchPorts = types.CSVString{fmt.Sprintf("Ethernet1/%d", 10+i)}
								nw.Attachments[secondIP] = attach
								networkRsc.Networks[nwName] = nw
							}
						}
					}

					// 2. Delete network 4 entirely
					nwName4 := helper.GetConfig("network").NDFC.NetPrefix + "4"
					delete(networkRsc.Networks, nwName4)

					// 3. From network 5, remove 2nd and 3rd attachments
					nwName5 := helper.GetConfig("network").NDFC.NetPrefix + "5"
					if nw5, ok := networkRsc.Networks[nwName5]; ok {
						delete(nw5.Attachments, helper.GetConfig("network").NDFC.SwitchIP[1])
						delete(nw5.Attachments, helper.GetConfig("network").NDFC.SwitchIP[2])
						networkRsc.Networks[nwName5] = nw5
					}

					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},
		}},
	)
}

// TestAccNetworkResourceIPKeyAttachLevelDeploy tests attachment-level deploy with IP-based keys
func TestAccNetworksResourceIPKeyAttachLevelDeploy(t *testing.T) {

	x := &map[string]string{
		"RscType":  ndfc.ResourceNetworks,
		"RscName":  "network_test",
		"User":     helper.GetConfig("network").NDFC.User,
		"Password": helper.GetConfig("network").NDFC.Password,
		"Host":     helper.GetConfig("network").NDFC.URL,
		"Insecure": helper.GetConfig("network").NDFC.Insecure,
	}
	networkRsc := new(resource_networks.NDFCNetworksModel)
	vrfRsc := new(resource_vrf_bulk.NDFCVrfBulkModel)
	stepCount := new(int)
	*stepCount = 0

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "network") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					t.Logf(logTestStep, t.Name(), *stepCount)
				},
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					tf_config := new(string)
					helper.GenerateSingleVrfObject(&vrfRsc, helper.GetConfig("network").NDFC.VrfPrefix, helper.GetConfig("network").NDFC.Fabric, 1, false, false, true, helper.GetConfig("network").NDFC.Switches)
					helper.GenerateNetworksObject(&networkRsc, helper.GetConfig("network").NDFC.Fabric,
						5, false, false, true, helper.GetConfig("network").NDFC.VrfPrefix+"1", []string{helper.GetConfig("network").NDFC.SwitchIP[0], helper.GetConfig("network").NDFC.SwitchIP[1]})
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},
		},
	})
}
