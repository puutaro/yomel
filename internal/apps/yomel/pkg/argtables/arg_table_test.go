package argtables_test

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtabledtos"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
	"github.com/puutaro/yomel/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

// Test_GenArgTable verifies that GenArgTable correctly converts a slice of ArgTableDto into a slice of ArgTable structures.
func Test_GenArgTable(t *testing.T) {
	tests := []struct {
		name  string
		input []argtabledtos.ArgTableDto
		want  []argtables.ArgTable
	}{
		{
			name: "should convert various DTO flags, options, and values to ArgTable correctly",
			input: []argtabledtos.ArgTableDto{
				{
					No:              1,
					IsVersion:       true,
					IsHelp:          false,
					IsGen:           true,
					IsDirect:        false,
					IsLogFilter:     false,
					IsErrLogFilter:  false,
					StageNo:         0,
					IsStage:         false,
					IsLog:           true,
					IsNoLog:         false,
					IsCmd:           false,
					IsSvc:           false,
					IsAct:           false,
					OptStr:          nil,
					LoptStr:         nil,
					ValueStr:        nil,
					ArgStr:          nil,
					QuoteTypeSignal: argtabledtos.NoQuote,
					UnknownOption:   "",
					Str:             nil,
				},
				{
					No:              2,
					IsVersion:       false,
					IsHelp:          true,
					IsGen:           false,
					IsDirect:        true,
					IsLogFilter:     true,
					IsErrLogFilter:  true,
					StageNo:         1,
					IsStage:         true,
					IsLog:           false,
					IsNoLog:         true,
					IsCmd:           true,
					IsSvc:           true,
					IsAct:           true,
					OptStr:          testutil.Ptr("a"),
					LoptStr:         testutil.Ptr("region"),
					ValueStr:        testutil.Ptr("val"),
					ArgStr:          testutil.Ptr("arg"),
					QuoteTypeSignal: argtabledtos.SingleQuote,
					UnknownOption:   "--unknown",
					Str:             testutil.Ptr("test-str"),
				},
			},
			want: []argtables.ArgTable{
				{
					No:              1,
					IsVersion:       true,
					IsHelp:          false,
					IsGen:           true,
					IsDirect:        false,
					IsLogFilter:     false,
					IsErrLogFilter:  false,
					StageNo:         0,
					IsStage:         false,
					IsLog:           true,
					IsNoLog:         false,
					IsCmd:           false,
					IsSvc:           false,
					IsAct:           false,
					IsOpt:           false,
					IsLopt:          false,
					IsValue:         false,
					IsArg:           false,
					QuoteTypeSignal: argtables.QuoteType(argtabledtos.NoQuote),
					UnkownOption:    "",
					Str:             nil,
				},
				{
					No:              2,
					IsVersion:       false,
					IsHelp:          true,
					IsGen:           false,
					IsDirect:        true,
					IsLogFilter:     true,
					IsErrLogFilter:  true,
					StageNo:         1,
					IsStage:         true,
					IsLog:           false,
					IsNoLog:         true,
					IsCmd:           true,
					IsSvc:           true,
					IsAct:           true,
					IsOpt:           true,
					IsLopt:          true,
					IsValue:         true,
					IsArg:           true,
					QuoteTypeSignal: argtables.QuoteType(argtabledtos.SingleQuote),
					UnkownOption:    "--unknown",
					Str:             testutil.Ptr("test-str"),
				},
			},
		},
		{
			name:  "should return empty slice when input DTOs slice is empty",
			input: []argtabledtos.ArgTableDto{},
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := argtables.GenArgTable(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
