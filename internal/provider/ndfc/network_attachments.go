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

func (c *NDFC) RscUpdateNetAttachments(ctx context.Context, dg *diag.Diagnostics, vPlan *resource_networks.NDFCNetworksModel,
	vState *resource_networks.NDFCNetworksModel) error {

	log.Printf("RscUpdateNetAttachments: fabricName=%s", vPlan.FabricName)

	actions := c.networkAttachmentsGetDiff(ctx, dg, vPlan, vState)

	updateNwAttach := actions["update"].(*rna.NDFCNetworkAttachments)

	if updateNwAttach != nil && len(updateNwAttach.NetworkAttachments) == 0 {
		tflog.Info(ctx, "RscUpdateNetAttachments: No attachments to update")
		return nil
	}

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

	tflog.Info(ctx, "RscUpdateNetAttachments: Successfully updated attachments")
	return nil
}

func (c NDFC) RscDeleteNetAttachments(ctx context.Context, dg *diag.Diagnostics, in *resource_networks.NDFCNetworksModel) error {
	payload := rna.NDFCNetworkAttachments{}
	tflog.Debug(ctx, fmt.Sprintf("RscDeleteNetAttachments: fabricName=%s", in.FabricName))

	in.FillAttachmentsPayloadFromModel(&payload, resource_networks.NwAttachmentDetach)
	data, err := json.Marshal(payload)
	if err != nil {
		tflog.Error(ctx, "RscDeleteNetAttachments: Error marshalling attachments", map[string]interface{}{"Err": err})
		return err
	}

	err = c.netAttachmentsPostPayload(ctx, in.FabricName, data)
	if err != nil {
		tflog.Error(ctx, "RscDeleteNetAttachments: Error deleting network attachments", map[string]interface{}{"Err": err})
		return err
	}

	return nil
}
