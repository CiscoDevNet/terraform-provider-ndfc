package ndfc

import (
	"context"
	"fmt"
	"log"
	"strings"
	"terraform-provider-ndfc/internal/provider/resources/resource_networks"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const ResourceNetworks = "networks"

func (c *NDFC) RscCreateNetworks(ctx context.Context, dg *diag.Diagnostics, in *resource_networks.NetworksModel) *resource_networks.NetworksModel {
	// Create API call logic
	tflog.Debug(ctx, fmt.Sprintf("RscCreateNetworks entry fabirc %s", in.FabricName.ValueString()))
	nw := in.GetModelData()
	if nw == nil {
		tflog.Error(ctx, "Data conversion from model failed")
		dg.AddError("Data conversion from model failed", "GetModelData returned empty")
		return nil
	}
	//form ID
	ID := c.RscCreateID(nw, ResourceNetworks)
	//create

	depMap := make(map[string][]string)

	retNets, err := c.networksIsPresent(ctx, ID)

	if err != nil {
		tflog.Error(ctx, "Error while getting Networks ", map[string]interface{}{"Err": err})
		dg.AddError("Network Read Failed", err.Error())
		return nil
	}

	var errs []string
	for i := range retNets {
		errs = append(errs, fmt.Sprintf("Network %s is already configured on %s", retNets[i], nw.FabricName))
	}
	if len(errs) > 0 {
		tflog.Error(ctx, "Networks exist", map[string]interface{}{"Err": errs})
		dg.AddError("Networks exist", strings.Join(errs, ","))
		return nil
	}

	//Part 1: Create Networks

	err = c.networksCreate(ctx, nw.FabricName, nw)
	if err != nil {
		tflog.Error(ctx, "Cannot create Network", map[string]interface{}{"Err": err})
		dg.AddError("Cannot create Network", err.Error())
		return nil
	}
	tflog.Info(ctx, fmt.Sprintf("Create Bulk VRF success ID %s", ID))
	//Part 2: Create Attachments if any

	if nw.DeployAllAttachments {
		depMap["global"] = append(depMap["global"], "all")
	}

	err = c.netAttachmentsAttach(ctx, nw)
	if err != nil {
		tflog.Error(ctx, "Network Attachments create failed")
		dg.AddError("Network Attachments create failed", err.Error())
		tflog.Error(ctx, "Rolling back the configurations...delete Networks")
		c.RscDeleteNetworks(ctx, dg, ID, in)
		return nil
	}
	//Part 3: Deploy Attachments if any
	//c.RscDeployAttachments(ctx, dg, nw)

	out := c.RscGetNetworks(ctx, dg, ID, &depMap)
	if out == nil {
		tflog.Error(ctx, "Failed to verify: Reading from NDFC after create failed")
		dg.AddError("Failed to verify", "Reading from NDFC after create failed")
		return nil
	}
	if ID != "" {
		out.Id = types.StringValue(ID)
	}
	return out
}

func (c *NDFC) RscGetNetworks(ctx context.Context, dg *diag.Diagnostics, ID string, depMap *map[string][]string) *resource_networks.NetworksModel {
	// Read API call logic
	var filterMap map[string]bool
	tflog.Debug(ctx, fmt.Sprintf("RscGetNetworks entry fabirc %s", ID))

	filterMap = make(map[string]bool)
	fabricName, networks := c.CreateFilterMap(ID, &filterMap)

	log.Printf("FilterMap: %v networks %v", filterMap, networks)

	if fabricName == "" {
		dg.AddError("ID format error", "ID is incorrect")
		return nil
	}
	ndNets, err := c.networksGet(ctx, fabricName)
	if err != nil {
		dg.AddError("Network Get Failed", err.Error())
		return nil
	}

	if len(filterMap) > 0 {
		log.Printf("Filtering is configured")
		//Set value filter - skip the nws that are not in ID
		for i, entry := range ndNets.Networks {
			if _, found := (filterMap)[entry.NetworkName]; !found {
				log.Printf("Filtering out Network %s", entry.NetworkName)
				entry.FilterThisValue = true
				ndNets.Networks[i] = entry
			}
		}
	}

	if _, ok := (*depMap)["global"]; ok {
		//This cannot be validated - as we don't know if all deployments were ok
		ndNets.DeployAllAttachments = true
	}

	//Get Attachments

	err = c.RscGetNetworkAttachments(ctx, ndNets)
	if err == nil {
		tflog.Debug(ctx, "Network Attachments read success")

		for i, nwEntry := range ndNets.Networks {
			if nwEntry.FilterThisValue {
				continue
			}
			vrfLevelDep := false
			if vl, vlOk := (*depMap)[i]; vlOk {
				//first element is vrf name if vrf level deploy is set
				if vl[0] == i {
					nwEntry.DeployAttachments = (nwEntry.NetworkStatus == "DEPLOYED")
					log.Printf("Setting Network level dep flag for %s to %v", i, nwEntry.DeployAttachments)
					vrfLevelDep = true
				}
			}
			for j, attachEntry := range nwEntry.Attachments {
				if attachEntry.FilterThisValue {
					continue
				}
				log.Printf("Attachment %s added to Network %s", j, i)
				if !vrfLevelDep {
					if attachEntry.AttachState == "DEPLOYED" {
						log.Printf("Attachment %s deployed", j)
						attachEntry.DeployThisAttachment = true
					}
				}

				if portOrder, ok := (*depMap)["SwitchPorts:"+i+"/"+j]; ok {
					log.Printf("SwitchPorts order %v", portOrder)
					processPortListOrder(&attachEntry.SwitchPorts, portOrder)
				}
				if portOrder, ok := (*depMap)["TorPorts:"+i+"/"+j]; ok {
					log.Printf("TorPorts order %v", portOrder)
					processPortListOrder(&attachEntry.TorPorts, portOrder)
				}
				//put modified entry back
				nwEntry.Attachments[j] = attachEntry
			}
			//put modified entry back
			ndNets.Networks[i] = nwEntry
		}

	} else {
		tflog.Error(ctx, "Network Attachments read failed", map[string]interface{}{"Err": err})
		dg.AddError("Network Attachments read failed", err.Error())
	}

	vModel := new(resource_networks.NetworksModel)
	vModel.Id = types.StringValue(ID)
	d := vModel.SetModelData(ndNets)
	if d != nil {
		dg.Append(d.Errors()...)
	}
	return vModel
}

// plan is updated at the end to reflect the current config and returned to TF
func (c *NDFC) RscUpdateNetworks(ctx context.Context, dg *diag.Diagnostics, ID string,
	plan *resource_networks.NetworksModel,
	state *resource_networks.NetworksModel,
	config *resource_networks.NetworksModel) {

	actions := c.networksGetDiff(ctx, dg, plan, state, config)

	delNw := actions["del"].(*resource_networks.NDFCNetworksModel)
	putNw := actions["put"].(*resource_networks.NDFCNetworksModel)
	newNw := actions["add"].(*resource_networks.NDFCNetworksModel)
	planNw := actions["plan"].(*resource_networks.NDFCNetworksModel)
	stateNw := actions["state"].(*resource_networks.NDFCNetworksModel)

	// Check if networks to be added exists
	nws, err := c.networksGet(ctx, plan.FabricName.ValueString())
	if err != nil {
		tflog.Error(ctx, "Network Get Failed", map[string]interface{}{"Err": err})
		dg.AddError("Network Get Failed", err.Error())
		return
	}
	for k, _ := range newNw.Networks {
		if _, ok := nws.Networks[k]; ok {
			// check if the network is marked for deletion
			if _, ok := delNw.Networks[k]; ok {
				tflog.Debug(ctx, "Network marked for deletion", map[string]interface{}{"Network": k})
				continue
			}
			tflog.Error(ctx, "Network already exists", map[string]interface{}{"Network": k})
			dg.AddError("Network already exists", fmt.Sprintf("Network %s already exists", k))
			return
		}
	}

	//Check if networks to be deleted exists

	for k, _ := range delNw.Networks {
		if _, ok := nws.Networks[k]; !ok {
			// Should we error out in this situation or just continue ?
			tflog.Error(ctx, "Network does not exist", map[string]interface{}{"Network": k})
			dg.AddWarning("Network does not exist", fmt.Sprintf("Network %s does not exist", k))
		}
	}

	// Check if networks to be updated exists

	for k, _ := range putNw.Networks {
		if _, ok := nws.Networks[k]; !ok {
			tflog.Error(ctx, "Network does not exist", map[string]interface{}{"Network": k})
			dg.AddError("Network does not exist", fmt.Sprintf("Network %s does not exist", k))
			return
		}
	}

	// Update API call logic
	// 1. Delete items marked for delete as well as re-create

	if len(delNw.Networks) > 0 {
		// Detach Attachments
		err := c.netAttachmentsDetach(ctx, delNw)
		if err != nil {
			tflog.Error(ctx, "Network Attachments delete failed", map[string]interface{}{"Err": err})
			dg.AddError("Network Attachments delete failed", err.Error())
			return
		}
		//TBD: Deploy so that attachments are re-deployed

		// Delete Networks
		err = c.networksDelete(ctx, delNw.FabricName, delNw.GetNetworksNames())
		if err != nil {
			tflog.Error(ctx, "Networks delete failed", map[string]interface{}{"Err": err})
			dg.AddError("Networks delete failed", err.Error())
			return
		}

	}

	//2. Update items marked for update
	if len(putNw.Networks) > 0 {
		c.networksUpdate(ctx, dg, putNw, 0)
		if dg.HasError() {
			//Roll back here is pretty complex - so not attempting
			tflog.Error(ctx, "Networks update failed", map[string]interface{}{"Err": err})
			dg.AddError("Networks update failed", "Rollback not attempted")
			return
		}
	}

	//3. Create new items
	if len(newNw.Networks) > 0 {
		err := c.networksCreate(ctx, newNw.FabricName, newNw)
		if err != nil {
			//On failure - creation api deletes any successful entry
			tflog.Error(ctx, "Networks create failed", map[string]interface{}{"Err": err})
			dg.AddError("Networks create failed", err.Error())
			return
		}
	}

	// Deal with attachments
	c.RscUpdateNetAttachments(ctx, dg, planNw, stateNw)
	if dg.HasError() {
		tflog.Error(ctx, "Network Attachments update failed")
		return
	}

	newID := c.RscCreateID(plan.GetModelData(), ResourceNetworks)
	tflog.Info(ctx, fmt.Sprintf("New ID after update %s", newID))

	depMap := make(map[string][]string)
	//4. Read the updated data from NDFC
	*plan = *(c.RscGetNetworks(ctx, dg, newID, &depMap))
}

func (c *NDFC) RscDeleteNetworks(ctx context.Context, dg *diag.Diagnostics, ID string, in *resource_networks.NetworksModel) {
	// Delete API call logic

	// Get Network names from ID
	rsList, err := c.networksIsPresent(ctx, ID)
	if err != nil {
		tflog.Error(ctx, "Error while getting Networks ", map[string]interface{}{"Err": err})
		dg.AddError("Network Read Failed", err.Error())
		return
	}

	if len(rsList) == 0 {
		tflog.Info(ctx, "No Networks to delete")
		return
	}

	results := c.RscBulkSplitID(ID)
	fabricName := results["fabric"][0]
	nwFromId := results["rsc"]

	if len(nwFromId) != len(rsList) {
		tflog.Error(ctx, "Mismatch in Network - ID not accurate", map[string]interface{}{"ID": ID})
		dg.AddError("Mismatch in network names in ID", "ID is incorrect")
		rsList = nwFromId
	}

	err = c.networksDelete(ctx, fabricName, rsList)
	if err != nil {
		tflog.Error(ctx, "Networks delete failed", map[string]interface{}{"Err": err})
		dg.AddError("Networks delete failed", err.Error())
		return
	}
	tflog.Info(ctx, "Networks delete success")

}
