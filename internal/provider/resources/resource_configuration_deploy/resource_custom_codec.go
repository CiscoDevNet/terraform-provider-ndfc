// This file is used to define the custom codec for the resource_configuration_deploy resource

package resource_configuration_deploy

import (
	"encoding/json"
)

type SwitchStatusDB struct {
	SerialNumMap map[string]SwitchStatus
}

type SwitchStatus struct {
	SerialNumber string `json:"serialNumber"`
	Status       string `json:"ccStatus"`
	SwitchName   string `json:"logicalName"`
}

func (m *SwitchStatusDB) UnmarshalJSON(data []byte) error {
	customModel := make([]SwitchStatus, 0)
	m.SerialNumMap = make(map[string]SwitchStatus)
	if err := json.Unmarshal(data, &customModel); err != nil {
		return err
	}
	for _, entry := range customModel {
		m.SerialNumMap[entry.SerialNumber] = entry
	}
	return nil
}
