// Code generated;  DO NOT EDIT.

package resource_interface_ethernet

import (
	"context"
	"log"
	"strconv"
	. "terraform-provider-ndfc/internal/provider/types"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"terraform-provider-ndfc/internal/provider/resources/resource_interface_common"
)

func (v *InterfaceEthernetModel) SetModelData(jsonData *resource_interface_common.NDFCInterfaceCommonModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.Policy != "" {
		v.Policy = types.StringValue(jsonData.Policy)
	} else {
		v.Policy = types.StringNull()
	}

	if jsonData.PolicyType != "" {
		v.PolicyType = types.StringValue(jsonData.PolicyType)
	} else {
		v.PolicyType = types.StringNull()
	}

	v.Deploy = types.BoolValue(jsonData.Deploy)
	if jsonData.SerialNumber != "" {
		v.SerialNumber = types.StringValue(jsonData.SerialNumber)
	} else {
		v.SerialNumber = types.StringNull()
	}

	if len(jsonData.Interfaces) == 0 {
		log.Printf("v.Interfaces is empty")
		v.Interfaces = types.MapNull(InterfacesValue{}.Type(context.Background()))
	} else {
		mapData := make(map[string]InterfacesValue)
		for key, item := range jsonData.Interfaces {
			if item.FilterThisValue {
				//Skip this entry - this parameter allows filtering
				continue
			}

			data := new(InterfacesValue)
			err = data.SetValue(&item)
			if err != nil {
				log.Printf("Error in InterfacesValue.SetValue")
				return err
			}
			data.state = attr.ValueStateKnown
			mapData[key] = *data
		}
		v.Interfaces, err = types.MapValueFrom(context.Background(), InterfacesValue{}.Type(context.Background()), mapData)
		if err != nil {
			log.Printf("Error in converting map[string]InterfacesValue to  Map")

		}
	}

	return err
}

func (v *InterfacesValue) SetValue(jsonData *resource_interface_common.NDFCInterfacesValue) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.SerialNumber != "" {
		v.SerialNumber = types.StringValue(jsonData.SerialNumber)
	} else {
		v.SerialNumber = types.StringNull()
	}

	if jsonData.InterfaceName != "" {
		v.InterfaceName = types.StringValue(jsonData.InterfaceName)
	} else {
		v.InterfaceName = types.StringNull()
	}

	if jsonData.NvPairs.FreeformConfig != "" {
		v.FreeformConfig = types.StringValue(jsonData.NvPairs.FreeformConfig)
	} else {
		v.FreeformConfig = types.StringNull()
	}

	if jsonData.NvPairs.AdminState != "" {
		x, _ := strconv.ParseBool(jsonData.NvPairs.AdminState)
		v.AdminState = types.BoolValue(x)
	} else {
		v.AdminState = types.BoolNull()
	}

	if jsonData.NvPairs.InterfaceDescription != "" {
		v.InterfaceDescription = types.StringValue(jsonData.NvPairs.InterfaceDescription)
	} else {
		v.InterfaceDescription = types.StringNull()
	}

	if jsonData.NvPairs.BpduGuard != "" {
		v.BpduGuard = types.StringValue(jsonData.NvPairs.BpduGuard)
	} else {
		v.BpduGuard = types.StringNull()
	}

	if jsonData.NvPairs.PortTypeFast != "" {
		x, _ := strconv.ParseBool(jsonData.NvPairs.PortTypeFast)
		v.PortTypeFast = types.BoolValue(x)
	} else {
		v.PortTypeFast = types.BoolNull()
	}

	if jsonData.NvPairs.Mtu != "" {
		v.Mtu = types.StringValue(jsonData.NvPairs.Mtu)
	} else {
		v.Mtu = types.StringNull()
	}

	if jsonData.NvPairs.Speed != "" {
		v.Speed = types.StringValue(jsonData.NvPairs.Speed)
	} else {
		v.Speed = types.StringNull()
	}

	if jsonData.NvPairs.AccessVlan != nil {
		if jsonData.NvPairs.AccessVlan.IsEmpty() {
			v.AccessVlan = types.Int64Null()
		} else {
			v.AccessVlan = types.Int64Value(int64(*jsonData.NvPairs.AccessVlan))
		}

	} else {
		v.AccessVlan = types.Int64Null()
	}

	if jsonData.NvPairs.OrphanPort != "" {
		x, _ := strconv.ParseBool(jsonData.NvPairs.OrphanPort)
		v.OrphanPort = types.BoolValue(x)
	} else {
		v.OrphanPort = types.BoolNull()
	}

	if jsonData.NvPairs.Ptp != "" {
		x, _ := strconv.ParseBool(jsonData.NvPairs.Ptp)
		v.Ptp = types.BoolValue(x)
	} else {
		v.Ptp = types.BoolNull()
	}

	if jsonData.NvPairs.Netflow != "" {
		x, _ := strconv.ParseBool(jsonData.NvPairs.Netflow)
		v.Netflow = types.BoolValue(x)
	} else {
		v.Netflow = types.BoolNull()
	}

	if jsonData.NvPairs.NetflowMonitor != "" {
		v.NetflowMonitor = types.StringValue(jsonData.NvPairs.NetflowMonitor)
	} else {
		v.NetflowMonitor = types.StringNull()
	}

	if jsonData.NvPairs.NetflowSampler != "" {
		v.NetflowSampler = types.StringValue(jsonData.NvPairs.NetflowSampler)
	} else {
		v.NetflowSampler = types.StringNull()
	}

	if jsonData.NvPairs.AllowedVlans != "" {
		v.AllowedVlans = types.StringValue(jsonData.NvPairs.AllowedVlans)
	} else {
		v.AllowedVlans = types.StringNull()
	}

	if jsonData.NvPairs.NativeVlan != nil {
		if jsonData.NvPairs.NativeVlan.IsEmpty() {
			v.NativeVlan = types.Int64Null()
		} else {
			v.NativeVlan = types.Int64Value(int64(*jsonData.NvPairs.NativeVlan))
		}

	} else {
		v.NativeVlan = types.Int64Null()
	}

	if jsonData.DeploymentStatus != "" {
		v.DeploymentStatus = types.StringValue(jsonData.DeploymentStatus)
	} else {
		v.DeploymentStatus = types.StringNull()
	}

	if len(jsonData.CustomPolicyParameters) == 0 {
		log.Printf("v.CustomPolicyParameters is empty")
		v.CustomPolicyParameters = types.MapNull(types.StringType)
	} else {
		mapData := make(map[string]attr.Value)
		for key, item := range jsonData.CustomPolicyParameters {
			mapData[key] = types.StringValue(item)
		}
		v.CustomPolicyParameters, err = types.MapValue(types.StringType, mapData)
		if err != nil {
			log.Printf("Error in converting map[string]string to  Map")
			return err
		}
	}

	return err
}

func (v InterfaceEthernetModel) GetModelData() *resource_interface_common.NDFCInterfaceCommonModel {
	var data = new(resource_interface_common.NDFCInterfaceCommonModel)

	//MARSHAL_BODY

	if !v.Policy.IsNull() && !v.Policy.IsUnknown() {
		data.Policy = v.Policy.ValueString()
	} else {
		data.Policy = ""
	}

	if !v.PolicyType.IsNull() && !v.PolicyType.IsUnknown() {
		data.PolicyType = v.PolicyType.ValueString()
	} else {
		data.PolicyType = ""
	}

	if !v.Deploy.IsNull() && !v.Deploy.IsUnknown() {
		data.Deploy = v.Deploy.ValueBool()
	}

	if !v.SerialNumber.IsNull() && !v.SerialNumber.IsUnknown() {
		data.SerialNumber = v.SerialNumber.ValueString()
	} else {
		data.SerialNumber = ""
	}

	if !v.Interfaces.IsNull() && !v.Interfaces.IsUnknown() {
		elements1 := make(map[string]InterfacesValue, len(v.Interfaces.Elements()))

		data.Interfaces = make(map[string]resource_interface_common.NDFCInterfacesValue)
		diag := v.Interfaces.ElementsAs(context.Background(), &elements1, false)
		if diag != nil {
			panic(diag)
		}
		for k1, ele1 := range elements1 {
			data1 := new(resource_interface_common.NDFCInterfacesValue)
			// filter_this_value | Bool| []| true
			// serial_number | String| []| false
			if !ele1.SerialNumber.IsNull() && !ele1.SerialNumber.IsUnknown() {

				data1.SerialNumber = ele1.SerialNumber.ValueString()
			} else {
				data1.SerialNumber = ""
			}

			// interface_name | String| []| false
			if !ele1.InterfaceName.IsNull() && !ele1.InterfaceName.IsUnknown() {

				data1.InterfaceName = ele1.InterfaceName.ValueString()
			} else {
				data1.InterfaceName = ""
			}

			// interface_type | String| []| true
			// interface_name | String| [nvPairs]| true
			// freeform_config | String| [nvPairs]| false
			if !ele1.FreeformConfig.IsNull() && !ele1.FreeformConfig.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.FreeformConfig = ele1.FreeformConfig.ValueString()
			} else {
				data1.NvPairs.FreeformConfig = ""
			}

			// admin_state | Bool| [nvPairs]| false
			if !ele1.AdminState.IsNull() && !ele1.AdminState.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.AdminState = strconv.FormatBool(ele1.AdminState.ValueBool())
			} else {
				data1.NvPairs.AdminState = ""
			}

			// interface_description | String| [nvPairs]| false
			if !ele1.InterfaceDescription.IsNull() && !ele1.InterfaceDescription.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.InterfaceDescription = ele1.InterfaceDescription.ValueString()
			} else {
				data1.NvPairs.InterfaceDescription = ""
			}

			// bpdu_guard | String| [nvPairs]| false
			if !ele1.BpduGuard.IsNull() && !ele1.BpduGuard.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.BpduGuard = ele1.BpduGuard.ValueString()
			} else {
				data1.NvPairs.BpduGuard = ""
			}

			// port_type_fast | Bool| [nvPairs]| false
			if !ele1.PortTypeFast.IsNull() && !ele1.PortTypeFast.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.PortTypeFast = strconv.FormatBool(ele1.PortTypeFast.ValueBool())
			} else {
				data1.NvPairs.PortTypeFast = ""
			}

			// mtu | String| [nvPairs]| false
			if !ele1.Mtu.IsNull() && !ele1.Mtu.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.Mtu = ele1.Mtu.ValueString()
			} else {
				data1.NvPairs.Mtu = ""
			}

			// speed | String| [nvPairs]| false
			if !ele1.Speed.IsNull() && !ele1.Speed.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.Speed = ele1.Speed.ValueString()
			} else {
				data1.NvPairs.Speed = ""
			}

			// access_vlan | Int64| [nvPairs]| false
			if !ele1.AccessVlan.IsNull() && !ele1.AccessVlan.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.AccessVlan = new(Int64Custom)
				*data1.NvPairs.AccessVlan = Int64Custom(ele1.AccessVlan.ValueInt64())
			} else {
				data1.NvPairs.AccessVlan = nil
			}

			// orphan_port | Bool| [nvPairs]| false
			if !ele1.OrphanPort.IsNull() && !ele1.OrphanPort.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.OrphanPort = strconv.FormatBool(ele1.OrphanPort.ValueBool())
			} else {
				data1.NvPairs.OrphanPort = ""
			}

			// ptp | Bool| [nvPairs]| false
			if !ele1.Ptp.IsNull() && !ele1.Ptp.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.Ptp = strconv.FormatBool(ele1.Ptp.ValueBool())
			} else {
				data1.NvPairs.Ptp = ""
			}

			// netflow | Bool| [nvPairs]| false
			if !ele1.Netflow.IsNull() && !ele1.Netflow.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.Netflow = strconv.FormatBool(ele1.Netflow.ValueBool())
			} else {
				data1.NvPairs.Netflow = ""
			}

			// netflow_monitor | String| [nvPairs]| false
			if !ele1.NetflowMonitor.IsNull() && !ele1.NetflowMonitor.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.NetflowMonitor = ele1.NetflowMonitor.ValueString()
			} else {
				data1.NvPairs.NetflowMonitor = ""
			}

			// netflow_sampler | String| [nvPairs]| false
			if !ele1.NetflowSampler.IsNull() && !ele1.NetflowSampler.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.NetflowSampler = ele1.NetflowSampler.ValueString()
			} else {
				data1.NvPairs.NetflowSampler = ""
			}

			// allowed_vlans | String| [nvPairs]| false
			if !ele1.AllowedVlans.IsNull() && !ele1.AllowedVlans.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.AllowedVlans = ele1.AllowedVlans.ValueString()
			} else {
				data1.NvPairs.AllowedVlans = ""
			}

			// native_vlan | Int64| [nvPairs]| false
			if !ele1.NativeVlan.IsNull() && !ele1.NativeVlan.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.NativeVlan = new(Int64Custom)
				*data1.NvPairs.NativeVlan = Int64Custom(ele1.NativeVlan.ValueInt64())
			} else {
				data1.NvPairs.NativeVlan = nil
			}

			// deployment_status | String| []| false
			// custom_policy_parameters | Map:String| []| false
			if !ele1.CustomPolicyParameters.IsNull() && !ele1.CustomPolicyParameters.IsUnknown() {

				mapStringData := make(map[string]string, len(ele1.CustomPolicyParameters.Elements()))
				dg := ele1.CustomPolicyParameters.ElementsAs(context.Background(), &mapStringData, false)
				if dg.HasError() {
					panic(dg.Errors())
				}
				data1.CustomPolicyParameters = make(map[string]string, len(mapStringData))
				for k, v := range mapStringData {
					data1.CustomPolicyParameters[k] = v
				}
			}

			data.Interfaces[k1] = *data1

		}
	}

	return data
}

func (v *InterfaceEthernetModel) SetDefaultValues() {

	if v.Policy.IsNull() || v.Policy.IsUnknown() {
		v.Policy = types.StringValue("int_trunk_host")
	}

	if v.PolicyType.IsNull() || v.PolicyType.IsUnknown() {
		v.PolicyType = types.StringValue("system")
	}

	if v.Deploy.IsNull() || v.Deploy.IsUnknown() {
		v.Deploy = types.BoolValue(false)
	}

}

func (v *InterfacesValue) SetDefaultValues() {

	if v.AdminState.IsNull() || v.AdminState.IsUnknown() {
		v.AdminState = types.BoolValue(true)
	}

	if v.BpduGuard.IsNull() || v.BpduGuard.IsUnknown() {
		v.BpduGuard = types.StringValue("true")
	}

	if v.PortTypeFast.IsNull() || v.PortTypeFast.IsUnknown() {
		v.PortTypeFast = types.BoolValue(true)
	}

	if v.Mtu.IsNull() || v.Mtu.IsUnknown() {
		v.Mtu = types.StringValue("jumbo")
	}

	if v.Speed.IsNull() || v.Speed.IsUnknown() {
		v.Speed = types.StringValue("Auto")
	}

	if v.OrphanPort.IsNull() || v.OrphanPort.IsUnknown() {
		v.OrphanPort = types.BoolValue(false)
	}

	if v.Ptp.IsNull() || v.Ptp.IsUnknown() {
		v.Ptp = types.BoolValue(false)
	}

	if v.Netflow.IsNull() || v.Netflow.IsUnknown() {
		v.Netflow = types.BoolValue(false)
	}

	if v.AllowedVlans.IsNull() || v.AllowedVlans.IsUnknown() {
		v.AllowedVlans = types.StringValue("none")
	}

}
