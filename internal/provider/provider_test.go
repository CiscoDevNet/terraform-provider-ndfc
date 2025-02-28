// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"os"
	"testing"
	"time"

	helper "terraform-provider-ndfc/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
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
	mockedServer := os.Getenv("NDFC_MOCKED_SERVER")
	helper.InitConfig(testConfigPath, mockedServer)
	ndProvider = NewNDFCProvider()
	res := m.Run()
	helper.StopMock()
	os.Exit(res)
}

func testAccPreCheck(t *testing.T, module string) {
	t.Logf("Starting testAccPreCheck for %s", module)
	if !helper.IsMocked() {
		return
	}
	go helper.StartMockServer(module)
	time.Sleep(10 * time.Second)
}

func TestAccProvider_Configure(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t, "provider") },
		Steps: []resource.TestStep{
			{
				Config: `
				terraform {
					required_providers {
						ndfc = {
						source = "registry.terraform.io/cisco/ndfc"
						}
					}
				}

				provider "ndfc" {
					host     = "example.com"
					username = "testuser"
					password = "testpassword"
					domain   = "example.com"
					insecure = true
				}`,
			},
		},
	})
}

func TestAccProviderIntegrationTest(t *testing.T) {
	rsList := []string{
		"ndfc_fabric_vxlan_evpn",
		"ndfc_vpc_pair",
		"ndfc_interface_ethernet",
		"ndfc_interface_loopback",
		"ndfc_interface_portchannel",
		"ndfc_interface_vlan",
		"ndfc_interface_vpc",
		"ndfc_networks",
		"ndfc_vrfs",
		"ndfc_policy",
		"ndfc_inventory_devices",
	}

	if helper.GetConfig("global").NDFC.Integration.Fabric == "" {
		t.Skip("Skipping TestAccProviderIntegrationTest, as configuration is not set in config file")
	}

	attrs := map[string]interface{}{
		"fabric":            helper.GetConfig("global").NDFC.Integration.Fabric,
		"user":              helper.GetConfig("global").NDFC.Integration.User,
		"password":          helper.GetConfig("global").NDFC.Integration.Password,
		"vpc_pair":          helper.GetConfig("global").NDFC.Integration.VpcPair,
		"inventory_devices": helper.GetConfig("global").NDFC.Integration.Inventory.GetDevices(),
		"inventory_roles":   helper.GetConfig("global").NDFC.Integration.Inventory.GetRoles(),
		"switches":          helper.GetConfig("global").NDFC.Integration.Switches,
	}

	providerAttrs := map[string]interface{}{
		"Host":     helper.GetConfig("global").NDFC.URL,
		"User":     helper.GetConfig("global").NDFC.User,
		"Password": helper.GetConfig("global").NDFC.Password,
		"Insecure": helper.GetConfig("global").NDFC.Insecure,
	}

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t, "provider") },
		Steps: []resource.TestStep{
			{
				Config: func() string {
					cfg := helper.GetProviderHeader()
					cfg = cfg + "\n\n" + helper.GetProviderConfig(providerAttrs)
					cfg = cfg + "\n\n" + helper.GetTFIntegrated(t.Name(), rsList, attrs, true)
					return cfg
				}(),
			},
		},
	})
}

func TestAccProviderIntegrationTestGlobalDeploy(t *testing.T) {
	rsList := []string{
		"ndfc_fabric_vxlan_evpn",
		"ndfc_vpc_pair",
		"ndfc_interface_ethernet",
		"ndfc_interface_loopback",
		"ndfc_interface_portchannel",
		"ndfc_interface_vlan",
		"ndfc_interface_vpc",
		"ndfc_networks",
		"ndfc_vrfs",
		"ndfc_policy",
		"ndfc_inventory_devices",
		"ndfc_configuration_deploy",
	}

	if helper.GetConfig("global").NDFC.Integration.Fabric == "" {
		t.Skip("Skipping TestAccProviderIntegrationTest, as configuration is not set in config file")
	}

	attrs := map[string]interface{}{
		"fabric":            helper.GetConfig("global").NDFC.Integration.Fabric,
		"user":              helper.GetConfig("global").NDFC.Integration.User,
		"password":          helper.GetConfig("global").NDFC.Integration.Password,
		"vpc_pair":          helper.GetConfig("global").NDFC.Integration.VpcPair,
		"inventory_devices": helper.GetConfig("global").NDFC.Integration.Inventory.GetDevices(),
		"inventory_roles":   helper.GetConfig("global").NDFC.Integration.Inventory.GetRoles(),
		"switches":          helper.GetConfig("global").NDFC.Integration.Switches,
	}

	providerAttrs := map[string]interface{}{
		"Host":     helper.GetConfig("global").NDFC.URL,
		"User":     helper.GetConfig("global").NDFC.User,
		"Password": helper.GetConfig("global").NDFC.Password,
		"Insecure": helper.GetConfig("global").NDFC.Insecure,
	}

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t, "provider") },
		Steps: []resource.TestStep{
			{
				Config: func() string {
					cfg := helper.GetProviderHeader()
					cfg = cfg + "\n\n" + helper.GetProviderConfig(providerAttrs)
					cfg = cfg + "\n\n" + helper.GetTFIntegrated(t.Name(), rsList, attrs, false)
					return cfg
				}(),
			},
		},
	})
}
