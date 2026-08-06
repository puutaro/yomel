package argtabledtos

import (
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
)

type ArgTableDto struct {
	No              int
	IsVersion       bool
	IsHelp          bool
	IsGen           bool
	IsDirect        bool
	IsNoLiveStdout  bool
	IsNoLiveStderr  bool
	IsTitle         bool
	IsLogFilter     bool
	IsErrLogFilter  bool
	StageNo         int
	IsStage         bool
	IsLog           bool
	IsNoLog         bool
	IsCmd           bool
	IsSvc           bool
	IsAct           bool
	IsOpt           bool
	IsLopt          bool
	IsValue         bool
	IsArg           bool
	QuoteTypeSignal argtables.QuoteType
	UnkownOption    string
	Str             *string
}

func GenArgTableDto(argTableDtos []argtables.ArgTableDto) []ArgTableDto {
	var argTables []ArgTableDto
	for _, argTableDto := range argTableDtos {
		argTable := ArgTableDto{
			No:              argTableDto.No,
			IsVersion:       argTableDto.IsVersion,
			IsHelp:          argTableDto.IsHelp,
			IsGen:           argTableDto.IsGen,
			IsDirect:        argTableDto.IsDirect,
			IsNoLiveStdout:  argTableDto.IsNoLiveStdout,
			IsNoLiveStderr:  argTableDto.IsNoLiveStderr,
			IsTitle:         argTableDto.IsTitle,
			IsLogFilter:     argTableDto.IsLogFilter,
			IsErrLogFilter:  argTableDto.IsErrLogFilter,
			StageNo:         argTableDto.StageNo,
			IsStage:         argTableDto.IsStage,
			IsLog:           argTableDto.IsLog,
			IsNoLog:         argTableDto.IsNoLog,
			IsCmd:           argTableDto.IsCmd,
			IsSvc:           argTableDto.IsSvc,
			IsAct:           argTableDto.IsAct,
			IsOpt:           isSetStr(argTableDto.OptStr),
			IsLopt:          isSetStr(argTableDto.LoptStr),
			IsValue:         isSetStr(argTableDto.ValueStr),
			IsArg:           isSetStr(argTableDto.ArgStr),
			QuoteTypeSignal: argTableDto.QuoteTypeSignal,
			UnkownOption:    argTableDto.UnknownOption,
			Str:             argTableDto.Str,
		}

		argTables = append(argTables, argTable)
	}
	return argTables
}

func isSetStr(ptrStr *string) bool {
	return ptrStr != nil
}
