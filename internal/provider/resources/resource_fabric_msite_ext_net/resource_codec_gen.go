// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package resource_fabric_msite_ext_net

import (
	"strconv"
	. "terraform-provider-ndfc/internal/provider/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"terraform-provider-ndfc/internal/provider/resources/resource_fabric_common"
)

func (v *FabricMsiteExtNetModel) SetModelData(jsonData *resource_fabric_common.NDFCFabricCommonModel) diag.Diagnostics {
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

	if jsonData.BgpAs != "" {
		v.BgpAs = types.StringValue(jsonData.BgpAs)
	} else {
		v.BgpAs = types.StringNull()
	}

	if jsonData.BootstrapConf != "" {
		v.BootstrapConf = types.StringValue(jsonData.BootstrapConf)
	} else {
		v.BootstrapConf = types.StringNull()
	}

	if jsonData.BootstrapConfXe != "" {
		v.BootstrapConfXe = types.StringValue(jsonData.BootstrapConfXe)
	} else {
		v.BootstrapConfXe = types.StringNull()
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

	if jsonData.DomainName != "" {
		v.DomainName = types.StringValue(jsonData.DomainName)
	} else {
		v.DomainName = types.StringNull()
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

	if jsonData.FabricFreeform != "" {
		v.FabricFreeform = types.StringValue(jsonData.FabricFreeform)
	} else {
		v.FabricFreeform = types.StringNull()
	}

	if jsonData.FeaturePtp != "" {
		x, _ := strconv.ParseBool(jsonData.FeaturePtp)
		v.FeaturePtp = types.BoolValue(x)
	} else {
		v.FeaturePtp = types.BoolNull()
	}

	if jsonData.InbandEnable != "" {
		x, _ := strconv.ParseBool(jsonData.InbandEnable)
		v.InbandEnable = types.BoolValue(x)
	} else {
		v.InbandEnable = types.BoolNull()
	}

	if jsonData.InbandMgmt != "" {
		x, _ := strconv.ParseBool(jsonData.InbandMgmt)
		v.InbandMgmt = types.BoolValue(x)
	} else {
		v.InbandMgmt = types.BoolNull()
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

	if jsonData.MgmtV6prefix != nil {
		if jsonData.MgmtV6prefix.IsEmpty() {
			v.MgmtV6prefix = types.Int64Null()
		} else {
			v.MgmtV6prefix = types.Int64Value(int64(*jsonData.MgmtV6prefix))
		}
	} else {
		v.MgmtV6prefix = types.Int64Null()
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

	if jsonData.PmEnable != "" {
		x, _ := strconv.ParseBool(jsonData.PmEnable)
		v.PmEnable = types.BoolValue(x)
	} else {
		v.PmEnable = types.BoolNull()
	}

	if jsonData.PnpEnable != "" {
		x, _ := strconv.ParseBool(jsonData.PnpEnable)
		v.PnpEnable = types.BoolValue(x)
	} else {
		v.PnpEnable = types.BoolNull()
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

func (v FabricMsiteExtNetModel) GetModelData() *resource_fabric_common.NDFCFabricCommonModel {
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

	if !v.BgpAs.IsNull() && !v.BgpAs.IsUnknown() {
		data.BgpAs = v.BgpAs.ValueString()
	} else {
		data.BgpAs = ""
	}

	if !v.BootstrapConf.IsNull() && !v.BootstrapConf.IsUnknown() {
		data.BootstrapConf = v.BootstrapConf.ValueString()
	} else {
		data.BootstrapConf = ""
	}

	if !v.BootstrapConfXe.IsNull() && !v.BootstrapConfXe.IsUnknown() {
		data.BootstrapConfXe = v.BootstrapConfXe.ValueString()
	} else {
		data.BootstrapConfXe = ""
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

	if !v.DomainName.IsNull() && !v.DomainName.IsUnknown() {
		data.DomainName = v.DomainName.ValueString()
	} else {
		data.DomainName = ""
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

	if !v.MgmtV6prefix.IsNull() && !v.MgmtV6prefix.IsUnknown() {
		data.MgmtV6prefix = new(Int64Custom)
		*data.MgmtV6prefix = Int64Custom(v.MgmtV6prefix.ValueInt64())
	} else {
		data.MgmtV6prefix = nil
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

	if !v.PmEnable.IsNull() && !v.PmEnable.IsUnknown() {
		data.PmEnable = strconv.FormatBool(v.PmEnable.ValueBool())
	} else {
		data.PmEnable = ""
	}

	if !v.PnpEnable.IsNull() && !v.PnpEnable.IsUnknown() {
		data.PnpEnable = strconv.FormatBool(v.PnpEnable.ValueBool())
	} else {
		data.PnpEnable = ""
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
