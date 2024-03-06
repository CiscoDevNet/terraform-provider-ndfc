package resource_vrf_attachments

import (
	"encoding/json"
	"log"
	"strconv"
)

type CustomNDFCInstanceValue NDFCInstanceValuesValue

func (v *NDFCInstanceValuesValue) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	if string(data) == "" || string(data) == "\"\"" {
		return nil
	}

	//valueMap := make(map[string]interface{})
	dataString, err := strconv.Unquote(string(data))
	if err != nil {

		log.Println("String conversion error", dataString, err)
		return err
	}

	newT := CustomNDFCInstanceValue{}

	err = json.Unmarshal([]byte(dataString), &newT)
	if err != nil {
		log.Println("unmarshall err", dataString, err)
		return err
	}

	v.LoopbackId = newT.LoopbackId
	v.LoopbackIpv4 = newT.LoopbackIpv4
	v.LoopbackIpv6 = newT.LoopbackIpv6

	return nil

}

func (v NDFCInstanceValuesValue) MarshalJSON() ([]byte, error) {
	if v.LoopbackId == nil && v.LoopbackIpv4 == "" && v.LoopbackIpv6 == "" {
		return []byte("\"\""), nil
	}
	retBytes, err :=  json.Marshal(&CustomNDFCInstanceValue{
		LoopbackId:   v.LoopbackId,
		LoopbackIpv4: v.LoopbackIpv4,
		LoopbackIpv6: v.LoopbackIpv6,
	})
	if err != nil {
		return []byte("\"\""), err
	}
	ret := strconv.Quote(string(retBytes))
	return []byte(ret), nil
}

type NDFCVrfAttachments struct {
	Attachments []NDFCVrfAttachmentsValue `json:"attachments"`
}

type NDFCVrfAttachmentsPayload struct {
	VrfName    string                `json:"vrfName"`
	AttachList []NDFCAttachListValue `json:"lanAttachList"`
}

type DeploymentState struct {
	State string
	Seen  bool
}

type NDFCVrfAttachmentsPayloads struct {
	GlobalDeploy   bool
	GlobalUndeploy bool
	FabricName     string
	VrfAttachments []NDFCVrfAttachmentsPayload // Attachment payload for NDFC
	DepMap         map[string][]string         // use for backfilling DeploymentFlag in TF state        // for deployment
}

func (v *NDFCVrfAttachmentsModel) FillVrfAttachmentsFromPayload(payload *NDFCVrfAttachmentsPayloads) {
	v.VrfAttachments = make(map[string]NDFCVrfAttachmentsValue)
	for i := range (*payload).VrfAttachments {
		vrfName := (*payload).VrfAttachments[i].VrfName
		vrfAttachEntry := NDFCVrfAttachmentsValue{}
		vrfAttachEntry.AttachList = make(map[string]NDFCAttachListValue)
		for j := range (*payload).VrfAttachments[i].AttachList {
			vrfAttachEntry.AttachList[(*payload).VrfAttachments[i].AttachList[j].SwitchSerialNo] = (*payload).VrfAttachments[i].AttachList[j]
		}
		v.VrfAttachments[vrfName] = vrfAttachEntry
	}
}

func (p *NDFCVrfAttachmentsPayloads) AddEntry(vrfName string, attachList []NDFCAttachListValue) {
	if len(attachList) == 0 {
		return
	}
	vrfAttachEntry := NDFCVrfAttachmentsPayload{}
	vrfAttachEntry.VrfName = vrfName
	vrfAttachEntry.AttachList = attachList
	p.VrfAttachments = append(p.VrfAttachments, vrfAttachEntry)
}
