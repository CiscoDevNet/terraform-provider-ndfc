// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package resource_fabric_vxlan_msd

import (
	"strconv"
	. "terraform-provider-ndfc/internal/provider/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"terraform-provider-ndfc/internal/provider/resources/resource_fabric_common"
)

func (v *FabricVxlanMsdModel) SetModelData(jsonData *resource_fabric_common.NDFCFabricCommonModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.FabricName != "" {
		v.FabricName = types.StringValue(jsonData.FabricName)
	} else {
		v.FabricName = types.StringNull()
	}

	if jsonData.AnycastGwMac != "" {
		v.AnycastGwMac = types.StringValue(jsonData.AnycastGwMac)
	} else {
		v.AnycastGwMac = types.StringNull()
	}

	if jsonData.BgpRpAsn != "" {
		v.BgpRpAsn = types.StringValue(jsonData.BgpRpAsn)
	} else {
		v.BgpRpAsn = types.StringNull()
	}

	if jsonData.BgwRoutingTag != nil {
		if jsonData.BgwRoutingTag.IsEmpty() {
			v.BgwRoutingTag = types.Int64Null()
		} else {
			v.BgwRoutingTag = types.Int64Value(int64(*jsonData.BgwRoutingTag))
		}
	} else {
		v.BgwRoutingTag = types.Int64Null()
	}

	if jsonData.BgwRoutingTagPrev != "" {
		v.BgwRoutingTagPrev = types.StringValue(jsonData.BgwRoutingTagPrev)
	} else {
		v.BgwRoutingTagPrev = types.StringNull()
	}

	if jsonData.BorderGwyConnections != "" {
		v.BorderGwyConnections = types.StringValue(jsonData.BorderGwyConnections)
	} else {
		v.BorderGwyConnections = types.StringNull()
	}

	if jsonData.CloudsecAlgorithm != "" {
		v.CloudsecAlgorithm = types.StringValue(jsonData.CloudsecAlgorithm)
	} else {
		v.CloudsecAlgorithm = types.StringNull()
	}

	if jsonData.CloudsecAutoconfig != "" {
		x, _ := strconv.ParseBool(jsonData.CloudsecAutoconfig)
		v.CloudsecAutoconfig = types.BoolValue(x)
	} else {
		v.CloudsecAutoconfig = types.BoolNull()
	}

	if jsonData.CloudsecEnforcement != "" {
		v.CloudsecEnforcement = types.StringValue(jsonData.CloudsecEnforcement)
	} else {
		v.CloudsecEnforcement = types.StringNull()
	}

	if jsonData.CloudsecKeyString != "" {
		v.CloudsecKeyString = types.StringValue(jsonData.CloudsecKeyString)
	} else {
		v.CloudsecKeyString = types.StringNull()
	}

	if jsonData.CloudsecReportTimer != nil {
		if jsonData.CloudsecReportTimer.IsEmpty() {
			v.CloudsecReportTimer = types.Int64Null()
		} else {
			v.CloudsecReportTimer = types.Int64Value(int64(*jsonData.CloudsecReportTimer))
		}
	} else {
		v.CloudsecReportTimer = types.Int64Null()
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

	if jsonData.DcnmId != "" {
		v.DcnmId = types.StringValue(jsonData.DcnmId)
	} else {
		v.DcnmId = types.StringNull()
	}

	if jsonData.DelayRestore != nil {
		if jsonData.DelayRestore.IsEmpty() {
			v.DelayRestore = types.Int64Null()
		} else {
			v.DelayRestore = types.Int64Value(int64(*jsonData.DelayRestore))
		}
	} else {
		v.DelayRestore = types.Int64Null()
	}

	if jsonData.EnableBgpBfd != "" {
		x, _ := strconv.ParseBool(jsonData.EnableBgpBfd)
		v.EnableBgpBfd = types.BoolValue(x)
	} else {
		v.EnableBgpBfd = types.BoolNull()
	}

	if jsonData.EnableBgpLogNeighborChange != "" {
		x, _ := strconv.ParseBool(jsonData.EnableBgpLogNeighborChange)
		v.EnableBgpLogNeighborChange = types.BoolValue(x)
	} else {
		v.EnableBgpLogNeighborChange = types.BoolNull()
	}

	if jsonData.EnableBgpSendComm != "" {
		x, _ := strconv.ParseBool(jsonData.EnableBgpSendComm)
		v.EnableBgpSendComm = types.BoolValue(x)
	} else {
		v.EnableBgpSendComm = types.BoolNull()
	}

	if jsonData.EnablePvlan != "" {
		x, _ := strconv.ParseBool(jsonData.EnablePvlan)
		v.EnablePvlan = types.BoolValue(x)
	} else {
		v.EnablePvlan = types.BoolNull()
	}

	if jsonData.EnablePvlanPrev != "" {
		x, _ := strconv.ParseBool(jsonData.EnablePvlanPrev)
		v.EnablePvlanPrev = types.BoolValue(x)
	} else {
		v.EnablePvlanPrev = types.BoolNull()
	}

	if jsonData.EnableRsRedistDirect != "" {
		x, _ := strconv.ParseBool(jsonData.EnableRsRedistDirect)
		v.EnableRsRedistDirect = types.BoolValue(x)
	} else {
		v.EnableRsRedistDirect = types.BoolNull()
	}

	if jsonData.EnableSgt != "" {
		v.EnableSgt = types.StringValue(jsonData.EnableSgt)
	} else {
		v.EnableSgt = types.StringNull()
	}

	if jsonData.EnableSgtPrev != "" {
		v.EnableSgtPrev = types.StringValue(jsonData.EnableSgtPrev)
	} else {
		v.EnableSgtPrev = types.StringNull()
	}

	if jsonData.EnableTrmTrmv6 != "" {
		x, _ := strconv.ParseBool(jsonData.EnableTrmTrmv6)
		v.EnableTrmTrmv6 = types.BoolValue(x)
	} else {
		v.EnableTrmTrmv6 = types.BoolNull()
	}

	if jsonData.EnableTrmTrmv6Prev != "" {
		x, _ := strconv.ParseBool(jsonData.EnableTrmTrmv6Prev)
		v.EnableTrmTrmv6Prev = types.BoolValue(x)
	} else {
		v.EnableTrmTrmv6Prev = types.BoolNull()
	}

	if jsonData.ExtFabricType != "" {
		v.ExtFabricType = types.StringValue(jsonData.ExtFabricType)
	} else {
		v.ExtFabricType = types.StringNull()
	}

	if jsonData.FabricType != "" {
		v.FabricType = types.StringValue(jsonData.FabricType)
	} else {
		v.FabricType = types.StringNull()
	}

	if jsonData.Ff != "" {
		v.Ff = types.StringValue(jsonData.Ff)
	} else {
		v.Ff = types.StringNull()
	}

	if jsonData.L2SegmentIdRange != "" {
		v.L2SegmentIdRange = types.StringValue(jsonData.L2SegmentIdRange)
	} else {
		v.L2SegmentIdRange = types.StringNull()
	}

	if jsonData.L3PartitionIdRange != "" {
		v.L3PartitionIdRange = types.StringValue(jsonData.L3PartitionIdRange)
	} else {
		v.L3PartitionIdRange = types.StringNull()
	}

	if jsonData.Loopback100Ipv6Range != "" {
		v.Loopback100Ipv6Range = types.StringValue(jsonData.Loopback100Ipv6Range)
	} else {
		v.Loopback100Ipv6Range = types.StringNull()
	}

	if jsonData.Loopback100IpRange != "" {
		v.Loopback100IpRange = types.StringValue(jsonData.Loopback100IpRange)
	} else {
		v.Loopback100IpRange = types.StringNull()
	}

	if jsonData.MsoControlerId != "" {
		v.MsoControlerId = types.StringValue(jsonData.MsoControlerId)
	} else {
		v.MsoControlerId = types.StringNull()
	}

	if jsonData.MsoSiteGroupName != "" {
		v.MsoSiteGroupName = types.StringValue(jsonData.MsoSiteGroupName)
	} else {
		v.MsoSiteGroupName = types.StringNull()
	}

	if jsonData.MsIfcBgpAuthKeyType != nil {
		if jsonData.MsIfcBgpAuthKeyType.IsEmpty() {
			v.MsIfcBgpAuthKeyType = types.Int64Null()
		} else {
			v.MsIfcBgpAuthKeyType = types.Int64Value(int64(*jsonData.MsIfcBgpAuthKeyType))
		}
	} else {
		v.MsIfcBgpAuthKeyType = types.Int64Null()
	}

	if jsonData.MsIfcBgpAuthKeyTypePrev != nil {
		if jsonData.MsIfcBgpAuthKeyTypePrev.IsEmpty() {
			v.MsIfcBgpAuthKeyTypePrev = types.Int64Null()
		} else {
			v.MsIfcBgpAuthKeyTypePrev = types.Int64Value(int64(*jsonData.MsIfcBgpAuthKeyTypePrev))
		}
	} else {
		v.MsIfcBgpAuthKeyTypePrev = types.Int64Null()
	}

	if jsonData.MsIfcBgpPassword != "" {
		v.MsIfcBgpPassword = types.StringValue(jsonData.MsIfcBgpPassword)
	} else {
		v.MsIfcBgpPassword = types.StringNull()
	}

	if jsonData.MsIfcBgpPasswordEnable != "" {
		x, _ := strconv.ParseBool(jsonData.MsIfcBgpPasswordEnable)
		v.MsIfcBgpPasswordEnable = types.BoolValue(x)
	} else {
		v.MsIfcBgpPasswordEnable = types.BoolNull()
	}

	if jsonData.MsIfcBgpPasswordEnablePrev != "" {
		x, _ := strconv.ParseBool(jsonData.MsIfcBgpPasswordEnablePrev)
		v.MsIfcBgpPasswordEnablePrev = types.BoolValue(x)
	} else {
		v.MsIfcBgpPasswordEnablePrev = types.BoolNull()
	}

	if jsonData.MsIfcBgpPasswordPrev != "" {
		v.MsIfcBgpPasswordPrev = types.StringValue(jsonData.MsIfcBgpPasswordPrev)
	} else {
		v.MsIfcBgpPasswordPrev = types.StringNull()
	}

	if jsonData.MsLoopbackId != nil {
		if jsonData.MsLoopbackId.IsEmpty() {
			v.MsLoopbackId = types.Int64Null()
		} else {
			v.MsLoopbackId = types.Int64Value(int64(*jsonData.MsLoopbackId))
		}
	} else {
		v.MsLoopbackId = types.Int64Null()
	}

	if jsonData.MsUnderlayAutoconfig != "" {
		x, _ := strconv.ParseBool(jsonData.MsUnderlayAutoconfig)
		v.MsUnderlayAutoconfig = types.BoolValue(x)
	} else {
		v.MsUnderlayAutoconfig = types.BoolNull()
	}

	if jsonData.ParentOnemanageFabric != "" {
		v.ParentOnemanageFabric = types.StringValue(jsonData.ParentOnemanageFabric)
	} else {
		v.ParentOnemanageFabric = types.StringNull()
	}

	if jsonData.PremsoParentFabric != "" {
		v.PremsoParentFabric = types.StringValue(jsonData.PremsoParentFabric)
	} else {
		v.PremsoParentFabric = types.StringNull()
	}

	if jsonData.RpServerIp != "" {
		v.RpServerIp = types.StringValue(jsonData.RpServerIp)
	} else {
		v.RpServerIp = types.StringNull()
	}

	if jsonData.RsRoutingTag != nil {
		if jsonData.RsRoutingTag.IsEmpty() {
			v.RsRoutingTag = types.Int64Null()
		} else {
			v.RsRoutingTag = types.Int64Value(int64(*jsonData.RsRoutingTag))
		}
	} else {
		v.RsRoutingTag = types.Int64Null()
	}

	if jsonData.SgtIdRange != "" {
		v.SgtIdRange = types.StringValue(jsonData.SgtIdRange)
	} else {
		v.SgtIdRange = types.StringNull()
	}

	if jsonData.SgtIdRangePrev != "" {
		v.SgtIdRangePrev = types.StringValue(jsonData.SgtIdRangePrev)
	} else {
		v.SgtIdRangePrev = types.StringNull()
	}

	if jsonData.SgtNamePrefix != "" {
		v.SgtNamePrefix = types.StringValue(jsonData.SgtNamePrefix)
	} else {
		v.SgtNamePrefix = types.StringNull()
	}

	if jsonData.SgtNamePrefixPrev != "" {
		v.SgtNamePrefixPrev = types.StringValue(jsonData.SgtNamePrefixPrev)
	} else {
		v.SgtNamePrefixPrev = types.StringNull()
	}

	if jsonData.SgtOperStatus != "" {
		v.SgtOperStatus = types.StringValue(jsonData.SgtOperStatus)
	} else {
		v.SgtOperStatus = types.StringNull()
	}

	if jsonData.SgtPreprovision != "" {
		x, _ := strconv.ParseBool(jsonData.SgtPreprovision)
		v.SgtPreprovision = types.BoolValue(x)
	} else {
		v.SgtPreprovision = types.BoolNull()
	}

	if jsonData.SgtPreprovisionPrev != "" {
		x, _ := strconv.ParseBool(jsonData.SgtPreprovisionPrev)
		v.SgtPreprovisionPrev = types.BoolValue(x)
	} else {
		v.SgtPreprovisionPrev = types.BoolNull()
	}

	if jsonData.SgtPreprovRecalcStatus != "" {
		v.SgtPreprovRecalcStatus = types.StringValue(jsonData.SgtPreprovRecalcStatus)
	} else {
		v.SgtPreprovRecalcStatus = types.StringNull()
	}

	if jsonData.SgtRecalcStatus != "" {
		v.SgtRecalcStatus = types.StringValue(jsonData.SgtRecalcStatus)
	} else {
		v.SgtRecalcStatus = types.StringNull()
	}

	if jsonData.TorAutoDeploy != "" {
		x, _ := strconv.ParseBool(jsonData.TorAutoDeploy)
		v.TorAutoDeploy = types.BoolValue(x)
	} else {
		v.TorAutoDeploy = types.BoolNull()
	}

	if jsonData.V6DciSubnetRange != "" {
		v.V6DciSubnetRange = types.StringValue(jsonData.V6DciSubnetRange)
	} else {
		v.V6DciSubnetRange = types.StringNull()
	}

	if jsonData.V6DciSubnetTargetMask != nil {
		if jsonData.V6DciSubnetTargetMask.IsEmpty() {
			v.V6DciSubnetTargetMask = types.Int64Null()
		} else {
			v.V6DciSubnetTargetMask = types.Int64Value(int64(*jsonData.V6DciSubnetTargetMask))
		}
	} else {
		v.V6DciSubnetTargetMask = types.Int64Null()
	}

	if jsonData.VxlanUnderlayIsV6 != "" {
		x, _ := strconv.ParseBool(jsonData.VxlanUnderlayIsV6)
		v.VxlanUnderlayIsV6 = types.BoolValue(x)
	} else {
		v.VxlanUnderlayIsV6 = types.BoolNull()
	}

	if jsonData.DefaultNetwork != "" {
		v.DefaultNetwork = types.StringValue(jsonData.DefaultNetwork)
	} else {
		v.DefaultNetwork = types.StringNull()
	}

	if jsonData.DefaultPvlanSecNetwork != "" {
		v.DefaultPvlanSecNetwork = types.StringValue(jsonData.DefaultPvlanSecNetwork)
	} else {
		v.DefaultPvlanSecNetwork = types.StringNull()
	}

	if jsonData.DefaultVrf != "" {
		v.DefaultVrf = types.StringValue(jsonData.DefaultVrf)
	} else {
		v.DefaultVrf = types.StringNull()
	}

	if jsonData.EnableScheduledBackup != "" {
		x, _ := strconv.ParseBool(jsonData.EnableScheduledBackup)
		v.EnableScheduledBackup = types.BoolValue(x)
	} else {
		v.EnableScheduledBackup = types.BoolNull()
	}

	if jsonData.NetworkExtensionTemplate != "" {
		v.NetworkExtensionTemplate = types.StringValue(jsonData.NetworkExtensionTemplate)
	} else {
		v.NetworkExtensionTemplate = types.StringNull()
	}

	if jsonData.ScheduledTime != "" {
		v.ScheduledTime = types.StringValue(jsonData.ScheduledTime)
	} else {
		v.ScheduledTime = types.StringNull()
	}

	if jsonData.VrfExtensionTemplate != "" {
		v.VrfExtensionTemplate = types.StringValue(jsonData.VrfExtensionTemplate)
	} else {
		v.VrfExtensionTemplate = types.StringNull()
	}

	v.Deploy = types.BoolValue(jsonData.Deploy)
	if jsonData.DeploymentStatus != "" {
		v.DeploymentStatus = types.StringValue(jsonData.DeploymentStatus)
	} else {
		v.DeploymentStatus = types.StringNull()
	}

	return err
}

func (v FabricVxlanMsdModel) GetModelData() *resource_fabric_common.NDFCFabricCommonModel {
	var data = new(resource_fabric_common.NDFCFabricCommonModel)

	//MARSHAL_BODY

	if !v.FabricName.IsNull() && !v.FabricName.IsUnknown() {
		data.FabricName = v.FabricName.ValueString()
	} else {
		data.FabricName = ""
	}

	if !v.AnycastGwMac.IsNull() && !v.AnycastGwMac.IsUnknown() {
		data.AnycastGwMac = v.AnycastGwMac.ValueString()
	} else {
		data.AnycastGwMac = ""
	}

	if !v.BgpRpAsn.IsNull() && !v.BgpRpAsn.IsUnknown() {
		data.BgpRpAsn = v.BgpRpAsn.ValueString()
	} else {
		data.BgpRpAsn = ""
	}

	if !v.BgwRoutingTag.IsNull() && !v.BgwRoutingTag.IsUnknown() {
		data.BgwRoutingTag = new(Int64Custom)
		*data.BgwRoutingTag = Int64Custom(v.BgwRoutingTag.ValueInt64())
	} else {
		data.BgwRoutingTag = nil
	}

	if !v.BorderGwyConnections.IsNull() && !v.BorderGwyConnections.IsUnknown() {
		data.BorderGwyConnections = v.BorderGwyConnections.ValueString()
	} else {
		data.BorderGwyConnections = ""
	}

	if !v.CloudsecAlgorithm.IsNull() && !v.CloudsecAlgorithm.IsUnknown() {
		data.CloudsecAlgorithm = v.CloudsecAlgorithm.ValueString()
	} else {
		data.CloudsecAlgorithm = ""
	}

	if !v.CloudsecAutoconfig.IsNull() && !v.CloudsecAutoconfig.IsUnknown() {
		data.CloudsecAutoconfig = strconv.FormatBool(v.CloudsecAutoconfig.ValueBool())
	} else {
		data.CloudsecAutoconfig = ""
	}

	if !v.CloudsecEnforcement.IsNull() && !v.CloudsecEnforcement.IsUnknown() {
		data.CloudsecEnforcement = v.CloudsecEnforcement.ValueString()
	} else {
		data.CloudsecEnforcement = ""
	}

	if !v.CloudsecKeyString.IsNull() && !v.CloudsecKeyString.IsUnknown() {
		data.CloudsecKeyString = v.CloudsecKeyString.ValueString()
	} else {
		data.CloudsecKeyString = ""
	}

	if !v.CloudsecReportTimer.IsNull() && !v.CloudsecReportTimer.IsUnknown() {
		data.CloudsecReportTimer = new(Int64Custom)
		*data.CloudsecReportTimer = Int64Custom(v.CloudsecReportTimer.ValueInt64())
	} else {
		data.CloudsecReportTimer = nil
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

	if !v.DelayRestore.IsNull() && !v.DelayRestore.IsUnknown() {
		data.DelayRestore = new(Int64Custom)
		*data.DelayRestore = Int64Custom(v.DelayRestore.ValueInt64())
	} else {
		data.DelayRestore = nil
	}

	if !v.EnableBgpBfd.IsNull() && !v.EnableBgpBfd.IsUnknown() {
		data.EnableBgpBfd = strconv.FormatBool(v.EnableBgpBfd.ValueBool())
	} else {
		data.EnableBgpBfd = ""
	}

	if !v.EnableBgpLogNeighborChange.IsNull() && !v.EnableBgpLogNeighborChange.IsUnknown() {
		data.EnableBgpLogNeighborChange = strconv.FormatBool(v.EnableBgpLogNeighborChange.ValueBool())
	} else {
		data.EnableBgpLogNeighborChange = ""
	}

	if !v.EnableBgpSendComm.IsNull() && !v.EnableBgpSendComm.IsUnknown() {
		data.EnableBgpSendComm = strconv.FormatBool(v.EnableBgpSendComm.ValueBool())
	} else {
		data.EnableBgpSendComm = ""
	}

	if !v.EnablePvlan.IsNull() && !v.EnablePvlan.IsUnknown() {
		data.EnablePvlan = strconv.FormatBool(v.EnablePvlan.ValueBool())
	} else {
		data.EnablePvlan = ""
	}

	if !v.EnableRsRedistDirect.IsNull() && !v.EnableRsRedistDirect.IsUnknown() {
		data.EnableRsRedistDirect = strconv.FormatBool(v.EnableRsRedistDirect.ValueBool())
	} else {
		data.EnableRsRedistDirect = ""
	}

	if !v.EnableSgt.IsNull() && !v.EnableSgt.IsUnknown() {
		data.EnableSgt = v.EnableSgt.ValueString()
	} else {
		data.EnableSgt = ""
	}

	if !v.EnableTrmTrmv6.IsNull() && !v.EnableTrmTrmv6.IsUnknown() {
		data.EnableTrmTrmv6 = strconv.FormatBool(v.EnableTrmTrmv6.ValueBool())
	} else {
		data.EnableTrmTrmv6 = ""
	}

	if !v.ExtFabricType.IsNull() && !v.ExtFabricType.IsUnknown() {
		data.ExtFabricType = v.ExtFabricType.ValueString()
	} else {
		data.ExtFabricType = ""
	}

	if !v.Ff.IsNull() && !v.Ff.IsUnknown() {
		data.Ff = v.Ff.ValueString()
	} else {
		data.Ff = ""
	}

	if !v.L2SegmentIdRange.IsNull() && !v.L2SegmentIdRange.IsUnknown() {
		data.L2SegmentIdRange = v.L2SegmentIdRange.ValueString()
	} else {
		data.L2SegmentIdRange = ""
	}

	if !v.L3PartitionIdRange.IsNull() && !v.L3PartitionIdRange.IsUnknown() {
		data.L3PartitionIdRange = v.L3PartitionIdRange.ValueString()
	} else {
		data.L3PartitionIdRange = ""
	}

	if !v.Loopback100Ipv6Range.IsNull() && !v.Loopback100Ipv6Range.IsUnknown() {
		data.Loopback100Ipv6Range = v.Loopback100Ipv6Range.ValueString()
	} else {
		data.Loopback100Ipv6Range = ""
	}

	if !v.Loopback100IpRange.IsNull() && !v.Loopback100IpRange.IsUnknown() {
		data.Loopback100IpRange = v.Loopback100IpRange.ValueString()
	} else {
		data.Loopback100IpRange = ""
	}

	if !v.MsIfcBgpAuthKeyType.IsNull() && !v.MsIfcBgpAuthKeyType.IsUnknown() {
		data.MsIfcBgpAuthKeyType = new(Int64Custom)
		*data.MsIfcBgpAuthKeyType = Int64Custom(v.MsIfcBgpAuthKeyType.ValueInt64())
	} else {
		data.MsIfcBgpAuthKeyType = nil
	}

	if !v.MsIfcBgpPassword.IsNull() && !v.MsIfcBgpPassword.IsUnknown() {
		data.MsIfcBgpPassword = v.MsIfcBgpPassword.ValueString()
	} else {
		data.MsIfcBgpPassword = ""
	}

	if !v.MsIfcBgpPasswordEnable.IsNull() && !v.MsIfcBgpPasswordEnable.IsUnknown() {
		data.MsIfcBgpPasswordEnable = strconv.FormatBool(v.MsIfcBgpPasswordEnable.ValueBool())
	} else {
		data.MsIfcBgpPasswordEnable = ""
	}

	if !v.MsLoopbackId.IsNull() && !v.MsLoopbackId.IsUnknown() {
		data.MsLoopbackId = new(Int64Custom)
		*data.MsLoopbackId = Int64Custom(v.MsLoopbackId.ValueInt64())
	} else {
		data.MsLoopbackId = nil
	}

	if !v.MsUnderlayAutoconfig.IsNull() && !v.MsUnderlayAutoconfig.IsUnknown() {
		data.MsUnderlayAutoconfig = strconv.FormatBool(v.MsUnderlayAutoconfig.ValueBool())
	} else {
		data.MsUnderlayAutoconfig = ""
	}

	if !v.RpServerIp.IsNull() && !v.RpServerIp.IsUnknown() {
		data.RpServerIp = v.RpServerIp.ValueString()
	} else {
		data.RpServerIp = ""
	}

	if !v.RsRoutingTag.IsNull() && !v.RsRoutingTag.IsUnknown() {
		data.RsRoutingTag = new(Int64Custom)
		*data.RsRoutingTag = Int64Custom(v.RsRoutingTag.ValueInt64())
	} else {
		data.RsRoutingTag = nil
	}

	if !v.SgtIdRange.IsNull() && !v.SgtIdRange.IsUnknown() {
		data.SgtIdRange = v.SgtIdRange.ValueString()
	} else {
		data.SgtIdRange = ""
	}

	if !v.SgtNamePrefix.IsNull() && !v.SgtNamePrefix.IsUnknown() {
		data.SgtNamePrefix = v.SgtNamePrefix.ValueString()
	} else {
		data.SgtNamePrefix = ""
	}

	if !v.SgtPreprovision.IsNull() && !v.SgtPreprovision.IsUnknown() {
		data.SgtPreprovision = strconv.FormatBool(v.SgtPreprovision.ValueBool())
	} else {
		data.SgtPreprovision = ""
	}

	if !v.TorAutoDeploy.IsNull() && !v.TorAutoDeploy.IsUnknown() {
		data.TorAutoDeploy = strconv.FormatBool(v.TorAutoDeploy.ValueBool())
	} else {
		data.TorAutoDeploy = ""
	}

	if !v.V6DciSubnetRange.IsNull() && !v.V6DciSubnetRange.IsUnknown() {
		data.V6DciSubnetRange = v.V6DciSubnetRange.ValueString()
	} else {
		data.V6DciSubnetRange = ""
	}

	if !v.V6DciSubnetTargetMask.IsNull() && !v.V6DciSubnetTargetMask.IsUnknown() {
		data.V6DciSubnetTargetMask = new(Int64Custom)
		*data.V6DciSubnetTargetMask = Int64Custom(v.V6DciSubnetTargetMask.ValueInt64())
	} else {
		data.V6DciSubnetTargetMask = nil
	}

	if !v.VxlanUnderlayIsV6.IsNull() && !v.VxlanUnderlayIsV6.IsUnknown() {
		data.VxlanUnderlayIsV6 = strconv.FormatBool(v.VxlanUnderlayIsV6.ValueBool())
	} else {
		data.VxlanUnderlayIsV6 = ""
	}

	if !v.DefaultNetwork.IsNull() && !v.DefaultNetwork.IsUnknown() {
		data.DefaultNetwork = v.DefaultNetwork.ValueString()
	} else {
		data.DefaultNetwork = ""
	}

	if !v.DefaultPvlanSecNetwork.IsNull() && !v.DefaultPvlanSecNetwork.IsUnknown() {
		data.DefaultPvlanSecNetwork = v.DefaultPvlanSecNetwork.ValueString()
	} else {
		data.DefaultPvlanSecNetwork = ""
	}

	if !v.DefaultVrf.IsNull() && !v.DefaultVrf.IsUnknown() {
		data.DefaultVrf = v.DefaultVrf.ValueString()
	} else {
		data.DefaultVrf = ""
	}

	if !v.EnableScheduledBackup.IsNull() && !v.EnableScheduledBackup.IsUnknown() {
		data.EnableScheduledBackup = strconv.FormatBool(v.EnableScheduledBackup.ValueBool())
	} else {
		data.EnableScheduledBackup = ""
	}

	if !v.NetworkExtensionTemplate.IsNull() && !v.NetworkExtensionTemplate.IsUnknown() {
		data.NetworkExtensionTemplate = v.NetworkExtensionTemplate.ValueString()
	} else {
		data.NetworkExtensionTemplate = ""
	}

	if !v.ScheduledTime.IsNull() && !v.ScheduledTime.IsUnknown() {
		data.ScheduledTime = v.ScheduledTime.ValueString()
	} else {
		data.ScheduledTime = ""
	}

	if !v.VrfExtensionTemplate.IsNull() && !v.VrfExtensionTemplate.IsUnknown() {
		data.VrfExtensionTemplate = v.VrfExtensionTemplate.ValueString()
	} else {
		data.VrfExtensionTemplate = ""
	}

	if !v.Deploy.IsNull() && !v.Deploy.IsUnknown() {
		data.Deploy = v.Deploy.ValueBool()
	}

	return data
}
