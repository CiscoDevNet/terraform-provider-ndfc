package ndfc

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	rva "terraform-provider-ndfc/internal/provider/resources/resource_vrf_attachments"
	"terraform-provider-ndfc/internal/provider/resources/resource_vrf_bulk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const ResourceVrfAttachments = "vrf_attachments"

/*
VRF Attachments Terraform Id is a combination of the following:
fabricName/vrf1{attachment1,attachment2,attachment3}/vrf2{attachment1,attachment2,attachment3}/vrf3{attachment1,attachment2,attachment3}

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
*/
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
func (c NDFC) RscCreateVrfAttachments(ctx context.Context, dg *diag.Diagnostics, va *rva.NDFCVrfAttachmentsPayloads) error {
	tflog.Debug(ctx, "RscCreateVrfAttachments: Entering Create")

	// Attach the VRF Attachments
	data, err := json.Marshal(va.VrfAttachments)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("RscCreateVrfAttachments: Error marshalling VRF Attachments %s", err.Error()))
		return err
	}
	err = c.vrfAttachmentsPost(ctx, va.FabricName, data)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("RscCreateVrfAttachments: Error creating VRF Attachments %s", err.Error()))
		dg.AddError("Attachments Failed", err.Error())
		return err
	}
	//To Do logic to read and verify attachments?

	return nil
}

func (c NDFC) RscGetVrfAttachments(ctx context.Context, dg *diag.Diagnostics, vrfs *resource_vrf_bulk.NDFCVrfBulkModel) error {
	tflog.Debug(ctx, "RscGetVrfAttachments: Entering Get")
	res, err := c.getVrfAttachments(ctx, dg, vrfs.FabricName, vrfs.GetVrfNames())
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("RscGetVrfAttachments: Error getting VRF Attachments %s", err.Error()))
		return err
	}

	vaPayload := rva.NDFCVrfAttachmentsPayloads{}

	err = json.Unmarshal(res, &vaPayload.VrfAttachments)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("getVrfAttachments: Error unmarshalling VRF Attachments %s", err.Error()))
		log.Printf("getVrfAttachments: Error unmarshalling VRF Attachments %s", string(res))
		return err
	}

	vrfs.FillAttachmentsFromPayload(&vaPayload)
	// Filter out the implicit attachments
	for i, vrfEntry := range vrfs.Vrfs {
		skip := 0
		for j, attachEntry := range vrfs.Vrfs[i].AttachList {
			if attachEntry.Attached != nil &&
				!(*(attachEntry.Attached)) {
				//This is auto entry, not created; skip it
				attachEntry.FilterThisValue = true
				log.Printf("getVrfAttachments: Filtering out implicit attachment %s/{%s}",
					vrfs.Vrfs[i].VrfName,
					vrfs.Vrfs[i].AttachList[j].SwitchName)
				skip++
			} else {
				log.Printf("getVrfAttachments: Keeping explicit attachment %s/{%s}",
					vrfs.Vrfs[i].VrfName,
					vrfs.Vrfs[i].AttachList[j].SwitchName)

			}
			vrfs.Vrfs[i].AttachList[j] = attachEntry
		}
		if skip == len(vrfEntry.AttachList) {
			//All entries are implicit, skip the VRF
			log.Printf("getVrfAttachments: No attachments in VRF %s", vrfs.Vrfs[i].VrfName)
			vrfEntry.AttachList = nil
		}
		vrfs.Vrfs[i] = vrfEntry
	}
	return nil
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
	plan *resource_vrf_bulk.NDFCVrfBulkModel,
	state *resource_vrf_bulk.NDFCVrfBulkModel) {

	tflog.Debug(ctx, "RscUpdateVrfAttachments: Entering Update")

	actionMap := c.diffVrfAttachments(ctx, plan, state)

	updateVA := actionMap["update"]
	deployVA := actionMap["deploy"]
	undeployVA := actionMap["undeploy"]

	// Create the VRF Attachments

	printSummary(ctx, actionMap)

	if updateVA == nil || len(updateVA.VrfAttachments) == 0 {
		tflog.Warn(ctx, "RscUpdateVrfAttachments: No VRF Attachments to update")
	} else {
		data, err := json.Marshal(updateVA.VrfAttachments)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("RscUpdateVrfAttachments: Error marshalling VRF Attachments %s", err.Error()))
			dg.AddError("Update Attachments Failed", err.Error())
			return
		}

		err = c.vrfAttachmentsPost(ctx, updateVA.FabricName, data)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("RscCreateVrfAttachments: Error creating VRF Attachments %s", err.Error()))
			dg.AddError("Update Attachments Failed", err.Error())
			return
		}
	}
	//Check and deploy
	if deployVA.GlobalDeploy {
		//Using update to deploy everything in update
		c.DeployFromPayload(ctx, dg, updateVA)
	} else if undeployVA.GlobalUndeploy {
		// This is pretty complex to handle
		// Detach everything in update, mark all true to false
		// deploy everything
		// Attach everything in update bring back whatever was true in plan
		tflog.Warn(ctx, "RscUpdateVrfAttachments: Global Undeploy not supported yet")
	} else {
		tflog.Info(ctx, "RscUpdateVrfAttachments: Deploying the changes")
		c.DeployFromPayload(ctx, dg, deployVA)
	}

	if len(undeployVA.VrfAttachments) > 0 {
		tflog.Info(ctx, "RscUpdateVrfAttachments: Undeploying the changes")
		// Step 1 : Detach
		// Step 2: Deploy
		// Step 3: Attach
		tflog.Warn(ctx, "RscUpdateVrfAttachments: Undeploy not supported yet")
	}

}

func printSummary(ctx context.Context, actionMap map[string]*rva.NDFCVrfAttachmentsPayloads) {
	log.Printf("==========================ATTACHMENT SUMMARY====================================start")
	log.Printf("Modified Attachments")
	for i := range actionMap["update"].VrfAttachments {
		for j := range actionMap["update"].VrfAttachments[i].AttachList {
			log.Printf("Attachment %s/%s",
				actionMap["update"].VrfAttachments[i].VrfName,
				actionMap["update"].VrfAttachments[i].AttachList[j].SerialNumber)
		}
		data, err := json.Marshal(actionMap["update"].VrfAttachments[i].AttachList)
		if err != nil {
			log.Printf("Marshal failed")
		} else {
			log.Printf("%s", string(data))
		}
	}
	log.Printf("List of undeploy")
	for i := range actionMap["undeploy"].VrfAttachments {
		for j := range actionMap["undeploy"].VrfAttachments[i].AttachList {
			log.Printf("Undeploy Attachment %s/%s",
				actionMap["undeploy"].VrfAttachments[i].VrfName,
				actionMap["undeploy"].VrfAttachments[i].AttachList[j].SerialNumber)
		}
		data, err := json.Marshal(actionMap["undeploy"].VrfAttachments[i].AttachList)
		if err != nil {
			log.Printf("Marshal failed")
		} else {
			log.Printf("%s", string(data))
		}
	}
	log.Printf("List of deploy")
	for i := range actionMap["deploy"].VrfAttachments {
		for j := range actionMap["deploy"].VrfAttachments[i].AttachList {
			log.Printf("Deploy Attachment %s/%s",
				actionMap["deploy"].VrfAttachments[i].VrfName,
				actionMap["deploy"].VrfAttachments[i].AttachList[j].SerialNumber)
		}
		data, err := json.Marshal(actionMap["deploy"].VrfAttachments[i].AttachList)
		if err != nil {
			log.Printf("Marshal failed")
		} else {
			log.Printf("%s", string(data))
		}
	}

	if actionMap["deploy"].GlobalDeploy {
		log.Printf("Global Deploy Changed from false to true")
	}

	if actionMap["undeploy"].GlobalUndeploy {
		log.Printf("Global Deploy Changed from true to false")
	}
	log.Printf("==========================ATTACHMENT SUMMARY====================================end")
}

func (c NDFC) RscDeleteVrfAttachments(ctx context.Context, dg *diag.Diagnostics, delVrf *resource_vrf_bulk.NDFCVrfBulkModel) error {
	tflog.Debug(ctx, "RscDeleteVrfAttachments: Entering Delete")
	va := delVrf.FillAttachPayloadFromModel(true)
	va.FabricName = delVrf.FabricName
	if len(va.VrfAttachments) == 0 {
		tflog.Info(ctx, "vrfAttachmentsDelete: No attachments to delete")
		return nil
	}
	data, err := json.Marshal(va.VrfAttachments)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("RscDeleteVrfAttachments: Error marshalling VRF Attachments %s", err.Error()))
		return err
	}
	err = c.vrfAttachmentsPost(ctx, va.FabricName, data)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("RscDeleteVrfAttachments: Error deleting VRF Attachments %s", err.Error()))
		//still proceed with deployment to remove any attachments that were removed correctly
		//return err
		dg.AddError("Error deleting VRF Attachments", err.Error())
	}
	c.RscDeployAttachments(ctx, dg, delVrf)
	return nil
}
