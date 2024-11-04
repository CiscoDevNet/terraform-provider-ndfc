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
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const UrlFabricGetAll = "/lan-fabric/rest/control/fabrics"

func (c NDFC) DSGetFabricBulk(ctx context.Context, dg *diag.Diagnostics) *datasource_fabric.FabricModel {
	//tflog.Debug(ctx, "DSGetFabricBulk entry")
	log.Printf("DSGetFabricBulk entry")
	time.Sleep(5 * time.Second)
	res, err := c.apiClient.GetRawJson(UrlFabricGetAll)

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
