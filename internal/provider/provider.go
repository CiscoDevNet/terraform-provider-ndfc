package provider

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"terraform-provider-ndfc/internal/provider/ndfc"
	"terraform-provider-ndfc/internal/provider/provider/provider_ndfc"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ provider.Provider = (*ndfcProvider)(nil)

func New() func() provider.Provider {
	return NewNDFCProvider
}

func NewNDFCProvider() provider.Provider {
	return &ndfcProvider{}
}

type ndfcProvider struct {
	Version types.String
}

func (p *ndfcProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = provider_ndfc.NdfcProviderSchema(ctx)
	tflog.Info(ctx, "Provider Schema  called")
	log.Printf("Provider Schema  called")
}

func (p *ndfcProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configure Provider   called")
	var config = provider_ndfc.NdfcModel{}
	diags := req.Config.Get(ctx, &config)
	log.Printf("Configure Provider   called")

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	ctx = tflog.SetField(ctx, "ndfc_host", config.Host.ValueString())
	ctx = tflog.SetField(ctx, "ndfc_username", config.Username.ValueString())
	ctx = tflog.SetField(ctx, "ndfc_password", config.Password.ValueString())
	ctx = tflog.SetField(ctx, "ndfc_domain", config.Domain.ValueString())
	ctx = tflog.SetField(ctx, "allow_insecure", config.Insecure.ValueBool())
	ctx = tflog.SetField(ctx, "timeout", config.Timeout.ValueInt64())

	// WORKAROUND default setting in provider schema is not supported
	timeout := int64(1000)
	var err error
	if os.Getenv("NDFC_TIMEOUT") != "" {
		timeout, err = strconv.ParseInt(os.Getenv("NDFC_TIMEOUT"), 10, 64)
		if err != nil {
			resp.Diagnostics.AddError(
				fmt.Sprintf("Invalid input '%s'", os.Getenv("NDFC_TIMEOUT")),
				"A int64 value must be provided for NDFC_TIMEOUT",
			)
			return
		}
	} else if !config.Timeout.IsUnknown() && !config.Timeout.IsNull() {
		timeout = config.Timeout.ValueInt64()
	}

	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "ndfc_password")

	tflog.Debug(ctx, "Creating NDFC client")

	host := os.Getenv("NDFC_HOST")
	domain := os.Getenv("NDFC_DOMAIN")
	user := os.Getenv("NDFC_USER")
	password := os.Getenv("NDFC_PASSWORD")

	if host == "" {
		host = config.Host.ValueString()
	}
	if domain == "" {
		domain = config.Domain.ValueString()
	}

	if user == "" {
		user = config.Username.ValueString()
	}

	if password == "" {
		password = config.Password.ValueString()
	}

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.
	client, err := ndfc.NewNDFCClient(host, user, password, domain, config.Insecure.ValueBool(), timeout)

	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create NDFC API client",
			"Unable to create nd client:\n\n"+err.Error(),
		)
		return
	}

	// For creating mutexes used to lock API access
	// General rule, do not access the same sub-system in NDFC in parallel
	ndfc.NewResource(ndfc.ResourceVrfBulk)
	ndfc.NewResource(ndfc.ResourceNetworks)
	ndfc.NewResource(ndfc.ResourceInterfaces)
	ndfc.NewResource(ndfc.ResourceTemplate)
	ndfc.NewResource(ndfc.ResourcePolicy)

	// Make the HashiCups client available during DataSource and Resource
	// type Configure methods.
	resp.DataSourceData = client
	resp.ResourceData = client
	tflog.Info(ctx, "Configured NDFC client", map[string]any{"success": true})

}

func (p *ndfcProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "ndfc"
	tflog.Info(ctx, "Provider metadata  called")
}

func (p *ndfcProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	tflog.Info(ctx, "Provider Datasources  called")
	return []func() datasource.DataSource{
		NewFabricDataSource,
		NewVrfBulkDataSource,
		NewNetworksDataSource,
		NewInterfacesDataSource,
		NewInventoryDevicesDataSource,
		NewInventoryReachabilityDataSource,
	}
}

func (p *ndfcProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewVrfBulkResource,
		NewNetworksResource,
		NewInterfaceEthernetResource,
		NewInterfaceLoopbackResource,
		NewInterfaceVlanResource,
		NewInventoryDevicesResource,
		NewConfigDeployResource,
		NewInterfacePortChannelResource,
		NewInterfaceVPCResource,
		NewTemplateResource,
		NewPolicyResource,
		NewTestRscResource,
	}
}
