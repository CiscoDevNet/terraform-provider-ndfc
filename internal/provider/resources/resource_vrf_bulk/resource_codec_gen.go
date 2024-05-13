// Code generated DO NOT EDIT.
package resource_vrf_bulk

import (
	"context"
	"log"
	"strconv"
	. "terraform-provider-ndfc/internal/provider/types"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"terraform-provider-ndfc/internal/provider/resources/resource_vrf_attachments"
)

type NDFCVrfBulkModel struct {
	FabricName           string                   `json:"fabric,omitempty"`
	DeployAllAttachments bool                     `json:"-"`
	Vrfs                 map[string]NDFCVrfsValue `json:"vrfs,omitempty"`
}

type NDFCVrfsValue struct {
	Id                   *int64                                                  `json:"id,omitempty"`
	FilterThisValue      bool                                                    `json:"-"`
	VrfName              string                                                  `json:"vrfName,omitempty"`
	FabricName           string                                                  `json:"fabric,omitempty"`
	VrfTemplate          string                                                  `json:"vrfTemplate,omitempty"`
	VrfExtensionTemplate string                                                  `json:"vrfExtensionTemplate,omitempty"`
	VrfId                *int64                                                  `json:"vrfId,omitempty"`
	VrfStatus            string                                                  `json:"vrfStatus,omitempty"`
	DeployAttachments    bool                                                    `json:"-"`
	AttachList           map[string]resource_vrf_attachments.NDFCAttachListValue `json:"lanAttachList,omitempty"`
	VrfTemplateConfig    NDFCVrfTemplateConfigValue                              `json:"vrfTemplateConfig,omitempty"`
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

func (v *VrfBulkModel) SetModelData(jsonData *NDFCVrfBulkModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.FabricName != "" {
		v.FabricName = types.StringValue(jsonData.FabricName)
	} else {
		v.FabricName = types.StringNull()
	}

	v.DeployAllAttachments = types.BoolValue(jsonData.DeployAllAttachments)
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
		if jsonData.VrfTemplateConfig.RpLoopbackId.IsEmpty() {
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

	v.DeployAttachments = types.BoolValue(jsonData.DeployAttachments)
	if len(jsonData.AttachList) == 0 {
		log.Printf("v.AttachList is empty")
		v.AttachList = types.MapNull(AttachListValue{}.Type(context.Background()))
	} else {
		mapData := make(map[string]AttachListValue)
		for key, item := range jsonData.AttachList {
			if item.FilterThisValue {
				//Skip this entry - this parameter allows filtering
				continue
			}

			data := new(AttachListValue)
			err = data.SetValue(&item)
			if err != nil {
				log.Printf("Error in AttachListValue.SetValue")
				return err
			}
			data.state = attr.ValueStateKnown
			mapData[key] = *data
		}
		v.AttachList, err = types.MapValueFrom(context.Background(), AttachListValue{}.Type(context.Background()), mapData)
		if err != nil {
			log.Printf("Error in converting map[string]AttachListValue to  Map")

		}
	}

	return err
}

func (v *AttachListValue) SetValue(jsonData *resource_vrf_attachments.NDFCAttachListValue) diag.Diagnostics {

	var err diag.Diagnostics
	err = nil

	if jsonData.SwitchName != "" {
		v.SwitchName = types.StringValue(jsonData.SwitchName)
	} else {
		v.SwitchName = types.StringNull()
	}

	if jsonData.Vlan != nil {
		if jsonData.Vlan.IsEmpty() {
			v.Vlan = types.Int64Null()
		} else {
			v.Vlan = types.Int64Value(int64(*jsonData.Vlan))
		}
	} else if jsonData.VlanId != nil {
		if jsonData.VlanId.IsEmpty() {
			v.Vlan = types.Int64Null()
		} else {
			v.Vlan = types.Int64Value(int64(*jsonData.VlanId))
		}
	} else {
		v.Vlan = types.Int64Null()
	}

	if jsonData.AttachState != "" {
		v.AttachState = types.StringValue(jsonData.AttachState)
	} else {
		v.AttachState = types.StringNull()
	}

	if jsonData.Attached != nil {
		v.Attached = types.BoolValue(*jsonData.Attached)

	} else {
		v.Attached = types.BoolNull()
	}

	if jsonData.FreeformConfig != "" {
		v.FreeformConfig = types.StringValue(jsonData.FreeformConfig)
	} else {
		v.FreeformConfig = types.StringNull()
	}

	v.DeployThisAttachment = types.BoolValue(jsonData.DeployThisAttachment)

	if jsonData.InstanceValues.LoopbackId != nil {
		if jsonData.InstanceValues.LoopbackId.IsEmpty() {
			v.LoopbackId = types.Int64Null()
		} else {
			v.LoopbackId = types.Int64Value(int64(*jsonData.InstanceValues.LoopbackId))
		}

	} else {
		v.LoopbackId = types.Int64Null()
	}

	if jsonData.InstanceValues.LoopbackIpv4 != "" {
		v.LoopbackIpv4 = types.StringValue(jsonData.InstanceValues.LoopbackIpv4)
	} else {
		v.LoopbackIpv4 = types.StringNull()
	}

	if jsonData.InstanceValues.LoopbackIpv6 != "" {
		v.LoopbackIpv6 = types.StringValue(jsonData.InstanceValues.LoopbackIpv6)
	} else {
		v.LoopbackIpv6 = types.StringNull()
	}

	return err
}

func (v VrfBulkModel) GetModelData() *NDFCVrfBulkModel {
	var data = new(NDFCVrfBulkModel)

	//MARSHAL_BODY

	if !v.FabricName.IsNull() && !v.FabricName.IsUnknown() {
		data.FabricName = v.FabricName.ValueString()
	} else {
		data.FabricName = ""
	}

	if !v.DeployAllAttachments.IsNull() && !v.DeployAllAttachments.IsUnknown() {
		data.DeployAllAttachments = v.DeployAllAttachments.ValueBool()
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

			// no_rp | Bool| [vrfTemplateConfig]| false
			if !ele1.NoRp.IsNull() && !ele1.NoRp.IsUnknown() {
				//-----inline nested----
				data1.VrfTemplateConfig.NoRp = strconv.FormatBool(ele1.NoRp.ValueBool())
			} else {
				data1.VrfTemplateConfig.NoRp = ""
			}

			// rp_external | Bool| [vrfTemplateConfig]| false
			if !ele1.RpExternal.IsNull() && !ele1.RpExternal.IsUnknown() {
				//-----inline nested----
				data1.VrfTemplateConfig.RpExternal = strconv.FormatBool(ele1.RpExternal.ValueBool())
			} else {
				data1.VrfTemplateConfig.RpExternal = ""
			}

			// rp_address | String| [vrfTemplateConfig]| false
			if !ele1.RpAddress.IsNull() && !ele1.RpAddress.IsUnknown() {
				//-----inline nested----
				data1.VrfTemplateConfig.RpAddress = ele1.RpAddress.ValueString()
			} else {
				data1.VrfTemplateConfig.RpAddress = ""
			}

			// rp_loopback_id | Int64| [vrfTemplateConfig]| false
			if !ele1.RpLoopbackId.IsNull() && !ele1.RpLoopbackId.IsUnknown() {
				//-----inline nested----
				data1.VrfTemplateConfig.RpLoopbackId = new(Int64Custom)
				*data1.VrfTemplateConfig.RpLoopbackId = Int64Custom(ele1.RpLoopbackId.ValueInt64())
			} else {
				data1.VrfTemplateConfig.RpLoopbackId = nil
			}

			// underlay_multicast_address | String| [vrfTemplateConfig]| false
			if !ele1.UnderlayMulticastAddress.IsNull() && !ele1.UnderlayMulticastAddress.IsUnknown() {
				//-----inline nested----
				data1.VrfTemplateConfig.UnderlayMulticastAddress = ele1.UnderlayMulticastAddress.ValueString()
			} else {
				data1.VrfTemplateConfig.UnderlayMulticastAddress = ""
			}

			// overlay_multicast_groups | String| [vrfTemplateConfig]| false
			if !ele1.OverlayMulticastGroups.IsNull() && !ele1.OverlayMulticastGroups.IsUnknown() {
				//-----inline nested----
				data1.VrfTemplateConfig.OverlayMulticastGroups = ele1.OverlayMulticastGroups.ValueString()
			} else {
				data1.VrfTemplateConfig.OverlayMulticastGroups = ""
			}

			// mvpn_inter_as | Bool| [vrfTemplateConfig]| false
			if !ele1.MvpnInterAs.IsNull() && !ele1.MvpnInterAs.IsUnknown() {
				//-----inline nested----
				data1.VrfTemplateConfig.MvpnInterAs = strconv.FormatBool(ele1.MvpnInterAs.ValueBool())
			} else {
				data1.VrfTemplateConfig.MvpnInterAs = ""
			}

			// trm_bgw_msite | Bool| [vrfTemplateConfig]| false
			if !ele1.TrmBgwMsite.IsNull() && !ele1.TrmBgwMsite.IsUnknown() {
				//-----inline nested----
				data1.VrfTemplateConfig.TrmBgwMsite = strconv.FormatBool(ele1.TrmBgwMsite.ValueBool())
			} else {
				data1.VrfTemplateConfig.TrmBgwMsite = ""
			}

			// advertise_host_routes | Bool| [vrfTemplateConfig]| false
			if !ele1.AdvertiseHostRoutes.IsNull() && !ele1.AdvertiseHostRoutes.IsUnknown() {
				//-----inline nested----
				data1.VrfTemplateConfig.AdvertiseHostRoutes = strconv.FormatBool(ele1.AdvertiseHostRoutes.ValueBool())
			} else {
				data1.VrfTemplateConfig.AdvertiseHostRoutes = ""
			}

			// advertise_default_route | Bool| [vrfTemplateConfig]| false
			if !ele1.AdvertiseDefaultRoute.IsNull() && !ele1.AdvertiseDefaultRoute.IsUnknown() {
				//-----inline nested----
				data1.VrfTemplateConfig.AdvertiseDefaultRoute = strconv.FormatBool(ele1.AdvertiseDefaultRoute.ValueBool())
			} else {
				data1.VrfTemplateConfig.AdvertiseDefaultRoute = ""
			}

			// configure_static_default_route | Bool| [vrfTemplateConfig]| false
			if !ele1.ConfigureStaticDefaultRoute.IsNull() && !ele1.ConfigureStaticDefaultRoute.IsUnknown() {
				//-----inline nested----
				data1.VrfTemplateConfig.ConfigureStaticDefaultRoute = strconv.FormatBool(ele1.ConfigureStaticDefaultRoute.ValueBool())
			} else {
				data1.VrfTemplateConfig.ConfigureStaticDefaultRoute = ""
			}

			// bgp_password | String| [vrfTemplateConfig]| false
			if !ele1.BgpPassword.IsNull() && !ele1.BgpPassword.IsUnknown() {
				//-----inline nested----
				data1.VrfTemplateConfig.BgpPassword = ele1.BgpPassword.ValueString()
			} else {
				data1.VrfTemplateConfig.BgpPassword = ""
			}

			// bgp_password_type | String| [vrfTemplateConfig]| false
			if !ele1.BgpPasswordType.IsNull() && !ele1.BgpPasswordType.IsUnknown() {
				//-----inline nested----
				data1.VrfTemplateConfig.BgpPasswordType = ele1.BgpPasswordType.ValueString()
			} else {
				data1.VrfTemplateConfig.BgpPasswordType = ""
			}

			// netflow | Bool| [vrfTemplateConfig]| false
			if !ele1.Netflow.IsNull() && !ele1.Netflow.IsUnknown() {
				//-----inline nested----
				data1.VrfTemplateConfig.Netflow = strconv.FormatBool(ele1.Netflow.ValueBool())
			} else {
				data1.VrfTemplateConfig.Netflow = ""
			}

			// netflow_monitor | String| [vrfTemplateConfig]| false
			if !ele1.NetflowMonitor.IsNull() && !ele1.NetflowMonitor.IsUnknown() {
				//-----inline nested----
				data1.VrfTemplateConfig.NetflowMonitor = ele1.NetflowMonitor.ValueString()
			} else {
				data1.VrfTemplateConfig.NetflowMonitor = ""
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

			// route_target_import_mvpn | String| [vrfTemplateConfig]| false
			if !ele1.RouteTargetImportMvpn.IsNull() && !ele1.RouteTargetImportMvpn.IsUnknown() {
				//-----inline nested----
				data1.VrfTemplateConfig.RouteTargetImportMvpn = ele1.RouteTargetImportMvpn.ValueString()
			} else {
				data1.VrfTemplateConfig.RouteTargetImportMvpn = ""
			}

			// route_target_export_mvpn | String| [vrfTemplateConfig]| false
			if !ele1.RouteTargetExportMvpn.IsNull() && !ele1.RouteTargetExportMvpn.IsUnknown() {
				//-----inline nested----
				data1.VrfTemplateConfig.RouteTargetExportMvpn = ele1.RouteTargetExportMvpn.ValueString()
			} else {
				data1.VrfTemplateConfig.RouteTargetExportMvpn = ""
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

			// deploy_attachments | Bool| []| false
			if !ele1.DeployAttachments.IsNull() && !ele1.DeployAttachments.IsUnknown() {

				data1.DeployAttachments = ele1.DeployAttachments.ValueBool()

			}

			// attach_list | MapNested| []| false

			if !ele1.AttachList.IsNull() && !ele1.AttachList.IsUnknown() {
				elements2 := make(map[string]AttachListValue, len(ele1.AttachList.Elements()))

				data1.AttachList = make(map[string]resource_vrf_attachments.NDFCAttachListValue)

				diag := ele1.AttachList.ElementsAs(context.Background(), &elements2, false)
				if diag != nil {
					panic(diag)
				}
				for k2, ele2 := range elements2 {
					data2 := new(resource_vrf_attachments.NDFCAttachListValue)

					// filter_this_value | Bool| []| true
					// id | Int64| []| true
					// fabric_name | String| []| true
					// vrf_name | String| []| true
					// serial_number | String| []| true
					// switch_name | String| []| false
					// vlan | Int64| []| false
					if !ele2.Vlan.IsNull() && !ele2.Vlan.IsUnknown() {
						data2.Vlan = new(Int64Custom)
						*data2.Vlan = Int64Custom(ele2.Vlan.ValueInt64())
					} else {
						data2.Vlan = nil
					}

					// deployment | Bool| []| true
					// attach_state | String| []| false
					// attached | Bool| []| false
					// freeform_config | String| []| false
					if !ele2.FreeformConfig.IsNull() && !ele2.FreeformConfig.IsUnknown() {

						data2.FreeformConfig = ele2.FreeformConfig.ValueString()
					} else {
						data2.FreeformConfig = ""
					}

					// deploy_this_attachment | Bool| []| false
					if !ele2.DeployThisAttachment.IsNull() && !ele2.DeployThisAttachment.IsUnknown() {

						data2.DeployThisAttachment = ele2.DeployThisAttachment.ValueBool()

					}

					// loopback_id | Int64| [instanceValues]| false
					if !ele2.LoopbackId.IsNull() && !ele2.LoopbackId.IsUnknown() {
						//-----inline nested----
						data2.InstanceValues.LoopbackId = new(Int64Custom)
						*data2.InstanceValues.LoopbackId = Int64Custom(ele2.LoopbackId.ValueInt64())
					} else {
						data2.InstanceValues.LoopbackId = nil
					}

					// loopback_ipv4 | String| [instanceValues]| false
					if !ele2.LoopbackIpv4.IsNull() && !ele2.LoopbackIpv4.IsUnknown() {
						//-----inline nested----
						data2.InstanceValues.LoopbackIpv4 = ele2.LoopbackIpv4.ValueString()
					} else {
						data2.InstanceValues.LoopbackIpv4 = ""
					}

					// loopback_ipv6 | String| [instanceValues]| false
					if !ele2.LoopbackIpv6.IsNull() && !ele2.LoopbackIpv6.IsUnknown() {
						//-----inline nested----
						data2.InstanceValues.LoopbackIpv6 = ele2.LoopbackIpv6.ValueString()
					} else {
						data2.InstanceValues.LoopbackIpv6 = ""
					}

					// update_action | BitMask| []| true
					data1.AttachList[k2] = *data2

				}
			}

			data.Vrfs[k1] = *data1

		}
	}

	return data
}
