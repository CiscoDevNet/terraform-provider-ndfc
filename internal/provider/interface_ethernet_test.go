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

var rsType string = "ethernet"

// Create Ethernet resouce with 10 entries, add 10 more
func TestAccInterfaceEthernetResourceCreateAndAdd(t *testing.T) {

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
		PreCheck:                 func() { testAccPreCheck(t, rsType) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		//CheckDestroy:             testAccCheckVrfBulkResourceDestroy(networkRsc),
		Steps: []resource.TestStep{
			{
				// Create a new resource with 10 interfaces
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					//Starting from 10 to consider anything pre-existing on ndfc
					helper.GenerateIntfResource(&intfRsc, 10, 10, "ethernet", true, helper.GetConfig("ethernet").NDFC.Switches, true, false)
					(*x)["RscName"] = "eth_intf_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfaceEthernetModelHelperStateCheck("ndfc_interface_ethernet.eth_intf_test", *intfRsc, path.Empty())...),
			},
			// Add some more intfs
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateIntfResource(&intfRsc, 20, 10, "ethernet", true, helper.GetConfig("ethernet").NDFC.Switches, true, true)
					(*x)["RscName"] = "eth_intf_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfaceEthernetModelHelperStateCheck("ndfc_interface_ethernet.eth_intf_test", *intfRsc, path.Empty())...),
			},
		},
	})
}

// Create with 20 entries, delete 10
func TestAccInterfaceEthernetResourceCreateAndReduce(t *testing.T) {
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
		PreCheck:                 func() { testAccPreCheck(t, rsType) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		//CheckDestroy:             testAccCheckVrfBulkResourceDestroy(networkRsc),
		Steps: []resource.TestStep{
			{
				// Create a new resource with 20 interfaces
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					//Starting from 20 to consider anything pre-existing on ndfc
					helper.GenerateIntfResource(&intfRsc, 10, 20, "ethernet", true, helper.GetConfig("ethernet").NDFC.Switches, true, false)
					(*x)["RscName"] = "eth_intf_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfaceEthernetModelHelperStateCheck("ndfc_interface_ethernet.eth_intf_test", *intfRsc, path.Empty())...),
			},
			// Delete 10 interfaces
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.GenerateIntfResource(&intfRsc, 10, -10, "ethernet", true, helper.GetConfig("ethernet").NDFC.Switches, true, true)
					(*x)["RscName"] = "eth_intf_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfaceEthernetModelHelperStateCheck("ndfc_interface_ethernet.eth_intf_test", *intfRsc, path.Empty())...),
			},
		},
	})
}

// Create with 20 entries, modify 5
func TestAccInterfaceEthernetResourceCreateAndModify(t *testing.T) {

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
		PreCheck:                 func() { testAccPreCheck(t, rsType) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		//CheckDestroy:             testAccCheckVrfBulkResourceDestroy(networkRsc),
		Steps: []resource.TestStep{
			{
				// Create a new resource with 20 interfaces
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					//Starting from 20 to consider anything pre-existing on ndfc
					helper.GenerateIntfResource(&intfRsc, 10, 20, "ethernet", true, helper.GetConfig("ethernet").NDFC.Switches, true, false)
					(*x)["RscName"] = "eth_intf_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfaceEthernetModelHelperStateCheck("ndfc_interface_ethernet.eth_intf_test", *intfRsc, path.Empty())...),
			},
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					helper.ModifyInterface(&intfRsc, 20, 5, "ethernet", map[string]interface{}{
						"admin_state": "down",
						"mtu":         "default",
						"speed":       "Auto",
						"bpduGuard":   "false",
						"accessVlan":  types.Int64Custom(1000),
						"nativeVlan":  types.Int64Custom(1500),
					})
					(*x)["RscName"] = "eth_intf_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfaceEthernetModelHelperStateCheck("ndfc_interface_ethernet.eth_intf_test", *intfRsc, path.Empty())...),
			},
		},
	})
}

// Create with 10 entries, add 10 delete 5 and modify 5 all in one shot
func TestAccInterfaceEthernetResourceCombinedUpdate(t *testing.T) {

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
		PreCheck:                 func() { testAccPreCheck(t, rsType) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		//CheckDestroy:             testAccCheckVrfBulkResourceDestroy(networkRsc),
		Steps: []resource.TestStep{
			{
				// Create a new resource with 20 interfaces
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					//Starting from 20 to consider anything pre-existing on ndfc
					helper.GenerateIntfResource(&intfRsc, 10, 10, "ethernet", true, helper.GetConfig("ethernet").NDFC.Switches, true, false)
					(*x)["RscName"] = "eth_intf_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfaceEthernetModelHelperStateCheck("ndfc_interface_ethernet.eth_intf_test", *intfRsc, path.Empty())...),
			},
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
					//Add 10
					helper.GenerateIntfResource(&intfRsc, 20, 10, "ethernet", true, helper.GetConfig("ethernet").NDFC.Switches, true, true)
					//Delete 5
					helper.GenerateIntfResource(&intfRsc, 10, -5, "ethernet", true, helper.GetConfig("ethernet").NDFC.Switches, true, true)
					//Modify 5
					helper.ModifyInterface(&intfRsc, 20, 5, "ethernet", map[string]interface{}{
						"admin_state":  "up",
						"mtu":          "jumbo",
						"speed":        "Auto",
						"bpduGuard":    "true",
						"accessVlan":   types.Int64Custom(1200),
						"nativeVlan":   types.Int64Custom(2400),
						"allowedVlans": "10-2000",
					})
					(*x)["RscName"] = "eth_intf_test"
					helper.GetTFConfigWithSingleResource(tName, *x, []interface{}{intfRsc}, &tf_config)
					return *tf_config
				}(),
				Check: resource.ComposeTestCheckFunc(InterfaceEthernetModelHelperStateCheck("ndfc_interface_ethernet.eth_intf_test", *intfRsc, path.Empty())...),
			},
		},
	})

}
