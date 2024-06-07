package ndfc

import (
	"context"
	"fmt"
	"log"
	"strings"

	"terraform-provider-ndfc/internal/provider/datasources/datasource_interfaces"
	"terraform-provider-ndfc/internal/provider/resources/resource_interface_common"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const ResourceInterfaces = "interfaces"

func (c NDFC) RscGetInterfaces(ctx context.Context, dg *diag.Diagnostics, in resource_interface_common.InterfaceModel) {
	// Get API call logic

	keyMap := make(map[string]string)
	inData := in.GetModelData()
	for i, intf := range inData.Interfaces {
		if inData.Interfaces[i].SerialNumber == "" {
			intf.SerialNumber = inData.SerialNumber
		}
		keyMap[intf.SerialNumber+":"+strings.ToLower(intf.InterfaceName)] = i
		log.Printf("Keymap: %s-%s", intf.SerialNumber+":"+strings.ToLower(intf.InterfaceName), i)
	}

	data := resource_interface_common.NDFCInterfaceCommonModel{}
	data.Interfaces = make(map[string]resource_interface_common.NDFCInterfacesValue)

	//out := resource_interface_common.InterfaceModel{}
	ID := in.GetID()
	//ifType := in.GetInterfaceType()

	data.Policy = inData.Policy
	data.SerialNumber = inData.SerialNumber
	data.Deploy = inData.Deploy
	ifMap := ifIdToMap(ID)

	for switchSerial, inList := range ifMap {
		ifSearchMap := make(map[string]bool)
		ifObj := c.NewInterfaceObject(in.GetInterfaceType(), &c.apiClient, c.GetLock(ResourceInterfaces))
		ifList := ifObj.GetInterface(ctx, dg, switchSerial, data.Policy)

		for i := range inList {
			ifSearchMap[strings.ToLower(inList[i])] = true
		}
		for i := range ifList {
			if _, ok := ifSearchMap[strings.ToLower(ifList[i].InterfaceName)]; ok {

				if ifList[i].NvPairs.FreeformConfig == " " {
					ifList[i].NvPairs.FreeformConfig = ""
				}
				key, ok := keyMap[ifList[i].SerialNumber+":"+ifList[i].InterfaceName]
				if !ok {
					panic(fmt.Sprintf("Key not found: %s", ifList[i].SerialNumber+":"+ifList[i].InterfaceName))
				}

				log.Printf("Found entry: key %s entry %s:%s", key, ifList[i].SerialNumber, ifList[i].InterfaceName)
				// Serial at resource level and per entry level are mutually exclusive
				// Set entry level to empty if resource level is set
				if inData.SerialNumber != "" {
					ifList[i].SerialNumber = ""
				}
				data.Interfaces[key] = ifList[i]
				log.Printf("Add entry %s:%v", key, ifList[i])
			} else {
				log.Printf("Skip entry: %s", ifList[i].InterfaceName)
			}
		}
	}
	err := in.SetModelData(&data)
	if err.HasError() {
		dg.Append(err.Errors()...)
	}
}

func (c NDFC) RscCreateInterfaces(ctx context.Context, dg *diag.Diagnostics, in resource_interface_common.InterfaceModel) {
	// Create API call logic
	tflog.Debug(ctx, fmt.Sprintf("RscCreateInterfaces: Creating interfaces for type %s", in.GetInterfaceType()))
	inData := in.GetModelData()
	c.IfTypeSet(inData, in.GetInterfaceType())
	c.IfPreProcess(inData)
	intfObj := c.NewInterfaceObject(in.GetInterfaceType(), &c.apiClient, c.GetLock(ResourceInterfaces))
	intfObj.CreateInterface(ctx, dg, inData)
	if dg.HasError() {
		tflog.Error(ctx, "Error creating interfaces")
		return
	}

	if inData.Deploy {
		tflog.Info(ctx, "Deploying interfaces")
		intfObj.DeployInterface(ctx, dg, inData)
		if dg.HasError() {
			tflog.Error(ctx, "Error deploying interfaces")
			return
		}
	}
	ID, _ := c.IfCreateID(ctx, inData)
	in.SetID(ID)
	c.RscGetInterfaces(ctx, dg, in)
}

func (c NDFC) RscUpdateInterfaces(ctx context.Context, dg *diag.Diagnostics, unique_id string,
	planData resource_interface_common.InterfaceModel,
	stateData resource_interface_common.InterfaceModel,
	configData resource_interface_common.InterfaceModel) {

	state := stateData.GetModelData()
	plan := planData.GetModelData()

	c.IfTypeSet(plan, planData.GetInterfaceType())
	c.IfTypeSet(state, stateData.GetInterfaceType())

	c.IfPreProcess(plan)
	c.IfPreProcess(state)

	actions := c.ifDiff(ctx, dg, state, plan)
	ifObj := c.NewInterfaceObject(planData.GetInterfaceType(), &c.apiClient, c.GetLock(ResourceInterfaces))

	//Delete any interfaces marked for delete
	tflog.Debug(ctx, "Deleting interfaces marked for delete")
	delIntf := actions["del"].(*resource_interface_common.NDFCInterfaceCommonModel)
	ifObj.DeleteInterface(ctx, dg, delIntf)
	if dg.HasError() {
		tflog.Error(ctx, "Error deleting interfaces")
		return
	}
	// Perform updates
	tflog.Debug(ctx, "Creating new interfaces")
	createIntf := actions["create"].(*resource_interface_common.NDFCInterfaceCommonModel)
	c.IfPreProcess(createIntf)
	ifObj.CreateInterface(ctx, dg, createIntf)
	if dg.HasError() {
		tflog.Error(ctx, "Error creating interfaces")
		return
	}

	tflog.Debug(ctx, "Updating interfaces")
	updateIntf := actions["update"].(*resource_interface_common.NDFCInterfaceCommonModel)
	c.IfPreProcess(updateIntf)
	ifObj.ModifyInterface(ctx, dg, updateIntf)
	if dg.HasError() {
		tflog.Error(ctx, "Error updating interfaces")
		return
	}

	if actions["deploy"].(bool) || plan.Deploy {
		tflog.Info(ctx, "Deploy flag is set  - deploy all interfaces in plan")
		ifObj.DeployInterface(ctx, dg, plan)
		//c.RscDeployInterfaces(ctx, dg, plan)
	}
	/*else {
		tflog.Info(ctx, "Deploy flag has not changed in plan - check and deploy modified interfaces")
		if plan.Deploy {
			if len(updateIntf.Interfaces) > 0 {
				tflog.Info(ctx, "Deploy flag is set. Deploying modified interfaces")
				ifObj.DeployInterface(ctx, dg, updateIntf)
				//c.RscDeployInterfaces(ctx, dg, updateIntf)
			} else {
				tflog.Info(ctx, "Deploy flag is set. No modified interfaces to deploy")
			}
		} else {
			tflog.Info(ctx, "Deploy flag is not set. Not deploying modified interfaces")
		}
	}
	*/
	ID, _ := c.IfCreateID(ctx, plan)
	planData.SetID(ID)
	// Fill resp with data from NDFC
	c.RscGetInterfaces(ctx, dg, planData)
}

func (c NDFC) RscDeleteInterfaces(ctx context.Context, dg *diag.Diagnostics, in resource_interface_common.InterfaceModel) {
	// Delete API call logic
	inData := in.GetModelData()
	c.IfTypeSet(inData, in.GetInterfaceType())
	c.IfPreProcess(inData)

	ifObj := c.NewInterfaceObject(in.GetInterfaceType(), &c.apiClient, c.GetLock(ResourceInterfaces))
	ifObj.DeleteInterface(ctx, dg, inData)
	if dg.HasError() {
		tflog.Error(ctx, "Error deleting interfaces")
		return
	}
	/*
		ifObj.DeployInterface(ctx, dg, inData)
		//c.RscDeployInterfaces(ctx, dg, in)
		if dg.HasError() {
			tflog.Error(ctx, "Error deploying interfaces")
			return
		}
	*/
}

func (c NDFC) DsGetInterfaces(ctx context.Context, dg *diag.Diagnostics, in *datasource_interfaces.NDFCInterfacesModel) {
	// Get API call logic
	tflog.Debug(ctx, "DsGetInterfaces: Getting interfaces")
	ifObj := c.NewInterfaceObject("datasource", &c.apiClient, c.GetLock(ResourceInterfaces))
	ifObj.(*NDFCInterfaceCommon).DsGetInterfaceDetails(ctx, dg, in)

}
