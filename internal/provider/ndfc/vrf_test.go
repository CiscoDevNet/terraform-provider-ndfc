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
	"reflect"
	"terraform-provider-ndfc/internal/provider/resources/resource_vrf_bulk"
	"terraform-provider-ndfc/tfutils/go-nd"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TestNDFC_RscCreateBulkVrf(t *testing.T) {
	type fields struct {
		url       string
		apiClient nd.Client
	}
	type args struct {
		ctx     context.Context
		dg      *diag.Diagnostics
		vrfBulk *resource_vrf_bulk.VrfBulkModel
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *resource_vrf_bulk.VrfBulkModel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NDFC{
				url:       tt.fields.url,
				apiClient: tt.fields.apiClient,
			}
			if got := c.RscCreateBulkVrf(tt.args.ctx, tt.args.dg, tt.args.vrfBulk); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NDFC.RscCreateBulkVrf() = %v, want %v", got, tt.want)
			}
		})
	}
}
