// Code generated DO NOT EDIT.
package resource_interface_vlan

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

func (v *InterfaceVlanModel) SetModelData(jsonData *resource_interface_common.NDFCInterfaceCommonModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.Policy != "" {
		v.Policy = types.StringValue(jsonData.Policy)
	} else {
		v.Policy = types.StringNull()
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

	if jsonData.NvPairs.Mtu != "" {
		v.Mtu = types.StringValue(jsonData.NvPairs.Mtu)
	} else {
		v.Mtu = types.StringNull()
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

	if jsonData.NvPairs.Vrf != "" {
		v.Vrf = types.StringValue(jsonData.NvPairs.Vrf)
	} else {
		v.Vrf = types.StringNull()
	}

	if jsonData.NvPairs.Ipv4Address != "" {
		v.Ipv4Address = types.StringValue(jsonData.NvPairs.Ipv4Address)
	} else {
		v.Ipv4Address = types.StringNull()
	}

	if jsonData.NvPairs.Ipv4PrefixLength != "" {
		v.Ipv4PrefixLength = types.StringValue(jsonData.NvPairs.Ipv4PrefixLength)
	} else {
		v.Ipv4PrefixLength = types.StringNull()
	}

	if jsonData.NvPairs.RoutingTag != "" {
		v.RoutingTag = types.StringValue(jsonData.NvPairs.RoutingTag)
	} else {
		v.RoutingTag = types.StringNull()
	}

	if jsonData.NvPairs.DisableIpRedirects != "" {
		x, _ := strconv.ParseBool(jsonData.NvPairs.DisableIpRedirects)
		v.DisableIpRedirects = types.BoolValue(x)
	} else {
		v.DisableIpRedirects = types.BoolNull()
	}

	if jsonData.NvPairs.EnableHsrp != "" {
		x, _ := strconv.ParseBool(jsonData.NvPairs.EnableHsrp)
		v.EnableHsrp = types.BoolValue(x)
	} else {
		v.EnableHsrp = types.BoolNull()
	}

	if jsonData.NvPairs.HsrpGroup != nil {
		if jsonData.NvPairs.HsrpGroup.IsEmpty() {
			v.HsrpGroup = types.Int64Null()
		} else {
			v.HsrpGroup = types.Int64Value(int64(*jsonData.NvPairs.HsrpGroup))
		}

	} else {
		v.HsrpGroup = types.Int64Null()
	}

	if jsonData.NvPairs.HsrpVip != "" {
		v.HsrpVip = types.StringValue(jsonData.NvPairs.HsrpVip)
	} else {
		v.HsrpVip = types.StringNull()
	}

	if jsonData.NvPairs.HsrpPriority != nil {
		if jsonData.NvPairs.HsrpPriority.IsEmpty() {
			v.HsrpPriority = types.Int64Null()
		} else {
			v.HsrpPriority = types.Int64Value(int64(*jsonData.NvPairs.HsrpPriority))
		}

	} else {
		v.HsrpPriority = types.Int64Null()
	}

	if jsonData.NvPairs.HsrpVersion != "" {
		v.HsrpVersion = types.StringValue(jsonData.NvPairs.HsrpVersion)
	} else {
		v.HsrpVersion = types.StringNull()
	}

	if jsonData.NvPairs.Preempt != "" {
		x, _ := strconv.ParseBool(jsonData.NvPairs.Preempt)
		v.Preempt = types.BoolValue(x)
	} else {
		v.Preempt = types.BoolNull()
	}

	if jsonData.NvPairs.Mac != "" {
		v.Mac = types.StringValue(jsonData.NvPairs.Mac)
	} else {
		v.Mac = types.StringNull()
	}

	if jsonData.NvPairs.DhcpServerAddr1 != "" {
		v.DhcpServerAddr1 = types.StringValue(jsonData.NvPairs.DhcpServerAddr1)
	} else {
		v.DhcpServerAddr1 = types.StringNull()
	}

	if jsonData.NvPairs.DhcpServerAddr2 != "" {
		v.DhcpServerAddr2 = types.StringValue(jsonData.NvPairs.DhcpServerAddr2)
	} else {
		v.DhcpServerAddr2 = types.StringNull()
	}

	if jsonData.NvPairs.DhcpServerAddr3 != "" {
		v.DhcpServerAddr3 = types.StringValue(jsonData.NvPairs.DhcpServerAddr3)
	} else {
		v.DhcpServerAddr3 = types.StringNull()
	}

	if jsonData.NvPairs.VrfDhcp1 != "" {
		v.VrfDhcp1 = types.StringValue(jsonData.NvPairs.VrfDhcp1)
	} else {
		v.VrfDhcp1 = types.StringNull()
	}

	if jsonData.NvPairs.VrfDhcp2 != "" {
		v.VrfDhcp2 = types.StringValue(jsonData.NvPairs.VrfDhcp2)
	} else {
		v.VrfDhcp2 = types.StringNull()
	}

	if jsonData.NvPairs.VrfDhcp3 != "" {
		v.VrfDhcp3 = types.StringValue(jsonData.NvPairs.VrfDhcp3)
	} else {
		v.VrfDhcp3 = types.StringNull()
	}

	if jsonData.NvPairs.AdvertiseSubnetInUnderlay != "" {
		x, _ := strconv.ParseBool(jsonData.NvPairs.AdvertiseSubnetInUnderlay)
		v.AdvertiseSubnetInUnderlay = types.BoolValue(x)
	} else {
		v.AdvertiseSubnetInUnderlay = types.BoolNull()
	}

	if jsonData.DeploymentStatus != "" {
		v.DeploymentStatus = types.StringValue(jsonData.DeploymentStatus)
	} else {
		v.DeploymentStatus = types.StringNull()
	}

	return err
}

func (v InterfaceVlanModel) GetModelData() *resource_interface_common.NDFCInterfaceCommonModel {
	var data = new(resource_interface_common.NDFCInterfaceCommonModel)

	//MARSHAL_BODY

	if !v.Policy.IsNull() && !v.Policy.IsUnknown() {
		data.Policy = v.Policy.ValueString()
	} else {
		data.Policy = ""
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

			// mtu | String| [nvPairs]| false
			if !ele1.Mtu.IsNull() && !ele1.Mtu.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.Mtu = ele1.Mtu.ValueString()
			} else {
				data1.NvPairs.Mtu = ""
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

			// vrf | String| [nvPairs]| false
			if !ele1.Vrf.IsNull() && !ele1.Vrf.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.Vrf = ele1.Vrf.ValueString()
			} else {
				data1.NvPairs.Vrf = ""
			}

			// ipv4_address | String| [nvPairs]| false
			if !ele1.Ipv4Address.IsNull() && !ele1.Ipv4Address.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.Ipv4Address = ele1.Ipv4Address.ValueString()
			} else {
				data1.NvPairs.Ipv4Address = ""
			}

			// ipv4_prefix_length | String| [nvPairs]| false
			if !ele1.Ipv4PrefixLength.IsNull() && !ele1.Ipv4PrefixLength.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.Ipv4PrefixLength = ele1.Ipv4PrefixLength.ValueString()
			} else {
				data1.NvPairs.Ipv4PrefixLength = ""
			}

			// routing_tag | String| [nvPairs]| false
			if !ele1.RoutingTag.IsNull() && !ele1.RoutingTag.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.RoutingTag = ele1.RoutingTag.ValueString()
			} else {
				data1.NvPairs.RoutingTag = ""
			}

			// disable_ip_redirects | Bool| [nvPairs]| false
			if !ele1.DisableIpRedirects.IsNull() && !ele1.DisableIpRedirects.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.DisableIpRedirects = strconv.FormatBool(ele1.DisableIpRedirects.ValueBool())
			} else {
				data1.NvPairs.DisableIpRedirects = ""
			}

			// enable_hsrp | Bool| [nvPairs]| false
			if !ele1.EnableHsrp.IsNull() && !ele1.EnableHsrp.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.EnableHsrp = strconv.FormatBool(ele1.EnableHsrp.ValueBool())
			} else {
				data1.NvPairs.EnableHsrp = ""
			}

			// hsrp_group | Int64| [nvPairs]| false
			if !ele1.HsrpGroup.IsNull() && !ele1.HsrpGroup.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.HsrpGroup = new(Int64Custom)
				*data1.NvPairs.HsrpGroup = Int64Custom(ele1.HsrpGroup.ValueInt64())
			} else {
				data1.NvPairs.HsrpGroup = nil
			}

			// hsrp_vip | String| [nvPairs]| false
			if !ele1.HsrpVip.IsNull() && !ele1.HsrpVip.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.HsrpVip = ele1.HsrpVip.ValueString()
			} else {
				data1.NvPairs.HsrpVip = ""
			}

			// hsrp_priority | Int64| [nvPairs]| false
			if !ele1.HsrpPriority.IsNull() && !ele1.HsrpPriority.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.HsrpPriority = new(Int64Custom)
				*data1.NvPairs.HsrpPriority = Int64Custom(ele1.HsrpPriority.ValueInt64())
			} else {
				data1.NvPairs.HsrpPriority = nil
			}

			// hsrp_version | String| [nvPairs]| false
			if !ele1.HsrpVersion.IsNull() && !ele1.HsrpVersion.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.HsrpVersion = ele1.HsrpVersion.ValueString()
			} else {
				data1.NvPairs.HsrpVersion = ""
			}

			// preempt | Bool| [nvPairs]| false
			if !ele1.Preempt.IsNull() && !ele1.Preempt.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.Preempt = strconv.FormatBool(ele1.Preempt.ValueBool())
			} else {
				data1.NvPairs.Preempt = ""
			}

			// mac | String| [nvPairs]| false
			if !ele1.Mac.IsNull() && !ele1.Mac.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.Mac = ele1.Mac.ValueString()
			} else {
				data1.NvPairs.Mac = ""
			}

			// dhcp_server_addr1 | String| [nvPairs]| false
			if !ele1.DhcpServerAddr1.IsNull() && !ele1.DhcpServerAddr1.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.DhcpServerAddr1 = ele1.DhcpServerAddr1.ValueString()
			} else {
				data1.NvPairs.DhcpServerAddr1 = ""
			}

			// dhcp_server_addr2 | String| [nvPairs]| false
			if !ele1.DhcpServerAddr2.IsNull() && !ele1.DhcpServerAddr2.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.DhcpServerAddr2 = ele1.DhcpServerAddr2.ValueString()
			} else {
				data1.NvPairs.DhcpServerAddr2 = ""
			}

			// dhcp_server_addr3 | String| [nvPairs]| false
			if !ele1.DhcpServerAddr3.IsNull() && !ele1.DhcpServerAddr3.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.DhcpServerAddr3 = ele1.DhcpServerAddr3.ValueString()
			} else {
				data1.NvPairs.DhcpServerAddr3 = ""
			}

			// vrf_dhcp1 | String| [nvPairs]| false
			if !ele1.VrfDhcp1.IsNull() && !ele1.VrfDhcp1.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.VrfDhcp1 = ele1.VrfDhcp1.ValueString()
			} else {
				data1.NvPairs.VrfDhcp1 = ""
			}

			// vrf_dhcp2 | String| [nvPairs]| false
			if !ele1.VrfDhcp2.IsNull() && !ele1.VrfDhcp2.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.VrfDhcp2 = ele1.VrfDhcp2.ValueString()
			} else {
				data1.NvPairs.VrfDhcp2 = ""
			}

			// vrf_dhcp3 | String| [nvPairs]| false
			if !ele1.VrfDhcp3.IsNull() && !ele1.VrfDhcp3.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.VrfDhcp3 = ele1.VrfDhcp3.ValueString()
			} else {
				data1.NvPairs.VrfDhcp3 = ""
			}

			// advertise_subnet_in_underlay | Bool| [nvPairs]| false
			if !ele1.AdvertiseSubnetInUnderlay.IsNull() && !ele1.AdvertiseSubnetInUnderlay.IsUnknown() {
				//-----inline nested----
				data1.NvPairs.AdvertiseSubnetInUnderlay = strconv.FormatBool(ele1.AdvertiseSubnetInUnderlay.ValueBool())
			} else {
				data1.NvPairs.AdvertiseSubnetInUnderlay = ""
			}

			// deployment_status | String| []| false
			data.Interfaces[k1] = *data1

		}
	}

	return data
}
