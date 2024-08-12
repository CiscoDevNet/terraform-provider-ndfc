// Code generated;  DO NOT EDIT.

package ndfc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"terraform-provider-ndfc/internal/provider/ndfc/api"
	"terraform-provider-ndfc/internal/provider/resources/resource_networks"
	. "terraform-provider-ndfc/internal/provider/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func (c NDFC) networksCreate(ctx context.Context, fabricName string, rscModel *resource_networks.NDFCNetworksModel) error {
	tflog.Info(ctx, fmt.Sprintf("Beginning Bulk Networks create in fabric %s", fabricName))
	payload := rscModel.FillNetworksPayloadFromModel()

	data, err := json.Marshal(payload.Networks)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Json Marshal failure %s", err.Error()))
		return err
	}
	log.Println("Data to be posted", string(data))

	rsObj := api.NewNetworksAPI(fabricName, c.GetLock(ResourceNetworks), &c.apiClient)
	res, err := rsObj.Post(data)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Error POST:  %s", err.Error()))
		okList, err1 := c.processBulkResponse(ctx, res)
		err2 := c.networksDelete(ctx, fabricName, okList)
		return errors.Join(err, err1, err2)
	}

	tflog.Info(ctx, fmt.Sprintf("resource create: Success res : %v", res.Str))
	return nil
}

func (c NDFC) networksDelete(ctx context.Context, fabricName string, rsList []string) error {
	if len(rsList) == 0 {
		return nil
	}
	tflog.Info(ctx, fmt.Sprintf("Attempting to delete resource fabric_name=%s, resources = %v", fabricName, rsList))
	rsObj := api.NewNetworksAPI(fabricName, c.GetLock(ResourceNetworks), &c.apiClient)
	rsObj.SetDeleteList(rsList)
	res, err := rsObj.Delete()
	if err != nil {
		_, err1 := c.processBulkResponse(ctx, res)
		return err1
	}
	tflog.Info(ctx, fmt.Sprintf("Deleting resources OK fabric_name=%s, resource names = %v", fabricName, rsList))
	return nil

}

func (c NDFC) networksGet(ctx context.Context, fabricName string) (*resource_networks.NDFCNetworksModel, error) {

	rsObj := api.NewNetworksAPI(fabricName, c.GetLock(ResourceNetworks), &c.apiClient)
	res, err := rsObj.Get()
	if err != nil {
		return nil, err
	}
	tflog.Debug(ctx, fmt.Sprintf("RscGetNetworks: result %s", string(res)))
	rsModel := resource_networks.NDFCNetworksModel{}
	payloads := resource_networks.NDFCNetworksPayload{}
	err = json.Unmarshal(res, &payloads.Networks)
	if err != nil {
		return nil, err
	} else {
		tflog.Debug(ctx, "resource_networks: Unmarshal OK")
	}
	rsModel.FillNetworksFromPayload(&payloads)
	return &rsModel, nil
}

func (c NDFC) networksIsPresent(ctx context.Context, ID string) ([]string, error) {
	var ret []string
	filterMap := make(map[string]bool)

	fabricName, _ := c.CreateFilterMap(ID, &filterMap)
	rsModel, err := c.networksGet(ctx, fabricName)
	if err != nil {
		tflog.Error(ctx, "Error while getting resource ", map[string]interface{}{"Err": err})
		return nil, err
	}
	for k, _ := range rsModel.Networks {
		ok, found := filterMap[k]
		if ok && found {
			ret = append(ret, k)
		}
	}
	return ret, nil
}

func (c NDFC) networksUpdate(ctx context.Context, dg *diag.Diagnostics, updateRsc *resource_networks.NDFCNetworksModel, retry int) {
	// PUT for each object
	payload := updateRsc.FillNetworksPayloadFromModel()

	if retry > 5 {
		dg.AddError("Update Failed", "Retry count exceeded")
		return
	}
	retryIndices := make([]int, 0)
	rsObj := api.NewNetworksAPI(updateRsc.FabricName, c.GetLock(ResourceNetworks), &c.apiClient)

	for i := range payload.Networks {
		data, err := json.Marshal(payload.Networks[i])
		if err != nil {
			dg.AddError("Marshal Failed", fmt.Sprintf("Resource %s Marshall error %v", payload.Networks[i].NetworkName, err))
			return
		}
		tflog.Info(ctx, fmt.Sprintf("Update resource %s", payload.Networks[i].NetworkName))
		rsObj.PutNetworkName = payload.Networks[i].NetworkName
		res, err := rsObj.Put(data)
		if err != nil {
			dg.AddError(fmt.Sprintf("Resource %s, Update failed", payload.Networks[i].NetworkName), fmt.Sprintf("Error %v, response %s", err, res.Str))
			return
		}
		// Get and check if they Mismatch
		// There is a bug in NDFC where PUT sometimes resets some values to DefaultValue
		// This is mostly seen with dhcpServers payload
		// Re-do ing PUT solves the problem in most cases
		// Hence do a GET nd verify if all params got updated
		// if not add them to a list and re-do PUT until its correct or retry count is exceeded
		rsObj.GetNetworkName = payload.Networks[i].NetworkName
		rs, err := rsObj.Get()
		if err != nil {
			tflog.Error(ctx, "Read resource after PUT Failed")
			dg.AddError(fmt.Sprintf("Resource %s, Get failed", payload.Networks[i].NetworkName), fmt.Sprintf("Error %v, response %s", err, res.Str))
			return
		}
		rsNewValue := resource_networks.NDFCNetworksValue{}
		err = json.Unmarshal(rs, &rsNewValue)
		if err != nil {
			tflog.Error(ctx, "Unmarshal Failed, networks GET followed by PUT")
			dg.AddError(fmt.Sprintf("Resource %s, Unmarshal failed", payload.Networks[i].NetworkName), fmt.Sprintf("Error %v, response %s", err, res.Str))
			return
		}
		if ValuesDeeplyEqual != rsNewValue.DeepEqual(payload.Networks[i]) {
			tflog.Error(ctx, "Mismatch in data retrieved after PUT - add to retry list")
			retryIndices = append(retryIndices, i)
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Update resource %s Successfull. Message %s", payload.Networks[i].NetworkName, res.Str))
	}
	if len(retryIndices) > 0 {
		tflog.Info(ctx, "Retrying network update due to mismatch", map[string]interface{}{"Err": "Mismatch in network data retrieved after PUT"})
		redoRsc := new(resource_networks.NDFCNetworksModel)
		redoRsc.FabricName = updateRsc.FabricName
		redoRsc.Networks = make(map[string]resource_networks.NDFCNetworksValue)
		for _, i := range retryIndices {
			redoRsc.Networks[payload.Networks[i].NetworkName] = payload.Networks[i]
		}
		c.networksUpdate(ctx, dg, redoRsc, retry+1)
	}
}

func (c NDFC) networksGetDiff(ctx context.Context, dg *diag.Diagnostics,
	vPlan *resource_networks.NetworksModel,
	vState *resource_networks.NetworksModel, vConfig *resource_networks.NetworksModel) map[string]interface{} {

	actions := make(map[string]interface{})
	rsState := vState.GetModelData()
	rsConfig := vPlan.GetModelData()

	//rsConfig := vConfig.GetModelData()
	//var delRscs []string
	var deployRscs []string

	putRscs := new(resource_networks.NDFCNetworksModel)
	putRscs.Networks = make(map[string]resource_networks.NDFCNetworksValue)
	putRscs.FabricName = vPlan.FabricName.ValueString()

	newRscs := new(resource_networks.NDFCNetworksModel)
	newRscs.Networks = make(map[string]resource_networks.NDFCNetworksValue)
	newRscs.FabricName = vPlan.FabricName.ValueString()

	delRscs := new(resource_networks.NDFCNetworksModel)
	delRscs.Networks = make(map[string]resource_networks.NDFCNetworksValue)
	delRscs.FabricName = vPlan.FabricName.ValueString()

	for sRsName, sRsc := range rsState.Networks {
		if rsPlan, ok := rsConfig.Networks[sRsName]; ok {
			rsPlan.FilterThisValue = true
			rsPlan.FabricName = newRscs.FabricName
			rsPlan.NetworkName = sRsName
			cf := false
			updateAction := rsPlan.CreatePlan(sRsc, &cf) //dummy cf
			if updateAction == ActionNone {
				//Case 1: Both networks entries are equal - no change
				tflog.Info(ctx, fmt.Sprintf("%s not changed", sRsName))

			} else if updateAction == RequiresReplace {
				//Case 2: attribute that cannot be modified in-place has changed - DELETE and Create
				tflog.Info(ctx, fmt.Sprintf("%s Needs to be replaced - Delete and Add", sRsName))
				//use the object in state for delete
				delRscs.Networks[rsPlan.NetworkName] = sRsc
				newRscs.Networks[rsPlan.NetworkName] = rsPlan
			} else if updateAction == ControlFlagUpdate {
				deployRscs = append(deployRscs, rsPlan.NetworkName)

			} else {
				//Case 3: attributes have changed - Do update
				putRscs.Networks[rsPlan.NetworkName] = rsPlan
				tflog.Info(ctx, fmt.Sprintf("%s has changed", rsPlan.NetworkName))
			}
			//put back updates
			rsConfig.Networks[sRsName] = rsPlan
		} else {
			//case 4: Rsc is missing in plan data - Delete it
			tflog.Info(ctx, fmt.Sprintf("%s Missing in Plan - Needs deletion", sRsName))
			delRscs.Networks[sRsName] = sRsc
		}
	}
	//case 5: Deal with New Rscs in plan - Add
	for k, v := range rsConfig.Networks {
		if !v.FilterThisValue {
			v.FabricName = newRscs.FabricName
			newRscs.Networks[k] = v
		}
	}
	actions["add"] = newRscs
	actions["put"] = putRscs
	actions["plan"] = rsConfig
	actions["state"] = rsState
	actions["del"] = delRscs
	actions["deploy"] = deployRscs

	return actions
}
