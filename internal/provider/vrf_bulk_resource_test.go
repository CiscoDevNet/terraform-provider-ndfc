package provider

import (
	"fmt"
	"regexp"
	"testing"

	//"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/hashicorp/terraform-plugin-framework/path"

	"terraform-provider-ndfc/internal/provider/ndfc"
	"terraform-provider-ndfc/internal/provider/resources/resource_vrf_bulk"

	helper "terraform-provider-ndfc/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

/*
When a test is ran,
Terraform runs plan, apply, refresh, and then final plan for each TestStep in the TestCase.
*/
func TestAccVrfBulkGenerate(t *testing.T) {
	x := &map[string]string{
		"RscType":  ndfc.ResourceVrfBulk,
		"RscName":  "vrf_test",
		"User":     "admin",
		"Password": "admin!@#",
		"Host":     "https://10.78.210.161",
		"Insecure": "true",
	}
	vrfScaledBulk := new(resource_vrf_bulk.NDFCVrfBulkModel)
	tf_config := new(string)
	helper.GenerateVrfBulkObject(&vrfScaledBulk, "test_evpn_vxlan",
		50, false, false, true, []string{"9FE076D8EJL", "9TQYTJSZ1VJ", "9QBCTIN0FMY"})
	helper.GetTFConfigWithSingleResource(t.Name(), *x, *vrfScaledBulk, &tf_config)
}

func TestAccVrfGenerate(t *testing.T) {
	testGenerateVrfMultipleResource(50, "vrf_scale", "vrf_test")
}

func TestAccNDFCVrfBulkResourceCRUD(t *testing.T) {

	x := &map[string]string{
		"RscType":  ndfc.ResourceVrfBulk,
		"RscName":  "vrf_test",
		"User":     "admin",
		"Password": "admin!@#",
		"Host":     "https://10.78.210.161",
		"Insecure": "true",
	}

	tf_config := new(string)
	*tf_config = `provider "ndfc" {
		host     = "https://"
		username = "admin"
		password = "admin!@#"
		domain   = "example.com"
		insecure = true
		}
		resource ndfc_vrf_bulk "vrf_test" {
			fabric_name = "dummy"
		}`

	stepCount := new(int)
	*stepCount = 0
	// Create a new instance of the NDFC client
	vrfBulk := new(resource_vrf_bulk.VrfBulkModel)
	vrfScaledBulk := new(resource_vrf_bulk.NDFCVrfBulkModel)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVrfBulkResourceDestroy(vrfBulk),
		Steps: []resource.TestStep{
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.GenerateVrfBulkObject(&vrfScaledBulk, "test_evpn_vxlan",
						10, false, false, false, nil)
					helper.GetTFConfigWithSingleResource(tName, *x, *vrfScaledBulk, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck("ndfc_vrf_bulk.vrf_test", *vrfScaledBulk, path.Empty())...),
			},
			{
				//Add 10 more VRFs
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.IncreaseVrfCount(&vrfScaledBulk,
						10, false, false, false, nil)
					helper.GetTFConfigWithSingleResource(tName, *x, *vrfScaledBulk, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck("ndfc_vrf_bulk.vrf_test", *vrfScaledBulk, path.Empty())...),
			},
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.DeleteVrfs(&vrfScaledBulk,
						11, 20)
					helper.GetTFConfigWithSingleResource(tName, *x, *vrfScaledBulk, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck("ndfc_vrf_bulk.vrf_test", *vrfScaledBulk, path.Empty())...),
			},
			{

				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					(*x)["RscName"] = "vrf_test_1"
					helper.GetTFConfigWithSingleResource(tName, *x, *vrfScaledBulk, &tf_config)
					return *tf_config
				}(),
				ExpectError: regexp.MustCompile(".*VRFs exist.*"),
			},
			{
				//Modify Few Params in VRFs
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					tf_config := new(string)
					helper.ModifyVrfBulkObject(&vrfScaledBulk, 1, map[string]interface{}{
						"vlan_id":              100,
						"vrf_description":      "test",
						"loopback_routing_tag": 2459,
						"mtu":                  9100,
						"max_bgp_paths":        2,
						"ipv6_link_local":      "fe80::9afb",
					})
					helper.ModifyVrfBulkObject(&vrfScaledBulk, 10, map[string]interface{}{
						"vlan_id":              110,
						"vrf_description":      "test",
						"loopback_routing_tag": 2459,
						"mtu":                  9100,
						"max_bgp_paths":        2,
						"ipv6_link_local":      "fe80::9afb",
					})

					helper.GetTFConfigWithSingleResource(tName, *x, *vrfScaledBulk, &tf_config)
					return *tf_config
				}(),
			},
		}})
}

func TestAccNDFCVrfBulkResourceAttachmentCRUD(t *testing.T) {

	x := &map[string]string{
		"RscType":  ndfc.ResourceVrfBulk,
		"RscName":  "vrf_test",
		"User":     "admin",
		"Password": "admin!@#",
		"Host":     "https://10.78.210.161",
		"Insecure": "true",
	}
	vrfScaledBulk := new(resource_vrf_bulk.NDFCVrfBulkModel)
	stepCount := new(int)
	*stepCount = 0

	resource.Test(t, resource.TestCase{

		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			//Create VRFs with 2 attachments _Attach
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					tf_config := new(string)
					helper.GenerateVrfBulkObject(&vrfScaledBulk, "test_evpn_vxlan",
						20, false, false, false, []string{"9FE076D8EJL", "9TQYTJSZ1VJ"})
					helper.GetTFConfigWithSingleResource(tName, *x, *vrfScaledBulk, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck("ndfc_vrf_bulk.vrf_test", *vrfScaledBulk, path.Empty())...),
			},
			{ // Remove both attachments _detach
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					tf_config := new(string)
					helper.VrfAttachmentsMod(&vrfScaledBulk, 1, len(vrfScaledBulk.Vrfs), nil, "", nil)
					helper.GetTFConfigWithSingleResource(tName, *x, *vrfScaledBulk, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck("ndfc_vrf_bulk.vrf_test", *vrfScaledBulk, path.Empty())...),
			},
			{
				// Add 2 attachments to all VRFs
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					tf_config := new(string)
					helper.VrfAttachmentsMod(&vrfScaledBulk, 1, len(vrfScaledBulk.Vrfs), []string{"9FE076D8EJL", "9TQYTJSZ1VJ"}, "", nil)
					helper.GetTFConfigWithSingleResource(tName, *x, *vrfScaledBulk, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck("ndfc_vrf_bulk.vrf_test", *vrfScaledBulk, path.Empty())...),
			},
			{
				// Add 3rd attachment half of them
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					tf_config := new(string)
					helper.VrfAttachmentsMod(&vrfScaledBulk, 1, len(vrfScaledBulk.Vrfs)/2, []string{"9FE076D8EJL", "9TQYTJSZ1VJ", "9QBCTIN0FMY"}, "", nil)
					helper.GetTFConfigWithSingleResource(tName, *x, *vrfScaledBulk, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck("ndfc_vrf_bulk.vrf_test", *vrfScaledBulk, path.Empty())...),
			},
			{
				// Add 3rd attachment remaining  half
				// Remove 3rd from others
				Config: func() string {
					tf_config := new(string)
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.VrfAttachmentsMod(&vrfScaledBulk, 1, len(vrfScaledBulk.Vrfs)/2, []string{"9FE076D8EJL", "9TQYTJSZ1VJ"}, "", nil)
					helper.VrfAttachmentsMod(&vrfScaledBulk, (len(vrfScaledBulk.Vrfs)/2)+1, len(vrfScaledBulk.Vrfs)/2, []string{"9FE076D8EJL", "9TQYTJSZ1VJ", "9QBCTIN0FMY"}, "", nil)
					helper.GetTFConfigWithSingleResource(tName, *x, *vrfScaledBulk, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck("ndfc_vrf_bulk.vrf_test", *vrfScaledBulk, path.Empty())...),
			},
			{
				//Modify params
				Config: func() string {
					tf_config := new(string)
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.VrfAttachmentsMod(&vrfScaledBulk, 1, 1, []string{"9FE076D8EJL", "9TQYTJSZ1VJ", "9QBCTIN0FMY"}, "9QBCTIN0FMY", map[string]interface{}{
						"vlan":          1001,
						"loopback_id":   1001,
						"loopback_ipv4": "10.1.1.1",
						"loopback_ipv6": "2001:db8::68",
					})

					helper.VrfAttachmentsMod(&vrfScaledBulk, 10, 10, []string{"9FE076D8EJL", "9TQYTJSZ1VJ", "9QBCTIN0FMY"}, "9QBCTIN0FMY", map[string]interface{}{
						"vlan":          1010,
						"loopback_id":   1010,
						"loopback_ipv4": "10.1.1.10",
						"loopback_ipv6": "2001:db8::610",
					})
					helper.GetTFConfigWithSingleResource(tName, *x, *vrfScaledBulk, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(VrfBulkModelHelperStateCheck("ndfc_vrf_bulk.vrf_test", *vrfScaledBulk, path.Empty())...),
			},
		}})
}

func TestAccNDFCMultiResourceWithDeploy(t *testing.T) {

	x := &map[string]string{
		"RscType":  ndfc.ResourceVrfBulk,
		"RscName":  "",
		"User":     "admin",
		"Password": "admin!@#",
		"Host":     "https://10.78.210.161",
		"Insecure": "true",
	}
	var vrfScaledBulk []*resource_vrf_bulk.NDFCVrfBulkModel
	stepCount := new(int)
	*stepCount = 0

	resource.Test(t, resource.TestCase{

		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			//Create VRFs with 2 attachments _Attach
			{
				PreConfig: func() {
					t.Logf("Starting Test Step %s_%d", t.Name(), *stepCount)
				},
				Config: func() string {
					vrfScaledBulk = make([]*resource_vrf_bulk.NDFCVrfBulkModel, 5)
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					tf_config := new(string)
					for i := 0; i < 5; i++ {
						vrfScaledBulk[i] = new(resource_vrf_bulk.NDFCVrfBulkModel)
						helper.GenerateSingleVrfObject(&(vrfScaledBulk[i]), "vrf_acc_", "test_evpn_vxlan",
							i+1, false, false, true, []string{"9FE076D8EJL", "9TQYTJSZ1VJ"})
						if i == 0 {
							(*x)["RscName"] = fmt.Sprintf("vrf_test_%d", i+1)
						} else {
							(*x)["RscName"] = fmt.Sprintf("%s,vrf_test_%d", (*x)["RscName"], i+1)
						}
					}
					helper.GetTFConfigWithMultipleResource(tName, *x, &vrfScaledBulk, &tf_config)
					return *tf_config
				}(),
				Check: func() resource.TestCheckFunc {
					var checks []resource.TestCheckFunc
					for i := 0; i < len(vrfScaledBulk); i++ {
						checks = append(checks, VrfBulkModelHelperStateCheck(fmt.Sprintf("ndfc_vrf_bulk.vrf_test_%d", i+1), *vrfScaledBulk[i], path.Empty())...)
					}
					return resource.ComposeTestCheckFunc(checks...)
				}(),
			},
		}})
}

func testAccCheckVrfBulkResourceDestroy(vrfBulk *resource_vrf_bulk.VrfBulkModel) resource.TestCheckFunc {
	return nil
}

func testGenerateVrfMultipleResource(count int, vrfName string, rscName string) string {
	x := &map[string]string{
		"RscType":  ndfc.ResourceVrfBulk,
		"RscName":  "vrf_test",
		"User":     "admin",
		"Password": "admin!@#",
		"Host":     "https://10.78.210.161",
		"Insecure": "true",
	}
	vrfScaledBulk := make([]*resource_vrf_bulk.NDFCVrfBulkModel, count)
	tf_config := new(string)
	for i := 0; i < count; i++ {
		vrfScaledBulk[i] = new(resource_vrf_bulk.NDFCVrfBulkModel)
		helper.GenerateSingleVrfObject(&(vrfScaledBulk[i]), vrfName, "test_evpn_vxlan",
			i+1, false, false, true, []string{"9FE076D8EJL", "9TQYTJSZ1VJ"})
		if (*x)["RscName"] == "" {
			(*x)["RscName"] = fmt.Sprintf("%s_%d", rscName, i+1)
		} else {
			(*x)["RscName"] = (*x)["RscName"] + "," + fmt.Sprintf("vrf_test_%d", i+1)
		}
	}
	helper.GetTFConfigWithMultipleResource("multiple_rsc_", *x, &vrfScaledBulk, &tf_config)
	return *tf_config
}
