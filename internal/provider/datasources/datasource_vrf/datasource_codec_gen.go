// Code generated DO NOT EDIT.
package datasource_vrf

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCVrfModel struct {
	FabricName           string                     `json:"fabric,omitempty"`
	VrfName              string                     `json:"vrfName,omitempty"`
	VrfTemplate          string                     `json:"vrfTemplate,omitempty"`
	VrfExtensionTemplate string                     `json:"vrfExtensionTemplate,omitempty"`
	VrfId                *int64                     `json:"vrfId,omitempty"`
	VrfTemplateConfig    NDFCVrfTemplateConfigValue `json:"vrfTemplateConfig,omitempty"`
}

type NDFCVrfTemplateConfigValue struct {
	VlanId                      *int64 `json:"vrfVlanId,string,omitempty"`
	VlanName                    string `json:"vrfVlanName,omitempty"`
	InterfaceDescription        string `json:"vrfIntfDescription,omitempty"`
	VrfDescription              string `json:"vrfDescription,omitempty"`
	Mtu                         *int64 `json:"mtu,string,omitempty"`
	LoopbackRoutingTag          *int64 `json:"tag,string,omitempty"`
	RedistributeDirectRouteMap  string `json:"vrfRouteMap,omitempty"`
	MaxBgpPaths                 *int64 `json:"maxBgpPaths,string,omitempty"`
	MaxIbgpPaths                *int64 `json:"maxIbgpPaths,string,omitempty"`
	Ipv6LinkLocal               string `json:"ipv6LinkLocalFlag,omitempty"`
	Trm                         string `json:"trmEnabled,omitempty"`
	NoRp                        string `json:"isRPAbsent,omitempty"`
	RpExternal                  string `json:"isRPExternal,omitempty"`
	RpAddress                   string `json:"rpAddress,omitempty"`
	RpLoopbackId                *int64 `json:"loopbackNumber,string,omitempty"`
	UnderlayMulticastAddress    string `json:"L3VniMcastGroup,omitempty"`
	OverlayMulticastGroups      string `json:"multicastGroup,omitempty"`
	MvpnInterAs                 string `json:"mvpnInterAs,omitempty"`
	TrmBgwMsite                 string `json:"trmBGWMSiteEnabled,omitempty"`
	AdvertiseHostRoutes         string `json:"advertiseHostRouteFlag,omitempty"`
	AdvertiseDefaultRoute       string `json:"advertiseDefaultRouteFlag,omitempty"`
	ConfigureStaticDefaultRoute string `json:"configureStaticDefaultRouteFlag,omitempty"`
	BgpPassword                 string `json:"bgpPassword,omitempty"`
	BgpPasswordType             string `json:"bgpPasswordKeyType,omitempty"`
	Netflow                     string `json:"ENABLE_NETFLOW,omitempty"`
	NetflowMonitor              string `json:"NETFLOW_MONITOR,omitempty"`
	DisableRtAuto               string `json:"disableRtAuto,omitempty"`
	RouteTargetImport           string `json:"routeTargetImport,omitempty"`
	RouteTargetExport           string `json:"routeTargetExport,omitempty"`
	RouteTargetImportEvpn       string `json:"routeTargetImportEvpn,omitempty"`
	RouteTargetExportEvpn       string `json:"routeTargetExportEvpn,omitempty"`
	RouteTargetImportMvpn       string `json:"routeTargetImportMvpn,omitempty"`
	RouteTargetExportMvpn       string `json:"routeTargetExportMvpn,omitempty"`
	RouteTargetImportCloudEvpn  string `json:"cloudRouteTargetImportEvpn,omitempty"`
	RouteTargetExportCloudEvpn  string `json:"cloudRouteTargetExportEvpn,omitempty"`
}

func (v *VrfModel) SetModelData(jsonData *NDFCVrfModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.FabricName != "" {
		v.FabricName = types.StringValue(jsonData.FabricName)
	} else {
		v.FabricName = types.StringNull()
	}

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
		v.VlanId = types.Int64Value(*jsonData.VrfTemplateConfig.VlanId)

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
		v.RpLoopbackId = types.Int64Value(*jsonData.VrfTemplateConfig.RpLoopbackId)

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

func (v VrfModel) GetModelData() *NDFCVrfModel {
	var data = new(NDFCVrfModel)
	if !v.FabricName.IsNull() && !v.FabricName.IsUnknown() {
		data.FabricName = v.FabricName.ValueString()
	} else {
		data.FabricName = ""
	}

	if !v.VrfName.IsNull() && !v.VrfName.IsUnknown() {
		data.VrfName = v.VrfName.ValueString()
	} else {
		data.VrfName = ""
	}

	if !v.VrfTemplate.IsNull() && !v.VrfTemplate.IsUnknown() {
		data.VrfTemplate = v.VrfTemplate.ValueString()
	} else {
		data.VrfTemplate = ""
	}

	if !v.VrfExtensionTemplate.IsNull() && !v.VrfExtensionTemplate.IsUnknown() {
		data.VrfExtensionTemplate = v.VrfExtensionTemplate.ValueString()
	} else {
		data.VrfExtensionTemplate = ""
	}

	if !v.VrfId.IsNull() && !v.VrfId.IsUnknown() {
		data.VrfId = new(int64)
		*data.VrfId = v.VrfId.ValueInt64()

	} else {
		data.VrfId = nil
	}

	if !v.VlanId.IsNull() && !v.VlanId.IsUnknown() {
		data.VrfTemplateConfig.VlanId = new(int64)
		*data.VrfTemplateConfig.VlanId = v.VlanId.ValueInt64()

	} else {
		data.VrfTemplateConfig.VlanId = nil
	}

	if !v.VlanName.IsNull() && !v.VlanName.IsUnknown() {
		data.VrfTemplateConfig.VlanName = v.VlanName.ValueString()
	} else {
		data.VrfTemplateConfig.VlanName = ""
	}

	if !v.InterfaceDescription.IsNull() && !v.InterfaceDescription.IsUnknown() {
		data.VrfTemplateConfig.InterfaceDescription = v.InterfaceDescription.ValueString()
	} else {
		data.VrfTemplateConfig.InterfaceDescription = ""
	}

	if !v.VrfDescription.IsNull() && !v.VrfDescription.IsUnknown() {
		data.VrfTemplateConfig.VrfDescription = v.VrfDescription.ValueString()
	} else {
		data.VrfTemplateConfig.VrfDescription = ""
	}

	if !v.Mtu.IsNull() && !v.Mtu.IsUnknown() {
		data.VrfTemplateConfig.Mtu = new(int64)
		*data.VrfTemplateConfig.Mtu = v.Mtu.ValueInt64()

	} else {
		data.VrfTemplateConfig.Mtu = nil
	}

	if !v.LoopbackRoutingTag.IsNull() && !v.LoopbackRoutingTag.IsUnknown() {
		data.VrfTemplateConfig.LoopbackRoutingTag = new(int64)
		*data.VrfTemplateConfig.LoopbackRoutingTag = v.LoopbackRoutingTag.ValueInt64()

	} else {
		data.VrfTemplateConfig.LoopbackRoutingTag = nil
	}

	if !v.RedistributeDirectRouteMap.IsNull() && !v.RedistributeDirectRouteMap.IsUnknown() {
		data.VrfTemplateConfig.RedistributeDirectRouteMap = v.RedistributeDirectRouteMap.ValueString()
	} else {
		data.VrfTemplateConfig.RedistributeDirectRouteMap = ""
	}

	if !v.MaxBgpPaths.IsNull() && !v.MaxBgpPaths.IsUnknown() {
		data.VrfTemplateConfig.MaxBgpPaths = new(int64)
		*data.VrfTemplateConfig.MaxBgpPaths = v.MaxBgpPaths.ValueInt64()

	} else {
		data.VrfTemplateConfig.MaxBgpPaths = nil
	}

	if !v.MaxIbgpPaths.IsNull() && !v.MaxIbgpPaths.IsUnknown() {
		data.VrfTemplateConfig.MaxIbgpPaths = new(int64)
		*data.VrfTemplateConfig.MaxIbgpPaths = v.MaxIbgpPaths.ValueInt64()

	} else {
		data.VrfTemplateConfig.MaxIbgpPaths = nil
	}

	if !v.Ipv6LinkLocal.IsNull() && !v.Ipv6LinkLocal.IsUnknown() {
		data.VrfTemplateConfig.Ipv6LinkLocal = strconv.FormatBool(v.Ipv6LinkLocal.ValueBool())
	} else {
		data.VrfTemplateConfig.Ipv6LinkLocal = ""
	}

	if !v.Trm.IsNull() && !v.Trm.IsUnknown() {
		data.VrfTemplateConfig.Trm = strconv.FormatBool(v.Trm.ValueBool())
	} else {
		data.VrfTemplateConfig.Trm = ""
	}

	if !v.NoRp.IsNull() && !v.NoRp.IsUnknown() {
		data.VrfTemplateConfig.NoRp = strconv.FormatBool(v.NoRp.ValueBool())
	} else {
		data.VrfTemplateConfig.NoRp = ""
	}

	if !v.RpExternal.IsNull() && !v.RpExternal.IsUnknown() {
		data.VrfTemplateConfig.RpExternal = strconv.FormatBool(v.RpExternal.ValueBool())
	} else {
		data.VrfTemplateConfig.RpExternal = ""
	}

	if !v.RpAddress.IsNull() && !v.RpAddress.IsUnknown() {
		data.VrfTemplateConfig.RpAddress = v.RpAddress.ValueString()
	} else {
		data.VrfTemplateConfig.RpAddress = ""
	}

	if !v.RpLoopbackId.IsNull() && !v.RpLoopbackId.IsUnknown() {
		data.VrfTemplateConfig.RpLoopbackId = new(int64)
		*data.VrfTemplateConfig.RpLoopbackId = v.RpLoopbackId.ValueInt64()

	} else {
		data.VrfTemplateConfig.RpLoopbackId = nil
	}

	if !v.UnderlayMulticastAddress.IsNull() && !v.UnderlayMulticastAddress.IsUnknown() {
		data.VrfTemplateConfig.UnderlayMulticastAddress = v.UnderlayMulticastAddress.ValueString()
	} else {
		data.VrfTemplateConfig.UnderlayMulticastAddress = ""
	}

	if !v.OverlayMulticastGroups.IsNull() && !v.OverlayMulticastGroups.IsUnknown() {
		data.VrfTemplateConfig.OverlayMulticastGroups = v.OverlayMulticastGroups.ValueString()
	} else {
		data.VrfTemplateConfig.OverlayMulticastGroups = ""
	}

	if !v.MvpnInterAs.IsNull() && !v.MvpnInterAs.IsUnknown() {
		data.VrfTemplateConfig.MvpnInterAs = strconv.FormatBool(v.MvpnInterAs.ValueBool())
	} else {
		data.VrfTemplateConfig.MvpnInterAs = ""
	}

	if !v.TrmBgwMsite.IsNull() && !v.TrmBgwMsite.IsUnknown() {
		data.VrfTemplateConfig.TrmBgwMsite = strconv.FormatBool(v.TrmBgwMsite.ValueBool())
	} else {
		data.VrfTemplateConfig.TrmBgwMsite = ""
	}

	if !v.AdvertiseHostRoutes.IsNull() && !v.AdvertiseHostRoutes.IsUnknown() {
		data.VrfTemplateConfig.AdvertiseHostRoutes = strconv.FormatBool(v.AdvertiseHostRoutes.ValueBool())
	} else {
		data.VrfTemplateConfig.AdvertiseHostRoutes = ""
	}

	if !v.AdvertiseDefaultRoute.IsNull() && !v.AdvertiseDefaultRoute.IsUnknown() {
		data.VrfTemplateConfig.AdvertiseDefaultRoute = strconv.FormatBool(v.AdvertiseDefaultRoute.ValueBool())
	} else {
		data.VrfTemplateConfig.AdvertiseDefaultRoute = ""
	}

	if !v.ConfigureStaticDefaultRoute.IsNull() && !v.ConfigureStaticDefaultRoute.IsUnknown() {
		data.VrfTemplateConfig.ConfigureStaticDefaultRoute = strconv.FormatBool(v.ConfigureStaticDefaultRoute.ValueBool())
	} else {
		data.VrfTemplateConfig.ConfigureStaticDefaultRoute = ""
	}

	if !v.BgpPassword.IsNull() && !v.BgpPassword.IsUnknown() {
		data.VrfTemplateConfig.BgpPassword = v.BgpPassword.ValueString()
	} else {
		data.VrfTemplateConfig.BgpPassword = ""
	}

	if !v.BgpPasswordType.IsNull() && !v.BgpPasswordType.IsUnknown() {
		data.VrfTemplateConfig.BgpPasswordType = v.BgpPasswordType.ValueString()
	} else {
		data.VrfTemplateConfig.BgpPasswordType = ""
	}

	if !v.Netflow.IsNull() && !v.Netflow.IsUnknown() {
		data.VrfTemplateConfig.Netflow = strconv.FormatBool(v.Netflow.ValueBool())
	} else {
		data.VrfTemplateConfig.Netflow = ""
	}

	if !v.NetflowMonitor.IsNull() && !v.NetflowMonitor.IsUnknown() {
		data.VrfTemplateConfig.NetflowMonitor = v.NetflowMonitor.ValueString()
	} else {
		data.VrfTemplateConfig.NetflowMonitor = ""
	}

	if !v.DisableRtAuto.IsNull() && !v.DisableRtAuto.IsUnknown() {
		data.VrfTemplateConfig.DisableRtAuto = strconv.FormatBool(v.DisableRtAuto.ValueBool())
	} else {
		data.VrfTemplateConfig.DisableRtAuto = ""
	}

	if !v.RouteTargetImport.IsNull() && !v.RouteTargetImport.IsUnknown() {
		data.VrfTemplateConfig.RouteTargetImport = v.RouteTargetImport.ValueString()
	} else {
		data.VrfTemplateConfig.RouteTargetImport = ""
	}

	if !v.RouteTargetExport.IsNull() && !v.RouteTargetExport.IsUnknown() {
		data.VrfTemplateConfig.RouteTargetExport = v.RouteTargetExport.ValueString()
	} else {
		data.VrfTemplateConfig.RouteTargetExport = ""
	}

	if !v.RouteTargetImportEvpn.IsNull() && !v.RouteTargetImportEvpn.IsUnknown() {
		data.VrfTemplateConfig.RouteTargetImportEvpn = v.RouteTargetImportEvpn.ValueString()
	} else {
		data.VrfTemplateConfig.RouteTargetImportEvpn = ""
	}

	if !v.RouteTargetExportEvpn.IsNull() && !v.RouteTargetExportEvpn.IsUnknown() {
		data.VrfTemplateConfig.RouteTargetExportEvpn = v.RouteTargetExportEvpn.ValueString()
	} else {
		data.VrfTemplateConfig.RouteTargetExportEvpn = ""
	}

	if !v.RouteTargetImportMvpn.IsNull() && !v.RouteTargetImportMvpn.IsUnknown() {
		data.VrfTemplateConfig.RouteTargetImportMvpn = v.RouteTargetImportMvpn.ValueString()
	} else {
		data.VrfTemplateConfig.RouteTargetImportMvpn = ""
	}

	if !v.RouteTargetExportMvpn.IsNull() && !v.RouteTargetExportMvpn.IsUnknown() {
		data.VrfTemplateConfig.RouteTargetExportMvpn = v.RouteTargetExportMvpn.ValueString()
	} else {
		data.VrfTemplateConfig.RouteTargetExportMvpn = ""
	}

	if !v.RouteTargetImportCloudEvpn.IsNull() && !v.RouteTargetImportCloudEvpn.IsUnknown() {
		data.VrfTemplateConfig.RouteTargetImportCloudEvpn = v.RouteTargetImportCloudEvpn.ValueString()
	} else {
		data.VrfTemplateConfig.RouteTargetImportCloudEvpn = ""
	}

	if !v.RouteTargetExportCloudEvpn.IsNull() && !v.RouteTargetExportCloudEvpn.IsUnknown() {
		data.VrfTemplateConfig.RouteTargetExportCloudEvpn = v.RouteTargetExportCloudEvpn.ValueString()
	} else {
		data.VrfTemplateConfig.RouteTargetExportCloudEvpn = ""
	}

	return data
}
