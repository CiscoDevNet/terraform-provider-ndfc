// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"testing"

	"terraform-provider-ndfc/internal/provider/ndfc"
	"terraform-provider-ndfc/internal/provider/resources/resource_links"
	helper "terraform-provider-ndfc/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

/*
// TestAccLinksIntraFabricNumbered tests the numbered intra-fabric links resource
func TestAccLinksIntraFabricNumbered(t *testing.T) {
	// Initialize test configuration
	x := &map[string]string{
		"RscType":  ndfc.ResourceLinks,
		"RscName":  "intra_numbered_link",
		"User":     helper.GetConfig("links").NDFC.User,
		"Password": helper.GetConfig("links").NDFC.Password,
		"Host":     helper.GetConfig("links").NDFC.URL,
		"Insecure": helper.GetConfig("links").NDFC.Insecure,
	}

	// Create a model for an intra-fabric numbered link
	linksResource := &resource_links.NDFCLinksModel{
		TemplateName:         "int_intra_fabric_num_link",
		SourceFabric:         "CML",
		DestinationFabric:    "CML",
		SourceDevice:         helper.GetConfig("links").NDFC.Switches[0],
		DestinationDevice:    helper.GetConfig("links").NDFC.Switches[1],
		SourceInterface:      "Ethernet1/1",
		DestinationInterface: "Ethernet1/1",
		Deploy:               false,
		LinkParameters: map[string]string{
			"PEER1_IP":    "192.168.1.1",
			"PEER2_IP":    "192.168.1.2",
			"ADMIN_STATE": "true",
			"MTU":         "9216",
			"SPEED":       "Auto",
		},
	}

	// Generate Terraform configuration
	tfConfig := new(string)
	helper.GetTFConfigWithSingleResource(t.Name(), *x, []interface{}{linksResource}, &tfConfig)

	// Run the test
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "links") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: *tfConfig,
				Check:  resource.ComposeTestCheckFunc(LinksModelHelperStateCheck("ndfc_links.intra_numbered_link", *linksResource, path.Empty())...),
			},
		},
	})
}

// TestAccLinksIntraFabricIPv6 tests the IPv6 link-local intra-fabric links resource
func TestAccLinksIntraFabricIPv6(t *testing.T) {
	// Initialize test configuration
	x := &map[string]string{
		"RscType":  ndfc.ResourceLinks,
		"RscName":  "intra_ipv6_link",
		"User":     helper.GetConfig("links").NDFC.User,
		"Password": helper.GetConfig("links").NDFC.Password,
		"Host":     helper.GetConfig("links").NDFC.URL,
		"Insecure": helper.GetConfig("links").NDFC.Insecure,
	}

	// Create a model for an intra-fabric IPv6 link
	linksResource := &resource_links.NDFCLinksModel{
		TemplateName:         "int_intra_fabric_ipv6_link_local",
		SourceFabric:         "CML",
		DestinationFabric:    "CML",
		SourceDevice:         helper.GetConfig("links").NDFC.Switches[0],
		DestinationDevice:    helper.GetConfig("links").NDFC.Switches[1],
		SourceInterface:      "Ethernet1/4",
		DestinationInterface: "Ethernet1/4",
		Deploy:               false,
		LinkParameters: map[string]string{
			"ADMIN_STATE": "true",
			"MTU":         "9216",
			"SPEED":       "Auto",
			"BFD_ENABLE":  "true",
		},
	}

	// Generate Terraform configuration
	tfConfig := new(string)
	helper.GetTFConfigWithSingleResource(t.Name(), *x, []interface{}{linksResource}, &tfConfig)

	// Run the test
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "links") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: *tfConfig,
				Check:  resource.ComposeTestCheckFunc(LinksModelHelperStateCheck("ndfc_links.intra_ipv6_link", *linksResource, path.Empty())...),
			},
			// ImportState testing
			{
				ResourceName:            "ndfc_links.intra_ipv6_link",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"last_updated"},
			},
		},
	})
}

// TestAccLinksIntraFabricUnnumbered tests the unnumbered intra-fabric links resource
func TestAccLinksIntraFabricUnnumbered(t *testing.T) {
	// Initialize test configuration
	x := &map[string]string{
		"RscType":  ndfc.ResourceLinks,
		"RscName":  "intra_unnumbered_link",
		"User":     helper.GetConfig("links").NDFC.User,
		"Password": helper.GetConfig("links").NDFC.Password,
		"Host":     helper.GetConfig("links").NDFC.URL,
		"Insecure": helper.GetConfig("links").NDFC.Insecure,
	}

	// Create a model for an intra-fabric unnumbered link
	linksResource := &resource_links.NDFCLinksModel{
		TemplateName:         "int_intra_fabric_unnum_link",
		SourceFabric:         "CML",
		DestinationFabric:    "CML",
		SourceDevice:         helper.GetConfig("links").NDFC.Switches[0],
		DestinationDevice:    helper.GetConfig("links").NDFC.Switches[1],
		SourceInterface:      "Ethernet1/7",
		DestinationInterface: "Ethernet1/7",
		Deploy:               false,
		LinkParameters: map[string]string{
			"ADMIN_STATE": "true",
			"MTU":         "9216",
			"SPEED":       "Auto",
		},
	}

	// Generate Terraform configuration
	tfConfig := new(string)
	helper.GetTFConfigWithSingleResource(t.Name(), *x, []interface{}{linksResource}, &tfConfig)

	// Run the test
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "links") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: *tfConfig,
				Check:  resource.ComposeTestCheckFunc(LinksModelHelperStateCheck("ndfc_links.intra_unnumbered_link", *linksResource, path.Empty())...),
			},
			// ImportState testing
			{
				ResourceName:            "ndfc_links.intra_unnumbered_link",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"last_updated"},
			},
		},
	})
}

// TestAccLinksIntraFabricVPC tests the VPC peer keep-alive links resource
func TestAccLinksIntraFabricVPC(t *testing.T) {
	// Initialize test configuration
	x := &map[string]string{
		"RscType":  ndfc.ResourceLinks,
		"RscName":  "intra_vpc_link",
		"User":     helper.GetConfig("links").NDFC.User,
		"Password": helper.GetConfig("links").NDFC.Password,
		"Host":     helper.GetConfig("links").NDFC.URL,
		"Insecure": helper.GetConfig("links").NDFC.Insecure,
	}

	// Create a model for a VPC peer keep-alive link
	linksResource := &resource_links.NDFCLinksModel{
		TemplateName:         "int_intra_vpc_peer_keep_alive_link",
		SourceFabric:         "CML",
		DestinationFabric:    "CML",
		SourceDevice:         helper.GetConfig("links").NDFC.Switches[0],
		DestinationDevice:    helper.GetConfig("links").NDFC.Switches[1],
		SourceInterface:      "Ethernet1/10",
		DestinationInterface: "Ethernet1/10",
		Deploy:               false,
		LinkParameters: map[string]string{
			"PEER1_IP":    "192.168.10.1",
			"PEER2_IP":    "192.168.10.2",
			"ADMIN_STATE": "true",
			"MTU":         "9216",
			"SPEED":       "Auto",
		},
	}

	// Generate Terraform configuration
	tfConfig := new(string)
	helper.GetTFConfigWithSingleResource(t.Name(), *x, []interface{}{linksResource}, &tfConfig)

	// Run the test
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "links") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: *tfConfig,
				Check:  resource.ComposeTestCheckFunc(LinksModelHelperStateCheck("ndfc_links.intra_vpc_link", *linksResource, path.Empty())...),
			},
			// ImportState testing
			{
				ResourceName:            "ndfc_links.intra_vpc_link",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"last_updated"},
			},
		},
	})
}

// TestAccLinksInterFabric tests the inter-fabric links resource
func TestAccLinksInterFabric(t *testing.T) {
	// Initialize test configuration
	x := &map[string]string{
		"RscType":  ndfc.ResourceLinks,
		"RscName":  "inter_fabric_link",
		"User":     helper.GetConfig("links").NDFC.User,
		"Password": helper.GetConfig("links").NDFC.Password,
		"Host":     helper.GetConfig("links").NDFC.URL,
		"Insecure": helper.GetConfig("links").NDFC.Insecure,
	}

	// Create a model for an inter-fabric numbered link
	linksResource := &resource_links.NDFCLinksModel{
		TemplateName:         "int_inter_fabric_num_link",
		SourceFabric:         "CML1",
		DestinationFabric:    "CML2",
		SourceDevice:         helper.GetConfig("links").NDFC.Switches[0],
		DestinationDevice:    helper.GetConfig("links").NDFC.Switches[1],
		SourceInterface:      "Ethernet1/14",
		DestinationInterface: "Ethernet1/14",
		Deploy:               false,
		LinkParameters: map[string]string{
			"PEER1_IP":    "192.168.14.1",
			"PEER2_IP":    "192.168.14.2",
			"ADMIN_STATE": "true",
			"MTU":         "9216",
			"SPEED":       "Auto",
		},
	}

	// Generate Terraform configuration
	tfConfig := new(string)
	helper.GetTFConfigWithSingleResource(t.Name(), *x, []interface{}{linksResource}, &tfConfig)

	// Run the test
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "links") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: *tfConfig,
				Check:  resource.ComposeTestCheckFunc(LinksModelHelperStateCheck("ndfc_links.inter_fabric_link", *linksResource, path.Empty())...),
			},
			// ImportState testing
			{
				ResourceName:            "ndfc_links.inter_fabric_link",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"last_updated"},
			},
		},
	})
}
*/

// TestAccLinksIntraFabricNumLink tests the intra-fabric numbered link configuration from sample test file
func TestAccLinksIntraFabricNumLink(t *testing.T) {
	// Initialize test configuration
	cfg := helper.GetConfig("links")
	if cfg.NDFC.Link.SrcFabric == "" {
		t.Skip("Skipping test as no link configuration is available")
	}
	x := &map[string]string{
		"RscType":  ndfc.ResourceLinks,
		"RscName":  "intra_fabric_numbered",
		"User":     helper.GetConfig("links").NDFC.User,
		"Password": helper.GetConfig("links").NDFC.Password,
		"Host":     helper.GetConfig("links").NDFC.URL,
		"Insecure": helper.GetConfig("links").NDFC.Insecure,
	}

	// Create a model for an intra-fabric numbered link matching the sample file
	linksResource := &resource_links.NDFCLinksModel{
		TemplateName:         "int_intra_fabric_num_link",
		SourceFabric:         helper.GetConfig("links").NDFC.Link.SrcFabric,
		DestinationFabric:    helper.GetConfig("links").NDFC.Link.SrcFabric,
		SourceDevice:         helper.GetConfig("links").NDFC.Switches[0],
		DestinationDevice:    helper.GetConfig("links").NDFC.Switches[1],
		SourceInterface:      "Ethernet1/1",
		DestinationInterface: "Ethernet1/1",
		LinkParameters: map[string]string{
			"PEER1_IP":                "40.1.1.1",
			"PEER2_IP":                "40.1.1.2",
			"PEER1_V6IP":              "",
			"PEER2_V6IP":              "",
			"ADMIN_STATE":             "true",
			"MTU":                     "9216",
			"SPEED":                   "Auto",
			"PEER1_BFD_ECHO_DISABLE":  "false",
			"PEER2_BFD_ECHO_DISABLE":  "false",
			"ENABLE_MACSEC":           "false",
			"PEER1_CONF":              "",
			"PEER2_CONF":              "",
			"PEER1_DESC":              "1",
			"PEER2_DESC":              "2",
			"ENABLE_PEER1_DHCP_RELAY": "false",
			"ENABLE_PEER2_DHCP_RELAY": "false",
		},
	}

	// Generate Terraform configuration
	tfConfig := new(string)
	helper.GetTFConfigWithSingleResource(t.Name(), *x, []interface{}{linksResource}, &tfConfig)

	// Run the test
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "links") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: *tfConfig,
				Check:  resource.ComposeTestCheckFunc(LinksModelHelperStateCheck("ndfc_links.intra_fabric_numbered", *linksResource, path.Empty())...),
			},
			/*
				// ImportState testing
				{
					ResourceName:            "ndfc_links.intra_fabric_numbered",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"last_updated"},
				},
			*/
		},
	})
}

// TestAccLinksInterFabricMultiSiteUnderlay tests the inter-fabric external multisite underlay setup
func TestAccLinksExtFabricSetup(t *testing.T) {
	cfg := helper.GetConfig("links")
	if cfg.NDFC.Link.SrcFabric == "" {
		t.Skip("Skipping test as no link configuration is available")
	}
	// Initialize test configuration
	x := &map[string]string{
		"RscType":  ndfc.ResourceLinks,
		"RscName":  "ext_fabric_ospf_link",
		"User":     helper.GetConfig("links").NDFC.User,
		"Password": helper.GetConfig("links").NDFC.Password,
		"Host":     helper.GetConfig("links").NDFC.URL,
		"Insecure": helper.GetConfig("links").NDFC.Insecure,
	}

	// Create a model for an OSPF external fabric setup link as in main.tf
	linksResource := &resource_links.NDFCLinksModel{
		TemplateName:         "ext_fabric_setup",
		SourceFabric:         cfg.NDFC.Link.SrcFabric,
		DestinationFabric:    cfg.NDFC.Link.DstFabric,
		SourceDevice:         cfg.NDFC.Link.SrcDevice,
		DestinationDevice:    cfg.NDFC.Link.DstDevice,
		SourceInterface:      "Ethernet1/2",
		DestinationInterface: "Ethernet1/2",
		LinkParameters: map[string]string{
			// Basic Parameters
			"asn":           "29500",
			"NEIGHBOR_ASN":  "36501",
			"IP_MASK":       "10.23.2.2/24",
			"NEIGHBOR_IP":   "10.23.2.4",
			"IPV6_MASK":     "",
			"IPV6_NEIGHBOR": "",

			// Interface Parameters
			"MTU":      "9216",
			"PRIORITY": "500",

			// VRF Parameters
			"AUTO_VRF_LITE_FLAG":            "false",
			"VRF_LITE_JYTHON_TEMPLATE":      "Ext_VRF_Lite_Jython",
			"DEFAULT_VRF_FLAG":              "false",
			"SYMMETRIC_DEFAULT_VRF_FLAG":    "false",
			"DEFAULT_VRF_REDIS_BGP_RMAP":    "",
			"DEFAULT_VRF_BGP_PASSWORD":      "",
			"DEFAULT_VRF_BGP_AUTH_KEY_TYPE": "",
			"DEFAULT_VRF_PEER_VRF_NAME":     "",

			// Advanced Parameters
			"ENABLE_DCI_TRACKING": "false",
			"ROUTING_TAG":         "",

			// Peer Descriptions
			"PEER1_DESC": "This is switch1",
			"PEER2_DESC": "This is switch2",
			"PEER1_CONF": "",
			"PEER2_CONF": "",
		},
	}

	// Generate Terraform configuration
	tfConfig := new(string)
	helper.GetTFConfigWithSingleResource(t.Name(), *x, []interface{}{linksResource}, &tfConfig)

	// Run the test
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "links") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: *tfConfig,
				Check:  resource.ComposeTestCheckFunc(LinksModelHelperStateCheck("ndfc_links.ext_fabric_ospf_link", *linksResource, path.Empty())...),
			},
			/*
				// ImportState testing
				{
					ResourceName:            "ndfc_links.ext_fabric_ospf_link",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"last_updated"},
				},
			*/
		},
	})
}

func TestAccLinksInterFabricMultiSiteUnderlay(t *testing.T) {
	cfg := helper.GetConfig("links")
	if cfg.NDFC.Link.SrcFabric == "" {
		t.Skip("Skipping test as no link configuration is available")
	}
	// Initialize test configuration
	x := &map[string]string{
		"RscType":  ndfc.ResourceLinks,
		"RscName":  "ext_multisite_underlay_link",
		"User":     helper.GetConfig("links").NDFC.User,
		"Password": helper.GetConfig("links").NDFC.Password,
		"Host":     helper.GetConfig("links").NDFC.URL,
		"Insecure": helper.GetConfig("links").NDFC.Insecure,
	}

	// Create a model for an inter-fabric external multisite underlay link
	linksResource := &resource_links.NDFCLinksModel{
		TemplateName:         "ext_multisite_underlay_setup",
		SourceFabric:         cfg.NDFC.Link.SrcFabric,
		DestinationFabric:    cfg.NDFC.Link.DstFabric,
		SourceDevice:         cfg.NDFC.Link.SrcDevice,
		DestinationDevice:    cfg.NDFC.Link.DstDevice,
		SourceInterface:      "Ethernet1/2",
		DestinationInterface: "Ethernet1/2",
		LinkParameters: map[string]string{
			"asn":                            "29500",
			"IP_MASK":                        "192.168.38.1/24",
			"NEIGHBOR_IP":                    "192.168.38.2",
			"NEIGHBOR_ASN":                   "36501",
			"MAX_PATHS":                      "1",
			"ROUTING_TAG":                    "54321",
			"MTU":                            "9216",
			"PEER1_DESC":                     "test1",
			"PEER2_DESC":                     "test2",
			"PEER1_CONF":                     "no shutdown",
			"PEER2_CONF":                     "no shutdown",
			"DEPLOY_DCI_TRACKING":            "false",
			"BGP_PASSWORD_ENABLE":            "true",
			"ENABLE_BGP_LOG_NEIGHBOR_CHANGE": "false",
			"ENABLE_BGP_BFD":                 "false",
			"ENABLE_BGP_SEND_COMM":           "false",
			"BGP_PASSWORD_INHERIT_FROM_MSD":  "true",
			"BGP_AUTH_KEY_TYPE":              "",
		},
	}

	// Generate Terraform configuration
	tfConfig := new(string)
	helper.GetTFConfigWithSingleResource(t.Name(), *x, []interface{}{linksResource}, &tfConfig)

	// Run the test
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "links") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: *tfConfig,
				Check:  resource.ComposeTestCheckFunc(LinksModelHelperStateCheck("ndfc_links.ext_multisite_underlay_link", *linksResource, path.Empty())...),
			},
			/*
				// ImportState testing
				{
					ResourceName:            "ndfc_links.ext_multisite_underlay_link",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"last_updated"},
				},
			*/
		},
	})
}
