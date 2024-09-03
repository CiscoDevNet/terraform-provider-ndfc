// Code generated;  DO NOT EDIT.

package resource_vpc_pair

import (
	"context"
	"log"
	"strconv"
	. "terraform-provider-ndfc/internal/provider/types"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCVpcPairModel struct {
	SerialNumbers      []string         `json:"-"`
	PeerOneId          string           `json:"peerOneId,omitempty"`
	PeerTwoId          string           `json:"peerTwoId,omitempty"`
	UseVirtualPeerlink *bool            `json:"useVirtualPeerlink,omitempty"`
	TemplateName       string           `json:"templateName,omitempty"`
	NvPairs            NDFCNvPairsValue `json:"nvPairs,omitempty"`
}

type NDFCNvPairsValue struct {
	DomainId              *Int64Custom `json:"DOMAIN_ID,omitempty"`
	Peer1KeepAliveLocalIp string       `json:"PEER1_KEEP_ALIVE_LOCAL_IP,omitempty"`
	Peer2KeepAliveLocalIp string       `json:"PEER2_KEEP_ALIVE_LOCAL_IP,omitempty"`
	KeepAliveVrf          string       `json:"KEEP_ALIVE_VRF,omitempty"`
	KeepAliveHoldTimeout  *int64       `json:"KEEP_ALIVE_HOLD_TIMEOUT,omitempty"`
	IsVpcPlus             string       `json:"isVpcPlus,omitempty"`
	FabricpathSwitchId    string       `json:"fabricPath_switch_id,omitempty"`
	Peer1SourceLoopback   string       `json:"PEER1_SOURCE_LOOPBACK,omitempty"`
	Peer2SourceLoopback   string       `json:"PEER2_SOURCE_LOOPBACK,omitempty"`
	Peer1PrimaryIp        string       `json:"PEER1_PRIMARY_IP,omitempty"`
	Peer2PrimaryIp        string       `json:"PEER2_PRIMARY_IP,omitempty"`
	LoopbackSecondaryIp   string       `json:"LOOPBACK_SECONDARY_IP,omitempty"`
	Peer1DomainConf       string       `json:"PEER1_DOMAIN_CONF,omitempty"`
	ClearPolicy           string       `json:"clear_policy,omitempty"`
	FabricName            string       `json:"FABRIC_NAME,omitempty"`
	Peer1Pcid             string       `json:"PEER1_PCID,omitempty"`
	Peer2Pcid             string       `json:"PEER2_PCID,omitempty"`
	Peer1MemberInterfaces string       `json:"PEER1_MEMBER_INTERFACES,omitempty"`
	Peer2MemberInterfaces string       `json:"PEER2_MEMBER_INTERFACES,omitempty"`
	PcMode                string       `json:"PC_MODE,omitempty"`
	Peer1PoDesc           string       `json:"PEER1_PO_DESC,omitempty"`
	Peer2PoDesc           string       `json:"PEER2_PO_DESC,omitempty"`
	AdminState            string       `json:"ADMIN_STATE,omitempty"`
	AllowedVlans          string       `json:"ALLOWED_VLANS,omitempty"`
	Peer1PoConf           string       `json:"PEER1_PO_CONF,omitempty"`
	Peer2PoConf           string       `json:"PEER2_PO_CONF,omitempty"`
	IsVteps               string       `json:"isVTEPS,omitempty"`
	NveInterface          string       `json:"NVE_INTERFACE,omitempty"`
}

func (v *VpcPairModel) SetModelData(jsonData *NDFCVpcPairModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if len(jsonData.SerialNumbers) == 0 {
		log.Printf("v.SerialNumbers is empty")
		v.SerialNumbers, err = types.SetValue(types.StringType, []attr.Value{})
		if err != nil {
			log.Printf("Error in converting []string to  List %v", err)
			return err
		}
	} else {
		listData := make([]attr.Value, len(jsonData.SerialNumbers))
		for i, item := range jsonData.SerialNumbers {
			listData[i] = types.StringValue(item)
		}
		v.SerialNumbers, err = types.SetValue(types.StringType, listData)
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}

	if jsonData.UseVirtualPeerlink != nil {
		v.UseVirtualPeerlink = types.BoolValue(*jsonData.UseVirtualPeerlink)

	} else {
		v.UseVirtualPeerlink = types.BoolNull()
	}

	if jsonData.TemplateName != "" {
		v.TemplateName = types.StringValue(jsonData.TemplateName)
	} else {
		v.TemplateName = types.StringNull()
	}

	if jsonData.NvPairs.DomainId != nil {
		if jsonData.NvPairs.DomainId.IsEmpty() {
			v.DomainId = types.Int64Null()
		} else {
			v.DomainId = types.Int64Value(int64(*jsonData.NvPairs.DomainId))
		}

	} else {
		v.DomainId = types.Int64Null()
	}

	if jsonData.NvPairs.Peer1KeepAliveLocalIp != "" {
		v.Peer1KeepAliveLocalIp = types.StringValue(jsonData.NvPairs.Peer1KeepAliveLocalIp)
	} else {
		v.Peer1KeepAliveLocalIp = types.StringNull()
	}

	if jsonData.NvPairs.Peer2KeepAliveLocalIp != "" {
		v.Peer2KeepAliveLocalIp = types.StringValue(jsonData.NvPairs.Peer2KeepAliveLocalIp)
	} else {
		v.Peer2KeepAliveLocalIp = types.StringNull()
	}

	if jsonData.NvPairs.KeepAliveVrf != "" {
		v.KeepAliveVrf = types.StringValue(jsonData.NvPairs.KeepAliveVrf)
	} else {
		v.KeepAliveVrf = types.StringNull()
	}

	if jsonData.NvPairs.KeepAliveHoldTimeout != nil {
		v.KeepAliveHoldTimeout = types.Int64Value(*jsonData.NvPairs.KeepAliveHoldTimeout)

	} else {
		v.KeepAliveHoldTimeout = types.Int64Null()
	}

	if jsonData.NvPairs.IsVpcPlus != "" {
		x, _ := strconv.ParseBool(jsonData.NvPairs.IsVpcPlus)
		v.IsVpcPlus = types.BoolValue(x)
	} else {
		v.IsVpcPlus = types.BoolNull()
	}

	if jsonData.NvPairs.FabricpathSwitchId != "" {
		v.FabricpathSwitchId = types.StringValue(jsonData.NvPairs.FabricpathSwitchId)
	} else {
		v.FabricpathSwitchId = types.StringNull()
	}

	if jsonData.NvPairs.Peer1SourceLoopback != "" {
		v.Peer1SourceLoopback = types.StringValue(jsonData.NvPairs.Peer1SourceLoopback)
	} else {
		v.Peer1SourceLoopback = types.StringNull()
	}

	if jsonData.NvPairs.Peer2SourceLoopback != "" {
		v.Peer2SourceLoopback = types.StringValue(jsonData.NvPairs.Peer2SourceLoopback)
	} else {
		v.Peer2SourceLoopback = types.StringNull()
	}

	if jsonData.NvPairs.Peer1PrimaryIp != "" {
		v.Peer1PrimaryIp = types.StringValue(jsonData.NvPairs.Peer1PrimaryIp)
	} else {
		v.Peer1PrimaryIp = types.StringNull()
	}

	if jsonData.NvPairs.Peer2PrimaryIp != "" {
		v.Peer2PrimaryIp = types.StringValue(jsonData.NvPairs.Peer2PrimaryIp)
	} else {
		v.Peer2PrimaryIp = types.StringNull()
	}

	if jsonData.NvPairs.LoopbackSecondaryIp != "" {
		v.LoopbackSecondaryIp = types.StringValue(jsonData.NvPairs.LoopbackSecondaryIp)
	} else {
		v.LoopbackSecondaryIp = types.StringNull()
	}

	if jsonData.NvPairs.Peer1DomainConf != "" {
		v.Peer1DomainConf = types.StringValue(jsonData.NvPairs.Peer1DomainConf)
	} else {
		v.Peer1DomainConf = types.StringNull()
	}

	if jsonData.NvPairs.ClearPolicy != "" {
		x, _ := strconv.ParseBool(jsonData.NvPairs.ClearPolicy)
		v.ClearPolicy = types.BoolValue(x)
	} else {
		v.ClearPolicy = types.BoolNull()
	}

	if jsonData.NvPairs.FabricName != "" {
		v.FabricName = types.StringValue(jsonData.NvPairs.FabricName)
	} else {
		v.FabricName = types.StringNull()
	}

	if jsonData.NvPairs.Peer1Pcid != "" {
		v.Peer1Pcid = types.StringValue(jsonData.NvPairs.Peer1Pcid)
	} else {
		v.Peer1Pcid = types.StringNull()
	}

	if jsonData.NvPairs.Peer2Pcid != "" {
		v.Peer2Pcid = types.StringValue(jsonData.NvPairs.Peer2Pcid)
	} else {
		v.Peer2Pcid = types.StringNull()
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

	if jsonData.NvPairs.PcMode != "" {
		v.PcMode = types.StringValue(jsonData.NvPairs.PcMode)
	} else {
		v.PcMode = types.StringNull()
	}

	if jsonData.NvPairs.Peer1PoDesc != "" {
		v.Peer1PoDesc = types.StringValue(jsonData.NvPairs.Peer1PoDesc)
	} else {
		v.Peer1PoDesc = types.StringNull()
	}

	if jsonData.NvPairs.Peer2PoDesc != "" {
		v.Peer2PoDesc = types.StringValue(jsonData.NvPairs.Peer2PoDesc)
	} else {
		v.Peer2PoDesc = types.StringNull()
	}

	if jsonData.NvPairs.AdminState != "" {
		x, _ := strconv.ParseBool(jsonData.NvPairs.AdminState)
		v.AdminState = types.BoolValue(x)
	} else {
		v.AdminState = types.BoolNull()
	}

	if jsonData.NvPairs.AllowedVlans != "" {
		v.AllowedVlans = types.StringValue(jsonData.NvPairs.AllowedVlans)
	} else {
		v.AllowedVlans = types.StringNull()
	}

	if jsonData.NvPairs.Peer1PoConf != "" {
		v.Peer1PoConf = types.StringValue(jsonData.NvPairs.Peer1PoConf)
	} else {
		v.Peer1PoConf = types.StringNull()
	}

	if jsonData.NvPairs.Peer2PoConf != "" {
		v.Peer2PoConf = types.StringValue(jsonData.NvPairs.Peer2PoConf)
	} else {
		v.Peer2PoConf = types.StringNull()
	}

	if jsonData.NvPairs.IsVteps != "" {
		x, _ := strconv.ParseBool(jsonData.NvPairs.IsVteps)
		v.IsVteps = types.BoolValue(x)
	} else {
		v.IsVteps = types.BoolNull()
	}

	if jsonData.NvPairs.NveInterface != "" {
		v.NveInterface = types.StringValue(jsonData.NvPairs.NveInterface)
	} else {
		v.NveInterface = types.StringNull()
	}

	return err
}

func (v VpcPairModel) GetModelData() *NDFCVpcPairModel {
	var data = new(NDFCVpcPairModel)

	//MARSHAL_BODY

	if !v.SerialNumbers.IsNull() && !v.SerialNumbers.IsUnknown() {
		listStringData := make([]string, len(v.SerialNumbers.Elements()))
		dg := v.SerialNumbers.ElementsAs(context.Background(), &listStringData, false)
		if dg.HasError() {
			panic(dg.Errors())
		}
		data.SerialNumbers = make([]string, len(listStringData))
		copy(data.SerialNumbers, listStringData)
	}

	if !v.UseVirtualPeerlink.IsNull() && !v.UseVirtualPeerlink.IsUnknown() {
		data.UseVirtualPeerlink = new(bool)
		*data.UseVirtualPeerlink = v.UseVirtualPeerlink.ValueBool()
	} else {
		data.UseVirtualPeerlink = nil
	}

	if !v.TemplateName.IsNull() && !v.TemplateName.IsUnknown() {
		data.TemplateName = v.TemplateName.ValueString()
	} else {
		data.TemplateName = ""
	}

	if !v.DomainId.IsNull() && !v.DomainId.IsUnknown() {
		data.NvPairs.DomainId = new(Int64Custom)
		*data.NvPairs.DomainId = Int64Custom(v.DomainId.ValueInt64())
	} else {
		data.NvPairs.DomainId = nil
	}

	if !v.Peer1KeepAliveLocalIp.IsNull() && !v.Peer1KeepAliveLocalIp.IsUnknown() {
		data.NvPairs.Peer1KeepAliveLocalIp = v.Peer1KeepAliveLocalIp.ValueString()
	} else {
		data.NvPairs.Peer1KeepAliveLocalIp = ""
	}

	if !v.Peer2KeepAliveLocalIp.IsNull() && !v.Peer2KeepAliveLocalIp.IsUnknown() {
		data.NvPairs.Peer2KeepAliveLocalIp = v.Peer2KeepAliveLocalIp.ValueString()
	} else {
		data.NvPairs.Peer2KeepAliveLocalIp = ""
	}

	if !v.KeepAliveVrf.IsNull() && !v.KeepAliveVrf.IsUnknown() {
		data.NvPairs.KeepAliveVrf = v.KeepAliveVrf.ValueString()
	} else {
		data.NvPairs.KeepAliveVrf = ""
	}

	if !v.KeepAliveHoldTimeout.IsNull() && !v.KeepAliveHoldTimeout.IsUnknown() {
		data.NvPairs.KeepAliveHoldTimeout = new(int64)
		*data.NvPairs.KeepAliveHoldTimeout = v.KeepAliveHoldTimeout.ValueInt64()

	} else {
		data.NvPairs.KeepAliveHoldTimeout = nil
	}

	if !v.IsVpcPlus.IsNull() && !v.IsVpcPlus.IsUnknown() {
		data.NvPairs.IsVpcPlus = strconv.FormatBool(v.IsVpcPlus.ValueBool())
	} else {
		data.NvPairs.IsVpcPlus = ""
	}

	if !v.FabricpathSwitchId.IsNull() && !v.FabricpathSwitchId.IsUnknown() {
		data.NvPairs.FabricpathSwitchId = v.FabricpathSwitchId.ValueString()
	} else {
		data.NvPairs.FabricpathSwitchId = ""
	}

	if !v.Peer1SourceLoopback.IsNull() && !v.Peer1SourceLoopback.IsUnknown() {
		data.NvPairs.Peer1SourceLoopback = v.Peer1SourceLoopback.ValueString()
	} else {
		data.NvPairs.Peer1SourceLoopback = ""
	}

	if !v.Peer2SourceLoopback.IsNull() && !v.Peer2SourceLoopback.IsUnknown() {
		data.NvPairs.Peer2SourceLoopback = v.Peer2SourceLoopback.ValueString()
	} else {
		data.NvPairs.Peer2SourceLoopback = ""
	}

	if !v.Peer1PrimaryIp.IsNull() && !v.Peer1PrimaryIp.IsUnknown() {
		data.NvPairs.Peer1PrimaryIp = v.Peer1PrimaryIp.ValueString()
	} else {
		data.NvPairs.Peer1PrimaryIp = ""
	}

	if !v.Peer2PrimaryIp.IsNull() && !v.Peer2PrimaryIp.IsUnknown() {
		data.NvPairs.Peer2PrimaryIp = v.Peer2PrimaryIp.ValueString()
	} else {
		data.NvPairs.Peer2PrimaryIp = ""
	}

	if !v.LoopbackSecondaryIp.IsNull() && !v.LoopbackSecondaryIp.IsUnknown() {
		data.NvPairs.LoopbackSecondaryIp = v.LoopbackSecondaryIp.ValueString()
	} else {
		data.NvPairs.LoopbackSecondaryIp = ""
	}

	if !v.Peer1DomainConf.IsNull() && !v.Peer1DomainConf.IsUnknown() {
		data.NvPairs.Peer1DomainConf = v.Peer1DomainConf.ValueString()
	} else {
		data.NvPairs.Peer1DomainConf = ""
	}

	if !v.ClearPolicy.IsNull() && !v.ClearPolicy.IsUnknown() {
		data.NvPairs.ClearPolicy = strconv.FormatBool(v.ClearPolicy.ValueBool())
	} else {
		data.NvPairs.ClearPolicy = ""
	}

	if !v.FabricName.IsNull() && !v.FabricName.IsUnknown() {
		data.NvPairs.FabricName = v.FabricName.ValueString()
	} else {
		data.NvPairs.FabricName = ""
	}

	if !v.Peer1Pcid.IsNull() && !v.Peer1Pcid.IsUnknown() {
		data.NvPairs.Peer1Pcid = v.Peer1Pcid.ValueString()
	} else {
		data.NvPairs.Peer1Pcid = ""
	}

	if !v.Peer2Pcid.IsNull() && !v.Peer2Pcid.IsUnknown() {
		data.NvPairs.Peer2Pcid = v.Peer2Pcid.ValueString()
	} else {
		data.NvPairs.Peer2Pcid = ""
	}

	if !v.Peer1MemberInterfaces.IsNull() && !v.Peer1MemberInterfaces.IsUnknown() {
		data.NvPairs.Peer1MemberInterfaces = v.Peer1MemberInterfaces.ValueString()
	} else {
		data.NvPairs.Peer1MemberInterfaces = ""
	}

	if !v.Peer2MemberInterfaces.IsNull() && !v.Peer2MemberInterfaces.IsUnknown() {
		data.NvPairs.Peer2MemberInterfaces = v.Peer2MemberInterfaces.ValueString()
	} else {
		data.NvPairs.Peer2MemberInterfaces = ""
	}

	if !v.PcMode.IsNull() && !v.PcMode.IsUnknown() {
		data.NvPairs.PcMode = v.PcMode.ValueString()
	} else {
		data.NvPairs.PcMode = ""
	}

	if !v.Peer1PoDesc.IsNull() && !v.Peer1PoDesc.IsUnknown() {
		data.NvPairs.Peer1PoDesc = v.Peer1PoDesc.ValueString()
	} else {
		data.NvPairs.Peer1PoDesc = ""
	}

	if !v.Peer2PoDesc.IsNull() && !v.Peer2PoDesc.IsUnknown() {
		data.NvPairs.Peer2PoDesc = v.Peer2PoDesc.ValueString()
	} else {
		data.NvPairs.Peer2PoDesc = ""
	}

	if !v.AdminState.IsNull() && !v.AdminState.IsUnknown() {
		data.NvPairs.AdminState = strconv.FormatBool(v.AdminState.ValueBool())
	} else {
		data.NvPairs.AdminState = ""
	}

	if !v.AllowedVlans.IsNull() && !v.AllowedVlans.IsUnknown() {
		data.NvPairs.AllowedVlans = v.AllowedVlans.ValueString()
	} else {
		data.NvPairs.AllowedVlans = ""
	}

	if !v.Peer1PoConf.IsNull() && !v.Peer1PoConf.IsUnknown() {
		data.NvPairs.Peer1PoConf = v.Peer1PoConf.ValueString()
	} else {
		data.NvPairs.Peer1PoConf = ""
	}

	if !v.Peer2PoConf.IsNull() && !v.Peer2PoConf.IsUnknown() {
		data.NvPairs.Peer2PoConf = v.Peer2PoConf.ValueString()
	} else {
		data.NvPairs.Peer2PoConf = ""
	}

	if !v.IsVteps.IsNull() && !v.IsVteps.IsUnknown() {
		data.NvPairs.IsVteps = strconv.FormatBool(v.IsVteps.ValueBool())
	} else {
		data.NvPairs.IsVteps = ""
	}

	if !v.NveInterface.IsNull() && !v.NveInterface.IsUnknown() {
		data.NvPairs.NveInterface = v.NveInterface.ValueString()
	} else {
		data.NvPairs.NveInterface = ""
	}

	return data
}
