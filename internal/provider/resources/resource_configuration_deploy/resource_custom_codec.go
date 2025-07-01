// This file is used to define the custom codec for the resource_configuration_deploy resource

package resource_configuration_deploy

import (
	"encoding/json"
	"time"
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

// DeployResponse represents a collection of deployment results
type DeployResponses []DeployResponse

// DeployResponse represents a single deployment operation result for a switch
type DeployResponse struct {
	IPAddress            string              `json:"ipaddress"`
	SerialNumber         string              `json:"serialnumber"`
	EntityName           string              `json:"entityName"`
	EntityType           string              `json:"entityType"`
	SecondaryEntityType  *string             `json:"secondaryEntityType"`
	SecondaryEntityName  *string             `json:"secondaryEntityName"`
	SubmittedTime        string              `json:"submittedTime"` // Format: "2025-06-20 08:07:59.372"
	CompletedTime        string              `json:"completedTime"`
	Source               string              `json:"source"`
	Status               string              `json:"status"` // SUCCESS, FAILED, etc.
	StatusDescription    string              `json:"statusDescription"`
	User                 string              `json:"user"`
	HostName             string              `json:"hostName"`
	TicketID             *string             `json:"ticketId"`
	ConfigResponseList   []CommandResponse   `json:"configResponseList"`
}

// CommandResponse represents the result of executing a single command
type CommandResponse struct {
	Command    string    `json:"command"`
	Response   string    `json:"cliResp"`
	StrStatus  string    `json:"strStatus"` // SUCCESS, FAILED, NOT_EXECUTED
	TimeMillis int64     `json:"time"`      // Unix timestamp in milliseconds
	Status     string    `json:"status"`
}

// ParseDeployResponses parses a JSON string into a slice of DeployResponse objects
func ParseDeployResponses(data []byte) (DeployResponses, error) {
	var responses DeployResponses
	if err := json.Unmarshal(data, &responses); err != nil {
		return nil, err
	}
	return responses, nil
}

// GetFormattedTime returns the submitted time as a time.Time object
func (dr *DeployResponse) GetSubmittedTime() (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05.000", dr.SubmittedTime)
}

// GetFormattedCompletedTime returns the completed time as a time.Time object
func (dr *DeployResponse) GetCompletedTime() (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05.000", dr.CompletedTime)
}

// IsSuccessful returns whether the overall deployment was successful
func (dr *DeployResponse) IsSuccessful() bool {
	return dr.Status == "SUCCESS"
}

// GetFailedCommands returns a list of commands that failed
func (dr *DeployResponse) GetFailedCommands() []CommandResponse {
	var failed []CommandResponse
	for _, cmd := range dr.ConfigResponseList {
		if cmd.Status == "FAILED" {
			failed = append(failed, cmd)
		}
	}
	return failed
}
