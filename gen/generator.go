// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

//go:build ignore

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"
)

const (
	definitionsPath   = "./gen/definitions/"
	providerTemplate  = "./gen/templates/provider.go"
	providerLocation  = "./internal/provider/provider.go"
	changelogTemplate = "./gen/templates/changelog.md.tmpl"
	changelogLocation = "./templates/guides/changelog.md.tmpl"
	changelogOriginal = "./CHANGELOG.md"
)

type t struct {
	path   string
	prefix string
	suffix string
}

var templates = []t{
	{
		path:   "./gen/templates/model.go",
		prefix: "./internal/provider/model_ndfc_",
		suffix: ".go",
	},
	{
		path:   "./gen/templates/data_source.go",
		prefix: "./internal/provider/data_source_ndfc_",
		suffix: ".go",
	},
	{
		path:   "./gen/templates/data_source_test.go",
		prefix: "./internal/provider/data_source_ndfc_",
		suffix: "_test.go",
	},
	{
		path:   "./gen/templates/resource.go",
		prefix: "./internal/provider/resource_ndfc_",
		suffix: ".go",
	},
	{
		path:   "./gen/templates/resource_test.go",
		prefix: "./internal/provider/resource_ndfc_",
		suffix: "_test.go",
	},
	{
		path:   "./gen/templates/data-source.tf",
		prefix: "./examples/data-sources/ndfc_",
		suffix: "/data-source.tf",
	},
	{
		path:   "./gen/templates/resource.tf",
		prefix: "./examples/resources/ndfc_",
		suffix: "/resource.tf",
	},
	{
		path:   "./gen/templates/import.sh",
		prefix: "./examples/resources/ndfc_",
		suffix: "/import.sh",
	},
}

type YamlConfig struct {
	Name              string                `yaml:"name"`
	Model             string                `yaml:"model"`
	RestEndpoint      string                `yaml:"rest_endpoint"`
	MinimumVersion    string                `yaml:"minimum_version"`
	DsDescription     string                `yaml:"ds_description"`
	ResDescription    string                `yaml:"res_description"`
	DocCategory       string                `yaml:"doc_category"`
	ExcludeTest       bool                  `yaml:"exclude_test"`
	Attributes        []YamlConfigAttribute `yaml:"attributes"`
	TestPrerequisites string                `yaml:"test_prerequisites"`
}

type YamlConfigAttribute struct {
	ModelName       string                `yaml:"model_name"`
	TfName          string                `yaml:"tf_name"`
	Type            string                `yaml:"type"`
	ModelTypeString bool                  `yaml:"model_type_string"`
	DataPath        []string              `yaml:"data_path"`
	Id              bool                  `yaml:"id"`
	Reference       bool                  `yaml:"reference"`
	Mandatory       bool                  `yaml:"mandatory"`
	Computed        bool                  `yaml:"computed"`
	WriteOnly       bool                  `yaml:"write_only"`
	TfOnly          bool                  `yaml:"tf_only"`
	ExcludeTest     bool                  `yaml:"exclude_test"`
	ExcludeExample  bool                  `yaml:"exclude_example"`
	Description     string                `yaml:"description"`
	Example         string                `yaml:"example"`
	EnumValues      []string              `yaml:"enum_values"`
	MinList         int64                 `yaml:"min_list"`
	MaxList         int64                 `yaml:"max_list"`
	MinInt          int64                 `yaml:"min_int"`
	MaxInt          int64                 `yaml:"max_int"`
	MinFloat        float64               `yaml:"min_float"`
	MaxFloat        float64               `yaml:"max_float"`
	StringPatterns  []string              `yaml:"string_patterns"`
	StringMinLength int64                 `yaml:"string_min_length"`
	StringMaxLength int64                 `yaml:"string_max_length"`
	DefaultValue    string                `yaml:"default_value"`
	Value           string                `yaml:"value"`
	TestValue       string                `yaml:"test_value"`
	Attributes      []YamlConfigAttribute `yaml:"attributes"`
}

type YamlConfigConditionalAttribute struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

// Templating helper function to convert TF name to GO name
func ToGoName(s string) string {
	var g []string

	p := strings.Split(s, "_")

	for _, value := range p {
		g = append(g, strings.Title(value))
	}
	s = strings.Join(g, "")
	return s
}

// Templating helper function to convert string to camel case
func CamelCase(s string) string {
	var g []string

	s = strings.ReplaceAll(s, "-", " ")
	p := strings.Fields(s)

	for _, value := range p {
		g = append(g, strings.Title(value))
	}
	return strings.Join(g, "")
}

// Templating helper function to convert string to snake case
func SnakeCase(s string) string {
	var g []string

	s = strings.ReplaceAll(s, "-", " ")
	p := strings.Fields(s)

	for _, value := range p {
		g = append(g, strings.ToLower(value))
	}
	return strings.Join(g, "_")
}

// Templating helper function to return true if reference included in attributes
func HasReference(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if attr.Reference {
			return true
		}
	}
	return false
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

// Templating helper function to return a list of numbers
func Iterate(count int) []int {
	var i int
	var Items []int
	for i = 0; i < count; i++ {
		Items = append(Items, i)
	}
	return Items
}

// Templating helper function to increment a number
func Increment(i int) int {
	return i + 1
}

// Map of templating functions
var functions = template.FuncMap{
	"toGoName":     ToGoName,
	"camelCase":    CamelCase,
	"snakeCase":    SnakeCase,
	"hasReference": HasReference,
	"iterate":      Iterate,
	"increment":    Increment,
}

func augmentAttribute(attr *YamlConfigAttribute) {
	if attr.TfName == "" {
		attr.TfName = SnakeCase(attr.ModelName)
	}
	if attr.Type == "List" || attr.Type == "Set" {
		for a := range attr.Attributes {
			augmentAttribute(&attr.Attributes[a])
		}
	}
}

func augmentConfig(config *YamlConfig) {
	for ia := range config.Attributes {
		augmentAttribute(&config.Attributes[ia])
	}
	if config.DsDescription == "" {
		config.DsDescription = fmt.Sprintf("This data source can read a %s.", config.Name)
	}
	if config.ResDescription == "" {
		config.ResDescription = fmt.Sprintf("This resource can manage a %s.", config.Name)
	}
}

func getTemplateSection(content, name string) string {
	scanner := bufio.NewScanner(strings.NewReader(content))
	result := ""
	foundSection := false
	beginRegex := regexp.MustCompile(`\/\/template:begin\s` + name + `$`)
	endRegex := regexp.MustCompile(`\/\/template:end\s` + name + `$`)
	for scanner.Scan() {
		line := scanner.Text()
		if !foundSection {
			match := beginRegex.MatchString(line)
			if match {
				foundSection = true
				result += line + "\n"
			}
		} else {
			result += line + "\n"
			match := endRegex.MatchString(line)
			if match {
				foundSection = false
			}
		}
	}
	return result
}

func renderTemplate(templatePath, outputPath string, config interface{}) {
	file, err := os.Open(templatePath)
	if err != nil {
		log.Fatalf("Error opening template: %v", err)
	}
	defer file.Close()

	// skip first line with 'build-ignore' directive for go files
	scanner := bufio.NewScanner(file)
	if strings.HasSuffix(templatePath, ".go") {
		scanner.Scan()
	}
	var temp string
	for scanner.Scan() {
		temp = temp + scanner.Text() + "\n"
	}

	template, err := template.New(path.Base(templatePath)).Funcs(functions).Parse(temp)
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	output := new(bytes.Buffer)
	err = template.Execute(output, config)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	outputFile := filepath.Join(outputPath)
	existingFile, err := os.Open(outputPath)
	if err != nil {
		os.MkdirAll(filepath.Dir(outputFile), 0755)
	} else if strings.HasSuffix(templatePath, ".go") {
		existingScanner := bufio.NewScanner(existingFile)
		var newContent string
		currentSectionName := ""
		beginRegex := regexp.MustCompile(`\/\/template:begin\s(.*?)$`)
		endRegex := regexp.MustCompile(`\/\/template:end\s(.*?)$`)
		for existingScanner.Scan() {
			line := existingScanner.Text()
			if currentSectionName == "" {
				matches := beginRegex.FindStringSubmatch(line)
				if len(matches) > 1 && matches[1] != "" {
					currentSectionName = matches[1]
				} else {
					newContent += line + "\n"
				}
			} else {
				matches := endRegex.FindStringSubmatch(line)
				if len(matches) > 1 && matches[1] == currentSectionName {
					currentSectionName = ""
					newSection := getTemplateSection(string(output.Bytes()), matches[1])
					newContent += newSection
				}
			}
		}
		output = bytes.NewBufferString(newContent)
	}
	// write to output file
	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	f.Write(output.Bytes())
}

func main() {
	providerConfig := make([]string, 0)

	files, _ := os.ReadDir(definitionsPath)
	configs := make([]YamlConfig, len(files))

	// Load configs
	for i, filename := range files {
		yamlFile, err := os.ReadFile(filepath.Join(definitionsPath, filename.Name()))
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}

		config := YamlConfig{}
		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			log.Fatalf("Error parsing yaml: %v", err)
		}
		configs[i] = config
	}

	for i := range configs {
		// Augment config
		augmentConfig(&configs[i])

		// Iterate over templates and render files
		for _, t := range templates {
			renderTemplate(t.path, t.prefix+SnakeCase(configs[i].Name)+t.suffix, configs[i])
		}
		providerConfig = append(providerConfig, configs[i].Name)
	}

	// render provider.go
	renderTemplate(providerTemplate, providerLocation, providerConfig)

	changelog, err := os.ReadFile(changelogOriginal)
	if err != nil {
		log.Fatalf("Error reading changelog: %v", err)
	}
	renderTemplate(changelogTemplate, changelogLocation, string(changelog))
}
