// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package provider_ndfc

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
)

func NdfcProviderSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"domain": schema.StringAttribute{
				Optional:            true,
				Description:         "NDFC Login credentials - domain. Enviroment variable `NDFC_DOMAIN` can be used to override the provider configuration.",
				MarkdownDescription: "NDFC Login credentials - domain. Enviroment variable `NDFC_DOMAIN` can be used to override the provider configuration.",
			},
			"insecure": schema.BoolAttribute{
				Optional:            true,
				Description:         "Controls whether ND server's certificate chain and host name is verified. This can also be set as the `NDFC_INSECURE` (true or false) environment variable.",
				MarkdownDescription: "Controls whether ND server's certificate chain and host name is verified. This can also be set as the `NDFC_INSECURE` (true or false) environment variable.",
			},
			"password": schema.StringAttribute{
				Required:            true,
				Sensitive:           true,
				Description:         "NDFC Login credentials - password. Enviroment variable `NDFC_PASSWORD` can be used to override the provider configuration.",
				MarkdownDescription: "NDFC Login credentials - password. Enviroment variable `NDFC_PASSWORD` can be used to override the provider configuration.",
			},
			"timeout": schema.Int64Attribute{
				Optional:            true,
				Description:         "NDFC HTTP request timeout - timeout. Enviroment variable `NDFC_TIMEOUT` can be used to override the provider configuration.",
				MarkdownDescription: "NDFC HTTP request timeout - timeout. Enviroment variable `NDFC_TIMEOUT` can be used to override the provider configuration.",
			},
			"url": schema.StringAttribute{
				Required:            true,
				Description:         "URL to connect to NDFC - Enviroment variable `NDFC_URL` can be used to override the provider configuration.",
				MarkdownDescription: "URL to connect to NDFC - Enviroment variable `NDFC_URL` can be used to override the provider configuration.",
			},
			"username": schema.StringAttribute{
				Required:            true,
				Description:         "NDFC Login credentials - user.  Enviroment variable `NDFC_USER` can be used to override the provider configuration.",
				MarkdownDescription: "NDFC Login credentials - user.  Enviroment variable `NDFC_USER` can be used to override the provider configuration.",
			},
		},
	}
}

type NdfcModel struct {
	Domain   types.String `tfsdk:"domain"`
	Insecure types.Bool   `tfsdk:"insecure"`
	Password types.String `tfsdk:"password"`
	Timeout  types.Int64  `tfsdk:"timeout"`
	Url      types.String `tfsdk:"url"`
	Username types.String `tfsdk:"username"`
}
