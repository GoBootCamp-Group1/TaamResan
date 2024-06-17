package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func ReadGeneric[T any](cfgPath string) (T, error) {
	var cfg T
	fullAbsPath, err := absPath(cfgPath)
	if err != nil {
		return cfg, err
	}

	content, err := os.ReadFile(fullAbsPath)
	if err != nil {
		return cfg, err
	}

	if err := json.Unmarshal(content, &cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}

func ReadStandard(cfgPath string) (Config, error) {
	return ReadGeneric[Config](cfgPath)
}

func absPath(cfgPath string) (string, error) {
	if !filepath.IsAbs(cfgPath) {
		return filepath.Abs(cfgPath)
	}
	return cfgPath, nil
}
