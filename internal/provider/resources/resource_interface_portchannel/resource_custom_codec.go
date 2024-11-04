// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_interface_portchannel

import "github.com/hashicorp/terraform-plugin-framework/types"

func (i *InterfacePortchannelModel) GetInterfaceType() string {
	return "portchannel"
}

func (i *InterfacePortchannelModel) GetID() string {
	return i.Id.ValueString()
}

func (i *InterfacePortchannelModel) SetID(id string) {
	if id == "" {
		i.Id = types.StringNull()
		return
	}
	i.Id = types.StringValue(id)
}
