// Code generated;  DO NOT EDIT.

package resource_test_rsc

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCTestRscModel struct {
	Policy       string                           `json:"policy,omitempty"`
	ParameterMap map[string]NDFCParameterMapValue `json:"parameter_map,omitempty"`
}

type NDFCParameterMapValue struct {
	SerialNumber              string            `json:"serialNumber,omitempty"`
	CustomAttributesNestedMap map[string]string `json:"custom_attributes,omitempty"`
}

func (v *TestRscModel) SetModelData(jsonData *NDFCTestRscModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.Policy != "" {
		v.Policy = types.StringValue(jsonData.Policy)
	} else {
		v.Policy = types.StringNull()
	}

	if len(jsonData.ParameterMap) == 0 {
		log.Printf("v.ParameterMap is empty")
		v.ParameterMap = types.MapNull(ParameterMapValue{}.Type(context.Background()))
	} else {
		mapData := make(map[string]ParameterMapValue)
		for key, item := range jsonData.ParameterMap {
			data := new(ParameterMapValue)
			err = data.SetValue(&item)
			if err != nil {
				log.Printf("Error in ParameterMapValue.SetValue")
				return err
			}
			data.state = attr.ValueStateKnown
			mapData[key] = *data
		}
		v.ParameterMap, err = types.MapValueFrom(context.Background(), ParameterMapValue{}.Type(context.Background()), mapData)
		if err != nil {
			log.Printf("Error in converting map[string]ParameterMapValue to  Map")

		}
	}

	return err
}

func (v *ParameterMapValue) SetValue(jsonData *NDFCParameterMapValue) diag.Diagnostics {

	var err diag.Diagnostics
	err = nil

	if jsonData.SerialNumber != "" {
		v.SerialNumber = types.StringValue(jsonData.SerialNumber)
	} else {
		v.SerialNumber = types.StringNull()
	}

	if len(jsonData.CustomAttributesNestedMap) == 0 {
		log.Printf("v.CustomAttributesNestedMap is empty")
		v.CustomAttributesNestedMap = types.MapNull(types.StringType)
	} else {
		mapData := make(map[string]attr.Value)
		for key, item := range jsonData.CustomAttributesNestedMap {
			mapData[key] = types.StringValue(item)
		}
		v.CustomAttributesNestedMap, err = types.MapValue(types.StringType, mapData)
		if err != nil {
			log.Printf("Error in converting map[string]string to  Map")
			return err
		}
	}

	return err
}

func (v TestRscModel) GetModelData() *NDFCTestRscModel {
	var data = new(NDFCTestRscModel)

	//MARSHAL_BODY

	if !v.Policy.IsNull() && !v.Policy.IsUnknown() {
		data.Policy = v.Policy.ValueString()
	} else {
		data.Policy = ""
	}

	if !v.ParameterMap.IsNull() && !v.ParameterMap.IsUnknown() {
		elements1 := make(map[string]ParameterMapValue, len(v.ParameterMap.Elements()))

		data.ParameterMap = make(map[string]NDFCParameterMapValue)

		diag := v.ParameterMap.ElementsAs(context.Background(), &elements1, false)
		if diag != nil {
			panic(diag)
		}
		for k1, ele1 := range elements1 {
			data1 := new(NDFCParameterMapValue)

			// serial_number | String| []| false
			if !ele1.SerialNumber.IsNull() && !ele1.SerialNumber.IsUnknown() {

				data1.SerialNumber = ele1.SerialNumber.ValueString()
			} else {
				data1.SerialNumber = ""
			}

			// custom_attributes_nested_map | Map:String| []| false
			if !ele1.CustomAttributesNestedMap.IsNull() && !ele1.CustomAttributesNestedMap.IsUnknown() {

				mapStringData := make(map[string]string, len(ele1.CustomAttributesNestedMap.Elements()))
				dg := ele1.CustomAttributesNestedMap.ElementsAs(context.Background(), &mapStringData, false)
				if dg.HasError() {
					panic(dg.Errors())
				}
				data1.CustomAttributesNestedMap = make(map[string]string, len(mapStringData))
				for k, v := range mapStringData {
					data1.CustomAttributesNestedMap[k] = v
				}
			}

			data.ParameterMap[k1] = *data1

		}
	}

	return data
}
