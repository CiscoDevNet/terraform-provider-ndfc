// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_networks

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	rna "terraform-provider-ndfc/internal/provider/resources/resource_network_attachments"
	. "terraform-provider-ndfc/internal/provider/types"
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

func (v *NDFCNetworksModel) FillAttachmentsFromPayload(payload *rna.NDFCNetworkAttachments) {
	for _, vv := range payload.NetworkAttachments {
		nw := vv.NetworkName
		if nwEntry, ok := v.Networks[nw]; ok {
			if len(nwEntry.Attachments) == 0 {
				nwEntry.Attachments = make(map[string]rna.NDFCAttachmentsValue)
			}
			for _, attachEntry := range vv.Attachments {
				nwEntry.Attachments[attachEntry.SwitchSerialNo] = attachEntry
			}
			//put it back
			v.Networks[nw] = nwEntry
		}
	}
}

func (v *NDFCNetworksValue) GetAttachmentsNames() []string {
	names := make([]string, 0)
	for name := range v.Attachments {
		names = append(names, name)
	}
	return names
}

func (v *NDFCNetworksModel) FillAttachmentsPayloadFromModel(payload *rna.NDFCNetworkAttachments, op NwAttachmentOperation) {

	payload.NetworkAttachments = make([]rna.NDFCNetworkAttachmentsPayload, 0)
	for nwName, nwEntry := range v.Networks {
		nwPayload := new(rna.NDFCNetworkAttachmentsPayload)
		nwPayload.NetworkName = nwName
		nwPayload.Attachments = make([]rna.NDFCAttachmentsValue, 0)
		for serial, attachEntry := range nwEntry.Attachments {
			log.Printf("Adding attachment %s/%s - operation %v", nwName, serial, op)
			attachEntry.NetworkName = nwName
			attachEntry.FabricName = v.FabricName
			attachEntry.SerialNumber = serial
			switch op {
			case NwAttachmentAttach:
				log.Printf("[DEBUG] Adding attachment %s/%s - operation true", nwName, serial)
				attachEntry.Deployment = "true"
			case NwAttachmentDetach:
				log.Printf("[DEBUG] Adding attachment %s/%s - operation false", nwName, serial)
				attachEntry.Deployment = "false"
			}
			nwPayload.Attachments = append(nwPayload.Attachments, attachEntry)
			nwEntry.Attachments[serial] = attachEntry
		}
		if len(nwPayload.Attachments) > 0 {
			payload.NetworkAttachments = append(payload.NetworkAttachments, *nwPayload)
		}
		v.Networks[nwName] = nwEntry
	}
}

func (v NDFCNetworksValue) GetAttachmentValues(filters uint16, attach string) []rna.NDFCAttachmentsValue {
	log.Printf("GetAttachmentValues: %s %s", v.NetworkName, v.FabricName)
	update := func(a *rna.NDFCAttachmentsValue, serial string) {
		log.Printf("GetAttachmentValues: update %s %s", v.NetworkName, v.FabricName)
		a.FabricName = v.FabricName
		a.NetworkName = v.NetworkName
		a.SerialNumber = serial
		if a.Vlan == nil {
			a.Vlan = new(Int64Custom)
			*a.Vlan = Int64Custom(-1)
		}
		if attach != "" {
			a.Deployment = attach
		}
	}
	attachmentValues := make([]rna.NDFCAttachmentsValue, 0)
	for serial, attachEntry := range v.Attachments {
		if filters != 0 {
			if attachEntry.UpdateAction&filters != 0 {
				update(&attachEntry, serial)
				attachmentValues = append(attachmentValues, attachEntry)
			}
		} else {
			update(&attachEntry, serial)
			attachmentValues = append(attachmentValues, attachEntry)
		}
		log.Printf("Loop GetAttachmentValues: %s %s %s %s", v.NetworkName, v.FabricName, serial, attachEntry.Deployment)
	}
	return attachmentValues
}
