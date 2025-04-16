package ndfc

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"strings"
	"terraform-provider-ndfc/internal/provider/datasources/datasource_rest_api"
	"terraform-provider-ndfc/internal/provider/ndfc/api"
	"terraform-provider-ndfc/internal/provider/resources/resource_rest_api"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func (c NDFC) RscCreateRestAPI(ctx context.Context, diags *diag.Diagnostics, data *resource_rest_api.RestApiModel) {
	// Create a new client
	tflog.Debug(ctx, "Creating rest resource", map[string]interface{}{
		"method": data.Method.ValueString(),
		"url":    data.Url.ValueString(),
	})
	restApi := api.NewRestAPI(c.GetLock(ResourceRestAPI), &c.apiClient)
	var err error
	var resp gjson.Result
	restApi.Url = data.Url.ValueString()
	switch data.Method.ValueString() {
	case "POST":
		resp, err = restApi.Post([]byte(data.Payload.ValueString()))
	case "PUT":
		resp, err = restApi.Put([]byte(data.Payload.ValueString()))
	case "DELETE":
		restApi.DeleteQp = data.DeleteParameters.ValueString()
		resp, err = restApi.Delete()
	default:
		diags.AddError("Invalid method", "The method must be one of POST, PUT, or DELETE")
	}
	if err != nil {
		log.Printf("[ERR] Failed to make API request: %s", err)
		diags.AddError(fmt.Sprintf("%s failed on API %s", data.Method.ValueString(), restApi.Url), err.Error())
		return
	}
	data.ResponseMessage = types.StringValue(resp.String())
	/* Set the checksum of payload as Id */

	data.Id = types.StringValue(computeMD5Sum(data.Payload.ValueString()))
}

func (c NDFC) DsGetRestAPI(ctx context.Context, data *datasource_rest_api.RestApiModel, dg *diag.Diagnostics) {
	if data.Url.IsNull() {
		dg.AddError("Url is required", "Url is a required field for the rest_api resource")
		return
	}

	restApi := api.NewRestAPI(c.GetLock(ResourceRestAPI), &c.apiClient)
	restApi.Url = data.Url.ValueString()

	if !data.QueryParameters.IsNull() {
		restApi.Url = restApi.Url + "?=" + data.QueryParameters.ValueString()
	}

	resp, err := restApi.Get()
	if err != nil {
		dg.AddError("Failed to make API request", err.Error())
		return
	}
	data.ResponseMessage = types.StringValue(string(resp))
}

func (c NDFC) RscGetRestAPI(ctx context.Context, data *resource_rest_api.RestApiModel, dg *diag.Diagnostics) {
	if data.Id.IsNull() || data.Id.IsUnknown() {
		dg.AddError("Id is required", "Id is a required field for the rest_api resource")
		return
	}

	if !data.Stateful.ValueBool() {
		log.Printf("[DEBUG] Stateful is false, not reading NDFC; setting empty payload")
		data.Payload = types.StringValue("")
		return
	}
	restApi := api.NewRestAPI(c.GetLock(ResourceRestAPI), &c.apiClient)
	if data.ReadUrl.ValueString() != "" {
		restApi.Url = data.ReadUrl.ValueString()
	} else {
		log.Printf("[DEBUG] Read URL not specified, not reading NDFC")
		return
	}
	if !data.ResponseMessage.IsNull() {
		restApi.Url = processTemplate(data.ResponseMessage.ValueString(), restApi.Url)
	}

	resp, err := restApi.Get()
	if err != nil {
		dg.AddError("Failed to make API request", err.Error())
		return
	}
	data.Payload = types.StringValue(string(resp))
}

func (c NDFC) RscUpdateRestAPI(ctx context.Context, diags *diag.Diagnostics,
	data *resource_rest_api.RestApiModel,
	state *resource_rest_api.RestApiModel) {
	// Create a new client
	tflog.Debug(ctx, "Updating rest resource", map[string]interface{}{
		"method": data.Method.ValueString(),
		"url":    data.Url.ValueString(),
	})
	restApi := api.NewRestAPI(c.GetLock(ResourceRestAPI), &c.apiClient)
	var err error
	var resp gjson.Result
	updateMethod := "PUT"
	payload := data.Payload.ValueString()

	restApi.Url = data.Url.ValueString()
	if !data.UpdateMethod.IsNull() {
		updateMethod = data.UpdateMethod.ValueString()
	}
	if !data.UpdatePayload.IsNull() {
		payload = data.UpdatePayload.ValueString()
	}

	if data.UpdateUrl.IsNull() {
		restApi.Url = data.Url.ValueString()
	} else {
		restApi.Url = data.UpdateUrl.ValueString()
	}

	// Process any templates
	restApi.Url = processTemplate(state.ResponseMessage.ValueString(), restApi.Url)
	payload = processTemplate(state.ResponseMessage.ValueString(), payload)

	switch updateMethod {
	case "POST":
		resp, err = restApi.Post([]byte(payload))
	case "PUT":
		resp, err = restApi.Put([]byte(payload))
	default:
		diags.AddError("Invalid method", "The method must be one of POST, PUT")
	}
	if err != nil {
		log.Printf("[ERR] Failed to make API request: %s", err)
		diags.AddError("Failed to make API request", err.Error())
		return
	}
	data.ResponseMessage = types.StringValue(resp.String())
	/* Set the checksum of payload as Id */
	data.Id = types.StringValue(computeMD5Sum(data.Payload.ValueString()))
}

func (c NDFC) RscDeleteRestAPI(ctx context.Context, data *resource_rest_api.RestApiModel, dg *diag.Diagnostics) {
	// Create a new client
	tflog.Debug(ctx, "Deleting rest resource", map[string]interface{}{
		"method":         data.Method.ValueString(),
		"url":            data.Url.ValueString(),
		"delete_url":     data.DeleteUrl.ValueString(),
		"delete_method":  data.DeleteMethod.ValueString(),
		"delete_qp":      data.DeleteParameters.ValueString(),
		"delete_payload": data.DeletePayload.ValueString(),
	})
	restApi := api.NewRestAPI(c.GetLock(ResourceRestAPI), &c.apiClient)
	if data.DeleteUrl.IsNull() {
		restApi.Url = data.Url.ValueString()
	} else {
		restApi.Url = data.DeleteUrl.ValueString()
	}

	// Process any templating in the URL
	restApi.Url = processTemplate(data.ResponseMessage.ValueString(), restApi.Url)
	deleteMethod := "DELETE"
	payload := ""

	if !data.DeleteMethod.IsNull() {
		deleteMethod = data.DeleteMethod.ValueString()
	}
	if !data.DeletePayload.IsNull() {
		payload = data.DeletePayload.ValueString()
	}

	// Process any templatws in the payload
	payload = processTemplate(data.ResponseMessage.ValueString(), payload)

	restApi.DeleteQp = data.DeleteParameters.ValueString()

	var resp gjson.Result
	var err error
	switch deleteMethod {
	case "DELETE":
		resp, err = restApi.Delete()
	case "POST":
		resp, err = restApi.Post([]byte(payload))
	case "PUT":
		resp, err = restApi.Put([]byte(payload))
	default:
		dg.AddError("Invalid method", "The method must be one of POST, PUT, or DELETE")
	}
	if err != nil {
		dg.AddError("Failed to execute API request", err.Error())
		return
	}
	log.Printf("Delete response: %s", resp.String())
}

func computeMD5Sum(payload string) string {
	hash := md5.Sum([]byte(payload))
	return hex.EncodeToString(hash[:])
}

func processTemplate(payload string, url string) string {
	if strings.Contains(url, "{{") {
		log.Printf("[DEBUG] Template variables present in : %s", url)
		tmpl, err := template.New("url").Parse(url)
		if err != nil {
			log.Printf("[ERR] Failed to parse template: %v", err.Error())
			return url
		}
		var buf strings.Builder
		tmplData := make(map[string]interface{})
		d := json.NewDecoder(strings.NewReader(payload))
		d.UseNumber()
		err = d.Decode(&tmplData)
		//err = json.Unmarshal([]byte(payload), &tmplData)
		if err != nil {
			log.Printf("[ERR] Failed to unmarshal response message %v", err.Error())
			return url
		}
		//log.Printf("[DEBUG] Template data: %v", tmplData)
		err = tmpl.Execute(&buf, &tmplData)
		if err != nil {
			log.Printf("[ERR] Failed to execute template %v", err.Error())
			return url
		}
		log.Printf("[DEBUG] Updated string after template processing: %s", buf.String())
		return buf.String()
	}
	log.Printf("[DEBUG] No Template variables: %s", url)
	return url
}
