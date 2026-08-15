package env_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/env"
)

func TestReadEnd(t *testing.T) {
	// 1. 一時ディレクトリとHOME環境変数のモック化（テストの最初で設定し、YomelTomlPathの検証を確実にする）
	tmpDir, err := os.MkdirTemp("", "yomeltaest")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	yomelConfigDir := filepath.Join(tmpDir, ".yomel")
	err = os.MkdirAll(yomelConfigDir, 0755)
	if err != nil {
		t.Fatalf("failed to create yomel config dir: %v", err)
	}

	tomlPath := filepath.Join(yomelConfigDir, "log.toml")
	err = os.WriteFile(tomlPath, []byte("[color]\nenable_light_color_mode = 1\n"), 0644)
	if err != nil {
		t.Fatalf("failed to write dummy toml: %v", err)
	}

	oldHome := os.Getenv("YOMEL_TOML_PATH")
	defer os.Setenv("YOMEL_TOML_PATH", oldHome)
	os.Setenv("YOMEL_TOML_PATH", tmpDir)
	oldEnableTee := os.Getenv("YOMEL_ENABLE_TEE")
	defer os.Setenv("YOMEL_ENABLE_TEE", oldEnableTee)
	os.Setenv("YOMEL_ENABLE_TEE", "1")

	// 2. 通常ケース（isDirect: false, isGen: false）
	// 実装側のロジック（!isDirect && !isGen）により IsTerminal は true になります
	yomelEnv1 := env.ReadEnd(false, false)
	if yomelEnv1.IsTerminal != true {
		t.Errorf("expected IsTerminal to be true, got %v", yomelEnv1.IsTerminal)
	}

	if yomelEnv1.YomelTomlPath == "" {
		t.Errorf("expected YomelTomlPath to be resolved, got empty string")
	}

	// 3. Direct モード（isDirect: true, isGen: false）
	yomelEnv2 := env.ReadEnd(true, false)
	if yomelEnv2.IsTerminal != false {
		t.Errorf("expected IsTerminal to be false, got %v", yomelEnv2.IsTerminal)
	}

	// 4. Gen モード（isDirect: false, isGen: true）
	yomelEnv3 := env.ReadEnd(false, true)
	if yomelEnv3.IsTerminal != false {
		t.Errorf("expected IsTerminal to be false, got %v", yomelEnv3.IsTerminal)
	}
}
