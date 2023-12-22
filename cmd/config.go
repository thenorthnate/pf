package cmd

import (
	"os"
	"path/filepath"

	toml "github.com/pelletier/go-toml/v2"
)

func initConfig() {
	dataDir, err := getDataPath()
	if err != nil {
		Fatalf("failed to find configuration path: %v", err)
	}
	configData, err := os.ReadFile(filepath.Join(dataDir, "config.toml"))
	if err != nil {
		Fatalf("failed to read configuration file: %v", err)
	}
	if err = toml.Unmarshal(configData, config); err != nil {
		Fatalf("failed to parse configuration file: %v", err)
	}
	accountsData, err := os.ReadFile(filepath.Join(dataDir, "accounts.toml"))
	if err != nil {
		Fatalf("failed to read accounts file: %v", err)
	}
	if err = toml.Unmarshal(accountsData, accounts); err != nil {
		Fatalf("failed to parse accounts file: %v", err)
	}
}

type Config struct {
	Categories  []string `toml:"categories"`
	Investments Invest   `toml:"invest"`
}

type Invest struct {
	Blend map[string]float64 `toml:"blend"`
}

type AccountFile struct {
	Accounts []Account `toml:"account"`
}

type Account struct {
	Name    string  `toml:"name"`
	Owner   string  `toml:"owner"`
	ID      string  `toml:"id"`
	Balance float64 `toml:"balance"`
}
