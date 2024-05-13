package resource_interface_loopback

import "github.com/hashicorp/terraform-plugin-framework/types"

func (i *InterfaceLoopbackModel) GetInterfaceType() string {
	return "loopback"
}

func (i *InterfaceLoopbackModel) GetID() string {
	return i.Id.ValueString()
}

func (i *InterfaceLoopbackModel) SetID(id string) {
	i.Id = types.StringValue(id)
}
