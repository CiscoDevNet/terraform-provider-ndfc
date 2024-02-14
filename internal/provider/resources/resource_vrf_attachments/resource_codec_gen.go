// Code generated DO NOT EDIT.
package resource_vrf_attachments

import (
	. "terraform-provider-ndfc/internal/provider/types"
)

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
