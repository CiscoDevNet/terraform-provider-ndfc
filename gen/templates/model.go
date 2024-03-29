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
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/terraform-provider-ndfc/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)
//template:end imports

//template:begin types
{{- $name := camelCase .Name}}
type {{camelCase .Name}} struct {
	Id types.String `tfsdk:"id"`
{{- range .Attributes}}
{{- if not .Value}}
{{- if or (eq .Type "List") (eq .Type "Set")}}
	{{toGoName .TfName}} []{{$name}}{{toGoName .TfName}} `tfsdk:"{{.TfName}}"`
{{- else if or (eq .Type "ListString") (eq .Type "Versions")}}
	{{toGoName .TfName}} types.List `tfsdk:"{{.TfName}}"`
{{- else if eq .Type "Version"}}
	{{toGoName .TfName}} types.Int64 `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
{{- end}}
}

{{ range .Attributes}}
{{- if not .Value}}
{{- $childName := toGoName .TfName}}
{{- if or (eq .Type "List") (eq .Type "Set")}}
type {{$name}}{{toGoName .TfName}} struct {
{{- range .Attributes}}
{{- if not .Value}}
{{- if or (eq .Type "List") (eq .Type "Set")}}
	{{toGoName .TfName}} []{{$name}}{{$childName}}{{toGoName .TfName}} `tfsdk:"{{.TfName}}"`
{{- else if or (eq .Type "ListString") (eq .Type "Versions")}}
	{{toGoName .TfName}} types.List `tfsdk:"{{.TfName}}"`
{{- else if eq .Type "Version"}}
	{{toGoName .TfName}} types.Int64 `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
{{- end}}
}
{{- end}}
{{- end}}
{{ end}}

{{ range .Attributes}}
{{- if not .Value}}
{{- $childName := toGoName .TfName}}
{{- if or (eq .Type "List") (eq .Type "Set")}}
{{ range .Attributes}}
{{- if not .Value}}
{{- $childChildName := toGoName .TfName}}
{{- if or (eq .Type "List") (eq .Type "Set")}}
type {{$name}}{{$childName}}{{toGoName .TfName}} struct {
{{- range .Attributes}}
{{- if not .Value}}
{{- if or (eq .Type "List") (eq .Type "Set")}}
	{{toGoName .TfName}} []{{$name}}{{$childName}}{{$childChildName}}{{toGoName .TfName}} `tfsdk:"{{.TfName}}"`
{{- else if or (eq .Type "ListString") (eq .Type "Versions")}}
	{{toGoName .TfName}} types.List `tfsdk:"{{.TfName}}"`
{{- else if eq .Type "Version"}}
	{{toGoName .TfName}} types.Int64 `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
{{- end}}
}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{ end}}

{{ range .Attributes}}
{{- if not .Value}}
{{- $childName := toGoName .TfName}}
{{- if or (eq .Type "List") (eq .Type "Set")}}
{{ range .Attributes}}
{{- if not .Value}}
{{- $childChildName := toGoName .TfName}}
{{- if or (eq .Type "List") (eq .Type "Set")}}
{{ range .Attributes}}
{{- if not .Value}}
{{- if or (eq .Type "List") (eq .Type "Set")}}
type {{$name}}{{$childName}}{{$childChildName}}{{toGoName .TfName}} struct {
{{- range .Attributes}}
{{- if not .Value}}
{{- if or (eq .Type "ListString") (eq .Type "Versions")}}
	{{toGoName .TfName}} types.List `tfsdk:"{{.TfName}}"`
{{- else if eq .Type "Version"}}
	{{toGoName .TfName}} types.Int64 `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
{{- end}}
}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{ end}}
//template:end types

//template:begin getPath
func (data {{camelCase .Name}}) getPath() string {
{{- if hasReference .Attributes}}
	return fmt.Sprintf("{{.RestEndpoint}}"{{range .Attributes}}{{if .Reference}}, url.QueryEscape(fmt.Sprintf("%v", data.{{toGoName .TfName}}.Value{{.Type}}())){{end}}{{end}})
{{- else}}
	return "{{.RestEndpoint}}"
{{- end}}
}
//template:end getPath

//template:begin toBody
func (data {{camelCase .Name}}) toBody(ctx context.Context) string {
	body := ""
	{{- range .Attributes}}
	{{- if .Value}}
	body, _ = sjson.Set(body, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
	{{- else if not .TfOnly}}
	{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64")}}
	if !data.{{toGoName .TfName}}.IsNull() && !data.{{toGoName .TfName}}.IsUnknown() {
		body, _ = sjson.Set(body, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if .ModelTypeString}}fmt.Sprint({{end}}data.{{toGoName .TfName}}.Value{{.Type}}(){{if .ModelTypeString}}){{end}})
	}
	{{- else if eq .Type "Bool"}}
	if !data.{{toGoName .TfName}}.IsNull() && !data.{{toGoName .TfName}}.IsUnknown() {
		body, _ = sjson.Set(body, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if .ModelTypeString}}fmt.Sprint({{end}}data.{{toGoName .TfName}}.Value{{.Type}}(){{if .ModelTypeString}}){{end}})
	}
	{{- else if eq .Type "ListString"}}
	if !data.{{toGoName .TfName}}.IsNull() && !data.{{toGoName .TfName}}.IsUnknown() {
		var values []string
		data.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", values)
	}
	{{- else if or (eq .Type "List") (eq .Type "Set")}}
	if len(data.{{toGoName .TfName}}) > 0 {
		body, _ = sjson.Set(body, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", []interface{}{})
		for _, item := range data.{{toGoName .TfName}} {
			itemBody := ""
			{{- range .Attributes}}
			{{- if .Value}}
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
			{{- else if not .TfOnly}}
			{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64")}}
			if !item.{{toGoName .TfName}}.IsNull() && !item.{{toGoName .TfName}}.IsUnknown() {
				itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if .ModelTypeString}}fmt.Sprint({{end}}item.{{toGoName .TfName}}.Value{{.Type}}(){{if .ModelTypeString}}){{end}})
			}
			{{- else if eq .Type "Bool"}}
			if !item.{{toGoName .TfName}}.IsNull() && !item.{{toGoName .TfName}}.IsUnknown() {
				itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if .ModelTypeString}}fmt.Sprint({{end}}item.{{toGoName .TfName}}.Value{{.Type}}(){{if .ModelTypeString}}){{end}})
			}
			{{- else if eq .Type "ListString"}}
			if !item.{{toGoName .TfName}}.IsNull() && !item.{{toGoName .TfName}}.IsUnknown() {
				var values []string
				item.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", values)
			}
			{{- else if or (eq .Type "List") (eq .Type "Set")}}
			if len(item.{{toGoName .TfName}}) > 0 {
				itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", []interface{}{})
				for _, childItem := range item.{{toGoName .TfName}} {
					itemChildBody := ""
					{{- range .Attributes}}
					{{- if .Value}}
					itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
					{{- else if not .TfOnly}}
					{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64")}}
					if !childItem.{{toGoName .TfName}}.IsNull() && !childItem.{{toGoName .TfName}}.IsUnknown() {
						itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if .ModelTypeString}}fmt.Sprint({{end}}childItem.{{toGoName .TfName}}.Value{{.Type}}(){{if .ModelTypeString}}){{end}})
					}
					{{- else if eq .Type "Bool"}}
					if !childItem.{{toGoName .TfName}}.IsNull() && !childItem.{{toGoName .TfName}}.IsUnknown() {
						itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if .ModelTypeString}}fmt.Sprint({{end}}childItem.{{toGoName .TfName}}.Value{{.Type}}(){{if .ModelTypeString}}){{end}})
					}
					{{- else if eq .Type "ListString"}}
					if !childItem.{{toGoName .TfName}}.IsNull() && !childItem.{{toGoName .TfName}}.IsUnknown() {
						var values []string
						childItem.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
						itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", values)
					}
					{{- else if or (eq .Type "List") (eq .Type "Set")}}
					if len(childItem.{{toGoName .TfName}}) > 0{{if ne .ConditionalAttribute.Name ""}} && childItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}} {
						itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", []interface{}{})
						for _, childChildItem := range childItem.{{toGoName .TfName}} {
							itemChildChildBody := ""
							{{- range .Attributes}}
							{{- if .Value}}
							itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
							{{- else if not .TfOnly}}
							{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64")}}
							if !childChildItem.{{toGoName .TfName}}.IsNull() && !childChildItem.{{toGoName .TfName}}.IsUnknown() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if .ModelTypeString}}fmt.Sprint({{end}}childChildItem.{{toGoName .TfName}}.Value{{.Type}}(){{if .ModelTypeString}}){{end}})
							}
							{{- else if eq .Type "Bool"}}
							if !childChildItem.{{toGoName .TfName}}.IsNull() && !childChildItem.{{toGoName .TfName}}.IsUnknown() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if .ModelTypeString}}fmt.Sprint({{end}}childChildItem.{{toGoName .TfName}}.Value{{.Type}}(){{if .ModelTypeString}}){{end}})
							}
							{{- else if eq .Type "ListString"}}
							if !childChildItem.{{toGoName .TfName}}.IsNull() && !childChildItem.{{toGoName .TfName}}.IsUnknown() {
								var values []string
								childChildItem.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", values)
							}
							{{- end}}
							{{- end}}
							{{- end}}
							itemChildBody, _ = sjson.SetRaw(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.-1", itemChildChildBody)
						}
					}
					{{- end}}
					{{- end}}
					{{- end}}
					itemBody, _ = sjson.SetRaw(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.-1", itemChildBody)
				}
			}
			{{- end}}
			{{- end}}
			{{- end}}
			body, _ = sjson.SetRaw(body, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.-1", itemBody)
		}
	}
	{{- end}}
	{{- end}}
	{{- end}}
	return body
}
//template:end toBody

//template:begin fromBody
func (data *{{camelCase .Name}}) fromBody(ctx context.Context, res gjson.Result) {
	{{- range .Attributes}}
	{{- if and (not .TfOnly) (not .Value)}}
	{{- $cname := toGoName .TfName}}
	{{- if eq .Type "String"}}
	if value := res.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); value.Exists() && value.String() != "" {
		data.{{toGoName .TfName}} = types.StringValue(value.String())
	} else {
		data.{{toGoName .TfName}} = types.StringNull()
	}
	{{- else if eq .Type "Int64"}}
	if value := res.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); value.Exists() && value.String() != "" {
		data.{{toGoName .TfName}} = types.Int64Value(value.Int())
	} else {
		data.{{toGoName .TfName}} = types.Int64Null()
	}
	{{- else if eq .Type "Float64"}}
	if value := res.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); value.Exists() && value.String() != "" {
		data.{{toGoName .TfName}} = types.Float64Value(value.Float())
	} else {
		data.{{toGoName .TfName}} = types.Float64Null()
	}
	{{- else if eq .Type "Bool"}}
	if value := res.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); value.Exists() && value.String() != "" {
		data.{{toGoName .TfName}} = types.BoolValue(value.Bool())
	} else {
		data.{{toGoName .TfName}} = types.BoolNull()
	}
	{{- else if eq .Type "ListString"}}
	if value := res.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); value.Exists() && value.String() != "" {
		data.{{toGoName .TfName}} = helpers.GetListString(value.Array())
	} else {
		data.{{toGoName .TfName}} = types.ListNull(types.StringType)
	}
	{{- else if or (eq .Type "List") (eq .Type "Set")}}
	if value := res.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); value.Exists() && value.String() != "" {
		data.{{toGoName .TfName}} = make([]{{$name}}{{toGoName .TfName}}, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := {{$name}}{{toGoName .TfName}}{}
			{{- range .Attributes}}
			{{- $ccname := toGoName .TfName}}
			{{- if and (not .TfOnly) (not .Value)}}
			{{- if eq .Type "String"}}
			if cValue := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); cValue.Exists() && cValue.String() != "" {
				item.{{toGoName .TfName}} = types.StringValue(cValue.String())
			} else {
				item.{{toGoName .TfName}} = types.StringNull()
			}
			{{- else if eq .Type "Int64"}}
			if cValue := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); cValue.Exists() && cValue.String() != "" {
				item.{{toGoName .TfName}} = types.Int64Value(cValue.Int())
			} else {
				item.{{toGoName .TfName}} = types.Int64Null()
			}
			{{- else if eq .Type "Float64"}}
			if cValue := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); cValue.Exists() && cValue.String() != "" {
				item.{{toGoName .TfName}} = types.Float64Value(cValue.Float())
			} else {
				item.{{toGoName .TfName}} = types.Float64Null()
			}
			{{- else if eq .Type "Bool"}}
			if cValue := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); cValue.Exists() && cValue.String() != "" {
				item.{{toGoName .TfName}} = types.BoolValue(cValue.Bool())
			} else {
				item.{{toGoName .TfName}} = types.BoolNull()
			}
			{{- else if eq .Type "ListString"}}
			if cValue := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); cValue.Exists() && cValue.String() != "" {
				item.{{toGoName .TfName}} = helpers.GetListString(cValue.Array())
			} else {
				item.{{toGoName .TfName}} = types.ListNull(types.StringType)
			}
			{{- else if or (eq .Type "List") (eq .Type "Set")}}
			if cValue := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); cValue.Exists() && cValue.String() != "" {
				item.{{toGoName .TfName}} = make([]{{$name}}{{$cname}}{{toGoName .TfName}}, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := {{$name}}{{$cname}}{{toGoName .TfName}}{}
					{{- range .Attributes}}
					{{- if and (not .TfOnly) (not .Value)}}
					{{- if eq .Type "String"}}
					if ccValue := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); ccValue.Exists() && ccValue.String() != "" {
						cItem.{{toGoName .TfName}} = types.StringValue(ccValue.String())
					} else {
						cItem.{{toGoName .TfName}} = types.StringNull()
					}
					{{- else if eq .Type "Int64"}}
					if ccValue := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); ccValue.Exists() && ccValue.String() != "" {
						cItem.{{toGoName .TfName}} = types.Int64Value(ccValue.Int())
					} else {
						cItem.{{toGoName .TfName}} = types.Int64Null()
					}
					{{- else if eq .Type "Float64"}}
					if ccValue := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); ccValue.Exists() && ccValue.String() != "" {
						cItem.{{toGoName .TfName}} = types.Float64Value(ccValue.Float())
					} else {
						cItem.{{toGoName .TfName}} = types.Float64Null()
					}
					{{- else if eq .Type "Bool"}}
					if ccValue := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); ccValue.Exists() && ccValue.String() != "" {
						cItem.{{toGoName .TfName}} = types.BoolValue(ccValue.Bool())
					} else {
						cItem.{{toGoName .TfName}} = types.BoolNull()
					}
					{{- else if eq .Type "ListString"}}
					if ccValue := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); ccValue.Exists() && ccValue.String() != "" {
						cItem.{{toGoName .TfName}} = helpers.GetListString(ccValue.Array())
					} else {
						cItem.{{toGoName .TfName}} = types.ListNull(types.StringType)
					}
					{{- else if or (eq .Type "List") (eq .Type "Set")}}
					if ccValue := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); ccValue.Exists() && ccValue.String() != "" {
						cItem.{{toGoName .TfName}} = make([]{{$name}}{{$cname}}{{$ccname}}{{toGoName .TfName}}, 0)
						ccValue.ForEach(func(cck, ccv gjson.Result) bool {
							ccItem := {{$name}}{{$cname}}{{$ccname}}{{toGoName .TfName}}{}
							{{- range .Attributes}}
							{{- if and (not .TfOnly) (not .Value)}}
							{{- if eq .Type "String"}}
							if cccValue := ccv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); cccValue.Exists() && cccValue.String() != "" {
								ccItem.{{toGoName .TfName}} = types.StringValue(cccValue.String())
							} else {
								ccItem.{{toGoName .TfName}} = types.StringNull()
							}
							{{- else if eq .Type "Int64"}}
							if cccValue := ccv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); cccValue.Exists() && cccValue.String() != "" {
								ccItem.{{toGoName .TfName}} = types.Int64Value(cccValue.Int())
							} else {
								ccItem.{{toGoName .TfName}} = types.Int64Null()
							}
							{{- else if eq .Type "Float64"}}
							if cccValue := ccv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); cccValue.Exists() && cccValue.String() != "" {
								ccItem.{{toGoName .TfName}} = types.Float64Value(cccValue.Float())
							} else {
								ccItem.{{toGoName .TfName}} = types.Float64Null()
							}
							{{- else if eq .Type "Bool"}}
							if cccValue := ccv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); cccValue.Exists() && cccValue.String() != "" {
								ccItem.{{toGoName .TfName}} = types.BoolValue(cccValue.Bool())
							} else {
								ccItem.{{toGoName .TfName}} = types.BoolNull()
							}
							{{- else if eq .Type "ListString"}}
							if cccValue := ccv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); cccValue.Exists() && cccValue.String() != "" {
								ccItem.{{toGoName .TfName}} = helpers.GetListString(cccValue.Array())
							} else {
								ccItem.{{toGoName .TfName}} = types.ListNull(types.StringType)
							}
							{{- end}}
							{{- end}}
							{{- end}}
							cItem.{{toGoName .TfName}} = append(cItem.{{toGoName .TfName}}, ccItem)
							return true
						})
					}
					{{- end}}
					{{- end}}
					{{- end}}
					item.{{toGoName .TfName}} = append(item.{{toGoName .TfName}}, cItem)
					return true
				})
			}
			{{- end}}
			{{- end}}
			{{- end}}
			data.{{toGoName .TfName}} = append(data.{{toGoName .TfName}}, item)
			return true
		})
	}
	{{- end}}
	{{- end}}
	{{- end}}
}
//template:end fromBody
