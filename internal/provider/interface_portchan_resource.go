package provider

import (
	"context"
	"fmt"
	"terraform-provider-ndfc/internal/provider/ndfc"
	"terraform-provider-ndfc/internal/provider/resources/resource_interface_portchannel"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = (*interfacePortChannelResource)(nil)

func NewInterfacePortChannelResource() resource.Resource {
	return &interfacePortChannelResource{}
}

type interfacePortChannelResource struct {
	client *ndfc.NDFC
}

func (r *interfacePortChannelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + ndfc.ResourcePortChannelInterface
}

func (r *interfacePortChannelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_interface_portchannel.InterfacePortchannelResourceSchema(ctx)
}

func (d *interfacePortChannelResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	tflog.Info(ctx, "interfaces vlan Configure")
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

func (r *interfacePortChannelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var in resource_interface_portchannel.InterfacePortchannelModel
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &in)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if r.client == nil {
		panic("Client is nil")
	}
	// Create API call logic
	r.client.RscCreateInterfaces(ctx, &resp.Diagnostics, &in)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Create Interfaces Failed")
		return
	}
	// Example data value setting
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, in)...)
}

func (r *interfacePortChannelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data resource_interface_portchannel.InterfacePortchannelModel

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
	tflog.Info(ctx, fmt.Sprintf("Incoming ID %s", unique_id))
	r.client.RscGetInterfaces(ctx, &resp.Diagnostics, &data)
	if resp.Diagnostics.HasError() {
		return
	}
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)

}

func (r *interfacePortChannelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	var planData resource_interface_portchannel.InterfacePortchannelModel
	var stateData resource_interface_portchannel.InterfacePortchannelModel
	var configData resource_interface_portchannel.InterfacePortchannelModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &planData)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)
	resp.Diagnostics.Append(req.Config.Get(ctx, &configData)...)

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
	r.client.RscUpdateInterfaces(ctx, &resp.Diagnostics, unique_id, &planData, &stateData, &configData)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Update interface Failed")
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &planData)...)
}

func (r *interfacePortChannelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data resource_interface_portchannel.InterfacePortchannelModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Id.IsNull() || data.Id.IsUnknown() {
		resp.Diagnostics.AddError("Delete: Id cannot be empty", "Id should be present")
		resp.State.RemoveResource(ctx)
		return
	}
	r.client.RscDeleteInterfaces(ctx, &resp.Diagnostics, &data)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Delete interface Failed")
		return
	}
	data.Id = types.StringNull()
	data.Policy = types.StringNull()
	data.Interfaces = types.MapNull(resource_interface_portchannel.InterfacesValue{}.Type(ctx))
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
	// Delete API call logic
}

/* DONOT use the deployment flags at different level at the same time */
/* use only one at a time */
/* eg. if global deploy_all_attachments is set, then do not use deploy_attachments or deploy_this_attachment at network/attachment level */
func (r interfacePortChannelResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data resource_interface_portchannel.InterfacePortchannelModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	inData := data.GetModelData()

	ifMap := make(map[string]string)
	for k, v := range inData.Interfaces {
		mapKey := v.SerialNumber + ":" + v.InterfaceName
		if dupKey, ok := ifMap[mapKey]; ok {
			resp.Diagnostics.AddError("Duplicate Interface", fmt.Sprintf("Duplicate Interface %s at keys  %s and %s", mapKey, k, dupKey))
		} else {
			ifMap[mapKey] = k
		}
	}
	if resp.Diagnostics.HasError() {
		return
	}

}
