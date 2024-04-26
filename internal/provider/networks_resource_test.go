package provider

import (
	"fmt"
	"regexp"
	"testing"

	//"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/hashicorp/terraform-plugin-framework/path"

	"terraform-provider-ndfc/internal/provider/ndfc"
	"terraform-provider-ndfc/internal/provider/resources/resource_networks"
	"terraform-provider-ndfc/internal/provider/resources/resource_vrf_bulk"
	"terraform-provider-ndfc/internal/provider/types"

	helper "terraform-provider-ndfc/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

/*
When a test is ran,
Terraform runs plan, apply, refresh, and then final plan for each TestStep in the TestCase.
*/

func TestAccNDFCNetworksResourceCRUD(t *testing.T) {

	x := &map[string]string{
		"RscType":  ndfc.ResourceNetworks,
		"RscName":  "network_test",
		"User":     helper.GetConfig().NDFC.User,
		"Password": helper.GetConfig().NDFC.Password,
		"Host":     helper.GetConfig().NDFC.URL,
		"Insecure": helper.GetConfig().NDFC.Insecure,
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

	networkRsc := new(resource_networks.NDFCNetworksModel)
	vrfRsc := new(resource_vrf_bulk.NDFCVrfBulkModel)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		//CheckDestroy:             testAccCheckVrfBulkResourceDestroy(networkRsc),
		Steps: []resource.TestStep{
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateSingleVrfObject(&vrfRsc, helper.GetConfig().NDFC.VrfPrefix, helper.GetConfig().NDFC.Fabric, 1, false, false, false, helper.GetConfig().NDFC.Switches)
					helper.GenerateNetworksObject(&networkRsc, helper.GetConfig().NDFC.Fabric,
						10, false, false, false, helper.GetConfig().NDFC.VrfPrefix+"1", nil)
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			{
				//Add 10 more Networks
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.IncreaseNetCount(&networkRsc,
						10, false, false, false, helper.GetConfig().NDFC.VrfPrefix+"1", nil)
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.DeleteNetworks(&networkRsc,
						11, 20)
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			{

				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					netNewRsc := new(resource_networks.NDFCNetworksModel)
					helper.GenerateNetworksObject(&netNewRsc, helper.GetConfig().NDFC.Fabric,
						10, false, false, false, helper.GetConfig().NDFC.VrfPrefix+"1", nil)

					(*x)["RscName"] = "vrf_test,network_test,network_test1"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc, netNewRsc}, &tf_config)
					return *tf_config
				}(),
				ExpectError: regexp.MustCompile(".*Networks exist.*"),
			},

			{
				//Modify Few Params in Nets
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					tf_config := new(string)
					helper.ModifyNetworksObject(&networkRsc, 1, map[string]interface{}{
						"vlan_id":                1220,
						"multicast_group":        "239.0.0.100",
						"dhcp_relay_loopback_id": 10,
						"arp_suppression":        "true",
					})
					helper.ModifyNetworksObject(&networkRsc, 10, map[string]interface{}{
						"vlan_id":                1102,
						"multicast_group":        "239.0.0.200",
						"dhcp_relay_loopback_id": 12,
						"arp_suppression":        "true",
						"secondary_gateway_1":    "192.168.100.1/24",
						"secondary_gateway_2":    "192.168.101.1/24",
					})
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},
		}})
}

func TestAccNDFCNetworkAttachmentCRUD(t *testing.T) {

	x := &map[string]string{
		"RscType":  ndfc.ResourceVrfBulk,
		"RscName":  "vrf_test,network_test",
		"User":     helper.GetConfig().NDFC.User,
		"Password": helper.GetConfig().NDFC.Password,
		"Host":     helper.GetConfig().NDFC.URL,
		"Insecure": helper.GetConfig().NDFC.Insecure,
	}
	stepCount := new(int)
	*stepCount = 0
	networkRsc := new(resource_networks.NDFCNetworksModel)
	vrfRsc := new(resource_vrf_bulk.NDFCVrfBulkModel)

	resource.Test(t, resource.TestCase{

		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			//Create VRFs with 2 attachments
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					tf_config := new(string)
					helper.GenerateSingleVrfObject(&vrfRsc, helper.GetConfig().NDFC.VrfPrefix, helper.GetConfig().NDFC.Fabric, 1, false, false, false, helper.GetConfig().NDFC.Switches)
					helper.GenerateNetworksObject(&networkRsc, helper.GetConfig().NDFC.Fabric,
						10, false, false, false, helper.GetConfig().NDFC.VrfPrefix+"1", []string{helper.GetConfig().NDFC.Switches[0], helper.GetConfig().NDFC.Switches[1]})
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			{ // Remove both attachments _detach
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					tf_config := new(string)
					//helper.VrfAttachmentsMod(&vrfRsc, 1, 1, nil, "", nil)
					helper.NetAttachmentsMod(&networkRsc, 1, len(networkRsc.Networks), nil, "", nil)
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				// Detach also detaches the entry in VRF in NDFC
				// This causes terraform to report VRF attachment as missing and need re-Add
				// This is a known limitation, due to the way the NDFC  works
				ExpectNonEmptyPlan: true,

				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},
			{
				// Add 2 attachments to all VRFs
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					tf_config := new(string)
					helper.NetAttachmentsMod(&networkRsc, 1, len(networkRsc.Networks), []string{helper.GetConfig().NDFC.Switches[0], helper.GetConfig().NDFC.Switches[1]}, "", nil)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			{
				// Add 3rd attachment half of them
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					tf_config := new(string)
					helper.NetAttachmentsMod(&networkRsc, 1, len(networkRsc.Networks)/2, helper.GetConfig().NDFC.Switches, "", nil)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			{
				// Add 3rd attachment remaining  half
				// Remove 3rd from others
				Config: func() string {
					tf_config := new(string)
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.NetAttachmentsMod(&networkRsc, 1, len(networkRsc.Networks)/2, []string{helper.GetConfig().NDFC.Switches[0], helper.GetConfig().NDFC.Switches[1]}, "", nil)
					helper.NetAttachmentsMod(&networkRsc, (len(networkRsc.Networks)/2)+1, len(networkRsc.Networks)/2, helper.GetConfig().NDFC.Switches, "", nil)
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				// Detach and attach also does the same in VRF entry in NDFC
				// This causes terraform to report VRF attachment as missing/added in plan
				// This is a known limitation, due to the way the NDFC works
				ExpectNonEmptyPlan: true,
				Check:              resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},

			{
				//Modify params
				Config: func() string {
					tf_config := new(string)
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.NetAttachmentsMod(&networkRsc, 1, 1, helper.GetConfig().NDFC.Switches, helper.GetConfig().NDFC.Switches[2], map[string]interface{}{
						"switch_ports": types.CSVString{"Ethernet1/1", "Ethernet1/2"},
					})
					helper.NetAttachmentsMod(&networkRsc, 10, 10, helper.GetConfig().NDFC.Switches, helper.GetConfig().NDFC.Switches[2], map[string]interface{}{
						"switch_ports": types.CSVString{"Ethernet1/1", "Ethernet1/2"},
					})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},
			{
				//Modify port list, remove one, add another
				Config: func() string {
					tf_config := new(string)
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.NetAttachmentsMod(&networkRsc, 1, 1, helper.GetConfig().NDFC.Switches, helper.GetConfig().NDFC.Switches[2], map[string]interface{}{
						"switch_ports": types.CSVString{"Ethernet1/2"},
					})
					helper.NetAttachmentsMod(&networkRsc, 10, 10, helper.GetConfig().NDFC.Switches, helper.GetConfig().NDFC.Switches[2], map[string]interface{}{
						"switch_ports": types.CSVString{"Ethernet1/2", "Ethernet1/3"},
					})
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},
		}})
}

func TestAccNDFCNetworksGlobalDeploy(t *testing.T) {
	x := &map[string]string{
		"RscType":  ndfc.ResourceNetworks,
		"RscName":  "network_test",
		"User":     helper.GetConfig().NDFC.User,
		"Password": helper.GetConfig().NDFC.Password,
		"Host":     helper.GetConfig().NDFC.URL,
		"Insecure": helper.GetConfig().NDFC.Insecure,
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

	networkRsc := new(resource_networks.NDFCNetworksModel)
	vrfRsc := new(resource_vrf_bulk.NDFCVrfBulkModel)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				//GlobalDeploy
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateSingleVrfObject(&vrfRsc, helper.GetConfig().NDFC.VrfPrefix, helper.GetConfig().NDFC.Fabric, 1, false, false, false, helper.GetConfig().NDFC.Switches)
					helper.GenerateNetworksObject(&networkRsc, helper.GetConfig().NDFC.Fabric,
						10, true, false, false, helper.GetConfig().NDFC.VrfPrefix+"1", helper.GetConfig().NDFC.Switches)
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				// Network Deploy also deploys VRF, so expect some change in plan
				ExpectNonEmptyPlan: true,
				Check:              resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},
		}})
}

func TestAccNetworkAttachmentDeployNetLevel(t *testing.T) {
	x := &map[string]string{
		"RscType":  ndfc.ResourceNetworks,
		"RscName":  "network_test",
		"User":     helper.GetConfig().NDFC.User,
		"Password": helper.GetConfig().NDFC.Password,
		"Host":     helper.GetConfig().NDFC.URL,
		"Insecure": helper.GetConfig().NDFC.Insecure,
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

	networkRsc := new(resource_networks.NDFCNetworksModel)
	vrfRsc := new(resource_vrf_bulk.NDFCVrfBulkModel)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				//Deploy Net Level
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateSingleVrfObject(&vrfRsc, helper.GetConfig().NDFC.VrfPrefix, helper.GetConfig().NDFC.Fabric, 1, false, false, false, helper.GetConfig().NDFC.Switches)
					helper.GenerateNetworksObject(&networkRsc, helper.GetConfig().NDFC.Fabric,
						10, false, true, false, helper.GetConfig().NDFC.VrfPrefix+"1", helper.GetConfig().NDFC.Switches)
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				// Network Deploy also deploys VRF, so expect some change in plan
				ExpectNonEmptyPlan: true,
				Check:              resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},
		}})
}

func TestAccNetworkAttachmentDeployAttachments(t *testing.T) {
	x := &map[string]string{
		"RscType":  ndfc.ResourceNetworks,
		"RscName":  "network_test",
		"User":     helper.GetConfig().NDFC.User,
		"Password": helper.GetConfig().NDFC.Password,
		"Host":     helper.GetConfig().NDFC.URL,
		"Insecure": helper.GetConfig().NDFC.Insecure,
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

	networkRsc := new(resource_networks.NDFCNetworksModel)
	vrfRsc := new(resource_vrf_bulk.NDFCVrfBulkModel)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				//Deploy Attachments
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateSingleVrfObject(&vrfRsc, helper.GetConfig().NDFC.VrfPrefix, helper.GetConfig().NDFC.Fabric, 1, false, false, false, helper.GetConfig().NDFC.Switches)
					helper.GenerateNetworksObject(&networkRsc, helper.GetConfig().NDFC.Fabric,
						10, false, false, true, helper.GetConfig().NDFC.VrfPrefix+"1", helper.GetConfig().NDFC.Switches)
					(*x)["RscName"] = "vrf_test,network_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{vrfRsc, networkRsc}, &tf_config)
					return *tf_config
				}(),
				// Network Deploy also deploys VRF, so expect some change in plan
				ExpectNonEmptyPlan: true,
				Check:              resource.ComposeTestCheckFunc(NetworksModelHelperStateCheck("ndfc_networks.network_test", *networkRsc, path.Empty())...),
			},
		}})
}
