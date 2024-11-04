// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"terraform-provider-ndfc/internal/provider/datasources/datasource_vrf_bulk"
	"terraform-provider-ndfc/internal/provider/ndfc"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*vrfBulkDataSource)(nil)

func NewVrfBulkDataSource() datasource.DataSource {
	return &vrfBulkDataSource{}
}

type vrfBulkDataSource struct {
	//version types.String
	client *ndfc.NDFC
}

func (d *vrfBulkDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "ndfc_vrf_bulk"
	tflog.Info(ctx, "DS Metadata called", map[string]interface{}{"ProTyNam": req.ProviderTypeName})
}

func (d *vrfBulkDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datasource_vrf_bulk.VrfBulkDataSourceSchema(ctx)
	tflog.Info(ctx, "DS Schema called")
}

func (d *vrfBulkDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	tflog.Info(ctx, "VRF Data source Configure")

	client, ok := req.ProviderData.(*ndfc.NDFC)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *nd.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *vrfBulkDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data datasource_vrf_bulk.VrfBulkModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	var vrf *datasource_vrf_bulk.VrfBulkModel
	url_path := ""
	if !data.FabricName.IsNull() && !data.FabricName.IsUnknown() {
		vrf = d.client.DSGetBulkVrf(ctx, &resp.Diagnostics, data.FabricName.ValueString())
		if vrf == nil {
			return
		}
		tflog.Info(ctx, "Url:"+url_path)
	} else {
		resp.Diagnostics.AddError("Fabric name is needed to read data", "Fabric name is needed to read data")
		return
	}
	// Example data value setting
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, vrf)...)
}
