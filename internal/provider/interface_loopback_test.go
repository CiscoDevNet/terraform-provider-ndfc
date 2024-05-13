package provider

import (
	"fmt"
	"terraform-provider-ndfc/internal/provider/resources/resource_interface_common"
	helper "terraform-provider-ndfc/internal/provider/testing"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccInterfaceLoopbackResourceBasic(t *testing.T) {

	x := &map[string]string{
		"RscType":    "ndfc_interface_loopback",
		"RscSubType": "loopback",
		"RscName":    "test_loopback",
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
					helper.GenerateIntfResource(&intfRsc, 30, 10, "loopback", true, helper.GetConfig().NDFC.Switches, true)
					(*x)["RscName"] = "lb_intf_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfaceLoopbackModelHelperStateCheck("ndfc_interface_loopback.lb_intf_test", *intfRsc, path.Empty())...),
			},
		},
	})

}
