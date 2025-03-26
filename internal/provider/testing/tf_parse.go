package testing

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"regexp"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

type TerraformConfig struct {
	Type    string
	RscType string
	RscName string
	File    *hclwrite.File
}

func (tc *TerraformConfig) AddContent(fc []byte) {

	var diags hcl.Diagnostics

	tc.File, diags = hclwrite.ParseConfig(fc, "unknown", hcl.Pos{Line: 1, Column: 1})
	if diags.HasErrors() {
		log.Fatalf("Failed to parse HCL: %v", diags)
	}

	// Find name of the resource block
	body := tc.File.Body()
	blocks := body.Blocks()
	for _, block := range blocks {
		if block.Type() == "resource" {
			//log.Printf("Found resource block: %s %v", block.Type(), block.Labels())
			tc.Type = block.Type()
			tc.RscType = block.Labels()[0]
			tc.RscName = block.Labels()[1]
			break
		}
	}

}

func (tc *TerraformConfig) ModifyMapValue(mapName string, key string, value interface{}) {
	// Locate the resource block
	body := tc.File.Body()
	blocks := body.Blocks()

	for _, block := range blocks {
		if block.Labels()[0] == tc.RscType && block.Labels()[1] == tc.RscName {
			// for deeper maps, GetAttribute approach won't work
			// scan entire body
			mapTokens := block.BuildTokens(nil)
			for i, token := range mapTokens {
				if token.Type == hclsyntax.TokenIdent && string(token.Bytes) == key {
					log.Printf("Token found: %s - %v", token.Type, string(token.Bytes))

					for j := i; j < len(mapTokens); j++ {
						if mapTokens[j].Type == hclsyntax.TokenStringLit {
							log.Printf("Replacing value: %s", string(mapTokens[j].Bytes))
							if v, ok := value.(bool); ok {
								mapTokens[j].Bytes = hclwrite.TokensForValue(cty.BoolVal(v)).Bytes()
							} else if v, ok := value.(string); ok {
								mapTokens[j].Bytes = []byte(v)
							} else {
								log.Fatalf("Unsupported type for new value: %T", value)
							}
							break
						}
					}
				}
			}
			tc.File.Body().Clear()
			tc.File.Body().AppendUnstructuredTokens(mapTokens)
			break
		}
	}
	tempFile := new(bytes.Buffer)
	_, err := tc.File.WriteTo(tempFile)
	if err != nil {
		log.Fatalf("Failed to write HCL: %v", err)
	}
	var dg hcl.Diagnostics
	tc.File, dg = hclwrite.ParseConfig(tempFile.Bytes(), "unknown", hcl.Pos{Line: 1, Column: 1})
	if dg.HasErrors() {
		log.Fatalf("Failed to parse HCL: %v", dg)
	}

}

func (tc *TerraformConfig) ModifyMapKey(mapName string, key string, newKey string) {
	log.Printf("Modifying map %s key %s to %s", mapName, key, newKey)
	// Locate the resource block
	body := tc.File.Body()
	blocks := body.Blocks()
	//log.Printf("Found %d blocks", len(blocks))
	for _, block := range blocks {
		if block.Labels()[0] == tc.RscType && block.Labels()[1] == tc.RscName {
			// Locate the "tags" attribute
			log.Printf("Processing type: %s labels %v", block.Type(), block.Labels())
			attr := block.Body().GetAttribute(mapName)
			if attr != nil {
				// Get the current map
				mapTokens := attr.Expr().BuildTokens(nil)
				log.Printf("Current value: %v", mapTokens)
				// Modify the "Environment" map entry
				for i, token := range mapTokens {
					log.Printf("Token type: %s - %v", token.Type, string(token.Bytes))
					if string(token.Bytes) == key {
						// Change the value of "Environment" to "production"
						log.Printf("Replacing %s to %s", key, newKey)
						mapTokens[i].Bytes = []byte(newKey)
						break
					}
				}
				// Set the modified map back to the "tags" attribute
				block.Body().SetAttributeRaw(mapName, mapTokens)
			} else {
				// Nested levels present - search every token for replace
				mapTokens := block.BuildTokens(nil)
				//log.Printf("Current value: %v", mapTokens)
				for i, token := range mapTokens {
					log.Printf("Token type: %s - %v", token.Type, string(token.Bytes))
					if string(token.Bytes) == key {
						// Change the value of "Environment" to "production"
						log.Printf("Token found: %s - %v; replace to %s", token.Type, string(token.Bytes), newKey)
						mapTokens[i].Bytes = []byte(newKey)
						break
					}
				}

				//block.Body().Clear()
				//block.Body().AppendUnstructuredTokens(mapTokens)
				//tc.File.Body().AppendBlock(block)
				tc.File.Body().Clear()
				tc.File.Body().AppendUnstructuredTokens(mapTokens)
			}
		}
	}
	tempFile := new(bytes.Buffer)
	_, err := tc.File.WriteTo(tempFile)
	if err != nil {
		log.Fatalf("Failed to write HCL: %v", err)
	}
	var dg hcl.Diagnostics
	tc.File, dg = hclwrite.ParseConfig(tempFile.Bytes(), "unknown", hcl.Pos{Line: 1, Column: 1})
	if dg.HasErrors() {
		log.Fatalf("Failed to parse HCL: %v", dg)
	}
}

func (tc *TerraformConfig) ModifyAttributeValue(attName string, value interface{}) {
	// Locate the resource block
	body := tc.File.Body()
	blocks := body.Blocks()

	var ctyValue cty.Value
	switch v := value.(type) {
	case string:
		ctyValue = cty.StringVal(v)
	case int:
		ctyValue = cty.NumberIntVal(int64(v))
	case float64:
		ctyValue = cty.NumberFloatVal(v)
	case bool:
		ctyValue = cty.BoolVal(v)
	case []string:
		list := make([]cty.Value, len(v))
		for i, s := range v {
			list[i] = cty.StringVal(s)
		}
		ctyValue = cty.ListVal(list)
	default:
		log.Fatalf("Unsupported type for new value: %T", v)
	}

	for _, block := range blocks {
		if block.Labels()[0] == tc.RscType && block.Labels()[1] == tc.RscName {
			// Locate the "tags" attribute
			attr := block.Body().GetAttribute(attName)
			if attr != nil {
				block.Body().SetAttributeValue(attName, ctyValue)
			}
		}
	}
}

func (tc *TerraformConfig) AddEntryToMap(mapName string, key string, replaceContent map[string]string, replaceEntry bool) {
	// Locate the resource block
	body := tc.File.Body()
	blocks := body.Blocks()
	//log.Printf("Found %d blocks", len(blocks))
	for _, block := range blocks {
		//log.Printf("Block type: %s labels %v", block.Type(), block.Labels())
		if block.Labels()[0] == tc.RscType && block.Labels()[1] == tc.RscName {
			// Locate the "tags" attribute
			attr := block.Body().GetAttribute(mapName)
			if attr != nil {
				// Get the current map and add a new entry
				existingMap := attr.Expr().BuildTokens(nil)
				copiedTokens := hclwrite.Tokens{}
				copyEnd := 0
				copyStart := 2
				replaceNext := ""
				for i, token := range existingMap {
					//log.Printf("%d: Token type: %s - %v", i, token.Type, string(token.Bytes))
					if i >= copyStart && token.Type != hclsyntax.TokenCBrace {
						tk := new(hclwrite.Token)
						if i == 3 && token.Type == hclsyntax.TokenQuotedLit {
							// replace key
							tk.Bytes = []byte(key)
							tk.Type = hclsyntax.TokenQuotedLit
							tk.SpacesBefore = token.SpacesBefore
							copiedTokens = append(copiedTokens, tk)
							continue
						}
						tk.Bytes = make([]byte, len(token.Bytes))
						if token.Type == hclsyntax.TokenIdent {
							replaceNext = replaceContent[string(token.Bytes)]
						}
						if token.Type == hclsyntax.TokenQuotedLit && replaceNext != "" {
							tk.Bytes = []byte(replaceNext)
							replaceNext = ""
						} else {
							copy(tk.Bytes, token.Bytes)
						}
						tk.Type = token.Type
						tk.SpacesBefore = token.SpacesBefore
						copiedTokens = append(copiedTokens, tk)
					}

					if token.Type == hclsyntax.TokenCBrace {
						copiedTokens = append(copiedTokens, token)
						copyEnd = i
						break
					}
				}
				var newMap hclwrite.Tokens
				if replaceEntry {
					newMap = append(existingMap[:copyStart], copiedTokens...)
					newMap = append(newMap, existingMap[copyEnd+1:]...)
				} else {
					newMap = append(existingMap[:len(existingMap)-1], copiedTokens...)
					newMap = append(newMap, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
					newMap = append(newMap, existingMap[len(existingMap)-1:]...)
				}
				log.Print(string(newMap.Bytes()))
				block.Body().SetAttributeRaw(mapName, newMap)
			}

		}
	}
}

func parseGraphWiz(dotContent string) (map[string][]string, error) {
	log.Printf("Parsing DOT file")
	dependencies := make(map[string][]string)
	file := new(bytes.Buffer)
	file.Write([]byte(dotContent))
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`"([^"]+)"\s*->\s*"([^"]+)"`)

	for scanner.Scan() {
		line := scanner.Text()
		// a->b a is dependant on b
		if matches := re.FindStringSubmatch(line); matches != nil {
			//log.Println(line)
			a := matches[1]
			b := matches[2]
			dependencies[a] = append(dependencies[a], b)
			log.Printf("Found dependency: %s -> %s", a, b)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading DOT file: %w", err)
	}
	log.Printf("End Parsing Dot")
	return dependencies, nil
}

// Function to update Terraform files with "depends_on"
func (tc *TerraformConfig) updateDependency(dependencies map[string][]string) {
	log.Printf("Updating dependencies %s.%s.%s", tc.Type, tc.RscType, tc.RscName)
	body := tc.File.Body()
	var block *hclwrite.Block
	for _, block = range body.Blocks() {
		if block.Type() == "resource" {
			log.Printf("Inside Block %s.%v", block.Type(), block.Labels())
			fullName := fmt.Sprintf("%s.%s", block.Labels()[0], block.Labels()[1])
			if deps, exists := dependencies[fullName]; exists {
				log.Printf("Dependencies for %s", fullName)
				depsTokens := make(hclwrite.Tokens, 0)

				depsTokens = append(depsTokens, hclwrite.Tokens{&hclwrite.Token{Type: hclsyntax.TokenOBrack, Bytes: []byte("["), SpacesBefore: 1}}...)
				for _, dep := range deps {
					log.Printf("Dependency  %s", dep)
					depsTokens = append(depsTokens, hclwrite.TokensForIdentifier(dep)...)
					depsTokens = append(depsTokens, hclwrite.Tokens{&hclwrite.Token{Type: hclsyntax.TokenComma, Bytes: []byte(",")}}...)
				}
				depsTokens = append(depsTokens, hclwrite.Tokens{&hclwrite.Token{Type: hclsyntax.TokenCBrack, Bytes: []byte("]")}}...)

				log.Printf("DepsTokens: %v", string(depsTokens.Bytes()))
				block.Body().SetAttributeRaw("depends_on", depsTokens)
			}
			break
		}
	}
}

/*
func main() {
	dotFile := "dependencies.dot"
	tfDir := "terraform_configs"

	// Parse the DOT file to get dependencies
	dependencies, err := parseDotFile(dotFile)
	if err != nil {
		log.Fatalf("Error parsing DOT file: %v", err)
	}

	// Update Terraform files with "depends_on"
	err = updateTerraformFiles(tfDir, dependencies)
	if err != nil {
		log.Fatalf("Error updating Terraform files: %v", err)
	}

	fmt.Println("Terraform files updated with dependencies.")
}
*/
