package datasource_networks

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
)

/*
  Payload format NDFC contains a nested dhcpserver array inside map, both named dhcpServers
  This is difficult to represent in the model struct autogeneration templates
  Hence we need overrides for marshalling and unmarshalling

 "dhcpServers": {
	"dhcpServers": [
		{
			"srvrAddr": "10.1.1.1",
			"srvrVrf": "management"
		}
	]
},
*/

type CustomNDFCDhcpRelayServersValues NDFCDhcpRelayServersValues

type NwAttachmentOperation uint8

const (
	NwAttachmentAttach NwAttachmentOperation = iota
	NwAttachmentDetach
	NwAttachmentNone
)

type NDFCDHCPRelayInnerPayload struct {
	ServersInner CustomNDFCDhcpRelayServersValues `json:"dhcpServers"`
}

func (v *NDFCDhcpRelayServersValues) UnmarshalJSON(data []byte) error {
	log.Printf("Unmarshalling NDFCDhcpRelayServersValues")
	if string(data) == "null" {
		return nil
	}
	if string(data) == "" || string(data) == "\"\"" {
		return nil
	}
	//valueMap := make(map[string]interface{})
	dataString, err := strconv.Unquote(string(data))
	if err != nil {
		log.Printf("String conversion error %s %s", string(data), err)
		dataString = string(data)
		err = nil
	}
	newServer := NDFCDHCPRelayInnerPayload{}
	if strings.HasPrefix(dataString, "{") {
		log.Printf("Got a map: %s", dataString)
		// This is a map
		err = json.Unmarshal([]byte(dataString), &newServer)
		if err != nil {
			log.Println("Error unmarshalling", dataString, err)
			return err
		}
	} else if strings.HasPrefix(dataString, "[") {
		log.Printf("Got an array: %s", dataString)
		// This is an array
		err = json.Unmarshal([]byte(dataString), &newServer.ServersInner)
		if err != nil {
			log.Println("Error unmarshalling", dataString, err)
			return err
		}
	}
	for _, vv := range newServer.ServersInner {
		xx := NDFCDhcpRelayServersValue{}
		xx.Address = vv.Address
		xx.Vrf = vv.Vrf
		*v = append(*v, xx)
	}
	return nil
}

func (v NDFCDhcpRelayServersValues) MarshalJSON() ([]byte, error) {
	log.Printf("Marshalling NDFCDhcpRelayServersValues")
	if len(v) == 0 {
		return []byte("\"\""), nil
	}
	relay := NDFCDHCPRelayInnerPayload{}

	for _, vv := range v {
		relay.ServersInner = append(relay.ServersInner, vv)
	}
	retBytes, err := json.Marshal(&relay)

	if err != nil {
		log.Printf("Marshalling error %s", err)
		return []byte("\"\""), err
	}
	return retBytes, nil
}
