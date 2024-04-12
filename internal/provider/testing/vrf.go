package testing

import (
	"log"
	"strconv"
	"strings"
	"terraform-provider-ndfc/internal/provider/resources/resource_vrf_attachments"
	"terraform-provider-ndfc/internal/provider/resources/resource_vrf_bulk"
	"terraform-provider-ndfc/internal/provider/types"
)

var tmpDir string

func GenerateVrfBulkObject(bulk **resource_vrf_bulk.NDFCVrfBulkModel, fabric string, vrfCount int,
	globaldeploy, vrf_deploy, deployNeeded bool, serials []string) {
	vrfBulk := new(resource_vrf_bulk.NDFCVrfBulkModel)
	vrfBulk.FabricName = fabric
	vrfBulk.DeployAllAttachments = globaldeploy
	vrfBulk.Vrfs = make(map[string]resource_vrf_bulk.NDFCVrfsValue)
	log.Printf("Creating Bulk VRF object VRF Count: %d", vrfCount)

	for i := 0; i < vrfCount; i++ {
		vrfName := GetConfig().NDFC.VrfPrefix + strconv.Itoa(i+1)
		log.Printf("Creating VRF: %s", vrfName)
		vrf := resource_vrf_bulk.NDFCVrfsValue{}
		//vrf.VrfTemplateConfig.VrfDescription = "VRF Description"
		//vrf.VrfName = vrfName
		vrf.DeployAttachments = vrf_deploy
		if len(serials) > 0 {
			vrf.AttachList = make(map[string]resource_vrf_attachments.NDFCAttachListValue)
			for j := range serials {
				attach := resource_vrf_attachments.NDFCAttachListValue{}
				attach.SerialNumber = serials[j]
				attach.DeployThisAttachment = deployNeeded
				vrf.AttachList[serials[j]] = attach
			}
		} else {
			vrf.AttachList = nil
		}
		vrfBulk.Vrfs[vrfName] = vrf
	}
	*bulk = vrfBulk
}

func GenerateSingleVrfObject(vrfDptr **resource_vrf_bulk.NDFCVrfBulkModel, namePrefix, fabric string, vrfNo int, globaldeploy, vrf_deploy, deployNeeded bool, serials []string) {
	vrfBulk := new(resource_vrf_bulk.NDFCVrfBulkModel)
	vrfBulk.FabricName = fabric
	vrfBulk.DeployAllAttachments = globaldeploy
	vrfBulk.Vrfs = make(map[string]resource_vrf_bulk.NDFCVrfsValue)
	vrfName := namePrefix + strconv.Itoa(vrfNo)
	vrf := resource_vrf_bulk.NDFCVrfsValue{}

	vrf.DeployAttachments = vrf_deploy
	if len(serials) > 0 {
		vrf.AttachList = make(map[string]resource_vrf_attachments.NDFCAttachListValue)
		for j := range serials {
			attach := resource_vrf_attachments.NDFCAttachListValue{}
			attach.SerialNumber = serials[j]
			attach.DeployThisAttachment = deployNeeded
			vrf.AttachList[serials[j]] = attach
		}
	} else {
		vrf.AttachList = nil
	}
	vrfBulk.Vrfs[vrfName] = vrf
	*vrfDptr = vrfBulk
}

func ModifyVrfBulkObject(vrfs **resource_vrf_bulk.NDFCVrfBulkModel, vrfNo int, values map[string]interface{}) {
	vrfBulk := *vrfs
	vrfName := GetConfig().NDFC.VrfPrefix + strconv.Itoa(vrfNo)
	vrf, ok := vrfBulk.Vrfs[vrfName]
	if !ok {
		for key, value := range values {
			switch key {
			case "vrf_description":
				vrf.VrfTemplateConfig.VrfDescription = value.(string)
			case "vlan_id":
				vrf.VrfTemplateConfig.VlanId = new(types.Int64Custom)
				*vrf.VrfTemplateConfig.VlanId = types.Int64Custom((value.(int)))
			case "loopback_routing_tag":
				vrf.VrfTemplateConfig.LoopbackRoutingTag = new(int64)
				*vrf.VrfTemplateConfig.LoopbackRoutingTag = int64(value.(int))
			case "mtu":
				vrf.VrfTemplateConfig.Mtu = new(int64)
				*vrf.VrfTemplateConfig.Mtu = int64(value.(int))
			case "max_bgp_paths":
				vrf.VrfTemplateConfig.MaxBgpPaths = new(int64)
				*vrf.VrfTemplateConfig.MaxBgpPaths = int64(value.(int))
			case "max_ibgp_paths":
				vrf.VrfTemplateConfig.MaxIbgpPaths = new(int64)
				*vrf.VrfTemplateConfig.MaxIbgpPaths = int64(value.(int))
			case "ipv6_link_local":
				vrf.VrfTemplateConfig.Ipv6LinkLocal = value.(string)
			case "vrf_id":
				vrf.VrfId = new(int64)
				*vrf.VrfId = int64(value.(int))
			}
		}
	}

}

func IncreaseVrfCount(vrf **resource_vrf_bulk.NDFCVrfBulkModel, vrfToAdd int,
	globaldeploy, vrf_deploy, deployNeeded bool, serials []string) {
	vrfBulk := *vrf
	vrfBulk.DeployAllAttachments = globaldeploy
	log.Printf("Add more VRFs  VRF Count: %d", vrfToAdd)

	currentCount := len(vrfBulk.Vrfs)
	for i := 0; i < vrfToAdd; i++ {
		vrfName := GetConfig().NDFC.VrfPrefix + strconv.Itoa(currentCount+i+1)
		log.Printf("Creating VRF: %s", vrfName)
		vrf := resource_vrf_bulk.NDFCVrfsValue{}
		//vrf.VrfTemplateConfig.VrfDescription = "VRF Description"
		//vrf.VrfName = vrfName
		vrf.DeployAttachments = vrf_deploy
		if len(serials) > 0 {
			vrf.AttachList = make(map[string]resource_vrf_attachments.NDFCAttachListValue)
			for j := range serials {
				attach := resource_vrf_attachments.NDFCAttachListValue{}
				attach.SerialNumber = serials[j]
				attach.DeployThisAttachment = deployNeeded
				vrf.AttachList[serials[j]] = attach
			}
		}
		vrfBulk.Vrfs[vrfName] = vrf
	}
}

func AddAttachments(vrfBulk *resource_vrf_bulk.NDFCVrfBulkModel, serials []string, deployNeeded bool, start, end int) *resource_vrf_bulk.NDFCVrfBulkModel {
	for vrfName, vrf := range vrfBulk.Vrfs {
		//vrf_acc_<id>
		id, err := strconv.Atoi(strings.Split(vrfName, "_")[2])
		if err != nil {
			panic(err)
		}

		if id >= start && id <= end {
			if len(serials) > 0 {
				vrf.AttachList = make(map[string]resource_vrf_attachments.NDFCAttachListValue)
				for j := range serials {
					attach := resource_vrf_attachments.NDFCAttachListValue{}
					attach.SerialNumber = serials[j]
					attach.DeployThisAttachment = deployNeeded
					vrf.AttachList[serials[j]] = attach
				}
			}
			vrfBulk.Vrfs[vrfName] = vrf
		}
	}

	return vrfBulk
}

// Delete attachCount attachments in vrfCount VRFs, if attachCount is -1, delete all attachments
func DeleteAttachments(vrfBulk *resource_vrf_bulk.NDFCVrfBulkModel, vrfCount int, attachCount int) {
	for vrfName, vrf := range vrfBulk.Vrfs {
		if vrfCount > 0 {
			for serial := range vrf.AttachList {
				if attachCount == -1 {
					vrf.AttachList = nil
					break
				}
				if attachCount > 0 {
					delete(vrf.AttachList, serial)
					attachCount--
				}
			}
			vrfCount--
			vrfBulk.Vrfs[vrfName] = vrf
		} else {
			break
		}
	}
}

func DeleteVrfs(vrf **resource_vrf_bulk.NDFCVrfBulkModel, start, end int) {
	vrfBulk := *vrf
	log.Printf("Delete VRFs: %d to %d", start, end)
	for vrfName := range vrfBulk.Vrfs {
		ids := strings.Split(vrfName, "_")
		id, err := strconv.Atoi(ids[len(ids)-1])
		if err != nil {
			panic(err)
		}
		if id >= start && id <= end {
			delete(vrfBulk.Vrfs, vrfName)
		}
	}
}
