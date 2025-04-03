// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package ndfc

import (
	"fmt"
	"log"
	"strings"
	"terraform-provider-ndfc/internal/provider/resources/resource_networks"
	"terraform-provider-ndfc/internal/provider/resources/resource_vrf_bulk"
)

func (c NDFC) RscCreateID(rsc interface{}, rscType string) string {

	uniqueID := "%s/%v"
	idList := make([]string, 0)
	idMap := make(map[string][]string)
	fname := ""
	switch rscType {
	case ResourceNetworks:
		ndNw := rsc.(*resource_networks.NDFCNetworksModel)
		fname = ndNw.FabricName
		for k := range ndNw.Networks {
			idList = append(idList, k)
			idMap[k] = make([]string, 0)
		}
	case ResourceVrfBulk:
		ndVrfs := rsc.(*resource_vrf_bulk.NDFCVrfBulkModel)
		fname = ndVrfs.FabricName
		for k, v := range ndVrfs.Vrfs {
			idList = append(idList, k)
			for kk := range v.AttachList {
				idMap[k] = append(idMap[k], kk)
			}
		}
	default:
		return ""
	}

	for i, rsName := range idList {
		if len(idMap[rsName]) > 0 {
			idList[i] = fmt.Sprintf("%s{%s}", rsName, strings.Join(idMap[rsName], ","))
		}
	}
	id_list_str := "[" + strings.Join(idList, ",") + "]"
	return fmt.Sprintf(uniqueID, fname, id_list_str)
}

func (c NDFC) CreateFilterMap(ID string, filterMap *map[string]bool) (string, []string) {

	result := c.RscBulkSplitID(ID)
	for _, v := range result["rsc"] {
		log.Printf("Set filtering of %s to true", v)
		(*filterMap)[v] = true
	}
	return result["fabric"][0], result["rsc"]
}

func (c NDFC) RscBulkSplitID(ID string) map[string][]string {
	// Split the ID into its components
	result := map[string][]string{}
	rsNames := make([]string, 0)
	components := strings.Split(ID, "/[")
	rscData := components[1]
	start := 0
	for i := 0; i < len(rscData); i++ {
		if rscData[i] == '{' {
			rsName := rscData[start:i]
			rsNames = append(rsNames, rsName)
			for j := i + 1; j < len(rscData); j++ {
				if rscData[j] == '}' {
					attachs := rscData[i+1 : j]
					fmt.Println(attachs)
					attachments := strings.Split(attachs, ",")
					result[rsName] = attachments
					i = j + 1
					start = j + 2
					break
				}
			}
		} else if rscData[i] == ',' || rscData[i] == ']' {
			rsName := rscData[start:i]
			rsNames = append(rsNames, rsName)
			result[rsName] = []string{}
			start = i + 1
		}
	}
	result["fabric"] = []string{components[0]}
	result["rsc"] = rsNames
	return result
}
