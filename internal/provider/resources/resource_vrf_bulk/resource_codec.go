package resource_vrf_bulk

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
	return json.Marshal(&CustomNDFCInstanceValue{
		LoopbackId:   v.LoopbackId,
		LoopbackIpv4: v.LoopbackIpv4,
		LoopbackIpv6: v.LoopbackIpv6,
	})
}

type CustomCNDFCVrfsValue NDFCVrfsValue

// Skip Attachlist in payload
func (v NDFCVrfsValue) MarshalJSON() ([]byte, error) {
	v1 := CustomCNDFCVrfsValue{}
	v1.Id = v.Id
	v1.VrfName = v.VrfName
	v1.VrfId = v.VrfId
	v1.FabricName = v.FabricName
	v1.VrfTemplate = v.VrfTemplate
	v1.VrfExtensionTemplate = v.VrfExtensionTemplate
	v1.VrfStatus = v.VrfStatus
	v1.VrfTemplateConfig = v.VrfTemplateConfig
	v1.AttachList = nil //v.AttachList
	return json.Marshal(&v1)
}

type NDFCVrfAttachments struct {
	Attachments []NDFCVrfAttachmentsValue `json:"attachments"`
}
