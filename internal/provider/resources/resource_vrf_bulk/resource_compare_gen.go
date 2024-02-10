package resource_vrf_bulk

import "log"

const (
	ValuesDeeplyEqual = iota
	RequiresReplace
	RequiresUpdate
	ControlFlagUpdate
)

func (v NDFCAttachListValue) DeepEqual(c NDFCAttachListValue) int {
	controlFlagUpdate := false
	if v.SerialNumber != c.SerialNumber {
		log.Printf("v.SerialNumber=%v, c.SerialNumber=%v", v.SerialNumber, c.SerialNumber)
		return RequiresReplace
	}

	if v.Vlan != nil && c.Vlan != nil {
		if *v.Vlan != *c.Vlan {
			log.Printf("v.Vlan=%v, c.Vlan=%v", *v.Vlan, *c.Vlan)
			return RequiresUpdate
		}
	} else {
		if v.Vlan != nil {
			log.Printf("v.Vlan=%v", *v.Vlan)
			return RequiresUpdate
		} else if c.Vlan != nil {
			log.Printf("c.Vlan=%v", *c.Vlan)
			return RequiresUpdate
		}
	}
	if v.FreeformConfig != c.FreeformConfig {
		log.Printf("v.FreeformConfig=%v, c.FreeformConfig=%v", v.FreeformConfig, c.FreeformConfig)
		return RequiresUpdate
	}
	if v.DeployThisAttachment != c.DeployThisAttachment {
		log.Printf("v.DeployThisAttachment=%v, c.DeployThisAttachment=%v", v.DeployThisAttachment, c.DeployThisAttachment)
		controlFlagUpdate = true
	}

	if v.InstanceValues.LoopbackId != nil && c.InstanceValues.LoopbackId != nil {
		if *v.InstanceValues.LoopbackId != *c.InstanceValues.LoopbackId {
			log.Printf("v.InstanceValues.LoopbackId=%v, c.InstanceValues.LoopbackId=%v", *v.InstanceValues.LoopbackId, *c.InstanceValues.LoopbackId)
			return RequiresUpdate
		}
	} else {
		if v.InstanceValues.LoopbackId != nil {
			log.Printf("v.InstanceValues.LoopbackId=%v", *v.InstanceValues.LoopbackId)
			return RequiresUpdate
		} else if c.InstanceValues.LoopbackId != nil {
			log.Printf("c.InstanceValues.LoopbackId=%v", *c.InstanceValues.LoopbackId)
			return RequiresUpdate
		}
	}
	if v.InstanceValues.LoopbackIpv4 != c.InstanceValues.LoopbackIpv4 {
		log.Printf("v.InstanceValues.LoopbackIpv4=%s, c.InstanceValues.LoopbackIpv4=%s", v.InstanceValues.LoopbackIpv4, c.InstanceValues.LoopbackIpv4)
		return RequiresUpdate
	}
	if v.InstanceValues.LoopbackIpv6 != c.InstanceValues.LoopbackIpv6 {
		log.Printf("v.InstanceValues.LoopbackIpv6=%s, c.InstanceValues.LoopbackIpv6=%s", v.InstanceValues.LoopbackIpv6, c.InstanceValues.LoopbackIpv6)
		return RequiresUpdate
	}

	if controlFlagUpdate {
		return ControlFlagUpdate
	}
	return ValuesDeeplyEqual
}

func (v *NDFCVrfsValue) CreateSearchMap() {
	v.AttachListMap = make(map[string]*NDFCAttachListValue)
	for i := range v.AttachList {
		key := ""
		if v.AttachList[i].SerialNumber == "" {
			key = v.AttachList[i].SwitchSerialNo
		} else {
			key = v.AttachList[i].SerialNumber
		}
		log.Printf("NDFCVrfsValue.CreateSearchMap: key=%s", key)
		v.AttachListMap[key] = &v.AttachList[i]
	}
}

func (v NDFCVrfsValue) DeepEqual(c NDFCVrfsValue) int {
	controlFlagUpdate := false
	if v.VrfName != c.VrfName {
		log.Printf("v.VrfName=%v, c.VrfName=%v", v.VrfName, c.VrfName)
		return RequiresReplace
	}
	if v.VrfTemplate != c.VrfTemplate {
		log.Printf("v.VrfTemplate=%v, c.VrfTemplate=%v", v.VrfTemplate, c.VrfTemplate)
		return RequiresUpdate
	}
	if v.VrfExtensionTemplate != c.VrfExtensionTemplate {
		log.Printf("v.VrfExtensionTemplate=%v, c.VrfExtensionTemplate=%v", v.VrfExtensionTemplate, c.VrfExtensionTemplate)
		return RequiresUpdate
	}

	if v.VrfId != nil && c.VrfId != nil {
		if *v.VrfId != *c.VrfId {
			log.Printf("v.VrfId=%v, c.VrfId=%v", *v.VrfId, *c.VrfId)
			return RequiresReplace
		}
	} else {
		if v.VrfId != nil {
			log.Printf("v.VrfId=%v", *v.VrfId)
			return RequiresReplace
		} else if c.VrfId != nil {
			log.Printf("c.VrfId=%v", *c.VrfId)
			return RequiresReplace
		}
	}

	if v.VrfTemplateConfig.VlanId != nil && c.VrfTemplateConfig.VlanId != nil {
		if *v.VrfTemplateConfig.VlanId != *c.VrfTemplateConfig.VlanId {
			log.Printf("v.VrfTemplateConfig.VlanId=%v, c.VrfTemplateConfig.VlanId=%v", *v.VrfTemplateConfig.VlanId, *c.VrfTemplateConfig.VlanId)
			return RequiresUpdate
		}
	} else {
		if v.VrfTemplateConfig.VlanId != nil {
			log.Printf("v.VrfTemplateConfig.VlanId=%v", *v.VrfTemplateConfig.VlanId)
			return RequiresUpdate
		} else if c.VrfTemplateConfig.VlanId != nil {
			log.Printf("c.VrfTemplateConfig.VlanId=%v", *c.VrfTemplateConfig.VlanId)
			return RequiresUpdate
		}
	}
	if v.VrfTemplateConfig.VlanName != c.VrfTemplateConfig.VlanName {
		log.Printf("v.VrfTemplateConfig.VlanName=%s, c.VrfTemplateConfig.VlanName=%s", v.VrfTemplateConfig.VlanName, c.VrfTemplateConfig.VlanName)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.InterfaceDescription != c.VrfTemplateConfig.InterfaceDescription {
		log.Printf("v.VrfTemplateConfig.InterfaceDescription=%s, c.VrfTemplateConfig.InterfaceDescription=%s", v.VrfTemplateConfig.InterfaceDescription, c.VrfTemplateConfig.InterfaceDescription)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.VrfDescription != c.VrfTemplateConfig.VrfDescription {
		log.Printf("v.VrfTemplateConfig.VrfDescription=%s, c.VrfTemplateConfig.VrfDescription=%s", v.VrfTemplateConfig.VrfDescription, c.VrfTemplateConfig.VrfDescription)
		return RequiresUpdate
	}

	if v.VrfTemplateConfig.Mtu != nil && c.VrfTemplateConfig.Mtu != nil {
		if *v.VrfTemplateConfig.Mtu != *c.VrfTemplateConfig.Mtu {
			log.Printf("v.VrfTemplateConfig.Mtu=%v, c.VrfTemplateConfig.Mtu=%v", *v.VrfTemplateConfig.Mtu, *c.VrfTemplateConfig.Mtu)
			return RequiresUpdate
		}
	} else {
		if v.VrfTemplateConfig.Mtu != nil {
			log.Printf("v.VrfTemplateConfig.Mtu=%v", *v.VrfTemplateConfig.Mtu)
			return RequiresUpdate
		} else if c.VrfTemplateConfig.Mtu != nil {
			log.Printf("c.VrfTemplateConfig.Mtu=%v", *c.VrfTemplateConfig.Mtu)
			return RequiresUpdate
		}
	}

	if v.VrfTemplateConfig.LoopbackRoutingTag != nil && c.VrfTemplateConfig.LoopbackRoutingTag != nil {
		if *v.VrfTemplateConfig.LoopbackRoutingTag != *c.VrfTemplateConfig.LoopbackRoutingTag {
			log.Printf("v.VrfTemplateConfig.LoopbackRoutingTag=%v, c.VrfTemplateConfig.LoopbackRoutingTag=%v", *v.VrfTemplateConfig.LoopbackRoutingTag, *c.VrfTemplateConfig.LoopbackRoutingTag)
			return RequiresUpdate
		}
	} else {
		if v.VrfTemplateConfig.LoopbackRoutingTag != nil {
			log.Printf("v.VrfTemplateConfig.LoopbackRoutingTag=%v", *v.VrfTemplateConfig.LoopbackRoutingTag)
			return RequiresUpdate
		} else if c.VrfTemplateConfig.LoopbackRoutingTag != nil {
			log.Printf("c.VrfTemplateConfig.LoopbackRoutingTag=%v", *c.VrfTemplateConfig.LoopbackRoutingTag)
			return RequiresUpdate
		}
	}
	if v.VrfTemplateConfig.RedistributeDirectRouteMap != c.VrfTemplateConfig.RedistributeDirectRouteMap {
		log.Printf("v.VrfTemplateConfig.RedistributeDirectRouteMap=%s, c.VrfTemplateConfig.RedistributeDirectRouteMap=%s", v.VrfTemplateConfig.RedistributeDirectRouteMap, c.VrfTemplateConfig.RedistributeDirectRouteMap)
		return RequiresUpdate
	}

	if v.VrfTemplateConfig.MaxBgpPaths != nil && c.VrfTemplateConfig.MaxBgpPaths != nil {
		if *v.VrfTemplateConfig.MaxBgpPaths != *c.VrfTemplateConfig.MaxBgpPaths {
			log.Printf("v.VrfTemplateConfig.MaxBgpPaths=%v, c.VrfTemplateConfig.MaxBgpPaths=%v", *v.VrfTemplateConfig.MaxBgpPaths, *c.VrfTemplateConfig.MaxBgpPaths)
			return RequiresUpdate
		}
	} else {
		if v.VrfTemplateConfig.MaxBgpPaths != nil {
			log.Printf("v.VrfTemplateConfig.MaxBgpPaths=%v", *v.VrfTemplateConfig.MaxBgpPaths)
			return RequiresUpdate
		} else if c.VrfTemplateConfig.MaxBgpPaths != nil {
			log.Printf("c.VrfTemplateConfig.MaxBgpPaths=%v", *c.VrfTemplateConfig.MaxBgpPaths)
			return RequiresUpdate
		}
	}

	if v.VrfTemplateConfig.MaxIbgpPaths != nil && c.VrfTemplateConfig.MaxIbgpPaths != nil {
		if *v.VrfTemplateConfig.MaxIbgpPaths != *c.VrfTemplateConfig.MaxIbgpPaths {
			log.Printf("v.VrfTemplateConfig.MaxIbgpPaths=%v, c.VrfTemplateConfig.MaxIbgpPaths=%v", *v.VrfTemplateConfig.MaxIbgpPaths, *c.VrfTemplateConfig.MaxIbgpPaths)
			return RequiresUpdate
		}
	} else {
		if v.VrfTemplateConfig.MaxIbgpPaths != nil {
			log.Printf("v.VrfTemplateConfig.MaxIbgpPaths=%v", *v.VrfTemplateConfig.MaxIbgpPaths)
			return RequiresUpdate
		} else if c.VrfTemplateConfig.MaxIbgpPaths != nil {
			log.Printf("c.VrfTemplateConfig.MaxIbgpPaths=%v", *c.VrfTemplateConfig.MaxIbgpPaths)
			return RequiresUpdate
		}
	}
	if v.VrfTemplateConfig.Ipv6LinkLocal != c.VrfTemplateConfig.Ipv6LinkLocal {
		log.Printf("v.VrfTemplateConfig.Ipv6LinkLocal=%s, c.VrfTemplateConfig.Ipv6LinkLocal=%s", v.VrfTemplateConfig.Ipv6LinkLocal, c.VrfTemplateConfig.Ipv6LinkLocal)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.Trm != c.VrfTemplateConfig.Trm {
		log.Printf("v.VrfTemplateConfig.Trm=%s, c.VrfTemplateConfig.Trm=%s", v.VrfTemplateConfig.Trm, c.VrfTemplateConfig.Trm)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.NoRp != c.VrfTemplateConfig.NoRp {
		log.Printf("v.VrfTemplateConfig.NoRp=%s, c.VrfTemplateConfig.NoRp=%s", v.VrfTemplateConfig.NoRp, c.VrfTemplateConfig.NoRp)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.RpExternal != c.VrfTemplateConfig.RpExternal {
		log.Printf("v.VrfTemplateConfig.RpExternal=%s, c.VrfTemplateConfig.RpExternal=%s", v.VrfTemplateConfig.RpExternal, c.VrfTemplateConfig.RpExternal)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.RpAddress != c.VrfTemplateConfig.RpAddress {
		log.Printf("v.VrfTemplateConfig.RpAddress=%s, c.VrfTemplateConfig.RpAddress=%s", v.VrfTemplateConfig.RpAddress, c.VrfTemplateConfig.RpAddress)
		return RequiresUpdate
	}

	if v.VrfTemplateConfig.RpLoopbackId != nil && c.VrfTemplateConfig.RpLoopbackId != nil {
		if *v.VrfTemplateConfig.RpLoopbackId != *c.VrfTemplateConfig.RpLoopbackId {
			log.Printf("v.VrfTemplateConfig.RpLoopbackId=%v, c.VrfTemplateConfig.RpLoopbackId=%v", *v.VrfTemplateConfig.RpLoopbackId, *c.VrfTemplateConfig.RpLoopbackId)
			return RequiresUpdate
		}
	} else {
		if v.VrfTemplateConfig.RpLoopbackId != nil {
			log.Printf("v.VrfTemplateConfig.RpLoopbackId=%v", *v.VrfTemplateConfig.RpLoopbackId)
			return RequiresUpdate
		} else if c.VrfTemplateConfig.RpLoopbackId != nil {
			log.Printf("c.VrfTemplateConfig.RpLoopbackId=%v", *c.VrfTemplateConfig.RpLoopbackId)
			return RequiresUpdate
		}
	}
	if v.VrfTemplateConfig.UnderlayMulticastAddress != c.VrfTemplateConfig.UnderlayMulticastAddress {
		log.Printf("v.VrfTemplateConfig.UnderlayMulticastAddress=%s, c.VrfTemplateConfig.UnderlayMulticastAddress=%s", v.VrfTemplateConfig.UnderlayMulticastAddress, c.VrfTemplateConfig.UnderlayMulticastAddress)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.OverlayMulticastGroups != c.VrfTemplateConfig.OverlayMulticastGroups {
		log.Printf("v.VrfTemplateConfig.OverlayMulticastGroups=%s, c.VrfTemplateConfig.OverlayMulticastGroups=%s", v.VrfTemplateConfig.OverlayMulticastGroups, c.VrfTemplateConfig.OverlayMulticastGroups)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.MvpnInterAs != c.VrfTemplateConfig.MvpnInterAs {
		log.Printf("v.VrfTemplateConfig.MvpnInterAs=%s, c.VrfTemplateConfig.MvpnInterAs=%s", v.VrfTemplateConfig.MvpnInterAs, c.VrfTemplateConfig.MvpnInterAs)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.TrmBgwMsite != c.VrfTemplateConfig.TrmBgwMsite {
		log.Printf("v.VrfTemplateConfig.TrmBgwMsite=%s, c.VrfTemplateConfig.TrmBgwMsite=%s", v.VrfTemplateConfig.TrmBgwMsite, c.VrfTemplateConfig.TrmBgwMsite)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.AdvertiseHostRoutes != c.VrfTemplateConfig.AdvertiseHostRoutes {
		log.Printf("v.VrfTemplateConfig.AdvertiseHostRoutes=%s, c.VrfTemplateConfig.AdvertiseHostRoutes=%s", v.VrfTemplateConfig.AdvertiseHostRoutes, c.VrfTemplateConfig.AdvertiseHostRoutes)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.AdvertiseDefaultRoute != c.VrfTemplateConfig.AdvertiseDefaultRoute {
		log.Printf("v.VrfTemplateConfig.AdvertiseDefaultRoute=%s, c.VrfTemplateConfig.AdvertiseDefaultRoute=%s", v.VrfTemplateConfig.AdvertiseDefaultRoute, c.VrfTemplateConfig.AdvertiseDefaultRoute)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.ConfigureStaticDefaultRoute != c.VrfTemplateConfig.ConfigureStaticDefaultRoute {
		log.Printf("v.VrfTemplateConfig.ConfigureStaticDefaultRoute=%s, c.VrfTemplateConfig.ConfigureStaticDefaultRoute=%s", v.VrfTemplateConfig.ConfigureStaticDefaultRoute, c.VrfTemplateConfig.ConfigureStaticDefaultRoute)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.BgpPassword != c.VrfTemplateConfig.BgpPassword {
		log.Printf("v.VrfTemplateConfig.BgpPassword=%s, c.VrfTemplateConfig.BgpPassword=%s", v.VrfTemplateConfig.BgpPassword, c.VrfTemplateConfig.BgpPassword)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.BgpPasswordType != c.VrfTemplateConfig.BgpPasswordType {
		log.Printf("v.VrfTemplateConfig.BgpPasswordType=%s, c.VrfTemplateConfig.BgpPasswordType=%s", v.VrfTemplateConfig.BgpPasswordType, c.VrfTemplateConfig.BgpPasswordType)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.Netflow != c.VrfTemplateConfig.Netflow {
		log.Printf("v.VrfTemplateConfig.Netflow=%s, c.VrfTemplateConfig.Netflow=%s", v.VrfTemplateConfig.Netflow, c.VrfTemplateConfig.Netflow)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.NetflowMonitor != c.VrfTemplateConfig.NetflowMonitor {
		log.Printf("v.VrfTemplateConfig.NetflowMonitor=%s, c.VrfTemplateConfig.NetflowMonitor=%s", v.VrfTemplateConfig.NetflowMonitor, c.VrfTemplateConfig.NetflowMonitor)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.DisableRtAuto != c.VrfTemplateConfig.DisableRtAuto {
		log.Printf("v.VrfTemplateConfig.DisableRtAuto=%s, c.VrfTemplateConfig.DisableRtAuto=%s", v.VrfTemplateConfig.DisableRtAuto, c.VrfTemplateConfig.DisableRtAuto)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.RouteTargetImport != c.VrfTemplateConfig.RouteTargetImport {
		log.Printf("v.VrfTemplateConfig.RouteTargetImport=%s, c.VrfTemplateConfig.RouteTargetImport=%s", v.VrfTemplateConfig.RouteTargetImport, c.VrfTemplateConfig.RouteTargetImport)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.RouteTargetExport != c.VrfTemplateConfig.RouteTargetExport {
		log.Printf("v.VrfTemplateConfig.RouteTargetExport=%s, c.VrfTemplateConfig.RouteTargetExport=%s", v.VrfTemplateConfig.RouteTargetExport, c.VrfTemplateConfig.RouteTargetExport)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.RouteTargetImportEvpn != c.VrfTemplateConfig.RouteTargetImportEvpn {
		log.Printf("v.VrfTemplateConfig.RouteTargetImportEvpn=%s, c.VrfTemplateConfig.RouteTargetImportEvpn=%s", v.VrfTemplateConfig.RouteTargetImportEvpn, c.VrfTemplateConfig.RouteTargetImportEvpn)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.RouteTargetExportEvpn != c.VrfTemplateConfig.RouteTargetExportEvpn {
		log.Printf("v.VrfTemplateConfig.RouteTargetExportEvpn=%s, c.VrfTemplateConfig.RouteTargetExportEvpn=%s", v.VrfTemplateConfig.RouteTargetExportEvpn, c.VrfTemplateConfig.RouteTargetExportEvpn)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.RouteTargetImportMvpn != c.VrfTemplateConfig.RouteTargetImportMvpn {
		log.Printf("v.VrfTemplateConfig.RouteTargetImportMvpn=%s, c.VrfTemplateConfig.RouteTargetImportMvpn=%s", v.VrfTemplateConfig.RouteTargetImportMvpn, c.VrfTemplateConfig.RouteTargetImportMvpn)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.RouteTargetExportMvpn != c.VrfTemplateConfig.RouteTargetExportMvpn {
		log.Printf("v.VrfTemplateConfig.RouteTargetExportMvpn=%s, c.VrfTemplateConfig.RouteTargetExportMvpn=%s", v.VrfTemplateConfig.RouteTargetExportMvpn, c.VrfTemplateConfig.RouteTargetExportMvpn)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.RouteTargetImportCloudEvpn != c.VrfTemplateConfig.RouteTargetImportCloudEvpn {
		log.Printf("v.VrfTemplateConfig.RouteTargetImportCloudEvpn=%s, c.VrfTemplateConfig.RouteTargetImportCloudEvpn=%s", v.VrfTemplateConfig.RouteTargetImportCloudEvpn, c.VrfTemplateConfig.RouteTargetImportCloudEvpn)
		return RequiresUpdate
	}
	if v.VrfTemplateConfig.RouteTargetExportCloudEvpn != c.VrfTemplateConfig.RouteTargetExportCloudEvpn {
		log.Printf("v.VrfTemplateConfig.RouteTargetExportCloudEvpn=%s, c.VrfTemplateConfig.RouteTargetExportCloudEvpn=%s", v.VrfTemplateConfig.RouteTargetExportCloudEvpn, c.VrfTemplateConfig.RouteTargetExportCloudEvpn)
		return RequiresUpdate
	}
	if v.DeployAttachments != c.DeployAttachments {
		log.Printf("v.DeployAttachments=%v, c.DeployAttachments=%v", v.DeployAttachments, c.DeployAttachments)
		controlFlagUpdate = true
	}

	if len(v.AttachList) != len(c.AttachList) {
		log.Printf("len(v.AttachList)=%d, len(c.AttachList)=%d", len(v.AttachList), len(c.AttachList))
		return RequiresUpdate
	}
	for i := range v.AttachList {
		retVal := v.AttachList[i].DeepEqual(c.AttachList[i])
		if retVal != ValuesDeeplyEqual {
			return retVal
		}
	}
	if controlFlagUpdate {
		return ControlFlagUpdate
	}
	return ValuesDeeplyEqual
}

func (v *NDFCVrfBulkModel) CreateSearchMap() {
	v.VrfsMap = make(map[string]*NDFCVrfsValue)
	for i := range v.Vrfs {
		key := ""
		key = v.Vrfs[i].VrfName
		log.Printf("NDFCVrfBulkModel.CreateSearchMap: key=%s", key)
		v.VrfsMap[key] = &v.Vrfs[i]
		v.Vrfs[i].CreateSearchMap()
	}
}
