// Code generated;  DO NOT EDIT.

package datasource_networks

import (
	"context"
	"log"
	"strconv"
	. "terraform-provider-ndfc/internal/provider/types"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCNetworksModel struct {
	FabricName  string                        `json:"fabric,omitempty"`
	Networks    NDFCNetworksValues            `json:"networks,omitempty"`
	NetworksMap map[string]*NDFCNetworksValue `json:"-"`
}

type NDFCNetworksValues []NDFCNetworksValue

type NDFCNetworksValue struct {
	NetworkName              string                         `json:"networkName,omitempty"`
	DisplayName              string                         `json:"displayName,omitempty"`
	NetworkId                *int64                         `json:"networkId,omitempty"`
	NetworkTemplate          string                         `json:"networkTemplate,omitempty"`
	NetworkExtensionTemplate string                         `json:"networkExtensionTemplate,omitempty"`
	VrfName                  string                         `json:"vrf,omitempty"`
	PrimaryNetworkId         *Int64Custom                   `json:"primaryNetworkId,omitempty"`
	NetworkType              string                         `json:"type,omitempty"`
	NetworkStatus            string                         `json:"networkStatus,omitempty"`
	Attachments              NDFCAttachmentsValues          `json:"lanAttachList,omitempty"`
	NetworkTemplateConfig    NDFCNetworkTemplateConfigValue `json:"networkTemplateConfig,omitempty"`
}

type NDFCAttachmentsValues []NDFCAttachmentsValue

type NDFCAttachmentsValue struct {
	SerialNumber   string       `json:"serialNumber,omitempty"`
	SwitchSerialNo string       `json:"switchSerialNo,omitempty"`
	SwitchName     string       `json:"switchName,omitempty"`
	DisplayName    string       `json:"displayName,omitempty"`
	Vlan           *Int64Custom `json:"vlan,omitempty"`
	VlanId         *Int64Custom `json:"vlanId,omitempty"`
	AttachState    string       `json:"lanAttachState,omitempty"`
	Attached       *bool        `json:"isLanAttached,omitempty"`
	FreeformConfig string       `json:"freeformconfig,omitempty"`
	SwitchPorts    CSVString    `json:"switchPorts,omitempty"`
	PortNames      string       `json:"portNames,omitempty"`
	TorPorts       CSVString    `json:"torPorts,omitempty"`
	InstanceValues string       `json:"instanceValues,omitempty"`
}

type NDFCNetworkTemplateConfigValue struct {
	GatewayIpv4Address   string                     `json:"gatewayIpAddress,omitempty"`
	GatewayIpv6Address   string                     `json:"gatewayIpV6Address,omitempty"`
	VlanId               *Int64Custom               `json:"vlanId,omitempty"`
	VlanName             string                     `json:"vlanName,omitempty"`
	Layer2Only           string                     `json:"isLayer2Only,omitempty"`
	InterfaceDescription string                     `json:"intfDescription,omitempty"`
	Mtu                  *Int64Custom               `json:"mtu,omitempty"`
	SecondaryGateway1    string                     `json:"secondaryGW1,omitempty"`
	SecondaryGateway2    string                     `json:"secondaryGW2,omitempty"`
	SecondaryGateway3    string                     `json:"secondaryGW3,omitempty"`
	SecondaryGateway4    string                     `json:"secondaryGW4,omitempty"`
	ArpSuppression       string                     `json:"suppressArp,omitempty"`
	IngressReplication   string                     `json:"enableIR,omitempty"`
	MulticastGroup       string                     `json:"mcastGroup,omitempty"`
	DhcpRelayServers     NDFCDhcpRelayServersValues `json:"dhcpServers,omitempty"`
	DhcpRelayLoopbackId  *Int64Custom               `json:"loopbackId,omitempty"`
	RoutingTag           *Int64Custom               `json:"tag,omitempty"`
	Trm                  string                     `json:"trmEnabled,omitempty"`
	RouteTargetBoth      string                     `json:"rtBothAuto,omitempty"`
	Netflow              string                     `json:"ENABLE_NETFLOW,omitempty"`
	SviNetflowMonitor    string                     `json:"SVI_NETFLOW_MONITOR,omitempty"`
	VlanNetflowMonitor   string                     `json:"VLAN_NETFLOW_MONITOR,omitempty"`
	L3GatwayBorder       string                     `json:"enableL3OnBorder,omitempty"`
	IgmpVersion          string                     `json:"igmpVerion,omitempty"`
}

type NDFCDhcpRelayServersValues []NDFCDhcpRelayServersValue

type NDFCDhcpRelayServersValue struct {
	Address string `json:"srvrAddr,omitempty"`
	Vrf     string `json:"srvrVrf,omitempty"`
}

func (v *NetworksModel) SetModelData(jsonData *NDFCNetworksModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.FabricName != "" {
		v.FabricName = types.StringValue(jsonData.FabricName)
	} else {
		v.FabricName = types.StringNull()
	}

	if len(jsonData.Networks) == 0 {
		log.Printf("v.Networks is empty")
		v.Networks = types.ListNull(NetworksValue{}.Type(context.Background()))
	} else {
		log.Printf("v.Networks contains %d elements", len(jsonData.Networks))
		listData := make([]NetworksValue, 0)
		for _, item := range jsonData.Networks {
			data := new(NetworksValue)
			err = data.SetValue(&item)
			if err != nil {
				log.Printf("Error in NetworksValue.SetValue")
				return err
			}
			data.state = attr.ValueStateKnown
			listData = append(listData, *data)
		}
		v.Networks, err = types.ListValueFrom(context.Background(), NetworksValue{}.Type(context.Background()), listData)
		if err != nil {
			log.Printf("Error in converting []NetworksValue to  List")
			return err
		}
	}

	return err
}

func (v *NetworksValue) SetValue(jsonData *NDFCNetworksValue) diag.Diagnostics {

	var err diag.Diagnostics
	err = nil

	if jsonData.NetworkName != "" {
		v.NetworkName = types.StringValue(jsonData.NetworkName)
	} else {
		v.NetworkName = types.StringNull()
	}

	if jsonData.DisplayName != "" {
		v.DisplayName = types.StringValue(jsonData.DisplayName)
	} else {
		v.DisplayName = types.StringNull()
	}

	if jsonData.NetworkId != nil {
		v.NetworkId = types.Int64Value(*jsonData.NetworkId)

	} else {
		v.NetworkId = types.Int64Null()
	}

	if jsonData.NetworkTemplate != "" {
		v.NetworkTemplate = types.StringValue(jsonData.NetworkTemplate)
	} else {
		v.NetworkTemplate = types.StringNull()
	}

	if jsonData.NetworkExtensionTemplate != "" {
		v.NetworkExtensionTemplate = types.StringValue(jsonData.NetworkExtensionTemplate)
	} else {
		v.NetworkExtensionTemplate = types.StringNull()
	}

	if jsonData.VrfName != "" {
		v.VrfName = types.StringValue(jsonData.VrfName)
	} else {
		v.VrfName = types.StringNull()
	}

	if jsonData.PrimaryNetworkId != nil {
		if jsonData.PrimaryNetworkId.IsEmpty() {
			v.PrimaryNetworkId = types.Int64Null()
		} else {
			v.PrimaryNetworkId = types.Int64Value(int64(*jsonData.PrimaryNetworkId))
		}
	} else {
		v.PrimaryNetworkId = types.Int64Null()
	}

	if jsonData.NetworkType != "" {
		v.NetworkType = types.StringValue(jsonData.NetworkType)
	} else {
		v.NetworkType = types.StringNull()
	}

	if jsonData.NetworkStatus != "" {
		v.NetworkStatus = types.StringValue(jsonData.NetworkStatus)
	} else {
		v.NetworkStatus = types.StringNull()
	}

	if jsonData.NetworkTemplateConfig.GatewayIpv4Address != "" {
		v.GatewayIpv4Address = types.StringValue(jsonData.NetworkTemplateConfig.GatewayIpv4Address)
	} else {
		v.GatewayIpv4Address = types.StringNull()
	}

	if jsonData.NetworkTemplateConfig.GatewayIpv6Address != "" {
		v.GatewayIpv6Address = types.StringValue(jsonData.NetworkTemplateConfig.GatewayIpv6Address)
	} else {
		v.GatewayIpv6Address = types.StringNull()
	}

	if jsonData.NetworkTemplateConfig.VlanId != nil {
		if jsonData.NetworkTemplateConfig.VlanId.IsEmpty() {
			v.VlanId = types.Int64Null()
		} else {
			v.VlanId = types.Int64Value(int64(*jsonData.NetworkTemplateConfig.VlanId))
		}

	} else {
		v.VlanId = types.Int64Null()
	}

	if jsonData.NetworkTemplateConfig.VlanName != "" {
		v.VlanName = types.StringValue(jsonData.NetworkTemplateConfig.VlanName)
	} else {
		v.VlanName = types.StringNull()
	}

	if jsonData.NetworkTemplateConfig.Layer2Only != "" {
		x, _ := strconv.ParseBool(jsonData.NetworkTemplateConfig.Layer2Only)
		v.Layer2Only = types.BoolValue(x)
	} else {
		v.Layer2Only = types.BoolNull()
	}

	if jsonData.NetworkTemplateConfig.InterfaceDescription != "" {
		v.InterfaceDescription = types.StringValue(jsonData.NetworkTemplateConfig.InterfaceDescription)
	} else {
		v.InterfaceDescription = types.StringNull()
	}

	if jsonData.NetworkTemplateConfig.Mtu != nil {
		if jsonData.NetworkTemplateConfig.Mtu.IsEmpty() {
			v.Mtu = types.Int64Null()
		} else {
			v.Mtu = types.Int64Value(int64(*jsonData.NetworkTemplateConfig.Mtu))
		}

	} else {
		v.Mtu = types.Int64Null()
	}

	if jsonData.NetworkTemplateConfig.SecondaryGateway1 != "" {
		v.SecondaryGateway1 = types.StringValue(jsonData.NetworkTemplateConfig.SecondaryGateway1)
	} else {
		v.SecondaryGateway1 = types.StringNull()
	}

	if jsonData.NetworkTemplateConfig.SecondaryGateway2 != "" {
		v.SecondaryGateway2 = types.StringValue(jsonData.NetworkTemplateConfig.SecondaryGateway2)
	} else {
		v.SecondaryGateway2 = types.StringNull()
	}

	if jsonData.NetworkTemplateConfig.SecondaryGateway3 != "" {
		v.SecondaryGateway3 = types.StringValue(jsonData.NetworkTemplateConfig.SecondaryGateway3)
	} else {
		v.SecondaryGateway3 = types.StringNull()
	}

	if jsonData.NetworkTemplateConfig.SecondaryGateway4 != "" {
		v.SecondaryGateway4 = types.StringValue(jsonData.NetworkTemplateConfig.SecondaryGateway4)
	} else {
		v.SecondaryGateway4 = types.StringNull()
	}

	if jsonData.NetworkTemplateConfig.ArpSuppression != "" {
		x, _ := strconv.ParseBool(jsonData.NetworkTemplateConfig.ArpSuppression)
		v.ArpSuppression = types.BoolValue(x)
	} else {
		v.ArpSuppression = types.BoolNull()
	}

	if jsonData.NetworkTemplateConfig.IngressReplication != "" {
		x, _ := strconv.ParseBool(jsonData.NetworkTemplateConfig.IngressReplication)
		v.IngressReplication = types.BoolValue(x)
	} else {
		v.IngressReplication = types.BoolNull()
	}

	if jsonData.NetworkTemplateConfig.MulticastGroup != "" {
		v.MulticastGroup = types.StringValue(jsonData.NetworkTemplateConfig.MulticastGroup)
	} else {
		v.MulticastGroup = types.StringNull()
	}

	if len(jsonData.NetworkTemplateConfig.DhcpRelayServers) == 0 {
		log.Printf("v.DhcpRelayServers is empty")
		v.DhcpRelayServers = types.ListNull(DhcpRelayServersValue{}.Type(context.Background()))
	} else {
		listData := make([]DhcpRelayServersValue, len(jsonData.NetworkTemplateConfig.DhcpRelayServers))
		for i, item := range jsonData.NetworkTemplateConfig.DhcpRelayServers {
			err = listData[i].SetValue(&item)
			if err != nil {
				return err
			}
			listData[i].state = attr.ValueStateKnown
		}
		v.DhcpRelayServers, err = types.ListValueFrom(context.Background(), DhcpRelayServersValue{}.Type(context.Background()), listData)

		if err != nil {
			return err
		}
	}

	if jsonData.NetworkTemplateConfig.DhcpRelayLoopbackId != nil {
		if jsonData.NetworkTemplateConfig.DhcpRelayLoopbackId.IsEmpty() {
			v.DhcpRelayLoopbackId = types.Int64Null()
		} else {
			v.DhcpRelayLoopbackId = types.Int64Value(int64(*jsonData.NetworkTemplateConfig.DhcpRelayLoopbackId))
		}

	} else {
		v.DhcpRelayLoopbackId = types.Int64Null()
	}

	if jsonData.NetworkTemplateConfig.RoutingTag != nil {
		if jsonData.NetworkTemplateConfig.RoutingTag.IsEmpty() {
			v.RoutingTag = types.Int64Null()
		} else {
			v.RoutingTag = types.Int64Value(int64(*jsonData.NetworkTemplateConfig.RoutingTag))
		}

	} else {
		v.RoutingTag = types.Int64Null()
	}

	if jsonData.NetworkTemplateConfig.Trm != "" {
		x, _ := strconv.ParseBool(jsonData.NetworkTemplateConfig.Trm)
		v.Trm = types.BoolValue(x)
	} else {
		v.Trm = types.BoolNull()
	}

	if jsonData.NetworkTemplateConfig.RouteTargetBoth != "" {
		x, _ := strconv.ParseBool(jsonData.NetworkTemplateConfig.RouteTargetBoth)
		v.RouteTargetBoth = types.BoolValue(x)
	} else {
		v.RouteTargetBoth = types.BoolNull()
	}

	if jsonData.NetworkTemplateConfig.Netflow != "" {
		x, _ := strconv.ParseBool(jsonData.NetworkTemplateConfig.Netflow)
		v.Netflow = types.BoolValue(x)
	} else {
		v.Netflow = types.BoolNull()
	}

	if jsonData.NetworkTemplateConfig.SviNetflowMonitor != "" {
		v.SviNetflowMonitor = types.StringValue(jsonData.NetworkTemplateConfig.SviNetflowMonitor)
	} else {
		v.SviNetflowMonitor = types.StringNull()
	}

	if jsonData.NetworkTemplateConfig.VlanNetflowMonitor != "" {
		v.VlanNetflowMonitor = types.StringValue(jsonData.NetworkTemplateConfig.VlanNetflowMonitor)
	} else {
		v.VlanNetflowMonitor = types.StringNull()
	}

	if jsonData.NetworkTemplateConfig.L3GatwayBorder != "" {
		x, _ := strconv.ParseBool(jsonData.NetworkTemplateConfig.L3GatwayBorder)
		v.L3GatwayBorder = types.BoolValue(x)
	} else {
		v.L3GatwayBorder = types.BoolNull()
	}

	if jsonData.NetworkTemplateConfig.IgmpVersion != "" {
		v.IgmpVersion = types.StringValue(jsonData.NetworkTemplateConfig.IgmpVersion)
	} else {
		v.IgmpVersion = types.StringNull()
	}

	if len(jsonData.Attachments) == 0 {
		log.Printf("v.Attachments is empty")
		v.Attachments = types.ListNull(AttachmentsValue{}.Type(context.Background()))
	} else {
		log.Printf("v.Attachments contains %d elements", len(jsonData.Attachments))
		listData := make([]AttachmentsValue, 0)
		for _, item := range jsonData.Attachments {
			data := new(AttachmentsValue)
			err = data.SetValue(&item)
			if err != nil {
				log.Printf("Error in AttachmentsValue.SetValue")
				return err
			}
			data.state = attr.ValueStateKnown
			listData = append(listData, *data)
		}
		v.Attachments, err = types.ListValueFrom(context.Background(), AttachmentsValue{}.Type(context.Background()), listData)
		if err != nil {
			log.Printf("Error in converting []AttachmentsValue to  List")
			return err
		}
	}

	return err
}

func (v *DhcpRelayServersValue) SetValue(jsonData *NDFCDhcpRelayServersValue) diag.Diagnostics {

	var err diag.Diagnostics
	err = nil

	if jsonData.Address != "" {
		v.Address = types.StringValue(jsonData.Address)
	} else {
		v.Address = types.StringNull()
	}

	if jsonData.Vrf != "" {
		v.Vrf = types.StringValue(jsonData.Vrf)
	} else {
		v.Vrf = types.StringNull()
	}

	return err
}

func (v *AttachmentsValue) SetValue(jsonData *NDFCAttachmentsValue) diag.Diagnostics {

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

	if jsonData.DisplayName != "" {
		v.DisplayName = types.StringValue(jsonData.DisplayName)
	} else {
		v.DisplayName = types.StringNull()
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

	if len(jsonData.SwitchPorts) == 0 {
		log.Printf("v.SwitchPorts is empty")
		v.SwitchPorts, err = types.SetValue(types.StringType, []attr.Value{})
		if err != nil {
			log.Printf("Error in converting []string to  List %v", err)
			return err
		}
	} else {
		listData := make([]attr.Value, len(jsonData.SwitchPorts))
		for i, item := range jsonData.SwitchPorts {
			listData[i] = types.StringValue(item)
		}
		v.SwitchPorts, err = types.SetValue(types.StringType, listData)
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}

	if len(jsonData.TorPorts) == 0 {
		log.Printf("v.TorPorts is empty")
		v.TorPorts, err = types.SetValue(types.StringType, []attr.Value{})
		if err != nil {
			log.Printf("Error in converting []string to  List %v", err)
			return err
		}
	} else {
		listData := make([]attr.Value, len(jsonData.TorPorts))
		for i, item := range jsonData.TorPorts {
			listData[i] = types.StringValue(item)
		}
		v.TorPorts, err = types.SetValue(types.StringType, listData)
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}
	if jsonData.InstanceValues != "" {
		v.InstanceValues = types.StringValue(jsonData.InstanceValues)
	} else {
		v.InstanceValues = types.StringNull()
	}

	return err
}

func (v NetworksModel) GetModelData() *NDFCNetworksModel {
	var data = new(NDFCNetworksModel)

	//MARSHAL_BODY

	if !v.FabricName.IsNull() && !v.FabricName.IsUnknown() {
		data.FabricName = v.FabricName.ValueString()
	} else {
		data.FabricName = ""
	}

	//MARSHALL_LIST

	data.NetworksMap = make(map[string]*NDFCNetworksValue)

	if !v.Networks.IsNull() && !v.Networks.IsUnknown() {
		elements := make([]NetworksValue, len(v.Networks.Elements()))
		data.Networks = make([]NDFCNetworksValue, len(v.Networks.Elements()))

		diag := v.Networks.ElementsAs(context.Background(), &elements, false)
		if diag != nil {
			panic(diag)
		}
		for i1, ele1 := range elements {
			if !ele1.NetworkName.IsNull() && !ele1.NetworkName.IsUnknown() {

				data.Networks[i1].NetworkName = ele1.NetworkName.ValueString()
			} else {
				data.Networks[i1].NetworkName = ""
			}

			if !ele1.DisplayName.IsNull() && !ele1.DisplayName.IsUnknown() {

				data.Networks[i1].DisplayName = ele1.DisplayName.ValueString()
			} else {
				data.Networks[i1].DisplayName = ""
			}

			if !ele1.NetworkId.IsNull() && !ele1.NetworkId.IsUnknown() {

				data.Networks[i1].NetworkId = new(int64)
				*data.Networks[i1].NetworkId = ele1.NetworkId.ValueInt64()

			} else {
				data.Networks[i1].NetworkId = nil
			}

			if !ele1.NetworkTemplate.IsNull() && !ele1.NetworkTemplate.IsUnknown() {

				data.Networks[i1].NetworkTemplate = ele1.NetworkTemplate.ValueString()
			} else {
				data.Networks[i1].NetworkTemplate = ""
			}

			if !ele1.NetworkExtensionTemplate.IsNull() && !ele1.NetworkExtensionTemplate.IsUnknown() {

				data.Networks[i1].NetworkExtensionTemplate = ele1.NetworkExtensionTemplate.ValueString()
			} else {
				data.Networks[i1].NetworkExtensionTemplate = ""
			}

			if !ele1.VrfName.IsNull() && !ele1.VrfName.IsUnknown() {

				data.Networks[i1].VrfName = ele1.VrfName.ValueString()
			} else {
				data.Networks[i1].VrfName = ""
			}

			if !ele1.PrimaryNetworkId.IsNull() && !ele1.PrimaryNetworkId.IsUnknown() {

				data.Networks[i1].PrimaryNetworkId = new(Int64Custom)
				*data.Networks[i1].PrimaryNetworkId = Int64Custom(ele1.PrimaryNetworkId.ValueInt64())
			} else {
				data.Networks[i1].PrimaryNetworkId = nil
			}

			if !ele1.NetworkType.IsNull() && !ele1.NetworkType.IsUnknown() {

				data.Networks[i1].NetworkType = ele1.NetworkType.ValueString()
			} else {
				data.Networks[i1].NetworkType = ""
			}

			//-----inline nesting Start----
			if !ele1.GatewayIpv4Address.IsNull() && !ele1.GatewayIpv4Address.IsUnknown() {
				data.Networks[i1].NetworkTemplateConfig.GatewayIpv4Address = ele1.GatewayIpv4Address.ValueString()
			} else {
				data.Networks[i1].NetworkTemplateConfig.GatewayIpv4Address = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.GatewayIpv6Address.IsNull() && !ele1.GatewayIpv6Address.IsUnknown() {
				data.Networks[i1].NetworkTemplateConfig.GatewayIpv6Address = ele1.GatewayIpv6Address.ValueString()
			} else {
				data.Networks[i1].NetworkTemplateConfig.GatewayIpv6Address = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.VlanId.IsNull() && !ele1.VlanId.IsUnknown() {
				data.Networks[i1].NetworkTemplateConfig.VlanId = new(Int64Custom)
				*data.Networks[i1].NetworkTemplateConfig.VlanId = Int64Custom(ele1.VlanId.ValueInt64())
			} else {
				data.Networks[i1].NetworkTemplateConfig.VlanId = nil
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.VlanName.IsNull() && !ele1.VlanName.IsUnknown() {
				data.Networks[i1].NetworkTemplateConfig.VlanName = ele1.VlanName.ValueString()
			} else {
				data.Networks[i1].NetworkTemplateConfig.VlanName = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.Layer2Only.IsNull() && !ele1.Layer2Only.IsUnknown() {
				// test signature1  NDFCType:
				data.Networks[i1].NetworkTemplateConfig.Layer2Only = strconv.FormatBool(ele1.Layer2Only.ValueBool())
			} else {
				data.Networks[i1].NetworkTemplateConfig.Layer2Only = ""
			}

			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.InterfaceDescription.IsNull() && !ele1.InterfaceDescription.IsUnknown() {
				data.Networks[i1].NetworkTemplateConfig.InterfaceDescription = ele1.InterfaceDescription.ValueString()
			} else {
				data.Networks[i1].NetworkTemplateConfig.InterfaceDescription = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.Mtu.IsNull() && !ele1.Mtu.IsUnknown() {
				data.Networks[i1].NetworkTemplateConfig.Mtu = new(Int64Custom)
				*data.Networks[i1].NetworkTemplateConfig.Mtu = Int64Custom(ele1.Mtu.ValueInt64())
			} else {
				data.Networks[i1].NetworkTemplateConfig.Mtu = nil
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.SecondaryGateway1.IsNull() && !ele1.SecondaryGateway1.IsUnknown() {
				data.Networks[i1].NetworkTemplateConfig.SecondaryGateway1 = ele1.SecondaryGateway1.ValueString()
			} else {
				data.Networks[i1].NetworkTemplateConfig.SecondaryGateway1 = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.SecondaryGateway2.IsNull() && !ele1.SecondaryGateway2.IsUnknown() {
				data.Networks[i1].NetworkTemplateConfig.SecondaryGateway2 = ele1.SecondaryGateway2.ValueString()
			} else {
				data.Networks[i1].NetworkTemplateConfig.SecondaryGateway2 = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.SecondaryGateway3.IsNull() && !ele1.SecondaryGateway3.IsUnknown() {
				data.Networks[i1].NetworkTemplateConfig.SecondaryGateway3 = ele1.SecondaryGateway3.ValueString()
			} else {
				data.Networks[i1].NetworkTemplateConfig.SecondaryGateway3 = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.SecondaryGateway4.IsNull() && !ele1.SecondaryGateway4.IsUnknown() {
				data.Networks[i1].NetworkTemplateConfig.SecondaryGateway4 = ele1.SecondaryGateway4.ValueString()
			} else {
				data.Networks[i1].NetworkTemplateConfig.SecondaryGateway4 = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.ArpSuppression.IsNull() && !ele1.ArpSuppression.IsUnknown() {
				// test signature1  NDFCType:
				data.Networks[i1].NetworkTemplateConfig.ArpSuppression = strconv.FormatBool(ele1.ArpSuppression.ValueBool())
			} else {
				data.Networks[i1].NetworkTemplateConfig.ArpSuppression = ""
			}

			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.IngressReplication.IsNull() && !ele1.IngressReplication.IsUnknown() {
				// test signature1  NDFCType:
				data.Networks[i1].NetworkTemplateConfig.IngressReplication = strconv.FormatBool(ele1.IngressReplication.ValueBool())
			} else {
				data.Networks[i1].NetworkTemplateConfig.IngressReplication = ""
			}

			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.MulticastGroup.IsNull() && !ele1.MulticastGroup.IsUnknown() {
				data.Networks[i1].NetworkTemplateConfig.MulticastGroup = ele1.MulticastGroup.ValueString()
			} else {
				data.Networks[i1].NetworkTemplateConfig.MulticastGroup = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----

			//MARSHALL_LIST

			if !ele1.DhcpRelayServers.IsNull() && !ele1.DhcpRelayServers.IsUnknown() {
				elements := make([]DhcpRelayServersValue, len(ele1.DhcpRelayServers.Elements()))
				data.Networks[i1].NetworkTemplateConfig.DhcpRelayServers = make([]NDFCDhcpRelayServersValue, len(ele1.DhcpRelayServers.Elements()))

				diag := ele1.DhcpRelayServers.ElementsAs(context.Background(), &elements, false)
				if diag != nil {
					panic(diag)
				}
				for i2, ele2 := range elements {
					if !ele2.Address.IsNull() && !ele2.Address.IsUnknown() {

						data.Networks[i1].NetworkTemplateConfig.DhcpRelayServers[i2].Address = ele2.Address.ValueString()
					} else {
						data.Networks[i1].NetworkTemplateConfig.DhcpRelayServers[i2].Address = ""
					}

					if !ele2.Vrf.IsNull() && !ele2.Vrf.IsUnknown() {

						data.Networks[i1].NetworkTemplateConfig.DhcpRelayServers[i2].Vrf = ele2.Vrf.ValueString()
					} else {
						data.Networks[i1].NetworkTemplateConfig.DhcpRelayServers[i2].Vrf = ""
					}

				}
			}

			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.DhcpRelayLoopbackId.IsNull() && !ele1.DhcpRelayLoopbackId.IsUnknown() {
				data.Networks[i1].NetworkTemplateConfig.DhcpRelayLoopbackId = new(Int64Custom)
				*data.Networks[i1].NetworkTemplateConfig.DhcpRelayLoopbackId = Int64Custom(ele1.DhcpRelayLoopbackId.ValueInt64())
			} else {
				data.Networks[i1].NetworkTemplateConfig.DhcpRelayLoopbackId = nil
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.RoutingTag.IsNull() && !ele1.RoutingTag.IsUnknown() {
				data.Networks[i1].NetworkTemplateConfig.RoutingTag = new(Int64Custom)
				*data.Networks[i1].NetworkTemplateConfig.RoutingTag = Int64Custom(ele1.RoutingTag.ValueInt64())
			} else {
				data.Networks[i1].NetworkTemplateConfig.RoutingTag = nil
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.Trm.IsNull() && !ele1.Trm.IsUnknown() {
				// test signature1  NDFCType:
				data.Networks[i1].NetworkTemplateConfig.Trm = strconv.FormatBool(ele1.Trm.ValueBool())
			} else {
				data.Networks[i1].NetworkTemplateConfig.Trm = ""
			}

			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.RouteTargetBoth.IsNull() && !ele1.RouteTargetBoth.IsUnknown() {
				// test signature1  NDFCType:
				data.Networks[i1].NetworkTemplateConfig.RouteTargetBoth = strconv.FormatBool(ele1.RouteTargetBoth.ValueBool())
			} else {
				data.Networks[i1].NetworkTemplateConfig.RouteTargetBoth = ""
			}

			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.Netflow.IsNull() && !ele1.Netflow.IsUnknown() {
				// test signature1  NDFCType:
				data.Networks[i1].NetworkTemplateConfig.Netflow = strconv.FormatBool(ele1.Netflow.ValueBool())
			} else {
				data.Networks[i1].NetworkTemplateConfig.Netflow = ""
			}

			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.SviNetflowMonitor.IsNull() && !ele1.SviNetflowMonitor.IsUnknown() {
				data.Networks[i1].NetworkTemplateConfig.SviNetflowMonitor = ele1.SviNetflowMonitor.ValueString()
			} else {
				data.Networks[i1].NetworkTemplateConfig.SviNetflowMonitor = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.VlanNetflowMonitor.IsNull() && !ele1.VlanNetflowMonitor.IsUnknown() {
				data.Networks[i1].NetworkTemplateConfig.VlanNetflowMonitor = ele1.VlanNetflowMonitor.ValueString()
			} else {
				data.Networks[i1].NetworkTemplateConfig.VlanNetflowMonitor = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.L3GatwayBorder.IsNull() && !ele1.L3GatwayBorder.IsUnknown() {
				// test signature1  NDFCType:
				data.Networks[i1].NetworkTemplateConfig.L3GatwayBorder = strconv.FormatBool(ele1.L3GatwayBorder.ValueBool())
			} else {
				data.Networks[i1].NetworkTemplateConfig.L3GatwayBorder = ""
			}

			//-----inline nesting end----

			//-----inline nesting Start----
			if !ele1.IgmpVersion.IsNull() && !ele1.IgmpVersion.IsUnknown() {
				data.Networks[i1].NetworkTemplateConfig.IgmpVersion = ele1.IgmpVersion.ValueString()
			} else {
				data.Networks[i1].NetworkTemplateConfig.IgmpVersion = ""
			}
			//-----inline nesting end----

			//MARSHALL_LIST

			if !ele1.Attachments.IsNull() && !ele1.Attachments.IsUnknown() {
				elements := make([]AttachmentsValue, len(ele1.Attachments.Elements()))
				data.Networks[i1].Attachments = make([]NDFCAttachmentsValue, len(ele1.Attachments.Elements()))

				diag := ele1.Attachments.ElementsAs(context.Background(), &elements, false)
				if diag != nil {
					panic(diag)
				}
				for i2, ele2 := range elements {
					if !ele2.SerialNumber.IsNull() && !ele2.SerialNumber.IsUnknown() {

						data.Networks[i1].Attachments[i2].SerialNumber = ele2.SerialNumber.ValueString()
					} else {
						data.Networks[i1].Attachments[i2].SerialNumber = ""
					}

					if !ele2.DisplayName.IsNull() && !ele2.DisplayName.IsUnknown() {

						data.Networks[i1].Attachments[i2].DisplayName = ele2.DisplayName.ValueString()
					} else {
						data.Networks[i1].Attachments[i2].DisplayName = ""
					}

					if !ele2.Vlan.IsNull() && !ele2.Vlan.IsUnknown() {

						data.Networks[i1].Attachments[i2].Vlan = new(Int64Custom)
						*data.Networks[i1].Attachments[i2].Vlan = Int64Custom(ele2.Vlan.ValueInt64())
					} else {
						data.Networks[i1].Attachments[i2].Vlan = nil
					}

					if !ele2.FreeformConfig.IsNull() && !ele2.FreeformConfig.IsUnknown() {

						data.Networks[i1].Attachments[i2].FreeformConfig = ele2.FreeformConfig.ValueString()
					} else {
						data.Networks[i1].Attachments[i2].FreeformConfig = ""
					}

					if !ele2.SwitchPorts.IsNull() && !ele2.SwitchPorts.IsUnknown() {

						listStringData := make([]string, len(ele2.SwitchPorts.Elements()))
						dg := ele2.SwitchPorts.ElementsAs(context.Background(), &listStringData, false)
						if dg.HasError() {
							panic(dg.Errors())
						}
						data.Networks[i1].Attachments[i2].SwitchPorts = make(CSVString, len(listStringData))
						copy(data.Networks[i1].Attachments[i2].SwitchPorts, listStringData)
					}

					if !ele2.TorPorts.IsNull() && !ele2.TorPorts.IsUnknown() {

						listStringData := make([]string, len(ele2.TorPorts.Elements()))
						dg := ele2.TorPorts.ElementsAs(context.Background(), &listStringData, false)
						if dg.HasError() {
							panic(dg.Errors())
						}
						data.Networks[i1].Attachments[i2].TorPorts = make(CSVString, len(listStringData))
						copy(data.Networks[i1].Attachments[i2].TorPorts, listStringData)
					}

					if !ele2.InstanceValues.IsNull() && !ele2.InstanceValues.IsUnknown() {

						data.Networks[i1].Attachments[i2].InstanceValues = ele2.InstanceValues.ValueString()
					} else {
						data.Networks[i1].Attachments[i2].InstanceValues = ""
					}

				}
			}

		}
	}

	return data
}
