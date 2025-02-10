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
	log.Println(output.String())
	*x = output.String()
	if tmpDir == "" {
		ct := time.Now()
		tmpDir = fmt.Sprintf("/tmp/tftest_%s", ct.Format("2006_01_02_15-04-05"))
		err = os.MkdirAll(tmpDir, 0755)
		if err != nil {
			panic(err)
		}
	}
	fp, err := os.Create(fmt.Sprintf("/%s/%s.tf", tmpDir, tt))
	if err != nil {
		panic(err)
	}
	fp.Write(output.Bytes())
	fp.Close()
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

	log.Println(output.String())
	*x = output.String()
	if tmpDir == "" {
		ct := time.Now()
		tmpDir = fmt.Sprintf("/tmp/tftest_%s", ct.Format("2006_01_02_15-04-05"))
		err = os.MkdirAll(tmpDir, 0755)
		if err != nil {
			panic(err)
		}
	}
	fp, err := os.Create(fmt.Sprintf("/%s/%s.tf", tmpDir, tt))
	if err != nil {
		panic(err)
	}
	fp.Write(output.Bytes())
	fp.Close()
	*out = x
}
