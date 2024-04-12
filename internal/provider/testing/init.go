package testing

import (
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	NDFC NDFCConfig `yaml:"ndfc"`
}

type NDFCConfig struct {
	URL       string   `yaml:"url"`
	User      string   `yaml:"user"`
	Password  string   `yaml:"pwd"`
	Insecure  string   `yaml:"insecure"`
	Fabric    string   `yaml:"fabric"`
	Switches  []string `yaml:"switches"`
	VrfPrefix string   `yaml:"vrf_prefix"`
	NetPrefix string   `yaml:"net_prefix"`
}

var config *Config

func LoadConfigFromYAML(yamlContent string) (*Config, error) {
	config = new(Config)
	err := yaml.Unmarshal([]byte(yamlContent), &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func InitConfig(path string) {
	if config != nil {
		return
	}
	config = new(Config)
	cfg, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	config, err = LoadConfigFromYAML(string(cfg))
	if err != nil {
		panic(err)
	}
}

func GetConfig() *Config {
	if config == nil {
		panic("Config not initialized")
	}
	return config
}
