// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package resource_interface_ethernet

import (
	"context"
	"log"
	"strconv"
	. "terraform-provider-ndfc/internal/provider/types"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"terraform-provider-ndfc/internal/provider/resources/resource_interface_common"
)

func (v *InterfaceEthernetModel) SetModelData(jsonData *resource_interface_common.NDFCInterfaceCommonModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.Policy != "" {
		v.Policy = types.StringValue(jsonData.Policy)
	} else {
		v.Policy = types.StringNull()
	}

	if jsonData.PolicyType != "" {
		v.PolicyType = types.StringValue(jsonData.PolicyType)
	} else {
		v.PolicyType = types.StringNull()
	}

	v.Deploy = types.BoolValue(jsonData.Deploy)
	if jsonData.SerialNumber != "" {
		v.SerialNumber = types.StringValue(jsonData.SerialNumber)
	} else {
		v.SerialNumber = types.StringNull()
	}

	if len(jsonData.Interfaces) == 0 {
		log.Printf("v.Interfaces is empty")
		v.Interfaces = types.MapNull(InterfacesValue{}.Type(context.Background()))
	} else {
		mapData := make(map[string]InterfacesValue)
		for key, item := range jsonData.Interfaces {
			if item.FilterThisValue {
				//Skip this entry - this parameter allows filtering
				continue
			}

			data := new(InterfacesValue)
			err = data.SetValue(&item)
			if err != nil {
				log.Printf("Error in InterfacesValue.SetValue")
				return err
			}
			data.state = attr.ValueStateKnown
			mapData[key] = *data
		}
		v.Interfaces, err = types.MapValueFrom(context.Background(), InterfacesValue{}.Type(context.Background()), mapData)
		if err != nil {
			log.Printf("Error in converting map[string]InterfacesValue to  Map")

		}
	}

	return err
}

func (v *InterfacesValue) SetValue(jsonData *resource_interface_common.NDFCInterfacesValue) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.PortChannelPolicy != "" {
		v.PortChannelPolicy = types.StringValue(jsonData.PortChannelPolicy)
	} else {
		v.PortChannelPolicy = types.StringNull()
	}

	if jsonData.NvPairs.PortChannelName != "" {
		v.PortChannelName = types.StringValue(jsonData.NvPairs.PortChannelName)
	} else {
		v.PortChannelName = types.StringNull()
	}

	if jsonData.SerialNumber != "" {
		v.SerialNumber = types.StringValue(jsonData.SerialNumber)
	} else {
		v.SerialNumber = types.StringNull()
	}

	if jsonData.InterfaceName != "" {
		v.InterfaceName = types.StringValue(jsonData.InterfaceName)
	} else {
		v.InterfaceName = types.StringNull()
	}

	if jsonData.NvPairs.FreeformConfig != "" {
		v.FreeformConfig = types.StringValue(jsonData.NvPairs.FreeformConfig)
	} else {
		v.FreeformConfig = types.StringNull()
	}

	if jsonData.NvPairs.AdminState != "" {
		x, _ := strconv.ParseBool(jsonData.NvPairs.AdminState)
		v.AdminState = types.BoolValue(x)
	} else {
		v.AdminState = types.BoolNull()
	}

	if jsonData.NvPairs.InterfaceDescription != "" {
		v.InterfaceDescription = types.StringValue(jsonData.NvPairs.InterfaceDescription)
	} else {
		v.InterfaceDescription = types.StringNull()
	}

	if jsonData.NvPairs.BpduGuard != "" {
		v.BpduGuard = types.StringValue(jsonData.NvPairs.BpduGuard)
	} else {
		v.BpduGuard = types.StringNull()
	}

	if jsonData.NvPairs.PortTypeFast != "" {
		x, _ := strconv.ParseBool(jsonData.NvPairs.PortTypeFast)
		v.PortTypeFast = types.BoolValue(x)
	} else {
		v.PortTypeFast = types.BoolNull()
	}

	if jsonData.NvPairs.Mtu != "" {
		v.Mtu = types.StringValue(jsonData.NvPairs.Mtu)
	} else {
		v.Mtu = types.StringNull()
	}

	if jsonData.NvPairs.Speed != "" {
		v.Speed = types.StringValue(jsonData.NvPairs.Speed)
	} else {
		v.Speed = types.StringNull()
	}

	if jsonData.NvPairs.AccessVlan != nil {
		if jsonData.NvPairs.AccessVlan.IsEmpty() {
			v.AccessVlan = types.Int64Null()
		} else {
			v.AccessVlan = types.Int64Value(int64(*jsonData.NvPairs.AccessVlan))
		}

	} else {
		v.AccessVlan = types.Int64Null()
	}

	if jsonData.NvPairs.OrphanPort != "" {
		x, _ := strconv.ParseBool(jsonData.NvPairs.OrphanPort)
		v.OrphanPort = types.BoolValue(x)
	} else {
		v.OrphanPort = types.BoolNull()
	}

	if jsonData.NvPairs.Ptp != "" {
		x, _ := strconv.ParseBool(jsonData.NvPairs.Ptp)
		v.Ptp = types.BoolValue(x)
	} else {
		v.Ptp = types.BoolNull()
	}

	if jsonData.NvPairs.Netflow != "" {
		x, _ := strconv.ParseBool(jsonData.NvPairs.Netflow)
		v.Netflow = types.BoolValue(x)
	} else {
		v.Netflow = types.BoolNull()
	}

	if jsonData.NvPairs.NetflowMonitor != "" {
		v.NetflowMonitor = types.StringValue(jsonData.NvPairs.NetflowMonitor)
	} else {
		v.NetflowMonitor = types.StringNull()
	}

	if jsonData.NvPairs.NetflowSampler != "" {
		v.NetflowSampler = types.StringValue(jsonData.NvPairs.NetflowSampler)
	} else {
		v.NetflowSampler = types.StringNull()
	}

	if jsonData.NvPairs.AllowedVlans != "" {
		v.AllowedVlans = types.StringValue(jsonData.NvPairs.AllowedVlans)
	} else {
		v.AllowedVlans = types.StringNull()
	}

	if jsonData.NvPairs.NativeVlan != nil {
		if jsonData.NvPairs.NativeVlan.IsEmpty() {
			v.NativeVlan = types.Int64Null()
		} else {
			v.NativeVlan = types.Int64Value(int64(*jsonData.NvPairs.NativeVlan))
		}

	} else {
		v.NativeVlan = types.Int64Null()
	}

	if jsonData.NvPairs.Vrf != "" {
		v.Vrf = types.StringValue(jsonData.NvPairs.Vrf)
	} else {
		v.Vrf = types.StringNull()
	}

	if jsonData.NvPairs.Ipv4Address != "" {
		v.Ipv4Address = types.StringValue(jsonData.NvPairs.Ipv4Address)
	} else {
		v.Ipv4Address = types.StringNull()
	}

	if jsonData.NvPairs.Ipv4PrefixLength != "" {
		v.Ipv4PrefixLength = types.StringValue(jsonData.NvPairs.Ipv4PrefixLength)
	} else {
		v.Ipv4PrefixLength = types.StringNull()
	}

	if jsonData.NvPairs.RoutingTag != "" {
		v.RoutingTag = types.StringValue(jsonData.NvPairs.RoutingTag)
	} else {
		v.RoutingTag = types.StringNull()
	}

	if jsonData.NvPairs.DisableIpRedirects != "" {
		x, _ := strconv.ParseBool(jsonData.NvPairs.DisableIpRedirects)
		v.DisableIpRedirects = types.BoolValue(x)
	} else {
		v.DisableIpRedirects = types.BoolNull()
	}

	if jsonData.NvPairs.EnablePimSparse != "" {
		x, _ := strconv.ParseBool(jsonData.NvPairs.EnablePimSparse)
		v.EnablePimSparse = types.BoolValue(x)
	} else {
		v.EnablePimSparse = types.BoolNull()
	}

	if jsonData.NvPairs.PimDrPriority != "" {
		v.PimDrPriority = types.StringValue(jsonData.NvPairs.PimDrPriority)
	} else {
		v.PimDrPriority = types.StringNull()
	}

	if jsonData.NvPairs.EnablePfc != "" {
		x, _ := strconv.ParseBool(jsonData.NvPairs.EnablePfc)
		v.EnablePfc = types.BoolValue(x)
	} else {
		v.EnablePfc = types.BoolNull()
	}

	if jsonData.NvPairs.EnableQos != "" {
		x, _ := strconv.ParseBool(jsonData.NvPairs.EnableQos)
		v.EnableQos = types.BoolValue(x)
	} else {
		v.EnableQos = types.BoolNull()
	}

	if jsonData.NvPairs.QosPolicy != "" {
		v.QosPolicy = types.StringValue(jsonData.NvPairs.QosPolicy)
	} else {
		v.QosPolicy = types.StringNull()
	}

	if jsonData.NvPairs.QueuingPolicy != "" {
		v.QueuingPolicy = types.StringValue(jsonData.NvPairs.QueuingPolicy)
	} else {
		v.QueuingPolicy = types.StringNull()
	}

	if jsonData.NvPairs.LinkStateRoutingProtocol != "" {
		v.LinkStateRoutingProtocol = types.StringValue(jsonData.NvPairs.LinkStateRoutingProtocol)
	} else {
		v.LinkStateRoutingProtocol = types.StringNull()
	}

	if jsonData.NvPairs.LinkStateRoutingTag != "" {
		v.LinkStateRoutingTag = types.StringValue(jsonData.NvPairs.LinkStateRoutingTag)
	} else {
		v.LinkStateRoutingTag = types.StringNull()
	}

	if jsonData.NvPairs.Ipv6Addr != "" {
		v.Ipv6Addr = types.StringValue(jsonData.NvPairs.Ipv6Addr)
	} else {
		v.Ipv6Addr = types.StringNull()
	}

	if jsonData.NvPairs.Ipv6PrefixLength != "" {
		v.Ipv6PrefixLength = types.StringValue(jsonData.NvPairs.Ipv6PrefixLength)
	} else {
		v.Ipv6PrefixLength = types.StringNull()
	}

	if jsonData.NvPairs.CdpEnable != "" {
		x, _ := strconv.ParseBool(jsonData.NvPairs.CdpEnable)
		v.CdpEnable = types.BoolValue(x)
	} else {
		v.CdpEnable = types.BoolNull()
	}

	if jsonData.NvPairs.PortDuplexMode != "" {
		v.PortDuplexMode = types.StringValue(jsonData.NvPairs.PortDuplexMode)
	} else {
		v.PortDuplexMode = types.StringNull()
	}

	if jsonData.NvPairs.EnableMonitor != "" {
		x, _ := strconv.ParseBool(jsonData.NvPairs.EnableMonitor)
		v.EnableMonitor = types.BoolValue(x)
	} else {
		v.EnableMonitor = types.BoolNull()
	}

	if jsonData.NvPairs.PvlanMode != "" {
		v.PvlanMode = types.StringValue(jsonData.NvPairs.PvlanMode)
	} else {
		v.PvlanMode = types.StringNull()
	}

	if jsonData.NvPairs.PvlanAllowedVlans != "" {
		v.PvlanAllowedVlans = types.StringValue(jsonData.NvPairs.PvlanAllowedVlans)
	} else {
		v.PvlanAllowedVlans = types.StringNull()
	}

	if jsonData.NvPairs.PvlanNativeVlan != "" {
		v.PvlanNativeVlan = types.StringValue(jsonData.NvPairs.PvlanNativeVlan)
	} else {
		v.PvlanNativeVlan = types.StringNull()
	}

	if jsonData.NvPairs.IgForFex != "" {
		x, _ := strconv.ParseBool(jsonData.NvPairs.IgForFex)
		v.IgForFex = types.BoolValue(x)
	} else {
		v.IgForFex = types.BoolNull()
	}

	if jsonData.NvPairs.AutoNegotiate != "" {
		v.AutoNegotiate = types.StringValue(jsonData.NvPairs.AutoNegotiate)
	} else {
		v.AutoNegotiate = types.StringNull()
	}

	if jsonData.NvPairs.PathCost != nil {
		v.PathCost = types.Int64Value(*jsonData.NvPairs.PathCost)

	} else {
		v.PathCost = types.Int64Null()
	}

	if jsonData.NvPairs.GuardMode != "" {
		v.GuardMode = types.StringValue(jsonData.NvPairs.GuardMode)
	} else {
		v.GuardMode = types.StringNull()
	}

	if jsonData.NvPairs.Ttag != "" {
		x, _ := strconv.ParseBool(jsonData.NvPairs.Ttag)
		v.Ttag = types.BoolValue(x)
	} else {
		v.Ttag = types.BoolNull()
	}

	if len(jsonData.NvPairs.PVlanMappingList) == 0 {
		log.Printf("v.PVlanMappingList is empty")
		v.PVlanMappingList = types.ListNull(PVlanMappingListValue{}.Type(context.Background()))
	} else {
		listData := make([]PVlanMappingListValue, len(jsonData.NvPairs.PVlanMappingList))
		for i, item := range jsonData.NvPairs.PVlanMappingList {
			err = listData[i].SetValue(&item)
			if err != nil {
				return err
			}
			listData[i].state = attr.ValueStateKnown
		}
		v.PVlanMappingList, err = types.ListValueFrom(context.Background(), PVlanMappingListValue{}.Type(context.Background()), listData)

		if err != nil {
			return err
		}
	}
	if len(jsonData.NvPairs.PVlanAssocList) == 0 {
		log.Printf("v.PVlanAssocList is empty")
		v.PVlanAssocList = types.ListNull(PVlanAssocListValue{}.Type(context.Background()))
	} else {
		listData := make([]PVlanAssocListValue, len(jsonData.NvPairs.PVlanAssocList))
		for i, item := range jsonData.NvPairs.PVlanAssocList {
			err = listData[i].SetValue(&item)
			if err != nil {
				return err
			}
			listData[i].state = attr.ValueStateKnown
		}
		v.PVlanAssocList, err = types.ListValueFrom(context.Background(), PVlanAssocListValue{}.Type(context.Background()), listData)

		if err != nil {
			return err
		}
	}
	if jsonData.DeploymentStatus != "" {
		v.DeploymentStatus = types.StringValue(jsonData.DeploymentStatus)
	} else {
		v.DeploymentStatus = types.StringNull()
	}

	if len(jsonData.CustomPolicyParameters) == 0 {
		log.Printf("v.CustomPolicyParameters is empty")
		v.CustomPolicyParameters = types.MapNull(types.StringType)
	} else {
		mapData := make(map[string]attr.Value)
		for key, item := range jsonData.CustomPolicyParameters {
			mapData[key] = types.StringValue(item)
		}
		v.CustomPolicyParameters, err = types.MapValue(types.StringType, mapData)
		if err != nil {
			log.Printf("Error in converting map[string]string to  Map")
			return err
		}
	}

	return err
}

func (v *PVlanMappingListValue) SetValue(jsonData *resource_interface_common.NDFCPVlanMappingListValue) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.SVlan != "" {
		v.SVlan = types.StringValue(jsonData.SVlan)
	} else {
		v.SVlan = types.StringNull()
	}

	if jsonData.PVlan != "" {
		v.PVlan = types.StringValue(jsonData.PVlan)
	} else {
		v.PVlan = types.StringNull()
	}

	return err
}

func (v *PVlanAssocListValue) SetValue(jsonData *resource_interface_common.NDFCPVlanAssocListValue) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.SVlan != "" {
		v.SVlan = types.StringValue(jsonData.SVlan)
	} else {
		v.SVlan = types.StringNull()
	}

	if jsonData.PVlan != "" {
		v.PVlan = types.StringValue(jsonData.PVlan)
	} else {
		v.PVlan = types.StringNull()
	}

	return err
}

func (v InterfaceEthernetModel) GetModelData() *resource_interface_common.NDFCInterfaceCommonModel {
	var data = new(resource_interface_common.NDFCInterfaceCommonModel)

	//MARSHAL_BODY

	if !v.Policy.IsNull() && !v.Policy.IsUnknown() {
		data.Policy = v.Policy.ValueString()
	} else {
		data.Policy = ""
	}

	if !v.PolicyType.IsNull() && !v.PolicyType.IsUnknown() {
		data.PolicyType = v.PolicyType.ValueString()
	} else {
		data.PolicyType = ""
	}

	if !v.Deploy.IsNull() && !v.Deploy.IsUnknown() {
		data.Deploy = v.Deploy.ValueBool()
	}

	if !v.SerialNumber.IsNull() && !v.SerialNumber.IsUnknown() {
		data.SerialNumber = v.SerialNumber.ValueString()
	} else {
		data.SerialNumber = ""
	}

	if !v.Interfaces.IsNull() && !v.Interfaces.IsUnknown() {
		elements1 := make(map[string]InterfacesValue, len(v.Interfaces.Elements()))

		data.Interfaces = make(map[string]resource_interface_common.NDFCInterfacesValue)
		diag := v.Interfaces.ElementsAs(context.Background(), &elements1, false)
		if diag != nil {
			panic(diag)
		}
		for k1, ele1 := range elements1 {
			data1 := new(resource_interface_common.NDFCInterfacesValue)
			// filter_this_value | Bool| []| true
			// port_channel_policy | String| []| false
			if !ele1.PortChannelPolicy.IsNull() && !ele1.PortChannelPolicy.IsUnknown() {

				data1.PortChannelPolicy = ele1.PortChannelPolicy.ValueString()
			} else {
				data1.PortChannelPolicy = ""
			}

			// port_channel_name | String| [nvPairs]| false
			if !ele1.PortChannelName.IsNull() && !ele1.PortChannelName.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.PortChannelName = ele1.PortChannelName.ValueString()
			} else {
				data1.NvPairs.PortChannelName = ""
			}

			// serial_number | String| []| false
			if !ele1.SerialNumber.IsNull() && !ele1.SerialNumber.IsUnknown() {

				data1.SerialNumber = ele1.SerialNumber.ValueString()
			} else {
				data1.SerialNumber = ""
			}

			// interface_name | String| []| false
			if !ele1.InterfaceName.IsNull() && !ele1.InterfaceName.IsUnknown() {

				data1.InterfaceName = ele1.InterfaceName.ValueString()
			} else {
				data1.InterfaceName = ""
			}

			// interface_type | String| []| true
			// interface_name | String| [nvPairs]| true
			// freeform_config | String| [nvPairs]| false
			if !ele1.FreeformConfig.IsNull() && !ele1.FreeformConfig.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.FreeformConfig = ele1.FreeformConfig.ValueString()
			} else {
				data1.NvPairs.FreeformConfig = ""
			}

			// admin_state | Bool| [nvPairs]| false
			if !ele1.AdminState.IsNull() && !ele1.AdminState.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.AdminState = strconv.FormatBool(ele1.AdminState.ValueBool())
			} else {
				data1.NvPairs.AdminState = ""
			}

			// interface_description | String| [nvPairs]| false
			if !ele1.InterfaceDescription.IsNull() && !ele1.InterfaceDescription.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.InterfaceDescription = ele1.InterfaceDescription.ValueString()
			} else {
				data1.NvPairs.InterfaceDescription = ""
			}

			// bpdu_guard | String| [nvPairs]| false
			if !ele1.BpduGuard.IsNull() && !ele1.BpduGuard.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.BpduGuard = ele1.BpduGuard.ValueString()
			} else {
				data1.NvPairs.BpduGuard = ""
			}

			// port_type_fast | Bool| [nvPairs]| false
			if !ele1.PortTypeFast.IsNull() && !ele1.PortTypeFast.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.PortTypeFast = strconv.FormatBool(ele1.PortTypeFast.ValueBool())
			} else {
				data1.NvPairs.PortTypeFast = ""
			}

			// mtu | String| [nvPairs]| false
			if !ele1.Mtu.IsNull() && !ele1.Mtu.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.Mtu = ele1.Mtu.ValueString()
			} else {
				data1.NvPairs.Mtu = ""
			}

			// speed | String| [nvPairs]| false
			if !ele1.Speed.IsNull() && !ele1.Speed.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.Speed = ele1.Speed.ValueString()
			} else {
				data1.NvPairs.Speed = ""
			}

			// access_vlan | Int64| [nvPairs]| false
			if !ele1.AccessVlan.IsNull() && !ele1.AccessVlan.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.AccessVlan = new(Int64Custom)
				*data1.NvPairs.AccessVlan = Int64Custom(ele1.AccessVlan.ValueInt64())
			} else {
				data1.NvPairs.AccessVlan = nil
			}

			// orphan_port | Bool| [nvPairs]| false
			if !ele1.OrphanPort.IsNull() && !ele1.OrphanPort.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.OrphanPort = strconv.FormatBool(ele1.OrphanPort.ValueBool())
			} else {
				data1.NvPairs.OrphanPort = ""
			}

			// ptp | Bool| [nvPairs]| false
			if !ele1.Ptp.IsNull() && !ele1.Ptp.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.Ptp = strconv.FormatBool(ele1.Ptp.ValueBool())
			} else {
				data1.NvPairs.Ptp = ""
			}

			// netflow | Bool| [nvPairs]| false
			if !ele1.Netflow.IsNull() && !ele1.Netflow.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.Netflow = strconv.FormatBool(ele1.Netflow.ValueBool())
			} else {
				data1.NvPairs.Netflow = ""
			}

			// netflow_monitor | String| [nvPairs]| false
			if !ele1.NetflowMonitor.IsNull() && !ele1.NetflowMonitor.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.NetflowMonitor = ele1.NetflowMonitor.ValueString()
			} else {
				data1.NvPairs.NetflowMonitor = ""
			}

			// netflow_sampler | String| [nvPairs]| false
			if !ele1.NetflowSampler.IsNull() && !ele1.NetflowSampler.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.NetflowSampler = ele1.NetflowSampler.ValueString()
			} else {
				data1.NvPairs.NetflowSampler = ""
			}

			// allowed_vlans | String| [nvPairs]| false
			if !ele1.AllowedVlans.IsNull() && !ele1.AllowedVlans.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.AllowedVlans = ele1.AllowedVlans.ValueString()
			} else {
				data1.NvPairs.AllowedVlans = ""
			}

			// native_vlan | Int64| [nvPairs]| false
			if !ele1.NativeVlan.IsNull() && !ele1.NativeVlan.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.NativeVlan = new(Int64Custom)
				*data1.NvPairs.NativeVlan = Int64Custom(ele1.NativeVlan.ValueInt64())
			} else {
				data1.NvPairs.NativeVlan = nil
			}

			// vrf | String| [nvPairs]| false
			if !ele1.Vrf.IsNull() && !ele1.Vrf.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.Vrf = ele1.Vrf.ValueString()
			} else {
				data1.NvPairs.Vrf = ""
			}

			// ipv4_address | String| [nvPairs]| false
			if !ele1.Ipv4Address.IsNull() && !ele1.Ipv4Address.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.Ipv4Address = ele1.Ipv4Address.ValueString()
			} else {
				data1.NvPairs.Ipv4Address = ""
			}

			// ipv4_prefix_length | String| [nvPairs]| false
			if !ele1.Ipv4PrefixLength.IsNull() && !ele1.Ipv4PrefixLength.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.Ipv4PrefixLength = ele1.Ipv4PrefixLength.ValueString()
			} else {
				data1.NvPairs.Ipv4PrefixLength = ""
			}

			// routing_tag | String| [nvPairs]| false
			if !ele1.RoutingTag.IsNull() && !ele1.RoutingTag.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.RoutingTag = ele1.RoutingTag.ValueString()
			} else {
				data1.NvPairs.RoutingTag = ""
			}

			// disable_ip_redirects | Bool| [nvPairs]| false
			if !ele1.DisableIpRedirects.IsNull() && !ele1.DisableIpRedirects.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.DisableIpRedirects = strconv.FormatBool(ele1.DisableIpRedirects.ValueBool())
			} else {
				data1.NvPairs.DisableIpRedirects = ""
			}

			// enable_pim_sparse | Bool| [nvPairs]| false
			if !ele1.EnablePimSparse.IsNull() && !ele1.EnablePimSparse.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.EnablePimSparse = strconv.FormatBool(ele1.EnablePimSparse.ValueBool())
			} else {
				data1.NvPairs.EnablePimSparse = ""
			}

			// pim_dr_priority | String| [nvPairs]| false
			if !ele1.PimDrPriority.IsNull() && !ele1.PimDrPriority.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.PimDrPriority = ele1.PimDrPriority.ValueString()
			} else {
				data1.NvPairs.PimDrPriority = ""
			}

			// enable_pfc | Bool| [nvPairs]| false
			if !ele1.EnablePfc.IsNull() && !ele1.EnablePfc.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.EnablePfc = strconv.FormatBool(ele1.EnablePfc.ValueBool())
			} else {
				data1.NvPairs.EnablePfc = ""
			}

			// enable_qos | Bool| [nvPairs]| false
			if !ele1.EnableQos.IsNull() && !ele1.EnableQos.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.EnableQos = strconv.FormatBool(ele1.EnableQos.ValueBool())
			} else {
				data1.NvPairs.EnableQos = ""
			}

			// qos_policy | String| [nvPairs]| false
			if !ele1.QosPolicy.IsNull() && !ele1.QosPolicy.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.QosPolicy = ele1.QosPolicy.ValueString()
			} else {
				data1.NvPairs.QosPolicy = ""
			}

			// queuing_policy | String| [nvPairs]| false
			if !ele1.QueuingPolicy.IsNull() && !ele1.QueuingPolicy.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.QueuingPolicy = ele1.QueuingPolicy.ValueString()
			} else {
				data1.NvPairs.QueuingPolicy = ""
			}

			// link_state_routing_protocol | String| [nvPairs]| false
			if !ele1.LinkStateRoutingProtocol.IsNull() && !ele1.LinkStateRoutingProtocol.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.LinkStateRoutingProtocol = ele1.LinkStateRoutingProtocol.ValueString()
			} else {
				data1.NvPairs.LinkStateRoutingProtocol = ""
			}

			// link_state_routing_tag | String| [nvPairs]| false
			if !ele1.LinkStateRoutingTag.IsNull() && !ele1.LinkStateRoutingTag.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.LinkStateRoutingTag = ele1.LinkStateRoutingTag.ValueString()
			} else {
				data1.NvPairs.LinkStateRoutingTag = ""
			}

			// ipv6_addr | String| [nvPairs]| false
			if !ele1.Ipv6Addr.IsNull() && !ele1.Ipv6Addr.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.Ipv6Addr = ele1.Ipv6Addr.ValueString()
			} else {
				data1.NvPairs.Ipv6Addr = ""
			}

			// ipv6_prefix_length | String| [nvPairs]| false
			if !ele1.Ipv6PrefixLength.IsNull() && !ele1.Ipv6PrefixLength.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.Ipv6PrefixLength = ele1.Ipv6PrefixLength.ValueString()
			} else {
				data1.NvPairs.Ipv6PrefixLength = ""
			}

			// cdp_enable | Bool| [nvPairs]| false
			if !ele1.CdpEnable.IsNull() && !ele1.CdpEnable.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.CdpEnable = strconv.FormatBool(ele1.CdpEnable.ValueBool())
			} else {
				data1.NvPairs.CdpEnable = ""
			}

			// port_duplex_mode | String| [nvPairs]| false
			if !ele1.PortDuplexMode.IsNull() && !ele1.PortDuplexMode.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.PortDuplexMode = ele1.PortDuplexMode.ValueString()
			} else {
				data1.NvPairs.PortDuplexMode = ""
			}

			// enable_monitor | Bool| [nvPairs]| false
			if !ele1.EnableMonitor.IsNull() && !ele1.EnableMonitor.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.EnableMonitor = strconv.FormatBool(ele1.EnableMonitor.ValueBool())
			} else {
				data1.NvPairs.EnableMonitor = ""
			}

			// pvlan_mode | String| [nvPairs]| false
			if !ele1.PvlanMode.IsNull() && !ele1.PvlanMode.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.PvlanMode = ele1.PvlanMode.ValueString()
			} else {
				data1.NvPairs.PvlanMode = ""
			}

			// pvlan_allowed_vlans | String| [nvPairs]| false
			if !ele1.PvlanAllowedVlans.IsNull() && !ele1.PvlanAllowedVlans.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.PvlanAllowedVlans = ele1.PvlanAllowedVlans.ValueString()
			} else {
				data1.NvPairs.PvlanAllowedVlans = ""
			}

			// pvlan_native_vlan | String| [nvPairs]| false
			if !ele1.PvlanNativeVlan.IsNull() && !ele1.PvlanNativeVlan.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.PvlanNativeVlan = ele1.PvlanNativeVlan.ValueString()
			} else {
				data1.NvPairs.PvlanNativeVlan = ""
			}

			// ig_for_fex | Bool| [nvPairs]| false
			if !ele1.IgForFex.IsNull() && !ele1.IgForFex.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.IgForFex = strconv.FormatBool(ele1.IgForFex.ValueBool())
			} else {
				data1.NvPairs.IgForFex = ""
			}

			// auto_negotiate | String| [nvPairs]| false
			if !ele1.AutoNegotiate.IsNull() && !ele1.AutoNegotiate.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.AutoNegotiate = ele1.AutoNegotiate.ValueString()
			} else {
				data1.NvPairs.AutoNegotiate = ""
			}

			// path_cost | Int64| [nvPairs]| false
			if !ele1.PathCost.IsNull() && !ele1.PathCost.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.PathCost = new(int64)
				*data1.NvPairs.PathCost = ele1.PathCost.ValueInt64()

			} else {
				data1.NvPairs.PathCost = nil
			}

			// guard_mode | String| [nvPairs]| false
			if !ele1.GuardMode.IsNull() && !ele1.GuardMode.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.GuardMode = ele1.GuardMode.ValueString()
			} else {
				data1.NvPairs.GuardMode = ""
			}

			// ttag | Bool| [nvPairs]| false
			if !ele1.Ttag.IsNull() && !ele1.Ttag.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.Ttag = strconv.FormatBool(ele1.Ttag.ValueBool())
			} else {
				data1.NvPairs.Ttag = ""
			}

			// p_vlan_mapping_list | List| [nvPairs]| false//-----inline nested----

			if !ele1.PVlanMappingList.IsNull() && !ele1.PVlanMappingList.IsUnknown() {

				data1.NvPairs.PVlanMappingList = make([]resource_interface_common.NDFCPVlanMappingListValue, len(ele1.PVlanMappingList.Elements()))
				elements2 := make([]PVlanMappingListValue, len(ele1.PVlanMappingList.Elements()))
				diag := ele1.PVlanMappingList.ElementsAs(context.Background(), &elements2, false)
				if diag.HasError() {
					panic(diag.Errors())
				}
				for i2, ele2 := range elements2 {
					data2 := new(resource_interface_common.NDFCPVlanMappingListValue)
					if !ele2.SVlan.IsNull() && !ele2.SVlan.IsUnknown() {
						data2.SVlan = ele2.SVlan.ValueString()
					} else {
						data2.SVlan = ""
					}
					data1.NvPairs.PVlanMappingList[i2] = *data2
					if !ele2.PVlan.IsNull() && !ele2.PVlan.IsUnknown() {
						data2.PVlan = ele2.PVlan.ValueString()
					} else {
						data2.PVlan = ""
					}
					data1.NvPairs.PVlanMappingList[i2] = *data2
				}
			}

			// p_vlan_assoc_list | List| [nvPairs]| false//-----inline nested----

			if !ele1.PVlanAssocList.IsNull() && !ele1.PVlanAssocList.IsUnknown() {

				data1.NvPairs.PVlanAssocList = make([]resource_interface_common.NDFCPVlanAssocListValue, len(ele1.PVlanAssocList.Elements()))
				elements2 := make([]PVlanAssocListValue, len(ele1.PVlanAssocList.Elements()))
				diag := ele1.PVlanAssocList.ElementsAs(context.Background(), &elements2, false)
				if diag.HasError() {
					panic(diag.Errors())
				}
				for i2, ele2 := range elements2 {
					data2 := new(resource_interface_common.NDFCPVlanAssocListValue)
					if !ele2.SVlan.IsNull() && !ele2.SVlan.IsUnknown() {
						data2.SVlan = ele2.SVlan.ValueString()
					} else {
						data2.SVlan = ""
					}
					data1.NvPairs.PVlanAssocList[i2] = *data2
					if !ele2.PVlan.IsNull() && !ele2.PVlan.IsUnknown() {
						data2.PVlan = ele2.PVlan.ValueString()
					} else {
						data2.PVlan = ""
					}
					data1.NvPairs.PVlanAssocList[i2] = *data2
				}
			}

			// deployment_status | String| []| false
			if !ele1.DeploymentStatus.IsNull() && !ele1.DeploymentStatus.IsUnknown() {

				data1.DeploymentStatus = ele1.DeploymentStatus.ValueString()
			} else {
				data1.DeploymentStatus = ""
			}

			// custom_policy_parameters | Map:String| []| false
			if !ele1.CustomPolicyParameters.IsNull() && !ele1.CustomPolicyParameters.IsUnknown() {

				mapStringData := make(map[string]string, len(ele1.CustomPolicyParameters.Elements()))
				dg := ele1.CustomPolicyParameters.ElementsAs(context.Background(), &mapStringData, false)
				if dg.HasError() {
					panic(dg.Errors())
				}
				data1.CustomPolicyParameters = make(map[string]string, len(mapStringData))
				for k, v := range mapStringData {
					data1.CustomPolicyParameters[k] = v
				}
			}

			data.Interfaces[k1] = *data1

		}
	}

	return data
}

func (v *InterfaceEthernetModel) SetDefaultValues() {

	if v.Policy.IsNull() || v.Policy.IsUnknown() {
		v.Policy = types.StringValue("int_trunk_host")
	}

	if v.PolicyType.IsNull() || v.PolicyType.IsUnknown() {
		v.PolicyType = types.StringValue("system")
	}

	if v.Deploy.IsNull() || v.Deploy.IsUnknown() {
		v.Deploy = types.BoolValue(false)
	}

}

func (v *InterfacesValue) SetDefaultValues() {

	if v.AdminState.IsNull() || v.AdminState.IsUnknown() {
		v.AdminState = types.BoolValue(true)
	}

	if v.BpduGuard.IsNull() || v.BpduGuard.IsUnknown() {
		v.BpduGuard = types.StringValue("true")
	}

	if v.PortTypeFast.IsNull() || v.PortTypeFast.IsUnknown() {
		v.PortTypeFast = types.BoolValue(true)
	}

	if v.Mtu.IsNull() || v.Mtu.IsUnknown() {
		v.Mtu = types.StringValue("jumbo")
	}

	if v.Speed.IsNull() || v.Speed.IsUnknown() {
		v.Speed = types.StringValue("Auto")
	}

	if v.OrphanPort.IsNull() || v.OrphanPort.IsUnknown() {
		v.OrphanPort = types.BoolValue(false)
	}

	if v.Ptp.IsNull() || v.Ptp.IsUnknown() {
		v.Ptp = types.BoolValue(false)
	}

	if v.Netflow.IsNull() || v.Netflow.IsUnknown() {
		v.Netflow = types.BoolValue(false)
	}

	if v.AllowedVlans.IsNull() || v.AllowedVlans.IsUnknown() {
		v.AllowedVlans = types.StringValue("none")
	}

	if v.DisableIpRedirects.IsNull() || v.DisableIpRedirects.IsUnknown() {
		v.DisableIpRedirects = types.BoolValue(false)
	}

	if v.EnablePimSparse.IsNull() || v.EnablePimSparse.IsUnknown() {
		v.EnablePimSparse = types.BoolValue(false)
	}

	if v.PimDrPriority.IsNull() || v.PimDrPriority.IsUnknown() {
		v.PimDrPriority = types.StringValue("1")
	}

	if v.EnablePfc.IsNull() || v.EnablePfc.IsUnknown() {
		v.EnablePfc = types.BoolValue(false)
	}

	if v.EnableQos.IsNull() || v.EnableQos.IsUnknown() {
		v.EnableQos = types.BoolValue(false)
	}

	if v.CdpEnable.IsNull() || v.CdpEnable.IsUnknown() {
		v.CdpEnable = types.BoolValue(false)
	}

	if v.PortDuplexMode.IsNull() || v.PortDuplexMode.IsUnknown() {
		v.PortDuplexMode = types.StringValue("auto")
	}

	if v.EnableMonitor.IsNull() || v.EnableMonitor.IsUnknown() {
		v.EnableMonitor = types.BoolValue(false)
	}

	if v.PvlanMode.IsNull() || v.PvlanMode.IsUnknown() {
		v.PvlanMode = types.StringValue("host")
	}

	if v.AutoNegotiate.IsNull() || v.AutoNegotiate.IsUnknown() {
		v.AutoNegotiate = types.StringValue("on")
	}

}

func (v *PVlanMappingListValue) SetDefaultValues() {

}

func (v *PVlanAssocListValue) SetDefaultValues() {

}
