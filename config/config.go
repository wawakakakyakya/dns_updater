package config

import "path/filepath"

type ConfigArray []*Config

type Config struct {
	// Excludes    *ArrayFlags
	Excludes    []string
	Src         string
	Dest        string
	Env         string
	Rotate      int
	IsRecursive bool //if bool, can't judg user input, because bool init to false
}

//convert config
func mutate(config *Config) error {
	absDest, err := filepath.Abs(config.Dest)
	if err != nil {
		return err
	}
	absSrc, err := filepath.Abs(config.Src)
	if err != nil {
		return err
	}
	config.Dest = absDest
	config.Src = absSrc
	return nil
}

func NewConfig() (*YamlConfigs, error) {
	cfgAr, err := LoadYamlConfig()
	if err != nil {
		return nil, err
	}

	// for _, c := range cfgAr {
	// 	if err := mutate(c); err != nil {
	// 		return nil, err
	// 	}
	// }

	return &cfgAr, nil
}
