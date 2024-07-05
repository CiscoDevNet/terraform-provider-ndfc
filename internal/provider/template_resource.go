package provider

import (
	"context"
	"fmt"
	"log"
	"strings"
	"terraform-provider-ndfc/internal/provider/ndfc"
	"terraform-provider-ndfc/internal/provider/resources/resource_template"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = (*templateResource)(nil)
var _ resource.ResourceWithImportState = (*templateResource)(nil)

func NewTemplateResource() resource.Resource {
	return &templateResource{}
}

type templateResource struct {
	client *ndfc.NDFC
}

func (r *templateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + ndfc.ResourceTemplate
}

func (r *templateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_template.TemplateResourceSchema(ctx)
	//resp.Schema.Attributes["new_attribute"] = schema.BoolAttribute{Optional: true}
}

func (d *templateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	tflog.Info(ctx, "Template Configure")
	client, ok := req.ProviderData.(*ndfc.NDFC)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected resource  Configure Type",
			fmt.Sprintf("Expected *nd.NDFC, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}
	d.client = client
}

func (r *templateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var in resource_template.TemplateModel
	// Read Terraform plan data into the model
	log.Printf("[TRACE] Create Template")
	resp.Diagnostics.Append(req.Plan.Get(ctx, &in)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Create API call logic
	inData := in.GetModelData()

	r.client.RscValidateTemplateContent(ctx, &resp.Diagnostics, inData.TemplateContent)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Template content validation failed")
		return
	}

	r.client.RscCreateTemplate(ctx, &resp.Diagnostics, &in)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Create Bulk Template Failed")
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, in)...)
}

func (r *templateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data resource_template.TemplateModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dataTmpl := data.GetModelData()
	tflog.Debug(ctx, fmt.Sprintf("Read Template with Name: %s", dataTmpl.TemplateName))
	r.client.RscGetTemplate(ctx, &resp.Diagnostics, dataTmpl.TemplateName, &data)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Read Template Failed")
		resp.Diagnostics.AddWarning("Read Failure", "No configuration found in NDFC")
		//resp.Diagnostics.AddError("Read Failure", "No data received from NDFC")

	}
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r *templateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	var planData resource_template.TemplateModel
	var stateData resource_template.TemplateModel
	var configData resource_template.TemplateModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &planData)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)
	resp.Diagnostics.Append(req.Config.Get(ctx, &configData)...)

	if resp.Diagnostics.HasError() {
		return
	}

	inData := planData.GetModelData()

	r.client.RscValidateTemplateContent(ctx, &resp.Diagnostics, inData.TemplateContent)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Template content validation failed")
		return
	}

	// Create API call logic
	r.client.RscUpdateTemplate(ctx, &resp.Diagnostics, &planData, &stateData, &configData)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Update Template Failed")
		return
	}
	//
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &planData)...)
}

func (r *templateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data resource_template.TemplateModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	r.client.RscDeleteTemplate(ctx, &resp.Diagnostics, &data)

	// Delete API call logic
}

func (r templateResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data resource_template.TemplateModel
	tflog.Debug(ctx, "ValidateConfig called")
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	inData := data.GetModelData()

	if inData.TemplateContent == "" {
		resp.Diagnostics.AddError("TemplateContent", "Template content is required")
	}
	if !strings.HasPrefix(inData.TemplateContent, "##template variables") {
		resp.Diagnostics.AddError("TemplateContent", "Template content must start with '##template variables'")
	}

}

func (r *templateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	//
	unique_id := req.ID
	tflog.Info(ctx, fmt.Sprintf("Incoming ID %s", unique_id))
	data := new(resource_template.TemplateModel)

	r.client.RscImportTemplate(ctx, &resp.Diagnostics, unique_id, data)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Read  Template Failed")
		resp.Diagnostics.AddWarning("Read Failure", "No configuration found in NDFC")
	}
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}
