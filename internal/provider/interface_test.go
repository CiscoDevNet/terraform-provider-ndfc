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
	"terraform-provider-ndfc/internal/provider/types"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// ============================================================================
// Ethernet Interface CRUD Test with IP Keys (Global Serial)
// ============================================================================

func TestAccInterfaceIPKeyGlobalEthernetCRUD(t *testing.T) {
	x := &map[string]string{
		"RscType":    "ndfc_interface_ethernet",
		"RscSubType": "ethernet",
		"RscName":    "test_ethernet",
		"User":       helper.GetConfig("ethernet").NDFC.User,
		"Password":   helper.GetConfig("ethernet").NDFC.Password,
		"Host":       helper.GetConfig("ethernet").NDFC.URL,
		"Insecure":   helper.GetConfig("ethernet").NDFC.Insecure,
	}

	tf_config := new(string)
	stepCount := new(int)
	*stepCount = 0
	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "ethernet") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create interfaces using IP keys
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Use SwitchIP instead of Switches for IP-based keys
					helper.GenerateIntfResource(&intfRsc, 10, 5, "ethernet", true, helper.GetConfig("ethernet").NDFC.SwitchIP, true, false)
					(*x)["RscName"] = "eth_intf_ipkey_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfaceEthernetModelHelperStateCheck("ndfc_interface_ethernet.eth_intf_ipkey_test", *intfRsc, path.Empty())...),
			},
			// Step 2: Add more interfaces
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateIntfResource(&intfRsc, 15, 5, "ethernet", true, helper.GetConfig("ethernet").NDFC.SwitchIP, true, true)
					(*x)["RscName"] = "eth_intf_ipkey_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfaceEthernetModelHelperStateCheck("ndfc_interface_ethernet.eth_intf_ipkey_test", *intfRsc, path.Empty())...),
			},
			// Step 3: Modify some interfaces
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.ModifyInterface(&intfRsc, 10, 3, "ethernet", map[string]interface{}{
						"admin_state": "down",
						"mtu":         "jumbo",
						"speed":       "Auto",
						"bpduGuard":   "true",
						"accessVlan":  types.Int64Custom(1000),
					})
					(*x)["RscName"] = "eth_intf_ipkey_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfaceEthernetModelHelperStateCheck("ndfc_interface_ethernet.eth_intf_ipkey_test", *intfRsc, path.Empty())...),
			},
			// Step 4: Delete some interfaces
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateIntfResource(&intfRsc, 10, -3, "ethernet", true, helper.GetConfig("ethernet").NDFC.SwitchIP, true, true)
					(*x)["RscName"] = "eth_intf_ipkey_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfaceEthernetModelHelperStateCheck("ndfc_interface_ethernet.eth_intf_ipkey_test", *intfRsc, path.Empty())...),
			},
		},
	})
}

// ============================================================================
// Port-Channel Interface CRUD Test with IP Keys (Global Serial)
// ============================================================================

func TestAccInterfaceIPKeyGlobalPortchannelCRUD(t *testing.T) {
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
	stepCount := new(int)
	*stepCount = 0
	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t, "portchannel")
			helper.EthIntf = 10
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create port-channels using IP keys
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Use SwitchIP instead of Switches for IP-based keys
					helper.GenerateIntfResource(&intfRsc, 200, 3, "portchannel", true, helper.GetConfig("portchannel").NDFC.SwitchIP, true, false)
					(*x)["RscName"] = "pc_intf_ipkey_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfacePortchannelModelHelperStateCheck("ndfc_interface_portchannel.pc_intf_ipkey_test", *intfRsc, path.Empty())...),
			},
			// Step 2: Add more port-channels
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateIntfResource(&intfRsc, 203, 2, "portchannel", true, helper.GetConfig("portchannel").NDFC.SwitchIP, true, true)
					(*x)["RscName"] = "pc_intf_ipkey_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfacePortchannelModelHelperStateCheck("ndfc_interface_portchannel.pc_intf_ipkey_test", *intfRsc, path.Empty())...),
			},
			// Step 3: Modify some port-channels
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.ModifyInterface(&intfRsc, 200, 2, "portchannel", map[string]interface{}{
						"portchannelMode":   "passive",
						"CopyPoDescription": "false",
					})
					(*x)["RscName"] = "pc_intf_ipkey_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfacePortchannelModelHelperStateCheck("ndfc_interface_portchannel.pc_intf_ipkey_test", *intfRsc, path.Empty())...),
			},
			// Step 4: Delete some port-channels
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateIntfResource(&intfRsc, 200, -2, "portchannel", true, helper.GetConfig("portchannel").NDFC.SwitchIP, true, true)
					(*x)["RscName"] = "pc_intf_ipkey_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfacePortchannelModelHelperStateCheck("ndfc_interface_portchannel.pc_intf_ipkey_test", *intfRsc, path.Empty())...),
			},
		},
	})
}

// ============================================================================
// SVI (VLAN) Interface CRUD Test with IP Keys (Global Serial)
// ============================================================================

func TestAccInterfaceIPKeyGlobalVlanCRUD(t *testing.T) {
	x := &map[string]string{
		"RscType":    "ndfc_interface_vlan",
		"RscSubType": "vlan",
		"RscName":    "test_vlan",
		"User":       helper.GetConfig("vlan").NDFC.User,
		"Password":   helper.GetConfig("vlan").NDFC.Password,
		"Host":       helper.GetConfig("vlan").NDFC.URL,
		"Insecure":   helper.GetConfig("vlan").NDFC.Insecure,
	}

	tf_config := new(string)
	stepCount := new(int)
	*stepCount = 0
	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "vlan") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create SVIs using IP keys
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Use SwitchIP instead of Switches for IP-based keys
					helper.GenerateIntfResource(&intfRsc, 200, 5, "vlan", true, helper.GetConfig("vlan").NDFC.SwitchIP, true, false)
					(*x)["RscName"] = "vlan_intf_ipkey_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfaceVlanModelHelperStateCheck("ndfc_interface_vlan.vlan_intf_ipkey_test", *intfRsc, path.Empty())...),
			},
			// Step 2: Add more SVIs
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateIntfResource(&intfRsc, 205, 5, "vlan", true, helper.GetConfig("vlan").NDFC.SwitchIP, true, true)
					(*x)["RscName"] = "vlan_intf_ipkey_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfaceVlanModelHelperStateCheck("ndfc_interface_vlan.vlan_intf_ipkey_test", *intfRsc, path.Empty())...),
			},
			// Step 3: Modify some SVIs
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.ModifyInterface(&intfRsc, 200, 3, "vlan", map[string]interface{}{
						"admin_state": "down",
						"vrf":         "default",
					})
					(*x)["RscName"] = "vlan_intf_ipkey_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfaceVlanModelHelperStateCheck("ndfc_interface_vlan.vlan_intf_ipkey_test", *intfRsc, path.Empty())...),
			},
			// Step 4: Delete some SVIs
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateIntfResource(&intfRsc, 200, -3, "vlan", true, helper.GetConfig("vlan").NDFC.SwitchIP, true, true)
					(*x)["RscName"] = "vlan_intf_ipkey_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfaceVlanModelHelperStateCheck("ndfc_interface_vlan.vlan_intf_ipkey_test", *intfRsc, path.Empty())...),
			},
		},
	})
}

// ============================================================================
// Loopback Interface CRUD Test with IP Keys (Global Serial)
// ============================================================================

func TestAccInterfaceIPKeyGlobalLoopbackCRUD(t *testing.T) {
	x := &map[string]string{
		"RscType":    "ndfc_interface_loopback",
		"RscSubType": "loopback",
		"RscName":    "test_loopback",
		"User":       helper.GetConfig("loopback").NDFC.User,
		"Password":   helper.GetConfig("loopback").NDFC.Password,
		"Host":       helper.GetConfig("loopback").NDFC.URL,
		"Insecure":   helper.GetConfig("loopback").NDFC.Insecure,
	}

	tf_config := new(string)
	stepCount := new(int)
	*stepCount = 0
	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "loopback") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create loopbacks using IP keys
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// Use SwitchIP instead of Switches for IP-based keys
					helper.GenerateIntfResource(&intfRsc, 30, 5, "loopback", true, helper.GetConfig("loopback").NDFC.SwitchIP, true, false)
					(*x)["RscName"] = "lb_intf_ipkey_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfaceLoopbackModelHelperStateCheck("ndfc_interface_loopback.lb_intf_ipkey_test", *intfRsc, path.Empty())...),
			},
			// Step 2: Add more loopbacks
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateIntfResource(&intfRsc, 35, 5, "loopback", true, helper.GetConfig("loopback").NDFC.SwitchIP, true, true)
					(*x)["RscName"] = "lb_intf_ipkey_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfaceLoopbackModelHelperStateCheck("ndfc_interface_loopback.lb_intf_ipkey_test", *intfRsc, path.Empty())...),
			},
			// Step 3: Modify some loopbacks
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.ModifyInterface(&intfRsc, 30, 3, "loopback", map[string]interface{}{
						"admin_state": "down",
					})
					(*x)["RscName"] = "lb_intf_ipkey_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfaceLoopbackModelHelperStateCheck("ndfc_interface_loopback.lb_intf_ipkey_test", *intfRsc, path.Empty())...),
			},
			// Step 4: Delete some loopbacks
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateIntfResource(&intfRsc, 30, -3, "loopback", true, helper.GetConfig("loopback").NDFC.SwitchIP, true, true)
					(*x)["RscName"] = "lb_intf_ipkey_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfaceLoopbackModelHelperStateCheck("ndfc_interface_loopback.lb_intf_ipkey_test", *intfRsc, path.Empty())...),
			},
		},
	})
}

// ============================================================================
// Interface Level IP Key Tests (globalSerial = false)
// IP is specified at interface level, not global level
// ============================================================================

// TestAccInterfaceIPKeyEthernetCRUD tests ethernet with IP at interface level
func TestAccInterfaceIPKeyEthernetCRUD(t *testing.T) {
	x := &map[string]string{
		"RscType":    "ndfc_interface_ethernet",
		"RscSubType": "ethernet",
		"RscName":    "test_ethernet",
		"User":       helper.GetConfig("ethernet").NDFC.User,
		"Password":   helper.GetConfig("ethernet").NDFC.Password,
		"Host":       helper.GetConfig("ethernet").NDFC.URL,
		"Insecure":   helper.GetConfig("ethernet").NDFC.Insecure,
	}

	tf_config := new(string)
	stepCount := new(int)
	*stepCount = 0
	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "ethernet") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create interfaces with IP at interface level (globalSerial=false)
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// globalSerial=false means IP is at interface level
					helper.GenerateIntfResource(&intfRsc, 10, 5, "ethernet", true, helper.GetConfig("ethernet").NDFC.SwitchIP, false, false)
					(*x)["RscName"] = "eth_intf_ipkey_iflevel_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfaceEthernetModelHelperStateCheck("ndfc_interface_ethernet.eth_intf_ipkey_iflevel_test", *intfRsc, path.Empty())...),
			},
			// Step 2: Add more interfaces
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateIntfResource(&intfRsc, 15, 3, "ethernet", true, helper.GetConfig("ethernet").NDFC.SwitchIP, false, true)
					(*x)["RscName"] = "eth_intf_ipkey_iflevel_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfaceEthernetModelHelperStateCheck("ndfc_interface_ethernet.eth_intf_ipkey_iflevel_test", *intfRsc, path.Empty())...),
			},
			// Step 3: Delete some interfaces
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateIntfResource(&intfRsc, 10, -2, "ethernet", true, helper.GetConfig("ethernet").NDFC.SwitchIP, false, true)
					(*x)["RscName"] = "eth_intf_ipkey_iflevel_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfaceEthernetModelHelperStateCheck("ndfc_interface_ethernet.eth_intf_ipkey_iflevel_test", *intfRsc, path.Empty())...),
			},
		},
	})
}

// TestAccInterfaceIPKeyPortchannelCRUD tests port-channel with IP at interface level
func TestAccInterfaceIPKeyPortchannelCRUD(t *testing.T) {
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
	stepCount := new(int)
	*stepCount = 0
	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t, "portchannel")
			helper.EthIntf = 10
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create port-channels with IP at interface level (globalSerial=false)
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// globalSerial=false means IP is at interface level
					helper.GenerateIntfResource(&intfRsc, 200, 3, "portchannel", true, helper.GetConfig("portchannel").NDFC.SwitchIP, false, false)
					(*x)["RscName"] = "pc_intf_ipkey_iflevel_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfacePortchannelModelHelperStateCheck("ndfc_interface_portchannel.pc_intf_ipkey_iflevel_test", *intfRsc, path.Empty())...),
			},
			// Step 2: Add more port-channels
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateIntfResource(&intfRsc, 203, 2, "portchannel", true, helper.GetConfig("portchannel").NDFC.SwitchIP, false, true)
					(*x)["RscName"] = "pc_intf_ipkey_iflevel_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfacePortchannelModelHelperStateCheck("ndfc_interface_portchannel.pc_intf_ipkey_iflevel_test", *intfRsc, path.Empty())...),
			},
			// Step 3: Delete some port-channels
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateIntfResource(&intfRsc, 200, -2, "portchannel", true, helper.GetConfig("portchannel").NDFC.SwitchIP, false, true)
					(*x)["RscName"] = "pc_intf_ipkey_iflevel_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfacePortchannelModelHelperStateCheck("ndfc_interface_portchannel.pc_intf_ipkey_iflevel_test", *intfRsc, path.Empty())...),
			},
		},
	})
}

// TestAccInterfaceIPKeyVlanCRUD tests SVI with IP at interface level
func TestAccInterfaceIPKeyVlanCRUD(t *testing.T) {
	x := &map[string]string{
		"RscType":    "ndfc_interface_vlan",
		"RscSubType": "vlan",
		"RscName":    "test_vlan",
		"User":       helper.GetConfig("vlan").NDFC.User,
		"Password":   helper.GetConfig("vlan").NDFC.Password,
		"Host":       helper.GetConfig("vlan").NDFC.URL,
		"Insecure":   helper.GetConfig("vlan").NDFC.Insecure,
	}

	tf_config := new(string)
	stepCount := new(int)
	*stepCount = 0
	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "vlan") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create SVIs with IP at interface level (globalSerial=false)
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// globalSerial=false means IP is at interface level
					helper.GenerateIntfResource(&intfRsc, 200, 5, "vlan", true, helper.GetConfig("vlan").NDFC.SwitchIP, false, false)
					(*x)["RscName"] = "vlan_intf_ipkey_iflevel_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfaceVlanModelHelperStateCheck("ndfc_interface_vlan.vlan_intf_ipkey_iflevel_test", *intfRsc, path.Empty())...),
			},
			// Step 2: Add more SVIs
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateIntfResource(&intfRsc, 205, 3, "vlan", true, helper.GetConfig("vlan").NDFC.SwitchIP, false, true)
					(*x)["RscName"] = "vlan_intf_ipkey_iflevel_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfaceVlanModelHelperStateCheck("ndfc_interface_vlan.vlan_intf_ipkey_iflevel_test", *intfRsc, path.Empty())...),
			},
			// Step 3: Delete some SVIs
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateIntfResource(&intfRsc, 200, -2, "vlan", true, helper.GetConfig("vlan").NDFC.SwitchIP, false, true)
					(*x)["RscName"] = "vlan_intf_ipkey_iflevel_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfaceVlanModelHelperStateCheck("ndfc_interface_vlan.vlan_intf_ipkey_iflevel_test", *intfRsc, path.Empty())...),
			},
		},
	})
}

// TestAccInterfaceIPKeyLoopbackCRUD tests loopback with IP at interface level
func TestAccInterfaceIPKeyLoopbackCRUD(t *testing.T) {
	x := &map[string]string{
		"RscType":    "ndfc_interface_loopback",
		"RscSubType": "loopback",
		"RscName":    "test_loopback",
		"User":       helper.GetConfig("loopback").NDFC.User,
		"Password":   helper.GetConfig("loopback").NDFC.Password,
		"Host":       helper.GetConfig("loopback").NDFC.URL,
		"Insecure":   helper.GetConfig("loopback").NDFC.Insecure,
	}

	tf_config := new(string)
	stepCount := new(int)
	*stepCount = 0
	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "loopback") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create loopbacks with IP at interface level (globalSerial=false)
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					// globalSerial=false means IP is at interface level
					helper.GenerateIntfResource(&intfRsc, 30, 5, "loopback", true, helper.GetConfig("loopback").NDFC.SwitchIP, false, false)
					(*x)["RscName"] = "lb_intf_ipkey_iflevel_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfaceLoopbackModelHelperStateCheck("ndfc_interface_loopback.lb_intf_ipkey_iflevel_test", *intfRsc, path.Empty())...),
			},
			// Step 2: Add more loopbacks
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateIntfResource(&intfRsc, 35, 3, "loopback", true, helper.GetConfig("loopback").NDFC.SwitchIP, false, true)
					(*x)["RscName"] = "lb_intf_ipkey_iflevel_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfaceLoopbackModelHelperStateCheck("ndfc_interface_loopback.lb_intf_ipkey_iflevel_test", *intfRsc, path.Empty())...),
			},
			// Step 3: Delete some loopbacks
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateIntfResource(&intfRsc, 30, -2, "loopback", true, helper.GetConfig("loopback").NDFC.SwitchIP, false, true)
					(*x)["RscName"] = "lb_intf_ipkey_iflevel_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfaceLoopbackModelHelperStateCheck("ndfc_interface_loopback.lb_intf_ipkey_iflevel_test", *intfRsc, path.Empty())...),
			},
		},
	})
}
