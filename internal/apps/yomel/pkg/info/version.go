package info

import (
	"errors"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtablecounter"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtabledtos"
)

func GetVersion(argTableDtos []argtabledtos.ArgTableDto) (*string, error) {
	stageNo := 0
	for _, argTableDto := range argTableDtos {
		stageNo += argtablecounter.IncrementStageNo(argTableDto.IsStage)
		if stageNo > 0 {
			return nil, nil
		}
		if argTableDto.IsVersion {
			return execGetVersion()
		}
	}
	return nil, nil
}
func execGetVersion() (*string, error) {
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
