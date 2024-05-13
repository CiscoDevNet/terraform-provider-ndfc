// Code generated DO NOT EDIT.
package resource_interface_common

type NDFCInterfaceCommonModel struct {
	Policy       string                         `json:"policy,omitempty"`
	Deploy       bool                           `json:"-"`
	SerialNumber string                         `json:"-"`
	Interfaces   map[string]NDFCInterfacesValue `json:"interfaces,omitempty"`
}

type NDFCInterfacesValue struct {
	FilterThisValue bool             `json:"-"`
	SerialNumber    string           `json:"serialNumber,omitempty"`
	InterfaceName   string           `json:"ifName,omitempty"`
	InterfaceType   string           `json:"interfaceType,omitempty"`
	NvPairs         NDFCNvPairsValue `json:"nvPairs,omitempty"`
}

type NDFCNvPairsValue struct {
	InterfaceName        string `json:"INTF_NAME,omitempty"`
	AdminState           string `json:"ADMIN_STATE,omitempty"`
	FreeformConfig       string `json:"CONF,omitempty"`
	InterfaceDescription string `json:"DESC,omitempty"`
	Vrf                  string `json:"INTF_VRF,omitempty"`
	Ipv4Address          string `json:"IP,omitempty"`
	Ipv6Address          string `json:"V6IP,omitempty"`
	RouteMapTag          string `json:"ROUTE_MAP_TAG,omitempty"`
	BpduGuard            string `json:"BPDUGUARD_ENABLED,omitempty"`
	PortTypeFast         string `json:"PORTTYPE_FAST_ENABLED,omitempty"`
	Mtu                  string `json:"MTU,omitempty"`
	Speed                string `json:"SPEED,omitempty"`
	AccessVlan           *int64 `json:"ACCESS_VLAN,string,omitempty"`
	OrphanPort           string `json:"ENABLE_ORPHAN_PORT,omitempty"`
	Ptp                  string `json:"PTP,omitempty"`
	Netflow              string `json:"ENABLE_NETFLOW,omitempty"`
	NetflowMonitor       string `json:"NETFLOW_MONITOR,omitempty"`
	NetflowSampler       string `json:"NETFLOW_SAMPLER,omitempty"`
	AllowedVlans         string `json:"ALLOWED_VLANS,omitempty"`
	NativeVlan           *int64 `json:"NATIVE_VLAN,string,omitempty"`
}
