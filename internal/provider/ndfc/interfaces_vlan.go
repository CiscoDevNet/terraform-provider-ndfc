package ndfc

import (
	"context"
	"fmt"
	"terraform-provider-ndfc/internal/provider/resources/resource_interface_common"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type NDFCVlanInterface struct {
	NDFCInterfaceCommon
}

const ResourceVlanInterface = "interface_vlan"

func (i *NDFCVlanInterface) CreateInterface(ctx context.Context, diags *diag.Diagnostics, inData *resource_interface_common.NDFCInterfaceCommonModel) {
	intfPayload := resource_interface_common.NDFCInterfacesPayload{}
	intfPayload.Policy = inData.Policy
	if len(inData.Interfaces) <= 0 {
		tflog.Debug(ctx, "No interfaces to create")
		return
	}
	for i, intf := range inData.Interfaces {
		inData.Interfaces[i] = intf

		intfPayload.Interfaces = append(intfPayload.Interfaces, intf)
	}
	i.createInterface(ctx, diags, &intfPayload)
}

func (i *NDFCVlanInterface) DeleteInterface(ctx context.Context, dg *diag.Diagnostics,
	inData *resource_interface_common.NDFCInterfaceCommonModel) {

	tflog.Debug(ctx, "Deleting interfaces")
	if len(inData.Interfaces) <= 0 {
		tflog.Debug(ctx, "No interfaces to delete")
		return
	}

	// DELETE and Deploy uses similar payload
	intfPayload := resource_interface_common.NDFCInterfacesDeploy{}

	//ifDeployPayload := resource_interface_common.NDFCInterfacesDeploy{}

	for _, intf := range inData.Interfaces {
		intfPayload = append(intfPayload, resource_interface_common.NDFCInterfaceDeploy{IfName: intf.InterfaceName,
			SerialNumber: intf.SerialNumber})
		tflog.Debug(ctx, fmt.Sprintf("Deleting interface: %s:%s", intf.SerialNumber, intf.InterfaceName))
	}

	i.deleteInterface(ctx, dg, &intfPayload)
	if dg.HasError() {
		tflog.Error(ctx, "Error deleting interfaces")
		return
	}
	i.deployInterface(ctx, dg, intfPayload)
}
