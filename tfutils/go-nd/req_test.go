// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package nd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSetRaw tests the Body::SetRaw method.
func TestSetRaw(t *testing.T) {
	name := Body{}.SetRaw("a", `{"name":"a"}`).Res().Get("a.name").Str
	assert.Equal(t, "a", name)
}

// TestDelete tests the Body::Delete method.
func TestDelete(t *testing.T) {
	body := Body{}
	body = body.SetRaw("a", `{"name":"a"}`)
	assert.Equal(t, "a", body.Res().Get("a.name").Str)
	body = body.Delete("a.name")
	assert.Equal(t, "", body.Res().Get("a.name").Str)
}
