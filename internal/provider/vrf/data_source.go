package vrf

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nd"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &NDFCVRFDataSource{}

func NewNDFCVRFDataSource() datasource.DataSource {
	return &NDFCVRFDataSource{}
}

// ExampleDataSource defines the data source implementation.
type NDFCVRFDataSource struct {
	client *nd.Client
}

// ExampleDataSourceModel describes the data source data model.
type NDFCVRFDataSourceModel struct {
	FabricName types.String `tfsdk:"fabric_name"`
	Vrfs       []VRF        `tfsdk:"vrfs"`
}

func (d *NDFCVRFDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vrfs"
}

func (d *NDFCVRFDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "VRFs on a Given Fabric",
		Attributes: map[string]schema.Attribute{
			"fabric_name": schema.StringAttribute{
				Computed: false,
				Required: true,
			},
			"vrfs": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"vrf_name": schema.StringAttribute{
							Computed: true,
						},
						"vrf_template": schema.StringAttribute{
							Computed: true,
						},
						"vrf_extension_template": schema.StringAttribute{
							Computed: true,
						},
						"vrf_template_config": schema.StringAttribute{
							Computed: true,
						},
						"id": schema.Int64Attribute{
							Computed: true,
						},
						"vrf_id": schema.Int64Attribute{
							Computed: true,
						},
						"vrf_status": schema.StringAttribute{
							Computed: true,
						},
						"fabric": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (d *NDFCVRFDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	tflog.Info(ctx, "VRF Data source Configure")

	client, ok := req.ProviderData.(*nd.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *nd.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *NDFCVRFDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data NDFCVRFDataSourceModel

	var vrf VRF

	// Read Terraform configuration data into the model

	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Error in reading config from tf")
		return
	}

	tflog.Info(ctx, fmt.Sprintf("Read request received fabric %s", data.FabricName))

	if resp.Diagnostics.HasError() {
		return
	}

	url_path := fmt.Sprintf("/lan-fabric/rest/top-down/fabrics/%v/vrfs", data.FabricName.ValueString())

	tflog.Info(ctx, fmt.Sprintf("Read URL %s", url_path))

	res, err := d.client.Get(url_path)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		tflog.Error(ctx, fmt.Sprintf("Client Error: %v", err))
		return
	}

	/*
	       test_out := `[
	   		{
	   		"fabric": "test_evpn_vxlan",
	   		"vrfName": "MyVRF_50003",
	   		"vrfTemplate": "Default_VRF_Universal",
	   		"vrfExtensionTemplate": "Default_VRF_Extension_Universal",
	   		"tenantName": null,
	   		"id": 14,
	   		"vrfId": 10001,
	   		"serviceVrfTemplate": null,
	   		"source": null,
	   		"vrfStatus": "NA",
	   		"hierarchicalKey": "test_evpn_vxlan",
	   		"vrfTemplateConfig": {
	   		"advertiseDefaultRouteFlag": "false",
	   		"routeTargetImport": "1:1",
	   		"vrfVlanId": "1500",
	   		"isRPExternal": "true",
	   		"vrfDescription": "My vrf description",
	   		"disableRtAuto": "true",
	   		"cloudRouteTargetImportEvpn": "1:1",
	   		"L3VniMcastGroup": "233.1.1.1",
	   		"maxBgpPaths": "2",
	   		"maxIbgpPaths": "3",
	   		"routeTargetExport": "1:1",
	   		"ipv6LinkLocalFlag": "false",
	   		"vrfRouteMap": "FABRIC-RMAP-REDIST",
	   		"ENABLE_NETFLOW": "false",
	   		"configureStaticDefaultRouteFlag": "false",
	   		"tag": "11111",
	   		"rpAddress": "1.2.3.4",
	   		"trmBGWMSiteEnabled": "true",
	   		"mvpnInterAs": "false",
	   		"nveId": "1",
	   		"routeTargetExportEvpn": "1:1",
	   		"NETFLOW_MONITOR": "MON1",
	   		"bgpPasswordKeyType": "7",
	   		"bgpPassword": "1234567890ABCDEF",
	   		"mtu": "9200",
	   		"multicastGroup": "234.0.0.0/8",
	   		"isRPAbsent": "false",
	   		"cloudRouteTargetExportEvpn": "1:1",
	   		"advertiseHostRouteFlag": "true",
	   		"vrfVlanName": "VLAN1500",
	   		"trmEnabled": "true",
	   		"loopbackNumber": "100",
	   		"asn": "64000",
	   		"vrfIntfDescription": "My int description",
	   		"routeTargetImportEvpn": "1:1"
	   		}
	   		}
	   		]`
	   	res = gjson.Get(test_out)
	*/
	tflog.Info(ctx, fmt.Sprintf("Retrieved data from NDFC %v", res.Raw))

	err = json.Unmarshal([]byte(res.Get("0").Raw), &vrf)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Json unmarshal error %v", err))
	}

	/*
		if value := res.Get("0.fabric"); value.Exists() && value.String() != "" {
			vrf.Fabric = types.StringValue(value.String())
		} else {
			vrf.Fabric = types.StringNull()
		}
		if value := res.Get("0.vrfName"); value.Exists() && value.String() != "" {
			vrf.VrfName = types.StringValue(value.String())
		} else {
			vrf.VrfName = types.StringNull()
		}
		if value := res.Get("0.vrfTemplate"); value.Exists() && value.String() != "" {
			vrf.VrfTemplate = types.StringValue(value.String())
		} else {
			vrf.VrfTemplate = types.StringNull()
		}
		if value := res.Get("0.vrfExtensionTemplate"); value.Exists() && value.String() != "" {
			vrf.VrfExtensionTemplate = types.StringValue(value.String())
		} else {
			vrf.VrfExtensionTemplate = types.StringNull()
		}
		if value := res.Get("0.vrfId"); value.Exists() && value.String() != "" {
			vrf.VrfId = types.Int64Value(value.Int())
		} else {
			vrf.VrfId = types.Int64Null()
		}
		if value := res.Get("0.vrfTemplateConfig"); value.Exists() && value.String() != "" {
			vrf.VrfTemplateConfig = types.StringValue(value.String())
		}

		if value := res.Get("0.vrfStatus"); value.Exists() && value.String() != "" {
			vrf.VrfStatus = types.StringValue(value.String())
		}

	*/

	data.Vrfs = append(data.Vrfs, vrf)
	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := d.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read example, got error: %s", err))
	//     return
	// }

	// For the purposes of this example code, hardcoding a response value to
	// save into the Terraform state.
	//data.Id = types.StringValue("example-id")

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "read a data source")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
