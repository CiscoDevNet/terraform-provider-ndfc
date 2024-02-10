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

// ID = fabricName/[vrf1{attachment1,attachment2,attachment3},vrf2{attachment1,attachment2,attachment3},vrf3{attachment1,attachment2,attachment3}]

func (c NDFC) VrfBulkCreateID(ndVrfs *resource_vrf_bulk.NDFCVrfBulkModel) string {
	uniqueID := "%s/%v"
	if ndVrfs != nil {
		fName := ndVrfs.FabricName
		vrf_list := make([]string, len(ndVrfs.Vrfs))
		for i, v := range ndVrfs.Vrfs {
			if len(v.AttachList) > 0 {
				serials := make([]string, len(v.AttachList))
				for j := range v.AttachList {
					serials[j] = v.AttachList[j].SerialNumber
					if v.AttachList[j].SerialNumber == "" {
						log.Panic("Serial number empty case")
					}
				}
				vrf_list[i] = fmt.Sprintf("%s{%s}", v.VrfName, strings.Join(serials, ","))
			} else {
				vrf_list[i] = v.VrfName
			}
		}
		vrf_list_str := "[" + strings.Join(vrf_list, ",") + "]"
		return fmt.Sprintf(uniqueID, fName, vrf_list_str)
	}
	return ""
}

func (c NDFC) VrfBulkSplitID(ID string) map[string][]string {
	// Split the ID into its components
	result := map[string][]string{}
	vrfs := make([]string, 0)
	components := strings.Split(ID, "/[")
	vrfData := components[1]
	start := 0
	for i := 0; i < len(vrfData); i++ {
		if vrfData[i] == '{' {
			vrfName := vrfData[start:i]
			vrfs = append(vrfs, vrfName)
			for j := i + 1; j < len(vrfData); j++ {
				if vrfData[j] == '}' {
					attachs := vrfData[i+1 : j]
					fmt.Println(attachs)
					attachments := strings.Split(attachs, ",")
					result[vrfName] = attachments
					i = j + 1
					start = j + 2
					break
				}
			}
		} else if vrfData[i] == ',' || vrfData[i] == ']' {
			vrfName := vrfData[start:i]
			vrfs = append(vrfs, vrfName)
			result[vrfName] = []string{}
			start = i + 1
		}
	}
	result["fabric"] = []string{components[0]}
	result["vrfs"] = vrfs
	return result
}

func (c NDFC) RscGetBulkVrf(ctx context.Context, dg *diag.Diagnostics, ID string, depMap *map[string][]string) *resource_vrf_bulk.VrfBulkModel {
	var filterMap map[string]bool
	tflog.Debug(ctx, fmt.Sprintf("RscGetBulkVrf entry fabirc %s", ID))

	filterMap = make(map[string]bool)
	fabricName, vrfs := c.vrfCreateFilterMap(ID, &filterMap)
	log.Printf("FilterMap: %v vrfs %v", filterMap, vrfs)
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
	attachOrder := c.VrfBulkSplitID(ID)

	ndVrfs.CreateSearchMap()
	// VRFs should be ordered as they appear in ID field if the Get
	// Except for create case - ideally they should come in the order of creation
	for i := range vrfs {
		vrf, ok := ndVrfs.VrfsMap[vrfs[i]]
		if ok {
			*vrf.Id = int64(i) //sort uses id - so put ascending values
		} else {
			tflog.Error(ctx, fmt.Sprintf("VRF %s missing in output", vrfs[i]))
			//log.Panicf("VRF %s missing in output", vrfs[i])
		}
	}

	if _, ok := (*depMap)["global"]; ok {
		//This cannot be validated - as we don't know if all deployments were ok
		ndVrfs.DeployAllAttachments = true
	}

	//Get Attachments
	va := c.RscGetVrfAttachments(ctx, dg, fabricName, vrfs)
	if va != nil {
		for i := range va.VrfAttachments {
			vrf, ok := ndVrfs.VrfsMap[va.VrfAttachments[i].VrfName]
			vrfLevelDep := false
			if ok {
				if vl, vlOk := (*depMap)[va.VrfAttachments[i].VrfName]; vlOk {
					if vl[0] != va.VrfAttachments[i].VrfName {
						vrf.DeployAttachments = (vrf.VrfStatus == "DEPLOYED")
						vrfLevelDep = true
					}
				}
				for j := range va.VrfAttachments[i].AttachList {
					if !va.VrfAttachments[i].AttachList[j].FilterThisValue {
						log.Printf("Attachment %s added to VRF %s", va.VrfAttachments[i].AttachList[j].SwitchSerialNo, va.VrfAttachments[i].VrfName)
						if !vrfLevelDep {
							if va.VrfAttachments[i].AttachList[j].AttachState == "DEPLOYED" {
								va.VrfAttachments[i].AttachList[j].DeployThisAttachment = true
							}
						}
						vrf.AttachList = append(vrf.AttachList, va.VrfAttachments[i].AttachList[j])
					}
				}
				if len(vrf.AttachList) > 0 {
					vrf.CreateSearchMap()
					for j := range attachOrder[vrf.VrfName] {
						attachEntry, found := vrf.AttachListMap[attachOrder[vrf.VrfName][j]]
						if found {
							log.Printf("Attachment Order %d: %s", j, attachOrder[vrf.VrfName][j])
							attachEntry.Id = new(int64)
							*attachEntry.Id = int64(j)
						} else {
							log.Panicf("RscGetBulkVrf: Attachment %s missing in output", attachOrder[vrf.VrfName][j])
						}
					}
					//Handle any entries missing in ID???
					for j := range vrf.AttachList {
						if vrf.AttachList[j].Id == nil {
							dg.AddWarning("Attachments mismatch", fmt.Sprintf("Attachment %s missing in ID", vrf.AttachList[j].SwitchSerialNo))
							vrf.AttachList[j].Id = new(int64)
							*vrf.AttachList[j].Id = int64(len(vrf.AttachList) + j)
						}
					}

					sort.Sort(vrf.AttachList)
				}
			} else {
				log.Panicf("RscGetBulkVrf: VRF %s missing in Attachment list", va.VrfAttachments[i].VrfName)
			}
		}
	}
	sort.Sort(ndVrfs.Vrfs)
	vModel := new(resource_vrf_bulk.VrfBulkModel)
	vModel.Id = types.StringValue(ID)
	d := vModel.SetModelData(&ndVrfs)
	if d != nil {
		dg.Append(d.Errors()...)
	}
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

	depMap := make(map[string][]string)

	// check if any of the vrfs exists
	err := c.vrfBulkCreateCheck(ctx, ID, vrfBulk)
	if err != nil {
		tflog.Error(ctx, "Cannot create VRF", map[string]interface{}{"Err": err})
		dg.AddError("Cannot create VRF ", err.Error())
		return nil
	}
	//Part 1: Create VRFs
	err = c.vrfCreateBulk(ctx, vrf.FabricName, vrf)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("VRF create failed %v", err))
		dg.AddError("VRF create failed", err.Error())
		return nil
	}
	tflog.Info(ctx, fmt.Sprintf("Create Bulk VRF success ID %s", ID))
	//Get from NDFC

	//Part 2: Create Attachments if any
	va := resource_vrf_bulk.NDFCVrfAttachmentsModel{}
	va.FabricName = vrf.FabricName
	va.VrfAttachments = make(resource_vrf_bulk.NDFCVrfAttachmentsValues, 0)
	vrfs := []string{}
	attachOrder := make(map[string][]string)

	if vrf.DeployAllAttachments {
		depMap["global"] = append(depMap["global"], "all")
	}

	for i := range vrf.Vrfs {
		vrfs = append(vrfs, vrf.Vrfs[i].VrfName)
		if vrf.Vrfs[i].DeployAttachments {
			depMap[vrf.Vrfs[i].VrfName] = append(depMap[vrf.Vrfs[i].VrfName], vrf.Vrfs[i].VrfName)
		}
		if len(vrf.Vrfs[i].AttachList) > 0 {
			attachOrder[vrf.Vrfs[i].VrfName] = make([]string, len(vrf.Vrfs[i].AttachList))
			vrfAttachVal := new(resource_vrf_bulk.NDFCVrfAttachmentsValue)
			vrfAttachVal.VrfName = vrf.Vrfs[i].VrfName
			vrfAttachVal.DeployAllAttachments = vrf.Vrfs[i].DeployAttachments

			log.Printf("DeployAllAttachments %s/%v", vrf.Vrfs[i].VrfName, vrf.Vrfs[i].DeployAttachments)
			vrfAttachVal.AttachList = vrf.Vrfs[i].AttachList

			va.VrfAttachments = append(va.VrfAttachments, *vrfAttachVal)
			for j := range vrf.Vrfs[i].AttachList {
				attachOrder[vrf.Vrfs[i].VrfName][j] = vrf.Vrfs[i].AttachList[j].SerialNumber
				if vrf.Vrfs[i].AttachList[j].DeployThisAttachment {
					depMap[vrf.Vrfs[i].VrfName] = append(depMap[vrf.Vrfs[i].VrfName], vrf.Vrfs[i].AttachList[j].SerialNumber)
				}
			}

		}
	}
	if len(va.VrfAttachments) > 0 {
		err := c.RscCreateVrfAttachments(ctx, dg, &va)
		if err != nil {
			tflog.Error(ctx, "VRF Attachments create failed")
			tflog.Error(ctx, "Rolling back the configurations...delete VRFs")
			err := c.vrfBulkDelete(ctx, va.FabricName, vrfs)
			if err != nil {
				tflog.Error(ctx, "VRF Attachments Delete failed, rollback failed")
				dg.AddError("Rollback vrf Delete failed", err.Error())
				return nil
			}
			return nil
		}
	}

	outVrf := c.RscGetBulkVrf(ctx, dg, ID, &depMap)
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
	results := c.VrfBulkSplitID(ID)
	fabricName := results["fabric"][0]
	vrfsFromId := results["vrfs"]
	if len(vrfs) != len(vrfsFromId) {
		errString := fmt.Sprintf("Mismatch in VRF data: fabric %s Read %v, from ID %v", fabricName, vrfs, vrfsFromId)
		tflog.Error(ctx, errString)
		dg.AddWarning(("Mismatch in VRF data"), errString)
		vrfsFromId = vrfs
	}

	delVrf := vrfBulk.GetModelData()
	delVA := resource_vrf_bulk.NDFCVrfAttachmentsModel{}
	delVA.FabricName = delVrf.FabricName
	for i := range delVrf.Vrfs {
		if len(delVrf.Vrfs[i].AttachList) > 0 {
			vrfAttachVal := new(resource_vrf_bulk.NDFCVrfAttachmentsValue)
			vrfAttachVal.VrfName = delVrf.Vrfs[i].VrfName
			vrfAttachVal.AttachList = delVrf.Vrfs[i].AttachList
			for j := range vrfAttachVal.AttachList {
				vrfAttachVal.AttachList[j].Deployment = "false"
			}
			delVA.VrfAttachments = append(delVA.VrfAttachments, *vrfAttachVal)
		}
	}
	if len(delVA.VrfAttachments) > 0 {
		err := c.RscDeleteVrfAttachments(ctx, dg, &delVA)
		if err != nil {
			tflog.Error(ctx, "VRF Attachments delete failed")
			dg.AddError("VRF Attachments delete failed", err.Error())
			return
		}
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

	actions := c.vrfBulkGetDiff(ctx, dg, vrfBulkPlan, vrfState)

	// Validate the Diff
	//Get the current VRFs from NDFC

	delVrfs := actions["del"].([]string)
	//deployVrfs := actions["deploy"].([]string)

	putVrfs := actions["put"].(*resource_vrf_bulk.NDFCVrfBulkModel)
	newVrfs := actions["add"].(*resource_vrf_bulk.NDFCVrfBulkModel)
	plan := actions["plan"].(*resource_vrf_bulk.NDFCVrfBulkModel)
	state := actions["state"].(*resource_vrf_bulk.NDFCVrfBulkModel)

	ndfcVRFs := c.vrfBulkGet(ctx, vrfBulkPlan.FabricName.ValueString())
	ndfcVRFs.CreateSearchMap()

	// Step 1 - Check if VRFs to update are present in NDFC
	for i := range putVrfs.Vrfs {
		name := putVrfs.Vrfs[i].VrfName
		_, ok := ndfcVRFs.VrfsMap[name]
		if !ok {
			errString := fmt.Sprintf("VRF %s to update is missing in NDFC", name)
			tflog.Error(ctx, errString)
			dg.AddError("In place update failed", errString)
			return
		}
	}

	updateVA := resource_vrf_bulk.NDFCVrfAttachmentsModel{}
	updateVA.FabricName = vrfBulkPlan.FabricName.ValueString()
	updateVA.VrfAttachments = make(resource_vrf_bulk.NDFCVrfAttachmentsValues, 0)

	// Step 2 - VRFs to delete are available
	for i := range delVrfs {
		vrfEntry, ok := ndfcVRFs.VrfsMap[delVrfs[i]]
		if !ok {
			// is this error a big deal? VRFs to be deleted missing - so ignore??
			delVrfs = append(delVrfs[:i], delVrfs[i+1:]...)
			tflog.Error(ctx, fmt.Sprintf("VRF to DELETE %s is missing in NDFC", delVrfs[i]))
		} else {
			//Check if attachments are present
			if len(vrfEntry.AttachList) > 0 {
				vrfAttach := new(resource_vrf_bulk.NDFCVrfAttachmentsValue)
				vrfAttach.VrfName = vrfEntry.VrfName
				vrfAttach.AttachList = vrfEntry.AttachList
				updateVA.VrfAttachments = append(updateVA.VrfAttachments, *vrfAttach)
			}

		}
	}
	// Step 3 - Check if VRFs (create-delete) are not present in NDFC

	for i := range newVrfs.Vrfs {
		name := newVrfs.Vrfs[i].VrfName
		_, ok := ndfcVRFs.VrfsMap[name]
		if ok {
			//VRF present in NDFC
			//Check if the VRF is marked for deletion as part of Delete-Create diff
			sort.Strings(delVrfs)
			if sort.SearchStrings(delVrfs, name) == 0 {
				tflog.Debug(ctx, "VRF marked for delete", map[string]interface{}{"vrfName": name})
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
		// Check and delete attachments
		err := c.RscDeleteVrfAttachments(ctx, dg, &updateVA)
		if err != nil {
			tflog.Error(ctx, "VRF Attachments delete failed")
			dg.AddError("VRF Attachments delete failed", err.Error())
			return
		}
		tflog.Info(ctx, fmt.Sprintf("Deleting VRFs %v, as part of Bulk Update", delVrfs))
		err = c.vrfBulkDelete(ctx, vrfBulkPlan.FabricName.ValueString(), delVrfs)
		if err != nil {
			dg.AddError("Bulk Delete failed", err.Error())
		}
	} else {
		tflog.Info(ctx, "Nothing to delete")
	}
	if putVrfs.Vrfs.Len() > 0 {
		tflog.Info(ctx, "Modifying VRFs , as part of Bulk Update")
		c.vrfBulkUpdate(ctx, dg, putVrfs)
		if dg.HasError() {
			tflog.Info(ctx, fmt.Sprintf("Modifying VRFs Failed %v", dg.Errors()))
			return
		}
	} else {
		tflog.Info(ctx, "Nothing to modify")
	}

	if newVrfs.Vrfs.Len() > 0 {
		tflog.Info(ctx, "Adding VRFs , as part of Bulk Update")
		err := c.vrfCreateBulk(ctx, vrfBulkPlan.FabricName.ValueString(), newVrfs)
		if err != nil {
			dg.AddError("VRF create failed", err.Error())
			return
		}
	}
	//Deal with attachments
	updateVA.VrfAttachments = make(resource_vrf_bulk.NDFCVrfAttachmentsValues, 0)
	copyVrfAttachments(plan, &updateVA)
	stateVA := resource_vrf_bulk.NDFCVrfAttachmentsModel{}
	copyVrfAttachments(state, &stateVA)
	c.RscUpdateVrfAttachments(ctx, dg, &updateVA, &stateVA)
	if dg.HasError() {
		tflog.Error(ctx, "Error during update")
		return
	}
	newID := c.VrfBulkCreateID(plan)
	depMap := make(map[string][]string)
	if plan.DeployAllAttachments {
		depMap["global"] = append(depMap["global"], "all")
	}
	for i := range plan.Vrfs {
		if plan.Vrfs[i].DeployAttachments {
			depMap[plan.Vrfs[i].VrfName] = append(depMap[plan.Vrfs[i].VrfName], plan.Vrfs[i].VrfName)
		}
		for j := range plan.Vrfs[i].AttachList {
			if plan.Vrfs[i].AttachList[j].DeployThisAttachment {
				depMap[plan.Vrfs[i].VrfName] = append(depMap[plan.Vrfs[i].VrfName], plan.Vrfs[i].AttachList[j].SerialNumber)
			}
		}
	}
	*(vrfBulkPlan) = *(c.RscGetBulkVrf(ctx, dg, newID, &depMap))
}

func copyVrfAttachments(src *resource_vrf_bulk.NDFCVrfBulkModel, dst *resource_vrf_bulk.NDFCVrfAttachmentsModel) {
	dst.FabricName = src.FabricName
	dst.DeployAllAttachments = src.DeployAllAttachments
	dst.VrfAttachments = make(resource_vrf_bulk.NDFCVrfAttachmentsValues, len(src.Vrfs))
	for i := range src.Vrfs {
		if len(src.Vrfs[i].AttachList) > 0 {
			vrfAttach := new(resource_vrf_bulk.NDFCVrfAttachmentsValue)
			vrfAttach.VrfName = src.Vrfs[i].VrfName
			vrfAttach.DeployAllAttachments = src.Vrfs[i].DeployAttachments
			vrfAttach.AttachList = src.Vrfs[i].AttachList
			for j := range vrfAttach.AttachList {
				vrfAttach.AttachList[j].FabricName = src.FabricName
				vrfAttach.AttachList[j].VrfName = src.Vrfs[i].VrfName
			}
			dst.VrfAttachments = append(dst.VrfAttachments, *vrfAttach)
		}
	}
}
