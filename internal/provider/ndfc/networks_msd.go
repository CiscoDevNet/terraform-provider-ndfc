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
	"net"
	"strings"
	rna "terraform-provider-ndfc/internal/provider/resources/resource_network_attachments"
	"terraform-provider-ndfc/internal/provider/resources/resource_networks"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// RscCreateNetworksMsd handles Network creation for both MSD parent and MSD child fabrics
// The isMsdParent and isMsdChild flags determine the behavior:
//
// MSD Parent fabric Network flow (isMsdParent=true):
// 1. Networks are created and owned by parent fabric
// 2. Networks that exist in NDFC are updated with new config (Network ID preserved)
// 3. Networks that don't exist in NDFC are created
// 4. Only parent-managed properties are allowed, setting child-specific properties will error
// 5. Attachments span multiple child fabrics
//
// MSD Child fabric Network flow (isMsdChild=true):
// 1. Networks are created and owned by parent fabric (via parent resource)
// 2. Update child-specific properties using child fabric name in API URL
// 3. Only child properties (dhcp, multicast, trm, netflow, etc.) are updated
// 4. Parent properties must match parent fabric configuration
// 5. Attachments are created on child fabric
func (c *NDFC) RscCreateNetworksMsd(ctx context.Context, dg *diag.Diagnostics, in *resource_networks.NetworksModel, isMsdParent bool, isMsdChild bool, parentFabric string) *resource_networks.NetworksModel {
	fabricName := in.FabricName.ValueString()

	if isMsdParent {
		tflog.Debug(ctx, fmt.Sprintf("RscCreateNetworksMsd entry for MSD Parent fabric %s", fabricName))
	} else if isMsdChild {
		tflog.Debug(ctx, fmt.Sprintf("RscCreateNetworksMsd entry for MSD Child fabric %s (parent: %s)", fabricName, parentFabric))
	}

	nw := in.GetModelData()
	if nw == nil {
		tflog.Error(ctx, "Data conversion from model failed")
		dg.AddError("Data conversion from model failed", "GetModelData returned empty")
		return nil
	}

	// Form ID
	ID := c.RscCreateID(nw, ResourceNetworks)
	depMap := make(map[string][]string)

	// Validate attachment fabric field based on fabric type
	c.validateNetworkAttachmentFabricField(ctx, dg, nw, isMsdParent, isMsdChild)
	if dg.HasError() {
		return nil
	}

	// Handle Network operations based on fabric type
	if isMsdParent {
		// Validate parent parameters for MSD parent fabric
		c.ValidateMsdParentNetworkParameters(ctx, dg, nw)
		if dg.HasError() {
			return nil
		}
		// Handle MSD parent Network operations (create/update Networks)
		err := c.handleMsdParentNetworkOperations(ctx, dg, nw)
		if err != nil {
			return nil
		}
		tflog.Info(ctx, fmt.Sprintf("MSD Parent: Network operations completed successfully ID %s", ID))
	} else if isMsdChild {
		// Handle MSD child Network operations (update child-specific properties)
		err := c.handleMsdChildNetworkOperations(ctx, dg, nw, parentFabric)
		if err != nil {
			return nil
		}
		tflog.Info(ctx, fmt.Sprintf("MSD Child: Network operations completed successfully ID %s", ID))
	}

	// Part 2: Create Attachments if any
	keyMap := c.netAttachmentSerialRemap(ctx, nw)

	// Build deployment map (after serial remap so SerialNumber is populated)
	if nw.DeployAllAttachments {
		depMap["global"] = append(depMap["global"], "all")
	}
	for i, entry := range nw.Networks {
		if entry.DeployAttachments {
			depMap[i] = append(depMap[i], i)
		}
		for _, attachEntry := range entry.Attachments {
			if attachEntry.DeployThisAttachment {
				depMap[i] = append(depMap[i], attachEntry.SerialNumber)
			}
		}
	}

	// Process attachments for MSD fabrics (Create flow)
	err := c.processNetworkAttachmentsForMsdCreate(ctx, dg, nw, isMsdParent)
	if err != nil {
		fabricType := "MSD Child"
		if isMsdParent {
			fabricType = "MSD Parent"
		}
		tflog.Error(ctx, fmt.Sprintf("%s: Network Attachments failed", fabricType))
		dg.AddError(fmt.Sprintf("%s: Network Attachments failed", fabricType), err.Error())
		tflog.Error(ctx, "Rolling back the configurations...delete Networks")
		c.RscDeleteNetworksMsd(ctx, dg, ID, in, isMsdParent, isMsdChild, parentFabric)
		return nil
	}

	// Part 3: Deploy Attachments if any
	c.RscDeployNetworkAttachments(ctx, dg, nw)

	out := c.RscGetNetworks(ctx, dg, ID, &depMap, &keyMap)
	if out == nil {
		tflog.Error(ctx, "Failed to verify: Reading from NDFC after create failed")
		dg.AddError("Failed to verify", "Reading from NDFC after create failed")
		return nil
	}
	if ID != "" {
		out.Id = types.StringValue(ID)
	}
	return out
}

// RscUpdateNetworksMsd handles Network update for both MSD parent and MSD child fabrics
// The isMsdParent and isMsdChild flags determine the behavior:
//
// MSD Parent fabric Network update flow (isMsdParent=true):
// 1. Networks are owned by parent fabric
// 2. Networks that exist are updated, new Networks are created
// 3. Only parent-managed properties are allowed
// 4. Attachments span multiple child fabrics
//
// MSD Child fabric Network update flow (isMsdChild=true):
// 1. Networks are owned by parent fabric
// 2. Only child-specific properties can be updated
// 3. Parent properties must match parent fabric configuration
// 4. Attachments are managed on child fabric
// 5. Network delete is a no-op (parent manages Network lifecycle)
func (c *NDFC) RscUpdateNetworksMsd(ctx context.Context, dg *diag.Diagnostics, ID string,
	plan *resource_networks.NetworksModel,
	state *resource_networks.NetworksModel,
	config *resource_networks.NetworksModel,
	isMsdParent bool, isMsdChild bool, parentFabric string) {

	fabricName := plan.FabricName.ValueString()

	if isMsdParent {
		tflog.Debug(ctx, fmt.Sprintf("RscUpdateNetworksMsd entry for MSD Parent fabric %s", fabricName))
	} else if isMsdChild {
		tflog.Debug(ctx, fmt.Sprintf("RscUpdateNetworksMsd entry for MSD Child fabric %s (parent: %s)", fabricName, parentFabric))
	}

	// Get config model for validation
	nw := config.GetModelData()
	if nw == nil {
		tflog.Error(ctx, "Data conversion from model failed")
		dg.AddError("Data conversion from model failed", "GetModelData returned empty")
		return
	}

	// Validate attachment fabric field based on fabric type
	c.validateNetworkAttachmentFabricField(ctx, dg, nw, isMsdParent, isMsdChild)
	if dg.HasError() {
		return
	}

	// Validate parent parameters for MSD parent fabric
	if isMsdParent {
		c.ValidateMsdParentNetworkParameters(ctx, dg, nw)
		if dg.HasError() {
			return
		}
	}

	// Get diff actions
	netActions := c.networksGetDiff(ctx, dg, plan, state, config)
	delNw := netActions["del"].(*resource_networks.NDFCNetworksModel)
	putNw := netActions["put"].(*resource_networks.NDFCNetworksModel)
	newNw := netActions["add"].(*resource_networks.NDFCNetworksModel)
	planNw := netActions["plan"].(*resource_networks.NDFCNetworksModel)
	stateNw := netActions["state"].(*resource_networks.NDFCNetworksModel)

	// Get existing Networks from NDFC
	nws, err := c.networksGet(ctx, fabricName)
	if err != nil {
		tflog.Error(ctx, "Network Get Failed", map[string]interface{}{"Err": err})
		dg.AddError("Network Get Failed", err.Error())
		return
	}

	// Check VRF dependencies of Networks to be added
	err = c.CheckNetworkVrfConfig(ctx, dg, planNw)
	if err != nil {
		tflog.Error(ctx, "CheckNetworkVrfConfig failed", map[string]interface{}{"Err": err})
		return
	}

	// Check if Networks to be added exist
	for k := range newNw.Networks {
		if ndfcNw, ok := nws.Networks[k]; ok {
			if _, ok := delNw.Networks[k]; ok {
				tflog.Debug(ctx, "Network marked for deletion", map[string]interface{}{"Network": k})
				continue
			}
			// For MSD fabrics, move existing networks to update list
			if isMsdParent {
				tflog.Info(ctx, fmt.Sprintf("MSD Parent: Network %s already exists in NDFC, moving to update list", k))
			} else {
				tflog.Info(ctx, fmt.Sprintf("MSD Child: Network %s already exists in NDFC (from parent), moving to update list", k))
			}
			nwEntry := newNw.Networks[k]
			nwEntry.NetworkId = ndfcNw.NetworkId
			putNw.Networks[k] = nwEntry
			delete(newNw.Networks, k)
		}
	}

	// For MSD child, validate parent properties for all Networks to be updated
	if isMsdChild && len(putNw.Networks) > 0 {
		tflog.Info(ctx, "MSD Child Update: Validating parent properties match NDFC")
		c.validateNetworksMsdParentPropertiesInChild(ctx, dg, putNw, nws, parentFabric)
		if dg.HasError() {
			return
		}
	}

	// Handle Networks to delete based on fabric type
	if isMsdChild {
		// MSD Child: Network delete is skipped (parent manages Network lifecycle)
		if len(delNw.Networks) > 0 {
			var networkNames []string
			for nwName := range delNw.Networks {
				networkNames = append(networkNames, nwName)
			}
			tflog.Info(ctx, fmt.Sprintf("MSD Child fabric: Skipping Network delete for %v - parent manages Network lifecycle", networkNames))
			// Clear the delete list for MSD child
			delNw.Networks = make(map[string]resource_networks.NDFCNetworksValue)
		}
	} else {
		// MSD Parent: Networks to delete are available
		for k := range delNw.Networks {
			if _, ok := nws.Networks[k]; !ok {
				delete(delNw.Networks, k)
				tflog.Error(ctx, fmt.Sprintf("Network to DELETE %s is missing in NDFC", k))
			}
		}
	}

	// Check if Networks to be updated exist
	for k := range putNw.Networks {
		if _, ok := nws.Networks[k]; !ok {
			tflog.Error(ctx, "Network does not exist", map[string]interface{}{"Network": k})
			dg.AddError("Network does not exist", fmt.Sprintf("Network %s does not exist", k))
			return
		}
	}

	// Delete items marked for delete
	if len(delNw.Networks) > 0 {
		tflog.Info(ctx, fmt.Sprintf("MSD Parent: Deleting %d Networks", len(delNw.Networks)))
		// Delete attachments first
		err := c.RscManageNetAttachmentsForMsdParent(ctx, dg, delNw, fabricName, true)
		if err != nil {
			tflog.Error(ctx, "MSD Parent: Network Attachments delete failed")
			dg.AddError("MSD Parent: Network Attachments delete failed", err.Error())
			return
		}
		// Delete Networks
		var nwNames []string
		for nwName := range delNw.Networks {
			nwNames = append(nwNames, nwName)
		}
		err = c.networksDelete(ctx, fabricName, nwNames)
		if err != nil {
			tflog.Error(ctx, "MSD Parent: Network delete failed", map[string]interface{}{"Err": err})
			dg.AddError("MSD Parent: Network delete failed", err.Error())
			return
		}
	}

	// Update existing Networks
	if len(putNw.Networks) > 0 {
		tflog.Info(ctx, fmt.Sprintf("MSD: Updating %d existing Networks", len(putNw.Networks)))
		// Merge current NDFC values with config values
		c.mergeCurrentNetworksWithConfig(ctx, dg, putNw)
		if dg.HasError() {
			return
		}
		c.networksUpdate(ctx, dg, putNw, 0)
		if dg.HasError() {
			tflog.Error(ctx, "MSD: Cannot update existing Networks")
			return
		}
	}

	// Create new Networks (MSD parent only)
	if isMsdParent && len(newNw.Networks) > 0 {
		tflog.Info(ctx, fmt.Sprintf("MSD Parent: Creating %d new Networks", len(newNw.Networks)))
		err = c.networksCreate(ctx, fabricName, newNw)
		if err != nil {
			tflog.Error(ctx, "MSD Parent: Cannot create Network", map[string]interface{}{"Err": err})
			dg.AddError("MSD Parent: Cannot create Network", err.Error())
			return
		}
	}

	// Handle attachment updates - processNetworkAttachmentsForMsdUpdate handles:
	// 1. DELETE (Detach): attachments in state but not in plan
	// 2. UPDATE: existing attachments with modifications
	// 3. CREATE (NewEntry): new attachments (with NDFC check to skip if already exists)
	// Deploy is handled separately below by RscDeployNetworkAttachments
	err = c.processNetworkAttachmentsForMsdUpdate(ctx, dg, planNw, stateNw, putNw, isMsdParent)
	if err != nil {
		fabricType := "MSD Child"
		if isMsdParent {
			fabricType = "MSD Parent"
		}
		tflog.Error(ctx, fmt.Sprintf("%s: Network Attachments update failed", fabricType))
		dg.AddError(fmt.Sprintf("%s: Network Attachments update failed", fabricType), err.Error())
		return
	}

	// Deploy attachments
	c.RscDeployNetworkAttachments(ctx, dg, planNw)

	// Update ID and read back
	newID := c.RscCreateID(plan.GetModelData(), ResourceNetworks)
	tflog.Info(ctx, fmt.Sprintf("New ID after update %s", newID))

	depMap := make(map[string][]string)
	if planNw.DeployAllAttachments {
		depMap["global"] = append(depMap["global"], "all")
	}
	keyMap := make(map[string]string)
	for i, entry := range planNw.Networks {
		if entry.DeployAttachments {
			depMap[i] = append(depMap[i], i)
		}
		for j, attachEntry := range entry.Attachments {
			serial := ""
			if net.ParseIP(j) != nil {
				serial = c.GetSerialFromIP(ctx, planNw.FabricName, j)
			} else {
				serial = j
			}
			keyMap[i+":"+serial] = j
			if attachEntry.DeployThisAttachment {
				depMap[i] = append(depMap[i], serial)
			}
		}
	}

	// Read the updated data from NDFC
	*plan = *(c.RscGetNetworks(ctx, dg, newID, &depMap, &keyMap))
}

// RscDeleteNetworksMsd handles Network deletion for both MSD parent and MSD child fabrics
// The isMsdParent and isMsdChild flags determine the behavior:
//
// MSD Parent fabric Network delete flow (isMsdParent=true):
// 1. Delete attachments from all child fabrics
// 2. Delete Networks from parent fabric
//
// MSD Child fabric Network delete flow (isMsdChild=true):
// 1. Delete is a no-op - parent fabric manages Network lifecycle
// 2. Networks and attachments are deleted when parent resource is destroyed
func (c *NDFC) RscDeleteNetworksMsd(ctx context.Context, dg *diag.Diagnostics, ID string, in *resource_networks.NetworksModel, isMsdParent bool, isMsdChild bool, parentFabric string) {
	fabricName := ""
	results := c.RscBulkSplitID(ID)
	if len(results["fabric"]) > 0 {
		fabricName = results["fabric"][0]
	}

	if isMsdParent {
		tflog.Info(ctx, fmt.Sprintf("RscDeleteNetworksMsd: MSD Parent fabric %s delete request %s", fabricName, ID))
	} else if isMsdChild {
		// For MSD child fabrics, delete is a no-op
		// The parent resource manages deletion of Networks and attachments
		tflog.Info(ctx, fmt.Sprintf("RscDeleteNetworksMsd: MSD Child fabric %s (parent: %s) delete request %s", fabricName, parentFabric, ID))
		tflog.Info(ctx, fmt.Sprintf("MSD Child fabric detected. Delete is no-op. Parent fabric %s manages Network deletion.", parentFabric))
		return
	}

	// MSD Parent: Proceed with deletion
	rsList, err := c.networksIsPresent(ctx, ID)
	if err != nil {
		tflog.Error(ctx, "Error while getting Networks", map[string]interface{}{"Err": err})
		dg.AddError("Network Read Failed", err.Error())
		return
	}

	if len(rsList) == 0 {
		tflog.Info(ctx, "No Networks to delete")
		return
	}

	nwFromId := results["rsc"]
	if len(nwFromId) != len(rsList) {
		errString := fmt.Sprintf("Mismatch in Network data: fabric %s Read %v, from ID %v", fabricName, rsList, nwFromId)
		tflog.Error(ctx, errString)
		dg.AddWarning("Mismatch in Network data", errString)
		nwFromId = rsList
	}

	// Delete attachments from child fabrics
	delNw := in.GetModelData()

	tflog.Info(ctx, "MSD Parent: Deleting attachments from child fabrics")
	err = c.RscManageNetAttachmentsForMsdParent(ctx, dg, delNw, fabricName, true)
	if err != nil {
		tflog.Error(ctx, "MSD Parent: Network Attachments delete failed")
		dg.AddError("Network Attachments delete failed", err.Error())
		return
	}

	// Delete the Networks
	err = c.networksDelete(ctx, fabricName, nwFromId)
	if err != nil {
		errString := fmt.Sprintf("Network delete failed on fabric %s networks %v", fabricName, nwFromId)
		tflog.Error(ctx, "Network delete failed")
		dg.AddError(errString, err.Error())
		return
	}

	tflog.Info(ctx, fmt.Sprintf("MSD Parent: Network Bulk Delete Success %s", ID))
}

// handleMsdParentNetworkOperations handles Network creation/update for MSD parent fabric
// MSD Parent fabric Network flow:
// 1. Networks are created and owned by parent fabric
// 2. Networks that exist in NDFC are updated with new config (Network ID preserved)
// 3. Networks that don't exist in NDFC are created
// 4. Only parent-managed properties are allowed
func (c *NDFC) handleMsdParentNetworkOperations(ctx context.Context, dg *diag.Diagnostics, nw *resource_networks.NDFCNetworksModel) error {
	tflog.Info(ctx, fmt.Sprintf("MSD Parent fabric detected: %s", nw.FabricName))

	// Get existing Networks from NDFC
	ndfcNetworks, err := c.networksGet(ctx, nw.FabricName)
	if err != nil {
		tflog.Error(ctx, "Error while getting Networks for MSD Parent fabric", map[string]interface{}{"Err": err})
		dg.AddError("Network Read Failed", err.Error())
		return err
	}

	// Separate Networks into those that exist (need update) and those that don't (need create)
	existingNetworks := resource_networks.NDFCNetworksModel{
		FabricName: nw.FabricName,
		Networks:   make(map[string]resource_networks.NDFCNetworksValue),
	}
	newNetworks := resource_networks.NDFCNetworksModel{
		FabricName: nw.FabricName,
		Networks:   make(map[string]resource_networks.NDFCNetworksValue),
	}

	for networkName, networkEntry := range nw.Networks {
		if ndfcNetworkEntry, found := ndfcNetworks.Networks[networkName]; found {
			// Network exists - add to update list and preserve Network ID
			networkEntry.NetworkId = ndfcNetworkEntry.NetworkId
			existingNetworks.Networks[networkName] = networkEntry
			tflog.Debug(ctx, fmt.Sprintf("MSD Parent: Network %s already exists, will update it", networkName))
		} else {
			// Network doesn't exist - add to create list
			newNetworks.Networks[networkName] = networkEntry
			tflog.Debug(ctx, fmt.Sprintf("MSD Parent: Network %s doesn't exist, will create it", networkName))
		}
	}

	// Check if VRFs referenced in the networks exist
	err = c.CheckNetworkVrfConfig(ctx, dg, nw)
	if err != nil {
		tflog.Error(ctx, "CheckNetworkVrfConfig failed", map[string]interface{}{"Err": err})
		return err
	}

	// Update existing Networks
	if len(existingNetworks.Networks) > 0 {
		tflog.Info(ctx, fmt.Sprintf("MSD Parent: Updating %d existing Networks", len(existingNetworks.Networks)))
		// Network update validates that config values match NDFC values using DeepEqual.
		// For MSD parent fabrics, networks may have been created earlier with default values
		// that aren't present in the current config. This causes DeepEqual to fail and triggers
		// infinite update retries. To prevent this, merge current NDFC values with config values
		// before calling update.
		c.mergeCurrentNetworksWithConfig(ctx, dg, &existingNetworks)
		if dg.HasError() {
			return fmt.Errorf("MSD parent network merge failed")
		}
		c.networksUpdate(ctx, dg, &existingNetworks, 0)
		if dg.HasError() {
			tflog.Error(ctx, "MSD Parent: Cannot update existing Networks")
			return fmt.Errorf("MSD parent network update failed")
		}
		tflog.Info(ctx, "MSD Parent: Update existing Networks success")
	}

	// Create new Networks
	if len(newNetworks.Networks) > 0 {
		tflog.Info(ctx, fmt.Sprintf("MSD Parent: Creating %d new Networks", len(newNetworks.Networks)))
		err = c.networksCreate(ctx, nw.FabricName, &newNetworks)
		if err != nil {
			tflog.Error(ctx, "MSD Parent: Cannot create Network", map[string]interface{}{"Err": err})
			dg.AddError("Cannot create Network", err.Error())
			return err
		}
		tflog.Info(ctx, "MSD Parent: Create new Networks success")
	}

	tflog.Info(ctx, "MSD Parent: Network operations completed successfully")
	return nil
}

// handleMsdChildNetworkOperations handles Network update for MSD child fabric
// MSD Child fabric Network flow:
// 1. Networks are created and owned by parent fabric (via parent resource)
// 2. Update child-specific properties using child fabric name in API URL
// 3. Only child properties (dhcp, multicast, trm, netflow, etc.) are updated
// 4. Parent properties must match parent fabric configuration
func (c *NDFC) handleMsdChildNetworkOperations(ctx context.Context, dg *diag.Diagnostics, nw *resource_networks.NDFCNetworksModel, parentFabric string) error {
	tflog.Info(ctx, fmt.Sprintf("MSD Child fabric detected: %s (parent: %s)", nw.FabricName, parentFabric))

	// Query child fabric to get existing Networks (Networks are owned by parent)
	ndfcNetworks, err := c.networksGet(ctx, nw.FabricName)
	if err != nil {
		tflog.Error(ctx, "Error while getting existing Networks for MSD Child fabric", map[string]interface{}{"Err": err})
		dg.AddError("Network Read Failed", err.Error())
		return err
	}

	// Check all networks in config are present in NDFC before we allow in-place update
	var missingNetworks []string
	for networkName := range nw.Networks {
		if _, found := ndfcNetworks.Networks[networkName]; !found {
			missingNetworks = append(missingNetworks, networkName)
		} else {
			networkEntry := nw.Networks[networkName]
			networkEntry.NetworkId = ndfcNetworks.Networks[networkName].NetworkId
			nw.Networks[networkName] = networkEntry
		}
	}
	if len(missingNetworks) > 0 {
		errString := fmt.Sprintf("Missing Network(s) %v in NDFC for child fabric update %s", missingNetworks, parentFabric)
		tflog.Error(ctx, errString)
		dg.AddError("Networks missing", errString)
		return fmt.Errorf("%s", errString)
	}

	// Validate parent properties in child config match what NDFC returns.
	// The child fabric GET already returns parent properties from the parent fabric,
	// so we validate against ndfcNetworks (no need for separate parent fabric GET).
	tflog.Info(ctx, "MSD Child: Validating parent properties match NDFC")
	c.validateNetworksMsdParentPropertiesInChild(ctx, dg, nw, ndfcNetworks, parentFabric)
	if dg.HasError() {
		return fmt.Errorf("MSD child parent property validation failed")
	}

	// Check if VRFs referenced in the networks exist
	err = c.CheckNetworkVrfConfig(ctx, dg, nw)
	if err != nil {
		tflog.Error(ctx, "CheckNetworkVrfConfig failed", map[string]interface{}{"Err": err})
		return err
	}

	// For MSD child fabrics, networks already exist in parent with default values that may
	// not be present in config. Merge current NDFC values with config to prevent DeepEqual
	// failures that cause infinite update retries.
	c.mergeCurrentNetworksWithConfig(ctx, dg, nw)
	if dg.HasError() {
		return fmt.Errorf("MSD child network merge failed")
	}

	// Update Networks
	c.networksUpdate(ctx, dg, nw, 0)
	if dg.HasError() {
		tflog.Info(ctx, fmt.Sprintf("Modifying MSD child Networks Failed %v", dg.Errors()))
		return fmt.Errorf("MSD child network update failed")
	}

	// Copy all parent properties from NDFC to config
	for networkName := range nw.Networks {
		if parentEntry, found := ndfcNetworks.Networks[networkName]; found {
			networkEntry := nw.Networks[networkName]
			// Copy parent-managed properties:
			// is_l2_only, vrf_name, net_id, vlan_id, vlan_name, int_desc, gw_ip_address,
			// gw_ipv6_address, secondary_ip_address, mtu_l3intf, route_tag, arp_supress, route_target_both
			networkEntry.NetworkId = parentEntry.NetworkId
			networkEntry.NetworkTemplate = parentEntry.NetworkTemplate
			networkEntry.NetworkExtensionTemplate = parentEntry.NetworkExtensionTemplate
			networkEntry.VrfName = parentEntry.VrfName
			networkEntry.NetworkTemplateConfig.VlanId = parentEntry.NetworkTemplateConfig.VlanId
			networkEntry.NetworkTemplateConfig.VlanName = parentEntry.NetworkTemplateConfig.VlanName
			networkEntry.NetworkTemplateConfig.GatewayIpv4Address = parentEntry.NetworkTemplateConfig.GatewayIpv4Address
			networkEntry.NetworkTemplateConfig.GatewayIpv6Address = parentEntry.NetworkTemplateConfig.GatewayIpv6Address
			networkEntry.NetworkTemplateConfig.InterfaceDescription = parentEntry.NetworkTemplateConfig.InterfaceDescription
			networkEntry.NetworkTemplateConfig.Mtu = parentEntry.NetworkTemplateConfig.Mtu
			networkEntry.NetworkTemplateConfig.SecondaryGateway1 = parentEntry.NetworkTemplateConfig.SecondaryGateway1
			networkEntry.NetworkTemplateConfig.SecondaryGateway2 = parentEntry.NetworkTemplateConfig.SecondaryGateway2
			networkEntry.NetworkTemplateConfig.SecondaryGateway3 = parentEntry.NetworkTemplateConfig.SecondaryGateway3
			networkEntry.NetworkTemplateConfig.SecondaryGateway4 = parentEntry.NetworkTemplateConfig.SecondaryGateway4
			networkEntry.NetworkTemplateConfig.ArpSuppression = parentEntry.NetworkTemplateConfig.ArpSuppression
			networkEntry.NetworkTemplateConfig.Layer2Only = parentEntry.NetworkTemplateConfig.Layer2Only
			networkEntry.NetworkTemplateConfig.RoutingTag = parentEntry.NetworkTemplateConfig.RoutingTag
			networkEntry.NetworkTemplateConfig.RouteTargetBoth = parentEntry.NetworkTemplateConfig.RouteTargetBoth
			// Note: Child-specific properties remain from child config:
			// dhcp_loopback_id, dhcp_servers, multicast_group_address, trm_enable,
			// netflow_enable, vlan_netflow_monitor, l3gw_on_border
			nw.Networks[networkName] = networkEntry
		}
	}

	tflog.Info(ctx, "MSD Child: Network update completed successfully")
	return nil
}

// validateNetworkAttachmentFabricField validates the fabric field in attachments based on fabric type:
// - MSD Parent: fabric field REQUIRED (specifies which child fabric the attachment belongs to)
// - MSD Child: fabric field should NOT be set (computed from resource's fabric_name)
// - Regular: fabric field should NOT be set (computed from resource's fabric_name)
func (c *NDFC) validateNetworkAttachmentFabricField(ctx context.Context, dg *diag.Diagnostics, nw *resource_networks.NDFCNetworksModel, isMsdParent bool, isMsdChild bool) {
	for networkName, networkEntry := range nw.Networks {
		for serialNum, attachment := range networkEntry.Attachments {
			if isMsdParent {
				// MSD parent: fabric field is REQUIRED
				if attachment.Fabric == "" {
					errMsg := fmt.Sprintf("Network '%s' attachment '%s': MSD parent fabric requires 'fabric' field to be set in attachments to specify which child fabric the attachment belongs to",
						networkName, serialNum)
					tflog.Error(ctx, errMsg)
					dg.AddError("Missing fabric field in attachment", errMsg)
				}
			} else {
				// MSD child or Regular fabric: fabric field should NOT be set
				if attachment.Fabric != "" {
					var fabricType string
					if isMsdChild {
						fabricType = "MSD child"
					} else {
						fabricType = "Regular (non-MSD)"
					}
					errMsg := fmt.Sprintf("Network '%s' attachment '%s': %s fabric should not have 'fabric' field set in attachments. The fabric is computed from the resource's fabric_name. Remove the 'fabric' field from the attachment configuration",
						networkName, serialNum, fabricType)
					tflog.Error(ctx, errMsg)
					dg.AddError("Invalid fabric field in attachment", errMsg)
				}
			}
		}
	}
}

// validateNetworksMsdParentPropertiesInChild validates that parent properties in MSD child config match the parent fabric
// For MSD child fabrics, parent-managed properties must match what's configured in the parent fabric
// Parent properties: is_l2_only, vrf_name, net_id, vlan_id, vlan_name, int_desc, gw_ip_address,
//
//	gw_ipv6_address, secondary_ip_address, mtu_l3intf, route_tag, arp_supress, route_target_both
//
// Child properties: dhcp_loopback_id, dhcp_servers, multicast_group_address, trm_enable,
//
//	netflow_enable, vlan_netflow_monitor, l3gw_on_border
func (c *NDFC) validateNetworksMsdParentPropertiesInChild(ctx context.Context, dg *diag.Diagnostics, configNws *resource_networks.NDFCNetworksModel, ndNws *resource_networks.NDFCNetworksModel, parentFabric string) {
	for networkName, childEntry := range configNws.Networks {
		parentEntry, found := ndNws.Networks[networkName]
		if !found {
			continue // Already validated in caller that Network exists
		}

		var mismatchedProps []string

		// Compare parent-managed properties (not child-specific properties)
		childCfg := childEntry.NetworkTemplateConfig
		parentCfg := parentEntry.NetworkTemplateConfig

		// Network ID (net_id - critical parent property)
		if childEntry.NetworkId != nil && *childEntry.NetworkId != 0 && parentEntry.NetworkId != nil && *childEntry.NetworkId != *parentEntry.NetworkId {
			mismatchedProps = append(mismatchedProps, fmt.Sprintf("network_id (child: %d, parent: %d)", *childEntry.NetworkId, *parentEntry.NetworkId))
		}

		// VRF Name (vrf_name - parent property)
		if childEntry.VrfName != "" && childEntry.VrfName != parentEntry.VrfName {
			mismatchedProps = append(mismatchedProps, fmt.Sprintf("vrf_name (child: %s, parent: %s)", childEntry.VrfName, parentEntry.VrfName))
		}

		// Network Template
		if childEntry.NetworkTemplate != "" && childEntry.NetworkTemplate != parentEntry.NetworkTemplate {
			mismatchedProps = append(mismatchedProps, fmt.Sprintf("network_template (child: %s, parent: %s)", childEntry.NetworkTemplate, parentEntry.NetworkTemplate))
		}

		// Network Extension Template
		if childEntry.NetworkExtensionTemplate != "" && childEntry.NetworkExtensionTemplate != parentEntry.NetworkExtensionTemplate {
			mismatchedProps = append(mismatchedProps, fmt.Sprintf("network_extension_template (child: %s, parent: %s)", childEntry.NetworkExtensionTemplate, parentEntry.NetworkExtensionTemplate))
		}

		// VLAN Name (vlan_name)
		if childCfg.VlanName != "" && childCfg.VlanName != parentCfg.VlanName {
			mismatchedProps = append(mismatchedProps, fmt.Sprintf("vlan_name (child: %s, parent: %s)", childCfg.VlanName, parentCfg.VlanName))
		}

		// Interface Description (int_desc)
		if childCfg.InterfaceDescription != "" && childCfg.InterfaceDescription != parentCfg.InterfaceDescription {
			mismatchedProps = append(mismatchedProps, fmt.Sprintf("intf_description (child: %s, parent: %s)", childCfg.InterfaceDescription, parentCfg.InterfaceDescription))
		}

		// Gateway IPv4 Address (gw_ip_address)
		if childCfg.GatewayIpv4Address != "" && childCfg.GatewayIpv4Address != parentCfg.GatewayIpv4Address {
			mismatchedProps = append(mismatchedProps, fmt.Sprintf("gateway_ipv4_address (child: %s, parent: %s)", childCfg.GatewayIpv4Address, parentCfg.GatewayIpv4Address))
		}

		// Gateway IPv6 Address (gw_ipv6_address)
		if childCfg.GatewayIpv6Address != "" && childCfg.GatewayIpv6Address != parentCfg.GatewayIpv6Address {
			mismatchedProps = append(mismatchedProps, fmt.Sprintf("gateway_ipv6_address (child: %s, parent: %s)", childCfg.GatewayIpv6Address, parentCfg.GatewayIpv6Address))
		}

		// MTU (mtu_l3intf)
		if childCfg.Mtu != nil && !childCfg.Mtu.IsEmpty() && parentCfg.Mtu != nil && !parentCfg.Mtu.IsEmpty() {
			if *childCfg.Mtu != *parentCfg.Mtu {
				mismatchedProps = append(mismatchedProps, fmt.Sprintf("mtu (child: %d, parent: %d)", *childCfg.Mtu, *parentCfg.Mtu))
			}
		}

		// Routing Tag (route_tag)
		if childCfg.RoutingTag != nil && !childCfg.RoutingTag.IsEmpty() && parentCfg.RoutingTag != nil && !parentCfg.RoutingTag.IsEmpty() {
			if *childCfg.RoutingTag != *parentCfg.RoutingTag {
				mismatchedProps = append(mismatchedProps, fmt.Sprintf("routing_tag (child: %d, parent: %d)", *childCfg.RoutingTag, *parentCfg.RoutingTag))
			}
		}

		// VLAN ID (vlan_id)
		if childCfg.VlanId != nil && !childCfg.VlanId.IsEmpty() && parentCfg.VlanId != nil && !parentCfg.VlanId.IsEmpty() {
			if *childCfg.VlanId != *parentCfg.VlanId {
				mismatchedProps = append(mismatchedProps, fmt.Sprintf("vlan_id (child: %d, parent: %d)", *childCfg.VlanId, *parentCfg.VlanId))
			}
		}

		// Layer2Only (is_l2_only)
		if childCfg.Layer2Only != "" && childCfg.Layer2Only != parentCfg.Layer2Only {
			mismatchedProps = append(mismatchedProps, fmt.Sprintf("layer2_only (child: %s, parent: %s)", childCfg.Layer2Only, parentCfg.Layer2Only))
		}

		// ARP Suppression (arp_supress)
		if childCfg.ArpSuppression != "" && childCfg.ArpSuppression != parentCfg.ArpSuppression {
			mismatchedProps = append(mismatchedProps, fmt.Sprintf("arp_suppression (child: %s, parent: %s)", childCfg.ArpSuppression, parentCfg.ArpSuppression))
		}

		// Route Target Both (route_target_both)
		if childCfg.RouteTargetBoth != "" && childCfg.RouteTargetBoth != parentCfg.RouteTargetBoth {
			mismatchedProps = append(mismatchedProps, fmt.Sprintf("route_target_both (child: %s, parent: %s)", childCfg.RouteTargetBoth, parentCfg.RouteTargetBoth))
		}

		// Secondary Gateways (secondary_ip_address)
		if childCfg.SecondaryGateway1 != "" && childCfg.SecondaryGateway1 != parentCfg.SecondaryGateway1 {
			mismatchedProps = append(mismatchedProps, fmt.Sprintf("secondary_gateway_1 (child: %s, parent: %s)", childCfg.SecondaryGateway1, parentCfg.SecondaryGateway1))
		}
		if childCfg.SecondaryGateway2 != "" && childCfg.SecondaryGateway2 != parentCfg.SecondaryGateway2 {
			mismatchedProps = append(mismatchedProps, fmt.Sprintf("secondary_gateway_2 (child: %s, parent: %s)", childCfg.SecondaryGateway2, parentCfg.SecondaryGateway2))
		}
		if childCfg.SecondaryGateway3 != "" && childCfg.SecondaryGateway3 != parentCfg.SecondaryGateway3 {
			mismatchedProps = append(mismatchedProps, fmt.Sprintf("secondary_gateway_3 (child: %s, parent: %s)", childCfg.SecondaryGateway3, parentCfg.SecondaryGateway3))
		}
		if childCfg.SecondaryGateway4 != "" && childCfg.SecondaryGateway4 != parentCfg.SecondaryGateway4 {
			mismatchedProps = append(mismatchedProps, fmt.Sprintf("secondary_gateway_4 (child: %s, parent: %s)", childCfg.SecondaryGateway4, parentCfg.SecondaryGateway4))
		}

		// If there are mismatches, report error
		if len(mismatchedProps) > 0 {
			errMsg := fmt.Sprintf("Network '%s': MSD child fabric configuration has parent properties that don't match the parent fabric. "+
				"Parent-managed properties must be identical to the parent fabric(%s) configuration. "+
				"Mismatched properties: %s",
				networkName, parentFabric, strings.Join(mismatchedProps, ", "))
			tflog.Error(ctx, errMsg)
			dg.AddError("Mismatch of config between parent and child attributes", errMsg)
		}
	}
}

// ValidateMsdParentNetworkParameters validates that child-specific properties are not set in MSD parent
func (c *NDFC) ValidateMsdParentNetworkParameters(ctx context.Context, dg *diag.Diagnostics, nw *resource_networks.NDFCNetworksModel) {
	// Validate Network parameters for MSD parent fabric
	// This function validates that child-specific properties are not set in MSD parent
	// Attachment fabric field validation is handled by validateNetworkAttachmentFabricField
	// Note: Caller must ensure this is only called for MSD parent fabrics
	tflog.Debug(ctx, fmt.Sprintf("ValidateMsdParentNetworkParameters: checking %d Networks in fabric %s", len(nw.Networks), nw.FabricName))

	for nwName, entry := range nw.Networks {
		tflog.Debug(ctx, fmt.Sprintf("ValidateMsdParentNetworkParameters: checking Network %s", nwName))
		switch {
		case entry.NetworkTemplateConfig.DhcpRelayServers != nil:
			tflog.Debug(ctx, "Network operation failed, DHCP Relay is managed per child fabric and should not be set in MSD parent")
			dg.AddError("DHCP Relay config not supported for MSD Parent Fabric",
				fmt.Sprintf("In network %s, DHCP Relay config should be done in the MSD child Fabric", entry.NetworkName))
		case entry.NetworkTemplateConfig.DhcpRelayLoopbackId != nil:
			tflog.Debug(ctx, "Network operation failed, DHCP Relay Loopback ID is managed per child fabric and should not be set in MSD parent")
			dg.AddError("DHCP Relay Loopback ID config not supported for MSD Parent Fabric",
				fmt.Sprintf("In network %s, DHCP Relay Loopback ID config should be done in the MSD child Fabric", entry.NetworkName))
		case len(entry.NetworkTemplateConfig.MulticastGroup) != 0:
			tflog.Debug(ctx, "Network operation failed, Multicast Group is managed per child fabric and should not be set in MSD parent")
			dg.AddError("Multicast Group config not supported for MSD Parent Fabric",
				fmt.Sprintf("In network %s, Multicast Group config should be done in the MSD child Fabric", entry.NetworkName))
		case len(entry.NetworkTemplateConfig.Trm) != 0:
			tflog.Debug(ctx, "Network operation failed, TRM is managed per child fabric and should not be set in MSD parent")
			dg.AddError("TRM Enabled config not supported for MSD Parent Fabric",
				fmt.Sprintf("In network %s, TRM Enabled config should be done in the MSD child Fabric", entry.NetworkName))
		case len(entry.NetworkTemplateConfig.Netflow) != 0:
			tflog.Debug(ctx, "Network operation failed, Netflow is managed per child fabric and should not be set in MSD parent")
			dg.AddError("Netflow config not supported for MSD Parent Fabric",
				fmt.Sprintf("In network %s, Netflow config should be done in the MSD child Fabric", entry.NetworkName))
		case len(entry.NetworkTemplateConfig.VlanNetflowMonitor) != 0:
			tflog.Debug(ctx, "Network operation failed, VLAN Netflow Monitor is managed per child fabric and should not be set in MSD parent")
			dg.AddError("VLAN Netflow Monitor config not supported for MSD Parent Fabric",
				fmt.Sprintf("In network %s, VLAN Netflow Monitor config should be done in the MSD child Fabric", entry.NetworkName))
		case len(entry.NetworkTemplateConfig.L3GatwayBorder) != 0:
			tflog.Debug(ctx, "Network operation failed, L3 Gateway Border is managed per child fabric and should not be set in MSD parent")
			dg.AddError("L3 Gateway Border config not supported for MSD Parent Fabric",
				fmt.Sprintf("In network %s, L3 Gateway Border config should be done in the MSD child Fabric", entry.NetworkName))
		case len(entry.NetworkTemplateConfig.IgmpVersion) != 0:
			tflog.Debug(ctx, "Network operation failed, IGMP Version is managed per child fabric and should not be set in MSD parent")
			dg.AddError("IGMP Version config not supported for MSD Parent Fabric",
				fmt.Sprintf("In network %s, IGMP Version config should be done in the MSD child Fabric", entry.NetworkName))
		}
	}
}

// mergeCurrentNetworksWithConfig gets current networks from NDFC and merges them with config values for MSD child fabrics
func (c *NDFC) mergeCurrentNetworksWithConfig(ctx context.Context, dg *diag.Diagnostics, nw *resource_networks.NDFCNetworksModel) {
	// Get current networks from NDFC to prefill values before update
	currentNdfcNetworks, err := c.networksGet(ctx, nw.FabricName)
	if err != nil {
		tflog.Error(ctx, "Error while getting current Networks from NDFC for MSD Child fabric", map[string]interface{}{"Err": err})
		dg.AddError("Network Read Failed", err.Error())
		return
	}

	// Merge current NDFC values with config values for MSD child
	for networkName, networkEntry := range nw.Networks {
		if currentEntry, found := currentNdfcNetworks.Networks[networkName]; found {
			// Start with current NDFC values
			mergedEntry := currentEntry

			// Overlay config values on top (preserve user-configured values)
			if networkEntry.DisplayName != "" {
				mergedEntry.DisplayName = networkEntry.DisplayName
			}
			if networkEntry.NetworkTemplate != "" {
				mergedEntry.NetworkTemplate = networkEntry.NetworkTemplate
			}
			if networkEntry.NetworkExtensionTemplate != "" {
				mergedEntry.NetworkExtensionTemplate = networkEntry.NetworkExtensionTemplate
			}
			if networkEntry.VrfName != "" {
				mergedEntry.VrfName = networkEntry.VrfName
			}

			// Merge NetworkTemplateConfig fields
			if networkEntry.NetworkTemplateConfig.VlanId != nil && !networkEntry.NetworkTemplateConfig.VlanId.IsEmpty() {
				mergedEntry.NetworkTemplateConfig.VlanId = networkEntry.NetworkTemplateConfig.VlanId
			}
			if networkEntry.NetworkTemplateConfig.VlanName != "" {
				mergedEntry.NetworkTemplateConfig.VlanName = networkEntry.NetworkTemplateConfig.VlanName
			}
			if networkEntry.NetworkTemplateConfig.GatewayIpv4Address != "" {
				mergedEntry.NetworkTemplateConfig.GatewayIpv4Address = networkEntry.NetworkTemplateConfig.GatewayIpv4Address
			}
			if networkEntry.NetworkTemplateConfig.GatewayIpv6Address != "" {
				mergedEntry.NetworkTemplateConfig.GatewayIpv6Address = networkEntry.NetworkTemplateConfig.GatewayIpv6Address
			}
			if networkEntry.NetworkTemplateConfig.InterfaceDescription != "" {
				mergedEntry.NetworkTemplateConfig.InterfaceDescription = networkEntry.NetworkTemplateConfig.InterfaceDescription
			}
			if networkEntry.NetworkTemplateConfig.Mtu != nil && !networkEntry.NetworkTemplateConfig.Mtu.IsEmpty() {
				mergedEntry.NetworkTemplateConfig.Mtu = networkEntry.NetworkTemplateConfig.Mtu
			}
			if networkEntry.NetworkTemplateConfig.RoutingTag != nil && !networkEntry.NetworkTemplateConfig.RoutingTag.IsEmpty() {
				mergedEntry.NetworkTemplateConfig.RoutingTag = networkEntry.NetworkTemplateConfig.RoutingTag
			}
			if networkEntry.NetworkTemplateConfig.Layer2Only != "" {
				mergedEntry.NetworkTemplateConfig.Layer2Only = networkEntry.NetworkTemplateConfig.Layer2Only
			}
			if networkEntry.NetworkTemplateConfig.ArpSuppression != "" {
				mergedEntry.NetworkTemplateConfig.ArpSuppression = networkEntry.NetworkTemplateConfig.ArpSuppression
			}
			if networkEntry.NetworkTemplateConfig.RouteTargetBoth != "" {
				mergedEntry.NetworkTemplateConfig.RouteTargetBoth = networkEntry.NetworkTemplateConfig.RouteTargetBoth
			}
			if networkEntry.NetworkTemplateConfig.SecondaryGateway1 != "" {
				mergedEntry.NetworkTemplateConfig.SecondaryGateway1 = networkEntry.NetworkTemplateConfig.SecondaryGateway1
			}
			if networkEntry.NetworkTemplateConfig.SecondaryGateway2 != "" {
				mergedEntry.NetworkTemplateConfig.SecondaryGateway2 = networkEntry.NetworkTemplateConfig.SecondaryGateway2
			}
			if networkEntry.NetworkTemplateConfig.SecondaryGateway3 != "" {
				mergedEntry.NetworkTemplateConfig.SecondaryGateway3 = networkEntry.NetworkTemplateConfig.SecondaryGateway3
			}
			if networkEntry.NetworkTemplateConfig.SecondaryGateway4 != "" {
				mergedEntry.NetworkTemplateConfig.SecondaryGateway4 = networkEntry.NetworkTemplateConfig.SecondaryGateway4
			}

			// Child-specific properties from config (only set if provided in config)
			if networkEntry.NetworkTemplateConfig.DhcpRelayServers != nil {
				mergedEntry.NetworkTemplateConfig.DhcpRelayServers = networkEntry.NetworkTemplateConfig.DhcpRelayServers
			} else {
				mergedEntry.NetworkTemplateConfig.DhcpRelayServers = nil
			}
			if networkEntry.NetworkTemplateConfig.DhcpRelayLoopbackId != nil {
				mergedEntry.NetworkTemplateConfig.DhcpRelayLoopbackId = networkEntry.NetworkTemplateConfig.DhcpRelayLoopbackId
			} else {
				mergedEntry.NetworkTemplateConfig.DhcpRelayLoopbackId = nil
			}
			if len(networkEntry.NetworkTemplateConfig.MulticastGroup) != 0 {
				mergedEntry.NetworkTemplateConfig.MulticastGroup = networkEntry.NetworkTemplateConfig.MulticastGroup
			} else {
				mergedEntry.NetworkTemplateConfig.MulticastGroup = ""
			}
			if len(networkEntry.NetworkTemplateConfig.Trm) != 0 {
				mergedEntry.NetworkTemplateConfig.Trm = networkEntry.NetworkTemplateConfig.Trm
			} else {
				mergedEntry.NetworkTemplateConfig.Trm = ""
			}
			if len(networkEntry.NetworkTemplateConfig.Netflow) != 0 {
				mergedEntry.NetworkTemplateConfig.Netflow = networkEntry.NetworkTemplateConfig.Netflow
			} else {
				mergedEntry.NetworkTemplateConfig.Netflow = ""
			}
			if len(networkEntry.NetworkTemplateConfig.VlanNetflowMonitor) != 0 {
				mergedEntry.NetworkTemplateConfig.VlanNetflowMonitor = networkEntry.NetworkTemplateConfig.VlanNetflowMonitor
			} else {
				mergedEntry.NetworkTemplateConfig.VlanNetflowMonitor = ""
			}
			if len(networkEntry.NetworkTemplateConfig.L3GatwayBorder) != 0 {
				mergedEntry.NetworkTemplateConfig.L3GatwayBorder = networkEntry.NetworkTemplateConfig.L3GatwayBorder
			} else {
				mergedEntry.NetworkTemplateConfig.L3GatwayBorder = ""
			}
			if len(networkEntry.NetworkTemplateConfig.IgmpVersion) != 0 {
				mergedEntry.NetworkTemplateConfig.IgmpVersion = networkEntry.NetworkTemplateConfig.IgmpVersion
			} else {
				mergedEntry.NetworkTemplateConfig.IgmpVersion = ""
			}

			// Preserve attachments from config
			mergedEntry.Attachments = networkEntry.Attachments
			mergedEntry.DeployAttachments = networkEntry.DeployAttachments

			// Update the network entry with merged values
			nw.Networks[networkName] = mergedEntry
		}
	}
}

// processNetworkAttachmentsForMsdCreate handles attachments for MSD fabrics in CREATE flow
// (MSD Parent and MSD Child fabrics - regular fabrics use standard flow in RscCreateNetworks)
// MSD Parent: check NDFC, create attachments that don't exist
// MSD Child: check NDFC, skip if all exist (created by parent), create if missing
func (c *NDFC) processNetworkAttachmentsForMsdCreate(ctx context.Context, dg *diag.Diagnostics,
	planNw *resource_networks.NDFCNetworksModel, isMsdParent bool) error {

	fabricType := "MSD Child"
	if isMsdParent {
		fabricType = "MSD Parent"
	}
	tflog.Info(ctx, fmt.Sprintf("%s: Processing attachments in CREATE flow", fabricType))
	// Step 1: Read current attachments from NDFC
	currentNw := &resource_networks.NDFCNetworksModel{
		FabricName: planNw.FabricName,
		Networks:   make(map[string]resource_networks.NDFCNetworksValue),
	}
	for networkName := range planNw.Networks {
		currentNw.Networks[networkName] = resource_networks.NDFCNetworksValue{
			NetworkName: networkName,
			FabricName:  planNw.FabricName,
			Attachments: make(map[string]rna.NDFCAttachmentsValue),
		}
	}
	err := c.RscGetNetworkAttachments(ctx, currentNw, nil)
	if err != nil {
		tflog.Warn(ctx, fmt.Sprintf("%s: Could not read existing attachments: %v", fabricType, err))
	}
	// Step 2: For MSD Child, check if all attachments already exist (created by parent)
	// If yes, skip attachment processing - only deploy
	if !isMsdParent {
		allAttachmentsExist := true
		for networkName, planNetwork := range planNw.Networks {
			if currentNetwork, exists := currentNw.Networks[networkName]; exists {
				for serial := range planNetwork.Attachments {
					if _, found := currentNetwork.Attachments[serial]; !found {
						allAttachmentsExist = false
						break
					}
				}
			} else {
				allAttachmentsExist = false
			}
			if !allAttachmentsExist {
				break
			}
		}
		if allAttachmentsExist {
			tflog.Info(ctx, "MSD Child: All attachments already exist (created by MSD parent), skipping attachment processing")
			return nil
		}
	}
	// Step 3: Identify attachments that need creation
	attachmentsToCreate := &resource_networks.NDFCNetworksModel{
		FabricName: planNw.FabricName,
		Networks:   make(map[string]resource_networks.NDFCNetworksValue),
	}
	for networkName, planNetwork := range planNw.Networks {
		currentNetwork, networkExistsInCurrent := currentNw.Networks[networkName]
		var newAttachments map[string]rna.NDFCAttachmentsValue
		for serial, planAttachment := range planNetwork.Attachments {
			attachmentExists := false
			if networkExistsInCurrent {
				if _, found := currentNetwork.Attachments[serial]; found {
					attachmentExists = true
				}
			}
			if !attachmentExists {
				if newAttachments == nil {
					newAttachments = make(map[string]rna.NDFCAttachmentsValue)
				}
				newAttachments[serial] = planAttachment
				tflog.Debug(ctx, fmt.Sprintf("%s: Network %s attachment %s needs to be created", fabricType, networkName, serial))
			} else {
				tflog.Debug(ctx, fmt.Sprintf("%s: Network %s attachment %s already exists, skipping", fabricType, networkName, serial))
			}
		}
		if len(newAttachments) > 0 {
			networkEntry := planNetwork
			networkEntry.Attachments = newAttachments
			attachmentsToCreate.Networks[networkName] = networkEntry
		}
	}
	// Step 4: Create new attachments if any
	if len(attachmentsToCreate.Networks) > 0 {
		tflog.Info(ctx, fmt.Sprintf("%s: Creating %d network(s) with new attachments", fabricType, len(attachmentsToCreate.Networks)))
		if isMsdParent {
			// MSD Parent: use RscManageNetAttachmentsForMsdParent
			err := c.RscManageNetAttachmentsForMsdParent(ctx, dg, attachmentsToCreate, planNw.FabricName, false)
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("%s: Failed to create attachments", fabricType))
				return err
			}
		} else {
			// MSD Child: use netAttachmentsAttach
			err := c.netAttachmentsAttach(ctx, attachmentsToCreate)
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("%s: Failed to create attachments", fabricType))
				return err
			}
		}
		tflog.Info(ctx, fmt.Sprintf("%s: Successfully created new attachments", fabricType))
	} else {
		tflog.Info(ctx, fmt.Sprintf("%s: No new attachments to create", fabricType))
	}
	return nil
}

// processNetworkAttachmentsForMsdUpdate handles attachment updates for MSD fabrics
// 1. DELETE (Detach): attachments in state but not in plan - POST directly with Deployment="false"
// 2. UPDATE: existing attachments with modifications - POST directly
// 3. CREATE (NewEntry): new attachments - check NDFC first for MSD (parent may have created)
func (c *NDFC) processNetworkAttachmentsForMsdUpdate(ctx context.Context, dg *diag.Diagnostics,
	planNw *resource_networks.NDFCNetworksModel,
	stateNw *resource_networks.NDFCNetworksModel,
	putNw *resource_networks.NDFCNetworksModel,
	isMsdParent bool) error {

	fabricType := "MSD Child"
	if isMsdParent {
		fabricType = "MSD Parent"
	}

	// Get attachment diff using existing function
	netAttachActions := c.networkAttachmentsGetDiff(ctx, dg, planNw, stateNw, putNw)
	updateNwAttach := netAttachActions["update"].(*rna.NDFCNetworkAttachments)

	if updateNwAttach == nil || len(updateNwAttach.NetworkAttachments) == 0 {
		tflog.Info(ctx, fmt.Sprintf("%s: No attachment changes detected", fabricType))
		return nil
	}

	// For MSD, we need to filter NewEntry attachments through NDFC check
	// Detach and Update can be posted directly
	tflog.Info(ctx, fmt.Sprintf("%s: Processing %d network attachment changes", fabricType, len(updateNwAttach.NetworkAttachments)))

	// Step 1: Read current attachments from NDFC (for NewEntry filtering)
	currentNw := &resource_networks.NDFCNetworksModel{
		FabricName: planNw.FabricName,
		Networks:   make(map[string]resource_networks.NDFCNetworksValue),
	}
	for networkName := range planNw.Networks {
		currentNw.Networks[networkName] = resource_networks.NDFCNetworksValue{
			NetworkName: networkName,
			FabricName:  planNw.FabricName,
			Attachments: make(map[string]rna.NDFCAttachmentsValue),
		}
	}
	err := c.RscGetNetworkAttachments(ctx, currentNw, nil)
	if err != nil {
		tflog.Warn(ctx, fmt.Sprintf("%s: Could not read existing attachments: %v", fabricType, err))
	}

	// Step 2: Filter attachments - skip NewEntry if already exists in NDFC
	filteredAttachments := &rna.NDFCNetworkAttachments{
		FabricName: updateNwAttach.FabricName,
	}

	for _, nwAttach := range updateNwAttach.NetworkAttachments {
		networkName := nwAttach.NetworkName
		currentNetwork, networkExistsInCurrent := currentNw.Networks[networkName]

		var filteredAttachList []rna.NDFCAttachmentsValue
		for _, attachVal := range nwAttach.Attachments {
			// Check if this is a NewEntry that needs NDFC filtering
			if attachVal.UpdateAction&NewEntry > 0 {
				attachmentExists := false
				if networkExistsInCurrent {
					if existingAttachment, found := currentNetwork.Attachments[attachVal.SerialNumber]; found {
						// Only consider as existing if it's actually attached (not filtered)
						if !existingAttachment.FilterThisValue {
							attachmentExists = true
						}
					}
				}
				if attachmentExists {
					tflog.Debug(ctx, fmt.Sprintf("%s: Network %s attachment %s already exists in NDFC, skipping create",
						fabricType, networkName, attachVal.SerialNumber))
					continue
				}
				tflog.Debug(ctx, fmt.Sprintf("%s: Network %s attachment %s is new, will create",
					fabricType, networkName, attachVal.SerialNumber))
			} else if attachVal.UpdateAction&Detach > 0 {
				tflog.Debug(ctx, fmt.Sprintf("%s: Network %s attachment %s will be detached",
					fabricType, networkName, attachVal.SerialNumber))
			} else if attachVal.UpdateAction&Update > 0 {
				tflog.Debug(ctx, fmt.Sprintf("%s: Network %s attachment %s will be updated",
					fabricType, networkName, attachVal.SerialNumber))
			}
			filteredAttachList = append(filteredAttachList, attachVal)
		}

		if len(filteredAttachList) > 0 {
			filteredAttachments.NetworkAttachments = append(filteredAttachments.NetworkAttachments,
				rna.NDFCNetworkAttachmentsPayload{
					NetworkName: networkName,
					Attachments: filteredAttachList,
				})
		}
	}

	// Step 3: Post filtered attachments
	if len(filteredAttachments.NetworkAttachments) == 0 {
		tflog.Info(ctx, fmt.Sprintf("%s: No attachment changes to apply after filtering", fabricType))
		return nil
	}

	tflog.Info(ctx, fmt.Sprintf("%s: Posting %d network attachment changes", fabricType, len(filteredAttachments.NetworkAttachments)))

	// Convert IP keys to serial numbers before API call
	c.netAttachmentSerialRemapFromPayload(ctx, filteredAttachments)

	data, err := json.Marshal(filteredAttachments.NetworkAttachments)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("%s: Error marshalling attachments: %v", fabricType, err))
		return err
	}

	err = c.netAttachmentsPostPayload(ctx, planNw.FabricName, data)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("%s: Error updating network attachments: %v", fabricType, err))
		return err
	}

	tflog.Info(ctx, fmt.Sprintf("%s: Successfully processed attachment updates", fabricType))
	return nil
}

// RscManageNetAttachmentsForMsdParent manages network attachments for MSD parent fabric
// It groups attachments by child fabric and processes them separately
func (c *NDFC) RscManageNetAttachmentsForMsdParent(ctx context.Context, dg *diag.Diagnostics, nw *resource_networks.NDFCNetworksModel, fabricName string, isDelete bool) error {
	// Group attachments by child fabric
	attachmentsByFabric := make(map[string]*resource_networks.NDFCNetworksModel)

	for networkName, networkEntry := range nw.Networks {
		for serialNumber, attachment := range networkEntry.Attachments {
			childFabric := attachment.Fabric
			if childFabric == "" {
				childFabric = attachment.FabricName
			}

			if _, exists := attachmentsByFabric[childFabric]; !exists {
				attachmentsByFabric[childFabric] = &resource_networks.NDFCNetworksModel{
					FabricName: childFabric,
					Networks:   make(map[string]resource_networks.NDFCNetworksValue),
				}
			}

			// Add or update network entry for this child fabric
			if _, exists := attachmentsByFabric[childFabric].Networks[networkName]; !exists {
				newEntry := networkEntry
				newEntry.Attachments = make(map[string]rna.NDFCAttachmentsValue)
				attachmentsByFabric[childFabric].Networks[networkName] = newEntry
			}

			// Add attachment to the child fabric's network entry (use key as serial number)
			childNet := attachmentsByFabric[childFabric].Networks[networkName]
			childNet.Attachments[serialNumber] = attachment
			attachmentsByFabric[childFabric].Networks[networkName] = childNet
		}
	}

	// Process attachments for each child fabric
	for childFabric, childNw := range attachmentsByFabric {
		if isDelete {
			tflog.Info(ctx, fmt.Sprintf("MSD Parent: Deleting attachments for child fabric %s", childFabric))
			err := c.RscDeleteNetAttachments(ctx, dg, childNw)
			if err != nil {
				return err
			}
		} else {
			tflog.Info(ctx, fmt.Sprintf("MSD Parent: Creating attachments for child fabric %s", childFabric))
			err := c.netAttachmentsAttach(ctx, childNw)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
