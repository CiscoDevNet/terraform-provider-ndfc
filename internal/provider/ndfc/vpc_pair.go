// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package ndfc

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"terraform-provider-ndfc/internal/provider/ndfc/api"
	"terraform-provider-ndfc/internal/provider/resources/resource_vpc_pair"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const ResourceVpcPair = "vpc_pair"

func (c *NDFC) RscReadVpcPair(ctx context.Context, resp *resource.ReadResponse, tf *resource_vpc_pair.VpcPairModel) {

	//  Rest API object for vPC Pair
	tflog.Debug(ctx, "RscReadVpcPair: Reading vPC Pair for type")
	vpcPairModel := c.rscGetVpcPair(ctx, tf)
	if vpcPairModel == nil {
		tflog.Error(ctx, "RscGetVpcPair: Failed to get vPC Pair")
		return
	}

	// Fill NDFC output data to terraform model
	vpcPairModel.Deploy = tf.GetModelData().Deploy
	rscCreateModelId(tf, vpcPairModel)
	log.Printf("[TRACE] Vpc pair model: SerialNumbers %v, Id %v", tf.SerialNumbers, tf.Id)

}
func (c *NDFC) RscDeleteVpcPair(ctx context.Context, dg *diag.Diagnostics, tf *resource_vpc_pair.VpcPairModel) {
	api := api.NewVpcPairAPI(c.GetLock(ResourceVpcPair), &c.apiClient)
	tflog.Debug(ctx, "RscDeleteVpcPair: Checking if vPC Pair is present")
	// Check if vPC Pair is already deleted
	vpcPairModel := c.rscGetVpcPair(ctx, tf)
	if vpcPairModel == nil {
		tflog.Error(ctx, "RscDeleteVpcPair: Vpc pair not present, might be already deleted")
		return
	}
	fabricName := vpcPairModel.PeerOneSwitchDetails.FabricName
	// Delete the vPC Pair if it is present
	api.VpcPairID = vpcPairModel.PeerOneId
	tflog.Debug(ctx, fmt.Sprint("RscDeleteVpcPair: Deleting vPC Pair with id: ", vpcPairModel.PeerOneId))

	res, err := api.Delete()
	if err != nil {
		tflog.Error(ctx, "Failed to destroy vPC Pair")
		dg.AddError("Failed to destroy vPC Pair", fmt.Sprintf("Error %v: %v", err, res.String()))
		return
	}

	// Check if the vPC Pair is deleted after some giving time
	time.Sleep(3 * time.Second)
	vpcPairModel = c.rscGetVpcPair(ctx, tf)
	if vpcPairModel != nil {
		dg.AddError("Failed to delete vPC Pair", fmt.Sprintf("Error %v: %v", err, res.String()))
		return
	}
	c.RscDeployVpcPair(ctx, dg, tf, fabricName)
	if dg.HasError() {
		return
	}

}

func (c *NDFC) RscCreateVpcPair(ctx context.Context, dg *diag.Diagnostics, tf *resource_vpc_pair.VpcPairModel) {
	err := c.rscCheckVpcPairRecommendations(ctx, tf)
	if err != nil {
		dg.AddError("vPC Pair creation failed", fmt.Sprintf("Error %v", err))
		return
	}

	api := api.NewVpcPairAPI(c.GetLock(ResourceVpcPair), &c.apiClient)
	nfdcVpcPairModel := tf.GetModelData()
	deploy := nfdcVpcPairModel.Deploy

	// Convert NDFC model to the json data
	payload, err := json.Marshal(nfdcVpcPairModel)
	if err != nil {
		tflog.Error(ctx, "RscCreateOrUpdate: RscCreateOrUpdate: Failed to marshal vPC Pair data")
		dg.AddError("Failed to marshal vPC Pair data", fmt.Sprintf("Error %v", err))
		return
	}

	// Check if the vPC Pair is already present
	nfdcVpcPairModel = c.rscGetVpcPair(ctx, tf)
	if nfdcVpcPairModel == nil {
		log.Printf("Vpc pair not present, creating new vPC Pair")
		res, err := api.Post(payload)
		if err != nil {
			tflog.Error(ctx, "RscCreateOrUpdate: Failed to create vPC Pair")
			dg.AddError("Failed to create vPC Pair", fmt.Sprintf("Error %v: %v", err, res.String()))
			return
		}
		// Check if the vPC Pair is present after creation
		nfdcVpcPairModel = c.rscGetVpcPair(ctx, tf)
		if nfdcVpcPairModel == nil {
			dg.AddError("Failed to create vPC Pair", "vPC Pair data empty in NDFC")
			return
		}
	} else {
		err := fmt.Errorf("vPC Pair already exists with serial numbers %s %s", nfdcVpcPairModel.PeerOneId, nfdcVpcPairModel.PeerTwoId)
		tflog.Error(ctx, "RscCreateVpcPair: Existing Vpc Pair present, destroy to proceed new one")
		dg.AddError("Failed to create vPC Pair", fmt.Sprintf("Error %v", err))
		return
	}
	fabricName := nfdcVpcPairModel.PeerOneSwitchDetails.FabricName
	log.Printf("fabric name %v", fabricName)
	if deploy {
		c.RscDeployVpcPair(ctx, dg, tf, fabricName)
		if dg.HasError() {
			return
		}
	}

	nfdcVpcPairModel.Deploy = deploy
	rscCreateModelId(tf, nfdcVpcPairModel)
	log.Printf("[TRACE] Vpc pair model: SerialNumbers %v, Id %v", tf.SerialNumbers, tf.Id)
}

func (c *NDFC) RscUpdateVpcPair(ctx context.Context, dg *diag.Diagnostics, tf *resource_vpc_pair.VpcPairModel) {

	err := c.rscCheckVpcPairRecommendations(ctx, tf)
	if err != nil {
		dg.AddError("vPC Pair update failed", fmt.Sprintf("Error %v", err))
		return
	}

	// Store the Deploy status and set it back, as it is not passed to NDFC
	nfdcVpcPairModel := tf.GetModelData()
	deploy := nfdcVpcPairModel.Deploy

	// Convert NDFC model to the json data
	payload, err := json.Marshal(nfdcVpcPairModel)
	if err != nil {
		tflog.Error(ctx, "RscCreateOrUpdate: RscCreateOrUpdate: Failed to marshal vPC Pair data")
		dg.AddError("Failed to marshal vPC Pair data", fmt.Sprintf("Error %v", err))
		return
	}

	// Directly call PUT , if it is update mode
	api := api.NewVpcPairAPI(c.GetLock(ResourceVpcPair), &c.apiClient)
	res, err := api.Put(payload)
	if err != nil {
		tflog.Error(ctx, "RscCreateOrUpdate: Failed to update vPC Pair")
		dg.AddError("Failed to update vPC Pair", fmt.Sprintf("Error %v: %v", err, res.String()))
		return
	}
	// Check if the vPC Pair is present after update
	nfdcVpcPairModel = c.rscGetVpcPair(ctx, tf)
	if nfdcVpcPairModel == nil {
		dg.AddError("Failed to create vPC Pair", "vPC Pair data empty in NDFC")
		return
	}
	fabricName := nfdcVpcPairModel.PeerOneSwitchDetails.FabricName
	if deploy {
		c.RscDeployVpcPair(ctx, dg, tf, fabricName)
		if dg.HasError() {
			return
		}
	}

	nfdcVpcPairModel.Deploy = deploy
	rscCreateModelId(tf, nfdcVpcPairModel)
	log.Printf("[TRACE] Vpc pair model: SerialNumbers %v, Id %v", tf.SerialNumbers, tf.Id)
}

func (c *NDFC) rscGetVpcPair(ctx context.Context,
	tf *resource_vpc_pair.VpcPairModel) *resource_vpc_pair.NDFCVpcPairModel {

	api := api.NewVpcPairAPI(c.GetLock(ResourceVpcPair), &c.apiClient)
	api.VpcPairID = tf.GetModelData().SerialNumbers[0]

	payload, erro := api.Get()
	modellist := []resource_vpc_pair.NDFCVpcPairModel{}

	if erro != nil || string(payload) == "[]" || payload == nil {
		tflog.Debug(ctx, "RscGetVpcPair: Failed to get vPC Pair")
		return nil
	}
	log.Printf("[TRACE] Vpc pair data: %v %v", string(payload), api.VpcPairID)
	err := json.Unmarshal(payload, &modellist)
	if err != nil {
		tflog.Error(ctx, "RscGetVpcPair: Failed to unmarshal vPC Pair data")
		return nil
	}
	vpcPairModel := rscValidateVpcPair(ctx, tf, modellist)
	return vpcPairModel
}

func rscCreateModelId(tf *resource_vpc_pair.VpcPairModel,
	vpcPairModel *resource_vpc_pair.NDFCVpcPairModel) {

	// Set the output data to terraform model
	tf.SetModelData(vpcPairModel)
	Id := vpcPairModel.PeerOneId + "/" + vpcPairModel.PeerTwoId
	tf.Id = types.StringValue(Id)
	log.Printf("Setting Id %v", tf.Id)

}

func rscValidateVpcPair(ctx context.Context,
	tf *resource_vpc_pair.VpcPairModel,
	vpcPairModelList []resource_vpc_pair.NDFCVpcPairModel) *resource_vpc_pair.NDFCVpcPairModel {
	// Validate the vPC Pair status
	nfdcVpcPairModel := tf.GetModelData()
	peerOneId := nfdcVpcPairModel.SerialNumbers[0]
	peerTwoId := nfdcVpcPairModel.SerialNumbers[1]

	for _, vpcPairModel := range vpcPairModelList {
		log.Printf("From Config: PeerOneId %v, PeerTwoId %v", peerOneId, peerTwoId)
		log.Printf("From NDFC: PeerOneId %v, PeerTwoId %v", vpcPairModel.PeerOneId, vpcPairModel.PeerTwoId)
		if (peerOneId == vpcPairModel.PeerOneId && peerTwoId == vpcPairModel.PeerTwoId) ||
			(peerOneId == vpcPairModel.PeerTwoId && peerTwoId == vpcPairModel.PeerOneId) {
			tflog.Debug(ctx, "RscValidateVpcPair: Vpc pair status is present and valid")
			return &vpcPairModel
		}
	}
	return nil
}
func (c *NDFC) rscCheckVpcPairRecommendations(ctx context.Context,
	tf *resource_vpc_pair.VpcPairModel) error {
	// Get the recommendations for the vPC Pair
	api := api.NewVpcPairAPI(c.GetLock(ResourceVpcPair), &c.apiClient)
	nfdcVpcPairModel := tf.GetModelData()
	recmdList := []resource_vpc_pair.NDFCVpcPairRecommendations{}

	api.VpcPairID = nfdcVpcPairModel.SerialNumbers[0]
	api.CheckStatus = make(map[string]bool)
	api.CheckStatus["recommendations"] = true
	api.VirtualPeerLink = *nfdcVpcPairModel.UseVirtualPeerlink

	payload, err := api.Get()
	api.CheckStatus["recommendations"] = false
	if err != nil {
		tflog.Error(ctx, "RscVpcPairRecommendations: Failed to get recommendations")
		err = fmt.Errorf("failed to get recommendations")
		return err
	}

	err = json.Unmarshal(payload, &recmdList)
	if err != nil {
		tflog.Error(ctx, "RscVpcPairRecommendations: Failed to unmarshal recommendations data")
		return err
	}
	for _, rec := range recmdList {
		log.Printf("SerialNumber %v, Recmd %v, RecReason %v", rec.SerialNumber, rec.Recommended, rec.RecommendationReason)
		if rec.SerialNumber == nfdcVpcPairModel.SerialNumbers[1] {
			if rec.Recommended {
				tflog.Debug(ctx, "NDFC Recommendation met")
				return nil
			} else {
				if rec.RecommendationReason == "Switches are not connected" {
					// This is a transient error, it will be really known when config apply is done.
					return nil
				} else {
					err = fmt.Errorf("%s (%s) : %s", nfdcVpcPairModel.SerialNumbers[1], rec.LogicalName, rec.RecommendationReason)
					tflog.Error(ctx, fmt.Sprintf("NDFC Recommendation not met for serial numbers  %s %s", nfdcVpcPairModel.SerialNumbers[0], nfdcVpcPairModel.SerialNumbers[1]))
					return err
				}
			}
		}
	}
	return nil
}
func (c *NDFC) RscImportVpcPairs(ctx context.Context,
	dg *diag.Diagnostics,
	tf *resource_vpc_pair.VpcPairModel) {
	// Get the vPC Pair data
	var err diag.Diagnostics
	id := tf.Id.ValueString()
	idSplit := strings.Split(id, ":")
	if len(idSplit) != 2 {
		dg.AddError("Import state failed", "ID should be in the format of serial number:serial number")
		return
	}
	tf.SerialNumbers, err = types.SetValue(types.StringType, []attr.Value{types.StringValue(idSplit[0]), types.StringValue(idSplit[1])})
	if err != nil {
		dg.AddError("Import state failed", "Failed to set serial numbers")
		return
	}
	vpcPairModel := c.rscGetVpcPair(ctx, tf)
	if vpcPairModel == nil {
		dg.AddError("Unable to get vpcPair", "vPC Pair data empty in NDFC")
		return
	}
	rscCreateModelId(tf, vpcPairModel)
}

/*
	    vPC pair does not have a module level deployment,
		so switch deploy is required
*/
func (c *NDFC) RscDeployVpcPair(ctx context.Context, dg *diag.Diagnostics, tf *resource_vpc_pair.VpcPairModel, fabricName string) {

	ndfcVpcPairModel := tf.GetModelData()
	serialNumbers := ndfcVpcPairModel.SerialNumbers
	c.SaveConfiguration(ctx, dg, fabricName)
	if dg.HasError() {
		return
	}
	c.DeployConfiguration(ctx, dg, fabricName, serialNumbers)
	if dg.HasError() {
		return
	}
}
