package provider

import (
	"terraform-provider-ndfc/internal/provider/resources/resource_inventory"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Inside TEST_HELPER_STATE_CHECK
func InventoryModelHelperStateCheck(RscName string, c resource_inventory.NDFCInventoryModel, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	for key, value := range c.SwitchDiscovery {
		attrNewPath := attrPath.AtName("switch_discovery").AtName(key)
		ret = append(ret, SwitchDiscoveryValueHelperStateCheck(RscName, value, attrNewPath)...)
	}
	return ret
}

// Inside TEST_HELPER_STATE_CHECK
func SwitchDiscoveryValueHelperStateCheck(RscName string, c resource_inventory.NDFCSwitchDiscoveryValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.UserName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("user_name").String(), c.UserName))
	}
	if c.Password != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("password").String(), c.Password))
	}
	if c.Role != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("role").String(), c.Role))
	}
	if c.PreserveConfig != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("preserve_config").String(), c.PreserveConfig))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("preserve_config").String(), "true"))
	}

	return ret
}
