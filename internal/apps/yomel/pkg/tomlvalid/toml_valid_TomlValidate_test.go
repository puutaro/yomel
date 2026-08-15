package tomlvalid_test

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/toml"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/tomlvalid"
)

func TestTomlValidate(t *testing.T) {
	tests := []struct {
		name      string
		logConfig toml.LogConfig
		tomlPath  string
		wantErr   bool
	}{
		{
			name: "Valid colors",
			logConfig: toml.LogConfig{
				Color: toml.ColorConfig{
					Color:             "red",
					BgColor:           "#000000",
					CommentColor:      "blue",
					TitleColor:        "green",
					TitleBgColor:      "#ffffff",
					TitleCommentColor: "yellow",
				},
			},
			tomlPath: "yomel.toml",
			wantErr:  false,
		},
		{
			name: "Valid colors with order operator",
			logConfig: toml.LogConfig{
				Color: toml.ColorConfig{
					Color:        "red > blue",
					BgColor:      "#123456",
					CommentColor: "green",
				},
			},
			tomlPath: "yomel.toml",
			wantErr:  false,
		},
		{
			name: "Invalid color string in Color",
			logConfig: toml.LogConfig{
				Color: toml.ColorConfig{
					Color: "invalid-color-string",
				},
			},
			tomlPath: "yomel.toml",
			wantErr:  true,
		},
		{
			name: "Invalid hex length in BgColor",
			logConfig: toml.LogConfig{
				Color: toml.ColorConfig{
					BgColor: "#123", // invalid length
				},
			},
			tomlPath: "yomel.toml",
			wantErr:  true,
		},
		{
			name: "Invalid color string in TitleColor",
			logConfig: toml.LogConfig{
				Color: toml.ColorConfig{
					TitleColor: "not-a-color",
				},
			},
			tomlPath: "yomel.toml",
			wantErr:  true,
		},
		{
			name: "Invalid hex in TitleBgColor",
			logConfig: toml.LogConfig{
				Color: toml.ColorConfig{
					TitleBgColor: "#xyz123", // non-hex characters
				},
			},
			tomlPath: "yomel.toml",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tomlvalid.TomlValidate(tt.logConfig, tt.tomlPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("TomlValidate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
