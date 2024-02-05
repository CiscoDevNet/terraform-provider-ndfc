package provider

import (
	"context"
	"fmt"
	"terraform-provider-ndfc/internal/provider/ndfc"
	"terraform-provider-ndfc/internal/provider/resources/resource_vrf_attachments"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = (*vrfAttachmentsResource)(nil)

func NewVrfAttachmentsResource() resource.Resource {
	return &vrfAttachmentsResource{}
}

type vrfAttachmentsResource struct {
	client *ndfc.NDFC
}

func (r *vrfAttachmentsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + ndfc.ResourceVrfAttachments
}

func (r *vrfAttachmentsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_vrf_attachments.VrfAttachmentsResourceSchema(ctx)
}

func (d *vrfAttachmentsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	tflog.Info(ctx, "vrf_attachments Configure")

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

// TODO Duplicate check in input - VRFs and Attachments

func (r *vrfAttachmentsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resource_vrf_attachments.VrfAttachmentsModel
	tflog.Info(ctx, "vrf_attachments: Create Called")
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Create API call logic
	setData := r.client.RscCreateVrfAttachments(ctx, &resp.Diagnostics, &data)
	if setData == nil {
		resp.Diagnostics.AddError("Create VrfAttachments", "Create VrfAttachments failed")
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, setData)...)
}

// TODO Duplicate check in input - VRFs and Attachments

func (r *vrfAttachmentsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data resource_vrf_attachments.VrfAttachmentsModel
	tflog.Info(ctx, "vrf_attachments: Read Called")
	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	retVA := r.client.RscGetVrfAttachments(ctx, &resp.Diagnostics, &data)
	// Read API call logic
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &retVA)...)
}

func (r *vrfAttachmentsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var planVA resource_vrf_attachments.VrfAttachmentsModel
	var stateVA resource_vrf_attachments.VrfAttachmentsModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &planVA)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &stateVA)...)
	if resp.Diagnostics.HasError() {
		return
	}

	updatedVA := r.client.RscUpdateVrfAttachments(ctx, &resp.Diagnostics, &planVA, &stateVA)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &updatedVA)...)
}

func (r *vrfAttachmentsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data resource_vrf_attachments.VrfAttachmentsModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	r.client.RscDeleteVrfAttachments(ctx, &resp.Diagnostics, &data)
	
	// Delete API call logic
}
