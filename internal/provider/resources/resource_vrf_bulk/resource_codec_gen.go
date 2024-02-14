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
	FabricName           string                    `json:"fabric,omitempty"`
	DeployAllAttachments bool                      `json:"-"`
	Vrfs                 NDFCVrfsValues            `json:"vrfs,omitempty"`
	VrfsMap              map[string]*NDFCVrfsValue `json:"-"`
}

type NDFCVrfsValues []NDFCVrfsValue

type NDFCVrfsValue struct {
	Id                   *int64                                                   `json:"id,omitempty"`
	FilterThisValue      bool                                                     `json:"-"`
	VrfName              string                                                   `json:"vrfName,omitempty"`
	FabricName           string                                                   `json:"fabric,omitempty"`
	VrfTemplate          string                                                   `json:"vrfTemplate,omitempty"`
	VrfExtensionTemplate string                                                   `json:"vrfExtensionTemplate,omitempty"`
	VrfId                *int64                                                   `json:"vrfId,omitempty"`
	VrfStatus            string                                                   `json:"vrfStatus,omitempty"`
	DeployAttachments    bool                                                     `json:"-"`
	AttachList           resource_vrf_attachments.NDFCAttachListValues            `json:"lanAttachList,omitempty"`
	AttachListMap        map[string]*resource_vrf_attachments.NDFCAttachListValue `json:"-"`

	VrfTemplateConfig NDFCVrfTemplateConfigValue `json:"vrfTemplateConfig,omitempty"`
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

	v.DeployAllAttachments = types.BoolValue(jsonData.DeployAllAttachments)
	if len(jsonData.Vrfs) == 0 {
		log.Printf("v.Vrfs is empty")
		v.Vrfs = types.ListNull(VrfsValue{}.Type(context.Background()))
	} else {
		log.Printf("v.Vrfs contains %d elements", len(jsonData.Vrfs))
		listData := make([]VrfsValue, 0)
		for _, item := range jsonData.Vrfs {
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

	v.DeployAttachments = types.BoolValue(jsonData.DeployAttachments)
	if len(jsonData.AttachList) == 0 {
		log.Printf("v.AttachList is empty")
		v.AttachList = types.ListNull(AttachListValue{}.Type(context.Background()))
	} else {
		log.Printf("v.AttachList contains %d elements", len(jsonData.AttachList))
		listData := make([]AttachListValue, 0)
		for _, item := range jsonData.AttachList {
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

func (v *AttachListValue) SetValue(jsonData *resource_vrf_attachments.NDFCAttachListValue) diag.Diagnostics {

	var err diag.Diagnostics
	err = nil

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
		if int64(*jsonData.Vlan) == -9223372036854775808 {
			v.Vlan = types.Int64Null()
		} else {
			v.Vlan = types.Int64Value(int64(*jsonData.Vlan))
		}
	} else if jsonData.VlanId != nil {
		if int64(*jsonData.VlanId) == -9223372036854775808 {
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
		if int64(*jsonData.InstanceValues.LoopbackId) == -9223372036854775808 {
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
	if !v.FabricName.IsNull() && !v.FabricName.IsUnknown() {
		data.FabricName = v.FabricName.ValueString()
	} else {
		data.FabricName = ""
	}

	if !v.DeployAllAttachments.IsNull() && !v.DeployAllAttachments.IsUnknown() {
		data.DeployAllAttachments = v.DeployAllAttachments.ValueBool()
	}

	data.VrfsMap = make(map[string]*NDFCVrfsValue)

	if !v.Vrfs.IsNull() && !v.Vrfs.IsUnknown() {
		elements := make([]VrfsValue, len(v.Vrfs.Elements()))
		data.Vrfs = make([]NDFCVrfsValue, len(v.Vrfs.Elements()))

		diag := v.Vrfs.ElementsAs(context.Background(), &elements, false)
		if diag != nil {
			panic(diag)
		}
		for i1, ele1 := range elements {

			if !ele1.VrfName.IsNull() && !ele1.VrfName.IsUnknown() {
				data.Vrfs[i1].VrfName = ele1.VrfName.ValueString()
			} else {
				data.Vrfs[i1].VrfName = ""
			}

			data.VrfsMap[data.Vrfs[i1].VrfName] = &data.Vrfs[i1]

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

			if !ele1.VrfStatus.IsNull() && !ele1.VrfStatus.IsUnknown() {
				data.Vrfs[i1].VrfStatus = ele1.VrfStatus.ValueString()
			} else {
				data.Vrfs[i1].VrfStatus = ""
			}

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

			if !ele1.DeployAttachments.IsNull() && !ele1.DeployAttachments.IsUnknown() {

				data.Vrfs[i1].DeployAttachments = ele1.DeployAttachments.ValueBool()
			}

			data.Vrfs[i1].AttachListMap = make(map[string]*resource_vrf_attachments.NDFCAttachListValue)

			if !ele1.AttachList.IsNull() && !ele1.AttachList.IsUnknown() {
				elements := make([]AttachListValue, len(ele1.AttachList.Elements()))
				data.Vrfs[i1].AttachList = make([]resource_vrf_attachments.NDFCAttachListValue, len(ele1.AttachList.Elements()))
				diag := ele1.AttachList.ElementsAs(context.Background(), &elements, false)
				if diag != nil {
					panic(diag)
				}
				for i2, ele2 := range elements {

					if !ele2.SerialNumber.IsNull() && !ele2.SerialNumber.IsUnknown() {
						data.Vrfs[i1].AttachList[i2].SerialNumber = ele2.SerialNumber.ValueString()
					} else {
						data.Vrfs[i1].AttachList[i2].SerialNumber = ""
					}

					data.Vrfs[i1].AttachListMap[data.Vrfs[i1].AttachList[i2].SerialNumber] = &data.Vrfs[i1].AttachList[i2]

					if !ele2.SwitchName.IsNull() && !ele2.SwitchName.IsUnknown() {
						data.Vrfs[i1].AttachList[i2].SwitchName = ele2.SwitchName.ValueString()
					} else {
						data.Vrfs[i1].AttachList[i2].SwitchName = ""
					}

					if !ele2.Vlan.IsNull() && !ele2.Vlan.IsUnknown() {
						data.Vrfs[i1].AttachList[i2].Vlan = new(Int64Custom)
						*data.Vrfs[i1].AttachList[i2].Vlan = Int64Custom(ele2.Vlan.ValueInt64())
					} else {
						data.Vrfs[i1].AttachList[i2].Vlan = nil
					}

					if !ele2.AttachState.IsNull() && !ele2.AttachState.IsUnknown() {
						data.Vrfs[i1].AttachList[i2].AttachState = ele2.AttachState.ValueString()
					} else {
						data.Vrfs[i1].AttachList[i2].AttachState = ""
					}

					if !ele2.Attached.IsNull() && !ele2.Attached.IsUnknown() {

						data.Vrfs[i1].AttachList[i2].Attached = new(bool)
						*data.Vrfs[i1].AttachList[i2].Attached = ele2.Attached.ValueBool()
					} else {
						data.Vrfs[i1].AttachList[i2].Attached = nil
					}

					if !ele2.FreeformConfig.IsNull() && !ele2.FreeformConfig.IsUnknown() {
						data.Vrfs[i1].AttachList[i2].FreeformConfig = ele2.FreeformConfig.ValueString()
					} else {
						data.Vrfs[i1].AttachList[i2].FreeformConfig = ""
					}

					if !ele2.DeployThisAttachment.IsNull() && !ele2.DeployThisAttachment.IsUnknown() {

						data.Vrfs[i1].AttachList[i2].DeployThisAttachment = ele2.DeployThisAttachment.ValueBool()
					}

					//-----inline nesting Start----
					if !ele2.LoopbackId.IsNull() && !ele2.LoopbackId.IsUnknown() {
						data.Vrfs[i1].AttachList[i2].InstanceValues.LoopbackId = new(Int64Custom)
						*data.Vrfs[i1].AttachList[i2].InstanceValues.LoopbackId = Int64Custom(ele2.LoopbackId.ValueInt64())
					} else {
						data.Vrfs[i1].AttachList[i2].InstanceValues.LoopbackId = nil
					}
					//-----inline nesting end----

					//-----inline nesting Start----
					if !ele2.LoopbackIpv4.IsNull() && !ele2.LoopbackIpv4.IsUnknown() {
						data.Vrfs[i1].AttachList[i2].InstanceValues.LoopbackIpv4 = ele2.LoopbackIpv4.ValueString()
					} else {
						data.Vrfs[i1].AttachList[i2].InstanceValues.LoopbackIpv4 = ""
					}
					//-----inline nesting end----

					//-----inline nesting Start----
					if !ele2.LoopbackIpv6.IsNull() && !ele2.LoopbackIpv6.IsUnknown() {
						data.Vrfs[i1].AttachList[i2].InstanceValues.LoopbackIpv6 = ele2.LoopbackIpv6.ValueString()
					} else {
						data.Vrfs[i1].AttachList[i2].InstanceValues.LoopbackIpv6 = ""
					}
					//-----inline nesting end----

				}
			}

		}
	}

	return data
}
