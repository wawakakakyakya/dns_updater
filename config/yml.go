package config

import (
	yml "github.com/wawakakakyakya/configloader/yml"
)

type YamlConfigs struct {
	Cfgs []YamlConfig `yaml:"lists"`
}

type GoogleConfig struct {
	Credential string `yaml:"credential"`
	ProjectID  string `yaml:"project_id"`
	ZoneName   string `yaml:"zone_name"`
	Domain     string `yaml:"domain"`
	RecordType string `yaml:"record_type"`
}

type MyDNSConfig struct {
	UserName string `yaml:"username"`
	Pass     string `yaml:"password"`
}

type GoogleDomainConfig struct {
	Domain   string `yaml:"domain"`
	UserName string `yaml:"username"`
	Pass     string `yaml:"password"`
}

type YamlConfig struct {
	Env          string             `yaml:"env"`
	Timeout      int                `yaml:"timeout"`
	Google       GoogleConfig       `yaml:"google"`
	MyDNS        MyDNSConfig        `yaml:"mydns"`
	GoogleDomain GoogleDomainConfig `yaml:"google_domain"`
}

func LoadYamlConfig() (YamlConfigs, error) {

	ycArray := YamlConfigs{}
	err := yml.Load(configPath, &ycArray)
	if err != nil {
		return ycArray, err
	}

	return ycArray, nil
}
