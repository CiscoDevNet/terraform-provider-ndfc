
resource "ndfc_fabric_vxlan_msd" "test_resource_fabric_vxlan_msd_1" {
  fabric_name                = "TF_FABRIC_VXLAN_MSD"
  anycast_gw_mac             = "2020.0000.00bb"
  bgw_routing_tag            = 34451
  border_gwy_connections     = "Manual"
  cloudsec_autoconfig        = false
  dci_subnet_range           = "10.10.1.0/24"
  dci_subnet_target_mask     = 30
  delay_restore              = 300
  enable_pvlan               = false
  l2_segment_id_range        = "30000-49000"
  l3_partition_id_range      = "50000-59000"
  loopback100_ip_range       = "10.10.0.0/24"
  ms_ifc_bgp_password_enable = false
  ms_loopback_id             = 100
  ms_underlay_autoconfig     = false
  tor_auto_deploy            = false
  default_network            = "Default_Network_Universal"
  default_vrf                = "Default_VRF_Universal"
  network_extension_template = "Default_Network_Extension_Universal"
  vrf_extension_template     = "Default_VRF_Extension_Universal"
  deploy                     = true
}
