package resource_vrf

import "log"

const (
	ValuesDeeplyEqual = iota
	RequiresReplace
	RequiresUpdate
)

func (v NDFCAttachmentsValue) DeepEqual(c NDFCAttachmentsValue) int {
	if v.SerialNumber != c.SerialNumber {
		log.Printf("v.SerialNumber=%s, c.SerialNumber=%s", v.SerialNumber, c.SerialNumber)
		return RequiresUpdate

	}

	if v.FreeformConfig != c.FreeformConfig {
		log.Printf("v.FreeformConfig=%s, c.FreeformConfig=%s", v.FreeformConfig, c.FreeformConfig)
		return RequiresUpdate

	}

	if v.LoopbackId != nil && c.LoopbackId != nil {
		if *v.LoopbackId != *c.LoopbackId {
			log.Printf("v.LoopbackId=%v, c.LoopbackId=%v", *v.LoopbackId, *c.LoopbackId)
			return RequiresUpdate

		}
	} else {
		if v.LoopbackId != nil {
			log.Printf("v.LoopbackId=%v", *v.LoopbackId)
			return RequiresUpdate

		} else if c.LoopbackId != nil {
			log.Printf("c.LoopbackId=%v", *c.LoopbackId)
			return RequiresUpdate

		}
	}

	if v.LoopbackIpv4 != c.LoopbackIpv4 {
		log.Printf("v.LoopbackIpv4=%s, c.LoopbackIpv4=%s", v.LoopbackIpv4, c.LoopbackIpv4)
		return RequiresUpdate

	}

	if v.LoopbackIpv6 != c.LoopbackIpv6 {
		log.Printf("v.LoopbackIpv6=%s, c.LoopbackIpv6=%s", v.LoopbackIpv6, c.LoopbackIpv6)
		return RequiresUpdate

	}

	return ValuesDeeplyEqual
}
