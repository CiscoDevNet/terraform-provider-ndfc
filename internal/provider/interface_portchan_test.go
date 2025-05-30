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
	"terraform-provider-ndfc/internal/provider/resources/resource_interface_common"
	helper "terraform-provider-ndfc/internal/provider/testing"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccInterfacePortchannelResourceBasic(t *testing.T) {
	x := &map[string]string{
		"RscType":    "ndfc_interface_portchannel",
		"RscSubType": "portchannel",
		"RscName":    "test_pc",
		"User":       helper.GetConfig("portchannel").NDFC.User,
		"Password":   helper.GetConfig("portchannel").NDFC.Password,
		"Host":       helper.GetConfig("portchannel").NDFC.URL,
		"Insecure":   helper.GetConfig("portchannel").NDFC.Insecure,
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
	// Create a new instance of the NDFC client
	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t, "portchannel")
			helper.EthIntf = 10
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateIntfResource(&intfRsc, 200, 5, "portchannel", true, helper.GetConfig("portchannel").NDFC.Switches, true, false)
					(*x)["RscName"] = "pc_intf_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfacePortchannelModelHelperStateCheck("ndfc_interface_portchannel.pc_intf_test", *intfRsc, path.Empty())...),
			},
		},
	})
}

// Create Vlan resource with 10 entries, add 10 more
func TestAccInterfacePortchannelResourceCreateAndAdd(t *testing.T) {

	x := &map[string]string{
		"RscType":    "ndfc_interface_portchannel",
		"RscSubType": "portchannel",
		"RscName":    "test_pc",
		"User":       helper.GetConfig("portchannel").NDFC.User,
		"Password":   helper.GetConfig("portchannel").NDFC.Password,
		"Host":       helper.GetConfig("portchannel").NDFC.URL,
		"Insecure":   helper.GetConfig("portchannel").NDFC.Insecure,
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
	// Create a new instance of the NDFC client

	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t, "portchannel")
			helper.EthIntf = 10
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		//CheckDestroy:             testAccCheckVrfBulkResourceDestroy(networkRsc),
		Steps: []resource.TestStep{
			{
				// Create a new resource with 10 interfaces
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					//Starting from 10 to consider anything pre-existing on ndfc
					helper.GenerateIntfResource(&intfRsc, 200, 5, "portchannel", true, helper.GetConfig("portchannel").NDFC.Switches, true, false)
					(*x)["RscName"] = "pc_intf_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfacePortchannelModelHelperStateCheck("ndfc_interface_portchannel.pc_intf_test", *intfRsc, path.Empty())...),
			},
			// Add some more intfs
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateIntfResource(&intfRsc, 205, 5, "portchannel", true, helper.GetConfig("portchannel").NDFC.Switches, true, true)
					(*x)["RscName"] = "pc_intf_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfacePortchannelModelHelperStateCheck("ndfc_interface_portchannel.pc_intf_test", *intfRsc, path.Empty())...),
			},
		},
	})
}

// Create with 20 entries, delete 10
func TestAccInterfacePortchannelResourceCreateAndReduce(t *testing.T) {
	x := &map[string]string{
		"RscType":    "ndfc_interface_portchannel",
		"RscSubType": "portchannel",
		"RscName":    "test_pc",
		"User":       helper.GetConfig("portchannel").NDFC.User,
		"Password":   helper.GetConfig("portchannel").NDFC.Password,
		"Host":       helper.GetConfig("portchannel").NDFC.URL,
		"Insecure":   helper.GetConfig("portchannel").NDFC.Insecure,
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
	// Create a new instance of the NDFC client

	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t, "portchannel")
			helper.EthIntf = 10
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		//CheckDestroy:             testAccCheckVrfBulkResourceDestroy(networkRsc),
		Steps: []resource.TestStep{
			{
				// Create a new resource with 20 interfaces
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					//Starting from 20 to consider anything pre-existing on ndfc
					helper.GenerateIntfResource(&intfRsc, 200, 10, "portchannel", true, helper.GetConfig("portchannel").NDFC.Switches, true, false)
					(*x)["RscName"] = "pc_intf_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfacePortchannelModelHelperStateCheck("ndfc_interface_portchannel.pc_intf_test", *intfRsc, path.Empty())...),
			},
			// Delete 10 interfaces
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateIntfResource(&intfRsc, 200, -5, "portchannel", true, helper.GetConfig("portchannel").NDFC.Switches, true, true)
					(*x)["RscName"] = "pc_intf_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfacePortchannelModelHelperStateCheck("ndfc_interface_portchannel.pc_intf_test", *intfRsc, path.Empty())...),
			},
		},
	})
}

// Create with 20 entries, modify 5
func TestAccInterfacePortchannelResourceCreateAndModify(t *testing.T) {

	x := &map[string]string{
		"RscType":    "ndfc_interface_portchannel",
		"RscSubType": "portchannel",
		"RscName":    "test_pc",
		"User":       helper.GetConfig("portchannel").NDFC.User,
		"Password":   helper.GetConfig("portchannel").NDFC.Password,
		"Host":       helper.GetConfig("portchannel").NDFC.URL,
		"Insecure":   helper.GetConfig("portchannel").NDFC.Insecure,
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
	// Create a new instance of the NDFC client

	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t, "portchannel")
			helper.EthIntf = 10
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		//CheckDestroy:             testAccCheckVrfBulkResourceDestroy(networkRsc),
		Steps: []resource.TestStep{
			{
				// Create a new resource with 20 interfaces
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					//Starting from 20 to consider anything pre-existing on ndfc
					helper.GenerateIntfResource(&intfRsc, 200, 20, "portchannel", true, helper.GetConfig("portchannel").NDFC.Switches, true, false)
					(*x)["RscName"] = "pc_intf_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfacePortchannelModelHelperStateCheck("ndfc_interface_portchannel.pc_intf_test", *intfRsc, path.Empty())...),
			},
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.ModifyInterface(&intfRsc, 200, 5, "portchannel", map[string]interface{}{
						"portchannelMode":   "passive",
						"CopyPoDescription": "false",
						"memberInterfaces":  "yes",
					})
					(*x)["RscName"] = "pc_intf_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfacePortchannelModelHelperStateCheck("ndfc_interface_portchannel.pc_intf_test", *intfRsc, path.Empty())...),
			},
		},
	})
}

// Create with 10 entries, add 10 delete 5 and modify 5 all in one shot
func TestAccInterfacePortchannelResourceCombinedUpdate(t *testing.T) {

	x := &map[string]string{
		"RscType":    "ndfc_interface_portchannel",
		"RscSubType": "portchannel",
		"RscName":    "test_pc",
		"User":       helper.GetConfig("portchannel").NDFC.User,
		"Password":   helper.GetConfig("portchannel").NDFC.Password,
		"Host":       helper.GetConfig("portchannel").NDFC.URL,
		"Insecure":   helper.GetConfig("portchannel").NDFC.Insecure,
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
	// Create a new instance of the NDFC client

	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t, "portchannel")
			helper.EthIntf = 10
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		//CheckDestroy:             testAccCheckVrfBulkResourceDestroy(networkRsc),
		Steps: []resource.TestStep{
			{
				// Create a new resource with 20 interfaces
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					//Starting from 20 to consider anything pre-existing on ndfc
					helper.GenerateIntfResource(&intfRsc, 10, 10, "portchannel", true, helper.GetConfig("portchannel").NDFC.Switches, true, false)
					(*x)["RscName"] = "pc_intf_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfacePortchannelModelHelperStateCheck("ndfc_interface_portchannel.pc_intf_test", *intfRsc, path.Empty())...),
			},
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					//Add 10
					helper.GenerateIntfResource(&intfRsc, 20, 10, "portchannel", true, helper.GetConfig("portchannel").NDFC.Switches, true, true)
					//Delete 5
					helper.GenerateIntfResource(&intfRsc, 10, -5, "portchannel", true, helper.GetConfig("portchannel").NDFC.Switches, true, true)
					//Modify 5
					helper.ModifyInterface(&intfRsc, 20, 5, "portchannel", map[string]interface{}{
						"admin_state": "up",
					})
					(*x)["RscName"] = "pc_intf_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfacePortchannelModelHelperStateCheck("ndfc_interface_portchannel.pc_intf_test", *intfRsc, path.Empty())...),
			},
		},
	})

}
