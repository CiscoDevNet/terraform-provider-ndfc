// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package ndfc

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"terraform-provider-ndfc/internal/provider/datasources/datasource_interfaces"
	"terraform-provider-ndfc/internal/provider/resources/resource_interface_common"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
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
		if intf.InterfaceName == "" {
			dg.AddWarning("State is corrupted", fmt.Sprintf("InterfaceName is empty for entry %s", i))
			tflog.Error(ctx, fmt.Sprintf("InterfaceName is empty for entry %s", i))
			continue
		}

		if intf.SerialNumber == "" {
			dg.AddWarning("State is corrupted", fmt.Sprintf("SerialNumber is empty for entry %s", i))
			tflog.Error(ctx, fmt.Sprintf("SerialNumber is empty for entry %s", i))
			continue
		}

		keyMap[intf.SerialNumber+":"+intf.InterfaceName] = i
		log.Printf("Keymap: %s-%s", intf.SerialNumber+":"+intf.InterfaceName, i)
	}

	data := resource_interface_common.NDFCInterfaceCommonModel{}
	data.Interfaces = make(map[string]resource_interface_common.NDFCInterfacesValue)

	//out := resource_interface_common.InterfaceModel{}
	ID := in.GetID()
	//ifType := in.GetInterfaceType()

	data.Policy = inData.Policy
	data.PolicyType = inData.PolicyType
	data.SerialNumber = inData.SerialNumber
	data.Deploy = inData.Deploy
	ifMap := ifIdToMap(ID)

	for switchSerial, inList := range ifMap {
		ifSearchMap := make(map[string]bool)
		// Get interfaces for each switch
		ifObj := c.NewInterfaceObject(in.GetInterfaceType(), &c.apiClient, c.GetLock(ResourceInterfaces))
		ifList := ifObj.GetInterface(ctx, dg, switchSerial, data.Policy)
		//c.processCustomIfPolicy(ctx, dg, &ifList, data.PolicyType)

		// use datasource API to get interface deploy getDeployStatus
		dsIfModel := datasource_interfaces.NDFCInterfacesModel{}
		dsIfModel.InterfaceTypes = c.NDFCIfType(in.GetInterfaceType())
		dsIfModel.SerialNumber = switchSerial

		c.DsGetInterfaces(ctx, dg, &dsIfModel)
		gotDeployStatus := true
		if dg.HasError() {
			tflog.Error(ctx, "Error getting deployment status")
			gotDeployStatus = false
			return
		}

		for i := range inList {
			ifSearchMap[inList[i]] = true
		}

		for i := range ifList {
			if _, ok := ifSearchMap[ifList[i].InterfaceName]; ok {

				if ifList[i].NvPairs.FreeformConfig == " " {
					ifList[i].NvPairs.FreeformConfig = ""
				}
				key, ok := keyMap[ifList[i].SerialNumber+":"+ifList[i].InterfaceName]
				if !ok {
					if strings.Contains(ifList[i].SerialNumber, "~") {
						serialNumber := strings.Split(ifList[i].SerialNumber, "~")
						newSerialNumber := serialNumber[1] + "~" + serialNumber[0]
						key, ok = keyMap[newSerialNumber+":"+ifList[i].InterfaceName]
						if !ok {
							tflog.Error(ctx, fmt.Sprintf("Key not found: %s",
								ifList[i].SerialNumber+":"+ifList[i].InterfaceName))
							continue
						}
					} else {
						tflog.Error(ctx, fmt.Sprintf("Key not found: %s",
							ifList[i].SerialNumber+":"+ifList[i].InterfaceName))
						continue
					}
				}

				log.Printf("Found entry: key %s entry %s:%s", key, ifList[i].SerialNumber, ifList[i].InterfaceName)
				// Serial at resource level and per entry level are mutually exclusive
				// Set entry level to empty if resource level is set
				if inData.SerialNumber != "" {
					ifList[i].SerialNumber = ""
				}
				c.processCustomIfPolicy(ctx, dg, &ifList[i], data.PolicyType, inData.Interfaces[key].CustomPolicyParameters)
				data.Interfaces[key] = ifList[i]
				log.Printf("Add entry %s:%v", key, ifList[i])
			} else {
				log.Printf("Skip entry: %s", ifList[i].InterfaceName)
			}
		}
		//Search in If Details for deploy status
		if gotDeployStatus {
			for i := range dsIfModel.Interfaces {
				found := false
				ifName := ""
				if _, ok := ifSearchMap[dsIfModel.Interfaces[i].InterfaceName]; ok {

					found = true
					ifName = dsIfModel.Interfaces[i].InterfaceName
				} else if _, ok := ifSearchMap[strings.ToLower(dsIfModel.Interfaces[i].InterfaceName)]; ok {
					found = true
					ifName = strings.ToLower(dsIfModel.Interfaces[i].InterfaceName)
				}
				if found {
					key, ok := keyMap[switchSerial+":"+ifName]
					if !ok {
						tflog.Error(ctx, fmt.Sprintf("key not found for entry: %s", switchSerial+":"+ifName))
						continue
					}
					intf, ok := data.Interfaces[key]
					if !ok {
						tflog.Error(ctx, fmt.Sprintf("Read Error: Not found in read data %s", key))
					}
					// After a deployment, the status string is "Success" for some time - consider it as "In-Sync"
					if dsIfModel.Interfaces[i].DeploymentStatus == "In-Sync" || dsIfModel.Interfaces[i].DeploymentStatus == "Success" {
						intf.DeploymentStatus = "In-Sync"
					} else {
						intf.DeploymentStatus = dsIfModel.Interfaces[i].DeploymentStatus
					}
					data.Interfaces[key] = intf
					if data.Deploy && intf.DeploymentStatus != "In-Sync" {
						tflog.Warn(ctx, fmt.Sprintf("Interface %s is not yet in deployed state", intf.InterfaceName))
						dg.AddWarning(fmt.Sprintf("Interface %s is not yet in deployed state", intf.InterfaceName), "")
					}
				} else {
					log.Println("Interface not found: ", dsIfModel.Interfaces[i].InterfaceName)
				}
			}
		}
	}
	err := in.SetModelData(&data)
	if err.HasError() {
		dg.Append(err.Errors()...)
	}
}

func (c NDFC) RscCreateInterfaces(ctx context.Context, resp *resource.CreateResponse,
	in resource_interface_common.InterfaceModel) {

	dg := &resp.Diagnostics
	// Create API call logic
	tflog.Debug(ctx, fmt.Sprintf("RscCreateInterfaces: Creating interfaces for type %s", in.GetInterfaceType()))
	inData := in.GetModelData()
	c.IfTypeSet(inData, in.GetInterfaceType())
	c.IfPreProcess(inData)
	intfObj := c.NewInterfaceObject(in.GetInterfaceType(), &c.apiClient, c.GetLock(ResourceInterfaces))
	intfObj.CreateInterface(ctx, dg, inData)
	if dg.HasError() {
		in.SetID("")
		tflog.Error(ctx, "Error creating interfaces")
		return
	}
	if inData.Deploy {
		tflog.Info(ctx, "Deploying interfaces")
		intfObj.DeployInterface(ctx, dg, inData)
		if dg.HasError() {
			tflog.Error(ctx, "Error deploying interfaces")
		}
	}
	ID, _ := c.IfCreateID(ctx, inData)
	in.SetID(ID)
	//NDFC Bug. It takes some time for deploy status to be updated. Delay GET by few seconds
	log.Printf(" LET the deployed data sync - waiting 5 seconds")
	time.Sleep(5 * time.Second)
	c.RscGetInterfaces(ctx, dg, in)
	dg.Append(resp.State.Set(ctx, in)...)

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
		if dg.HasError() {
			tflog.Error(ctx, "Error deploying interfaces")
		}
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

func (c NDFC) RscImportInterfaces(ctx context.Context, dg *diag.Diagnostics, in resource_interface_common.InterfaceModel) {
	// Import API call logic
	tflog.Debug(ctx, "RscImportInterfaces: Importing interfaces")
	// Format 1 => policy:serial1[if1,if2],Serial2[if1,if2] => Selected interfaces
	// Format 2 => policy:serial => all interfaces in switch with serial
	ID := in.GetID()
	IdSplit := strings.Split(ID, ":")
	if len(IdSplit) != 2 {
		dg.AddError("Invalid ID", "Policy:Serial or policy:serial[ifName1, ifName2] format is expected from the ID")
		return
	}
	deployFlag := true

	IfData := IdSplit[1]
	data := resource_interface_common.NDFCInterfaceCommonModel{}
	data.Interfaces = make(map[string]resource_interface_common.NDFCInterfacesValue)
	data.Policy = IdSplit[0]
	log.Printf("Policy %s, IfData %s", data.Policy, IfData)

	ifMap := ifIdToMap(IfData)

	if len(ifMap) == 0 {
		//policy:serial format
		tflog.Debug(ctx, fmt.Sprintf("Importing all interfaces for switch: %s", IfData))
		ifMap[IfData] = []string{}
	}
	switchCount := 0
	for switchSerial, inputIfList := range ifMap {
		switchCount++
		ifSearchMap := make(map[string]bool)
		// Get interfaces for each switch
		ifObj := c.NewInterfaceObject(in.GetInterfaceType(), &c.apiClient, c.GetLock(ResourceInterfaces))
		ndfcIfList := ifObj.GetInterface(ctx, dg, switchSerial, data.Policy)
		// use datasource API to get interface deploy getDeployStatus
		dsIfModel := datasource_interfaces.NDFCInterfacesModel{}
		dsIfModel.InterfaceTypes = c.NDFCIfType(in.GetInterfaceType())
		dsIfModel.SerialNumber = switchSerial

		c.DsGetInterfaces(ctx, dg, &dsIfModel)
		gotDeployStatus := true
		if dg.HasError() {
			tflog.Error(ctx, "Error getting deployment status")
			gotDeployStatus = false
			//return
		}

		for i := range inputIfList {
			ifSearchMap[inputIfList[i]] = true
		}

		for i := range ndfcIfList {
			addIf := false
			if len(inputIfList) == 0 {
				addIf = true
			} else if _, ok := ifSearchMap[ndfcIfList[i].InterfaceName]; ok {
				addIf = true
			} else if IsFirstLetterLC(ndfcIfList[i].InterfaceName) {
				// NDFC returns ifName in lowercase for loopback, port-channel, vpc etc
				if _, ok := ifSearchMap[ToTitleCase(ndfcIfList[i].InterfaceName)]; ok {
					addIf = true
				}
			}
			if addIf {

				if ndfcIfList[i].NvPairs.FreeformConfig == " " {
					ndfcIfList[i].NvPairs.FreeformConfig = ""
				}
				key := getIFImportMapKey(switchSerial, ndfcIfList[i].InterfaceName, false)
				log.Printf("Import Entry entry: %s", key)
				// Serial at resource level and per entry level are mutually exclusive
				// Set entry level to empty if resource level is set
				data.Interfaces[key] = ndfcIfList[i]

			} else {
				log.Printf("Skip entry: %s:%s as entry is not found in input intf list", ndfcIfList[i].SerialNumber, ndfcIfList[i].InterfaceName)
			}
		}
		//Search in If Details for deploy status
		if gotDeployStatus {
			for i := range dsIfModel.Interfaces {
				ifName := dsIfModel.Interfaces[i].InterfaceName
				changeCase := false
			searchAgain:
				searchKey := getIFImportMapKey(switchSerial, ifName, changeCase)
				if ifEntry, ok := data.Interfaces[searchKey]; ok {
					ifEntry.DeploymentStatus = dsIfModel.Interfaces[i].DeploymentStatus
					data.Interfaces[searchKey] = ifEntry
					if ifEntry.DeploymentStatus != "In-Sync" {
						log.Printf("Interface %s DeploymentStatus |%s|", ifName, ifEntry.DeploymentStatus)
						// set to deploy to false as the interface is not in deployed state
						deployFlag = false
					}
				} else if !changeCase {
					changeCase = true
					// NDFC typically change ifName to lower case, so need to convert to upper and check
					log.Printf("Interface %s not found: changing case", ifName)
					goto searchAgain
				} else {
					log.Printf("Entry %s:%s doesn't exist - skipping: ", switchSerial, ifName)
				}
			}
		}
	}
	data.Deploy = deployFlag

	ID, _ = c.IfCreateID(ctx, &data)
	in.SetID(ID)
	err := in.SetModelData(&data)
	if err.HasError() {
		dg.Append(err.Errors()...)
	}
}

/* TF import has a bug where map keys starting with numbers are NOK
 * So, we are using a different key for map entries
 */
func getIFImportMapKey(serial, ifName string, changeCase bool) string {

	if changeCase {
		if IsFirstLetterLC(ifName) {
			ifName = ToTitleCase(ifName)
		} else {
			ifName = strings.ToLower(ifName)
		}
	}
	//change / to _ in ifName
	s := strings.Replace(ifName, "/", "_", -1)
	return s + "_" + serial
}

func (c NDFC) processCustomIfPolicy(ctx context.Context, diags *diag.Diagnostics,
	payload *resource_interface_common.NDFCInterfacesValue, pType string, custParams map[string]string) error {
	log.Printf("[DEBUG] Processing custom policy")

	if pType == "user-defined" {
		// Reset NvPairs as this is a custom template
		log.Printf("[DEBUG] Resetting NvPairs for user-defined policy")
		(*payload).NvPairs = resource_interface_common.NDFCNvPairsValue{}
		// Filter the custom values
		for k, _ := range payload.CustomPolicyParameters {
			if _, ok := (custParams)[k]; !ok {
				log.Printf("[DEBUG] Removing custom parameter %s", k)
				delete(payload.CustomPolicyParameters, k)
			}
		}
	} else {
		// Reset the custom template parameters
		log.Printf("[DEBUG] Resetting custom parameters for system policy")
		clear((*payload).CustomPolicyParameters)
	}

	return nil
}
