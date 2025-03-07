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
	helper "terraform-provider-ndfc/internal/provider/testing"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Contains acceptance tests for all fabric types except ipfm.
func TestAccFabricVxlanEvpnResource(t *testing.T) {
	cfg := map[string]string{
		"FabricType": "fabric_vxlan_evpn",
		"User":       helper.GetConfig("ndfc_fabric_vxlan_evpn").NDFC.User,
		"Password":   helper.GetConfig("ndfc_fabric_vxlan_evpn").NDFC.Password,
		"Host":       helper.GetConfig("ndfc_fabric_vxlan_evpn").NDFC.URL,
		"Insecure":   helper.GetConfig("ndfc_fabric_vxlan_evpn").NDFC.Insecure,
	}
	stepCount := new(int)
	*stepCount = 0
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "fabric_vxlan_evpn") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{ // Create fabric and update bfd_enable
				Config: func() string {
					*stepCount++
					tt := fmt.Sprintf("%s%d", t.Name(), *stepCount)
					return helper.GenerateFabricConfig(tt, cfg, helper.Base_file)
				}(),
				//Check: resource.ComposeTestCheckFunc(FabricVxlanEvpnModelHelperStateCheck("ndfc_fabric_vxlan_evpn.test_resource_fabric_vxlan_evpn_1", FabricRsc, path.Empty())...),
			}, { // update bgp_as, it will fail as bgp as cannot be updated
				Config: func() string {
					*stepCount++
					tt := fmt.Sprintf("%s%d", t.Name(), *stepCount)
					return helper.GenerateFabricConfig(tt, cfg, helper.Bgp_as_change_file)
				}(),
				//Check:       resource.ComposeTestCheckFunc(FabricVxlanEvpnModelHelperStateCheck("ndfc_fabric_vxlan_evpn.test_resource_fabric_vxlan_evpn_1", FabricRsc, path.Empty())...),
				ExpectError: regexp.MustCompile("bgp_as cannot be updated"),
			}, { // modify fields like CDP_ENABLE, BFD_ENABLE, enable_trm with wrong configs and MTU with different values and deploy true
				Config: func() string {
					*stepCount++
					tt := fmt.Sprintf("%s%d", t.Name(), *stepCount)
					return helper.GenerateFabricConfig(tt, cfg, helper.Enable_trm_config_file)
				}(),
				//Check: resource.ComposeTestCheckFunc(FabricVxlanEvpnModelHelperStateCheck("ndfc_fabric_vxlan_evpn.test_resource_fabric_vxlan_evpn_1", FabricRsc, path.Empty())...),
				ExpectError: regexp.MustCompile("l3vni_mcast_group is required for TRM"),
			}, { // modify fields like CDP_ENABLE, BFD_ENABLE, enable_trm with correct configs and MTU with different values and deploy true
				Config: func() string {
					*stepCount++
					tt := fmt.Sprintf("%s%d", t.Name(), *stepCount)
					return helper.GenerateFabricConfig(tt, cfg, helper.Modified_file)
				}(),
				//Check: resource.ComposeTestCheckFunc(FabricVxlanEvpnModelHelperStateCheck("ndfc_fabric_vxlan_evpn.test_resource_fabric_vxlan_evpn_1", FabricRsc, path.Empty())...),
			},
		},
	})

}

func TestAccFabricVxlanMsdResource(t *testing.T) {
	cfg := map[string]string{
		"FabricType": "fabric_vxlan_msd",
		"User":       helper.GetConfig("ndfc_fabric_vxlan_msd").NDFC.User,
		"Password":   helper.GetConfig("ndfc_fabric_vxlan_msd").NDFC.Password,
		"Host":       helper.GetConfig("ndfc_fabric_vxlan_msd").NDFC.URL,
		"Insecure":   helper.GetConfig("ndfc_fabric_vxlan_msd").NDFC.Insecure,
	}
	stepCount := new(int)
	*stepCount = 0
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "fabric_vxlan_msd") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{ // Create fabric
				Config: func() string {
					*stepCount++
					tt := fmt.Sprintf("%s%d", t.Name(), *stepCount)
					return helper.GenerateFabricConfig(tt, cfg, helper.Base_file)
				}(),
				//Check: resource.ComposeTestCheckFunc(FabricVxlanMsdModelHelperStateCheck("ndfc_fabric_vxlan_msd.test_resource_fabric_vxlan_msd_1", FabricRsc, path.Empty())...),
			}, { // modify fields like ANYCAST_GW_MAC, BGW_ROUTING_TAG and DELAY_RESTORE with different values and deploy true
				Config: func() string {
					*stepCount++
					tt := fmt.Sprintf("%s%d", t.Name(), *stepCount)
					return helper.GenerateFabricConfig(tt, cfg, helper.Modified_file)
				}(),
				//Check: resource.ComposeTestCheckFunc(FabricVxlanMsdModelHelperStateCheck("ndfc_fabric_vxlan_msd.test_resource_fabric_vxlan_msd_1", FabricRsc, path.Empty())...),
			},
		},
	})

}

func TestAccFabricLanClassicResource(t *testing.T) {
	cfg := map[string]string{
		"FabricType": "fabric_lan_classic",
		"User":       helper.GetConfig("ndfc_fabric_lan_classic").NDFC.User,
		"Password":   helper.GetConfig("ndfc_fabric_lan_classic").NDFC.Password,
		"Host":       helper.GetConfig("ndfc_fabric_lan_classic").NDFC.URL,
		"Insecure":   helper.GetConfig("ndfc_fabric_lan_classic").NDFC.Insecure,
	}
	stepCount := new(int)
	*stepCount = 0
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "fabric_lan_classic") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{ // Create fabric and update bfd_enable
				Config: func() string {
					*stepCount++
					tt := fmt.Sprintf("%s%d", t.Name(), *stepCount)
					return helper.GenerateFabricConfig(tt, cfg, helper.Base_file)
				}(),
				//Check: resource.ComposeTestCheckFunc(FabricLanClassicModelHelperStateCheck("ndfc_fabric_lan_classic.test_resource_fabric_lan_classic_1", FabricRsc, path.Empty())...),
			}, { // modify fields like CDP_ENABLE, BFD_ENABLE and MTU with different values
				Config: func() string {
					*stepCount++
					tt := fmt.Sprintf("%s%d", t.Name(), *stepCount)
					return helper.GenerateFabricConfig(tt, cfg, helper.Modified_file)
				}(),
				//Check: resource.ComposeTestCheckFunc(FabricLanClassicModelHelperStateCheck("ndfc_fabric_lan_classic.test_resource_fabric_lan_classic_1", FabricRsc, path.Empty())...),
			},
		},
	})

}
func TestAccFabricMsiteExtNetResource(t *testing.T) {

	cfg := map[string]string{
		"FabricType": "fabric_msite_ext_net",
		"User":       helper.GetConfig("ndfc_fabric_msite_ext_net").NDFC.User,
		"Password":   helper.GetConfig("ndfc_fabric_msite_ext_net").NDFC.Password,
		"Host":       helper.GetConfig("ndfc_fabric_msite_ext_net").NDFC.URL,
		"Insecure":   helper.GetConfig("ndfc_fabric_msite_ext_net").NDFC.Insecure,
	}
	stepCount := new(int)
	*stepCount = 0
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "fabric_msite_ext_net") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{ // Create fabric and update bfd_enable
				Config: func() string {
					*stepCount++
					tt := fmt.Sprintf("%s%d", t.Name(), *stepCount)
					return helper.GenerateFabricConfig(tt, cfg, helper.Base_file)
				}(),
				//Check: resource.ComposeTestCheckFunc(FabricMsiteExtNetModelHelperStateCheck("ndfc_fabric_msite_ext_net.test_resource_fabric_msite_ext_net_1", FabricRsc, path.Empty())...),
			}, { // update bgp_as, it capture failure
				Config: func() string {
					*stepCount++
					tt := fmt.Sprintf("%s%d", t.Name(), *stepCount)
					return helper.GenerateFabricConfig(tt, cfg, helper.Bgp_as_change_file)
				}(),
				//Check:       resource.ComposeTestCheckFunc(FabricMsiteExtNetModelHelperStateCheck("ndfc_fabric_msite_ext_net.test_resource_fabric_msite_ext_net_1", FabricRsc, path.Empty())...),
				ExpectError: regexp.MustCompile("bgp_as cannot be updated"),
			}, { // modify fields like CDP_ENABLE, BFD_ENABLE and MTU with different values
				Config: func() string {
					*stepCount++
					tt := fmt.Sprintf("%s%d", t.Name(), *stepCount)
					return helper.GenerateFabricConfig(tt, cfg, helper.Modified_file)
				}(),
				//Check: resource.ComposeTestCheckFunc(FabricMsiteExtNetModelHelperStateCheck("ndfc_fabric_msite_ext_net.test_resource_fabric_msite_ext_net_1", FabricRsc, path.Empty())...),
			},
		},
	})

}
