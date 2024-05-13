package provider

import (
	"fmt"
	"terraform-provider-ndfc/internal/provider/resources/resource_interface_common"
	helper "terraform-provider-ndfc/internal/provider/testing"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccInterfaceEthernetResourceBasic(t *testing.T) {

	x := &map[string]string{
		"RscType":    "ndfc_interface_ethernet",
		"RscSubType": "ethernet",
		"RscName":    "test_ethernet",
		"User":       helper.GetConfig().NDFC.User,
		"Password":   helper.GetConfig().NDFC.Password,
		"Host":       helper.GetConfig().NDFC.URL,
		"Insecure":   helper.GetConfig().NDFC.Insecure,
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
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		//CheckDestroy:             testAccCheckVrfBulkResourceDestroy(networkRsc),
		Steps: []resource.TestStep{
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					//Starting from 10 to consider anything pre-existing on ndfc
					helper.GenerateIntfResource(&intfRsc, 10, 10, "ethernet", true, helper.GetConfig().NDFC.Switches, true)
					(*x)["RscName"] = "eth_intf_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfaceEthernetModelHelperStateCheck("ndfc_interface_ethernet.eth_intf_test", *intfRsc, path.Empty())...),
			},
		},
	})

}
