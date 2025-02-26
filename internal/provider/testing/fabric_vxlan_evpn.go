// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package testing

import (
	"log"
	"os"
)

const (
	Base_file = iota
	Modified_file
	Bgp_as_change_file
	Enable_trm_config_file
)

func GenerateFabricConfig(cfg map[string]string, fileType int) string {

	var tf_config = `
		terraform {
			required_providers {
				ndfc = {
					source = "registry.terraform.io/cisco/ndfc"
				}
			}
		}
		provider "ndfc" {
			host     = "` + cfg["Host"] + `"
			username = "` + cfg["User"] + `"
			password = "` + cfg["Password"] + `"
			insecure = true
		}
	`
	var filePath string
	folder := os.Getenv("GOPATH") + "/src/terraform-provider-ndfc/testdata/" + cfg["FabricType"]
	switch fileType {
	case Base_file:
		filePath = folder + "/resource.tf"
	case Modified_file:
		filePath = folder + "/resource_modified.tf"
	case Bgp_as_change_file:
		filePath = folder + "/resource_bgp_as_change.tf"
	case Enable_trm_config_file:
		filePath = folder + "/resource_enable_trm.tf"
	default:
		log.Fatalf("Invalid file type: %d", fileType)
	}
	rscFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to open resource config file: %s", err)
	}

	tf_config += string(rscFile)
	return tf_config
}
