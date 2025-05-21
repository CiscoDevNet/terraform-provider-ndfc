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
	"log"
	"os"
)

const (
	Base_file = iota
	Modified_file
	Bgp_as_change_file
	Different_fabric_profiles
)

func GenerateFabricConfig(tt string, cfg map[string]string, fileType int) string {

	var tf_config = `
		terraform {
			required_providers {
				ndfc = {
					source = "registry.terraform.io/cisco/ndfc"
				}
			}
		}
		provider "ndfc" {
			url     = "` + cfg["Host"] + `"
			username = "` + cfg["User"] + `"
			password = "` + cfg["Password"] + `"
			insecure = true
		}
	`
	var filePath string
	folder := os.Getenv("GOPATH") + "/src/terraform-provider-ndfc/testing/data/" + cfg["FabricType"]
	switch fileType {
	case Base_file:
		filePath = folder + "/resource.tf"
	case Modified_file:
		filePath = folder + "/resource_modified.tf"
	case Bgp_as_change_file:
		filePath = folder + "/resource_bgp_as_change.tf"
	case Different_fabric_profiles:
		filePath = folder + "/resource_different_fabric_profiles.tf"
	default:
		log.Fatalf("Invalid file type: %d", fileType)
	}
	rscFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to open resource config file: %s", err)
	}

	tf_config += string(rscFile)
	out := new(bytes.Buffer)
	out.Write(rscFile)
	WriteConfigToFile(tt, out)
	return tf_config
}
