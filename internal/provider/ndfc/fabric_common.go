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
	"terraform-provider-ndfc/internal/provider/datasources/datasource_fabric"
	"terraform-provider-ndfc/internal/provider/ndfc/api"
	"terraform-provider-ndfc/internal/provider/resources/resource_fabric_common"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const (
	ResourceFabrics          = "fabric"
	ResourceVxlanEvpnFabric  = "fabric_vxlan_evpn"
	ResourceVxlanMsdFabric   = "fabric_vxlan_msd"
	ResourceLanClassicFabric = "fabric_lan_classic"
	ResourceIsnFabric        = "fabric_msite_ext_net"
	ResourceIpfmFabric       = "fabric_ipfm"
	ResourceIpfmFabricType   = "Easy_Fabric_IPFM"
	ResourceIsnFabricType    = "External_Fabric"
	ResourceLanClassicType   = "LAN_Classic"
	ResourceVxlanEvpnType    = "Easy_Fabric"
	ResourceVxlanMsdType     = "MSD_Fabric"
)

func (f *NDFC) RscReadFabric(ctx context.Context, dg *diag.Diagnostics, tf resource_fabric_common.FabricModel, fType string) {
	tflog.Info(ctx, "Read Fabric")
	var nvPairsModel resource_fabric_common.NdfcFabricPayload
	model := tf.GetModelData()
	fapi, _ := f.RscGetFabricApiDetails(ctx, dg, model, fType)
	if dg.HasError() {
		return
	}

	payload, err := fapi.Get()
	if len(payload) == 0 {
		if err == nil {
			err = fmt.Errorf("fabric not found in NDFC")
		}
		tflog.Error(ctx, "RscReadFabric: Failed to get fabric")
		dg.AddError("Failed to get fabric", fmt.Sprintf("Error: %q", err.Error()))
		tf.SetModelData(&nvPairsModel.NdfcFabricNvPairs)
		return
	}
	tflog.Debug(ctx, fmt.Sprintf("RscReadFabric: payload %s", string(payload)))
	err = json.Unmarshal(payload, &nvPairsModel)
	if err != nil {
		tflog.Error(ctx, "RscReadFabric: Failed to unmarshal Fabric data")
		dg.AddError("Failed to unmarshal fabric data", fmt.Sprintf("Error: %q", err.Error()))
		return
	}
	tflog.Debug(ctx, fmt.Sprintf("RscReadFabric: nvPairsModel %v", nvPairsModel))
	tf.SetModelData(&nvPairsModel.NdfcFabricNvPairs)
}

func (f *NDFC) RscCreateFabric(ctx context.Context, dg *diag.Diagnostics, tf resource_fabric_common.FabricModel, fType string) {
	tflog.Info(ctx, "Create Fabric")
	model := tf.GetModelData()
	fapi, payload := f.RscGetFabricApiDetails(ctx, dg, model, fType)
	if dg.HasError() {
		return
	}

	ret, _ := fapi.Get()
	tflog.Debug(ctx, fmt.Sprintf("RscCreateFabric: response %s", ret))
	if len(ret) > 0 && string(ret) != "[]" && string(ret) != "" {
		tflog.Error(ctx, "RscCreateFabric: Fabric already exists")
		dg.AddError("Fabric already exists", fmt.Sprintf("Fabric %s already exists", model.FabricName))
		return
	}

	resp, err := fapi.Post(payload)
	if err != nil {
		tflog.Error(ctx, "RscCreateFabric: POST failed with payload %s", map[string]interface{}{"Payload": payload})
		dg.AddError("Failed to create fabric", fmt.Sprintf("Error: %q Response :%s", err.Error(), resp.String()))
		return
	}

	if model.Deploy {
		err = fmt.Errorf("no switches found in the fabric for deployment")
		dg.AddWarning("Fabric created but not deployed", fmt.Sprintf("Reason: %q", err.Error()))
	}
	f.RscReadFabric(ctx, dg, tf, fType)

}
func (f *NDFC) RscUpdateFabric(ctx context.Context, dg *diag.Diagnostics, tf resource_fabric_common.FabricModel, fType string) {
	tflog.Info(ctx, "Update Fabric")
	model := tf.GetModelData()
	fabricName := model.FabricName
	deploy := model.Deploy
	fapi, payload := f.RscGetFabricApiDetails(ctx, dg, model, fType)
	if dg.HasError() {
		return
	}

	res, err := fapi.Put(payload)
	if err != nil {
		tflog.Error(ctx, "RscUpdateFabric: PUT failed with payload %s", map[string]interface{}{"Payload": payload})
		dg.AddError("Failed to update fabric", fmt.Sprintf("Error: %q %q", err.Error(), res.String()))
		return
	}
	f.RscDeployFabric(ctx, dg, fabricName, deploy)
	if dg.HasError() {
		return
	}
	f.RscReadFabric(ctx, dg, tf, fType)
}
func (f *NDFC) RscDeleteFabric(ctx context.Context, dg *diag.Diagnostics, tf resource_fabric_common.FabricModel, fType string) {
	tflog.Info(ctx, "Delete Fabric")
	model := tf.GetModelData()
	fapi, _ := f.RscGetFabricApiDetails(ctx, dg, model, fType)
	if dg.HasError() {
		tflog.Error(ctx, "RscDeleteFabric: Failed to get fabric api details")
		return
	}
	tflog.Debug(ctx, fmt.Sprintf("RscDeleteFabric: fapi %v", fapi))
	/* Check if switches are present in the fabric before deleting the fabric
	   and throw an error if switches are present */

	payload, err := f.GetSwitchesInFabric(ctx, model.FabricName)
	tflog.Debug(ctx, fmt.Sprintf("RscDeployFabric: payload %s", string(payload)))
	if !(len(payload) == 0 || string(payload) == "[]") {
		if err == nil {
			err = fmt.Errorf("switches in fabric needs to be deleted before deleting the fabric")
		}
		tflog.Error(ctx, "RscDeployFabric: Failed to get switches in fabric")
		dg.AddError("Fabric updated but not deployed", fmt.Sprintf("Reason: %q", err.Error()))
		return
	}
	res, err := fapi.Delete()
	if err != nil {
		tflog.Error(ctx, "RscDeleteFabric: DELETE failed")
		dg.AddError("Failed to delete fabric", fmt.Sprintf("Error: %s %s", err.Error(), res.String()))
		return
	}
	tflog.Info(ctx, "RscDeleteFabric: Fabric delete initiated")
	time.Sleep(3 * time.Second) // Wait for DB to sync
	payload, err = fapi.Get()
	tflog.Debug(ctx, fmt.Sprintf("RscDeleteFabric: payload %s", string(payload)))
	if err != nil || len(payload) == 0 || string(payload) == "[]" {
		tflog.Info(ctx, "RscDeleteFabric: Fabric deleted")
		/* Fabric is deleted. */
		return
	} else {
		dg.AddError("Fabric delete failed", fmt.Sprintf("Fabric %s is not deleted", model.FabricName))
	}
}
func (f NDFC) RscImportFabric(ctx context.Context, dg *diag.Diagnostics, tf resource_fabric_common.FabricModel, fType string) {
	tflog.Info(ctx, "Import  Fabric")
	f.RscReadFabric(ctx, dg, tf, fType)
}

func (f NDFC) RscGetFabricApiDetails(ctx context.Context, dg *diag.Diagnostics, model *resource_fabric_common.NDFCFabricCommonModel, fType string) (*api.FabricAPI, []byte) {
	tflog.Info(ctx, "Retrieve payload and fabric api object")
	fapi := api.NewFabricAPI(f.GetLock(ResourceFabrics), &f.apiClient)
	fapi.FabricType = fType
	fapi.FabricName = model.FabricName
	payload, err := json.Marshal(model)
	if err != nil {
		tflog.Error(ctx, "RscGetFabricApiDetails: Failed to marshal fabric data")
		dg.AddError("Failed to marshal fabric data", fmt.Sprintf("Error: %q", err.Error()))
		return nil, nil
	}
	return fapi, payload
}
func (f *NDFC) RscDeployFabric(ctx context.Context, dg *diag.Diagnostics, fabricName string, deploy bool) {
	if deploy {
		payload, err := f.GetSwitchesInFabric(ctx, fabricName)
		tflog.Debug(ctx, fmt.Sprintf("RscDeployFabric: payload %s", string(payload)))
		if len(payload) == 0 || string(payload) == "[]" {
			if err == nil {
				err = fmt.Errorf("no switches found in the fabric for deployment")
			}
			tflog.Error(ctx, "RscDeployFabric: Failed to get switches in fabric")
			dg.AddWarning("Fabric not deployed", fmt.Sprintf("Reason: %q", err.Error()))
			return
		}
		f.RecalculateAndDeploy(ctx, dg, fabricName, true, deploy, nil)
		if dg.HasError() {
			return
		}
	}
}
func (f *NDFC) GetSwitchesInFabric(ctx context.Context, fabricName string) ([]byte, error) {
	fapi := api.NewFabricAPI(f.GetLock(ResourceFabrics), &f.apiClient)
	fapi.FabricName = fabricName
	fapi.GetSwitchesInFabric = true
	return fapi.Get()
}
func (f *NDFC) GetFabricName(ctx context.Context, serialNumber string) string {
	fapi := api.NewFabricAPI(f.GetLock(ResourceFabrics), &f.apiClient)
	fapi.Serialnumber = serialNumber
	payload, err := fapi.Get()
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("GetFabricName: Failed to get fabric name for serial number %s", serialNumber))
		return ""
	}
	var FabricNamePayload resource_fabric_common.NdfcFabricNamePayload
	err = json.Unmarshal(payload, &FabricNamePayload)
	if err != nil {
		return ""
	}
	return FabricNamePayload.FabricName

}
func (f *NDFC) DSGetFabric(ctx context.Context, dg *diag.Diagnostics, fabricName string) *datasource_fabric.FabricModel {
	tflog.Debug(ctx, "DSGetFabricBulk entry")
	fapi := api.NewFabricAPI(f.GetLock(ResourceFabrics), &f.apiClient)
	fapi.FabricName = fabricName
	res, err := fapi.Get()
	if err != nil {
		dg.AddError("Get failed", err.Error())
		return nil
	} else if string(res) == "[]" || len(res) == 0 {
		dg.AddError("Fabric not found", "Fabric not found")
		return nil
	} else {
		tflog.Info(ctx, "Url:"+f.url+" Read success ")
		tflog.Info(ctx, string(res))
	}
	ndFabric := datasource_fabric.NDFCFabricDataSourceModel{}
	log.Printf("[DEBUG] DSGetFabric: res %v", string(res))
	err = json.Unmarshal(res, &ndFabric)
	if err != nil {
		dg.AddError("datasource_fabric: unmarshal failed ", err.Error())
		return nil
	} else {
		tflog.Debug(ctx, "datasource_fabric: Unmarshal OK")
	}
	log.Printf("[DEBUG] DSGetFabric: ndFabric %v", ndFabric)
	data := new(datasource_fabric.FabricModel)
	d := data.SetModelData(&ndFabric.NvPairs)
	if d != nil {
		*dg = d
		return nil
	} else {
		tflog.Debug(ctx, fmt.Sprintf("DSGetFabric: SetModelData OK for fabric %s", data.FabricName.ValueString()))
	}
	return data
}
