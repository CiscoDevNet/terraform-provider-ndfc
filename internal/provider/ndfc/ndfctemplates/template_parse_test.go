// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package ndfctemplates

import (
	"testing"
)

func TestParseTemplateFromActualTemplate(t *testing.T) {
	// Use actual template content from ext_multisite_underlay_setup.tmpl
	// This is a subset to focus on the key features
	templateContent := `##template variables
#    Copyright (c) 2019-2020 by Cisco Systems, Inc.
#    All rights reserved.
@(IsMandatory=true, IsInternal=true, IsSourceFabric=true)
string SOURCE_FABRIC_NAME;
@(IsMandatory=true, IsInternal=true, IsDestinationFabric=true)
string DEST_FABRIC_NAME;
@(IsMandatory=false, IsInternal=true, IsSourceDevice=true)
string SOURCE_SERIAL_NUMBER;
@(IsMandatory=true, IsInternal=true, IsDestinationDevice=true)
string DEST_SERIAL_NUMBER;
@(IsMandatory=true, IsInternal=true, IsSourceSwitchName=true)
string SOURCE_SWITCH_NAME;
@(IsMandatory=true, IsInternal=true, IsDestinationSwitchName=true)
string DEST_SWITCH_NAME;
@(IsMandatory=true, IsInternal=true, IsSourceInterface=true)
interface SOURCE_IF_NAME;
@(IsMandatory=true, IsInternal=true, IsDestinationInterface=true)
interface DEST_IF_NAME;
@(IsMandatory=true, IsInternal=false)
string LINK_UUID
{
    defaultValue = LINK_UUID_1;
};
@(IsMandatory=true, IsAsn=true, Description="BGP Autonomous System Number in Source Fabric", DisplayName="Source BGP ASN")
string asn{
minLength=1;
maxLength=11;
regularExpr=^(((\+)?[1-9]{1}[0-9]{0,8}|(\+)?[1-3]{1}[0-9]{1,9}|(\+)?[4]{1}([0-1]{1}[0-9]{8}|[2]{1}([0-8]{1}[0-9]{7}|[9]{1}([0-3]{1}[0-9]{6}|[4]{1}([0-8]{1}[0-9]{5}|[9]{1}([0-5]{1}[0-9]{4}|[6]{1}([0-6]{1}[0-9]{3}|[7]{1}([0-1]{1}[0-9]{2}|[2]{1}([0-8]{1}[0-9]{1}|[9]{1}[0-5]{1})))))))))|([1-5]\d{4}|[1-9]\d{0,3}|6[0-4]\d{3}|65[0-4]\d{2}|655[0-2]\d|6553[0-5])(\.([1-5]\d{4}|[1-9]\d{0,3}|6[0-4]\d{3}|65[0-4]\d{2}|655[0-2]\d|6553[0-5]|0))?)$;
};
@(IsMandatory=true, Description="IP address with mask (e.g. 192.168.10.1/24) of Source Interface", DisplayName="Source <br/>IP Address/Mask")
ipV4AddressWithSubnet IP_MASK;
@(IsMandatory=true, Description="IP address of Destination Interface", DisplayName="Destination IP")
ipV4Address NEIGHBOR_IP;
##template content
// Implementation code
SOURCE = LINK_UUID
`

	template := NewNDFCTemplate()
	err := template.ParseTemplate(templateContent)

	if err != nil {
		t.Fatalf("Failed to parse template: %v", err)
	}

	// Test 1: Check the total number of fields extracted
	expectedFieldCount := 12 // Count all fields in the template
	if len(template.Fields) != expectedFieldCount {
		t.Errorf("Expected %d fields, got %d", expectedFieldCount, len(template.Fields))
	}

	// Test 2: Check field with complex description containing parentheses
	ipMaskField, exists := template.Fields["IP_MASK"]
	if !exists {
		t.Error("IP_MASK field not found")
	} else {
		expectedDesc := "IP address with mask (e.g. 192.168.10.1/24) of Source Interface"
		actualDesc, exists := ipMaskField.Flags["Description"]
		if !exists {
			t.Error("Description not found for IP_MASK field")
		} else if actualDesc != expectedDesc {
			t.Errorf("Description not parsed correctly for IP_MASK.\nGot: %s\nWant: %s", actualDesc, expectedDesc)
		}

		// Check type
		if ipMaskField.Type != "ipV4AddressWithSubnet" {
			t.Errorf("Field type not parsed correctly. Got %s, want ipV4AddressWithSubnet", ipMaskField.Type)
		}
	}

	// Test 3: Check field with body (defaultValue)
	linkUuidField, exists := template.Fields["LINK_UUID"]
	if !exists {
		t.Error("LINK_UUID field not found")
	} else {
		if linkUuidField.DefaultValue != "LINK_UUID_1" {
			t.Errorf("Default value not parsed correctly. Got '%s', want 'LINK_UUID_1'", linkUuidField.DefaultValue)
		}
	}

	// Test 4: Check field with complex regex
	asnField, exists := template.Fields["asn"]
	if !exists {
		t.Error("asn field not found")
	} else {
		if asnField.MinLength != 1 {
			t.Errorf("MinLength not parsed correctly for asn. Got %d, want 1", asnField.MinLength)
		}

		if asnField.MaxLength != 11 {
			t.Errorf("MaxLength not parsed correctly for asn. Got %d, want 11", asnField.MaxLength)
		}

		if asnField.Regex != `^(((\+)?[1-9]{1}[0-9]{0,8}|(\+)?[1-3]{1}[0-9]{1,9}|(\+)?[4]{1}([0-1]{1}[0-9]{8}|[2]{1}([0-8]{1}[0-9]{7}|[9]{1}([0-3]{1}[0-9]{6}|[4]{1}([0-8]{1}[0-9]{5}|[9]{1}([0-5]{1}[0-9]{4}|[6]{1}([0-6]{1}[0-9]{3}|[7]{1}([0-1]{1}[0-9]{2}|[2]{1}([0-8]{1}[0-9]{1}|[9]{1}[0-5]{1})))))))))|([1-5]\d{4}|[1-9]\d{0,3}|6[0-4]\d{3}|65[0-4]\d{2}|655[0-2]\d|6553[0-5])(\.([1-5]\d{4}|[1-9]\d{0,3}|6[0-4]\d{3}|65[0-4]\d{2}|655[0-2]\d|6553[0-5]|0))?)$` {
			t.Error("Regex not parsed for asn field")
		}
	}

	// Test 5: Check interface type field
	sourceIfField, exists := template.Fields["SOURCE_IF_NAME"]
	if !exists {
		t.Error("SOURCE_IF_NAME field not found")
	} else {
		if sourceIfField.Type != "interface" {
			t.Errorf("Field type not parsed correctly. Got %s, want interface", sourceIfField.Type)
		}
	}

	// Test 6: Ensure SOURCE is not extracted as it's only in the implementation part
	if _, exists := template.Fields["SOURCE"]; exists {
		t.Error("SOURCE field incorrectly extracted from implementation section")
	}
}

func TestValidatePayloadUsingActualTemplate(t *testing.T) {
	// Use actual template content from ext_multisite_underlay_setup.tmpl
	// Simplified version focusing on testing validation
	templateContent := `##template variables
@(IsMandatory=true, Description="BGP Autonomous System Number in Source Fabric", DisplayName="Source BGP ASN")
string asn{
minLength=1;
maxLength=11;
regularExpr=^(((\+)?[1-9]{1}[0-9]{0,8}|(\+)?[1-3]{1}[0-9]{1,9}|(\+)?[4]{1}([0-1]{1}[0-9]{8}|[2]{1}([0-8]{1}[0-9]{7}|[9]{1}([0-3]{1}[0-9]{6}|[4]{1}([0-8]{1}[0-9]{5}|[9]{1}([0-5]{1}[0-9]{4}|[6]{1}([0-6]{1}[0-9]{3}|[7]{1}([0-1]{1}[0-9]{2}|[2]{1}([0-8]{1}[0-9]{1}|[9]{1}[0-5]{1})))))))))|([1-5]\d{4}|[1-9]\d{0,3}|6[0-4]\d{3}|65[0-4]\d{2}|655[0-2]\d|6553[0-5])(\.([1-5]\d{4}|[1-9]\d{0,3}|6[0-4]\d{3}|65[0-4]\d{2}|655[0-2]\d|6553[0-5]|0))?)$;
};
@(IsMandatory=true, Description="IP address with mask (e.g. 192.168.10.1/24) of Source Interface", DisplayName="Source <br/>IP Address/Mask")
ipV4AddressWithSubnet IP_MASK;
@(IsMandatory=false, Description="Enable BGP password authentication", DisplayName="Enable BGP Password Authentication")
boolean BGP_PASSWORD_ENABLE;
@(IsMandatory=true, IsShow="BGP_PASSWORD_ENABLE==true", Description="BGP Password", DisplayName="BGP Password")
string BGP_PASSWORD;
##template content
// Implementation code
SOURCE = LINK_UUID
`

	template := NewNDFCTemplate()
	err := template.ParseTemplate(templateContent)

	if err != nil {
		t.Fatalf("Failed to parse template: %v", err)
	}

	// Test 1: Valid payload with all required fields
	validPayloadString := `{
		"asn": "65001",
		"IP_MASK": "192.168.10.1/24",
		"BGP_PASSWORD_ENABLE": "false"
	}`
	valid, errs := template.ValidatePayload([]byte(validPayloadString))
	if !valid {
		t.Errorf("ValidatePayload() should be valid: %v", errs)
	}

	// Test 2: Valid payload with conditional field (BGP_PASSWORD when enabled)
	validConditionalString := `{
		"asn": "65001",
		"IP_MASK": "192.168.10.1/24",
		"BGP_PASSWORD_ENABLE": "true",
		"BGP_PASSWORD": "mysecretpassword"
	}`
	valid, errs = template.ValidatePayload([]byte(validConditionalString))
	if !valid {
		t.Errorf("ValidatePayload() should be valid with conditional field: %v", errs)
	}

	// Test 3: Invalid payload - password required but not provided
	invalidConditionalString := `{
		"asn": "65001",
		"IP_MASK": "192.168.10.1/24",
		"BGP_PASSWORD_ENABLE": "true"
	}`
	valid, _ = template.ValidatePayload([]byte(invalidConditionalString))
	if valid {
		t.Error("ValidatePayload() should be invalid when required conditional field missing")
	}

	// Test 4: Invalid payload - missing mandatory field
	invalidMissingFieldString := `{
		"BGP_PASSWORD_ENABLE": "false"
	}`
	valid, _ = template.ValidatePayload([]byte(invalidMissingFieldString))
	if valid {
		t.Error("ValidatePayload() should be invalid when mandatory field missing")
	}

	// Test 5: Custom field isn't required (SOURCE)
	validCustomFieldString := `{
		"asn": "65001",
		"IP_MASK": "192.168.10.1/24",
		"BGP_PASSWORD_ENABLE": "false",
		"SOURCE": "something"
	}`
	valid, errs = template.ValidatePayload([]byte(validCustomFieldString))
	if !valid {
		t.Errorf("ValidatePayload() should be valid with extra SOURCE field: %v", errs)
	}
}

func TestParseComplexAttributesFromActualTemplate(t *testing.T) {
	// Test direct parsing of complex attributes from the actual template

	// Test 1: Field with parentheses in description
	attrStr := `@(IsMandatory=true, Description="IP address with mask (e.g. 192.168.10.1/24) of Source Interface", DisplayName="Source <br/>IP Address/Mask")`
	attrs := parseAttributes(attrStr)

	expectedDescription := "IP address with mask (e.g. 192.168.10.1/24) of Source Interface"
	if attrs["Description"] != expectedDescription {
		t.Errorf("Failed to parse attribute with parentheses.\nGot: %s\nWant: %s", attrs["Description"], expectedDescription)
	}

	expectedDisplayName := "Source <br/>IP Address/Mask"
	if attrs["DisplayName"] != expectedDisplayName {
		t.Errorf("Failed to parse attribute with HTML tags.\nGot: %s\nWant: %s", attrs["DisplayName"], expectedDisplayName)
	}

	// Test 2: IsShow condition
	attrStr = `@(IsMandatory=true, IsShow="BGP_PASSWORD_ENABLE==true", Description="BGP Password", DisplayName="BGP Password")`
	attrs = parseAttributes(attrStr)

	expectedIsShow := "BGP_PASSWORD_ENABLE==true"
	if attrs["IsShow"] != expectedIsShow {
		t.Errorf("Failed to parse IsShow condition.\nGot: %s\nWant: %s", attrs["IsShow"], expectedIsShow)
	}
}

func TestValidatePayloadUsingComplexTemplate(t *testing.T) {
	templateStr := `##template variables
@(IsMandatory=true, IsShow="BGP_PASSWORD_ENABLE==true && BGP_PASSWORD_INHERIT_FROM_MSD==false", DisplayName="eBGP Password", Description="Encrypted eBGP Password Hex String", Section="Advanced")
string BGP_PASSWORD {
minLength = 1;
maxLength = 130;
regularExpr=^[a-fA-F0-9]+$;
};`

	template := NewNDFCTemplate()
	err := template.ParseTemplate(templateStr)

	if err != nil {
		t.Fatalf("Failed to parse template: %v", err)
	}

	// Test 1: Valid payload with all required fields
	validPayloadString := `{"BGP_PASSWORD": "1234567890abcdef", "BGP_PASSWORD_ENABLE": "true"}`
	valid, errs := template.ValidatePayload([]byte(validPayloadString))
	if !valid {
		t.Errorf("ValidatePayload() should be valid: %v", errs)
	}

	// Test 2: Invalid payload - missing required field
	invalidPayloadString := `{"BGP_PASSWORD_INHERIT_FROM_MSD": "true", "BGP_PASSWORD_ENABLE": "true", "BGP_PASSWORD": ""}`
	valid, _ = template.ValidatePayload([]byte(invalidPayloadString))
	if !valid {
		t.Error("ValidatePayload() should be valid when BGP_PASSWORD_INHERIT_FROM_MSD is true")
	}
}
