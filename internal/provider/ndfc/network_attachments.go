// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package ndfc

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	rna "terraform-provider-ndfc/internal/provider/resources/resource_network_attachments"
	"terraform-provider-ndfc/internal/provider/resources/resource_networks"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func (c *NDFC) RscGetNetworkAttachments(ctx context.Context, nw *resource_networks.NDFCNetworksModel) error {
	log.Printf("RscGetNetworkAttachments: fabricName=%s", nw.FabricName)

	nwAttachPayload, err := c.netAttachmentsGet(ctx, nw.FabricName, nw.GetNetworksNames())
	if err != nil {
		tflog.Error(ctx, "RscGetNetworkAttachments: Error getting network attachments", map[string]interface{}{"Err": err})
		return err
	}
    c.createVpcPairMap(ctx,nw.FabricName)
	nw.FillAttachmentsFromPayload(nwAttachPayload)

	for netName, nwEntry := range nw.Networks {
		skip := 0
		for serial, attachEntry := range nwEntry.Attachments {
			if *attachEntry.Attached {
				tflog.Debug(ctx, fmt.Sprintf("RscGetNetworkAttachments: Attached entry %s/%s", netName, serial))
				processPortList(&attachEntry)
			} else {
				tflog.Debug(ctx, fmt.Sprintf("RscGetNetworkAttachments: Skip entry: %s/%s as its not attached", netName, serial))
				attachEntry.FilterThisValue = true
				skip++
			}
			//put the updated entry back
			nwEntry.Attachments[serial] = attachEntry
		}
		if skip == len(nwEntry.Attachments) {
			nwEntry.Attachments = nil
		}
		//put it back
		nw.Networks[netName] = nwEntry
	}
	return nil
}

func (c NDFC) RscGetPendingNetAttachments(ctx context.Context, nw *resource_networks.NDFCNetworksModel) {
	log.Printf("RscGetPendingNetAttachments: fabricName=%s", nw.FabricName)

	nwAttachPayload, err := c.netAttachmentsGet(ctx, nw.FabricName, nw.GetNetworksNames())
	if err != nil {
		tflog.Error(ctx, "RscGetPendingNetAttachments: Error getting network attachments", map[string]interface{}{"Err": err})
		return
	}
	nw.FillAttachmentsFromPayload(nwAttachPayload)
	for netName, nwEntry := range nw.Networks {
		for serial, attachEntry := range nwEntry.Attachments {
			if attachEntry.AttachState != NDFCStateNA {
				tflog.Error(ctx, fmt.Sprintf("RscGetPendingNetAttachments: Attachment %s/%s is in PENDING, needs deploy", netName, serial))
				nwEntry.Attachments[serial] = attachEntry
			}
		}
		nw.Networks[netName] = nwEntry
	}

}

func (c *NDFC) RscUpdateNetAttachments(ctx context.Context, dg *diag.Diagnostics, actions map[string]interface{}) error {

	vPlan := actions["plan"].(*resource_networks.NDFCNetworksModel)
	//vState := actions["state"].(*resource_networks.NDFCNetworksModel)

	log.Printf("RscUpdateNetAttachments: fabricName=%s", vPlan.FabricName)
	updateNwAttach := actions["update"].(*rna.NDFCNetworkAttachments)

	if updateNwAttach != nil && len(updateNwAttach.NetworkAttachments) == 0 {
		tflog.Info(ctx, "RscUpdateNetAttachments: No attachments to update")
	} else {
		log.Printf("Network Attachments: %v", updateNwAttach.NetworkAttachments)
		data, err := json.Marshal(updateNwAttach.NetworkAttachments)
		if err != nil {
			tflog.Error(ctx, "RscUpdateNetAttachments: Error marshalling attachments", map[string]interface{}{"Err": err})
			dg.AddError("Error marshalling attachments", err.Error())
			return err
		}
		err = c.netAttachmentsPostPayload(ctx, vPlan.FabricName, data)
		if err != nil {
			tflog.Error(ctx, "RscUpdateNetAttachments: Error updating network attachments", map[string]interface{}{"Err": err})
			dg.AddError("Error updating network attachments", err.Error())
			return err
		}
	}

	naDeploy := actions["deploy"].(*rna.NDFCNetworkAttachments)
	naUndeploy := actions["undeploy"].(*rna.NDFCNetworkAttachments)

	if naDeploy.GlobalDeploy {
		//Using update to deploy everything in update
		tflog.Debug(ctx, "RscUpdateNetAttachments: Global Deploy set")
		c.RscDeployNetworkAttachments(ctx, dg, updateNwAttach)

	} else if naUndeploy.GlobalUndeploy {
		// This is pretty complex to handle
		// Detach everything in update, mark all true to false
		// deploy everything
		// Attach everything in update bring back whatever was true in plan
		tflog.Warn(ctx, "RscUpdateVrfAttachments: Global Undeploy not supported yet")
	} else {
		tflog.Info(ctx, "RscUpdateVrfAttachments: Deploying the changes")
		c.RscDeployNetworkAttachments(ctx, dg, naDeploy)
		//c.DeployVrfFromPayload(ctx, dg, deployVA)
	}

	if len(naUndeploy.NetworkAttachments) > 0 {
		tflog.Info(ctx, "RscUpdateVrfAttachments: Undeploying the changes")
		// Step 1 : Detach
		// Step 2: Deploy
		// Step 3: Attach
		tflog.Warn(ctx, "RscUpdateVrfAttachments: Undeploy not supported yet")
	}

	tflog.Info(ctx, "RscUpdateNetAttachments: Successfully updated attachments")
	return nil
}

func (c NDFC) RscDeleteNetAttachments(ctx context.Context, dg *diag.Diagnostics, in *resource_networks.NDFCNetworksModel) error {

	tflog.Debug(ctx, fmt.Sprintf("RscDeleteNetAttachments: fabricName=%s", in.FabricName))

	err, count := c.netAttachmentsDetach(ctx, in)
	if err != nil {
		tflog.Error(ctx, "RscDeleteNetAttachments: Error detaching network attachments", map[string]interface{}{"Err": err})
		dg.AddError("Error detaching network attachments", err.Error())
		return err
	}

	// It may be better to get the current status of the network attachments
	// to handle any cases where attachment is missing in input, but "PENDING" state in NDFC
	if count == 0 {
		tflog.Info(ctx, "RscDeleteNetAttachments: Nothing was detached - still checking for any attachments in pending state")
		c.RscGetPendingNetAttachments(ctx, in)
		for netName, nwEntry := range in.Networks {
			for serial, attachEntry := range nwEntry.Attachments {
				if attachEntry.AttachState != NDFCStateNA {
					tflog.Error(ctx, fmt.Sprintf("RscDeleteNetAttachments: Attachment %s/%s is in PENDING, needs deploy", netName, serial))
					attachEntry.Deployment = "false"
					nwEntry.Attachments[serial] = attachEntry
				}
			}
			in.Networks[netName] = nwEntry
		}
	}
	c.RscDeployNetworkAttachments(ctx, dg, in)
	return nil
}
