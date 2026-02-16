# =============================================================================
# Child Fabrics (VXLAN EVPN) — managed externally, not by this module.
# =============================================================================

# =============================================================================
# MSD Parent Fabric
# =============================================================================

resource "ndfc_fabric_vxlan_msd" "parent" {

  fabric_name   = var.msd_fabric_name
  child_fabrics = keys(var.child_fabrics)

  deploy = var.msd_fabric_config.deploy

  anycast_gw_mac                 = try(var.msd_fabric_config.anycast_gw_mac, null)
  bgp_rp_asn                     = try(var.msd_fabric_config.bgp_rp_asn, null)
  bgw_routing_tag                = try(var.msd_fabric_config.bgw_routing_tag, null)
  border_gwy_connections         = try(var.msd_fabric_config.border_gwy_connections, null)
  cloudsec_algorithm             = try(var.msd_fabric_config.cloudsec_algorithm, null)
  cloudsec_autoconfig            = try(var.msd_fabric_config.cloudsec_autoconfig, null)
  cloudsec_enforcement           = try(var.msd_fabric_config.cloudsec_enforcement, null)
  cloudsec_key_string            = try(var.msd_fabric_config.cloudsec_key_string, null)
  cloudsec_report_timer          = try(var.msd_fabric_config.cloudsec_report_timer, null)
  dci_subnet_range               = try(var.msd_fabric_config.dci_subnet_range, null)
  dci_subnet_target_mask         = try(var.msd_fabric_config.dci_subnet_target_mask, null)
  default_network                = try(var.msd_fabric_config.default_network, null)
  default_pvlan_sec_network      = try(var.msd_fabric_config.default_pvlan_sec_network, null)
  default_vrf                    = try(var.msd_fabric_config.default_vrf, null)
  delay_restore                  = try(var.msd_fabric_config.delay_restore, null)
  enable_bgp_bfd                 = try(var.msd_fabric_config.enable_bgp_bfd, null)
  enable_bgp_log_neighbor_change = try(var.msd_fabric_config.enable_bgp_log_neighbor_change, null)
  enable_bgp_send_comm           = try(var.msd_fabric_config.enable_bgp_send_comm, null)
  enable_pvlan                   = try(var.msd_fabric_config.enable_pvlan, null)
  enable_rs_redist_direct        = try(var.msd_fabric_config.enable_rs_redist_direct, null)
  enable_scheduled_backup        = try(var.msd_fabric_config.enable_scheduled_backup, null)
  enable_sgt                     = try(var.msd_fabric_config.enable_sgt, null)
  enable_trm_trmv6               = try(var.msd_fabric_config.enable_trm_trmv6, null)
  ext_fabric_type                = try(var.msd_fabric_config.ext_fabric_type, null)
  l2_segment_id_range            = try(var.msd_fabric_config.l2_segment_id_range, null)
  l3_partition_id_range          = try(var.msd_fabric_config.l3_partition_id_range, null)
  loopback100_ip_range           = try(var.msd_fabric_config.loopback100_ip_range, null)
  loopback100_ipv6_range         = try(var.msd_fabric_config.loopback100_ipv6_range, null)
  ms_ifc_bgp_auth_key_type       = try(var.msd_fabric_config.ms_ifc_bgp_auth_key_type, null)
  ms_ifc_bgp_password            = try(var.msd_fabric_config.ms_ifc_bgp_password, null)
  ms_ifc_bgp_password_enable     = try(var.msd_fabric_config.ms_ifc_bgp_password_enable, null)
  ms_loopback_id                 = try(var.msd_fabric_config.ms_loopback_id, null)
  ms_underlay_autoconfig         = try(var.msd_fabric_config.ms_underlay_autoconfig, null)
  network_extension_template     = try(var.msd_fabric_config.network_extension_template, null)
  rp_server_ip                   = try(var.msd_fabric_config.rp_server_ip, null)
  rs_routing_tag                 = try(var.msd_fabric_config.rs_routing_tag, null)
  scheduled_time                 = try(var.msd_fabric_config.scheduled_time, null)
  sgt_id_range                   = try(var.msd_fabric_config.sgt_id_range, null)
  sgt_name_prefix                = try(var.msd_fabric_config.sgt_name_prefix, null)
  sgt_preprovision               = try(var.msd_fabric_config.sgt_preprovision, null)
  tor_auto_deploy                = try(var.msd_fabric_config.tor_auto_deploy, null)
  v6_dci_subnet_range            = try(var.msd_fabric_config.v6_dci_subnet_range, null)
  v6_dci_subnet_target_mask      = try(var.msd_fabric_config.v6_dci_subnet_target_mask, null)
  vrf_extension_template         = try(var.msd_fabric_config.vrf_extension_template, null)
  vxlan_underlay_is_v6           = try(var.msd_fabric_config.vxlan_underlay_is_v6, null)
}
