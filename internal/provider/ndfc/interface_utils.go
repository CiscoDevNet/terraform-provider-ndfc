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
	"fmt"
	"log"
	"regexp"
	"strings"
	"terraform-provider-ndfc/internal/provider/resources/resource_interface_common"
	"terraform-provider-ndfc/internal/provider/types"
	"unicode"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func (c NDFC) IfCreateID(ctx context.Context, inData *resource_interface_common.NDFCInterfaceCommonModel) (string, map[string][]string) {
	// Create ID logic
	ID := ""

	ifMap := make(map[string][]string)
	// Get the data into the map
	for i := range inData.Interfaces {
		ifMap[inData.Interfaces[i].SerialNumber] = append(ifMap[inData.Interfaces[i].SerialNumber], inData.Interfaces[i].InterfaceName)
	}
	for k, v := range ifMap {
		ID += k + "["
		for i := range v {
			ID += v[i] + ","
		}
		ID += "],"
	}
	ID = strings.Replace(ID, ",]", "]", -1)
	ID = strings.TrimSuffix(ID, ",")
	log.Printf("Created ID: %s", ID)
	return ID, ifMap
}

// <serial_number>[<interface_name>,<interface_name>],<serial_number>[<interface_name>]
// to map
func ifIdToMap(ID string) map[string][]string {
	// Convert ID to map logic
	ifMap := make(map[string][]string)
	// match - SerialNumber[InterfaceName,InterfaceName],...
	// SerialNumber:
	//	(\w => [a-zA-Z0-9_])
	//  ~ => catch vpc serial numbets (a~b)
	// InterfaceName can have - and /
	pattern := `([\w|~]+)\[([\w|\,|\/|-]+)\]`

	//str := "ABC123[Eth0,Eth1],DEF456[Vl0,Vl1]"

	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(ID, -1)
	for _, match := range matches {
		ifMap[match[1]] = strings.Split(match[2], ",")
	}
	log.Printf("Decoded ID: %v", ifMap)
	return ifMap
}

func (c NDFC) IfTypeSet(inData *resource_interface_common.NDFCInterfaceCommonModel, ifType string) {
	switch ifType {
	case "ethernet":
		fallthrough
	case "INTERFACE_ETHERNET":
		ifType = "INTERFACE_ETHERNET"
	case "loopback":
		fallthrough
	case "INTERFACE_LOOPBACK":
		ifType = "INTERFACE_LOOPBACK"
	case "vlan":
		fallthrough
	case "INTERFACE_VLAN":
		ifType = "INTERFACE_VLAN"
	case "portchannel":
		fallthrough
	case "INTERFACE_PORT_CHANNEL":
		ifType = "INTERFACE_PORT_CHANNEL"
	case "vpc":
		fallthrough
	case "INTERFACE_VPC":
		ifType = "INTERFACE_VPC"
	default:
		log.Panicf("Interface type not supported: %s", ifType)
	}

	for k, intf := range inData.Interfaces {
		intf.InterfaceType = ifType
		inData.Interfaces[k] = intf
	}
}

func (c NDFC) NDFCIfType(ifType string) string {
	switch ifType {
	case "ethernet":
		fallthrough
	case "INTERFACE_ETHERNET":
		ifType = "INTERFACE_ETHERNET"
	case "loopback":
		fallthrough
	case "INTERFACE_LOOPBACK":
		ifType = "INTERFACE_LOOPBACK"
	case "vlan":
		fallthrough
	case "INTERFACE_VLAN":
		ifType = "INTERFACE_VLAN"
	case "portchannel":
		fallthrough
	case "INTERFACE_PORT_CHANNEL":
		ifType = "INTERFACE_PORT_CHANNEL"
	case "vpc":
		fallthrough
	case "INTERFACE_VPC":
		ifType = "INTERFACE_VPC"
	default:
		log.Panicf("Interface type not supported: %s", ifType)
	}
	return ifType
}

// inData - to be sent to NDFC
func (c NDFC) IfPreProcess(inData *resource_interface_common.NDFCInterfaceCommonModel) {
	for k, intf := range inData.Interfaces {
		if intf.SerialNumber == "" {
			intf.SerialNumber = inData.SerialNumber
		}
		intf.NvPairs.InterfaceName = intf.InterfaceName
		inData.Interfaces[k] = intf
	}
}

// inData - data received from NDFC
func (c NDFC) IfPostProcess(inData *resource_interface_common.NDFCInterfaceCommonModel) {
	for k, intf := range inData.Interfaces {
		inData.Interfaces[k] = intf
	}
}

func (c NDFC) ifDiff(ctx context.Context,
	state *resource_interface_common.NDFCInterfaceCommonModel,
	plan *resource_interface_common.NDFCInterfaceCommonModel) map[string]interface{} {

	createIntf := new(resource_interface_common.NDFCInterfaceCommonModel)
	deleteIntf := new(resource_interface_common.NDFCInterfaceCommonModel)
	updateIntf := new(resource_interface_common.NDFCInterfaceCommonModel)
	deleteIntf.Policy = plan.Policy
	createIntf.Policy = plan.Policy
	updateIntf.Policy = plan.Policy
	createIntf.Interfaces = make(map[string]resource_interface_common.NDFCInterfacesValue)
	deleteIntf.Interfaces = make(map[string]resource_interface_common.NDFCInterfacesValue)
	updateIntf.Interfaces = make(map[string]resource_interface_common.NDFCInterfacesValue)

	action := make(map[string]interface{})

	for k, intf := range plan.Interfaces {
		// Check in state
		if stateIntf, ok := state.Interfaces[k]; !ok {
			// New interface, create it
			tflog.Debug(ctx, fmt.Sprintf("New Interface in plan: %s:%s", intf.SerialNumber, intf.InterfaceName))
			createIntf.Interfaces[k] = intf
		} else {
			stateIntf.FilterThisValue = true
			var ctrl bool
			tflog.Debug(ctx, fmt.Sprintf("Existing Interface in plan: %s:%s", intf.SerialNumber, intf.InterfaceName))
			action := intf.CreatePlan(stateIntf, &ctrl)
			tflog.Debug(ctx, fmt.Sprintf("CreatePlan: Action: %v", action))
			if action == types.RequiresUpdate {
				tflog.Debug(ctx, fmt.Sprintf("Interface requires update: %s:%s", intf.SerialNumber, intf.InterfaceName))
				updateIntf.Interfaces[k] = intf
			} else if action == types.RequiresReplace {
				tflog.Debug(ctx, fmt.Sprintf("Interface requires replace: %s:%s", intf.SerialNumber, intf.InterfaceName))
				deleteIntf.Interfaces[k] = stateIntf
				createIntf.Interfaces[k] = intf
			} else {
				tflog.Debug(ctx, fmt.Sprintf("Interface does not require update: %s:%s", intf.SerialNumber, intf.InterfaceName))
			}
			state.Interfaces[k] = stateIntf
		}
		plan.Interfaces[k] = intf

	}
	if plan.Deploy {
		if len(updateIntf.Interfaces) > 0 {
			// Deploy flag is set - hence deploy all modified interfaces
			updateIntf.Deploy = true
		}
		if len(createIntf.Interfaces) > 0 {
			// Deploy flag is set - hence deploy all new interfaces
			createIntf.Deploy = true
		}
	}

	// check if deploy flag is changed in plan
	if !state.Deploy && plan.Deploy {
		// all interfaces should be deployed
		tflog.Debug(ctx, "Deployment flag has changed")
		action["deploy"] = true
	} else {
		// no flag changed
		action["deploy"] = false
	}
	action["create"] = createIntf
	action["update"] = updateIntf

	// check for deleted entries
	for k, intf := range state.Interfaces {
		if intf.FilterThisValue {
			continue
		}
		if _, ok := plan.Interfaces[k]; !ok {
			// Interface deleted
			tflog.Debug(ctx, fmt.Sprintf("Interface deleted: %s:%s", intf.SerialNumber, intf.InterfaceName))
			intf.NvPairs.AdminState = "false"
			deleteIntf.Interfaces[k] = intf
		}
	}
	action["del"] = deleteIntf
	return action
}

/*
 * Just convert first letter to upper case
 * strings.Title() is deprecated and equivalent text/case api looks complex for
 * this simple use case
 */
func ToTitleCase(s string) string {
	if len(s) == 0 {
		return s
	}
	// next two lines - typecasting to convert first letter
	p := ([]rune(s))
	p[0] = rune(strings.ToUpper(string(p[0]))[0])
	return string(p)
}

func IsFirstLetterLC(s string) bool {
	if len(s) == 0 {
		return false
	}
	return unicode.IsLower(rune(s[0]))
}

func (c NDFC) IntfRscModified(ctx context.Context, dg *diag.Diagnostics, plan *resource_interface_common.NDFCInterfaceCommonModel,
	state *resource_interface_common.NDFCInterfaceCommonModel) map[string]bool {

	ret := make(map[string]bool)
	for k, intf := range plan.Interfaces {
		// Check in state
		if stateIntf, ok := state.Interfaces[k]; !ok {
			// New interface, create it
			tflog.Debug(ctx, fmt.Sprintf("IntfRscModified: New Interface in plan: %s:%s", intf.SerialNumber, intf.InterfaceName))
			ret[k] = true
		} else {
			stateIntf.FilterThisValue = true
			tflog.Debug(ctx, fmt.Sprintf("Existing Interface in plan: %s:%s", intf.SerialNumber, intf.InterfaceName))
			action := intf.DeepEqual(stateIntf)
			tflog.Debug(ctx, fmt.Sprintf("CreatePlan: Action: %v", action))
			if action != types.ValuesDeeplyEqual {
				tflog.Debug(ctx, fmt.Sprintf("Interface requires update: %s:%s", intf.SerialNumber, intf.InterfaceName))
				ret[k] = true
			} else {
				tflog.Debug(ctx, fmt.Sprintf("Interface does not require update: %s:%s", intf.SerialNumber, intf.InterfaceName))
				ret[k] = false
			}
			state.Interfaces[k] = stateIntf
		}

	}
	// check for deleted entries
	for k, intf := range state.Interfaces {
		if intf.FilterThisValue {
			continue
		}
		if _, ok := plan.Interfaces[k]; !ok {
			// Interface deleted
			tflog.Debug(ctx, fmt.Sprintf("Interface deleted: %s:%s", intf.SerialNumber, intf.InterfaceName))
			ret[k] = true
		}
	}
	return ret
}
