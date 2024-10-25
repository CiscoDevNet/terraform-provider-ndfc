package ndfc

import (
	"context"
	"fmt"
	"strings"
	"terraform-provider-ndfc/internal/provider/resources/resource_networks"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func (c NDFC) CheckNetworkVrfConfig(ctx context.Context, dg *diag.Diagnostics, nw *resource_networks.NDFCNetworksModel) error {

	// Get VRFs referenced in Network resource
	vrf := make(map[string][]string)
	for _, network := range nw.Networks {
		if strings.ToLower(network.NetworkTemplateConfig.Layer2Only) == "false" {
			vrf[network.VrfName] = make([]string, 0)
		} else {
			tflog.Debug(ctx, fmt.Sprintf("CheckNetworkVrfConfig: Skipping VRF check for network %s", network.NetworkName))
		}
	}

	// Step 1 - Check if VRFs to update are present in NDFC
	ndfcVRFs, err := c.vrfBulkGet(ctx, nw.FabricName)

	if err != nil {
		tflog.Error(ctx, "CheckNetworkVrfConfig: Failed to Read existing VRFs")
		dg.AddError("VRF Read Failed", err.Error())
		return err
	}

	// Check if the VRF is present
	for vrfName := range vrf {
		vrfEntry, ok := ndfcVRFs.Vrfs[vrfName]
		if !ok {
			dg.AddError("VRF not found", fmt.Sprintf("VRF %s not found", vrfName))
			return fmt.Errorf("CheckNetworkVrfConfig: VRF %s not found", vrfName)
		}
		vrfEntry.FilterThisValue = true
		ndfcVRFs.Vrfs[vrfName] = vrfEntry
	}

	//Delete vrf entries that are not needed
	for vrfName, vrfEntry := range ndfcVRFs.Vrfs {
		if !vrfEntry.FilterThisValue {
			delete(ndfcVRFs.Vrfs, vrfName)
		}
	}

	// Step 2 - Check if VRFs are attached
	err = c.RscGetVrfAttachments(ctx, dg, ndfcVRFs)
	if err != nil {
		tflog.Error(ctx, "CheckNetworkVrfConfig: Failed to Read existing VRF Attachments")
		dg.AddError("VRF Attachments Read Failed", err.Error())
		return err
	}
	// Check if network attachments are also attached in corresponding VRF
	for nwName, network := range nw.Networks {
		vrfEntry, ok := ndfcVRFs.Vrfs[network.VrfName]
		if !ok {
			tflog.Info(ctx, fmt.Sprintf("CheckNetworkVrfConfig: VRF %s not found", network.VrfName))
			continue
		}
		for serial := range network.Attachments {
			if _, ok := vrfEntry.AttachList[serial]; !ok {
				//path := path.
				dg.AddAttributeError(path.Root("ndfc_networks").AtName("networks").AtMapKey(nwName).AtName("attachments").AtMapKey(serial),
					"Config error: Attachment not found",
					fmt.Sprintf("Attachment %s not found in VRF %s", serial, network.VrfName))
				return fmt.Errorf("CheckNetworkVrfConfig: Attachment %s not found in VRF %s", serial, network.VrfName)
			} else {
				tflog.Debug(ctx, fmt.Sprintf("CheckNetworkVrfConfig: Attachment %s found in VRF %s", serial, network.VrfName))
			}
		}
	}
	// Check if the VRF configuration matches
	tflog.Debug(ctx, "CheckNetworkVrfConfig: Done checking VRF config")

	return nil
}
