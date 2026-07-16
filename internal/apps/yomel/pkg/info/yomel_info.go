package info

import (
	_ "embed"
)

type YomelInfo struct {
	Yomel struct {
		Version     string `toml:"version"`
		Name        string `toml:"name"`
		Description string `toml:"description"`
	} `toml:"yomel"`
}
