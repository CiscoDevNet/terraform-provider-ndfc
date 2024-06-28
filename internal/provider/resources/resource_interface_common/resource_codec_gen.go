// Code generated DO NOT EDIT.
package resource_interface_common

import (
	. "terraform-provider-ndfc/internal/provider/types"
)

type NDFCInterfaceCommonModel struct {
	Policy       string                         `json:"policy,omitempty"`
	Deploy       bool                           `json:"-"`
	SerialNumber string                         `json:"-"`
	Interfaces   map[string]NDFCInterfacesValue `json:"interfaces,omitempty"`
}

type NDFCInterfacesValue struct {
	FilterThisValue  bool             `json:"-"`
	SerialNumber     string           `json:"serialNumber,omitempty"`
	InterfaceName    string           `json:"ifName,omitempty"`
	InterfaceType    string           `json:"interfaceType,omitempty"`
	DeploymentStatus string           `json:"-"`
	NvPairs          NDFCNvPairsValue `json:"nvPairs,omitempty"`
}

type NDFCNvPairsValue struct {
	InterfaceName             string       `json:"INTF_NAME,omitempty"`
	AdminState                string       `json:"ADMIN_STATE,omitempty"`
	FreeformConfig            string       `json:"CONF"`
	InterfaceDescription      string       `json:"DESC"`
	Vrf                       string       `json:"INTF_VRF,omitempty"`
	Ipv4Address               string       `json:"IP,omitempty"`
	Ipv6Address               string       `json:"V6IP,omitempty"`
	RouteMapTag               string       `json:"ROUTE_MAP_TAG,omitempty"`
	BpduGuard                 string       `json:"BPDUGUARD_ENABLED,omitempty"`
	PortTypeFast              string       `json:"PORTTYPE_FAST_ENABLED,omitempty"`
	Mtu                       string       `json:"MTU"`
	Speed                     string       `json:"SPEED,omitempty"`
	AccessVlan                *Int64Custom `json:"ACCESS_VLAN,omitempty"`
	OrphanPort                string       `json:"ENABLE_ORPHAN_PORT,omitempty"`
	Ptp                       string       `json:"PTP,omitempty"`
	Netflow                   string       `json:"ENABLE_NETFLOW,omitempty"`
	NetflowMonitor            string       `json:"NETFLOW_MONITOR,omitempty"`
	NetflowSampler            string       `json:"NETFLOW_SAMPLER,omitempty"`
	AllowedVlans              string       `json:"ALLOWED_VLANS,omitempty"`
	NativeVlan                *Int64Custom `json:"NATIVE_VLAN,omitempty"`
	Ipv4PrefixLength          string       `json:"PREFIX,omitempty"`
	RoutingTag                string       `json:"ROUTING_TAG"`
	DisableIpRedirects        string       `json:"DISABLE_IP_REDIRECTS,omitempty"`
	EnableHsrp                string       `json:"ENABLE_HSRP,omitempty"`
	HsrpGroup                 *Int64Custom `json:"HSRP_GROUP,omitempty"`
	HsrpVip                   string       `json:"HSRP_VIP,omitempty"`
	HsrpPriority              *Int64Custom `json:"HSRP_PRIORITY,omitempty"`
	HsrpVersion               string       `json:"HSRP_VERSION,omitempty"`
	Preempt                   string       `json:"PREEMPT,omitempty"`
	Mac                       string       `json:"MAC,omitempty"`
	DhcpServerAddr1           string       `json:"dhcpServerAddr1,omitempty"`
	DhcpServerAddr2           string       `json:"dhcpServerAddr2,omitempty"`
	DhcpServerAddr3           string       `json:"dhcpServerAddr3,omitempty"`
	VrfDhcp1                  string       `json:"vrfDhcp1,omitempty"`
	VrfDhcp2                  string       `json:"vrfDhcp2,omitempty"`
	VrfDhcp3                  string       `json:"vrfDhcp3,omitempty"`
	AdvertiseSubnetInUnderlay string       `json:"advSubnetInUnderlay,omitempty"`
	PortChannelName           string       `json:"PO_ID,omitempty"`
	CopyPoDescription         string       `json:"COPY_DESC,omitempty"`
	PortchannelMode           string       `json:"PC_MODE,omitempty"`
	MemberInterfaces          string       `json:"MEMBER_INTERFACES,omitempty"`
	MirrorConfig              string       `json:"ENABLE_MIRROR_CONFIG,omitempty"`
	Peer1PoFreeformConfig     string       `json:"PEER1_PO_CONF"`
	Peer2PoFreeformConfig     string       `json:"PEER2_PO_CONF"`
	Peer1PoDescription        string       `json:"PEER1_PO_DESC"`
	Peer2PoDescription        string       `json:"PEER2_PO_DESC"`
	Peer1AllowedVlans         string       `json:"PEER1_ALLOWED_VLANS,omitempty"`
	Peer2AllowedVlans         string       `json:"PEER2_ALLOWED_VLANS,omitempty"`
	Peer1NativeVlan           *Int64Custom `json:"PEER1_NATIVE_VLAN,omitempty"`
	Peer2NativeVlan           *Int64Custom `json:"PEER2_NATIVE_VLAN,omitempty"`
	Peer1MemberInterfaces     string       `json:"PEER1_MEMBER_INTERFACES,omitempty"`
	Peer2MemberInterfaces     string       `json:"PEER2_MEMBER_INTERFACES,omitempty"`
	Peer1PortChannelId        *int64       `json:"PEER1_PCID,string,omitempty"`
	Peer2PortChannelId        *Int64Custom `json:"PEER2_PCID,string,omitempty"`
}
