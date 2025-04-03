// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package types

import (
	"strconv"
	"strings"
)

const (
	ActionNone = iota
	ValuesDeeplyEqual
	RequiresReplace
	RequiresUpdate
	ControlFlagUpdate
	PortListUpdate
)

type CSVString []string

func (i *CSVString) UnmarshalJSON(data []byte) error {
	if string(data) == "" || string(data) == "\"\"" {
		*i = make(CSVString, 0)
		return nil
	}
	ss := string(data)
	// If the string is quoted, remove the quotes
	ssUn, err := strconv.Unquote(ss)
	if err == nil {
		// Quote removed
		ss = ssUn
	}
	*i = strings.Split(ss, ",")

	return nil
}

func (i CSVString) MarshalJSON() ([]byte, error) {
	res := ""
	res = strings.Join(i, ",")
	return []byte(strconv.Quote(res)), nil

}

type Int64Custom int64

func (i *Int64Custom) IsEmpty() bool {
	if i == nil {
		return true
	}
	return *i == -9223372036854775808
}

func (i *Int64Custom) UnmarshalJSON(data []byte) error {
	if string(data) == "" || string(data) == "\"\"" {
		*i = -9223372036854775808
	} else {
		ss := string(data)
		// If the string is quoted, remove the quotes
		ssUn, err := strconv.Unquote(ss)
		if err == nil {
			// Quote removed
			ss = ssUn
		}
		ii, _ := strconv.ParseInt(ss, 10, 64)
		*i = Int64Custom(ii)
	}
	return nil
}

func (i Int64Custom) MarshalJSON() ([]byte, error) {
	res := ""
	res = strconv.FormatInt(int64(i), 10)
	return []byte(strconv.Quote(res)), nil

}
