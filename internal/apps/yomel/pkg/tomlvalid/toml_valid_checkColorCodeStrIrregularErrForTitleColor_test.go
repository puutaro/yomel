package tomlvalid_test

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/toml"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/tomlvalid"
)

func TestTomlValid_CheckColorCodeStrIrregularErrForTitleColor(t *testing.T) {
	tests := []struct {
		name      string
		logConfig toml.LogConfig
		wantErr   bool
	}{
		{
			name: "Valid title color hex",
			logConfig: toml.LogConfig{
				Color: toml.ColorConfig{
					TitleColor:        "#ff0000",
					TitleBgColor:      "#00ff00",
					TitleCommentColor: "#0000ff",
				},
			},
			wantErr: false,
		},
		{
			name: "Valid empty title colors",
			logConfig: toml.LogConfig{
				Color: toml.ColorConfig{
					TitleColor:        "",
					TitleBgColor:      "",
					TitleCommentColor: "",
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid title color hex length",
			logConfig: toml.LogConfig{
				Color: toml.ColorConfig{
					TitleColor: "#ff00",
				},
			},
			wantErr: true,
		},
		{
			name: "Invalid title background color hex",
			logConfig: toml.LogConfig{
				Color: toml.ColorConfig{
					TitleBgColor: "invalid-color",
				},
			},
			wantErr: true,
		},
		{
			name: "Invalid title comment color hex length",
			logConfig: toml.LogConfig{
				Color: toml.ColorConfig{
					TitleCommentColor: "#123",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tomlvalid.TomlValidate(tt.logConfig, "dummy.toml")
			if (err != nil) != tt.wantErr {
				t.Errorf("TomlValidate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
