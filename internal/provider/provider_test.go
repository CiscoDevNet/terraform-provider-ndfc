// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"os"
	"testing"

	helper "terraform-provider-ndfc/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	//"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

var ndProvider provider.Provider

// var dnsClient *DNSClient
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"ndfc": providerserver.NewProtocol6WithError(ndProvider),
}

func TestMain(m *testing.M) {
	accTest := os.Getenv("TF_ACC")
	if accTest == "" {
		//UT run
		os.Exit(m.Run())
	}
	// This is TF acceptance tests, initialize NDFC details
	testConfigPath := os.Getenv("NDFC_TEST_CONFIG_FILE")
	if testConfigPath == "" {
		panic("NDFC_TEST_CONFIG_FILE env variable not set")
	}
	helper.InitConfig(testConfigPath)
	ndProvider = NewNDFCProvider()
	os.Exit(m.Run())
}

func testAccPreCheck(t *testing.T) {

}

/*

func TestAccProvider_Configure(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: `
				provider "ndfc" {
					host     = "example.com"
					username = "testuser"
					password = "testpassword"
					domain   = "example.com"
					insecure = true
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckProviderAttr("provider.ndfc", "host", "example.com"),
					resource.TestCheckResourceAttr("provider.ndfc", "username", "testuser"),
					resource.TestCheckResourceAttr("provider.ndfc", "password", "testpassword"),
					resource.TestCheckResourceAttr("provider.ndfc", "domain", "example.com"),
					resource.TestCheckResourceAttr("provider.ndfc", "insecure", "true"),
				),
			},
		},
	})
}
*/
