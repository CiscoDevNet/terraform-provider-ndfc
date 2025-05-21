// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package resource_fabric_lan_classic

import (
	"strconv"
	. "terraform-provider-ndfc/internal/provider/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"terraform-provider-ndfc/internal/provider/resources/resource_fabric_common"
)

func (v *FabricLanClassicModel) SetModelData(jsonData *resource_fabric_common.NDFCFabricCommonModel) diag.Diagnostics {
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

	if jsonData.AllowNxc != "" {
		x, _ := strconv.ParseBool(jsonData.AllowNxc)
		v.AllowNxc = types.BoolValue(x)
	} else {
		v.AllowNxc = types.BoolNull()
	}

	if jsonData.AllowNxcPrev != "" {
		x, _ := strconv.ParseBool(jsonData.AllowNxcPrev)
		v.AllowNxcPrev = types.BoolValue(x)
	} else {
		v.AllowNxcPrev = types.BoolNull()
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

	if jsonData.BootstrapMultisubnetInternal != "" {
		v.BootstrapMultisubnetInternal = types.StringValue(jsonData.BootstrapMultisubnetInternal)
	} else {
		v.BootstrapMultisubnetInternal = types.StringNull()
	}

	if jsonData.CdpEnable != "" {
		x, _ := strconv.ParseBool(jsonData.CdpEnable)
		v.CdpEnable = types.BoolValue(x)
	} else {
		v.CdpEnable = types.BoolNull()
	}

	if jsonData.DciSubnetRange != "" {
		v.DciSubnetRange = types.StringValue(jsonData.DciSubnetRange)
	} else {
		v.DciSubnetRange = types.StringNull()
	}

	if jsonData.DciSubnetTargetMask != nil {
		if jsonData.DciSubnetTargetMask.IsEmpty() {
			v.DciSubnetTargetMask = types.Int64Null()
		} else {
			v.DciSubnetTargetMask = types.Int64Value(int64(*jsonData.DciSubnetTargetMask))
		}
	} else {
		v.DciSubnetTargetMask = types.Int64Null()
	}

	if jsonData.DeploymentFreeze != "" {
		x, _ := strconv.ParseBool(jsonData.DeploymentFreeze)
		v.DeploymentFreeze = types.BoolValue(x)
	} else {
		v.DeploymentFreeze = types.BoolNull()
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

	if jsonData.DhcpEndInternal != "" {
		v.DhcpEndInternal = types.StringValue(jsonData.DhcpEndInternal)
	} else {
		v.DhcpEndInternal = types.StringNull()
	}

	if jsonData.DhcpIpv6Enable != "" {
		v.DhcpIpv6Enable = types.StringValue(jsonData.DhcpIpv6Enable)
	} else {
		v.DhcpIpv6Enable = types.StringNull()
	}

	if jsonData.DhcpIpv6EnableInternal != "" {
		v.DhcpIpv6EnableInternal = types.StringValue(jsonData.DhcpIpv6EnableInternal)
	} else {
		v.DhcpIpv6EnableInternal = types.StringNull()
	}

	if jsonData.DhcpStart != "" {
		v.DhcpStart = types.StringValue(jsonData.DhcpStart)
	} else {
		v.DhcpStart = types.StringNull()
	}

	if jsonData.DhcpStartInternal != "" {
		v.DhcpStartInternal = types.StringValue(jsonData.DhcpStartInternal)
	} else {
		v.DhcpStartInternal = types.StringNull()
	}

	if jsonData.EnableAaa != "" {
		x, _ := strconv.ParseBool(jsonData.EnableAaa)
		v.EnableAaa = types.BoolValue(x)
	} else {
		v.EnableAaa = types.BoolNull()
	}

	if jsonData.EnableNetflow != "" {
		x, _ := strconv.ParseBool(jsonData.EnableNetflow)
		v.EnableNetflow = types.BoolValue(x)
	} else {
		v.EnableNetflow = types.BoolNull()
	}

	if jsonData.EnableNetflowPrev != "" {
		x, _ := strconv.ParseBool(jsonData.EnableNetflowPrev)
		v.EnableNetflowPrev = types.BoolValue(x)
	} else {
		v.EnableNetflowPrev = types.BoolNull()
	}

	if jsonData.EnableNxapi != "" {
		x, _ := strconv.ParseBool(jsonData.EnableNxapi)
		v.EnableNxapi = types.BoolValue(x)
	} else {
		v.EnableNxapi = types.BoolNull()
	}

	if jsonData.EnableNxapiHttp != "" {
		x, _ := strconv.ParseBool(jsonData.EnableNxapiHttp)
		v.EnableNxapiHttp = types.BoolValue(x)
	} else {
		v.EnableNxapiHttp = types.BoolNull()
	}

	if jsonData.EnableRtIntfStats != "" {
		x, _ := strconv.ParseBool(jsonData.EnableRtIntfStats)
		v.EnableRtIntfStats = types.BoolValue(x)
	} else {
		v.EnableRtIntfStats = types.BoolNull()
	}

	if jsonData.ExtFabricType != "" {
		v.ExtFabricType = types.StringValue(jsonData.ExtFabricType)
	} else {
		v.ExtFabricType = types.StringNull()
	}

	if jsonData.FabricFreeform != "" {
		v.FabricFreeform = types.StringValue(jsonData.FabricFreeform)
	} else {
		v.FabricFreeform = types.StringNull()
	}

	if jsonData.FabricTechnology != "" {
		v.FabricTechnology = types.StringValue(jsonData.FabricTechnology)
	} else {
		v.FabricTechnology = types.StringNull()
	}

	if jsonData.FabricType != "" {
		v.FabricType = types.StringValue(jsonData.FabricType)
	} else {
		v.FabricType = types.StringNull()
	}

	if jsonData.FeaturePtp != "" {
		x, _ := strconv.ParseBool(jsonData.FeaturePtp)
		v.FeaturePtp = types.BoolValue(x)
	} else {
		v.FeaturePtp = types.BoolNull()
	}

	if jsonData.FeaturePtpInternal != "" {
		x, _ := strconv.ParseBool(jsonData.FeaturePtpInternal)
		v.FeaturePtpInternal = types.BoolValue(x)
	} else {
		v.FeaturePtpInternal = types.BoolNull()
	}

	if jsonData.Ff != "" {
		v.Ff = types.StringValue(jsonData.Ff)
	} else {
		v.Ff = types.StringNull()
	}

	if jsonData.InbandEnable != "" {
		x, _ := strconv.ParseBool(jsonData.InbandEnable)
		v.InbandEnable = types.BoolValue(x)
	} else {
		v.InbandEnable = types.BoolNull()
	}

	if jsonData.InbandEnablePrev != "" {
		x, _ := strconv.ParseBool(jsonData.InbandEnablePrev)
		v.InbandEnablePrev = types.BoolValue(x)
	} else {
		v.InbandEnablePrev = types.BoolNull()
	}

	if jsonData.InbandMgmt != "" {
		x, _ := strconv.ParseBool(jsonData.InbandMgmt)
		v.InbandMgmt = types.BoolValue(x)
	} else {
		v.InbandMgmt = types.BoolNull()
	}

	if jsonData.InbandMgmtPrev != "" {
		x, _ := strconv.ParseBool(jsonData.InbandMgmtPrev)
		v.InbandMgmtPrev = types.BoolValue(x)
	} else {
		v.InbandMgmtPrev = types.BoolNull()
	}

	if jsonData.IntfStatLoadInterval != nil {
		if jsonData.IntfStatLoadInterval.IsEmpty() {
			v.IntfStatLoadInterval = types.Int64Null()
		} else {
			v.IntfStatLoadInterval = types.Int64Value(int64(*jsonData.IntfStatLoadInterval))
		}
	} else {
		v.IntfStatLoadInterval = types.Int64Null()
	}

	if jsonData.IsReadOnly != "" {
		x, _ := strconv.ParseBool(jsonData.IsReadOnly)
		v.IsReadOnly = types.BoolValue(x)
	} else {
		v.IsReadOnly = types.BoolNull()
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

	if jsonData.MgmtGwInternal != "" {
		v.MgmtGwInternal = types.StringValue(jsonData.MgmtGwInternal)
	} else {
		v.MgmtGwInternal = types.StringNull()
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

	if jsonData.MgmtPrefixInternal != nil {
		if jsonData.MgmtPrefixInternal.IsEmpty() {
			v.MgmtPrefixInternal = types.Int64Null()
		} else {
			v.MgmtPrefixInternal = types.Int64Value(int64(*jsonData.MgmtPrefixInternal))
		}
	} else {
		v.MgmtPrefixInternal = types.Int64Null()
	}

	if jsonData.MgmtV6prefix != nil {
		if jsonData.MgmtV6prefix.IsEmpty() {
			v.MgmtV6prefix = types.Int64Null()
		} else {
			v.MgmtV6prefix = types.Int64Value(int64(*jsonData.MgmtV6prefix))
		}
	} else {
		v.MgmtV6prefix = types.Int64Null()
	}

	if jsonData.MgmtV6prefixInternal != nil {
		if jsonData.MgmtV6prefixInternal.IsEmpty() {
			v.MgmtV6prefixInternal = types.Int64Null()
		} else {
			v.MgmtV6prefixInternal = types.Int64Value(int64(*jsonData.MgmtV6prefixInternal))
		}
	} else {
		v.MgmtV6prefixInternal = types.Int64Null()
	}

	if jsonData.MplsHandoff != "" {
		x, _ := strconv.ParseBool(jsonData.MplsHandoff)
		v.MplsHandoff = types.BoolValue(x)
	} else {
		v.MplsHandoff = types.BoolNull()
	}

	if jsonData.MplsLbId != nil {
		if jsonData.MplsLbId.IsEmpty() {
			v.MplsLbId = types.Int64Null()
		} else {
			v.MplsLbId = types.Int64Value(int64(*jsonData.MplsLbId))
		}
	} else {
		v.MplsLbId = types.Int64Null()
	}

	if jsonData.MplsLoopbackIpRange != "" {
		v.MplsLoopbackIpRange = types.StringValue(jsonData.MplsLoopbackIpRange)
	} else {
		v.MplsLoopbackIpRange = types.StringNull()
	}

	if jsonData.NetflowExporterList != "" {
		v.NetflowExporterList = types.StringValue(jsonData.NetflowExporterList)
	} else {
		v.NetflowExporterList = types.StringNull()
	}

	if jsonData.NetflowMonitorList != "" {
		v.NetflowMonitorList = types.StringValue(jsonData.NetflowMonitorList)
	} else {
		v.NetflowMonitorList = types.StringNull()
	}

	if jsonData.NetflowRecordList != "" {
		v.NetflowRecordList = types.StringValue(jsonData.NetflowRecordList)
	} else {
		v.NetflowRecordList = types.StringNull()
	}

	if jsonData.NetflowSamplerList != "" {
		v.NetflowSamplerList = types.StringValue(jsonData.NetflowSamplerList)
	} else {
		v.NetflowSamplerList = types.StringNull()
	}

	if jsonData.NxapiHttpsPort != nil {
		if jsonData.NxapiHttpsPort.IsEmpty() {
			v.NxapiHttpsPort = types.Int64Null()
		} else {
			v.NxapiHttpsPort = types.Int64Value(int64(*jsonData.NxapiHttpsPort))
		}
	} else {
		v.NxapiHttpsPort = types.Int64Null()
	}

	if jsonData.NxapiHttpPort != nil {
		if jsonData.NxapiHttpPort.IsEmpty() {
			v.NxapiHttpPort = types.Int64Null()
		} else {
			v.NxapiHttpPort = types.Int64Value(int64(*jsonData.NxapiHttpPort))
		}
	} else {
		v.NxapiHttpPort = types.Int64Null()
	}

	if jsonData.NxcDestVrf != "" {
		v.NxcDestVrf = types.StringValue(jsonData.NxcDestVrf)
	} else {
		v.NxcDestVrf = types.StringNull()
	}

	if jsonData.NxcProxyPort != nil {
		if jsonData.NxcProxyPort.IsEmpty() {
			v.NxcProxyPort = types.Int64Null()
		} else {
			v.NxcProxyPort = types.Int64Value(int64(*jsonData.NxcProxyPort))
		}
	} else {
		v.NxcProxyPort = types.Int64Null()
	}

	if jsonData.NxcProxyServer != "" {
		v.NxcProxyServer = types.StringValue(jsonData.NxcProxyServer)
	} else {
		v.NxcProxyServer = types.StringNull()
	}

	if jsonData.NxcSrcIntf != "" {
		v.NxcSrcIntf = types.StringValue(jsonData.NxcSrcIntf)
	} else {
		v.NxcSrcIntf = types.StringNull()
	}

	if jsonData.OverwriteGlobalNxc != "" {
		x, _ := strconv.ParseBool(jsonData.OverwriteGlobalNxc)
		v.OverwriteGlobalNxc = types.BoolValue(x)
	} else {
		v.OverwriteGlobalNxc = types.BoolNull()
	}

	if jsonData.PmEnable != "" {
		x, _ := strconv.ParseBool(jsonData.PmEnable)
		v.PmEnable = types.BoolValue(x)
	} else {
		v.PmEnable = types.BoolNull()
	}

	if jsonData.PmEnablePrev != "" {
		x, _ := strconv.ParseBool(jsonData.PmEnablePrev)
		v.PmEnablePrev = types.BoolValue(x)
	} else {
		v.PmEnablePrev = types.BoolNull()
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

	if jsonData.SnmpServerHostTrap != "" {
		x, _ := strconv.ParseBool(jsonData.SnmpServerHostTrap)
		v.SnmpServerHostTrap = types.BoolValue(x)
	} else {
		v.SnmpServerHostTrap = types.BoolNull()
	}

	if jsonData.SubinterfaceRange != "" {
		v.SubinterfaceRange = types.StringValue(jsonData.SubinterfaceRange)
	} else {
		v.SubinterfaceRange = types.StringNull()
	}

	if jsonData.EnableRealTimeBackup != "" {
		x, _ := strconv.ParseBool(jsonData.EnableRealTimeBackup)
		v.EnableRealTimeBackup = types.BoolValue(x)
	} else {
		v.EnableRealTimeBackup = types.BoolNull()
	}

	if jsonData.EnableScheduledBackup != "" {
		x, _ := strconv.ParseBool(jsonData.EnableScheduledBackup)
		v.EnableScheduledBackup = types.BoolValue(x)
	} else {
		v.EnableScheduledBackup = types.BoolNull()
	}

	if jsonData.ScheduledTime != "" {
		v.ScheduledTime = types.StringValue(jsonData.ScheduledTime)
	} else {
		v.ScheduledTime = types.StringNull()
	}

	v.Deploy = types.BoolValue(jsonData.Deploy)
	if jsonData.DeploymentStatus != "" {
		v.DeploymentStatus = types.StringValue(jsonData.DeploymentStatus)
	} else {
		v.DeploymentStatus = types.StringNull()
	}

	return err
}

func (v FabricLanClassicModel) GetModelData() *resource_fabric_common.NDFCFabricCommonModel {
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

	if !v.AllowNxc.IsNull() && !v.AllowNxc.IsUnknown() {
		data.AllowNxc = strconv.FormatBool(v.AllowNxc.ValueBool())
	} else {
		data.AllowNxc = ""
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

	if !v.DciSubnetRange.IsNull() && !v.DciSubnetRange.IsUnknown() {
		data.DciSubnetRange = v.DciSubnetRange.ValueString()
	} else {
		data.DciSubnetRange = ""
	}

	if !v.DciSubnetTargetMask.IsNull() && !v.DciSubnetTargetMask.IsUnknown() {
		data.DciSubnetTargetMask = new(Int64Custom)
		*data.DciSubnetTargetMask = Int64Custom(v.DciSubnetTargetMask.ValueInt64())
	} else {
		data.DciSubnetTargetMask = nil
	}

	if !v.DeploymentFreeze.IsNull() && !v.DeploymentFreeze.IsUnknown() {
		data.DeploymentFreeze = strconv.FormatBool(v.DeploymentFreeze.ValueBool())
	} else {
		data.DeploymentFreeze = ""
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

	if !v.EnableAaa.IsNull() && !v.EnableAaa.IsUnknown() {
		data.EnableAaa = strconv.FormatBool(v.EnableAaa.ValueBool())
	} else {
		data.EnableAaa = ""
	}

	if !v.EnableNetflow.IsNull() && !v.EnableNetflow.IsUnknown() {
		data.EnableNetflow = strconv.FormatBool(v.EnableNetflow.ValueBool())
	} else {
		data.EnableNetflow = ""
	}

	if !v.EnableNxapi.IsNull() && !v.EnableNxapi.IsUnknown() {
		data.EnableNxapi = strconv.FormatBool(v.EnableNxapi.ValueBool())
	} else {
		data.EnableNxapi = ""
	}

	if !v.EnableNxapiHttp.IsNull() && !v.EnableNxapiHttp.IsUnknown() {
		data.EnableNxapiHttp = strconv.FormatBool(v.EnableNxapiHttp.ValueBool())
	} else {
		data.EnableNxapiHttp = ""
	}

	if !v.EnableRtIntfStats.IsNull() && !v.EnableRtIntfStats.IsUnknown() {
		data.EnableRtIntfStats = strconv.FormatBool(v.EnableRtIntfStats.ValueBool())
	} else {
		data.EnableRtIntfStats = ""
	}

	if !v.ExtFabricType.IsNull() && !v.ExtFabricType.IsUnknown() {
		data.ExtFabricType = v.ExtFabricType.ValueString()
	} else {
		data.ExtFabricType = ""
	}

	if !v.FabricFreeform.IsNull() && !v.FabricFreeform.IsUnknown() {
		data.FabricFreeform = v.FabricFreeform.ValueString()
	} else {
		data.FabricFreeform = ""
	}

	if !v.FeaturePtp.IsNull() && !v.FeaturePtp.IsUnknown() {
		data.FeaturePtp = strconv.FormatBool(v.FeaturePtp.ValueBool())
	} else {
		data.FeaturePtp = ""
	}

	if !v.Ff.IsNull() && !v.Ff.IsUnknown() {
		data.Ff = v.Ff.ValueString()
	} else {
		data.Ff = ""
	}

	if !v.InbandEnable.IsNull() && !v.InbandEnable.IsUnknown() {
		data.InbandEnable = strconv.FormatBool(v.InbandEnable.ValueBool())
	} else {
		data.InbandEnable = ""
	}

	if !v.InbandMgmt.IsNull() && !v.InbandMgmt.IsUnknown() {
		data.InbandMgmt = strconv.FormatBool(v.InbandMgmt.ValueBool())
	} else {
		data.InbandMgmt = ""
	}

	if !v.IntfStatLoadInterval.IsNull() && !v.IntfStatLoadInterval.IsUnknown() {
		data.IntfStatLoadInterval = new(Int64Custom)
		*data.IntfStatLoadInterval = Int64Custom(v.IntfStatLoadInterval.ValueInt64())
	} else {
		data.IntfStatLoadInterval = nil
	}

	if !v.IsReadOnly.IsNull() && !v.IsReadOnly.IsUnknown() {
		data.IsReadOnly = strconv.FormatBool(v.IsReadOnly.ValueBool())
	} else {
		data.IsReadOnly = ""
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

	if !v.MplsHandoff.IsNull() && !v.MplsHandoff.IsUnknown() {
		data.MplsHandoff = strconv.FormatBool(v.MplsHandoff.ValueBool())
	} else {
		data.MplsHandoff = ""
	}

	if !v.MplsLbId.IsNull() && !v.MplsLbId.IsUnknown() {
		data.MplsLbId = new(Int64Custom)
		*data.MplsLbId = Int64Custom(v.MplsLbId.ValueInt64())
	} else {
		data.MplsLbId = nil
	}

	if !v.MplsLoopbackIpRange.IsNull() && !v.MplsLoopbackIpRange.IsUnknown() {
		data.MplsLoopbackIpRange = v.MplsLoopbackIpRange.ValueString()
	} else {
		data.MplsLoopbackIpRange = ""
	}

	if !v.NetflowExporterList.IsNull() && !v.NetflowExporterList.IsUnknown() {
		data.NetflowExporterList = v.NetflowExporterList.ValueString()
	} else {
		data.NetflowExporterList = ""
	}

	if !v.NetflowMonitorList.IsNull() && !v.NetflowMonitorList.IsUnknown() {
		data.NetflowMonitorList = v.NetflowMonitorList.ValueString()
	} else {
		data.NetflowMonitorList = ""
	}

	if !v.NetflowRecordList.IsNull() && !v.NetflowRecordList.IsUnknown() {
		data.NetflowRecordList = v.NetflowRecordList.ValueString()
	} else {
		data.NetflowRecordList = ""
	}

	if !v.NetflowSamplerList.IsNull() && !v.NetflowSamplerList.IsUnknown() {
		data.NetflowSamplerList = v.NetflowSamplerList.ValueString()
	} else {
		data.NetflowSamplerList = ""
	}

	if !v.NxapiHttpsPort.IsNull() && !v.NxapiHttpsPort.IsUnknown() {
		data.NxapiHttpsPort = new(Int64Custom)
		*data.NxapiHttpsPort = Int64Custom(v.NxapiHttpsPort.ValueInt64())
	} else {
		data.NxapiHttpsPort = nil
	}

	if !v.NxapiHttpPort.IsNull() && !v.NxapiHttpPort.IsUnknown() {
		data.NxapiHttpPort = new(Int64Custom)
		*data.NxapiHttpPort = Int64Custom(v.NxapiHttpPort.ValueInt64())
	} else {
		data.NxapiHttpPort = nil
	}

	if !v.NxcDestVrf.IsNull() && !v.NxcDestVrf.IsUnknown() {
		data.NxcDestVrf = v.NxcDestVrf.ValueString()
	} else {
		data.NxcDestVrf = ""
	}

	if !v.NxcProxyPort.IsNull() && !v.NxcProxyPort.IsUnknown() {
		data.NxcProxyPort = new(Int64Custom)
		*data.NxcProxyPort = Int64Custom(v.NxcProxyPort.ValueInt64())
	} else {
		data.NxcProxyPort = nil
	}

	if !v.NxcProxyServer.IsNull() && !v.NxcProxyServer.IsUnknown() {
		data.NxcProxyServer = v.NxcProxyServer.ValueString()
	} else {
		data.NxcProxyServer = ""
	}

	if !v.NxcSrcIntf.IsNull() && !v.NxcSrcIntf.IsUnknown() {
		data.NxcSrcIntf = v.NxcSrcIntf.ValueString()
	} else {
		data.NxcSrcIntf = ""
	}

	if !v.OverwriteGlobalNxc.IsNull() && !v.OverwriteGlobalNxc.IsUnknown() {
		data.OverwriteGlobalNxc = strconv.FormatBool(v.OverwriteGlobalNxc.ValueBool())
	} else {
		data.OverwriteGlobalNxc = ""
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

	if !v.SnmpServerHostTrap.IsNull() && !v.SnmpServerHostTrap.IsUnknown() {
		data.SnmpServerHostTrap = strconv.FormatBool(v.SnmpServerHostTrap.ValueBool())
	} else {
		data.SnmpServerHostTrap = ""
	}

	if !v.SubinterfaceRange.IsNull() && !v.SubinterfaceRange.IsUnknown() {
		data.SubinterfaceRange = v.SubinterfaceRange.ValueString()
	} else {
		data.SubinterfaceRange = ""
	}

	if !v.EnableRealTimeBackup.IsNull() && !v.EnableRealTimeBackup.IsUnknown() {
		data.EnableRealTimeBackup = strconv.FormatBool(v.EnableRealTimeBackup.ValueBool())
	} else {
		data.EnableRealTimeBackup = ""
	}

	if !v.EnableScheduledBackup.IsNull() && !v.EnableScheduledBackup.IsUnknown() {
		data.EnableScheduledBackup = strconv.FormatBool(v.EnableScheduledBackup.ValueBool())
	} else {
		data.EnableScheduledBackup = ""
	}

	if !v.ScheduledTime.IsNull() && !v.ScheduledTime.IsUnknown() {
		data.ScheduledTime = v.ScheduledTime.ValueString()
	} else {
		data.ScheduledTime = ""
	}

	if !v.Deploy.IsNull() && !v.Deploy.IsUnknown() {
		data.Deploy = v.Deploy.ValueBool()
	}

	return data
}
