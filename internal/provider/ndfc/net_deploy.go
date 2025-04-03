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
	"log"

	"fmt"

	rna "terraform-provider-ndfc/internal/provider/resources/resource_network_attachments"
	"terraform-provider-ndfc/internal/provider/resources/resource_networks"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Called from Create/Delete Flows
func (c NDFC) RscDeployNetworkAttachments(ctx context.Context, dg *diag.Diagnostics, attachment interface{}) {

	//Deploy all attachments
	detach_present := false
	var d *NDFCVrfNetworkDeployment
	if va, ok := attachment.(*resource_networks.NDFCNetworksModel); ok {
		d = NewVrfNetworkDeployment(&c, va.FabricName, ResourceNetworks)
		c.fillNetDeploymentDBFromModel(ctx, va, d, &detach_present)
	} else if payload, ok := attachment.(*rna.NDFCNetworkAttachments); ok {
		log.Printf("RscDeployNetworkAttachments: deploy from payload; payload=%v", *payload)
		d = NewVrfNetworkDeployment(&c, payload.FabricName, ResourceNetworks)
		c.fillNetDeploymentDBFromPayload(ctx, payload, d, &detach_present)
	}

	if d.GetDeployPendingCount() == 0 {
		tflog.Info(ctx, "RscDeployNetworkAttachments: No attachments to deploy")
		//dg.AddWarning("Deployment not done", "No attachments to deploy")
		return
	}
	// If detach is present in the list
	// have to wait for deploy complete so that subsequent ops like delete can be taken up
	if detach_present || d.ctrlr.WaitForDeployComplete {
		tflog.Info(ctx, "RscDeployNetworkAttachments: Deploying attachments and wait for completion", map[string]interface{}{
			"detach_present":        detach_present,
			"WaitForDeployComplete": d.ctrlr.WaitForDeployComplete,
		})
		d.DeployFSM(ctx, dg)
	} else {
		tflog.Info(ctx, "RscDeployNetworkAttachments: Deploying attachments - not waiting for completion", map[string]interface{}{
			"detach_present":        detach_present,
			"WaitForDeployComplete": d.ctrlr.WaitForDeployComplete,
		})
		err := c.DeployBulk(ctx, dg, d)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("RscDeployNetworkAttachments: Deploying attachments - not waiting for completion, error: %v", err))
		}
	}
}

func (c NDFC) fillNetDeploymentDBFromModel(ctx context.Context, va *resource_networks.NDFCNetworksModel,
	d *NDFCVrfNetworkDeployment, detach_present *bool) {
	deployAll := false
	*detach_present = false

	tflog.Debug(ctx, "fillDeploymentDBFromModel entry")

	if va.DeployAllAttachments {
		tflog.Info(ctx, "fillNetDeploymentDBFromModel: Deploying all attachments")
		deployAll = true
	}
	for nwName, nwEntry := range va.Networks {
		for serial, attachEntry := range nwEntry.Attachments {
			if attachEntry.Deployment == "false" {
				tflog.Info(ctx, fmt.Sprintf("fillNetDeploymentDBFromModel: Deploying Attachment %s/%s due to detach", nwName, serial))
				*detach_present = true
			}
			if attachEntry.Deployment == "false" || deployAll ||
				nwEntry.DeployAttachments || attachEntry.DeployThisAttachment {
				tflog.Info(ctx, fmt.Sprintf("fillNetDeploymentDBFromModel: Deploying Attachment %s/%s", nwName, serial))
				d.updateDeploymentDB(serial, nwName, attachEntry.Deployment)
			}
		}
	}
}

func (c NDFC) fillNetDeploymentDBFromPayload(ctx context.Context, payload *rna.NDFCNetworkAttachments,
	deployment *NDFCVrfNetworkDeployment, detach_present *bool) {
	tflog.Debug(ctx, "fillNetDeploymentDBFromPayload entry")
	for _, nwEntry := range payload.NetworkAttachments {
		for _, attachEntry := range nwEntry.Attachments {

			if attachEntry.Deployment == "false" {
				*detach_present = true
			}
			tflog.Info(ctx, fmt.Sprintf("fillNetDeploymentDBFromPayload: Deploying Attachment %s/%s", nwEntry.NetworkName, attachEntry.SerialNumber))
			deployment.updateDeploymentDB(attachEntry.SerialNumber, nwEntry.NetworkName, attachEntry.Deployment)
		}
	}
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
