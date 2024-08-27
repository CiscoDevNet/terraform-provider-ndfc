// Code generated;  DO NOT EDIT.

package resource_test_rsc

import (
	"log"
	. "terraform-provider-ndfc/internal/provider/types"
)

func (v NDFCParameterMapValue) DeepEqual(c NDFCParameterMapValue) int {
	cf := false
	if v.SerialNumber != c.SerialNumber {
		log.Printf("v.SerialNumber=%v, c.SerialNumber=%v", v.SerialNumber, c.SerialNumber)
		return RequiresUpdate
	}

	if cf {
		return ControlFlagUpdate
	}
	return ValuesDeeplyEqual
}

func (v *NDFCParameterMapValue) CreatePlan(c NDFCParameterMapValue, cf *bool) int {
	action := ActionNone

	if v.SerialNumber != "" {

		if v.SerialNumber != c.SerialNumber {
			log.Printf("Update: v.SerialNumber=%v, c.SerialNumber=%v", v.SerialNumber, c.SerialNumber)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}

	} else {
		//v empty, fill with c
		log.Printf("Copy from state: v.SerialNumber=%v, c.SerialNumber=%v", v.SerialNumber, c.SerialNumber)
		v.SerialNumber = c.SerialNumber
	}

	if len(v.CustomAttributesNestedMap) != len(c.CustomAttributesNestedMap) {
		log.Printf("Update: len(v.CustomAttributesNestedMap)=%d, len(c.CustomAttributesNestedMap)=%d", len(v.CustomAttributesNestedMap), len(c.CustomAttributesNestedMap))
		return RequiresUpdate
	}
	for kk, vv := range v.CustomAttributesNestedMap {
		cc, ok := c.CustomAttributesNestedMap[kk]
		if !ok {
			log.Printf("Update: v.CustomAttributesNestedMap[%s]=%s, c.CustomAttributesNestedMap[%s]=nil", kk, vv, kk)
			return RequiresUpdate
		}
		if vv != cc {
			log.Printf("Update: v.CustomAttributesNestedMap[%s]=%s, c.CustomAttributesNestedMap[%s]=%s", kk, vv, kk, cc)
			return RequiresUpdate
		}
	}

	return action
}
