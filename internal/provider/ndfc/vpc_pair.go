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
const ErrVpcPairNotFound = "vPC Pair data not found"

func (c *NDFC) RscReadVpcPair(ctx context.Context, resp *resource.ReadResponse, tf *resource_vpc_pair.VpcPairModel) error {

	//  Rest API object for vPC Pair
	tflog.Debug(ctx, "Reading vPC Pair")
	vpcPairModel, err := c.rscGetVpcPair(ctx, tf)

	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("RscReadVpcPair: Err in getting vPC Pair from NDFC: %v", err))
		resp.Diagnostics.AddError("Failed to get vPC Pair", fmt.Sprintf("Error %v", err))
		return err
	}
	tflog.Debug(ctx, fmt.Sprintf("RscReadVpcPair: vPC Pair model: %v", vpcPairModel))
	if vpcPairModel == nil {
		tflog.Error(ctx, "RscGetVpcPair: Failed to get vPC Pair")
		resp.Diagnostics.AddError("Failed to get vPC Pair", "vPC Pair data empty in NDFC")
		err = fmt.Errorf("%s", ErrVpcPairNotFound)
		return err
	}
	// Fill NDFC output data to terraform model
	vpcPairModel.Deploy = tf.GetModelData().Deploy
	rscCreateTfModelAndId(tf, vpcPairModel)
	log.Printf("[TRACE] vPC pair model: SerialNumbers %v, Id %v", tf.SerialNumbers, tf.Id)
	return nil

}

func (c *NDFC) RscDeleteVpcPair(ctx context.Context, dg *diag.Diagnostics, tf *resource_vpc_pair.VpcPairModel) {
	api := api.NewVpcPairAPI(c.GetLock(ResourceVpcPair), &c.apiClient)
	tflog.Debug(ctx, "RscDeleteVpcPair: Checking if vPC Pair is present")
	// Check if vPC Pair is already deleted
	vpcPairModel, _ := c.rscGetVpcPair(ctx, tf)
	if vpcPairModel == nil {
		tflog.Error(ctx, "RscDeleteVpcPair: vPC pair not present, might be already deleted")
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
	vpcPairModel, err = c.rscGetVpcPair(ctx, tf)
	if err != nil {
		tflog.Error(ctx, "RscDeleteVpcPair: Failed to get vPC Pair after delete")
		dg.AddError("Failed to get vPC Pair after delete", fmt.Sprintf("Error %v", err))
		return
	}
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

	//Check if the serial number is valid and present in the fabric before creating vPC Pair
	c.rscValidateSerialNumber(ctx, dg, tf)
	if dg.HasError() {
		return
	}

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
		tflog.Error(ctx, "RscCreateVpcPair: Failed to marshal vPC Pair data")
		dg.AddError("Failed to marshal vPC Pair data", fmt.Sprintf("Error %v", err))
		return
	}

	// Check if the vPC Pair is already present
	nfdcVpcPairModel, _ = c.rscGetVpcPair(ctx, tf)
	if nfdcVpcPairModel == nil {
		log.Printf("vPC pair not present, creating new vPC Pair")
		res, err := api.Post(payload)
		if err != nil {
			tflog.Error(ctx, "RscCreateVpcPair: Failed to create vPC Pair")
			dg.AddError("Failed to create vPC Pair", fmt.Sprintf("Error %v: %v", err, res.String()))
			return
		}
		// Check if the vPC Pair is present after creation
		nfdcVpcPairModel, err = c.rscGetVpcPair(ctx, tf)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("RscCreateVpcPair: Failed to get vPC Pair after create: %v", err))
			dg.AddError("Failed to create vPC Pair", fmt.Sprintf("Failed to get vPC Pair after create: Error %v", err))
			return
		}
		if nfdcVpcPairModel == nil {
			tflog.Error(ctx, "RscCreateVpcPair: Failed to get vPC Pair after create")
			dg.AddError("Failed to create vPC Pair", "vPC Pair data empty in NDFC")
			return
		}
	} else {
		err := fmt.Errorf("vPC Pair already exists with serial numbers %s %s", nfdcVpcPairModel.PeerOneId, nfdcVpcPairModel.PeerTwoId)
		tflog.Error(ctx, "RscCreateVpcPair: Existing vPC Pair present, destroy to proceed new one")
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
	rscCreateTfModelAndId(tf, nfdcVpcPairModel)
	log.Printf("[TRACE] vPC pair model: SerialNumbers %v, Id %v", tf.SerialNumbers, tf.Id)
}

func (c *NDFC) RscUpdateVpcPair(ctx context.Context, dg *diag.Diagnostics, tf *resource_vpc_pair.VpcPairModel) {

	//Check if the serial number is valid and present in the fabric before updating vPC Pair
	c.rscValidateSerialNumber(ctx, dg, tf)
	if dg.HasError() {
		return
	}

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
		tflog.Error(ctx, "RscUpdateVpcPair: Failed to marshal vPC Pair data")
		dg.AddError("Failed to marshal vPC Pair data", fmt.Sprintf("Error %v", err))
		return
	}

	// Directly call PUT , if it is update mode
	api := api.NewVpcPairAPI(c.GetLock(ResourceVpcPair), &c.apiClient)
	res, err := api.Put(payload)
	if err != nil {
		tflog.Error(ctx, "RscUpdateVpcPair: Failed to update vPC Pair")
		dg.AddError("Failed to update vPC Pair", fmt.Sprintf("Error %v: %v", err, res.String()))
		return
	}
	// Check if the vPC Pair is present after update
	nfdcVpcPairModel, err = c.rscGetVpcPair(ctx, tf)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("RscUpdateVpcPair: Failed to get vPC Pair after update: %v", err))
		dg.AddError("Failed to update vPC Pair", fmt.Sprintf("Failed to get vPC Pair after update: Error %v", err))
		return
	}
	if nfdcVpcPairModel == nil {
		tflog.Error(ctx, "RscUpdateVpcPair: Failed to get vPC Pair after update")
		dg.AddError("Failed to update vPC Pair", "vPC Pair data empty in NDFC")
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
	rscCreateTfModelAndId(tf, nfdcVpcPairModel)
	log.Printf("[TRACE] vPC pair model: SerialNumbers %v, Id %v", tf.SerialNumbers, tf.Id)
}

func (c *NDFC) rscGetVpcPair(ctx context.Context, tf *resource_vpc_pair.VpcPairModel) (*resource_vpc_pair.NDFCVpcPairModel, error) {

	api := api.NewVpcPairAPI(c.GetLock(ResourceVpcPair), &c.apiClient)
	api.VpcPairID = tf.GetModelData().SerialNumbers[0]

	payload, erro := api.Get()
	modellist := []resource_vpc_pair.NDFCVpcPairModel{}

	if erro == nil && (string(payload) == "[]" || payload == nil) {
		tflog.Debug(ctx, "RscGetVpcPair: No vPC Pair data present in NDFC")
		return nil, nil
	} else if erro != nil {
		tflog.Debug(ctx, "RscGetVpcPair: Failed to get vPC Pair")
		return nil, erro
	}
	log.Printf("[TRACE] vPC pair data: %v %v", string(payload), api.VpcPairID)
	err := json.Unmarshal(payload, &modellist)
	if err != nil {
		tflog.Error(ctx, "RscGetVpcPair: Failed to unmarshal vPC Pair data")
		err = fmt.Errorf("failed to unmarshal vPC Pair data")
		return nil, err
	}
	vpcPairModel := rscValidateVpcPair(ctx, tf, modellist)
	if vpcPairModel == nil {
		err = fmt.Errorf("%s", ErrVpcPairNotFound)
		return nil, err
	}
	return vpcPairModel, nil
}

func rscCreateTfModelAndId(tf *resource_vpc_pair.VpcPairModel,
	vpcPairModel *resource_vpc_pair.NDFCVpcPairModel) {

	// Set the output data to terraform model
	tf.SetModelData(vpcPairModel)
	Id := vpcPairModel.PeerOneId + "/" + vpcPairModel.PeerTwoId
	tf.Id = types.StringValue(Id)
	log.Printf("Setting Id %v", tf.Id)

}

func rscValidateVpcPair(ctx context.Context, tf *resource_vpc_pair.VpcPairModel, vpcPairModelList []resource_vpc_pair.NDFCVpcPairModel) *resource_vpc_pair.NDFCVpcPairModel {
	// Validate the vPC Pair status
	nfdcVpcPairModel := tf.GetModelData()
	peerOneId := nfdcVpcPairModel.SerialNumbers[0]
	peerTwoId := nfdcVpcPairModel.SerialNumbers[1]

	for _, vpcPairModel := range vpcPairModelList {
		log.Printf("From Config: PeerOneId %v, PeerTwoId %v", peerOneId, peerTwoId)
		log.Printf("From NDFC: PeerOneId %v, PeerTwoId %v", vpcPairModel.PeerOneId, vpcPairModel.PeerTwoId)
		if (peerOneId == vpcPairModel.PeerOneId && peerTwoId == vpcPairModel.PeerTwoId) ||
			(peerOneId == vpcPairModel.PeerTwoId && peerTwoId == vpcPairModel.PeerOneId) {
			tflog.Debug(ctx, "RscValidateVpcPair: vPC pair status is present and valid")
			return &vpcPairModel
		}
	}
	return nil
}
func (c *NDFC) rscCheckVpcPairRecommendations(ctx context.Context, tf *resource_vpc_pair.VpcPairModel) error {
	// Get the recommendations for the vPC Pair
	api := api.NewVpcPairAPI(c.GetLock(ResourceVpcPair), &c.apiClient)
	nfdcVpcPairModel := tf.GetModelData()
	recmdList := []resource_vpc_pair.NDFCVpcPairRecommendations{}

	api.VpcPairID = nfdcVpcPairModel.SerialNumbers[0]
	api.GetRecommendations = true
	api.VirtualPeerLink = *nfdcVpcPairModel.UseVirtualPeerlink
	payload, err := api.Get()
	if err != nil {
		tflog.Error(ctx, "RscVpcPairRecommendations: Failed to get recommendations")
		err = fmt.Errorf("failed to get recommendations")
		return err
	}
	tflog.Debug(ctx, fmt.Sprintf("Recommendations payload: %s", payload))
	api.GetRecommendations = false
	err = json.Unmarshal(payload, &recmdList)
	if err != nil {
		tflog.Error(ctx, "RscVpcPairRecommendations: Failed to unmarshal recommendations data")
		return err
	}
	tflog.Debug(ctx, fmt.Sprintf("Recommendations: %v", recmdList))
	for _, rec := range recmdList {
		log.Printf("SerialNumber %v, Recmd %v, RecReason %v", rec.SerialNumber, rec.Recommended, rec.RecommendationReason)
		if rec.SerialNumber == nfdcVpcPairModel.SerialNumbers[1] {
			if rec.Recommended {
				tflog.Debug(ctx, "NDFC Recommendation met")
				return nil
			} else {
				if rec.RecommendationReason == "Switches are not connected" {
					tflog.Debug(ctx, "NDFC Recommendation not met, but it is a transient error")
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
func (c *NDFC) RscImportVpcPairs(ctx context.Context, dg *diag.Diagnostics, tf *resource_vpc_pair.VpcPairModel) {
	// Get the vPC Pair data
	var err diag.Diagnostics
	id := tf.Id.ValueString()
	idSplit := strings.Split(id, ":")
	if len(idSplit) != 2 {
		dg.AddError("Import state failed", "ID should be in the format of serial number:serial number")
		return
	}
	tf.SerialNumbers, err = types.SetValue(types.StringType, []attr.Value{types.StringValue(idSplit[0]), types.StringValue(idSplit[1])})
	tflog.Debug(ctx, fmt.Sprintf("SerialNumbers %v", tf.SerialNumbers))
	if err != nil {
		dg.AddError("Import state failed", "Failed to set serial numbers")
		return
	}
	vpcPairModel, erro := c.rscGetVpcPair(ctx, tf)
	if erro != nil {
		dg.AddError("Import state failed", fmt.Sprintf("Failed to get vPC Pair: %v", erro))
		return
	}
	if vpcPairModel == nil {
		dg.AddError("Unable to get vpcPair", "vPC Pair data empty in NDFC")
		return
	}
	tflog.Debug(ctx, fmt.Sprintf("vPC Pair model: %v", vpcPairModel))
	rscCreateTfModelAndId(tf, vpcPairModel)
}

/*
	    vPC pair does not have a module level deployment,
		so switch deploy is required
*/
func (c *NDFC) RscDeployVpcPair(ctx context.Context, dg *diag.Diagnostics, tf *resource_vpc_pair.VpcPairModel, fabricName string) {

	ndfcVpcPairModel := tf.GetModelData()
	serialNumbers := ndfcVpcPairModel.SerialNumbers
	tflog.Debug(ctx, fmt.Sprintf("RscDeployVpcPair: Deploying vPC Pair with serial numbers: %v", serialNumbers))
	tflog.Debug(ctx, "Performing config save and deploy for vPC Pair")

	c.RecalculateAndDeploy(ctx, dg, fabricName, true, true, serialNumbers)

	if dg.HasError() {
		return
	}
}

func (c *NDFC) rscValidateSerialNumber(ctx context.Context, dg *diag.Diagnostics, tf *resource_vpc_pair.VpcPairModel) {
	vapi := api.NewFabricAPI(c.GetLock(ResourceFabrics), &c.apiClient)
	// Check if the serial number is valid and present a the fabric
	vapi.Serialnumber = tf.GetModelData().SerialNumbers[0]
	payload, err := vapi.Get()
	if err != nil || payload == nil || string(payload) == "[]" {
		tflog.Error(ctx, fmt.Sprintf("RscValidateSerialNumber: Failed to get fabric name for serial number %s", vapi.Serialnumber))
		dg.AddError("Serial number not valid", "Serial number is not part of any fabric")
		return
	}
	vapi.Serialnumber = tf.GetModelData().SerialNumbers[1]
	payload, err = vapi.Get()
	if err != nil || payload == nil || string(payload) == "[]" {
		tflog.Error(ctx, fmt.Sprintf("RscValidateSerialNumber: Failed to get fabric name for serial number %s", vapi.Serialnumber))
		dg.AddError("Serial number not valid", "Serial number is not part of any fabric")
		return
	}

}
