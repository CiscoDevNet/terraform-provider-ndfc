terraform {
  required_providers {
    ndfc = {
      source = "registry.terraform.io/cisco/ndfc"
    }
  }
}

provider "ndfc" {
  username = "test"
  password = "test"
  host     = "https://10.78.210.161"
  insecure = true
}

import {
  id = "int_trunk_host:9FE076D8EJL"
  to = ndfc_interface_ethernet.test_evpn_vxlan_eth_imported
}
import {
  id = "int_loopback:9FE076D8EJL"
  to = ndfc_interface_loopback.test_evpn_vxlan_lb_imported_1
}
import {
  id = "int_vlan:9FE076D8EJL"
  to = ndfc_interface_vlan.test_evpn_vxlan_vlan_imported_1
}

import {
  id = "9D0ZV7JBFNM:9N14Y8PVD2Y"
  to = ndfc_vpc_pair.test_vpc_pair
}
resource "ndfc_interface_ethernet" "test_evpn_vxlan_eth_imported" {
}

resource "ndfc_interface_loopback" "test_evpn_vxlan_lb_imported_1" {
}

resource "ndfc_interface_vlan" "test_evpn_vxlan_vlan_imported_1" {
}

resource "ndfc_interface_portchannel" "test_evpn_vxlan_pc_imported_1" {
}
resource "ndfc_interface_vpc" "test_evpn_vxlan_vpc_imported_1" {
}

resource "ndfc_vpc_pair" "test_vpc_pair" {
}

