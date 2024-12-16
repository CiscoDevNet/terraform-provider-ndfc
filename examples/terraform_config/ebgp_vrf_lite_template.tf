
resource "ndfc_template" "ebgp_vrf_lite_template" {
  template_name     = "tf_ebgp_vrf_lite_template"
  description       = "tf_ebgp_vrf_lite_template"
  template_type     = "POLICY"
  template_sub_type = "DEVICE"
  content_type      = "TEMPLATE_CLI"
  tags              = ["tf_test_ebgp_vrf_lite_template", "my_template_tag_4", "tag4"]
  template_content  = <<EOF
##template variables
#    Copyright (c) 2020 by Cisco Systems, Inc.
#    All rights reserved.
@(IsMandatory=true, DisplayName="VRF Name")
string vrfName;
@(IsMandatory=false, DisplayName="VRF Description")
string vrfDescription;
@(IsMandatory=true, DisplayName="Local ASN")
string asn;
@(IsMandatory=false, DisplayName="Neighbor IPv4 Address")
ipV4Address NEIGHBOR_IP;
@(IsMandatory=true, DisplayName="Neighbor ASN")
string NEIGHBOR_ASN;
@(IsMandatory=false, DisplayName="BGP Neighbor Password", Description="Hex String")
string bgpPassword;
@(IsMandatory=false, DisplayName="BGP Password<br/>Key Encryption Type", Description="BGP Key Encryption Type: 3 - 3DES, 7 - Cisco")
enum bgpPasswordKeyType {
    validValues=3,7;
    defaultValue=3;
};

@(IsMandatory=false, DisplayName="Neighbor IPv6 Address")
ipV6Address IPV6_NEIGHBOR;
##
##template content
if ($$vrfName$$ != "default" && $$vrfName$$ != "") {
vrf context $$vrfName$$
 if ($$vrfDescription$$ != "") {
  description $$vrfDescription$$
 }
 if ($$NEIGHBOR_IP$$ !="") {
  address-family ipv4 unicast
 }
 if ($$IPV6_NEIGHBOR$$ !="") {
  address-family ipv6 unicast
 }
}

router bgp $$asn$$
#handle non-default vrf case
if ($$vrfName$$ != "default" && $$vrfName$$ != "") {
  vrf $$vrfName$$
  if ($$NEIGHBOR_IP$$ !="") {
    address-family ipv4 unicast
  }
  if ($$IPV6_NEIGHBOR$$ !="") {
    address-family ipv6 unicast
  }    
  if ($$NEIGHBOR_IP$$ != "" && $$NEIGHBOR_ASN$$ !="") {
    neighbor $$NEIGHBOR_IP$$ remote-as $$NEIGHBOR_ASN$$
      if ($$bgpPassword$$ != "") {
      password $$bgpPasswordKeyType$$ $$bgpPassword$$
      }
      address-family ipv4 unicast
        send-community
        send-community extended
  }
  if ($$IPV6_NEIGHBOR$$ != "" && $$NEIGHBOR_ASN$$ !="") {
    neighbor $$IPV6_NEIGHBOR$$ remote-as $$NEIGHBOR_ASN$$
      if ($$bgpPassword$$ != "") {
      password $$bgpPasswordKeyType$$ $$bgpPassword$$
      }
      address-family ipv6 unicast
        send-community
        send-community extended
  }
}
else {
#handle default vrf case
 if ($$NEIGHBOR_IP$$ != "" && $$NEIGHBOR_ASN$$ !="") {
  neighbor $$NEIGHBOR_IP$$ remote-as $$NEIGHBOR_ASN$$
    if ($$bgpPassword$$ != "") {
    password $$bgpPasswordKeyType$$ $$bgpPassword$$
    }
    address-family ipv4 unicast
      send-community
      send-community extended
 }
 if ($$IPV6_NEIGHBOR$$ != "" && $$NEIGHBOR_ASN$$ !="") {
  neighbor $$IPV6_NEIGHBOR$$ remote-as $$NEIGHBOR_ASN$$
    if ($$bgpPassword$$ != "") {
    password $$bgpPasswordKeyType$$ $$bgpPassword$$
    }
    address-family ipv6 unicast
      send-community
      send-community extended
 }
}
##
EOF
}