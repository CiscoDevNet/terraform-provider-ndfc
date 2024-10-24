// Code generated;  DO NOT EDIT.

package provider

import (
	"strconv"
	"terraform-provider-ndfc/internal/provider/resources/resource_vpc_pair"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func VpcPairModelHelperStateCheck(RscName string, c resource_vpc_pair.NDFCVpcPairModel, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.UseVirtualPeerlink != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("use_virtual_peerlink").String(), strconv.FormatBool(*c.UseVirtualPeerlink)))
	}
	if c.NvPairs.FabricName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_name").String(), c.NvPairs.FabricName))
	}
	if c.Deploy {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("deploy").String(), "true"))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("deploy").String(), "false"))
	}
	return ret
}
