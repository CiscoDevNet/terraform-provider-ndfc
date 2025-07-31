// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"terraform-provider-ndfc/internal/provider/resources/resource_interface_common"
	"terraform-provider-ndfc/internal/provider/resources/resource_vpc_pair"
	helper "terraform-provider-ndfc/internal/provider/testing"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

var rsVpcType string = "vpc"

// Helper function to set up VPC pair required for VPC interfaces
func setupVpcPair(tname string, configPtr *string) *resource_vpc_pair.NDFCVpcPairModel {
	x := &map[string]string{
		"RscType":    "ndfc_vpc_pair",
		"RscSubType": "vpc_pair",
		"RscName":    "test_vpc_pair",
		"User":       helper.GetConfig("vpc_pair").NDFC.User,
		"Password":   helper.GetConfig("vpc_pair").NDFC.Password,
		"Host":       helper.GetConfig("vpc_pair").NDFC.URL,
		"Insecure":   helper.GetConfig("vpc_pair").NDFC.Insecure,
	}
	vpcPairRsc := new(resource_vpc_pair.NDFCVpcPairModel)
	vpcPairRsc.UseVirtualPeerlink = new(bool)
	*vpcPairRsc.UseVirtualPeerlink = false
	helper.GenerateVpcPairResource(&vpcPairRsc, helper.GetConfig("vpc_pair").NDFC.VpcPair, false, true)
	helper.GetTFConfigWithSingleResource(tname, *x, []any{vpcPairRsc}, &configPtr)
	return vpcPairRsc
}

// Create VPC resource with 3 entries, add 3 more
func TestAccInterfaceVpcResourceCreateAndAdd(t *testing.T) {
	x := &map[string]string{
		"RscType":    "ndfc_interface_vpc",
		"RscSubType": "vpc",
		"RscName":    "vpc_pair,test_vpc",
		"User":       helper.GetConfig("vpc").NDFC.User,
		"Password":   helper.GetConfig("vpc").NDFC.Password,
		"Host":       helper.GetConfig("vpc").NDFC.URL,
		"Insecure":   helper.GetConfig("vpc").NDFC.Insecure,
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

	stepCount := new(int)
	*stepCount = 0

	// Create the interface resources
	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "vpc") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					tName := fmt.Sprintf("%s_%d", t.Name(), 1)
					*stepCount = 1

					// First set up the required VPC pair
					vpcPairRsc := new(resource_vpc_pair.NDFCVpcPairModel)
					vpcPairRsc.UseVirtualPeerlink = new(bool)
					*vpcPairRsc.UseVirtualPeerlink = false
					helper.GenerateVpcPairResource(&vpcPairRsc, helper.GetConfig("vpc_pair").NDFC.VpcPair, false, true)
					// Now set up the VPC interfaces
					// We need to use the serial numbers from the VPC pair
					// Format for VPC interfaces is "switch1~switch2"
					vpcSerial := fmt.Sprintf("%s~%s", vpcPairRsc.SerialNumbers[0], vpcPairRsc.SerialNumbers[1])
					serials := []string{vpcSerial}

					// Create 3 VPC interfaces
					// Global serial
					helper.GenerateIntfResource(&intfRsc, 1, 3, rsVpcType, true, serials, true, false)
					helper.GetTFConfigWithSingleResource(tName, *x, []any{vpcPairRsc, intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: func() resource.TestCheckFunc {
					funcs := InterfaceVpcModelHelperStateCheck("ndfc_interface_vpc.test_vpc", *intfRsc, path.Empty())
					return resource.ComposeTestCheckFunc(funcs...)
				}(),
			},
			{
				Config: func() string {
					tName := fmt.Sprintf("%s_%d", t.Name(), 2)
					*stepCount = 2
					// First set up the required VPC pair
					vpcPairRsc := new(resource_vpc_pair.NDFCVpcPairModel)
					vpcPairRsc.UseVirtualPeerlink = new(bool)
					*vpcPairRsc.UseVirtualPeerlink = false
					helper.GenerateVpcPairResource(&vpcPairRsc, helper.GetConfig("vpc_pair").NDFC.VpcPair, false, true)

					// Format for VPC interfaces is "switch1~switch2"
					vpcSerial := fmt.Sprintf("%s~%s", vpcPairRsc.SerialNumbers[0], vpcPairRsc.SerialNumbers[1])
					serials := []string{vpcSerial}
					t.Logf("Serials: %v", serials)

					// Add 3 more VPC interfaces (total 6)
					helper.GenerateIntfResource(&intfRsc, 4, 3, rsVpcType, true, serials, true, true)
					helper.GetTFConfigWithSingleResource(tName, *x, []any{vpcPairRsc, intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: func() resource.TestCheckFunc {
					funcs := InterfaceVpcModelHelperStateCheck("ndfc_interface_vpc.test_vpc", *intfRsc, path.Empty())
					return resource.ComposeTestCheckFunc(funcs...)
				}(),
			},
		},
	})
}

func TestAccInterfaceVpcResourceCheckInfLevelSerials(t *testing.T) {
	x := &map[string]string{
		"RscType":    "ndfc_interface_vpc",
		"RscSubType": "vpc",
		"RscName":    "vpc_pair,test_vpc",
		"User":       helper.GetConfig("vpc").NDFC.User,
		"Password":   helper.GetConfig("vpc").NDFC.Password,
		"Host":       helper.GetConfig("vpc").NDFC.URL,
		"Insecure":   helper.GetConfig("vpc").NDFC.Insecure,
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

	stepCount := new(int)
	*stepCount = 0

	// Create the interface resources
	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "vpc") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					tName := fmt.Sprintf("%s_%d", t.Name(), 1)
					*stepCount = 1

					// First set up the required VPC pair
					vpcPairRsc := new(resource_vpc_pair.NDFCVpcPairModel)
					vpcPairRsc.UseVirtualPeerlink = new(bool)
					*vpcPairRsc.UseVirtualPeerlink = false
					helper.GenerateVpcPairResource(&vpcPairRsc, helper.GetConfig("vpc_pair").NDFC.VpcPair, false, true)
					// Now set up the VPC interfaces
					// We need to use the serial numbers from the VPC pair
					// Format for VPC interfaces is "switch1~switch2"
					vpcSerial := fmt.Sprintf("%s~%s", vpcPairRsc.SerialNumbers[0], vpcPairRsc.SerialNumbers[1])
					serials := []string{vpcSerial}

					// Create 3 VPC interfaces
					// Global serial false
					helper.GenerateIntfResource(&intfRsc, 1, 3, rsVpcType, true, serials, false, false)
					helper.GetTFConfigWithSingleResource(tName, *x, []any{vpcPairRsc, intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: func() resource.TestCheckFunc {
					funcs := InterfaceVpcModelHelperStateCheck("ndfc_interface_vpc.test_vpc", *intfRsc, path.Empty())
					return resource.ComposeTestCheckFunc(funcs...)
				}(),
			},
			{
				Config: func() string {
					tName := fmt.Sprintf("%s_%d", t.Name(), 2)
					*stepCount = 2
					// First set up the required VPC pair
					vpcPairRsc := new(resource_vpc_pair.NDFCVpcPairModel)
					vpcPairRsc.UseVirtualPeerlink = new(bool)
					*vpcPairRsc.UseVirtualPeerlink = false
					helper.GenerateVpcPairResource(&vpcPairRsc, helper.GetConfig("vpc_pair").NDFC.VpcPair, false, true)

					// Format for VPC interfaces is "switch1~switch2"
					vpcSerial := fmt.Sprintf("%s~%s", vpcPairRsc.SerialNumbers[0], vpcPairRsc.SerialNumbers[1])
					serials := []string{vpcSerial}
					t.Logf("Serials: %v", serials)

					// Add 3 more VPC interfaces (total 6)
					helper.GenerateIntfResource(&intfRsc, 4, 3, rsVpcType, true, serials, false, true)
					helper.GetTFConfigWithSingleResource(tName, *x, []any{vpcPairRsc, intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: func() resource.TestCheckFunc {
					funcs := InterfaceVpcModelHelperStateCheck("ndfc_interface_vpc.test_vpc", *intfRsc, path.Empty())
					return resource.ComposeTestCheckFunc(funcs...)
				}(),
			},
		},
	})
}
