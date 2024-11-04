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
	"terraform-provider-ndfc/internal/provider/datasources/datasource_networks"
	"terraform-provider-ndfc/internal/provider/ndfc"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*networksDatasource)(nil)

func NewNetworksDataSource() datasource.DataSource {
	return &networksDatasource{}
}

type networksDatasource struct {
	client *ndfc.NDFC
}

func (r *networksDatasource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + ndfc.ResourceNetworks
}

func (r *networksDatasource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datasource_networks.NetworksDataSourceSchema(ctx)
	//resp.Schema.Attributes["new_attribute"] = schema.BoolAttribute{Optional: true}
}

func (d *networksDatasource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	tflog.Info(ctx, "networks datasource Configure")
	client, ok := req.ProviderData.(*ndfc.NDFC)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected datasource  Configure Type",
			fmt.Sprintf("Expected *nd.NDFC, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}
	d.client = client
}

func (r *networksDatasource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data datasource_networks.NetworksModel

	// Read Terraform config data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	if data.FabricName.IsNull() || data.FabricName.IsUnknown() {
		resp.Diagnostics.AddError("FabricName cannot be empty", "should be present")
		return
	}

	dd := r.client.DsGetNetworks(ctx, &resp.Diagnostics, data.FabricName.ValueString())
	if dd == nil {
		tflog.Error(ctx, "Read Networks Failed")
		//resp.Diagnostics.AddError("Read Failure", "No data received from NDFC")

	}
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, dd)...)

}
