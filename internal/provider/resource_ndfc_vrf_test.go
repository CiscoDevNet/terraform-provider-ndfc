// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

//template:begin imports
import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

//template:end imports

//template:begin testAcc
func TestAccNdfcVRF(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNdfcVRFConfigMinimal,
			},
			{
				Config: testAccNdfcVRFConfigAll,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ndfc_vrf.test", "fabric_name", "CML"),
					resource.TestCheckResourceAttr("ndfc_vrf.test", "vrf_name", "VRF1"),
					resource.TestCheckResourceAttr("ndfc_vrf.test", "vrf_template", "Default_VRF_Universal"),
					resource.TestCheckResourceAttr("ndfc_vrf.test", "vrf_extension_template", "Default_VRF_Extension_Universal"),
					resource.TestCheckResourceAttr("ndfc_vrf.test", "vrf_id", "50000"),
					resource.TestCheckResourceAttr("ndfc_vrf.test", "vlan_id", "1500"),
					resource.TestCheckResourceAttr("ndfc_vrf.test", "vlan_name", "VLAN1500"),
					resource.TestCheckResourceAttr("ndfc_vrf.test", "interface_description", "My int description"),
					resource.TestCheckResourceAttr("ndfc_vrf.test", "vrf_description", "My vrf description"),
					resource.TestCheckResourceAttr("ndfc_vrf.test", "mtu", "9200"),
					resource.TestCheckResourceAttr("ndfc_vrf.test", "loopback_routing_tag", "11111"),
					resource.TestCheckResourceAttr("ndfc_vrf.test", "redistribute_direct_route_map", "FABRIC-RMAP-REDIST"),
					resource.TestCheckResourceAttr("ndfc_vrf.test", "max_bgp_paths", "2"),
					resource.TestCheckResourceAttr("ndfc_vrf.test", "max_ibgp_paths", "3"),
					resource.TestCheckResourceAttr("ndfc_vrf.test", "ipv6_link_local", "false"),
					resource.TestCheckResourceAttr("ndfc_vrf.test", "trm", "true"),
					resource.TestCheckResourceAttr("ndfc_vrf.test", "no_rp", "false"),
					resource.TestCheckResourceAttr("ndfc_vrf.test", "rp_external", "true"),
					resource.TestCheckResourceAttr("ndfc_vrf.test", "rp_address", "1.2.3.4"),
					resource.TestCheckResourceAttr("ndfc_vrf.test", "rp_loopback_id", "100"),
					resource.TestCheckResourceAttr("ndfc_vrf.test", "underlay_multicast_address", "233.1.1.1"),
					resource.TestCheckResourceAttr("ndfc_vrf.test", "overlay_multicast_groups", "234.0.0.0/8"),
					resource.TestCheckResourceAttr("ndfc_vrf.test", "mvpn_inter_as", "false"),
					resource.TestCheckResourceAttr("ndfc_vrf.test", "trm_bgw_msite", "true"),
					resource.TestCheckResourceAttr("ndfc_vrf.test", "advertise_host_routes", "true"),
					resource.TestCheckResourceAttr("ndfc_vrf.test", "advertise_default_route", "false"),
					resource.TestCheckResourceAttr("ndfc_vrf.test", "configure_static_default_route", "false"),
					resource.TestCheckResourceAttr("ndfc_vrf.test", "bgp_password", "1234567890ABCDEF"),
					resource.TestCheckResourceAttr("ndfc_vrf.test", "bgp_password_type", "7"),
					resource.TestCheckResourceAttr("ndfc_vrf.test", "netflow", "false"),
					resource.TestCheckResourceAttr("ndfc_vrf.test", "netflow_monitor", "MON1"),
					resource.TestCheckResourceAttr("ndfc_vrf.test", "disable_rt_auto", "true"),
					resource.TestCheckResourceAttr("ndfc_vrf.test", "route_target_import", "1:1"),
					resource.TestCheckResourceAttr("ndfc_vrf.test", "route_target_export", "1:1"),
					resource.TestCheckResourceAttr("ndfc_vrf.test", "route_target_import_evpn", "1:1"),
					resource.TestCheckResourceAttr("ndfc_vrf.test", "route_target_export_evpn", "1:1"),
					resource.TestCheckResourceAttr("ndfc_vrf.test", "route_target_import_cloud_evpn", "1:1"),
					resource.TestCheckResourceAttr("ndfc_vrf.test", "route_target_export_cloud_evpn", "1:1"),
					resource.TestCheckResourceAttr("ndfc_vrf.test", "attachments.0.serial_number", "9DBYO6WQJ46"),
					resource.TestCheckResourceAttr("ndfc_vrf.test", "attachments.0.vlan_id", "2000"),
				),
			},
			{
				ResourceName:  "ndfc_vrf.test",
				ImportState:   true,
				ImportStateId: "CML:VRF1",
			},
		},
	})
}

//template:end testAcc

//template:begin testAccConfigMinimal
const testAccNdfcVRFConfigMinimal = `

resource "ndfc_vrf" "test" {
	fabric_name = "CML"
	vrf_name = "VRF1"
}
`

//template:end testAccConfigMinimal

//template:begin testAccConfigAll
const testAccNdfcVRFConfigAll = `

resource "ndfc_vrf" "test" {
	fabric_name = "CML"
	vrf_name = "VRF1"
	vrf_template = "Default_VRF_Universal"
	vrf_extension_template = "Default_VRF_Extension_Universal"
	vrf_id = 50000
	vlan_id = 1500
	vlan_name = "VLAN1500"
	interface_description = "My int description"
	vrf_description = "My vrf description"
	mtu = 9200
	loopback_routing_tag = 11111
	redistribute_direct_route_map = "FABRIC-RMAP-REDIST"
	max_bgp_paths = 2
	max_ibgp_paths = 3
	ipv6_link_local = false
	trm = true
	no_rp = false
	rp_external = true
	rp_address = "1.2.3.4"
	rp_loopback_id = 100
	underlay_multicast_address = "233.1.1.1"
	overlay_multicast_groups = "234.0.0.0/8"
	mvpn_inter_as = false
	trm_bgw_msite = true
	advertise_host_routes = true
	advertise_default_route = false
	configure_static_default_route = false
	bgp_password = "1234567890ABCDEF"
	bgp_password_type = "7"
	netflow = false
	netflow_monitor = "MON1"
	disable_rt_auto = true
	route_target_import = "1:1"
	route_target_export = "1:1"
	route_target_import_evpn = "1:1"
	route_target_export_evpn = "1:1"
	route_target_import_cloud_evpn = "1:1"
	route_target_export_cloud_evpn = "1:1"
	attachments = [{
		serial_number = "9DBYO6WQJ46"
		vlan_id = 2000
	}]
}
`

//template:end testAccConfigAll
