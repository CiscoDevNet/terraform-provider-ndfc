// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

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
	SerialNumber   string            `json:"serialNumber,omitempty"`
	SwitchSerialNo string            `json:"switchSerialNo,omitempty"`
	SwitchName     string            `json:"switchName,omitempty"`
	DisplayName    string            `json:"displayName,omitempty"`
	Vlan           *Int64Custom      `json:"vlan,omitempty"`
	VlanId         *Int64Custom      `json:"vlanId,omitempty"`
	AttachState    string            `json:"lanAttachState,omitempty"`
	Attached       *bool             `json:"isLanAttached,omitempty"`
	FreeformConfig string            `json:"freeformconfig,omitempty"`
	SwitchPorts    CSVString         `json:"switchPorts,omitempty"`
	PortNames      string            `json:"portNames,omitempty"`
	TorPorts       CSVString         `json:"torPorts,omitempty"`
	InstanceValues map[string]string `json:"instanceValues,omitempty"`
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
	IgmpVersion          string                     `json:"igmpVersion,omitempty"`
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

	if len(jsonData.InstanceValues) == 0 {
		log.Printf("v.InstanceValues is empty")
		v.InstanceValues = types.MapNull(types.StringType)
	} else {
		mapData := make(map[string]attr.Value)
		for key, item := range jsonData.InstanceValues {
			mapData[key] = types.StringValue(item)
		}
		v.InstanceValues, err = types.MapValue(types.StringType, mapData)
		if err != nil {
			log.Printf("Error in converting map[string]string to  Map")
			return err
		}
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

	return data
}
