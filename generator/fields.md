# Auto Generator Input Help

generator/defs/`resource_name`.yaml file acts as input to generate provider code

## Resource Yaml file structure

```
---
resource: 
  name: <resource_name>
  generate_tf_resource: <true/false>
  inherited_model: <true/false>
  attributes:
    - model_name: id
      tf_name: id
      description: Terraform unique Id for the resource
      type: String
      id: true
      ndfc_type: string
      computed: true
      mandatory: false
      optional: false
      payload_hide: true
      tf_hide: false
      include_empty: flase
      requires_replace: false
      tf_requires_replace: false
      use_state: true
      example: "1"
      ...

datasource:
  name: <ds_name>
  generate_tf_resource: <true/false>
  inherited_model: <true/false>
  attributes:
    - model_name: id
      tf_name: id
      description: Terraform unique Id for the resource
      type: String
      id: true
      ndfc_type: string
      computed: true
      mandatory: false
      optional: false
      payload_hide: true
      tf_hide: false
      include_empty: flase
     
      ...

```

## Attribute Fields

### Mandatory Computed Optional Default Values

**Terraform behaviour**
Following describes the terraform behaviour when config fields are marked with Required/Optional/Computed

*Here config file is the TF file*

**Required:**   Any field marked by `Required` must be set in the config file   
Use  `mandatory: true` in the yaml file to mark an attribute as `Required`
**Optional:**   The field may be set in the configuration, but it is not required.    
Use `optional: true` in the yaml file     
*Note:* This attribute cannot be filled in by provider   
**Computed:**   The field is computed and filled in by the provider and cannot be set in the tf file by the user.
**Combinations:**

	**computed_optional:** User can set the value in config, but if not set, it will be computed by provider    
	 * Also need to use PlanModifiers - use_state_for_unknown     
	 * or else it will be empty in plan data triggering provider to think something has changed     
     Use     
     ```optional: true
        computed: true 
     ```    
	**computed_optional with default values:** User can set the value, but if not, a default value will be used    
	* Provider cannot fill this value to a non user configured/default value     
	* No need of planmodifiers as default value is used if its not set
    Use `default_value:`     

**Truth Table**   
```

|---------|----------|----------|---------|---------------------|
|Mandatory| Computed | Optional | Default |  Result             |
|---------|----------|----------|---------|---------------------|
|   true  |   x      |   x      |  x      |  required           |   x => Don't care
|   x     |   x      |   x      |  true   |  computed_optional  |
|   false |   true   |   false  |  false  |  computed  		    	|
|   false |   false  |   true   |  false  |  optional  		    	|
|   false |   true   |   true   |  false  |  computed_optional  |
|---------------------------------------------------------------|
|	            Special Case                |                     |
|		    *** If all are false     ***      |  optional           |
|---------------------------------------------------------------|

```

### model_name    
Parameter name Used in payload of NDFC REST calls    

### tf_name
Attribute name used in Terraform TF file    

### type, ndfc_type
Attribute type to be used in Terraform TF file    
Using `ndfc_type` it can be represented as another type in payload    

*Default Type Mapping*    
```
-------------------------------------------------------------
| type        |  ndfc_type        |  Payload type  (Json)   |
------------------------------------------------------------|
| String      | empty             |  string                 |
| Int64       | empty             |  int64                  |
| Int64       | string            |  string                 |
| Bool        | empty             |  string                 |
| Bool        | bool              |  bool                   |
-------------------------------------------------------------
```
*Special and nested types
```
------------------------------------------------------------------
| type        |  ndfc_type        |  Payload type  (Json)        |
------------------------------------------------------------------
| Map         | empty             |  map[string]<Value Type>     |
| List        | empty             |  []<tf_name>                 |
| Set         | string            |  []<tf_name>                 |
| List:String | empty             |  <tf_name> []string          |
| List:String | csvstring         |  <tf_name> []string          | => Converted to csv on marshall
| Map:String  | empty             |  <tf_name> map[string]string |                
------------------------------------------------------------------
```

### handle_empty
Only applicable for `int64` types
NDFC responses typically contain  empty strings when values of ints are not set
JSON lib work well if its  0 or "0" or attribute itself not present in the payload
But the presense of `attribute: ""` in the json payload will cause the decoding (unmarshall) to error out
when `handle_empty` is set, a customtype is used `Int64Custom` and the marshallers are overrided to handle this cases correctly.

### include_empty
In the go structs generated, the json tags would have `omitempty` so that empty fields are omitted from the payload. But due to inconsistent behaviour across multiple APIs in NDFC, some fields in some APIs are always needed even though they are empty.
Setting this flag would generate the struct tag without `omitempty` so that empty values are included in marshalled payload

  
### requires_replace
Use for bulk resource.
If an  attribute change needs a delete and create, use this flag 

### tf_requires_replace
For non bulk resources, if an attribute change needs a destroy/create, set this to `true`

### use_state
For optional_computed parameters use this flag so that `UseState` plan modifier is added to schema.
Set this flag for   all `computed_optional` cases

### ndfc_nested - name of nested structure
If NDFC payload has nested structs, for making it simpler for users, in tf config, it can be represented flat.
Use this flag for all such attributes in the nested payload

### tf_hide
Control values for provider plugin if any; not exposed to terraform

### payload_hide
Control values that are not send directly to NDFC, but used as a trigger to provider plugin for custom actions    
example:  `deploy:true` used by provider code for trggering deploy operation

### filtered, bulk, sort, list_key
Used for provider code generation 
  `filtered`: Set this field and use to filter out the entry when sending to terraform    
  `SetModeldata` ignores array entries when this field is set to true     

  `bulk`: If the array has large number of entries and need efficient searching, set this field     
        A map is created map[`list_key`]*array_entry for faster searching/accessing the entry     
  `sort`: When a field is set as `sort: true` sort helpers are generated so that go sort library can act on the array.      
        

### unmarshal_tag
Used in special cases where NDFC has different variable names in the payload if POST and GET     
`model_name` is used for POST    
`unmarshal_tag` if present is used to intepret data retrieved via GET APIs    
A duplicate member will be generated in the payload structure to handle this case

**Documentation related fields**
### example
`example` is a very important field that helps with the documentation
A `resource.tf/datasource.tf` file is generated inside `examples/resources(datasources)/<name>/resource.tf(datasource.tf)`
with the content constructed using `example` fields of various attributes
For proper generation of documentation make sure this field is filled in properly

**Note**
`example` in a map entry is used as the map key in the generated example

### import_id
This is a resource level attribute     
Field used to document a sample import ID of the resource      
This is an array and can list multiple IDs if supported    

### import_desc
Explanation of the format of import ID that needs to go into documentation       
This is also array and can provide descriptions of each entry inside import_id

### descriptions
Description field of each attribute goes into schema and into the documentation of that attribute.     
Give relevant information in the field




