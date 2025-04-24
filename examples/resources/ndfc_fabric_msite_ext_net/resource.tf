
resource "ndfc_fabric_msite_ext_net" "test_resource_fabric_msite_ext_net_1" {
  fabric_name           = "TF_FABRIC_MSITE_EXT_NET"
  aaa_remote_ip_enabled = false
  bgp_as                = "65000"
  bootstrap_multisubnet = "#Scope_Start_IP, Scope_End_IP, Scope_Default_Gateway, Scope_Subnet_Prefix"
  cdp_enable            = false
  dhcp_enable           = false
  enable_netflow        = false
  enable_nxapi          = false
  enable_nxapi_http     = false
  feature_ptp           = false
  inband_mgmt           = false
  is_read_only          = true
  mpls_handoff          = false
  netflow_exporter_list = jsonencode(
    {
      "NETFLOW_EXPORTER_LIST" : [
        {
          "EXPORTER_NAME" : "Test2",
          "IP" : "10.1.1.1",
          "VRF" : "",
          "SRC_IF_NAME" : "eth1/1",
          "UDP_PORT" : "800"
        }
      ]
    }
  )
  netflow_monitor_list = jsonencode(
    {
      "NETFLOW_MONITOR_LIST" : [
        {
          "MONITOR_NAME" : "Test",
          "RECORD_NAME" : "Test1",
          "EXPORTER1" : "Test2",
          "EXPORTER2" : ""
        }
      ]
    }
  )
  netflow_record_list = jsonencode(
    {
      "NETFLOW_RECORD_LIST" : [
        {
          "RECORD_NAME" : "Test1",
          "RECORD_TEMPLATE" : "netflow_ipv4_record",
          "LAYER2_RECORD" : "false"
        }
      ]
    }
  )
  netflow_sampler_list = jsonencode(
    {
      "NETFLOW_SAMPLER_LIST" : [
        {
          "SAMPLER_NAME" : "Test1",
          "NUM_SAMPLES" : 12,
          "SAMPLING_RATE" : 10
        }
      ]
    }
  )
  nxapi_https_port      = 443
  nxapi_http_port       = 80
  pm_enable             = false
  power_redundancy_mode = "ps-redundant"
  snmp_server_host_trap = true
  subinterface_range    = "2-511"
  deploy                = false
}