// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package testing

import (
	"fmt"
	"os"
	"os/exec"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	NDFC NDFCConfig `yaml:"ndfc"`
}

type InventoryDevice struct {
	Device string `yaml:"device"`
	Role   string `yaml:"role"`
}

type InventoryDevices []InventoryDevice
type IntegratedConfig struct {
	Fabric    string           `yaml:"fabric"`
	Switches  []string         `yaml:"switches"`
	VpcPair   []string         `yaml:"vpc_pair"`
	User      string           `yaml:"user"`
	Password  string           `yaml:"pwd"`
	Inventory InventoryDevices `yaml:"inventory_devices"`
}

func (i InventoryDevices) GetDevices() []string {
	var devices []string
	for _, d := range i {
		devices = append(devices, d.Device)
	}
	return devices
}

func (i InventoryDevices) GetRoles() []string {
	var roles []string
	for _, d := range i {
		roles = append(roles, d.Role)
	}
	return roles
}

type NDFCConfig struct {
	URL               string           `yaml:"url"`
	User              string           `yaml:"user"`
	Password          string           `yaml:"pwd"`
	Insecure          string           `yaml:"insecure"`
	Fabric            string           `yaml:"fabric"`
	Switches          []string         `yaml:"switches"`
	VrfPrefix         string           `yaml:"vrf_prefix"`
	NetPrefix         string           `yaml:"net_prefix"`
	VpcPair           []string         `yaml:"vpc_pair"`
	Integration       IntegratedConfig `yaml:"integration_test"`
	mockPort          int
	mockServerStarted bool
	mockConfigFile    string
}

var config map[string]*Config
var testModules = []string{
	"vrf",
	"network",
	"ethernet",
	"loopback",
	"vlan",
	"portchannel",
	"vpc_pair",
}

func LoadConfigFromYAML(yamlContent string) (*Config, error) {
	cf := new(Config)
	err := yaml.Unmarshal([]byte(yamlContent), cf)
	if err != nil {
		return nil, err
	}
	cf.NDFC.mockPort = 0
	cf.NDFC.mockServerStarted = false
	return cf, nil
}

func InitConfig(path string, mock string) {
	if config != nil {
		return
	}
	config = make(map[string]*Config)
	cfg, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	if mock == "" {
		config["global"], err = LoadConfigFromYAML(string(cfg))
		if err != nil {
			panic(err)
		}
		return
	}
	for i, module := range testModules {

		config[module], err = LoadConfigFromYAML(string(cfg))
		if err != nil {
			panic(err)
		}
		config[module].NDFC.mockPort = 3000 + i
		config[module].NDFC.URL = fmt.Sprintf("https://localhost:%d", config[module].NDFC.mockPort)
		switch module {
		case "ethernet":
			fallthrough
		case "loopback":
			fallthrough
		case "vlan":
			fallthrough
		case "portchannel":
			config[module].NDFC.mockConfigFile = "./testdata/interfaces.json"
		default:
			fmt.Println("Mocking not supported for this module")
		}
	}
}

func GetConfig(module string) *Config {
	if len(config) == 0 {
		panic("Config not initialized")
	}
	if len(config) == 1 {
		return config["global"]
	}
	return config[module]
}

func IsMocked() bool {
	return len(config) > 1
}

func StartMockServer(module string) {
	// not mocked
	if len(config) == 1 {
		return
	}
	if config[module].NDFC.mockServerStarted {
		return
	}
	rpath, _ := os.Getwd()
	config[module].NDFC.mockServerStarted = true
	mockScript := rpath + "/../../mock.sh"
	fmt.Println(rpath, rpath+"../../"+config[module].NDFC.mockConfigFile)
	cmd := exec.Command("/bin/bash", mockScript, "start",
		rpath+"/../../"+config[module].NDFC.mockConfigFile,
		fmt.Sprint(config[module].NDFC.mockPort))
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err, string(out))
		panic(err)
	}
	fmt.Println(string(out))

}

func stopMockServer(module string) {
	rpath, _ := os.Getwd()
	rpath = rpath + "/../../mock.sh"

	if len(config) == 1 {
		return
	}
	if !config[module].NDFC.mockServerStarted {
		return
	}
	fmt.Println(rpath)
	cmd := exec.Command("/bin/bash", rpath, "stop",
		fmt.Sprint(config[module].NDFC.mockPort))
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
	config[module].NDFC.mockServerStarted = false

}

func StopMock() {
	if len(config) == 1 {
		return
	}
	for _, module := range testModules {
		stopMockServer(module)
	}
}
