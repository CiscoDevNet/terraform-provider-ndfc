// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated DO NOT EDIT.
package resource_vrf_bulk

import (
	"testing"
	."terraform-provider-ndfc/internal/provider/types"
	"github.com/stretchr/testify/assert"
)

func TestInt64Custom_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		i       *Int64Custom
		args    args
		wantErr bool
		value  int64
	}{
		{
			name: "Normal",
			i: new(Int64Custom),
			args: args{data: []byte("\"2000\"")},
			wantErr: false,
			value: 2000,

		},
		{
			name: "empty",
			i: new(Int64Custom),
			args: args{data: []byte("")},
			wantErr: false,
			value: -9223372036854775808,

		},
		{
			name: "quoted empty",
			i: new(Int64Custom),
			args: args{data: []byte("\"\"")},
			wantErr: false,
			value: -9223372036854775808,

		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.i.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Int64Custom.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Logf("%d", *tt.i)
			assert.Equal(t, tt.value, int64(*tt.i))
		})
	}
}
