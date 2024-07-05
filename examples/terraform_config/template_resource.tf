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

resource "ndfc_template" "test_policy_template_1" {
  template_name = "tf_test_policy_template_1"
  description = "tf_test_policy_template_1"
  template_type = "POLICY"
  template_sub_type = "INTERFACE_ETHERNET"
  content_type = "TEMPLATE_CLI"
  tags = ["tf_test_policy_template_1", "my_template_tag_1", "tag3"]
  template_content = <<EOF
##template variables

#    Copyright (c) 2018 by Cisco Systems, Inc.
#    All rights reserved.

@(DisplayName="Access Interface", Description="Access host port")
interface INTF_NAME;

##
##template content

interface $$INTF_NAME$$
  switchport mode access
  no shutdown

##
EOF
}

resource "ndfc_template" "test_policy_ext_vrf_lite_svi" {
  template_name = "tf_test_policy_ext_vrf_lite_svi"
  description = "tf_test_policy_ext_vrf_lite_svi created from TF"
  template_type = "POLICY"
  template_sub_type = "DEVICE"
  content_type = "PYTHON"
  tags = ["tf_test_policy_ext_vrf_lite_svi", "my_template_tag_2", "tag2"]
  template_content = <<EOF
##template variables

#    Copyright (c) 2021 by Cisco Systems, Inc.
#    All rights reserved.

@(IsInternal=true)
string SERIAL_NUMBER;

@(IsInternal=true)
string SOURCE;

@(IsInternal=true)
string POLICY_ID;
#this is a test change

@(DisplayName="VLAN ID", Description="2-4094, value needs to be in switch allowed range")
integer DOT1Q_ID {
  min = 2;
  max = 4094;
};

@(IsMandatory=false, DisplayName="VRF Name", Description="SVI VRF name, default VRF if not specified")
string VRF_NAME {
  minLength = 1;
  maxLength = 32;
};

@(IsMandatory=false, DisplayName="SVI IPv4 Address/Netmask", Description="For IPv4 VRF Lite peering")
ipV4AddressWithSubnet IP_MASK;

@(IsMandatory="IP_MASK!=null", IsShow="IP_MASK!=null", DisplayName="Neighbor IPv4 Address", Description="BGP Peer IPv4 Address")
ipV4Address NEIGHBOR_IP;

@(IsMandatory=false, DisplayName="SVI IPv6 Address/Netmask", Description="For IPv6 VRF Lite peering")
ipV6AddressWithSubnet IPv6_MASK;

@(IsMandatory="IPv6_MASK!=null", IsShow="IPv6_MASK!=null", DisplayName="Neighbor IPv6 Address", Description="BGP Peer IPv6 Address")
ipV6Address NEIGHBOR_IPv6;

@(DisplayName="Neighbor ASN", Description="BGP ASN of IPv4/IPv6 Neighbor")
string NEIGHBOR_ASN;

@(IsMandatory=false, DisplayName="Local ASN", Description="The fabric ASN will be used if not specified")
string asn;

@(IsMandatory=false, DisplayName="BGP Neighbor Password", Description="Hex String")
string bgpPassword {
minLength = 1;
maxLength = 130;
regularExpr=^[a-fA-F0-9]+$;
};

@(IsMandatory=false, IsShow="bgpPassword!=null", DisplayName="BGP Password Key Encryption Type", Description="BGP Key Encryption Type: 3 - 3DES, 7 - Cisco")
enum bgpPasswordKeyType {
    validValues=3,7;
    defaultValue=3;
};

@(IsMandatory=false, DisplayName="Additional BGP Neighbors", Description="Add peerings to multiple neighbors over the same SVI")
struct BGP_NEIGHBOR {
  @(IsMandatory=false, DisplayName="Neighbor IPv4 Address", Description="BGP Peer IPv4 Address")
  ipV4Address NEIGHBOR_IP;
  @(IsMandatory=false, DisplayName="Neighbor IPv6 Address", Description="BGP Peer IPv6 Address")
  ipV6Address NEIGHBOR_IPv6;
  @(DisplayName="Neighbor ASN", Description="BGP ASN of IPv4/IPv6 Neighbor")
  string NEIGHBOR_ASN;
  @(IsMandatory=false, DisplayName="BGP Neighbor Password", Description="Hex String")
  string bgpPassword {
    minLength = 1;
    maxLength = 130;
    regularExpr=^[a-fA-F0-9]+$;
  };
  @(IsMandatory=false, IsShow="bgpPassword!=null", DisplayName="BGP Password Key Encryption Type", Description="BGP Key Encryption Type: 3 - 3DES, 7 - Cisco")
  enum bgpPasswordKeyType {
    validValues=3,7;
    defaultValue=3;
  };
} BGP_NEIGHBOR_LIST[];

@(DisplayName="SVI MTU", IsMTU=true, Description="Layer-3 MTU (Min:576, Max:9216)")
integer MTU {
min = 576;
max = 9216;
defaultValue=9216;
};

@(IsMandatory=false, DisplayName="Routing TAG", Description="Routing tag associated with interface IP", Section = "Advanced")
string ROUTING_TAG;

@(IsMandatory=false, DisplayName="Disable IP redirects", Description="Disables both IPv4/IPv6 redirects on the interface", Section = "Advanced")
boolean DISABLE_IP_REDIRECTS {
  defaultValue=false;  
};

@(IsMandatory=false, IsShow="IP_MASK!=null", DisplayName="IPv4 Inbound Route Map", Description="No route-map if blank", Section = "Advanced")
string ROUTE_MAP_IN;

@(IsMandatory=false, IsShow="IP_MASK!=null", DisplayName="IPv4 Outbound Route Map", Description="No route-map if blank", Section = "Advanced")
string ROUTE_MAP_OUT;

@(IsMandatory=false, IsShow="IPv6_MASK!=null", DisplayName="IPv6 Inbound Route Map", Description="No route-map if blank", Section = "Advanced")
string IPV6_ROUTE_MAP_IN;

@(IsMandatory=false, IsShow="IPv6_MASK!=null", DisplayName="IPv6 Outbound Route Map", Description="No route-map if blank", Section = "Advanced")
string IPV6_ROUTE_MAP_OUT;

@(IsMandatory=false, DisplayName="Interface Description", Description="Add description to the interface (Max Size 254)", Section = "Advanced")
string DESC {
  minLength = 1;
  maxLength = 254;
};

@(IsMandatory=false, IsMultiLineString=true, DisplayName="Freeform Config", Description="Additional CLI for the interface", Section = "Advanced")
string CONF;

@(DisplayName="Interface Admin State", Description="Admin state of the interface", Section = "Advanced")
boolean ADMIN_STATE {
  defaultValue=true;
};

##
##template content

from com.cisco.dcbu.vinci.rest.services.jython import PTIWrapper
from com.cisco.dcbu.vinci.rest.services.jython import Wrapper
from com.cisco.dcbu.vinci.rest.services.jython import WrappersResp
from com.cisco.dcbu.vinci.rest.services.jython import Helper
from com.cisco.dcbu.vinci.rest.services.jython import InventoryWrapper
from com.cisco.dcbu.vinci.rest.services.jython import InterfaceManagerWrapper
from utility import *
import ast

def add():
    try:
        
        #if called by topdown, SOURCE will be set. Then user will not be able to edit from View/Edit policies
        jython_source  = ""
        try:
            jython_source = SOURCE
        except:
            jython_source = POLICY_ID
            pass
        
        SOURCE = jython_source
        eBGPPassword = ""
        eBGPPasswordKeyType = "3"
        try:
            eBGPPassword = bgpPassword
        except:
            pass
        try:
            eBGPPasswordKeyType = bgpPasswordKeyType
        except:
            pass

        Wrapper.print("Source of VRF Lite Jython is [%s] for switch [%s]. BGP Password [%s]/[%s]" %(SOURCE, SERIAL_NUMBER, 
                                            eBGPPassword, eBGPPasswordKeyType))

        src_intf = "Vlan" + str(DOT1Q_ID)
        Wrapper.print("Interface provided %s" % (src_intf))
        
        srchOpt = CtrlPolicySearch()
        srchOpt.setSerialNumber(SERIAL_NUMBER)
        srchOpt.setSource(SOURCE)
        srchOpt.setTemplateName( "int_vlan_internal")
        ptiList = Util.exe(PTIWrapper.getPTIs(srchOpt))
        if ptiList:
            for pti in ptiList:
                oldInterface = pti.getEntityName()
                #Compare the interface, it should be the same
                if oldInterface != src_intf:
                    Wrapper.print("Error case: Old: %s New %s" %(oldInterface, src_intf))
                    errorMsg = ("New parent interface %s not matching old one %s" % (oldInterface, src_intf))
                    respObj = WrappersResp.getRespObj()
                    respObj.addErrorReport(("Error"), errorMsg)
                    respObj.setFailureRetCode()
                    return respObj

                break
            
        #Validate neighbor ASN format
        respObj = Helper.isValidAsn (NEIGHBOR_ASN)
        if respObj.isRetCodeFailure():
            errorMsg = ("Invalid Neighbor ASN %s" % (NEIGHBOR_ASN))
            respObj.addErrorReport(("Error"), errorMsg)
            return respObj
        
        # get fabric for switch with this serial number
        inventoryResp = InventoryWrapper.getFabricNameForSwitch(SERIAL_NUMBER)
        if inventoryResp.isRetCodeSuccess():
            fabricName = inventoryResp.getValue()
            Wrapper.print("Inventory wrapper response fabric name is %s" %(fabricName))
        else:
            Wrapper.print("Invalid fabric for switch %s" %(SERIAL_NUMBER))
            return inventoryResp
                                           
        # try to honor the ASN passed in
        fabricAsn = local_asn = ""
        config_local_as = False
        try:
            fabricAsn = local_asn = asn
        except:
            pass

        if ((fabricAsn == None) or (fabricAsn == "")):
            # the ASN does not seem valid... use the current fabric ASN
            helperResp = Helper.getFabricAsn(fabricName)
            if helperResp.isRetCodeSuccess():
                fabricAsn = helperResp.getValue()
            else:
                Wrapper.print("Invalid ASN for Fabric %s" %(fabricName))
                return helperResp
        else:
            # local ASN is provided. Check whether needs to override fabric ASN
            fabric_type = Util.exe(Helper.getFabricType(fabricName))
            if fabric_type == "Switch_Fabric":
                helperResp = Helper.getFabricAsn(fabricName)
                if helperResp.isRetCodeSuccess():
                    fabricAsn = helperResp.getValue()
                if local_asn != fabricAsn:
                    config_local_as = True
            
        if ((fabricAsn == None) or (fabricAsn == "")):
            # we still have an invalid ASN
            errorMsg = ("Could not determine the BGP ASN for fabric [%s]" % (fabricName))
            respObj = WrappersResp.getRespObj()
            respObj.addErrorReport("VRF-LITE-AUTO_EXT:INVALID_ASN", errorMsg)
            respObj.setFailureRetCode()
            return respObj

        try:
            if IP_MASK != "":
                ipAddress = IP_MASK.split("/")
                IP = ipAddress[0]
                PREFIX = ipAddress[1]
    
                subnetStr1 = Util.getSubnetStringWithPrefix(IP, PREFIX)
                subnetStr2 = Util.getSubnetStringWithPrefix(NEIGHBOR_IP, PREFIX)
                if subnetStr1 != subnetStr2:
                    errorMsg = ("Local IP address %s and BGP Peer address %s are not in same subnet"%
                                (IP_MASK, NEIGHBOR_IP))
                    respObj.addErrorReport("VRF-LITE-EXT:INVALID_IP", errorMsg)
                    respObj.setFailureRetCode()
                    return respObj
                IP_NEIGHBOR = NEIGHBOR_IP
            else:
                IP = PREFIX = IP_NEIGHBOR = ""
        except:
            IP = PREFIX = IP_NEIGHBOR = ""

        try:
            if IPv6_MASK != "":
                ipv6Address = IPv6_MASK.split("/")
                IPv6 = ipv6Address[0]
                IPv6_PREFIX = ipv6Address[1]

                subnetStr1 = Util.getV6SubnetStringWithPrefix(IPv6, IPv6_PREFIX)
                subnetStr2 = Util.getV6SubnetStringWithPrefix(NEIGHBOR_IPv6, IPv6_PREFIX)
                if subnetStr1 != subnetStr2:
                    errorMsg = ("Local IPv6 address %s and BGP Peer address %s are not in same subnet"%
                                (IPv6_MASK, NEIGHBOR_IPv6))
                    respObj.addErrorReport("VRF-LITE-EXT:INVALID_IP", errorMsg)
                    respObj.setFailureRetCode()
                    return respObj
                IPv6_NEIGHBOR = NEIGHBOR_IPv6
            else:
                IPv6 = IPv6_PREFIX = IPv6_NEIGHBOR = ""
        except:
            IPv6 = IPv6_PREFIX = IPv6_NEIGHBOR = ""

        if IP == "" and IPv6 == "":
            errorMsg = ("Neither SVI IPv4 nor IPv6 Address is configured, at least one must be configured")
            respObj.addErrorReport("VRF-LITE-EXT:INVALID_IP", errorMsg)
            respObj.setFailureRetCode()
            return respObj

        try:
            bgp_neighbor_list = BGP_NEIGHBOR_LIST
        except:
            bgp_neighbor_list = ""

        if bgp_neighbor_list != "":
            bgpRespObj = WrappersResp.getRespObj()
            bgpRespObj.setSuccessRetCode()
            bgpNeighborList = ast.literal_eval(bgp_neighbor_list)
            if IP != "" and PREFIX != "":
                subnetStr1 = Util.getSubnetStringWithPrefix(IP, PREFIX)
            if IPv6 != "" and IPv6_PREFIX != "":
                subnetV6Str1 = Util.getV6SubnetStringWithPrefix(IPv6, IPv6_PREFIX)
            for item in bgpNeighborList["BGP_NEIGHBOR_LIST"]:
                neighborAsn = item["NEIGHBOR_ASN"]
                respObj = Helper.isValidAsn(neighborAsn)
                if respObj.isRetCodeFailure():
                    errorMsg = ("Invalid Neighbor ASN %s" % (neighborAsn))
                    bgpRespObj.addErrorReport(("Error"), errorMsg)
                    bgpRespObj.setFailureRetCode()
                  
                neighborIp = item["NEIGHBOR_IP"]
                neighborIpV6 = item["NEIGHBOR_IPv6"]
                if neighborIp == "" and neighborIpV6 == "":
                    errorMsg = ("Neither BGP Peer IPv4 nor IPv6 Address is configured for the additional neighbor(s), "
                                "at least one must be configured for peering with a BGP neighbor")
                    bgpRespObj.addErrorReport("VRF-LITE-EXT:INVALID_IP", errorMsg)
                    bgpRespObj.setFailureRetCode()
                    continue

                if neighborIp != "":
                    if IP == "":
                        errorMsg = ("BGP Peer IPv4 address %s is configured but Local IPv4 address is not, "
                                    "Local IPv4 address must also be configured"% (neighborIp))
                        bgpRespObj.addErrorReport("VRF-LITE-EXT:INVALID_IP", errorMsg)
                        bgpRespObj.setFailureRetCode()
                    elif PREFIX != "":
                        subnetStr2 = Util.getSubnetStringWithPrefix(neighborIp, PREFIX)
                        if subnetStr1 != subnetStr2:
                            errorMsg = ("Local IP address %s and BGP Peer address %s are not in same subnet"%
                                        (IP_MASK, neighborIp))
                            bgpRespObj.addErrorReport("VRF-LITE-EXT:INVALID_IP", errorMsg)
                            bgpRespObj.setFailureRetCode()

                if neighborIpV6 != "":
                    if IPv6 == "":
                        errorMsg = ("BGP Peer IPv6 address %s is configured but Local IPv6 address is not, "
                                    "Local IPv6 address must also be configured"% (neighborIpV6))
                        bgpRespObj.addErrorReport("VRF-LITE-EXT:INVALID_IP", errorMsg)
                        bgpRespObj.setFailureRetCode()
                    elif IPv6_PREFIX != "":
                        subnetV6Str2 = Util.getV6SubnetStringWithPrefix(neighborIpV6, IPv6_PREFIX)
                        if subnetV6Str1 != subnetV6Str2:
                            errorMsg = ("Local IPv6 address %s and BGP Peer address %s are not in same subnet"%
                                        (IPv6_MASK, neighborIpV6))
                            bgpRespObj.addErrorReport("VRF-LITE-EXT:INVALID_IP", errorMsg)
                            bgpRespObj.setFailureRetCode()

            if bgpRespObj.isRetCodeFailure():
                return bgpRespObj
            
        # modify to be done, calling delete now to clean up PTIs before add
        Wrapper.print("Calling delete of PTIs for SOURCE %s" %(SOURCE))
        delete(True)

        #SVI Interface
        intfName = "Vlan" + str(DOT1Q_ID)
        
        # For default vrf in external fabric, should RM be done?
        rmIds = []
        
        intfVrf = "default" if VRF_NAME == "" or VRF_NAME.lower() == "default" else VRF_NAME
    
        # Create SVI for VRF Lite Extension
        new_nvPairs = {"INTF_NAME" : intfName,
                       "VLAN" : DOT1Q_ID,
                       "IP" : IP,
                       "PREFIX" : PREFIX,
                       "IPv6" : IPv6,
                       "PREFIXv6" : IPv6_PREFIX,
                       "INTF_VRF" : intfVrf,
                       "MTU" : MTU,
                       "ROUTING_TAG": ROUTING_TAG,
                       "DISABLE_IP_REDIRECTS": DISABLE_IP_REDIRECTS,
                       "ADMIN_STATE" : ADMIN_STATE,
                       "DESC" : DESC,
                       "CONF" : CONF
                      }
        Wrapper.print("adding SVI %s, nvPairs %s" % (new_nvPairs["INTF_NAME"], new_nvPairs))

        Util.exe(PTIWrapper.createOrUpdate(SERIAL_NUMBER, "INTERFACE", new_nvPairs["INTF_NAME"], SOURCE,
                                           ConfigPriority.CONFIG_PRIO_INTF, rmIds,
                                           "int_vlan_internal", new_nvPairs))
                                           
        #Create VRF (if non-default VRF) and eBGP peering for VRF Lite
        nvPairs = {"vrfName": intfVrf, "vrfDescription": "", "asn": fabricAsn,
                   "NEIGHBOR_ASN": NEIGHBOR_ASN, "bgpPassword" : eBGPPassword, "bgpPasswordKeyType"  : eBGPPasswordKeyType,
                   "NEIGHBOR_IP": IP_NEIGHBOR, "IPV6_NEIGHBOR": IPv6_NEIGHBOR}
        
        Wrapper.print("Creating External_VRF_Lite_eBGP with BGP neighbor %s, nvPairs %s" % (IP_NEIGHBOR, nvPairs))
        Util.exe(PTIWrapper.create(SERIAL_NUMBER, "SWITCH", "SWITCH", SOURCE,
                                   ConfigPriority.CONFIG_PRIO_UNDERLAY_TOP_LVL1 + 1,
                                   "External_VRF_Lite_eBGP", nvPairs))

        if config_local_as:
            if IP != "":
                nvPairs = {"BGP_AS": fabricAsn,
                           "VRF_NAME": intfVrf,
                           "BGP_NEIGHBOR_IP": IP_NEIGHBOR,
                           "BGP_LOCAL_AS": local_asn}
                Util.exe(PTIWrapper.create(SERIAL_NUMBER, "SWITCH", "SWITCH", SOURCE,
                                       ConfigPriority.CONFIG_PRIO_UNDERLAY_TOP_LVL1 + 2,
                                       "bgp_neighbor_ebgp_local_as",  nvPairs))

            if IPv6 != "":
                nvPairs = {"BGP_AS": fabricAsn,
                           "VRF_NAME": intfVrf,
                           "BGP_NEIGHBOR_IPV6": IPv6_NEIGHBOR,
                           "BGP_LOCAL_AS": local_asn}
                Util.exe(PTIWrapper.create(SERIAL_NUMBER, "SWITCH", "SWITCH", SOURCE,
                                       ConfigPriority.CONFIG_PRIO_UNDERLAY_TOP_LVL1 + 2,
                                       "bgp_neighbor_ebgp_v6_local_as",  nvPairs))

        try:
            if IP != "":
                if ROUTE_MAP_IN != "":
                    nvPairs = {"BGP_AS": fabricAsn,
                               "VRF_NAME": intfVrf,
                               "BGP_NEIGHBOR_IP": IP_NEIGHBOR,
                               "IP_TYPE": "ipv4",
                               "ROUTE_MAP": ROUTE_MAP_IN,
                               "DIRECTION": "in"}
                    Util.exe(PTIWrapper.create(SERIAL_NUMBER, "SWITCH", "SWITCH", SOURCE,
                                           ConfigPriority.CONFIG_PRIO_UNDERLAY_TOP_LVL1 + 2,
                                           "bgp_neighbor_route_map",  nvPairs))
                if ROUTE_MAP_OUT != "":
                    nvPairs = {"BGP_AS": fabricAsn,
                               "VRF_NAME": intfVrf,
                               "BGP_NEIGHBOR_IP": IP_NEIGHBOR,
                               "IP_TYPE": "ipv4",
                               "ROUTE_MAP": ROUTE_MAP_OUT,
                               "DIRECTION": "out"}
                    Util.exe(PTIWrapper.create(SERIAL_NUMBER, "SWITCH", "SWITCH", SOURCE,
                                           ConfigPriority.CONFIG_PRIO_UNDERLAY_TOP_LVL1 + 2,
                                           "bgp_neighbor_route_map",  nvPairs))

            if IPv6 != "":
                if IPV6_ROUTE_MAP_IN != "":
                    nvPairs = {"BGP_AS": fabricAsn,
                               "VRF_NAME": intfVrf,
                               "BGP_NEIGHBOR_IP": IPv6_NEIGHBOR,
                               "IP_TYPE": "ipv6",
                               "ROUTE_MAP": IPV6_ROUTE_MAP_IN,
                               "DIRECTION": "in"}
                    Util.exe(PTIWrapper.create(SERIAL_NUMBER, "SWITCH", "SWITCH", SOURCE,
                                           ConfigPriority.CONFIG_PRIO_UNDERLAY_TOP_LVL1 + 2,
                                           "bgp_neighbor_route_map",  nvPairs))
                if IPV6_ROUTE_MAP_OUT != "":
                    nvPairs = {"BGP_AS": fabricAsn,
                               "VRF_NAME": intfVrf,
                               "BGP_NEIGHBOR_IP": IPv6_NEIGHBOR,
                               "IP_TYPE": "ipv6",
                               "ROUTE_MAP": IPV6_ROUTE_MAP_OUT,
                               "DIRECTION": "out"}
                    Util.exe(PTIWrapper.create(SERIAL_NUMBER, "SWITCH", "SWITCH", SOURCE,
                                           ConfigPriority.CONFIG_PRIO_UNDERLAY_TOP_LVL1 + 2,
                                           "bgp_neighbor_route_map",  nvPairs))
        except:
            pass

        if bgp_neighbor_list != "":
            bgpNeighborList = ast.literal_eval(bgp_neighbor_list)
            for item in bgpNeighborList["BGP_NEIGHBOR_LIST"]:
                neighborIp = item["NEIGHBOR_IP"]
                neighborIpV6 = item["NEIGHBOR_IPv6"]
                neighborAsn = item["NEIGHBOR_ASN"]
                eBGPPassword = item["bgpPassword"]
                eBGPPasswordKeyType = item["bgpPasswordKeyType"]

                #Create VRF (if non-default VRF) and eBGP peering for VRF Lite
                nvPairs = {"vrfName": intfVrf, "vrfDescription": "", "asn": fabricAsn,
                           "NEIGHBOR_ASN": neighborAsn, "bgpPassword": eBGPPassword, "bgpPasswordKeyType": eBGPPasswordKeyType,
                           "NEIGHBOR_IP": neighborIp, "IPV6_NEIGHBOR": neighborIpV6}
        
                Wrapper.print("Creating External_VRF_Lite_eBGP with BGP neighbor %s, nvPairs %s" % (neighborIp, nvPairs))
                Util.exe(PTIWrapper.create(SERIAL_NUMBER, "SWITCH", "SWITCH", SOURCE,
                                           ConfigPriority.CONFIG_PRIO_UNDERLAY_TOP_LVL1 + 1,
                                           "External_VRF_Lite_eBGP", nvPairs))

                if config_local_as:
                    if neighborIp != "":
                        nvPairs = {"BGP_AS": fabricAsn,
                                   "VRF_NAME": intfVrf,
                                   "BGP_NEIGHBOR_IP": neighborIp,
                                   "BGP_LOCAL_AS": local_asn}
                        Util.exe(PTIWrapper.create(SERIAL_NUMBER, "SWITCH", "SWITCH", SOURCE,
                                                   ConfigPriority.CONFIG_PRIO_UNDERLAY_TOP_LVL1 + 2,
                                                   "bgp_neighbor_ebgp_local_as", nvPairs))
    
                    if neighborIpV6 != "":
                        nvPairs = {"BGP_AS": fabricAsn,
                                   "VRF_NAME": intfVrf,
                                   "BGP_NEIGHBOR_IPV6": neighborIpV6,
                                   "BGP_LOCAL_AS": local_asn}
                        Util.exe(PTIWrapper.create(SERIAL_NUMBER, "SWITCH", "SWITCH", SOURCE,
                                                   ConfigPriority.CONFIG_PRIO_UNDERLAY_TOP_LVL1 + 2,
                                                   "bgp_neighbor_ebgp_v6_local_as", nvPairs))
    
                try:
                    if neighborIp != "":
                        if ROUTE_MAP_IN != "":
                            nvPairs = {"BGP_AS": fabricAsn,
                                       "VRF_NAME": intfVrf,
                                       "BGP_NEIGHBOR_IP": neighborIp,
                                       "IP_TYPE": "ipv4",
                                       "ROUTE_MAP": ROUTE_MAP_IN,
                                       "DIRECTION": "in"}
                            Util.exe(PTIWrapper.create(SERIAL_NUMBER, "SWITCH", "SWITCH", SOURCE,
                                                       ConfigPriority.CONFIG_PRIO_UNDERLAY_TOP_LVL1 + 2,
                                                       "bgp_neighbor_route_map", nvPairs))
                        if ROUTE_MAP_OUT != "":
                            nvPairs = {"BGP_AS": fabricAsn,
                                       "VRF_NAME": intfVrf,
                                       "BGP_NEIGHBOR_IP": neighborIp,
                                       "IP_TYPE": "ipv4",
                                       "ROUTE_MAP": ROUTE_MAP_OUT,
                                       "DIRECTION": "out"}
                            Util.exe(PTIWrapper.create(SERIAL_NUMBER, "SWITCH", "SWITCH", SOURCE,
                                                       ConfigPriority.CONFIG_PRIO_UNDERLAY_TOP_LVL1 + 2,
                                                       "bgp_neighbor_route_map", nvPairs))
        
                    if neighborIpV6 != "":
                        if IPV6_ROUTE_MAP_IN != "":
                            nvPairs = {"BGP_AS": fabricAsn,
                                       "VRF_NAME": intfVrf,
                                       "BGP_NEIGHBOR_IP": neighborIpV6,
                                       "IP_TYPE": "ipv6",
                                       "ROUTE_MAP": IPV6_ROUTE_MAP_IN,
                                       "DIRECTION": "in"}
                            Util.exe(PTIWrapper.create(SERIAL_NUMBER, "SWITCH", "SWITCH", SOURCE,
                                                       ConfigPriority.CONFIG_PRIO_UNDERLAY_TOP_LVL1 + 2,
                                                       "bgp_neighbor_route_map", nvPairs))
                        if IPV6_ROUTE_MAP_OUT != "":
                            nvPairs = {"BGP_AS": fabricAsn,
                                       "VRF_NAME": intfVrf,
                                       "BGP_NEIGHBOR_IP": neighborIpV6,
                                       "IP_TYPE": "ipv6",
                                       "ROUTE_MAP": IPV6_ROUTE_MAP_OUT,
                                       "DIRECTION": "out"}
                            Util.exe(PTIWrapper.create(SERIAL_NUMBER, "SWITCH", "SWITCH", SOURCE,
                                                       ConfigPriority.CONFIG_PRIO_UNDERLAY_TOP_LVL1 + 2,
                                                       "bgp_neighbor_route_map", nvPairs))
                except:
                    pass

        respObj = WrappersResp.getRespObj()
        respObj.setSuccessRetCode()
        return respObj
    except respObjError as e:
        return e.value

def delete(Force=False):
    try:
        if Force:
            try:
                if IP_MASK != "":
                    IP, PREFIX = IP_MASK.split("/")
                    Wrapper.print("IPMask %s PREFIX %s" % (IP_MASK, PREFIX))

                    v4SkipDel = True if PREFIX and (int(PREFIX) < 8 or int(PREFIX) > 31) else False
                    if v4SkipDel:
                        Wrapper.print("Ext_VRF_Lite_Jython (intf %s): upgrade case, just skip so that template validation "
                                      "error reports warning and no change on ptis" % (src_intf))
                        respObj = WrappersResp.getRespObj()
                        respObj.setSuccessRetCode()
                        return respObj
            except:
                pass

        intfName = "Vlan" + str(DOT1Q_ID)

        #if called by topdown, SOURCE will be set. Then user will not be able to edit from View/Edit policies
        jython_source  = ""
        try:
            jython_source = SOURCE
        except:
            jython_source = POLICY_ID
            pass
        
        SOURCE = jython_source
        Wrapper.print("Source of VRF Lite Jython is %s for switch %s" %(SOURCE, SERIAL_NUMBER))
        
        #Get all PTIs for the SOURCE VRF Lite Extension
        wrappersResp = PTIWrapper.get(SERIAL_NUMBER, SOURCE)
        if wrappersResp.isRetCodeSuccess():
            vrfLitePolicies = wrappersResp.getValue()
            for policy in vrfLitePolicies:
                Wrapper.print("VRF Lite policy for External Fabric DELETE: %s %s" % (policy.getPolicyId(), policy.getTemplateName()))
                if policy.getTemplateName() == "int_vlan_internal":
                    PTIWrapper.deleteInstance(policy.getPolicyId())
                else:
                    PTIWrapper.markDeleteInstance(policy.getPolicyId())

        #Get the PTI for the SVI if any exists, this is to take care of the case where SVI is already created and then VRF Lite Jython is called
        srchOpt = CtrlPolicySearch()
        srchOpt.setSerialNumber(SERIAL_NUMBER)
        srchOpt.setEntityName(intfName)
        srchOpt.setTemplateContentType("PYTHON")
        SVIPolicies = Util.exe(PTIWrapper.getPTIs(srchOpt))
        if SVIPolicies:
            for policy in SVIPolicies:
                Wrapper.print("SVI policy for External Fabric DELETE: %s %s" % (policy.getPolicyId(), policy.getTemplateName()))
                PTIWrapper.deleteInstance(policy.getPolicyId())

        respObj = WrappersResp.getRespObj()
        respObj.setSuccessRetCode()
        return respObj
    except respObjError as e:
        return e.value

##

EOF
}