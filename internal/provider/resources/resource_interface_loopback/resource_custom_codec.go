// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_interface_loopback

import "github.com/hashicorp/terraform-plugin-framework/types"

func (i *InterfaceLoopbackModel) GetInterfaceType() string {
	return "loopback"
}

func (i *InterfaceLoopbackModel) GetID() string {
	return i.Id.ValueString()
}

func (i *InterfaceLoopbackModel) SetID(id string) {
	if id == "" {
		i.Id = types.StringNull()
		return
	}
	i.Id = types.StringValue(id)
}
