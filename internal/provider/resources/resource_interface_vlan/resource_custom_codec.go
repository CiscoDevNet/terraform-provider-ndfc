package resource_interface_vlan

import "github.com/hashicorp/terraform-plugin-framework/types"

func (i *InterfaceVlanModel) GetInterfaceType() string {
	return "vlan"
}

func (i *InterfaceVlanModel) GetID() string {
	return i.Id.ValueString()
}

func (i *InterfaceVlanModel) SetID(id string) {
	i.Id = types.StringValue(id)
}
