// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package resource_ip_fabric_media

import (
	"strconv"
	. "terraform-provider-ndfc/internal/provider/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"terraform-provider-ndfc/internal/provider/resources/resource_fabric_common"
)

func (v *IpFabricMediaModel) SetModelData(jsonData *resource_fabric_common.NDFCFabricCommonModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.FabricName != "" {
		v.FabricName = types.StringValue(jsonData.FabricName)
	} else {
		v.FabricName = types.StringNull()
	}

	if jsonData.AaaRemoteIpEnabled != "" {
		x, _ := strconv.ParseBool(jsonData.AaaRemoteIpEnabled)
		v.AaaRemoteIpEnabled = types.BoolValue(x)
	} else {
		v.AaaRemoteIpEnabled = types.BoolNull()
	}

	if jsonData.AaaServerConf != "" {
		v.AaaServerConf = types.StringValue(jsonData.AaaServerConf)
	} else {
		v.AaaServerConf = types.StringNull()
	}

	if jsonData.AsmGroupRanges != "" {
		v.AsmGroupRanges = types.StringValue(jsonData.AsmGroupRanges)
	} else {
		v.AsmGroupRanges = types.StringNull()
	}

	if jsonData.BootstrapConf != "" {
		v.BootstrapConf = types.StringValue(jsonData.BootstrapConf)
	} else {
		v.BootstrapConf = types.StringNull()
	}

	if jsonData.BootstrapEnable != "" {
		x, _ := strconv.ParseBool(jsonData.BootstrapEnable)
		v.BootstrapEnable = types.BoolValue(x)
	} else {
		v.BootstrapEnable = types.BoolNull()
	}

	if jsonData.BootstrapMultisubnet != "" {
		v.BootstrapMultisubnet = types.StringValue(jsonData.BootstrapMultisubnet)
	} else {
		v.BootstrapMultisubnet = types.StringNull()
	}

	if jsonData.CdpEnable != "" {
		x, _ := strconv.ParseBool(jsonData.CdpEnable)
		v.CdpEnable = types.BoolValue(x)
	} else {
		v.CdpEnable = types.BoolNull()
	}

	if jsonData.DhcpEnable != "" {
		x, _ := strconv.ParseBool(jsonData.DhcpEnable)
		v.DhcpEnable = types.BoolValue(x)
	} else {
		v.DhcpEnable = types.BoolNull()
	}

	if jsonData.DhcpEnd != "" {
		v.DhcpEnd = types.StringValue(jsonData.DhcpEnd)
	} else {
		v.DhcpEnd = types.StringNull()
	}

	if jsonData.DhcpIpv6Enable != "" {
		v.DhcpIpv6Enable = types.StringValue(jsonData.DhcpIpv6Enable)
	} else {
		v.DhcpIpv6Enable = types.StringNull()
	}

	if jsonData.DhcpStart != "" {
		v.DhcpStart = types.StringValue(jsonData.DhcpStart)
	} else {
		v.DhcpStart = types.StringNull()
	}

	if jsonData.DnsServerIpList != "" {
		v.DnsServerIpList = types.StringValue(jsonData.DnsServerIpList)
	} else {
		v.DnsServerIpList = types.StringNull()
	}

	if jsonData.DnsServerVrf != "" {
		v.DnsServerVrf = types.StringValue(jsonData.DnsServerVrf)
	} else {
		v.DnsServerVrf = types.StringNull()
	}

	if jsonData.EnableAaa != "" {
		x, _ := strconv.ParseBool(jsonData.EnableAaa)
		v.EnableAaa = types.BoolValue(x)
	} else {
		v.EnableAaa = types.BoolNull()
	}

	if jsonData.EnableAsm != "" {
		x, _ := strconv.ParseBool(jsonData.EnableAsm)
		v.EnableAsm = types.BoolValue(x)
	} else {
		v.EnableAsm = types.BoolNull()
	}

	if jsonData.EnableNbmPassive != "" {
		x, _ := strconv.ParseBool(jsonData.EnableNbmPassive)
		v.EnableNbmPassive = types.BoolValue(x)
	} else {
		v.EnableNbmPassive = types.BoolNull()
	}

	if jsonData.ExtraConfIntraLinks != "" {
		v.ExtraConfIntraLinks = types.StringValue(jsonData.ExtraConfIntraLinks)
	} else {
		v.ExtraConfIntraLinks = types.StringNull()
	}

	if jsonData.ExtraConfLeaf != "" {
		v.ExtraConfLeaf = types.StringValue(jsonData.ExtraConfLeaf)
	} else {
		v.ExtraConfLeaf = types.StringNull()
	}

	if jsonData.ExtraConfSpine != "" {
		v.ExtraConfSpine = types.StringValue(jsonData.ExtraConfSpine)
	} else {
		v.ExtraConfSpine = types.StringNull()
	}

	if jsonData.FabricInterfaceType != "" {
		v.FabricInterfaceType = types.StringValue(jsonData.FabricInterfaceType)
	} else {
		v.FabricInterfaceType = types.StringNull()
	}

	if jsonData.FabricMtu != nil {
		if jsonData.FabricMtu.IsEmpty() {
			v.FabricMtu = types.Int64Null()
		} else {
			v.FabricMtu = types.Int64Value(int64(*jsonData.FabricMtu))
		}
	} else {
		v.FabricMtu = types.Int64Null()
	}

	if jsonData.FeaturePtp != "" {
		x, _ := strconv.ParseBool(jsonData.FeaturePtp)
		v.FeaturePtp = types.BoolValue(x)
	} else {
		v.FeaturePtp = types.BoolNull()
	}

	if jsonData.IsisAuthEnable != "" {
		x, _ := strconv.ParseBool(jsonData.IsisAuthEnable)
		v.IsisAuthEnable = types.BoolValue(x)
	} else {
		v.IsisAuthEnable = types.BoolNull()
	}

	if jsonData.IsisAuthKey != "" {
		v.IsisAuthKey = types.StringValue(jsonData.IsisAuthKey)
	} else {
		v.IsisAuthKey = types.StringNull()
	}

	if jsonData.IsisAuthKeychainKeyId != nil {
		if jsonData.IsisAuthKeychainKeyId.IsEmpty() {
			v.IsisAuthKeychainKeyId = types.Int64Null()
		} else {
			v.IsisAuthKeychainKeyId = types.Int64Value(int64(*jsonData.IsisAuthKeychainKeyId))
		}
	} else {
		v.IsisAuthKeychainKeyId = types.Int64Null()
	}

	if jsonData.IsisAuthKeychainName != "" {
		v.IsisAuthKeychainName = types.StringValue(jsonData.IsisAuthKeychainName)
	} else {
		v.IsisAuthKeychainName = types.StringNull()
	}

	if jsonData.IsisLevel != "" {
		v.IsisLevel = types.StringValue(jsonData.IsisLevel)
	} else {
		v.IsisLevel = types.StringNull()
	}

	if jsonData.IsisP2pEnable != "" {
		x, _ := strconv.ParseBool(jsonData.IsisP2pEnable)
		v.IsisP2pEnable = types.BoolValue(x)
	} else {
		v.IsisP2pEnable = types.BoolNull()
	}

	if jsonData.L2HostIntfMtu != nil {
		if jsonData.L2HostIntfMtu.IsEmpty() {
			v.L2HostIntfMtu = types.Int64Null()
		} else {
			v.L2HostIntfMtu = types.Int64Value(int64(*jsonData.L2HostIntfMtu))
		}
	} else {
		v.L2HostIntfMtu = types.Int64Null()
	}

	if jsonData.LinkStateRouting != "" {
		v.LinkStateRouting = types.StringValue(jsonData.LinkStateRouting)
	} else {
		v.LinkStateRouting = types.StringNull()
	}

	if jsonData.LinkStateRoutingTag != "" {
		v.LinkStateRoutingTag = types.StringValue(jsonData.LinkStateRoutingTag)
	} else {
		v.LinkStateRoutingTag = types.StringNull()
	}

	if jsonData.Loopback0IpRange != "" {
		v.Loopback0IpRange = types.StringValue(jsonData.Loopback0IpRange)
	} else {
		v.Loopback0IpRange = types.StringNull()
	}

	if jsonData.MgmtGw != "" {
		v.MgmtGw = types.StringValue(jsonData.MgmtGw)
	} else {
		v.MgmtGw = types.StringNull()
	}

	if jsonData.MgmtPrefix != nil {
		if jsonData.MgmtPrefix.IsEmpty() {
			v.MgmtPrefix = types.Int64Null()
		} else {
			v.MgmtPrefix = types.Int64Value(int64(*jsonData.MgmtPrefix))
		}
	} else {
		v.MgmtPrefix = types.Int64Null()
	}

	if jsonData.NtpServerIpList != "" {
		v.NtpServerIpList = types.StringValue(jsonData.NtpServerIpList)
	} else {
		v.NtpServerIpList = types.StringNull()
	}

	if jsonData.NtpServerVrf != "" {
		v.NtpServerVrf = types.StringValue(jsonData.NtpServerVrf)
	} else {
		v.NtpServerVrf = types.StringNull()
	}

	if jsonData.NxapiVrf != "" {
		v.NxapiVrf = types.StringValue(jsonData.NxapiVrf)
	} else {
		v.NxapiVrf = types.StringNull()
	}

	if jsonData.OspfAreaId != "" {
		v.OspfAreaId = types.StringValue(jsonData.OspfAreaId)
	} else {
		v.OspfAreaId = types.StringNull()
	}

	if jsonData.OspfAuthEnable != "" {
		x, _ := strconv.ParseBool(jsonData.OspfAuthEnable)
		v.OspfAuthEnable = types.BoolValue(x)
	} else {
		v.OspfAuthEnable = types.BoolNull()
	}

	if jsonData.OspfAuthKey != "" {
		v.OspfAuthKey = types.StringValue(jsonData.OspfAuthKey)
	} else {
		v.OspfAuthKey = types.StringNull()
	}

	if jsonData.OspfAuthKeyId != nil {
		if jsonData.OspfAuthKeyId.IsEmpty() {
			v.OspfAuthKeyId = types.Int64Null()
		} else {
			v.OspfAuthKeyId = types.Int64Value(int64(*jsonData.OspfAuthKeyId))
		}
	} else {
		v.OspfAuthKeyId = types.Int64Null()
	}

	if jsonData.PimHelloAuthEnable != "" {
		x, _ := strconv.ParseBool(jsonData.PimHelloAuthEnable)
		v.PimHelloAuthEnable = types.BoolValue(x)
	} else {
		v.PimHelloAuthEnable = types.BoolNull()
	}

	if jsonData.PimHelloAuthKey != "" {
		v.PimHelloAuthKey = types.StringValue(jsonData.PimHelloAuthKey)
	} else {
		v.PimHelloAuthKey = types.StringNull()
	}

	if jsonData.PmEnable != "" {
		x, _ := strconv.ParseBool(jsonData.PmEnable)
		v.PmEnable = types.BoolValue(x)
	} else {
		v.PmEnable = types.BoolNull()
	}

	if jsonData.PowerRedundancyMode != "" {
		v.PowerRedundancyMode = types.StringValue(jsonData.PowerRedundancyMode)
	} else {
		v.PowerRedundancyMode = types.StringNull()
	}

	if jsonData.PtpDomainId != nil {
		if jsonData.PtpDomainId.IsEmpty() {
			v.PtpDomainId = types.Int64Null()
		} else {
			v.PtpDomainId = types.Int64Value(int64(*jsonData.PtpDomainId))
		}
	} else {
		v.PtpDomainId = types.Int64Null()
	}

	if jsonData.PtpLbId != nil {
		if jsonData.PtpLbId.IsEmpty() {
			v.PtpLbId = types.Int64Null()
		} else {
			v.PtpLbId = types.Int64Value(int64(*jsonData.PtpLbId))
		}
	} else {
		v.PtpLbId = types.Int64Null()
	}

	if jsonData.PtpProfile != "" {
		v.PtpProfile = types.StringValue(jsonData.PtpProfile)
	} else {
		v.PtpProfile = types.StringNull()
	}

	if jsonData.RoutingLbId != nil {
		if jsonData.RoutingLbId.IsEmpty() {
			v.RoutingLbId = types.Int64Null()
		} else {
			v.RoutingLbId = types.Int64Value(int64(*jsonData.RoutingLbId))
		}
	} else {
		v.RoutingLbId = types.Int64Null()
	}

	if jsonData.RpIpRange != "" {
		v.RpIpRange = types.StringValue(jsonData.RpIpRange)
	} else {
		v.RpIpRange = types.StringNull()
	}

	if jsonData.RpLbId != nil {
		if jsonData.RpLbId.IsEmpty() {
			v.RpLbId = types.Int64Null()
		} else {
			v.RpLbId = types.Int64Value(int64(*jsonData.RpLbId))
		}
	} else {
		v.RpLbId = types.Int64Null()
	}

	if jsonData.SnmpServerHostTrap != "" {
		x, _ := strconv.ParseBool(jsonData.SnmpServerHostTrap)
		v.SnmpServerHostTrap = types.BoolValue(x)
	} else {
		v.SnmpServerHostTrap = types.BoolNull()
	}

	if jsonData.StaticUnderlayIpAlloc != "" {
		x, _ := strconv.ParseBool(jsonData.StaticUnderlayIpAlloc)
		v.StaticUnderlayIpAlloc = types.BoolValue(x)
	} else {
		v.StaticUnderlayIpAlloc = types.BoolNull()
	}

	if jsonData.SubnetRange != "" {
		v.SubnetRange = types.StringValue(jsonData.SubnetRange)
	} else {
		v.SubnetRange = types.StringNull()
	}

	if jsonData.SubnetTargetMask != nil {
		if jsonData.SubnetTargetMask.IsEmpty() {
			v.SubnetTargetMask = types.Int64Null()
		} else {
			v.SubnetTargetMask = types.Int64Value(int64(*jsonData.SubnetTargetMask))
		}
	} else {
		v.SubnetTargetMask = types.Int64Null()
	}

	if jsonData.SyslogServerIpList != "" {
		v.SyslogServerIpList = types.StringValue(jsonData.SyslogServerIpList)
	} else {
		v.SyslogServerIpList = types.StringNull()
	}

	if jsonData.SyslogServerVrf != "" {
		v.SyslogServerVrf = types.StringValue(jsonData.SyslogServerVrf)
	} else {
		v.SyslogServerVrf = types.StringNull()
	}

	if jsonData.SyslogSev != "" {
		v.SyslogSev = types.StringValue(jsonData.SyslogSev)
	} else {
		v.SyslogSev = types.StringNull()
	}

	v.Deploy = types.BoolValue(jsonData.Deploy)
	if jsonData.DeploymentStatus != "" {
		v.DeploymentStatus = types.StringValue(jsonData.DeploymentStatus)
	} else {
		v.DeploymentStatus = types.StringNull()
	}

	return err
}

func (v IpFabricMediaModel) GetModelData() *resource_fabric_common.NDFCFabricCommonModel {
	var data = new(resource_fabric_common.NDFCFabricCommonModel)

	//MARSHAL_BODY

	if !v.FabricName.IsNull() && !v.FabricName.IsUnknown() {
		data.FabricName = v.FabricName.ValueString()
	} else {
		data.FabricName = ""
	}

	if !v.AaaRemoteIpEnabled.IsNull() && !v.AaaRemoteIpEnabled.IsUnknown() {
		data.AaaRemoteIpEnabled = strconv.FormatBool(v.AaaRemoteIpEnabled.ValueBool())
	} else {
		data.AaaRemoteIpEnabled = ""
	}

	if !v.AaaServerConf.IsNull() && !v.AaaServerConf.IsUnknown() {
		data.AaaServerConf = v.AaaServerConf.ValueString()
	} else {
		data.AaaServerConf = ""
	}

	if !v.AsmGroupRanges.IsNull() && !v.AsmGroupRanges.IsUnknown() {
		data.AsmGroupRanges = v.AsmGroupRanges.ValueString()
	} else {
		data.AsmGroupRanges = ""
	}

	if !v.BootstrapConf.IsNull() && !v.BootstrapConf.IsUnknown() {
		data.BootstrapConf = v.BootstrapConf.ValueString()
	} else {
		data.BootstrapConf = ""
	}

	if !v.BootstrapEnable.IsNull() && !v.BootstrapEnable.IsUnknown() {
		data.BootstrapEnable = strconv.FormatBool(v.BootstrapEnable.ValueBool())
	} else {
		data.BootstrapEnable = ""
	}

	if !v.BootstrapMultisubnet.IsNull() && !v.BootstrapMultisubnet.IsUnknown() {
		data.BootstrapMultisubnet = v.BootstrapMultisubnet.ValueString()
	} else {
		data.BootstrapMultisubnet = ""
	}

	if !v.CdpEnable.IsNull() && !v.CdpEnable.IsUnknown() {
		data.CdpEnable = strconv.FormatBool(v.CdpEnable.ValueBool())
	} else {
		data.CdpEnable = ""
	}

	if !v.DhcpEnable.IsNull() && !v.DhcpEnable.IsUnknown() {
		data.DhcpEnable = strconv.FormatBool(v.DhcpEnable.ValueBool())
	} else {
		data.DhcpEnable = ""
	}

	if !v.DhcpEnd.IsNull() && !v.DhcpEnd.IsUnknown() {
		data.DhcpEnd = v.DhcpEnd.ValueString()
	} else {
		data.DhcpEnd = ""
	}

	if !v.DhcpIpv6Enable.IsNull() && !v.DhcpIpv6Enable.IsUnknown() {
		data.DhcpIpv6Enable = v.DhcpIpv6Enable.ValueString()
	} else {
		data.DhcpIpv6Enable = ""
	}

	if !v.DhcpStart.IsNull() && !v.DhcpStart.IsUnknown() {
		data.DhcpStart = v.DhcpStart.ValueString()
	} else {
		data.DhcpStart = ""
	}

	if !v.DnsServerIpList.IsNull() && !v.DnsServerIpList.IsUnknown() {
		data.DnsServerIpList = v.DnsServerIpList.ValueString()
	} else {
		data.DnsServerIpList = ""
	}

	if !v.DnsServerVrf.IsNull() && !v.DnsServerVrf.IsUnknown() {
		data.DnsServerVrf = v.DnsServerVrf.ValueString()
	} else {
		data.DnsServerVrf = ""
	}

	if !v.EnableAaa.IsNull() && !v.EnableAaa.IsUnknown() {
		data.EnableAaa = strconv.FormatBool(v.EnableAaa.ValueBool())
	} else {
		data.EnableAaa = ""
	}

	if !v.EnableAsm.IsNull() && !v.EnableAsm.IsUnknown() {
		data.EnableAsm = strconv.FormatBool(v.EnableAsm.ValueBool())
	} else {
		data.EnableAsm = ""
	}

	if !v.EnableNbmPassive.IsNull() && !v.EnableNbmPassive.IsUnknown() {
		data.EnableNbmPassive = strconv.FormatBool(v.EnableNbmPassive.ValueBool())
	} else {
		data.EnableNbmPassive = ""
	}

	if !v.ExtraConfIntraLinks.IsNull() && !v.ExtraConfIntraLinks.IsUnknown() {
		data.ExtraConfIntraLinks = v.ExtraConfIntraLinks.ValueString()
	} else {
		data.ExtraConfIntraLinks = ""
	}

	if !v.ExtraConfLeaf.IsNull() && !v.ExtraConfLeaf.IsUnknown() {
		data.ExtraConfLeaf = v.ExtraConfLeaf.ValueString()
	} else {
		data.ExtraConfLeaf = ""
	}

	if !v.ExtraConfSpine.IsNull() && !v.ExtraConfSpine.IsUnknown() {
		data.ExtraConfSpine = v.ExtraConfSpine.ValueString()
	} else {
		data.ExtraConfSpine = ""
	}

	if !v.FabricInterfaceType.IsNull() && !v.FabricInterfaceType.IsUnknown() {
		data.FabricInterfaceType = v.FabricInterfaceType.ValueString()
	} else {
		data.FabricInterfaceType = ""
	}

	if !v.FabricMtu.IsNull() && !v.FabricMtu.IsUnknown() {
		data.FabricMtu = new(Int64Custom)
		*data.FabricMtu = Int64Custom(v.FabricMtu.ValueInt64())
	} else {
		data.FabricMtu = nil
	}

	if !v.FeaturePtp.IsNull() && !v.FeaturePtp.IsUnknown() {
		data.FeaturePtp = strconv.FormatBool(v.FeaturePtp.ValueBool())
	} else {
		data.FeaturePtp = ""
	}

	if !v.IsisAuthEnable.IsNull() && !v.IsisAuthEnable.IsUnknown() {
		data.IsisAuthEnable = strconv.FormatBool(v.IsisAuthEnable.ValueBool())
	} else {
		data.IsisAuthEnable = ""
	}

	if !v.IsisAuthKey.IsNull() && !v.IsisAuthKey.IsUnknown() {
		data.IsisAuthKey = v.IsisAuthKey.ValueString()
	} else {
		data.IsisAuthKey = ""
	}

	if !v.IsisAuthKeychainKeyId.IsNull() && !v.IsisAuthKeychainKeyId.IsUnknown() {
		data.IsisAuthKeychainKeyId = new(Int64Custom)
		*data.IsisAuthKeychainKeyId = Int64Custom(v.IsisAuthKeychainKeyId.ValueInt64())
	} else {
		data.IsisAuthKeychainKeyId = nil
	}

	if !v.IsisAuthKeychainName.IsNull() && !v.IsisAuthKeychainName.IsUnknown() {
		data.IsisAuthKeychainName = v.IsisAuthKeychainName.ValueString()
	} else {
		data.IsisAuthKeychainName = ""
	}

	if !v.IsisLevel.IsNull() && !v.IsisLevel.IsUnknown() {
		data.IsisLevel = v.IsisLevel.ValueString()
	} else {
		data.IsisLevel = ""
	}

	if !v.IsisP2pEnable.IsNull() && !v.IsisP2pEnable.IsUnknown() {
		data.IsisP2pEnable = strconv.FormatBool(v.IsisP2pEnable.ValueBool())
	} else {
		data.IsisP2pEnable = ""
	}

	if !v.L2HostIntfMtu.IsNull() && !v.L2HostIntfMtu.IsUnknown() {
		data.L2HostIntfMtu = new(Int64Custom)
		*data.L2HostIntfMtu = Int64Custom(v.L2HostIntfMtu.ValueInt64())
	} else {
		data.L2HostIntfMtu = nil
	}

	if !v.LinkStateRouting.IsNull() && !v.LinkStateRouting.IsUnknown() {
		data.LinkStateRouting = v.LinkStateRouting.ValueString()
	} else {
		data.LinkStateRouting = ""
	}

	if !v.LinkStateRoutingTag.IsNull() && !v.LinkStateRoutingTag.IsUnknown() {
		data.LinkStateRoutingTag = v.LinkStateRoutingTag.ValueString()
	} else {
		data.LinkStateRoutingTag = ""
	}

	if !v.Loopback0IpRange.IsNull() && !v.Loopback0IpRange.IsUnknown() {
		data.Loopback0IpRange = v.Loopback0IpRange.ValueString()
	} else {
		data.Loopback0IpRange = ""
	}

	if !v.MgmtGw.IsNull() && !v.MgmtGw.IsUnknown() {
		data.MgmtGw = v.MgmtGw.ValueString()
	} else {
		data.MgmtGw = ""
	}

	if !v.MgmtPrefix.IsNull() && !v.MgmtPrefix.IsUnknown() {
		data.MgmtPrefix = new(Int64Custom)
		*data.MgmtPrefix = Int64Custom(v.MgmtPrefix.ValueInt64())
	} else {
		data.MgmtPrefix = nil
	}

	if !v.NtpServerIpList.IsNull() && !v.NtpServerIpList.IsUnknown() {
		data.NtpServerIpList = v.NtpServerIpList.ValueString()
	} else {
		data.NtpServerIpList = ""
	}

	if !v.NtpServerVrf.IsNull() && !v.NtpServerVrf.IsUnknown() {
		data.NtpServerVrf = v.NtpServerVrf.ValueString()
	} else {
		data.NtpServerVrf = ""
	}

	if !v.NxapiVrf.IsNull() && !v.NxapiVrf.IsUnknown() {
		data.NxapiVrf = v.NxapiVrf.ValueString()
	} else {
		data.NxapiVrf = ""
	}

	if !v.OspfAreaId.IsNull() && !v.OspfAreaId.IsUnknown() {
		data.OspfAreaId = v.OspfAreaId.ValueString()
	} else {
		data.OspfAreaId = ""
	}

	if !v.OspfAuthEnable.IsNull() && !v.OspfAuthEnable.IsUnknown() {
		data.OspfAuthEnable = strconv.FormatBool(v.OspfAuthEnable.ValueBool())
	} else {
		data.OspfAuthEnable = ""
	}

	if !v.OspfAuthKey.IsNull() && !v.OspfAuthKey.IsUnknown() {
		data.OspfAuthKey = v.OspfAuthKey.ValueString()
	} else {
		data.OspfAuthKey = ""
	}

	if !v.OspfAuthKeyId.IsNull() && !v.OspfAuthKeyId.IsUnknown() {
		data.OspfAuthKeyId = new(Int64Custom)
		*data.OspfAuthKeyId = Int64Custom(v.OspfAuthKeyId.ValueInt64())
	} else {
		data.OspfAuthKeyId = nil
	}

	if !v.PimHelloAuthEnable.IsNull() && !v.PimHelloAuthEnable.IsUnknown() {
		data.PimHelloAuthEnable = strconv.FormatBool(v.PimHelloAuthEnable.ValueBool())
	} else {
		data.PimHelloAuthEnable = ""
	}

	if !v.PimHelloAuthKey.IsNull() && !v.PimHelloAuthKey.IsUnknown() {
		data.PimHelloAuthKey = v.PimHelloAuthKey.ValueString()
	} else {
		data.PimHelloAuthKey = ""
	}

	if !v.PmEnable.IsNull() && !v.PmEnable.IsUnknown() {
		data.PmEnable = strconv.FormatBool(v.PmEnable.ValueBool())
	} else {
		data.PmEnable = ""
	}

	if !v.PowerRedundancyMode.IsNull() && !v.PowerRedundancyMode.IsUnknown() {
		data.PowerRedundancyMode = v.PowerRedundancyMode.ValueString()
	} else {
		data.PowerRedundancyMode = ""
	}

	if !v.PtpDomainId.IsNull() && !v.PtpDomainId.IsUnknown() {
		data.PtpDomainId = new(Int64Custom)
		*data.PtpDomainId = Int64Custom(v.PtpDomainId.ValueInt64())
	} else {
		data.PtpDomainId = nil
	}

	if !v.PtpLbId.IsNull() && !v.PtpLbId.IsUnknown() {
		data.PtpLbId = new(Int64Custom)
		*data.PtpLbId = Int64Custom(v.PtpLbId.ValueInt64())
	} else {
		data.PtpLbId = nil
	}

	if !v.PtpProfile.IsNull() && !v.PtpProfile.IsUnknown() {
		data.PtpProfile = v.PtpProfile.ValueString()
	} else {
		data.PtpProfile = ""
	}

	if !v.RoutingLbId.IsNull() && !v.RoutingLbId.IsUnknown() {
		data.RoutingLbId = new(Int64Custom)
		*data.RoutingLbId = Int64Custom(v.RoutingLbId.ValueInt64())
	} else {
		data.RoutingLbId = nil
	}

	if !v.RpIpRange.IsNull() && !v.RpIpRange.IsUnknown() {
		data.RpIpRange = v.RpIpRange.ValueString()
	} else {
		data.RpIpRange = ""
	}

	if !v.RpLbId.IsNull() && !v.RpLbId.IsUnknown() {
		data.RpLbId = new(Int64Custom)
		*data.RpLbId = Int64Custom(v.RpLbId.ValueInt64())
	} else {
		data.RpLbId = nil
	}

	if !v.SnmpServerHostTrap.IsNull() && !v.SnmpServerHostTrap.IsUnknown() {
		data.SnmpServerHostTrap = strconv.FormatBool(v.SnmpServerHostTrap.ValueBool())
	} else {
		data.SnmpServerHostTrap = ""
	}

	if !v.StaticUnderlayIpAlloc.IsNull() && !v.StaticUnderlayIpAlloc.IsUnknown() {
		data.StaticUnderlayIpAlloc = strconv.FormatBool(v.StaticUnderlayIpAlloc.ValueBool())
	} else {
		data.StaticUnderlayIpAlloc = ""
	}

	if !v.SubnetRange.IsNull() && !v.SubnetRange.IsUnknown() {
		data.SubnetRange = v.SubnetRange.ValueString()
	} else {
		data.SubnetRange = ""
	}

	if !v.SubnetTargetMask.IsNull() && !v.SubnetTargetMask.IsUnknown() {
		data.SubnetTargetMask = new(Int64Custom)
		*data.SubnetTargetMask = Int64Custom(v.SubnetTargetMask.ValueInt64())
	} else {
		data.SubnetTargetMask = nil
	}

	if !v.SyslogServerIpList.IsNull() && !v.SyslogServerIpList.IsUnknown() {
		data.SyslogServerIpList = v.SyslogServerIpList.ValueString()
	} else {
		data.SyslogServerIpList = ""
	}

	if !v.SyslogServerVrf.IsNull() && !v.SyslogServerVrf.IsUnknown() {
		data.SyslogServerVrf = v.SyslogServerVrf.ValueString()
	} else {
		data.SyslogServerVrf = ""
	}

	if !v.SyslogSev.IsNull() && !v.SyslogSev.IsUnknown() {
		data.SyslogSev = v.SyslogSev.ValueString()
	} else {
		data.SyslogSev = ""
	}

	if !v.Deploy.IsNull() && !v.Deploy.IsUnknown() {
		data.Deploy = v.Deploy.ValueBool()
	}

	return data
}
