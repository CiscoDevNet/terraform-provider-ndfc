// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package resource_vpc_pair

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func VpcPairResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"deploy": schema.BoolAttribute{
				Required:            true,
				Description:         "Deploy vPC pair",
				MarkdownDescription: "Deploy vPC pair",
			},
			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "The Terraform Unique Identifier for the vPC pair resource",
				MarkdownDescription: "The Terraform Unique Identifier for the vPC pair resource",
			},
			"serial_numbers": schema.SetAttribute{
				ElementType:         types.StringType,
				Required:            true,
				Description:         "Serial numbers of the switches in the vPC pair. Must contain exactly two serial numbers.",
				MarkdownDescription: "Serial numbers of the switches in the vPC pair. Must contain exactly two serial numbers.",
				Validators: []validator.Set{
					setvalidator.SizeBetween(2, 2),
				},
			},
			"use_virtual_peerlink": schema.BoolAttribute{
				Required:            true,
				Description:         "Set to true to use virtual peer link",
				MarkdownDescription: "Set to true to use virtual peer link",
			},
		},
		Description:         "Resource to configure vPC pair on a switch. Note only VXLAN EVPN fabric is supported",
		MarkdownDescription: "Resource to configure vPC pair on a switch. Note only VXLAN EVPN fabric is supported",
	}
}

type VpcPairModel struct {
	Deploy             types.Bool   `tfsdk:"deploy"`
	Id                 types.String `tfsdk:"id"`
	SerialNumbers      types.Set    `tfsdk:"serial_numbers"`
	UseVirtualPeerlink types.Bool   `tfsdk:"use_virtual_peerlink"`
}
