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
	"terraform-provider-ndfc/internal/provider/resources/resource_vpc_pair"
	helper "terraform-provider-ndfc/internal/provider/testing"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

var ErrorVirtualLinkNotSupported = "doesn't support Virtual Fabric Peering"

func TestAccCreateVpcPair(t *testing.T) {

	x := &map[string]string{
		"RscType":    "ndfc_vpc_pair",
		"RscSubType": "vpc_pair",
		"RscName":    "test_vpc_pair",
		"User":       helper.GetConfig("vpc_pair").NDFC.User,
		"Password":   helper.GetConfig("vpc_pair").NDFC.Password,
		"Host":       helper.GetConfig("vpc_pair").NDFC.URL,
		"Insecure":   helper.GetConfig("vpc_pair").NDFC.Insecure,
	}

	tf_config := new(string)
	*tf_config = `provider "ndfc" {
		host     = "https://"
		username = "admin"
		password = "admin!@#"
		domain   = "example.com"
		insecure = true
		}
		resource ndfc_vrf_bulk "net_test" {
			fabric_name = "dummy"
		}`

	// Create a new instance of the NDFC client

	vpcPairRsc := new(resource_vpc_pair.NDFCVpcPairModel)
	tfvpcPairRsc := new(resource_vpc_pair.VpcPairModel)
	vpcPairRsc.UseVirtualPeerlink = new(bool)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "vpc_pair") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpcPairResourceDestroy(tfvpcPairRsc),
		Steps: []resource.TestStep{
			{ // Create vpc pair and check resource state
				Config: func() string {
					tName := fmt.Sprintf("%s_%d", t.Name(), 1)
					helper.GenerateVpcPairResource(&vpcPairRsc, helper.GetConfig("vpc_pair").NDFC.Switches, false)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vpcPairRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(VpcPairModelHelperStateCheck("ndfc_vpc_pair.test_vpc_pair", *vpcPairRsc, path.Empty())...),
			}, { // create and destroy vpc pair
				Config: func() string {
					tName := fmt.Sprintf("%s_%d", t.Name(), 1)
					helper.GenerateVpcPairResource(&vpcPairRsc, helper.GetConfig("vpc_pair").NDFC.Switches, false)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vpcPairRsc}, &tf_config)
					return *tf_config
				}(),
				Destroy: true,
			}, { // create with virtual link true
				Config: func() string {
					tName := fmt.Sprintf("%s_%d", t.Name(), 1)
					helper.GenerateVpcPairResource(&vpcPairRsc, helper.GetConfig("vpc_pair").NDFC.Switches, true)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vpcPairRsc}, &tf_config)
					return *tf_config
				}(),
				Check:       resource.ComposeTestCheckFunc(VpcPairModelHelperStateCheck("ndfc_vpc_pair.test_vpc_pair", *vpcPairRsc, path.Empty())...),
				ExpectError: regexp.MustCompile("doesn't support Virtual Fabric Peering"),
			},
			{ // create with virtual link true
				Config: func() string {
					tName := fmt.Sprintf("%s_%d", t.Name(), 1)
					helper.GenerateVpcPairResource(&vpcPairRsc, helper.GetConfig("vpc_pair").NDFC.Switches, true)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vpcPairRsc}, &tf_config)
					return *tf_config
				}(),
				Destroy: true,
			},
		},
	})

}
func testAccCheckVpcPairResourceDestroy(vpcPair *resource_vpc_pair.VpcPairModel) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources["ndfc_vpc_pair.test_vpc_pair"]
		if !ok {
			return nil
		}
		if rs.Primary.ID == "" {
			return nil
		}
		return fmt.Errorf("VPC Pair still exists")
	}

}
