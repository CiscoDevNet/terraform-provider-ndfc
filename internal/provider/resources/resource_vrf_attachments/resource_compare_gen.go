package resource_vrf_attachments

import (
	"log"
	. "terraform-provider-ndfc/internal/provider/types"
)

func (v NDFCAttachListValue) DeepEqual(c NDFCAttachListValue) int {
	cf := false

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
		cf = true
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

	if cf {
		return ControlFlagUpdate
	}
	return ValuesDeeplyEqual
}

func (v NDFCVrfAttachmentsValue) DeepEqual(c NDFCVrfAttachmentsValue) int {
	cf := false
	if v.VrfName != c.VrfName {
		log.Printf("v.VrfName=%v, c.VrfName=%v", v.VrfName, c.VrfName)
		return RequiresReplace
	}
	if v.DeployAllAttachments != c.DeployAllAttachments {
		log.Printf("v.DeployAllAttachments=%v, c.DeployAllAttachments=%v", v.DeployAllAttachments, c.DeployAllAttachments)
		cf = true
	}

	if cf {
		return ControlFlagUpdate
	}
	return ValuesDeeplyEqual
}

func (v *NDFCAttachListValue) CreatePlan(c NDFCAttachListValue) int {
	action := ActionNone
	controlFlagUpdate := false

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

	if controlFlagUpdate {
		return ControlFlagUpdate
	}
	return action
}
