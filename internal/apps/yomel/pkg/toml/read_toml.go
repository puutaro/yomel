package toml

import (
	"log"

	"github.com/BurntSushi/toml"
)

const (
	ColorColor             = "color"
	ColorBgColor           = "bg_color"
	ColorCommentColor      = "comment_color"
	ColorTitleColor        = "title_color"
	ColorTitleBgColor      = "title_bg_color"
	ColorTitleCommentColor = "title_comment_color"
	ColorEnableLightMode   = "enable_light_color_mode"

	LogColor  = "color"
	LogStream = "stream"
)

func ReadToml(yomelTomlPath string) LogConfig {
	if yomelTomlPath == "" {
		return LogConfig{}
	}
	var logConfig LogConfig
	_, err := toml.DecodeFile(yomelTomlPath, &logConfig)
	if err != nil {
		log.Fatalf("fail to read yoml toml: %v", err)
	}
	return logConfig
}

type LogConfig struct {
	Color  ColorConfig  `toml:"color"`
	Stream StreamConfig `toml:"stream"`
}

type ColorConfig struct {
	Color             string `toml:"color"`
	BgColor           string `toml:"bg_color"`
	TitleColor        string `toml:"title_color"`
	TitleBgColor      string `toml:"title_bg_color"`
	CommentColor      string `toml:"comment_color"`
	TitleCommentColor string `toml:"title_comment_color"`
	EnableLightMode   int    `toml:"enable_light_color_mode"`
}
type StreamConfig struct {
	LogFilterShell    string `toml:"log_filter_shell"`
	ErrLogFilterShell string `toml:"err_log_filter_shell"`
	EnableTee         int    `toml:"enable_tee"`
}
