package testing

import (
	"terraform-provider-ndfc/internal/provider/resources/resource_vpc_pair"
)

func GenerateVpcPairResource(VpcPairObj **resource_vpc_pair.NDFCVpcPairModel, serials []string, VirtualPeerLink bool) {
	vpcPair := *VpcPairObj
	vpcPair.SerialNumbers = serials[:2]
	*vpcPair.UseVirtualPeerlink = VirtualPeerLink
}
func GenerateVpcPairResourceWithMoreSerials(VpcPairObj **resource_vpc_pair.NDFCVpcPairModel, serials []string, VirtualPeerLink bool) {
	vpcPair := *VpcPairObj
	vpcPair.SerialNumbers = serials
	*vpcPair.UseVirtualPeerlink = VirtualPeerLink
}
