package ndfc

import (
	"context"
	"fmt"
	"log"
	"terraform-provider-ndfc/internal/provider/resources/resource_interface_common"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type NDFCEthernetInterface struct {
	NDFCInterfaceCommon
}

const ResourceEthernetInterface = "interface_ethernet"

func (i *NDFCEthernetInterface) CreateInterface(ctx context.Context, diags *diag.Diagnostics, inData *resource_interface_common.NDFCInterfaceCommonModel) {
	intfPayload := resource_interface_common.NDFCInterfacesPayload{}
	intfPayload.Policy = inData.Policy
	for i, intf := range inData.Interfaces {
		inData.Interfaces[i] = intf
		intfPayload.Interfaces = append(intfPayload.Interfaces, intf)
	}
	i.modifyInterface(ctx, diags, &intfPayload)

}

func (i *NDFCEthernetInterface) DeleteInterface(ctx context.Context, dg *diag.Diagnostics,
	inData *resource_interface_common.NDFCInterfaceCommonModel) {

	tflog.Debug(ctx, "Deleting interfaces")
	if len(inData.Interfaces) <= 0 {
		tflog.Debug(ctx, "No interfaces to delete")
		return
	}
	intfPayload := resource_interface_common.NDFCInterfacesPayload{}
	intfPayload.Policy = inData.Policy

	ifDeployPayload := resource_interface_common.NDFCInterfacesDeploy{}

	for k, intf := range inData.Interfaces {
		tflog.Debug(ctx, fmt.Sprintf("Deleting interface: %s:%s", intf.SerialNumber, intf.InterfaceName))
		intf.NvPairs.AdminState = "false"
		intf.NvPairs.InterfaceDescription = "DELETED BY TERRAFORM"
		intf.NvPairs.FreeformConfig = " "
		intf.NvPairs.Speed = "Auto"
		intf.NvPairs.Mtu = "jumbo"
		inData.Interfaces[k] = intf
		intfPayload.Interfaces = append(intfPayload.Interfaces, intf)
		ifDeployPayload = append(ifDeployPayload, resource_interface_common.NDFCInterfaceDeploy{
			IfName:       intf.InterfaceName,
			SerialNumber: intf.SerialNumber,
		})

	}
	i.modifyInterface(ctx, dg, &intfPayload)
	if dg.HasError() {
		tflog.Error(ctx, "Error deleting interfaces")
		return
	}
	i.deployInterface(ctx, dg, ifDeployPayload)
}

func printModel(model resource_interface_common.InterfaceModel) {
	log.Printf("Model: %v", model)
	inData := model.GetModelData()
	log.Printf("Model Data: %v", inData)
	for i := range inData.Interfaces {
		log.Printf("Interface: |%v|", inData.Interfaces[i])
		log.Printf("Interface Name: |%s|", inData.Interfaces[i].InterfaceName)
		log.Printf("Serial Number: |%s|", inData.Interfaces[i].SerialNumber)
		log.Printf("Freeform Config: |%s|", inData.Interfaces[i].NvPairs.FreeformConfig)
		log.Printf("Speed: |%s|", inData.Interfaces[i].NvPairs.Speed)
		log.Printf("Mtu: |%s|", inData.Interfaces[i].NvPairs.Mtu)
		log.Printf("Port Type Fast: |%s|", inData.Interfaces[i].NvPairs.PortTypeFast)
		log.Printf("Bpdu Guard: |%s|", inData.Interfaces[i].NvPairs.BpduGuard)
		if inData.Interfaces[i].NvPairs.AccessVlan != nil {
			log.Printf("Access Vlan: |%d|", *inData.Interfaces[i].NvPairs.AccessVlan)
		}
		log.Printf("Interface Description: |%s|", inData.Interfaces[i].NvPairs.InterfaceDescription)
		log.Printf("Orphan Port: |%s|", inData.Interfaces[i].NvPairs.OrphanPort)
		log.Printf("AdminState: |%s|", inData.Interfaces[i].NvPairs.AdminState)
		log.Printf("Ptp : |%s|", inData.Interfaces[i].NvPairs.Ptp)
		log.Printf("Netflow : |%s|", inData.Interfaces[i].NvPairs.Netflow)
		log.Printf("NetflowMonitor : |%s|", inData.Interfaces[i].NvPairs.NetflowMonitor)
		log.Printf("NetflowSampler : |%s|", inData.Interfaces[i].NvPairs.NetflowSampler)
		log.Printf("AllowedVlans : |%s|", inData.Interfaces[i].NvPairs.AllowedVlans)
		log.Printf("NativeVlan : |%v|", *inData.Interfaces[i].NvPairs.NativeVlan)
		log.Printf("Interface Type: |%s|", model.GetInterfaceType())
		log.Printf("Policy: |%s|", inData.Policy)
	}
}
