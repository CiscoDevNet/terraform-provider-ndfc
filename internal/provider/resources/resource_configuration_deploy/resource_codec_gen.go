// Code generated DO NOT EDIT.
package resource_configuration_deploy

import (
	"log"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCConfigurationDeployModel struct {
	FabricName    string   `json:"fabricName,string,omitempty"`
	SerialNumbers []string `json:"serial_numbers,omitempty"`
}

func (v *ConfigurationDeployModel) SetModelData(jsonData *NDFCConfigurationDeployModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.FabricName != "" {
		v.FabricName = types.StringValue(jsonData.FabricName)
	} else {
		v.FabricName = types.StringNull()
	}

	if len(jsonData.SerialNumbers) == 0 {
		log.Printf("v.SerialNumbers is empty")
		v.SerialNumbers, err = types.SetValue(types.StringType, []attr.Value{})
		if err != nil {
			log.Printf("Error in converting []string to  List %v", err)
			return err
		}
	} else {
		listData := make([]attr.Value, len(jsonData.SerialNumbers))
		for i, item := range jsonData.SerialNumbers {
			listData[i] = types.StringValue(item)
		}
		v.SerialNumbers, err = types.SetValue(types.StringType, listData)
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}

	return err
}

func (v ConfigurationDeployModel) GetModelData() *NDFCConfigurationDeployModel {
	var data = new(NDFCConfigurationDeployModel)

	//MARSHAL_BODY

	if !v.FabricName.IsNull() && !v.FabricName.IsUnknown() {
		data.FabricName = v.FabricName.ValueString()
	} else {
		data.FabricName = ""
	}

	return data
}
