package config

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestReadStandard(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "config_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	cfgFilePath := filepath.Join(tmpDir, "config.json")
	cfgContent := `{
		"server": {
			"http_port": 8080,
			"host": "localhost",
			"token_exp_minutes": 60,
			"refresh_token_exp_minute": 1440,
			"token_secret": "secret"
		}
	}`

	if err := os.WriteFile(cfgFilePath, []byte(cfgContent), os.ModePerm); err != nil {
		t.Fatal(err)
	}

	expectedConfig := Config{
		Server: Server{
			HttpPort:               8080,
			Host:                   "localhost",
			TokenExpMinutes:        60,
			RefreshTokenExpMinutes: 1440,
			TokenSecret:            "secret",
		},
	}

	config, err := ReadStandard(cfgFilePath)
	if err != nil {
		t.Fatalf("ReadStandard failed: %v", err)
	}

	if !reflect.DeepEqual(config, expectedConfig) {
		t.Errorf("Expected config: %+v, \ngot: %+v", expectedConfig, config)
	}
}
