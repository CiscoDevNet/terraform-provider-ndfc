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
	"strconv"
	"terraform-provider-ndfc/internal/provider/ndfc"
	"terraform-provider-ndfc/internal/provider/resources/resource_vrf_attachments"
	"terraform-provider-ndfc/internal/provider/resources/resource_vrf_bulk"
	helper "terraform-provider-ndfc/internal/provider/testing"
	"terraform-provider-ndfc/internal/provider/types"
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
		"Password": "password",
		"Host":     "ndfc.com",
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

func TestAccVRFResourceAddRemoveVrfs(t *testing.T) {
	x := &map[string]string{
		"RscType":  ndfc.ResourceVrfBulk,
		"RscName":  "vrf_test",
		"User":     helper.GetConfig("vrf").NDFC.User,
		"Password": helper.GetConfig("vrf").NDFC.Password,
		"Host":     helper.GetConfig("vrf").NDFC.URL,
		"Insecure": helper.GetConfig("vrf").NDFC.Insecure,
	}

	tfConfig := new(string)
	stepCount := new(int)
	*stepCount = 0
	vrfScaledBulk := new(resource_vrf_bulk.NDFCVrfBulkModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "vrf") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create 3 VRFs with 2 attachments each
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateVrfBulkObject(&vrfScaledBulk, helper.GetConfig("vrf").NDFC.Fabric,
						3, false, false, false, []string{helper.GetConfig("vrf").NDFC.Switches[0], helper.GetConfig("vrf").NDFC.Switches[1]})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 2: Remove the 3rd VRF entirely
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.DeleteVrfs(&vrfScaledBulk, 3, 3)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 3: Add the 3rd VRF back
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Regenerate all 3 VRFs to add VRF 3 back
					helper.GenerateVrfBulkObject(&vrfScaledBulk, helper.GetConfig("vrf").NDFC.Fabric,
						3, false, false, false, []string{helper.GetConfig("vrf").NDFC.Switches[0], helper.GetConfig("vrf").NDFC.Switches[1]})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 4: Remove the 1st VRF
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.DeleteVrfs(&vrfScaledBulk, 1, 1)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 5: Remove the 2nd VRF
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.DeleteVrfs(&vrfScaledBulk, 2, 2)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 6: Add both 1st and 2nd VRFs back
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Regenerate all 3 VRFs to add VRFs 1 and 2 back
					helper.GenerateVrfBulkObject(&vrfScaledBulk, helper.GetConfig("vrf").NDFC.Fabric,
						3, false, false, false, []string{helper.GetConfig("vrf").NDFC.Switches[0], helper.GetConfig("vrf").NDFC.Switches[1]})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 7: Remove VRFs 2 and 3
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.DeleteVrfs(&vrfScaledBulk, 2, 3)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},
		}},
	)
}

func TestAccVRFResourceAddRemoveVrfsWithDeploy(t *testing.T) {
	x := &map[string]string{
		"RscType":  ndfc.ResourceVrfBulk,
		"RscName":  "vrf_test",
		"User":     helper.GetConfig("vrf").NDFC.User,
		"Password": helper.GetConfig("vrf").NDFC.Password,
		"Host":     helper.GetConfig("vrf").NDFC.URL,
		"Insecure": helper.GetConfig("vrf").NDFC.Insecure,
	}

	tfConfig := new(string)
	stepCount := new(int)
	*stepCount = 0
	vrfScaledBulk := new(resource_vrf_bulk.NDFCVrfBulkModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "vrf") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create 3 VRFs with 2 attachments each (with deploy)
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateVrfBulkObject(&vrfScaledBulk, helper.GetConfig("vrf").NDFC.Fabric,
						3, false, false, true, []string{helper.GetConfig("vrf").NDFC.Switches[0], helper.GetConfig("vrf").NDFC.Switches[1]})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 2: Remove the 3rd VRF entirely
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.DeleteVrfs(&vrfScaledBulk, 3, 3)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 3: Add the 3rd VRF back
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Regenerate all 3 VRFs to add VRF 3 back
					helper.GenerateVrfBulkObject(&vrfScaledBulk, helper.GetConfig("vrf").NDFC.Fabric,
						3, false, false, true, []string{helper.GetConfig("vrf").NDFC.Switches[0], helper.GetConfig("vrf").NDFC.Switches[1]})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 4: Remove the 1st VRF
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.DeleteVrfs(&vrfScaledBulk, 1, 1)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 5: Remove the 2nd VRF
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.DeleteVrfs(&vrfScaledBulk, 2, 2)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 6: Add both 1st and 2nd VRFs back
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Regenerate all 3 VRFs to add VRFs 1 and 2 back
					helper.GenerateVrfBulkObject(&vrfScaledBulk, helper.GetConfig("vrf").NDFC.Fabric,
						3, false, false, true, []string{helper.GetConfig("vrf").NDFC.Switches[0], helper.GetConfig("vrf").NDFC.Switches[1]})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 7: Remove VRFs 2 and 3
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.DeleteVrfs(&vrfScaledBulk, 2, 3)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},
		}},
	)
}

func TestAccVRFResourceAddRemoveVrfsWithVrfDeployFlag(t *testing.T) {
	x := &map[string]string{
		"RscType":  ndfc.ResourceVrfBulk,
		"RscName":  "vrf_test",
		"User":     helper.GetConfig("vrf").NDFC.User,
		"Password": helper.GetConfig("vrf").NDFC.Password,
		"Host":     helper.GetConfig("vrf").NDFC.URL,
		"Insecure": helper.GetConfig("vrf").NDFC.Insecure,
	}

	tfConfig := new(string)
	stepCount := new(int)
	*stepCount = 0
	vrfScaledBulk := new(resource_vrf_bulk.NDFCVrfBulkModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "vrf") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create 3 VRFs with 2 attachments each (with VRF-level deploy)
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateVrfBulkObject(&vrfScaledBulk, helper.GetConfig("vrf").NDFC.Fabric,
						3, false, true, false, []string{helper.GetConfig("vrf").NDFC.Switches[0], helper.GetConfig("vrf").NDFC.Switches[1]})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 4: Remove the 1st VRF
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.DeleteVrfs(&vrfScaledBulk, 1, 1)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 6: Add both 1st and 2nd VRFs back
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Regenerate all 3 VRFs to add VRFs 1 and 2 back
					helper.GenerateVrfBulkObject(&vrfScaledBulk, helper.GetConfig("vrf").NDFC.Fabric,
						3, false, true, false, []string{helper.GetConfig("vrf").NDFC.Switches[0], helper.GetConfig("vrf").NDFC.Switches[1]})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 7: Remove VRFs 2 and 3
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.DeleteVrfs(&vrfScaledBulk, 2, 3)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},
		}},
	)
}

func TestAccVRFResourceAddRemoveVrfsWithGlobalDeployFlag(t *testing.T) {
	x := &map[string]string{
		"RscType":  ndfc.ResourceVrfBulk,
		"RscName":  "vrf_test",
		"User":     helper.GetConfig("vrf").NDFC.User,
		"Password": helper.GetConfig("vrf").NDFC.Password,
		"Host":     helper.GetConfig("vrf").NDFC.URL,
		"Insecure": helper.GetConfig("vrf").NDFC.Insecure,
	}

	tfConfig := new(string)
	stepCount := new(int)
	*stepCount = 0
	vrfScaledBulk := new(resource_vrf_bulk.NDFCVrfBulkModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "vrf") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create 3 VRFs with 2 attachments each (with global deploy)
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateVrfBulkObject(&vrfScaledBulk, helper.GetConfig("vrf").NDFC.Fabric,
						3, true, false, false, []string{helper.GetConfig("vrf").NDFC.Switches[0], helper.GetConfig("vrf").NDFC.Switches[1]})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 4: Remove the 1st VRF
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.DeleteVrfs(&vrfScaledBulk, 1, 1)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 6: Add both 1st and 2nd VRFs back
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Regenerate all 3 VRFs to add VRFs 1 and 2 back
					helper.GenerateVrfBulkObject(&vrfScaledBulk, helper.GetConfig("vrf").NDFC.Fabric,
						3, true, false, false, []string{helper.GetConfig("vrf").NDFC.Switches[0], helper.GetConfig("vrf").NDFC.Switches[1]})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 7: Remove VRFs 2 and 3
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.DeleteVrfs(&vrfScaledBulk, 2, 3)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},
		}},
	)
}

func TestAccVRFResourceAddRemoveVrfsComboDeploy(t *testing.T) {
	x := &map[string]string{
		"RscType":  ndfc.ResourceVrfBulk,
		"RscName":  "vrf_test",
		"User":     helper.GetConfig("vrf").NDFC.User,
		"Password": helper.GetConfig("vrf").NDFC.Password,
		"Host":     helper.GetConfig("vrf").NDFC.URL,
		"Insecure": helper.GetConfig("vrf").NDFC.Insecure,
	}

	tfConfig := new(string)
	stepCount := new(int)
	*stepCount = 0
	vrfScaledBulk := new(resource_vrf_bulk.NDFCVrfBulkModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "vrf") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create 3 VRFs with 2 attachments each (deploy OFF)
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateVrfBulkObject(&vrfScaledBulk, helper.GetConfig("vrf").NDFC.Fabric,
						3, false, false, false, []string{helper.GetConfig("vrf").NDFC.Switches[0], helper.GetConfig("vrf").NDFC.Switches[1]})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 2: Delete VRF 3, deploy remaining ones using deploy_this_attachment
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.DeleteVrfs(&vrfScaledBulk, 3, 3)
					// Turn on deploy_this_attachment for VRFs 1 and 2
					for i := 1; i <= 2; i++ {
						vrfName := helper.GetConfig("vrf").NDFC.VrfPrefix + fmt.Sprintf("%d", i)
						if vrf, ok := vrfScaledBulk.Vrfs[vrfName]; ok {
							for serial, attach := range vrf.AttachList {
								attach.DeployThisAttachment = true
								vrf.AttachList[serial] = attach
							}
							vrfScaledBulk.Vrfs[vrfName] = vrf
						}
					}
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 3: Add VRF 3 back with deploy OFF
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Regenerate all 3 VRFs - VRFs 1,2 keep deploy=true from step 2
					helper.GenerateVrfBulkObject(&vrfScaledBulk, helper.GetConfig("vrf").NDFC.Fabric,
						3, false, false, false, []string{helper.GetConfig("vrf").NDFC.Switches[0], helper.GetConfig("vrf").NDFC.Switches[1]})
					// Restore deploy_this_attachment=true for VRFs 1 and 2
					for i := 1; i <= 2; i++ {
						vrfName := helper.GetConfig("vrf").NDFC.VrfPrefix + fmt.Sprintf("%d", i)
						if vrf, ok := vrfScaledBulk.Vrfs[vrfName]; ok {
							for serial, attach := range vrf.AttachList {
								attach.DeployThisAttachment = true
								vrf.AttachList[serial] = attach
							}
							vrfScaledBulk.Vrfs[vrfName] = vrf
						}
					}
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 4: Turn on deploy flag for VRF 3
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Turn on deploy_this_attachment for VRF 3
					vrfName := helper.GetConfig("vrf").NDFC.VrfPrefix + "3"
					if vrf, ok := vrfScaledBulk.Vrfs[vrfName]; ok {
						for serial, attach := range vrf.AttachList {
							attach.DeployThisAttachment = true
							vrf.AttachList[serial] = attach
						}
						vrfScaledBulk.Vrfs[vrfName] = vrf
					}
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 5: Remove attachments from VRF 3
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.VrfAttachmentsMod(&vrfScaledBulk, 3, 3, nil, "", nil)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 6: Add attachments back to VRF 3 with deploy=true
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Add attachments back
					vrfName := helper.GetConfig("vrf").NDFC.VrfPrefix + "3"
					if vrf, ok := vrfScaledBulk.Vrfs[vrfName]; ok {
						vrf.AttachList = make(map[string]resource_vrf_attachments.NDFCAttachListValue)
						for _, serial := range []string{helper.GetConfig("vrf").NDFC.Switches[0], helper.GetConfig("vrf").NDFC.Switches[1]} {
							attach := resource_vrf_attachments.NDFCAttachListValue{}
							attach.SerialNumber = serial
							attach.DeployThisAttachment = true
							vrf.AttachList[serial] = attach
						}
						vrfScaledBulk.Vrfs[vrfName] = vrf
					}
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 7: Remove VRFs 2 and 3
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.DeleteVrfs(&vrfScaledBulk, 2, 3)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},
		}},
	)
}

// TestAccVRFResourceSimultaneousModifications tests complex updates where VRF properties,
// attachments, and deployment flags are all modified in a single apply operation
func TestAccVRFResourceSimultaneousModifications(t *testing.T) {
	x := &map[string]string{
		"RscType":  ndfc.ResourceVrfBulk,
		"RscName":  "vrf_test",
		"User":     helper.GetConfig("vrf").NDFC.User,
		"Password": helper.GetConfig("vrf").NDFC.Password,
		"Host":     helper.GetConfig("vrf").NDFC.URL,
		"Insecure": helper.GetConfig("vrf").NDFC.Insecure,
	}

	tfConfig := new(string)
	stepCount := new(int)
	*stepCount = 0
	vrfScaledBulk := new(resource_vrf_bulk.NDFCVrfBulkModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "vrf") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create 10 VRFs with 2 attachments each, all deployment flags false
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateVrfBulkObject(&vrfScaledBulk, helper.GetConfig("vrf").NDFC.Fabric,
						10, false, false, false, []string{helper.GetConfig("vrf").NDFC.Switches[0], helper.GetConfig("vrf").NDFC.Switches[1]})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 2: Add 3rd attachment with deploy=false, toggle existing 2 attachments to deploy=true
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Add 3rd attachment and enable deployment on existing attachments for all VRFs
					for i := 1; i <= 10; i++ {
						vrfName := helper.GetConfig("vrf").NDFC.VrfPrefix + strconv.Itoa(i)
						if vrf, ok := vrfScaledBulk.Vrfs[vrfName]; ok {
							// Toggle existing 2 attachments to deploy=true
							for serial, attach := range vrf.AttachList {
								attach.DeployThisAttachment = true
								vrf.AttachList[serial] = attach
							}
							// Add 3rd attachment with deploy=false
							thirdAttach := resource_vrf_attachments.NDFCAttachListValue{}
							thirdAttach.SerialNumber = helper.GetConfig("vrf").NDFC.Switches[2]
							thirdAttach.DeployThisAttachment = false
							vrf.AttachList[helper.GetConfig("vrf").NDFC.Switches[2]] = thirdAttach
							vrfScaledBulk.Vrfs[vrfName] = vrf
						}
					}
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 3: Toggle 3rd attachment to deploy=true
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Toggle 3rd attachment to deploy=true for all VRFs
					for i := 1; i <= 10; i++ {
						vrfName := helper.GetConfig("vrf").NDFC.VrfPrefix + strconv.Itoa(i)
						if vrf, ok := vrfScaledBulk.Vrfs[vrfName]; ok {
							thirdSerial := helper.GetConfig("vrf").NDFC.Switches[2]
							if attach, ok := vrf.AttachList[thirdSerial]; ok {
								attach.DeployThisAttachment = true
								vrf.AttachList[thirdSerial] = attach
								vrfScaledBulk.Vrfs[vrfName] = vrf
							}
						}
					}
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 4: Complex simultaneous modifications across multiple VRFs
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					// 1. Modify vlan parameters on attachment 2 (Switches[1]) for VRFs 1-5
					for i := 1; i <= 5; i++ {
						vrfName := helper.GetConfig("vrf").NDFC.VrfPrefix + strconv.Itoa(i)
						if vrf, ok := vrfScaledBulk.Vrfs[vrfName]; ok {
							secondSerial := helper.GetConfig("vrf").NDFC.Switches[1]
							if attach, ok := vrf.AttachList[secondSerial]; ok {
								attach.Vlan = new(types.Int64Custom)
								*attach.Vlan = types.Int64Custom(1000 + i) // 1001, 1002, 1003, 1004, 1005
								vrf.AttachList[secondSerial] = attach
								vrfScaledBulk.Vrfs[vrfName] = vrf
							}
						}
					}

					// 2. Delete VRF 6 entirely
					vrfName6 := helper.GetConfig("vrf").NDFC.VrfPrefix + "6"
					delete(vrfScaledBulk.Vrfs, vrfName6)

					// 3. From VRF 7, remove 2nd and 3rd attachments (Switches[1] and Switches[2])
					vrfName7 := helper.GetConfig("vrf").NDFC.VrfPrefix + "7"
					if vrf7, ok := vrfScaledBulk.Vrfs[vrfName7]; ok {
						delete(vrf7.AttachList, helper.GetConfig("vrf").NDFC.Switches[1])
						delete(vrf7.AttachList, helper.GetConfig("vrf").NDFC.Switches[2])
						vrfScaledBulk.Vrfs[vrfName7] = vrf7
					}

					// 4. From VRF 8, set VRF-level parameter (modify mtu and description)
					helper.ModifyVrfBulkObject(&vrfScaledBulk, 8, map[string]interface{}{
						"mtu":             9100,
						"vrf_description": "Modified VRF 8 in Step 4",
						"max_bgp_paths":   4,
					})

					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},
		}},
	)
}

// func TestAccVRFResourceSimultaneousModifications_OLD(t *testing.T) {
// 	x := &map[string]string{
// 		"RscType":  ndfc.ResourceVrfBulk,
// 		"RscName":  "vrf_test",
// 		"User":     helper.GetConfig("vrf").NDFC.User,
// 		"Password": helper.GetConfig("vrf").NDFC.Password,
// 		"Host":     helper.GetConfig("vrf").NDFC.URL,
// 		"Insecure": helper.GetConfig("vrf").NDFC.Insecure,
// 	}

// 	tfConfig := new(string)
// 	stepCount := new(int)
// 	*stepCount = 0
// 	vrfScaledBulk := new(resource_vrf_bulk.NDFCVrfBulkModel)

// 	resource.Test(t, resource.TestCase{
// 		PreCheck:                 func() { testAccPreCheck(t, "vrf") },
// 		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Step 1: Create 3 VRFs with 2 attachments each, no deployment
// 			{
// 				Config: func() string {
// 					*stepCount++
// 					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
// 					helper.GenerateVrfBulkObject(&vrfScaledBulk, helper.GetConfig("vrf").NDFC.Fabric,
// 						3, false, false, false, []string{helper.GetConfig("vrf").NDFC.Switches[0], helper.GetConfig("vrf").NDFC.Switches[1]})
// 					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
// 					return *tfConfig
// 				}(),
// 				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
// 			},

// 			// Step 2: Simultaneously modify VRF params, add 3rd attachment, and enable deployment
// 			{
// 				Config: func() string {
// 					*stepCount++
// 					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
// 					// Modify VRF 1: Update properties + add attachment + enable deploy
// 					helper.ModifyVrfBulkObject(&vrfScaledBulk, 1, map[string]interface{}{
// 						"vlan_id":              200,
// 						"vrf_description":      "Modified VRF 1",
// 						"mtu":                  9000,
// 						"loopback_routing_tag": 2000,
// 					})
// 					helper.VrfAttachmentsMod(&vrfScaledBulk, 1, 1, helper.GetConfig("vrf").NDFC.Switches, "", nil)
// 					vrfName1 := helper.GetConfig("vrf").NDFC.VrfPrefix + "1"
// 					if vrf, ok := vrfScaledBulk.Vrfs[vrfName1]; ok {
// 						for serial, attach := range vrf.AttachList {
// 							attach.DeployThisAttachment = true
// 							vrf.AttachList[serial] = attach
// 						}
// 						vrfScaledBulk.Vrfs[vrfName1] = vrf
// 					}

// 					// Modify VRF 2: Update properties + remove one attachment
// 					helper.ModifyVrfBulkObject(&vrfScaledBulk, 2, map[string]interface{}{
// 						"vlan_id":         300,
// 						"vrf_description": "Modified VRF 2",
// 						"max_bgp_paths":   3,
// 					})
// 					helper.VrfAttachmentsMod(&vrfScaledBulk, 2, 2, []string{helper.GetConfig("vrf").NDFC.Switches[0]}, "", nil)

// 					// VRF 3: Only add 3rd attachment, no property changes
// 					helper.VrfAttachmentsMod(&vrfScaledBulk, 3, 3, helper.GetConfig("vrf").NDFC.Switches, "", nil)

// 					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
// 					return *tfConfig
// 				}(),
// 				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
// 			},

// 			// Step 3: Simultaneously modify attachments, update different VRF properties, enable deployment for VRF 2
// 			{
// 				Config: func() string {
// 					*stepCount++
// 					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

// 					// VRF 1: Modify attachment params + update VRF properties (keep deploy=true from Step 2)
// 					helper.ModifyVrfBulkObject(&vrfScaledBulk, 1, map[string]interface{}{
// 						"mtu":             9100,
// 						"max_ibgp_paths":  4,
// 						"ipv6_link_local": "false",
// 					})
// 					helper.VrfAttachmentsMod(&vrfScaledBulk, 1, 1, helper.GetConfig("vrf").NDFC.Switches, helper.GetConfig("vrf").NDFC.Switches[2], map[string]interface{}{
// 						"vlan":          3100,
// 						"loopback_id":   100,
// 						"loopback_ipv4": "192.168.1.1",
// 					})
// 					// VRF 1 keeps DeployThisAttachment = true from Step 2

// 					// VRF 2: Swap attachments (remove existing, add new) + modify properties + enable deploy (false->true)
// 					helper.ModifyVrfBulkObject(&vrfScaledBulk, 2, map[string]interface{}{
// 						"vlan_id":              400,
// 						"loopback_routing_tag": 3000,
// 					})
// 					helper.VrfAttachmentsMod(&vrfScaledBulk, 2, 2, []string{helper.GetConfig("vrf").NDFC.Switches[1], helper.GetConfig("vrf").NDFC.Switches[2]}, "", nil)
// 					// Enable deployment for VRF 2
// 					vrfName2 := helper.GetConfig("vrf").NDFC.VrfPrefix + "2"
// 					if vrf, ok := vrfScaledBulk.Vrfs[vrfName2]; ok {
// 						for serial, attach := range vrf.AttachList {
// 							attach.DeployThisAttachment = true
// 							vrf.AttachList[serial] = attach
// 						}
// 						vrfScaledBulk.Vrfs[vrfName2] = vrf
// 					}

// 					// VRF 3: Only modify properties, keep attachments same
// 					helper.ModifyVrfBulkObject(&vrfScaledBulk, 3, map[string]interface{}{
// 						"vrf_description":      "Updated VRF 3",
// 						"loopback_routing_tag": 4000,
// 					})

// 					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
// 					return *tfConfig
// 				}(),
// 				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
// 			},

// 			// Step 4: Remove all attachments from VRF 2, modify VRF 1 & 3 properties
// 			{
// 				Config: func() string {
// 					*stepCount++
// 					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

// 					// VRF 1: Just property update
// 					helper.ModifyVrfBulkObject(&vrfScaledBulk, 1, map[string]interface{}{
// 						"vrf_description": "Final VRF 1 state",
// 					})

// 					// VRF 2: Remove all attachments + modify properties
// 					helper.ModifyVrfBulkObject(&vrfScaledBulk, 2, map[string]interface{}{
// 						"vrf_description": "Detached VRF 2",
// 						"mtu":             9216,
// 					})
// 					helper.VrfAttachmentsMod(&vrfScaledBulk, 2, 2, nil, "", nil)

// 					// VRF 3: Add back attachment that was removed + modify properties
// 					helper.ModifyVrfBulkObject(&vrfScaledBulk, 3, map[string]interface{}{
// 						"max_bgp_paths": 4,
// 					})

// 					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
// 					return *tfConfig
// 				}(),
// 				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
// 			},
// 		}},
// 	)
// }

// TestAccVRFResourceEmptyStateTransition tests the ability to transition
// to and from an empty VRF state (removing all VRFs and adding them back)
func TestAccVRFResourceEmptyStateTransition(t *testing.T) {
	x := &map[string]string{
		"RscType":  ndfc.ResourceVrfBulk,
		"RscName":  "vrf_test",
		"User":     helper.GetConfig("vrf").NDFC.User,
		"Password": helper.GetConfig("vrf").NDFC.Password,
		"Host":     helper.GetConfig("vrf").NDFC.URL,
		"Insecure": helper.GetConfig("vrf").NDFC.Insecure,
	}

	tfConfig := new(string)
	stepCount := new(int)
	*stepCount = 0
	vrfScaledBulk := new(resource_vrf_bulk.NDFCVrfBulkModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "vrf") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create 10 VRFs with 2 attachments each, all deployment flags false
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateVrfBulkObject(&vrfScaledBulk, helper.GetConfig("vrf").NDFC.Fabric,
						10, false, false, false, []string{helper.GetConfig("vrf").NDFC.Switches[0], helper.GetConfig("vrf").NDFC.Switches[1]})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 2: Empty State Transition - Remove ALL VRFs
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Clear all VRFs - empty resource
					vrfScaledBulk.Vrfs = make(map[string]resource_vrf_bulk.NDFCVrfsValue)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(vrfTestRscName, "vrfs.%", "0"),
				),
				// Currently code returns null VRF map, whereas TF expects an empty map.
				// This scenario is unlikely in field; if all vrfs are removed, might as well destroy the resource
				// So marking an expected failure here. Next step should carry on
				// ExpectError: regexp.MustCompile("Error: Provider produced inconsistent result after apply*"),
			},

			// Step 3: Add different set of VRFs back (test recovery from empty state)
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Generate 3 new VRFs with different names and 2 attachments, all deploy=false
					helper.GenerateVrfBulkObject(&vrfScaledBulk, helper.GetConfig("vrf").NDFC.Fabric,
						3, false, false, false, []string{helper.GetConfig("vrf").NDFC.Switches[0], helper.GetConfig("vrf").NDFC.Switches[1]})
					// Modify VRF properties to distinguish from Step 1
					for i := 1; i <= 3; i++ {
						helper.ModifyVrfBulkObject(&vrfScaledBulk, i, map[string]interface{}{
							"vrf_description": fmt.Sprintf("New VRF %d after empty state", i),
							"mtu":             9216,
						})
					}
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},
		}},
	)
}

// TestAccVRFResourceMixedReplaceAndUpdate tests mixed replace + update operations
// Modifies vrf_id (requires_replace=true) on some VRFs while updating others in-place
func TestAccVRFResourceMixedReplaceAndUpdate(t *testing.T) {
	x := &map[string]string{
		"RscType":  ndfc.ResourceVrfBulk,
		"RscName":  "vrf_test",
		"User":     helper.GetConfig("vrf").NDFC.User,
		"Password": helper.GetConfig("vrf").NDFC.Password,
		"Host":     helper.GetConfig("vrf").NDFC.URL,
		"Insecure": helper.GetConfig("vrf").NDFC.Insecure,
	}

	tfConfig := new(string)
	stepCount := new(int)
	*stepCount = 0
	vrfScaledBulk := new(resource_vrf_bulk.NDFCVrfBulkModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "vrf") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create 5 VRFs with specific vrf_ids and 2 attachments each
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateVrfBulkObject(&vrfScaledBulk, helper.GetConfig("vrf").NDFC.Fabric,
						5, false, false, false, []string{helper.GetConfig("vrf").NDFC.Switches[0], helper.GetConfig("vrf").NDFC.Switches[1]})
					// Set specific vrf_ids for each VRF
					for i := 1; i <= 5; i++ {
						helper.ModifyVrfBulkObject(&vrfScaledBulk, i, map[string]interface{}{
							"vrf_id":          51000 + i, // VRF1: 51001, VRF2: 51002, etc.
							"vrf_description": fmt.Sprintf("VRF %d initial", i),
						})
					}
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 2: Mixed operations - Replace VRF1 & VRF2 (vrf_id change), Update VRF3 & VRF4 (in-place), Keep VRF5 unchanged
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					// VRF 1: Change vrf_id (triggers replace)
					helper.ModifyVrfBulkObject(&vrfScaledBulk, 1, map[string]interface{}{
						"vrf_id":          52001, // Changed from 51001
						"vrf_description": "VRF 1 replaced",
					})

					// VRF 2: Change vrf_id (triggers replace)
					helper.ModifyVrfBulkObject(&vrfScaledBulk, 2, map[string]interface{}{
						"vrf_id":          52002, // Changed from 51002
						"vrf_description": "VRF 2 replaced",
					})

					// VRF 3: Update in-place (no vrf_id change)
					helper.ModifyVrfBulkObject(&vrfScaledBulk, 3, map[string]interface{}{
						"vrf_id":          51003, // Same as before
						"vrf_description": "VRF 3 updated in-place",
						"mtu":             9100,
						"max_bgp_paths":   2,
					})

					// VRF 4: Update in-place (no vrf_id change)
					helper.ModifyVrfBulkObject(&vrfScaledBulk, 4, map[string]interface{}{
						"vrf_id":               51004, // Same as before
						"vrf_description":      "VRF 4 updated in-place",
						"loopback_routing_tag": 2500,
					})

					// VRF 5: Keep unchanged
					helper.ModifyVrfBulkObject(&vrfScaledBulk, 5, map[string]interface{}{
						"vrf_id":          51005,           // Same as before
						"vrf_description": "VRF 5 initial", // Same as before
					})

					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 3: Another mixed operation - Replace VRF3 (vrf_id change), add attachment to VRF5
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					// VRF 3: Change vrf_id again (triggers replace)
					helper.ModifyVrfBulkObject(&vrfScaledBulk, 3, map[string]interface{}{
						"vrf_id":          53003, // Changed from 51003
						"vrf_description": "VRF 3 replaced again",
						"mtu":             9216,
					})

					// VRF 5: Add 3rd attachment (in-place update)
					helper.VrfAttachmentsMod(&vrfScaledBulk, 5, 5, helper.GetConfig("vrf").NDFC.Switches, "", nil)

					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},
		}},
	)
}

// ==================== IP-KEY BASED TESTS ====================
// These tests use IP addresses instead of serial numbers as attachment keys

// TestAccVRFResourceIPKeyAttachmentCRUD tests attachment CRUD operations using IP addresses as keys
func TestAccVRFResourceIPKeyAttachmentCRUD(t *testing.T) {

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
			// Step 1: Create VRFs with 2 attachments using IP keys
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					tfConfig := new(string)
					helper.GenerateVrfBulkObject(&vrfScaledBulk, helper.GetConfig("vrf").NDFC.Fabric,
						5, false, false, false, []string{helper.GetConfig("vrf").NDFC.SwitchIP[0], helper.GetConfig("vrf").NDFC.SwitchIP[1]})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},
			// Step 2: Remove both attachments (detach)
			{
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
			// Step 3: Add 2 attachments back using IP keys
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					tfConfig := new(string)
					helper.VrfAttachmentsMod(&vrfScaledBulk, 1, len(vrfScaledBulk.Vrfs), []string{helper.GetConfig("vrf").NDFC.SwitchIP[0], helper.GetConfig("vrf").NDFC.SwitchIP[1]}, "", nil)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},
			// Step 4: Add 3rd attachment using IP key
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					tfConfig := new(string)
					helper.VrfAttachmentsMod(&vrfScaledBulk, 1, len(vrfScaledBulk.Vrfs), helper.GetConfig("vrf").NDFC.SwitchIP, "", nil)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},
			// Step 5: Modify params on specific attachment using IP key
			{
				Config: func() string {
					tfConfig := new(string)
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.VrfAttachmentsMod(&vrfScaledBulk, 1, 1, helper.GetConfig("vrf").NDFC.SwitchIP, helper.GetConfig("vrf").NDFC.SwitchIP[2], map[string]interface{}{
						"vlan":          3001,
						"loopback_id":   1001,
						"loopback_ipv4": "10.1.1.1",
						"loopback_ipv6": "2001:db8::68",
					})

					helper.VrfAttachmentsMod(&vrfScaledBulk, 3, 3, helper.GetConfig("vrf").NDFC.SwitchIP, helper.GetConfig("vrf").NDFC.SwitchIP[2], map[string]interface{}{
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

// TestAccVRFResourceIPKeyGlobalDeploy tests global deploy with IP-based attachment keys
func TestAccVRFResourceIPKeyGlobalDeploy(t *testing.T) {

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
						5, true, false, false, []string{helper.GetConfig("vrf").NDFC.SwitchIP[0], helper.GetConfig("vrf").NDFC.SwitchIP[1]})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},
			{
				// Add 3rd Attachment using IP key
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					tfConfig := new(string)
					helper.VrfAttachmentsMod(&vrfScaledBulk, 1, len(vrfScaledBulk.Vrfs), helper.GetConfig("vrf").NDFC.SwitchIP, "", nil)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},
		},
	})
}

// TestAccVRFResourceIPKeyVrfAttachLevelDeploy tests attachment-level deploy with IP-based keys
func TestAccVRFResourceIPKeyVrfAttachLevelDeploy(t *testing.T) {

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
						5, false, false, true, []string{helper.GetConfig("vrf").NDFC.SwitchIP[0], helper.GetConfig("vrf").NDFC.SwitchIP[1]})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},
		},
	})
}

// TestAccVRFResourceIPKeyAddRemoveVrfsWithDeploy tests add/remove VRFs with deploy using IP-based keys
func TestAccVRFResourceIPKeyAddRemoveVrfsWithDeploy(t *testing.T) {
	x := &map[string]string{
		"RscType":  ndfc.ResourceVrfBulk,
		"RscName":  "vrf_test",
		"User":     helper.GetConfig("vrf").NDFC.User,
		"Password": helper.GetConfig("vrf").NDFC.Password,
		"Host":     helper.GetConfig("vrf").NDFC.URL,
		"Insecure": helper.GetConfig("vrf").NDFC.Insecure,
	}

	tfConfig := new(string)
	stepCount := new(int)
	*stepCount = 0
	vrfScaledBulk := new(resource_vrf_bulk.NDFCVrfBulkModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "vrf") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create 3 VRFs with 2 attachments each using IP keys (with deploy)
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateVrfBulkObject(&vrfScaledBulk, helper.GetConfig("vrf").NDFC.Fabric,
						3, false, false, true, []string{helper.GetConfig("vrf").NDFC.SwitchIP[0], helper.GetConfig("vrf").NDFC.SwitchIP[1]})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 2: Remove the 3rd VRF entirely
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.DeleteVrfs(&vrfScaledBulk, 3, 3)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 3: Add the 3rd VRF back using IP keys
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateVrfBulkObject(&vrfScaledBulk, helper.GetConfig("vrf").NDFC.Fabric,
						3, false, false, true, []string{helper.GetConfig("vrf").NDFC.SwitchIP[0], helper.GetConfig("vrf").NDFC.SwitchIP[1]})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 4: Remove the 1st VRF
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.DeleteVrfs(&vrfScaledBulk, 1, 1)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 5: Add both 1st and 2nd VRFs back using IP keys
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateVrfBulkObject(&vrfScaledBulk, helper.GetConfig("vrf").NDFC.Fabric,
						3, false, false, true, []string{helper.GetConfig("vrf").NDFC.SwitchIP[0], helper.GetConfig("vrf").NDFC.SwitchIP[1]})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},
		}},
	)
}

// TestAccVRFResourceIPKeySimultaneousModifications tests complex updates with IP-based keys
func TestAccVRFResourceIPKeySimultaneousModifications(t *testing.T) {
	x := &map[string]string{
		"RscType":  ndfc.ResourceVrfBulk,
		"RscName":  "vrf_test",
		"User":     helper.GetConfig("vrf").NDFC.User,
		"Password": helper.GetConfig("vrf").NDFC.Password,
		"Host":     helper.GetConfig("vrf").NDFC.URL,
		"Insecure": helper.GetConfig("vrf").NDFC.Insecure,
	}

	tfConfig := new(string)
	stepCount := new(int)
	*stepCount = 0
	vrfScaledBulk := new(resource_vrf_bulk.NDFCVrfBulkModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "vrf") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create 5 VRFs with 2 attachments each using IP keys
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateVrfBulkObject(&vrfScaledBulk, helper.GetConfig("vrf").NDFC.Fabric,
						5, false, false, false, []string{helper.GetConfig("vrf").NDFC.SwitchIP[0], helper.GetConfig("vrf").NDFC.SwitchIP[1]})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 2: Add 3rd attachment with deploy=false, toggle existing 2 attachments to deploy=true
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Add 3rd attachment and enable deployment on existing attachments for all VRFs
					for i := 1; i <= 5; i++ {
						vrfName := helper.GetConfig("vrf").NDFC.VrfPrefix + strconv.Itoa(i)
						if vrf, ok := vrfScaledBulk.Vrfs[vrfName]; ok {
							// Toggle existing 2 attachments to deploy=true
							for ipKey, attach := range vrf.AttachList {
								attach.DeployThisAttachment = true
								vrf.AttachList[ipKey] = attach
							}
							// Add 3rd attachment with deploy=false using IP key
							thirdAttach := resource_vrf_attachments.NDFCAttachListValue{}
							thirdAttach.SerialNumber = helper.GetConfig("vrf").NDFC.SwitchIP[2]
							thirdAttach.DeployThisAttachment = false
							vrf.AttachList[helper.GetConfig("vrf").NDFC.SwitchIP[2]] = thirdAttach
							vrfScaledBulk.Vrfs[vrfName] = vrf
						}
					}
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 3: Toggle 3rd attachment to deploy=true
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Toggle 3rd attachment to deploy=true for all VRFs
					for i := 1; i <= 5; i++ {
						vrfName := helper.GetConfig("vrf").NDFC.VrfPrefix + strconv.Itoa(i)
						if vrf, ok := vrfScaledBulk.Vrfs[vrfName]; ok {
							thirdIP := helper.GetConfig("vrf").NDFC.SwitchIP[2]
							if attach, ok := vrf.AttachList[thirdIP]; ok {
								attach.DeployThisAttachment = true
								vrf.AttachList[thirdIP] = attach
								vrfScaledBulk.Vrfs[vrfName] = vrf
							}
						}
					}
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 4: Complex simultaneous modifications across multiple VRFs
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					// 1. Modify vlan parameters on attachment 2 (SwitchIP[1]) for VRFs 1-3
					for i := 1; i <= 3; i++ {
						vrfName := helper.GetConfig("vrf").NDFC.VrfPrefix + strconv.Itoa(i)
						if vrf, ok := vrfScaledBulk.Vrfs[vrfName]; ok {
							secondIP := helper.GetConfig("vrf").NDFC.SwitchIP[1]
							if attach, ok := vrf.AttachList[secondIP]; ok {
								attach.Vlan = new(types.Int64Custom)
								*attach.Vlan = types.Int64Custom(1000 + i) // 1001, 1002, 1003
								vrf.AttachList[secondIP] = attach
								vrfScaledBulk.Vrfs[vrfName] = vrf
							}
						}
					}

					// 2. Delete VRF 4 entirely
					vrfName4 := helper.GetConfig("vrf").NDFC.VrfPrefix + "4"
					delete(vrfScaledBulk.Vrfs, vrfName4)

					// 3. From VRF 5, remove 2nd and 3rd attachments
					vrfName5 := helper.GetConfig("vrf").NDFC.VrfPrefix + "5"
					if vrf5, ok := vrfScaledBulk.Vrfs[vrfName5]; ok {
						delete(vrf5.AttachList, helper.GetConfig("vrf").NDFC.SwitchIP[1])
						delete(vrf5.AttachList, helper.GetConfig("vrf").NDFC.SwitchIP[2])
						vrfScaledBulk.Vrfs[vrfName5] = vrf5
					}

					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},
		}},
	)
}

// TestAccVRFResourceIPKeyVrfLevelDeploy tests VRF-level deploy (deploy_attachments) with IP-based keys
func TestAccVRFResourceIPKeyVrfLevelDeploy(t *testing.T) {

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
			// Step 1: Create 5 VRFs with deploy_attachments=true, 2 IP-keyed attachments each
			{
				PreConfig: func() {
					t.Logf(logTestStep, t.Name(), *stepCount)
				},
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					tfConfig := new(string)
					helper.GenerateVrfBulkObject(&vrfScaledBulk, helper.GetConfig("vrf").NDFC.Fabric,
						5, false, true, false, []string{helper.GetConfig("vrf").NDFC.SwitchIP[0], helper.GetConfig("vrf").NDFC.SwitchIP[1]})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},
			// Step 2: Add 3rd attachment using IP key (auto-deployed via VRF-level flag)
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					tfConfig := new(string)
					helper.VrfAttachmentsMod(&vrfScaledBulk, 1, len(vrfScaledBulk.Vrfs), helper.GetConfig("vrf").NDFC.SwitchIP, "", nil)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},
		},
	})
}

// TestAccVRFResourceIPKeyMixedKeys tests attachments with mixed keys (serial + IP) in same VRF
func TestAccVRFResourceIPKeyMixedKeys(t *testing.T) {

	x := &map[string]string{
		"RscType":  ndfc.ResourceVrfBulk,
		"RscName":  "vrf_test",
		"User":     helper.GetConfig("vrf").NDFC.User,
		"Password": helper.GetConfig("vrf").NDFC.Password,
		"Host":     helper.GetConfig("vrf").NDFC.URL,
		"Insecure": helper.GetConfig("vrf").NDFC.Insecure,
	}

	tfConfig := new(string)
	stepCount := new(int)
	*stepCount = 0
	vrfScaledBulk := new(resource_vrf_bulk.NDFCVrfBulkModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "vrf") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create 3 VRFs with mixed keys - 1st attachment uses serial, 2nd uses IP
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					// Create VRFs with mixed attachment keys
					vrfScaledBulk.FabricName = helper.GetConfig("vrf").NDFC.Fabric
					vrfScaledBulk.DeployAllAttachments = false
					vrfScaledBulk.Vrfs = make(map[string]resource_vrf_bulk.NDFCVrfsValue)

					for i := 1; i <= 3; i++ {
						vrfName := helper.GetConfig("vrf").NDFC.VrfPrefix + strconv.Itoa(i)
						vrf := resource_vrf_bulk.NDFCVrfsValue{}
						vrf.DeployAttachments = true
						vrf.AttachList = make(map[string]resource_vrf_attachments.NDFCAttachListValue)

						// First attachment uses serial number as key
						attach1 := resource_vrf_attachments.NDFCAttachListValue{}
						//attach1.SerialNumber = helper.GetConfig("vrf").NDFC.Switches[0]
						vrf.AttachList[helper.GetConfig("vrf").NDFC.Switches[0]] = attach1

						// Second attachment uses IP address as key
						attach2 := resource_vrf_attachments.NDFCAttachListValue{}
						//attach2.SerialNumber = helper.GetConfig("vrf").NDFC.SwitchIP[1]
						vrf.AttachList[helper.GetConfig("vrf").NDFC.SwitchIP[1]] = attach2

						vrfScaledBulk.Vrfs[vrfName] = vrf
					}

					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},
			// Step 2: Add 3rd attachment using IP key
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					for vrfName, vrf := range vrfScaledBulk.Vrfs {
						attach3 := resource_vrf_attachments.NDFCAttachListValue{}
						//attach3.SerialNumber = helper.GetConfig("vrf").NDFC.SwitchIP[2]
						vrf.AttachList[helper.GetConfig("vrf").NDFC.SwitchIP[2]] = attach3
						vrfScaledBulk.Vrfs[vrfName] = vrf
					}

					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},
			// Step 3: Remove the serial-based attachment, keep IP-based ones
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					for vrfName, vrf := range vrfScaledBulk.Vrfs {
						delete(vrf.AttachList, helper.GetConfig("vrf").NDFC.Switches[0])
						vrfScaledBulk.Vrfs[vrfName] = vrf
					}

					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},
		},
	})
}

// TestAccVRFResourceFreeformConfig tests freeform config on VRF attachments
// Tests setting, updating, and removing freeform config
func TestAccVRFResourceFreeformConfig(t *testing.T) {
	x := &map[string]string{
		"RscType":  ndfc.ResourceVrfBulk,
		"RscName":  "vrf_test",
		"User":     helper.GetConfig("vrf").NDFC.User,
		"Password": helper.GetConfig("vrf").NDFC.Password,
		"Host":     helper.GetConfig("vrf").NDFC.URL,
		"Insecure": helper.GetConfig("vrf").NDFC.Insecure,
	}

	tfConfig := new(string)
	stepCount := new(int)
	*stepCount = 0
	vrfScaledBulk := new(resource_vrf_bulk.NDFCVrfBulkModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "vrf") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create 3 VRFs with 2 attachments each (no freeform config initially)
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateVrfBulkObject(&vrfScaledBulk, helper.GetConfig("vrf").NDFC.Fabric,
						3, false, false, false, []string{helper.GetConfig("vrf").NDFC.Switches[0], helper.GetConfig("vrf").NDFC.Switches[1]})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 2: Add freeform_config to first attachment of VRF 1
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.VrfAttachmentsMod(&vrfScaledBulk, 1, 1, nil, helper.GetConfig("vrf").NDFC.Switches[0], map[string]interface{}{
						"freeform_config": "router bgp 29500\\n  vrf {vrfName}\\n    address-family ipv4 unicast\\n      maximum-paths 2\\ninterface Ethernet1/21.1010\\n    encapsulation dot1q 1010\\n    vrf member {vrfName}\\n    mtu 9216\\n    ip address 10.1.0.1/30\\n    no shutdown",
					})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				//Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 3: Add freeform_config to multiple attachments across different VRFs
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// VRF 1, Switch 1: Already has config from previous step
					// VRF 1, Switch 2: Add new freeform config
					helper.VrfAttachmentsMod(&vrfScaledBulk, 1, 1, nil, helper.GetConfig("vrf").NDFC.Switches[1], map[string]interface{}{
						"freeform_config": "router bgp 29500\\n  vrf {vrfName}\\n    address-family ipv4 unicast\\n      maximum-paths 2\\ninterface Ethernet1/22.1011\\n    encapsulation dot1q 1011\\n    vrf member {vrfName}\\n    mtu 9216\\n    ip address 10.1.1.1/30\\n    no shutdown",
					})
					// VRF 2, Switch 1: Add freeform config
					helper.VrfAttachmentsMod(&vrfScaledBulk, 2, 2, nil, helper.GetConfig("vrf").NDFC.Switches[0], map[string]interface{}{
						"freeform_config": "router bgp 29500\\n  vrf {vrfName}\\n    address-family ipv4 unicast\\n      maximum-paths 3\\ninterface Ethernet1/21.1020\\n    encapsulation dot1q 1020\\n    vrf member {vrfName}\\n    mtu 9216\\n    ip address 10.2.0.1/30\\n    no shutdown",
					})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				//Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 4: Update existing freeform_config (modify content)
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.VrfAttachmentsMod(&vrfScaledBulk, 1, 1, nil, helper.GetConfig("vrf").NDFC.Switches[0], map[string]interface{}{
						"freeform_config": "router bgp 29500\\n  vrf {vrfName}\\n    address-family ipv4 unicast\\n      maximum-paths 4\\ninterface Ethernet1/21.1012\\n    encapsulation dot1q 1012\\n    vrf member {vrfName}\\n    mtu 9000\\n    ip address 10.1.0.5/30\\n    no shutdown",
					})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				//Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 5: Remove freeform_config from one attachment (set to empty string)
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.VrfAttachmentsMod(&vrfScaledBulk, 2, 2, nil, helper.GetConfig("vrf").NDFC.Switches[0], map[string]interface{}{
						"freeform_config": "",
					})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				//Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},

			// Step 6: Add freeform_config along with other attachment parameters
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.VrfAttachmentsMod(&vrfScaledBulk, 3, 3, nil, helper.GetConfig("vrf").NDFC.Switches[0], map[string]interface{}{
						"freeform_config": "router bgp 29500\\n  vrf {vrfName}\\n    address-family ipv4 unicast\\n      maximum-paths 2\\ninterface Ethernet1/21.1030\\n    encapsulation dot1q 1030\\n    vrf member {vrfName}\\n    mtu 9216\\n    ip address 10.3.0.1/30\\n    no shutdown",
						"vlan":            300,
					})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfScaledBulk}, &tfConfig)
					return *tfConfig
				}(),
				//Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck(vrfTestRscName, *vrfScaledBulk, path.Empty())...),
			},
		}},
	)
}
