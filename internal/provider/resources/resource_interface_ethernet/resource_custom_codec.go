package resource_interface_ethernet

import "github.com/hashicorp/terraform-plugin-framework/types"

func (i *InterfaceEthernetModel) GetInterfaceType() string {
	return "ethernet"
}

func (i *InterfaceEthernetModel) GetID() string {
	return i.Id.ValueString()
}

func (i *InterfaceEthernetModel) SetID(id string) {
	i.Id = types.StringValue(id)
}
