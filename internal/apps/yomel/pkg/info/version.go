package info

import (
	_ "embed"
	"errors"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/parser"
)

type YomelInfo struct {
	Yomel struct {
		Version     string `toml:"version"`
		Name        string `toml:"name"`
		Description string `toml:"description"`
	} `toml:"yomel"`
}

func GetVersion(ctrl parser.Control) (*string, error) {
	if !ctrl.IsVersion {
		return nil, nil
	}
	var info YomelInfo
	if _, err := toml.Decode(YomelInfoRaw, &info); err != nil {
		return nil, fmt.Errorf("failed to parse yomel.toml: %v\n", err)
	}
	version := info.Yomel.Version
	if version == "" {
		return nil, errors.New("unknown")
	}
	return &version, nil
}
