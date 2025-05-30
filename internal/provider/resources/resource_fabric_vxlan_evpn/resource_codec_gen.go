// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package resource_fabric_vxlan_evpn

import (
	"strconv"
	. "terraform-provider-ndfc/internal/provider/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"terraform-provider-ndfc/internal/provider/resources/resource_fabric_common"
)

func (v *FabricVxlanEvpnModel) SetModelData(jsonData *resource_fabric_common.NDFCFabricCommonModel) diag.Diagnostics {
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

	if jsonData.ActiveMigration != "" {
		x, _ := strconv.ParseBool(jsonData.ActiveMigration)
		v.ActiveMigration = types.BoolValue(x)
	} else {
		v.ActiveMigration = types.BoolNull()
	}

	if jsonData.AdvertisePipBgp != "" {
		x, _ := strconv.ParseBool(jsonData.AdvertisePipBgp)
		v.AdvertisePipBgp = types.BoolValue(x)
	} else {
		v.AdvertisePipBgp = types.BoolNull()
	}

	if jsonData.AdvertisePipOnBorder != "" {
		x, _ := strconv.ParseBool(jsonData.AdvertisePipOnBorder)
		v.AdvertisePipOnBorder = types.BoolValue(x)
	} else {
		v.AdvertisePipOnBorder = types.BoolNull()
	}

	if jsonData.AgentIntf != "" {
		v.AgentIntf = types.StringValue(jsonData.AgentIntf)
	} else {
		v.AgentIntf = types.StringNull()
	}

	if jsonData.AggAccVpcPoIdRange != "" {
		v.AggAccVpcPoIdRange = types.StringValue(jsonData.AggAccVpcPoIdRange)
	} else {
		v.AggAccVpcPoIdRange = types.StringNull()
	}

	if jsonData.AiMlQosPolicy != "" {
		v.AiMlQosPolicy = types.StringValue(jsonData.AiMlQosPolicy)
	} else {
		v.AiMlQosPolicy = types.StringNull()
	}

	if jsonData.AllowL3vniNoVlan != "" {
		x, _ := strconv.ParseBool(jsonData.AllowL3vniNoVlan)
		v.AllowL3vniNoVlan = types.BoolValue(x)
	} else {
		v.AllowL3vniNoVlan = types.BoolNull()
	}

	if jsonData.AllowL3vniNoVlanPrev != "" {
		x, _ := strconv.ParseBool(jsonData.AllowL3vniNoVlanPrev)
		v.AllowL3vniNoVlanPrev = types.BoolValue(x)
	} else {
		v.AllowL3vniNoVlanPrev = types.BoolNull()
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

	if jsonData.AnycastBgwAdvertisePip != "" {
		x, _ := strconv.ParseBool(jsonData.AnycastBgwAdvertisePip)
		v.AnycastBgwAdvertisePip = types.BoolValue(x)
	} else {
		v.AnycastBgwAdvertisePip = types.BoolNull()
	}

	if jsonData.AnycastGwMac != "" {
		v.AnycastGwMac = types.StringValue(jsonData.AnycastGwMac)
	} else {
		v.AnycastGwMac = types.StringNull()
	}

	if jsonData.AnycastLbId != nil {
		if jsonData.AnycastLbId.IsEmpty() {
			v.AnycastLbId = types.Int64Null()
		} else {
			v.AnycastLbId = types.Int64Value(int64(*jsonData.AnycastLbId))
		}
	} else {
		v.AnycastLbId = types.Int64Null()
	}

	if jsonData.AnycastRpIpRange != "" {
		v.AnycastRpIpRange = types.StringValue(jsonData.AnycastRpIpRange)
	} else {
		v.AnycastRpIpRange = types.StringNull()
	}

	if jsonData.AnycastRpIpRangeInternal != "" {
		v.AnycastRpIpRangeInternal = types.StringValue(jsonData.AnycastRpIpRangeInternal)
	} else {
		v.AnycastRpIpRangeInternal = types.StringNull()
	}

	if jsonData.AutoSymmetricDefaultVrf != "" {
		x, _ := strconv.ParseBool(jsonData.AutoSymmetricDefaultVrf)
		v.AutoSymmetricDefaultVrf = types.BoolValue(x)
	} else {
		v.AutoSymmetricDefaultVrf = types.BoolNull()
	}

	if jsonData.AutoSymmetricVrfLite != "" {
		x, _ := strconv.ParseBool(jsonData.AutoSymmetricVrfLite)
		v.AutoSymmetricVrfLite = types.BoolValue(x)
	} else {
		v.AutoSymmetricVrfLite = types.BoolNull()
	}

	if jsonData.AutoUniqueVrfLiteIpPrefix != "" {
		x, _ := strconv.ParseBool(jsonData.AutoUniqueVrfLiteIpPrefix)
		v.AutoUniqueVrfLiteIpPrefix = types.BoolValue(x)
	} else {
		v.AutoUniqueVrfLiteIpPrefix = types.BoolNull()
	}

	if jsonData.AutoUniqueVrfLiteIpPrefixPrev != "" {
		x, _ := strconv.ParseBool(jsonData.AutoUniqueVrfLiteIpPrefixPrev)
		v.AutoUniqueVrfLiteIpPrefixPrev = types.BoolValue(x)
	} else {
		v.AutoUniqueVrfLiteIpPrefixPrev = types.BoolNull()
	}

	if jsonData.AutoVrfliteIfcDefaultVrf != "" {
		x, _ := strconv.ParseBool(jsonData.AutoVrfliteIfcDefaultVrf)
		v.AutoVrfliteIfcDefaultVrf = types.BoolValue(x)
	} else {
		v.AutoVrfliteIfcDefaultVrf = types.BoolNull()
	}

	if jsonData.Banner != "" {
		v.Banner = types.StringValue(jsonData.Banner)
	} else {
		v.Banner = types.StringNull()
	}

	if jsonData.BfdAuthEnable != "" {
		x, _ := strconv.ParseBool(jsonData.BfdAuthEnable)
		v.BfdAuthEnable = types.BoolValue(x)
	} else {
		v.BfdAuthEnable = types.BoolNull()
	}

	if jsonData.BfdAuthKey != "" {
		v.BfdAuthKey = types.StringValue(jsonData.BfdAuthKey)
	} else {
		v.BfdAuthKey = types.StringNull()
	}

	if jsonData.BfdAuthKeyId != nil {
		if jsonData.BfdAuthKeyId.IsEmpty() {
			v.BfdAuthKeyId = types.Int64Null()
		} else {
			v.BfdAuthKeyId = types.Int64Value(int64(*jsonData.BfdAuthKeyId))
		}
	} else {
		v.BfdAuthKeyId = types.Int64Null()
	}

	if jsonData.BfdEnable != "" {
		x, _ := strconv.ParseBool(jsonData.BfdEnable)
		v.BfdEnable = types.BoolValue(x)
	} else {
		v.BfdEnable = types.BoolNull()
	}

	if jsonData.BfdEnablePrev != "" {
		x, _ := strconv.ParseBool(jsonData.BfdEnablePrev)
		v.BfdEnablePrev = types.BoolValue(x)
	} else {
		v.BfdEnablePrev = types.BoolNull()
	}

	if jsonData.BfdIbgpEnable != "" {
		x, _ := strconv.ParseBool(jsonData.BfdIbgpEnable)
		v.BfdIbgpEnable = types.BoolValue(x)
	} else {
		v.BfdIbgpEnable = types.BoolNull()
	}

	if jsonData.BfdIsisEnable != "" {
		x, _ := strconv.ParseBool(jsonData.BfdIsisEnable)
		v.BfdIsisEnable = types.BoolValue(x)
	} else {
		v.BfdIsisEnable = types.BoolNull()
	}

	if jsonData.BfdOspfEnable != "" {
		x, _ := strconv.ParseBool(jsonData.BfdOspfEnable)
		v.BfdOspfEnable = types.BoolValue(x)
	} else {
		v.BfdOspfEnable = types.BoolNull()
	}

	if jsonData.BfdPimEnable != "" {
		x, _ := strconv.ParseBool(jsonData.BfdPimEnable)
		v.BfdPimEnable = types.BoolValue(x)
	} else {
		v.BfdPimEnable = types.BoolNull()
	}

	if jsonData.BgpAs != "" {
		v.BgpAs = types.StringValue(jsonData.BgpAs)
	} else {
		v.BgpAs = types.StringNull()
	}

	if jsonData.BgpAsPrev != "" {
		v.BgpAsPrev = types.StringValue(jsonData.BgpAsPrev)
	} else {
		v.BgpAsPrev = types.StringNull()
	}

	if jsonData.BgpAuthEnable != "" {
		x, _ := strconv.ParseBool(jsonData.BgpAuthEnable)
		v.BgpAuthEnable = types.BoolValue(x)
	} else {
		v.BgpAuthEnable = types.BoolNull()
	}

	if jsonData.BgpAuthKey != "" {
		v.BgpAuthKey = types.StringValue(jsonData.BgpAuthKey)
	} else {
		v.BgpAuthKey = types.StringNull()
	}

	if jsonData.BgpAuthKeyType != nil {
		if jsonData.BgpAuthKeyType.IsEmpty() {
			v.BgpAuthKeyType = types.Int64Null()
		} else {
			v.BgpAuthKeyType = types.Int64Value(int64(*jsonData.BgpAuthKeyType))
		}
	} else {
		v.BgpAuthKeyType = types.Int64Null()
	}

	if jsonData.BgpLbId != nil {
		if jsonData.BgpLbId.IsEmpty() {
			v.BgpLbId = types.Int64Null()
		} else {
			v.BgpLbId = types.Int64Value(int64(*jsonData.BgpLbId))
		}
	} else {
		v.BgpLbId = types.Int64Null()
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

	if jsonData.BootstrapEnablePrev != "" {
		x, _ := strconv.ParseBool(jsonData.BootstrapEnablePrev)
		v.BootstrapEnablePrev = types.BoolValue(x)
	} else {
		v.BootstrapEnablePrev = types.BoolNull()
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

	if jsonData.BrfieldDebugFlag != "" {
		v.BrfieldDebugFlag = types.StringValue(jsonData.BrfieldDebugFlag)
	} else {
		v.BrfieldDebugFlag = types.StringNull()
	}

	if jsonData.BrownfieldNetworkNameFormat != "" {
		v.BrownfieldNetworkNameFormat = types.StringValue(jsonData.BrownfieldNetworkNameFormat)
	} else {
		v.BrownfieldNetworkNameFormat = types.StringNull()
	}

	if jsonData.BrownfieldSkipOverlayNetworkAttachments != "" {
		x, _ := strconv.ParseBool(jsonData.BrownfieldSkipOverlayNetworkAttachments)
		v.BrownfieldSkipOverlayNetworkAttachments = types.BoolValue(x)
	} else {
		v.BrownfieldSkipOverlayNetworkAttachments = types.BoolNull()
	}

	if jsonData.CdpEnable != "" {
		x, _ := strconv.ParseBool(jsonData.CdpEnable)
		v.CdpEnable = types.BoolValue(x)
	} else {
		v.CdpEnable = types.BoolNull()
	}

	if jsonData.CoppPolicy != "" {
		v.CoppPolicy = types.StringValue(jsonData.CoppPolicy)
	} else {
		v.CoppPolicy = types.StringNull()
	}

	if jsonData.DciMacsecAlgorithm != "" {
		v.DciMacsecAlgorithm = types.StringValue(jsonData.DciMacsecAlgorithm)
	} else {
		v.DciMacsecAlgorithm = types.StringNull()
	}

	if jsonData.DciMacsecCipherSuite != "" {
		v.DciMacsecCipherSuite = types.StringValue(jsonData.DciMacsecCipherSuite)
	} else {
		v.DciMacsecCipherSuite = types.StringNull()
	}

	if jsonData.DciMacsecFallbackAlgorithm != "" {
		v.DciMacsecFallbackAlgorithm = types.StringValue(jsonData.DciMacsecFallbackAlgorithm)
	} else {
		v.DciMacsecFallbackAlgorithm = types.StringNull()
	}

	if jsonData.DciMacsecFallbackKeyString != "" {
		v.DciMacsecFallbackKeyString = types.StringValue(jsonData.DciMacsecFallbackKeyString)
	} else {
		v.DciMacsecFallbackKeyString = types.StringNull()
	}

	if jsonData.DciMacsecKeyString != "" {
		v.DciMacsecKeyString = types.StringValue(jsonData.DciMacsecKeyString)
	} else {
		v.DciMacsecKeyString = types.StringNull()
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

	if jsonData.DefaultQueuingPolicyCloudscale != "" {
		v.DefaultQueuingPolicyCloudscale = types.StringValue(jsonData.DefaultQueuingPolicyCloudscale)
	} else {
		v.DefaultQueuingPolicyCloudscale = types.StringNull()
	}

	if jsonData.DefaultQueuingPolicyOther != "" {
		v.DefaultQueuingPolicyOther = types.StringValue(jsonData.DefaultQueuingPolicyOther)
	} else {
		v.DefaultQueuingPolicyOther = types.StringNull()
	}

	if jsonData.DefaultQueuingPolicyRSeries != "" {
		v.DefaultQueuingPolicyRSeries = types.StringValue(jsonData.DefaultQueuingPolicyRSeries)
	} else {
		v.DefaultQueuingPolicyRSeries = types.StringNull()
	}

	if jsonData.DefaultVrfRedisBgpRmap != "" {
		v.DefaultVrfRedisBgpRmap = types.StringValue(jsonData.DefaultVrfRedisBgpRmap)
	} else {
		v.DefaultVrfRedisBgpRmap = types.StringNull()
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

	if jsonData.EnableAgent != "" {
		x, _ := strconv.ParseBool(jsonData.EnableAgent)
		v.EnableAgent = types.BoolValue(x)
	} else {
		v.EnableAgent = types.BoolNull()
	}

	if jsonData.EnableAggAccIdRange != "" {
		x, _ := strconv.ParseBool(jsonData.EnableAggAccIdRange)
		v.EnableAggAccIdRange = types.BoolValue(x)
	} else {
		v.EnableAggAccIdRange = types.BoolNull()
	}

	if jsonData.EnableAiMlQosPolicy != "" {
		x, _ := strconv.ParseBool(jsonData.EnableAiMlQosPolicy)
		v.EnableAiMlQosPolicy = types.BoolValue(x)
	} else {
		v.EnableAiMlQosPolicy = types.BoolNull()
	}

	if jsonData.EnableAiMlQosPolicyFlap != "" {
		x, _ := strconv.ParseBool(jsonData.EnableAiMlQosPolicyFlap)
		v.EnableAiMlQosPolicyFlap = types.BoolValue(x)
	} else {
		v.EnableAiMlQosPolicyFlap = types.BoolNull()
	}

	if jsonData.EnableDciMacsec != "" {
		x, _ := strconv.ParseBool(jsonData.EnableDciMacsec)
		v.EnableDciMacsec = types.BoolValue(x)
	} else {
		v.EnableDciMacsec = types.BoolNull()
	}

	if jsonData.EnableDciMacsecPrev != "" {
		x, _ := strconv.ParseBool(jsonData.EnableDciMacsecPrev)
		v.EnableDciMacsecPrev = types.BoolValue(x)
	} else {
		v.EnableDciMacsecPrev = types.BoolNull()
	}

	if jsonData.EnableDefaultQueuingPolicy != "" {
		x, _ := strconv.ParseBool(jsonData.EnableDefaultQueuingPolicy)
		v.EnableDefaultQueuingPolicy = types.BoolValue(x)
	} else {
		v.EnableDefaultQueuingPolicy = types.BoolNull()
	}

	if jsonData.EnableEvpn != "" {
		x, _ := strconv.ParseBool(jsonData.EnableEvpn)
		v.EnableEvpn = types.BoolValue(x)
	} else {
		v.EnableEvpn = types.BoolNull()
	}

	if jsonData.EnableFabricVpcDomainId != "" {
		x, _ := strconv.ParseBool(jsonData.EnableFabricVpcDomainId)
		v.EnableFabricVpcDomainId = types.BoolValue(x)
	} else {
		v.EnableFabricVpcDomainId = types.BoolNull()
	}

	if jsonData.EnableFabricVpcDomainIdPrev != "" {
		x, _ := strconv.ParseBool(jsonData.EnableFabricVpcDomainIdPrev)
		v.EnableFabricVpcDomainIdPrev = types.BoolValue(x)
	} else {
		v.EnableFabricVpcDomainIdPrev = types.BoolNull()
	}

	if jsonData.EnableL3vniNoVlan != "" {
		x, _ := strconv.ParseBool(jsonData.EnableL3vniNoVlan)
		v.EnableL3vniNoVlan = types.BoolValue(x)
	} else {
		v.EnableL3vniNoVlan = types.BoolNull()
	}

	if jsonData.EnableMacsec != "" {
		x, _ := strconv.ParseBool(jsonData.EnableMacsec)
		v.EnableMacsec = types.BoolValue(x)
	} else {
		v.EnableMacsec = types.BoolNull()
	}

	if jsonData.EnableMacsecPrev != "" {
		x, _ := strconv.ParseBool(jsonData.EnableMacsecPrev)
		v.EnableMacsecPrev = types.BoolValue(x)
	} else {
		v.EnableMacsecPrev = types.BoolNull()
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

	if jsonData.EnableNgoam != "" {
		x, _ := strconv.ParseBool(jsonData.EnableNgoam)
		v.EnableNgoam = types.BoolValue(x)
	} else {
		v.EnableNgoam = types.BoolNull()
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

	if jsonData.EnablePbr != "" {
		x, _ := strconv.ParseBool(jsonData.EnablePbr)
		v.EnablePbr = types.BoolValue(x)
	} else {
		v.EnablePbr = types.BoolNull()
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

	if jsonData.EnableQkd != "" {
		x, _ := strconv.ParseBool(jsonData.EnableQkd)
		v.EnableQkd = types.BoolValue(x)
	} else {
		v.EnableQkd = types.BoolNull()
	}

	if jsonData.EnableRtIntfStats != "" {
		x, _ := strconv.ParseBool(jsonData.EnableRtIntfStats)
		v.EnableRtIntfStats = types.BoolValue(x)
	} else {
		v.EnableRtIntfStats = types.BoolNull()
	}

	if jsonData.EnableSgt != "" {
		x, _ := strconv.ParseBool(jsonData.EnableSgt)
		v.EnableSgt = types.BoolValue(x)
	} else {
		v.EnableSgt = types.BoolNull()
	}

	if jsonData.EnableSgtPrev != "" {
		x, _ := strconv.ParseBool(jsonData.EnableSgtPrev)
		v.EnableSgtPrev = types.BoolValue(x)
	} else {
		v.EnableSgtPrev = types.BoolNull()
	}

	if jsonData.EnableTenantDhcp != "" {
		x, _ := strconv.ParseBool(jsonData.EnableTenantDhcp)
		v.EnableTenantDhcp = types.BoolValue(x)
	} else {
		v.EnableTenantDhcp = types.BoolNull()
	}

	if jsonData.EnableTrm != "" {
		x, _ := strconv.ParseBool(jsonData.EnableTrm)
		v.EnableTrm = types.BoolValue(x)
	} else {
		v.EnableTrm = types.BoolNull()
	}

	if jsonData.EnableTrmv6 != "" {
		x, _ := strconv.ParseBool(jsonData.EnableTrmv6)
		v.EnableTrmv6 = types.BoolValue(x)
	} else {
		v.EnableTrmv6 = types.BoolNull()
	}

	if jsonData.EnableVpcPeerLinkNativeVlan != "" {
		x, _ := strconv.ParseBool(jsonData.EnableVpcPeerLinkNativeVlan)
		v.EnableVpcPeerLinkNativeVlan = types.BoolValue(x)
	} else {
		v.EnableVpcPeerLinkNativeVlan = types.BoolNull()
	}

	if jsonData.EnableVriIdRealloc != "" {
		x, _ := strconv.ParseBool(jsonData.EnableVriIdRealloc)
		v.EnableVriIdRealloc = types.BoolValue(x)
	} else {
		v.EnableVriIdRealloc = types.BoolNull()
	}

	if jsonData.EsrOption != "" {
		v.EsrOption = types.StringValue(jsonData.EsrOption)
	} else {
		v.EsrOption = types.StringNull()
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

	if jsonData.ExtraConfTor != "" {
		v.ExtraConfTor = types.StringValue(jsonData.ExtraConfTor)
	} else {
		v.ExtraConfTor = types.StringNull()
	}

	if jsonData.ExtFabricType != "" {
		v.ExtFabricType = types.StringValue(jsonData.ExtFabricType)
	} else {
		v.ExtFabricType = types.StringNull()
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

	if jsonData.FabricMtuPrev != nil {
		if jsonData.FabricMtuPrev.IsEmpty() {
			v.FabricMtuPrev = types.Int64Null()
		} else {
			v.FabricMtuPrev = types.Int64Value(int64(*jsonData.FabricMtuPrev))
		}
	} else {
		v.FabricMtuPrev = types.Int64Null()
	}

	if jsonData.FabricType != "" {
		v.FabricType = types.StringValue(jsonData.FabricType)
	} else {
		v.FabricType = types.StringNull()
	}

	if jsonData.FabricVpcDomainId != nil {
		if jsonData.FabricVpcDomainId.IsEmpty() {
			v.FabricVpcDomainId = types.Int64Null()
		} else {
			v.FabricVpcDomainId = types.Int64Value(int64(*jsonData.FabricVpcDomainId))
		}
	} else {
		v.FabricVpcDomainId = types.Int64Null()
	}

	if jsonData.FabricVpcDomainIdPrev != nil {
		if jsonData.FabricVpcDomainIdPrev.IsEmpty() {
			v.FabricVpcDomainIdPrev = types.Int64Null()
		} else {
			v.FabricVpcDomainIdPrev = types.Int64Value(int64(*jsonData.FabricVpcDomainIdPrev))
		}
	} else {
		v.FabricVpcDomainIdPrev = types.Int64Null()
	}

	if jsonData.FabricVpcQos != "" {
		x, _ := strconv.ParseBool(jsonData.FabricVpcQos)
		v.FabricVpcQos = types.BoolValue(x)
	} else {
		v.FabricVpcQos = types.BoolNull()
	}

	if jsonData.FabricVpcQosPolicyName != "" {
		v.FabricVpcQosPolicyName = types.StringValue(jsonData.FabricVpcQosPolicyName)
	} else {
		v.FabricVpcQosPolicyName = types.StringNull()
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

	if jsonData.GrfieldDebugFlag != "" {
		v.GrfieldDebugFlag = types.StringValue(jsonData.GrfieldDebugFlag)
	} else {
		v.GrfieldDebugFlag = types.StringNull()
	}

	if jsonData.HdTime != nil {
		if jsonData.HdTime.IsEmpty() {
			v.HdTime = types.Int64Null()
		} else {
			v.HdTime = types.Int64Value(int64(*jsonData.HdTime))
		}
	} else {
		v.HdTime = types.Int64Null()
	}

	if jsonData.HostIntfAdminState != "" {
		x, _ := strconv.ParseBool(jsonData.HostIntfAdminState)
		v.HostIntfAdminState = types.BoolValue(x)
	} else {
		v.HostIntfAdminState = types.BoolNull()
	}

	if jsonData.IbgpPeerTemplate != "" {
		v.IbgpPeerTemplate = types.StringValue(jsonData.IbgpPeerTemplate)
	} else {
		v.IbgpPeerTemplate = types.StringNull()
	}

	if jsonData.IbgpPeerTemplateLeaf != "" {
		v.IbgpPeerTemplateLeaf = types.StringValue(jsonData.IbgpPeerTemplateLeaf)
	} else {
		v.IbgpPeerTemplateLeaf = types.StringNull()
	}

	if jsonData.IgnoreCert != "" {
		x, _ := strconv.ParseBool(jsonData.IgnoreCert)
		v.IgnoreCert = types.BoolValue(x)
	} else {
		v.IgnoreCert = types.BoolNull()
	}

	if jsonData.InbandDhcpServers != "" {
		v.InbandDhcpServers = types.StringValue(jsonData.InbandDhcpServers)
	} else {
		v.InbandDhcpServers = types.StringNull()
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

	if jsonData.Ipv6AnycastRpIpRange != "" {
		v.Ipv6AnycastRpIpRange = types.StringValue(jsonData.Ipv6AnycastRpIpRange)
	} else {
		v.Ipv6AnycastRpIpRange = types.StringNull()
	}

	if jsonData.Ipv6AnycastRpIpRangeInternal != "" {
		v.Ipv6AnycastRpIpRangeInternal = types.StringValue(jsonData.Ipv6AnycastRpIpRangeInternal)
	} else {
		v.Ipv6AnycastRpIpRangeInternal = types.StringNull()
	}

	if jsonData.Ipv6MulticastGroupSubnet != "" {
		v.Ipv6MulticastGroupSubnet = types.StringValue(jsonData.Ipv6MulticastGroupSubnet)
	} else {
		v.Ipv6MulticastGroupSubnet = types.StringNull()
	}

	if jsonData.IsisAreaNum != "" {
		v.IsisAreaNum = types.StringValue(jsonData.IsisAreaNum)
	} else {
		v.IsisAreaNum = types.StringNull()
	}

	if jsonData.IsisAreaNumPrev != "" {
		v.IsisAreaNumPrev = types.StringValue(jsonData.IsisAreaNumPrev)
	} else {
		v.IsisAreaNumPrev = types.StringNull()
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

	if jsonData.IsisOverloadElapseTime != nil {
		if jsonData.IsisOverloadElapseTime.IsEmpty() {
			v.IsisOverloadElapseTime = types.Int64Null()
		} else {
			v.IsisOverloadElapseTime = types.Int64Value(int64(*jsonData.IsisOverloadElapseTime))
		}
	} else {
		v.IsisOverloadElapseTime = types.Int64Null()
	}

	if jsonData.IsisOverloadEnable != "" {
		x, _ := strconv.ParseBool(jsonData.IsisOverloadEnable)
		v.IsisOverloadEnable = types.BoolValue(x)
	} else {
		v.IsisOverloadEnable = types.BoolNull()
	}

	if jsonData.IsisP2pEnable != "" {
		x, _ := strconv.ParseBool(jsonData.IsisP2pEnable)
		v.IsisP2pEnable = types.BoolValue(x)
	} else {
		v.IsisP2pEnable = types.BoolNull()
	}

	if jsonData.KmeServerIp != "" {
		v.KmeServerIp = types.StringValue(jsonData.KmeServerIp)
	} else {
		v.KmeServerIp = types.StringNull()
	}

	if jsonData.KmeServerPort != nil {
		if jsonData.KmeServerPort.IsEmpty() {
			v.KmeServerPort = types.Int64Null()
		} else {
			v.KmeServerPort = types.Int64Value(int64(*jsonData.KmeServerPort))
		}
	} else {
		v.KmeServerPort = types.Int64Null()
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

	if jsonData.L2HostIntfMtuPrev != nil {
		if jsonData.L2HostIntfMtuPrev.IsEmpty() {
			v.L2HostIntfMtuPrev = types.Int64Null()
		} else {
			v.L2HostIntfMtuPrev = types.Int64Value(int64(*jsonData.L2HostIntfMtuPrev))
		}
	} else {
		v.L2HostIntfMtuPrev = types.Int64Null()
	}

	if jsonData.L2SegmentIdRange != "" {
		v.L2SegmentIdRange = types.StringValue(jsonData.L2SegmentIdRange)
	} else {
		v.L2SegmentIdRange = types.StringNull()
	}

	if jsonData.L3vniIpv6McastGroup != "" {
		v.L3vniIpv6McastGroup = types.StringValue(jsonData.L3vniIpv6McastGroup)
	} else {
		v.L3vniIpv6McastGroup = types.StringNull()
	}

	if jsonData.L3vniMcastGroup != "" {
		v.L3vniMcastGroup = types.StringValue(jsonData.L3vniMcastGroup)
	} else {
		v.L3vniMcastGroup = types.StringNull()
	}

	if jsonData.L3PartitionIdRange != "" {
		v.L3PartitionIdRange = types.StringValue(jsonData.L3PartitionIdRange)
	} else {
		v.L3PartitionIdRange = types.StringNull()
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

	if jsonData.LinkStateRoutingTagPrev != "" {
		v.LinkStateRoutingTagPrev = types.StringValue(jsonData.LinkStateRoutingTagPrev)
	} else {
		v.LinkStateRoutingTagPrev = types.StringNull()
	}

	if jsonData.Loopback0Ipv6Range != "" {
		v.Loopback0Ipv6Range = types.StringValue(jsonData.Loopback0Ipv6Range)
	} else {
		v.Loopback0Ipv6Range = types.StringNull()
	}

	if jsonData.Loopback0IpRange != "" {
		v.Loopback0IpRange = types.StringValue(jsonData.Loopback0IpRange)
	} else {
		v.Loopback0IpRange = types.StringNull()
	}

	if jsonData.Loopback1Ipv6Range != "" {
		v.Loopback1Ipv6Range = types.StringValue(jsonData.Loopback1Ipv6Range)
	} else {
		v.Loopback1Ipv6Range = types.StringNull()
	}

	if jsonData.Loopback1IpRange != "" {
		v.Loopback1IpRange = types.StringValue(jsonData.Loopback1IpRange)
	} else {
		v.Loopback1IpRange = types.StringNull()
	}

	if jsonData.MacsecAlgorithm != "" {
		v.MacsecAlgorithm = types.StringValue(jsonData.MacsecAlgorithm)
	} else {
		v.MacsecAlgorithm = types.StringNull()
	}

	if jsonData.MacsecCipherSuite != "" {
		v.MacsecCipherSuite = types.StringValue(jsonData.MacsecCipherSuite)
	} else {
		v.MacsecCipherSuite = types.StringNull()
	}

	if jsonData.MacsecFallbackAlgorithm != "" {
		v.MacsecFallbackAlgorithm = types.StringValue(jsonData.MacsecFallbackAlgorithm)
	} else {
		v.MacsecFallbackAlgorithm = types.StringNull()
	}

	if jsonData.MacsecFallbackKeyString != "" {
		v.MacsecFallbackKeyString = types.StringValue(jsonData.MacsecFallbackKeyString)
	} else {
		v.MacsecFallbackKeyString = types.StringNull()
	}

	if jsonData.MacsecKeyString != "" {
		v.MacsecKeyString = types.StringValue(jsonData.MacsecKeyString)
	} else {
		v.MacsecKeyString = types.StringNull()
	}

	if jsonData.MacsecReportTimer != nil {
		if jsonData.MacsecReportTimer.IsEmpty() {
			v.MacsecReportTimer = types.Int64Null()
		} else {
			v.MacsecReportTimer = types.Int64Value(int64(*jsonData.MacsecReportTimer))
		}
	} else {
		v.MacsecReportTimer = types.Int64Null()
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

	if jsonData.MplsIsisAreaNum != "" {
		v.MplsIsisAreaNum = types.StringValue(jsonData.MplsIsisAreaNum)
	} else {
		v.MplsIsisAreaNum = types.StringNull()
	}

	if jsonData.MplsIsisAreaNumPrev != "" {
		v.MplsIsisAreaNumPrev = types.StringValue(jsonData.MplsIsisAreaNumPrev)
	} else {
		v.MplsIsisAreaNumPrev = types.StringNull()
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

	if jsonData.MsoConnectivityDeployed != "" {
		v.MsoConnectivityDeployed = types.StringValue(jsonData.MsoConnectivityDeployed)
	} else {
		v.MsoConnectivityDeployed = types.StringNull()
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

	if jsonData.MsoSiteId != "" {
		v.MsoSiteId = types.StringValue(jsonData.MsoSiteId)
	} else {
		v.MsoSiteId = types.StringNull()
	}

	if jsonData.MstInstanceRange != "" {
		v.MstInstanceRange = types.StringValue(jsonData.MstInstanceRange)
	} else {
		v.MstInstanceRange = types.StringNull()
	}

	if jsonData.MulticastGroupSubnet != "" {
		v.MulticastGroupSubnet = types.StringValue(jsonData.MulticastGroupSubnet)
	} else {
		v.MulticastGroupSubnet = types.StringNull()
	}

	if jsonData.MvpnVriIdRange != "" {
		v.MvpnVriIdRange = types.StringValue(jsonData.MvpnVriIdRange)
	} else {
		v.MvpnVriIdRange = types.StringNull()
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

	if jsonData.NetworkVlanRange != "" {
		v.NetworkVlanRange = types.StringValue(jsonData.NetworkVlanRange)
	} else {
		v.NetworkVlanRange = types.StringNull()
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

	if jsonData.NveLbId != nil {
		if jsonData.NveLbId.IsEmpty() {
			v.NveLbId = types.Int64Null()
		} else {
			v.NveLbId = types.Int64Value(int64(*jsonData.NveLbId))
		}
	} else {
		v.NveLbId = types.Int64Null()
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

	if jsonData.ObjectTrackingNumberRange != "" {
		v.ObjectTrackingNumberRange = types.StringValue(jsonData.ObjectTrackingNumberRange)
	} else {
		v.ObjectTrackingNumberRange = types.StringNull()
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

	if jsonData.OverlayMode != "" {
		v.OverlayMode = types.StringValue(jsonData.OverlayMode)
	} else {
		v.OverlayMode = types.StringNull()
	}

	if jsonData.OverlayModePrev != "" {
		v.OverlayModePrev = types.StringValue(jsonData.OverlayModePrev)
	} else {
		v.OverlayModePrev = types.StringNull()
	}

	if jsonData.OverwriteGlobalNxc != "" {
		x, _ := strconv.ParseBool(jsonData.OverwriteGlobalNxc)
		v.OverwriteGlobalNxc = types.BoolValue(x)
	} else {
		v.OverwriteGlobalNxc = types.BoolNull()
	}

	if jsonData.PerVrfLoopbackAutoProvision != "" {
		x, _ := strconv.ParseBool(jsonData.PerVrfLoopbackAutoProvision)
		v.PerVrfLoopbackAutoProvision = types.BoolValue(x)
	} else {
		v.PerVrfLoopbackAutoProvision = types.BoolNull()
	}

	if jsonData.PerVrfLoopbackAutoProvisionPrev != "" {
		x, _ := strconv.ParseBool(jsonData.PerVrfLoopbackAutoProvisionPrev)
		v.PerVrfLoopbackAutoProvisionPrev = types.BoolValue(x)
	} else {
		v.PerVrfLoopbackAutoProvisionPrev = types.BoolNull()
	}

	if jsonData.PerVrfLoopbackAutoProvisionV6 != "" {
		x, _ := strconv.ParseBool(jsonData.PerVrfLoopbackAutoProvisionV6)
		v.PerVrfLoopbackAutoProvisionV6 = types.BoolValue(x)
	} else {
		v.PerVrfLoopbackAutoProvisionV6 = types.BoolNull()
	}

	if jsonData.PerVrfLoopbackAutoProvisionV6Prev != "" {
		x, _ := strconv.ParseBool(jsonData.PerVrfLoopbackAutoProvisionV6Prev)
		v.PerVrfLoopbackAutoProvisionV6Prev = types.BoolValue(x)
	} else {
		v.PerVrfLoopbackAutoProvisionV6Prev = types.BoolNull()
	}

	if jsonData.PerVrfLoopbackIpRange != "" {
		v.PerVrfLoopbackIpRange = types.StringValue(jsonData.PerVrfLoopbackIpRange)
	} else {
		v.PerVrfLoopbackIpRange = types.StringNull()
	}

	if jsonData.PerVrfLoopbackIpRangeV6 != "" {
		v.PerVrfLoopbackIpRangeV6 = types.StringValue(jsonData.PerVrfLoopbackIpRangeV6)
	} else {
		v.PerVrfLoopbackIpRangeV6 = types.StringNull()
	}

	if jsonData.PfcWatchInt != nil {
		if jsonData.PfcWatchInt.IsEmpty() {
			v.PfcWatchInt = types.Int64Null()
		} else {
			v.PfcWatchInt = types.Int64Value(int64(*jsonData.PfcWatchInt))
		}
	} else {
		v.PfcWatchInt = types.Int64Null()
	}

	if jsonData.PfcWatchIntPrev != nil {
		if jsonData.PfcWatchIntPrev.IsEmpty() {
			v.PfcWatchIntPrev = types.Int64Null()
		} else {
			v.PfcWatchIntPrev = types.Int64Value(int64(*jsonData.PfcWatchIntPrev))
		}
	} else {
		v.PfcWatchIntPrev = types.Int64Null()
	}

	if jsonData.PhantomRpLbId1 != nil {
		if jsonData.PhantomRpLbId1.IsEmpty() {
			v.PhantomRpLbId1 = types.Int64Null()
		} else {
			v.PhantomRpLbId1 = types.Int64Value(int64(*jsonData.PhantomRpLbId1))
		}
	} else {
		v.PhantomRpLbId1 = types.Int64Null()
	}

	if jsonData.PhantomRpLbId2 != nil {
		if jsonData.PhantomRpLbId2.IsEmpty() {
			v.PhantomRpLbId2 = types.Int64Null()
		} else {
			v.PhantomRpLbId2 = types.Int64Value(int64(*jsonData.PhantomRpLbId2))
		}
	} else {
		v.PhantomRpLbId2 = types.Int64Null()
	}

	if jsonData.PhantomRpLbId3 != nil {
		if jsonData.PhantomRpLbId3.IsEmpty() {
			v.PhantomRpLbId3 = types.Int64Null()
		} else {
			v.PhantomRpLbId3 = types.Int64Value(int64(*jsonData.PhantomRpLbId3))
		}
	} else {
		v.PhantomRpLbId3 = types.Int64Null()
	}

	if jsonData.PhantomRpLbId4 != nil {
		if jsonData.PhantomRpLbId4.IsEmpty() {
			v.PhantomRpLbId4 = types.Int64Null()
		} else {
			v.PhantomRpLbId4 = types.Int64Value(int64(*jsonData.PhantomRpLbId4))
		}
	} else {
		v.PhantomRpLbId4 = types.Int64Null()
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

	if jsonData.PremsoParentFabric != "" {
		v.PremsoParentFabric = types.StringValue(jsonData.PremsoParentFabric)
	} else {
		v.PremsoParentFabric = types.StringNull()
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

	if jsonData.PtpVlanId != nil {
		if jsonData.PtpVlanId.IsEmpty() {
			v.PtpVlanId = types.Int64Null()
		} else {
			v.PtpVlanId = types.Int64Value(int64(*jsonData.PtpVlanId))
		}
	} else {
		v.PtpVlanId = types.Int64Null()
	}

	if jsonData.QkdProfileName != "" {
		v.QkdProfileName = types.StringValue(jsonData.QkdProfileName)
	} else {
		v.QkdProfileName = types.StringNull()
	}

	if jsonData.QkdProfileNamePrev != "" {
		x, _ := strconv.ParseBool(jsonData.QkdProfileNamePrev)
		v.QkdProfileNamePrev = types.BoolValue(x)
	} else {
		v.QkdProfileNamePrev = types.BoolNull()
	}

	if jsonData.ReplicationMode != "" {
		v.ReplicationMode = types.StringValue(jsonData.ReplicationMode)
	} else {
		v.ReplicationMode = types.StringNull()
	}

	if jsonData.RouterIdRange != "" {
		v.RouterIdRange = types.StringValue(jsonData.RouterIdRange)
	} else {
		v.RouterIdRange = types.StringNull()
	}

	if jsonData.RouteMapSequenceNumberRange != "" {
		v.RouteMapSequenceNumberRange = types.StringValue(jsonData.RouteMapSequenceNumberRange)
	} else {
		v.RouteMapSequenceNumberRange = types.StringNull()
	}

	if jsonData.RpCount != nil {
		if jsonData.RpCount.IsEmpty() {
			v.RpCount = types.Int64Null()
		} else {
			v.RpCount = types.Int64Value(int64(*jsonData.RpCount))
		}
	} else {
		v.RpCount = types.Int64Null()
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

	if jsonData.RpMode != "" {
		v.RpMode = types.StringValue(jsonData.RpMode)
	} else {
		v.RpMode = types.StringNull()
	}

	if jsonData.RrCount != nil {
		if jsonData.RrCount.IsEmpty() {
			v.RrCount = types.Int64Null()
		} else {
			v.RrCount = types.Int64Value(int64(*jsonData.RrCount))
		}
	} else {
		v.RrCount = types.Int64Null()
	}

	if jsonData.SeedSwitchCoreInterfaces != "" {
		v.SeedSwitchCoreInterfaces = types.StringValue(jsonData.SeedSwitchCoreInterfaces)
	} else {
		v.SeedSwitchCoreInterfaces = types.StringNull()
	}

	if jsonData.ServiceNetworkVlanRange != "" {
		v.ServiceNetworkVlanRange = types.StringValue(jsonData.ServiceNetworkVlanRange)
	} else {
		v.ServiceNetworkVlanRange = types.StringNull()
	}

	if jsonData.SgtIdRange != "" {
		v.SgtIdRange = types.StringValue(jsonData.SgtIdRange)
	} else {
		v.SgtIdRange = types.StringNull()
	}

	if jsonData.SgtNamePrefix != "" {
		v.SgtNamePrefix = types.StringValue(jsonData.SgtNamePrefix)
	} else {
		v.SgtNamePrefix = types.StringNull()
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

	if jsonData.SiteId != "" {
		v.SiteId = types.StringValue(jsonData.SiteId)
	} else {
		v.SiteId = types.StringNull()
	}

	if jsonData.SiteIdPolicyId != nil {
		if jsonData.SiteIdPolicyId.IsEmpty() {
			v.SiteIdPolicyId = types.Int64Null()
		} else {
			v.SiteIdPolicyId = types.Int64Value(int64(*jsonData.SiteIdPolicyId))
		}
	} else {
		v.SiteIdPolicyId = types.Int64Null()
	}

	if jsonData.SlaIdRange != "" {
		v.SlaIdRange = types.StringValue(jsonData.SlaIdRange)
	} else {
		v.SlaIdRange = types.StringNull()
	}

	if jsonData.SnmpServerHostTrap != "" {
		x, _ := strconv.ParseBool(jsonData.SnmpServerHostTrap)
		v.SnmpServerHostTrap = types.BoolValue(x)
	} else {
		v.SnmpServerHostTrap = types.BoolNull()
	}

	if jsonData.SpineCount != nil {
		if jsonData.SpineCount.IsEmpty() {
			v.SpineCount = types.Int64Null()
		} else {
			v.SpineCount = types.Int64Value(int64(*jsonData.SpineCount))
		}
	} else {
		v.SpineCount = types.Int64Null()
	}

	if jsonData.SpineSwitchCoreInterfaces != "" {
		v.SpineSwitchCoreInterfaces = types.StringValue(jsonData.SpineSwitchCoreInterfaces)
	} else {
		v.SpineSwitchCoreInterfaces = types.StringNull()
	}

	if jsonData.SspineAddDelDebugFlag != "" {
		v.SspineAddDelDebugFlag = types.StringValue(jsonData.SspineAddDelDebugFlag)
	} else {
		v.SspineAddDelDebugFlag = types.StringNull()
	}

	if jsonData.SspineCount != nil {
		if jsonData.SspineCount.IsEmpty() {
			v.SspineCount = types.Int64Null()
		} else {
			v.SspineCount = types.Int64Value(int64(*jsonData.SspineCount))
		}
	} else {
		v.SspineCount = types.Int64Null()
	}

	if jsonData.StaticUnderlayIpAlloc != "" {
		x, _ := strconv.ParseBool(jsonData.StaticUnderlayIpAlloc)
		v.StaticUnderlayIpAlloc = types.BoolValue(x)
	} else {
		v.StaticUnderlayIpAlloc = types.BoolNull()
	}

	if jsonData.StpBridgePriority != nil {
		if jsonData.StpBridgePriority.IsEmpty() {
			v.StpBridgePriority = types.Int64Null()
		} else {
			v.StpBridgePriority = types.Int64Value(int64(*jsonData.StpBridgePriority))
		}
	} else {
		v.StpBridgePriority = types.Int64Null()
	}

	if jsonData.StpRootOption != "" {
		v.StpRootOption = types.StringValue(jsonData.StpRootOption)
	} else {
		v.StpRootOption = types.StringNull()
	}

	if jsonData.StpVlanRange != "" {
		v.StpVlanRange = types.StringValue(jsonData.StpVlanRange)
	} else {
		v.StpVlanRange = types.StringNull()
	}

	if jsonData.StrictCcMode != "" {
		x, _ := strconv.ParseBool(jsonData.StrictCcMode)
		v.StrictCcMode = types.BoolValue(x)
	} else {
		v.StrictCcMode = types.BoolNull()
	}

	if jsonData.SubinterfaceRange != "" {
		v.SubinterfaceRange = types.StringValue(jsonData.SubinterfaceRange)
	} else {
		v.SubinterfaceRange = types.StringNull()
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

	if jsonData.TcamAllocation != "" {
		x, _ := strconv.ParseBool(jsonData.TcamAllocation)
		v.TcamAllocation = types.BoolValue(x)
	} else {
		v.TcamAllocation = types.BoolNull()
	}

	if jsonData.TopdownConfigRmTracking != "" {
		v.TopdownConfigRmTracking = types.StringValue(jsonData.TopdownConfigRmTracking)
	} else {
		v.TopdownConfigRmTracking = types.StringNull()
	}

	if jsonData.TrustpointLabel != "" {
		v.TrustpointLabel = types.StringValue(jsonData.TrustpointLabel)
	} else {
		v.TrustpointLabel = types.StringNull()
	}

	if jsonData.UnderlayIsV6 != "" {
		x, _ := strconv.ParseBool(jsonData.UnderlayIsV6)
		v.UnderlayIsV6 = types.BoolValue(x)
	} else {
		v.UnderlayIsV6 = types.BoolNull()
	}

	if jsonData.UnderlayIsV6Prev != "" {
		x, _ := strconv.ParseBool(jsonData.UnderlayIsV6Prev)
		v.UnderlayIsV6Prev = types.BoolValue(x)
	} else {
		v.UnderlayIsV6Prev = types.BoolNull()
	}

	if jsonData.UnnumBootstrapLbId != nil {
		if jsonData.UnnumBootstrapLbId.IsEmpty() {
			v.UnnumBootstrapLbId = types.Int64Null()
		} else {
			v.UnnumBootstrapLbId = types.Int64Value(int64(*jsonData.UnnumBootstrapLbId))
		}
	} else {
		v.UnnumBootstrapLbId = types.Int64Null()
	}

	if jsonData.UnnumDhcpEndInternal != "" {
		v.UnnumDhcpEndInternal = types.StringValue(jsonData.UnnumDhcpEndInternal)
	} else {
		v.UnnumDhcpEndInternal = types.StringNull()
	}

	if jsonData.UnnumDhcpStartInternal != "" {
		v.UnnumDhcpStartInternal = types.StringValue(jsonData.UnnumDhcpStartInternal)
	} else {
		v.UnnumDhcpStartInternal = types.StringNull()
	}

	if jsonData.UpgradeFromVersion != "" {
		v.UpgradeFromVersion = types.StringValue(jsonData.UpgradeFromVersion)
	} else {
		v.UpgradeFromVersion = types.StringNull()
	}

	if jsonData.UseLinkLocal != "" {
		x, _ := strconv.ParseBool(jsonData.UseLinkLocal)
		v.UseLinkLocal = types.BoolValue(x)
	} else {
		v.UseLinkLocal = types.BoolNull()
	}

	if jsonData.V6SubnetRange != "" {
		v.V6SubnetRange = types.StringValue(jsonData.V6SubnetRange)
	} else {
		v.V6SubnetRange = types.StringNull()
	}

	if jsonData.V6SubnetTargetMask != nil {
		if jsonData.V6SubnetTargetMask.IsEmpty() {
			v.V6SubnetTargetMask = types.Int64Null()
		} else {
			v.V6SubnetTargetMask = types.Int64Value(int64(*jsonData.V6SubnetTargetMask))
		}
	} else {
		v.V6SubnetTargetMask = types.Int64Null()
	}

	if jsonData.VpcAutoRecoveryTime != nil {
		if jsonData.VpcAutoRecoveryTime.IsEmpty() {
			v.VpcAutoRecoveryTime = types.Int64Null()
		} else {
			v.VpcAutoRecoveryTime = types.Int64Value(int64(*jsonData.VpcAutoRecoveryTime))
		}
	} else {
		v.VpcAutoRecoveryTime = types.Int64Null()
	}

	if jsonData.VpcDelayRestore != nil {
		if jsonData.VpcDelayRestore.IsEmpty() {
			v.VpcDelayRestore = types.Int64Null()
		} else {
			v.VpcDelayRestore = types.Int64Value(int64(*jsonData.VpcDelayRestore))
		}
	} else {
		v.VpcDelayRestore = types.Int64Null()
	}

	if jsonData.VpcDelayRestoreTime != nil {
		if jsonData.VpcDelayRestoreTime.IsEmpty() {
			v.VpcDelayRestoreTime = types.Int64Null()
		} else {
			v.VpcDelayRestoreTime = types.Int64Value(int64(*jsonData.VpcDelayRestoreTime))
		}
	} else {
		v.VpcDelayRestoreTime = types.Int64Null()
	}

	if jsonData.VpcDomainIdRange != "" {
		v.VpcDomainIdRange = types.StringValue(jsonData.VpcDomainIdRange)
	} else {
		v.VpcDomainIdRange = types.StringNull()
	}

	if jsonData.VpcEnableIpv6NdSync != "" {
		x, _ := strconv.ParseBool(jsonData.VpcEnableIpv6NdSync)
		v.VpcEnableIpv6NdSync = types.BoolValue(x)
	} else {
		v.VpcEnableIpv6NdSync = types.BoolNull()
	}

	if jsonData.VpcPeerKeepAliveOption != "" {
		v.VpcPeerKeepAliveOption = types.StringValue(jsonData.VpcPeerKeepAliveOption)
	} else {
		v.VpcPeerKeepAliveOption = types.StringNull()
	}

	if jsonData.VpcPeerLinkPo != nil {
		if jsonData.VpcPeerLinkPo.IsEmpty() {
			v.VpcPeerLinkPo = types.Int64Null()
		} else {
			v.VpcPeerLinkPo = types.Int64Value(int64(*jsonData.VpcPeerLinkPo))
		}
	} else {
		v.VpcPeerLinkPo = types.Int64Null()
	}

	if jsonData.VpcPeerLinkVlan != nil {
		if jsonData.VpcPeerLinkVlan.IsEmpty() {
			v.VpcPeerLinkVlan = types.Int64Null()
		} else {
			v.VpcPeerLinkVlan = types.Int64Value(int64(*jsonData.VpcPeerLinkVlan))
		}
	} else {
		v.VpcPeerLinkVlan = types.Int64Null()
	}

	if jsonData.VrfLiteAutoconfig != "" {
		v.VrfLiteAutoconfig = types.StringValue(jsonData.VrfLiteAutoconfig)
	} else {
		v.VrfLiteAutoconfig = types.StringNull()
	}

	if jsonData.VrfVlanRange != "" {
		v.VrfVlanRange = types.StringValue(jsonData.VrfVlanRange)
	} else {
		v.VrfVlanRange = types.StringNull()
	}

	if jsonData.AbstractAnycastRp != "" {
		v.AbstractAnycastRp = types.StringValue(jsonData.AbstractAnycastRp)
	} else {
		v.AbstractAnycastRp = types.StringNull()
	}

	if jsonData.AbstractBgp != "" {
		v.AbstractBgp = types.StringValue(jsonData.AbstractBgp)
	} else {
		v.AbstractBgp = types.StringNull()
	}

	if jsonData.AbstractBgpNeighbor != "" {
		v.AbstractBgpNeighbor = types.StringValue(jsonData.AbstractBgpNeighbor)
	} else {
		v.AbstractBgpNeighbor = types.StringNull()
	}

	if jsonData.AbstractBgpRr != "" {
		v.AbstractBgpRr = types.StringValue(jsonData.AbstractBgpRr)
	} else {
		v.AbstractBgpRr = types.StringNull()
	}

	if jsonData.AbstractDhcp != "" {
		v.AbstractDhcp = types.StringValue(jsonData.AbstractDhcp)
	} else {
		v.AbstractDhcp = types.StringNull()
	}

	if jsonData.AbstractExtraConfigBootstrap != "" {
		v.AbstractExtraConfigBootstrap = types.StringValue(jsonData.AbstractExtraConfigBootstrap)
	} else {
		v.AbstractExtraConfigBootstrap = types.StringNull()
	}

	if jsonData.AbstractExtraConfigLeaf != "" {
		v.AbstractExtraConfigLeaf = types.StringValue(jsonData.AbstractExtraConfigLeaf)
	} else {
		v.AbstractExtraConfigLeaf = types.StringNull()
	}

	if jsonData.AbstractExtraConfigSpine != "" {
		v.AbstractExtraConfigSpine = types.StringValue(jsonData.AbstractExtraConfigSpine)
	} else {
		v.AbstractExtraConfigSpine = types.StringNull()
	}

	if jsonData.AbstractExtraConfigTor != "" {
		v.AbstractExtraConfigTor = types.StringValue(jsonData.AbstractExtraConfigTor)
	} else {
		v.AbstractExtraConfigTor = types.StringNull()
	}

	if jsonData.AbstractFeatureLeaf != "" {
		v.AbstractFeatureLeaf = types.StringValue(jsonData.AbstractFeatureLeaf)
	} else {
		v.AbstractFeatureLeaf = types.StringNull()
	}

	if jsonData.AbstractFeatureSpine != "" {
		v.AbstractFeatureSpine = types.StringValue(jsonData.AbstractFeatureSpine)
	} else {
		v.AbstractFeatureSpine = types.StringNull()
	}

	if jsonData.AbstractIsis != "" {
		v.AbstractIsis = types.StringValue(jsonData.AbstractIsis)
	} else {
		v.AbstractIsis = types.StringNull()
	}

	if jsonData.AbstractIsisInterface != "" {
		v.AbstractIsisInterface = types.StringValue(jsonData.AbstractIsisInterface)
	} else {
		v.AbstractIsisInterface = types.StringNull()
	}

	if jsonData.AbstractLoopbackInterface != "" {
		v.AbstractLoopbackInterface = types.StringValue(jsonData.AbstractLoopbackInterface)
	} else {
		v.AbstractLoopbackInterface = types.StringNull()
	}

	if jsonData.AbstractMulticast != "" {
		v.AbstractMulticast = types.StringValue(jsonData.AbstractMulticast)
	} else {
		v.AbstractMulticast = types.StringNull()
	}

	if jsonData.AbstractOspf != "" {
		v.AbstractOspf = types.StringValue(jsonData.AbstractOspf)
	} else {
		v.AbstractOspf = types.StringNull()
	}

	if jsonData.AbstractOspfInterface != "" {
		v.AbstractOspfInterface = types.StringValue(jsonData.AbstractOspfInterface)
	} else {
		v.AbstractOspfInterface = types.StringNull()
	}

	if jsonData.AbstractPimInterface != "" {
		v.AbstractPimInterface = types.StringValue(jsonData.AbstractPimInterface)
	} else {
		v.AbstractPimInterface = types.StringNull()
	}

	if jsonData.AbstractRouteMap != "" {
		v.AbstractRouteMap = types.StringValue(jsonData.AbstractRouteMap)
	} else {
		v.AbstractRouteMap = types.StringNull()
	}

	if jsonData.AbstractRoutedHost != "" {
		v.AbstractRoutedHost = types.StringValue(jsonData.AbstractRoutedHost)
	} else {
		v.AbstractRoutedHost = types.StringNull()
	}

	if jsonData.AbstractVlanInterface != "" {
		v.AbstractVlanInterface = types.StringValue(jsonData.AbstractVlanInterface)
	} else {
		v.AbstractVlanInterface = types.StringNull()
	}

	if jsonData.AbstractVpcDomain != "" {
		v.AbstractVpcDomain = types.StringValue(jsonData.AbstractVpcDomain)
	} else {
		v.AbstractVpcDomain = types.StringNull()
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

	v.Deploy = types.BoolValue(jsonData.Deploy)
	if jsonData.DeploymentStatus != "" {
		v.DeploymentStatus = types.StringValue(jsonData.DeploymentStatus)
	} else {
		v.DeploymentStatus = types.StringNull()
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

	if jsonData.TempAnycastGateway != "" {
		v.TempAnycastGateway = types.StringValue(jsonData.TempAnycastGateway)
	} else {
		v.TempAnycastGateway = types.StringNull()
	}

	if jsonData.TempVpcDomainMgmt != "" {
		v.TempVpcDomainMgmt = types.StringValue(jsonData.TempVpcDomainMgmt)
	} else {
		v.TempVpcDomainMgmt = types.StringNull()
	}

	if jsonData.TempVpcPeerLink != "" {
		v.TempVpcPeerLink = types.StringValue(jsonData.TempVpcPeerLink)
	} else {
		v.TempVpcPeerLink = types.StringNull()
	}

	if jsonData.VrfExtensionTemplate != "" {
		v.VrfExtensionTemplate = types.StringValue(jsonData.VrfExtensionTemplate)
	} else {
		v.VrfExtensionTemplate = types.StringNull()
	}

	return err
}

func (v FabricVxlanEvpnModel) GetModelData() *resource_fabric_common.NDFCFabricCommonModel {
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

	if !v.ActiveMigration.IsNull() && !v.ActiveMigration.IsUnknown() {
		data.ActiveMigration = strconv.FormatBool(v.ActiveMigration.ValueBool())
	} else {
		data.ActiveMigration = ""
	}

	if !v.AdvertisePipBgp.IsNull() && !v.AdvertisePipBgp.IsUnknown() {
		data.AdvertisePipBgp = strconv.FormatBool(v.AdvertisePipBgp.ValueBool())
	} else {
		data.AdvertisePipBgp = ""
	}

	if !v.AdvertisePipOnBorder.IsNull() && !v.AdvertisePipOnBorder.IsUnknown() {
		data.AdvertisePipOnBorder = strconv.FormatBool(v.AdvertisePipOnBorder.ValueBool())
	} else {
		data.AdvertisePipOnBorder = ""
	}

	if !v.AgentIntf.IsNull() && !v.AgentIntf.IsUnknown() {
		data.AgentIntf = v.AgentIntf.ValueString()
	} else {
		data.AgentIntf = ""
	}

	if !v.AggAccVpcPoIdRange.IsNull() && !v.AggAccVpcPoIdRange.IsUnknown() {
		data.AggAccVpcPoIdRange = v.AggAccVpcPoIdRange.ValueString()
	} else {
		data.AggAccVpcPoIdRange = ""
	}

	if !v.AiMlQosPolicy.IsNull() && !v.AiMlQosPolicy.IsUnknown() {
		data.AiMlQosPolicy = v.AiMlQosPolicy.ValueString()
	} else {
		data.AiMlQosPolicy = ""
	}

	if !v.AllowL3vniNoVlan.IsNull() && !v.AllowL3vniNoVlan.IsUnknown() {
		data.AllowL3vniNoVlan = strconv.FormatBool(v.AllowL3vniNoVlan.ValueBool())
	} else {
		data.AllowL3vniNoVlan = ""
	}

	if !v.AllowNxc.IsNull() && !v.AllowNxc.IsUnknown() {
		data.AllowNxc = strconv.FormatBool(v.AllowNxc.ValueBool())
	} else {
		data.AllowNxc = ""
	}

	if !v.AnycastBgwAdvertisePip.IsNull() && !v.AnycastBgwAdvertisePip.IsUnknown() {
		data.AnycastBgwAdvertisePip = strconv.FormatBool(v.AnycastBgwAdvertisePip.ValueBool())
	} else {
		data.AnycastBgwAdvertisePip = ""
	}

	if !v.AnycastGwMac.IsNull() && !v.AnycastGwMac.IsUnknown() {
		data.AnycastGwMac = v.AnycastGwMac.ValueString()
	} else {
		data.AnycastGwMac = ""
	}

	if !v.AnycastLbId.IsNull() && !v.AnycastLbId.IsUnknown() {
		data.AnycastLbId = new(Int64Custom)
		*data.AnycastLbId = Int64Custom(v.AnycastLbId.ValueInt64())
	} else {
		data.AnycastLbId = nil
	}

	if !v.AnycastRpIpRange.IsNull() && !v.AnycastRpIpRange.IsUnknown() {
		data.AnycastRpIpRange = v.AnycastRpIpRange.ValueString()
	} else {
		data.AnycastRpIpRange = ""
	}

	if !v.AutoSymmetricDefaultVrf.IsNull() && !v.AutoSymmetricDefaultVrf.IsUnknown() {
		data.AutoSymmetricDefaultVrf = strconv.FormatBool(v.AutoSymmetricDefaultVrf.ValueBool())
	} else {
		data.AutoSymmetricDefaultVrf = ""
	}

	if !v.AutoSymmetricVrfLite.IsNull() && !v.AutoSymmetricVrfLite.IsUnknown() {
		data.AutoSymmetricVrfLite = strconv.FormatBool(v.AutoSymmetricVrfLite.ValueBool())
	} else {
		data.AutoSymmetricVrfLite = ""
	}

	if !v.AutoUniqueVrfLiteIpPrefix.IsNull() && !v.AutoUniqueVrfLiteIpPrefix.IsUnknown() {
		data.AutoUniqueVrfLiteIpPrefix = strconv.FormatBool(v.AutoUniqueVrfLiteIpPrefix.ValueBool())
	} else {
		data.AutoUniqueVrfLiteIpPrefix = ""
	}

	if !v.AutoVrfliteIfcDefaultVrf.IsNull() && !v.AutoVrfliteIfcDefaultVrf.IsUnknown() {
		data.AutoVrfliteIfcDefaultVrf = strconv.FormatBool(v.AutoVrfliteIfcDefaultVrf.ValueBool())
	} else {
		data.AutoVrfliteIfcDefaultVrf = ""
	}

	if !v.Banner.IsNull() && !v.Banner.IsUnknown() {
		data.Banner = v.Banner.ValueString()
	} else {
		data.Banner = ""
	}

	if !v.BfdAuthEnable.IsNull() && !v.BfdAuthEnable.IsUnknown() {
		data.BfdAuthEnable = strconv.FormatBool(v.BfdAuthEnable.ValueBool())
	} else {
		data.BfdAuthEnable = ""
	}

	if !v.BfdAuthKey.IsNull() && !v.BfdAuthKey.IsUnknown() {
		data.BfdAuthKey = v.BfdAuthKey.ValueString()
	} else {
		data.BfdAuthKey = ""
	}

	if !v.BfdAuthKeyId.IsNull() && !v.BfdAuthKeyId.IsUnknown() {
		data.BfdAuthKeyId = new(Int64Custom)
		*data.BfdAuthKeyId = Int64Custom(v.BfdAuthKeyId.ValueInt64())
	} else {
		data.BfdAuthKeyId = nil
	}

	if !v.BfdEnable.IsNull() && !v.BfdEnable.IsUnknown() {
		data.BfdEnable = strconv.FormatBool(v.BfdEnable.ValueBool())
	} else {
		data.BfdEnable = ""
	}

	if !v.BfdIbgpEnable.IsNull() && !v.BfdIbgpEnable.IsUnknown() {
		data.BfdIbgpEnable = strconv.FormatBool(v.BfdIbgpEnable.ValueBool())
	} else {
		data.BfdIbgpEnable = ""
	}

	if !v.BfdIsisEnable.IsNull() && !v.BfdIsisEnable.IsUnknown() {
		data.BfdIsisEnable = strconv.FormatBool(v.BfdIsisEnable.ValueBool())
	} else {
		data.BfdIsisEnable = ""
	}

	if !v.BfdOspfEnable.IsNull() && !v.BfdOspfEnable.IsUnknown() {
		data.BfdOspfEnable = strconv.FormatBool(v.BfdOspfEnable.ValueBool())
	} else {
		data.BfdOspfEnable = ""
	}

	if !v.BfdPimEnable.IsNull() && !v.BfdPimEnable.IsUnknown() {
		data.BfdPimEnable = strconv.FormatBool(v.BfdPimEnable.ValueBool())
	} else {
		data.BfdPimEnable = ""
	}

	if !v.BgpAs.IsNull() && !v.BgpAs.IsUnknown() {
		data.BgpAs = v.BgpAs.ValueString()
	} else {
		data.BgpAs = ""
	}

	if !v.BgpAuthEnable.IsNull() && !v.BgpAuthEnable.IsUnknown() {
		data.BgpAuthEnable = strconv.FormatBool(v.BgpAuthEnable.ValueBool())
	} else {
		data.BgpAuthEnable = ""
	}

	if !v.BgpAuthKey.IsNull() && !v.BgpAuthKey.IsUnknown() {
		data.BgpAuthKey = v.BgpAuthKey.ValueString()
	} else {
		data.BgpAuthKey = ""
	}

	if !v.BgpAuthKeyType.IsNull() && !v.BgpAuthKeyType.IsUnknown() {
		data.BgpAuthKeyType = new(Int64Custom)
		*data.BgpAuthKeyType = Int64Custom(v.BgpAuthKeyType.ValueInt64())
	} else {
		data.BgpAuthKeyType = nil
	}

	if !v.BgpLbId.IsNull() && !v.BgpLbId.IsUnknown() {
		data.BgpLbId = new(Int64Custom)
		*data.BgpLbId = Int64Custom(v.BgpLbId.ValueInt64())
	} else {
		data.BgpLbId = nil
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

	if !v.BrfieldDebugFlag.IsNull() && !v.BrfieldDebugFlag.IsUnknown() {
		data.BrfieldDebugFlag = v.BrfieldDebugFlag.ValueString()
	} else {
		data.BrfieldDebugFlag = ""
	}

	if !v.BrownfieldNetworkNameFormat.IsNull() && !v.BrownfieldNetworkNameFormat.IsUnknown() {
		data.BrownfieldNetworkNameFormat = v.BrownfieldNetworkNameFormat.ValueString()
	} else {
		data.BrownfieldNetworkNameFormat = ""
	}

	if !v.BrownfieldSkipOverlayNetworkAttachments.IsNull() && !v.BrownfieldSkipOverlayNetworkAttachments.IsUnknown() {
		data.BrownfieldSkipOverlayNetworkAttachments = strconv.FormatBool(v.BrownfieldSkipOverlayNetworkAttachments.ValueBool())
	} else {
		data.BrownfieldSkipOverlayNetworkAttachments = ""
	}

	if !v.CdpEnable.IsNull() && !v.CdpEnable.IsUnknown() {
		data.CdpEnable = strconv.FormatBool(v.CdpEnable.ValueBool())
	} else {
		data.CdpEnable = ""
	}

	if !v.CoppPolicy.IsNull() && !v.CoppPolicy.IsUnknown() {
		data.CoppPolicy = v.CoppPolicy.ValueString()
	} else {
		data.CoppPolicy = ""
	}

	if !v.DciMacsecAlgorithm.IsNull() && !v.DciMacsecAlgorithm.IsUnknown() {
		data.DciMacsecAlgorithm = v.DciMacsecAlgorithm.ValueString()
	} else {
		data.DciMacsecAlgorithm = ""
	}

	if !v.DciMacsecCipherSuite.IsNull() && !v.DciMacsecCipherSuite.IsUnknown() {
		data.DciMacsecCipherSuite = v.DciMacsecCipherSuite.ValueString()
	} else {
		data.DciMacsecCipherSuite = ""
	}

	if !v.DciMacsecFallbackAlgorithm.IsNull() && !v.DciMacsecFallbackAlgorithm.IsUnknown() {
		data.DciMacsecFallbackAlgorithm = v.DciMacsecFallbackAlgorithm.ValueString()
	} else {
		data.DciMacsecFallbackAlgorithm = ""
	}

	if !v.DciMacsecFallbackKeyString.IsNull() && !v.DciMacsecFallbackKeyString.IsUnknown() {
		data.DciMacsecFallbackKeyString = v.DciMacsecFallbackKeyString.ValueString()
	} else {
		data.DciMacsecFallbackKeyString = ""
	}

	if !v.DciMacsecKeyString.IsNull() && !v.DciMacsecKeyString.IsUnknown() {
		data.DciMacsecKeyString = v.DciMacsecKeyString.ValueString()
	} else {
		data.DciMacsecKeyString = ""
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

	if !v.DefaultQueuingPolicyCloudscale.IsNull() && !v.DefaultQueuingPolicyCloudscale.IsUnknown() {
		data.DefaultQueuingPolicyCloudscale = v.DefaultQueuingPolicyCloudscale.ValueString()
	} else {
		data.DefaultQueuingPolicyCloudscale = ""
	}

	if !v.DefaultQueuingPolicyOther.IsNull() && !v.DefaultQueuingPolicyOther.IsUnknown() {
		data.DefaultQueuingPolicyOther = v.DefaultQueuingPolicyOther.ValueString()
	} else {
		data.DefaultQueuingPolicyOther = ""
	}

	if !v.DefaultQueuingPolicyRSeries.IsNull() && !v.DefaultQueuingPolicyRSeries.IsUnknown() {
		data.DefaultQueuingPolicyRSeries = v.DefaultQueuingPolicyRSeries.ValueString()
	} else {
		data.DefaultQueuingPolicyRSeries = ""
	}

	if !v.DefaultVrfRedisBgpRmap.IsNull() && !v.DefaultVrfRedisBgpRmap.IsUnknown() {
		data.DefaultVrfRedisBgpRmap = v.DefaultVrfRedisBgpRmap.ValueString()
	} else {
		data.DefaultVrfRedisBgpRmap = ""
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

	if !v.EnableAgent.IsNull() && !v.EnableAgent.IsUnknown() {
		data.EnableAgent = strconv.FormatBool(v.EnableAgent.ValueBool())
	} else {
		data.EnableAgent = ""
	}

	if !v.EnableAggAccIdRange.IsNull() && !v.EnableAggAccIdRange.IsUnknown() {
		data.EnableAggAccIdRange = strconv.FormatBool(v.EnableAggAccIdRange.ValueBool())
	} else {
		data.EnableAggAccIdRange = ""
	}

	if !v.EnableAiMlQosPolicy.IsNull() && !v.EnableAiMlQosPolicy.IsUnknown() {
		data.EnableAiMlQosPolicy = strconv.FormatBool(v.EnableAiMlQosPolicy.ValueBool())
	} else {
		data.EnableAiMlQosPolicy = ""
	}

	if !v.EnableDciMacsec.IsNull() && !v.EnableDciMacsec.IsUnknown() {
		data.EnableDciMacsec = strconv.FormatBool(v.EnableDciMacsec.ValueBool())
	} else {
		data.EnableDciMacsec = ""
	}

	if !v.EnableDefaultQueuingPolicy.IsNull() && !v.EnableDefaultQueuingPolicy.IsUnknown() {
		data.EnableDefaultQueuingPolicy = strconv.FormatBool(v.EnableDefaultQueuingPolicy.ValueBool())
	} else {
		data.EnableDefaultQueuingPolicy = ""
	}

	if !v.EnableFabricVpcDomainId.IsNull() && !v.EnableFabricVpcDomainId.IsUnknown() {
		data.EnableFabricVpcDomainId = strconv.FormatBool(v.EnableFabricVpcDomainId.ValueBool())
	} else {
		data.EnableFabricVpcDomainId = ""
	}

	if !v.EnableL3vniNoVlan.IsNull() && !v.EnableL3vniNoVlan.IsUnknown() {
		data.EnableL3vniNoVlan = strconv.FormatBool(v.EnableL3vniNoVlan.ValueBool())
	} else {
		data.EnableL3vniNoVlan = ""
	}

	if !v.EnableMacsec.IsNull() && !v.EnableMacsec.IsUnknown() {
		data.EnableMacsec = strconv.FormatBool(v.EnableMacsec.ValueBool())
	} else {
		data.EnableMacsec = ""
	}

	if !v.EnableNetflow.IsNull() && !v.EnableNetflow.IsUnknown() {
		data.EnableNetflow = strconv.FormatBool(v.EnableNetflow.ValueBool())
	} else {
		data.EnableNetflow = ""
	}

	if !v.EnableNgoam.IsNull() && !v.EnableNgoam.IsUnknown() {
		data.EnableNgoam = strconv.FormatBool(v.EnableNgoam.ValueBool())
	} else {
		data.EnableNgoam = ""
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

	if !v.EnablePbr.IsNull() && !v.EnablePbr.IsUnknown() {
		data.EnablePbr = strconv.FormatBool(v.EnablePbr.ValueBool())
	} else {
		data.EnablePbr = ""
	}

	if !v.EnablePvlan.IsNull() && !v.EnablePvlan.IsUnknown() {
		data.EnablePvlan = strconv.FormatBool(v.EnablePvlan.ValueBool())
	} else {
		data.EnablePvlan = ""
	}

	if !v.EnableQkd.IsNull() && !v.EnableQkd.IsUnknown() {
		data.EnableQkd = strconv.FormatBool(v.EnableQkd.ValueBool())
	} else {
		data.EnableQkd = ""
	}

	if !v.EnableRtIntfStats.IsNull() && !v.EnableRtIntfStats.IsUnknown() {
		data.EnableRtIntfStats = strconv.FormatBool(v.EnableRtIntfStats.ValueBool())
	} else {
		data.EnableRtIntfStats = ""
	}

	if !v.EnableSgt.IsNull() && !v.EnableSgt.IsUnknown() {
		data.EnableSgt = strconv.FormatBool(v.EnableSgt.ValueBool())
	} else {
		data.EnableSgt = ""
	}

	if !v.EnableTenantDhcp.IsNull() && !v.EnableTenantDhcp.IsUnknown() {
		data.EnableTenantDhcp = strconv.FormatBool(v.EnableTenantDhcp.ValueBool())
	} else {
		data.EnableTenantDhcp = ""
	}

	if !v.EnableTrm.IsNull() && !v.EnableTrm.IsUnknown() {
		data.EnableTrm = strconv.FormatBool(v.EnableTrm.ValueBool())
	} else {
		data.EnableTrm = ""
	}

	if !v.EnableTrmv6.IsNull() && !v.EnableTrmv6.IsUnknown() {
		data.EnableTrmv6 = strconv.FormatBool(v.EnableTrmv6.ValueBool())
	} else {
		data.EnableTrmv6 = ""
	}

	if !v.EnableVpcPeerLinkNativeVlan.IsNull() && !v.EnableVpcPeerLinkNativeVlan.IsUnknown() {
		data.EnableVpcPeerLinkNativeVlan = strconv.FormatBool(v.EnableVpcPeerLinkNativeVlan.ValueBool())
	} else {
		data.EnableVpcPeerLinkNativeVlan = ""
	}

	if !v.EnableVriIdRealloc.IsNull() && !v.EnableVriIdRealloc.IsUnknown() {
		data.EnableVriIdRealloc = strconv.FormatBool(v.EnableVriIdRealloc.ValueBool())
	} else {
		data.EnableVriIdRealloc = ""
	}

	if !v.EsrOption.IsNull() && !v.EsrOption.IsUnknown() {
		data.EsrOption = v.EsrOption.ValueString()
	} else {
		data.EsrOption = ""
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

	if !v.ExtraConfTor.IsNull() && !v.ExtraConfTor.IsUnknown() {
		data.ExtraConfTor = v.ExtraConfTor.ValueString()
	} else {
		data.ExtraConfTor = ""
	}

	if !v.ExtFabricType.IsNull() && !v.ExtFabricType.IsUnknown() {
		data.ExtFabricType = v.ExtFabricType.ValueString()
	} else {
		data.ExtFabricType = ""
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

	if !v.FabricVpcDomainId.IsNull() && !v.FabricVpcDomainId.IsUnknown() {
		data.FabricVpcDomainId = new(Int64Custom)
		*data.FabricVpcDomainId = Int64Custom(v.FabricVpcDomainId.ValueInt64())
	} else {
		data.FabricVpcDomainId = nil
	}

	if !v.FabricVpcQos.IsNull() && !v.FabricVpcQos.IsUnknown() {
		data.FabricVpcQos = strconv.FormatBool(v.FabricVpcQos.ValueBool())
	} else {
		data.FabricVpcQos = ""
	}

	if !v.FabricVpcQosPolicyName.IsNull() && !v.FabricVpcQosPolicyName.IsUnknown() {
		data.FabricVpcQosPolicyName = v.FabricVpcQosPolicyName.ValueString()
	} else {
		data.FabricVpcQosPolicyName = ""
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

	if !v.GrfieldDebugFlag.IsNull() && !v.GrfieldDebugFlag.IsUnknown() {
		data.GrfieldDebugFlag = v.GrfieldDebugFlag.ValueString()
	} else {
		data.GrfieldDebugFlag = ""
	}

	if !v.HdTime.IsNull() && !v.HdTime.IsUnknown() {
		data.HdTime = new(Int64Custom)
		*data.HdTime = Int64Custom(v.HdTime.ValueInt64())
	} else {
		data.HdTime = nil
	}

	if !v.HostIntfAdminState.IsNull() && !v.HostIntfAdminState.IsUnknown() {
		data.HostIntfAdminState = strconv.FormatBool(v.HostIntfAdminState.ValueBool())
	} else {
		data.HostIntfAdminState = ""
	}

	if !v.IbgpPeerTemplate.IsNull() && !v.IbgpPeerTemplate.IsUnknown() {
		data.IbgpPeerTemplate = v.IbgpPeerTemplate.ValueString()
	} else {
		data.IbgpPeerTemplate = ""
	}

	if !v.IbgpPeerTemplateLeaf.IsNull() && !v.IbgpPeerTemplateLeaf.IsUnknown() {
		data.IbgpPeerTemplateLeaf = v.IbgpPeerTemplateLeaf.ValueString()
	} else {
		data.IbgpPeerTemplateLeaf = ""
	}

	if !v.IgnoreCert.IsNull() && !v.IgnoreCert.IsUnknown() {
		data.IgnoreCert = strconv.FormatBool(v.IgnoreCert.ValueBool())
	} else {
		data.IgnoreCert = ""
	}

	if !v.InbandDhcpServers.IsNull() && !v.InbandDhcpServers.IsUnknown() {
		data.InbandDhcpServers = v.InbandDhcpServers.ValueString()
	} else {
		data.InbandDhcpServers = ""
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

	if !v.Ipv6AnycastRpIpRange.IsNull() && !v.Ipv6AnycastRpIpRange.IsUnknown() {
		data.Ipv6AnycastRpIpRange = v.Ipv6AnycastRpIpRange.ValueString()
	} else {
		data.Ipv6AnycastRpIpRange = ""
	}

	if !v.Ipv6MulticastGroupSubnet.IsNull() && !v.Ipv6MulticastGroupSubnet.IsUnknown() {
		data.Ipv6MulticastGroupSubnet = v.Ipv6MulticastGroupSubnet.ValueString()
	} else {
		data.Ipv6MulticastGroupSubnet = ""
	}

	if !v.IsisAreaNum.IsNull() && !v.IsisAreaNum.IsUnknown() {
		data.IsisAreaNum = v.IsisAreaNum.ValueString()
	} else {
		data.IsisAreaNum = ""
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

	if !v.IsisOverloadElapseTime.IsNull() && !v.IsisOverloadElapseTime.IsUnknown() {
		data.IsisOverloadElapseTime = new(Int64Custom)
		*data.IsisOverloadElapseTime = Int64Custom(v.IsisOverloadElapseTime.ValueInt64())
	} else {
		data.IsisOverloadElapseTime = nil
	}

	if !v.IsisOverloadEnable.IsNull() && !v.IsisOverloadEnable.IsUnknown() {
		data.IsisOverloadEnable = strconv.FormatBool(v.IsisOverloadEnable.ValueBool())
	} else {
		data.IsisOverloadEnable = ""
	}

	if !v.IsisP2pEnable.IsNull() && !v.IsisP2pEnable.IsUnknown() {
		data.IsisP2pEnable = strconv.FormatBool(v.IsisP2pEnable.ValueBool())
	} else {
		data.IsisP2pEnable = ""
	}

	if !v.KmeServerIp.IsNull() && !v.KmeServerIp.IsUnknown() {
		data.KmeServerIp = v.KmeServerIp.ValueString()
	} else {
		data.KmeServerIp = ""
	}

	if !v.KmeServerPort.IsNull() && !v.KmeServerPort.IsUnknown() {
		data.KmeServerPort = new(Int64Custom)
		*data.KmeServerPort = Int64Custom(v.KmeServerPort.ValueInt64())
	} else {
		data.KmeServerPort = nil
	}

	if !v.L2HostIntfMtu.IsNull() && !v.L2HostIntfMtu.IsUnknown() {
		data.L2HostIntfMtu = new(Int64Custom)
		*data.L2HostIntfMtu = Int64Custom(v.L2HostIntfMtu.ValueInt64())
	} else {
		data.L2HostIntfMtu = nil
	}

	if !v.L2SegmentIdRange.IsNull() && !v.L2SegmentIdRange.IsUnknown() {
		data.L2SegmentIdRange = v.L2SegmentIdRange.ValueString()
	} else {
		data.L2SegmentIdRange = ""
	}

	if !v.L3vniIpv6McastGroup.IsNull() && !v.L3vniIpv6McastGroup.IsUnknown() {
		data.L3vniIpv6McastGroup = v.L3vniIpv6McastGroup.ValueString()
	} else {
		data.L3vniIpv6McastGroup = ""
	}

	if !v.L3vniMcastGroup.IsNull() && !v.L3vniMcastGroup.IsUnknown() {
		data.L3vniMcastGroup = v.L3vniMcastGroup.ValueString()
	} else {
		data.L3vniMcastGroup = ""
	}

	if !v.L3PartitionIdRange.IsNull() && !v.L3PartitionIdRange.IsUnknown() {
		data.L3PartitionIdRange = v.L3PartitionIdRange.ValueString()
	} else {
		data.L3PartitionIdRange = ""
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

	if !v.Loopback0Ipv6Range.IsNull() && !v.Loopback0Ipv6Range.IsUnknown() {
		data.Loopback0Ipv6Range = v.Loopback0Ipv6Range.ValueString()
	} else {
		data.Loopback0Ipv6Range = ""
	}

	if !v.Loopback0IpRange.IsNull() && !v.Loopback0IpRange.IsUnknown() {
		data.Loopback0IpRange = v.Loopback0IpRange.ValueString()
	} else {
		data.Loopback0IpRange = ""
	}

	if !v.Loopback1Ipv6Range.IsNull() && !v.Loopback1Ipv6Range.IsUnknown() {
		data.Loopback1Ipv6Range = v.Loopback1Ipv6Range.ValueString()
	} else {
		data.Loopback1Ipv6Range = ""
	}

	if !v.Loopback1IpRange.IsNull() && !v.Loopback1IpRange.IsUnknown() {
		data.Loopback1IpRange = v.Loopback1IpRange.ValueString()
	} else {
		data.Loopback1IpRange = ""
	}

	if !v.MacsecAlgorithm.IsNull() && !v.MacsecAlgorithm.IsUnknown() {
		data.MacsecAlgorithm = v.MacsecAlgorithm.ValueString()
	} else {
		data.MacsecAlgorithm = ""
	}

	if !v.MacsecCipherSuite.IsNull() && !v.MacsecCipherSuite.IsUnknown() {
		data.MacsecCipherSuite = v.MacsecCipherSuite.ValueString()
	} else {
		data.MacsecCipherSuite = ""
	}

	if !v.MacsecFallbackAlgorithm.IsNull() && !v.MacsecFallbackAlgorithm.IsUnknown() {
		data.MacsecFallbackAlgorithm = v.MacsecFallbackAlgorithm.ValueString()
	} else {
		data.MacsecFallbackAlgorithm = ""
	}

	if !v.MacsecFallbackKeyString.IsNull() && !v.MacsecFallbackKeyString.IsUnknown() {
		data.MacsecFallbackKeyString = v.MacsecFallbackKeyString.ValueString()
	} else {
		data.MacsecFallbackKeyString = ""
	}

	if !v.MacsecKeyString.IsNull() && !v.MacsecKeyString.IsUnknown() {
		data.MacsecKeyString = v.MacsecKeyString.ValueString()
	} else {
		data.MacsecKeyString = ""
	}

	if !v.MacsecReportTimer.IsNull() && !v.MacsecReportTimer.IsUnknown() {
		data.MacsecReportTimer = new(Int64Custom)
		*data.MacsecReportTimer = Int64Custom(v.MacsecReportTimer.ValueInt64())
	} else {
		data.MacsecReportTimer = nil
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

	if !v.MplsIsisAreaNum.IsNull() && !v.MplsIsisAreaNum.IsUnknown() {
		data.MplsIsisAreaNum = v.MplsIsisAreaNum.ValueString()
	} else {
		data.MplsIsisAreaNum = ""
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

	if !v.MstInstanceRange.IsNull() && !v.MstInstanceRange.IsUnknown() {
		data.MstInstanceRange = v.MstInstanceRange.ValueString()
	} else {
		data.MstInstanceRange = ""
	}

	if !v.MulticastGroupSubnet.IsNull() && !v.MulticastGroupSubnet.IsUnknown() {
		data.MulticastGroupSubnet = v.MulticastGroupSubnet.ValueString()
	} else {
		data.MulticastGroupSubnet = ""
	}

	if !v.MvpnVriIdRange.IsNull() && !v.MvpnVriIdRange.IsUnknown() {
		data.MvpnVriIdRange = v.MvpnVriIdRange.ValueString()
	} else {
		data.MvpnVriIdRange = ""
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

	if !v.NetworkVlanRange.IsNull() && !v.NetworkVlanRange.IsUnknown() {
		data.NetworkVlanRange = v.NetworkVlanRange.ValueString()
	} else {
		data.NetworkVlanRange = ""
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

	if !v.NveLbId.IsNull() && !v.NveLbId.IsUnknown() {
		data.NveLbId = new(Int64Custom)
		*data.NveLbId = Int64Custom(v.NveLbId.ValueInt64())
	} else {
		data.NveLbId = nil
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

	if !v.ObjectTrackingNumberRange.IsNull() && !v.ObjectTrackingNumberRange.IsUnknown() {
		data.ObjectTrackingNumberRange = v.ObjectTrackingNumberRange.ValueString()
	} else {
		data.ObjectTrackingNumberRange = ""
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

	if !v.OverlayMode.IsNull() && !v.OverlayMode.IsUnknown() {
		data.OverlayMode = v.OverlayMode.ValueString()
	} else {
		data.OverlayMode = ""
	}

	if !v.OverwriteGlobalNxc.IsNull() && !v.OverwriteGlobalNxc.IsUnknown() {
		data.OverwriteGlobalNxc = strconv.FormatBool(v.OverwriteGlobalNxc.ValueBool())
	} else {
		data.OverwriteGlobalNxc = ""
	}

	if !v.PerVrfLoopbackAutoProvision.IsNull() && !v.PerVrfLoopbackAutoProvision.IsUnknown() {
		data.PerVrfLoopbackAutoProvision = strconv.FormatBool(v.PerVrfLoopbackAutoProvision.ValueBool())
	} else {
		data.PerVrfLoopbackAutoProvision = ""
	}

	if !v.PerVrfLoopbackAutoProvisionV6.IsNull() && !v.PerVrfLoopbackAutoProvisionV6.IsUnknown() {
		data.PerVrfLoopbackAutoProvisionV6 = strconv.FormatBool(v.PerVrfLoopbackAutoProvisionV6.ValueBool())
	} else {
		data.PerVrfLoopbackAutoProvisionV6 = ""
	}

	if !v.PerVrfLoopbackIpRange.IsNull() && !v.PerVrfLoopbackIpRange.IsUnknown() {
		data.PerVrfLoopbackIpRange = v.PerVrfLoopbackIpRange.ValueString()
	} else {
		data.PerVrfLoopbackIpRange = ""
	}

	if !v.PerVrfLoopbackIpRangeV6.IsNull() && !v.PerVrfLoopbackIpRangeV6.IsUnknown() {
		data.PerVrfLoopbackIpRangeV6 = v.PerVrfLoopbackIpRangeV6.ValueString()
	} else {
		data.PerVrfLoopbackIpRangeV6 = ""
	}

	if !v.PfcWatchInt.IsNull() && !v.PfcWatchInt.IsUnknown() {
		data.PfcWatchInt = new(Int64Custom)
		*data.PfcWatchInt = Int64Custom(v.PfcWatchInt.ValueInt64())
	} else {
		data.PfcWatchInt = nil
	}

	if !v.PhantomRpLbId1.IsNull() && !v.PhantomRpLbId1.IsUnknown() {
		data.PhantomRpLbId1 = new(Int64Custom)
		*data.PhantomRpLbId1 = Int64Custom(v.PhantomRpLbId1.ValueInt64())
	} else {
		data.PhantomRpLbId1 = nil
	}

	if !v.PhantomRpLbId2.IsNull() && !v.PhantomRpLbId2.IsUnknown() {
		data.PhantomRpLbId2 = new(Int64Custom)
		*data.PhantomRpLbId2 = Int64Custom(v.PhantomRpLbId2.ValueInt64())
	} else {
		data.PhantomRpLbId2 = nil
	}

	if !v.PhantomRpLbId3.IsNull() && !v.PhantomRpLbId3.IsUnknown() {
		data.PhantomRpLbId3 = new(Int64Custom)
		*data.PhantomRpLbId3 = Int64Custom(v.PhantomRpLbId3.ValueInt64())
	} else {
		data.PhantomRpLbId3 = nil
	}

	if !v.PhantomRpLbId4.IsNull() && !v.PhantomRpLbId4.IsUnknown() {
		data.PhantomRpLbId4 = new(Int64Custom)
		*data.PhantomRpLbId4 = Int64Custom(v.PhantomRpLbId4.ValueInt64())
	} else {
		data.PhantomRpLbId4 = nil
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

	if !v.PtpVlanId.IsNull() && !v.PtpVlanId.IsUnknown() {
		data.PtpVlanId = new(Int64Custom)
		*data.PtpVlanId = Int64Custom(v.PtpVlanId.ValueInt64())
	} else {
		data.PtpVlanId = nil
	}

	if !v.QkdProfileName.IsNull() && !v.QkdProfileName.IsUnknown() {
		data.QkdProfileName = v.QkdProfileName.ValueString()
	} else {
		data.QkdProfileName = ""
	}

	if !v.ReplicationMode.IsNull() && !v.ReplicationMode.IsUnknown() {
		data.ReplicationMode = v.ReplicationMode.ValueString()
	} else {
		data.ReplicationMode = ""
	}

	if !v.RouterIdRange.IsNull() && !v.RouterIdRange.IsUnknown() {
		data.RouterIdRange = v.RouterIdRange.ValueString()
	} else {
		data.RouterIdRange = ""
	}

	if !v.RouteMapSequenceNumberRange.IsNull() && !v.RouteMapSequenceNumberRange.IsUnknown() {
		data.RouteMapSequenceNumberRange = v.RouteMapSequenceNumberRange.ValueString()
	} else {
		data.RouteMapSequenceNumberRange = ""
	}

	if !v.RpCount.IsNull() && !v.RpCount.IsUnknown() {
		data.RpCount = new(Int64Custom)
		*data.RpCount = Int64Custom(v.RpCount.ValueInt64())
	} else {
		data.RpCount = nil
	}

	if !v.RpLbId.IsNull() && !v.RpLbId.IsUnknown() {
		data.RpLbId = new(Int64Custom)
		*data.RpLbId = Int64Custom(v.RpLbId.ValueInt64())
	} else {
		data.RpLbId = nil
	}

	if !v.RpMode.IsNull() && !v.RpMode.IsUnknown() {
		data.RpMode = v.RpMode.ValueString()
	} else {
		data.RpMode = ""
	}

	if !v.RrCount.IsNull() && !v.RrCount.IsUnknown() {
		data.RrCount = new(Int64Custom)
		*data.RrCount = Int64Custom(v.RrCount.ValueInt64())
	} else {
		data.RrCount = nil
	}

	if !v.SeedSwitchCoreInterfaces.IsNull() && !v.SeedSwitchCoreInterfaces.IsUnknown() {
		data.SeedSwitchCoreInterfaces = v.SeedSwitchCoreInterfaces.ValueString()
	} else {
		data.SeedSwitchCoreInterfaces = ""
	}

	if !v.ServiceNetworkVlanRange.IsNull() && !v.ServiceNetworkVlanRange.IsUnknown() {
		data.ServiceNetworkVlanRange = v.ServiceNetworkVlanRange.ValueString()
	} else {
		data.ServiceNetworkVlanRange = ""
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

	if !v.SiteId.IsNull() && !v.SiteId.IsUnknown() {
		data.SiteId = v.SiteId.ValueString()
	} else {
		data.SiteId = ""
	}

	if !v.SlaIdRange.IsNull() && !v.SlaIdRange.IsUnknown() {
		data.SlaIdRange = v.SlaIdRange.ValueString()
	} else {
		data.SlaIdRange = ""
	}

	if !v.SnmpServerHostTrap.IsNull() && !v.SnmpServerHostTrap.IsUnknown() {
		data.SnmpServerHostTrap = strconv.FormatBool(v.SnmpServerHostTrap.ValueBool())
	} else {
		data.SnmpServerHostTrap = ""
	}

	if !v.SpineSwitchCoreInterfaces.IsNull() && !v.SpineSwitchCoreInterfaces.IsUnknown() {
		data.SpineSwitchCoreInterfaces = v.SpineSwitchCoreInterfaces.ValueString()
	} else {
		data.SpineSwitchCoreInterfaces = ""
	}

	if !v.SspineAddDelDebugFlag.IsNull() && !v.SspineAddDelDebugFlag.IsUnknown() {
		data.SspineAddDelDebugFlag = v.SspineAddDelDebugFlag.ValueString()
	} else {
		data.SspineAddDelDebugFlag = ""
	}

	if !v.StaticUnderlayIpAlloc.IsNull() && !v.StaticUnderlayIpAlloc.IsUnknown() {
		data.StaticUnderlayIpAlloc = strconv.FormatBool(v.StaticUnderlayIpAlloc.ValueBool())
	} else {
		data.StaticUnderlayIpAlloc = ""
	}

	if !v.StpBridgePriority.IsNull() && !v.StpBridgePriority.IsUnknown() {
		data.StpBridgePriority = new(Int64Custom)
		*data.StpBridgePriority = Int64Custom(v.StpBridgePriority.ValueInt64())
	} else {
		data.StpBridgePriority = nil
	}

	if !v.StpRootOption.IsNull() && !v.StpRootOption.IsUnknown() {
		data.StpRootOption = v.StpRootOption.ValueString()
	} else {
		data.StpRootOption = ""
	}

	if !v.StpVlanRange.IsNull() && !v.StpVlanRange.IsUnknown() {
		data.StpVlanRange = v.StpVlanRange.ValueString()
	} else {
		data.StpVlanRange = ""
	}

	if !v.StrictCcMode.IsNull() && !v.StrictCcMode.IsUnknown() {
		data.StrictCcMode = strconv.FormatBool(v.StrictCcMode.ValueBool())
	} else {
		data.StrictCcMode = ""
	}

	if !v.SubinterfaceRange.IsNull() && !v.SubinterfaceRange.IsUnknown() {
		data.SubinterfaceRange = v.SubinterfaceRange.ValueString()
	} else {
		data.SubinterfaceRange = ""
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

	if !v.TcamAllocation.IsNull() && !v.TcamAllocation.IsUnknown() {
		data.TcamAllocation = strconv.FormatBool(v.TcamAllocation.ValueBool())
	} else {
		data.TcamAllocation = ""
	}

	if !v.TrustpointLabel.IsNull() && !v.TrustpointLabel.IsUnknown() {
		data.TrustpointLabel = v.TrustpointLabel.ValueString()
	} else {
		data.TrustpointLabel = ""
	}

	if !v.UnderlayIsV6.IsNull() && !v.UnderlayIsV6.IsUnknown() {
		data.UnderlayIsV6 = strconv.FormatBool(v.UnderlayIsV6.ValueBool())
	} else {
		data.UnderlayIsV6 = ""
	}

	if !v.UnnumBootstrapLbId.IsNull() && !v.UnnumBootstrapLbId.IsUnknown() {
		data.UnnumBootstrapLbId = new(Int64Custom)
		*data.UnnumBootstrapLbId = Int64Custom(v.UnnumBootstrapLbId.ValueInt64())
	} else {
		data.UnnumBootstrapLbId = nil
	}

	if !v.UseLinkLocal.IsNull() && !v.UseLinkLocal.IsUnknown() {
		data.UseLinkLocal = strconv.FormatBool(v.UseLinkLocal.ValueBool())
	} else {
		data.UseLinkLocal = ""
	}

	if !v.V6SubnetRange.IsNull() && !v.V6SubnetRange.IsUnknown() {
		data.V6SubnetRange = v.V6SubnetRange.ValueString()
	} else {
		data.V6SubnetRange = ""
	}

	if !v.V6SubnetTargetMask.IsNull() && !v.V6SubnetTargetMask.IsUnknown() {
		data.V6SubnetTargetMask = new(Int64Custom)
		*data.V6SubnetTargetMask = Int64Custom(v.V6SubnetTargetMask.ValueInt64())
	} else {
		data.V6SubnetTargetMask = nil
	}

	if !v.VpcAutoRecoveryTime.IsNull() && !v.VpcAutoRecoveryTime.IsUnknown() {
		data.VpcAutoRecoveryTime = new(Int64Custom)
		*data.VpcAutoRecoveryTime = Int64Custom(v.VpcAutoRecoveryTime.ValueInt64())
	} else {
		data.VpcAutoRecoveryTime = nil
	}

	if !v.VpcDelayRestore.IsNull() && !v.VpcDelayRestore.IsUnknown() {
		data.VpcDelayRestore = new(Int64Custom)
		*data.VpcDelayRestore = Int64Custom(v.VpcDelayRestore.ValueInt64())
	} else {
		data.VpcDelayRestore = nil
	}

	if !v.VpcDelayRestoreTime.IsNull() && !v.VpcDelayRestoreTime.IsUnknown() {
		data.VpcDelayRestoreTime = new(Int64Custom)
		*data.VpcDelayRestoreTime = Int64Custom(v.VpcDelayRestoreTime.ValueInt64())
	} else {
		data.VpcDelayRestoreTime = nil
	}

	if !v.VpcDomainIdRange.IsNull() && !v.VpcDomainIdRange.IsUnknown() {
		data.VpcDomainIdRange = v.VpcDomainIdRange.ValueString()
	} else {
		data.VpcDomainIdRange = ""
	}

	if !v.VpcEnableIpv6NdSync.IsNull() && !v.VpcEnableIpv6NdSync.IsUnknown() {
		data.VpcEnableIpv6NdSync = strconv.FormatBool(v.VpcEnableIpv6NdSync.ValueBool())
	} else {
		data.VpcEnableIpv6NdSync = ""
	}

	if !v.VpcPeerKeepAliveOption.IsNull() && !v.VpcPeerKeepAliveOption.IsUnknown() {
		data.VpcPeerKeepAliveOption = v.VpcPeerKeepAliveOption.ValueString()
	} else {
		data.VpcPeerKeepAliveOption = ""
	}

	if !v.VpcPeerLinkPo.IsNull() && !v.VpcPeerLinkPo.IsUnknown() {
		data.VpcPeerLinkPo = new(Int64Custom)
		*data.VpcPeerLinkPo = Int64Custom(v.VpcPeerLinkPo.ValueInt64())
	} else {
		data.VpcPeerLinkPo = nil
	}

	if !v.VpcPeerLinkVlan.IsNull() && !v.VpcPeerLinkVlan.IsUnknown() {
		data.VpcPeerLinkVlan = new(Int64Custom)
		*data.VpcPeerLinkVlan = Int64Custom(v.VpcPeerLinkVlan.ValueInt64())
	} else {
		data.VpcPeerLinkVlan = nil
	}

	if !v.VrfLiteAutoconfig.IsNull() && !v.VrfLiteAutoconfig.IsUnknown() {
		data.VrfLiteAutoconfig = v.VrfLiteAutoconfig.ValueString()
	} else {
		data.VrfLiteAutoconfig = ""
	}

	if !v.VrfVlanRange.IsNull() && !v.VrfVlanRange.IsUnknown() {
		data.VrfVlanRange = v.VrfVlanRange.ValueString()
	} else {
		data.VrfVlanRange = ""
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

	if !v.Deploy.IsNull() && !v.Deploy.IsUnknown() {
		data.Deploy = v.Deploy.ValueBool()
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

	return data
}
