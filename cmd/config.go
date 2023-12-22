package cmd

import (
	"os"
	"path/filepath"

	toml "github.com/pelletier/go-toml/v2"
)

func initConfig() {
	configDir, err := getConfigPath()
	if err != nil {
		Fatalf("failed to find configuration path: %v", err)
	}
	data, err := os.ReadFile(configDir)
	if err != nil {
		Fatalf("failed to read configuration file: %v", err)
	}
	if err = toml.Unmarshal(data, config); err != nil {
		Fatalf("failed to parse configuration file: %v", err)
	}
}

func getConfigPath() (string, error) {
	if configPath != "" {
		return configPath, nil
	}
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homedir, ".pf", "config.toml"), nil
}

type Config struct {
	Investments Invest `toml:"invest"`
}

type Invest struct {
	Blend map[string]float64 `toml:"blend"`
}
