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
	"terraform-provider-ndfc/internal/provider/datasources/datasource_interfaces"
	"terraform-provider-ndfc/internal/provider/ndfc"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*interfacesDatasource)(nil)

func NewInterfacesDataSource() datasource.DataSource {
	return &interfacesDatasource{}
}

type interfacesDatasource struct {
	client *ndfc.NDFC
}

func (r *interfacesDatasource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + ndfc.ResourceInterfaces
}

func (r *interfacesDatasource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datasource_interfaces.InterfacesDataSourceSchema(ctx)
	//resp.Schema.Attributes["new_attribute"] = schema.BoolAttribute{Optional: true}
}

func (d *interfacesDatasource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	tflog.Info(ctx, "interfaces datasource Configure")
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

func (r *interfacesDatasource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data datasource_interfaces.InterfacesModel

	// Read Terraform config data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	if data.SerialNumber.IsNull() || data.SerialNumber.IsUnknown() {
		resp.Diagnostics.AddError("SerialNumber cannot be empty", "should be present")
		return
	}
	inData := data.GetModelData()
	r.client.DsGetInterfaces(ctx, &resp.Diagnostics, inData)
	if resp.Diagnostics.HasError() {
		return
	}
	data.SetModelData(inData)
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)

}
