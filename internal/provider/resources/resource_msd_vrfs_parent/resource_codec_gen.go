// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package resource_msd_vrfs_parent

import (
	"context"
	"log"
	"strconv"
	. "terraform-provider-ndfc/internal/provider/types"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCMsdVrfsParentModel struct {
	FabricName string                   `json:"fabric,omitempty"`
	Vrfs       map[string]NDFCVrfsValue `json:"vrfs,omitempty"`
}

type NDFCVrfsValue struct {
	Id                   *int64                     `json:"id,omitempty"`
	FilterThisValue      bool                       `json:"-"`
	VrfName              string                     `json:"vrfName,omitempty"`
	FabricName           string                     `json:"fabric,omitempty"`
	VrfTemplate          string                     `json:"vrfTemplate,omitempty"`
	VrfExtensionTemplate string                     `json:"vrfExtensionTemplate,omitempty"`
	VrfId                *int64                     `json:"vrfId,omitempty"`
	VrfStatus            string                     `json:"vrfStatus,omitempty"`
	VrfTemplateConfig    NDFCVrfTemplateConfigValue `json:"vrfTemplateConfig,omitempty"`
}

type NDFCVrfTemplateConfigValue struct {
	VlanId                     *Int64Custom `json:"vrfVlanId,omitempty"`
	VlanName                   string       `json:"vrfVlanName,omitempty"`
	InterfaceDescription       string       `json:"vrfIntfDescription,omitempty"`
	VrfDescription             string       `json:"vrfDescription,omitempty"`
	Mtu                        *int64       `json:"mtu,string,omitempty"`
	LoopbackRoutingTag         *int64       `json:"tag,string,omitempty"`
	RedistributeDirectRouteMap string       `json:"vrfRouteMap,omitempty"`
	MaxBgpPaths                *int64       `json:"maxBgpPaths,string,omitempty"`
	MaxIbgpPaths               *int64       `json:"maxIbgpPaths,string,omitempty"`
	Ipv6LinkLocal              string       `json:"ipv6LinkLocalFlag,omitempty"`
	Trm                        string       `json:"trmEnabled,omitempty"`
	RpExternal                 string       `json:"isRPExternal,omitempty"`
	MvpnInterAs                string       `json:"mvpnInterAs,omitempty"`
	DisableRtAuto              string       `json:"disableRtAuto,omitempty"`
	RouteTargetImport          string       `json:"routeTargetImport,omitempty"`
	RouteTargetExport          string       `json:"routeTargetExport,omitempty"`
	RouteTargetImportEvpn      string       `json:"routeTargetImportEvpn,omitempty"`
	RouteTargetExportEvpn      string       `json:"routeTargetExportEvpn,omitempty"`
	RouteTargetImportCloudEvpn string       `json:"cloudRouteTargetImportEvpn,omitempty"`
	RouteTargetExportCloudEvpn string       `json:"cloudRouteTargetExportEvpn,omitempty"`
}

func (v *MsdVrfsParentModel) SetModelData(jsonData *NDFCMsdVrfsParentModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.FabricName != "" {
		v.FabricName = types.StringValue(jsonData.FabricName)
	} else {
		v.FabricName = types.StringNull()
	}

	if len(jsonData.Vrfs) == 0 {
		log.Printf("v.Vrfs is empty")
		v.Vrfs = types.MapNull(VrfsValue{}.Type(context.Background()))
	} else {
		mapData := make(map[string]VrfsValue)
		for key, item := range jsonData.Vrfs {
			if item.FilterThisValue {
				//Skip this entry - this parameter allows filtering
				continue
			}

			data := new(VrfsValue)
			err = data.SetValue(&item)
			if err != nil {
				log.Printf("Error in VrfsValue.SetValue")
				return err
			}
			data.state = attr.ValueStateKnown
			mapData[key] = *data
		}
		v.Vrfs, err = types.MapValueFrom(context.Background(), VrfsValue{}.Type(context.Background()), mapData)
		if err != nil {
			log.Printf("Error in converting map[string]VrfsValue to  Map")

		}
	}

	return err
}

func (v *VrfsValue) SetValue(jsonData *NDFCVrfsValue) diag.Diagnostics {

	var err diag.Diagnostics
	err = nil

	if jsonData.VrfTemplate != "" {
		v.VrfTemplate = types.StringValue(jsonData.VrfTemplate)
	} else {
		v.VrfTemplate = types.StringNull()
	}

	if jsonData.VrfExtensionTemplate != "" {
		v.VrfExtensionTemplate = types.StringValue(jsonData.VrfExtensionTemplate)
	} else {
		v.VrfExtensionTemplate = types.StringNull()
	}

	if jsonData.VrfId != nil {
		v.VrfId = types.Int64Value(*jsonData.VrfId)

	} else {
		v.VrfId = types.Int64Null()
	}

	if jsonData.VrfTemplateConfig.VlanId != nil {
		if jsonData.VrfTemplateConfig.VlanId.IsEmpty() {
			v.VlanId = types.Int64Null()
		} else {
			v.VlanId = types.Int64Value(int64(*jsonData.VrfTemplateConfig.VlanId))
		}

	} else {
		v.VlanId = types.Int64Null()
	}

	if jsonData.VrfTemplateConfig.VlanName != "" {
		v.VlanName = types.StringValue(jsonData.VrfTemplateConfig.VlanName)
	} else {
		v.VlanName = types.StringNull()
	}

	if jsonData.VrfTemplateConfig.InterfaceDescription != "" {
		v.InterfaceDescription = types.StringValue(jsonData.VrfTemplateConfig.InterfaceDescription)
	} else {
		v.InterfaceDescription = types.StringNull()
	}

	if jsonData.VrfTemplateConfig.VrfDescription != "" {
		v.VrfDescription = types.StringValue(jsonData.VrfTemplateConfig.VrfDescription)
	} else {
		v.VrfDescription = types.StringNull()
	}

	if jsonData.VrfTemplateConfig.Mtu != nil {
		v.Mtu = types.Int64Value(*jsonData.VrfTemplateConfig.Mtu)

	} else {
		v.Mtu = types.Int64Null()
	}

	if jsonData.VrfStatus != "" {
		v.VrfStatus = types.StringValue(jsonData.VrfStatus)
	} else {
		v.VrfStatus = types.StringNull()
	}

	if jsonData.VrfTemplateConfig.LoopbackRoutingTag != nil {
		v.LoopbackRoutingTag = types.Int64Value(*jsonData.VrfTemplateConfig.LoopbackRoutingTag)

	} else {
		v.LoopbackRoutingTag = types.Int64Null()
	}

	if jsonData.VrfTemplateConfig.RedistributeDirectRouteMap != "" {
		v.RedistributeDirectRouteMap = types.StringValue(jsonData.VrfTemplateConfig.RedistributeDirectRouteMap)
	} else {
		v.RedistributeDirectRouteMap = types.StringNull()
	}

	if jsonData.VrfTemplateConfig.MaxBgpPaths != nil {
		v.MaxBgpPaths = types.Int64Value(*jsonData.VrfTemplateConfig.MaxBgpPaths)

	} else {
		v.MaxBgpPaths = types.Int64Null()
	}

	if jsonData.VrfTemplateConfig.MaxIbgpPaths != nil {
		v.MaxIbgpPaths = types.Int64Value(*jsonData.VrfTemplateConfig.MaxIbgpPaths)

	} else {
		v.MaxIbgpPaths = types.Int64Null()
	}

	if jsonData.VrfTemplateConfig.Ipv6LinkLocal != "" {
		x, _ := strconv.ParseBool(jsonData.VrfTemplateConfig.Ipv6LinkLocal)
		v.Ipv6LinkLocal = types.BoolValue(x)
	} else {
		v.Ipv6LinkLocal = types.BoolNull()
	}

	if jsonData.VrfTemplateConfig.Trm != "" {
		x, _ := strconv.ParseBool(jsonData.VrfTemplateConfig.Trm)
		v.Trm = types.BoolValue(x)
	} else {
		v.Trm = types.BoolNull()
	}

	if jsonData.VrfTemplateConfig.RpExternal != "" {
		x, _ := strconv.ParseBool(jsonData.VrfTemplateConfig.RpExternal)
		v.RpExternal = types.BoolValue(x)
	} else {
		v.RpExternal = types.BoolNull()
	}

	if jsonData.VrfTemplateConfig.MvpnInterAs != "" {
		x, _ := strconv.ParseBool(jsonData.VrfTemplateConfig.MvpnInterAs)
		v.MvpnInterAs = types.BoolValue(x)
	} else {
		v.MvpnInterAs = types.BoolNull()
	}

	if jsonData.VrfTemplateConfig.DisableRtAuto != "" {
		x, _ := strconv.ParseBool(jsonData.VrfTemplateConfig.DisableRtAuto)
		v.DisableRtAuto = types.BoolValue(x)
	} else {
		v.DisableRtAuto = types.BoolNull()
	}

	if jsonData.VrfTemplateConfig.RouteTargetImport != "" {
		v.RouteTargetImport = types.StringValue(jsonData.VrfTemplateConfig.RouteTargetImport)
	} else {
		v.RouteTargetImport = types.StringNull()
	}

	if jsonData.VrfTemplateConfig.RouteTargetExport != "" {
		v.RouteTargetExport = types.StringValue(jsonData.VrfTemplateConfig.RouteTargetExport)
	} else {
		v.RouteTargetExport = types.StringNull()
	}

	if jsonData.VrfTemplateConfig.RouteTargetImportEvpn != "" {
		v.RouteTargetImportEvpn = types.StringValue(jsonData.VrfTemplateConfig.RouteTargetImportEvpn)
	} else {
		v.RouteTargetImportEvpn = types.StringNull()
	}

	if jsonData.VrfTemplateConfig.RouteTargetExportEvpn != "" {
		v.RouteTargetExportEvpn = types.StringValue(jsonData.VrfTemplateConfig.RouteTargetExportEvpn)
	} else {
		v.RouteTargetExportEvpn = types.StringNull()
	}

	if jsonData.VrfTemplateConfig.RouteTargetImportCloudEvpn != "" {
		v.RouteTargetImportCloudEvpn = types.StringValue(jsonData.VrfTemplateConfig.RouteTargetImportCloudEvpn)
	} else {
		v.RouteTargetImportCloudEvpn = types.StringNull()
	}

	if jsonData.VrfTemplateConfig.RouteTargetExportCloudEvpn != "" {
		v.RouteTargetExportCloudEvpn = types.StringValue(jsonData.VrfTemplateConfig.RouteTargetExportCloudEvpn)
	} else {
		v.RouteTargetExportCloudEvpn = types.StringNull()
	}

	return err
}

func (v MsdVrfsParentModel) GetModelData() *NDFCMsdVrfsParentModel {
	var data = new(NDFCMsdVrfsParentModel)

	//MARSHAL_BODY

	if !v.FabricName.IsNull() && !v.FabricName.IsUnknown() {
		data.FabricName = v.FabricName.ValueString()
	} else {
		data.FabricName = ""
	}

	if !v.Vrfs.IsNull() && !v.Vrfs.IsUnknown() {
		elements1 := make(map[string]VrfsValue, len(v.Vrfs.Elements()))

		data.Vrfs = make(map[string]NDFCVrfsValue)

		diag := v.Vrfs.ElementsAs(context.Background(), &elements1, false)
		if diag != nil {
			panic(diag)
		}
		for k1, ele1 := range elements1 {
			data1 := new(NDFCVrfsValue)

			// id | Int64| []| true
			// filter_this_value | Bool| []| true
			// vrf_name | String| []| true
			// fabric_name | String| []| true
			// vrf_template | String| []| false
			if !ele1.VrfTemplate.IsNull() && !ele1.VrfTemplate.IsUnknown() {

				data1.VrfTemplate = ele1.VrfTemplate.ValueString()
			} else {
				data1.VrfTemplate = ""
			}

			// vrf_extension_template | String| []| false
			if !ele1.VrfExtensionTemplate.IsNull() && !ele1.VrfExtensionTemplate.IsUnknown() {

				data1.VrfExtensionTemplate = ele1.VrfExtensionTemplate.ValueString()
			} else {
				data1.VrfExtensionTemplate = ""
			}

			// vrf_id | Int64| []| false
			if !ele1.VrfId.IsNull() && !ele1.VrfId.IsUnknown() {

				data1.VrfId = new(int64)
				*data1.VrfId = ele1.VrfId.ValueInt64()

			} else {
				data1.VrfId = nil
			}

			// vlan_id | Int64| [vrfTemplateConfig]| false
			if !ele1.VlanId.IsNull() && !ele1.VlanId.IsUnknown() {
				//-----inline nested----
				data1.VrfTemplateConfig.VlanId = new(Int64Custom)
				*data1.VrfTemplateConfig.VlanId = Int64Custom(ele1.VlanId.ValueInt64())
			} else {
				data1.VrfTemplateConfig.VlanId = nil
			}

			// vlan_name | String| [vrfTemplateConfig]| false
			if !ele1.VlanName.IsNull() && !ele1.VlanName.IsUnknown() {
				//-----inline nested----
				data1.VrfTemplateConfig.VlanName = ele1.VlanName.ValueString()
			} else {
				data1.VrfTemplateConfig.VlanName = ""
			}

			// interface_description | String| [vrfTemplateConfig]| false
			if !ele1.InterfaceDescription.IsNull() && !ele1.InterfaceDescription.IsUnknown() {
				//-----inline nested----
				data1.VrfTemplateConfig.InterfaceDescription = ele1.InterfaceDescription.ValueString()
			} else {
				data1.VrfTemplateConfig.InterfaceDescription = ""
			}

			// vrf_description | String| [vrfTemplateConfig]| false
			if !ele1.VrfDescription.IsNull() && !ele1.VrfDescription.IsUnknown() {
				//-----inline nested----
				data1.VrfTemplateConfig.VrfDescription = ele1.VrfDescription.ValueString()
			} else {
				data1.VrfTemplateConfig.VrfDescription = ""
			}

			// mtu | Int64| [vrfTemplateConfig]| false
			if !ele1.Mtu.IsNull() && !ele1.Mtu.IsUnknown() {
				//-----inline nested----
				data1.VrfTemplateConfig.Mtu = new(int64)
				*data1.VrfTemplateConfig.Mtu = ele1.Mtu.ValueInt64()

			} else {
				data1.VrfTemplateConfig.Mtu = nil
			}

			// vrf_status | String| []| false
			// loopback_routing_tag | Int64| [vrfTemplateConfig]| false
			if !ele1.LoopbackRoutingTag.IsNull() && !ele1.LoopbackRoutingTag.IsUnknown() {
				//-----inline nested----
				data1.VrfTemplateConfig.LoopbackRoutingTag = new(int64)
				*data1.VrfTemplateConfig.LoopbackRoutingTag = ele1.LoopbackRoutingTag.ValueInt64()

			} else {
				data1.VrfTemplateConfig.LoopbackRoutingTag = nil
			}

			// redistribute_direct_route_map | String| [vrfTemplateConfig]| false
			if !ele1.RedistributeDirectRouteMap.IsNull() && !ele1.RedistributeDirectRouteMap.IsUnknown() {
				//-----inline nested----
				data1.VrfTemplateConfig.RedistributeDirectRouteMap = ele1.RedistributeDirectRouteMap.ValueString()
			} else {
				data1.VrfTemplateConfig.RedistributeDirectRouteMap = ""
			}

			// max_bgp_paths | Int64| [vrfTemplateConfig]| false
			if !ele1.MaxBgpPaths.IsNull() && !ele1.MaxBgpPaths.IsUnknown() {
				//-----inline nested----
				data1.VrfTemplateConfig.MaxBgpPaths = new(int64)
				*data1.VrfTemplateConfig.MaxBgpPaths = ele1.MaxBgpPaths.ValueInt64()

			} else {
				data1.VrfTemplateConfig.MaxBgpPaths = nil
			}

			// max_ibgp_paths | Int64| [vrfTemplateConfig]| false
			if !ele1.MaxIbgpPaths.IsNull() && !ele1.MaxIbgpPaths.IsUnknown() {
				//-----inline nested----
				data1.VrfTemplateConfig.MaxIbgpPaths = new(int64)
				*data1.VrfTemplateConfig.MaxIbgpPaths = ele1.MaxIbgpPaths.ValueInt64()

			} else {
				data1.VrfTemplateConfig.MaxIbgpPaths = nil
			}

			// ipv6_link_local | Bool| [vrfTemplateConfig]| false
			if !ele1.Ipv6LinkLocal.IsNull() && !ele1.Ipv6LinkLocal.IsUnknown() {
				//-----inline nested----
				data1.VrfTemplateConfig.Ipv6LinkLocal = strconv.FormatBool(ele1.Ipv6LinkLocal.ValueBool())
			} else {
				data1.VrfTemplateConfig.Ipv6LinkLocal = ""
			}

			// trm | Bool| [vrfTemplateConfig]| false
			if !ele1.Trm.IsNull() && !ele1.Trm.IsUnknown() {
				//-----inline nested----
				data1.VrfTemplateConfig.Trm = strconv.FormatBool(ele1.Trm.ValueBool())
			} else {
				data1.VrfTemplateConfig.Trm = ""
			}

			// rp_external | Bool| [vrfTemplateConfig]| false
			if !ele1.RpExternal.IsNull() && !ele1.RpExternal.IsUnknown() {
				//-----inline nested----
				data1.VrfTemplateConfig.RpExternal = strconv.FormatBool(ele1.RpExternal.ValueBool())
			} else {
				data1.VrfTemplateConfig.RpExternal = ""
			}

			// mvpn_inter_as | Bool| [vrfTemplateConfig]| false
			if !ele1.MvpnInterAs.IsNull() && !ele1.MvpnInterAs.IsUnknown() {
				//-----inline nested----
				data1.VrfTemplateConfig.MvpnInterAs = strconv.FormatBool(ele1.MvpnInterAs.ValueBool())
			} else {
				data1.VrfTemplateConfig.MvpnInterAs = ""
			}

			// disable_rt_auto | Bool| [vrfTemplateConfig]| false
			if !ele1.DisableRtAuto.IsNull() && !ele1.DisableRtAuto.IsUnknown() {
				//-----inline nested----
				data1.VrfTemplateConfig.DisableRtAuto = strconv.FormatBool(ele1.DisableRtAuto.ValueBool())
			} else {
				data1.VrfTemplateConfig.DisableRtAuto = ""
			}

			// route_target_import | String| [vrfTemplateConfig]| false
			if !ele1.RouteTargetImport.IsNull() && !ele1.RouteTargetImport.IsUnknown() {
				//-----inline nested----
				data1.VrfTemplateConfig.RouteTargetImport = ele1.RouteTargetImport.ValueString()
			} else {
				data1.VrfTemplateConfig.RouteTargetImport = ""
			}

			// route_target_export | String| [vrfTemplateConfig]| false
			if !ele1.RouteTargetExport.IsNull() && !ele1.RouteTargetExport.IsUnknown() {
				//-----inline nested----
				data1.VrfTemplateConfig.RouteTargetExport = ele1.RouteTargetExport.ValueString()
			} else {
				data1.VrfTemplateConfig.RouteTargetExport = ""
			}

			// route_target_import_evpn | String| [vrfTemplateConfig]| false
			if !ele1.RouteTargetImportEvpn.IsNull() && !ele1.RouteTargetImportEvpn.IsUnknown() {
				//-----inline nested----
				data1.VrfTemplateConfig.RouteTargetImportEvpn = ele1.RouteTargetImportEvpn.ValueString()
			} else {
				data1.VrfTemplateConfig.RouteTargetImportEvpn = ""
			}

			// route_target_export_evpn | String| [vrfTemplateConfig]| false
			if !ele1.RouteTargetExportEvpn.IsNull() && !ele1.RouteTargetExportEvpn.IsUnknown() {
				//-----inline nested----
				data1.VrfTemplateConfig.RouteTargetExportEvpn = ele1.RouteTargetExportEvpn.ValueString()
			} else {
				data1.VrfTemplateConfig.RouteTargetExportEvpn = ""
			}

			// route_target_import_cloud_evpn | String| [vrfTemplateConfig]| false
			if !ele1.RouteTargetImportCloudEvpn.IsNull() && !ele1.RouteTargetImportCloudEvpn.IsUnknown() {
				//-----inline nested----
				data1.VrfTemplateConfig.RouteTargetImportCloudEvpn = ele1.RouteTargetImportCloudEvpn.ValueString()
			} else {
				data1.VrfTemplateConfig.RouteTargetImportCloudEvpn = ""
			}

			// route_target_export_cloud_evpn | String| [vrfTemplateConfig]| false
			if !ele1.RouteTargetExportCloudEvpn.IsNull() && !ele1.RouteTargetExportCloudEvpn.IsUnknown() {
				//-----inline nested----
				data1.VrfTemplateConfig.RouteTargetExportCloudEvpn = ele1.RouteTargetExportCloudEvpn.ValueString()
			} else {
				data1.VrfTemplateConfig.RouteTargetExportCloudEvpn = ""
			}

			data.Vrfs[k1] = *data1

		}
	}

	return data
}
