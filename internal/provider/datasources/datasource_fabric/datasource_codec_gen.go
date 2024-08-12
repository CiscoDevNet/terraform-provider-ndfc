// Code generated;  DO NOT EDIT.

package datasource_fabric

import (
	"context"
	"log"

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

	if len(jsonData.Fabrics) == 0 {
		log.Printf("v.Fabrics is empty")
		v.Fabrics = types.ListNull(FabricsValue{}.Type(context.Background()))
	} else {
		log.Printf("v.Fabrics contains %d elements", len(jsonData.Fabrics))
		listData := make([]FabricsValue, 0)
		for _, item := range jsonData.Fabrics {
			data := new(FabricsValue)
			err = data.SetValue(&item)
			if err != nil {
				log.Printf("Error in FabricsValue.SetValue")
				return err
			}
			data.state = attr.ValueStateKnown
			listData = append(listData, *data)
		}
		v.Fabrics, err = types.ListValueFrom(context.Background(), FabricsValue{}.Type(context.Background()), listData)
		if err != nil {
			log.Printf("Error in converting []FabricsValue to  List")
			return err
		}
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

	//MARSHAL_BODY

	//MARSHALL_LIST

	if !v.Fabrics.IsNull() && !v.Fabrics.IsUnknown() {
		elements := make([]FabricsValue, len(v.Fabrics.Elements()))
		data.Fabrics = make([]NDFCFabricsValue, len(v.Fabrics.Elements()))

		diag := v.Fabrics.ElementsAs(context.Background(), &elements, false)
		if diag != nil {
			panic(diag)
		}
		for i1, ele1 := range elements {
			if !ele1.FabricId.IsNull() && !ele1.FabricId.IsUnknown() {

				data.Fabrics[i1].FabricId = ele1.FabricId.ValueString()
			} else {
				data.Fabrics[i1].FabricId = ""
			}

			if !ele1.FabricName.IsNull() && !ele1.FabricName.IsUnknown() {

				data.Fabrics[i1].FabricName = ele1.FabricName.ValueString()
			} else {
				data.Fabrics[i1].FabricName = ""
			}

			if !ele1.FabricType.IsNull() && !ele1.FabricType.IsUnknown() {

				data.Fabrics[i1].FabricType = ele1.FabricType.ValueString()
			} else {
				data.Fabrics[i1].FabricType = ""
			}

			if !ele1.FabricTechnology.IsNull() && !ele1.FabricTechnology.IsUnknown() {

				data.Fabrics[i1].FabricTechnology = ele1.FabricTechnology.ValueString()
			} else {
				data.Fabrics[i1].FabricTechnology = ""
			}

			if !ele1.ProvisionMode.IsNull() && !ele1.ProvisionMode.IsUnknown() {

				data.Fabrics[i1].ProvisionMode = ele1.ProvisionMode.ValueString()
			} else {
				data.Fabrics[i1].ProvisionMode = ""
			}

			if !ele1.DeviceType.IsNull() && !ele1.DeviceType.IsUnknown() {

				data.Fabrics[i1].DeviceType = ele1.DeviceType.ValueString()
			} else {
				data.Fabrics[i1].DeviceType = ""
			}

			if !ele1.AsNumber.IsNull() && !ele1.AsNumber.IsUnknown() {

				data.Fabrics[i1].AsNumber = ele1.AsNumber.ValueString()
			} else {
				data.Fabrics[i1].AsNumber = ""
			}

			if !ele1.SiteId.IsNull() && !ele1.SiteId.IsUnknown() {

				data.Fabrics[i1].SiteId = ele1.SiteId.ValueString()
			} else {
				data.Fabrics[i1].SiteId = ""
			}

		}
	}

	return data
}
