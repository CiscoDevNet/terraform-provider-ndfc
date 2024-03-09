package ndfc

//VRF deployment in NDFC implementation

import (
	"context"

	"fmt"

	rva "terraform-provider-ndfc/internal/provider/resources/resource_vrf_attachments"
	"terraform-provider-ndfc/internal/provider/resources/resource_vrf_bulk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Called from Create/Delete Flows
func (c NDFC) RscDeployAttachments(ctx context.Context, dg *diag.Diagnostics, va *resource_vrf_bulk.NDFCVrfBulkModel) {
	d := NewVrfDeployment(&c, va.FabricName)
	deployAllVrf := false
	detach_present := false

	if va.DeployAllAttachments {
		tflog.Info(ctx, "RscDeployAttachments: Deploying all attachments")
		deployAllVrf = true
	}
	//Deploy all attachments
	for vrfName, vrfEntry := range va.Vrfs {
		for serial, attachEntry := range vrfEntry.AttachList {
			if attachEntry.Deployment == "false" {
				tflog.Info(ctx, fmt.Sprintf("RscDeployAttachments: Deploying Attachment %s/%s due to detach", vrfName, serial))
				detach_present = true
			}
			if attachEntry.Deployment == "false" || deployAllVrf ||
				vrfEntry.DeployAttachments || attachEntry.DeployThisAttachment {
				tflog.Info(ctx, fmt.Sprintf("RscDeployAttachments: Deploying Attachment %s/%s", vrfName, serial))
				d.updateDeploymentDB(serial, vrfName, attachEntry.Deployment)
			}
		}
	}
	if d.GetDeployPendingCount() == 0 {
		tflog.Info(ctx, "RscDeployAttachments: No attachments to deploy")
		//dg.AddWarning("Deployment not done", "No attachments to deploy")
		return
	}
	// If detach is present in the list
	// have to wait for deploy complete so that subsequent ops like delete can be taken up
	if detach_present || d.ctrlr.WaitForDeployComplete {
		tflog.Info(ctx, "RscDeployAttachments: Deploying attachments and wait for completion", map[string]interface{}{
			"detach_present":        detach_present,
			"WaitForDeployComplete": d.ctrlr.WaitForDeployComplete,
		})
		d.DeployFSM(ctx, dg)
	} else {
		tflog.Info(ctx, "RscDeployAttachments: Deploying attachments - not waiting for completion", map[string]interface{}{
			"detach_present":        detach_present,
			"WaitForDeployComplete": d.ctrlr.WaitForDeployComplete,
		})
		c.DeployBulk(ctx, dg, d)
	}
}

// Called from Update
func (c NDFC) DeployFromPayload(ctx context.Context, dg *diag.Diagnostics, payload *rva.NDFCVrfAttachmentsPayloads) {
	deployment := NewVrfDeployment(&c, payload.FabricName)
	detach_present := false
	for _, vrfEntry := range payload.VrfAttachments {
		for _, attachEntry := range vrfEntry.AttachList {
			if attachEntry.Deployment == "false" {
				detach_present = true
			}
			deployment.updateDeploymentDB(attachEntry.SerialNumber, vrfEntry.VrfName, attachEntry.Deployment)
		}
	}
	if len(deployment.DeployMap) == 0 {
		tflog.Info(ctx, "vrfAttachmentsDeploy: No attachments to deploy")
		//dg.AddWarning("Deployment not done", "No attachments to deploy")
		return
	}
	if detach_present || c.WaitForDeployComplete {
		tflog.Info(ctx, "vrfAttachmentsDeploy: Deploying attachments and wait for completion", map[string]interface{}{
			"detach_present":        detach_present,
			"WaitForDeployComplete": c.WaitForDeployComplete,
		})
		deployment.DeployFSM(ctx, dg)
	} else {
		tflog.Info(ctx, "vrfAttachmentsDeploy: Deploying attachments - not waiting for completion", map[string]interface{}{
			"detach_present":        detach_present,
			"WaitForDeployComplete": c.WaitForDeployComplete,
		})
		c.DeployBulk(ctx, dg, deployment)
	}

}

func (c NDFC) DeployBulk(ctx context.Context, dg *diag.Diagnostics, deployment *NDFCVrfDeployment) error {
	tflog.Debug(ctx, "DeployBulk: Entering")
	deployment.Deploy(ctx, dg, nil, true)
	return nil
}
