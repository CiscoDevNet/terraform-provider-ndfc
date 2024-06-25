package resource_interface_portchannel

import "github.com/hashicorp/terraform-plugin-framework/types"

func (i *InterfacePortchannelModel) GetInterfaceType() string {
	return "portchannel"
}

func (i *InterfacePortchannelModel) GetID() string {
	return i.Id.ValueString()
}

func (i *InterfacePortchannelModel) SetID(id string) {
	if id == "" {
		i.Id = types.StringNull()
		return
	}
	i.Id = types.StringValue(id)
}
