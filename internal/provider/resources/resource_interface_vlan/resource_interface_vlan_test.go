// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package resource_interface_vlan

import (
	"encoding/json"
	"os"
	"terraform-provider-ndfc/internal/provider/resources/resource_interface_common"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResourceInterfaceVlan(t *testing.T) {
	type args struct {
		rscType  string
		rscName  string
		dataFile string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test_resource_interface_vlan",
			args: args{
				rscType:  "resource",
				rscName:  "interface_vlan",
				dataFile: "/examples/ndfc_payloads/data_interface_vlan.json",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name+"model_read", func(t *testing.T) {
			fileName := os.Getenv("GOPATH") + "/src/terraform-provider-ndfc" + tt.args.dataFile
			RsType := tt.args.rscType
			rscName := tt.args.rscName
			dataFromFile, err := os.ReadFile(fileName)
			if err != nil {
				t.Errorf("File read failure %v", err)
				return
			}
			modelData := resource_interface_common.NDFCInterfaceCommonModel{}
			v := InterfaceVlanModel{}

			err = json.Unmarshal(dataFromFile, &modelData)
			if err != nil {
				t.Errorf("Json Unmarshal failed %s_%s: %v", RsType, rscName, err)
			}
			if err := v.SetModelData(&modelData); err != nil {
				t.Errorf("SetModelData failed %s_%s: %v", RsType, rscName, err)
			}
			t.Logf("%s_%s Read and Set ok", RsType, rscName)

			var dataFromModel []byte

			modelDataRead := v.GetModelData()

			dataFromModel, err = json.Marshal(&modelDataRead)
			if err != nil {
				t.Errorf("Json marshal failed %s_%s: %v", RsType, rscName, err)
			}
			t.Logf("%s_%s Marshall ok", RsType, rscName)

			require.JSONEq(t, string(dataFromModel), string(dataFromFile))

		})
	}
}
