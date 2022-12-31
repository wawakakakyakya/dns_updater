package config

import (
	"flag"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "config.yml", "config path, default ./config.yml")
	flag.Parse()
}
