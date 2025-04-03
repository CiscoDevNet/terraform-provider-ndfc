// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package ndfc

import (
	"terraform-provider-ndfc/internal/provider/resources/resource_networks"
	"terraform-provider-ndfc/internal/provider/resources/resource_vrf_attachments"
	"terraform-provider-ndfc/internal/provider/resources/resource_vrf_bulk"
	"testing"
)

func TestNDFC_RscCreateID(t *testing.T) {

	type args struct {
		rsc     interface{}
		rscType string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Networks",
			args: args{
				rsc: &resource_networks.NDFCNetworksModel{
					FabricName: "fabric1",
					Networks: map[string]resource_networks.NDFCNetworksValue{
						"network1": {
							NetworkName: "network1",
						},
						"network2": {
							NetworkName: "network2",
						},
						"network3": {
							NetworkName: "network3",
						},
					},
				},
				rscType: ResourceNetworks,
			},
			want: "fabric1/[network1,network2,network3]",
		},
		{
			name: "VrfBulk",
			args: args{
				rsc: &resource_vrf_bulk.NDFCVrfBulkModel{
					FabricName: "fabric1",
					Vrfs: map[string]resource_vrf_bulk.NDFCVrfsValue{
						"vrf1": {
							VrfName: "vrf1",
							AttachList: map[string]resource_vrf_attachments.NDFCAttachListValue{
								"attach1": {
									SerialNumber: "attach1",
								},
								"attach2": {
									SerialNumber: "attach2",
								},
							},
						},
					},
				},
				rscType: ResourceVrfBulk,
			},
			want: "fabric1/[vrf1{attach1,attach2}]",
		},
		{
			name: "Default",
			args: args{
				rsc:     nil,
				rscType: "",
			},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NDFC{}
			if got := c.RscCreateID(tt.args.rsc, tt.args.rscType); got != tt.want {
				t.Errorf("NDFC.RscCreateID() = %v, want %v", got, tt.want)
			}
		})
	}
}
