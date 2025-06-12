
resource "ndfc_links" "test_resource_links_1" {
  source_fabric         = "CML1"
  destination_fabric    = "CML2"
  source_device         = "SERIAL_1"
  destination_device    = "SERIAL_2"
  source_interface      = "Ethernet1/1"
  destination_interface = "Ethernet1/1"
  template_name         = "ext_fabric_setup"
  link_parameters = {
    # Basic Parameters
    "asn"           = "65001"
    "NEIGHBOR_ASN"  = "65002"
    "IP_MASK"       = "192.168.1.1/24"
    "NEIGHBOR_IP"   = "192.168.1.2"
    "IPV6_MASK"     = "2001:db8::1/64"
    "IPV6_NEIGHBOR" = "2001:db8::2"

    # Interface Parameters
    "MTU"        = "9216"
    "DOT1Q_ID"   = "2"
    "PEER1_DESC" = "Link to remote site A"
    "PEER2_DESC" = "Link to remote site B"

    # VRF Parameters
    "AUTO_VRF_LITE_FLAG"            = "false"
    "DEFAULT_VRF_FLAG"              = "false"
    "SYMMETRIC_DEFAULT_VRF_FLAG"    = "false"
    "DEFAULT_VRF_REDIS_BGP_RMAP"    = "extcon-rmap-filter"
    "DEFAULT_VRF_BGP_AUTH_KEY_TYPE" = "3"
    "ENABLE_DCI_TRACKING"           = "false"
    "VRF_LITE_JYTHON_TEMPLATE"      = "Ext_VRF_Lite_Jython"

    # Advanced Parameters
    "PEER1_CONF"                = "no shutdown"
    "PEER2_CONF"                = "no shutdown"
    "PEER_VRF_NAME"             = "default"
    "DEFAULT_VRF_PEER_VRF_NAME" = "default"
  }

}