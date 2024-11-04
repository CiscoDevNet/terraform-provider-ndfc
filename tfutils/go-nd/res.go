// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package nd

import (
	"github.com/tidwall/gjson"
)

// Res is an API response returned by client requests.
// This is a GJSON result, which offers advanced and safe parsing capabilities.
// https://github.com/tidwall/gjson
type Res = gjson.Result
