package testing

import (
	"strconv"
	"terraform-provider-ndfc/internal/provider/resources/resource_vrf_attachments"
	"terraform-provider-ndfc/internal/provider/resources/resource_vrf_bulk"
	"terraform-provider-ndfc/internal/provider/types"
)

func VrfAttachmentsMod(vrfs **resource_vrf_bulk.NDFCVrfBulkModel, start, end int, serials []string, serial string, x map[string]interface{}) {
	vrfBulk := *vrfs
	var vrf []string
	var vrfName string
	for i := start; i <= end; i++ {
		vrfName = GetConfig("network").NDFC.VrfPrefix + strconv.Itoa(i)
		_, ok := vrfBulk.Vrfs[vrfName]
		if !ok {
			panic("VRF not found")
		}
		vrf = append(vrf, vrfName)
	}

	if len(serials) > 0 {
		for i := range vrf {
			vrfName := vrf[i]
			vrfEntry, ok := vrfBulk.Vrfs[vrfName]
			if !ok {
				panic("VRF not found")
			}
			vrfEntry.AttachList = make(map[string]resource_vrf_attachments.NDFCAttachListValue)
			for j := range serials {
				attach := resource_vrf_attachments.NDFCAttachListValue{}
				attach.SerialNumber = serials[j]
				//attach.DeployThisAttachment = true
				vrfEntry.AttachList[serials[j]] = attach
			}
			vrfBulk.Vrfs[vrfName] = vrfEntry
		}
	} else {
		for i := range vrf {
			vrfName := vrf[i]
			vrfEntry, ok := vrfBulk.Vrfs[vrfName]
			if !ok {
				panic("VRF not found")
			}
			vrfEntry.AttachList = nil
			vrfBulk.Vrfs[vrfName] = vrfEntry
		}
	}

	if serial != "" {
		if len(vrf) != 1 {
			panic("Only one attachment entry can be modified at a time")
		}
		vrfName := vrf[0]
		vrfEntry, ok := vrfBulk.Vrfs[vrfName]
		if !ok {
			panic("VRF not found")
		}
		attachEntry, ok := vrfEntry.AttachList[serial]
		if !ok {
			panic("Serial not found in AttachList")
		}

		for key, value := range x {
			switch key {
			case "vlan":
				attachEntry.Vlan = new(types.Int64Custom)
				*attachEntry.Vlan = types.Int64Custom(value.(int))
			case "freeform_config":
				attachEntry.FreeformConfig = value.(string)
			case "loopback_id":
				attachEntry.InstanceValues.LoopbackId = new(types.Int64Custom)
				*attachEntry.InstanceValues.LoopbackId = types.Int64Custom(value.(int))
			case "loopback_ipv4":
				attachEntry.InstanceValues.LoopbackIpv4 = value.(string)
			case "loopback_ipv6":
				attachEntry.InstanceValues.LoopbackIpv6 = value.(string)
			}
		}
		vrfEntry.AttachList[serial] = attachEntry
		vrfBulk.Vrfs[vrfName] = vrfEntry
	}
	*vrfs = vrfBulk
}
