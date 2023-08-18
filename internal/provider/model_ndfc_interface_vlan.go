// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

//template:begin imports
import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

//template:end imports

//template:begin types
type InterfaceVlan struct {
	Id                        types.String `tfsdk:"id"`
	SerialNumber              types.String `tfsdk:"serial_number"`
	InterfaceName             types.String `tfsdk:"interface_name"`
	Policy                    types.String `tfsdk:"policy"`
	Vrf                       types.String `tfsdk:"vrf"`
	Ipv4Address               types.String `tfsdk:"ipv4_address"`
	Ipv4PrefixLength          types.Int64  `tfsdk:"ipv4_prefix_length"`
	Mtu                       types.Int64  `tfsdk:"mtu"`
	RoutingTag                types.String `tfsdk:"routing_tag"`
	DisableIpRedirects        types.Bool   `tfsdk:"disable_ip_redirects"`
	InterfaceDescription      types.String `tfsdk:"interface_description"`
	FreeformConfig            types.String `tfsdk:"freeform_config"`
	AdminState                types.Bool   `tfsdk:"admin_state"`
	Hsrp                      types.Bool   `tfsdk:"hsrp"`
	HsrpVip                   types.String `tfsdk:"hsrp_vip"`
	HsrpGroup                 types.Int64  `tfsdk:"hsrp_group"`
	HsrpVersion               types.String `tfsdk:"hsrp_version"`
	HsrpPriority              types.Int64  `tfsdk:"hsrp_priority"`
	HsrpPreempt               types.Bool   `tfsdk:"hsrp_preempt"`
	HsrpMac                   types.String `tfsdk:"hsrp_mac"`
	DhcpServer1               types.String `tfsdk:"dhcp_server_1"`
	DhcpServer1Vrf            types.String `tfsdk:"dhcp_server_1_vrf"`
	DhcpServer2               types.String `tfsdk:"dhcp_server_2"`
	DhcpServer2Vrf            types.String `tfsdk:"dhcp_server_2_vrf"`
	DhcpServer3               types.String `tfsdk:"dhcp_server_3"`
	DhcpServer3Vrf            types.String `tfsdk:"dhcp_server_3_vrf"`
	AdvertiseSubnetInUnderlay types.Bool   `tfsdk:"advertise_subnet_in_underlay"`
	Netflow                   types.Bool   `tfsdk:"netflow"`
	NetflowMonitor            types.String `tfsdk:"netflow_monitor"`
	NetflowSampler            types.String `tfsdk:"netflow_sampler"`
}

//template:end types

//template:begin getPath
func (data InterfaceVlan) getPath() string {
	return "/lan-fabric/rest/interface"
}

//template:end getPath

func (data InterfaceVlan) toBody(ctx context.Context) string {
	body := ""
	if !data.SerialNumber.IsNull() && !data.SerialNumber.IsUnknown() {
		body, _ = sjson.Set(body, "interfaces.0.serialNumber", data.SerialNumber.ValueString())
	}
	if !data.InterfaceName.IsNull() && !data.InterfaceName.IsUnknown() {
		body, _ = sjson.Set(body, "interfaces.0.ifName", data.InterfaceName.ValueString())
	}
	if !data.Policy.IsNull() && !data.Policy.IsUnknown() {
		body, _ = sjson.Set(body, "policy", data.Policy.ValueString())
	}
	body, _ = sjson.Set(body, "interfaceType", "INTERFACE_VLAN")
	body, _ = sjson.Set(body, "interfaces.0.nvPairs.INTF_NAME", data.InterfaceName.ValueString())
	body, _ = sjson.Set(body, "interfaces.0.nvPairs.INTF_VRF", data.Vrf.ValueString())
	body, _ = sjson.Set(body, "interfaces.0.nvPairs.IP", data.Ipv4Address.ValueString())
	if !data.Ipv4PrefixLength.IsNull() && !data.Ipv4PrefixLength.IsUnknown() {
		body, _ = sjson.Set(body, "interfaces.0.nvPairs.PREFIX", fmt.Sprint(data.Ipv4PrefixLength.ValueInt64()))
	}
	if !data.Mtu.IsNull() && !data.Mtu.IsUnknown() && data.Mtu.ValueInt64() != 0 {
		body, _ = sjson.Set(body, "interfaces.0.nvPairs.MTU", fmt.Sprint(data.Mtu.ValueInt64()))
	} else {
		body, _ = sjson.Set(body, "interfaces.0.nvPairs.MTU", "")
	}
	if !data.RoutingTag.IsNull() && !data.RoutingTag.IsUnknown() {
		body, _ = sjson.Set(body, "interfaces.0.nvPairs.ROUTING_TAG", data.RoutingTag.ValueString())
	}
	if !data.DisableIpRedirects.IsNull() && !data.DisableIpRedirects.IsUnknown() {
		body, _ = sjson.Set(body, "interfaces.0.nvPairs.DISABLE_IP_REDIRECTS", fmt.Sprint(data.DisableIpRedirects.ValueBool()))
	}
	body, _ = sjson.Set(body, "interfaces.0.nvPairs.DESC", data.InterfaceDescription.ValueString())
	body, _ = sjson.Set(body, "interfaces.0.nvPairs.CONF", data.FreeformConfig.ValueString())
	if !data.AdminState.IsNull() && !data.AdminState.IsUnknown() {
		body, _ = sjson.Set(body, "interfaces.0.nvPairs.ADMIN_STATE", fmt.Sprint(data.AdminState.ValueBool()))
	}
	if !data.Hsrp.IsNull() && !data.Hsrp.IsUnknown() {
		body, _ = sjson.Set(body, "interfaces.0.nvPairs.ENABLE_HSRP", fmt.Sprint(data.Hsrp.ValueBool()))
	} else {
		body, _ = sjson.Set(body, "interfaces.0.nvPairs.ENABLE_HSRP", "")
	}
	if !data.HsrpVip.IsNull() && !data.HsrpVip.IsUnknown() {
		body, _ = sjson.Set(body, "interfaces.0.nvPairs.HSRP_VIP", data.HsrpVip.ValueString())
	}
	if !data.HsrpGroup.IsNull() && !data.HsrpGroup.IsUnknown() {
		body, _ = sjson.Set(body, "interfaces.0.nvPairs.HSRP_GROUP", data.HsrpGroup.ValueInt64())
	}
	if !data.HsrpVersion.IsNull() && !data.HsrpVersion.IsUnknown() {
		body, _ = sjson.Set(body, "interfaces.0.nvPairs.HSRP_VERSION", data.HsrpVersion.ValueString())
	}
	if !data.HsrpPriority.IsNull() && !data.HsrpPriority.IsUnknown() {
		body, _ = sjson.Set(body, "interfaces.0.nvPairs.HSRP_PRIORITY", data.HsrpPriority.ValueInt64())
	}
	if !data.HsrpPreempt.IsNull() && !data.HsrpPreempt.IsUnknown() {
		body, _ = sjson.Set(body, "interfaces.0.nvPairs.PREEMPT", fmt.Sprint(data.HsrpPreempt.ValueBool()))
	}
	if !data.HsrpMac.IsNull() && !data.HsrpMac.IsUnknown() {
		body, _ = sjson.Set(body, "interfaces.0.nvPairs.MAC", data.HsrpMac.ValueString())
	}
	if !data.DhcpServer1.IsNull() && !data.DhcpServer1.IsUnknown() {
		body, _ = sjson.Set(body, "interfaces.0.nvPairs.dhcpServerAddr1", data.DhcpServer1.ValueString())
	}
	if !data.DhcpServer1Vrf.IsNull() && !data.DhcpServer1Vrf.IsUnknown() {
		body, _ = sjson.Set(body, "interfaces.0.nvPairs.vrfDhcp1", data.DhcpServer1Vrf.ValueString())
	}
	if !data.DhcpServer2.IsNull() && !data.DhcpServer2.IsUnknown() {
		body, _ = sjson.Set(body, "interfaces.0.nvPairs.dhcpServerAddr2", data.DhcpServer2.ValueString())
	}
	if !data.DhcpServer2Vrf.IsNull() && !data.DhcpServer2Vrf.IsUnknown() {
		body, _ = sjson.Set(body, "interfaces.0.nvPairs.vrfDhcp2", data.DhcpServer2Vrf.ValueString())
	}
	if !data.DhcpServer3.IsNull() && !data.DhcpServer3.IsUnknown() {
		body, _ = sjson.Set(body, "interfaces.0.nvPairs.dhcpServerAddr3", data.DhcpServer3.ValueString())
	}
	if !data.DhcpServer3Vrf.IsNull() && !data.DhcpServer3Vrf.IsUnknown() {
		body, _ = sjson.Set(body, "interfaces.0.nvPairs.vrfDhcp3", data.DhcpServer3Vrf.ValueString())
	}
	if !data.AdvertiseSubnetInUnderlay.IsNull() && !data.AdvertiseSubnetInUnderlay.IsUnknown() {
		body, _ = sjson.Set(body, "interfaces.0.nvPairs.advSubnetInUnderlay", fmt.Sprint(data.AdvertiseSubnetInUnderlay.ValueBool()))
	}
	if !data.Netflow.IsNull() && !data.Netflow.IsUnknown() {
		body, _ = sjson.Set(body, "interfaces.0.nvPairs.ENABLE_NETFLOW", fmt.Sprint(data.Netflow.ValueBool()))
	}
	if !data.NetflowMonitor.IsNull() && !data.NetflowMonitor.IsUnknown() {
		body, _ = sjson.Set(body, "interfaces.0.nvPairs.NETFLOW_MONITOR", data.NetflowMonitor.ValueString())
	}
	if !data.NetflowSampler.IsNull() && !data.NetflowSampler.IsUnknown() {
		body, _ = sjson.Set(body, "interfaces.0.nvPairs.NETFLOW_SAMPLER", data.NetflowSampler.ValueString())
	}
	return body
}

func (data *InterfaceVlan) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("0.interfaces.0.serialNumber"); value.Exists() && value.String() != "" {
		data.SerialNumber = types.StringValue(value.String())
	} else {
		data.SerialNumber = types.StringNull()
	}
	if value := res.Get("0.interfaces.0.ifName"); value.Exists() && value.String() != "" {
		data.InterfaceName = types.StringValue(value.String())
	} else {
		data.InterfaceName = types.StringNull()
	}
	if value := res.Get("0.policy"); value.Exists() && value.String() != "" {
		data.Policy = types.StringValue(value.String())
	} else {
		data.Policy = types.StringNull()
	}
	if value := res.Get("0.interfaces.0.nvPairs.INTF_VRF"); value.Exists() && value.String() != "" {
		data.Vrf = types.StringValue(value.String())
	} else {
		data.Vrf = types.StringNull()
	}
	if value := res.Get("0.interfaces.0.nvPairs.IP"); value.Exists() && value.String() != "" {
		data.Ipv4Address = types.StringValue(value.String())
	} else {
		data.Ipv4Address = types.StringNull()
	}
	if value := res.Get("0.interfaces.0.nvPairs.PREFIX"); value.Exists() && value.String() != "" {
		data.Ipv4PrefixLength = types.Int64Value(value.Int())
	} else {
		data.Ipv4PrefixLength = types.Int64Null()
	}
	if value := res.Get("0.interfaces.0.nvPairs.MTU"); value.Exists() && value.String() != "" {
		data.Mtu = types.Int64Value(value.Int())
	} else {
		data.Mtu = types.Int64Null()
	}
	if value := res.Get("0.interfaces.0.nvPairs.ROUTING_TAG"); value.Exists() && value.String() != "" {
		data.RoutingTag = types.StringValue(value.String())
	} else {
		data.RoutingTag = types.StringNull()
	}
	if value := res.Get("0.interfaces.0.nvPairs.DISABLE_IP_REDIRECTS"); value.Exists() && value.String() != "" {
		data.DisableIpRedirects = types.BoolValue(value.Bool())
	} else {
		data.DisableIpRedirects = types.BoolNull()
	}
	if value := res.Get("0.interfaces.0.nvPairs.DESC"); value.Exists() && value.String() != "" {
		data.InterfaceDescription = types.StringValue(value.String())
	} else {
		data.InterfaceDescription = types.StringNull()
	}
	if value := res.Get("0.interfaces.0.nvPairs.CONF"); value.Exists() && value.String() != "" {
		data.FreeformConfig = types.StringValue(value.String())
	} else {
		data.FreeformConfig = types.StringNull()
	}
	if value := res.Get("0.interfaces.0.nvPairs.ADMIN_STATE"); value.Exists() && value.String() != "" {
		data.AdminState = types.BoolValue(value.Bool())
	} else {
		data.AdminState = types.BoolNull()
	}
	if value := res.Get("0.interfaces.0.nvPairs.ENABLE_HSRP"); value.Exists() && value.String() != "" {
		data.Hsrp = types.BoolValue(value.Bool())
	} else {
		data.Hsrp = types.BoolNull()
	}
	if value := res.Get("0.interfaces.0.nvPairs.HSRP_VIP"); value.Exists() && value.String() != "" {
		data.HsrpVip = types.StringValue(value.String())
	} else {
		data.HsrpVip = types.StringNull()
	}
	if value := res.Get("0.interfaces.0.nvPairs.HSRP_GROUP"); value.Exists() && value.String() != "" {
		data.HsrpGroup = types.Int64Value(value.Int())
	} else {
		data.HsrpGroup = types.Int64Null()
	}
	if value := res.Get("0.interfaces.0.nvPairs.HSRP_VERSION"); value.Exists() && value.String() != "" {
		data.HsrpVersion = types.StringValue(value.String())
	} else {
		data.HsrpVersion = types.StringNull()
	}
	if value := res.Get("0.interfaces.0.nvPairs.HSRP_PRIORITY"); value.Exists() && value.String() != "" {
		data.HsrpPriority = types.Int64Value(value.Int())
	} else {
		data.HsrpPriority = types.Int64Null()
	}
	if value := res.Get("0.interfaces.0.nvPairs.PREEMPT"); value.Exists() && value.String() != "" {
		data.HsrpPreempt = types.BoolValue(value.Bool())
	} else {
		data.HsrpPreempt = types.BoolNull()
	}
	if value := res.Get("0.interfaces.0.nvPairs.MAC"); value.Exists() && value.String() != "" {
		data.HsrpMac = types.StringValue(value.String())
	} else {
		data.HsrpMac = types.StringNull()
	}
	if value := res.Get("0.interfaces.0.nvPairs.dhcpServerAddr1"); value.Exists() && value.String() != "" {
		data.DhcpServer1 = types.StringValue(value.String())
	} else {
		data.DhcpServer1 = types.StringNull()
	}
	if value := res.Get("0.interfaces.0.nvPairs.vrfDhcp1"); value.Exists() && value.String() != "" {
		data.DhcpServer1Vrf = types.StringValue(value.String())
	} else {
		data.DhcpServer1Vrf = types.StringNull()
	}
	if value := res.Get("0.interfaces.0.nvPairs.dhcpServerAddr2"); value.Exists() && value.String() != "" {
		data.DhcpServer2 = types.StringValue(value.String())
	} else {
		data.DhcpServer2 = types.StringNull()
	}
	if value := res.Get("0.interfaces.0.nvPairs.vrfDhcp2"); value.Exists() && value.String() != "" {
		data.DhcpServer2Vrf = types.StringValue(value.String())
	} else {
		data.DhcpServer2Vrf = types.StringNull()
	}
	if value := res.Get("0.interfaces.0.nvPairs.dhcpServerAddr3"); value.Exists() && value.String() != "" {
		data.DhcpServer3 = types.StringValue(value.String())
	} else {
		data.DhcpServer3 = types.StringNull()
	}
	if value := res.Get("0.interfaces.0.nvPairs.vrfDhcp3"); value.Exists() && value.String() != "" {
		data.DhcpServer3Vrf = types.StringValue(value.String())
	} else {
		data.DhcpServer3Vrf = types.StringNull()
	}
	if value := res.Get("0.interfaces.0.nvPairs.advSubnetInUnderlay"); value.Exists() && value.String() != "" {
		data.AdvertiseSubnetInUnderlay = types.BoolValue(value.Bool())
	} else {
		data.AdvertiseSubnetInUnderlay = types.BoolNull()
	}
	if value := res.Get("0.interfaces.0.nvPairs.ENABLE_NETFLOW"); value.Exists() && value.String() != "" {
		data.Netflow = types.BoolValue(value.Bool())
	} else {
		data.Netflow = types.BoolNull()
	}
	if value := res.Get("0.interfaces.0.nvPairs.NETFLOW_MONITOR"); value.Exists() && value.String() != "" {
		data.NetflowMonitor = types.StringValue(value.String())
	} else {
		data.NetflowMonitor = types.StringNull()
	}
	if value := res.Get("0.interfaces.0.nvPairs.NETFLOW_SAMPLER"); value.Exists() && value.String() != "" {
		data.NetflowSampler = types.StringValue(value.String())
	} else {
		data.NetflowSampler = types.StringNull()
	}
}
