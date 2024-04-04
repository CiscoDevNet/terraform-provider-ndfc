package resource_network_attachments

type NDFCNetworkAttachmentsPayload struct {
	NetworkName string                 `json:"networkName,omitempty"`
	Attachments []NDFCAttachmentsValue `json:"lanAttachList,omitempty"`
}

type NDFCNetworkAttachments struct {
	GlobalDeploy       bool
	GlobalUndeploy     bool
	FabricName         string
	NetworkAttachments []NDFCNetworkAttachmentsPayload // Attachment payload for NDFC
	DepMap             map[string][]string             // use for backfilling DeploymentFlag in TF state        // for deployment
}

func (p *NDFCNetworkAttachments) AddEntry(nwName string, attachList []NDFCAttachmentsValue) {
	if len(attachList) == 0 {
		return
	}
	nwAttachEntry := NDFCNetworkAttachmentsPayload{NetworkName: nwName}
	nwAttachEntry.Attachments = attachList
	p.NetworkAttachments = append(p.NetworkAttachments, nwAttachEntry)
}
