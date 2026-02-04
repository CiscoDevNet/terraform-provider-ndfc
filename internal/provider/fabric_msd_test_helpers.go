// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	helper "terraform-provider-ndfc/internal/provider/testing"
)

// initializeFabricsAndInventory initializes the VXLAN fabrics and inventory devices configuration
// This creates the child fabrics and adds switches from testbed.yaml
func initializeFabricsAndInventory(cfg *helper.Config, numFabrics int) *helper.MsdComprehensiveTestConfig {
	testConfig := &helper.MsdComprehensiveTestConfig{
		VxlanFabrics:     make([]helper.VxlanFabricConfig, 0),
		InventoryDevices: make([]helper.InventoryDevicesConfig, 0),
	}

	// Create VXLAN fabrics for each child fabric that has a BGP AS configured
	for i := 0; i < numFabrics && i < len(cfg.NDFC.Msd.ChildFabrics); i++ {
		childFabric := cfg.NDFC.Msd.ChildFabrics[i]
		if childFabric.BgpAs == "" {
			continue
		}

		// Add VXLAN fabric
		testConfig.VxlanFabrics = append(testConfig.VxlanFabrics, helper.VxlanFabricConfig{
			ResourceName: fmt.Sprintf("vxlan_fabric_%d", i+1),
			FabricName:   childFabric.Name,
			BgpAs:        childFabric.BgpAs,
			Deploy:       false,
		})

		// Add inventory devices if this fabric has them configured
		if len(childFabric.InventoryDevices) > 0 {
			devices := make(map[string]helper.InventoryDeviceEntry)
			for _, dev := range childFabric.InventoryDevices {
				devices[dev.Device] = helper.InventoryDeviceEntry{
					Role:                  dev.Role,
					DiscoveryType:         "discover",
					DiscoveryAuthProtocol: "md5",
				}
			}

			testConfig.InventoryDevices = append(testConfig.InventoryDevices, helper.InventoryDevicesConfig{
				ResourceName:   fmt.Sprintf("inventory_%d", i+1),
				DependsOn:      fmt.Sprintf("resource.ndfc_fabric_vxlan_evpn.vxlan_fabric_%d", i+1),
				FabricName:     childFabric.Name,
				AuthProtocol:   "md5",
				Username:       cfg.NDFC.User,
				Password:       cfg.NDFC.Password,
				MaxHops:        0,
				PreserveConfig: false,
				Save:           true,
				Deploy:         true,
				Retries:        300,
				RetryWaitTime:  20,
				Devices:        devices,
			})
		}
	}

	return testConfig
}

// initializeVrfConfig initializes VRF configuration within the unified test config structure
func initializeVrfConfig(testConfig *helper.MsdComprehensiveTestConfig, cfg *helper.Config) {
	if testConfig.MsdFabric == nil {
		return
	}

	// Build a map of active child fabrics for quick lookup
	activeFabricMap := make(map[string]bool)
	for _, fabricName := range testConfig.MsdFabric.ChildFabrics {
		activeFabricMap[fabricName] = true
	}

	// Build child VRF configs only for active child fabrics
	childVrfs := make([]helper.MsdVrfsConfig, 0)
	for i, childFabric := range cfg.NDFC.Msd.ChildFabrics {
		if !activeFabricMap[childFabric.Name] {
			continue
		}
		childVrfs = append(childVrfs, helper.MsdVrfsConfig{
			ResourceName:         fmt.Sprintf("child_vrfs%d", i+1),
			DependsOn:            "resource.ndfc_vrfs.msd_vrfs",
			FabricName:           childFabric.Name,
			DeployAllAttachments: false,
			Vrfs:                 make(map[string]helper.MsdVrfEntry),
		})
	}

	testConfig.MsdParentVrfs = &helper.MsdVrfsConfig{
		ResourceName:         "msd_vrfs",
		DependsOn:            "resource.ndfc_fabric_vxlan_msd.msd_fabric",
		FabricName:           cfg.NDFC.Msd.MsdFabricName,
		DeployAllAttachments: false,
		Vrfs:                 make(map[string]helper.MsdVrfEntry),
	}
	testConfig.MsdChildVrfs = childVrfs
}

// initializeNetworkConfig initializes network configuration within the unified test config structure
func initializeNetworkConfig(testConfig *helper.MsdComprehensiveTestConfig, cfg *helper.Config) {
	if testConfig.MsdFabric == nil {
		return
	}

	// Build a map of active child fabrics for quick lookup
	activeFabricMap := make(map[string]bool)
	for _, fabricName := range testConfig.MsdFabric.ChildFabrics {
		activeFabricMap[fabricName] = true
	}

	// Build depends_on string for active child VRFs that have switches
	childVrfDeps := "resource.ndfc_vrfs.msd_vrfs"
	for i, childFabric := range cfg.NDFC.Msd.ChildFabrics {
		if activeFabricMap[childFabric.Name] && len(childFabric.Switches) > 0 {
			childVrfDeps += fmt.Sprintf(", resource.ndfc_vrfs.child_vrfs%d", i+1)
		}
	}

	testConfig.MsdParentNetworks = &helper.MsdNetworksConfig{
		ResourceName:         "msd_networks",
		DependsOn:            childVrfDeps,
		FabricName:           cfg.NDFC.Msd.MsdFabricName,
		DeployAllAttachments: false,
		Networks:             make(map[string]helper.MsdNetworkEntry),
	}

	// Build child network configs only for active child fabrics
	childNetworks := make([]helper.MsdNetworksConfig, 0)
	for i, childFabric := range cfg.NDFC.Msd.ChildFabrics {
		if !activeFabricMap[childFabric.Name] {
			continue
		}
		childNetworks = append(childNetworks, helper.MsdNetworksConfig{
			ResourceName:         fmt.Sprintf("child_networks%d", i+1),
			DependsOn:            "resource.ndfc_networks.msd_networks" + childVrfDeps[len("resource.ndfc_vrfs.msd_vrfs"):],
			FabricName:           childFabric.Name,
			DeployAllAttachments: false,
			Networks:             make(map[string]helper.MsdNetworkEntry),
		})
	}
	testConfig.MsdChildNetworks = childNetworks
}

// createVrf creates a VRF with attachments using the unified test config structure
func createVrf(testConfig *helper.MsdComprehensiveTestConfig, cfg *helper.Config, vrfIndex int) {
	vrfName := fmt.Sprintf("%s%02d", cfg.NDFC.Msd.VrfPrefix, vrfIndex)

	// Build a map of active child fabric names based on existing child resources
	activeFabricMap := make(map[string]bool)
	for _, childVrf := range testConfig.MsdChildVrfs {
		activeFabricMap[childVrf.FabricName] = true
	}

	// Parent VRF - Build attach list only for active child fabrics
	parentVrfAttachList := make(map[string]helper.MsdAttachment)
	for _, childFabric := range cfg.NDFC.Msd.ChildFabrics {
		if !activeFabricMap[childFabric.Name] {
			continue
		}
		for _, switchSerial := range childFabric.Switches {
			parentVrfAttachList[switchSerial] = helper.MsdAttachment{
				DeployThisAttachment: false,
				Fabric:               childFabric.Name,
			}
		}
	}

	testConfig.MsdParentVrfs.Vrfs[vrfName] = helper.MsdVrfEntry{
		VrfId:              50000 + vrfIndex,
		VlanId:             3500 + vrfIndex,
		Mtu:                9216,
		LoopbackRoutingTag: 12345 + vrfIndex,
		MaxBgpPaths:        1,
		MaxIbgpPaths:       2,
		AttachList:         parentVrfAttachList,
	}

	// Create VRF in each child resource
	for i := range testConfig.MsdChildVrfs {
		childFabricName := testConfig.MsdChildVrfs[i].FabricName
		childVrfAttachList := make(map[string]helper.MsdAttachment)
		for _, childFabric := range cfg.NDFC.Msd.ChildFabrics {
			if childFabric.Name != childFabricName {
				continue
			}
			for _, switchSerial := range childFabric.Switches {
				childVrfAttachList[switchSerial] = helper.MsdAttachment{
					DeployThisAttachment: false,
				}
			}
		}
		testConfig.MsdChildVrfs[i].Vrfs[vrfName] = helper.MsdVrfEntry{
			VrfId:              50000 + vrfIndex,
			VlanId:             3500 + vrfIndex,
			Mtu:                9216,
			LoopbackRoutingTag: 12345 + vrfIndex,
			MaxBgpPaths:        1,
			MaxIbgpPaths:       2,
			AttachList:         childVrfAttachList,
		}
	}
}

// deleteVrfs removes VRFs from the unified test config structure
func deleteVrfs(testConfig *helper.MsdComprehensiveTestConfig, cfg *helper.Config, vrfIndices []int) {
	for _, i := range vrfIndices {
		vrfName := fmt.Sprintf("%s%02d", cfg.NDFC.Msd.VrfPrefix, i)
		delete(testConfig.MsdParentVrfs.Vrfs, vrfName)
		for j := range testConfig.MsdChildVrfs {
			delete(testConfig.MsdChildVrfs[j].Vrfs, vrfName)
		}
	}
}

// createNetwork creates a network with attachments using the unified test config structure
func createNetwork(testConfig *helper.MsdComprehensiveTestConfig, cfg *helper.Config, networkIndex int, vrfIndex int) {
	nwName := fmt.Sprintf("%s%d", cfg.NDFC.Msd.NetworkPrefix, networkIndex)
	vrfName := fmt.Sprintf("%s%02d", cfg.NDFC.Msd.VrfPrefix, vrfIndex)

	// Build a map of active child fabric names based on existing child resources
	activeFabricMap := make(map[string]bool)
	for _, childNetwork := range testConfig.MsdChildNetworks {
		activeFabricMap[childNetwork.FabricName] = true
	}

	// Parent Network - Build attach list only for active child fabrics
	parentNwAttachList := make(map[string]helper.MsdAttachment)
	for _, childFabric := range cfg.NDFC.Msd.ChildFabrics {
		if !activeFabricMap[childFabric.Name] {
			continue
		}
		for _, switchSerial := range childFabric.Switches {
			parentNwAttachList[switchSerial] = helper.MsdAttachment{
				DeployThisAttachment: false,
				Fabric:               childFabric.Name,
			}
		}
	}

	testConfig.MsdParentNetworks.Networks[nwName] = helper.MsdNetworkEntry{
		VrfName:              vrfName,
		VlanId:               300 + networkIndex,
		NetworkId:            30000 + networkIndex,
		Mtu:                  9100,
		RoutingTag:           54321 + networkIndex,
		InterfaceDescription: nwName + " SVI",
		Attachments:          parentNwAttachList,
	}

	// Create Network in each child resource
	for i := range testConfig.MsdChildNetworks {
		childFabricName := testConfig.MsdChildNetworks[i].FabricName
		childNwAttachList := make(map[string]helper.MsdAttachment)
		for _, childFabric := range cfg.NDFC.Msd.ChildFabrics {
			if childFabric.Name != childFabricName {
				continue
			}
			for _, switchSerial := range childFabric.Switches {
				childNwAttachList[switchSerial] = helper.MsdAttachment{
					DeployThisAttachment: false,
				}
			}
		}
		testConfig.MsdChildNetworks[i].Networks[nwName] = helper.MsdNetworkEntry{
			VrfName:              vrfName,
			VlanId:               300 + networkIndex,
			NetworkId:            30000 + networkIndex,
			Mtu:                  9100,
			RoutingTag:           54321 + networkIndex,
			InterfaceDescription: nwName + " SVI",
			Trm:                  true,
			MulticastGroup:       fmt.Sprintf("239.1.1.%d", networkIndex),
			IgmpVersion:          3,
			Attachments:          childNwAttachList,
		}
	}
}

// deleteNetworks removes networks from the unified test config structure
func deleteNetworks(testConfig *helper.MsdComprehensiveTestConfig, cfg *helper.Config, networkIndices []int) {
	for _, i := range networkIndices {
		nwName := fmt.Sprintf("%s%d", cfg.NDFC.Msd.NetworkPrefix, i)
		delete(testConfig.MsdParentNetworks.Networks, nwName)
		for j := range testConfig.MsdChildNetworks {
			delete(testConfig.MsdChildNetworks[j].Networks, nwName)
		}
	}
}

// modifyVrfParameters modifies VRF parameters using the unified test config structure
func modifyVrfParameters(testConfig *helper.MsdComprehensiveTestConfig, cfg *helper.Config, vrfIndices []int, mtu int, routingTagBase int, maxBgpPaths int, maxIbgpPaths int) {
	for _, i := range vrfIndices {
		vrfName := fmt.Sprintf("%s%02d", cfg.NDFC.Msd.VrfPrefix, i)

		// Modify parent VRF
		vrf := testConfig.MsdParentVrfs.Vrfs[vrfName]
		vrf.Mtu = mtu
		vrf.LoopbackRoutingTag = routingTagBase + i
		vrf.MaxBgpPaths = maxBgpPaths
		vrf.MaxIbgpPaths = maxIbgpPaths
		testConfig.MsdParentVrfs.Vrfs[vrfName] = vrf

		// Modify all child VRFs
		for j := range testConfig.MsdChildVrfs {
			childVrf := testConfig.MsdChildVrfs[j].Vrfs[vrfName]
			childVrf.Mtu = mtu
			childVrf.LoopbackRoutingTag = routingTagBase + i
			childVrf.MaxBgpPaths = maxBgpPaths
			childVrf.MaxIbgpPaths = maxIbgpPaths
			testConfig.MsdChildVrfs[j].Vrfs[vrfName] = childVrf
		}
	}
}

// modifyNetworkParameters modifies network parameters using the unified test config structure
func modifyNetworkParameters(testConfig *helper.MsdComprehensiveTestConfig, cfg *helper.Config, networkIndices []int, vlanBase int, mtu int, routingTagBase int, descSuffix string, trm bool, multicastBase string, igmpVersion int) {
	for _, i := range networkIndices {
		nwName := fmt.Sprintf("%s%d", cfg.NDFC.Msd.NetworkPrefix, i)

		// Modify parent network
		nw := testConfig.MsdParentNetworks.Networks[nwName]
		nw.VlanId = vlanBase + i
		nw.Mtu = mtu
		nw.RoutingTag = routingTagBase + i
		nw.InterfaceDescription = nwName + descSuffix
		testConfig.MsdParentNetworks.Networks[nwName] = nw

		// Modify all child networks
		for j := range testConfig.MsdChildNetworks {
			childNw := testConfig.MsdChildNetworks[j].Networks[nwName]
			childNw.VlanId = vlanBase + i
			childNw.Mtu = mtu
			childNw.RoutingTag = routingTagBase + i
			childNw.InterfaceDescription = nwName + descSuffix
			childNw.Trm = trm
			childNw.MulticastGroup = fmt.Sprintf(multicastBase+"%d", i)
			childNw.IgmpVersion = igmpVersion
			testConfig.MsdChildNetworks[j].Networks[nwName] = childNw
		}
	}
}

// setGlobalDeployFlags sets global deployment flags using the unified test config structure
func setGlobalDeployFlags(testConfig *helper.MsdComprehensiveTestConfig, deployVrfs bool, deployNetworks bool) {
	if testConfig.MsdParentVrfs != nil {
		testConfig.MsdParentVrfs.DeployAllAttachments = deployVrfs
	}
	if testConfig.MsdParentNetworks != nil {
		testConfig.MsdParentNetworks.DeployAllAttachments = deployNetworks
	}
}

// createVrfWithInvalidFabricField creates a VRF with fabric field set in attachment,
// which is INVALID for non-MSD-parent fabrics. This is used for error testing.
// Error expected: "Invalid fabric field in attachment"
func createVrfWithInvalidFabricField(testConfig *helper.MsdComprehensiveTestConfig, cfg *helper.Config, fabricIndex int) {
	if fabricIndex >= len(cfg.NDFC.Msd.ChildFabrics) {
		return
	}

	childFabric := cfg.NDFC.Msd.ChildFabrics[fabricIndex]

	// Create VRF attachments with fabric field set - this is INVALID for individual fabric
	vrfAttachList := make(map[string]helper.MsdAttachment)
	for _, switchSerial := range childFabric.Switches {
		vrfAttachList[switchSerial] = helper.MsdAttachment{
			DeployThisAttachment: false,
			Fabric:               childFabric.Name, // This is INVALID - fabric field not allowed for non-MSD
		}
	}

	vrfName := fmt.Sprintf("%s%02d", cfg.NDFC.Msd.VrfPrefix, 1)

	// Find the inventory resource name for depends_on
	dependsOn := ""
	for _, inv := range testConfig.InventoryDevices {
		if inv.FabricName == childFabric.Name {
			dependsOn = fmt.Sprintf("resource.ndfc_inventory_devices.%s", inv.ResourceName)
			break
		}
	}

	// Use ChildVrfs - same structure works for individual fabrics and MSD children
	testConfig.MsdChildVrfs = append(testConfig.MsdChildVrfs, helper.MsdVrfsConfig{
		ResourceName:         fmt.Sprintf("child_vrfs%d", fabricIndex+1),
		DependsOn:            dependsOn,
		FabricName:           childFabric.Name,
		DeployAllAttachments: false,
		Vrfs: map[string]helper.MsdVrfEntry{
			vrfName: {
				VrfId:              50001,
				VlanId:             3501,
				Mtu:                9216,
				LoopbackRoutingTag: 12346,
				MaxBgpPaths:        1,
				MaxIbgpPaths:       2,
				AttachList:         vrfAttachList,
			},
		},
	})
}

// createVrfOnIndividualFabric creates valid VRFs on individual fabrics (before MSD is created).
// Uses ChildVrfs with depends_on pointing to inventory. When MSD is created later,
// the depends_on will be updated to point to msd_vrfs.
func createVrfOnIndividualFabric(testConfig *helper.MsdComprehensiveTestConfig, cfg *helper.Config, fabricIndex int, vrfIndices []int) {
	if fabricIndex >= len(cfg.NDFC.Msd.ChildFabrics) {
		return
	}

	childFabric := cfg.NDFC.Msd.ChildFabrics[fabricIndex]

	// Find the inventory resource name for depends_on
	dependsOn := ""
	for _, inv := range testConfig.InventoryDevices {
		if inv.FabricName == childFabric.Name {
			dependsOn = fmt.Sprintf("resource.ndfc_inventory_devices.%s", inv.ResourceName)
			break
		}
	}

	vrfs := make(map[string]helper.MsdVrfEntry)
	for _, vrfIndex := range vrfIndices {
		vrfName := fmt.Sprintf("%s%02d", cfg.NDFC.Msd.VrfPrefix, vrfIndex)

		// Create VRF attachments WITHOUT fabric field - valid for individual fabric
		vrfAttachList := make(map[string]helper.MsdAttachment)
		for _, switchSerial := range childFabric.Switches {
			vrfAttachList[switchSerial] = helper.MsdAttachment{
				DeployThisAttachment: false,
			}
		}

		vrfs[vrfName] = helper.MsdVrfEntry{
			VrfId:              50000 + vrfIndex,
			VlanId:             3500 + vrfIndex,
			Mtu:                9216,
			LoopbackRoutingTag: 12345 + vrfIndex,
			MaxBgpPaths:        1,
			MaxIbgpPaths:       2,
			AttachList:         vrfAttachList,
		}
	}

	// Use ChildVrfs - same structure works for individual fabrics and MSD children
	testConfig.MsdChildVrfs = append(testConfig.MsdChildVrfs, helper.MsdVrfsConfig{
		ResourceName:         fmt.Sprintf("child_vrfs%d", fabricIndex+1),
		DependsOn:            dependsOn,
		FabricName:           childFabric.Name,
		DeployAllAttachments: false,
		Vrfs:                 vrfs,
	})
}

// setChildPropertiesOnParentVrf sets child-only properties on a parent VRF entry
// This should trigger validation error as these properties are not allowed on MSD parent
func setChildPropertiesOnParentVrf(testConfig *helper.MsdComprehensiveTestConfig, vrfIndex int, cfg *helper.Config) {
	vrfName := fmt.Sprintf("%s%02d", cfg.NDFC.Msd.VrfPrefix, vrfIndex)
	if vrfEntry, exists := testConfig.MsdParentVrfs.Vrfs[vrfName]; exists {
		vrfEntry.AdvertiseHostRoutes = true
		testConfig.MsdParentVrfs.Vrfs[vrfName] = vrfEntry
	}
}

// clearChildPropertiesOnParentVrf clears child-only properties from a parent VRF entry
func clearChildPropertiesOnParentVrf(testConfig *helper.MsdComprehensiveTestConfig, vrfIndex int, cfg *helper.Config) {
	vrfName := fmt.Sprintf("%s%02d", cfg.NDFC.Msd.VrfPrefix, vrfIndex)
	if vrfEntry, exists := testConfig.MsdParentVrfs.Vrfs[vrfName]; exists {
		vrfEntry.AdvertiseHostRoutes = false
		vrfEntry.AdvertiseDefaultRoute = false
		testConfig.MsdParentVrfs.Vrfs[vrfName] = vrfEntry
	}
}

// setChildPropertiesOnParentNetwork sets child-only properties on a parent Network entry
// This should trigger validation error as these properties are not allowed on MSD parent
func setChildPropertiesOnParentNetwork(testConfig *helper.MsdComprehensiveTestConfig, networkIndex int, cfg *helper.Config) {
	nwName := fmt.Sprintf("%s%d", cfg.NDFC.Msd.NetworkPrefix, networkIndex)
	if nwEntry, exists := testConfig.MsdParentNetworks.Networks[nwName]; exists {
		nwEntry.Trm = true
		testConfig.MsdParentNetworks.Networks[nwName] = nwEntry
	}
}

// clearChildPropertiesOnParentNetwork clears child-only properties from a parent Network entry
func clearChildPropertiesOnParentNetwork(testConfig *helper.MsdComprehensiveTestConfig, networkIndex int, cfg *helper.Config) {
	nwName := fmt.Sprintf("%s%d", cfg.NDFC.Msd.NetworkPrefix, networkIndex)
	if nwEntry, exists := testConfig.MsdParentNetworks.Networks[nwName]; exists {
		nwEntry.Trm = false
		nwEntry.MulticastGroup = ""
		nwEntry.Netflow = false
		nwEntry.L3GatewayOnBorder = false
		testConfig.MsdParentNetworks.Networks[nwName] = nwEntry
	}
}

// setMismatchedParentPropertyOnChildVrf sets a parent-managed property on a child VRF that
// doesn't match the parent fabric config. This should trigger validation error:
// "Mismatch of config between parent and child attributes"
func setMismatchedParentPropertyOnChildVrf(testConfig *helper.MsdComprehensiveTestConfig, vrfIndex int, cfg *helper.Config) {
	vrfName := fmt.Sprintf("%s%02d", cfg.NDFC.Msd.VrfPrefix, vrfIndex)

	// Set mismatched MTU on all child VRFs - different from what parent has
	for i := range testConfig.MsdChildVrfs {
		if vrfEntry, exists := testConfig.MsdChildVrfs[i].Vrfs[vrfName]; exists {
			// Parent has MTU 9216, set child to something different
			vrfEntry.Mtu = 1500
			testConfig.MsdChildVrfs[i].Vrfs[vrfName] = vrfEntry
		}
	}
}

// clearMismatchedParentPropertyOnChildVrf restores the parent-managed property on child VRF
// to match the parent fabric config
func clearMismatchedParentPropertyOnChildVrf(testConfig *helper.MsdComprehensiveTestConfig, vrfIndex int, cfg *helper.Config) {
	vrfName := fmt.Sprintf("%s%02d", cfg.NDFC.Msd.VrfPrefix, vrfIndex)

	// Restore MTU to match parent
	for i := range testConfig.MsdChildVrfs {
		if vrfEntry, exists := testConfig.MsdChildVrfs[i].Vrfs[vrfName]; exists {
			vrfEntry.Mtu = 9216 // Same as parent
			testConfig.MsdChildVrfs[i].Vrfs[vrfName] = vrfEntry
		}
	}
}

// setMismatchedParentPropertyOnChildNetwork sets a parent-managed property on a child Network that
// doesn't match the parent fabric config. This should trigger validation error:
// "Mismatch of config between parent and child attributes"
func setMismatchedParentPropertyOnChildNetwork(testConfig *helper.MsdComprehensiveTestConfig, networkIndex int, cfg *helper.Config) {
	nwName := fmt.Sprintf("%s%d", cfg.NDFC.Msd.NetworkPrefix, networkIndex)

	// Set mismatched MTU on all child Networks - different from what parent has
	for i := range testConfig.MsdChildNetworks {
		if nwEntry, exists := testConfig.MsdChildNetworks[i].Networks[nwName]; exists {
			// Parent has MTU 9100, set child to something different
			nwEntry.Mtu = 1500
			testConfig.MsdChildNetworks[i].Networks[nwName] = nwEntry
		}
	}
}

// clearMismatchedParentPropertyOnChildNetwork restores the parent-managed property on child Network
// to match the parent fabric config
func clearMismatchedParentPropertyOnChildNetwork(testConfig *helper.MsdComprehensiveTestConfig, networkIndex int, cfg *helper.Config) {
	nwName := fmt.Sprintf("%s%d", cfg.NDFC.Msd.NetworkPrefix, networkIndex)

	// Restore MTU to match parent
	for i := range testConfig.MsdChildNetworks {
		if nwEntry, exists := testConfig.MsdChildNetworks[i].Networks[nwName]; exists {
			nwEntry.Mtu = 9100 // Same as parent
			testConfig.MsdChildNetworks[i].Networks[nwName] = nwEntry
		}
	}
}
