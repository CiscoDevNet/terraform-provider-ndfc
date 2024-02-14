package resource_vrf_bulk

import "encoding/json"

type CustomCNDFCVrfsValue NDFCVrfsValue

// Skip Attachlist in payload
func (v NDFCVrfsValue) MarshalJSON() ([]byte, error) {
	v1 := CustomCNDFCVrfsValue{}
	v1.Id = v.Id
	v1.VrfName = v.VrfName
	v1.VrfId = v.VrfId
	v1.FabricName = v.FabricName
	v1.VrfTemplate = v.VrfTemplate
	v1.VrfExtensionTemplate = v.VrfExtensionTemplate
	v1.VrfStatus = v.VrfStatus
	v1.VrfTemplateConfig = v.VrfTemplateConfig
	v1.AttachList = nil //v.AttachList
	return json.Marshal(&v1)
}
