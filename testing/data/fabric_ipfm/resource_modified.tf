
resource "ndfc_fabric_ipfm" "test_resource_fabric_ipfm_1" {
  fabric_name              = "IP_FABRIC_MEDIA"
  aaa_remote_ip_enabled    = false
  bootstrap_enable         = false
  bootstrap_multisubnet    = "#Scope_Start_IP, Scope_End_IP, Scope_Default_Gateway, Scope_Subnet_Prefix"
  cdp_enable               = true
  dhcp_enable              = false
  enable_aaa               = false
  enable_asm               = false
  enable_nbm_passive       = false
  fabric_interface_type    = "p2p"
  fabric_mtu               = 9230
  feature_ptp              = false
  isis_auth_enable         = false
  isis_level               = "level-2"
  l2_host_intf_mtu         = 9200
  link_state_routing       = "ospf"
  link_state_routing_tag   = "1"
  loopback0_ip_range       = "10.2.0.0/22"
  nxapi_vrf                = "management"
  ospf_area_id             = "0.0.0.0"
  ospf_auth_enable         = false
  pim_hello_auth_enable    = false
  pm_enable                = true
  power_redundancy_mode    = "ps-redundant"
  routing_lb_id            = 0
  snmp_server_host_trap    = true
  static_underlay_ip_alloc = false
  subnet_range             = "10.4.0.0/16"
  subnet_target_mask       = 30
  deploy                   = true
}
