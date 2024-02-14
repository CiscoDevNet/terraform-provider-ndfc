package ndfc

import (
	"context"
	"fmt"
	"log"
	"strings"
	rva "terraform-provider-ndfc/internal/provider/resources/resource_vrf_attachments"
	. "terraform-provider-ndfc/internal/provider/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
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
func (c NDFC) RscCreateVrfAttachments(ctx context.Context, dg *diag.Diagnostics, va *rva.NDFCVrfAttachmentsModel) error {
	tflog.Debug(ctx, "RscCreateVrfAttachments: Entering Create")
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
				va.VrfAttachments[i].AttachList[j].Vlan = new(Int64Custom)
				*va.VrfAttachments[i].AttachList[j].Vlan = Int64Custom(-1)
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
		dg.AddError("Attachments Failed", err.Error())
		return err
	}

	//Check and deploy
	err = c.vrfAttachmentsDeploy(ctx, dg, va)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("RscCreateVrfAttachments: Error deploying VRF Attachments %s", err.Error()))
		return err
	}
	return nil
}

func (c NDFC) RscGetVrfAttachments(ctx context.Context, dg *diag.Diagnostics, fabricName string, vrfs []string) *rva.NDFCVrfAttachmentsModel {
	tflog.Debug(ctx, "RscGetVrfAttachments: Entering Get")
	return c.getVrfAttachments(ctx, dg, fabricName, vrfs)
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
	planVA *rva.NDFCVrfAttachmentsModel,
	stateVA *rva.NDFCVrfAttachmentsModel) *rva.NDFCVrfAttachmentsModel {

	tflog.Debug(ctx, "RscUpdateVrfAttachments: Entering Update")

	actionMap, _ := c.diffVrfAttachments(ctx, planVA, stateVA)

	updateVA := actionMap["update"]
	// Create the VRF Attachments

	printSummary(ctx, actionMap)
	if updateVA == nil || len(updateVA.VrfAttachments) == 0 {
		tflog.Warn(ctx, "RscUpdateVrfAttachments: No VRF Attachments to update")
	} else {
		err := c.vrfAttachmentsPost(ctx, updateVA)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("RscCreateVrfAttachments: Error creating VRF Attachments %s", err.Error()))
			dg.AddError("Update Attachments Failed", err.Error())
			return nil
		}
	}
	//Check and deploy
	err := c.vrfAttachmentsDeploy(ctx, dg, updateVA)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("RscCreateVrfAttachments: Error deploying VRF Attachments %s", err.Error()))
		return nil
	}

	deployVA := actionMap["deploy"]
	if deployVA != nil && len((*deployVA).VrfAttachments) != 0 {
		err := c.vrfAttachmentsDeploy(ctx, dg, deployVA)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("RscCreateVrfAttachments: Error deploying VRF Attachments %s", err.Error()))
			return nil
		}
	}
	return updateVA

}

func printSummary(ctx context.Context, actionMap map[string]*rva.NDFCVrfAttachmentsModel) {
	log.Printf("Summary of changes to attachments")
	log.Printf("Modified Attachments")
	for i := range actionMap["update"].VrfAttachments {
		for j := range actionMap["update"].VrfAttachments[i].AttachList {
			log.Printf("Attachment %s/%s",
				actionMap["update"].VrfAttachments[i].VrfName,
				actionMap["update"].VrfAttachments[i].AttachList[j].SerialNumber)
		}
	}

	log.Printf("List of undeploy")
	for i := range actionMap["undeploy"].VrfAttachments {
		for j := range actionMap["undeploy"].VrfAttachments[i].AttachList {
			log.Printf("Undeploy Attachment %s/%s",
				actionMap["undeploy"].VrfAttachments[i].VrfName,
				actionMap["undeploy"].VrfAttachments[i].AttachList[j].SerialNumber)
		}
	}

	log.Printf("List of deploy")
	for i := range actionMap["deploy"].VrfAttachments {
		for j := range actionMap["deploy"].VrfAttachments[i].AttachList {
			log.Printf("Deploy Attachment %s/%s",
				actionMap["deploy"].VrfAttachments[i].VrfName,
				actionMap["deploy"].VrfAttachments[i].AttachList[j].SerialNumber)
		}
	}

	if actionMap["deploy"].DeployAllAttachments {
		log.Printf("Global Deploy Changed from false to true")
	}

	if actionMap["undeploy"].DeployAllAttachments {
		log.Printf("Global Deploy Changed from true to false")
	}
}

func (c NDFC) RscDeleteVrfAttachments(ctx context.Context, dg *diag.Diagnostics, va *rva.NDFCVrfAttachmentsModel) error {
	tflog.Debug(ctx, "RscDeleteVrfAttachments: Entering Delete")

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
	old := va.DeployAllAttachments
	defer func() { va.DeployAllAttachments = old }()

	//Delete forces a redeploy
	va.DeployAllAttachments = true
	err = c.vrfAttachmentsDeploy(ctx, dg, va)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("RscCreateVrfAttachments: Error deploying VRF Attachments %s", err.Error()))
		return nil
	}

	return nil
}
