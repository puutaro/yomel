package toml_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/toml"
)

func TestReadToml(t *testing.T) {
	// 1. Test empty path case
	t.Run("empty path returns empty config", func(t *testing.T) {
		cfg := toml.ReadToml("")
		if cfg.Color.Color != "" || cfg.Stream.EnableTee != 0 {
			t.Errorf("expected empty LogConfig, got %+v", cfg)
		}
	})

	// 2. Test valid TOML file case
	t.Run("valid toml file", func(t *testing.T) {
		tempDir := t.TempDir()
		tomlContent := []byte(`
[color]
color = "red"
bg_color = "green"
title_color = "blue"
title_bg_color = "white"
comment_color = "gray"
title_comment_color = "black"
enable_light_color_mode = 1

[stream]
log_filter_shell = "echo test"
err_log_filter_shell = "echo err"
enable_tee = 1
`)
		tmpFilePath := filepath.Join(tempDir, "yomel.toml")
		if err := os.WriteFile(tmpFilePath, tomlContent, 0644); err != nil {
			t.Fatalf("failed to write temp file: %v", err)
		}

		cfg := toml.ReadToml(tmpFilePath)

		if cfg.Color.Color != "red" {
			t.Errorf("expected color 'red', got '%s'", cfg.Color.Color)
		}
		if cfg.Color.BgColor != "green" {
			t.Errorf("expected bg_color 'green', got '%s'", cfg.Color.BgColor)
		}
		if cfg.Color.TitleColor != "blue" {
			t.Errorf("expected title_color 'blue', got '%s'", cfg.Color.TitleColor)
		}
		if cfg.Color.TitleBgColor != "white" {
			t.Errorf("expected title_bg_color 'white', got '%s'", cfg.Color.TitleBgColor)
		}
		if cfg.Color.CommentColor != "gray" {
			t.Errorf("expected comment_color 'gray', got '%s'", cfg.Color.CommentColor)
		}
		if cfg.Color.TitleCommentColor != "black" {
			t.Errorf("expected title_comment_color 'black', got '%s'", cfg.Color.TitleCommentColor)
		}
		if cfg.Color.EnableLightMode != 1 {
			t.Errorf("expected enable_light_color_mode 1, got %d", cfg.Color.EnableLightMode)
		}
		if cfg.Stream.LogFilterShell != "echo test" {
			t.Errorf("expected log_filter_shell 'echo test', got '%s'", cfg.Stream.LogFilterShell)
		}
		if cfg.Stream.ErrLogFilterShell != "echo err" {
			t.Errorf("expected err_log_filter_shell 'echo err', got '%s'", cfg.Stream.ErrLogFilterShell)
		}
		if cfg.Stream.EnableTee != 1 {
			t.Errorf("expected enable_tee 1, got %d", cfg.Stream.EnableTee)
		}
	})

	// 3. Test invalid TOML file case using subprocess
	t.Run("invalid toml file calls log.Fatalf", func(t *testing.T) {
		tempDir := t.TempDir()
		invalidTomlContent := []byte(`
[color
color = "invalid_syntax
`)
		tmpFilePath := filepath.Join(tempDir, "invalid.toml")
		if err := os.WriteFile(tmpFilePath, invalidTomlContent, 0644); err != nil {
			t.Fatalf("failed to write temp file: %v", err)
		}

		// サブプロセスとして自身を実行した際の処理
		if os.Getenv("BE_CRASHER") == "1" {
			toml.ReadToml(tmpFilePath)
			return
		}

		// 親プロセス：自分自身のテストを特定のフラグと環境変数で再実行する
		cmd := exec.Command(os.Args[0], "-test.run=TestReadToml/invalid_toml_file_calls_log.Fatalf")
		cmd.Env = append(os.Environ(), "BE_CRASHER=1")
		err := cmd.Run()

		// log.Fatalf によってプロセスが異常終了（Exit Code != 0）したことを確認する
		if exitErr, ok := err.(*exec.ExitError); ok && !exitErr.Success() {
			return
		}
		t.Fatalf("process ran with err %v, want exit status 1", err)
	})
}
