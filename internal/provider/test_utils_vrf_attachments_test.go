// Code generated;  DO NOT EDIT.

package provider

import (
	"strconv"
	"terraform-provider-ndfc/internal/provider/resources/resource_vrf_attachments"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func VrfAttachmentsModelHelperStateCheck(RscName string, c resource_vrf_attachments.NDFCVrfAttachmentsModel, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.FabricName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_name").String(), c.FabricName))
	}
	if c.DeployAllAttachments {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("deploy_all_attachments").String(), "true"))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("deploy_all_attachments").String(), "false"))
	}
	for key, value := range c.VrfAttachments {
		attrNewPath := attrPath.AtName("vrf_attachments").AtName(key)
		ret = append(ret, VrfAttachmentsValueHelperStateCheck(RscName, value, attrNewPath)...)
	}
	return ret
}

func AttachListValueHelperStateCheck(RscName string, c resource_vrf_attachments.NDFCAttachListValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.SwitchName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("switch_name").String(), c.SwitchName))
	}
	if c.Vlan != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vlan").String(), strconv.Itoa(int(*c.Vlan))))
	}

	if c.AttachState != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("attach_state").String(), c.AttachState))
	}
	if c.Attached != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("attached").String(), strconv.FormatBool(*c.Attached)))
	}
	if c.FreeformConfig != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("freeform_config").String(), c.FreeformConfig))
	}
	if c.DeployThisAttachment {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("deploy_this_attachment").String(), "true"))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("deploy_this_attachment").String(), "false"))
	}
	if c.InstanceValues.LoopbackId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("loopback_id").String(), strconv.Itoa(int(*c.InstanceValues.LoopbackId))))
	}
	if c.InstanceValues.LoopbackIpv4 != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("loopback_ipv4").String(), c.InstanceValues.LoopbackIpv4))
	}
	if c.InstanceValues.LoopbackIpv6 != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("loopback_ipv6").String(), c.InstanceValues.LoopbackIpv6))
	}

	return ret
}

func VrfAttachmentsValueHelperStateCheck(RscName string, c resource_vrf_attachments.NDFCVrfAttachmentsValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.VrfName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_name").String(), c.VrfName))
	}
	if c.DeployAllAttachments {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("deploy_all_attachments").String(), "true"))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("deploy_all_attachments").String(), "false"))
	}
	for key, value := range c.AttachList {
		attrNewPath := attrPath.AtName("attach_list").AtName(key)
		ret = append(ret, AttachListValueHelperStateCheck(RscName, value, attrNewPath)...)
	}
	return ret
}
