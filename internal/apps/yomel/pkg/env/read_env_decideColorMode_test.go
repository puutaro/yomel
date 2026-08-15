package env

import (
	"os"
	"strconv"
	"testing"
)

func TestDecideColorMode(t *testing.T) {
	// Save original environment variable to restore later
	originalEnv := os.Getenv(envYomelLightColorMode)
	defer func() {
		if originalEnv == "" {
			os.Unsetenv(envYomelLightColorMode)
		} else {
			os.Setenv(envYomelLightColorMode, originalEnv)
		}
	}()

	tests := []struct {
		name     string
		envVal   string
		expected bool
	}{
		{
			name:     "When env matches Light mode value",
			envVal:   strconv.Itoa(int(Light)),
			expected: true,
		},
		{
			name:     "When env does not match Light mode value",
			envVal:   "999",
			expected: false,
		},
		{
			name:     "When env is empty",
			envVal:   "",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envVal == "" {
				os.Unsetenv(envYomelLightColorMode)
			} else {
				os.Setenv(envYomelLightColorMode, tt.envVal)
			}

			actual := decideColorMode()
			if actual != tt.expected {
				t.Errorf("decideColorMode() = %v, expected %v", actual, tt.expected)
			}
		})
	}
}
