package ndfc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"
	"terraform-provider-ndfc/internal/provider/resources/resource_vrf_bulk"
	. "terraform-provider-ndfc/internal/provider/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
)

const UrlVrfGetBulk = "/lan-fabric/rest/top-down/v2/fabrics/%s/vrfs"
const UrlVrfCreateBulk = "/lan-fabric/rest/top-down/v2/bulk-create/vrfs"
const UrlVrfDeleteBulk = "/lan-fabric/rest/top-down/v2/fabrics/%s/bulk-delete/vrfs"
const UrlVrfGet = "/lan-fabric/rest/top-down/v2/fabrics/%s/vrf/%s"
const UrlVrfUpdate = "/lan-fabric/rest/top-down/v2/fabrics/%s/vrfs/%s"

func (c NDFC) vrfGetAll(ctx context.Context, fabricName string) ([]byte, error) {
	c.GetLock(ResourceVrfBulk).Lock()
	defer c.GetLock(ResourceVrfBulk).Unlock()
	res, err := c.apiClient.GetRawJson(fmt.Sprintf(UrlVrfGetBulk, fabricName))
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c NDFC) vrfCreateBulk(ctx context.Context, fabricName string, vrfs *resource_vrf_bulk.NDFCVrfBulkModel) error {
	unlockOnce := sync.Once{}
	lock := c.GetLock(ResourceVrfBulk)
	tflog.Info(ctx, fmt.Sprintf("Beginning Bulk VRF create in fabric %s", fabricName))
	data, err := json.Marshal(vrfs.Vrfs)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Json Marshal failure %s", err.Error()))
		return err
	}
	log.Println("Data to be posted", string(data))
	lock.Lock()
	defer unlockOnce.Do(lock.Unlock)

	res, err := c.apiClient.Post(UrlVrfCreateBulk, string(data))
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Error POST:  %s", err.Error()))
		okList, err1 := c.processBulkResponse(ctx, res)
		unlockOnce.Do(lock.Unlock)
		err2 := c.vrfBulkDelete(ctx, fabricName, okList)
		return errors.Join(err, err1, err2)
	}

	tflog.Info(ctx, fmt.Sprintf("vrfCreateBulk: Success res : %v", res.Str))
	return nil
}

func (c NDFC) vrfBulkDelete(ctx context.Context, fabricName string, vrfList []string) error {
	if len(vrfList) == 0 {
		return nil
	}
	tflog.Info(ctx, fmt.Sprintf("Attempting to delete VRFs fabric_name=%s, vrfs = %v", fabricName, vrfList))
	qp := []string{"vrf-names=" + strings.Join(vrfList, ",")}
	c.GetLock(ResourceVrfBulk).Lock()
	defer c.GetLock(ResourceVrfBulk).Unlock()
	res, err := c.apiClient.DeleteRaw(fmt.Sprintf(UrlVrfDeleteBulk, fabricName), qp)
	if err != nil {
		_, err1 := c.processBulkResponse(ctx, res)
		return err1
	}
	tflog.Info(ctx, fmt.Sprintf("Deleting VRFs OK fabric_name=%s, vrfs = %v", fabricName, vrfList))
	return nil

}

func (c NDFC) vrfBulkGet(ctx context.Context, fabricName string) *resource_vrf_bulk.NDFCVrfBulkModel {

	data, err := c.vrfGetAll(ctx, fabricName)
	if err != nil {
		tflog.Error(ctx, "Error retrieving existing VRFs", map[string]interface{}{"Err": err})
		return nil
	}

	ndVrfs := new(resource_vrf_bulk.NDFCVrfBulkModel)
	err = json.Unmarshal(data, &ndVrfs.Vrfs)
	if err != nil {
		tflog.Error(ctx, "resource_vrf_bulk: unmarshal failed ", map[string]interface{}{"Err": err})
		return nil
	} else {
		tflog.Debug(ctx, "resource_vrf_bulk: Unmarshal OK")
	}
	return ndVrfs
}

func (c NDFC) vrfBulkIsPresent(ctx context.Context, ID string, fabricName string) ([]string, error) {
	var retVrfs []string
	filterMap := make(map[string]bool)

	ndVrfs := c.vrfBulkGet(ctx, fabricName)
	c.vrfCreateFilterMap(ID, &filterMap)

	for i := range ndVrfs.Vrfs {
		ok, found := filterMap[ndVrfs.Vrfs[i].VrfName]
		if ok && found {
			retVrfs = append(retVrfs, ndVrfs.Vrfs[i].VrfName)
		}
	}
	return retVrfs, nil

}

func (c NDFC) vrfBulkCreateCheck(ctx context.Context, ID string, vrfs *resource_vrf_bulk.VrfBulkModel) error {

	retVrfs, err := c.vrfBulkIsPresent(ctx, ID, vrfs.FabricName.ValueString())
	if err != nil {
		tflog.Error(ctx, "Error while getting VRFs ", map[string]interface{}{"Err": err})
		return err
	}

	var errs []error
	for i := range retVrfs {
		errs = append(errs, fmt.Errorf("VRF %s is already configured on %s", retVrfs[i], vrfs.FabricName.ValueString()))
	}

	if len(errs) > 0 {
		tflog.Error(ctx, "VRFs exist", map[string]interface{}{"Err": errs})
		return errors.Join(errs...)
	}

	tflog.Info(ctx, "VRFs were not found in NDFC, ok to create")
	return nil
}

/*
Failure Response format
{"failureList":

	     [{"name":"Murali_Bulk_01","message":"VRF\tMurali_Bulk_01\talready exists","status":"Failed"},
		  {"name":"Murali_Bulk_02","message":"VRF\tMurali_Bulk_02\talready exists","status":"Failed"}
		  ],
		  "successList":[]
		  {"successList":
		  [{"name":"Murali_Bulk_01","message":"VRF is successfully created.","status":"Success"},
		   {"name":"Murali_Bulk_02","message":"VRF is successfully created.","status":"Success"}]}
	}
*/
func (c NDFC) processBulkResponse(ctx context.Context, res gjson.Result) ([]string, error) {
	tflog.Error(ctx, res.String())
	flist := res.Get("failureList")
	var failed []map[string]string
	var errs []error
	e := json.Unmarshal([]byte(flist.Raw), &failed)
	if e != nil {
		log.Println("Error unmarshalling response from NDFC")
		errs = append(errs, e)
	}
	for _, v := range failed {
		errs = append(errs, fmt.Errorf("VRF=%s Status=%s Message=%s", v["name"], v["status"], v["message"]))
	}
	var arr []string
	if res.Get("successList.#.name").Exists() {
		slist := res.Get("successList.#.name").Value().([]interface{})
		for _, v := range slist {
			arr = append(arr, v.(string))
		}
	}
	return arr, errors.Join(errs...)
}

func (c NDFC) vrfCreateFilterMap(ID string, filterMap *map[string]bool) (string, []string) {

	result := c.VrfBulkSplitID(ID)
	for _, v := range result["vrfs"] {
		log.Printf("Set filtering of %s to true", v)
		(*filterMap)[v] = true
	}
	return result["fabric"][0], result["vrfs"]
}

/*
	func (c NDFC) vrfBulkSplitID(ID string) (string, []string) {
		idSplit := strings.Split(ID, "/")
		fabricName := idSplit[0]
		var vrfs []string
		if len(idSplit) > 1 {
			vrfs_string := idSplit[1]
			log.Println(vrfs_string)
			if vrfs_string[0] == '{' && vrfs_string[len(vrfs_string)-1] == '}' {
				vrfs = strings.Split(vrfs_string[1:len(vrfs_string)-1], ",")
			}
		}
		return fabricName, vrfs
	}
*/
func (c NDFC) vrfBulkGetDiff(ctx context.Context, dg *diag.Diagnostics,
	vPlan *resource_vrf_bulk.VrfBulkModel,
	vState *resource_vrf_bulk.VrfBulkModel, vConfig *resource_vrf_bulk.VrfBulkModel) map[string]interface{} {

	actions := make(map[string]interface{})
	v1 := vState.GetModelData()
	v2 := vPlan.GetModelData()
	v3 := vConfig.GetModelData()

	var delVrfs []string
	var deployVrfs []string

	a1 := []*resource_vrf_bulk.NDFCVrfBulkModel{v1, v2, v3}
	log.Printf("Update Dump=====================================start")
	for i := range a1 {
		data, err := json.Marshal(a1[i].Vrfs)
		if err != nil {
			log.Printf("vrfBulkGetDiff: Marshal failed state")
		} else {
			log.Printf("%s", string(data))
		}
		for j := range a1[i].Vrfs {
			data, err := json.Marshal(a1[i].Vrfs[j].AttachList)
			if err == nil {
				log.Printf("%s", string(data))
			}
		}
	}
	log.Printf("Update Dump=====================================end")

	putVRFs := new(resource_vrf_bulk.NDFCVrfBulkModel)
	putVRFs.FabricName = vPlan.FabricName.ValueString()

	newVRFs := new(resource_vrf_bulk.NDFCVrfBulkModel)
	newVRFs.FabricName = vPlan.FabricName.ValueString()
	for i := range v1.Vrfs {
		if vrf, ok := v2.VrfsMap[v1.Vrfs[i].VrfName]; ok {
			vrf.FilterThisValue = true
			updateAction := v1.Vrfs[i].DeepEqual(*vrf)
			if updateAction == ValuesDeeplyEqual {
				//Case 1: Both VRFs are equal - no change to the VRF entry
				tflog.Info(ctx, fmt.Sprintf("%s not changed", v1.Vrfs[i].VrfName))

			} else if updateAction == RequiresReplace {
				//Case 2: attribute that cannot be modified in-place has changed - DELETE and Create
				tflog.Info(ctx, fmt.Sprintf("%s Needs to be replaced - Delete and Add", v1.Vrfs[i].VrfName))
				delVrfs = append(delVrfs, vrf.VrfName)
				vrf.FabricName = newVRFs.FabricName
				newVRFs.Vrfs = append(newVRFs.Vrfs, *vrf)
			} else if updateAction == ControlFlagUpdate {
				deployVrfs = append(deployVrfs, vrf.VrfName)

			} else {
				//Case 3: attributes have changed - Do update
				vrf.FabricName = newVRFs.FabricName
				putVRFs.Vrfs = append(putVRFs.Vrfs, *vrf)
				tflog.Info(ctx, fmt.Sprintf("%s has changed", v1.Vrfs[i].VrfName))
			}
		} else {
			//case 4: VRF is missing in plan data - Delete it
			tflog.Info(ctx, fmt.Sprintf("%s Missing in Plan - Needs deletion", v1.Vrfs[i].VrfName))
			delVrfs = append(delVrfs, v1.Vrfs[i].VrfName)
		}
	}
	//case 5: Deal with New VRFs in plan - Add
	for i := range v2.Vrfs {
		if !v2.Vrfs[i].FilterThisValue {
			v2.Vrfs[i].FabricName = newVRFs.FabricName
			newVRFs.Vrfs = append(newVRFs.Vrfs, v2.Vrfs[i])
		}
	}
	actions["add"] = newVRFs
	actions["put"] = putVRFs
	actions["plan"] = v2
	actions["state"] = v1
	actions["del"] = delVrfs
	actions["deploy"] = deployVrfs

	return actions

}

func (c NDFC) vrfBulkUpdate(ctx context.Context, dg *diag.Diagnostics, ndVRFs *resource_vrf_bulk.NDFCVrfBulkModel) {
	// PUT for each vrf

	for i := range ndVRFs.Vrfs {
		ndVRFs.Vrfs[i].FabricName = ndVRFs.FabricName
		data, err := json.Marshal(ndVRFs.Vrfs[i])
		if err != nil {
			dg.AddError("Marshal Failed", fmt.Sprintf("VRF %s Marshall error %v", ndVRFs.Vrfs[i].VrfName, err))
			return
		}
		tflog.Info(ctx, fmt.Sprintf("Update VRF %s", ndVRFs.Vrfs[i].VrfName))

		res, err := c.apiClient.Put(fmt.Sprintf(UrlVrfUpdate, ndVRFs.FabricName, ndVRFs.Vrfs[i].VrfName), string(data))
		if err != nil {
			dg.AddError(fmt.Sprintf("VRF %s, Update failed", ndVRFs.Vrfs[i].VrfName), fmt.Sprintf("Error %v, response %s", err, res.Str))
			return
		}
		tflog.Info(ctx, fmt.Sprintf("Update VRF %s Successfull. Message %s", ndVRFs.Vrfs[i].VrfName, res.Str))
	}

}
