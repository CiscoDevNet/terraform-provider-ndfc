package provider

import (
	"terraform-provider-ndfc/internal/provider/resources/resource_configuration_deploy"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func ConfigurationDeployModelHelperStateCheck(RscName string, c resource_configuration_deploy.NDFCConfigurationDeployModel, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.FabricName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_name").String(), c.FabricName))
	}
	return ret
}
