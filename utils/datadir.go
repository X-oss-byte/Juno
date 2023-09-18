package utils

import (
	"errors"
	"os"
	"path/filepath"
)

const junoDirName = "juno"

func DefaultDataDir() (string, error) {
	if userDataDir := os.Getenv("XDG_DATA_HOME"); userDataDir != "" {
		return filepath.Join(userDataDir, junoDirName), nil
	}

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	if userHomeDir == "" {
		return "", errors.New("home directory not found")
	}
	return filepath.Join(userHomeDir, ".local", "share", junoDirName), nil
}
