package provider

import (
	"context"
	"fmt"
	"log"
	"terraform-provider-ndfc/internal/provider/ndfc"
	"terraform-provider-ndfc/internal/provider/resources/resource_networks"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = (*networkBulkResource)(nil)

func NewNetworksResource() resource.Resource {
	return &networkBulkResource{}
}

type networkBulkResource struct {
	client *ndfc.NDFC
}

func (r *networkBulkResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + ndfc.ResourceNetworks
}

func (r *networkBulkResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_networks.NetworksResourceSchema(ctx)
	//resp.Schema.Attributes["new_attribute"] = schema.BoolAttribute{Optional: true}
}

func (d *networkBulkResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	tflog.Info(ctx, "network_bulk Configure")
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

func (r *networkBulkResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var in resource_networks.NetworksModel
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &in)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if r.client == nil {
		panic("Client is nil")
	}
	// Create API call logic
	networkDone := r.client.RscCreateNetworks(ctx, &resp.Diagnostics, &in)
	if networkDone == nil {
		tflog.Error(ctx, "Create Networks Failed")
		return
	}
	// Example data value setting
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, networkDone)...)
}

func (r *networkBulkResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data resource_networks.NetworksModel

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
	// unique_id = fabric_name/[network1,network2,network3...]
	dataNetwork := data.GetModelData()
	deployMap := make(map[string][]string)
	if dataNetwork.DeployAllAttachments {
		deployMap["global"] = append(deployMap["global"], "all")
	}
	for nwName, v := range dataNetwork.Networks {
		v.NetworkName = nwName

		if v.DeployAttachments {
			//first element is the network itself - means deploy enabled at network level
			deployMap[v.NetworkName] = append(deployMap[v.NetworkName], v.NetworkName)
		}

		for serial, s := range v.Attachments {
			s.SerialNumber = serial
			if s.DeployThisAttachment {
				deployMap[v.NetworkName] = append(deployMap[v.NetworkName], s.SerialNumber)
			}
			if len(s.SwitchPorts) > 0 {
				//Store the order of ports as they appear
				key := "SwitchPorts:" + v.NetworkName + "/" + s.SerialNumber
				deployMap[key] = append(deployMap[key], s.SwitchPorts...)
			}
			if len(s.TorPorts) > 0 {
				//Store the order of ports as they appear
				key := "TorPorts:" + v.NetworkName + "/" + s.SerialNumber
				deployMap[key] = append(deployMap[key], s.TorPorts...)
			}
			v.Attachments[serial] = s
		}
		dataNetwork.Networks[nwName] = v
		log.Printf("DeployMap = %v", deployMap)

	}

	tflog.Info(ctx, fmt.Sprintf("Incoming ID %s", unique_id))
	dd := r.client.RscGetNetworks(ctx, &resp.Diagnostics, unique_id, &deployMap)
	if dd == nil {
		tflog.Error(ctx, "Read Networks Failed")
		resp.Diagnostics.AddWarning("Read Failure", "No configuration found in NDFC")
		//resp.Diagnostics.AddError("Read Failure", "No data received from NDFC")

	}
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, dd)...)

}

func (r *networkBulkResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	var planData resource_networks.NetworksModel
	var stateData resource_networks.NetworksModel
	var configData resource_networks.NetworksModel

	// Read Terraform plan data into the model
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
	r.client.RscUpdateNetworks(ctx, &resp.Diagnostics, unique_id, &planData, &stateData, &configData)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Create Bulk VRF Failed")
		return
	}
	//

	// Save updated data into Terraform state

	resp.Diagnostics.Append(resp.State.Set(ctx, &planData)...)
}

func (r *networkBulkResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data resource_networks.NetworksModel

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

	r.client.RscDeleteNetworks(ctx, &resp.Diagnostics, data.Id.ValueString(), &data)

	// Delete API call logic
}

/* DONOT use the deployment flags at different level at the same time */
/* use only one at a time */
/* eg. if global deploy_all_attachments is set, then do not use deploy_attachments or deploy_this_attachment at network/attachment level */
func (r networkBulkResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data resource_networks.NetworksModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	globalDeploy := false
	if !data.DeployAllAttachments.IsNull() || !data.DeployAllAttachments.IsUnknown() {
		tflog.Debug(ctx, "DeployAllAttachments", map[string]interface{}{"Value": data.DeployAllAttachments.ValueBool()})
		globalDeploy = data.DeployAllAttachments.ValueBool()
	}

	elements1 := make(map[string]resource_networks.NetworksValue)
	dg := data.Networks.ElementsAs(ctx, &elements1, false)
	if dg.HasError() {
		resp.Diagnostics.AddError("Networks", "Error in reading Networks")
		return
	}

	for network, v := range elements1 {
		networkDeploy := false
		if !v.DeployAttachments.IsNull() && !v.DeployAttachments.IsUnknown() {
			tflog.Debug(ctx, fmt.Sprintf("DeployAttachments is set for network %s", network))
			networkDeploy = v.DeployAttachments.ValueBool()
			if globalDeploy && networkDeploy {
				tflog.Error(ctx, "Conflicting Deployment Flags - Global & network level")
				resp.Diagnostics.AddAttributeError(
					path.Root("deploy_all_attachments"),
					"Conflicting Deployment Flags",
					"Use only one deployment flag at a time",
				)
				resp.Diagnostics.AddAttributeError(
					path.Root("networks").AtMapKey(network).AtName("deploy_attachments"),
					"Conflicting Deployment Flags",
					"Use only one deployment flag at a time",
				)
				return
			}
		}

		elements2 := make(map[string]resource_networks.AttachmentsValue)
		dg := v.Attachments.ElementsAs(ctx, &elements2, false)
		if dg.HasError() {
			resp.Diagnostics.AddError("AttachList", "Error in reading AttachList")
			return
		}
		for serial, s := range elements2 {
			if !s.DeployThisAttachment.IsNull() && !s.DeployThisAttachment.IsUnknown() {
				attachDeploy := s.DeployThisAttachment.ValueBool()
				tflog.Debug(ctx, fmt.Sprintf("DeployThisAttachment is set for %s/%s", network, serial))
				if globalDeploy && attachDeploy {
					tflog.Error(ctx, "Conflicting Deployment Flags - Global & attachment level")
					resp.Diagnostics.AddAttributeError(
						path.Root("deploy_all_attachments"),
						"Conflicting Deployment Flags",
						"Use only one deployment flag at a time",
					)
					resp.Diagnostics.AddAttributeError(
						path.Root("networks").AtMapKey(network).AtName("attach_list").AtMapKey(serial).AtName("deploy_this_attachment"),
						"Conflicting Deployment Flags",
						"Use only one deployment flag at a time",
					)
					return
				}
				if networkDeploy && attachDeploy {
					tflog.Error(ctx, "Conflicting Deployment Flags - network & attachment level")
					resp.Diagnostics.AddAttributeError(
						path.Root("networks").AtMapKey(network).AtName("deploy_attachments"),
						"Conflicting Deployment Flags",
						"Use only one deployment flag at a time",
					)
					resp.Diagnostics.AddAttributeError(
						path.Root("networks").AtMapKey(network).AtName("attach_list").AtMapKey(serial).AtName("deploy_this_attachment"),
						"Conflicting Deployment Flags",
						"Use only one deployment flag at a time",
					)
					return
				}
				//resp.Diagnostics.AddWarning("Attachments", fmt.Sprintf("VRF %s must have attachment %s configured", v.VrfName.ValueString(), serial))
			}
		}

	}

}
