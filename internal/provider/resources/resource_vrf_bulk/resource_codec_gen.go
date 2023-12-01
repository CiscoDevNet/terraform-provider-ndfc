// Code generated DO NOT EDIT.
package resource_vrf_bulk

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Int64Custom int64

func (i *Int64Custom) UnmarshalJSON(data []byte) error {
	if string(data) == "" || string(data) == "\"\"" {
		*i = -9223372036854775808
	} else {
		ss := string(data)
		ss, _ = strconv.Unquote(ss)
		ii, _ := strconv.ParseInt(ss, 10, 64)
		*i = Int64Custom(ii)
	}

	return nil
}

func (i Int64Custom) MarshalJSON() ([]byte, error) {
	res := ""
	res = strconv.FormatInt(int64(i), 10)
	return []byte(strconv.Quote(res)), nil

}

type NDFCVrfBulkModel struct {
	FabricName string                    `json:"fabric,omitempty"`
	Vrfs       NDFCVrfsValues            `json:"vrfs,omitempty"`
	VrfsMap    map[string]*NDFCVrfsValue `json:"-"`
}

type NDFCVrfsValues []NDFCVrfsValue

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
	VlanId                      *Int64Custom `json:"vrfVlanId,omitempty"`
	VlanName                    string       `json:"vrfVlanName,omitempty"`
	InterfaceDescription        string       `json:"vrfIntfDescription,omitempty"`
	VrfDescription              string       `json:"vrfDescription,omitempty"`
	Mtu                         *int64       `json:"mtu,string,omitempty"`
	LoopbackRoutingTag          *int64       `json:"tag,string,omitempty"`
	RedistributeDirectRouteMap  string       `json:"vrfRouteMap,omitempty"`
	MaxBgpPaths                 *int64       `json:"maxBgpPaths,string,omitempty"`
	MaxIbgpPaths                *int64       `json:"maxIbgpPaths,string,omitempty"`
	Ipv6LinkLocal               string       `json:"ipv6LinkLocalFlag,omitempty"`
	Trm                         string       `json:"trmEnabled,omitempty"`
	NoRp                        string       `json:"isRPAbsent,omitempty"`
	RpExternal                  string       `json:"isRPExternal,omitempty"`
	RpAddress                   string       `json:"rpAddress,omitempty"`
	RpLoopbackId                *Int64Custom `json:"loopbackNumber,omitempty"`
	UnderlayMulticastAddress    string       `json:"L3VniMcastGroup,omitempty"`
	OverlayMulticastGroups      string       `json:"multicastGroup,omitempty"`
	MvpnInterAs                 string       `json:"mvpnInterAs,omitempty"`
	TrmBgwMsite                 string       `json:"trmBGWMSiteEnabled,omitempty"`
	AdvertiseHostRoutes         string       `json:"advertiseHostRouteFlag,omitempty"`
	AdvertiseDefaultRoute       string       `json:"advertiseDefaultRouteFlag,omitempty"`
	ConfigureStaticDefaultRoute string       `json:"configureStaticDefaultRouteFlag,omitempty"`
	BgpPassword                 string       `json:"bgpPassword,omitempty"`
	BgpPasswordType             string       `json:"bgpPasswordKeyType,omitempty"`
	Netflow                     string       `json:"ENABLE_NETFLOW,omitempty"`
	NetflowMonitor              string       `json:"NETFLOW_MONITOR,omitempty"`
	DisableRtAuto               string       `json:"disableRtAuto,omitempty"`
	RouteTargetImport           string       `json:"routeTargetImport,omitempty"`
	RouteTargetExport           string       `json:"routeTargetExport,omitempty"`
	RouteTargetImportEvpn       string       `json:"routeTargetImportEvpn,omitempty"`
	RouteTargetExportEvpn       string       `json:"routeTargetExportEvpn,omitempty"`
	RouteTargetImportMvpn       string       `json:"routeTargetImportMvpn,omitempty"`
	RouteTargetExportMvpn       string       `json:"routeTargetExportMvpn,omitempty"`
	RouteTargetImportCloudEvpn  string       `json:"cloudRouteTargetImportEvpn,omitempty"`
	RouteTargetExportCloudEvpn  string       `json:"cloudRouteTargetExportEvpn,omitempty"`
}

func (s NDFCVrfsValues) Len() int {
	return len(s)
}

func (s NDFCVrfsValues) Less(i, j int) bool {
	return (*s[i].Id < *s[j].Id)

}

func (s NDFCVrfsValues) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (v *VrfBulkModel) SetModelData(jsonData *NDFCVrfBulkModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.FabricName != "" {
		v.FabricName = types.StringValue(jsonData.FabricName)
	} else {
		v.FabricName = types.StringNull()
	}

	listData := make([]VrfsValue, 0)
	for _, item := range jsonData.Vrfs {
		if item.FilterThisValue {
			//Skip this entry - this parameter allows filtering
			continue
		}

		data := new(VrfsValue)
		err = data.SetValue(&item)
		if err != nil {
			return err
		}
		data.state = attr.ValueStateKnown
		listData = append(listData, *data)
	}
	v.Vrfs, err = types.ListValueFrom(context.Background(), VrfsValue{}.Type(context.Background()), listData)
	if err != nil {
		return err
	}

	return err
}

func (v *VrfsValue) SetValue(jsonData *NDFCVrfsValue) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.VrfName != "" {
		v.VrfName = types.StringValue(jsonData.VrfName)
	} else {
		v.VrfName = types.StringNull()
	}

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
		if int64(*jsonData.VrfTemplateConfig.VlanId) == -9223372036854775808 {
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

	if jsonData.VrfTemplateConfig.NoRp != "" {
		x, _ := strconv.ParseBool(jsonData.VrfTemplateConfig.NoRp)
		v.NoRp = types.BoolValue(x)
	} else {
		v.NoRp = types.BoolNull()
	}

	if jsonData.VrfTemplateConfig.RpExternal != "" {
		x, _ := strconv.ParseBool(jsonData.VrfTemplateConfig.RpExternal)
		v.RpExternal = types.BoolValue(x)
	} else {
		v.RpExternal = types.BoolNull()
	}

	if jsonData.VrfTemplateConfig.RpAddress != "" {
		v.RpAddress = types.StringValue(jsonData.VrfTemplateConfig.RpAddress)
	} else {
		v.RpAddress = types.StringNull()
	}

	if jsonData.VrfTemplateConfig.RpLoopbackId != nil {
		if int64(*jsonData.VrfTemplateConfig.RpLoopbackId) == -9223372036854775808 {
			v.RpLoopbackId = types.Int64Null()
		} else {
			v.RpLoopbackId = types.Int64Value(int64(*jsonData.VrfTemplateConfig.RpLoopbackId))
		}

	} else {
		v.RpLoopbackId = types.Int64Null()
	}

	if jsonData.VrfTemplateConfig.UnderlayMulticastAddress != "" {
		v.UnderlayMulticastAddress = types.StringValue(jsonData.VrfTemplateConfig.UnderlayMulticastAddress)
	} else {
		v.UnderlayMulticastAddress = types.StringNull()
	}

	if jsonData.VrfTemplateConfig.OverlayMulticastGroups != "" {
		v.OverlayMulticastGroups = types.StringValue(jsonData.VrfTemplateConfig.OverlayMulticastGroups)
	} else {
		v.OverlayMulticastGroups = types.StringNull()
	}

	if jsonData.VrfTemplateConfig.MvpnInterAs != "" {
		x, _ := strconv.ParseBool(jsonData.VrfTemplateConfig.MvpnInterAs)
		v.MvpnInterAs = types.BoolValue(x)
	} else {
		v.MvpnInterAs = types.BoolNull()
	}

	if jsonData.VrfTemplateConfig.TrmBgwMsite != "" {
		x, _ := strconv.ParseBool(jsonData.VrfTemplateConfig.TrmBgwMsite)
		v.TrmBgwMsite = types.BoolValue(x)
	} else {
		v.TrmBgwMsite = types.BoolNull()
	}

	if jsonData.VrfTemplateConfig.AdvertiseHostRoutes != "" {
		x, _ := strconv.ParseBool(jsonData.VrfTemplateConfig.AdvertiseHostRoutes)
		v.AdvertiseHostRoutes = types.BoolValue(x)
	} else {
		v.AdvertiseHostRoutes = types.BoolNull()
	}

	if jsonData.VrfTemplateConfig.AdvertiseDefaultRoute != "" {
		x, _ := strconv.ParseBool(jsonData.VrfTemplateConfig.AdvertiseDefaultRoute)
		v.AdvertiseDefaultRoute = types.BoolValue(x)
	} else {
		v.AdvertiseDefaultRoute = types.BoolNull()
	}

	if jsonData.VrfTemplateConfig.ConfigureStaticDefaultRoute != "" {
		x, _ := strconv.ParseBool(jsonData.VrfTemplateConfig.ConfigureStaticDefaultRoute)
		v.ConfigureStaticDefaultRoute = types.BoolValue(x)
	} else {
		v.ConfigureStaticDefaultRoute = types.BoolNull()
	}

	if jsonData.VrfTemplateConfig.BgpPassword != "" {
		v.BgpPassword = types.StringValue(jsonData.VrfTemplateConfig.BgpPassword)
	} else {
		v.BgpPassword = types.StringNull()
	}

	if jsonData.VrfTemplateConfig.BgpPasswordType != "" {
		v.BgpPasswordType = types.StringValue(jsonData.VrfTemplateConfig.BgpPasswordType)
	} else {
		v.BgpPasswordType = types.StringNull()
	}

	if jsonData.VrfTemplateConfig.Netflow != "" {
		x, _ := strconv.ParseBool(jsonData.VrfTemplateConfig.Netflow)
		v.Netflow = types.BoolValue(x)
	} else {
		v.Netflow = types.BoolNull()
	}

	if jsonData.VrfTemplateConfig.NetflowMonitor != "" {
		v.NetflowMonitor = types.StringValue(jsonData.VrfTemplateConfig.NetflowMonitor)
	} else {
		v.NetflowMonitor = types.StringNull()
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

	if jsonData.VrfTemplateConfig.RouteTargetImportMvpn != "" {
		v.RouteTargetImportMvpn = types.StringValue(jsonData.VrfTemplateConfig.RouteTargetImportMvpn)
	} else {
		v.RouteTargetImportMvpn = types.StringNull()
	}

	if jsonData.VrfTemplateConfig.RouteTargetExportMvpn != "" {
		v.RouteTargetExportMvpn = types.StringValue(jsonData.VrfTemplateConfig.RouteTargetExportMvpn)
	} else {
		v.RouteTargetExportMvpn = types.StringNull()
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

func (v VrfBulkModel) GetModelData() *NDFCVrfBulkModel {
	var data = new(NDFCVrfBulkModel)
	if !v.FabricName.IsNull() && !v.FabricName.IsUnknown() {
		data.FabricName = v.FabricName.ValueString()
	} else {
		data.FabricName = ""
	}

	data.VrfsMap = make(map[string]*NDFCVrfsValue)

	if !v.Vrfs.IsNull() && !v.Vrfs.IsUnknown() {
		elements := make([]VrfsValue, len(v.Vrfs.Elements()))
		data.Vrfs = make([]NDFCVrfsValue, len(v.Vrfs.Elements()))
		diag := v.Vrfs.ElementsAs(context.Background(), &elements, false)
		if diag != nil {
			panic(diag)
		}
		for i, ele := range elements {

			if !ele.VrfName.IsNull() && !ele.VrfName.IsUnknown() {
				data.Vrfs[i].VrfName = ele.VrfName.ValueString()
			} else {
				data.Vrfs[i].VrfName = ""
			}

			data.VrfsMap[data.Vrfs[i].VrfName] = &data.Vrfs[i]

			if !ele.VrfTemplate.IsNull() && !ele.VrfTemplate.IsUnknown() {
				data.Vrfs[i].VrfTemplate = ele.VrfTemplate.ValueString()
			} else {
				data.Vrfs[i].VrfTemplate = ""
			}

			if !ele.VrfExtensionTemplate.IsNull() && !ele.VrfExtensionTemplate.IsUnknown() {
				data.Vrfs[i].VrfExtensionTemplate = ele.VrfExtensionTemplate.ValueString()
			} else {
				data.Vrfs[i].VrfExtensionTemplate = ""
			}

			if !ele.VrfId.IsNull() && !ele.VrfId.IsUnknown() {
				data.Vrfs[i].VrfId = new(int64)
				*data.Vrfs[i].VrfId = ele.VrfId.ValueInt64()

			} else {
				data.Vrfs[i].VrfId = nil
			}

			//-----inline nesting Start----
			if !ele.VlanId.IsNull() && !ele.VlanId.IsUnknown() {
				data.Vrfs[i].VrfTemplateConfig.VlanId = new(Int64Custom)
				*data.Vrfs[i].VrfTemplateConfig.VlanId = Int64Custom(ele.VlanId.ValueInt64())
			} else {
				data.Vrfs[i].VrfTemplateConfig.VlanId = nil
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele.VlanName.IsNull() && !ele.VlanName.IsUnknown() {
				data.Vrfs[i].VrfTemplateConfig.VlanName = ele.VlanName.ValueString()
			} else {
				data.Vrfs[i].VrfTemplateConfig.VlanName = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele.InterfaceDescription.IsNull() && !ele.InterfaceDescription.IsUnknown() {
				data.Vrfs[i].VrfTemplateConfig.InterfaceDescription = ele.InterfaceDescription.ValueString()
			} else {
				data.Vrfs[i].VrfTemplateConfig.InterfaceDescription = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele.VrfDescription.IsNull() && !ele.VrfDescription.IsUnknown() {
				data.Vrfs[i].VrfTemplateConfig.VrfDescription = ele.VrfDescription.ValueString()
			} else {
				data.Vrfs[i].VrfTemplateConfig.VrfDescription = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele.Mtu.IsNull() && !ele.Mtu.IsUnknown() {
				data.Vrfs[i].VrfTemplateConfig.Mtu = new(int64)
				*data.Vrfs[i].VrfTemplateConfig.Mtu = ele.Mtu.ValueInt64()

			} else {
				data.Vrfs[i].VrfTemplateConfig.Mtu = nil
			}
			//-----inline nesting end----

			if !ele.VrfStatus.IsNull() && !ele.VrfStatus.IsUnknown() {
				data.Vrfs[i].VrfStatus = ele.VrfStatus.ValueString()
			} else {
				data.Vrfs[i].VrfStatus = ""
			}

			//-----inline nesting Start----
			if !ele.LoopbackRoutingTag.IsNull() && !ele.LoopbackRoutingTag.IsUnknown() {
				data.Vrfs[i].VrfTemplateConfig.LoopbackRoutingTag = new(int64)
				*data.Vrfs[i].VrfTemplateConfig.LoopbackRoutingTag = ele.LoopbackRoutingTag.ValueInt64()

			} else {
				data.Vrfs[i].VrfTemplateConfig.LoopbackRoutingTag = nil
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele.RedistributeDirectRouteMap.IsNull() && !ele.RedistributeDirectRouteMap.IsUnknown() {
				data.Vrfs[i].VrfTemplateConfig.RedistributeDirectRouteMap = ele.RedistributeDirectRouteMap.ValueString()
			} else {
				data.Vrfs[i].VrfTemplateConfig.RedistributeDirectRouteMap = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele.MaxBgpPaths.IsNull() && !ele.MaxBgpPaths.IsUnknown() {
				data.Vrfs[i].VrfTemplateConfig.MaxBgpPaths = new(int64)
				*data.Vrfs[i].VrfTemplateConfig.MaxBgpPaths = ele.MaxBgpPaths.ValueInt64()

			} else {
				data.Vrfs[i].VrfTemplateConfig.MaxBgpPaths = nil
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele.MaxIbgpPaths.IsNull() && !ele.MaxIbgpPaths.IsUnknown() {
				data.Vrfs[i].VrfTemplateConfig.MaxIbgpPaths = new(int64)
				*data.Vrfs[i].VrfTemplateConfig.MaxIbgpPaths = ele.MaxIbgpPaths.ValueInt64()

			} else {
				data.Vrfs[i].VrfTemplateConfig.MaxIbgpPaths = nil
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele.Ipv6LinkLocal.IsNull() && !ele.Ipv6LinkLocal.IsUnknown() {
				data.Vrfs[i].VrfTemplateConfig.Ipv6LinkLocal = strconv.FormatBool(ele.Ipv6LinkLocal.ValueBool())
			} else {
				data.Vrfs[i].VrfTemplateConfig.Ipv6LinkLocal = ""
			}

			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele.Trm.IsNull() && !ele.Trm.IsUnknown() {
				data.Vrfs[i].VrfTemplateConfig.Trm = strconv.FormatBool(ele.Trm.ValueBool())
			} else {
				data.Vrfs[i].VrfTemplateConfig.Trm = ""
			}

			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele.NoRp.IsNull() && !ele.NoRp.IsUnknown() {
				data.Vrfs[i].VrfTemplateConfig.NoRp = strconv.FormatBool(ele.NoRp.ValueBool())
			} else {
				data.Vrfs[i].VrfTemplateConfig.NoRp = ""
			}

			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele.RpExternal.IsNull() && !ele.RpExternal.IsUnknown() {
				data.Vrfs[i].VrfTemplateConfig.RpExternal = strconv.FormatBool(ele.RpExternal.ValueBool())
			} else {
				data.Vrfs[i].VrfTemplateConfig.RpExternal = ""
			}

			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele.RpAddress.IsNull() && !ele.RpAddress.IsUnknown() {
				data.Vrfs[i].VrfTemplateConfig.RpAddress = ele.RpAddress.ValueString()
			} else {
				data.Vrfs[i].VrfTemplateConfig.RpAddress = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele.RpLoopbackId.IsNull() && !ele.RpLoopbackId.IsUnknown() {
				data.Vrfs[i].VrfTemplateConfig.RpLoopbackId = new(Int64Custom)
				*data.Vrfs[i].VrfTemplateConfig.RpLoopbackId = Int64Custom(ele.RpLoopbackId.ValueInt64())
			} else {
				data.Vrfs[i].VrfTemplateConfig.RpLoopbackId = nil
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele.UnderlayMulticastAddress.IsNull() && !ele.UnderlayMulticastAddress.IsUnknown() {
				data.Vrfs[i].VrfTemplateConfig.UnderlayMulticastAddress = ele.UnderlayMulticastAddress.ValueString()
			} else {
				data.Vrfs[i].VrfTemplateConfig.UnderlayMulticastAddress = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele.OverlayMulticastGroups.IsNull() && !ele.OverlayMulticastGroups.IsUnknown() {
				data.Vrfs[i].VrfTemplateConfig.OverlayMulticastGroups = ele.OverlayMulticastGroups.ValueString()
			} else {
				data.Vrfs[i].VrfTemplateConfig.OverlayMulticastGroups = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele.MvpnInterAs.IsNull() && !ele.MvpnInterAs.IsUnknown() {
				data.Vrfs[i].VrfTemplateConfig.MvpnInterAs = strconv.FormatBool(ele.MvpnInterAs.ValueBool())
			} else {
				data.Vrfs[i].VrfTemplateConfig.MvpnInterAs = ""
			}

			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele.TrmBgwMsite.IsNull() && !ele.TrmBgwMsite.IsUnknown() {
				data.Vrfs[i].VrfTemplateConfig.TrmBgwMsite = strconv.FormatBool(ele.TrmBgwMsite.ValueBool())
			} else {
				data.Vrfs[i].VrfTemplateConfig.TrmBgwMsite = ""
			}

			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele.AdvertiseHostRoutes.IsNull() && !ele.AdvertiseHostRoutes.IsUnknown() {
				data.Vrfs[i].VrfTemplateConfig.AdvertiseHostRoutes = strconv.FormatBool(ele.AdvertiseHostRoutes.ValueBool())
			} else {
				data.Vrfs[i].VrfTemplateConfig.AdvertiseHostRoutes = ""
			}

			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele.AdvertiseDefaultRoute.IsNull() && !ele.AdvertiseDefaultRoute.IsUnknown() {
				data.Vrfs[i].VrfTemplateConfig.AdvertiseDefaultRoute = strconv.FormatBool(ele.AdvertiseDefaultRoute.ValueBool())
			} else {
				data.Vrfs[i].VrfTemplateConfig.AdvertiseDefaultRoute = ""
			}

			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele.ConfigureStaticDefaultRoute.IsNull() && !ele.ConfigureStaticDefaultRoute.IsUnknown() {
				data.Vrfs[i].VrfTemplateConfig.ConfigureStaticDefaultRoute = strconv.FormatBool(ele.ConfigureStaticDefaultRoute.ValueBool())
			} else {
				data.Vrfs[i].VrfTemplateConfig.ConfigureStaticDefaultRoute = ""
			}

			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele.BgpPassword.IsNull() && !ele.BgpPassword.IsUnknown() {
				data.Vrfs[i].VrfTemplateConfig.BgpPassword = ele.BgpPassword.ValueString()
			} else {
				data.Vrfs[i].VrfTemplateConfig.BgpPassword = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele.BgpPasswordType.IsNull() && !ele.BgpPasswordType.IsUnknown() {
				data.Vrfs[i].VrfTemplateConfig.BgpPasswordType = ele.BgpPasswordType.ValueString()
			} else {
				data.Vrfs[i].VrfTemplateConfig.BgpPasswordType = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele.Netflow.IsNull() && !ele.Netflow.IsUnknown() {
				data.Vrfs[i].VrfTemplateConfig.Netflow = strconv.FormatBool(ele.Netflow.ValueBool())
			} else {
				data.Vrfs[i].VrfTemplateConfig.Netflow = ""
			}

			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele.NetflowMonitor.IsNull() && !ele.NetflowMonitor.IsUnknown() {
				data.Vrfs[i].VrfTemplateConfig.NetflowMonitor = ele.NetflowMonitor.ValueString()
			} else {
				data.Vrfs[i].VrfTemplateConfig.NetflowMonitor = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele.DisableRtAuto.IsNull() && !ele.DisableRtAuto.IsUnknown() {
				data.Vrfs[i].VrfTemplateConfig.DisableRtAuto = strconv.FormatBool(ele.DisableRtAuto.ValueBool())
			} else {
				data.Vrfs[i].VrfTemplateConfig.DisableRtAuto = ""
			}

			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele.RouteTargetImport.IsNull() && !ele.RouteTargetImport.IsUnknown() {
				data.Vrfs[i].VrfTemplateConfig.RouteTargetImport = ele.RouteTargetImport.ValueString()
			} else {
				data.Vrfs[i].VrfTemplateConfig.RouteTargetImport = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele.RouteTargetExport.IsNull() && !ele.RouteTargetExport.IsUnknown() {
				data.Vrfs[i].VrfTemplateConfig.RouteTargetExport = ele.RouteTargetExport.ValueString()
			} else {
				data.Vrfs[i].VrfTemplateConfig.RouteTargetExport = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele.RouteTargetImportEvpn.IsNull() && !ele.RouteTargetImportEvpn.IsUnknown() {
				data.Vrfs[i].VrfTemplateConfig.RouteTargetImportEvpn = ele.RouteTargetImportEvpn.ValueString()
			} else {
				data.Vrfs[i].VrfTemplateConfig.RouteTargetImportEvpn = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele.RouteTargetExportEvpn.IsNull() && !ele.RouteTargetExportEvpn.IsUnknown() {
				data.Vrfs[i].VrfTemplateConfig.RouteTargetExportEvpn = ele.RouteTargetExportEvpn.ValueString()
			} else {
				data.Vrfs[i].VrfTemplateConfig.RouteTargetExportEvpn = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele.RouteTargetImportMvpn.IsNull() && !ele.RouteTargetImportMvpn.IsUnknown() {
				data.Vrfs[i].VrfTemplateConfig.RouteTargetImportMvpn = ele.RouteTargetImportMvpn.ValueString()
			} else {
				data.Vrfs[i].VrfTemplateConfig.RouteTargetImportMvpn = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele.RouteTargetExportMvpn.IsNull() && !ele.RouteTargetExportMvpn.IsUnknown() {
				data.Vrfs[i].VrfTemplateConfig.RouteTargetExportMvpn = ele.RouteTargetExportMvpn.ValueString()
			} else {
				data.Vrfs[i].VrfTemplateConfig.RouteTargetExportMvpn = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele.RouteTargetImportCloudEvpn.IsNull() && !ele.RouteTargetImportCloudEvpn.IsUnknown() {
				data.Vrfs[i].VrfTemplateConfig.RouteTargetImportCloudEvpn = ele.RouteTargetImportCloudEvpn.ValueString()
			} else {
				data.Vrfs[i].VrfTemplateConfig.RouteTargetImportCloudEvpn = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele.RouteTargetExportCloudEvpn.IsNull() && !ele.RouteTargetExportCloudEvpn.IsUnknown() {
				data.Vrfs[i].VrfTemplateConfig.RouteTargetExportCloudEvpn = ele.RouteTargetExportCloudEvpn.ValueString()
			} else {
				data.Vrfs[i].VrfTemplateConfig.RouteTargetExportCloudEvpn = ""
			}
			//-----inline nesting end----

		}
	}

	return data
}
