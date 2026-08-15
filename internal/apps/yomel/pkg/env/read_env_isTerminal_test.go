package env

import (
	"os"
	"strconv"
	"testing"
)

func TestIsTerminal(t *testing.T) {
	// 1. 環境変数 envYomelEnableTee が TeeOn の場合テスト
	t.Run("EnableTee is TeeOn", func(t *testing.T) {
		t.Setenv(envYomelEnableTee, strconv.Itoa(int(TeeOn)))

		// 適当な一時ファイルを作成して渡す
		tmpFile, err := os.CreateTemp("", "test_terminal_*.txt")
		if err != nil {
			t.Fatalf("failed to create temp file: %v", err)
		}
		defer os.Remove(tmpFile.Name())
		defer tmpFile.Close()

		if !isTerminal(tmpFile) {
			t.Errorf("expected isTerminal to be true when EnableTee is TeeOn")
		}
	})

	// 2. f.Stat() がエラーになる場合 (閉じられたファイルなどを渡す) のテスト
	t.Run("File Stat error", func(t *testing.T) {
		t.Setenv(envYomelEnableTee, "")

		tmpFile, err := os.CreateTemp("", "test_terminal_*.txt")
		if err != nil {
			t.Fatalf("failed to create temp file: %v", err)
		}
		tmpFile.Close() // 閉じることで Stat() でエラーが発生するようにする
		os.Remove(tmpFile.Name())

		if isTerminal(tmpFile) {
			t.Errorf("expected isTerminal to be false when file is closed and Stat fails")
		}
	})

	// 3. 通常ファイル (非キャラクタデバイス) の場合のテスト
	t.Run("Regular file is not terminal", func(t *testing.T) {
		t.Setenv(envYomelEnableTee, "")

		tmpFile, err := os.CreateTemp("", "test_terminal_*.txt")
		if err != nil {
			t.Fatalf("failed to create temp file: %v", err)
		}
		defer os.Remove(tmpFile.Name())
		defer tmpFile.Close()

		// 通常のファイルは通常 ModeCharDevice が立たないため false になる
		if isTerminal(tmpFile) {
			t.Errorf("expected isTerminal to be false for a regular temp file")
		}
	})
}
