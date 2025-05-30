// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package resource_configuration_deploy

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func ConfigurationDeployResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"config_save": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Save the configuration",
				MarkdownDescription: "Save the configuration",
				Default:             booldefault.StaticBool(false),
			},
			"fabric_name": schema.StringAttribute{
				Required:            true,
				Description:         "The name of the fabric",
				MarkdownDescription: "The name of the fabric",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "The Terraform Unique Identifier for the Inventory Devices resource.",
				MarkdownDescription: "The Terraform Unique Identifier for the Inventory Devices resource.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"serial_numbers": schema.SetAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Description:         "Value 'ALL' if all switches in the fabric are to be deployed, or a list of serial numbers of the switches to be deployed.",
				MarkdownDescription: "Value 'ALL' if all switches in the fabric are to be deployed, or a list of serial numbers of the switches to be deployed.",
			},
			"trigger_deploy_on_update": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Default set to false, config deploy will be only triggered on create of resource. If set to true in resource update, the configurations are deployed to the switches and the flag will be toggled back to false after the deployment is completed, when terraform refresh is performed. Terraform plan will always show in-place update for this field when set to true.",
				MarkdownDescription: "Default set to false, config deploy will be only triggered on create of resource. If set to true in resource update, the configurations are deployed to the switches and the flag will be toggled back to false after the deployment is completed, when terraform refresh is performed. Terraform plan will always show in-place update for this field when set to true.",
				Default:             booldefault.StaticBool(false),
			},
		},
		Description:         "This resource allows configuration deployment operations across specified switches or all switches in an NDFC-managed fabric.",
		MarkdownDescription: "This resource allows configuration deployment operations across specified switches or all switches in an NDFC-managed fabric.",
	}
}

type ConfigurationDeployModel struct {
	ConfigSave            types.Bool   `tfsdk:"config_save"`
	FabricName            types.String `tfsdk:"fabric_name"`
	Id                    types.String `tfsdk:"id"`
	SerialNumbers         types.Set    `tfsdk:"serial_numbers"`
	TriggerDeployOnUpdate types.Bool   `tfsdk:"trigger_deploy_on_update"`
}
