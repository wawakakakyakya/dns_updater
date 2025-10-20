package config

import (
	yml "github.com/wawakakakyakya/configloader/yml"
)

type YamlConfigs struct {
	Cfgs      []*YamlConfig `yaml:"lists"`
	GlobalCfg GlobalConfig  `yaml:"global_config"`
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

type CloudDNS struct {
	Credential string `yaml:"credential"`
	ProjectID  string `yaml:"project_id"`
	ZoneName   string `yaml:"zone_name"`
	Name       string `yaml:"name"`
	RecordType string `yaml:"record_type"`
}

type MyDNSConfig struct {
	UserName string `yaml:"username"`
	Pass     string `yaml:"password"`
}

type GoogleDomainConfig struct {
	Name     string `yaml:"name"`
	UserName string `yaml:"username"`
	Pass     string `yaml:"password"`
}

type CloudFlareConfig struct {
	Token   string `yaml:"token"`
	ZoneId  string `yaml:"zone_id"`
	Name    string `yaml:"name"`
	Type    string `yaml:"type"`
	Comment string `yaml:"comment"`
}

type YamlConfig struct {
	Env          string             `yaml:"env"`
	Timeout      int                `yaml:"timeout"`
	CloudDNS     CloudDNS           `yaml:"cloudDNS"`
	MyDNS        MyDNSConfig        `yaml:"mydns"`
	GoogleDomain GoogleDomainConfig `yaml:"googleDomain"`
	CloudFlare   CloudFlareConfig   `yaml:"cloudflare"`
	SkipVefiry   bool               `yaml:"skipVerify"`
}

func LoadYamlConfig() (YamlConfigs, error) {
	defaultLogConfig := NewDefaultLogConfig()
	ycArray := YamlConfigs{GlobalCfg: GlobalConfig{Log: defaultLogConfig}}
	err := yml.Load(configPath, &ycArray)
	if err != nil {
		return ycArray, err
	}

	return ycArray, nil
}
