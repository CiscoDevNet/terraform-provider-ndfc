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
	"terraform-provider-ndfc/internal/provider/ndfc"
	"terraform-provider-ndfc/internal/provider/resources/resource_vrf_bulk"
	helper "terraform-provider-ndfc/internal/provider/testing"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

/*
When a test is ran,
Terraform runs plan, apply, refresh, and then final plan for each TestStep in the TestCase.
*/

const vrfTestRscName = "ndfc_vrfs.vrf_test"
const logTestStep = "Starting Test Step %s_%d"

func TestAccVRFResourceCRUD(t *testing.T) {

	x := &map[string]string{
		"RscType":  ndfc.ResourceVrfBulk,
		"RscName":  "vrf_test",
		"User":     helper.GetConfig("vrf").NDFC.User,
		"Password": helper.GetConfig("vrf").NDFC.Password,
		"Host":     helper.GetConfig("vrf").NDFC.URL,
		"Insecure": helper.GetConfig("vrf").NDFC.Insecure,
	}

	tfConfig := new(string)
	*tfConfig = `provider "ndfc" {
		host     = "https://"
		username = "admin"
		password = "admin!@#"
		domain   = "example.com"
		insecure = true
		}
		resource ndfc_vrfs "vrf_test" {
			fabric_name = "dummy"
		}`

	stepCount := new(int)
	*stepCount = 0
	// Create a new instance of the NDFC client
	vrfScaledBulk := new(resource_vrf_bulk.NDFCVrfBulkModel)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "vrf") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.GenerateVrfBulkObject(&vrfScaledBulk, helper.GetConfig("vrf").NDFC.Fabric,
						10, false, false, false, nil)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},
			{
				//Add 10 more VRFs
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.IncreaseVrfCount(&vrfScaledBulk,
						10, false, false, false, nil)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.DeleteVrfs(&vrfScaledBulk,
						11, 20)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},
			{

				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					(*x)["RscName"] = "vrf_test,vrf_test_1"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk, vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				ExpectError: regexp.MustCompile(".*VRFs exist.*"),
			},
			{
				//Modify Few Params in VRFs
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					tfConfig := new(string)
					helper.ModifyVrfBulkObject(&vrfScaledBulk, 1, map[string]interface{}{
						"vlan_id":              100,
						"vrf_description":      "test",
						"loopback_routing_tag": 2459,
						"mtu":                  9100,
						"max_bgp_paths":        2,
						"ipv6_link_local":      "true",
					})
					helper.ModifyVrfBulkObject(&vrfScaledBulk, 10, map[string]interface{}{
						"vlan_id":              110,
						"vrf_description":      "test",
						"loopback_routing_tag": 2459,
						"mtu":                  9100,
						"max_bgp_paths":        2,
						"ipv6_link_local":      "true",
					})
					(*x)["RscName"] = "vrf_test"

					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},
		}})
}

func TestAccVRFResourceAttachmentCRUD(t *testing.T) {

	x := &map[string]string{
		"RscType":  ndfc.ResourceVrfBulk,
		"RscName":  "vrf_test",
		"User":     helper.GetConfig("vrf").NDFC.User,
		"Password": helper.GetConfig("vrf").NDFC.Password,
		"Host":     helper.GetConfig("vrf").NDFC.URL,
		"Insecure": helper.GetConfig("vrf").NDFC.Insecure,
	}
	vrfScaledBulk := new(resource_vrf_bulk.NDFCVrfBulkModel)
	stepCount := new(int)
	*stepCount = 0

	resource.Test(t, resource.TestCase{

		PreCheck:                 func() { testAccPreCheck(t, "vrf") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			//Create VRFs with 2 attachments _Attach
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					tfConfig := new(string)
					helper.GenerateVrfBulkObject(&vrfScaledBulk, helper.GetConfig("vrf").NDFC.Fabric,
						20, false, false, false, []string{helper.GetConfig("vrf").NDFC.Switches[0], helper.GetConfig("vrf").NDFC.Switches[1]})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},
			{ // Remove both attachments _detach
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					tfConfig := new(string)
					helper.VrfAttachmentsMod(&vrfScaledBulk, 1, len(vrfScaledBulk.Vrfs), nil, "", nil)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},
			{
				// Add 2 attachments to all VRFs
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					tfConfig := new(string)
					helper.VrfAttachmentsMod(&vrfScaledBulk, 1, len(vrfScaledBulk.Vrfs), []string{helper.GetConfig("vrf").NDFC.Switches[0], helper.GetConfig("vrf").NDFC.Switches[1]}, "", nil)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},
			{
				// Add 3rd attachment half of them
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					tfConfig := new(string)
					helper.VrfAttachmentsMod(&vrfScaledBulk, 1, len(vrfScaledBulk.Vrfs)/2, helper.GetConfig("vrf").NDFC.Switches, "", nil)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},
			{
				// Add 3rd attachment remaining  half
				// Remove 3rd from others
				Config: func() string {
					tfConfig := new(string)
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.VrfAttachmentsMod(&vrfScaledBulk, 1, len(vrfScaledBulk.Vrfs)/2, []string{helper.GetConfig("vrf").NDFC.Switches[0], helper.GetConfig("vrf").NDFC.Switches[1]}, "", nil)
					helper.VrfAttachmentsMod(&vrfScaledBulk, (len(vrfScaledBulk.Vrfs)/2)+1, len(vrfScaledBulk.Vrfs)/2, helper.GetConfig("vrf").NDFC.Switches, "", nil)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},
			{
				//Modify params
				Config: func() string {
					tfConfig := new(string)
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.VrfAttachmentsMod(&vrfScaledBulk, 1, 1, helper.GetConfig("vrf").NDFC.Switches, helper.GetConfig("vrf").NDFC.Switches[2], map[string]interface{}{
						"vlan":          3001,
						"loopback_id":   1001,
						"loopback_ipv4": "10.1.1.1",
						"loopback_ipv6": "2001:db8::68",
					})

					helper.VrfAttachmentsMod(&vrfScaledBulk, 10, 10, helper.GetConfig("vrf").NDFC.Switches, helper.GetConfig("vrf").NDFC.Switches[2], map[string]interface{}{
						"vlan":          3010,
						"loopback_id":   1010,
						"loopback_ipv4": "10.1.1.10",
						"loopback_ipv6": "2001:db8::610",
					})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},
		}})
}

// GLOBAL_DEPLOY_TEST Add 10 VRFs with 2 attachments, and global deployment
func TestAccVRFResourceGlobalDeploy(t *testing.T) {

	x := &map[string]string{
		"RscType":  ndfc.ResourceVrfBulk,
		"RscName":  "vrf_test",
		"User":     helper.GetConfig("vrf").NDFC.User,
		"Password": helper.GetConfig("vrf").NDFC.Password,
		"Host":     helper.GetConfig("vrf").NDFC.URL,
		"Insecure": helper.GetConfig("vrf").NDFC.Insecure,
	}
	vrfScaledBulk := new(resource_vrf_bulk.NDFCVrfBulkModel)
	stepCount := new(int)
	*stepCount = 0

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "vrf") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					t.Logf(logTestStep, t.Name(), *stepCount)
				},
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					tfConfig := new(string)
					helper.GenerateVrfBulkObject(&vrfScaledBulk, helper.GetConfig("vrf").NDFC.Fabric,
						5, true, false, false, []string{helper.GetConfig("vrf").NDFC.Switches[0], helper.GetConfig("vrf").NDFC.Switches[1]})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},
			{
				// Add 3rd Attachment
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					tfConfig := new(string)
					helper.VrfAttachmentsMod(&vrfScaledBulk, 1, len(vrfScaledBulk.Vrfs), helper.GetConfig("vrf").NDFC.Switches, "", nil)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},
		},
	})
}

// GLOBAL_DEPLOY_TEST Add 10 VRFs with 2 attachments, VRF level deployment
func TestAccVRFResourceVrfLevelDeploy(t *testing.T) {

	x := &map[string]string{
		"RscType":  ndfc.ResourceVrfBulk,
		"RscName":  "vrf_test",
		"User":     helper.GetConfig("vrf").NDFC.User,
		"Password": helper.GetConfig("vrf").NDFC.Password,
		"Host":     helper.GetConfig("vrf").NDFC.URL,
		"Insecure": helper.GetConfig("vrf").NDFC.Insecure,
	}
	vrfScaledBulk := new(resource_vrf_bulk.NDFCVrfBulkModel)
	stepCount := new(int)
	*stepCount = 0

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "vrf") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					t.Logf(logTestStep, t.Name(), *stepCount)
				},
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					tfConfig := new(string)
					helper.GenerateVrfBulkObject(&vrfScaledBulk, helper.GetConfig("vrf").NDFC.Fabric,
						10, false, true, false, []string{helper.GetConfig("vrf").NDFC.Switches[0], helper.GetConfig("vrf").NDFC.Switches[1]})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},
			{
				// Add 3rd Attachment
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					tfConfig := new(string)
					helper.VrfAttachmentsMod(&vrfScaledBulk, 1, len(vrfScaledBulk.Vrfs), helper.GetConfig("vrf").NDFC.Switches, "", nil)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},
		},
	})
}

func TestAccVRFResourceVrfAttachLevelDeploy(t *testing.T) {

	x := &map[string]string{
		"RscType":  ndfc.ResourceVrfBulk,
		"RscName":  "vrf_test",
		"User":     helper.GetConfig("vrf").NDFC.User,
		"Password": helper.GetConfig("vrf").NDFC.Password,
		"Host":     helper.GetConfig("vrf").NDFC.URL,
		"Insecure": helper.GetConfig("vrf").NDFC.Insecure,
	}
	vrfScaledBulk := new(resource_vrf_bulk.NDFCVrfBulkModel)
	stepCount := new(int)
	*stepCount = 0

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "vrf") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					t.Logf(logTestStep, t.Name(), *stepCount)
				},
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					tfConfig := new(string)
					helper.GenerateVrfBulkObject(&vrfScaledBulk, helper.GetConfig("vrf").NDFC.Fabric,
						5, false, false, true, []string{helper.GetConfig("vrf").NDFC.Switches[0], helper.GetConfig("vrf").NDFC.Switches[1]})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},
		},
	})
}

func TestAccVRFResourceMultiResourceWithDeploy(t *testing.T) {

	x := &map[string]string{
		"RscType":  ndfc.ResourceVrfBulk,
		"RscName":  "",
		"User":     helper.GetConfig("vrf").NDFC.User,
		"Password": helper.GetConfig("vrf").NDFC.Password,
		"Host":     helper.GetConfig("vrf").NDFC.URL,
		"Insecure": helper.GetConfig("vrf").NDFC.Insecure,
	}
	var vrfScaledBulk []*resource_vrf_bulk.NDFCVrfBulkModel
	stepCount := new(int)
	*stepCount = 0

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "vrf") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			//Create VRFs with 2 attachments _Attach
			{
				PreConfig: func() {
					t.Logf(logTestStep, t.Name(), *stepCount)
				},
				Config: func() string {
					vrfScaledBulk = make([]*resource_vrf_bulk.NDFCVrfBulkModel, 5)
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					tfConfig := new(string)
					for i := 0; i < 5; i++ {
						vrfScaledBulk[i] = new(resource_vrf_bulk.NDFCVrfBulkModel)
						helper.GenerateSingleVrfObject(&(vrfScaledBulk[i]), helper.GetConfig("vrf").NDFC.VrfPrefix, helper.GetConfig("vrf").NDFC.Fabric,
							i+1, false, false, true, []string{helper.GetConfig("vrf").NDFC.Switches[0], helper.GetConfig("vrf").NDFC.Switches[1]})
						if i == 0 {
							(*x)["RscName"] = fmt.Sprintf("vrf_test_%d", i+1)
						} else {
							(*x)["RscName"] = fmt.Sprintf("%s,vrf_test_%d", (*x)["RscName"], i+1)
						}
					}
					helper.GetVRFTFConfigWithMultipleResource(tName, *x, &vrfScaledBulk, &tfConfig)
					return *tfConfig
				}(),
				Check: func() resource.TestCheckFunc {
					var checks []resource.TestCheckFunc
					for i := 0; i < len(vrfScaledBulk); i++ {
						checks = append(checks, VrfBulkModelHelperStateCheck(fmt.Sprintf("ndfc_vrfs.vrf_test_%d", i+1), *vrfScaledBulk[i], path.Empty())...)
					}
					return resource.ComposeTestCheckFunc(checks...)
				}(),
			},
		}})
}

/*
func testAccCheckVrfBulkResourceDestroy(vrfBulk *resource_vrf_bulk.VrfBulkModel) resource.TestCheckFunc {
	return nil
}

func testGenerateVrfMultipleResource(count int, vrfName string, rscName string) string {
	x := &map[string]string{
		"RscType":  ndfc.ResourceVrfBulk,
		"RscName":  "vrf_test",
		"User":     "admin",
		"Password": "admin!@#",
		"Host":     "https://10.78.210.161",
		"Insecure": "true",
	}
	vrfScaledBulk := make([]*resource_vrf_bulk.NDFCVrfBulkModel, count)
	tfConfig := new(string)
	for i := 0; i < count; i++ {
		vrfScaledBulk[i] = new(resource_vrf_bulk.NDFCVrfBulkModel)
		helper.GenerateSingleVrfObject(&(vrfScaledBulk[i]), vrfName, helper.GetConfig("vrf").NDFC.Fabric,
			i+1, false, false, true, []string{helper.GetConfig("vrf").NDFC.Switches[0], helper.GetConfig("vrf").NDFC.Switches[1]})
		if (*x)["RscName"] == "" {
			(*x)["RscName"] = fmt.Sprintf("%s_%d", rscName, i+1)
		} else {
			(*x)["RscName"] = (*x)["RscName"] + "," + fmt.Sprintf("vrf_test_%d", i+1)
		}
	}
	helper.GetVRFTFConfigWithMultipleResource("multiple_rsc_", *x, &vrfScaledBulk, &tfConfig)
	return *tfConfig
}
*/
