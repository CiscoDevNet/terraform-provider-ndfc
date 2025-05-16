// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package datasource_fabric

import (
	"strconv"
	. "terraform-provider-ndfc/internal/provider/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCFabricModel struct {
	FabricName                              string       `json:"FABRIC_NAME,omitempty"`
	AaaRemoteIpEnabled                      string       `json:"AAA_REMOTE_IP_ENABLED,omitempty"`
	AaaServerConf                           string       `json:"AAA_SERVER_CONF,omitempty"`
	AdvertisePipBgp                         string       `json:"ADVERTISE_PIP_BGP,omitempty"`
	AdvertisePipOnBorder                    string       `json:"ADVERTISE_PIP_ON_BORDER,omitempty"`
	AnycastBgwAdvertisePip                  string       `json:"ANYCAST_BGW_ADVERTISE_PIP,omitempty"`
	AnycastGwMac                            string       `json:"ANYCAST_GW_MAC,omitempty"`
	AnycastLbId                             *Int64Custom `json:"ANYCAST_LB_ID,omitempty"`
	AnycastRpIpRange                        string       `json:"ANYCAST_RP_IP_RANGE,omitempty"`
	AutoSymmetricDefaultVrf                 string       `json:"AUTO_SYMMETRIC_DEFAULT_VRF,omitempty"`
	AutoSymmetricVrfLite                    string       `json:"AUTO_SYMMETRIC_VRF_LITE,omitempty"`
	AutoUniqueVrfLiteIpPrefix               string       `json:"AUTO_UNIQUE_VRF_LITE_IP_PREFIX,omitempty"`
	AutoVrfliteIfcDefaultVrf                string       `json:"AUTO_VRFLITE_IFC_DEFAULT_VRF,omitempty"`
	Banner                                  string       `json:"BANNER,omitempty"`
	BfdAuthEnable                           string       `json:"BFD_AUTH_ENABLE,omitempty"`
	BfdAuthKey                              string       `json:"BFD_AUTH_KEY,omitempty"`
	BfdAuthKeyId                            *Int64Custom `json:"BFD_AUTH_KEY_ID,omitempty"`
	BfdEnable                               string       `json:"BFD_ENABLE,omitempty"`
	BfdIbgpEnable                           string       `json:"BFD_IBGP_ENABLE,omitempty"`
	BfdIsisEnable                           string       `json:"BFD_ISIS_ENABLE,omitempty"`
	BfdOspfEnable                           string       `json:"BFD_OSPF_ENABLE,omitempty"`
	BfdPimEnable                            string       `json:"BFD_PIM_ENABLE,omitempty"`
	BgpAs                                   string       `json:"BGP_AS,omitempty"`
	BgpAuthEnable                           string       `json:"BGP_AUTH_ENABLE,omitempty"`
	BgpAuthKey                              string       `json:"BGP_AUTH_KEY,omitempty"`
	BgpAuthKeyType                          *Int64Custom `json:"BGP_AUTH_KEY_TYPE,omitempty"`
	BgpLbId                                 *Int64Custom `json:"BGP_LB_ID,omitempty"`
	BootstrapConf                           string       `json:"BOOTSTRAP_CONF,omitempty"`
	BootstrapEnable                         string       `json:"BOOTSTRAP_ENABLE,omitempty"`
	BootstrapMultisubnet                    string       `json:"BOOTSTRAP_MULTISUBNET,omitempty"`
	BrownfieldNetworkNameFormat             string       `json:"BROWNFIELD_NETWORK_NAME_FORMAT,omitempty"`
	BrownfieldSkipOverlayNetworkAttachments string       `json:"BROWNFIELD_SKIP_OVERLAY_NETWORK_ATTACHMENTS,omitempty"`
	CdpEnable                               string       `json:"CDP_ENABLE,omitempty"`
	CoppPolicy                              string       `json:"COPP_POLICY,omitempty"`
	DciSubnetRange                          string       `json:"DCI_SUBNET_RANGE,omitempty"`
	DciSubnetTargetMask                     *int64       `json:"DCI_SUBNET_TARGET_MASK,omitempty"`
	DefaultQueuingPolicyCloudscale          string       `json:"DEAFULT_QUEUING_POLICY_CLOUDSCALE,omitempty"`
	DefaultQueuingPolicyOther               string       `json:"DEAFULT_QUEUING_POLICY_OTHER,omitempty"`
	DefaultQueuingPolicyRSeries             string       `json:"DEAFULT_QUEUING_POLICY_R_SERIES,omitempty"`
	DefaultVrfRedisBgpRmap                  string       `json:"DEFAULT_VRF_REDIS_BGP_RMAP,omitempty"`
	DhcpEnable                              string       `json:"DHCP_ENABLE,omitempty"`
	DhcpEnd                                 string       `json:"DHCP_END,omitempty"`
	DhcpIpv6Enable                          string       `json:"DHCP_IPV6_ENABLE,omitempty"`
	DhcpStart                               string       `json:"DHCP_START,omitempty"`
	DnsServerIpList                         string       `json:"DNS_SERVER_IP_LIST,omitempty"`
	DnsServerVrf                            string       `json:"DNS_SERVER_VRF,omitempty"`
	EnableAaa                               string       `json:"ENABLE_AAA,omitempty"`
	EnableDefaultQueuingPolicy              string       `json:"ENABLE_DEFAULT_QUEUING_POLICY,omitempty"`
	EnableFabricVpcDomainId                 string       `json:"ENABLE_FABRIC_VPC_DOMAIN_ID,omitempty"`
	EnableMacsec                            string       `json:"ENABLE_MACSEC,omitempty"`
	EnableNetflow                           string       `json:"ENABLE_NETFLOW,omitempty"`
	EnableNgoam                             string       `json:"ENABLE_NGOAM,omitempty"`
	EnableNxapi                             string       `json:"ENABLE_NXAPI,omitempty"`
	EnableNxapiHttp                         string       `json:"ENABLE_NXAPI_HTTP,omitempty"`
	EnablePbr                               string       `json:"ENABLE_PBR,omitempty"`
	EnablePvlan                             string       `json:"ENABLE_PVLAN,omitempty"`
	EnableTenantDhcp                        string       `json:"ENABLE_TENANT_DHCP,omitempty"`
	EnableTrm                               string       `json:"ENABLE_TRM,omitempty"`
	EnableVpcPeerLinkNativeVlan             string       `json:"ENABLE_VPC_PEER_LINK_NATIVE_VLAN,omitempty"`
	ExtraConfIntraLinks                     string       `json:"EXTRA_CONF_INTRA_LINKS,omitempty"`
	ExtraConfLeaf                           string       `json:"EXTRA_CONF_LEAF,omitempty"`
	ExtraConfSpine                          string       `json:"EXTRA_CONF_SPINE,omitempty"`
	ExtraConfTor                            string       `json:"EXTRA_CONF_TOR,omitempty"`
	FabricInterfaceType                     string       `json:"FABRIC_INTERFACE_TYPE,omitempty"`
	FabricMtu                               *Int64Custom `json:"FABRIC_MTU,omitempty"`
	FabricVpcDomainId                       *Int64Custom `json:"FABRIC_VPC_DOMAIN_ID,omitempty"`
	FabricVpcQos                            string       `json:"FABRIC_VPC_QOS,omitempty"`
	FabricVpcQosPolicyName                  string       `json:"FABRIC_VPC_QOS_POLICY_NAME,omitempty"`
	FeaturePtp                              string       `json:"FEATURE_PTP,omitempty"`
	GrfieldDebugFlag                        string       `json:"GRFIELD_DEBUG_FLAG,omitempty"`
	HdTime                                  *Int64Custom `json:"HD_TIME,omitempty"`
	HostIntfAdminState                      string       `json:"HOST_INTF_ADMIN_STATE,omitempty"`
	IbgpPeerTemplate                        string       `json:"IBGP_PEER_TEMPLATE,omitempty"`
	IbgpPeerTemplateLeaf                    string       `json:"IBGP_PEER_TEMPLATE_LEAF,omitempty"`
	InbandDhcpServers                       string       `json:"INBAND_DHCP_SERVERS,omitempty"`
	InbandMgmt                              string       `json:"INBAND_MGMT,omitempty"`
	IsisAuthEnable                          string       `json:"ISIS_AUTH_ENABLE,omitempty"`
	IsisAuthKey                             string       `json:"ISIS_AUTH_KEY,omitempty"`
	IsisAuthKeychainKeyId                   *Int64Custom `json:"ISIS_AUTH_KEYCHAIN_KEY_ID,omitempty"`
	IsisAuthKeychainName                    string       `json:"ISIS_AUTH_KEYCHAIN_NAME,omitempty"`
	IsisLevel                               string       `json:"ISIS_LEVEL,omitempty"`
	IsisOverloadElapseTime                  *Int64Custom `json:"ISIS_OVERLOAD_ELAPSE_TIME,omitempty"`
	IsisOverloadEnable                      string       `json:"ISIS_OVERLOAD_ENABLE,omitempty"`
	IsisP2pEnable                           string       `json:"ISIS_P2P_ENABLE,omitempty"`
	L2HostIntfMtu                           *Int64Custom `json:"L2_HOST_INTF_MTU,omitempty"`
	L2SegmentIdRange                        string       `json:"L2_SEGMENT_ID_RANGE,omitempty"`
	L3vniMcastGroup                         string       `json:"L3VNI_MCAST_GROUP,omitempty"`
	L3PartitionIdRange                      string       `json:"L3_PARTITION_ID_RANGE,omitempty"`
	LinkStateRouting                        string       `json:"LINK_STATE_ROUTING,omitempty"`
	LinkStateRoutingTag                     string       `json:"LINK_STATE_ROUTING_TAG,omitempty"`
	Loopback0Ipv6Range                      string       `json:"LOOPBACK0_IPV6_RANGE,omitempty"`
	Loopback0IpRange                        string       `json:"LOOPBACK0_IP_RANGE,omitempty"`
	Loopback1Ipv6Range                      string       `json:"LOOPBACK1_IPV6_RANGE,omitempty"`
	Loopback1IpRange                        string       `json:"LOOPBACK1_IP_RANGE,omitempty"`
	MacsecAlgorithm                         string       `json:"MACSEC_ALGORITHM,omitempty"`
	MacsecCipherSuite                       string       `json:"MACSEC_CIPHER_SUITE,omitempty"`
	MacsecFallbackAlgorithm                 string       `json:"MACSEC_FALLBACK_ALGORITHM,omitempty"`
	MacsecFallbackKeyString                 string       `json:"MACSEC_FALLBACK_KEY_STRING,omitempty"`
	MacsecKeyString                         string       `json:"MACSEC_KEY_STRING,omitempty"`
	MacsecReportTimer                       *Int64Custom `json:"MACSEC_REPORT_TIMER,omitempty"`
	MgmtGw                                  string       `json:"MGMT_GW,omitempty"`
	MgmtPrefix                              *Int64Custom `json:"MGMT_PREFIX,omitempty"`
	MgmtV6prefix                            *Int64Custom `json:"MGMT_V6PREFIX,omitempty"`
	MplsHandoff                             string       `json:"MPLS_HANDOFF,omitempty"`
	MplsLbId                                *Int64Custom `json:"MPLS_LB_ID,omitempty"`
	MplsLoopbackIpRange                     string       `json:"MPLS_LOOPBACK_IP_RANGE,omitempty"`
	MstInstanceRange                        string       `json:"MST_INSTANCE_RANGE,omitempty"`
	MulticastGroupSubnet                    string       `json:"MULTICAST_GROUP_SUBNET,omitempty"`
	NetflowExporterList                     string       `json:"NETFLOW_EXPORTER_LIST,omitempty"`
	NetflowMonitorList                      string       `json:"NETFLOW_MONITOR_LIST,omitempty"`
	NetflowRecordList                       string       `json:"NETFLOW_RECORD_LIST,omitempty"`
	NetworkVlanRange                        string       `json:"NETWORK_VLAN_RANGE,omitempty"`
	NtpServerIpList                         string       `json:"NTP_SERVER_IP_LIST,omitempty"`
	NtpServerVrf                            string       `json:"NTP_SERVER_VRF,omitempty"`
	NveLbId                                 *Int64Custom `json:"NVE_LB_ID,omitempty"`
	NxapiHttpsPort                          *Int64Custom `json:"NXAPI_HTTPS_PORT,omitempty"`
	NxapiHttpPort                           *Int64Custom `json:"NXAPI_HTTP_PORT,omitempty"`
	ObjectTrackingNumberRange               string       `json:"OBJECT_TRACKING_NUMBER_RANGE,omitempty"`
	OspfAreaId                              string       `json:"OSPF_AREA_ID,omitempty"`
	OspfAuthEnable                          string       `json:"OSPF_AUTH_ENABLE,omitempty"`
	OspfAuthKey                             string       `json:"OSPF_AUTH_KEY,omitempty"`
	OspfAuthKeyId                           *Int64Custom `json:"OSPF_AUTH_KEY_ID,omitempty"`
	OverlayMode                             string       `json:"OVERLAY_MODE,omitempty"`
	PerVrfLoopbackAutoProvision             string       `json:"PER_VRF_LOOPBACK_AUTO_PROVISION,omitempty"`
	PerVrfLoopbackIpRange                   string       `json:"PER_VRF_LOOPBACK_IP_RANGE,omitempty"`
	PhantomRpLbId1                          *Int64Custom `json:"PHANTOM_RP_LB_ID1,omitempty"`
	PhantomRpLbId2                          *Int64Custom `json:"PHANTOM_RP_LB_ID2,omitempty"`
	PhantomRpLbId3                          *Int64Custom `json:"PHANTOM_RP_LB_ID3,omitempty"`
	PhantomRpLbId4                          *Int64Custom `json:"PHANTOM_RP_LB_ID4,omitempty"`
	PimHelloAuthEnable                      string       `json:"PIM_HELLO_AUTH_ENABLE,omitempty"`
	PimHelloAuthKey                         string       `json:"PIM_HELLO_AUTH_KEY,omitempty"`
	PmEnable                                string       `json:"PM_ENABLE,omitempty"`
	PowerRedundancyMode                     string       `json:"POWER_REDUNDANCY_MODE,omitempty"`
	PtpDomainId                             *Int64Custom `json:"PTP_DOMAIN_ID,omitempty"`
	PtpLbId                                 *Int64Custom `json:"PTP_LB_ID,omitempty"`
	ReplicationMode                         string       `json:"REPLICATION_MODE,omitempty"`
	RouterIdRange                           string       `json:"ROUTER_ID_RANGE,omitempty"`
	RouteMapSequenceNumberRange             string       `json:"ROUTE_MAP_SEQUENCE_NUMBER_RANGE,omitempty"`
	RpCount                                 *Int64Custom `json:"RP_COUNT,omitempty"`
	RpLbId                                  *Int64Custom `json:"RP_LB_ID,omitempty"`
	RpMode                                  string       `json:"RP_MODE,omitempty"`
	RrCount                                 *Int64Custom `json:"RR_COUNT,omitempty"`
	SeedSwitchCoreInterfaces                string       `json:"SEED_SWITCH_CORE_INTERFACES,omitempty"`
	ServiceNetworkVlanRange                 string       `json:"SERVICE_NETWORK_VLAN_RANGE,omitempty"`
	SiteId                                  string       `json:"SITE_ID,omitempty"`
	SlaIdRange                              string       `json:"SLA_ID_RANGE,omitempty"`
	SnmpServerHostTrap                      string       `json:"SNMP_SERVER_HOST_TRAP,omitempty"`
	SpineSwitchCoreInterfaces               string       `json:"SPINE_SWITCH_CORE_INTERFACES,omitempty"`
	StaticUnderlayIpAlloc                   string       `json:"STATIC_UNDERLAY_IP_ALLOC,omitempty"`
	StpBridgePriority                       *Int64Custom `json:"STP_BRIDGE_PRIORITY,omitempty"`
	StpRootOption                           string       `json:"STP_ROOT_OPTION,omitempty"`
	StpVlanRange                            string       `json:"STP_VLAN_RANGE,omitempty"`
	StrictCcMode                            string       `json:"STRICT_CC_MODE,omitempty"`
	SubinterfaceRange                       string       `json:"SUBINTERFACE_RANGE,omitempty"`
	SubnetRange                             string       `json:"SUBNET_RANGE,omitempty"`
	SubnetTargetMask                        *Int64Custom `json:"SUBNET_TARGET_MASK,omitempty"`
	SyslogServerIpList                      string       `json:"SYSLOG_SERVER_IP_LIST,omitempty"`
	SyslogServerVrf                         string       `json:"SYSLOG_SERVER_VRF,omitempty"`
	SyslogSev                               string       `json:"SYSLOG_SEV,omitempty"`
	TcamAllocation                          string       `json:"TCAM_ALLOCATION,omitempty"`
	UnderlayIsV6                            string       `json:"UNDERLAY_IS_V6,omitempty"`
	UnnumBootstrapLbId                      *Int64Custom `json:"UNNUM_BOOTSTRAP_LB_ID,omitempty"`
	UnnumDhcpEnd                            string       `json:"UNNUM_DHCP_END,omitempty"`
	UnnumDhcpStart                          string       `json:"UNNUM_DHCP_START,omitempty"`
	UseLinkLocal                            string       `json:"USE_LINK_LOCAL,omitempty"`
	V6SubnetRange                           string       `json:"V6_SUBNET_RANGE,omitempty"`
	V6SubnetTargetMask                      *Int64Custom `json:"V6_SUBNET_TARGET_MASK,omitempty"`
	VpcAutoRecoveryTime                     *Int64Custom `json:"VPC_AUTO_RECOVERY_TIME,omitempty"`
	VpcDelayRestore                         *Int64Custom `json:"VPC_DELAY_RESTORE,omitempty"`
	VpcDomainIdRange                        string       `json:"VPC_DOMAIN_ID_RANGE,omitempty"`
	VpcEnableIpv6NdSync                     string       `json:"VPC_ENABLE_IPv6_ND_SYNC,omitempty"`
	VpcPeerKeepAliveOption                  string       `json:"VPC_PEER_KEEP_ALIVE_OPTION,omitempty"`
	VpcPeerLinkPo                           *Int64Custom `json:"VPC_PEER_LINK_PO,omitempty"`
	VpcPeerLinkVlan                         *Int64Custom `json:"VPC_PEER_LINK_VLAN,omitempty"`
	VrfLiteAutoconfig                       string       `json:"VRF_LITE_AUTOCONFIG,omitempty"`
	VrfVlanRange                            string       `json:"VRF_VLAN_RANGE,omitempty"`
	DefaultNetwork                          string       `json:"default_network,omitempty"`
	DefaultPvlanSecNetwork                  string       `json:"default_pvlan_sec_network,omitempty"`
	DefaultVrf                              string       `json:"default_vrf,omitempty"`
	EnableRealTimeBackup                    string       `json:"enableRealTimeBackup,omitempty"`
	EnableScheduledBackup                   string       `json:"enableScheduledBackup,omitempty"`
	NetworkExtensionTemplate                string       `json:"network_extension_template,omitempty"`
	ScheduledTime                           string       `json:"scheduledTime,omitempty"`
	VrfExtensionTemplate                    string       `json:"vrf_extension_template,omitempty"`
	DeploymentStatus                        string       `json:"deployment_status,omitempty"`
	Ipv6MulticastGroupSubnet                string       `json:"IPv6_MULTICAST_GROUP_SUBNET,omitempty"`
	EnableTrmv6                             string       `json:"ENABLE_TRMv6,omitempty"`
	L3vniIpv6McastGroup                     string       `json:"L3VNI_IPv6_MCAST_GROUP,omitempty"`
	MvpnVriIdRange                          string       `json:"MVPN_VRI_ID_RANGE,omitempty"`
	EnableVriIdRealloc                      string       `json:"ENABLE_VRI_ID_REALLOC,omitempty"`
	EnableAggAccIdRange                     string       `json:"ENABLE_AGG_ACC_ID_RANGE,omitempty"`
	AggAccVpcPoIdRange                      string       `json:"AGG_ACC_VPC_PO_ID_RANGE,omitempty"`
	IsisAreaNum                             string       `json:"ISIS_AREA_NUM,omitempty"`
	EnableSgt                               string       `json:"ENABLE_SGT,omitempty"`
	SgtNamePrefix                           string       `json:"SGT_NAME_PREFIX,omitempty"`
	SgtIdRange                              string       `json:"SGT_ID_RANGE,omitempty"`
	SgtPreprovision                         string       `json:"SGT_PREPROVISION,omitempty"`
	EnableDciMacsec                         string       `json:"ENABLE_DCI_MACSEC,omitempty"`
	EnableQkd                               string       `json:"ENABLE_QKD,omitempty"`
	AllowL3vniNoVlan                        string       `json:"ALLOW_L3VNI_NO_VLAN,omitempty"`
	EnableL3vniNoVlan                       string       `json:"ENABLE_L3VNI_NO_VLAN,omitempty"`
	MplsIsisAreaNum                         string       `json:"MPLS_ISIS_AREA_NUM,omitempty"`
	EnableAiMlQosPolicy                     string       `json:"ENABLE_AI_ML_QOS_POLICY,omitempty"`
	AiMlQosPolicy                           string       `json:"AI_ML_QOS_POLICY,omitempty"`
	PfcWatchInt                             *Int64Custom `json:"PFC_WATCH_INT,omitempty"`
	EnableRtIntfStats                       string       `json:"ENABLE_RT_INTF_STATS,omitempty"`
	Ipv6AnycastRpIpRange                    string       `json:"IPv6_ANYCAST_RP_IP_RANGE,omitempty"`
	KmeServerIp                             string       `json:"KME_SERVER_IP,omitempty"`
	KmeServerPort                           *Int64Custom `json:"KME_SERVER_PORT,omitempty"`
	TrustpointLabel                         string       `json:"TRUSTPOINT_LABEL,omitempty"`
	IgnoreCert                              string       `json:"IGNORE_CERT,omitempty"`
	DeploymentFreeze                        string       `json:"DEPLOYMENT_FREEZE,omitempty"`
	QkdProfileName                          string       `json:"QKD_PROFILE_NAME,omitempty"`
	PerVrfLoopbackAutoProvisionV6           string       `json:"PER_VRF_LOOPBACK_AUTO_PROVISION_V6,omitempty"`
	PerVrfLoopbackIpRangeV6                 string       `json:"PER_VRF_LOOPBACK_IP_RANGE_V6,omitempty"`
	EsrOption                               string       `json:"ESR_OPTION,omitempty"`
	PtpVlanId                               *Int64Custom `json:"PTP_VLAN_ID,omitempty"`
	AllowNxc                                string       `json:"ALLOW_NXC,omitempty"`
	OverwriteGlobalNxc                      string       `json:"OVERWRITE_GLOBAL_NXC,omitempty"`
	NxcDestVrf                              string       `json:"NXC_DEST_VRF,omitempty"`
	NxcSrcIntf                              string       `json:"NXC_SRC_INTF,omitempty"`
	NxcProxyServer                          string       `json:"NXC_PROXY_SERVER,omitempty"`
	NxcProxyPort                            *Int64Custom `json:"NXC_PROXY_PORT,omitempty"`
	VpcDelayRestoreTime                     *Int64Custom `json:"VPC_DELAY_RESTORE_TIME,omitempty"`
	FabricType                              string       `json:"FABRIC_TYPE,omitempty"`
	ExtFabricType                           string       `json:"EXT_FABRIC_TYPE,omitempty"`
	EnableAgent                             string       `json:"ENABLE_AGENT,omitempty"`
	AgentIntf                               string       `json:"AGENT_INTF,omitempty"`
	SspineAddDelDebugFlag                   string       `json:"SSPINE_ADD_DEL_DEBUG_FLAG,omitempty"`
	BrfieldDebugFlag                        string       `json:"BRFIELD_DEBUG_FLAG,omitempty"`
	ActiveMigration                         string       `json:"ACTIVE_MIGRATION,omitempty"`
	Ff                                      string       `json:"FF,omitempty"`
	BgpAsPrev                               string       `json:"BGP_AS_PREV,omitempty"`
	UnderlayIsV6Prev                        string       `json:"UNDERLAY_IS_V6_PREV,omitempty"`
	PmEnablePrev                            string       `json:"PM_ENABLE_PREV,omitempty"`
	EnableFabricVpcDomainIdPrev             string       `json:"ENABLE_FABRIC_VPC_DOMAIN_ID_PREV,omitempty"`
	OverlayModePrev                         string       `json:"OVERLAY_MODE_PREV,omitempty"`
	AllowL3vniNoVlanPrev                    string       `json:"ALLOW_L3VNI_NO_VLAN_PREV,omitempty"`
	EnablePvlanPrev                         string       `json:"ENABLE_PVLAN_PREV,omitempty"`
	AutoUniqueVrfLiteIpPrefixPrev           string       `json:"AUTO_UNIQUE_VRF_LITE_IP_PREFIX_PREV,omitempty"`
	PerVrfLoopbackAutoProvisionPrev         string       `json:"PER_VRF_LOOPBACK_AUTO_PROVISION_PREV,omitempty"`
	PerVrfLoopbackAutoProvisionV6Prev       string       `json:"PER_VRF_LOOPBACK_AUTO_PROVISION_V6_PREV,omitempty"`
	MsoSiteId                               string       `json:"MSO_SITE_ID,omitempty"`
	MsoControlerId                          string       `json:"MSO_CONTROLER_ID,omitempty"`
	MsoSiteGroupName                        string       `json:"MSO_SITE_GROUP_NAME,omitempty"`
	PremsoParentFabric                      string       `json:"PREMSO_PARENT_FABRIC,omitempty"`
	MsoConnectivityDeployed                 string       `json:"MSO_CONNECTIVITY_DEPLOYED,omitempty"`
	AnycastRpIpRangeInternal                string       `json:"ANYCAST_RP_IP_RANGE_INTERNAL,omitempty"`
	Ipv6AnycastRpIpRangeInternal            string       `json:"IPv6_ANYCAST_RP_IP_RANGE_INTERNAL,omitempty"`
	DhcpStartInternal                       string       `json:"DHCP_START_INTERNAL,omitempty"`
	DhcpEndInternal                         string       `json:"DHCP_END_INTERNAL,omitempty"`
	MgmtGwInternal                          string       `json:"MGMT_GW_INTERNAL,omitempty"`
	MgmtPrefixInternal                      *Int64Custom `json:"MGMT_PREFIX_INTERNAL,omitempty"`
	BootstrapMultisubnetInternal            string       `json:"BOOTSTRAP_MULTISUBNET_INTERNAL,omitempty"`
	MgmtV6prefixInternal                    *Int64Custom `json:"MGMT_V6PREFIX_INTERNAL,omitempty"`
	DhcpIpv6EnableInternal                  string       `json:"DHCP_IPV6_ENABLE_INTERNAL,omitempty"`
	UnnumDhcpStartInternal                  string       `json:"UNNUM_DHCP_START_INTERNAL,omitempty"`
	UnnumDhcpEndInternal                    string       `json:"UNNUM_DHCP_END_INTERNAL,omitempty"`
	EnableEvpn                              string       `json:"ENABLE_EVPN,omitempty"`
	FeaturePtpInternal                      string       `json:"FEATURE_PTP_INTERNAL,omitempty"`
	SspineCount                             *Int64Custom `json:"SSPINE_COUNT,omitempty"`
	SpineCount                              *Int64Custom `json:"SPINE_COUNT,omitempty"`
	AbstractFeatureLeaf                     string       `json:"abstract_feature_leaf,omitempty"`
	AbstractFeatureSpine                    string       `json:"abstract_feature_spine,omitempty"`
	AbstractDhcp                            string       `json:"abstract_dhcp,omitempty"`
	AbstractMulticast                       string       `json:"abstract_multicast,omitempty"`
	AbstractAnycastRp                       string       `json:"abstract_anycast_rp,omitempty"`
	AbstractLoopbackInterface               string       `json:"abstract_loopback_interface,omitempty"`
	AbstractIsis                            string       `json:"abstract_isis,omitempty"`
	AbstractOspf                            string       `json:"abstract_ospf,omitempty"`
	AbstractVpcDomain                       string       `json:"abstract_vpc_domain,omitempty"`
	AbstractVlanInterface                   string       `json:"abstract_vlan_interface,omitempty"`
	AbstractIsisInterface                   string       `json:"abstract_isis_interface,omitempty"`
	AbstractOspfInterface                   string       `json:"abstract_ospf_interface,omitempty"`
	AbstractPimInterface                    string       `json:"abstract_pim_interface,omitempty"`
	AbstractRouteMap                        string       `json:"abstract_route_map,omitempty"`
	AbstractBgp                             string       `json:"abstract_bgp,omitempty"`
	AbstractBgpRr                           string       `json:"abstract_bgp_rr,omitempty"`
	AbstractBgpNeighbor                     string       `json:"abstract_bgp_neighbor,omitempty"`
	AbstractExtraConfigLeaf                 string       `json:"abstract_extra_config_leaf,omitempty"`
	AbstractExtraConfigSpine                string       `json:"abstract_extra_config_spine,omitempty"`
	AbstractExtraConfigTor                  string       `json:"abstract_extra_config_tor,omitempty"`
	AbstractExtraConfigBootstrap            string       `json:"abstract_extra_config_bootstrap,omitempty"`
	TempAnycastGateway                      string       `json:"temp_anycast_gateway,omitempty"`
	TempVpcDomainMgmt                       string       `json:"temp_vpc_domain_mgmt,omitempty"`
	TempVpcPeerLink                         string       `json:"temp_vpc_peer_link,omitempty"`
	AbstractRoutedHost                      string       `json:"abstract_routed_host,omitempty"`
	UpgradeFromVersion                      string       `json:"UPGRADE_FROM_VERSION,omitempty"`
	TopdownConfigRmTracking                 string       `json:"TOPDOWN_CONFIG_RM_TRACKING,omitempty"`
	SiteIdPolicyId                          *Int64Custom `json:"SITE_ID_POLICY_ID,omitempty"`
	FabricVpcDomainIdPrev                   *Int64Custom `json:"FABRIC_VPC_DOMAIN_ID_PREV,omitempty"`
	LinkStateRoutingTagPrev                 string       `json:"LINK_STATE_ROUTING_TAG_PREV,omitempty"`
	BfdEnablePrev                           string       `json:"BFD_ENABLE_PREV,omitempty"`
	EnableSgtPrev                           string       `json:"ENABLE_SGT_PREV,omitempty"`
	SgtPreprovisionPrev                     string       `json:"SGT_PREPROVISION_PREV,omitempty"`
	SgtPreprovRecalcStatus                  string       `json:"SGT_PREPROV_RECALC_STATUS,omitempty"`
	SgtRecalcStatus                         string       `json:"SGT_RECALC_STATUS,omitempty"`
	SgtOperStatus                           string       `json:"SGT_OPER_STATUS,omitempty"`
	EnableMacsecPrev                        string       `json:"ENABLE_MACSEC_PREV,omitempty"`
	EnableDciMacsecPrev                     string       `json:"ENABLE_DCI_MACSEC_PREV,omitempty"`
	DciMacsecFallbackKeyString              string       `json:"DCI_MACSEC_FALLBACK_KEY_STRING,omitempty"`
	DciMacsecFallbackAlgorithm              string       `json:"DCI_MACSEC_FALLBACK_ALGORITHM,omitempty"`
	DciMacsecAlgorithm                      string       `json:"DCI_MACSEC_ALGORITHM,omitempty"`
	DciMacsecKeyString                      string       `json:"DCI_MACSEC_KEY_STRING,omitempty"`
	DciMacsecCipherSuite                    string       `json:"DCI_MACSEC_CIPHER_SUITE,omitempty"`
	QkdProfileNamePrev                      string       `json:"QKD_PROFILE_NAME_PREV,omitempty"`
	FabricMtuPrev                           *Int64Custom `json:"FABRIC_MTU_PREV,omitempty"`
	L2HostIntfMtuPrev                       *Int64Custom `json:"L2_HOST_INTF_MTU_PREV,omitempty"`
	MplsIsisAreaNumPrev                     string       `json:"MPLS_ISIS_AREA_NUM_PREV,omitempty"`
	IsisAreaNumPrev                         string       `json:"ISIS_AREA_NUM_PREV,omitempty"`
	EnableAiMlQosPolicyFlap                 string       `json:"ENABLE_AI_ML_QOS_POLICY_FLAP,omitempty"`
	PfcWatchIntPrev                         *Int64Custom `json:"PFC_WATCH_INT_PREV,omitempty"`
	InbandMgmtPrev                          string       `json:"INBAND_MGMT_PREV,omitempty"`
	BootstrapEnablePrev                     string       `json:"BOOTSTRAP_ENABLE_PREV,omitempty"`
	EnableNetflowPrev                       string       `json:"ENABLE_NETFLOW_PREV,omitempty"`
	AllowNxcPrev                            string       `json:"ALLOW_NXC_PREV,omitempty"`
	EnableNbmPassivePrev                    string       `json:"ENABLE_NBM_PASSIVE_PREV,omitempty"`
	FabricTechnology                        string       `json:"FABRIC_TECHNOLOGY,omitempty"`
	InterfaceEthernetDefaultPolicy          string       `json:"INTERFACE_ETHERNET_DEFAULT_POLICY,omitempty"`
	InterfaceLoopbackDefaultPolicy          string       `json:"INTERFACE_LOOPBACK_DEFAULT_POLICY,omitempty"`
	InterfacePortChannelDefaultPolicy       string       `json:"INTERFACE_PORT_CHANNEL_DEFAULT_POLICY,omitempty"`
	InterfaceVlanDefaultPolicy              string       `json:"INTERFACE_VLAN_DEFAULT_POLICY,omitempty"`
	RpIpRangeInternal                       string       `json:"RP_IP_RANGE_INTERNAL,omitempty"`
	InbandEnablePrev                        string       `json:"INBAND_ENABLE_PREV,omitempty"`
	EnableAsm                               string       `json:"ENABLE_ASM,omitempty"`
	DomainNameInternal                      string       `json:"DOMAIN_NAME_INTERNAL,omitempty"`
	PnpEnableInternal                       string       `json:"PNP_ENABLE_INTERNAL,omitempty"`
	BgwRoutingTag                           *Int64Custom `json:"BGW_ROUTING_TAG,omitempty"`
	DcnmId                                  string       `json:"DCNM_ID,omitempty"`
	EnableTrmTrmv6                          string       `json:"ENABLE_TRM_TRMv6,omitempty"`
	EnableTrmTrmv6Prev                      string       `json:"ENABLE_TRM_TRMv6_PREV,omitempty"`
	Loopback100Ipv6Range                    string       `json:"LOOPBACK100_IPV6_RANGE,omitempty"`
	BgwRoutingTagPrev                       string       `json:"BGW_ROUTING_TAG_PREV,omitempty"`
	MsIfcBgpAuthKeyType                     *Int64Custom `json:"MS_IFC_BGP_AUTH_KEY_TYPE,omitempty"`
	MsIfcBgpAuthKeyTypePrev                 *Int64Custom `json:"MS_IFC_BGP_AUTH_KEY_TYPE_PREV,omitempty"`
	MsIfcBgpPasswordEnablePrev              string       `json:"MS_IFC_BGP_PASSWORD_ENABLE_PREV,omitempty"`
	MsIfcBgpPasswordPrev                    string       `json:"MS_IFC_BGP_PASSWORD_PREV,omitempty"`
	ParentOnemanageFabric                   string       `json:"PARENT_ONEMANAGE_FABRIC,omitempty"`
	SgtIdRangePrev                          string       `json:"SGT_ID_RANGE_PREV,omitempty"`
	SgtNamePrefixPrev                       string       `json:"SGT_NAME_PREFIX_PREV,omitempty"`
	V6DciSubnetRange                        string       `json:"V6_DCI_SUBNET_RANGE,omitempty"`
	V6DciSubnetTargetMask                   *Int64Custom `json:"V6_DCI_SUBNET_TARGET_MASK,omitempty"`
	VxlanUnderlayIsV6                       string       `json:"VXLAN_UNDERLAY_IS_V6,omitempty"`
}

func (v *FabricModel) SetModelData(jsonData *NDFCFabricModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.FabricName != "" {
		v.FabricName = types.StringValue(jsonData.FabricName)
	} else {
		v.FabricName = types.StringNull()
	}

	if jsonData.AaaRemoteIpEnabled != "" {
		x, _ := strconv.ParseBool(jsonData.AaaRemoteIpEnabled)
		v.AaaRemoteIpEnabled = types.BoolValue(x)
	} else {
		v.AaaRemoteIpEnabled = types.BoolNull()
	}

	if jsonData.AaaServerConf != "" {
		v.AaaServerConf = types.StringValue(jsonData.AaaServerConf)
	} else {
		v.AaaServerConf = types.StringNull()
	}

	if jsonData.AdvertisePipBgp != "" {
		x, _ := strconv.ParseBool(jsonData.AdvertisePipBgp)
		v.AdvertisePipBgp = types.BoolValue(x)
	} else {
		v.AdvertisePipBgp = types.BoolNull()
	}

	if jsonData.AdvertisePipOnBorder != "" {
		x, _ := strconv.ParseBool(jsonData.AdvertisePipOnBorder)
		v.AdvertisePipOnBorder = types.BoolValue(x)
	} else {
		v.AdvertisePipOnBorder = types.BoolNull()
	}

	if jsonData.AnycastBgwAdvertisePip != "" {
		x, _ := strconv.ParseBool(jsonData.AnycastBgwAdvertisePip)
		v.AnycastBgwAdvertisePip = types.BoolValue(x)
	} else {
		v.AnycastBgwAdvertisePip = types.BoolNull()
	}

	if jsonData.AnycastGwMac != "" {
		v.AnycastGwMac = types.StringValue(jsonData.AnycastGwMac)
	} else {
		v.AnycastGwMac = types.StringNull()
	}

	if jsonData.AnycastLbId != nil {
		if jsonData.AnycastLbId.IsEmpty() {
			v.AnycastLbId = types.Int64Null()
		} else {
			v.AnycastLbId = types.Int64Value(int64(*jsonData.AnycastLbId))
		}
	} else {
		v.AnycastLbId = types.Int64Null()
	}

	if jsonData.AnycastRpIpRange != "" {
		v.AnycastRpIpRange = types.StringValue(jsonData.AnycastRpIpRange)
	} else {
		v.AnycastRpIpRange = types.StringNull()
	}

	if jsonData.AutoSymmetricDefaultVrf != "" {
		x, _ := strconv.ParseBool(jsonData.AutoSymmetricDefaultVrf)
		v.AutoSymmetricDefaultVrf = types.BoolValue(x)
	} else {
		v.AutoSymmetricDefaultVrf = types.BoolNull()
	}

	if jsonData.AutoSymmetricVrfLite != "" {
		x, _ := strconv.ParseBool(jsonData.AutoSymmetricVrfLite)
		v.AutoSymmetricVrfLite = types.BoolValue(x)
	} else {
		v.AutoSymmetricVrfLite = types.BoolNull()
	}

	if jsonData.AutoUniqueVrfLiteIpPrefix != "" {
		x, _ := strconv.ParseBool(jsonData.AutoUniqueVrfLiteIpPrefix)
		v.AutoUniqueVrfLiteIpPrefix = types.BoolValue(x)
	} else {
		v.AutoUniqueVrfLiteIpPrefix = types.BoolNull()
	}

	if jsonData.AutoVrfliteIfcDefaultVrf != "" {
		x, _ := strconv.ParseBool(jsonData.AutoVrfliteIfcDefaultVrf)
		v.AutoVrfliteIfcDefaultVrf = types.BoolValue(x)
	} else {
		v.AutoVrfliteIfcDefaultVrf = types.BoolNull()
	}

	if jsonData.Banner != "" {
		v.Banner = types.StringValue(jsonData.Banner)
	} else {
		v.Banner = types.StringNull()
	}

	if jsonData.BfdAuthEnable != "" {
		x, _ := strconv.ParseBool(jsonData.BfdAuthEnable)
		v.BfdAuthEnable = types.BoolValue(x)
	} else {
		v.BfdAuthEnable = types.BoolNull()
	}

	if jsonData.BfdAuthKey != "" {
		v.BfdAuthKey = types.StringValue(jsonData.BfdAuthKey)
	} else {
		v.BfdAuthKey = types.StringNull()
	}

	if jsonData.BfdAuthKeyId != nil {
		if jsonData.BfdAuthKeyId.IsEmpty() {
			v.BfdAuthKeyId = types.Int64Null()
		} else {
			v.BfdAuthKeyId = types.Int64Value(int64(*jsonData.BfdAuthKeyId))
		}
	} else {
		v.BfdAuthKeyId = types.Int64Null()
	}

	if jsonData.BfdEnable != "" {
		x, _ := strconv.ParseBool(jsonData.BfdEnable)
		v.BfdEnable = types.BoolValue(x)
	} else {
		v.BfdEnable = types.BoolNull()
	}

	if jsonData.BfdIbgpEnable != "" {
		x, _ := strconv.ParseBool(jsonData.BfdIbgpEnable)
		v.BfdIbgpEnable = types.BoolValue(x)
	} else {
		v.BfdIbgpEnable = types.BoolNull()
	}

	if jsonData.BfdIsisEnable != "" {
		x, _ := strconv.ParseBool(jsonData.BfdIsisEnable)
		v.BfdIsisEnable = types.BoolValue(x)
	} else {
		v.BfdIsisEnable = types.BoolNull()
	}

	if jsonData.BfdOspfEnable != "" {
		x, _ := strconv.ParseBool(jsonData.BfdOspfEnable)
		v.BfdOspfEnable = types.BoolValue(x)
	} else {
		v.BfdOspfEnable = types.BoolNull()
	}

	if jsonData.BfdPimEnable != "" {
		x, _ := strconv.ParseBool(jsonData.BfdPimEnable)
		v.BfdPimEnable = types.BoolValue(x)
	} else {
		v.BfdPimEnable = types.BoolNull()
	}

	if jsonData.BgpAs != "" {
		v.BgpAs = types.StringValue(jsonData.BgpAs)
	} else {
		v.BgpAs = types.StringNull()
	}

	if jsonData.BgpAuthEnable != "" {
		x, _ := strconv.ParseBool(jsonData.BgpAuthEnable)
		v.BgpAuthEnable = types.BoolValue(x)
	} else {
		v.BgpAuthEnable = types.BoolNull()
	}

	if jsonData.BgpAuthKey != "" {
		v.BgpAuthKey = types.StringValue(jsonData.BgpAuthKey)
	} else {
		v.BgpAuthKey = types.StringNull()
	}

	if jsonData.BgpAuthKeyType != nil {
		if jsonData.BgpAuthKeyType.IsEmpty() {
			v.BgpAuthKeyType = types.Int64Null()
		} else {
			v.BgpAuthKeyType = types.Int64Value(int64(*jsonData.BgpAuthKeyType))
		}
	} else {
		v.BgpAuthKeyType = types.Int64Null()
	}

	if jsonData.BgpLbId != nil {
		if jsonData.BgpLbId.IsEmpty() {
			v.BgpLbId = types.Int64Null()
		} else {
			v.BgpLbId = types.Int64Value(int64(*jsonData.BgpLbId))
		}
	} else {
		v.BgpLbId = types.Int64Null()
	}

	if jsonData.BootstrapConf != "" {
		v.BootstrapConf = types.StringValue(jsonData.BootstrapConf)
	} else {
		v.BootstrapConf = types.StringNull()
	}

	if jsonData.BootstrapEnable != "" {
		x, _ := strconv.ParseBool(jsonData.BootstrapEnable)
		v.BootstrapEnable = types.BoolValue(x)
	} else {
		v.BootstrapEnable = types.BoolNull()
	}

	if jsonData.BootstrapMultisubnet != "" {
		v.BootstrapMultisubnet = types.StringValue(jsonData.BootstrapMultisubnet)
	} else {
		v.BootstrapMultisubnet = types.StringNull()
	}

	if jsonData.BrownfieldNetworkNameFormat != "" {
		v.BrownfieldNetworkNameFormat = types.StringValue(jsonData.BrownfieldNetworkNameFormat)
	} else {
		v.BrownfieldNetworkNameFormat = types.StringNull()
	}

	if jsonData.BrownfieldSkipOverlayNetworkAttachments != "" {
		x, _ := strconv.ParseBool(jsonData.BrownfieldSkipOverlayNetworkAttachments)
		v.BrownfieldSkipOverlayNetworkAttachments = types.BoolValue(x)
	} else {
		v.BrownfieldSkipOverlayNetworkAttachments = types.BoolNull()
	}

	if jsonData.CdpEnable != "" {
		x, _ := strconv.ParseBool(jsonData.CdpEnable)
		v.CdpEnable = types.BoolValue(x)
	} else {
		v.CdpEnable = types.BoolNull()
	}

	if jsonData.CoppPolicy != "" {
		v.CoppPolicy = types.StringValue(jsonData.CoppPolicy)
	} else {
		v.CoppPolicy = types.StringNull()
	}

	if jsonData.DciSubnetRange != "" {
		v.DciSubnetRange = types.StringValue(jsonData.DciSubnetRange)
	} else {
		v.DciSubnetRange = types.StringNull()
	}

	if jsonData.DciSubnetTargetMask != nil {
		v.DciSubnetTargetMask = types.Int64Value(*jsonData.DciSubnetTargetMask)

	} else {
		v.DciSubnetTargetMask = types.Int64Null()
	}

	if jsonData.DefaultQueuingPolicyCloudscale != "" {
		v.DefaultQueuingPolicyCloudscale = types.StringValue(jsonData.DefaultQueuingPolicyCloudscale)
	} else {
		v.DefaultQueuingPolicyCloudscale = types.StringNull()
	}

	if jsonData.DefaultQueuingPolicyOther != "" {
		v.DefaultQueuingPolicyOther = types.StringValue(jsonData.DefaultQueuingPolicyOther)
	} else {
		v.DefaultQueuingPolicyOther = types.StringNull()
	}

	if jsonData.DefaultQueuingPolicyRSeries != "" {
		v.DefaultQueuingPolicyRSeries = types.StringValue(jsonData.DefaultQueuingPolicyRSeries)
	} else {
		v.DefaultQueuingPolicyRSeries = types.StringNull()
	}

	if jsonData.DefaultVrfRedisBgpRmap != "" {
		v.DefaultVrfRedisBgpRmap = types.StringValue(jsonData.DefaultVrfRedisBgpRmap)
	} else {
		v.DefaultVrfRedisBgpRmap = types.StringNull()
	}

	if jsonData.DhcpEnable != "" {
		x, _ := strconv.ParseBool(jsonData.DhcpEnable)
		v.DhcpEnable = types.BoolValue(x)
	} else {
		v.DhcpEnable = types.BoolNull()
	}

	if jsonData.DhcpEnd != "" {
		v.DhcpEnd = types.StringValue(jsonData.DhcpEnd)
	} else {
		v.DhcpEnd = types.StringNull()
	}

	if jsonData.DhcpIpv6Enable != "" {
		v.DhcpIpv6Enable = types.StringValue(jsonData.DhcpIpv6Enable)
	} else {
		v.DhcpIpv6Enable = types.StringNull()
	}

	if jsonData.DhcpStart != "" {
		v.DhcpStart = types.StringValue(jsonData.DhcpStart)
	} else {
		v.DhcpStart = types.StringNull()
	}

	if jsonData.DnsServerIpList != "" {
		v.DnsServerIpList = types.StringValue(jsonData.DnsServerIpList)
	} else {
		v.DnsServerIpList = types.StringNull()
	}

	if jsonData.DnsServerVrf != "" {
		v.DnsServerVrf = types.StringValue(jsonData.DnsServerVrf)
	} else {
		v.DnsServerVrf = types.StringNull()
	}

	if jsonData.EnableAaa != "" {
		x, _ := strconv.ParseBool(jsonData.EnableAaa)
		v.EnableAaa = types.BoolValue(x)
	} else {
		v.EnableAaa = types.BoolNull()
	}

	if jsonData.EnableDefaultQueuingPolicy != "" {
		x, _ := strconv.ParseBool(jsonData.EnableDefaultQueuingPolicy)
		v.EnableDefaultQueuingPolicy = types.BoolValue(x)
	} else {
		v.EnableDefaultQueuingPolicy = types.BoolNull()
	}

	if jsonData.EnableFabricVpcDomainId != "" {
		x, _ := strconv.ParseBool(jsonData.EnableFabricVpcDomainId)
		v.EnableFabricVpcDomainId = types.BoolValue(x)
	} else {
		v.EnableFabricVpcDomainId = types.BoolNull()
	}

	if jsonData.EnableMacsec != "" {
		x, _ := strconv.ParseBool(jsonData.EnableMacsec)
		v.EnableMacsec = types.BoolValue(x)
	} else {
		v.EnableMacsec = types.BoolNull()
	}

	if jsonData.EnableNetflow != "" {
		x, _ := strconv.ParseBool(jsonData.EnableNetflow)
		v.EnableNetflow = types.BoolValue(x)
	} else {
		v.EnableNetflow = types.BoolNull()
	}

	if jsonData.EnableNgoam != "" {
		x, _ := strconv.ParseBool(jsonData.EnableNgoam)
		v.EnableNgoam = types.BoolValue(x)
	} else {
		v.EnableNgoam = types.BoolNull()
	}

	if jsonData.EnableNxapi != "" {
		x, _ := strconv.ParseBool(jsonData.EnableNxapi)
		v.EnableNxapi = types.BoolValue(x)
	} else {
		v.EnableNxapi = types.BoolNull()
	}

	if jsonData.EnableNxapiHttp != "" {
		x, _ := strconv.ParseBool(jsonData.EnableNxapiHttp)
		v.EnableNxapiHttp = types.BoolValue(x)
	} else {
		v.EnableNxapiHttp = types.BoolNull()
	}

	if jsonData.EnablePbr != "" {
		x, _ := strconv.ParseBool(jsonData.EnablePbr)
		v.EnablePbr = types.BoolValue(x)
	} else {
		v.EnablePbr = types.BoolNull()
	}

	if jsonData.EnablePvlan != "" {
		x, _ := strconv.ParseBool(jsonData.EnablePvlan)
		v.EnablePvlan = types.BoolValue(x)
	} else {
		v.EnablePvlan = types.BoolNull()
	}

	if jsonData.EnableTenantDhcp != "" {
		x, _ := strconv.ParseBool(jsonData.EnableTenantDhcp)
		v.EnableTenantDhcp = types.BoolValue(x)
	} else {
		v.EnableTenantDhcp = types.BoolNull()
	}

	if jsonData.EnableTrm != "" {
		x, _ := strconv.ParseBool(jsonData.EnableTrm)
		v.EnableTrm = types.BoolValue(x)
	} else {
		v.EnableTrm = types.BoolNull()
	}

	if jsonData.EnableVpcPeerLinkNativeVlan != "" {
		x, _ := strconv.ParseBool(jsonData.EnableVpcPeerLinkNativeVlan)
		v.EnableVpcPeerLinkNativeVlan = types.BoolValue(x)
	} else {
		v.EnableVpcPeerLinkNativeVlan = types.BoolNull()
	}

	if jsonData.ExtraConfIntraLinks != "" {
		v.ExtraConfIntraLinks = types.StringValue(jsonData.ExtraConfIntraLinks)
	} else {
		v.ExtraConfIntraLinks = types.StringNull()
	}

	if jsonData.ExtraConfLeaf != "" {
		v.ExtraConfLeaf = types.StringValue(jsonData.ExtraConfLeaf)
	} else {
		v.ExtraConfLeaf = types.StringNull()
	}

	if jsonData.ExtraConfSpine != "" {
		v.ExtraConfSpine = types.StringValue(jsonData.ExtraConfSpine)
	} else {
		v.ExtraConfSpine = types.StringNull()
	}

	if jsonData.ExtraConfTor != "" {
		v.ExtraConfTor = types.StringValue(jsonData.ExtraConfTor)
	} else {
		v.ExtraConfTor = types.StringNull()
	}

	if jsonData.FabricInterfaceType != "" {
		v.FabricInterfaceType = types.StringValue(jsonData.FabricInterfaceType)
	} else {
		v.FabricInterfaceType = types.StringNull()
	}

	if jsonData.FabricMtu != nil {
		if jsonData.FabricMtu.IsEmpty() {
			v.FabricMtu = types.Int64Null()
		} else {
			v.FabricMtu = types.Int64Value(int64(*jsonData.FabricMtu))
		}
	} else {
		v.FabricMtu = types.Int64Null()
	}

	if jsonData.FabricVpcDomainId != nil {
		if jsonData.FabricVpcDomainId.IsEmpty() {
			v.FabricVpcDomainId = types.Int64Null()
		} else {
			v.FabricVpcDomainId = types.Int64Value(int64(*jsonData.FabricVpcDomainId))
		}
	} else {
		v.FabricVpcDomainId = types.Int64Null()
	}

	if jsonData.FabricVpcQos != "" {
		x, _ := strconv.ParseBool(jsonData.FabricVpcQos)
		v.FabricVpcQos = types.BoolValue(x)
	} else {
		v.FabricVpcQos = types.BoolNull()
	}

	if jsonData.FabricVpcQosPolicyName != "" {
		v.FabricVpcQosPolicyName = types.StringValue(jsonData.FabricVpcQosPolicyName)
	} else {
		v.FabricVpcQosPolicyName = types.StringNull()
	}

	if jsonData.FeaturePtp != "" {
		x, _ := strconv.ParseBool(jsonData.FeaturePtp)
		v.FeaturePtp = types.BoolValue(x)
	} else {
		v.FeaturePtp = types.BoolNull()
	}

	if jsonData.GrfieldDebugFlag != "" {
		v.GrfieldDebugFlag = types.StringValue(jsonData.GrfieldDebugFlag)
	} else {
		v.GrfieldDebugFlag = types.StringNull()
	}

	if jsonData.HdTime != nil {
		if jsonData.HdTime.IsEmpty() {
			v.HdTime = types.Int64Null()
		} else {
			v.HdTime = types.Int64Value(int64(*jsonData.HdTime))
		}
	} else {
		v.HdTime = types.Int64Null()
	}

	if jsonData.HostIntfAdminState != "" {
		x, _ := strconv.ParseBool(jsonData.HostIntfAdminState)
		v.HostIntfAdminState = types.BoolValue(x)
	} else {
		v.HostIntfAdminState = types.BoolNull()
	}

	if jsonData.IbgpPeerTemplate != "" {
		v.IbgpPeerTemplate = types.StringValue(jsonData.IbgpPeerTemplate)
	} else {
		v.IbgpPeerTemplate = types.StringNull()
	}

	if jsonData.IbgpPeerTemplateLeaf != "" {
		v.IbgpPeerTemplateLeaf = types.StringValue(jsonData.IbgpPeerTemplateLeaf)
	} else {
		v.IbgpPeerTemplateLeaf = types.StringNull()
	}

	if jsonData.InbandDhcpServers != "" {
		v.InbandDhcpServers = types.StringValue(jsonData.InbandDhcpServers)
	} else {
		v.InbandDhcpServers = types.StringNull()
	}

	if jsonData.InbandMgmt != "" {
		x, _ := strconv.ParseBool(jsonData.InbandMgmt)
		v.InbandMgmt = types.BoolValue(x)
	} else {
		v.InbandMgmt = types.BoolNull()
	}

	if jsonData.IsisAuthEnable != "" {
		x, _ := strconv.ParseBool(jsonData.IsisAuthEnable)
		v.IsisAuthEnable = types.BoolValue(x)
	} else {
		v.IsisAuthEnable = types.BoolNull()
	}

	if jsonData.IsisAuthKey != "" {
		v.IsisAuthKey = types.StringValue(jsonData.IsisAuthKey)
	} else {
		v.IsisAuthKey = types.StringNull()
	}

	if jsonData.IsisAuthKeychainKeyId != nil {
		if jsonData.IsisAuthKeychainKeyId.IsEmpty() {
			v.IsisAuthKeychainKeyId = types.Int64Null()
		} else {
			v.IsisAuthKeychainKeyId = types.Int64Value(int64(*jsonData.IsisAuthKeychainKeyId))
		}
	} else {
		v.IsisAuthKeychainKeyId = types.Int64Null()
	}

	if jsonData.IsisAuthKeychainName != "" {
		v.IsisAuthKeychainName = types.StringValue(jsonData.IsisAuthKeychainName)
	} else {
		v.IsisAuthKeychainName = types.StringNull()
	}

	if jsonData.IsisLevel != "" {
		v.IsisLevel = types.StringValue(jsonData.IsisLevel)
	} else {
		v.IsisLevel = types.StringNull()
	}

	if jsonData.IsisOverloadElapseTime != nil {
		if jsonData.IsisOverloadElapseTime.IsEmpty() {
			v.IsisOverloadElapseTime = types.Int64Null()
		} else {
			v.IsisOverloadElapseTime = types.Int64Value(int64(*jsonData.IsisOverloadElapseTime))
		}
	} else {
		v.IsisOverloadElapseTime = types.Int64Null()
	}

	if jsonData.IsisOverloadEnable != "" {
		x, _ := strconv.ParseBool(jsonData.IsisOverloadEnable)
		v.IsisOverloadEnable = types.BoolValue(x)
	} else {
		v.IsisOverloadEnable = types.BoolNull()
	}

	if jsonData.IsisP2pEnable != "" {
		x, _ := strconv.ParseBool(jsonData.IsisP2pEnable)
		v.IsisP2pEnable = types.BoolValue(x)
	} else {
		v.IsisP2pEnable = types.BoolNull()
	}

	if jsonData.L2HostIntfMtu != nil {
		if jsonData.L2HostIntfMtu.IsEmpty() {
			v.L2HostIntfMtu = types.Int64Null()
		} else {
			v.L2HostIntfMtu = types.Int64Value(int64(*jsonData.L2HostIntfMtu))
		}
	} else {
		v.L2HostIntfMtu = types.Int64Null()
	}

	if jsonData.L2SegmentIdRange != "" {
		v.L2SegmentIdRange = types.StringValue(jsonData.L2SegmentIdRange)
	} else {
		v.L2SegmentIdRange = types.StringNull()
	}

	if jsonData.L3vniMcastGroup != "" {
		v.L3vniMcastGroup = types.StringValue(jsonData.L3vniMcastGroup)
	} else {
		v.L3vniMcastGroup = types.StringNull()
	}

	if jsonData.L3PartitionIdRange != "" {
		v.L3PartitionIdRange = types.StringValue(jsonData.L3PartitionIdRange)
	} else {
		v.L3PartitionIdRange = types.StringNull()
	}

	if jsonData.LinkStateRouting != "" {
		v.LinkStateRouting = types.StringValue(jsonData.LinkStateRouting)
	} else {
		v.LinkStateRouting = types.StringNull()
	}

	if jsonData.LinkStateRoutingTag != "" {
		v.LinkStateRoutingTag = types.StringValue(jsonData.LinkStateRoutingTag)
	} else {
		v.LinkStateRoutingTag = types.StringNull()
	}

	if jsonData.Loopback0Ipv6Range != "" {
		v.Loopback0Ipv6Range = types.StringValue(jsonData.Loopback0Ipv6Range)
	} else {
		v.Loopback0Ipv6Range = types.StringNull()
	}

	if jsonData.Loopback0IpRange != "" {
		v.Loopback0IpRange = types.StringValue(jsonData.Loopback0IpRange)
	} else {
		v.Loopback0IpRange = types.StringNull()
	}

	if jsonData.Loopback1Ipv6Range != "" {
		v.Loopback1Ipv6Range = types.StringValue(jsonData.Loopback1Ipv6Range)
	} else {
		v.Loopback1Ipv6Range = types.StringNull()
	}

	if jsonData.Loopback1IpRange != "" {
		v.Loopback1IpRange = types.StringValue(jsonData.Loopback1IpRange)
	} else {
		v.Loopback1IpRange = types.StringNull()
	}

	if jsonData.MacsecAlgorithm != "" {
		v.MacsecAlgorithm = types.StringValue(jsonData.MacsecAlgorithm)
	} else {
		v.MacsecAlgorithm = types.StringNull()
	}

	if jsonData.MacsecCipherSuite != "" {
		v.MacsecCipherSuite = types.StringValue(jsonData.MacsecCipherSuite)
	} else {
		v.MacsecCipherSuite = types.StringNull()
	}

	if jsonData.MacsecFallbackAlgorithm != "" {
		v.MacsecFallbackAlgorithm = types.StringValue(jsonData.MacsecFallbackAlgorithm)
	} else {
		v.MacsecFallbackAlgorithm = types.StringNull()
	}

	if jsonData.MacsecFallbackKeyString != "" {
		v.MacsecFallbackKeyString = types.StringValue(jsonData.MacsecFallbackKeyString)
	} else {
		v.MacsecFallbackKeyString = types.StringNull()
	}

	if jsonData.MacsecKeyString != "" {
		v.MacsecKeyString = types.StringValue(jsonData.MacsecKeyString)
	} else {
		v.MacsecKeyString = types.StringNull()
	}

	if jsonData.MacsecReportTimer != nil {
		if jsonData.MacsecReportTimer.IsEmpty() {
			v.MacsecReportTimer = types.Int64Null()
		} else {
			v.MacsecReportTimer = types.Int64Value(int64(*jsonData.MacsecReportTimer))
		}
	} else {
		v.MacsecReportTimer = types.Int64Null()
	}

	if jsonData.MgmtGw != "" {
		v.MgmtGw = types.StringValue(jsonData.MgmtGw)
	} else {
		v.MgmtGw = types.StringNull()
	}

	if jsonData.MgmtPrefix != nil {
		if jsonData.MgmtPrefix.IsEmpty() {
			v.MgmtPrefix = types.Int64Null()
		} else {
			v.MgmtPrefix = types.Int64Value(int64(*jsonData.MgmtPrefix))
		}
	} else {
		v.MgmtPrefix = types.Int64Null()
	}

	if jsonData.MgmtV6prefix != nil {
		if jsonData.MgmtV6prefix.IsEmpty() {
			v.MgmtV6prefix = types.Int64Null()
		} else {
			v.MgmtV6prefix = types.Int64Value(int64(*jsonData.MgmtV6prefix))
		}
	} else {
		v.MgmtV6prefix = types.Int64Null()
	}

	if jsonData.MplsHandoff != "" {
		x, _ := strconv.ParseBool(jsonData.MplsHandoff)
		v.MplsHandoff = types.BoolValue(x)
	} else {
		v.MplsHandoff = types.BoolNull()
	}

	if jsonData.MplsLbId != nil {
		if jsonData.MplsLbId.IsEmpty() {
			v.MplsLbId = types.Int64Null()
		} else {
			v.MplsLbId = types.Int64Value(int64(*jsonData.MplsLbId))
		}
	} else {
		v.MplsLbId = types.Int64Null()
	}

	if jsonData.MplsLoopbackIpRange != "" {
		v.MplsLoopbackIpRange = types.StringValue(jsonData.MplsLoopbackIpRange)
	} else {
		v.MplsLoopbackIpRange = types.StringNull()
	}

	if jsonData.MstInstanceRange != "" {
		v.MstInstanceRange = types.StringValue(jsonData.MstInstanceRange)
	} else {
		v.MstInstanceRange = types.StringNull()
	}

	if jsonData.MulticastGroupSubnet != "" {
		v.MulticastGroupSubnet = types.StringValue(jsonData.MulticastGroupSubnet)
	} else {
		v.MulticastGroupSubnet = types.StringNull()
	}

	if jsonData.NetflowExporterList != "" {
		v.NetflowExporterList = types.StringValue(jsonData.NetflowExporterList)
	} else {
		v.NetflowExporterList = types.StringNull()
	}

	if jsonData.NetflowMonitorList != "" {
		v.NetflowMonitorList = types.StringValue(jsonData.NetflowMonitorList)
	} else {
		v.NetflowMonitorList = types.StringNull()
	}

	if jsonData.NetflowRecordList != "" {
		v.NetflowRecordList = types.StringValue(jsonData.NetflowRecordList)
	} else {
		v.NetflowRecordList = types.StringNull()
	}

	if jsonData.NetworkVlanRange != "" {
		v.NetworkVlanRange = types.StringValue(jsonData.NetworkVlanRange)
	} else {
		v.NetworkVlanRange = types.StringNull()
	}

	if jsonData.NtpServerIpList != "" {
		v.NtpServerIpList = types.StringValue(jsonData.NtpServerIpList)
	} else {
		v.NtpServerIpList = types.StringNull()
	}

	if jsonData.NtpServerVrf != "" {
		v.NtpServerVrf = types.StringValue(jsonData.NtpServerVrf)
	} else {
		v.NtpServerVrf = types.StringNull()
	}

	if jsonData.NveLbId != nil {
		if jsonData.NveLbId.IsEmpty() {
			v.NveLbId = types.Int64Null()
		} else {
			v.NveLbId = types.Int64Value(int64(*jsonData.NveLbId))
		}
	} else {
		v.NveLbId = types.Int64Null()
	}

	if jsonData.NxapiHttpsPort != nil {
		if jsonData.NxapiHttpsPort.IsEmpty() {
			v.NxapiHttpsPort = types.Int64Null()
		} else {
			v.NxapiHttpsPort = types.Int64Value(int64(*jsonData.NxapiHttpsPort))
		}
	} else {
		v.NxapiHttpsPort = types.Int64Null()
	}

	if jsonData.NxapiHttpPort != nil {
		if jsonData.NxapiHttpPort.IsEmpty() {
			v.NxapiHttpPort = types.Int64Null()
		} else {
			v.NxapiHttpPort = types.Int64Value(int64(*jsonData.NxapiHttpPort))
		}
	} else {
		v.NxapiHttpPort = types.Int64Null()
	}

	if jsonData.ObjectTrackingNumberRange != "" {
		v.ObjectTrackingNumberRange = types.StringValue(jsonData.ObjectTrackingNumberRange)
	} else {
		v.ObjectTrackingNumberRange = types.StringNull()
	}

	if jsonData.OspfAreaId != "" {
		v.OspfAreaId = types.StringValue(jsonData.OspfAreaId)
	} else {
		v.OspfAreaId = types.StringNull()
	}

	if jsonData.OspfAuthEnable != "" {
		x, _ := strconv.ParseBool(jsonData.OspfAuthEnable)
		v.OspfAuthEnable = types.BoolValue(x)
	} else {
		v.OspfAuthEnable = types.BoolNull()
	}

	if jsonData.OspfAuthKey != "" {
		v.OspfAuthKey = types.StringValue(jsonData.OspfAuthKey)
	} else {
		v.OspfAuthKey = types.StringNull()
	}

	if jsonData.OspfAuthKeyId != nil {
		if jsonData.OspfAuthKeyId.IsEmpty() {
			v.OspfAuthKeyId = types.Int64Null()
		} else {
			v.OspfAuthKeyId = types.Int64Value(int64(*jsonData.OspfAuthKeyId))
		}
	} else {
		v.OspfAuthKeyId = types.Int64Null()
	}

	if jsonData.OverlayMode != "" {
		v.OverlayMode = types.StringValue(jsonData.OverlayMode)
	} else {
		v.OverlayMode = types.StringNull()
	}

	if jsonData.PerVrfLoopbackAutoProvision != "" {
		x, _ := strconv.ParseBool(jsonData.PerVrfLoopbackAutoProvision)
		v.PerVrfLoopbackAutoProvision = types.BoolValue(x)
	} else {
		v.PerVrfLoopbackAutoProvision = types.BoolNull()
	}

	if jsonData.PerVrfLoopbackIpRange != "" {
		v.PerVrfLoopbackIpRange = types.StringValue(jsonData.PerVrfLoopbackIpRange)
	} else {
		v.PerVrfLoopbackIpRange = types.StringNull()
	}

	if jsonData.PhantomRpLbId1 != nil {
		if jsonData.PhantomRpLbId1.IsEmpty() {
			v.PhantomRpLbId1 = types.Int64Null()
		} else {
			v.PhantomRpLbId1 = types.Int64Value(int64(*jsonData.PhantomRpLbId1))
		}
	} else {
		v.PhantomRpLbId1 = types.Int64Null()
	}

	if jsonData.PhantomRpLbId2 != nil {
		if jsonData.PhantomRpLbId2.IsEmpty() {
			v.PhantomRpLbId2 = types.Int64Null()
		} else {
			v.PhantomRpLbId2 = types.Int64Value(int64(*jsonData.PhantomRpLbId2))
		}
	} else {
		v.PhantomRpLbId2 = types.Int64Null()
	}

	if jsonData.PhantomRpLbId3 != nil {
		if jsonData.PhantomRpLbId3.IsEmpty() {
			v.PhantomRpLbId3 = types.Int64Null()
		} else {
			v.PhantomRpLbId3 = types.Int64Value(int64(*jsonData.PhantomRpLbId3))
		}
	} else {
		v.PhantomRpLbId3 = types.Int64Null()
	}

	if jsonData.PhantomRpLbId4 != nil {
		if jsonData.PhantomRpLbId4.IsEmpty() {
			v.PhantomRpLbId4 = types.Int64Null()
		} else {
			v.PhantomRpLbId4 = types.Int64Value(int64(*jsonData.PhantomRpLbId4))
		}
	} else {
		v.PhantomRpLbId4 = types.Int64Null()
	}

	if jsonData.PimHelloAuthEnable != "" {
		x, _ := strconv.ParseBool(jsonData.PimHelloAuthEnable)
		v.PimHelloAuthEnable = types.BoolValue(x)
	} else {
		v.PimHelloAuthEnable = types.BoolNull()
	}

	if jsonData.PimHelloAuthKey != "" {
		v.PimHelloAuthKey = types.StringValue(jsonData.PimHelloAuthKey)
	} else {
		v.PimHelloAuthKey = types.StringNull()
	}

	if jsonData.PmEnable != "" {
		x, _ := strconv.ParseBool(jsonData.PmEnable)
		v.PmEnable = types.BoolValue(x)
	} else {
		v.PmEnable = types.BoolNull()
	}

	if jsonData.PowerRedundancyMode != "" {
		v.PowerRedundancyMode = types.StringValue(jsonData.PowerRedundancyMode)
	} else {
		v.PowerRedundancyMode = types.StringNull()
	}

	if jsonData.PtpDomainId != nil {
		if jsonData.PtpDomainId.IsEmpty() {
			v.PtpDomainId = types.Int64Null()
		} else {
			v.PtpDomainId = types.Int64Value(int64(*jsonData.PtpDomainId))
		}
	} else {
		v.PtpDomainId = types.Int64Null()
	}

	if jsonData.PtpLbId != nil {
		if jsonData.PtpLbId.IsEmpty() {
			v.PtpLbId = types.Int64Null()
		} else {
			v.PtpLbId = types.Int64Value(int64(*jsonData.PtpLbId))
		}
	} else {
		v.PtpLbId = types.Int64Null()
	}

	if jsonData.ReplicationMode != "" {
		v.ReplicationMode = types.StringValue(jsonData.ReplicationMode)
	} else {
		v.ReplicationMode = types.StringNull()
	}

	if jsonData.RouterIdRange != "" {
		v.RouterIdRange = types.StringValue(jsonData.RouterIdRange)
	} else {
		v.RouterIdRange = types.StringNull()
	}

	if jsonData.RouteMapSequenceNumberRange != "" {
		v.RouteMapSequenceNumberRange = types.StringValue(jsonData.RouteMapSequenceNumberRange)
	} else {
		v.RouteMapSequenceNumberRange = types.StringNull()
	}

	if jsonData.RpCount != nil {
		if jsonData.RpCount.IsEmpty() {
			v.RpCount = types.Int64Null()
		} else {
			v.RpCount = types.Int64Value(int64(*jsonData.RpCount))
		}
	} else {
		v.RpCount = types.Int64Null()
	}

	if jsonData.RpLbId != nil {
		if jsonData.RpLbId.IsEmpty() {
			v.RpLbId = types.Int64Null()
		} else {
			v.RpLbId = types.Int64Value(int64(*jsonData.RpLbId))
		}
	} else {
		v.RpLbId = types.Int64Null()
	}

	if jsonData.RpMode != "" {
		v.RpMode = types.StringValue(jsonData.RpMode)
	} else {
		v.RpMode = types.StringNull()
	}

	if jsonData.RrCount != nil {
		if jsonData.RrCount.IsEmpty() {
			v.RrCount = types.Int64Null()
		} else {
			v.RrCount = types.Int64Value(int64(*jsonData.RrCount))
		}
	} else {
		v.RrCount = types.Int64Null()
	}

	if jsonData.SeedSwitchCoreInterfaces != "" {
		v.SeedSwitchCoreInterfaces = types.StringValue(jsonData.SeedSwitchCoreInterfaces)
	} else {
		v.SeedSwitchCoreInterfaces = types.StringNull()
	}

	if jsonData.ServiceNetworkVlanRange != "" {
		v.ServiceNetworkVlanRange = types.StringValue(jsonData.ServiceNetworkVlanRange)
	} else {
		v.ServiceNetworkVlanRange = types.StringNull()
	}

	if jsonData.SiteId != "" {
		v.SiteId = types.StringValue(jsonData.SiteId)
	} else {
		v.SiteId = types.StringNull()
	}

	if jsonData.SlaIdRange != "" {
		v.SlaIdRange = types.StringValue(jsonData.SlaIdRange)
	} else {
		v.SlaIdRange = types.StringNull()
	}

	if jsonData.SnmpServerHostTrap != "" {
		x, _ := strconv.ParseBool(jsonData.SnmpServerHostTrap)
		v.SnmpServerHostTrap = types.BoolValue(x)
	} else {
		v.SnmpServerHostTrap = types.BoolNull()
	}

	if jsonData.SpineSwitchCoreInterfaces != "" {
		v.SpineSwitchCoreInterfaces = types.StringValue(jsonData.SpineSwitchCoreInterfaces)
	} else {
		v.SpineSwitchCoreInterfaces = types.StringNull()
	}

	if jsonData.StaticUnderlayIpAlloc != "" {
		x, _ := strconv.ParseBool(jsonData.StaticUnderlayIpAlloc)
		v.StaticUnderlayIpAlloc = types.BoolValue(x)
	} else {
		v.StaticUnderlayIpAlloc = types.BoolNull()
	}

	if jsonData.StpBridgePriority != nil {
		if jsonData.StpBridgePriority.IsEmpty() {
			v.StpBridgePriority = types.Int64Null()
		} else {
			v.StpBridgePriority = types.Int64Value(int64(*jsonData.StpBridgePriority))
		}
	} else {
		v.StpBridgePriority = types.Int64Null()
	}

	if jsonData.StpRootOption != "" {
		v.StpRootOption = types.StringValue(jsonData.StpRootOption)
	} else {
		v.StpRootOption = types.StringNull()
	}

	if jsonData.StpVlanRange != "" {
		v.StpVlanRange = types.StringValue(jsonData.StpVlanRange)
	} else {
		v.StpVlanRange = types.StringNull()
	}

	if jsonData.StrictCcMode != "" {
		x, _ := strconv.ParseBool(jsonData.StrictCcMode)
		v.StrictCcMode = types.BoolValue(x)
	} else {
		v.StrictCcMode = types.BoolNull()
	}

	if jsonData.SubinterfaceRange != "" {
		v.SubinterfaceRange = types.StringValue(jsonData.SubinterfaceRange)
	} else {
		v.SubinterfaceRange = types.StringNull()
	}

	if jsonData.SubnetRange != "" {
		v.SubnetRange = types.StringValue(jsonData.SubnetRange)
	} else {
		v.SubnetRange = types.StringNull()
	}

	if jsonData.SubnetTargetMask != nil {
		if jsonData.SubnetTargetMask.IsEmpty() {
			v.SubnetTargetMask = types.Int64Null()
		} else {
			v.SubnetTargetMask = types.Int64Value(int64(*jsonData.SubnetTargetMask))
		}
	} else {
		v.SubnetTargetMask = types.Int64Null()
	}

	if jsonData.SyslogServerIpList != "" {
		v.SyslogServerIpList = types.StringValue(jsonData.SyslogServerIpList)
	} else {
		v.SyslogServerIpList = types.StringNull()
	}

	if jsonData.SyslogServerVrf != "" {
		v.SyslogServerVrf = types.StringValue(jsonData.SyslogServerVrf)
	} else {
		v.SyslogServerVrf = types.StringNull()
	}

	if jsonData.SyslogSev != "" {
		v.SyslogSev = types.StringValue(jsonData.SyslogSev)
	} else {
		v.SyslogSev = types.StringNull()
	}

	if jsonData.TcamAllocation != "" {
		x, _ := strconv.ParseBool(jsonData.TcamAllocation)
		v.TcamAllocation = types.BoolValue(x)
	} else {
		v.TcamAllocation = types.BoolNull()
	}

	if jsonData.UnderlayIsV6 != "" {
		x, _ := strconv.ParseBool(jsonData.UnderlayIsV6)
		v.UnderlayIsV6 = types.BoolValue(x)
	} else {
		v.UnderlayIsV6 = types.BoolNull()
	}

	if jsonData.UnnumBootstrapLbId != nil {
		if jsonData.UnnumBootstrapLbId.IsEmpty() {
			v.UnnumBootstrapLbId = types.Int64Null()
		} else {
			v.UnnumBootstrapLbId = types.Int64Value(int64(*jsonData.UnnumBootstrapLbId))
		}
	} else {
		v.UnnumBootstrapLbId = types.Int64Null()
	}

	if jsonData.UnnumDhcpEnd != "" {
		v.UnnumDhcpEnd = types.StringValue(jsonData.UnnumDhcpEnd)
	} else {
		v.UnnumDhcpEnd = types.StringNull()
	}

	if jsonData.UnnumDhcpStart != "" {
		v.UnnumDhcpStart = types.StringValue(jsonData.UnnumDhcpStart)
	} else {
		v.UnnumDhcpStart = types.StringNull()
	}

	if jsonData.UseLinkLocal != "" {
		x, _ := strconv.ParseBool(jsonData.UseLinkLocal)
		v.UseLinkLocal = types.BoolValue(x)
	} else {
		v.UseLinkLocal = types.BoolNull()
	}

	if jsonData.V6SubnetRange != "" {
		v.V6SubnetRange = types.StringValue(jsonData.V6SubnetRange)
	} else {
		v.V6SubnetRange = types.StringNull()
	}

	if jsonData.V6SubnetTargetMask != nil {
		if jsonData.V6SubnetTargetMask.IsEmpty() {
			v.V6SubnetTargetMask = types.Int64Null()
		} else {
			v.V6SubnetTargetMask = types.Int64Value(int64(*jsonData.V6SubnetTargetMask))
		}
	} else {
		v.V6SubnetTargetMask = types.Int64Null()
	}

	if jsonData.VpcAutoRecoveryTime != nil {
		if jsonData.VpcAutoRecoveryTime.IsEmpty() {
			v.VpcAutoRecoveryTime = types.Int64Null()
		} else {
			v.VpcAutoRecoveryTime = types.Int64Value(int64(*jsonData.VpcAutoRecoveryTime))
		}
	} else {
		v.VpcAutoRecoveryTime = types.Int64Null()
	}

	if jsonData.VpcDelayRestore != nil {
		if jsonData.VpcDelayRestore.IsEmpty() {
			v.VpcDelayRestore = types.Int64Null()
		} else {
			v.VpcDelayRestore = types.Int64Value(int64(*jsonData.VpcDelayRestore))
		}
	} else {
		v.VpcDelayRestore = types.Int64Null()
	}

	if jsonData.VpcDomainIdRange != "" {
		v.VpcDomainIdRange = types.StringValue(jsonData.VpcDomainIdRange)
	} else {
		v.VpcDomainIdRange = types.StringNull()
	}

	if jsonData.VpcEnableIpv6NdSync != "" {
		x, _ := strconv.ParseBool(jsonData.VpcEnableIpv6NdSync)
		v.VpcEnableIpv6NdSync = types.BoolValue(x)
	} else {
		v.VpcEnableIpv6NdSync = types.BoolNull()
	}

	if jsonData.VpcPeerKeepAliveOption != "" {
		v.VpcPeerKeepAliveOption = types.StringValue(jsonData.VpcPeerKeepAliveOption)
	} else {
		v.VpcPeerKeepAliveOption = types.StringNull()
	}

	if jsonData.VpcPeerLinkPo != nil {
		if jsonData.VpcPeerLinkPo.IsEmpty() {
			v.VpcPeerLinkPo = types.Int64Null()
		} else {
			v.VpcPeerLinkPo = types.Int64Value(int64(*jsonData.VpcPeerLinkPo))
		}
	} else {
		v.VpcPeerLinkPo = types.Int64Null()
	}

	if jsonData.VpcPeerLinkVlan != nil {
		if jsonData.VpcPeerLinkVlan.IsEmpty() {
			v.VpcPeerLinkVlan = types.Int64Null()
		} else {
			v.VpcPeerLinkVlan = types.Int64Value(int64(*jsonData.VpcPeerLinkVlan))
		}
	} else {
		v.VpcPeerLinkVlan = types.Int64Null()
	}

	if jsonData.VrfLiteAutoconfig != "" {
		v.VrfLiteAutoconfig = types.StringValue(jsonData.VrfLiteAutoconfig)
	} else {
		v.VrfLiteAutoconfig = types.StringNull()
	}

	if jsonData.VrfVlanRange != "" {
		v.VrfVlanRange = types.StringValue(jsonData.VrfVlanRange)
	} else {
		v.VrfVlanRange = types.StringNull()
	}

	if jsonData.DefaultNetwork != "" {
		v.DefaultNetwork = types.StringValue(jsonData.DefaultNetwork)
	} else {
		v.DefaultNetwork = types.StringNull()
	}

	if jsonData.DefaultPvlanSecNetwork != "" {
		v.DefaultPvlanSecNetwork = types.StringValue(jsonData.DefaultPvlanSecNetwork)
	} else {
		v.DefaultPvlanSecNetwork = types.StringNull()
	}

	if jsonData.DefaultVrf != "" {
		v.DefaultVrf = types.StringValue(jsonData.DefaultVrf)
	} else {
		v.DefaultVrf = types.StringNull()
	}

	if jsonData.EnableRealTimeBackup != "" {
		x, _ := strconv.ParseBool(jsonData.EnableRealTimeBackup)
		v.EnableRealTimeBackup = types.BoolValue(x)
	} else {
		v.EnableRealTimeBackup = types.BoolNull()
	}

	if jsonData.EnableScheduledBackup != "" {
		x, _ := strconv.ParseBool(jsonData.EnableScheduledBackup)
		v.EnableScheduledBackup = types.BoolValue(x)
	} else {
		v.EnableScheduledBackup = types.BoolNull()
	}

	if jsonData.NetworkExtensionTemplate != "" {
		v.NetworkExtensionTemplate = types.StringValue(jsonData.NetworkExtensionTemplate)
	} else {
		v.NetworkExtensionTemplate = types.StringNull()
	}

	if jsonData.ScheduledTime != "" {
		v.ScheduledTime = types.StringValue(jsonData.ScheduledTime)
	} else {
		v.ScheduledTime = types.StringNull()
	}

	if jsonData.VrfExtensionTemplate != "" {
		v.VrfExtensionTemplate = types.StringValue(jsonData.VrfExtensionTemplate)
	} else {
		v.VrfExtensionTemplate = types.StringNull()
	}

	if jsonData.DeploymentStatus != "" {
		v.DeploymentStatus = types.StringValue(jsonData.DeploymentStatus)
	} else {
		v.DeploymentStatus = types.StringNull()
	}

	if jsonData.Ipv6MulticastGroupSubnet != "" {
		v.Ipv6MulticastGroupSubnet = types.StringValue(jsonData.Ipv6MulticastGroupSubnet)
	} else {
		v.Ipv6MulticastGroupSubnet = types.StringNull()
	}

	if jsonData.EnableTrmv6 != "" {
		x, _ := strconv.ParseBool(jsonData.EnableTrmv6)
		v.EnableTrmv6 = types.BoolValue(x)
	} else {
		v.EnableTrmv6 = types.BoolNull()
	}

	if jsonData.L3vniIpv6McastGroup != "" {
		v.L3vniIpv6McastGroup = types.StringValue(jsonData.L3vniIpv6McastGroup)
	} else {
		v.L3vniIpv6McastGroup = types.StringNull()
	}

	if jsonData.MvpnVriIdRange != "" {
		v.MvpnVriIdRange = types.StringValue(jsonData.MvpnVriIdRange)
	} else {
		v.MvpnVriIdRange = types.StringNull()
	}

	if jsonData.EnableVriIdRealloc != "" {
		x, _ := strconv.ParseBool(jsonData.EnableVriIdRealloc)
		v.EnableVriIdRealloc = types.BoolValue(x)
	} else {
		v.EnableVriIdRealloc = types.BoolNull()
	}

	if jsonData.EnableAggAccIdRange != "" {
		x, _ := strconv.ParseBool(jsonData.EnableAggAccIdRange)
		v.EnableAggAccIdRange = types.BoolValue(x)
	} else {
		v.EnableAggAccIdRange = types.BoolNull()
	}

	if jsonData.AggAccVpcPoIdRange != "" {
		v.AggAccVpcPoIdRange = types.StringValue(jsonData.AggAccVpcPoIdRange)
	} else {
		v.AggAccVpcPoIdRange = types.StringNull()
	}

	if jsonData.IsisAreaNum != "" {
		v.IsisAreaNum = types.StringValue(jsonData.IsisAreaNum)
	} else {
		v.IsisAreaNum = types.StringNull()
	}

	if jsonData.EnableSgt != "" {
		x, _ := strconv.ParseBool(jsonData.EnableSgt)
		v.EnableSgt = types.BoolValue(x)
	} else {
		v.EnableSgt = types.BoolNull()
	}

	if jsonData.SgtNamePrefix != "" {
		v.SgtNamePrefix = types.StringValue(jsonData.SgtNamePrefix)
	} else {
		v.SgtNamePrefix = types.StringNull()
	}

	if jsonData.SgtIdRange != "" {
		v.SgtIdRange = types.StringValue(jsonData.SgtIdRange)
	} else {
		v.SgtIdRange = types.StringNull()
	}

	if jsonData.SgtPreprovision != "" {
		x, _ := strconv.ParseBool(jsonData.SgtPreprovision)
		v.SgtPreprovision = types.BoolValue(x)
	} else {
		v.SgtPreprovision = types.BoolNull()
	}

	if jsonData.EnableDciMacsec != "" {
		x, _ := strconv.ParseBool(jsonData.EnableDciMacsec)
		v.EnableDciMacsec = types.BoolValue(x)
	} else {
		v.EnableDciMacsec = types.BoolNull()
	}

	if jsonData.EnableQkd != "" {
		x, _ := strconv.ParseBool(jsonData.EnableQkd)
		v.EnableQkd = types.BoolValue(x)
	} else {
		v.EnableQkd = types.BoolNull()
	}

	if jsonData.AllowL3vniNoVlan != "" {
		x, _ := strconv.ParseBool(jsonData.AllowL3vniNoVlan)
		v.AllowL3vniNoVlan = types.BoolValue(x)
	} else {
		v.AllowL3vniNoVlan = types.BoolNull()
	}

	if jsonData.EnableL3vniNoVlan != "" {
		x, _ := strconv.ParseBool(jsonData.EnableL3vniNoVlan)
		v.EnableL3vniNoVlan = types.BoolValue(x)
	} else {
		v.EnableL3vniNoVlan = types.BoolNull()
	}

	if jsonData.MplsIsisAreaNum != "" {
		v.MplsIsisAreaNum = types.StringValue(jsonData.MplsIsisAreaNum)
	} else {
		v.MplsIsisAreaNum = types.StringNull()
	}

	if jsonData.EnableAiMlQosPolicy != "" {
		x, _ := strconv.ParseBool(jsonData.EnableAiMlQosPolicy)
		v.EnableAiMlQosPolicy = types.BoolValue(x)
	} else {
		v.EnableAiMlQosPolicy = types.BoolNull()
	}

	if jsonData.AiMlQosPolicy != "" {
		v.AiMlQosPolicy = types.StringValue(jsonData.AiMlQosPolicy)
	} else {
		v.AiMlQosPolicy = types.StringNull()
	}

	if jsonData.PfcWatchInt != nil {
		if jsonData.PfcWatchInt.IsEmpty() {
			v.PfcWatchInt = types.Int64Null()
		} else {
			v.PfcWatchInt = types.Int64Value(int64(*jsonData.PfcWatchInt))
		}
	} else {
		v.PfcWatchInt = types.Int64Null()
	}

	if jsonData.EnableRtIntfStats != "" {
		x, _ := strconv.ParseBool(jsonData.EnableRtIntfStats)
		v.EnableRtIntfStats = types.BoolValue(x)
	} else {
		v.EnableRtIntfStats = types.BoolNull()
	}

	if jsonData.Ipv6AnycastRpIpRange != "" {
		v.Ipv6AnycastRpIpRange = types.StringValue(jsonData.Ipv6AnycastRpIpRange)
	} else {
		v.Ipv6AnycastRpIpRange = types.StringNull()
	}

	if jsonData.KmeServerIp != "" {
		v.KmeServerIp = types.StringValue(jsonData.KmeServerIp)
	} else {
		v.KmeServerIp = types.StringNull()
	}

	if jsonData.KmeServerPort != nil {
		if jsonData.KmeServerPort.IsEmpty() {
			v.KmeServerPort = types.Int64Null()
		} else {
			v.KmeServerPort = types.Int64Value(int64(*jsonData.KmeServerPort))
		}
	} else {
		v.KmeServerPort = types.Int64Null()
	}

	if jsonData.TrustpointLabel != "" {
		v.TrustpointLabel = types.StringValue(jsonData.TrustpointLabel)
	} else {
		v.TrustpointLabel = types.StringNull()
	}

	if jsonData.IgnoreCert != "" {
		x, _ := strconv.ParseBool(jsonData.IgnoreCert)
		v.IgnoreCert = types.BoolValue(x)
	} else {
		v.IgnoreCert = types.BoolNull()
	}

	if jsonData.DeploymentFreeze != "" {
		x, _ := strconv.ParseBool(jsonData.DeploymentFreeze)
		v.DeploymentFreeze = types.BoolValue(x)
	} else {
		v.DeploymentFreeze = types.BoolNull()
	}

	if jsonData.QkdProfileName != "" {
		v.QkdProfileName = types.StringValue(jsonData.QkdProfileName)
	} else {
		v.QkdProfileName = types.StringNull()
	}

	if jsonData.PerVrfLoopbackAutoProvisionV6 != "" {
		x, _ := strconv.ParseBool(jsonData.PerVrfLoopbackAutoProvisionV6)
		v.PerVrfLoopbackAutoProvisionV6 = types.BoolValue(x)
	} else {
		v.PerVrfLoopbackAutoProvisionV6 = types.BoolNull()
	}

	if jsonData.PerVrfLoopbackIpRangeV6 != "" {
		v.PerVrfLoopbackIpRangeV6 = types.StringValue(jsonData.PerVrfLoopbackIpRangeV6)
	} else {
		v.PerVrfLoopbackIpRangeV6 = types.StringNull()
	}

	if jsonData.EsrOption != "" {
		v.EsrOption = types.StringValue(jsonData.EsrOption)
	} else {
		v.EsrOption = types.StringNull()
	}

	if jsonData.PtpVlanId != nil {
		if jsonData.PtpVlanId.IsEmpty() {
			v.PtpVlanId = types.Int64Null()
		} else {
			v.PtpVlanId = types.Int64Value(int64(*jsonData.PtpVlanId))
		}
	} else {
		v.PtpVlanId = types.Int64Null()
	}

	if jsonData.AllowNxc != "" {
		x, _ := strconv.ParseBool(jsonData.AllowNxc)
		v.AllowNxc = types.BoolValue(x)
	} else {
		v.AllowNxc = types.BoolNull()
	}

	if jsonData.OverwriteGlobalNxc != "" {
		x, _ := strconv.ParseBool(jsonData.OverwriteGlobalNxc)
		v.OverwriteGlobalNxc = types.BoolValue(x)
	} else {
		v.OverwriteGlobalNxc = types.BoolNull()
	}

	if jsonData.NxcDestVrf != "" {
		v.NxcDestVrf = types.StringValue(jsonData.NxcDestVrf)
	} else {
		v.NxcDestVrf = types.StringNull()
	}

	if jsonData.NxcSrcIntf != "" {
		v.NxcSrcIntf = types.StringValue(jsonData.NxcSrcIntf)
	} else {
		v.NxcSrcIntf = types.StringNull()
	}

	if jsonData.NxcProxyServer != "" {
		v.NxcProxyServer = types.StringValue(jsonData.NxcProxyServer)
	} else {
		v.NxcProxyServer = types.StringNull()
	}

	if jsonData.NxcProxyPort != nil {
		if jsonData.NxcProxyPort.IsEmpty() {
			v.NxcProxyPort = types.Int64Null()
		} else {
			v.NxcProxyPort = types.Int64Value(int64(*jsonData.NxcProxyPort))
		}
	} else {
		v.NxcProxyPort = types.Int64Null()
	}

	if jsonData.VpcDelayRestoreTime != nil {
		if jsonData.VpcDelayRestoreTime.IsEmpty() {
			v.VpcDelayRestoreTime = types.Int64Null()
		} else {
			v.VpcDelayRestoreTime = types.Int64Value(int64(*jsonData.VpcDelayRestoreTime))
		}
	} else {
		v.VpcDelayRestoreTime = types.Int64Null()
	}

	if jsonData.FabricType != "" {
		v.FabricType = types.StringValue(jsonData.FabricType)
	} else {
		v.FabricType = types.StringNull()
	}

	if jsonData.ExtFabricType != "" {
		v.ExtFabricType = types.StringValue(jsonData.ExtFabricType)
	} else {
		v.ExtFabricType = types.StringNull()
	}

	if jsonData.EnableAgent != "" {
		x, _ := strconv.ParseBool(jsonData.EnableAgent)
		v.EnableAgent = types.BoolValue(x)
	} else {
		v.EnableAgent = types.BoolNull()
	}

	if jsonData.AgentIntf != "" {
		v.AgentIntf = types.StringValue(jsonData.AgentIntf)
	} else {
		v.AgentIntf = types.StringNull()
	}

	if jsonData.SspineAddDelDebugFlag != "" {
		v.SspineAddDelDebugFlag = types.StringValue(jsonData.SspineAddDelDebugFlag)
	} else {
		v.SspineAddDelDebugFlag = types.StringNull()
	}

	if jsonData.BrfieldDebugFlag != "" {
		v.BrfieldDebugFlag = types.StringValue(jsonData.BrfieldDebugFlag)
	} else {
		v.BrfieldDebugFlag = types.StringNull()
	}

	if jsonData.ActiveMigration != "" {
		x, _ := strconv.ParseBool(jsonData.ActiveMigration)
		v.ActiveMigration = types.BoolValue(x)
	} else {
		v.ActiveMigration = types.BoolNull()
	}

	if jsonData.Ff != "" {
		v.Ff = types.StringValue(jsonData.Ff)
	} else {
		v.Ff = types.StringNull()
	}

	if jsonData.BgpAsPrev != "" {
		v.BgpAsPrev = types.StringValue(jsonData.BgpAsPrev)
	} else {
		v.BgpAsPrev = types.StringNull()
	}

	if jsonData.UnderlayIsV6Prev != "" {
		x, _ := strconv.ParseBool(jsonData.UnderlayIsV6Prev)
		v.UnderlayIsV6Prev = types.BoolValue(x)
	} else {
		v.UnderlayIsV6Prev = types.BoolNull()
	}

	if jsonData.PmEnablePrev != "" {
		x, _ := strconv.ParseBool(jsonData.PmEnablePrev)
		v.PmEnablePrev = types.BoolValue(x)
	} else {
		v.PmEnablePrev = types.BoolNull()
	}

	if jsonData.EnableFabricVpcDomainIdPrev != "" {
		x, _ := strconv.ParseBool(jsonData.EnableFabricVpcDomainIdPrev)
		v.EnableFabricVpcDomainIdPrev = types.BoolValue(x)
	} else {
		v.EnableFabricVpcDomainIdPrev = types.BoolNull()
	}

	if jsonData.OverlayModePrev != "" {
		v.OverlayModePrev = types.StringValue(jsonData.OverlayModePrev)
	} else {
		v.OverlayModePrev = types.StringNull()
	}

	if jsonData.AllowL3vniNoVlanPrev != "" {
		x, _ := strconv.ParseBool(jsonData.AllowL3vniNoVlanPrev)
		v.AllowL3vniNoVlanPrev = types.BoolValue(x)
	} else {
		v.AllowL3vniNoVlanPrev = types.BoolNull()
	}

	if jsonData.EnablePvlanPrev != "" {
		x, _ := strconv.ParseBool(jsonData.EnablePvlanPrev)
		v.EnablePvlanPrev = types.BoolValue(x)
	} else {
		v.EnablePvlanPrev = types.BoolNull()
	}

	if jsonData.AutoUniqueVrfLiteIpPrefixPrev != "" {
		x, _ := strconv.ParseBool(jsonData.AutoUniqueVrfLiteIpPrefixPrev)
		v.AutoUniqueVrfLiteIpPrefixPrev = types.BoolValue(x)
	} else {
		v.AutoUniqueVrfLiteIpPrefixPrev = types.BoolNull()
	}

	if jsonData.PerVrfLoopbackAutoProvisionPrev != "" {
		x, _ := strconv.ParseBool(jsonData.PerVrfLoopbackAutoProvisionPrev)
		v.PerVrfLoopbackAutoProvisionPrev = types.BoolValue(x)
	} else {
		v.PerVrfLoopbackAutoProvisionPrev = types.BoolNull()
	}

	if jsonData.PerVrfLoopbackAutoProvisionV6Prev != "" {
		x, _ := strconv.ParseBool(jsonData.PerVrfLoopbackAutoProvisionV6Prev)
		v.PerVrfLoopbackAutoProvisionV6Prev = types.BoolValue(x)
	} else {
		v.PerVrfLoopbackAutoProvisionV6Prev = types.BoolNull()
	}

	if jsonData.MsoSiteId != "" {
		v.MsoSiteId = types.StringValue(jsonData.MsoSiteId)
	} else {
		v.MsoSiteId = types.StringNull()
	}

	if jsonData.MsoControlerId != "" {
		v.MsoControlerId = types.StringValue(jsonData.MsoControlerId)
	} else {
		v.MsoControlerId = types.StringNull()
	}

	if jsonData.MsoSiteGroupName != "" {
		v.MsoSiteGroupName = types.StringValue(jsonData.MsoSiteGroupName)
	} else {
		v.MsoSiteGroupName = types.StringNull()
	}

	if jsonData.PremsoParentFabric != "" {
		v.PremsoParentFabric = types.StringValue(jsonData.PremsoParentFabric)
	} else {
		v.PremsoParentFabric = types.StringNull()
	}

	if jsonData.MsoConnectivityDeployed != "" {
		v.MsoConnectivityDeployed = types.StringValue(jsonData.MsoConnectivityDeployed)
	} else {
		v.MsoConnectivityDeployed = types.StringNull()
	}

	if jsonData.AnycastRpIpRangeInternal != "" {
		v.AnycastRpIpRangeInternal = types.StringValue(jsonData.AnycastRpIpRangeInternal)
	} else {
		v.AnycastRpIpRangeInternal = types.StringNull()
	}

	if jsonData.Ipv6AnycastRpIpRangeInternal != "" {
		v.Ipv6AnycastRpIpRangeInternal = types.StringValue(jsonData.Ipv6AnycastRpIpRangeInternal)
	} else {
		v.Ipv6AnycastRpIpRangeInternal = types.StringNull()
	}

	if jsonData.DhcpStartInternal != "" {
		v.DhcpStartInternal = types.StringValue(jsonData.DhcpStartInternal)
	} else {
		v.DhcpStartInternal = types.StringNull()
	}

	if jsonData.DhcpEndInternal != "" {
		v.DhcpEndInternal = types.StringValue(jsonData.DhcpEndInternal)
	} else {
		v.DhcpEndInternal = types.StringNull()
	}

	if jsonData.MgmtGwInternal != "" {
		v.MgmtGwInternal = types.StringValue(jsonData.MgmtGwInternal)
	} else {
		v.MgmtGwInternal = types.StringNull()
	}

	if jsonData.MgmtPrefixInternal != nil {
		if jsonData.MgmtPrefixInternal.IsEmpty() {
			v.MgmtPrefixInternal = types.Int64Null()
		} else {
			v.MgmtPrefixInternal = types.Int64Value(int64(*jsonData.MgmtPrefixInternal))
		}
	} else {
		v.MgmtPrefixInternal = types.Int64Null()
	}

	if jsonData.BootstrapMultisubnetInternal != "" {
		v.BootstrapMultisubnetInternal = types.StringValue(jsonData.BootstrapMultisubnetInternal)
	} else {
		v.BootstrapMultisubnetInternal = types.StringNull()
	}

	if jsonData.MgmtV6prefixInternal != nil {
		if jsonData.MgmtV6prefixInternal.IsEmpty() {
			v.MgmtV6prefixInternal = types.Int64Null()
		} else {
			v.MgmtV6prefixInternal = types.Int64Value(int64(*jsonData.MgmtV6prefixInternal))
		}
	} else {
		v.MgmtV6prefixInternal = types.Int64Null()
	}

	if jsonData.DhcpIpv6EnableInternal != "" {
		v.DhcpIpv6EnableInternal = types.StringValue(jsonData.DhcpIpv6EnableInternal)
	} else {
		v.DhcpIpv6EnableInternal = types.StringNull()
	}

	if jsonData.UnnumDhcpStartInternal != "" {
		v.UnnumDhcpStartInternal = types.StringValue(jsonData.UnnumDhcpStartInternal)
	} else {
		v.UnnumDhcpStartInternal = types.StringNull()
	}

	if jsonData.UnnumDhcpEndInternal != "" {
		v.UnnumDhcpEndInternal = types.StringValue(jsonData.UnnumDhcpEndInternal)
	} else {
		v.UnnumDhcpEndInternal = types.StringNull()
	}

	if jsonData.EnableEvpn != "" {
		x, _ := strconv.ParseBool(jsonData.EnableEvpn)
		v.EnableEvpn = types.BoolValue(x)
	} else {
		v.EnableEvpn = types.BoolNull()
	}

	if jsonData.FeaturePtpInternal != "" {
		x, _ := strconv.ParseBool(jsonData.FeaturePtpInternal)
		v.FeaturePtpInternal = types.BoolValue(x)
	} else {
		v.FeaturePtpInternal = types.BoolNull()
	}

	if jsonData.SspineCount != nil {
		if jsonData.SspineCount.IsEmpty() {
			v.SspineCount = types.Int64Null()
		} else {
			v.SspineCount = types.Int64Value(int64(*jsonData.SspineCount))
		}
	} else {
		v.SspineCount = types.Int64Null()
	}

	if jsonData.SpineCount != nil {
		if jsonData.SpineCount.IsEmpty() {
			v.SpineCount = types.Int64Null()
		} else {
			v.SpineCount = types.Int64Value(int64(*jsonData.SpineCount))
		}
	} else {
		v.SpineCount = types.Int64Null()
	}

	if jsonData.AbstractFeatureLeaf != "" {
		v.AbstractFeatureLeaf = types.StringValue(jsonData.AbstractFeatureLeaf)
	} else {
		v.AbstractFeatureLeaf = types.StringNull()
	}

	if jsonData.AbstractFeatureSpine != "" {
		v.AbstractFeatureSpine = types.StringValue(jsonData.AbstractFeatureSpine)
	} else {
		v.AbstractFeatureSpine = types.StringNull()
	}

	if jsonData.AbstractDhcp != "" {
		v.AbstractDhcp = types.StringValue(jsonData.AbstractDhcp)
	} else {
		v.AbstractDhcp = types.StringNull()
	}

	if jsonData.AbstractMulticast != "" {
		v.AbstractMulticast = types.StringValue(jsonData.AbstractMulticast)
	} else {
		v.AbstractMulticast = types.StringNull()
	}

	if jsonData.AbstractAnycastRp != "" {
		v.AbstractAnycastRp = types.StringValue(jsonData.AbstractAnycastRp)
	} else {
		v.AbstractAnycastRp = types.StringNull()
	}

	if jsonData.AbstractLoopbackInterface != "" {
		v.AbstractLoopbackInterface = types.StringValue(jsonData.AbstractLoopbackInterface)
	} else {
		v.AbstractLoopbackInterface = types.StringNull()
	}

	if jsonData.AbstractIsis != "" {
		v.AbstractIsis = types.StringValue(jsonData.AbstractIsis)
	} else {
		v.AbstractIsis = types.StringNull()
	}

	if jsonData.AbstractOspf != "" {
		v.AbstractOspf = types.StringValue(jsonData.AbstractOspf)
	} else {
		v.AbstractOspf = types.StringNull()
	}

	if jsonData.AbstractVpcDomain != "" {
		v.AbstractVpcDomain = types.StringValue(jsonData.AbstractVpcDomain)
	} else {
		v.AbstractVpcDomain = types.StringNull()
	}

	if jsonData.AbstractVlanInterface != "" {
		v.AbstractVlanInterface = types.StringValue(jsonData.AbstractVlanInterface)
	} else {
		v.AbstractVlanInterface = types.StringNull()
	}

	if jsonData.AbstractIsisInterface != "" {
		v.AbstractIsisInterface = types.StringValue(jsonData.AbstractIsisInterface)
	} else {
		v.AbstractIsisInterface = types.StringNull()
	}

	if jsonData.AbstractOspfInterface != "" {
		v.AbstractOspfInterface = types.StringValue(jsonData.AbstractOspfInterface)
	} else {
		v.AbstractOspfInterface = types.StringNull()
	}

	if jsonData.AbstractPimInterface != "" {
		v.AbstractPimInterface = types.StringValue(jsonData.AbstractPimInterface)
	} else {
		v.AbstractPimInterface = types.StringNull()
	}

	if jsonData.AbstractRouteMap != "" {
		v.AbstractRouteMap = types.StringValue(jsonData.AbstractRouteMap)
	} else {
		v.AbstractRouteMap = types.StringNull()
	}

	if jsonData.AbstractBgp != "" {
		v.AbstractBgp = types.StringValue(jsonData.AbstractBgp)
	} else {
		v.AbstractBgp = types.StringNull()
	}

	if jsonData.AbstractBgpRr != "" {
		v.AbstractBgpRr = types.StringValue(jsonData.AbstractBgpRr)
	} else {
		v.AbstractBgpRr = types.StringNull()
	}

	if jsonData.AbstractBgpNeighbor != "" {
		v.AbstractBgpNeighbor = types.StringValue(jsonData.AbstractBgpNeighbor)
	} else {
		v.AbstractBgpNeighbor = types.StringNull()
	}

	if jsonData.AbstractExtraConfigLeaf != "" {
		v.AbstractExtraConfigLeaf = types.StringValue(jsonData.AbstractExtraConfigLeaf)
	} else {
		v.AbstractExtraConfigLeaf = types.StringNull()
	}

	if jsonData.AbstractExtraConfigSpine != "" {
		v.AbstractExtraConfigSpine = types.StringValue(jsonData.AbstractExtraConfigSpine)
	} else {
		v.AbstractExtraConfigSpine = types.StringNull()
	}

	if jsonData.AbstractExtraConfigTor != "" {
		v.AbstractExtraConfigTor = types.StringValue(jsonData.AbstractExtraConfigTor)
	} else {
		v.AbstractExtraConfigTor = types.StringNull()
	}

	if jsonData.AbstractExtraConfigBootstrap != "" {
		v.AbstractExtraConfigBootstrap = types.StringValue(jsonData.AbstractExtraConfigBootstrap)
	} else {
		v.AbstractExtraConfigBootstrap = types.StringNull()
	}

	if jsonData.TempAnycastGateway != "" {
		v.TempAnycastGateway = types.StringValue(jsonData.TempAnycastGateway)
	} else {
		v.TempAnycastGateway = types.StringNull()
	}

	if jsonData.TempVpcDomainMgmt != "" {
		v.TempVpcDomainMgmt = types.StringValue(jsonData.TempVpcDomainMgmt)
	} else {
		v.TempVpcDomainMgmt = types.StringNull()
	}

	if jsonData.TempVpcPeerLink != "" {
		v.TempVpcPeerLink = types.StringValue(jsonData.TempVpcPeerLink)
	} else {
		v.TempVpcPeerLink = types.StringNull()
	}

	if jsonData.AbstractRoutedHost != "" {
		v.AbstractRoutedHost = types.StringValue(jsonData.AbstractRoutedHost)
	} else {
		v.AbstractRoutedHost = types.StringNull()
	}

	if jsonData.UpgradeFromVersion != "" {
		v.UpgradeFromVersion = types.StringValue(jsonData.UpgradeFromVersion)
	} else {
		v.UpgradeFromVersion = types.StringNull()
	}

	if jsonData.TopdownConfigRmTracking != "" {
		v.TopdownConfigRmTracking = types.StringValue(jsonData.TopdownConfigRmTracking)
	} else {
		v.TopdownConfigRmTracking = types.StringNull()
	}

	if jsonData.SiteIdPolicyId != nil {
		if jsonData.SiteIdPolicyId.IsEmpty() {
			v.SiteIdPolicyId = types.Int64Null()
		} else {
			v.SiteIdPolicyId = types.Int64Value(int64(*jsonData.SiteIdPolicyId))
		}
	} else {
		v.SiteIdPolicyId = types.Int64Null()
	}

	if jsonData.FabricVpcDomainIdPrev != nil {
		if jsonData.FabricVpcDomainIdPrev.IsEmpty() {
			v.FabricVpcDomainIdPrev = types.Int64Null()
		} else {
			v.FabricVpcDomainIdPrev = types.Int64Value(int64(*jsonData.FabricVpcDomainIdPrev))
		}
	} else {
		v.FabricVpcDomainIdPrev = types.Int64Null()
	}

	if jsonData.LinkStateRoutingTagPrev != "" {
		v.LinkStateRoutingTagPrev = types.StringValue(jsonData.LinkStateRoutingTagPrev)
	} else {
		v.LinkStateRoutingTagPrev = types.StringNull()
	}

	if jsonData.BfdEnablePrev != "" {
		x, _ := strconv.ParseBool(jsonData.BfdEnablePrev)
		v.BfdEnablePrev = types.BoolValue(x)
	} else {
		v.BfdEnablePrev = types.BoolNull()
	}

	if jsonData.EnableSgtPrev != "" {
		x, _ := strconv.ParseBool(jsonData.EnableSgtPrev)
		v.EnableSgtPrev = types.BoolValue(x)
	} else {
		v.EnableSgtPrev = types.BoolNull()
	}

	if jsonData.SgtPreprovisionPrev != "" {
		x, _ := strconv.ParseBool(jsonData.SgtPreprovisionPrev)
		v.SgtPreprovisionPrev = types.BoolValue(x)
	} else {
		v.SgtPreprovisionPrev = types.BoolNull()
	}

	if jsonData.SgtPreprovRecalcStatus != "" {
		v.SgtPreprovRecalcStatus = types.StringValue(jsonData.SgtPreprovRecalcStatus)
	} else {
		v.SgtPreprovRecalcStatus = types.StringNull()
	}

	if jsonData.SgtRecalcStatus != "" {
		v.SgtRecalcStatus = types.StringValue(jsonData.SgtRecalcStatus)
	} else {
		v.SgtRecalcStatus = types.StringNull()
	}

	if jsonData.SgtOperStatus != "" {
		v.SgtOperStatus = types.StringValue(jsonData.SgtOperStatus)
	} else {
		v.SgtOperStatus = types.StringNull()
	}

	if jsonData.EnableMacsecPrev != "" {
		x, _ := strconv.ParseBool(jsonData.EnableMacsecPrev)
		v.EnableMacsecPrev = types.BoolValue(x)
	} else {
		v.EnableMacsecPrev = types.BoolNull()
	}

	if jsonData.EnableDciMacsecPrev != "" {
		x, _ := strconv.ParseBool(jsonData.EnableDciMacsecPrev)
		v.EnableDciMacsecPrev = types.BoolValue(x)
	} else {
		v.EnableDciMacsecPrev = types.BoolNull()
	}

	if jsonData.DciMacsecFallbackKeyString != "" {
		v.DciMacsecFallbackKeyString = types.StringValue(jsonData.DciMacsecFallbackKeyString)
	} else {
		v.DciMacsecFallbackKeyString = types.StringNull()
	}

	if jsonData.DciMacsecFallbackAlgorithm != "" {
		v.DciMacsecFallbackAlgorithm = types.StringValue(jsonData.DciMacsecFallbackAlgorithm)
	} else {
		v.DciMacsecFallbackAlgorithm = types.StringNull()
	}

	if jsonData.DciMacsecAlgorithm != "" {
		v.DciMacsecAlgorithm = types.StringValue(jsonData.DciMacsecAlgorithm)
	} else {
		v.DciMacsecAlgorithm = types.StringNull()
	}

	if jsonData.DciMacsecKeyString != "" {
		v.DciMacsecKeyString = types.StringValue(jsonData.DciMacsecKeyString)
	} else {
		v.DciMacsecKeyString = types.StringNull()
	}

	if jsonData.DciMacsecCipherSuite != "" {
		v.DciMacsecCipherSuite = types.StringValue(jsonData.DciMacsecCipherSuite)
	} else {
		v.DciMacsecCipherSuite = types.StringNull()
	}

	if jsonData.QkdProfileNamePrev != "" {
		x, _ := strconv.ParseBool(jsonData.QkdProfileNamePrev)
		v.QkdProfileNamePrev = types.BoolValue(x)
	} else {
		v.QkdProfileNamePrev = types.BoolNull()
	}

	if jsonData.FabricMtuPrev != nil {
		if jsonData.FabricMtuPrev.IsEmpty() {
			v.FabricMtuPrev = types.Int64Null()
		} else {
			v.FabricMtuPrev = types.Int64Value(int64(*jsonData.FabricMtuPrev))
		}
	} else {
		v.FabricMtuPrev = types.Int64Null()
	}

	if jsonData.L2HostIntfMtuPrev != nil {
		if jsonData.L2HostIntfMtuPrev.IsEmpty() {
			v.L2HostIntfMtuPrev = types.Int64Null()
		} else {
			v.L2HostIntfMtuPrev = types.Int64Value(int64(*jsonData.L2HostIntfMtuPrev))
		}
	} else {
		v.L2HostIntfMtuPrev = types.Int64Null()
	}

	if jsonData.MplsIsisAreaNumPrev != "" {
		v.MplsIsisAreaNumPrev = types.StringValue(jsonData.MplsIsisAreaNumPrev)
	} else {
		v.MplsIsisAreaNumPrev = types.StringNull()
	}

	if jsonData.IsisAreaNumPrev != "" {
		v.IsisAreaNumPrev = types.StringValue(jsonData.IsisAreaNumPrev)
	} else {
		v.IsisAreaNumPrev = types.StringNull()
	}

	if jsonData.EnableAiMlQosPolicyFlap != "" {
		x, _ := strconv.ParseBool(jsonData.EnableAiMlQosPolicyFlap)
		v.EnableAiMlQosPolicyFlap = types.BoolValue(x)
	} else {
		v.EnableAiMlQosPolicyFlap = types.BoolNull()
	}

	if jsonData.PfcWatchIntPrev != nil {
		if jsonData.PfcWatchIntPrev.IsEmpty() {
			v.PfcWatchIntPrev = types.Int64Null()
		} else {
			v.PfcWatchIntPrev = types.Int64Value(int64(*jsonData.PfcWatchIntPrev))
		}
	} else {
		v.PfcWatchIntPrev = types.Int64Null()
	}

	if jsonData.InbandMgmtPrev != "" {
		x, _ := strconv.ParseBool(jsonData.InbandMgmtPrev)
		v.InbandMgmtPrev = types.BoolValue(x)
	} else {
		v.InbandMgmtPrev = types.BoolNull()
	}

	if jsonData.BootstrapEnablePrev != "" {
		x, _ := strconv.ParseBool(jsonData.BootstrapEnablePrev)
		v.BootstrapEnablePrev = types.BoolValue(x)
	} else {
		v.BootstrapEnablePrev = types.BoolNull()
	}

	if jsonData.EnableNetflowPrev != "" {
		x, _ := strconv.ParseBool(jsonData.EnableNetflowPrev)
		v.EnableNetflowPrev = types.BoolValue(x)
	} else {
		v.EnableNetflowPrev = types.BoolNull()
	}

	if jsonData.AllowNxcPrev != "" {
		x, _ := strconv.ParseBool(jsonData.AllowNxcPrev)
		v.AllowNxcPrev = types.BoolValue(x)
	} else {
		v.AllowNxcPrev = types.BoolNull()
	}

	if jsonData.EnableNbmPassivePrev != "" {
		x, _ := strconv.ParseBool(jsonData.EnableNbmPassivePrev)
		v.EnableNbmPassivePrev = types.BoolValue(x)
	} else {
		v.EnableNbmPassivePrev = types.BoolNull()
	}

	if jsonData.FabricTechnology != "" {
		v.FabricTechnology = types.StringValue(jsonData.FabricTechnology)
	} else {
		v.FabricTechnology = types.StringNull()
	}

	if jsonData.InterfaceEthernetDefaultPolicy != "" {
		v.InterfaceEthernetDefaultPolicy = types.StringValue(jsonData.InterfaceEthernetDefaultPolicy)
	} else {
		v.InterfaceEthernetDefaultPolicy = types.StringNull()
	}

	if jsonData.InterfaceLoopbackDefaultPolicy != "" {
		v.InterfaceLoopbackDefaultPolicy = types.StringValue(jsonData.InterfaceLoopbackDefaultPolicy)
	} else {
		v.InterfaceLoopbackDefaultPolicy = types.StringNull()
	}

	if jsonData.InterfacePortChannelDefaultPolicy != "" {
		v.InterfacePortChannelDefaultPolicy = types.StringValue(jsonData.InterfacePortChannelDefaultPolicy)
	} else {
		v.InterfacePortChannelDefaultPolicy = types.StringNull()
	}

	if jsonData.InterfaceVlanDefaultPolicy != "" {
		v.InterfaceVlanDefaultPolicy = types.StringValue(jsonData.InterfaceVlanDefaultPolicy)
	} else {
		v.InterfaceVlanDefaultPolicy = types.StringNull()
	}

	if jsonData.RpIpRangeInternal != "" {
		v.RpIpRangeInternal = types.StringValue(jsonData.RpIpRangeInternal)
	} else {
		v.RpIpRangeInternal = types.StringNull()
	}

	if jsonData.InbandEnablePrev != "" {
		x, _ := strconv.ParseBool(jsonData.InbandEnablePrev)
		v.InbandEnablePrev = types.BoolValue(x)
	} else {
		v.InbandEnablePrev = types.BoolNull()
	}

	if jsonData.EnableAsm != "" {
		x, _ := strconv.ParseBool(jsonData.EnableAsm)
		v.EnableAsm = types.BoolValue(x)
	} else {
		v.EnableAsm = types.BoolNull()
	}

	if jsonData.DomainNameInternal != "" {
		v.DomainNameInternal = types.StringValue(jsonData.DomainNameInternal)
	} else {
		v.DomainNameInternal = types.StringNull()
	}

	if jsonData.PnpEnableInternal != "" {
		x, _ := strconv.ParseBool(jsonData.PnpEnableInternal)
		v.PnpEnableInternal = types.BoolValue(x)
	} else {
		v.PnpEnableInternal = types.BoolNull()
	}

	if jsonData.BgwRoutingTag != nil {
		if jsonData.BgwRoutingTag.IsEmpty() {
			v.BgwRoutingTag = types.Int64Null()
		} else {
			v.BgwRoutingTag = types.Int64Value(int64(*jsonData.BgwRoutingTag))
		}
	} else {
		v.BgwRoutingTag = types.Int64Null()
	}

	if jsonData.DcnmId != "" {
		v.DcnmId = types.StringValue(jsonData.DcnmId)
	} else {
		v.DcnmId = types.StringNull()
	}

	if jsonData.EnableTrmTrmv6 != "" {
		x, _ := strconv.ParseBool(jsonData.EnableTrmTrmv6)
		v.EnableTrmTrmv6 = types.BoolValue(x)
	} else {
		v.EnableTrmTrmv6 = types.BoolNull()
	}

	if jsonData.EnableTrmTrmv6Prev != "" {
		x, _ := strconv.ParseBool(jsonData.EnableTrmTrmv6Prev)
		v.EnableTrmTrmv6Prev = types.BoolValue(x)
	} else {
		v.EnableTrmTrmv6Prev = types.BoolNull()
	}

	if jsonData.Loopback100Ipv6Range != "" {
		v.Loopback100Ipv6Range = types.StringValue(jsonData.Loopback100Ipv6Range)
	} else {
		v.Loopback100Ipv6Range = types.StringNull()
	}

	if jsonData.BgwRoutingTagPrev != "" {
		v.BgwRoutingTagPrev = types.StringValue(jsonData.BgwRoutingTagPrev)
	} else {
		v.BgwRoutingTagPrev = types.StringNull()
	}

	if jsonData.MsIfcBgpAuthKeyType != nil {
		if jsonData.MsIfcBgpAuthKeyType.IsEmpty() {
			v.MsIfcBgpAuthKeyType = types.Int64Null()
		} else {
			v.MsIfcBgpAuthKeyType = types.Int64Value(int64(*jsonData.MsIfcBgpAuthKeyType))
		}
	} else {
		v.MsIfcBgpAuthKeyType = types.Int64Null()
	}

	if jsonData.MsIfcBgpAuthKeyTypePrev != nil {
		if jsonData.MsIfcBgpAuthKeyTypePrev.IsEmpty() {
			v.MsIfcBgpAuthKeyTypePrev = types.Int64Null()
		} else {
			v.MsIfcBgpAuthKeyTypePrev = types.Int64Value(int64(*jsonData.MsIfcBgpAuthKeyTypePrev))
		}
	} else {
		v.MsIfcBgpAuthKeyTypePrev = types.Int64Null()
	}

	if jsonData.MsIfcBgpPasswordEnablePrev != "" {
		x, _ := strconv.ParseBool(jsonData.MsIfcBgpPasswordEnablePrev)
		v.MsIfcBgpPasswordEnablePrev = types.BoolValue(x)
	} else {
		v.MsIfcBgpPasswordEnablePrev = types.BoolNull()
	}

	if jsonData.MsIfcBgpPasswordPrev != "" {
		v.MsIfcBgpPasswordPrev = types.StringValue(jsonData.MsIfcBgpPasswordPrev)
	} else {
		v.MsIfcBgpPasswordPrev = types.StringNull()
	}

	if jsonData.ParentOnemanageFabric != "" {
		v.ParentOnemanageFabric = types.StringValue(jsonData.ParentOnemanageFabric)
	} else {
		v.ParentOnemanageFabric = types.StringNull()
	}

	if jsonData.SgtIdRangePrev != "" {
		v.SgtIdRangePrev = types.StringValue(jsonData.SgtIdRangePrev)
	} else {
		v.SgtIdRangePrev = types.StringNull()
	}

	if jsonData.SgtNamePrefixPrev != "" {
		v.SgtNamePrefixPrev = types.StringValue(jsonData.SgtNamePrefixPrev)
	} else {
		v.SgtNamePrefixPrev = types.StringNull()
	}

	if jsonData.V6DciSubnetRange != "" {
		v.V6DciSubnetRange = types.StringValue(jsonData.V6DciSubnetRange)
	} else {
		v.V6DciSubnetRange = types.StringNull()
	}

	if jsonData.V6DciSubnetTargetMask != nil {
		if jsonData.V6DciSubnetTargetMask.IsEmpty() {
			v.V6DciSubnetTargetMask = types.Int64Null()
		} else {
			v.V6DciSubnetTargetMask = types.Int64Value(int64(*jsonData.V6DciSubnetTargetMask))
		}
	} else {
		v.V6DciSubnetTargetMask = types.Int64Null()
	}

	if jsonData.VxlanUnderlayIsV6 != "" {
		x, _ := strconv.ParseBool(jsonData.VxlanUnderlayIsV6)
		v.VxlanUnderlayIsV6 = types.BoolValue(x)
	} else {
		v.VxlanUnderlayIsV6 = types.BoolNull()
	}

	return err
}

func (v FabricModel) GetModelData() *NDFCFabricModel {
	var data = new(NDFCFabricModel)

	//MARSHAL_BODY

	if !v.FabricName.IsNull() && !v.FabricName.IsUnknown() {
		data.FabricName = v.FabricName.ValueString()
	} else {
		data.FabricName = ""
	}

	return data
}
