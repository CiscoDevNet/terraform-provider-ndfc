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
	return json.Marshal(&CustomNDFCInstanceValue{
		LoopbackId:   v.LoopbackId,
		LoopbackIpv4: v.LoopbackIpv4,
		LoopbackIpv6: v.LoopbackIpv6,
	})
}

