package resource_vpc_pair

import (
	"encoding/json"
)

type NDFCVpcPairRecommendations struct {
	SerialNumber         string  `json:"serialNumber"`
	RecommendationReason string  `json:"recommendationReason"`
    LogicalName 		 string  `json:"logicalName"`
	UseVirtualPeerlink   bool    `json:"useVirtualPeerlink"`
	Recommended          bool    `json:"recommended"`
}

type CustomNDFCVpcPairModel NDFCVpcPairModel

func (m *NDFCVpcPairModel) UnmarshalJSON(data []byte) error {
	var customModel CustomNDFCVpcPairModel
	if err := json.Unmarshal(data, &customModel); err != nil {
		return err
	}
	m.SerialNumbers = []string{customModel.PeerOneId, customModel.PeerTwoId}
	m.UseVirtualPeerlink = customModel.UseVirtualPeerlink
	m.PeerOneId = customModel.PeerOneId
	m.PeerTwoId = customModel.PeerTwoId
	return nil
}
func (m *NDFCVpcPairModel) MarshalJSON() ([]byte, error) {
	var customModel CustomNDFCVpcPairModel
	customModel.PeerOneId = m.SerialNumbers[0]
	customModel.PeerTwoId = m.SerialNumbers[1]
	customModel.UseVirtualPeerlink = m.UseVirtualPeerlink
	customModel.SerialNumbers = m.SerialNumbers
	return json.Marshal(customModel)
}
