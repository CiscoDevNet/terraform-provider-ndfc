package provider

import (
	"context"
	"fmt"
	"terraform-provider-ndfc/internal/provider/ndfc"
	"terraform-provider-ndfc/internal/provider/resources/resource_configuration_deploy"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = &ConfigDeployResource{}

var loggingConfigDeploy = "Configuration Deploy Resource"

func NewConfigDeployResource() resource.Resource {
	return &ConfigDeployResource{}
}

type ConfigDeployResource struct {
	client *ndfc.NDFC
}

type ConfigDeployResourceModel struct {
	Id            types.String `tfsdk:"id"`
	FabricName    types.String `tfsdk:"fabric_name"`
	SerialNumbers types.Set    `tfsdk:"serial_numbers"`
}

func (r *ConfigDeployResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	tflog.Debug(ctx, fmt.Sprintf("Start of %s Metadata", loggingConfigDeploy))
	resp.TypeName = req.ProviderTypeName + "_configuration_deploy"
	tflog.Debug(ctx, fmt.Sprintf("End of %s Metadata", loggingConfigDeploy))
}

func (r *ConfigDeployResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	tflog.Debug(ctx, fmt.Sprintf("Start of %s Schema", loggingConfigDeploy))
	resp.Schema = resource_configuration_deploy.ConfigurationDeployResourceSchema(ctx)
	tflog.Debug(ctx, fmt.Sprintf("End of %s Schema", loggingConfigDeploy))
}

func (r *ConfigDeployResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Debug(ctx, fmt.Sprintf("Start of %s Configure", loggingConfigDeploy))
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*ndfc.NDFC)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *ndfc.NDFC, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
	tflog.Debug(ctx, fmt.Sprintf("End of %s Configure", loggingConfigDeploy))
}

func (r *ConfigDeployResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, fmt.Sprintf("Start of %s Create", loggingConfigDeploy))
	var data ConfigDeployResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var serialNumbers []string
	data.SerialNumbers.ElementsAs(ctx, &serialNumbers, false)
	r.client.DeployConfiguration(ctx, &resp.Diagnostics, data.FabricName.ValueString(), serialNumbers)

	data.Id = data.FabricName

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End of %s Create", loggingConfigDeploy))
}

func (r *ConfigDeployResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, fmt.Sprintf("Start of %s Read", loggingConfigDeploy))
	tflog.Debug(ctx, fmt.Sprintf("End of %s Read", loggingConfigDeploy))
}

func (r *ConfigDeployResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, fmt.Sprintf("Start of %s Update", loggingConfigDeploy))
	var data ConfigDeployResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var serialNumbers []string
	data.SerialNumbers.ElementsAs(ctx, &serialNumbers, false)
	r.client.DeployConfiguration(ctx, &resp.Diagnostics, data.FabricName.ValueString(), serialNumbers)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End of %s Update", loggingConfigDeploy))
}

func (r *ConfigDeployResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, fmt.Sprintf("Start of %s Delete", loggingConfigDeploy))
	tflog.Debug(ctx, fmt.Sprintf("End of %s Delete", loggingConfigDeploy))
}
