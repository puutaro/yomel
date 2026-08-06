package argtables

import (
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtabledtos"
)

type ArgTable struct {
	No              int
	IsVersion       bool
	IsHelp          bool
	IsGen           bool
	IsDirect        bool
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
	QuoteTypeSignal argtabledtos.QuoteType
	UnkownOption    string
	Str             *string
}

func GenArgTable(argTableDtos []argtabledtos.ArgTableDto) []ArgTable {
	var argTables []ArgTable
	for _, argTableDto := range argTableDtos {
		argTable := ArgTable{
			No:              argTableDto.No,
			IsVersion:       argTableDto.IsVersion,
			IsHelp:          argTableDto.IsHelp,
			IsGen:           argTableDto.IsGen,
			IsDirect:        argTableDto.IsDirect,
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
