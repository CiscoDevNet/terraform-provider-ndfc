package ndfc

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strings"
	rva "terraform-provider-ndfc/internal/provider/resources/resource_vrf_attachments"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const UrlVrfAttachmentsCreate = "/lan-fabric/rest/top-down/v2/fabrics/%s/vrfs/attachments"
const UrlVrfAttachmentsGet = "/lan-fabric/rest/top-down/fabrics/%s/vrfs/attachments"
const UrlQP = "?%s=%s"
const UrlVrfAttachmentsDeploy = "/lan-fabric/rest/top-down/v2/vrfs/deploy"

func (c NDFC) vrfAttachmentsGet(ctx context.Context, fabricName string, vrfs []string) ([]byte, error) {

	url := fmt.Sprintf(UrlVrfAttachmentsGet, fabricName)
	if len(vrfs) > 0 {
		qp := fmt.Sprintf(UrlQP, "vrf-names", strings.Join(vrfs, ","))
		url += qp
	}
	tflog.Debug(ctx, fmt.Sprintf("vrfAttachmentsGet: url %s", url))
	c.GetLock(ResourceVrfAttachments).Lock()
	defer c.GetLock(ResourceVrfAttachments).Unlock()
	res, err := c.apiClient.GetRawJson(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c NDFC) vrfAttachmentsIsPresent(ctx context.Context, dg *diag.Diagnostics, fabricName string, vrfs []string) bool {
	tflog.Debug(ctx, "vrfAttachmentsIsPresent: Entering")

	res, err := c.vrfAttachmentsGet(ctx, fabricName, vrfs)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("vrfAttachmentsIsPresent: Error getting VRF Attachments %s", err.Error()))
		dg.AddError("Error in getting VRF Attachments from NDFC", err.Error())
		return false
	}
	ndVrfs := rva.NDFCVrfAttachmentsModel{}
	err = json.Unmarshal(res, &ndVrfs.VrfAttachments)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("vrfAttachmentsIsPresent: Error unmarshalling VRF Attachments %s", err.Error()))
		log.Printf("vrfAttachmentsIsPresent: Error unmarshalling VRF Attachments %s", string(res))
		dg.AddError("Error in Payload format,when retrieving attachments from NDFC", err.Error())
		return false
	}
	if len(ndVrfs.VrfAttachments) > 0 {
		tflog.Info(ctx, fmt.Sprintf("vrfAttachmentsIsPresent: VRF Attachments exists on %s", fabricName))
		// Check if they were created and not implicit attachments present in NDFC GET
		for i := range ndVrfs.VrfAttachments {
			for j := range ndVrfs.VrfAttachments[i].AttachList {
				if ndVrfs.VrfAttachments[i].AttachList[j].Attached != nil &&
					*(ndVrfs.VrfAttachments[i].AttachList[j].Attached) {
					tflog.Info(ctx, fmt.Sprintf("vrfAttachmentsIsPresent: VRF Attachments were created on %s", fabricName))
					dg.AddError("VRF Attachments were created on "+fabricName,
						fmt.Sprintf("VRF %s Attachment {Switch: %s, Attached: %v, State:%s}",
							ndVrfs.VrfAttachments[i].VrfName,
							ndVrfs.VrfAttachments[i].AttachList[j].SwitchName,
							*ndVrfs.VrfAttachments[i].AttachList[j].Attached,
							ndVrfs.VrfAttachments[i].AttachList[j].AttachState))
					return true
				}
			}
		}
	}
	tflog.Info(ctx, fmt.Sprintf("vrfAttachmentsIsPresent: VRF Attachments does not exist on %s", fabricName))
	return false
}

func (c NDFC) vrfAttachmentsPost(ctx context.Context, va *rva.NDFCVrfAttachmentsModel) error {
	data, err := json.Marshal(va.VrfAttachments)
	if err != nil {
		return err
	}
	log.Println("Data to be posted", string(data))
	c.GetLock(ResourceVrfAttachments).Lock()
	defer c.GetLock(ResourceVrfAttachments).Unlock()
	res, err := c.apiClient.Post(fmt.Sprintf(UrlVrfAttachmentsCreate, va.FabricName), string(data))
	if err != nil {
		return err
	}
	tflog.Info(ctx, fmt.Sprintf("vrfAttachmentsCreate: Success res : %v", res.Str))
	return nil
}

func (c NDFC) getVrfAttachments(ctx context.Context, dg *diag.Diagnostics,
	Id string, in *rva.NDFCVrfAttachmentsModel) *rva.VrfAttachmentsModel {
	tflog.Debug(ctx, fmt.Sprintf("getVrfAttachments: Entering Id %s", Id))
	fabricName, vrfs := VrfAttachmentsSplitID(Id)
	log.Printf("getVrfAttachments:%+v", *in)
	tflog.Debug(ctx, fmt.Sprintf("getVrfAttachments: fabricName %s vrfs %v", fabricName, vrfs))

	// Get the VRF Attachments
	res, err := c.vrfAttachmentsGet(ctx, fabricName, vrfs)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("getVrfAttachments: Error getting VRF Attachments %s", err.Error()))
		return nil
	}
	tflog.Debug(ctx, fmt.Sprintf("getVrfAttachments: data read from NDFC: %s", string(res)))

	ndVrfs := rva.NDFCVrfAttachmentsModel{}
	err = json.Unmarshal(res, &ndVrfs.VrfAttachments)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("getVrfAttachments: Error unmarshalling VRF Attachments %s", err.Error()))
		log.Printf("getVrfAttachments: Error unmarshalling VRF Attachments %s", string(res))
		return nil
	}
	// Filter out the implicit attachments
	for i := range ndVrfs.VrfAttachments {
		skip := 0
		for j := range ndVrfs.VrfAttachments[i].AttachList {
			if ndVrfs.VrfAttachments[i].AttachList[j].Attached != nil &&
				!(*(ndVrfs.VrfAttachments[i].AttachList[j].Attached)) {
				//This is auto entry, not created; skip it
				ndVrfs.VrfAttachments[i].AttachList[j].FilterThisValue = true
				log.Printf("getVrfAttachments: Filtering out implicit attachment %s/{%s}",
					ndVrfs.VrfAttachments[i].VrfName,
					ndVrfs.VrfAttachments[i].AttachList[j].SwitchName)
				skip++
			}
		}
		if skip == len(ndVrfs.VrfAttachments[i].AttachList) {
			//All entries are implicit, skip the VRF
			log.Printf("getVrfAttachments: Filtering out implicit VRF %s", ndVrfs.VrfAttachments[i].VrfName)
			ndVrfs.VrfAttachments[i].FilterThisValue = true
		}
	}

	ndVrfs.CreateSearchMap()
	for i := range vrfs {
		if _, found := ndVrfs.VrfAttachmentsMap[vrfs[i]]; found {
			ndVrfs.VrfAttachmentsMap[vrfs[i]].Id = new(int64)
			log.Printf("getVrfAttachments: Setting ID for %s to %d", vrfs[i], i)
			*(ndVrfs.VrfAttachmentsMap[vrfs[i]].Id) = int64(i)
		}
	}
	sort.Sort(&ndVrfs.VrfAttachments)

	//Fill all control params from input, that are not in NDFC payload
	// - here deployment
	in.CreateSearchMap()
	ndVrfs.DeployAllAttachments = in.DeployAllAttachments

	for i := range ndVrfs.VrfAttachments {
		if ndVrfs.VrfAttachments[i].Id != nil {
			log.Printf("getVrfAttachments: Sorted ID for %s is %d ", ndVrfs.VrfAttachments[i].VrfName, *(ndVrfs.VrfAttachments[i].Id))
		}
		vrfEntry, found := in.VrfAttachmentsMap[ndVrfs.VrfAttachments[i].VrfName]
		if found {
			ndVrfs.VrfAttachments[i].DeployAllAttachments = vrfEntry.DeployAllAttachments
			vrfEntry.CreateSearchMap()
		} else {
			log.Printf("This is not expected, VRF entry %s not found in input", ndVrfs.VrfAttachments[i].VrfName)
		}

		for j := range ndVrfs.VrfAttachments[i].AttachList {
			attachEntry, found := vrfEntry.AttachListMap[ndVrfs.VrfAttachments[i].AttachList[j].SerialNumber]
			if found {
				ndVrfs.VrfAttachments[i].AttachList[j].DeployThisAttachment = attachEntry.DeployThisAttachment
			} else {
				log.Printf("This is not expected, Attachment entry %s{%s} not found in input", ndVrfs.VrfAttachments[i].VrfName, ndVrfs.VrfAttachments[i].AttachList[j].SerialNumber)
			}

			if ndVrfs.VrfAttachments[i].AttachList[j].VlanId != nil {
				log.Printf("getVrfAttachments: Sorted VlanId for %s is %v", ndVrfs.VrfAttachments[i].VrfName, *ndVrfs.VrfAttachments[i].AttachList[j].VlanId)
			}
			if ndVrfs.VrfAttachments[i].AttachList[j].Vlan != nil {
				log.Printf("getVrfAttachments: Sorted Vlan for %s is %v", ndVrfs.VrfAttachments[i].VrfName, *ndVrfs.VrfAttachments[i].AttachList[j].Vlan)
			}

			//log.Printf("getVrfAttachments: Sorted VlanId for %s is %v %v", ndVrfs.VrfAttachments[i].VrfName, *ndVrfs.VrfAttachments[i].AttachList[j].VlanId, *ndVrfs.VrfAttachments[i].AttachList[j].Vlan)
		}
	}
	ndVrfs.FabricName = fabricName
	ret := rva.VrfAttachmentsModel{}
	diagErr := ret.SetModelData(&ndVrfs)
	if diagErr.HasError() {
		tflog.Error(ctx, fmt.Sprintf("getVrfAttachments: Error setting model data %v", diagErr.Errors()))
		return nil
	}
	ret.Id = types.StringValue(Id)
	//Set entries that are not in Payload

	return &ret
}

func (c NDFC) diffVrfAttachments(ctx context.Context, plan *rva.VrfAttachmentsModel,
	state *rva.VrfAttachmentsModel) (
	*rva.NDFCVrfAttachmentsModel, string) {

	tflog.Debug(ctx, "diffVrfAttachments: Entering")
	planData := plan.GetModelData()
	stateData := state.GetModelData()
	ID, _ := c.VrfAttachmentsCreateID(planData)
	updateVA := new(rva.NDFCVrfAttachmentsModel)
	updateVA.VrfAttachmentsMap = make(map[string]*rva.NDFCVrfAttachmentsValue)
	updateVA.FabricName = planData.FabricName
	addToUpdate := func(vrf string, x ...rva.NDFCAttachListValue) {
		if len(updateVA.VrfAttachments) == 0 {
			updateVA.VrfAttachments = make(rva.NDFCVrfAttachmentsValues, 0)
		}
		for i := range x {
			x[i].VrfName = vrf
			x[i].FilterThisValue = false
			x[i].FabricName = planData.FabricName
			/* NDFCBUG: throws 500 Server error if vlan field is empty */
			/* Setting to -1 to avoid this - UI does the same      */
			if x[i].Vlan == nil {
				x[i].Vlan = new(rva.Int64Custom)
				*x[i].Vlan = rva.Int64Custom(-1)
			}
		}
		vrfAttachEntry, found := updateVA.VrfAttachmentsMap[vrf]
		if !found {
			vrfAttachEntry = new(rva.NDFCVrfAttachmentsValue)
			vrfAttachEntry.VrfName = vrf
			vrfAttachEntry.AttachList = make(rva.NDFCAttachListValues, 0)
			vrfAttachEntry.AttachList = append(vrfAttachEntry.AttachList, x...)
			updateVA.VrfAttachments = append(updateVA.VrfAttachments, *vrfAttachEntry)
			updateVA.VrfAttachmentsMap[vrf] = &updateVA.VrfAttachments[len(updateVA.VrfAttachments)-1]
		} else {
			vrfAttachEntry.AttachList = append(vrfAttachEntry.AttachList, x...)
		}
		vrfAttachEntry.CreateSearchMap()
	}

	planData.CreateSearchMap()
	stateData.CreateSearchMap()

	for i := range planData.VrfAttachments {
		stateVA, found := stateData.VrfAttachmentsMap[planData.VrfAttachments[i].VrfName]
		if !found {
			//newAttachments[planData.VrfAttachments[i].VrfName] = planData.VrfAttachments[i].AttachList
			tflog.Debug(ctx, fmt.Sprintf("diffVrfAttachments: New VRF %s in plan", planData.VrfAttachments[i].VrfName))
			for j := range planData.VrfAttachments[i].AttachList {
				//attach
				planData.VrfAttachments[i].AttachList[j].Deployment = "true"
			}
			addToUpdate(planData.VrfAttachments[i].VrfName, planData.VrfAttachments[i].AttachList...)
			//New VRF entry in plan
		} else {
			stateVA.FilterThisValue = true
			//Existing VRF entry in plan
			tflog.Debug(ctx, fmt.Sprintf("diffVrfAttachments: Existing VRF %s in plan",
				planData.VrfAttachments[i].VrfName))
			//Check if there are new attachments
			planData.VrfAttachments[i].CreateSearchMap()
			for j := range planData.VrfAttachments[i].AttachList {
				stateAttachment, found := stateVA.AttachListMap[planData.VrfAttachments[i].AttachList[j].SerialNumber]
				if !found {
					//New attachment in plan list
					tflog.Debug(ctx, fmt.Sprintf("diffVrfAttachments: New attachment %s/%s in plan",
						planData.VrfAttachments[i].VrfName,
						planData.VrfAttachments[i].AttachList[j].SerialNumber))
					//attach
					planData.VrfAttachments[i].AttachList[j].Deployment = "true"
					addToUpdate(planData.VrfAttachments[i].VrfName, planData.VrfAttachments[i].AttachList[j])
				} else {
					stateAttachment.FilterThisValue = true
					//Check if parameters are different
					retVal := planData.VrfAttachments[i].AttachList[j].DeepEqual(*stateAttachment)
					if retVal != rva.ValuesDeeplyEqual {
						//Modified attachment in plan list
						tflog.Debug(ctx, fmt.Sprintf("diffVrfAttachments: Modified attachment %s/%s in plan",
							planData.VrfAttachments[i].VrfName,
							planData.VrfAttachments[i].AttachList[j].SerialNumber))
						//attach
						planData.VrfAttachments[i].AttachList[j].Deployment = "true"
						addToUpdate(planData.VrfAttachments[i].VrfName, planData.VrfAttachments[i].AttachList[j])

					}
				}
			}
		}
	}
	for i := range stateData.VrfAttachments {
		if stateData.VrfAttachments[i].FilterThisValue {
			//Look inside attachlist
			for j := range stateData.VrfAttachments[i].AttachList {
				if !stateData.VrfAttachments[i].AttachList[j].FilterThisValue {
					//attachment to be deleted as its missing in plan
					tflog.Debug(ctx, fmt.Sprintf("diffVrfAttachments: To be Deleted attachment %s",
						stateData.VrfAttachments[i].AttachList[j].SerialNumber))
					//Detach
					stateData.VrfAttachments[i].AttachList[j].Deployment = "false"
					addToUpdate(stateData.VrfAttachments[i].VrfName, stateData.VrfAttachments[i].AttachList[j])
				}
			}

		} else {
			tflog.Info(ctx, fmt.Sprintf("diffVrfAttachments: To be Deleted VrfAttachment %s", stateData.VrfAttachments[i].VrfName))
			// VrfAttachments in state that are not in plan - to be detached
			for j := range stateData.VrfAttachments[i].AttachList {
				tflog.Debug(ctx, fmt.Sprintf("diffVrfAttachments: To be Deleted attachment %s",
					stateData.VrfAttachments[i].AttachList[j].SerialNumber))
				//Detach
				stateData.VrfAttachments[i].AttachList[j].Deployment = "false"
			}
			addToUpdate(stateData.VrfAttachments[i].VrfName, stateData.VrfAttachments[i].AttachList...)
		}
	}
	return updateVA, ID
}

func (c NDFC) vrfAttachmentsDeploy(ctx context.Context, dg *diag.Diagnostics, va *rva.NDFCVrfAttachmentsModel) error {

	deploy_map := make(map[string][]string)
	deployVA := new(rva.NDFCVrfAttachmentsModel)
	deployVA.VrfAttachments = make(rva.NDFCVrfAttachmentsValues, 0)
	deployVA.FabricName = va.FabricName
	deployRequired := false

	tflog.Debug(ctx, "vrfAttachmentsDeploy: Entering")

	if va.DeployAllAttachments {
		deployRequired = true
		//Deploy all attachments
		for i := range va.VrfAttachments {
			for j := range va.VrfAttachments[i].AttachList {
				deploy_map[va.VrfAttachments[i].AttachList[j].SerialNumber] =
					append(deploy_map[va.VrfAttachments[i].AttachList[j].SerialNumber], va.VrfAttachments[i].VrfName)
			}
		}
	} else {
		// Deploy VRFs in the resource if flag is set
		for i := range va.VrfAttachments {
			if va.VrfAttachments[i].DeployAllAttachments {
				deployRequired = true
				//Deploy all attachments in the VRF
				for j := range va.VrfAttachments[i].AttachList {
					deploy_map[va.VrfAttachments[i].AttachList[j].SerialNumber] =
						append(deploy_map[va.VrfAttachments[i].AttachList[j].SerialNumber], va.VrfAttachments[i].VrfName)
				}
			} else {
				//Deploy specific attachments in the VRF
				for j := range va.VrfAttachments[i].AttachList {
					if va.VrfAttachments[i].AttachList[j].DeployThisAttachment {
						deployRequired = true
						deploy_map[va.VrfAttachments[i].AttachList[j].SerialNumber] =
							append(deploy_map[va.VrfAttachments[i].AttachList[j].SerialNumber], va.VrfAttachments[i].VrfName)
					}
				}

			}
		}
	}

	if !deployRequired {
		tflog.Info(ctx, "vrfAttachmentsDeploy: No attachments to deploy")
		return nil
	}

	deploy_post_payload := "{"
	for k, v := range deploy_map {
		deploy_post_payload += "\"" + k + "\":\""
		deploy_post_payload += strings.Join(v, ",")
		deploy_post_payload += "\","
	}
	deploy_post_payload += "}"
	deploy_post_payload = strings.ReplaceAll(deploy_post_payload, ",}", "}")

	tflog.Info(ctx, fmt.Sprintf("vrfAttachmentsDeploy: Deploying Attachments %s", deploy_post_payload))
	c.GetLock(ResourceVrfAttachments).Lock()
	defer c.GetLock(ResourceVrfAttachments).Unlock()
	res, err := c.apiClient.Post(UrlVrfAttachmentsDeploy, deploy_post_payload)
	if err != nil {
		return err
	}
	tflog.Info(ctx, fmt.Sprintf("vrfAttachmentsDeploy: Success res : %v", res.Str))
	return nil
}
