package cmd

import (
	"fmt"
	"os"
	"path/filepath"
)

func Fatalf(message string, args ...any) {
	panic(fmt.Sprintf(message, args...))
}

func getDataPath() (string, error) {
	if dataDir != "" {
		if _, err := os.Lstat(dataDir); err != nil {
			return "", err
		}
		return dataDir, nil
	}
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homedir, ".pf"), nil
}
