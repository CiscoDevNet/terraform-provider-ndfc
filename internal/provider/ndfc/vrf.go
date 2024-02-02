package ndfc

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strings"
	"terraform-provider-ndfc/internal/provider/datasources/datasource_vrf"
	"terraform-provider-ndfc/internal/provider/datasources/datasource_vrf_bulk"
	"terraform-provider-ndfc/internal/provider/resources/resource_vrf_bulk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const ResourceVrfBulk = "vrf_bulk"

func (c NDFC) DSGetBulkVrf(ctx context.Context, dg *diag.Diagnostics, fabricName string) *datasource_vrf_bulk.VrfBulkModel {
	log.Printf("DSGetBulkVrf entry fabirc %s", fabricName)
	//tflog.Debug(ctx, fmt.Sprintf("DSGetBulkVrf entry fabirc %s", fabricName))
	//time.Sleep(1 * time.Second)
	res, err := c.vrfGetAll(ctx, fabricName)
	if err != nil {
		dg.AddError("VRF Get Failed", err.Error())
		return nil
	}

	ndVrfs := datasource_vrf_bulk.NDFCVrfBulkModel{}
	err = json.Unmarshal(res, &ndVrfs.Vrfs)
	if err != nil {
		dg.AddError("datasource_vrf_bulk: unmarshal failed ", err.Error())
		return nil
	} else {
		tflog.Debug(ctx, "datasource_vrf_bulk: Unmarshal OK")
	}
	if len(ndVrfs.Vrfs) > 0 {
		ndVrfs.FabricName = ndVrfs.Vrfs[0].FabricName
	}
	data := new(datasource_vrf_bulk.VrfBulkModel)
	d := data.SetModelData(&ndVrfs)
	if d != nil {
		*dg = d
		return nil
	} else {
		tflog.Debug(ctx, "datasource_vrf_bulk: SetModelData OK")
	}
	return data
}

func (c NDFC) DSGetVrf(ctx context.Context, dg *diag.Diagnostics, fabricName string, vrfName string) *datasource_vrf.NDFCVrfModel {
	tflog.Debug(ctx, fmt.Sprintf("DSGetVrf entry fabirc %s vrfName %s", fabricName, vrfName))

	return nil
}

func (c NDFC) VrfBulkCreateID(ndVrfs *resource_vrf_bulk.NDFCVrfBulkModel) string {
	uniqueID := "%s/%v"
	if ndVrfs != nil {
		fName := ndVrfs.FabricName
		vrf_list := make([]string, len(ndVrfs.Vrfs))
		for i, v := range ndVrfs.Vrfs {
			vrf_list[i] = v.VrfName
		}
		vrf_list_str := "{" + strings.Join(vrf_list, ",") + "}"
		return fmt.Sprintf(uniqueID, fName, vrf_list_str)
	}
	return ""
}

func (c NDFC) RscGetBulkVrf(ctx context.Context, dg *diag.Diagnostics, ID string, create bool) *resource_vrf_bulk.VrfBulkModel {
	var filterMap map[string]bool
	tflog.Debug(ctx, fmt.Sprintf("RscGetBulkVrf entry fabirc %s", ID))

	filterMap = make(map[string]bool)
	fabricName, vrfs := c.vrfCreateFilterMap(ID, &filterMap)
	if fabricName == "" {
		dg.AddError("ID format error", "ID is incorrect")
		return nil
	}
	res, err := c.vrfGetAll(ctx, fabricName)
	if err != nil {
		dg.AddError("VRF Get Failed", err.Error())
		return nil
	}
	tflog.Debug(ctx, fmt.Sprintf("RscGetBulkVrf: result %s", string(res)))
	ndVrfs := resource_vrf_bulk.NDFCVrfBulkModel{}
	err = json.Unmarshal(res, &ndVrfs.Vrfs)
	if err != nil {
		dg.AddError("resource_vrf_bulk: unmarshal failed ", err.Error())
		return nil
	} else {
		tflog.Debug(ctx, "resource_vrf_bulk: Unmarshal OK")
	}
	if len(ndVrfs.Vrfs) > 0 {
		ndVrfs.FabricName = ndVrfs.Vrfs[0].FabricName
	} else {
		tflog.Error(ctx, "resource_vrf_bulk: No VRFs found")
		return nil
	}
	if len(filterMap) > 0 {
		log.Printf("Filtering is configured")
		//Set value filter - skip the vrfs that are not in ID
		for i := range ndVrfs.Vrfs {
			if _, found := (filterMap)[ndVrfs.Vrfs[i].VrfName]; !found {
				log.Printf("Filtering out VRF %s", ndVrfs.Vrfs[i].VrfName)
				ndVrfs.Vrfs[i].FilterThisValue = true
			}
		}
	}
	if !create {
		ndVrfs.CreateSearchMap()
		// VRFs should be ordered as they appear in ID field if the Get
		//Except for create case - ideally they should come in the order of creation
		for i := range vrfs {
			vrf, ok := ndVrfs.VrfsMap[vrfs[i]]
			if ok {
				*vrf.Id = int64(i) //sort uses id - so put ascending values
			} else {
				tflog.Error(ctx, fmt.Sprintf("VRF %s missing in output", vrfs[i]))
				//log.Panicf("VRF %s missing in output", vrfs[i])
			}
		}
	}
	sort.Sort(ndVrfs.Vrfs)
	vModel := new(resource_vrf_bulk.VrfBulkModel)
	vModel.Id = types.StringValue(ID)
	vModel.SetModelData(&ndVrfs)
	return vModel
}

/*
Usecase: Bulk VRF Creation

Case 1: No VRFs from list exist on NDFC already

	action: create everything

Case 2: Some VRFs already exists
  - user created some manually
  - vrfs were part of another resource (bulk/single)
    Action : Don't proceed - infra tainted.
    User can manually cleanup or update the resource to exclude existing

Case 3: During Creation, some failed - probably due to a param problem or some other issues

	VRFs partially created
	Action:
	Delete what we created - rollback
	return failure - reporting the issues

case 4: Nothing was created

	return failure, report the errors coming from NDFC
*/
func (c NDFC) RscCreateBulkVrf(ctx context.Context, dg *diag.Diagnostics, vrfBulk *resource_vrf_bulk.VrfBulkModel) *resource_vrf_bulk.VrfBulkModel {
	tflog.Debug(ctx, fmt.Sprintf("RscCreateBulkVrf entry fabirc %s", vrfBulk.FabricName.ValueString()))
	vrf := vrfBulk.GetModelData()
	if vrf == nil {
		tflog.Error(ctx, "Data conversion from model failed")
		dg.AddError("Data conversion from model failed", "GetModelData returned empty")
		return nil
	}
	//Fill fabricName in each VRF
	for i := range vrf.Vrfs {
		vrf.Vrfs[i].FabricName = vrf.FabricName
	}
	//form ID
	ID := c.VrfBulkCreateID(vrf)
	//create

	// check if any of the vrfs exists
	err := c.vrfBulkCreateCheck(ctx, ID, vrfBulk)
	if err != nil {
		tflog.Error(ctx, "Cannot create VRF", map[string]interface{}{"Err": err})
		dg.AddError("Cannot create VRF ", err.Error())
		return nil
	}

	err = c.vrfCreateBulk(ctx, vrf.FabricName, vrf)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("VRF create failed %v", err))
		dg.AddError("VRF create failed", err.Error())
		return nil
	}
	tflog.Info(ctx, fmt.Sprintf("Create Bulk VRF success ID %s", ID))
	//Get from NDFC
	outVrf := c.RscGetBulkVrf(ctx, dg, ID, true)
	if outVrf == nil {
		tflog.Error(ctx, "Failed to verify: Reading from NDFC after create failed")
		dg.AddError("Failed to verify", "Reading from NDFC after create failed")
		return nil
	}
	if ID != "" {
		outVrf.Id = types.StringValue(ID)
	}
	return outVrf

}

func (c NDFC) RscDeleteBulkVrf(ctx context.Context, dg *diag.Diagnostics, ID string, vrfBulk *resource_vrf_bulk.VrfBulkModel) {
	tflog.Info(ctx, fmt.Sprintf("VRF Bulk Delete request %s", ID))
	vrfs, err := c.vrfBulkIsPresent(ctx, ID, vrfBulk.FabricName.ValueString())
	if err != nil {
		tflog.Error(ctx, "Failed to Read existing VRFs")
		dg.AddError("VRF Read Failed", err.Error())
		return
	}
	fabricName, vrfsFromId := c.vrfBulkSplitID(ID)
	if len(vrfs) != len(vrfsFromId) {
		errString := fmt.Sprintf("Mismatch in VRF data: fabric %s Read %v, presented for delete %v", fabricName, vrfs, vrfsFromId)
		tflog.Error(ctx, errString)
		dg.AddError("VRF Delete Failed", errString)
		return
	}

	err = c.vrfBulkDelete(ctx, fabricName, vrfsFromId)
	if err != nil {
		errString := fmt.Sprintf("VRF delete failed on fabric %s vrfs %v", fabricName, vrfsFromId)
		tflog.Error(ctx, "VRF delete failed")
		dg.AddError(errString, err.Error())
		return
	}

	tflog.Info(ctx, fmt.Sprintf("VRF Bulk Delete Success %s", ID))

}

/*
1. Compare state and plan
2. Do bulk update for those VRFs changed
*/
func (c NDFC) RscUpdateBulkVrf(ctx context.Context,
	dg *diag.Diagnostics, ID string,
	vrfBulkPlan *resource_vrf_bulk.VrfBulkModel,
	vrfState *resource_vrf_bulk.VrfBulkModel) {

	actions, delVrfs := c.vrfBulkGetDiff(ctx, dg, vrfBulkPlan, vrfState)

	// Validate the Diff
	//Get the current VRFs from NDFC

	ndfcVRFs := c.vrfBulkGet(ctx, vrfBulkPlan.FabricName.ValueString())
	ndfcVRFs.CreateSearchMap()

	// Step 1 - Check if VRFs to update are present in NDFC
	for i := range actions["put"].Vrfs {
		name := actions["put"].Vrfs[i].VrfName
		_, ok := ndfcVRFs.VrfsMap[name]
		if !ok {
			errString := fmt.Sprintf("VRF %s to update is missing in NDFC", name)
			tflog.Error(ctx, errString)
			dg.AddError("In place update failed", errString)
			return
		}
	}

	// Step 2 - VRFs to delete are available
	for i := range delVrfs {
		_, ok := ndfcVRFs.VrfsMap[delVrfs[i]]
		if !ok {
			// is this error a big deal? VRFs to be deleted missing - so ignore??
			delVrfs = append(delVrfs[:i], delVrfs[i+1:]...)
			tflog.Error(ctx, fmt.Sprintf("VRF to DELETE %s is missing in NDFC", delVrfs[i]))
		}
	}
	// Step 3 - Check if VRFs (create-delete) are not present in NDFC

	for i := range actions["add"].Vrfs {
		name := actions["add"].Vrfs[i].VrfName
		_, ok := ndfcVRFs.VrfsMap[name]
		if ok {
			//VRF present in NDFC
			//Check if the VRF is marked for deletion as part of Delete-Create diff
			sort.Strings(delVrfs)
			if sort.SearchStrings(delVrfs, name) == 0 {
				tflog.Debug(ctx, "VRF marked for delete")
				continue
			}
			errString := fmt.Sprintf("VRF %s to Create is Already Present in NDFC. Update cancelled", name)
			tflog.Error(ctx, errString)
			dg.AddError("In place update failed", errString)
			return
		}
	}

	//Begin update - No rollback from here on

	if len(delVrfs) > 0 {
		tflog.Info(ctx, fmt.Sprintf("Deleting VRFs %v, as part of Bulk Update", delVrfs))
		err := c.vrfBulkDelete(ctx, vrfBulkPlan.FabricName.ValueString(), delVrfs)
		if err != nil {
			dg.AddError("Bulk Delete failed", err.Error())
		}
	} else {
		tflog.Info(ctx, "Nothing to delete")
	}
	if actions["put"].Vrfs.Len() > 0 {
		tflog.Info(ctx, "Modifying VRFs , as part of Bulk Update")
		c.vrfBulkUpdate(ctx, dg, actions["put"])
		if dg.HasError() {
			tflog.Info(ctx, fmt.Sprintf("Modifying VRFs Failed %v", dg.Errors()))
			return
		}
	} else {
		tflog.Info(ctx, "Nothing to modify")
	}

	if actions["add"].Vrfs.Len() > 0 {
		tflog.Info(ctx, "Adding VRFs , as part of Bulk Update")
		err := c.vrfCreateBulk(ctx, vrfBulkPlan.FabricName.ValueString(), actions["add"])
		if err != nil {
			dg.AddError("VRF create failed", err.Error())
			return
		}
	}

	if dg.HasError() {
		tflog.Error(ctx, "Error during update")
		return
	}

	newID := c.VrfBulkCreateID(actions["plan"])

	*(vrfBulkPlan) = *(c.RscGetBulkVrf(ctx, dg, newID, false))
}
