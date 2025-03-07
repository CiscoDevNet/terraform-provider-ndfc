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
	"terraform-provider-ndfc/internal/provider/resources/resource_interface_common"
	"terraform-provider-ndfc/internal/provider/resources/resource_networks"
	"terraform-provider-ndfc/internal/provider/resources/resource_vpc_pair"
	"terraform-provider-ndfc/internal/provider/resources/resource_vrf_bulk"
	"terraform-provider-ndfc/internal/provider/types"
	"text/template"
	"time"
)

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
	host     = "{{.Host}}"
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
			tt.ModifyAttributeValue("device_serial_number", switches[0])
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
		tt.File.WriteTo(tfConfig)
		//tfConfig.Write(tt.File.Bytes())
		tfConfig.Write([]byte("\n\n"))
	}
	WriteConfigToFile(ts, tfConfig)
	return tfConfig.String()
}

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
	fp.Write(tfConfig.Bytes())
	fp.Close()
}
