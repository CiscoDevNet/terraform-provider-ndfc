// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_template

import (
	"bytes"
	"encoding/json"
	"log"
	"strings"
	"text/template"
)

var template_payload = `##template properties \nname={{.TemplateName}};\ndescription={{.Description}};\ntags={{TrimmedArray .Tags}};\nsupportedPlatforms={{.SupportedPlatforms}};\ntemplateType={{.TemplateType}};\ntemplateSubType={{.TemplateSubType}};\ncontentType={{.ContentType}};\r\n{{.TemplateContent}}`

type NDFCTEmplatePayload struct {
	TemplateName string `json:"templatename"`
	Content      string `json:"content"`
}

type CustomNDFCTemplateModel struct {
	InstanceId         *int64 `json:"instanceClassId,omitempty"`
	TemplateName       string `json:"name,omitempty"`
	Description        string `json:"description,omitempty"`
	Tags               string `json:"tags,omitempty"`
	SupportedPlatforms string `json:"supportedPlatforms,omitempty"`
	FileName           string `json:"fileName,omitempty"`
	TemplateType       string `json:"templateType,omitempty"`
	TemplateContent    string `json:"newContent,omitempty"`
	ContentType        string `json:"contentType,omitempty"`
	TemplateSubType    string `json:"templateSubType,omitempty"`
}

var functions = template.FuncMap{
	"TrimmedArray": TrimmedArray,
}

func TrimmedArray(tags []string) string {
	return strings.Join(tags, ",")
}

func (t *NDFCTemplateModel) MarshalJSON() ([]byte, error) {
	output := new(bytes.Buffer)
	tpl, err := template.New("Template_content").Funcs(functions).Parse(template_payload)
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		return nil, err
	}
	err = tpl.Execute(output, t)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		return nil, err
	}

	log.Println("Output: ", output.String())
	log.Println("Tags: ", t.Tags)
	outPayload := NDFCTEmplatePayload{
		TemplateName: t.TemplateName,
		Content:      output.String(),
	}
	return json.Marshal(outPayload)
}

func (t *NDFCTemplateModel) ReformatContent() {
	switch t.ContentType {
	case "TEMPLATE_CLI":
		fallthrough
	case "PYTHON":
		// Include only the script, not properties
		// NDFC inconsistency: response sometimes includes properties
		// in both content and newcontent fields and hence needs to be removed
		startPos := strings.Index(t.TemplateContent, "##template variables")
		if startPos == -1 {
			return
		}
		t.TemplateContent = t.TemplateContent[startPos:]
		// NDFC inconsistency - extra spaces are added to beginning of some fields in NDFC output
		// Hence trim it off

	}
	t.Description = strings.Trim(t.Description, "\n \t")
}

/* Custom unmarshall is needed due to following reasons
* NDFC inconsistency 1: templatename
  * templatename is the property name for template name in POST
  * it is `name` in GET response for the same field
  * using unmarshall tag in generator could solve this by having two different fieds in generated struct
  * As a unmarshall override is needed for other reasons (See below) fixing it here
* NDFC inconsistency 2: tags field
  * tags field is a string in GET response with array brackets included
  * In POST it is sent as csv strings which seems to be the correct way
  * tags field is Set in Terraform and hence needs conversion
  * json unmarshall fails to decode a stringified array to a slice
*/
func (t *NDFCTemplateModel) UnmarshalJSON(data []byte) error {
	newT := CustomNDFCTemplateModel{}
	err := json.Unmarshal(data, &newT)
	if err != nil {
		log.Println("unmarshall err", err)
		return err
	}

	t.InstanceId = newT.InstanceId
	t.TemplateName = newT.TemplateName
	t.Description = newT.Description
	t.SupportedPlatforms = newT.SupportedPlatforms
	t.FileName = newT.FileName
	t.TemplateType = newT.TemplateType
	t.TemplateContent = newT.TemplateContent
	t.ContentType = newT.ContentType
	t.TemplateSubType = newT.TemplateSubType
	t.Tags = strings.Split(strings.Trim(newT.Tags, "[]"), ", ")
	log.Printf("Tags: %v", t.Tags)
	t.ReformatContent()
	return nil
}
