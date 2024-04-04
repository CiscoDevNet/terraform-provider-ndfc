// Code generated DO NOT EDIT.
package datasource_vrf_bulk

import (
	"context"
	"log"
	"strconv"
	. "terraform-provider-ndfc/internal/provider/types"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCVrfBulkModel struct {
	FabricName string         `json:"fabric,omitempty"`
	Vrfs       NDFCVrfsValues `json:"vrfs,omitempty"`
}

type NDFCVrfsValues []NDFCVrfsValue

type NDFCVrfsValue struct {
	Id                   *int64                     `json:"Id,omitempty"`
	VrfName              string                     `json:"vrfName,omitempty"`
	VrfTemplate          string                     `json:"vrfTemplate,omitempty"`
	VrfStatus            string                     `json:"vrfStatus,omitempty"`
	VrfExtensionTemplate string                     `json:"vrfExtensionTemplate,omitempty"`
	VrfId                *int64                     `json:"vrfId,omitempty"`
	AttachList           NDFCAttachListValues       `json:"lanAttachList,omitempty"`
	VrfTemplateConfig    NDFCVrfTemplateConfigValue `json:"vrfTemplateConfig,omitempty"`
}

type NDFCAttachListValues []NDFCAttachListValue

type NDFCAttachListValue struct {
	Id             *int64                  `json:"id,omitempty"`
	SerialNumber   string                  `json:"serialNumber,omitempty"`
	SwitchSerialNo string                  `json:"switchSerialNo,omitempty"`
	SwitchName     string                  `json:"switchName,omitempty"`
	Vlan           *Int64Custom            `json:"vlan,omitempty"`
	VlanId         *Int64Custom            `json:"vlanId,omitempty"`
	AttachState    string                  `json:"lanAttachState,omitempty"`
	Attached       *bool                   `json:"isLanAttached,omitempty"`
	FreeformConfig string                  `json:"freeformconfig,omitempty"`
	InstanceValues NDFCInstanceValuesValue `json:"instanceValues,omitempty"`
}

type NDFCInstanceValuesValue struct {
	LoopbackId   *Int64Custom `json:"loopbackId,omitempty"`
	LoopbackIpv4 string       `json:"loopbackIpAddress,omitempty"`
	LoopbackIpv6 string       `json:"loopbackIpv6Address,omitempty"`
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

	if len(jsonData.Vrfs) == 0 {
		log.Printf("v.Vrfs is empty")
		v.Vrfs = types.ListNull(VrfsValue{}.Type(context.Background()))
	} else {
		log.Printf("v.Vrfs contains %d elements", len(jsonData.Vrfs))
		listData := make([]VrfsValue, 0)
		for _, item := range jsonData.Vrfs {
			data := new(VrfsValue)
			err = data.SetValue(&item)
			if err != nil {
				log.Printf("Error in VrfsValue.SetValue")
				return err
			}
			data.state = attr.ValueStateKnown
			listData = append(listData, *data)
		}
		v.Vrfs, err = types.ListValueFrom(context.Background(), VrfsValue{}.Type(context.Background()), listData)
		if err != nil {
			log.Printf("Error in converting []VrfsValue to  List")
			return err
		}
	}

	return err
}

func (v *VrfsValue) SetValue(jsonData *NDFCVrfsValue) diag.Diagnostics {

	var err diag.Diagnostics
	err = nil

	if jsonData.Id != nil {
		v.Id = types.Int64Value(*jsonData.Id)

	} else {
		v.Id = types.Int64Null()
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

	if jsonData.VrfStatus != "" {
		v.VrfStatus = types.StringValue(jsonData.VrfStatus)
	} else {
		v.VrfStatus = types.StringNull()
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

	if len(jsonData.AttachList) == 0 {
		log.Printf("v.AttachList is empty")
		v.AttachList = types.ListNull(AttachListValue{}.Type(context.Background()))
	} else {
		log.Printf("v.AttachList contains %d elements", len(jsonData.AttachList))
		listData := make([]AttachListValue, 0)
		for _, item := range jsonData.AttachList {
			data := new(AttachListValue)
			err = data.SetValue(&item)
			if err != nil {
				log.Printf("Error in AttachListValue.SetValue")
				return err
			}
			data.state = attr.ValueStateKnown
			listData = append(listData, *data)
		}
		v.AttachList, err = types.ListValueFrom(context.Background(), AttachListValue{}.Type(context.Background()), listData)
		if err != nil {
			log.Printf("Error in converting []AttachListValue to  List")
			return err
		}
	}

	return err
}

func (v *AttachListValue) SetValue(jsonData *NDFCAttachListValue) diag.Diagnostics {

	var err diag.Diagnostics
	err = nil

	if jsonData.Id != nil {
		v.Id = types.Int64Value(*jsonData.Id)

	} else {
		v.Id = types.Int64Null()
	}

	if jsonData.SerialNumber != "" {
		v.SerialNumber = types.StringValue(jsonData.SerialNumber)
	} else if jsonData.SwitchSerialNo != "" {
		v.SerialNumber = types.StringValue(jsonData.SwitchSerialNo)
	} else {
		v.SerialNumber = types.StringNull()
	}

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

	//MARSHALL_LIST

	if !v.Vrfs.IsNull() && !v.Vrfs.IsUnknown() {
		elements := make([]VrfsValue, len(v.Vrfs.Elements()))
		data.Vrfs = make([]NDFCVrfsValue, len(v.Vrfs.Elements()))

		diag := v.Vrfs.ElementsAs(context.Background(), &elements, false)
		if diag != nil {
			panic(diag)
		}
		for i1, ele1 := range elements {
			if !ele1.Id.IsNull() && !ele1.Id.IsUnknown() {
				data.Vrfs[i1].Id = new(int64)
				*data.Vrfs[i1].Id = ele1.Id.ValueInt64()

			} else {
				data.Vrfs[i1].Id = nil
			}

			if !ele1.VrfName.IsNull() && !ele1.VrfName.IsUnknown() {
				data.Vrfs[i1].VrfName = ele1.VrfName.ValueString()
			} else {
				data.Vrfs[i1].VrfName = ""
			}

			if !ele1.VrfTemplate.IsNull() && !ele1.VrfTemplate.IsUnknown() {
				data.Vrfs[i1].VrfTemplate = ele1.VrfTemplate.ValueString()
			} else {
				data.Vrfs[i1].VrfTemplate = ""
			}

			if !ele1.VrfExtensionTemplate.IsNull() && !ele1.VrfExtensionTemplate.IsUnknown() {
				data.Vrfs[i1].VrfExtensionTemplate = ele1.VrfExtensionTemplate.ValueString()
			} else {
				data.Vrfs[i1].VrfExtensionTemplate = ""
			}

			if !ele1.VrfId.IsNull() && !ele1.VrfId.IsUnknown() {
				data.Vrfs[i1].VrfId = new(int64)
				*data.Vrfs[i1].VrfId = ele1.VrfId.ValueInt64()

			} else {
				data.Vrfs[i1].VrfId = nil
			}

			//-----inline nesting Start----
			if !ele1.VlanId.IsNull() && !ele1.VlanId.IsUnknown() {
				data.Vrfs[i1].VrfTemplateConfig.VlanId = new(Int64Custom)
				*data.Vrfs[i1].VrfTemplateConfig.VlanId = Int64Custom(ele1.VlanId.ValueInt64())
			} else {
				data.Vrfs[i1].VrfTemplateConfig.VlanId = nil
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.VlanName.IsNull() && !ele1.VlanName.IsUnknown() {
				data.Vrfs[i1].VrfTemplateConfig.VlanName = ele1.VlanName.ValueString()
			} else {
				data.Vrfs[i1].VrfTemplateConfig.VlanName = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.InterfaceDescription.IsNull() && !ele1.InterfaceDescription.IsUnknown() {
				data.Vrfs[i1].VrfTemplateConfig.InterfaceDescription = ele1.InterfaceDescription.ValueString()
			} else {
				data.Vrfs[i1].VrfTemplateConfig.InterfaceDescription = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.VrfDescription.IsNull() && !ele1.VrfDescription.IsUnknown() {
				data.Vrfs[i1].VrfTemplateConfig.VrfDescription = ele1.VrfDescription.ValueString()
			} else {
				data.Vrfs[i1].VrfTemplateConfig.VrfDescription = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.Mtu.IsNull() && !ele1.Mtu.IsUnknown() {
				data.Vrfs[i1].VrfTemplateConfig.Mtu = new(int64)
				*data.Vrfs[i1].VrfTemplateConfig.Mtu = ele1.Mtu.ValueInt64()

			} else {
				data.Vrfs[i1].VrfTemplateConfig.Mtu = nil
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.LoopbackRoutingTag.IsNull() && !ele1.LoopbackRoutingTag.IsUnknown() {
				data.Vrfs[i1].VrfTemplateConfig.LoopbackRoutingTag = new(int64)
				*data.Vrfs[i1].VrfTemplateConfig.LoopbackRoutingTag = ele1.LoopbackRoutingTag.ValueInt64()

			} else {
				data.Vrfs[i1].VrfTemplateConfig.LoopbackRoutingTag = nil
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.RedistributeDirectRouteMap.IsNull() && !ele1.RedistributeDirectRouteMap.IsUnknown() {
				data.Vrfs[i1].VrfTemplateConfig.RedistributeDirectRouteMap = ele1.RedistributeDirectRouteMap.ValueString()
			} else {
				data.Vrfs[i1].VrfTemplateConfig.RedistributeDirectRouteMap = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.MaxBgpPaths.IsNull() && !ele1.MaxBgpPaths.IsUnknown() {
				data.Vrfs[i1].VrfTemplateConfig.MaxBgpPaths = new(int64)
				*data.Vrfs[i1].VrfTemplateConfig.MaxBgpPaths = ele1.MaxBgpPaths.ValueInt64()

			} else {
				data.Vrfs[i1].VrfTemplateConfig.MaxBgpPaths = nil
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.MaxIbgpPaths.IsNull() && !ele1.MaxIbgpPaths.IsUnknown() {
				data.Vrfs[i1].VrfTemplateConfig.MaxIbgpPaths = new(int64)
				*data.Vrfs[i1].VrfTemplateConfig.MaxIbgpPaths = ele1.MaxIbgpPaths.ValueInt64()

			} else {
				data.Vrfs[i1].VrfTemplateConfig.MaxIbgpPaths = nil
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.Ipv6LinkLocal.IsNull() && !ele1.Ipv6LinkLocal.IsUnknown() { // test signature1  NDFCType:
				data.Vrfs[i1].VrfTemplateConfig.Ipv6LinkLocal = strconv.FormatBool(ele1.Ipv6LinkLocal.ValueBool())
			} else {
				data.Vrfs[i1].VrfTemplateConfig.Ipv6LinkLocal = ""
			}

			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.Trm.IsNull() && !ele1.Trm.IsUnknown() { // test signature1  NDFCType:
				data.Vrfs[i1].VrfTemplateConfig.Trm = strconv.FormatBool(ele1.Trm.ValueBool())
			} else {
				data.Vrfs[i1].VrfTemplateConfig.Trm = ""
			}

			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.NoRp.IsNull() && !ele1.NoRp.IsUnknown() { // test signature1  NDFCType:
				data.Vrfs[i1].VrfTemplateConfig.NoRp = strconv.FormatBool(ele1.NoRp.ValueBool())
			} else {
				data.Vrfs[i1].VrfTemplateConfig.NoRp = ""
			}

			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.RpExternal.IsNull() && !ele1.RpExternal.IsUnknown() { // test signature1  NDFCType:
				data.Vrfs[i1].VrfTemplateConfig.RpExternal = strconv.FormatBool(ele1.RpExternal.ValueBool())
			} else {
				data.Vrfs[i1].VrfTemplateConfig.RpExternal = ""
			}

			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.RpAddress.IsNull() && !ele1.RpAddress.IsUnknown() {
				data.Vrfs[i1].VrfTemplateConfig.RpAddress = ele1.RpAddress.ValueString()
			} else {
				data.Vrfs[i1].VrfTemplateConfig.RpAddress = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.RpLoopbackId.IsNull() && !ele1.RpLoopbackId.IsUnknown() {
				data.Vrfs[i1].VrfTemplateConfig.RpLoopbackId = new(Int64Custom)
				*data.Vrfs[i1].VrfTemplateConfig.RpLoopbackId = Int64Custom(ele1.RpLoopbackId.ValueInt64())
			} else {
				data.Vrfs[i1].VrfTemplateConfig.RpLoopbackId = nil
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.UnderlayMulticastAddress.IsNull() && !ele1.UnderlayMulticastAddress.IsUnknown() {
				data.Vrfs[i1].VrfTemplateConfig.UnderlayMulticastAddress = ele1.UnderlayMulticastAddress.ValueString()
			} else {
				data.Vrfs[i1].VrfTemplateConfig.UnderlayMulticastAddress = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.OverlayMulticastGroups.IsNull() && !ele1.OverlayMulticastGroups.IsUnknown() {
				data.Vrfs[i1].VrfTemplateConfig.OverlayMulticastGroups = ele1.OverlayMulticastGroups.ValueString()
			} else {
				data.Vrfs[i1].VrfTemplateConfig.OverlayMulticastGroups = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.MvpnInterAs.IsNull() && !ele1.MvpnInterAs.IsUnknown() { // test signature1  NDFCType:
				data.Vrfs[i1].VrfTemplateConfig.MvpnInterAs = strconv.FormatBool(ele1.MvpnInterAs.ValueBool())
			} else {
				data.Vrfs[i1].VrfTemplateConfig.MvpnInterAs = ""
			}

			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.TrmBgwMsite.IsNull() && !ele1.TrmBgwMsite.IsUnknown() { // test signature1  NDFCType:
				data.Vrfs[i1].VrfTemplateConfig.TrmBgwMsite = strconv.FormatBool(ele1.TrmBgwMsite.ValueBool())
			} else {
				data.Vrfs[i1].VrfTemplateConfig.TrmBgwMsite = ""
			}

			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.AdvertiseHostRoutes.IsNull() && !ele1.AdvertiseHostRoutes.IsUnknown() { // test signature1  NDFCType:
				data.Vrfs[i1].VrfTemplateConfig.AdvertiseHostRoutes = strconv.FormatBool(ele1.AdvertiseHostRoutes.ValueBool())
			} else {
				data.Vrfs[i1].VrfTemplateConfig.AdvertiseHostRoutes = ""
			}

			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.AdvertiseDefaultRoute.IsNull() && !ele1.AdvertiseDefaultRoute.IsUnknown() { // test signature1  NDFCType:
				data.Vrfs[i1].VrfTemplateConfig.AdvertiseDefaultRoute = strconv.FormatBool(ele1.AdvertiseDefaultRoute.ValueBool())
			} else {
				data.Vrfs[i1].VrfTemplateConfig.AdvertiseDefaultRoute = ""
			}

			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.ConfigureStaticDefaultRoute.IsNull() && !ele1.ConfigureStaticDefaultRoute.IsUnknown() { // test signature1  NDFCType:
				data.Vrfs[i1].VrfTemplateConfig.ConfigureStaticDefaultRoute = strconv.FormatBool(ele1.ConfigureStaticDefaultRoute.ValueBool())
			} else {
				data.Vrfs[i1].VrfTemplateConfig.ConfigureStaticDefaultRoute = ""
			}

			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.BgpPassword.IsNull() && !ele1.BgpPassword.IsUnknown() {
				data.Vrfs[i1].VrfTemplateConfig.BgpPassword = ele1.BgpPassword.ValueString()
			} else {
				data.Vrfs[i1].VrfTemplateConfig.BgpPassword = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.BgpPasswordType.IsNull() && !ele1.BgpPasswordType.IsUnknown() {
				data.Vrfs[i1].VrfTemplateConfig.BgpPasswordType = ele1.BgpPasswordType.ValueString()
			} else {
				data.Vrfs[i1].VrfTemplateConfig.BgpPasswordType = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.Netflow.IsNull() && !ele1.Netflow.IsUnknown() { // test signature1  NDFCType:
				data.Vrfs[i1].VrfTemplateConfig.Netflow = strconv.FormatBool(ele1.Netflow.ValueBool())
			} else {
				data.Vrfs[i1].VrfTemplateConfig.Netflow = ""
			}

			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.NetflowMonitor.IsNull() && !ele1.NetflowMonitor.IsUnknown() {
				data.Vrfs[i1].VrfTemplateConfig.NetflowMonitor = ele1.NetflowMonitor.ValueString()
			} else {
				data.Vrfs[i1].VrfTemplateConfig.NetflowMonitor = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.DisableRtAuto.IsNull() && !ele1.DisableRtAuto.IsUnknown() { // test signature1  NDFCType:
				data.Vrfs[i1].VrfTemplateConfig.DisableRtAuto = strconv.FormatBool(ele1.DisableRtAuto.ValueBool())
			} else {
				data.Vrfs[i1].VrfTemplateConfig.DisableRtAuto = ""
			}

			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.RouteTargetImport.IsNull() && !ele1.RouteTargetImport.IsUnknown() {
				data.Vrfs[i1].VrfTemplateConfig.RouteTargetImport = ele1.RouteTargetImport.ValueString()
			} else {
				data.Vrfs[i1].VrfTemplateConfig.RouteTargetImport = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.RouteTargetExport.IsNull() && !ele1.RouteTargetExport.IsUnknown() {
				data.Vrfs[i1].VrfTemplateConfig.RouteTargetExport = ele1.RouteTargetExport.ValueString()
			} else {
				data.Vrfs[i1].VrfTemplateConfig.RouteTargetExport = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.RouteTargetImportEvpn.IsNull() && !ele1.RouteTargetImportEvpn.IsUnknown() {
				data.Vrfs[i1].VrfTemplateConfig.RouteTargetImportEvpn = ele1.RouteTargetImportEvpn.ValueString()
			} else {
				data.Vrfs[i1].VrfTemplateConfig.RouteTargetImportEvpn = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.RouteTargetExportEvpn.IsNull() && !ele1.RouteTargetExportEvpn.IsUnknown() {
				data.Vrfs[i1].VrfTemplateConfig.RouteTargetExportEvpn = ele1.RouteTargetExportEvpn.ValueString()
			} else {
				data.Vrfs[i1].VrfTemplateConfig.RouteTargetExportEvpn = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.RouteTargetImportMvpn.IsNull() && !ele1.RouteTargetImportMvpn.IsUnknown() {
				data.Vrfs[i1].VrfTemplateConfig.RouteTargetImportMvpn = ele1.RouteTargetImportMvpn.ValueString()
			} else {
				data.Vrfs[i1].VrfTemplateConfig.RouteTargetImportMvpn = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.RouteTargetExportMvpn.IsNull() && !ele1.RouteTargetExportMvpn.IsUnknown() {
				data.Vrfs[i1].VrfTemplateConfig.RouteTargetExportMvpn = ele1.RouteTargetExportMvpn.ValueString()
			} else {
				data.Vrfs[i1].VrfTemplateConfig.RouteTargetExportMvpn = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.RouteTargetImportCloudEvpn.IsNull() && !ele1.RouteTargetImportCloudEvpn.IsUnknown() {
				data.Vrfs[i1].VrfTemplateConfig.RouteTargetImportCloudEvpn = ele1.RouteTargetImportCloudEvpn.ValueString()
			} else {
				data.Vrfs[i1].VrfTemplateConfig.RouteTargetImportCloudEvpn = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.RouteTargetExportCloudEvpn.IsNull() && !ele1.RouteTargetExportCloudEvpn.IsUnknown() {
				data.Vrfs[i1].VrfTemplateConfig.RouteTargetExportCloudEvpn = ele1.RouteTargetExportCloudEvpn.ValueString()
			} else {
				data.Vrfs[i1].VrfTemplateConfig.RouteTargetExportCloudEvpn = ""
			}
			//-----inline nesting end----

		}
	}

	return data
}
