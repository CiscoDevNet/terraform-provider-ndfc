// Code generated DO NOT EDIT.
package resource_vrf_attachments

import (
	"context"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Int64Custom int64

func (i *Int64Custom) UnmarshalJSON(data []byte) error {
	if string(data) == "" || string(data) == "\"\"" {
		*i = -9223372036854775808
	} else {
		ss := string(data)
		// If the string is quoted, remove the quotes
		ssUn, err := strconv.Unquote(ss)
		if err == nil {
			// Quote removed
			ss = ssUn
		}
		ii, _ := strconv.ParseInt(ss, 10, 64)
		*i = Int64Custom(ii)
	}

	return nil
}

func (i Int64Custom) MarshalJSON() ([]byte, error) {
	res := ""
	res = strconv.FormatInt(int64(i), 10)
	return []byte(strconv.Quote(res)), nil

}

type NDFCVrfAttachmentsModel struct {
	FabricName           string                              `json:"fabric,omitempty"`
	DeployAllAttachments bool                                `json:"-"`
	VrfAttachments       NDFCVrfAttachmentsValues            `json:"attachments,omitempty"`
	VrfAttachmentsMap    map[string]*NDFCVrfAttachmentsValue `json:"-"`
}

type NDFCVrfAttachmentsValues []NDFCVrfAttachmentsValue

type NDFCVrfAttachmentsValue struct {
	Id                   *int64                          `json:"-"`
	FilterThisValue      bool                            `json:"-"`
	VrfName              string                          `json:"vrfName,omitempty"`
	DeployAllAttachments bool                            `json:"-"`
	AttachList           NDFCAttachListValues            `json:"lanAttachList,omitempty"`
	AttachListMap        map[string]*NDFCAttachListValue `json:"-"`
}

type NDFCAttachListValues []NDFCAttachListValue

type NDFCAttachListValue struct {
	FilterThisValue      bool                    `json:"-"`
	Id                   *int64                  `json:"-"`
	FabricName           string                  `json:"fabric,omitempty"`
	VrfName              string                  `json:"vrfName,omitempty"`
	SerialNumber         string                  `json:"serialNumber,omitempty"`
	SwitchSerialNo       string                  `json:"switchSerialNo,omitempty"`
	SwitchName           string                  `json:"switchName,omitempty"`
	Vlan                 *Int64Custom            `json:"vlan,omitempty"`
	VlanId               *Int64Custom            `json:"vlanId,omitempty"`
	Deployment           string                  `json:"deployment,omitempty"`
	AttachState          string                  `json:"lanAttachState,omitempty"`
	Attached             *bool                   `json:"isLanAttached,omitempty"`
	FreeformConfig       string                  `json:"freeformconfig,omitempty"`
	DeployThisAttachment bool                    `json:"-"`
	InstanceValues       NDFCInstanceValuesValue `json:"instanceValues,omitempty"`
}

type NDFCInstanceValuesValue struct {
	LoopbackId   *Int64Custom `json:"loopbackId,omitempty"`
	LoopbackIpv4 string       `json:"loopbackIpv4,omitempty"`
	LoopbackIpv6 string       `json:"loopbackIpv6,omitempty"`
}

func (s NDFCAttachListValues) Len() int {
	return len(s)
}

func (s NDFCAttachListValues) Less(i, j int) bool {
	return (*s[i].Id < *s[j].Id)

}

func (s NDFCAttachListValues) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s NDFCVrfAttachmentsValues) Len() int {
	return len(s)
}

func (s NDFCVrfAttachmentsValues) Less(i, j int) bool {
	return (*s[i].Id < *s[j].Id)

}

func (s NDFCVrfAttachmentsValues) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (v *VrfAttachmentsModel) SetModelData(jsonData *NDFCVrfAttachmentsModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.FabricName != "" {
		v.FabricName = types.StringValue(jsonData.FabricName)
	} else {
		v.FabricName = types.StringNull()
	}

	v.DeployAllAttachments = types.BoolValue(jsonData.DeployAllAttachments)
	if len(jsonData.VrfAttachments) == 0 {
		log.Printf("v.VrfAttachments is empty")
		v.VrfAttachments = types.ListNull(VrfAttachmentsValue{}.Type(context.Background()))
	} else {
		log.Printf("v.VrfAttachments contains %d elements", len(jsonData.VrfAttachments))
		listData := make([]VrfAttachmentsValue, 0)
		for _, item := range jsonData.VrfAttachments {
			if item.FilterThisValue {
				//Skip this entry - this parameter allows filtering
				continue
			}

			data := new(VrfAttachmentsValue)
			err = data.SetValue(&item)
			if err != nil {
				log.Printf("Error in VrfAttachmentsValue.SetValue")
				return err
			}
			data.state = attr.ValueStateKnown
			listData = append(listData, *data)
		}
		v.VrfAttachments, err = types.ListValueFrom(context.Background(), VrfAttachmentsValue{}.Type(context.Background()), listData)
		if err != nil {
			log.Printf("Error in converting []VrfAttachmentsValue to  List")
			return err
		}
	}

	return err
}

func (v *VrfAttachmentsValue) SetValue(jsonData *NDFCVrfAttachmentsValue) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.VrfName != "" {
		v.VrfName = types.StringValue(jsonData.VrfName)
	} else {
		v.VrfName = types.StringNull()
	}

	v.DeployAllAttachments = types.BoolValue(jsonData.DeployAllAttachments)
	if len(jsonData.AttachList) == 0 {
		log.Printf("v.AttachList is empty")
		v.AttachList = types.ListNull(AttachListValue{}.Type(context.Background()))
	} else {
		log.Printf("v.AttachList contains %d elements", len(jsonData.AttachList))
		listData := make([]AttachListValue, 0)
		for _, item := range jsonData.AttachList {
			if item.FilterThisValue {
				//Skip this entry - this parameter allows filtering
				continue
			}

			data := new(AttachListValue)
			err = data.SetValue(&item)
			if err != nil {
				log.Printf("Error in AttachListValue.SetValue")
				return err
			}
			data.state = attr.ValueStateKnown
			listData = append(listData, *data)
		}
		v.AttachList, err = types.ListValueFrom(context.Background(), AttachListValue{}.Type(context.Background()), listData)
		if err != nil {
			log.Printf("Error in converting []AttachListValue to  List")
			return err
		}
	}

	return err
}

func (v *AttachListValue) SetValue(jsonData *NDFCAttachListValue) diag.Diagnostics {
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

	if jsonData.Vlan != nil {
		if int64(*jsonData.Vlan) == -9223372036854775808 {
			v.Vlan = types.Int64Null()
		} else {
			v.Vlan = types.Int64Value(int64(*jsonData.Vlan))
		}
	} else if jsonData.VlanId != nil {
		if int64(*jsonData.VlanId) == -9223372036854775808 {
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

	v.DeployThisAttachment = types.BoolValue(jsonData.DeployThisAttachment)

	if jsonData.InstanceValues.LoopbackId != nil {
		if int64(*jsonData.InstanceValues.LoopbackId) == -9223372036854775808 {
			v.LoopbackId = types.Int64Null()
		} else {
			v.LoopbackId = types.Int64Value(int64(*jsonData.InstanceValues.LoopbackId))
		}

	} else {
		v.LoopbackId = types.Int64Null()
	}

	if jsonData.InstanceValues.LoopbackIpv4 != "" {
		v.LoopbackIpv4 = types.StringValue(jsonData.InstanceValues.LoopbackIpv4)
	} else {
		v.LoopbackIpv4 = types.StringNull()
	}

	if jsonData.InstanceValues.LoopbackIpv6 != "" {
		v.LoopbackIpv6 = types.StringValue(jsonData.InstanceValues.LoopbackIpv6)
	} else {
		v.LoopbackIpv6 = types.StringNull()
	}

	return err
}

func (v VrfAttachmentsModel) GetModelData() *NDFCVrfAttachmentsModel {
	var data = new(NDFCVrfAttachmentsModel)
	if !v.FabricName.IsNull() && !v.FabricName.IsUnknown() {
		data.FabricName = v.FabricName.ValueString()
	} else {
		data.FabricName = ""
	}

	if !v.DeployAllAttachments.IsNull() && !v.DeployAllAttachments.IsUnknown() {
		data.DeployAllAttachments = v.DeployAllAttachments.ValueBool()
	}

	data.VrfAttachmentsMap = make(map[string]*NDFCVrfAttachmentsValue)

	if !v.VrfAttachments.IsNull() && !v.VrfAttachments.IsUnknown() {
		elements := make([]VrfAttachmentsValue, len(v.VrfAttachments.Elements()))
		data.VrfAttachments = make([]NDFCVrfAttachmentsValue, len(v.VrfAttachments.Elements()))
		diag := v.VrfAttachments.ElementsAs(context.Background(), &elements, false)
		if diag != nil {
			panic(diag)
		}
		for i1, ele1 := range elements {

			if !ele1.VrfName.IsNull() && !ele1.VrfName.IsUnknown() {
				data.VrfAttachments[i1].VrfName = ele1.VrfName.ValueString()
			} else {
				data.VrfAttachments[i1].VrfName = ""
			}

			data.VrfAttachmentsMap[data.VrfAttachments[i1].VrfName] = &data.VrfAttachments[i1]

			if !ele1.DeployAllAttachments.IsNull() && !ele1.DeployAllAttachments.IsUnknown() {

				data.VrfAttachments[i1].DeployAllAttachments = ele1.DeployAllAttachments.ValueBool()
			}

			data.VrfAttachments[i1].AttachListMap = make(map[string]*NDFCAttachListValue)

			if !ele1.AttachList.IsNull() && !ele1.AttachList.IsUnknown() {
				elements := make([]AttachListValue, len(ele1.AttachList.Elements()))
				data.VrfAttachments[i1].AttachList = make([]NDFCAttachListValue, len(ele1.AttachList.Elements()))
				diag := ele1.AttachList.ElementsAs(context.Background(), &elements, false)
				if diag != nil {
					panic(diag)
				}
				for i2, ele2 := range elements {

					if !ele2.SerialNumber.IsNull() && !ele2.SerialNumber.IsUnknown() {
						data.VrfAttachments[i1].AttachList[i2].SerialNumber = ele2.SerialNumber.ValueString()
					} else {
						data.VrfAttachments[i1].AttachList[i2].SerialNumber = ""
					}

					data.VrfAttachments[i1].AttachListMap[data.VrfAttachments[i1].AttachList[i2].SerialNumber] = &data.VrfAttachments[i1].AttachList[i2]

					if !ele2.SwitchName.IsNull() && !ele2.SwitchName.IsUnknown() {
						data.VrfAttachments[i1].AttachList[i2].SwitchName = ele2.SwitchName.ValueString()
					} else {
						data.VrfAttachments[i1].AttachList[i2].SwitchName = ""
					}

					if !ele2.Vlan.IsNull() && !ele2.Vlan.IsUnknown() {
						data.VrfAttachments[i1].AttachList[i2].Vlan = new(Int64Custom)
						*data.VrfAttachments[i1].AttachList[i2].Vlan = Int64Custom(ele2.Vlan.ValueInt64())
					} else {
						data.VrfAttachments[i1].AttachList[i2].Vlan = nil
					}

					if !ele2.AttachState.IsNull() && !ele2.AttachState.IsUnknown() {
						data.VrfAttachments[i1].AttachList[i2].AttachState = ele2.AttachState.ValueString()
					} else {
						data.VrfAttachments[i1].AttachList[i2].AttachState = ""
					}

					if !ele2.Attached.IsNull() && !ele2.Attached.IsUnknown() {

						data.VrfAttachments[i1].AttachList[i2].Attached = new(bool)
						*data.VrfAttachments[i1].AttachList[i2].Attached = ele2.Attached.ValueBool()
					} else {
						data.VrfAttachments[i1].AttachList[i2].Attached = nil
					}

					if !ele2.FreeformConfig.IsNull() && !ele2.FreeformConfig.IsUnknown() {
						data.VrfAttachments[i1].AttachList[i2].FreeformConfig = ele2.FreeformConfig.ValueString()
					} else {
						data.VrfAttachments[i1].AttachList[i2].FreeformConfig = ""
					}

					if !ele2.DeployThisAttachment.IsNull() && !ele2.DeployThisAttachment.IsUnknown() {

						data.VrfAttachments[i1].AttachList[i2].DeployThisAttachment = ele2.DeployThisAttachment.ValueBool()
					}

					//-----inline nesting Start----
					if !ele2.LoopbackId.IsNull() && !ele2.LoopbackId.IsUnknown() {
						data.VrfAttachments[i1].AttachList[i2].InstanceValues.LoopbackId = new(Int64Custom)
						*data.VrfAttachments[i1].AttachList[i2].InstanceValues.LoopbackId = Int64Custom(ele2.LoopbackId.ValueInt64())
					} else {
						data.VrfAttachments[i1].AttachList[i2].InstanceValues.LoopbackId = nil
					}
					//-----inline nesting end----

					//-----inline nesting Start----
					if !ele2.LoopbackIpv4.IsNull() && !ele2.LoopbackIpv4.IsUnknown() {
						data.VrfAttachments[i1].AttachList[i2].InstanceValues.LoopbackIpv4 = ele2.LoopbackIpv4.ValueString()
					} else {
						data.VrfAttachments[i1].AttachList[i2].InstanceValues.LoopbackIpv4 = ""
					}
					//-----inline nesting end----

					//-----inline nesting Start----
					if !ele2.LoopbackIpv6.IsNull() && !ele2.LoopbackIpv6.IsUnknown() {
						data.VrfAttachments[i1].AttachList[i2].InstanceValues.LoopbackIpv6 = ele2.LoopbackIpv6.ValueString()
					} else {
						data.VrfAttachments[i1].AttachList[i2].InstanceValues.LoopbackIpv6 = ""
					}
					//-----inline nesting end----

				}
			}

		}
	}

	return data
}
