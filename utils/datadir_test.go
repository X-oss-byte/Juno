package utils_test

import (
	"testing"

	"github.com/NethermindEth/juno/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDataDir(t *testing.T) {
	tests := map[string]struct {
		dataDir     string
		homeDir     string
		expectedDir string
	}{
		"data dir exists without home dir": {
			dataDir:     "dataDir",
			homeDir:     "",
			expectedDir: "dataDir/juno",
		},
		"data dir and home dir exist": {
			dataDir:     "dataDir",
			homeDir:     "homeDir",
			expectedDir: "dataDir/juno",
		},
		"home dir exists without data dir": {
			dataDir:     "",
			homeDir:     "homeDir",
			expectedDir: "homeDir/.local/share/juno",
		},
		"neither home dir nor data dir exist": {
			dataDir:     "",
			homeDir:     "",
			expectedDir: "",
		},
	}

	for description, test := range tests {
		t.Run(description, func(t *testing.T) {
			t.Setenv("XDG_DATA_HOME", test.dataDir)
			t.Setenv("HOME", test.homeDir)
			got, err := utils.DefaultDataDir()
			if test.expectedDir == "" {
				require.Error(t, err)
			}
			assert.Equal(t, test.expectedDir, got)
		})
	}
}
