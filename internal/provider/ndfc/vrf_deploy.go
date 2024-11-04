// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

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
func (c NDFC) RscDeployVrfAttachments(ctx context.Context, dg *diag.Diagnostics, attachment interface{}) {

	//Deploy all attachments
	detach_present := false
	var d *NDFCVrfNetworkDeployment
	if va, ok := attachment.(*resource_vrf_bulk.NDFCVrfBulkModel); ok {
		d = NewVrfNetworkDeployment(&c, va.FabricName, "vrfs")
		c.fillDeploymentDBFromModel(ctx, dg, va, d, &detach_present)
	} else if payload, ok := attachment.(*rva.NDFCVrfAttachmentsPayloads); ok {
		d = NewVrfNetworkDeployment(&c, payload.FabricName, "vrfs")
		c.fillDeploymentDBFromPayload(ctx, dg, payload, d, &detach_present)
	}

	if d.GetDeployPendingCount() == 0 {
		tflog.Info(ctx, "RscDeployVrfAttachments: No attachments to deploy")
		//dg.AddWarning("Deployment not done", "No attachments to deploy")
		return
	}
	// If detach is present in the list
	// have to wait for deploy complete so that subsequent ops like delete can be taken up
	if detach_present || d.ctrlr.WaitForDeployComplete {
		tflog.Info(ctx, "RscDeployVrfAttachments: Deploying attachments and wait for completion", map[string]interface{}{
			"detach_present":        detach_present,
			"WaitForDeployComplete": d.ctrlr.WaitForDeployComplete,
		})
		d.DeployFSM(ctx, dg)
	} else {
		tflog.Info(ctx, "RscDeployVrfAttachments: Deploying attachments - not waiting for completion", map[string]interface{}{
			"detach_present":        detach_present,
			"WaitForDeployComplete": d.ctrlr.WaitForDeployComplete,
		})
		c.DeployBulk(ctx, dg, d)
	}
}

func (c NDFC) fillDeploymentDBFromModel(ctx context.Context, dg *diag.Diagnostics, va *resource_vrf_bulk.NDFCVrfBulkModel,
	d *NDFCVrfNetworkDeployment, detach_present *bool) {
	deployAllVrf := false
	*detach_present = false

	if va.DeployAllAttachments {
		tflog.Info(ctx, "RscDeployVrfAttachments: Deploying all attachments")
		deployAllVrf = true
	}
	for vrfName, vrfEntry := range va.Vrfs {
		for serial, attachEntry := range vrfEntry.AttachList {
			if attachEntry.Deployment == "false" {
				tflog.Info(ctx, fmt.Sprintf("RscDeployVrfAttachments: Deploying Attachment %s/%s due to detach", vrfName, serial))
				*detach_present = true
			}
			if attachEntry.Deployment == "false" || deployAllVrf ||
				vrfEntry.DeployAttachments || attachEntry.DeployThisAttachment {
				tflog.Info(ctx, fmt.Sprintf("RscDeployVrfAttachments: Deploying Attachment %s/%s", vrfName, serial))
				d.updateDeploymentDB(serial, vrfName, attachEntry.Deployment)
			}
		}
	}
}

func (c NDFC) fillDeploymentDBFromPayload(ctx context.Context, dg *diag.Diagnostics, payload *rva.NDFCVrfAttachmentsPayloads,
	deployment *NDFCVrfNetworkDeployment, detach_present *bool) {
	for _, vrfEntry := range payload.VrfAttachments {
		for _, attachEntry := range vrfEntry.AttachList {
			if attachEntry.Deployment == "false" {
				*detach_present = true
			}
			deployment.updateDeploymentDB(attachEntry.SerialNumber, vrfEntry.VrfName, attachEntry.Deployment)
		}
	}
}

func (c NDFC) DeployBulk(ctx context.Context, dg *diag.Diagnostics, deployment *NDFCVrfNetworkDeployment) error {
	tflog.Debug(ctx, "DeployBulk: Entering")
	deployment.Deploy(ctx, dg, nil, true)
	return nil
}

/*
func (c NDFC) deployFromPayload(ctx context.Context, dg *diag.Diagnostics, deployment *NDFCVrfNetworkDeployment, detach_present bool) {
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
*/
