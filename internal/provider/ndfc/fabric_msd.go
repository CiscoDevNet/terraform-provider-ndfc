package ndfc

import (
	"context"
	"encoding/json"
	"fmt"
	"slices"
	"terraform-provider-ndfc/internal/provider/ndfc/api"
	"terraform-provider-ndfc/internal/provider/resources/resource_fabric_common"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func (m *NDFC) AddChildFabricsToMsd(ctx context.Context, dg *diag.Diagnostics, tf resource_fabric_common.FabricModel) {
	model := tf.GetModelData()
	if len(model.ChildFabrics) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("AddChildFabricsToMsd: child fabrics %v", model.ChildFabrics))
		m.ManageChildFabricsInMsd(ctx, dg, model.FabricName, model.ChildFabrics, api.MSD_OPERATION_ADD)
		return
	}
	tflog.Debug(ctx, "AddChildFabricsToMsd: No child fabrics to add")
}
func (m *NDFC) RemoveChildFabricsFromMsd(ctx context.Context, dg *diag.Diagnostics, tf resource_fabric_common.FabricModel) {
	model := tf.GetModelData()
	if len(model.ChildFabrics) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("RemoveChildFabricsFromMsd: child fabrics %v", model.ChildFabrics))
		m.ManageChildFabricsInMsd(ctx, dg, model.FabricName, model.ChildFabrics, api.MSD_OPERATION_REMOVE)
		if dg.HasError() {
			return
		}
		// Deploy the removed child fabrics to ensure configuration is applied
		for _, childFabric := range model.ChildFabrics {
			tflog.Info(ctx, fmt.Sprintf("RemoveChildFabricsFromMsd: Deploying removed child fabric %s", childFabric))
			m.RscDeployFabric(ctx, dg, childFabric)
			if dg.HasError() {
				tflog.Error(ctx, fmt.Sprintf("RemoveChildFabricsFromMsd: Failed to deploy removed child fabric %s", childFabric))
				return
			}
		}
		return
	}
	tflog.Debug(ctx, "RemoveChildFabricsFromMsd: No child fabrics to remove")
}

func (m *NDFC) UpdateChildFabricsToMsd(ctx context.Context, dg *diag.Diagnostics, plan resource_fabric_common.FabricModel, state resource_fabric_common.FabricModel) {

	p := plan.GetModelData()
	s := state.GetModelData()
	parentFabric := p.FabricName
	// Find fabrics to add (in plan but not in state)
	addList := make([]string, 0)
	for _, newList := range p.ChildFabrics {
		found := slices.Contains(s.ChildFabrics, newList)
		if !found {
			addList = append(addList, newList)
		}
	}

	// Find fabrics to delete (in state but not in plan)
	delList := make([]string, 0)
	for _, oldList := range s.ChildFabrics {
		found := slices.Contains(p.ChildFabrics, oldList)
		if !found {
			delList = append(delList, oldList)
		}
	}

	// Add new child fabrics
	if len(addList) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("UpdateChildFabricsToMsd: Adding child fabrics %v", addList))
		m.ManageChildFabricsInMsd(ctx, dg, parentFabric, addList, api.MSD_OPERATION_ADD)
	}

	// Remove old child fabrics
	if len(delList) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("UpdateChildFabricsToMsd: Removing child fabrics %v", delList))
		m.ManageChildFabricsInMsd(ctx, dg, parentFabric, delList, api.MSD_OPERATION_REMOVE)
	}
}

func (m *NDFC) ManageChildFabricsInMsd(ctx context.Context, dg *diag.Diagnostics, parentFabric string, childFabrics []string, op string) {

	fapi := api.NewFabricAPI(m.GetLock(ResourceFabrics), &m.apiClient)
	fapi.MsdOperation = op
	var msd resource_fabric_common.NdfcMsdFabricPayload

	// For add operations, validate that all child fabrics exist before proceeding
	if op == api.MSD_OPERATION_ADD {
		for _, childFabric := range childFabrics {
			if !m.FabricExists(ctx, childFabric) {
				tflog.Error(ctx, fmt.Sprintf("ManageChildFabricsInMsd: Child fabric %s does not exist", childFabric))
				dg.AddError("Child fabric does not exist",
					fmt.Sprintf("Cannot add fabric '%s' to MSD '%s': fabric '%s' does not exist. Please create the fabric first before adding it to the MSD.", childFabric, parentFabric, childFabric))
				return
			}
		}
	}

	for _, childFabric := range childFabrics {
		msd.DstFabricName = parentFabric
		msd.SrcFabricName = childFabric
		tflog.Debug(ctx, fmt.Sprintf("ManageChildFabricsInMsd: msd %v", msd))
		payload, err := json.Marshal(msd)
		if err != nil {
			tflog.Error(ctx, "ManageChildFabricsInMsd: Failed to marshal msd fabric data")
			dg.AddError("Failed to marshal msd fabric data", fmt.Sprintf("Error: %q", err.Error()))
			return
		}
		res, err := fapi.Post(payload)
		if err != nil {
			tflog.Error(ctx, "ManageChildFabricsInMsd: POST failed with payload %s", map[string]any{"Payload": payload})
			dg.AddError("Failed to add child fabric to msd", fmt.Sprintf("Error: %q %q", err.Error(), res.String()))
			return
		}
	}
}

func (m *NDFC) GetMsdChildFabricAssociations(ctx context.Context, dg *diag.Diagnostics, parentFabric string) []string {
	fapi := api.NewFabricAPI(m.GetLock(ResourceFabrics), &m.apiClient)
	fapi.MsdOperation = api.MSD_OPERATION_GET
	var fabricAscs []resource_fabric_common.NdfcMsdFabricAssociations
	res, err := fapi.Get()
	if err != nil {
		tflog.Error(ctx, "GetMsdFabricAssociations: GET failed")
		dg.AddError("Failed to get msd fabric associations", fmt.Sprintf("Error: %q", err.Error()))
		return nil
	}
	err = json.Unmarshal(res, &fabricAscs)
	if err != nil {
		tflog.Error(ctx, "GetMsdFabricAssociations: Failed to unmarshal response")
		dg.AddError("Failed to unmarshal msd fabric associations", fmt.Sprintf("Error: %q", err.Error()))
		return nil
	}
	childFabrics := make([]string, 0)
	for _, asc := range fabricAscs {
		if asc.FabricParent == parentFabric {
			childFabrics = append(childFabrics, asc.FabricName)
		}
	}
	return childFabrics
}

// CheckIsFabricMsdType checks the MSD type of a fabric.
// Returns:
//   - (true, "") if the fabric is an MSD parent
//   - (false, parentName) if the fabric is an MSD child (parentName is non-empty)
//   - (false, "") if the fabric is a regular (non-MSD) fabric
func (m *NDFC) CheckIsFabricMsdType(ctx context.Context, dg *diag.Diagnostics, fabricName string) (isMsdParent bool, parentFabric string) {
	// First check if this is an MSD parent fabric
	fType := m.GetFabricTemplateType(ctx, dg, fabricName)
	if dg.HasError() {
		return false, ""
	}

	if fType == ResourceVxlanMsdType {
		tflog.Debug(ctx, fmt.Sprintf("Fabric %s is MSD parent fabric", fabricName))
		return true, ""
	}

	// Check if this is an MSD child fabric
	fapi := api.NewFabricAPI(m.GetLock(ResourceFabrics), &m.apiClient)
	fapi.MsdOperation = api.MSD_OPERATION_GET
	var fabricAscs []resource_fabric_common.NdfcMsdFabricAssociations
	res, err := fapi.Get()
	if err != nil {
		tflog.Error(ctx, "GetMsdFabricAssociations: GET failed")
		dg.AddError("Failed to get msd fabric associations", fmt.Sprintf("Error: %q", err.Error()))
		return false, ""
	}
	err = json.Unmarshal(res, &fabricAscs)
	if err != nil {
		tflog.Error(ctx, "GetMsdFabricAssociations: Failed to unmarshal response")
		dg.AddError("Failed to unmarshal msd fabric associations", fmt.Sprintf("Error: %q", err.Error()))
		return false, ""
	}

	for _, asc := range fabricAscs {
		if asc.FabricParent != "" && asc.FabricParent != "None" {
			if asc.FabricName == fabricName {
				tflog.Debug(ctx, fmt.Sprintf("Fabric %s is MSD child fabric (parent: %s)", fabricName, asc.FabricParent))
				return false, asc.FabricParent
			}
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("Fabric %s is a regular (non-MSD) fabric", fabricName))
	return false, ""
}
