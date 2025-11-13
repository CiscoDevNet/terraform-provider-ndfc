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
	"maps"
	"terraform-provider-ndfc/internal/provider/resources/resource_vrf_bulk"
	. "terraform-provider-ndfc/internal/provider/types"

	api "terraform-provider-ndfc/internal/provider/ndfc/api"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// vrfApplyMsdChildFabricAttributes updates MSD child fabric attributes for each VRF and child fabric combination
// by applying the configured values to the corresponding child fabrics in NDFC.
func (c NDFC) vrfApplyMsdChildFabricAttributes(ctx context.Context, dg *diag.Diagnostics, vrf *resource_vrf_bulk.NDFCVrfBulkModel) {
	tflog.Info(ctx, fmt.Sprintf("Beginning MSD child fabric attributes update for fabric %s", vrf.FabricName))

	// Walk through each VRF
	for vrfName, vrfEntry := range vrf.Vrfs {
		// Check if this VRF has MSD child fabric attributes configured
		if len(vrfEntry.MsdChildFabricAttributes) == 0 {
			tflog.Debug(ctx, fmt.Sprintf("VRF %s has no MSD child fabric attributes, skipping", vrfName))
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Processing MSD child fabric attributes for VRF %s", vrfName))

		// Apply configured values for explicitly specified child fabrics
		for childFabricName, childAttrs := range vrfEntry.MsdChildFabricAttributes {
			tflog.Debug(ctx, fmt.Sprintf("Applying configured attributes for VRF %s in child fabric %s", vrfName, childFabricName))
			err := c.vrfUpdateMsdChildFabricAttributes(ctx, vrfName, childFabricName, &childAttrs, "update")
			if err != nil {
				dg.AddError(
					fmt.Sprintf("MSD child fabric attributes update failed for VRF %s in child fabric %s", vrfName, childFabricName),
					err.Error(),
				)
				return
			}
		}

	}

	tflog.Info(ctx, "MSD child fabric attributes update completed successfully")
}

// vrfUpdateMsdChildFabricAttributes performs a PUT operation to update VRF attributes in a child fabric
// This is a common helper used by both update and delete operations
func (c NDFC) vrfUpdateMsdChildFabricAttributes(ctx context.Context, vrfName, childFabricName string,
	childAttrs *resource_vrf_bulk.NDFCMsdChildFabricAttributesValue, operation string) error {

	tflog.Debug(ctx, fmt.Sprintf("Fetching existing VRF %s from child fabric %s for %s", vrfName, childFabricName, operation))

	// Fetch existing VRF payload from the child fabric
	existingVrfPayloadInterface, err := c.vrfGetFromChildFabric(ctx, childFabricName, vrfName)
	if err != nil {
		return fmt.Errorf("failed to fetch VRF %s from child fabric %s: %w", vrfName, childFabricName, err)
	}

	// Type assert to map[string]interface{}
	existingVrfPayload, ok := existingVrfPayloadInterface.(map[string]any)
	if !ok {
		return fmt.Errorf("invalid VRF payload type for VRF %s from child fabric %s", vrfName, childFabricName)
	}

	tflog.Debug(ctx, fmt.Sprintf("Successfully fetched existing VRF %s from child fabric %s", vrfName, childFabricName))

	// Build the payload by merging existing data with new values
	payload := c.buildMsdChildFabricPayload(ctx, childAttrs, existingVrfPayload)

	tflog.Debug(ctx, fmt.Sprintf("Performing %s operation for VRF %s in child fabric %s", operation, vrfName, childFabricName))

	// Marshal the payload
	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal failed for VRF %s, child fabric %s: %w", vrfName, childFabricName, err)
	}

	tflog.Debug(ctx, fmt.Sprintf("MSD child fabric attributes payload for VRF %s, child fabric %s: %s", vrfName, childFabricName, string(data)))

	// Use the VRF API to update - using the child fabric name as the fabric context
	vrfObj := api.NewVrfAPI(childFabricName, c.GetLock(ResourceVrfBulk), &c.apiClient)
	vrfObj.PutVrf = vrfName

	res, err := vrfObj.Put(data)
	if err != nil {
		return fmt.Errorf("error %v, response %s", err, res.Str)
	}

	tflog.Info(ctx, fmt.Sprintf("MSD child fabric attributes %s successful for VRF %s in child fabric %s. Message: %s", operation, vrfName, childFabricName, res.Str))
	return nil
}

// buildMsdChildFabricPayload builds the payload for updating MSD child fabric attributes
// by starting with the existing VRF payload and overriding with user-provided values
func (c NDFC) buildMsdChildFabricPayload(ctx context.Context, childAttrs *resource_vrf_bulk.NDFCMsdChildFabricAttributesValue, existingVrfPayload map[string]interface{}) map[string]interface{} {
	// Start with a copy of the existing payload
	payload := make(map[string]any)
	maps.Copy(payload, existingVrfPayload)

	// Get the existing vrfTemplateConfig
	vrfTemplateConfig, ok := payload["vrfTemplateConfig"].(map[string]any)
	if !ok {
		// If it doesn't exist or is not a map, create a new one
		vrfTemplateConfig = make(map[string]any)
	}

	// Override with user-provided child fabric attributes
	if childAttrs.AdvertiseHostRoutes != "" {
		vrfTemplateConfig["advertiseHostRouteFlag"] = childAttrs.AdvertiseHostRoutes
	}
	if childAttrs.AdvertiseDefaultRoute != "" {
		vrfTemplateConfig["advertiseDefaultRouteFlag"] = childAttrs.AdvertiseDefaultRoute
	}
	if childAttrs.ConfigureStaticDefaultRoute != "" {
		vrfTemplateConfig["configureStaticDefaultRouteFlag"] = childAttrs.ConfigureStaticDefaultRoute
	}
	if childAttrs.BgpPassword != "" {
		vrfTemplateConfig["bgpPassword"] = childAttrs.BgpPassword
	}
	if childAttrs.BgpPasswordType != "" {
		vrfTemplateConfig["bgpPasswordKeyType"] = childAttrs.BgpPasswordType
	}
	if childAttrs.Netflow != "" {
		vrfTemplateConfig["ENABLE_NETFLOW"] = childAttrs.Netflow
	}
	if childAttrs.NetflowMonitor != "" {
		vrfTemplateConfig["NETFLOW_MONITOR"] = childAttrs.NetflowMonitor
	}
	if childAttrs.Trm != "" {
		vrfTemplateConfig["trmEnabled"] = childAttrs.Trm
	}
	if childAttrs.TrmBgwMsite != "" {
		vrfTemplateConfig["trmBGWMSiteEnabled"] = childAttrs.TrmBgwMsite
	}
	if childAttrs.NoRp != "" {
		vrfTemplateConfig["isRPAbsent"] = childAttrs.NoRp
	}
	if childAttrs.RpAddress != "" {
		vrfTemplateConfig["rpAddress"] = childAttrs.RpAddress
	}
	if childAttrs.RpLoopbackId != nil {
		vrfTemplateConfig["loopbackNumber"] = *childAttrs.RpLoopbackId
	}
	if childAttrs.UnderlayMulticastAddress != "" {
		vrfTemplateConfig["L3VniMcastGroup"] = childAttrs.UnderlayMulticastAddress
	}
	if childAttrs.OverlayMulticastGroups != "" {
		vrfTemplateConfig["multicastGroup"] = childAttrs.OverlayMulticastGroups
	}
	if childAttrs.RouteTargetImportMvpn != "" {
		vrfTemplateConfig["routeTargetImportMvpn"] = childAttrs.RouteTargetImportMvpn
	}
	if childAttrs.RouteTargetExportMvpn != "" {
		vrfTemplateConfig["routeTargetExportMvpn"] = childAttrs.RouteTargetExportMvpn
	}

	// Put the updated vrfTemplateConfig back
	payload["vrfTemplateConfig"] = vrfTemplateConfig

	return payload
}

// vrfGetFromChildFabric fetches VRF payloads from a specific child fabric
// If vrfName is empty, returns all VRFs as []map[string]interface{}
// If vrfName is provided, returns the specific VRF as map[string]interface{} (or nil if not found)
// The return type is interface{} - caller should type assert based on whether vrfName was provided
func (c NDFC) vrfGetFromChildFabric(ctx context.Context, childFabricName string, vrfName string) (interface{}, error) {
	vrfObj := api.NewVrfAPI(childFabricName, c.GetLock(ResourceVrfBulk), &c.apiClient)

	// Get all VRFs from the child fabric
	res, err := vrfObj.Get()
	if err != nil {
		return nil, fmt.Errorf("failed to get VRFs from child fabric %s: %w", childFabricName, err)
	}

	// Unmarshal as generic map to preserve all fields
	var vrfPayloads []map[string]interface{}
	err = json.Unmarshal(res, &vrfPayloads)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal VRF response from child fabric %s: %w", childFabricName, err)
	}

	tflog.Debug(ctx, fmt.Sprintf("Found %d VRFs in child fabric %s", len(vrfPayloads), childFabricName))

	// If no specific VRF name provided, return all VRFs
	if vrfName == "" {
		return vrfPayloads, nil
	}

	// Find the specific VRF
	for _, vrfPayload := range vrfPayloads {
		if vrfPayload["vrfName"] == vrfName {
			tflog.Debug(ctx, fmt.Sprintf("Found VRF %s in child fabric %s", vrfName, childFabricName))
			return vrfPayload, nil
		}
	}

	// VRF not found in this child fabric
	tflog.Debug(ctx, fmt.Sprintf("VRF %s not found in child fabric %s", vrfName, childFabricName))
	return nil, nil
}

// vrfFetchMsdChildFabricAttributes fetches the actual MSD child fabric attributes
// from child fabrics and populates them in the provided VRF model
// Note: This function should only be called if the fabric is MSD type (checked at call site)
// Fetches all child fabric attributes for all VRFs from NDFC
// ndVrfs: the VRF model to populate with fetched data (typically from NDFC API read)
func (c NDFC) vrfFetchMsdChildFabricAttributes(ctx context.Context, dg *diag.Diagnostics, ndVrfs *resource_vrf_bulk.NDFCVrfBulkModel, parentFabricName string) {
	tflog.Info(ctx, fmt.Sprintf("Fetching MSD child fabric attributes for parent fabric %s", parentFabricName))

	// Get child fabrics for this MSD parent
	childFabrics := c.GetMsdChildFabricAssociations(ctx, dg, parentFabricName)
	if dg.HasError() {
		return
	}

	if len(childFabrics) == 0 {
		tflog.Debug(ctx, fmt.Sprintf("No child fabrics found for MSD parent fabric %s", parentFabricName))
		return
	}

	tflog.Info(ctx, fmt.Sprintf("Found %d child fabric(s) for MSD parent %s: %v", len(childFabrics), parentFabricName, childFabrics))

	// Get all VRFs from each child fabric and process them
	for _, childFabricName := range childFabrics {
		tflog.Debug(ctx, fmt.Sprintf("Fetching all VRFs from child fabric %s", childFabricName))

		vrfPayloadsInterface, err := c.vrfGetFromChildFabric(ctx, childFabricName, "")
		if err != nil {
			tflog.Warn(ctx, fmt.Sprintf("Could not fetch VRFs from child fabric %s: %v", childFabricName, err))
			continue
		}

		// Type assert to []map[string]interface{}
		vrfPayloads, ok := vrfPayloadsInterface.([]map[string]any)
		if !ok {
			tflog.Warn(ctx, fmt.Sprintf("Invalid VRF payloads type from child fabric %s", childFabricName))
			continue
		}

		// Process each VRF found in this child fabric
		for _, vrfPayload := range vrfPayloads {
			vrfName, ok := vrfPayload["vrfName"].(string)
			if !ok || vrfName == "" {
				tflog.Warn(ctx, fmt.Sprintf("Invalid VRF name in child fabric %s", childFabricName))
				continue
			}

			// Get VRF entry from main VRF list (from parent fabric)
			vrfEntry, exists := ndVrfs.Vrfs[vrfName]
			if !exists {
				tflog.Debug(ctx, fmt.Sprintf("VRF %s found in child fabric %s but not in parent fabric %s, skipping", vrfName, childFabricName, parentFabricName))
				continue
			}

			// Initialize the map if not already done
			if vrfEntry.MsdChildFabricAttributes == nil {
				vrfEntry.MsdChildFabricAttributes = make(map[string]resource_vrf_bulk.NDFCMsdChildFabricAttributesValue)
			}

			tflog.Debug(ctx, fmt.Sprintf("Processing VRF %s from child fabric %s", vrfName, childFabricName))

			// Process this VRF payload
			c.processVrfPayloadForChildFabric(ctx, childFabricName, vrfName, vrfPayload, &vrfEntry)

			// Put the updated vrfEntry back into the map
			ndVrfs.Vrfs[vrfName] = vrfEntry
		}
	}

	tflog.Info(ctx, "Completed fetching MSD child fabric attributes")
}

// processVrfPayloadForChildFabric processes a NDFC VRF payload from a child fabric
// and populates the MSD child fabric attributes for the given VRF entry
func (c NDFC) processVrfPayloadForChildFabric(ctx context.Context, childFabricName string, vrfName string, vrfPayload map[string]interface{}, vrfEntry *resource_vrf_bulk.NDFCVrfsValue) {
	// Extract vrfTemplateConfig
	vrfTemplateConfig, ok := vrfPayload["vrfTemplateConfig"].(map[string]any)
	if !ok {
		tflog.Warn(ctx, fmt.Sprintf("Could not extract vrfTemplateConfig for VRF %s from child fabric %s", vrfName, childFabricName))
		return
	}

	// Build the child fabric attributes from the fetched data
	childAttrs := resource_vrf_bulk.NDFCMsdChildFabricAttributesValue{
		FabricName: childFabricName,
	}

	// Extract and populate the attributes - handle both bool and string types
	// Always capture all values to show complete current state

	// advertiseHostRouteFlag
	if val, ok := vrfTemplateConfig["advertiseHostRouteFlag"].(bool); ok {
		if val {
			childAttrs.AdvertiseHostRoutes = "true"
		} else {
			childAttrs.AdvertiseHostRoutes = "false"
		}
	} else if val, ok := vrfTemplateConfig["advertiseHostRouteFlag"].(string); ok {
		childAttrs.AdvertiseHostRoutes = val
	}

	// advertiseDefaultRouteFlag
	if val, ok := vrfTemplateConfig["advertiseDefaultRouteFlag"].(bool); ok {
		if val {
			childAttrs.AdvertiseDefaultRoute = "true"
		} else {
			childAttrs.AdvertiseDefaultRoute = "false"
		}
	} else if val, ok := vrfTemplateConfig["advertiseDefaultRouteFlag"].(string); ok {
		childAttrs.AdvertiseDefaultRoute = val
	}

	// configureStaticDefaultRouteFlag
	if val, ok := vrfTemplateConfig["configureStaticDefaultRouteFlag"].(bool); ok {
		if val {
			childAttrs.ConfigureStaticDefaultRoute = "true"
		} else {
			childAttrs.ConfigureStaticDefaultRoute = "false"
		}
	} else if val, ok := vrfTemplateConfig["configureStaticDefaultRouteFlag"].(string); ok {
		childAttrs.ConfigureStaticDefaultRoute = val
	}

	// bgpPassword
	if val, ok := vrfTemplateConfig["bgpPassword"].(string); ok {
		childAttrs.BgpPassword = val
	}

	// bgpPasswordKeyType
	if val, ok := vrfTemplateConfig["bgpPasswordKeyType"].(string); ok {
		childAttrs.BgpPasswordType = val
	}

	// ENABLE_NETFLOW
	if val, ok := vrfTemplateConfig["ENABLE_NETFLOW"].(bool); ok {
		if val {
			childAttrs.Netflow = "true"
		} else {
			childAttrs.Netflow = "false"
		}
	} else if val, ok := vrfTemplateConfig["ENABLE_NETFLOW"].(string); ok {
		childAttrs.Netflow = val
	}

	// NETFLOW_MONITOR
	if val, ok := vrfTemplateConfig["NETFLOW_MONITOR"].(string); ok {
		childAttrs.NetflowMonitor = val
	}

	// trmEnabled
	if val, ok := vrfTemplateConfig["trmEnabled"].(bool); ok {
		if val {
			childAttrs.Trm = "true"
		} else {
			childAttrs.Trm = "false"
		}
	} else if val, ok := vrfTemplateConfig["trmEnabled"].(string); ok {
		childAttrs.Trm = val
	}

	// trmBGWMSiteEnabled
	if val, ok := vrfTemplateConfig["trmBGWMSiteEnabled"].(bool); ok {
		if val {
			childAttrs.TrmBgwMsite = "true"
		} else {
			childAttrs.TrmBgwMsite = "false"
		}
	} else if val, ok := vrfTemplateConfig["trmBGWMSiteEnabled"].(string); ok {
		childAttrs.TrmBgwMsite = val
	}

	// isRPAbsent
	if val, ok := vrfTemplateConfig["isRPAbsent"].(bool); ok {
		if val {
			childAttrs.NoRp = "true"
		} else {
			childAttrs.NoRp = "false"
		}
	} else if val, ok := vrfTemplateConfig["isRPAbsent"].(string); ok {
		childAttrs.NoRp = val
	}

	// rpAddress
	if val, ok := vrfTemplateConfig["rpAddress"].(string); ok {
		childAttrs.RpAddress = val
	}

	// loopbackNumber - handle numeric values
	if val, ok := vrfTemplateConfig["loopbackNumber"]; ok && val != nil {
		switch v := val.(type) {
		case float64:
			intVal := Int64Custom(int64(v))
			childAttrs.RpLoopbackId = &intVal
		case int:
			intVal := Int64Custom(int64(v))
			childAttrs.RpLoopbackId = &intVal
		case int64:
			intVal := Int64Custom(v)
			childAttrs.RpLoopbackId = &intVal
		}
	}

	// L3VniMcastGroup
	if val, ok := vrfTemplateConfig["L3VniMcastGroup"].(string); ok {
		childAttrs.UnderlayMulticastAddress = val
	}

	// multicastGroup
	if val, ok := vrfTemplateConfig["multicastGroup"].(string); ok {
		childAttrs.OverlayMulticastGroups = val
	}

	// routeTargetImportMvpn
	if val, ok := vrfTemplateConfig["routeTargetImportMvpn"].(string); ok {
		childAttrs.RouteTargetImportMvpn = val
	}

	// routeTargetExportMvpn
	if val, ok := vrfTemplateConfig["routeTargetExportMvpn"].(string); ok {
		childAttrs.RouteTargetExportMvpn = val
	}

	// Always store child fabric attributes to show complete current state
	// Since msd_child_fabric_attributes is computed and optional, it should always reflect the actual NDFC state
	vrfEntry.MsdChildFabricAttributes[childFabricName] = childAttrs
	tflog.Debug(ctx, fmt.Sprintf("Successfully fetched and populated MSD child fabric attributes for VRF %s from child fabric %s", vrfName, childFabricName))
}

// vrfGetMsdChildFabricAttributesDiff computes the differences in MSD child fabric attributes
// between plan and state for VRFs that exist in both (existing VRFs)
// Returns maps of child fabric attributes to add, update, and delete
// Note: checkVrfs parameter is typically 'plan' to check all VRFs, not just those with main attribute changes
func (c NDFC) vrfGetMsdChildFabricAttributesDiff(ctx context.Context,
	planVrfs *resource_vrf_bulk.NDFCVrfBulkModel,
	stateVrfs *resource_vrf_bulk.NDFCVrfBulkModel,
	checkVrfs *resource_vrf_bulk.NDFCVrfBulkModel) (
	updateAttrs *resource_vrf_bulk.NDFCVrfBulkModel,
	deleteChildFabrics map[string][]string) {

	tflog.Debug(ctx, "Computing MSD child fabric attributes differences")

	updateAttrs = &resource_vrf_bulk.NDFCVrfBulkModel{
		FabricName: planVrfs.FabricName,
		Vrfs:       make(map[string]resource_vrf_bulk.NDFCVrfsValue),
	}
	deleteChildFabrics = make(map[string][]string)

	// Check all VRFs in checkVrfs (typically all VRFs in plan) for MSD child fabric attribute changes
	for vrfName := range checkVrfs.Vrfs {
		stateVrfEntry, stateExists := stateVrfs.Vrfs[vrfName]
		planVrfEntry, planExists := planVrfs.Vrfs[vrfName]

		if !planExists || !stateExists {
			continue
		}

		// Track child fabrics to update (new or modified)
		var needsUpdate bool
		updatedVrfEntry := resource_vrf_bulk.NDFCVrfsValue{
			VrfName:                  vrfName,
			FabricName:               planVrfs.FabricName,
			MsdChildFabricAttributes: make(map[string]resource_vrf_bulk.NDFCMsdChildFabricAttributesValue),
		}

		// Check for new or modified child fabric attributes in plan
		for childFabricName, planChildAttrs := range planVrfEntry.MsdChildFabricAttributes {
			stateChildAttrs, existsInState := stateVrfEntry.MsdChildFabricAttributes[childFabricName]

			if !existsInState {
				// New child fabric attribute - add to update list
				tflog.Debug(ctx, fmt.Sprintf("VRF %s: New MSD child fabric %s detected", vrfName, childFabricName))
				updatedVrfEntry.MsdChildFabricAttributes[childFabricName] = planChildAttrs
				needsUpdate = true
			} else {
				// Child fabric exists in both - check if attributes changed
				if !c.msdChildFabricAttributesEqual(planChildAttrs, stateChildAttrs) {
					tflog.Debug(ctx, fmt.Sprintf("VRF %s: MSD child fabric %s attributes changed", vrfName, childFabricName))
					updatedVrfEntry.MsdChildFabricAttributes[childFabricName] = planChildAttrs
					needsUpdate = true
				}
			}
		}

		// Check for deleted child fabric attributes (exists in state but not in plan)
		for childFabricName := range stateVrfEntry.MsdChildFabricAttributes {
			if _, existsInPlan := planVrfEntry.MsdChildFabricAttributes[childFabricName]; !existsInPlan {
				tflog.Debug(ctx, fmt.Sprintf("VRF %s: MSD child fabric %s removed from config", vrfName, childFabricName))
				deleteChildFabrics[vrfName] = append(deleteChildFabrics[vrfName], childFabricName)
			}
		}

		// If there are updates for this VRF, add it to the result
		if needsUpdate {
			updateAttrs.Vrfs[vrfName] = updatedVrfEntry
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("MSD child fabric attributes diff computed: %d VRFs to update, %d VRFs with deletions",
		len(updateAttrs.Vrfs), len(deleteChildFabrics)))

	return updateAttrs, deleteChildFabrics
}

// msdChildFabricAttributesEqual compares two MSD child fabric attribute values
func (c NDFC) msdChildFabricAttributesEqual(a, b resource_vrf_bulk.NDFCMsdChildFabricAttributesValue) bool {
	return a.AdvertiseHostRoutes == b.AdvertiseHostRoutes &&
		a.AdvertiseDefaultRoute == b.AdvertiseDefaultRoute &&
		a.ConfigureStaticDefaultRoute == b.ConfigureStaticDefaultRoute &&
		a.BgpPassword == b.BgpPassword &&
		a.BgpPasswordType == b.BgpPasswordType &&
		a.Netflow == b.Netflow &&
		a.NetflowMonitor == b.NetflowMonitor &&
		a.Trm == b.Trm &&
		a.TrmBgwMsite == b.TrmBgwMsite &&
		a.NoRp == b.NoRp &&
		a.RpAddress == b.RpAddress &&
		((a.RpLoopbackId == nil && b.RpLoopbackId == nil) ||
			(a.RpLoopbackId != nil && b.RpLoopbackId != nil && *a.RpLoopbackId == *b.RpLoopbackId)) &&
		a.UnderlayMulticastAddress == b.UnderlayMulticastAddress &&
		a.OverlayMulticastGroups == b.OverlayMulticastGroups &&
		a.RouteTargetImportMvpn == b.RouteTargetImportMvpn &&
		a.RouteTargetExportMvpn == b.RouteTargetExportMvpn
}

// getDefaultMsdChildFabricAttributes returns the default values for child fabric attributes
// These default values are used when a child fabric entry is removed from the configuration,
// which means the values should be reset to defaults (not deleted from the VRF)
// Based on the default vrfTemplateConfig from NDFC API
func (c NDFC) getDefaultMsdChildFabricAttributes() resource_vrf_bulk.NDFCMsdChildFabricAttributesValue {
	return resource_vrf_bulk.NDFCMsdChildFabricAttributesValue{
		AdvertiseDefaultRoute:       "true",
		AdvertiseHostRoutes:         "false",
		ConfigureStaticDefaultRoute: "true",
		BgpPassword:                 "",
		BgpPasswordType:             "",
		Netflow:                     "false",
		NetflowMonitor:              "",
		Trm:                         "false",
		TrmBgwMsite:                 "false",
		NoRp:                        "false",
		RpAddress:                   "",
		RpLoopbackId:                nil,
		UnderlayMulticastAddress:    "",
		OverlayMulticastGroups:      "",
		RouteTargetImportMvpn:       "",
		RouteTargetExportMvpn:       "",
	}
}

// vrfDeleteMsdChildFabricAttributes deletes MSD child fabric attributes
// by resetting them to default values in the child fabric
// This is called when child fabrics are removed from configuration
// Note: Deletion here means resetting to default values, not removing the VRF
func (c NDFC) vrfDeleteMsdChildFabricAttributes(ctx context.Context, dg *diag.Diagnostics,
	deleteMap map[string][]string) {

	if len(deleteMap) == 0 {
		tflog.Debug(ctx, "No MSD child fabric attributes to delete")
		return
	}

	tflog.Info(ctx, "Beginning MSD child fabric attributes deletion (reset to defaults)")

	for vrfName, childFabrics := range deleteMap {
		for _, childFabricName := range childFabrics {
			// Reset to default values using the default child fabric attributes
			defaultAttrs := c.getDefaultMsdChildFabricAttributes()
			tflog.Debug(ctx, fmt.Sprintf("Resetting VRF %s in child fabric %s to default values", vrfName, childFabricName))

			err := c.vrfUpdateMsdChildFabricAttributes(ctx, vrfName, childFabricName, &defaultAttrs, "delete")
			if err != nil {
				dg.AddError(
					fmt.Sprintf("MSD child fabric attributes deletion failed for VRF %s in child fabric %s", vrfName, childFabricName),
					err.Error(),
				)
				return
			}
			tflog.Info(ctx, fmt.Sprintf("Successfully reset VRF %s in child fabric %s to default values", vrfName, childFabricName))
		}
	}

	tflog.Info(ctx, "MSD child fabric attributes deletion (reset to defaults) completed successfully")
}

// vrfFilterVrfsWithMsdChildFabricAttributes filters VRFs that have MSD child fabric attributes
// Returns a new NDFCVrfBulkModel containing only VRFs with MSD child fabric attributes
func (c NDFC) vrfFilterVrfsWithMsdChildFabricAttributes(vrfs *resource_vrf_bulk.NDFCVrfBulkModel) *resource_vrf_bulk.NDFCVrfBulkModel {
	filtered := &resource_vrf_bulk.NDFCVrfBulkModel{
		FabricName: vrfs.FabricName,
		Vrfs:       make(map[string]resource_vrf_bulk.NDFCVrfsValue),
	}

	for vrfName, vrfEntry := range vrfs.Vrfs {
		if len(vrfEntry.MsdChildFabricAttributes) > 0 {
			filtered.Vrfs[vrfName] = vrfEntry
		}
	}

	return filtered
}
