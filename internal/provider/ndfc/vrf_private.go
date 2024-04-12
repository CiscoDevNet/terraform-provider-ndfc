package ndfc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"terraform-provider-ndfc/internal/provider/resources/resource_vrf_bulk"
	. "terraform-provider-ndfc/internal/provider/types"

	api "terraform-provider-ndfc/internal/provider/ndfc/api"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
)

func (c NDFC) vrfCreateBulk(ctx context.Context, fabricName string, vrfsPayload *resource_vrf_bulk.NDFCBulkVrfPayload) error {
	tflog.Info(ctx, fmt.Sprintf("Beginning Bulk VRF create in fabric %s", fabricName))
	data, err := json.Marshal(vrfsPayload.Vrfs)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Json Marshal failure %s", err.Error()))
		return err
	}
	log.Println("Data to be posted", string(data))

	vrfObj := api.NewVrfAPI(fabricName, c.GetLock(ResourceVrfBulk), &c.apiClient)
	res, err := vrfObj.Post(data)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Error POST:  %s", err.Error()))
		okList, err1 := c.processBulkResponse(ctx, res)
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

	vrfObj := api.NewVrfAPI(fabricName, c.GetLock(ResourceVrfBulk), &c.apiClient)
	vrfObj.SetDeleteList(vrfList)
	res, err := vrfObj.Delete()
	if err != nil {

		_, err1 := c.processBulkResponse(ctx, res)
		return err1
	}
	tflog.Info(ctx, fmt.Sprintf("Deleting VRFs OK fabric_name=%s, vrfs = %v", fabricName, vrfList))
	return nil

}

func (c NDFC) vrfBulkGet(ctx context.Context, fabricName string) (*resource_vrf_bulk.NDFCVrfBulkModel, error) {

	vrfObj := api.NewVrfAPI(fabricName, c.GetLock(ResourceVrfBulk), &c.apiClient)
	res, err := vrfObj.Get()
	if err != nil {
		return nil, err
	}
	tflog.Debug(ctx, fmt.Sprintf("RscGetBulkVrf: result %s", string(res)))
	ndVrfs := resource_vrf_bulk.NDFCVrfBulkModel{}
	vrfPayloads := resource_vrf_bulk.NDFCBulkVrfPayload{}
	err = json.Unmarshal(res, &vrfPayloads.Vrfs)
	if err != nil {
		return nil, err
	} else {
		tflog.Debug(ctx, "resource_vrf_bulk: Unmarshal OK")
	}
	ndVrfs.FillVrfsFromPayload(&vrfPayloads)
	return &ndVrfs, nil
}

/*
func (c NDFC) vrfBulkCreateCheck(ctx context.Context, ID string) error {

	retVrfs, err := c.vrfBulkIsPresent(ctx, ID)
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
*/
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
	vrfState := vState.GetModelData()
	vrfConfig := vPlan.GetModelData()
	//vrfConfig := vConfig.GetModelData()

	//var delVrfs []string
	var deployVrfs []string

	putVRFs := new(resource_vrf_bulk.NDFCVrfBulkModel)
	putVRFs.Vrfs = make(map[string]resource_vrf_bulk.NDFCVrfsValue)
	putVRFs.FabricName = vPlan.FabricName.ValueString()

	newVRFs := new(resource_vrf_bulk.NDFCVrfBulkModel)
	newVRFs.Vrfs = make(map[string]resource_vrf_bulk.NDFCVrfsValue)
	newVRFs.FabricName = vPlan.FabricName.ValueString()

	delVrfs := new(resource_vrf_bulk.NDFCVrfBulkModel)
	delVrfs.Vrfs = make(map[string]resource_vrf_bulk.NDFCVrfsValue)
	delVrfs.FabricName = vPlan.FabricName.ValueString()

	for sVrfName, sVrf := range vrfState.Vrfs {
		if vrf, ok := vrfConfig.Vrfs[sVrfName]; ok {

			vrf.FilterThisValue = true
			vrf.FabricName = newVRFs.FabricName
			vrf.VrfName = sVrfName
			updateAction := vrf.CreatePlan(sVrf) //vrfState.Vrfs[i].DeepEqual(*vrf)
			if updateAction == ActionNone {
				//Case 1: Both VRFs are equal - no change to the VRF entry
				tflog.Info(ctx, fmt.Sprintf("%s not changed", sVrfName))

			} else if updateAction == RequiresReplace {
				//Case 2: attribute that cannot be modified in-place has changed - DELETE and Create
				tflog.Info(ctx, fmt.Sprintf("%s Needs to be replaced - Delete and Add |%s|", sVrfName, vrf.VrfName))
				//use the object in state for delete
				delVrfs.Vrfs[vrf.VrfName] = sVrf
				newVRFs.Vrfs[vrf.VrfName] = vrf
			} else if updateAction == ControlFlagUpdate {
				deployVrfs = append(deployVrfs, vrf.VrfName)

			} else {
				//Case 3: attributes have changed - Do update
				putVRFs.Vrfs[vrf.VrfName] = vrf
				tflog.Info(ctx, fmt.Sprintf("%s has changed", vrf.VrfName))
			}
			//put back updates
			vrfConfig.Vrfs[sVrfName] = vrf
		} else {
			//case 4: VRF is missing in plan data - Delete it
			tflog.Info(ctx, fmt.Sprintf("%s Missing in Plan - Needs deletion", sVrfName))
			delVrfs.Vrfs[sVrfName] = sVrf
		}
	}
	//case 5: Deal with New VRFs in plan - Add
	for k, v := range vrfConfig.Vrfs {
		if !v.FilterThisValue {
			v.FabricName = newVRFs.FabricName
			newVRFs.Vrfs[k] = v
		}
	}
	actions["add"] = newVRFs
	actions["put"] = putVRFs
	actions["plan"] = vrfConfig
	actions["state"] = vrfState
	actions["del"] = delVrfs
	actions["deploy"] = deployVrfs

	return actions
}

func (c NDFC) vrfBulkUpdate(ctx context.Context, dg *diag.Diagnostics, ndVRFs *resource_vrf_bulk.NDFCVrfBulkModel) {
	// PUT for each vrf
	payload := ndVRFs.FillVrfPayloadFromModel(nil)

	vrfObj := api.NewVrfAPI(ndVRFs.FabricName, c.GetLock(ResourceVrfBulk), &c.apiClient)

	for i := range payload.Vrfs {
		data, err := json.Marshal(payload.Vrfs[i])
		if err != nil {
			dg.AddError("Marshal Failed", fmt.Sprintf("VRF %s Marshall error %v", payload.Vrfs[i].VrfName, err))
			return
		}
		tflog.Info(ctx, fmt.Sprintf("Update VRF %s", payload.Vrfs[i].VrfName))
		vrfObj.PutVrf = payload.Vrfs[i].VrfName
		res, err := vrfObj.Put(data)
		if err != nil {
			dg.AddError(fmt.Sprintf("VRF %s, Update failed", payload.Vrfs[i].VrfName), fmt.Sprintf("Error %v, response %s", err, res.Str))
			return
		}
		tflog.Info(ctx, fmt.Sprintf("Update VRF %s Successfull. Message %s", payload.Vrfs[i].VrfName, res.Str))
	}
}
