// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package resource_interface_vpc

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

func (v *InterfaceVpcModel) SetModelData(jsonData *resource_interface_common.NDFCInterfaceCommonModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.Policy != "" {
		v.Policy = types.StringValue(jsonData.Policy)
	} else {
		v.Policy = types.StringNull()
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

	if jsonData.NvPairs.AdminState != "" {
		x, _ := strconv.ParseBool(jsonData.NvPairs.AdminState)
		v.AdminState = types.BoolValue(x)
	} else {
		v.AdminState = types.BoolNull()
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

	if jsonData.NvPairs.CopyPoDescription != "" {
		x, _ := strconv.ParseBool(jsonData.NvPairs.CopyPoDescription)
		v.CopyPoDescription = types.BoolValue(x)
	} else {
		v.CopyPoDescription = types.BoolNull()
	}

	if jsonData.NvPairs.PortchannelMode != "" {
		v.PortchannelMode = types.StringValue(jsonData.NvPairs.PortchannelMode)
	} else {
		v.PortchannelMode = types.StringNull()
	}

	if jsonData.NvPairs.Peer1PoFreeformConfig != "" {
		v.Peer1PoFreeformConfig = types.StringValue(jsonData.NvPairs.Peer1PoFreeformConfig)
	} else {
		v.Peer1PoFreeformConfig = types.StringNull()
	}

	if jsonData.NvPairs.Peer2PoFreeformConfig != "" {
		v.Peer2PoFreeformConfig = types.StringValue(jsonData.NvPairs.Peer2PoFreeformConfig)
	} else {
		v.Peer2PoFreeformConfig = types.StringNull()
	}

	if jsonData.NvPairs.Peer1PoDescription != "" {
		v.Peer1PoDescription = types.StringValue(jsonData.NvPairs.Peer1PoDescription)
	} else {
		v.Peer1PoDescription = types.StringNull()
	}

	if jsonData.NvPairs.Peer2PoDescription != "" {
		v.Peer2PoDescription = types.StringValue(jsonData.NvPairs.Peer2PoDescription)
	} else {
		v.Peer2PoDescription = types.StringNull()
	}

	if jsonData.NvPairs.Peer1AllowedVlans != "" {
		v.Peer1AllowedVlans = types.StringValue(jsonData.NvPairs.Peer1AllowedVlans)
	} else {
		v.Peer1AllowedVlans = types.StringNull()
	}

	if jsonData.NvPairs.Peer2AllowedVlans != "" {
		v.Peer2AllowedVlans = types.StringValue(jsonData.NvPairs.Peer2AllowedVlans)
	} else {
		v.Peer2AllowedVlans = types.StringNull()
	}

	if jsonData.NvPairs.Peer1NativeVlan != nil {
		if jsonData.NvPairs.Peer1NativeVlan.IsEmpty() {
			v.Peer1NativeVlan = types.Int64Null()
		} else {
			v.Peer1NativeVlan = types.Int64Value(int64(*jsonData.NvPairs.Peer1NativeVlan))
		}

	} else {
		v.Peer1NativeVlan = types.Int64Null()
	}

	if jsonData.NvPairs.Peer2NativeVlan != nil {
		if jsonData.NvPairs.Peer2NativeVlan.IsEmpty() {
			v.Peer2NativeVlan = types.Int64Null()
		} else {
			v.Peer2NativeVlan = types.Int64Value(int64(*jsonData.NvPairs.Peer2NativeVlan))
		}

	} else {
		v.Peer2NativeVlan = types.Int64Null()
	}

	if jsonData.NvPairs.Peer1MemberInterfaces != "" {
		v.Peer1MemberInterfaces = types.StringValue(jsonData.NvPairs.Peer1MemberInterfaces)
	} else {
		v.Peer1MemberInterfaces = types.StringNull()
	}

	if jsonData.NvPairs.Peer2MemberInterfaces != "" {
		v.Peer2MemberInterfaces = types.StringValue(jsonData.NvPairs.Peer2MemberInterfaces)
	} else {
		v.Peer2MemberInterfaces = types.StringNull()
	}

	if jsonData.NvPairs.Peer1PortChannelId != nil {
		v.Peer1PortChannelId = types.Int64Value(*jsonData.NvPairs.Peer1PortChannelId)

	} else {
		v.Peer1PortChannelId = types.Int64Null()
	}

	if jsonData.NvPairs.Peer2PortChannelId != nil {
		if jsonData.NvPairs.Peer2PortChannelId.IsEmpty() {
			v.Peer2PortChannelId = types.Int64Null()
		} else {
			v.Peer2PortChannelId = types.Int64Value(int64(*jsonData.NvPairs.Peer2PortChannelId))
		}

	} else {
		v.Peer2PortChannelId = types.Int64Null()
	}

	if jsonData.DeploymentStatus != "" {
		v.DeploymentStatus = types.StringValue(jsonData.DeploymentStatus)
	} else {
		v.DeploymentStatus = types.StringNull()
	}

	return err
}

func (v InterfaceVpcModel) GetModelData() *resource_interface_common.NDFCInterfaceCommonModel {
	var data = new(resource_interface_common.NDFCInterfaceCommonModel)

	//MARSHAL_BODY

	if !v.Policy.IsNull() && !v.Policy.IsUnknown() {
		data.Policy = v.Policy.ValueString()
	} else {
		data.Policy = ""
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
			// admin_state | Bool| [nvPairs]| false
			if !ele1.AdminState.IsNull() && !ele1.AdminState.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.AdminState = strconv.FormatBool(ele1.AdminState.ValueBool())
			} else {
				data1.NvPairs.AdminState = ""
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

			// copy_po_description | Bool| [nvPairs]| false
			if !ele1.CopyPoDescription.IsNull() && !ele1.CopyPoDescription.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.CopyPoDescription = strconv.FormatBool(ele1.CopyPoDescription.ValueBool())
			} else {
				data1.NvPairs.CopyPoDescription = ""
			}

			// portchannel_mode | String| [nvPairs]| false
			if !ele1.PortchannelMode.IsNull() && !ele1.PortchannelMode.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.PortchannelMode = ele1.PortchannelMode.ValueString()
			} else {
				data1.NvPairs.PortchannelMode = ""
			}

			// mirror_config | Bool| [nvPairs]| true
			// peer1_po_freeform_config | String| [nvPairs]| false
			if !ele1.Peer1PoFreeformConfig.IsNull() && !ele1.Peer1PoFreeformConfig.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.Peer1PoFreeformConfig = ele1.Peer1PoFreeformConfig.ValueString()
			} else {
				data1.NvPairs.Peer1PoFreeformConfig = ""
			}

			// peer2_po_freeform_config | String| [nvPairs]| false
			if !ele1.Peer2PoFreeformConfig.IsNull() && !ele1.Peer2PoFreeformConfig.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.Peer2PoFreeformConfig = ele1.Peer2PoFreeformConfig.ValueString()
			} else {
				data1.NvPairs.Peer2PoFreeformConfig = ""
			}

			// peer1_po_description | String| [nvPairs]| false
			if !ele1.Peer1PoDescription.IsNull() && !ele1.Peer1PoDescription.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.Peer1PoDescription = ele1.Peer1PoDescription.ValueString()
			} else {
				data1.NvPairs.Peer1PoDescription = ""
			}

			// peer2_po_description | String| [nvPairs]| false
			if !ele1.Peer2PoDescription.IsNull() && !ele1.Peer2PoDescription.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.Peer2PoDescription = ele1.Peer2PoDescription.ValueString()
			} else {
				data1.NvPairs.Peer2PoDescription = ""
			}

			// peer1_allowed_vlans | String| [nvPairs]| false
			if !ele1.Peer1AllowedVlans.IsNull() && !ele1.Peer1AllowedVlans.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.Peer1AllowedVlans = ele1.Peer1AllowedVlans.ValueString()
			} else {
				data1.NvPairs.Peer1AllowedVlans = ""
			}

			// peer2_allowed_vlans | String| [nvPairs]| false
			if !ele1.Peer2AllowedVlans.IsNull() && !ele1.Peer2AllowedVlans.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.Peer2AllowedVlans = ele1.Peer2AllowedVlans.ValueString()
			} else {
				data1.NvPairs.Peer2AllowedVlans = ""
			}

			// peer1_native_vlan | Int64| [nvPairs]| false
			if !ele1.Peer1NativeVlan.IsNull() && !ele1.Peer1NativeVlan.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.Peer1NativeVlan = new(Int64Custom)
				*data1.NvPairs.Peer1NativeVlan = Int64Custom(ele1.Peer1NativeVlan.ValueInt64())
			} else {
				data1.NvPairs.Peer1NativeVlan = nil
			}

			// peer2_native_vlan | Int64| [nvPairs]| false
			if !ele1.Peer2NativeVlan.IsNull() && !ele1.Peer2NativeVlan.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.Peer2NativeVlan = new(Int64Custom)
				*data1.NvPairs.Peer2NativeVlan = Int64Custom(ele1.Peer2NativeVlan.ValueInt64())
			} else {
				data1.NvPairs.Peer2NativeVlan = nil
			}

			// peer1_member_interfaces | String| [nvPairs]| false
			if !ele1.Peer1MemberInterfaces.IsNull() && !ele1.Peer1MemberInterfaces.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.Peer1MemberInterfaces = ele1.Peer1MemberInterfaces.ValueString()
			} else {
				data1.NvPairs.Peer1MemberInterfaces = ""
			}

			// peer2_member_interfaces | String| [nvPairs]| false
			if !ele1.Peer2MemberInterfaces.IsNull() && !ele1.Peer2MemberInterfaces.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.Peer2MemberInterfaces = ele1.Peer2MemberInterfaces.ValueString()
			} else {
				data1.NvPairs.Peer2MemberInterfaces = ""
			}

			// peer1_port_channel_id | Int64| [nvPairs]| false
			if !ele1.Peer1PortChannelId.IsNull() && !ele1.Peer1PortChannelId.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.Peer1PortChannelId = new(int64)
				*data1.NvPairs.Peer1PortChannelId = ele1.Peer1PortChannelId.ValueInt64()

			} else {
				data1.NvPairs.Peer1PortChannelId = nil
			}

			// peer2_port_channel_id | Int64| [nvPairs]| false
			if !ele1.Peer2PortChannelId.IsNull() && !ele1.Peer2PortChannelId.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.Peer2PortChannelId = new(Int64Custom)
				*data1.NvPairs.Peer2PortChannelId = Int64Custom(ele1.Peer2PortChannelId.ValueInt64())
			} else {
				data1.NvPairs.Peer2PortChannelId = nil
			}

			// deployment_status | String| []| false
			data.Interfaces[k1] = *data1

		}
	}

	return data
}
