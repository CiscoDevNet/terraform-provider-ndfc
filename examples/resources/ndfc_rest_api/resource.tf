
resource "ndfc_rest_api" "test_resource_rest_api_1" {
  url    = "/lan-fabric/rest/control/policies"
  method = "POST"
  payload = jsonencode(
    {
      "nvPairs" : {
        "IF_NAME" : "ethernet1/20",
        "VRF_NAME" : "red",
        "IP_MASK" : "20.1.1.1/24",
        "NEIGHBOR_IP" : "20.1.1.12",
        "IPv6_MASK" : "",
        "NEIGHBOR_IPv6" : "",
        "NEIGHBOR_ASN" : "86000",
        "asn" : "",
        "bgpPassword" : "",
        "bgpPasswordKeyType" : "",
        "MTU" : "9216",
        "SERIAL_NUMBER" : "",
        "SOURCE" : "",
        "POLICY_ID" : "",
        "ROUTING_TAG" : "",
        "ROUTE_MAP_IN" : "",
        "ROUTE_MAP_OUT" : "",
        "IPV6_ROUTE_MAP_IN" : "",
        "IPV6_ROUTE_MAP_OUT" : "",
        "DESC" : "",
        "CONF" : "",
        "ADMIN_STATE" : "true"
      },
      "entityName" : "SWITCH",
      "entityType" : "SWITCH",
      "source" : "",
      "priority" : 500,
      "description" : "",
      "templateName" : "Ext_VRF_Lite_Routed",
      "serialNumber" : "9TQYTJSZ1VJ"
    }

  )
  stateful      = false
  update_method = "PUT"
  update_payload = jsonencode(
    {
      "nvPairs" : {
        "IF_NAME" : "ethernet1/20",
        "VRF_NAME" : "red",
        "IP_MASK" : "20.1.1.1/24",
        "NEIGHBOR_IP" : "20.1.1.12",
        "IPv6_MASK" : "",
        "NEIGHBOR_IPv6" : "",
        "NEIGHBOR_ASN" : "86000",
        "asn" : "",
        "bgpPassword" : "",
        "bgpPasswordKeyType" : "",
        "MTU" : "9216",
        "SERIAL_NUMBER" : "",
        "SOURCE" : "",
        "POLICY_ID" : "",
        "ROUTING_TAG" : "",
        "ROUTE_MAP_IN" : "",
        "ROUTE_MAP_OUT" : "",
        "IPV6_ROUTE_MAP_IN" : "",
        "IPV6_ROUTE_MAP_OUT" : "",
        "DESC" : "",
        "CONF" : "",
        "ADMIN_STATE" : "true"
      },
      "id" : "{{.id}}",
      "entityName" : "SWITCH",
      "entityType" : "SWITCH",
      "source" : "",
      "priority" : 500,
      "description" : "New description",
      "templateName" : "Ext_VRF_Lite_Routed",
      "serialNumber" : "9TQYTJSZ1VJ"
    }

  )
  update_url = "/lan-fabric/rest/control/policies/{{.policyId}}"
  delete_url = "/lan-fabric/rest/control/policies/{{.policyId}}"
  delete_payload = jsonencode(
    {
      "nvPairs" : {
        "IF_NAME" : "ethernet1/20",
        "VRF_NAME" : "red",
        "IP_MASK" : "20.1.1.1/24",
        "NEIGHBOR_IP" : "20.1.1.12",
        "IPv6_MASK" : "",
        "NEIGHBOR_IPv6" : "",
        "NEIGHBOR_ASN" : "86000",
        "asn" : "",
        "bgpPassword" : "",
        "bgpPasswordKeyType" : "",
        "MTU" : "9216",
        "SERIAL_NUMBER" : "",
        "SOURCE" : "",
        "POLICY_ID" : "",
        "ROUTING_TAG" : "",
        "ROUTE_MAP_IN" : "",
        "ROUTE_MAP_OUT" : "",
        "IPV6_ROUTE_MAP_IN" : "",
        "IPV6_ROUTE_MAP_OUT" : "",
        "DESC" : "",
        "CONF" : "",
        "ADMIN_STATE" : "true"
      },
      "id" : "{{.id}}",
      "deleted" : true,
      "entityName" : "SWITCH",
      "entityType" : "SWITCH",
      "source" : "",
      "priority" : 500,
      "description" : "New description",
      "templateName" : "Ext_VRF_Lite_Routed",
      "serialNumber" : "9TQYTJSZ1VJ"
    }

  )
}