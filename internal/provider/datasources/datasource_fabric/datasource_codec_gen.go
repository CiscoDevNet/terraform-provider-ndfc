// Code generated DO NOT EDIT.
package datasource_fabric

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCFabricModel struct {
	Fabrics NDFCFabricsValues `json:"fabrics,omitempty"`
}

type NDFCFabricsValues []NDFCFabricsValue

type NDFCFabricsValue struct {
	FabricId         string `json:"fabricId,omitempty"`
	FabricName       string `json:"fabricName,omitempty"`
	FabricType       string `json:"fabricType,omitempty"`
	FabricTechnology string `json:"fabricTechnology,omitempty"`
	ProvisionMode    string `json:"provisionMode,omitempty"`
	DeviceType       string `json:"deviceType,omitempty"`
	AsNumber         string `json:"asn,omitempty"`
	SiteId           string `json:"siteId,omitempty"`
}

func (v *FabricModel) SetModelData(jsonData *NDFCFabricModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	listData := make([]FabricsValue, 0)
	for _, item := range jsonData.Fabrics {
		data := new(FabricsValue)
		err = data.SetValue(&item)
		if err != nil {
			return err
		}
		data.state = attr.ValueStateKnown
		listData = append(listData, *data)
	}
	v.Fabrics, err = types.ListValueFrom(context.Background(), FabricsValue{}.Type(context.Background()), listData)
	if err != nil {
		return err
	}

	return err
}

func (v *FabricsValue) SetValue(jsonData *NDFCFabricsValue) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.FabricId != "" {
		v.FabricId = types.StringValue(jsonData.FabricId)
	} else {
		v.FabricId = types.StringNull()
	}

	if jsonData.FabricName != "" {
		v.FabricName = types.StringValue(jsonData.FabricName)
	} else {
		v.FabricName = types.StringNull()
	}

	if jsonData.FabricType != "" {
		v.FabricType = types.StringValue(jsonData.FabricType)
	} else {
		v.FabricType = types.StringNull()
	}

	if jsonData.FabricTechnology != "" {
		v.FabricTechnology = types.StringValue(jsonData.FabricTechnology)
	} else {
		v.FabricTechnology = types.StringNull()
	}

	if jsonData.ProvisionMode != "" {
		v.ProvisionMode = types.StringValue(jsonData.ProvisionMode)
	} else {
		v.ProvisionMode = types.StringNull()
	}

	if jsonData.DeviceType != "" {
		v.DeviceType = types.StringValue(jsonData.DeviceType)
	} else {
		v.DeviceType = types.StringNull()
	}

	if jsonData.AsNumber != "" {
		v.AsNumber = types.StringValue(jsonData.AsNumber)
	} else {
		v.AsNumber = types.StringNull()
	}

	if jsonData.SiteId != "" {
		v.SiteId = types.StringValue(jsonData.SiteId)
	} else {
		v.SiteId = types.StringNull()
	}

	return err
}

func (v FabricModel) GetModelData() *NDFCFabricModel {
	var data = new(NDFCFabricModel)

	if !v.Fabrics.IsNull() && !v.Fabrics.IsUnknown() {
		elements := make([]FabricsValue, len(v.Fabrics.Elements()))
		data.Fabrics = make([]NDFCFabricsValue, len(v.Fabrics.Elements()))
		diag := v.Fabrics.ElementsAs(context.Background(), &elements, false)
		if diag != nil {
			panic(diag)
		}
		for i, ele := range elements {
			if !ele.FabricId.IsNull() && !ele.FabricId.IsUnknown() {
				data.Fabrics[i].FabricId = ele.FabricId.ValueString()
			} else {
				data.Fabrics[i].FabricId = ""
			}

			if !ele.FabricName.IsNull() && !ele.FabricName.IsUnknown() {
				data.Fabrics[i].FabricName = ele.FabricName.ValueString()
			} else {
				data.Fabrics[i].FabricName = ""
			}

			if !ele.FabricType.IsNull() && !ele.FabricType.IsUnknown() {
				data.Fabrics[i].FabricType = ele.FabricType.ValueString()
			} else {
				data.Fabrics[i].FabricType = ""
			}

			if !ele.FabricTechnology.IsNull() && !ele.FabricTechnology.IsUnknown() {
				data.Fabrics[i].FabricTechnology = ele.FabricTechnology.ValueString()
			} else {
				data.Fabrics[i].FabricTechnology = ""
			}

			if !ele.ProvisionMode.IsNull() && !ele.ProvisionMode.IsUnknown() {
				data.Fabrics[i].ProvisionMode = ele.ProvisionMode.ValueString()
			} else {
				data.Fabrics[i].ProvisionMode = ""
			}

			if !ele.DeviceType.IsNull() && !ele.DeviceType.IsUnknown() {
				data.Fabrics[i].DeviceType = ele.DeviceType.ValueString()
			} else {
				data.Fabrics[i].DeviceType = ""
			}

			if !ele.AsNumber.IsNull() && !ele.AsNumber.IsUnknown() {
				data.Fabrics[i].AsNumber = ele.AsNumber.ValueString()
			} else {
				data.Fabrics[i].AsNumber = ""
			}

			if !ele.SiteId.IsNull() && !ele.SiteId.IsUnknown() {
				data.Fabrics[i].SiteId = ele.SiteId.ValueString()
			} else {
				data.Fabrics[i].SiteId = ""
			}

		}
	}

	return data
}
