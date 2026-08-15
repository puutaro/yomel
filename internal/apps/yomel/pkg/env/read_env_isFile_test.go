package env

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIsFile(t *testing.T) {
	// 1. パスが空文字の場合のテスト
	if isFile("") {
		t.Errorf("expected false for empty path, got true")
	}

	// 2. 存在しないパスの場合のテスト
	nonExistentPath := filepath.Join(t.TempDir(), "non_existent_file.txt")
	if isFile(nonExistentPath) {
		t.Errorf("expected false for non-existent path: %s, got true", nonExistentPath)
	}

	// 3. 存在するファイルの場合のテスト
	tmpFile, err := os.CreateTemp(t.TempDir(), "test_file_*.txt")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	tmpFile.Close()

	if !isFile(tmpFile.Name()) {
		t.Errorf("expected true for existing file: %s, got false", tmpFile.Name())
	}
}
