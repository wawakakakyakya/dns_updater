package config

import (
	yml "github.com/wawakakakyakya/configloader/yml"
)

type YamlConfigs struct {
	Cfgs      []YamlConfig `yaml:"lists"`
	GlobalCfg GlobalConfig `yaml:"global_config"`
}

type GlobalConfig struct {
	Log LogConfig `yaml:"log"`
}

type LogConfig struct {
	Level      int    `yaml:"level"`
	Path       string `yaml:"path"`
	MaxSize    int    `yaml:"max_size"`
	MaxBackups int    `yaml:"max_backups"`
	MaxAge     int    `yaml:"max_age"`
	Compress   bool   `yaml:"compress"`
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
