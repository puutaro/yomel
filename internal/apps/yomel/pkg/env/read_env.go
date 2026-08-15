package env

import (
	"os"
	"path/filepath"
	"strconv"
)

const (
	envYomelLightColorMode = "YOMEL_LIGHT_COLOR_MODE"
	envYomelEnableTee      = "YOMEL_ENABLE_TEE"
	envYomelTomlPath       = "YOMEL_TOML_PATH"
)

const (
	yomelHiddenDirName = ".yomel"
	yomelTomlName      = "yomel.toml"
)

type YomelEnv struct {
	IsTerminal       bool
	IsLightColorMode bool
	YomelTomlPath    string
}

type EnableTee int

const (
	TeeOn EnableTee = 1
)

type EnableLightColorMode int

const (
	Light EnableLightColorMode = 1
)

func ReadEnd(isDirect bool, isGen bool) YomelEnv {
	if isDirect || isGen {
		return YomelEnv{}
	}
	yomelTomlPath := decideYomelTomlPath()
	return YomelEnv{
		IsTerminal:       isTerminal(os.Stderr),
		IsLightColorMode: decideColorMode(),
		YomelTomlPath:    yomelTomlPath,
	}
}

func isTerminal(f *os.File) bool {
	enableTeeValue := os.Getenv(envYomelEnableTee)
	if enableTeeValue == strconv.Itoa(int(TeeOn)) {
		return true
	}
	fi, err := f.Stat()
	if err != nil {
		return false
	}
	return (fi.Mode() & os.ModeCharDevice) != 0
}

func decideColorMode() bool {
	yomelColorMode := os.Getenv(envYomelLightColorMode)
	return yomelColorMode == strconv.Itoa(int(Light))
}
func decideYomelTomlPath() string {
	envYTomlPath := os.Getenv(envYomelTomlPath)
	if isFile(envYTomlPath) {
		return envYTomlPath
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	homeYomlPath := filepath.Join(
		homeDir,
		yomelHiddenDirName,
		yomelTomlName,
	)
	if isFile(homeYomlPath) {
		return envYTomlPath
	}
	return ""
}

func isFile(path string) bool {
	if path == "" {
		return false
	}
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}
