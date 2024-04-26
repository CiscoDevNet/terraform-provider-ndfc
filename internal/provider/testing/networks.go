package testing

import (
	"log"
	"strconv"
	"strings"
	rna "terraform-provider-ndfc/internal/provider/resources/resource_network_attachments"
	"terraform-provider-ndfc/internal/provider/resources/resource_networks"
	"terraform-provider-ndfc/internal/provider/types"
)

//var tmpDir string

func GenerateNetworksObject(bulk **resource_networks.NDFCNetworksModel, fabric string, count int,
	globaldeploy, net_deploy, deployNeeded bool, vrf string, serials []string) {
	nets := new(resource_networks.NDFCNetworksModel)
	nets.FabricName = fabric
	nets.DeployAllAttachments = globaldeploy
	nets.Networks = make(map[string]resource_networks.NDFCNetworksValue)
	log.Printf("Creating Bulk Network object Network Count: %d", count)

	for i := 0; i < count; i++ {
		nwName := GetConfig().NDFC.NetPrefix + strconv.Itoa(i+1)
		log.Printf("Creating Network: %s", nwName)
		nw := resource_networks.NDFCNetworksValue{}
		//nw.VrfTemplateConfig.VrfDescription = "Network Description"
		//nw.NetworkName = nwName
		nw.DeployAttachments = net_deploy
		nw.VrfName = vrf
		if len(serials) > 0 {
			nw.Attachments = make(map[string]rna.NDFCAttachmentsValue)

			for j := range serials {
				attach := rna.NDFCAttachmentsValue{}
				attach.SerialNumber = serials[j]
				attach.DeployThisAttachment = deployNeeded
				nw.Attachments[serials[j]] = attach
			}
		} else {
			nw.Attachments = nil
		}
		nets.Networks[nwName] = nw
	}
	*bulk = nets
}

func GenerateSingleNetworkObject(nwDptr **resource_networks.NDFCNetworksModel, namePrefix, fabric string, nwNo int, globaldeploy, net_deploy, deployNeeded bool, serials []string) {
	nets := new(resource_networks.NDFCNetworksModel)
	nets.FabricName = fabric
	nets.DeployAllAttachments = globaldeploy
	nets.Networks = make(map[string]resource_networks.NDFCNetworksValue)
	nwName := namePrefix + strconv.Itoa(nwNo)
	nw := resource_networks.NDFCNetworksValue{}

	nw.DeployAttachments = net_deploy
	if len(serials) > 0 {
		nw.Attachments = make(map[string]rna.NDFCAttachmentsValue)
		for j := range serials {
			attach := rna.NDFCAttachmentsValue{}
			attach.SerialNumber = serials[j]
			attach.DeployThisAttachment = deployNeeded
			nw.Attachments[serials[j]] = attach
		}
	} else {
		nw.Attachments = nil
	}
	nets.Networks[nwName] = nw
	*nwDptr = nets
}

func ModifyNetworksObject(nws **resource_networks.NDFCNetworksModel, nwNo int, values map[string]interface{}) {
	nets := *nws
	nwName := GetConfig().NDFC.NetPrefix + strconv.Itoa(nwNo)
	nw, ok := nets.Networks[nwName]
	if ok {
		for key, value := range values {
			switch key {

			case "vlan_id":
				nw.NetworkTemplateConfig.VlanId = new(types.Int64Custom)
				*nw.NetworkTemplateConfig.VlanId = types.Int64Custom((value.(int)))
			case "routing_tag":
				nw.NetworkTemplateConfig.RoutingTag = new(types.Int64Custom)
				*nw.NetworkTemplateConfig.RoutingTag = types.Int64Custom((value.(int)))
			case "mtu":
				nw.NetworkTemplateConfig.Mtu = new(types.Int64Custom)
				*nw.NetworkTemplateConfig.Mtu = types.Int64Custom((value.(int)))
			case "secondary_gateway_1":
				nw.NetworkTemplateConfig.SecondaryGateway1 = value.(string)
			case "secondary_gateway_2":
				nw.NetworkTemplateConfig.SecondaryGateway2 = value.(string)
			case "secondary_gateway_3":
				nw.NetworkTemplateConfig.SecondaryGateway3 = value.(string)
			case "secondary_gateway_4":
				nw.NetworkTemplateConfig.SecondaryGateway4 = value.(string)
			case "arp_suppression":
				nw.NetworkTemplateConfig.ArpSuppression = value.(string)
			case "ingress_replication":
				nw.NetworkTemplateConfig.IngressReplication = value.(string)
			case "multicast_group":
				nw.NetworkTemplateConfig.MulticastGroup = value.(string)
			case "dhcp_relay_loopback_id":
				nw.NetworkTemplateConfig.DhcpRelayLoopbackId = new(types.Int64Custom)
				*nw.NetworkTemplateConfig.DhcpRelayLoopbackId = types.Int64Custom((value.(int)))
			case "trm":
				nw.NetworkTemplateConfig.Trm = value.(string)
			case "route_target_both":
				nw.NetworkTemplateConfig.RouteTargetBoth = value.(string)
			case "netflow":
				nw.NetworkTemplateConfig.Netflow = value.(string)
			case "nw_id":
				nw.NetworkId = new(int64)
				*nw.NetworkId = int64(value.(int))
			}
		}
		nets.Networks[nwName] = nw
	}

}

func IncreaseNetCount(nw **resource_networks.NDFCNetworksModel, nwToAdd int,
	globaldeploy, net_deploy, deployNeeded bool, nwName string, serials []string) {
	nets := *nw
	nets.DeployAllAttachments = globaldeploy
	log.Printf("Add more Networks  Network Count: %d", nwToAdd)

	currentCount := len(nets.Networks)
	for i := 0; i < nwToAdd; i++ {
		nwName := GetConfig().NDFC.NetPrefix + strconv.Itoa(currentCount+i+1)
		log.Printf("Creating Network: %s", nwName)
		nw := resource_networks.NDFCNetworksValue{}
		//nw.VrfTemplateConfig.VrfDescription = "Network Description"
		//nw.NetworkName = nwName
		nw.DeployAttachments = net_deploy
		nw.VrfName = nwName
		if len(serials) > 0 {
			nw.Attachments = make(map[string]rna.NDFCAttachmentsValue)
			for j := range serials {
				attach := rna.NDFCAttachmentsValue{}
				attach.SerialNumber = serials[j]
				attach.DeployThisAttachment = deployNeeded
				nw.Attachments[serials[j]] = attach
			}
		}
		nets.Networks[nwName] = nw
	}
}

func AddNetAttachments(nets *resource_networks.NDFCNetworksModel, serials []string, deployNeeded bool, start, end int) *resource_networks.NDFCNetworksModel {
	for nwName, nw := range nets.Networks {
		//nw_acc_<id>

		ids := strings.Split(nwName, "_")
		id, err := strconv.Atoi(ids[len(ids)-1])
		if err != nil {
			panic(err)
		}

		if id >= start && id <= end {
			if len(serials) > 0 {
				nw.Attachments = make(map[string]rna.NDFCAttachmentsValue)
				for j := range serials {
					attach := rna.NDFCAttachmentsValue{}
					attach.SerialNumber = serials[j]
					attach.DeployThisAttachment = deployNeeded
					nw.Attachments[serials[j]] = attach
				}
			}
			nets.Networks[nwName] = nw
		}
	}

	return nets
}

// Delete attachCount attachments in count Networks, if attachCount is -1, delete all attachments
func DeleteNetAttachments(nets *resource_networks.NDFCNetworksModel, count int, attachCount int) {
	for nwName, nw := range nets.Networks {
		if count > 0 {
			for serial := range nw.Attachments {
				if attachCount == -1 {
					nw.Attachments = nil
					break
				}
				if attachCount > 0 {
					delete(nw.Attachments, serial)
					attachCount--
				}
			}
			count--
			nets.Networks[nwName] = nw
		} else {
			break
		}
	}
}

func DeleteNetworks(nw **resource_networks.NDFCNetworksModel, start, end int) {
	nets := *nw
	log.Printf("Delete Networks: %d to %d", start, end)
	for nwName := range nets.Networks {
		ids := strings.Split(nwName, "_")
		id, err := strconv.Atoi(ids[len(ids)-1])
		if err != nil {
			panic(err)
		}
		if id >= start && id <= end {
			delete(nets.Networks, nwName)
		}
	}
}
func NetAttachmentsMod(nw **resource_networks.NDFCNetworksModel, start, end int, serials []string, serial string, x map[string]interface{}) {
	nwRsc := *nw
	var net []string
	var nwName string
	for i := start; i <= end; i++ {
		nwName = GetConfig().NDFC.NetPrefix + strconv.Itoa(i)
		_, ok := nwRsc.Networks[nwName]
		if !ok {
			panic("Net not found")
		}
		net = append(net, nwName)
	}

	if len(serials) > 0 {
		for i := range net {
			nwName := net[i]
			netEntry, ok := nwRsc.Networks[nwName]
			if !ok {
				panic("Net not found")
			}
			netEntry.Attachments = make(map[string]rna.NDFCAttachmentsValue)
			for j := range serials {
				attach := rna.NDFCAttachmentsValue{}
				attach.SerialNumber = serials[j]
				//attach.DeployThisAttachment = true
				netEntry.Attachments[serials[j]] = attach
			}
			nwRsc.Networks[nwName] = netEntry
		}
	} else {
		for i := range net {
			nwName := net[i]
			netEntry, ok := nwRsc.Networks[nwName]
			if !ok {
				panic("Net not found")
			}
			netEntry.Attachments = nil
			nwRsc.Networks[nwName] = netEntry
		}
	}

	if serial != "" {
		if len(net) != 1 {
			panic("Only one attachment entry can be modified at a time")
		}
		nwName := net[0]
		netEntry, ok := nwRsc.Networks[nwName]
		if !ok {
			panic("VRF not found")
		}
		attachEntry, ok := netEntry.Attachments[serial]
		if !ok {
			panic("Serial not found in Attachments")
		}

		for key, value := range x {
			switch key {
			case "vlan":
				attachEntry.Vlan = new(types.Int64Custom)
				*attachEntry.Vlan = types.Int64Custom(value.(int))
			case "freeform_config":
				attachEntry.FreeformConfig = value.(string)
			case "switch_ports":
				attachEntry.SwitchPorts = value.(types.CSVString)
			}
		}
		netEntry.Attachments[serial] = attachEntry
		nwRsc.Networks[nwName] = netEntry
	}
	*nw = nwRsc
}
