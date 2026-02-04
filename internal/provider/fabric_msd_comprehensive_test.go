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
	"strings"
	"terraform-provider-ndfc/internal/provider/ndfc"
	helper "terraform-provider-ndfc/internal/provider/testing"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// TestAccMsdFabricComprehensive tests comprehensive MSD fabric scenarios
// Covers VRF and Network CRUD operations, attachments, deployments, and parameter modifications
func TestAccMsdFabricComprehensive(t *testing.T) {
	cfg := helper.GetConfig("fabric_vxlan_msd")
	if cfg.NDFC.Msd.MsdFabricName == "" {
		t.Skip("Skipping test as no MSD fabric configuration is available")
	}
	x := &map[string]string{
		"RscType":  ndfc.ResourceVxlanMsdFabric,
		"RscName":  "msd_fabric",
		"User":     cfg.NDFC.User,
		"Password": cfg.NDFC.Password,
		"Host":     cfg.NDFC.URL,
		"Insecure": cfg.NDFC.Insecure,
	}

	testConfig := new(helper.MsdComprehensiveTestConfig)
	tfConfig := new(string)
	stepCount := new(int)
	*stepCount = 0

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "fabric_vxlan_msd") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create 5 VXLAN EVPN fabrics and add inventory devices
			{
				Config: func() string {
					*stepCount++
					t_name := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					testConfig = initializeFabricsAndInventory(cfg, 5)
					helper.GetTFConfigWithSingleResource(t_name, *x, []interface{}{testConfig}, &tfConfig)
					// fmt.Println("Step 1 - Terraform Config:\n", *tfConfig)
					return *tfConfig
				}(),
			},

			// Step 2: Error case - fabric field in VRF attachment on individual (non-MSD) fabric
			// Expects error: "Invalid fabric field in attachment" - fabric field only valid for MSD parent
			{
				Config: func() string {
					*stepCount++
					t_name := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Create VRF on individual fabric with fabric field set - should fail
					createVrfWithInvalidFabricField(testConfig, cfg, 0)
					helper.GetTFConfigWithSingleResource(t_name, *x, []interface{}{testConfig}, &tfConfig)
					// fmt.Println("Step 2 - Terraform Config (expect error):\n", *tfConfig)
					return *tfConfig
				}(),
				ExpectError: regexp.MustCompile(`Invalid fabric field in attachment`),
			},

			// Step 3: Create VRFs on individual fabrics with valid config (no fabric field)
			// These VRFs will be managed when fabrics later join MSD
			{
				Config: func() string {
					*stepCount++
					t_name := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Clear the invalid VRF config from Step 2
					testConfig.MsdChildVrfs = nil

					// Create VRFs 1-3 on first 2 fabrics (matching what we'll use in MSD later)
					createVrfOnIndividualFabric(testConfig, cfg, 0, []int{1, 2, 3})
					createVrfOnIndividualFabric(testConfig, cfg, 1, []int{1, 2, 3})

					helper.GetTFConfigWithSingleResource(t_name, *x, []interface{}{testConfig}, &tfConfig)
					// fmt.Println("Step 3 - Terraform Config:\n", *tfConfig)
					return *tfConfig
				}(),
			},

			// Step 4: Create MSD fabric with 4 child fabrics (fabrics already have VRFs from Step 3)
			{
				Config: func() string {
					*stepCount++
					t_name := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// ChildVrfs already exist from Step 3, just need to add MSD fabric
					// The VRFs will be "adopted" by MSD when fabrics become children

					// Build DependsOn string for all inventory resources
					var dependsOnParts []string
					for _, inv := range testConfig.InventoryDevices {
						dependsOnParts = append(dependsOnParts, fmt.Sprintf("resource.ndfc_inventory_devices.%s", inv.ResourceName))
					}
					dependsOn := ""
					if len(dependsOnParts) > 0 {
						dependsOn = strings.Join(dependsOnParts, ", ")
					}

					testConfig.MsdFabric = &helper.MsdFabricConfig{
						ResourceName: "msd_fabric",
						DependsOn:    dependsOn,
						FabricName:   cfg.NDFC.Msd.MsdFabricName,
						Deploy:       false,
						ChildFabrics: []string{
							cfg.NDFC.Msd.ChildFabrics[0].Name,
							cfg.NDFC.Msd.ChildFabrics[2].Name,
							cfg.NDFC.Msd.ChildFabrics[3].Name,
							cfg.NDFC.Msd.ChildFabrics[4].Name,
						},
					}
					helper.GetTFConfigWithSingleResource(t_name, *x, []interface{}{testConfig}, &tfConfig)
					// fmt.Println("Step 4 - Terraform Config:\n", *tfConfig)
					return *tfConfig
				}(),
			},

			// Step 5: Update to only 2 child fabrics (0 and 1 - both have pre-existing VRFs)
			{
				Config: func() string {
					*stepCount++
					t_name := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					testConfig.MsdFabric.ChildFabrics = []string{
						cfg.NDFC.Msd.ChildFabrics[0].Name,
						cfg.NDFC.Msd.ChildFabrics[1].Name,
					}
					helper.GetTFConfigWithSingleResource(t_name, *x, []interface{}{testConfig}, &tfConfig)
					// fmt.Println("Step 5 - Terraform Config:\n", *tfConfig)
					return *tfConfig
				}(),
			},

			// Step 6: Update pre-existing VRFs into MSD management (VRFs 1-3 already exist from Step 3)
			{
				Config: func() string {
					*stepCount++
					t_name := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					initializeVrfConfig(testConfig, cfg)

					// Create VRFs 1-3 with same config as individual VRFs
					for i := 1; i <= 3; i++ {
						createVrf(testConfig, cfg, i)
					}

					helper.GetTFConfigWithSingleResource(t_name, *x, []interface{}{testConfig}, &tfConfig)
					// fmt.Println("Step 6 - Terraform Config:\n", *tfConfig)
					return *tfConfig
				}(),
			},

			// Step 7: Error case - setting child-specific properties on MSD parent VRF
			// Expects error: "AdvertiseHostRoutes is managed per child fabric and should not be set in MSD parent"
			{
				Config: func() string {
					*stepCount++
					t_name := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Set child-only property on parent VRF - should fail
					setChildPropertiesOnParentVrf(testConfig, 1, cfg)
					helper.GetTFConfigWithSingleResource(t_name, *x, []interface{}{testConfig}, &tfConfig)
					//fmt.Println("Step 7 - Terraform Config (expect error):\n", *tfConfig)
					return *tfConfig
				}(),
				ExpectError: regexp.MustCompile(`AdvertiseHostRoutes is managed per child fabric`),
			},

			// Step 8: Clear child properties and add 7 more VRFs (total 10)
			{
				Config: func() string {
					*stepCount++
					t_name := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Clear the invalid child property from Step 7
					clearChildPropertiesOnParentVrf(testConfig, 1, cfg)

					for i := 4; i <= 10; i++ {
						createVrf(testConfig, cfg, i)
					}

					helper.GetTFConfigWithSingleResource(t_name, *x, []interface{}{testConfig}, &tfConfig)
					// fmt.Println("Step 8 - Terraform Config:\n", *tfConfig)
					return *tfConfig
				}(),
			},

			// Step 9: Error case - setting mismatched parent property on child VRF
			// Expects error: "Mismatch of config between parent and child attributes"
			{
				Config: func() string {
					*stepCount++
					t_name := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Set mismatched MTU on child VRF - different from parent
					setMismatchedParentPropertyOnChildVrf(testConfig, 1, cfg)
					helper.GetTFConfigWithSingleResource(t_name, *x, []interface{}{testConfig}, &tfConfig)
					// fmt.Println("Step 9 - Terraform Config (expect error):\n", *tfConfig)
					return *tfConfig
				}(),
				ExpectError: regexp.MustCompile(`Mismatch of config between parent and child attributes`),
			},

			// Step 10: Clear mismatched property and remove VRFs 6-10 (keep 1-5)
			{
				Config: func() string {
					*stepCount++
					t_name := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Clear the mismatched property from Step 9
					clearMismatchedParentPropertyOnChildVrf(testConfig, 1, cfg)

					deleteVrfs(testConfig, cfg, []int{6, 7, 8, 9, 10})

					helper.GetTFConfigWithSingleResource(t_name, *x, []interface{}{testConfig}, &tfConfig)
					// fmt.Println("Step 8 - Terraform Config:\n", *tfConfig)
					return *tfConfig
				}(),
			},

			// Step 11: Add Networks for the first 3 VRFs
			{
				Config: func() string {
					*stepCount++
					t_name := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					initializeNetworkConfig(testConfig, cfg)

					// Create 3 networks
					for i := 1; i <= 3; i++ {
						createNetwork(testConfig, cfg, i, i)
					}

					helper.GetTFConfigWithSingleResource(t_name, *x, []interface{}{testConfig}, &tfConfig)
					// fmt.Println("Step 9 - Terraform Config:\n", *tfConfig)
					return *tfConfig
				}(),
			},

			// Step 12: Remove all networks to prepare for full recreation
			{
				Config: func() string {
					*stepCount++
					t_name := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					testConfig.MsdParentNetworks = nil
					testConfig.MsdChildNetworks = nil

					helper.GetTFConfigWithSingleResource(t_name, *x, []interface{}{testConfig}, &tfConfig)
					// fmt.Println("Step 10 - Terraform Config:\n", *tfConfig)
					return *tfConfig
				}(),
				ExpectNonEmptyPlan: true,
			},

			// Step 13: Create all 10 networks fresh
			{
				Config: func() string {
					*stepCount++
					t_name := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					initializeNetworkConfig(testConfig, cfg)

					// Create all 10 networks
					for i := 1; i <= 10; i++ {
						vrfIdx := ((i - 1) % 5) + 1 // Cycle through VRFs 1-5
						createNetwork(testConfig, cfg, i, vrfIdx)
					}

					helper.GetTFConfigWithSingleResource(t_name, *x, []interface{}{testConfig}, &tfConfig)
					// fmt.Println("Step 11 - Terraform Config:\n", *tfConfig)
					return *tfConfig
				}(),
			},

			// Step 14: Error case - setting child-specific properties on MSD parent Network
			// Expects error: "TRM Enabled config not supported for MSD Parent Fabric"
			{
				Config: func() string {
					*stepCount++
					t_name := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Set child-only property on parent Network - should fail
					setChildPropertiesOnParentNetwork(testConfig, 1, cfg)
					helper.GetTFConfigWithSingleResource(t_name, *x, []interface{}{testConfig}, &tfConfig)
					// fmt.Println("Step 13 - Terraform Config (expect error):\n", *tfConfig)
					return *tfConfig
				}(),
				ExpectError: regexp.MustCompile(`TRM Enabled config not supported for MSD Parent Fabric`),
			},

			// Step 15: Clear child properties and add mismatch error test
			{
				Config: func() string {
					*stepCount++
					t_name := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Clear the invalid child property from Step 14
					clearChildPropertiesOnParentNetwork(testConfig, 1, cfg)

					helper.GetTFConfigWithSingleResource(t_name, *x, []interface{}{testConfig}, &tfConfig)
					// fmt.Println("Step 15 - Terraform Config:\n", *tfConfig)
					return *tfConfig
				}(),
			},

			// Step 16: Error case - setting mismatched parent property on child Network
			// Expects error: "Mismatch of config between parent and child attributes"
			{
				Config: func() string {
					*stepCount++
					t_name := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Set mismatched MTU on child Network - different from parent
					setMismatchedParentPropertyOnChildNetwork(testConfig, 1, cfg)
					helper.GetTFConfigWithSingleResource(t_name, *x, []interface{}{testConfig}, &tfConfig)
					// fmt.Println("Step 16 - Terraform Config (expect error):\n", *tfConfig)
					return *tfConfig
				}(),
				ExpectError: regexp.MustCompile(`Mismatch of config between parent and child attributes`),
			},

			// Step 17: Clear mismatch and remove networks 6-10 (keep 1-5)
			{
				Config: func() string {
					*stepCount++
					t_name := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Clear the mismatched property from Step 16
					clearMismatchedParentPropertyOnChildNetwork(testConfig, 1, cfg)

					deleteNetworks(testConfig, cfg, []int{6, 7, 8, 9, 10})

					helper.GetTFConfigWithSingleResource(t_name, *x, []interface{}{testConfig}, &tfConfig)
					// fmt.Println("Step 12 - Terraform Config:\n", *tfConfig)
					return *tfConfig
				}(),
			},

			// Step 18: Cleanup - Remove all networks to test network deletion
			{
				Config: func() string {
					*stepCount++
					t_name := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					testConfig.MsdParentNetworks = nil
					testConfig.MsdChildNetworks = nil

					helper.GetTFConfigWithSingleResource(t_name, *x, []interface{}{testConfig}, &tfConfig)
					// fmt.Println("Step 13 - Terraform Config:\n", *tfConfig)
					return *tfConfig
				}(),
				ExpectNonEmptyPlan: true,
			},

			// Step 19: Recreate 3 networks afresh
			{
				Config: func() string {
					*stepCount++
					t_name := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					initializeNetworkConfig(testConfig, cfg)

					// Create 3 networks
					for i := 1; i <= 3; i++ {
						createNetwork(testConfig, cfg, i, i)
					}

					helper.GetTFConfigWithSingleResource(t_name, *x, []interface{}{testConfig}, &tfConfig)
					// fmt.Println("Step 14 - Terraform Config:\n", *tfConfig)
					return *tfConfig
				}(),
			},

			// Step 20: Modify VRF and network parameters
			{
				Config: func() string {
					*stepCount++
					t_name := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					modifyVrfParameters(testConfig, cfg, []int{1, 5}, 9000, 54000, 2, 3)
					modifyNetworkParameters(testConfig, cfg, []int{1, 3}, 400, 9200, 54000, " SVI Modified", false, "239.2.2.", 2)
					setGlobalDeployFlags(testConfig, true, false)

					helper.GetTFConfigWithSingleResource(t_name, *x, []interface{}{testConfig}, &tfConfig)
					// fmt.Println("Step 15 - Terraform Config:\n", *tfConfig)
					return *tfConfig
				}(),
			},

			// Step 21: Remove network attachments from ONE switch for ALL networks belonging to VRF 2
			// This prepares for Step 22 which will remove the VRF attachment from that switch.
			// NDFC requires all networks to be detached from a switch before VRF can be detached.
			{
				Config: func() string {
					*stepCount++
					t_name := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					vrf2Name := fmt.Sprintf("%s%02d", cfg.NDFC.Msd.VrfPrefix, 2)

					// Find all networks that belong to VRF 2
					var networksInVrf2 []string
					for nwName, nw := range testConfig.MsdParentNetworks.Networks {
						if nw.VrfName == vrf2Name {
							networksInVrf2 = append(networksInVrf2, nwName)
						}
					}

					// Pick one switch to remove all network attachments from
					// Use the first network to identify a switch
					var removedSerial, removedFabric string
					if len(networksInVrf2) > 0 {
						firstNw := testConfig.MsdParentNetworks.Networks[networksInVrf2[0]]
						for switchSerial, att := range firstNw.Attachments {
							removedSerial = switchSerial
							removedFabric = att.Fabric
							break
						}
					}

					if removedSerial == "" {
						panic("Step 21: No network attachments found for VRF 2")
					}

					// Remove this switch's attachment from ALL networks in VRF 2 (parent)
					for _, nwName := range networksInVrf2 {
						nw := testConfig.MsdParentNetworks.Networks[nwName]
						delete(nw.Attachments, removedSerial)
						testConfig.MsdParentNetworks.Networks[nwName] = nw
					}

					// Remove the SAME switch's attachment from ALL networks in VRF 2 (child fabric)
					for i := range testConfig.MsdChildNetworks {
						if testConfig.MsdChildNetworks[i].FabricName != removedFabric {
							continue
						}
						for _, nwName := range networksInVrf2 {
							childNw := testConfig.MsdChildNetworks[i].Networks[nwName]
							delete(childNw.Attachments, removedSerial)
							testConfig.MsdChildNetworks[i].Networks[nwName] = childNw
						}
						break
					}

					helper.GetTFConfigWithSingleResource(t_name, *x, []interface{}{testConfig}, &tfConfig)
					// fmt.Println("Step 16 - Terraform Config:\n", *tfConfig)
					return *tfConfig
				}(),
				ExpectNonEmptyPlan: true,
			},

			// Step 22: Remove one attachment from VRF 2 (keep at least one)
			// Parent resource runs first, so we remove from parent and corresponding child
			// IMPORTANT: We can only remove VRF attachments from switches that have NO network attachments
			// for that VRF. NDFC requires networks to be detached before VRF can be detached.
			{
				Config: func() string {
					*stepCount++
					t_name := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					vrf2Name := fmt.Sprintf("%s%02d", cfg.NDFC.Msd.VrfPrefix, 2)

					// Build a set of switches that still have network attachments for networks in this VRF
					switchesWithNetworkAttachments := make(map[string]bool)
					for _, nw := range testConfig.MsdParentNetworks.Networks {
						if nw.VrfName == vrf2Name {
							for switchSerial := range nw.Attachments {
								switchesWithNetworkAttachments[switchSerial] = true
							}
						}
					}

					// Pick a VRF attachment from a switch that has no network attachments
					vrf := testConfig.MsdParentVrfs.Vrfs[vrf2Name]
					var removedSerial, removedFabric string
					for switchSerial, att := range vrf.AttachList {
						if switchesWithNetworkAttachments[switchSerial] {
							continue // Skip - this switch still has networks attached
						}
						removedSerial = switchSerial
						removedFabric = att.Fabric
						delete(vrf.AttachList, switchSerial)
						break // Remove only one
					}

					// If we couldn't find a switch without network attachments, the test has a problem
					if removedSerial == "" {
						// All VRF-attached switches still have network attachments
						// This means Step 15 didn't remove a network attachment from a switch
						// that allows us to remove the VRF attachment
						panic("Step 22: Cannot find a VRF attachment to remove - all switches still have network attachments")
					}

					testConfig.MsdParentVrfs.Vrfs[vrf2Name] = vrf

					// Remove the SAME attachment from the corresponding child fabric
					for i := range testConfig.MsdChildVrfs {
						if testConfig.MsdChildVrfs[i].FabricName != removedFabric {
							continue
						}
						childVrf := testConfig.MsdChildVrfs[i].Vrfs[vrf2Name]
						delete(childVrf.AttachList, removedSerial)
						testConfig.MsdChildVrfs[i].Vrfs[vrf2Name] = childVrf
						break
					}

					helper.GetTFConfigWithSingleResource(t_name, *x, []interface{}{testConfig}, &tfConfig)
					// fmt.Println("Step 17 - Terraform Config:\n", *tfConfig)
					return *tfConfig
				}(),
			},

			// Step 23: Re-add attachments to VRF 2 and all Networks in VRF 2
			// Since Step 15 removed network attachments from ALL networks in VRF 2,
			// we need to re-add attachments to all of them.
			{
				Config: func() string {
					*stepCount++
					t_name := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					vrf2Name := fmt.Sprintf("%s%02d", cfg.NDFC.Msd.VrfPrefix, 2)

					// Build map of active child fabrics
					activeFabricMap := make(map[string]bool)
					for _, childVrfConfig := range testConfig.MsdChildVrfs {
						activeFabricMap[childVrfConfig.FabricName] = true
					}

					// Re-add parent VRF attachments dynamically for active child fabrics only
					vrf := testConfig.MsdParentVrfs.Vrfs[vrf2Name]
					vrf.AttachList = make(map[string]helper.MsdAttachment)
					for _, childFabric := range cfg.NDFC.Msd.ChildFabrics {
						if !activeFabricMap[childFabric.Name] {
							continue
						}
						for _, switchSerial := range childFabric.Switches {
							vrf.AttachList[switchSerial] = helper.MsdAttachment{
								DeployThisAttachment: false,
								Fabric:               childFabric.Name,
							}
						}
					}
					testConfig.MsdParentVrfs.Vrfs[vrf2Name] = vrf

					// Re-add child VRF attachments dynamically for each active child fabric
					for i := range testConfig.MsdChildVrfs {
						childVrfConfig := &testConfig.MsdChildVrfs[i]
						// Find the corresponding child fabric config
						var childFabric *helper.MsdChildFabric
						for j := range cfg.NDFC.Msd.ChildFabrics {
							if cfg.NDFC.Msd.ChildFabrics[j].Name == childVrfConfig.FabricName {
								childFabric = &cfg.NDFC.Msd.ChildFabrics[j]
								break
							}
						}
						if childFabric == nil {
							continue
						}

						childVrf := childVrfConfig.Vrfs[vrf2Name]
						childVrf.AttachList = make(map[string]helper.MsdAttachment)
						for _, switchSerial := range childFabric.Switches {
							childVrf.AttachList[switchSerial] = helper.MsdAttachment{
								DeployThisAttachment: false,
							}
						}
						childVrfConfig.Vrfs[vrf2Name] = childVrf
					}

					// Find all networks belonging to VRF 2
					var networksInVrf2 []string
					for nwName, nw := range testConfig.MsdParentNetworks.Networks {
						if nw.VrfName == vrf2Name {
							networksInVrf2 = append(networksInVrf2, nwName)
						}
					}

					// Re-add parent network attachments for ALL networks in VRF 2
					for _, nwName := range networksInVrf2 {
						nw := testConfig.MsdParentNetworks.Networks[nwName]
						nw.Attachments = make(map[string]helper.MsdAttachment)
						for _, childFabric := range cfg.NDFC.Msd.ChildFabrics {
							if !activeFabricMap[childFabric.Name] {
								continue
							}
							for _, switchSerial := range childFabric.Switches {
								nw.Attachments[switchSerial] = helper.MsdAttachment{
									DeployThisAttachment: false,
									Fabric:               childFabric.Name,
								}
							}
						}
						testConfig.MsdParentNetworks.Networks[nwName] = nw
					}

					// Re-add child network attachments for ALL networks in VRF 2
					for i := range testConfig.MsdChildNetworks {
						childNetworkConfig := &testConfig.MsdChildNetworks[i]
						// Find the corresponding child fabric config
						var childFabric *helper.MsdChildFabric
						for j := range cfg.NDFC.Msd.ChildFabrics {
							if cfg.NDFC.Msd.ChildFabrics[j].Name == childNetworkConfig.FabricName {
								childFabric = &cfg.NDFC.Msd.ChildFabrics[j]
								break
							}
						}
						if childFabric == nil {
							continue
						}

						for _, nwName := range networksInVrf2 {
							childNw := childNetworkConfig.Networks[nwName]
							childNw.Attachments = make(map[string]helper.MsdAttachment)
							for _, switchSerial := range childFabric.Switches {
								childNw.Attachments[switchSerial] = helper.MsdAttachment{
									DeployThisAttachment: false,
								}
							}
							childNetworkConfig.Networks[nwName] = childNw
						}
					}

					helper.GetTFConfigWithSingleResource(t_name, *x, []interface{}{testConfig}, &tfConfig)
					// fmt.Println("Step 18 - Terraform Config:\n", *tfConfig)
					return *tfConfig
				}(),
			},

			// Step 24: Remove all NETWORK attachments first
			{
				Config: func() string {
					*stepCount++
					t_name := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Remove all parent network attachments
					for nwName, nw := range testConfig.MsdParentNetworks.Networks {
						nw.Attachments = make(map[string]helper.MsdAttachment)
						testConfig.MsdParentNetworks.Networks[nwName] = nw
					}

					// Remove all child network attachments
					for i := range testConfig.MsdChildNetworks {
						for nwName, childNw := range testConfig.MsdChildNetworks[i].Networks {
							childNw.Attachments = make(map[string]helper.MsdAttachment)
							testConfig.MsdChildNetworks[i].Networks[nwName] = childNw
						}
					}

					helper.GetTFConfigWithSingleResource(t_name, *x, []interface{}{testConfig}, &tfConfig)
					// fmt.Println("Step 19 - Terraform Config:\n", *tfConfig)
					return *tfConfig
				}(),
				ExpectNonEmptyPlan: true,
			},

			// Step 25: Remove all VRF attachments (now safe since networks are detached)
			{
				Config: func() string {
					*stepCount++
					t_name := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Remove all parent VRF attachments
					for vrfName, vrf := range testConfig.MsdParentVrfs.Vrfs {
						vrf.AttachList = make(map[string]helper.MsdAttachment)
						testConfig.MsdParentVrfs.Vrfs[vrfName] = vrf
					}

					// Remove all child VRF attachments
					for i := range testConfig.MsdChildVrfs {
						for vrfName, childVrf := range testConfig.MsdChildVrfs[i].Vrfs {
							childVrf.AttachList = make(map[string]helper.MsdAttachment)
							testConfig.MsdChildVrfs[i].Vrfs[vrfName] = childVrf
						}
					}

					helper.GetTFConfigWithSingleResource(t_name, *x, []interface{}{testConfig}, &tfConfig)
					// fmt.Println("Step 20 - Terraform Config:\n", *tfConfig)
					return *tfConfig
				}(),
			},

			// Step 26: Re-add all attachments with deploy flags set to false
			{
				Config: func() string {
					*stepCount++
					t_name := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Build map of active child fabrics
					activeFabricMap := make(map[string]bool)
					for _, childVrfConfig := range testConfig.MsdChildVrfs {
						activeFabricMap[childVrfConfig.FabricName] = true
					}

					// Re-add all parent VRF attachments with deploy=false
					for vrfName, vrf := range testConfig.MsdParentVrfs.Vrfs {
						vrf.AttachList = make(map[string]helper.MsdAttachment)
						for _, childFabric := range cfg.NDFC.Msd.ChildFabrics {
							if !activeFabricMap[childFabric.Name] {
								continue
							}
							for _, switchSerial := range childFabric.Switches {
								vrf.AttachList[switchSerial] = helper.MsdAttachment{
									DeployThisAttachment: false,
									Fabric:               childFabric.Name,
								}
							}
						}
						testConfig.MsdParentVrfs.Vrfs[vrfName] = vrf
					}

					// Re-add all child VRF attachments with deploy=false
					for i := range testConfig.MsdChildVrfs {
						childVrfConfig := &testConfig.MsdChildVrfs[i]
						var childFabric *helper.MsdChildFabric
						for j := range cfg.NDFC.Msd.ChildFabrics {
							if cfg.NDFC.Msd.ChildFabrics[j].Name == childVrfConfig.FabricName {
								childFabric = &cfg.NDFC.Msd.ChildFabrics[j]
								break
							}
						}
						if childFabric == nil {
							continue
						}

						for vrfName, childVrf := range childVrfConfig.Vrfs {
							childVrf.AttachList = make(map[string]helper.MsdAttachment)
							for _, switchSerial := range childFabric.Switches {
								childVrf.AttachList[switchSerial] = helper.MsdAttachment{
									DeployThisAttachment: false,
								}
							}
							childVrfConfig.Vrfs[vrfName] = childVrf
						}
					}

					// Re-add all parent network attachments with deploy=false
					for nwName, nw := range testConfig.MsdParentNetworks.Networks {
						nw.Attachments = make(map[string]helper.MsdAttachment)
						for _, childFabric := range cfg.NDFC.Msd.ChildFabrics {
							if !activeFabricMap[childFabric.Name] {
								continue
							}
							for _, switchSerial := range childFabric.Switches {
								nw.Attachments[switchSerial] = helper.MsdAttachment{
									DeployThisAttachment: false,
									Fabric:               childFabric.Name,
								}
							}
						}
						testConfig.MsdParentNetworks.Networks[nwName] = nw
					}

					// Re-add all child network attachments with deploy=false
					for i := range testConfig.MsdChildNetworks {
						childNetworkConfig := &testConfig.MsdChildNetworks[i]
						var childFabric *helper.MsdChildFabric
						for j := range cfg.NDFC.Msd.ChildFabrics {
							if cfg.NDFC.Msd.ChildFabrics[j].Name == childNetworkConfig.FabricName {
								childFabric = &cfg.NDFC.Msd.ChildFabrics[j]
								break
							}
						}
						if childFabric == nil {
							continue
						}

						for nwName, childNw := range childNetworkConfig.Networks {
							childNw.Attachments = make(map[string]helper.MsdAttachment)
							for _, switchSerial := range childFabric.Switches {
								childNw.Attachments[switchSerial] = helper.MsdAttachment{
									DeployThisAttachment: false,
								}
							}
							childNetworkConfig.Networks[nwName] = childNw
						}
					}

					helper.GetTFConfigWithSingleResource(t_name, *x, []interface{}{testConfig}, &tfConfig)
					// fmt.Println("Step 21 - Terraform Config:\n", *tfConfig)
					return *tfConfig
				}(),
			},
			// Step 27: Deploy both VRFs and networks using global flag
			{
				Config: func() string {
					*stepCount++
					t_name := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					modifyVrfParameters(testConfig, cfg, []int{1, 5}, 9000, 54000, 2, 3)
					modifyNetworkParameters(testConfig, cfg, []int{1, 3}, 400, 9100, 54001, " SVI Modified", false, "239.2.2.", 2)
					setGlobalDeployFlags(testConfig, true, true)

					helper.GetTFConfigWithSingleResource(t_name, *x, []interface{}{testConfig}, &tfConfig)
					// fmt.Println("Step 22 - Terraform Config:\n", *tfConfig)
					return *tfConfig
				}(),
			},

			// Step 28: Clean up - Remove all networks
			{
				Config: func() string {
					*stepCount++
					t_name := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					testConfig.MsdParentNetworks = nil
					testConfig.MsdChildNetworks = nil

					helper.GetTFConfigWithSingleResource(t_name, *x, []interface{}{testConfig}, &tfConfig)
					// fmt.Println("Step 23 - Terraform Config:\n", *tfConfig)
					return *tfConfig
				}(),
				ExpectNonEmptyPlan: true,
			},

			// Step 29: Clean up - Remove all VRFs
			{
				Config: func() string {
					*stepCount++
					t_name := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					testConfig.MsdParentVrfs = nil
					testConfig.MsdChildVrfs = nil

					helper.GetTFConfigWithSingleResource(t_name, *x, []interface{}{testConfig}, &tfConfig)
					// fmt.Println("Step 24 - Terraform Config:\n", *tfConfig)
					return *tfConfig
				}(),
			},
		},
	})
}
