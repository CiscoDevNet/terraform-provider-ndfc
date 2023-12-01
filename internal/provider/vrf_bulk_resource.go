package provider

import (
	"context"
	"fmt"
	"terraform-provider-ndfc/internal/provider/ndfc"
	"terraform-provider-ndfc/internal/provider/resources/resource_vrf_bulk"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = (*vrfBulkResource)(nil)

func NewVrfBulkResource() resource.Resource {
	return &vrfBulkResource{}
}

type vrfBulkResource struct {
	client *ndfc.NDFC
}

func (r *vrfBulkResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + ndfc.ResourceVrfBulk
}

func (r *vrfBulkResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_vrf_bulk.VrfBulkResourceSchema(ctx)
}

func (d *vrfBulkResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	tflog.Info(ctx, "vrf_bulk Configure")

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

func (r *vrfBulkResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var in resource_vrf_bulk.VrfBulkModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &in)...)

	if resp.Diagnostics.HasError() {
		return
	}
	// Create API call logic
	vrfDone := r.client.RscCreateBulkVrf(ctx, &resp.Diagnostics, &in)
	if vrfDone == nil {
		tflog.Error(ctx, "Create Bulk VRF Failed")
		return
	}
	// Example data value setting

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, vrfDone)...)
}

func (r *vrfBulkResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data resource_vrf_bulk.VrfBulkModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() || data.Id.IsUnknown() {
		resp.Diagnostics.AddError("Id cannot be empty", "Id should be present")
		resp.State.RemoveResource(ctx)
		return
	}

	unique_id := data.Id.ValueString()
	// unique_id = fabric_name/[vrf1,vrf2,vrf3...]

	tflog.Info(ctx, fmt.Sprintf("Incoming ID %s", unique_id))
	dd := r.client.RscGetBulkVrf(ctx, &resp.Diagnostics, unique_id, false)
	if dd == nil {
		resp.Diagnostics.AddError("Read Failure", "No data received from NDFC")
		return
	}
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, dd)...)
}

func (r *vrfBulkResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var planData resource_vrf_bulk.VrfBulkModel
	var stateData resource_vrf_bulk.VrfBulkModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &planData)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

	if resp.Diagnostics.HasError() {
		return
	}
	unique_id := stateData.Id.ValueString()
	tflog.Info(ctx, fmt.Sprintf("Incoming ID %s", unique_id))
	if unique_id == "" {
		resp.Diagnostics.AddError("ID cannot be empty for update", "Id is mandatory - State may be corrupted")
		return
	}
	// Update API call logic

	// Create API call logic
	r.client.RscUpdateBulkVrf(ctx, &resp.Diagnostics, unique_id, &planData, &stateData)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Create Bulk VRF Failed")
		return
	}
	//

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &planData)...)
}

func (r *vrfBulkResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data resource_vrf_bulk.VrfBulkModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Id.IsNull() || data.Id.IsUnknown() {
		resp.Diagnostics.AddError("Delete: Id cannot be empty", "Id should be present")
		resp.State.RemoveResource(ctx)
		return
	}

	r.client.RscDeleteBulkVrf(ctx, &resp.Diagnostics, data.Id.ValueString(), &data)

	// Delete API call logic
}
