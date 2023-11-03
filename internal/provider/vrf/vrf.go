package vrf

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type VRF struct {
	VrfName              types.String `tfsdk:"vrf_name"`
	VrfTemplate          types.String `tfsdk:"vrf_template"`
	VrfExtensionTemplate types.String `tfsdk:"vrf_extension_template"`
	VrfTemplateConfig    types.String `tfsdk:"vrf_template_config"`
	Id                   types.Int64  `tfsdk:"id"`
	VrfId                types.Int64  `tfsdk:"vrf_id"`
	VrfStatus            types.String `tfsdk:"vrf_status"`
	Fabric               types.String `tfsdk:"fabric"`
}

type vrfJSON struct {
	VrfName              string `json:"vrfName,omitempty"`
	VrfTemplate          string `json:"vrfTemplate,omitempty"`
	VrfExtensionTemplate string `json:"vrfExtensionTemplate,omitempty"`
	VrfTemplateConfig    string `json:"vrfTemplateConfig,omitempty"`
	Id                   int64  `json:"id,omitempty"`
	VrfId                int64  `json:"vrfId,omitempty"`
	VrfStatus            string `json:"vrfStatus,omitempty"`
	Fabric               string `json:"fabric,omitempty"`
}

func NewVrfJSON(v VRF) *vrfJSON {
	return &vrfJSON{
		VrfName:              v.VrfName.ValueString(),
		VrfTemplate:          v.VrfTemplate.ValueString(),
		VrfExtensionTemplate: v.VrfExtensionTemplate.ValueString(),
		VrfTemplateConfig:    v.VrfTemplateConfig.ValueString(),
		Id:                   v.Id.ValueInt64(),
		VrfId:                v.Id.ValueInt64(),
		VrfStatus:            v.VrfStatus.ValueString(),
		Fabric:               v.Fabric.ValueString(),
	}

}
func (v VRF) MarshalJSON() ([]byte, error) {
	return json.Marshal(NewVrfJSON(v))
}

func (v *VRF) UnmarshalJSON(data []byte) error {
	var jv vrfJSON
	if err := json.Unmarshal(data, &jv); err != nil {
		return err
	}
	v.Fabric = types.StringValue(jv.Fabric)
	v.VrfName = types.StringValue(jv.VrfName)
	v.Id = types.Int64Value(jv.Id)
	v.VrfTemplate = types.StringValue(jv.VrfTemplate)
	v.VrfTemplateConfig = types.StringValue(jv.VrfTemplateConfig)
	v.VrfId = types.Int64Value(jv.VrfId)
	v.VrfExtensionTemplate = types.StringValue(jv.VrfExtensionTemplate)
	v.VrfStatus = types.StringValue(jv.VrfStatus)

	return nil
}
