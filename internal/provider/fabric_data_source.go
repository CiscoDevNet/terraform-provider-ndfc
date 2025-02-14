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
	"terraform-provider-ndfc/internal/provider/datasources/datasource_fabric"
	"terraform-provider-ndfc/internal/provider/ndfc"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*fabricDataSource)(nil)

func NewFabricDataSource() datasource.DataSource {
	return &fabricDataSource{}
}

type fabricDataSource struct {
	client *ndfc.NDFC
}

func (d *fabricDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + ndfc.DataSourceFabric
}

func (d *fabricDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datasource_fabric.FabricDataSourceSchema(ctx)
}

func (d *fabricDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	tflog.Info(ctx, "Fabric Data source Configure")

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
func (d *fabricDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data datasource_fabric.FabricModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	fb := d.client.DSGetFabricBulk(ctx, &resp.Diagnostics)
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, fb)...)
}
