package ndfctemplates

// Most of the logic here is generated using AI
// Code is edited to fix certain bugs - most of the cases are tested using UT code
import (
	"encoding/json"
	"fmt"
	"log"

	// "log" - Commented out as part of debug log cleanup
	"regexp"
	"strconv"
	"strings"
)

// FieldDefinition represents a field in the template
type FieldDefinition struct {
	Name         string
	Type         string
	IsMandatory  bool
	IsInternal   bool
	DefaultValue string
	MinLength    int
	MaxLength    int
	Regex        string
	Description  string
	Flags        map[string]string // Store additional flags like IsSourceFabric, IsAsn, etc.
}

// NewFieldDefinition creates a new FieldDefinition with initialized Flags map
func NewFieldDefinition() *FieldDefinition {
	return &FieldDefinition{
		Flags: make(map[string]string),
	}
}

// Template represents the parsed template
type NDFCTemplate struct {
	Fields map[string]FieldDefinition
}

// NewTemplate creates a new Template
func NewNDFCTemplate() *NDFCTemplate {
	return &NDFCTemplate{
		Fields: make(map[string]FieldDefinition),
	}
}

// ParseTemplate parses the template content
func (t *NDFCTemplate) ParseTemplate(templateStr string) error {
	// Find the start of the template variables section
	varStart := strings.Index(templateStr, "##template variables")
	if varStart == -1 {
		//log.Printf("Template variables section not found - parsing entire template")
		// Default to parsing the entire template if no markers are found
		varStart = 0
	} else {
		// Move to the line after ##template variables
		varStart = varStart + len("##template variables")
		// Make sure we start at the beginning of the next line
		nextLine := strings.Index(templateStr[varStart:], "\n")
		if nextLine != -1 {
			varStart = varStart + nextLine + 1
		}
	}

	// Find the end of the variables section - either ##template content or end of string
	varEndIndex := strings.Index(templateStr[varStart:], "##template content")
	varEnd := 0
	if varEndIndex == -1 {
		// No template content marker found, use the whole string from start
		//log.Printf("Template content marker not found - parsing to end of template")
		varEnd = len(templateStr)
	} else {
		varEnd = varStart + varEndIndex
	}

	// Extract the variables section
	variablesSection := templateStr[varStart:varEnd]

	// Split content into lines for processing
	lines := strings.Split(variablesSection, "\n")

	// Variables to track the current field being processed
	var currentField *FieldDefinition
	var inFieldBody bool
	var fieldBodyText strings.Builder

	// Flags to track if we're between an attribute and type definition
	var currentAttrs map[string]string
	var waitingForTypeDef bool

	//log.Printf("Parsing template section between ##template variables and ##template content. %d lines found", len(lines))

	// Process each line of the template
	for _, line := range lines {
		origLine := line
		line = strings.TrimSpace(line)

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		//log.Printf("Processing line %d: %s", i+1, line)

		// Check if we're starting a new field definition with attributes
		if strings.HasPrefix(line, "@(") {
			//log.Printf("Starting new field definition with attributes: %s", line)
			// If we were processing a field body, finish it up
			if inFieldBody && currentField != nil {
				// Process the accumulated field body
				processFieldBody(currentField, fieldBodyText.String())
				// Store the completed field
				t.Fields[currentField.Name] = *currentField
				//log.Printf("Completed field: %s", currentField.Name)
				// Reset field body tracking
				fieldBodyText.Reset()
				inFieldBody = false
			}

			// Process attribute section if line begins with @(...) format

			// Find the matching closing parenthesis for @(...) using proper parenthesis tracking
			/*
				attrEndPos := strings.LastIndex(line, ")")
					if attrEndPos == -1 {
						log.Printf("Invalid field definition (missing matching parenthesis): %s", line)
						continue
					}
					log.Printf("Using fallback parenthesis matching for attribute string: %s", line)
				}
			*/

			// Extract only the attribute part with the surrounding @(...) included
			//attrStr := line[:attrEndPos+1]
			// Parse attributes right away (our parseAttributes function expects the @(...) format)
			currentAttrs = parseAttributes(line)
			//log.Printf("Parsed attributes: %+v", currentAttrs)
			/*
				// Check if the type definition is on the same line, after the attributes
				rest := strings.TrimSpace(line[strings.LastIndex(line, ")")+1:])
				if rest != "" {
					// Process inline type definition
					currentField = processTypeDefinition(t, rest, currentAttrs)
					// Reset attribute tracking
					currentAttrs = nil
					waitingForTypeDef = false
				} else {
			*/
			// Attributes only, wait for type definition on next line
			waitingForTypeDef = true

		} else if waitingForTypeDef {
			// Process type definition that follows an attribute declaration
			//log.Printf("Processing type definition: %s", line)
			currentField = processTypeDefinition(t, line, currentAttrs)

			// Check if this line includes a field body
			if strings.Contains(line, "{") {
				//log.Printf("Field body starts: %s", line)
				inFieldBody = true
				fieldBodyText.WriteString(line[strings.Index(line, "{"):]) // Start from the opening brace
			} else if !strings.Contains(line, ";") {
				// Line does not end with ; this attribute has field body
				inFieldBody = true
				// Field body probably starts in next line as there is no {  here
				fieldBodyText.Reset()
			}
			// Reset attribute tracking
			currentAttrs = nil
			waitingForTypeDef = false
		} else if inFieldBody && currentField != nil {
			//log.Printf("Continuation of fieldBody: %s", origLine)
			// Continuing with a field body from previous lines
			fieldBodyText.WriteString(origLine) // Use original line with whitespace
			fieldBodyText.WriteString("\n")     // Preserve newlines

			// Check if field body is complete
			if strings.HasSuffix(line, "};") {
				// Process the accumulated field body
				processFieldBody(currentField, fieldBodyText.String())
				// Store the completed field
				t.Fields[currentField.Name] = *currentField
				//log.Printf("Completed field with body: %s", currentField.Name)
				// Reset
				currentField = nil
				fieldBodyText.Reset()
				inFieldBody = false
			}
		}
	}

	// Handle any final field that might be still in process
	if inFieldBody && currentField != nil {
		// Process the accumulated field body
		processFieldBody(currentField, fieldBodyText.String())
		// Store the completed field
		t.Fields[currentField.Name] = *currentField
		//log.Printf("Completed final field: %s", currentField.Name)
	}

	// Dump the parsed template for debugging
	fieldCount := len(t.Fields)
	//log.Printf("Finished parsing template with %d fields", fieldCount)
	if fieldCount > 0 {
		t.DumpTemplate()
	} else {
		//log.Printf("WARNING: No fields were parsed from the template. Check your template format.")
	}
	return nil
}

// processTypeDefinition processes a type definition line and returns a pointer to the created field
func processTypeDefinition(t *NDFCTemplate, line string, attrs map[string]string) *FieldDefinition {
	// Split the line to get the type and name
	// Handle any trailing semicolon or opening brace
	baseLine := line

	// Remove opening brace and anything after it if present
	if idx := strings.Index(baseLine, "{"); idx != -1 {
		baseLine = strings.TrimSpace(baseLine[:idx])
	}

	// Remove trailing semicolon
	baseLine = strings.TrimSuffix(baseLine, ";")

	// Split by whitespace to get type and name
	parts := strings.Fields(baseLine)
	if len(parts) < 2 {
		//log.Printf("Invalid type definition format: %s", line)
		return nil
	}

	// Extract type and name
	fieldType := parts[0]
	fieldName := parts[1]

	// Create field
	field := NewFieldDefinition()
	field.Type = fieldType
	field.Name = fieldName

	// Apply attributes
	for k, v := range attrs {
		switch k {
		case "IsMandatory":
			field.IsMandatory = v == "true"
			//log.Printf("Set IsMandatory=%v for field %s", field.IsMandatory, field.Name)
		case "IsInternal":
			field.IsInternal = v == "true"
			//log.Printf("Set IsInternal=%v for field %s", field.IsInternal, field.Name)
		default:
			field.Flags[k] = v
			//log.Printf("Set flag %s=%s for field %s", k, v, field.Name)
		}
	}

	// Store in template map
	t.Fields[field.Name] = *field

	// Return pointer to field for further processing
	copy := t.Fields[field.Name] // Make a copy that we can take the address of
	return &copy
}

// processFieldBody extracts attributes from a field body
func processFieldBody(field *FieldDefinition, body string) {
	// Extract content between curly braces
	body = strings.TrimSpace(body)
	// Remove the opening and closing braces and trailing semicolon
	body = strings.TrimPrefix(body, "{")
	body = strings.TrimSuffix(body, "};")
	body = strings.TrimSuffix(body, "}")
	//log.Printf("FieldBody: %s", body)
	// Split by semicolons
	parts := strings.Split(body, ";")

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		// Try to parse as key=value
		kv := strings.SplitN(part, "=", 2)
		if len(kv) == 2 {
			key := strings.TrimSpace(kv[0])
			value := strings.TrimSpace(kv[1])

			switch key {
			case "defaultValue":
				field.DefaultValue = value
				//log.Printf("Set defaultValue=%s for field %s", value, field.Name)
			case "minLength":
				if i, err := strconv.Atoi(value); err == nil {
					field.MinLength = i
				}
				//log.Printf("Set minLength=%d for field %s", field.MinLength, field.Name)
			case "maxLength":
				if i, err := strconv.Atoi(value); err == nil {
					field.MaxLength = i
				}
				//log.Printf("Set maxLength=%d for field %s", field.MaxLength, field.Name)
			case "regularExpr":
				field.Regex = value
				//log.Printf("Set regex for field %s", field.Name)
			default:
				field.Flags[key] = value
				//log.Printf("Added field attribute %s=%s for field %s", key, value, field.Name)
			}
		}
	}
}

// DumpTemplate outputs the template field definitions to the log
func (t *NDFCTemplate) DumpTemplate() {
	log.Printf("************Start Dumping the template content ****************")
	for _, field := range t.Fields {
		log.Printf("Field: %s", field.Name)
		log.Printf("  Type: %s", field.Type)
		log.Printf("  IsMandatory: %v", field.IsMandatory)
		log.Printf("  IsInternal: %v", field.IsInternal)
		log.Printf("  Flags:")
		// Print all flags to help with debugging
		for k, v := range field.Flags {
			log.Printf("    %s: %s", k, v)
		}
	}
	log.Printf("************End Dumping the template content ****************")
}

// parseAttributes parses the attributes from @() annotation
func parseAttributes(attrStr string) map[string]string {
	//log.Printf("parseAttributes input: %s", attrStr)
	attrs := make(map[string]string)

	// Extract the content between @( and the matching closing )
	// Need to find the real closing parenthesis by tracking nesting
	if !strings.HasPrefix(attrStr, "@(") {
		//log.Printf("Attribute string doesn't start with @(")
		return attrs
	}

	// Remove the opening @(
	attrStr = attrStr[2:]

	// Now find the matching closing parenthesis
	parenCount := 1 // We've already consumed opening parenthesis
	inDoubleQuote := false
	inSingleQuote := false
	endPos := len(attrStr)

	for i, r := range attrStr {
		// Track quotes - they toggle quote state
		if r == '"' && !inSingleQuote {
			inDoubleQuote = !inDoubleQuote
		} else if r == '\'' && !inDoubleQuote {
			inSingleQuote = !inSingleQuote
		}

		// Only count parens if we're outside quotes
		if !inDoubleQuote && !inSingleQuote {
			if r == '(' {
				parenCount++
			} else if r == ')' {
				parenCount--
				// Found the closing parenthesis
				if parenCount == 0 {
					endPos = i
					break
				}
			}
		}
	}

	// Extract the actual attribute content
	attrStr = attrStr[:endPos]

	// Modified approach: Split into key-value pairs using a state machine to respect quoted content
	// The goal is to split on commas, but only when they're outside of quotes or parentheses
	var parts []string
	inDoubleQuote = false
	inSingleQuote = false
	nestedParens := 0
	current := ""

	// Process character by character to properly handle nested structures
	for i, r := range attrStr {
		// Track quotes state - account for escaped quotes
		if r == '"' && !inSingleQuote {
			// Check if escaped (preceded by backslash)
			if i > 0 && attrStr[i-1] != '\\' {
				inDoubleQuote = !inDoubleQuote
			}
		} else if r == '\'' && !inDoubleQuote {
			// Check if escaped (preceded by backslash)
			if i > 0 && attrStr[i-1] != '\\' {
				inSingleQuote = !inSingleQuote
			}
		}

		// Track nested parentheses, but only outside quotes
		if !inDoubleQuote && !inSingleQuote {
			if r == '(' {
				nestedParens++
			} else if r == ')' {
				nestedParens--
			}
		}

		// Split on commas, but only when outside quotes and not in nested parentheses
		if r == ',' && !inDoubleQuote && !inSingleQuote && nestedParens == 0 {
			if strings.TrimSpace(current) != "" {
				parts = append(parts, strings.TrimSpace(current))
			}
			current = ""
			continue
		}

		// Add character to current part
		current += string(r)
	}

	// Add the last part if not empty
	if strings.TrimSpace(current) != "" {
		parts = append(parts, strings.TrimSpace(current))
	}

	//log.Printf("Parsed attribute parts: %v", parts)

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		// Handle key=value pairs
		kv := strings.SplitN(part, "=", 2)
		if len(kv) == 2 {
			// Remove quotes if present
			key := strings.TrimSpace(kv[0])
			value := strings.TrimSpace(kv[1])

			// Better handling of quoted values - make sure we only unquote properly paired quotes
			if len(value) >= 2 {
				if (value[0] == '"' && value[len(value)-1] == '"') ||
					(value[0] == '\'' && value[len(value)-1] == '\'') {
					// Remove the surrounding quotes but keep the content intact
					value = value[1 : len(value)-1]
				}
			}

			// Special handling for boolean values
			if value == "true" || value == "false" {
				attrs[key] = value
			} else {
				attrs[key] = value
			}
			//log.Printf("Added attribute: %s = %s", key, attrs[key])

			// Also add a flag for boolean attributes
			if value == "true" {
				attrs["Is"+key] = "true"
			}
		} else {
			// Handle standalone flags (boolean true)
			key := strings.TrimSpace(part)
			if key != "" {
				attrs[key] = "true"
				//log.Printf("Added boolean flag: %s = true", key)
			}
		}
	}

	//log.Printf("Final attributes: %+v", attrs)

	// Add a debug log for any bugs with the attribute parsing
	if _, exists := attrs["Description"]; exists {
		if strings.Contains(attrs["Description"], ")") {
			//log.Printf("Description contains a closing parenthesis: %s", attrs["Description"])
		}
	}

	return attrs
}

// parseFieldAttributes parses attributes from field body
func parseFieldAttributes(line string) map[string]string {
	//log.Printf("parseFieldAttributes input: %s", line)
	attrs := make(map[string]string)

	// Trim spaces and braces
	line = strings.TrimSpace(line)
	line = strings.TrimPrefix(line, "{")
	line = strings.TrimSuffix(line, "}")

	// Split by semicolon but handle quoted strings
	var parts []string
	inQuotes := false
	current := ""

	for i, r := range line {
		switch r {
		case '\'':
			inQuotes = !inQuotes
		case ';':
			if !inQuotes {
				part := strings.TrimSpace(current)
				if part != "" {
					parts = append(parts, part)
				}
				current = ""
				continue
			}
		}
		current += string(r)
		// Handle the last part
		if i == len(line)-1 && current != "" {
			parts = append(parts, strings.TrimSpace(current))
		}
	}

	//log.Printf("Parsed field attribute parts: %v", parts)

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		kv := strings.SplitN(part, "=", 2)
		if len(kv) == 2 {
			// Remove quotes if present
			key := strings.TrimSpace(kv[0])
			value := strings.TrimSpace(kv[1])
			value = strings.Trim(value, "'\"")
			attrs[key] = value
			//log.Printf("Added field attribute: %s = %s", key, value)
		} else {
			// Handle standalone attributes
			key := strings.TrimSpace(part)
			if key != "" {
				attrs[key] = "true"
				//log.Printf("Added boolean field attribute: %s = true", key)
			}
		}
	}

	//log.Printf("Final field attributes: %+v", attrs)
	return attrs
}

// ValidatePayload validates a JSON payload against the template
func (t *NDFCTemplate) ValidatePayload(payload []byte) (bool, []string) {
	var errors []string

	var model map[string]interface{}
	if err := json.Unmarshal(payload, &model); err != nil {
		return false, []string{fmt.Sprintf("Failed to unmarshal payload: %s", err.Error())}
	}

	// We only care about fields defined in the template itself
	// This is important because implementation code may use variables that
	// are not defined in the template variables section

	// Convert model to string map for easier comparison
	stringModel := toStringMap(model)

	// First pass: Check if field should be shown based on IsShow conditions
	fieldVisible := make(map[string]bool)
	for name, field := range t.Fields {
		// Default is visible unless there's an IsShow condition
		isVisible := true

		// Check if field has IsShow condition
		if showCondition, hasShowCondition := field.Flags["IsShow"]; hasShowCondition && showCondition != "" {
			isVisible = evaluateShowCondition(showCondition, model, stringModel)
		}

		fieldVisible[name] = isVisible
	}

	// Second pass: Validate required fields and values
	//log.Printf("[DEBUG] ValidatePayload: Second pass: Validate required fields and values")

	// Loop through defined template fields - we ONLY validate fields actually defined in the template
	// This intentionally ignores any fields used only in implementation code but not defined in template variables
	for name, field := range t.Fields {
		//log.Printf("[DEBUG] ValidatePayload: Field: %s", name)

		// Skip fields that aren't visible due to IsShow conditions
		if visible, exists := fieldVisible[name]; exists && !visible {
			continue
		}

		value, exists := model[name]
		if !exists {
			// Skip internal fields
			if field.IsInternal {
				continue
			}

			// Only require fields that should be visible
			if field.IsMandatory && field.DefaultValue == "" {
				errors = append(errors, fmt.Sprintf("missing required field: %s", name))
			}
			continue
		}

		// Type checking
		strValue := stringModel[name]
		if field.Type != "interface" && field.Type != "boolean" && value == nil {
			errors = append(errors, fmt.Sprintf("field %s must not be null", name))
			continue
		}

		// Length validation for string fields
		if _, isString := value.(string); isString {
			if field.MinLength > 0 && len(strValue) < field.MinLength {
				errors = append(errors,
					fmt.Sprintf("field %s is too short (min %d characters)", name, field.MinLength))
			}

			if field.MaxLength > 0 && len(strValue) > field.MaxLength {
				errors = append(errors,
					fmt.Sprintf("field %s is too long (max %d characters)", name, field.MaxLength))
			}

			// Regex validation
			if field.Regex != "" {
				re, err := regexp.Compile(field.Regex)
				if err != nil {
					errors = append(errors,
						fmt.Sprintf("invalid regex for field %s: %v", name, err))
					continue
				}
				if !re.MatchString(strValue) {
					errors = append(errors,
						fmt.Sprintf("field %s does not match required pattern", name))
				}
			}
		}
	}

	return len(errors) == 0, errors
}

func (t *NDFCTemplate) IsInternal(fieldName string) bool {
	field, exists := t.Fields[fieldName]
	if !exists {
		return false
	}
	return field.IsInternal
}

func (t *NDFCTemplate) IsMandatory(fieldName string) bool {
	field, exists := t.Fields[fieldName]
	if !exists {
		return false
	}
	return field.IsMandatory
}

func (t *NDFCTemplate) FillDefaultValues(payload *map[string]string) []string {
	for name, field := range t.Fields {
		_, exists := (*payload)[name]
		if !exists && !field.IsInternal {
			if field.DefaultValue != "" {
				(*payload)[name] = field.DefaultValue
			}
		}
	}
	return nil
}

func (t *NDFCTemplate) GetFieldWithFlag(flagName string) string {
	//log.Printf("Searching for field with flag: %s", flagName)

	// Try different variations of the flag name
	variations := []string{
		flagName,                  // exact match
		"Is" + flagName,           // IsSourceDevice
		strings.ToLower(flagName), // sourcedevice
		strings.ToUpper(flagName), // SOURCEDEVICE
	}

	// For each field in the template
	for name, field := range t.Fields {
		//log.Printf("Checking field: %s", name)
		//log.Printf("  Field type: %s, Mandatory: %v, Internal: %v", field.Type, field.IsMandatory, field.IsInternal)
		//log.Printf("  Field flags: %+v", field.Flags)

		// Check each variation of the flag name
		for _, variation := range variations {
			// Check for exact match
			if value, ok := field.Flags[variation]; ok {
				//log.Printf("  Found matching flag '%s' with value: %s", variation, value)
				if value == "true" || value == "" {
					//log.Printf("  Returning field: %s", name)
					return name
				}
			}

			// Check if the flag is part of a key-value pair
			for k, v := range field.Flags {
				// Check if the key contains our variation (case insensitive)
				if strings.Contains(strings.ToLower(k), strings.ToLower(variation)) {
					//log.Printf("  Found partial match in key '%s' with value: %s", k, v)
					if v == "true" || v == "" {
						//log.Printf("  Returning field based on partial match: %s", name)
						return name
					}
				}
			}
		}

		// Special case: Check if the field name itself matches the flag we're looking for
		if strings.Contains(strings.ToLower(name), strings.ToLower(flagName)) {
			//log.Printf("  Field name '%s' matches flag '%s'", name, flagName)
			return name
		}
	}

	// If we get here, we didn't find the flag
	//log.Printf("No field found with flag: %s", flagName)
	//log.Printf("Available fields: %v", getFieldNames(t.Fields))
	return ""
}

// Helper function to get field names from a map of fields
func getFieldNames(fields map[string]FieldDefinition) []string {
	names := make([]string, 0, len(fields))
	for name := range fields {
		names = append(names, name)
	}
	return names
}

// findMatchingParenthesis locates the matching closing parenthesis
// startPos should be the position after the opening parenthesis
func findMatchingParenthesis(s string, startPos int) int {
	// Input validation
	if startPos >= len(s) || startPos < 0 {
		return -1
	}

	// For NDFC template parsing, where the template attributes often
	// contain parentheses inside quoted strings that shouldn't count
	// for balancing. We determine if this is an attribute string by
	// examining the string prefix.
	isAttribute := false
	if len(s) >= 2 && s[0] == '@' && s[1] == '(' {
		isAttribute = true
		//log.Printf("Processing attribute string with length %d", len(s))
	}

	// Initialize parsing state
	parenCount := 1 // We've already encountered the opening parenthesis
	inSingleQuote := false
	inDoubleQuote := false

	// Special handling for template attribute strings with quoted parentheses
	if isAttribute {
		// Scan from the start position to the end, tracking quote state and parenthesis balance
		for i := startPos; i < len(s); i++ {
			c := s[i]

			// Handle quote state
			if c == '"' {
				// Only toggle double quote state when outside single quotes and not escaped
				if !inSingleQuote {
					// Check for escaping
					if i > 0 && s[i-1] != '\\' {
						inDoubleQuote = !inDoubleQuote
					}
				}
			} else if c == '\'' {
				// Only toggle single quote state when outside double quotes and not escaped
				if !inDoubleQuote {
					// Check for escaping
					if i > 0 && s[i-1] != '\\' {
						inSingleQuote = !inSingleQuote
					}
				}
			}

			// Only count parentheses when outside quotes
			if !inSingleQuote && !inDoubleQuote {
				if c == '(' {
					parenCount++
				} else if c == ')' {
					parenCount--
					if parenCount == 0 {
						return i // Found matching closing parenthesis
					}
				}
			}
		}

		// Special case fallback for NDFC templates if we couldn't find a proper match
		// Often the attribute strings are complex and may not properly balance
		// In such cases, find the closing parenthesis of the attribute declaration
		if lastIdx := strings.LastIndex(s, ")"); lastIdx > startPos {
			return lastIdx
		}

		//log.Printf("Failed to find matching parenthesis in attribute string: %s", s)
		return -1
	}

	// Standard parenthesis matching for non-attribute strings
	for i := startPos; i < len(s); i++ {
		c := s[i]

		// Track quote state
		if c == '"' && !inSingleQuote {
			// Check if quote is escaped
			if i > 0 && s[i-1] != '\\' {
				inDoubleQuote = !inDoubleQuote
			}
		} else if c == '\'' && !inDoubleQuote {
			// Check if quote is escaped
			if i > 0 && s[i-1] != '\\' {
				inSingleQuote = !inSingleQuote
			}
		}

		// Only count parentheses if outside of quotes
		if !inDoubleQuote && !inSingleQuote {
			if c == '(' {
				parenCount++
			} else if c == ')' {
				parenCount--
				if parenCount == 0 {
					return i // Found the matching closing parenthesis
				}
			}
		}
	}

	//log.Printf("No matching parenthesis found in string: %s", s)
	return -1 // No matching parenthesis found
}

// toStringMap converts a map with interface{} values to a map with string values
func toStringMap(data map[string]interface{}) map[string]string {
	result := make(map[string]string)
	for k, v := range data {
		switch val := v.(type) {
		case string:
			result[k] = val
		case bool:
			result[k] = strconv.FormatBool(val)
		case int:
			result[k] = strconv.Itoa(val)
		case float64:
			result[k] = strconv.FormatFloat(val, 'f', -1, 64)
		default:
			if v != nil {
				result[k] = fmt.Sprintf("%v", val)
			} else {
				result[k] = ""
			}
		}
	}
	return result
}

// evaluateShowCondition evaluates IsShow condition expressions like "FIELD==value"
// or "FIELD!=value" to determine if a field should be shown based on other field values
func evaluateShowCondition(condition string, rawData map[string]interface{}, stringData map[string]string) bool {
	//log.Printf("Evaluating condition: %s", condition)

	// Handle complex expressions with && and || operators
	if strings.Contains(condition, "&&") {
		parts := strings.Split(condition, "&&")
		for _, part := range parts {
			if !evaluateShowCondition(strings.TrimSpace(part), rawData, stringData) {
				return false
			}
		}
		return true
	}

	if strings.Contains(condition, "||") {
		parts := strings.Split(condition, "||")
		for _, part := range parts {
			if evaluateShowCondition(strings.TrimSpace(part), rawData, stringData) {
				return true
			}
		}
		return false
	}

	// Handle equality check: FIELD==value
	if strings.Contains(condition, "==") {
		parts := strings.SplitN(condition, "==", 2)
		fieldName := strings.TrimSpace(parts[0])
		expectedValue := strings.TrimSpace(parts[1])
		actualValue, exists := stringData[fieldName]

		// For boolean values, handle special cases
		if expectedValue == "true" || expectedValue == "false" {
			boolVal, exists := rawData[fieldName]
			if exists {
				if bVal, ok := boolVal.(bool); ok {
					return (expectedValue == "true") == bVal
				}
			}
		}

		return exists && actualValue == expectedValue
	}

	// Handle inequality check: FIELD!=value
	if strings.Contains(condition, "!=") {
		parts := strings.SplitN(condition, "!=", 2)
		fieldName := strings.TrimSpace(parts[0])
		expectedValue := strings.TrimSpace(parts[1])
		actualValue, exists := stringData[fieldName]

		// For boolean values, handle special cases
		if expectedValue == "true" || expectedValue == "false" {
			boolVal, exists := rawData[fieldName]
			if exists {
				if bVal, ok := boolVal.(bool); ok {
					return (expectedValue == "true") != bVal
				}
			}
		}

		return !exists || actualValue != expectedValue
	}

	//log.Printf("Warning: Could not evaluate condition: %s", condition)
	return true // Default to showing the field if condition can't be parsed
}

/*
func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: template_validator <template_file> <json_file>")
		os.Exit(1)
	}

	// Read template file
	templateContent, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error reading template file: %v\n", err)
		os.Exit(1)
	}

	// Parse template
	template := NewNDFCTemplate()
	if err := template.ParseTemplate(string(templateContent)); err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		os.Exit(1)
	}

	// Read and parse JSON file
	jsonContent, err := os.ReadFile(os.Args[2])
	if err != nil {
		fmt.Printf("Error reading JSON file: %v\n", err)
		os.Exit(1)
	}

	var payload map[string]interface{}
	if err := json.Unmarshal(jsonContent, &payload); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	// Validate payload
	valid, errors := template.ValidatePayload(payload)
	if valid {
		fmt.Println("✅ Payload is valid!")
	} else {
		fmt.Println("❌ Validation errors:")
		for _, err := range errors {
			fmt.Printf("  - %s\n", err)
		}
		os.Exit(1)
	}
}
*/
