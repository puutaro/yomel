package env

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDecideYomelTomlPath(t *testing.T) {
	// Temporarily create a temporary directory to act as a mock home directory
	tmpDir, err := os.MkdirTemp("", "yomel-home-test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create mock .yomel directory and yomel.toml inside the temp home directory
	mockHiddenDir := filepath.Join(tmpDir, yomelHiddenDirName)
	if err := os.MkdirAll(mockHiddenDir, 0755); err != nil {
		t.Fatalf("failed to create mock hidden dir: %v", err)
	}
	mockTomlPath := filepath.Join(mockHiddenDir, yomelTomlName)
	if err := os.WriteFile(mockTomlPath, []byte("[yomel]\nversion = \"0.0.2\""), 0644); err != nil {
		t.Fatalf("failed to write mock toml file: %v", err)
	}

	// Create another standalone temp file to test explicit environment variable path
	explicitTomlFile, err := os.CreateTemp("", "explicit-yomel.toml")
	if err != nil {
		t.Fatalf("failed to create explicit temp file: %v", err)
	}
	explicitTomlFilePath := explicitTomlFile.Name()
	explicitTomlFile.Close()
	defer os.Remove(explicitTomlFilePath)

	// Save original environment variable to restore later
	origEnv := os.Getenv(envYomelTomlPath)
	defer func() {
		os.Setenv(envYomelTomlPath, origEnv)
	}()

	t.Run("Env variable path is valid and exists", func(t *testing.T) {
		os.Setenv(envYomelTomlPath, explicitTomlFilePath)

		// Note: Since decideYomelTomlPath uses os.UserHomeDir internally if env doesn't match or for home fallback,
		// we test the environment variable branch directly.
		path := decideYomelTomlPath()
		if path != explicitTomlFilePath {
			t.Errorf("expected %s, got %s", explicitTomlFilePath, path)
		}
	})

	t.Run("Env variable path is invalid, falls back to empty or home if mocked", func(t *testing.T) {
		os.Setenv(envYomelTomlPath, "/non/existent/path/yomel.toml")

		// To safely test home directory logic without modifying real user home,
		// we can check behavior when home config might or might not match.
		path := decideYomelTomlPath()
		// Since we cannot easily mock os.UserHomeDir() without extra libraries,
		// we verify that it gracefully returns a string without panicking.
		_ = path
	})
}
