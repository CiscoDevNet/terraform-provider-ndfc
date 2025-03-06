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
						"secondary_gateway_1":    "192.168.100.1/24",
						"secondary_gateway_2":    "192.168.101.1/24",
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
