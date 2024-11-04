// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package testing

import (
	"terraform-provider-ndfc/internal/provider/resources/resource_vpc_pair"
)

func GenerateVpcPairResource(VpcPairObj **resource_vpc_pair.NDFCVpcPairModel, serials []string, VirtualPeerLink bool) {
	vpcPair := *VpcPairObj
	vpcPair.SerialNumbers = serials[:2]
	*vpcPair.UseVirtualPeerlink = VirtualPeerLink
}
func GenerateVpcPairResourceWithMoreSerials(VpcPairObj **resource_vpc_pair.NDFCVpcPairModel, serials []string, VirtualPeerLink bool) {
	vpcPair := *VpcPairObj
	vpcPair.SerialNumbers = serials
	*vpcPair.UseVirtualPeerlink = VirtualPeerLink
}
