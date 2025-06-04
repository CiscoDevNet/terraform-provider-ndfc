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
)

var ErrorVirtualLinkNotSupported = "doesn't support Virtual Fabric Peering"

func TestAccVPCPairResourceCreateVpcPair(t *testing.T) {

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
		resource ndfc_vrfs "net_test" {
			fabric_name = "dummy"
		}`

	// Create a new instance of the NDFC client

	vpcPairRsc := new(resource_vpc_pair.NDFCVpcPairModel)
	vpcPairRsc.UseVirtualPeerlink = new(bool)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "vpc_pair") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					tName := fmt.Sprintf("%s_%d", t.Name(), 1)
					helper.GenerateVpcPairResource(&vpcPairRsc, helper.GetConfig("vpc_pair").NDFC.VpcPair, false, true)
					helper.GetTFConfigWithSingleResource(tName, *x, []any{vpcPairRsc}, &tf_config)
					return *tf_config
				}(),
				Check: func() resource.TestCheckFunc {
					funcs1 := VpcPairModelHelperStateCheck("ndfc_vpc_pair.test_vpc_pair_1", *vpcPairRsc, path.Empty())
					//funcs2 := VpcPairModelHelperStateCheck("ndfc_vpc_pair.test_vpc_pair_2", *vpcPairRsc, path.Empty())
					//allFuncs := append(funcs1, funcs2...)
					return resource.ComposeTestCheckFunc(funcs1...)
				}(),
			},
			{
				Config: func() string {
					tName := fmt.Sprintf("%s_%d", t.Name(), 1)
					helper.GenerateVpcPairResource(&vpcPairRsc, helper.GetConfig("vpc_pair").NDFC.VpcPair, true, true)
					helper.GetTFConfigWithSingleResource(tName, *x, []any{vpcPairRsc}, &tf_config)
					return *tf_config
				}(),
				Check: func() resource.TestCheckFunc {
					funcs1 := VpcPairModelHelperStateCheck("ndfc_vpc_pair.test_vpc_pair_1", *vpcPairRsc, path.Empty())
					// funcs2 := VpcPairModelHelperStateCheck("ndfc_vpc_pair.test_vpc_pair_2", *vpcPairRsc, path.Empty())
					//allFuncs := append(funcs1, funcs2...)
					return resource.ComposeTestCheckFunc(funcs1...)
				}(),
				ExpectError: regexp.MustCompile(".*doesn't support Virtual Fabric.*"),
			},
		},
	})

}
