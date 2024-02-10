package resource_vrf_bulk

import "log"

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

func (s NDFCVrfAttachmentsValues) Len() int {
	return len(s)
}

func (s NDFCVrfAttachmentsValues) Less(i, j int) bool {
	return (*s[i].Id < *s[j].Id)

}

func (s NDFCVrfAttachmentsValues) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (v *NDFCVrfAttachmentsValue) CreateSearchMap() {
	v.AttachListMap = make(map[string]*NDFCAttachListValue)
	for i := range v.AttachList {
		key := ""
		if v.AttachList[i].SerialNumber == "" {
			key = v.AttachList[i].SwitchSerialNo
		} else {
			key = v.AttachList[i].SerialNumber
		}
		log.Printf("NDFCVrfAttachmentsValue.CreateSearchMap: key=%s", key)
		v.AttachListMap[key] = &v.AttachList[i]
	}
}

func (v NDFCVrfAttachmentsValue) DeepEqual(c NDFCVrfAttachmentsValue) int {
	controlFlagUpdate := false
	if v.VrfName != c.VrfName {
		log.Printf("v.VrfName=%v, c.VrfName=%v", v.VrfName, c.VrfName)
		return RequiresReplace
	}
	if v.DeployAllAttachments != c.DeployAllAttachments {
		log.Printf("v.DeployAllAttachments=%v, c.DeployAllAttachments=%v", v.DeployAllAttachments, c.DeployAllAttachments)
		controlFlagUpdate = true
	}

	if len(v.AttachList) != len(c.AttachList) {
		log.Printf("len(v.AttachList)=%d, len(c.AttachList)=%d", len(v.AttachList), len(c.AttachList))
		return RequiresUpdate
	}
	for i := range v.AttachList {
		retVal := v.AttachList[i].DeepEqual(c.AttachList[i])
		if retVal != ValuesDeeplyEqual {
			return retVal
		}
	}
	if controlFlagUpdate {
		return ControlFlagUpdate
	}
	return ValuesDeeplyEqual
}

func (v *NDFCVrfAttachmentsModel) CreateSearchMap() {
	v.VrfAttachmentsMap = make(map[string]*NDFCVrfAttachmentsValue)
	for i := range v.VrfAttachments {
		key := ""
		key = v.VrfAttachments[i].VrfName
		log.Printf("NDFCVrfAttachmentsModel.CreateSearchMap: key=%s", key)
		v.VrfAttachmentsMap[key] = &v.VrfAttachments[i]
		v.VrfAttachments[i].CreateSearchMap()
	}
}
