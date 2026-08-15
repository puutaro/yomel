package info

import (
	"errors"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtablecounter"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
)

func GetVersion(argtables []argtables.ArgTable) (*string, error) {
	stageNo := 0
	for _, argTable := range argtables {
		stageNo += argtablecounter.IncrementStageNo(argTable.IsStage)
		if stageNo > 0 {
			return nil, nil
		}
		if argTable.IsVersion {
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
