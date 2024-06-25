package resource_interface_vpc

import "github.com/hashicorp/terraform-plugin-framework/types"

func (i *InterfaceVpcModel) GetInterfaceType() string {
	return "vpc"
}

func (i *InterfaceVpcModel) GetID() string {
	return i.Id.ValueString()
}

func (i *InterfaceVpcModel) SetID(id string) {
	if id == "" {
		i.Id = types.StringNull()
		return
	}
	i.Id = types.StringValue(id)
}
