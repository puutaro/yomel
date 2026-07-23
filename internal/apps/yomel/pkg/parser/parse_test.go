package parser

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/args"
	"github.com/puutaro/yomel/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

// Test_Parse verifies the top-level Parse function correctly transforms argTables into Yomel structures.
func Test_Parse(t *testing.T) {
	tests := []struct {
		name  string
		input []args.ArgTable
		want  Yomel
	}{
		{
			name: "should parse control flags and multi-stage pipeline configuration successfully",
			input: []args.ArgTable{
				{No: 1, StageNo: 0, IsLog: true},
				{No: 2, StageNo: 1, IsStage: true},
				{No: 3, StageNo: 1, Str: testutil.Ptr("first-stage")},
				{No: 4, StageNo: 1, IsCmd: true},
				{No: 5, StageNo: 1, Str: testutil.Ptr("echo")},
				{No: 6, StageNo: 1, IsArg: true},
				{No: 7, StageNo: 1, QuoteTypeSignal: args.NoQuote},
				{No: 8, StageNo: 1, Str: testutil.Ptr("test-message")},
			},
			want: Yomel{
				Ctrl: Control{
					IsLog:        true,
					LogFilter:    "",
					ErrLogFilter: "",
					IsVersion:    false,
					IsHelp:       false,
				},
				Stages: []Stage{
					{
						No:           1,
						Desc:         "first-stage",
						Cmd:          "echo",
						CmdOpArgs:    []string{"test-message"},
						Svc:          "",
						SvcOpArgs:    nil,
						Act:          "",
						ActOpArgs:    nil,
						LogFilter:    "",
						ErrLogFilter: "",
					},
				},
			},
		},
		// {
		// 	name:  "should return empty stages when input argTables is empty",
		// 	input: []args.ArgTable{},
		// 	want: Yomel{
		// 		Ctrl: Control{
		// 			IsLog:        false,
		// 			LogFilter:    "",
		// 			ErrLogFilter: "",
		// 			IsVersion:    false,
		// 			IsHelp:       false,
		// 		},
		// 		Stages: []Stage{},
		// 	},
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Parse(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
