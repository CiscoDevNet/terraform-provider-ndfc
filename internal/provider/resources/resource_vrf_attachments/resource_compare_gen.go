package resource_vrf_attachments

import (
	"log"
	. "terraform-provider-ndfc/internal/provider/types"
)

func (v NDFCAttachListValue) DeepEqual(c NDFCAttachListValue) int {
	controlFlagUpdate := false
	if v.SerialNumber != c.SerialNumber {
		log.Printf("v.SerialNumber=%v, c.SerialNumber=%v", v.SerialNumber, c.SerialNumber)
		return RequiresReplace
	}

	if v.Vlan != nil && c.Vlan != nil {
		if *v.Vlan != *c.Vlan {
			log.Printf("v.Vlan=%v, c.Vlan=%v", *v.Vlan, *c.Vlan)
			return RequiresUpdate
		}
	} else {
		if v.Vlan != nil {
			log.Printf("v.Vlan=%v", *v.Vlan)
			return RequiresUpdate
		} else if c.Vlan != nil {
			log.Printf("c.Vlan=%v", *c.Vlan)
			return RequiresUpdate
		}
	}
	if v.FreeformConfig != c.FreeformConfig {
		log.Printf("v.FreeformConfig=%v, c.FreeformConfig=%v", v.FreeformConfig, c.FreeformConfig)
		return RequiresUpdate
	}
	if v.DeployThisAttachment != c.DeployThisAttachment {
		log.Printf("v.DeployThisAttachment=%v, c.DeployThisAttachment=%v", v.DeployThisAttachment, c.DeployThisAttachment)
		controlFlagUpdate = true
	}

	if v.InstanceValues.LoopbackId != nil && c.InstanceValues.LoopbackId != nil {
		if *v.InstanceValues.LoopbackId != *c.InstanceValues.LoopbackId {
			log.Printf("v.InstanceValues.LoopbackId=%v, c.InstanceValues.LoopbackId=%v", *v.InstanceValues.LoopbackId, *c.InstanceValues.LoopbackId)
			return RequiresUpdate
		}
	} else {
		if v.InstanceValues.LoopbackId != nil {
			log.Printf("v.InstanceValues.LoopbackId=%v", *v.InstanceValues.LoopbackId)
			return RequiresUpdate
		} else if c.InstanceValues.LoopbackId != nil {
			log.Printf("c.InstanceValues.LoopbackId=%v", *c.InstanceValues.LoopbackId)
			return RequiresUpdate
		}
	}
	if v.InstanceValues.LoopbackIpv4 != c.InstanceValues.LoopbackIpv4 {
		log.Printf("v.InstanceValues.LoopbackIpv4=%s, c.InstanceValues.LoopbackIpv4=%s", v.InstanceValues.LoopbackIpv4, c.InstanceValues.LoopbackIpv4)
		return RequiresUpdate
	}
	if v.InstanceValues.LoopbackIpv6 != c.InstanceValues.LoopbackIpv6 {
		log.Printf("v.InstanceValues.LoopbackIpv6=%s, c.InstanceValues.LoopbackIpv6=%s", v.InstanceValues.LoopbackIpv6, c.InstanceValues.LoopbackIpv6)
		return RequiresUpdate
	}

	if controlFlagUpdate {
		return ControlFlagUpdate
	}
	return ValuesDeeplyEqual
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

func (v *NDFCAttachListValue) CreatePlan(c NDFCAttachListValue) int {
	action := ActionNone
	controlFlagUpdate := false
	if v.SerialNumber != "" {

		if v.SerialNumber != c.SerialNumber {
			log.Printf("Config value in Plan: v.SerialNumber=%v, c.SerialNumber=%v", v.SerialNumber, c.SerialNumber)
			if action == ActionNone || action == RequiresUpdate {

				action = RequiresReplace
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("State value in Plan: v.SerialNumber=%v, c.SerialNumber=%v", v.SerialNumber, c.SerialNumber)
		v.SerialNumber = c.SerialNumber
	}

	if v.Vlan != nil {
		if c.Vlan != nil && (*v.Vlan != *c.Vlan) {
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
			log.Printf("Config value in plan: v.Vlan=%v, c.Vlan=%v", *v.Vlan, *c.Vlan)
		}
	} else if c.Vlan != nil {
		v.Vlan = new(Int64Custom)
		log.Printf("State value in plan: v.Vlan=%v, c.Vlan=%v", *v.Vlan, *c.Vlan)
		*v.Vlan = *c.Vlan
	}
	if v.FreeformConfig != "" {

		if v.FreeformConfig != c.FreeformConfig {
			log.Printf("Config value in Plan: v.FreeformConfig=%v, c.FreeformConfig=%v", v.FreeformConfig, c.FreeformConfig)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("State value in Plan: v.FreeformConfig=%v, c.FreeformConfig=%v", v.FreeformConfig, c.FreeformConfig)
		v.FreeformConfig = c.FreeformConfig
	}

	if v.DeployThisAttachment != c.DeployThisAttachment {
		log.Printf("Config value in Plan: v.DeployThisAttachment=%v, c.DeployThisAttachment=%v", v.DeployThisAttachment, c.DeployThisAttachment)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.InstanceValues.LoopbackId != nil {
		if c.InstanceValues.LoopbackId != nil &&
			(*v.InstanceValues.LoopbackId != *c.InstanceValues.LoopbackId) {
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
			log.Printf("Config value in plan: v.InstanceValues.LoopbackId=%v, c.InstanceValues.LoopbackId=%v", *v.InstanceValues.LoopbackId, *c.InstanceValues.LoopbackId)
		}
	} else if c.InstanceValues.LoopbackId != nil {
		v.InstanceValues.LoopbackId = new(Int64Custom)
		log.Printf("State value in plan: v.InstanceValues.LoopbackId=%v, c.InstanceValues.LoopbackId=%v", *v.InstanceValues.LoopbackId, *c.InstanceValues.LoopbackId)
		*v.InstanceValues.LoopbackId = *c.InstanceValues.LoopbackId
	}

	if v.InstanceValues.LoopbackIpv4 != "" {
		if v.InstanceValues.LoopbackIpv4 != c.InstanceValues.LoopbackIpv4 {
			log.Printf("Config value in Plan: v.InstanceValues.LoopbackIpv4=%v, c.InstanceValues.LoopbackIpv4=%v", v.InstanceValues.LoopbackIpv4, c.InstanceValues.LoopbackIpv4)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("State value in Plan: v.InstanceValues.LoopbackIpv4=%v, c.InstanceValues.LoopbackIpv4=%v", v.InstanceValues.LoopbackIpv4, c.InstanceValues.LoopbackIpv4)
		v.InstanceValues.LoopbackIpv4 = c.InstanceValues.LoopbackIpv4
	}

	if v.InstanceValues.LoopbackIpv6 != "" {
		if v.InstanceValues.LoopbackIpv6 != c.InstanceValues.LoopbackIpv6 {
			log.Printf("Config value in Plan: v.InstanceValues.LoopbackIpv6=%v, c.InstanceValues.LoopbackIpv6=%v", v.InstanceValues.LoopbackIpv6, c.InstanceValues.LoopbackIpv6)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("State value in Plan: v.InstanceValues.LoopbackIpv6=%v, c.InstanceValues.LoopbackIpv6=%v", v.InstanceValues.LoopbackIpv6, c.InstanceValues.LoopbackIpv6)
		v.InstanceValues.LoopbackIpv6 = c.InstanceValues.LoopbackIpv6
	}

	if controlFlagUpdate {
		return ControlFlagUpdate
	}
	return action
}

func (v *NDFCVrfAttachmentsValue) CreatePlan(c NDFCVrfAttachmentsValue) int {
	action := ActionNone
	controlFlagUpdate := false
	if v.VrfName != "" {

		if v.VrfName != c.VrfName {
			log.Printf("Config value in Plan: v.VrfName=%v, c.VrfName=%v", v.VrfName, c.VrfName)
			if action == ActionNone || action == RequiresUpdate {

				action = RequiresReplace
			}
		}
	} else {
		//v empty, fill with c
		log.Printf("State value in Plan: v.VrfName=%v, c.VrfName=%v", v.VrfName, c.VrfName)
		v.VrfName = c.VrfName
	}

	if v.DeployAllAttachments != c.DeployAllAttachments {
		log.Printf("Config value in Plan: v.DeployAllAttachments=%v, c.DeployAllAttachments=%v", v.DeployAllAttachments, c.DeployAllAttachments)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	for i := range v.AttachList {
		retVal := v.AttachList[i].CreatePlan(c.AttachList[i])
		if retVal != ValuesDeeplyEqual {
			return retVal
		}
	}
	if controlFlagUpdate {
		return ControlFlagUpdate
	}
	return action
}
