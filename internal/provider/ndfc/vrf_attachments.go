package ndfc

import (
	"context"
	"fmt"
	"strings"
	rva "terraform-provider-ndfc/internal/provider/resources/resource_vrf_attachments"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const ResourceVrfAttachments = "vrf_attachments"

/*
VRF Attachments Terraform Id is a combination of the following:
fabricName/vrf1{attachment1,attachment2,attachment3}/vrf2{attachment1,attachment2,attachment3}/vrf3{attachment1,attachment2,attachment3}
*/
func (c NDFC) VrfAttachmentsCreateID(ndVA *rva.NDFCVrfAttachmentsModel) (string, []string) {
	uniqueID := ndVA.FabricName
	vrfs := make([]string, 0)

	tflog.Debug(context.Background(), "vrfAttachmentsGetID: Entering")
	// Loop through Attachments
	for va := range ndVA.VrfAttachments {
		uniqueID += "/"
		uniqueID += ndVA.VrfAttachments[va].VrfName + "{"
		vrfs = append(vrfs, ndVA.VrfAttachments[va].VrfName)
		// Loop through AttachList
		for la := range ndVA.VrfAttachments[va].AttachList {
			if la > 0 {
				uniqueID += ","
			}
			uniqueID += ndVA.VrfAttachments[va].AttachList[la].SerialNumber
		}
		uniqueID += "}"
	}
	tflog.Info(context.Background(), fmt.Sprintf("vrfAttachmentsGetID: uniqueID %s", uniqueID))
	return uniqueID, vrfs
}

func VrfAttachmentsSplitID(id string) (string, []string) {
	tflog.Debug(context.Background(), fmt.Sprintf("vrfAttachmentsSplitID: Entering id %s", id))
	// Split the ID into its components
	components := strings.Split(id, "/")
	fabricName := components[0]
	vrfs := make([]string, 0)
	for i := 1; i < len(components); i++ {
		// Split the VRFs into their components
		vrfComponents := strings.Split(components[i], "{")
		vrfs = append(vrfs, vrfComponents[0])
	}
	return fabricName, vrfs
}

// TODO - roll back if there is partial failure
// Creating VRF Attachment resource involves attaching all the VRFs/Switches present in the resource
func (c NDFC) RscCreateVrfAttachments(ctx context.Context, dg *diag.Diagnostics, data *rva.VrfAttachmentsModel) *rva.VrfAttachmentsModel {
	tflog.Debug(ctx, "RscCreateVrfAttachments: Entering Create")

	va := data.GetModelData()
	ID, vrfs := c.VrfAttachmentsCreateID(va)

	tflog.Debug(ctx, fmt.Sprintf("RscCreateVrfAttachments: ID %s", ID))

	// Set the FabricName and VrfName in the AttachList - this is required for the API call

	for i := range va.VrfAttachments {
		for j := range va.VrfAttachments[i].AttachList {
			va.VrfAttachments[i].AttachList[j].FabricName = va.FabricName
			va.VrfAttachments[i].AttachList[j].VrfName = va.VrfAttachments[i].VrfName
			va.VrfAttachments[i].AttachList[j].Deployment = "true"
			/* NDFCBUG: throws 500 Server error if vlan field is empty */
			/* Setting to -1 to avoid this - UI does the same */
			if va.VrfAttachments[i].AttachList[j].Vlan == nil {
				va.VrfAttachments[i].AttachList[j].Vlan = new(rva.Int64Custom)
				*va.VrfAttachments[i].AttachList[j].Vlan = rva.Int64Custom(-1)
			}
		}
	}
	if c.vrfAttachmentsIsPresent(ctx, dg, va.FabricName, vrfs) {
		tflog.Error(ctx, fmt.Sprintf("RscCreateVrfAttachments: VRF Attachments already exist on %s", va.FabricName))
		return nil
	}
	// Attach the VRF Attachments
	err := c.vrfAttachmentsPost(ctx, va)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("RscCreateVrfAttachments: Error creating VRF Attachments %s", err.Error()))
		return nil
	}

	//Check and deploy
	err = c.vrfAttachmentsDeploy(ctx, dg, va)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("RscCreateVrfAttachments: Error deploying VRF Attachments %s", err.Error()))
		return nil
	}
	//Get attachments after set
	data.Id = types.StringValue(ID)
	retVA := c.getVrfAttachments(ctx, dg, ID, va)
	return retVA
}

func (c NDFC) RscGetVrfAttachments(ctx context.Context, dg *diag.Diagnostics, in *rva.VrfAttachmentsModel) *rva.VrfAttachmentsModel {
	tflog.Debug(ctx, "RscGetVrfAttachments: Entering Get")
	ID := ""
	va := in.GetModelData()
	if in.Id.IsNull() || in.Id.IsUnknown() {
		tflog.Error(ctx, "RscGetVrfAttachments: Id is unknown")
		ID, _ = c.VrfAttachmentsCreateID(va)
	} else {
		ID = in.Id.ValueString()
	}
	return c.getVrfAttachments(ctx, dg, ID, va)
}

/*
VRF Attachments Update scenarios
Compare the state and plan objects
New Attachments in Plan
=======================
+ New attachment on existing VRF
+ New VRF and new attachments
Actions:
Check if attachment is attached state => Do nothing
Attachments are in detached state => Fresh attach
Missing Attachments in Plan
===========================
+ Missing on existing VRF
+ VRF itself missing
Action:  Check if attached => Detach it

	Already detached => Do nothing

Attachments match - but params have changed
===========================================
Action: Just attach again or detach and attach?
Get the final list for all the VRFs in plan
Sort and return it
*/
func (c NDFC) RscUpdateVrfAttachments(ctx context.Context, dg *diag.Diagnostics,
	planVA *rva.VrfAttachmentsModel,
	stateVA *rva.VrfAttachmentsModel) *rva.VrfAttachmentsModel {

	tflog.Debug(ctx, "RscUpdateVrfAttachments: Entering Update")

	updateVA, ID := c.diffVrfAttachments(ctx, planVA, stateVA)

	// Create the VRF Attachments
	if updateVA == nil || len(updateVA.VrfAttachments) == 0 {
		tflog.Warn(ctx, "RscUpdateVrfAttachments: No VRF Attachments to update")
	} else {
		err := c.vrfAttachmentsPost(ctx, updateVA)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("RscCreateVrfAttachments: Error creating VRF Attachments %s", err.Error()))
			return nil
		}
	}
	//Check and deploy
	err := c.vrfAttachmentsDeploy(ctx, dg, updateVA)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("RscCreateVrfAttachments: Error deploying VRF Attachments %s", err.Error()))
		return nil
	}
	//Get attachments after set
	retVA := c.getVrfAttachments(ctx, dg, ID, updateVA)

	return retVA
}

func (c NDFC) RscDeleteVrfAttachments(ctx context.Context, dg *diag.Diagnostics, data *rva.VrfAttachmentsModel) error {
	tflog.Debug(ctx, "RscDeleteVrfAttachments: Entering Delete")
	va := data.GetModelData()
	// Delete the VRF Attachments
	for i := range va.VrfAttachments {
		for j := range va.VrfAttachments[i].AttachList {
			va.VrfAttachments[i].AttachList[j].FabricName = va.FabricName
			va.VrfAttachments[i].AttachList[j].VrfName = va.VrfAttachments[i].VrfName
			va.VrfAttachments[i].AttachList[j].Deployment = "false"
		}
	}
	if len(va.VrfAttachments) == 0 {
		tflog.Info(ctx, "vrfAttachmentsDelete: No attachments to delete")
		return nil
	}
	err := c.vrfAttachmentsPost(ctx, va)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("RscDeleteVrfAttachments: Error deleting VRF Attachments %s", err.Error()))
		return err
	}

	//Delete forces a redeploy
	va.DeployAllAttachments = true
	err = c.vrfAttachmentsDeploy(ctx, dg, va)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("RscCreateVrfAttachments: Error deploying VRF Attachments %s", err.Error()))
		return nil
	}

	return nil
}
