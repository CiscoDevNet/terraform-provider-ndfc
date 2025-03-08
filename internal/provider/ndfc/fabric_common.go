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
	"terraform-provider-ndfc/internal/provider/datasources/datasource_fabric"
	"terraform-provider-ndfc/internal/provider/ndfc/api"
	"terraform-provider-ndfc/internal/provider/resources/resource_fabric_common"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const (
	DataSourceFabric         = "fabric_data_source"
	ResourceFabrics          = "fabrics"
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

func (c *NDFC) RscReadFabric(ctx context.Context, dg *diag.Diagnostics, tf resource_fabric_common.FabricModel, fabricType string) {
	tflog.Info(ctx, "Read Fabric")
	var nvPairsModel resource_fabric_common.NdfcFabricPayload
	ndfcFabricModel := tf.GetModelData()
	fabricApi, _ := c.RscGetFabricApiDetails(ctx, dg, ndfcFabricModel, fabricType)
	if dg.HasError() {
		return
	}
	fabricApi.FabricName = ndfcFabricModel.FabricName
	payload, err := fabricApi.Get()
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

func (c *NDFC) RscCreateFabric(ctx context.Context, dg *diag.Diagnostics, tf resource_fabric_common.FabricModel, fabricType string) {
	tflog.Info(ctx, "Create Fabric")
	ndfcFabricModel := tf.GetModelData()
	fabricApi, payload := c.RscGetFabricApiDetails(ctx, dg, ndfcFabricModel, fabricType)
	if dg.HasError() {
		return
	}

	fabricApi.FabricName = ndfcFabricModel.FabricName
	ret, _ := fabricApi.Get()
	tflog.Debug(ctx, fmt.Sprintf("RscCreateFabric: response %s", ret))
	if len(ret) != 0 {
		tflog.Error(ctx, "RscCreateFabric: Fabric already exists")
		dg.AddError("Fabric already exists", fmt.Sprintf("Fabric with Name: %s already exists", ndfcFabricModel.FabricName))
		return
	}

	_, err := fabricApi.Post(payload)
	if err != nil {
		tflog.Error(ctx, "RscCreateFabric: POST failed with payload %s", map[string]interface{}{"Payload": payload})
		dg.AddError("Failed to create fabric", fmt.Sprintf("Error: %q", err.Error()))
		return
	}

	if ndfcFabricModel.Deploy {
		err = fmt.Errorf("no switches found in the fabric for deployment")
		dg.AddWarning("Fabric created but not deployed", fmt.Sprintf("Reason: %q", err.Error()))
	}
	c.RscReadFabric(ctx, dg, tf, fabricType)

}
func (c *NDFC) RscUpdateFabric(ctx context.Context, dg *diag.Diagnostics, tf resource_fabric_common.FabricModel, fabricType string) {
	tflog.Info(ctx, "Update Fabric")
	ndfcFabricModel := tf.GetModelData()
	fabricName := ndfcFabricModel.FabricName
	deploy := ndfcFabricModel.Deploy
	fabricApi, payload := c.RscGetFabricApiDetails(ctx, dg, ndfcFabricModel, fabricType)
	if dg.HasError() {
		return
	}
	fabricApi.FabricName = fabricName
	res, err := fabricApi.Put(payload)
	if err != nil {
		tflog.Error(ctx, "RscUpdateFabric: PUT failed with payload %s", map[string]interface{}{"Payload": payload})
		dg.AddError("Failed to update fabric", fmt.Sprintf("Error: %q %q", err.Error(), res.String()))
		return
	}
	c.RscDeployFabric(ctx, dg, fabricName, deploy)
	if dg.HasError() {
		return
	}
	c.RscReadFabric(ctx, dg, tf, fabricType)
}
func (c *NDFC) RscDeleteFabric(ctx context.Context, dg *diag.Diagnostics, tf resource_fabric_common.FabricModel, fabricType string) {
	tflog.Info(ctx, "Delete Fabric")
	ndfcFabricModel := tf.GetModelData()
	fabricApi, _ := c.RscGetFabricApiDetails(ctx, dg, ndfcFabricModel, fabricType)
	if dg.HasError() {
		tflog.Error(ctx, "RscDeleteFabric: Failed to get fabric api details")
		return
	}
	tflog.Debug(ctx, fmt.Sprintf("RscDeleteFabric: fabricApi %v", fabricApi))
	/* Check if switches are present in the fabric before deleting the fabric
	   and throw an error if switches are present */
	fabricApi.FabricName = ndfcFabricModel.FabricName
	fabricApi.GetSwitchesInFabric = true
	payload, err := fabricApi.Get()
	tflog.Debug(ctx, fmt.Sprintf("RscDeployFabric: payload %s", string(payload)))
	if !(len(payload) == 0 || string(payload) == "[]") {
		if err == nil {
			err = fmt.Errorf("switches in fabric needs to be deleted before deleting the fabric")
		}
		tflog.Error(ctx, "RscDeployFabric: Failed to get switches in fabric")
		dg.AddError("Fabric updated but not deployed", fmt.Sprintf("Reason: %q", err.Error()))
		return
	}
	res, err := fabricApi.Delete()
	if err != nil {
		tflog.Error(ctx, "RscDeleteFabric: DELETE failed")
		dg.AddError("Failed to delete fabric", fmt.Sprintf("Error: %s %s", err.Error(), res.String()))
		return
	}
	tflog.Info(ctx, "RscDeleteFabric: Fabric delete initiated")
	time.Sleep(3 * time.Second) // Wait for DB to sync
	_, err = fabricApi.Get()
	if err != nil {
		tflog.Info(ctx, "RscDeleteFabric: Fabric deleted")
		/* Fabric is deleted. */
		return
	} else {
		dg.AddError("Fabric delete failed", fmt.Sprintf("Fabric with Name: %s is not deleted", ndfcFabricModel.FabricName))
	}
}
func (c NDFC) RscImportFabric(ctx context.Context, dg *diag.Diagnostics, tf resource_fabric_common.FabricModel, fabricType string) {
	tflog.Info(ctx, "Import  Fabric")
	c.RscReadFabric(ctx, dg, tf, fabricType)
}

func (c NDFC) RscGetFabricApiDetails(ctx context.Context, dg *diag.Diagnostics, ndfcFabricModel *resource_fabric_common.NDFCFabricCommonModel, fabricType string) (*api.FabricAPI, []byte) {
	tflog.Info(ctx, "Retrieve payload and fabric api object")
	fabricApi := api.NewFabricAPI(c.GetLock(ResourceFabrics), &c.apiClient)
	fabricApi.FabricType = fabricType
	payload, err := json.Marshal(ndfcFabricModel)
	if err != nil {
		tflog.Error(ctx, "RscGetFabricApiDetails: Failed to marshal fabric data")
		dg.AddError("Failed to marshal fabric data", fmt.Sprintf("Error: %q", err.Error()))
		return nil, nil
	}
	return fabricApi, payload
}
func (c *NDFC) RscDeployFabric(ctx context.Context, dg *diag.Diagnostics, fabricName string, deploy bool) {
	if deploy {
		payload, err := c.GetSwitchesInFabric(ctx, fabricName)
		tflog.Debug(ctx, fmt.Sprintf("RscDeployFabric: payload %s", string(payload)))
		if len(payload) == 0 || string(payload) == "[]" {
			if err == nil {
				err = fmt.Errorf("no switches found in the fabric for deployment")
			}
			tflog.Error(ctx, "RscDeployFabric: Failed to get switches in fabric")
			dg.AddWarning("Fabric not deployed", fmt.Sprintf("Reason: %q", err.Error()))
			return
		}
		c.RecalculateAndDeploy(ctx, dg, fabricName, true, deploy, nil)
		if dg.HasError() {
			return
		}
	}
}
func (c *NDFC) GetSwitchesInFabric(ctx context.Context, fabricName string) ([]byte, error) {
	fabricApi := api.NewFabricAPI(c.GetLock(ResourceFabrics), &c.apiClient)
	fabricApi.FabricName = fabricName
	fabricApi.GetSwitchesInFabric = true
	return fabricApi.Get()
}
func (c *NDFC) GetFabricName(ctx context.Context, serialNumber string) string {
	fabricApi := api.NewFabricAPI(c.GetLock(ResourceFabrics), &c.apiClient)
	fabricApi.Serialnumber = serialNumber
	payload, err := fabricApi.Get()
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
func (c *NDFC) DSGetFabricBulk(ctx context.Context, dg *diag.Diagnostics) *datasource_fabric.FabricModel {
	tflog.Debug(ctx, "DSGetFabricBulk entry")
	fabricApi := api.NewFabricAPI(c.GetLock(ResourceFabrics), &c.apiClient)
	res, err := fabricApi.Get()
	if err != nil {
		dg.AddError("Get failed", err.Error())
		return nil

	} else {
		tflog.Info(ctx, "Url:"+c.url+" Read success ")
		tflog.Info(ctx, string(res))
	}
	ndFabric := datasource_fabric.NDFCFabricModel{}
	err = json.Unmarshal(res, &ndFabric.Fabrics)
	if err != nil {
		dg.AddError("datasource_fabric: unmarshal failed ", err.Error())
		return nil
	} else {
		tflog.Debug(ctx, "datasource_fabric: Unmarshal OK")
	}
	if len(ndFabric.Fabrics) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("datasource_fabric: Retrieved %d fabrics", len(ndFabric.Fabrics)))
	}
	data := new(datasource_fabric.FabricModel)
	d := data.SetModelData(&ndFabric)
	if d != nil {
		*dg = d
		return nil
	} else {
		tflog.Debug(ctx, "datasource_vrf_bulk: SetModelData OK")
	}
	return data
}