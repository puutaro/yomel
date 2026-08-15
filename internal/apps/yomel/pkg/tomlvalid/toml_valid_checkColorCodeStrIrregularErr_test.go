package tomlvalid_test

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/toml"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/tomlvalid"
)

func TestCheckColorCodeStrIrregularErrForTitleColor(t *testing.T) {
	// 正常系: すべて有効なカラーコード（または空文字列）の場合
	validConfig := toml.LogConfig{
		Color: toml.ColorConfig{
			TitleColor:        "#123456",
			TitleBgColor:      "#654321",
			TitleCommentColor: "#abcdef",
		},
	}

	err := tomlvalid.TomlValidate(validConfig, "yomel.toml")
	if err != nil {
		t.Fatalf("expected no error for valid title color config, got %v", err)
	}

	// 異常系1: TitleColor が不正なカラーコードの場合
	invalidTitleColorConfig := toml.LogConfig{
		Color: toml.ColorConfig{
			TitleColor:        "invalid-color",
			TitleBgColor:      "#654321",
			TitleCommentColor: "#abcdef",
		},
	}

	err = tomlvalid.TomlValidate(invalidTitleColorConfig, "yomel.toml")
	if err == nil {
		t.Fatal("expected error for invalid TitleColor, got nil")
	}

	// 異常系2: TitleBgColor が不正なカラーコードの場合
	invalidTitleBgColorConfig := toml.LogConfig{
		Color: toml.ColorConfig{
			TitleColor:        "#123456",
			TitleBgColor:      "invalid-color",
			TitleCommentColor: "#abcdef",
		},
	}

	err = tomlvalid.TomlValidate(invalidTitleBgColorConfig, "yomel.toml")
	if err == nil {
		t.Fatal("expected error for invalid TitleBgColor, got nil")
	}

	// 異常系3: TitleCommentColor が不正なカラーコードの場合
	invalidTitleCommentColorConfig := toml.LogConfig{
		Color: toml.ColorConfig{
			TitleColor:        "#123456",
			TitleBgColor:      "#654321",
			TitleCommentColor: "invalid-color",
		},
	}

	err = tomlvalid.TomlValidate(invalidTitleCommentColorConfig, "yomel.toml")
	if err == nil {
		t.Fatal("expected error for invalid TitleCommentColor, got nil")
	}
}
