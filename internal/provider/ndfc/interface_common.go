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
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"
	"terraform-provider-ndfc/internal/provider/datasources/datasource_interfaces"
	"terraform-provider-ndfc/internal/provider/ndfc/api"
	"terraform-provider-ndfc/internal/provider/resources/resource_interface_common"

	"github.com/netascode/go-nd"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Common Implementation - override for changing behaviour
type NDFCInterface interface {
	CreateInterface(ctx context.Context, diags *diag.Diagnostics, inData *resource_interface_common.NDFCInterfaceCommonModel)
	DeleteInterface(ctx context.Context, diags *diag.Diagnostics, inData *resource_interface_common.NDFCInterfaceCommonModel)
	ModifyInterface(ctx context.Context, diags *diag.Diagnostics, inData *resource_interface_common.NDFCInterfaceCommonModel)
	GetInterface(ctx context.Context, diags *diag.Diagnostics, serial string, policy string) []resource_interface_common.NDFCInterfacesValue
	DeployInterface(ctx context.Context, diags *diag.Diagnostics, inData *resource_interface_common.NDFCInterfaceCommonModel)
	GetPayload(ctx context.Context, diags *diag.Diagnostics, inData *resource_interface_common.NDFCInterfacesPayload) ([]byte, error)
}

type NDFCInterfaceCommon struct {
	NDFCInterface
	client *nd.Client
	lock   *sync.Mutex
	ifType string
}

func (c NDFC) NewInterfaceObject(ifType string, client *nd.Client, lock *sync.Mutex) NDFCInterface {
	switch ifType {
	case "ethernet":
		intf := new(NDFCEthernetInterface)
		intf.client = client
		intf.lock = lock
		intf.NDFCInterface = intf
		intf.ifType = ifType
		return intf
	case "loopback":
		intf := new(NDFCLoopbackInterface)
		intf.client = client
		intf.lock = lock
		intf.NDFCInterface = intf
		intf.ifType = ifType
		return intf
	case "datasource":
		intf := new(NDFCInterfaceCommon)
		intf.client = client
		intf.lock = lock
		intf.NDFCInterface = intf
		intf.ifType = ifType
		return intf
	case "vlan":
		intf := new(NDFCVlanInterface)
		intf.client = client
		intf.lock = lock
		intf.NDFCInterface = intf
		intf.ifType = ifType
		return intf
	case "portchannel":
		intf := new(NDFCPortChannelInterface)
		intf.client = client
		intf.lock = lock
		intf.NDFCInterface = intf
		intf.ifType = ifType
		return intf
	case "vpc":
		intf := new(NDFCVPCInterface)
		intf.client = client
		intf.lock = lock
		intf.NDFCInterface = intf
		intf.ifType = ifType
		return intf

	default:
		log.Panicf("Interface type not supported: %s", ifType)
	}
	return nil
}

func (i *NDFCInterfaceCommon) CreateInterface(ctx context.Context, dg *diag.Diagnostics,
	inData *resource_interface_common.NDFCInterfaceCommonModel) {
	log.Panicf("CreateInterface not implemented in common level")
}

func (i *NDFCInterfaceCommon) DeleteInterface(ctx context.Context, diags *diag.Diagnostics,
	inData *resource_interface_common.NDFCInterfaceCommonModel) {
	log.Panicf("DeleteInterface not implemented in common level")
}

// Common modify logic for all interface types
func (i *NDFCInterfaceCommon) ModifyInterface(ctx context.Context, dg *diag.Diagnostics,
	inData *resource_interface_common.NDFCInterfaceCommonModel) {
	if len(inData.Interfaces) > 0 {
		updateIntf := new(resource_interface_common.NDFCInterfacesPayload)
		updateIntf.Policy = inData.Policy
		for i := range inData.Interfaces {
			updateIntf.Interfaces = append(updateIntf.Interfaces, inData.Interfaces[i])
		}
		i.modifyInterface(ctx, dg, updateIntf)
		if dg.HasError() {
			tflog.Error(ctx, "Error updating interfaces")
			return
		}
	} else {
		tflog.Debug(ctx, "No interfaces to modify")
	}
}

func (i *NDFCInterfaceCommon) GetInterface(ctx context.Context, diags *diag.Diagnostics, serial string,
	policy string) []resource_interface_common.NDFCInterfacesValue {
	return i.getInterfaces(ctx, diags, serial, policy)
}

func (i *NDFCInterfaceCommon) DeployInterface(ctx context.Context, dg *diag.Diagnostics,
	in *resource_interface_common.NDFCInterfaceCommonModel) {
	log.Printf("Deploying interfaces on : %v", in.SerialNumber)
	payload := resource_interface_common.NDFCInterfacesDeploy{}
	for i := range in.Interfaces {
		ifEntry := resource_interface_common.NDFCInterfaceDeploy{
			IfName:       in.Interfaces[i].InterfaceName,
			SerialNumber: in.Interfaces[i].SerialNumber,
		}
		payload = append(payload, ifEntry)
	}
	i.deployInterface(ctx, dg, payload)
}

func (i *NDFCInterfaceCommon) GetPayload(ctx context.Context, diags *diag.Diagnostics,
	intfPayload *resource_interface_common.NDFCInterfacesPayload) ([]byte, error) {
	return json.Marshal(intfPayload)
}

func (i *NDFCInterfaceCommon) deployInterface(ctx context.Context, dg *diag.Diagnostics, payload resource_interface_common.NDFCInterfacesDeploy) {
	GlobalDeployLock("interface")
	defer GlobalDeployUnlock("interface")
	data, err := json.Marshal(payload)
	if err != nil {
		dg.AddError("Error marshalling data", err.Error())
		return
	}
	ifApi := api.NewInterfaceAPI(i.lock, i.client)
	ifApi.SetAPI(api.PostInterfaceDeploy)
	//ifApi.SetDeployLocked()
	res, err := ifApi.DeployPost(data)
	//res, err := i.client.Post(UrlInterfaceDeploy, string(data))
	if err != nil {
		dg.AddError("Error deploying interface:", fmt.Sprintf("%s:%s", err.Error(), res))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Deploy Response: %s", res))
}

func (i *NDFCInterfaceCommon) getInterfaces(ctx context.Context, diags *diag.Diagnostics, serial string, policy string) []resource_interface_common.NDFCInterfacesValue {
	ifList := make([]resource_interface_common.NDFCInterfacesPayload, 0)
	ifApi := api.NewInterfaceAPI(i.lock, i.client)
	ifApi.SwitchSerial = serial
	ifApi.Policy = policy
	res, err := ifApi.Get()
	//res, err := c.apiClient.GetRawJsonWithQueryString(url, []string{fmt.Sprintf("serialNumber=%s", serial), fmt.Sprintf("templateName=%s", policy)})
	if err != nil {
		tflog.Error(ctx, "Error getting interfaces")
		diags.AddError("Error getting interfaces", err.Error())
		return nil
	}
	log.Printf("Response=%s", string(res))
	err = json.Unmarshal((res), &ifList)
	if err != nil {
		diags.AddError("Error unmarshalling data", err.Error())
		return nil
	}
	if len(ifList) == 0 {
		diags.AddWarning("No interfaces found", "")
		return nil
	}
	return ifList[0].Interfaces
}

func (i *NDFCInterfaceCommon) modifyInterface(ctx context.Context, diags *diag.Diagnostics,
	intfPayload *resource_interface_common.NDFCInterfacesPayload) {

	data, err := i.NDFCInterface.GetPayload(ctx, diags, intfPayload)
	if err != nil {
		diags.AddError("Error marshalling data", err.Error())
		return
	}
	ifApi := api.NewInterfaceAPI(i.lock, i.client)
	res, err := ifApi.Put(data)
	//res, err := c.apiClient.Put("/lan-fabric/rest/interface", string(data))
	if err != nil {
		diags.AddError("Error updating interface:", fmt.Sprintf("%s:%s", err.Error(), res))
		return
	}
	tflog.Debug(ctx, fmt.Sprintf("Response: %s", res))

}

/*
func (i *NDFCInterfaceCommon) getDeployStatus(ctx context.Context, diags *diag.Diagnostics, serial string) {
	ifApi := api.NewInterfaceAPI(i.lock, i.client)
	ifApi.SwitchSerial = serial
	res, err := ifApi.Get()
	//res, err := c.apiClient.GetRawJsonWithQueryString(url, []string{fmt.Sprintf("serialNumber=%s", serial)})
	if err != nil {
		diags.AddError("Error getting interfaces", err.Error())
		return
	}
	log.Printf("Response=%s", string(res))
}
*/

func (i *NDFCInterfaceCommon) createInterface(ctx context.Context, diags *diag.Diagnostics,
	intfPayload *resource_interface_common.NDFCInterfacesPayload) {

	data, err := i.NDFCInterface.GetPayload(ctx, diags, intfPayload)
	if err != nil {
		diags.AddError("Error marshalling data", err.Error())
		return
	}

	ifApi := api.NewInterfaceAPI(i.lock, i.client)
	ifApi.SetAPI(api.PostInterfaceCreate)
	res, err := ifApi.Post(data)
	//res, err := c.apiClient.Post("/lan-fabric/rest/interface", string(data))
	if err != nil {
		diags.AddError("Error creating interface:", fmt.Sprintf("%s:%s", err.Error(), res))
		return
	}
	tflog.Debug(ctx, fmt.Sprintf("Response: %s", res))
}

func (i *NDFCInterfaceCommon) deleteInterface(ctx context.Context, diags *diag.Diagnostics,
	intfPayload *resource_interface_common.NDFCInterfacesDeploy) {

	data, err := json.Marshal(intfPayload)
	if err != nil {
		diags.AddError("Error marshalling data", err.Error())
		return
	}

	ifApi := api.NewInterfaceAPI(i.lock, i.client)
	//ifApi.SetAPI(api.DeleteInterface)
	res, err := ifApi.DeleteWithPayload(data)
	//res, err := c.apiClient.Delete("/lan-fabric/rest/interface", string(data))
	if err != nil {
		diags.AddError("Error deleting interface:", fmt.Sprintf("%s:%s", err.Error(), res))
		return
	}
	tflog.Debug(ctx, fmt.Sprintf("Response: %s", res))
}

func (i *NDFCInterfaceCommon) DsGetInterfaceDetails(ctx context.Context, diags *diag.Diagnostics,
	inData *datasource_interfaces.NDFCInterfacesModel) {

	ifApi := api.NewInterfaceAPI(i.lock, i.client)
	//For vPC interface serial no is of form <peer-1>~<peer-2>
	ifApi.SwitchSerial = strings.Split(inData.SerialNumber, "~")[0]
	ifApi.PortMode = inData.PortModes
	ifApi.Excludes = inData.Excludes
	ifApi.IfTypes = inData.InterfaceTypes
	ifApi.SetAPI(api.GetInterfaceDetailed)
	tflog.Debug(ctx, fmt.Sprintf("DsGetInterfaceDetails: Getting interfaces for serial=|%s|,PortMode=|%s| ifTypes=|%s| excludes=|%s|",
		inData.SerialNumber,
		inData.PortModes,
		inData.InterfaceTypes,
		inData.Excludes))
	res, err := ifApi.Get()
	if err != nil {
		diags.AddError("Error getting interfaces", err.Error()+string(res))
		tflog.Error(ctx, fmt.Sprintf("Error getting interfaces: %s: %v", err.Error(), string(res)))
		return
	}
	log.Printf("Response=%s", string(res))
	err = json.Unmarshal((res), &inData.Interfaces)
	if err != nil {
		diags.AddError("Error unmarshalling data", err.Error())
		return
	}
}
