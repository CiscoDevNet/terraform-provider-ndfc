//go:build ignore
// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

//template:begin imports
import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nd"
	"github.com/netascode/terraform-provider-ndfc/internal/provider/helpers"
)
//template:end imports

//template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &{{camelCase .Name}}DataSource{}
	_ datasource.DataSourceWithConfigure = &{{camelCase .Name}}DataSource{}
)

func New{{camelCase .Name}}DataSource() datasource.DataSource {
	return &{{camelCase .Name}}DataSource{}
}

type {{camelCase .Name}}DataSource struct {
	client *nd.Client
}

func (d *{{camelCase .Name}}DataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_{{snakeCase .Name}}"
}

func (d *{{camelCase .Name}}DataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "{{.DsDescription}}",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			{{- range  .Attributes}}
			{{- if not .Value}}
			"{{.TfName}}": schema.{{if or (eq .Type "List") (eq .Type "Set")}}{{.Type}}Nested{{else if eq .Type "ListString"}}List{{else}}{{.Type}}{{end}}Attribute{
				MarkdownDescription: "{{.Description}}",
				{{- if eq .Type "ListString"}}
				ElementType:         types.StringType,
				{{- end}}
				{{- if or .Id .Reference}}
				Required:            true,
				{{- else}}
				Computed:            true,
				{{- end}}
				{{- if or (eq .Type "List") (eq .Type "Set")}}
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						{{- range  .Attributes}}
						{{- if not .Value}}
						"{{.TfName}}": schema.{{if or (eq .Type "List") (eq .Type "Set")}}{{.Type}}Nested{{else if eq .Type "ListString"}}List{{else}}{{.Type}}{{end}}Attribute{
							MarkdownDescription: "{{.Description}}",
							{{- if eq .Type "ListString"}}
							ElementType:         types.StringType,
							{{- end}}
							{{- if or .Id .Reference}}
							Required:            true,
							{{- else}}
							Computed:            true,
							{{- end}}
							{{- if or (eq .Type "List") (eq .Type "Set")}}
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									{{- range  .Attributes}}
									{{- if not .Value}}
									"{{.TfName}}": schema.{{if or (eq .Type "List") (eq .Type "Set")}}{{.Type}}Nested{{else if eq .Type "ListString"}}List{{else}}{{.Type}}{{end}}Attribute{
										MarkdownDescription: "{{.Description}}",
										{{- if eq .Type "ListString"}}
										ElementType:         types.StringType,
										{{- end}}
										{{- if or .Id .Reference}}
										Required:            true,
										{{- else}}
										Computed:            true,
										{{- end}}
										{{- if or (eq .Type "List") (eq .Type "Set")}}
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												{{- range  .Attributes}}
												{{- if not .Value}}
												"{{.TfName}}": schema.{{if or (eq .Type "List") (eq .Type "Set")}}{{.Type}}Nested{{else if eq .Type "ListString"}}List{{else}}{{.Type}}{{end}}Attribute{
													MarkdownDescription: "{{.Description}}",
													{{- if eq .Type "ListString"}}
													ElementType:         types.StringType,
													{{- end}}
													{{- if or .Id .Reference}}
													Required:            true,
													{{- else}}
													Computed:            true,
													{{- end}}
												},
												{{- end}}
												{{- end}}
											},
										},
										{{- end}}
									},
									{{- end}}
									{{- end}}
								},
							},
							{{- end}}
						},
						{{- end}}
						{{- end}}
					},
				},
				{{- end}}
			},
			{{- end}}
			{{- end}}
		},
	}
}

func (d *{{camelCase .Name}}DataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*NdfcProviderData).Client
}
//template:end model

//template:begin read
func (d *{{camelCase .Name}}DataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config {{camelCase .Name}}

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	res, err := d.client.Get(fmt.Sprintf("%v%v", config.getPath(), {{range .Attributes}}{{if .Id}}config.{{toGoName .TfName}}.Value{{.Type}}(){{end}}{{end}}))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)
	config.Id = types.StringValue({{$first := true}}{{range .Attributes}}{{if or .Id .Reference}}{{if not $first}}+"/"+{{end}}{{$first = false}}config.{{toGoName .TfName}}.Value{{.Type}}(){{end}}{{end}})

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
//template:end read
