// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package testing

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"terraform-provider-ndfc/internal/provider/resources/resource_fabric_vxlan_msd"
	"terraform-provider-ndfc/internal/provider/resources/resource_interface_common"
	"terraform-provider-ndfc/internal/provider/resources/resource_links"
	"terraform-provider-ndfc/internal/provider/resources/resource_networks"
	"terraform-provider-ndfc/internal/provider/resources/resource_policy"
	"terraform-provider-ndfc/internal/provider/resources/resource_vpc_pair"
	"terraform-provider-ndfc/internal/provider/resources/resource_vrf_bulk"
	"terraform-provider-ndfc/internal/provider/types"
	"text/template"
	"time"
)

// MsdComprehensiveTestConfig is a unified structure for MSD comprehensive test
// It encapsulates all resources: VXLAN fabrics, inventory, MSD fabric, VRFs, and Networks
type MsdComprehensiveTestConfig struct {
	// VXLAN Fabrics (child fabrics for MSD)
	VxlanFabrics []VxlanFabricConfig
	// Inventory devices for each child fabric
	InventoryDevices []InventoryDevicesConfig
	// MSD Fabric configuration
	MsdFabric *MsdFabricConfig
	// MSD Parent VRFs in MSD fabric
	MsdParentVrfs *MsdVrfsConfig
	// MSD Child VRFs in each child fabric
	MsdChildVrfs []MsdVrfsConfig
	// MSD Parent Networks in MSD fabric
	MsdParentNetworks *MsdNetworksConfig
	// MSD Child Networks in each child fabric
	MsdChildNetworks []MsdNetworksConfig
}

// MsdFabricConfig represents the MSD fabric resource configuration
type MsdFabricConfig struct {
	ResourceName string
	DependsOn    string
	FabricName   string
	Deploy       bool
	ChildFabrics []string
}

// MSD VRFs and Networks structures (kept for backward compatibility)
type MsdVrfsNetworks struct {
	MsdParentVrfs     *MsdVrfsConfig
	MsdChildVrfs      []MsdVrfsConfig
	MsdParentNetworks *MsdNetworksConfig
	MsdChildNetworks  []MsdNetworksConfig
}

type MsdVrfsConfig struct {
	ResourceName         string
	DependsOn            string
	FabricName           string
	DeployAllAttachments bool
	Vrfs                 map[string]MsdVrfEntry
}

type MsdVrfEntry struct {
	VrfId                      int
	VlanId                     int
	DisableRtAuto              bool
	Ipv6LinkLocal              bool
	LoopbackRoutingTag         int
	MaxBgpPaths                int
	MaxIbgpPaths               int
	Mtu                        int
	RedistributeDirectRouteMap string
	VrfExtensionTemplate       string
	VrfTemplate                string
	InterfaceDescription       string
	VlanName                   string
	VrfDescription             string
	AdvertiseDefaultRoute      bool
	AdvertiseHostRoutes        bool
	AttachList                 map[string]MsdAttachment
}

type MsdNetworksConfig struct {
	ResourceName         string
	DependsOn            string
	FabricName           string
	DeployAllAttachments bool
	Networks             map[string]MsdNetworkEntry
}

type MsdNetworkEntry struct {
	VrfName              string
	NetworkId            int
	GatewayIpv4Address   string
	GatewayIpv6Address   string
	VlanId               int
	VlanName             string
	Layer2Only           bool
	InterfaceDescription string
	Mtu                  int
	SecondaryGateway1    string
	SecondaryGateway2    string
	SecondaryGateway3    string
	SecondaryGateway4    string
	RoutingTag           int
	ArpSuppression       bool
	RouteTargetBoth      bool
	DhcpLoopbackId       int
	MulticastGroup       string
	Trm                  bool
	Netflow              bool
	VlanNetflowMonitor   string
	L3GatewayOnBorder    bool
	IgmpVersion          int
	Attachments          map[string]MsdAttachment
}

type MsdAttachment struct {
	DeployThisAttachment bool
	Fabric               string
}

// VXLAN Fabric and Inventory structures for MSD test setup
type MsdFabricsAndInventory struct {
	VxlanFabrics     []VxlanFabricConfig
	InventoryDevices []InventoryDevicesConfig
}

type VxlanFabricConfig struct {
	ResourceName string
	FabricName   string
	BgpAs        string
	Deploy       bool
}

type InventoryDevicesConfig struct {
	ResourceName   string
	DependsOn      string
	FabricName     string
	AuthProtocol   string
	Username       string
	Password       string
	MaxHops        int
	PreserveConfig bool
	Save           bool
	Deploy         bool
	Retries        int
	RetryWaitTime  int
	Devices        map[string]InventoryDeviceEntry
}

type InventoryDeviceEntry struct {
	Role                  string
	DiscoveryType         string
	DiscoveryAuthProtocol string
}

const tfHeader = `
terraform {
	required_providers {
		ndfc = {
		source = "registry.terraform.io/cisco/ndfc"
		}
	}
}`

const tfProviderConfig = `
provider "ndfc" {
	username = "{{.User}}"
	password = "{{.Password}}"
	url     = "{{.Host}}"
	insecure = {{.Insecure}}
}
`

const tfDependenciesDot = `
digraph G {
  rankdir = "RL";
  node [shape = rect, fontname = "sans-serif"];
  "ndfc_configuration_deploy.test_resource_configuration_deploy_1" [label="ndfc_configuration_deploy.test_resource_configuration_deploy_1"];
  "ndfc_interface_ethernet.test_resource_interface_ethernet_1" [label="ndfc_interface_ethernet.test_resource_interface_ethernet_1"];
  "ndfc_interface_loopback.test_resource_interface_loopback_1" [label="ndfc_interface_loopback.test_resource_interface_loopback_1"];
  "ndfc_interface_portchannel.test_resource_interface_portchannel_1" [label="ndfc_interface_portchannel.test_resource_interface_portchannel_1"];
  "ndfc_interface_vlan.test_resource_interface_vlan_1" [label="ndfc_interface_vlan.test_resource_interface_vlan_1"];
  "ndfc_interface_vpc.test_resource_interface_vpc_1" [label="ndfc_interface_vpc.test_resource_interface_vpc_1"];
  "ndfc_inventory_devices.test_resource_inventory_devices_1" [label="ndfc_inventory_devices.test_resource_inventory_devices_1"];
  "ndfc_networks.test_resource_networks_1" [label="ndfc_networks.test_resource_networks_1"];
  "ndfc_policy.test_resource_policy_1" [label="ndfc_policy.test_resource_policy_1"];
  "ndfc_vpc_pair.test_resource_vpc_pair_1" [label="ndfc_vpc_pair.test_resource_vpc_pair_1"];
  "ndfc_vrfs.test_resource_vrf_bulk_1" [label="ndfc_vrfs.test_resource_vrf_bulk_1"];
  "ndfc_vxlan_evpn_fabric.test_resource_fabric_vxlan_evpn_1" [label="ndfc_vxlan_evpn_fabric.test_resource_fabric_vxlan_evpn_1"];
  "ndfc_interface_ethernet.test_resource_interface_ethernet_1" -> "ndfc_vpc_pair.test_resource_vpc_pair_1";
  "ndfc_interface_loopback.test_resource_interface_loopback_1" -> "ndfc_vpc_pair.test_resource_vpc_pair_1";
  "ndfc_interface_portchannel.test_resource_interface_portchannel_1" -> "ndfc_interface_ethernet.test_resource_interface_ethernet_1";
  "ndfc_interface_vlan.test_resource_interface_vlan_1" -> "ndfc_vpc_pair.test_resource_vpc_pair_1";
  "ndfc_interface_vpc.test_resource_interface_vpc_1" -> "ndfc_inventory_devices.test_resource_inventory_devices_1";
   "ndfc_interface_vpc.test_resource_interface_vpc_1" -> "ndfc_vpc_pair.test_resource_vpc_pair_1"
  "ndfc_inventory_devices.test_resource_inventory_devices_1" -> "ndfc_fabric_vxlan_evpn.test_resource_fabric_vxlan_evpn_1";
  "ndfc_networks.test_resource_networks_1" -> "ndfc_vrfs.test_resource_vrfs_1";
  "ndfc_policy.test_resource_policy_1" -> "ndfc_networks.test_resource_networks_1";
  "ndfc_vpc_pair.test_resource_vpc_pair_1" -> "ndfc_inventory_devices.test_resource_inventory_devices_1";
  "ndfc_vrfs.test_resource_vrfs_1" -> "ndfc_vpc_pair.test_resource_vpc_pair_1";
}`

const tgDependenciesGlobalDeploy = `
digraph G {
  rankdir = "RL";
  node [shape = rect, fontname = "sans-serif"];
  "ndfc_configuration_deploy.test_resource_configuration_deploy_1" [label="ndfc_configuration_deploy.test_resource_configuration_deploy_1"];
  "ndfc_fabric_vxlan_evpn.test_resource_fabric_vxlan_evpn_1" [label="ndfc_fabric_vxlan_evpn.test_resource_fabric_vxlan_evpn_1"];
  "ndfc_interface_ethernet.test_resource_interface_ethernet_1" [label="ndfc_interface_ethernet.test_resource_interface_ethernet_1"];
  "ndfc_interface_loopback.test_resource_interface_loopback_1" [label="ndfc_interface_loopback.test_resource_interface_loopback_1"];
  "ndfc_interface_portchannel.test_resource_interface_portchannel_1" [label="ndfc_interface_portchannel.test_resource_interface_portchannel_1"];
  "ndfc_interface_vlan.test_resource_interface_vlan_1" [label="ndfc_interface_vlan.test_resource_interface_vlan_1"];
  "ndfc_interface_vpc.test_resource_interface_vpc_1" [label="ndfc_interface_vpc.test_resource_interface_vpc_1"];
  "ndfc_inventory_devices.test_resource_inventory_devices_1" [label="ndfc_inventory_devices.test_resource_inventory_devices_1"];
  "ndfc_networks.test_resource_networks_1" [label="ndfc_networks.test_resource_networks_1"];
  "ndfc_policy.test_resource_policy_1" [label="ndfc_policy.test_resource_policy_1"];
  "ndfc_vpc_pair.test_resource_vpc_pair_1" [label="ndfc_vpc_pair.test_resource_vpc_pair_1"];
  "ndfc_vrfs.test_resource_vrfs_1" [label="ndfc_vrfs.test_resource_vrfs_1"];
  "ndfc_configuration_deploy.test_resource_configuration_deploy_1" -> "ndfc_interface_loopback.test_resource_interface_loopback_1";
  "ndfc_configuration_deploy.test_resource_configuration_deploy_1" -> "ndfc_interface_portchannel.test_resource_interface_portchannel_1";
  "ndfc_configuration_deploy.test_resource_configuration_deploy_1" -> "ndfc_interface_vlan.test_resource_interface_vlan_1";
  "ndfc_configuration_deploy.test_resource_configuration_deploy_1" -> "ndfc_interface_vpc.test_resource_interface_vpc_1";
  "ndfc_configuration_deploy.test_resource_configuration_deploy_1" -> "ndfc_networks.test_resource_networks_1";
  "ndfc_configuration_deploy.test_resource_configuration_deploy_1" -> "ndfc_policy.test_resource_policy_1";
  "ndfc_interface_ethernet.test_resource_interface_ethernet_1" -> "ndfc_vpc_pair.test_resource_vpc_pair_1";
  "ndfc_interface_loopback.test_resource_interface_loopback_1" -> "ndfc_vpc_pair.test_resource_vpc_pair_1";
  "ndfc_interface_portchannel.test_resource_interface_portchannel_1" -> "ndfc_interface_ethernet.test_resource_interface_ethernet_1";
  "ndfc_interface_vlan.test_resource_interface_vlan_1" -> "ndfc_vpc_pair.test_resource_vpc_pair_1";
  "ndfc_interface_vpc.test_resource_interface_vpc_1" -> "ndfc_vpc_pair.test_resource_vpc_pair_1";
  "ndfc_inventory_devices.test_resource_inventory_devices_1" -> "ndfc_fabric_vxlan_evpn.test_resource_fabric_vxlan_evpn_1";
  "ndfc_networks.test_resource_networks_1" -> "ndfc_vrfs.test_resource_vrfs_1";
  "ndfc_policy.test_resource_policy_1" -> "ndfc_inventory_devices.test_resource_inventory_devices_1";
  "ndfc_vpc_pair.test_resource_vpc_pair_1" -> "ndfc_inventory_devices.test_resource_inventory_devices_1";
  "ndfc_vrfs.test_resource_vrfs_1" -> "ndfc_vpc_pair.test_resource_vpc_pair_1";
}`

func GetTFConfigWithSingleResource(tt string, cfg map[string]string, rscs []interface{}, out **string) {
	x := new(string)
	args := map[string]interface{}{
		"User":     cfg["User"],
		"Password": cfg["Password"],
		"Host":     cfg["Host"],
		"Insecure": cfg["Insecure"],
		"RscType":  cfg["RscType"],
		"RscName":  cfg["RscName"],
	}
	functions := template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"deref": func(a *types.Int64Custom) int64 {
			if a == nil {
				return 0
			}
			return int64(*a)
		},
		"deref_bool": func(a *bool) bool {
			if a == nil {
				return false
			}
			return *a
		},
		"trimSuffix": func(s, suffix string) string {
			return strings.TrimSuffix(s, suffix)
		},
	}

	root_path, _ := os.Getwd()
	tmpl := bytes.Buffer{}
	files, err := os.ReadDir(root_path + "/testing/")
	if err != nil {
		log.Panicf("Err reading dir %v", err)
	}
	for _, file := range files {
		if strings.Contains(file.Name(), ".gotmpl") {
			tmplFile, err := os.ReadFile(root_path + "/testing/" + file.Name())
			if err != nil {
				log.Panicf("Err reading file %v", err)
			}
			tmpl.Write(tmplFile)
		}
	}

	t, err := template.New("config").Funcs(functions).Parse(tmpl.String())
	if err != nil {
		panic(err)
	}
	output := bytes.Buffer{}
	err = t.ExecuteTemplate(&output, "HEADER", args)
	if err != nil {
		panic(err)
	}
	err = t.ExecuteTemplate(&output, "NDFC", args)
	if err != nil {
		panic(err)
	}

	if len(rscs) == 0 {
		panic("Empty arr")
	}
	vrfRscName := ""
	rsNames := strings.Split(cfg["RscName"], ",")
	for i, rsc := range rscs {

		vrfBulk, ok := rsc.(*resource_vrf_bulk.NDFCVrfBulkModel)
		if ok {
			args["Vrf"] = vrfBulk
			args["RscName"] = rsNames[i]
			args["RscType"] = "vrfs"
			vrfRscName = rsNames[i]
			err = t.ExecuteTemplate(&output, "NDFC_VRF_RESOURCE", args)
			if err != nil {
				panic(err)
			}
		}
		nwRsc, ok := rsc.(*resource_networks.NDFCNetworksModel)
		if ok {
			args["Network"] = nwRsc
			args["RscName"] = rsNames[i]
			args["RscType"] = "networks"
			args["VrfRscName"] = vrfRscName
			err = t.ExecuteTemplate(&output, "NDFC_NETWORK_RESOURCE", args)
			if err != nil {
				panic(err)
			}
		}

		ifRsc, ok := rsc.(*resource_interface_common.NDFCInterfaceCommonModel)
		if ok {
			args["Interface"] = ifRsc
			args["RscName"] = rsNames[i]
			vpc, ok := cfg["RscSubType"]
			if ok && vpc == "vpc" {
				args["depends"] = "resource.ndfc_vpc_pair." + rsNames[i-1]
				args["RscSubType"] = "vpc"
			}
			args["RscType"] = "interface_" + cfg["RscSubType"]
			err = t.ExecuteTemplate(&output, "NDFC_INT_RSC", args)
			if err != nil {
				panic(err)
			}
		}

		vpcRsc, ok := rsc.(*resource_vpc_pair.NDFCVpcPairModel)
		if ok {
			args["VpcPair"] = vpcRsc
			args["RscName"] = rsNames[i]
			args["RscType"] = "vpc_pair"
			err = t.ExecuteTemplate(&output, "NDFC_VPCPAIR_RSC", args)
			if err != nil {
				panic(err)
			}
		}

		policyRsc, ok := rsc.(*resource_policy.NDFCPolicyModel)
		if ok {
			args["Policy"] = policyRsc
			args["RscName"] = rsNames[i]
			args["RscType"] = "policy"
			err = t.ExecuteTemplate(&output, "NDFC_POLICY_RSC", args)
			if err != nil {
				panic(err)
			}
		}

		linksRsc, ok := rsc.(*resource_links.NDFCLinksModel)
		if ok {
			args["Links"] = linksRsc
			args["RscName"] = rsNames[i]
			args["RscType"] = "links"
			err = t.ExecuteTemplate(&output, "NDFC_LINKS_RSC", args)
			if err != nil {
				panic(err)
			}
		}

		fabricMsdRsc, ok := rsc.(*resource_fabric_vxlan_msd.FabricVxlanMsdModel)
		if ok {
			args["FabricVxlanMsd"] = fabricMsdRsc
			args["RscName"] = rsNames[i]
			args["RscType"] = "fabric_vxlan_msd"
			err = t.ExecuteTemplate(&output, "NDFC_FABRIC_VXLAN_MSD_RSC", args)
			if err != nil {
				panic(err)
			}
		}

		msdVrfsNetworks, ok := rsc.(*MsdVrfsNetworks)
		if ok {
			args["MsdParentVrfs"] = msdVrfsNetworks.MsdParentVrfs
			args["MsdChildVrfs"] = msdVrfsNetworks.MsdChildVrfs
			args["MsdParentNetworks"] = msdVrfsNetworks.MsdParentNetworks
			args["MsdChildNetworks"] = msdVrfsNetworks.MsdChildNetworks
			err = t.ExecuteTemplate(&output, "NDFC_MSD_VRFS_NETWORKS", args)
			if err != nil {
				panic(err)
			}
		}

		fabricsAndInventory, ok := rsc.(*MsdFabricsAndInventory)
		if ok {
			args["VxlanFabrics"] = fabricsAndInventory.VxlanFabrics
			args["InventoryDevices"] = fabricsAndInventory.InventoryDevices
			err = t.ExecuteTemplate(&output, "NDFC_VXLAN_FABRICS_AND_INVENTORY", args)
			if err != nil {
				panic(err)
			}
		}

		// Unified MSD Comprehensive Test Config - handles all MSD test resources in one structure
		msdTestConfig, ok := rsc.(*MsdComprehensiveTestConfig)
		if ok {
			args["VxlanFabrics"] = msdTestConfig.VxlanFabrics
			args["InventoryDevices"] = msdTestConfig.InventoryDevices
			args["MsdFabric"] = msdTestConfig.MsdFabric
			args["MsdParentVrfs"] = msdTestConfig.MsdParentVrfs
			args["MsdChildVrfs"] = msdTestConfig.MsdChildVrfs
			args["MsdParentNetworks"] = msdTestConfig.MsdParentNetworks
			args["MsdChildNetworks"] = msdTestConfig.MsdChildNetworks
			err = t.ExecuteTemplate(&output, "NDFC_MSD_COMPREHENSIVE_TEST", args)
			if err != nil {
				panic(err)
			}
		}
	}
	//log.Println(output.String())
	*x = output.String()
	WriteConfigToFile(tt, &output)
	*out = x
}

func GetVRFTFConfigWithMultipleResource(tt string, cfg map[string]string, vrfBulk *[]*resource_vrf_bulk.NDFCVrfBulkModel, out **string) {
	x := new(string)
	args := map[string]interface{}{
		"User":     cfg["User"],
		"Password": cfg["Password"],
		"Host":     cfg["Host"],
		"Insecure": cfg["Insecure"],
		"Vrf":      &vrfBulk,
		"RscType":  cfg["RscType"],
		"RscName":  cfg["RscName"],
	}

	functions := template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"deref": func(a *int64) int64 {
			if a == nil {
				return 0
			}
			return *a
		},
		"deref_bool": func(a *bool) bool {
			if a == nil {
				return false
			}
			return *a
		},
	}
	tmpl := bytes.Buffer{}
	root_path, _ := os.Getwd()
	files, err := os.ReadDir(root_path + "/testing/")
	if err != nil {
		log.Panicf("Err reading dir %v", err)
	}
	for _, file := range files {
		if strings.Contains(file.Name(), ".gotmpl") {
			tmplFile, err := os.ReadFile(root_path + "/testing/" + file.Name())
			if err != nil {
				log.Panicf("Err reading file %v", err)
			}
			tmpl.Write(tmplFile)
		}
	}
	t, err := template.New("config").Funcs(functions).Parse(tmpl.String())
	if err != nil {
		panic(err)
	}
	output := bytes.Buffer{}

	err = t.ExecuteTemplate(&output, "HEADER", args)
	if err != nil {
		panic(err)
	}
	err = t.ExecuteTemplate(&output, "NDFC", args)
	if err != nil {
		panic(err)
	}
	rscNames := strings.Split(cfg["RscName"], ",")

	for i := range *vrfBulk {
		args["Vrf"] = &(*vrfBulk)[i]
		args["RscName"] = rscNames[i]
		err = t.ExecuteTemplate(&output, "NDFC_VRF_RESOURCE", args)
		if err != nil {
			panic(err)
		}
	}
	*x = output.String()
	WriteConfigToFile(tt, &output)
	*out = x
}

func GetProviderHeader() string {

	return tfHeader

}

func GetProviderConfig(attr map[string]interface{}) string {
	tmpl, err := template.New("test").Parse(tfProviderConfig)
	if err != nil {
		panic(err)
	}
	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, attr)
	if err != nil {
		panic(err)
	}
	return tpl.String()
}

func GetTFIntegrated(ts string, rsList []string, attrs map[string]interface{}, rsDeploy bool) string {

	switches := attrs["switches"].([]string)
	vpcPair := attrs["vpc_pair"].([]string)
	fabricName := attrs["fabric"].(string)

	inventorySwitches := attrs["inventory_devices"].([]string)
	inventoryRoles := attrs["inventory_roles"].([]string)

	user := attrs["user"].(string)
	password := attrs["password"].(string)

	// Read resource.tf from each item in rsList
	//rscTfList := make([]TerraformConfig, len(rsList))
	var deps map[string][]string
	var err error
	tfConfig := new(bytes.Buffer)
	// rsDeploy=> True indicates resource level deploy needs to be set; otherwise global deploy is enabled
	if rsDeploy {
		deps, err = parseGraphWiz(tfDependenciesDot)
	} else {
		deps, err = parseGraphWiz(tgDependenciesGlobalDeploy)
	}
	if err != nil {
		panic(err)
	}
	exampleFolder := os.Getenv("GOPATH") + "/src/terraform-provider-ndfc/examples/resources"
	for _, rs := range rsList {
		tt := TerraformConfig{}
		rsTFContent, err := os.ReadFile(fmt.Sprintf("%s/%s/resource.tf", exampleFolder, rs))
		if err != nil {
			panic(err)
		}
		tt.AddContent(rsTFContent)
		tt.updateDependency(deps)
		switch rs {
		case "ndfc_fabric_vxlan_evpn":
			tt.ModifyAttributeValue("fabric_name", fabricName)

		case "ndfc_vrfs":
			tt.ModifyAttributeValue("fabric_name", fabricName)
			if !rsDeploy {
				tt.ModifyAttributeValue("deploy_all_attachments", false)
				tt.ModifyMapValue("attach_list", "deploy_this_attachment", false)
			}
			tt.ModifyMapKey("attach_list", "SWITCH_SERIAL_NO", switches[0])

		case "ndfc_networks":
			tt.ModifyAttributeValue("fabric_name", fabricName)

			if !rsDeploy {
				tt.ModifyAttributeValue("deploy_all_attachments", false)
				tt.ModifyMapValue("attachments", "deploy_this_attachment", false)
			}
			tt.ModifyMapKey("attachments", "SWITCH_SERIAL_NO", switches[0])

		case "ndfc_policy":
			tt.ModifyAttributeValue("serial_numbers", []string{switches[0]})
			if !rsDeploy {
				tt.ModifyAttributeValue("deploy", false)
			}

		case "ndfc_interface_ethernet":
			fallthrough
		case "ndfc_interface_loopback":
			fallthrough
		case "ndfc_interface_portchannel":
			fallthrough
		case "ndfc_interface_vlan":
			tt.ModifyAttributeValue("serial_number", switches[0])
			if !rsDeploy {
				tt.ModifyAttributeValue("deploy", false)
			}

		case "ndfc_interface_vpc":
			tt.ModifyAttributeValue("serial_number", vpcPair[0]+"~"+vpcPair[1])
			if !rsDeploy {
				tt.ModifyAttributeValue("deploy", false)
			}

		case "ndfc_vpc_pair":
			tt.ModifyAttributeValue("serial_numbers", vpcPair)
			if !rsDeploy {
				tt.ModifyAttributeValue("deploy", false)
			}

		case "ndfc_inventory_devices":
			tt.ModifyAttributeValue("fabric_name", fabricName)
			tt.ModifyAttributeValue("username", user)
			tt.ModifyAttributeValue("password", password)
			for i, sw := range inventorySwitches {
				if i == 0 {
					tt.AddEntryToMap("devices", sw, map[string]string{"role": inventoryRoles[i]}, true)
				} else {
					tt.AddEntryToMap("devices", sw, map[string]string{"role": inventoryRoles[i]}, false)
				}
			}
			tt.ModifyAttributeValue("deploy", true)
		case "ndfc_configuration_deploy":
			tt.ModifyAttributeValue("fabric_name", fabricName)
			tt.ModifyAttributeValue("serial_numbers", []string{"ALL"})
		}
		_, err = tt.File.WriteTo(tfConfig)
		if err != nil {
			panic(err)
		}
		//tfConfig.Write(tt.File.Bytes())
		tfConfig.Write([]byte("\n\n"))
	}
	WriteConfigToFile(ts, tfConfig)
	return tfConfig.String()
}

/*
func GetTFDataSourceIntegrated(ts string, rsList []string, attrs map[string]interface{}) string {

		switches := attrs["switches"].([]string)
		fabricName := attrs["fabric"].(string)


		//user := attrs["user"].(string)
		//password := attrs["password"].(string)

		// Read data-source.tf from each item in rsList
		//rscTfList := make([]TerraformConfig, len(rsList))

		tfConfig := new(bytes.Buffer)

		exampleFolder := os.Getenv("GOPATH") + "/src/terraform-provider-ndfc/examples/data-sources"
		for _, rs := range rsList {
			tt := TerraformConfig{}
			rsTFContent, err := os.ReadFile(fmt.Sprintf("%s/%s/data-source.tf", exampleFolder, rs))
			if err != nil {
				panic(err)
			}
			tt.AddContent(rsTFContent)
			switch rs {
			case "ndfc_fabric":
				tt.ModifyAttributeValue("fabric_name", fabricName)

			case "ndfc_vrfs":
				tt.ModifyAttributeValue("fabric_name", fabricName)

			case "ndfc_networks":
				tt.ModifyAttributeValue("fabric_name", fabricName)

			case "ndfc_interfaces":
				tt.ModifyAttributeValue("serial_number", switches[0])
			}
			_, err = tt.File.WriteTo(tfConfig)
			if err != nil {
				panic(err)
			}
			//tfConfig.Write(tt.File.Bytes())
			tfConfig.Write([]byte("\n\n"))
		}
		WriteConfigToFile(ts, tfConfig)
		return tfConfig.String()
	}
*/
func WriteConfigToFile(ts string, tfConfig *bytes.Buffer) {
	if tmpDir == "" {
		ct := time.Now()
		tmpDir = fmt.Sprintf("/tmp/tftest_%s", ct.Format("2006_01_02_15-04-05"))
		err := os.MkdirAll(tmpDir, 0755)
		if err != nil {
			panic(err)
		}
	}
	fp, err := os.Create(fmt.Sprintf("/%s/%s.tf", tmpDir, ts))
	if err != nil {
		panic(err)
	}
	_, err = fp.Write(tfConfig.Bytes())
	if err != nil {
		panic(err)
	}
	fp.Close()
}
