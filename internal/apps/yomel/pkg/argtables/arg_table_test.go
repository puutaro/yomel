package argtables_test

import (
	"os"
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/arglist"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
	"github.com/puutaro/yomel/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func Test_GenArgTable(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []argtables.ArgTable
	}{
		{
			name: "make argTable from args",
			input: []string{
				"yomel",
				"--log",
				"--no-log",
				"--version",
				"--help",
				"--log-filter", "grep log",
				"--err-log-filter", "grep err_log",
				"stage", "test1",
				"--log-filter", "grep \"log aws1\"",
				"-cmd", "aws",
				"--opt", "a",
				"--val", "--s", "aaa",
				"--opt", "b",
				"--val", "--n", "bbb",
				"--lop", "c",
				"--arg", "--s", "awsawsaws1",
				"--arg", "--n", "awsawsaws2",
				"--single",
				"--no-quote",
				"-svc", "a3api",
				"--opt", "e",
				"--val", "--s", "eeee",
				"--opt", "f",
				"--val", "--s", "ffff",
				"--arg", "--s", "svcsvcsvc1",
				"--arg", "--s", "svcsvcsvc2",
				"-act", "list-objects",
				"--lop", "s",
				"--val", "--s", "sss",
				"--lop", "t",
				"--val", "--n", "ttt",
				"--arg", "agagagaga1",
				"--arg", "agagagaga2",
				"stage", "sed",
				"-cmd", "sed",
				"--opt", "e",
				"--arg", "/aa/bb/",
			},
			want: []argtables.ArgTable{
				// parse --log option
				{StageNo: 0, IsLog: true},
				// parse --no-log option
				{StageNo: 0, IsNoLog: true},
				// parse --version option
				{StageNo: 0, IsVersion: true},
				// parse --help option
				{StageNo: 0, IsHelp: true},
				// parse --log-filter option with pattern
				{StageNo: 0, IsLogFilter: true},
				{StageNo: 0, Str: testutil.Ptr("grep log")},
				// parse --err-log-filter option with pattern
				{StageNo: 0, IsErrLogFilter: true},
				{StageNo: 0, Str: testutil.Ptr("grep err_log")},

				// parse stage definition for test1
				{StageNo: 1, IsStage: true},
				{StageNo: 1, Str: testutil.Ptr("test1")},
				// parse stage level log-filter
				{StageNo: 1, IsLogFilter: true},
				{StageNo: 1, Str: testutil.Ptr("grep \"log aws1\"")},
				// parse command option
				{StageNo: 1, IsCmd: true},
				{StageNo: 1, Str: testutil.Ptr("aws")},
				// parse option a
				{StageNo: 1, IsOpt: true},
				{StageNo: 1, Str: testutil.Ptr("a")},
				// parse value with single quote signal
				{StageNo: 1, IsValue: true},
				{StageNo: 1, QuoteTypeSignal: argtables.SingleQuote},
				{StageNo: 1, Str: testutil.Ptr("aaa")},
				// parse option b
				{StageNo: 1, IsOpt: true},
				{StageNo: 1, Str: testutil.Ptr("b")},
				// parse value with no quote signal
				{StageNo: 1, IsValue: true},
				{StageNo: 1, QuoteTypeSignal: argtables.NoQuote},
				{StageNo: 1, Str: testutil.Ptr("bbb")},
				// parse long option c
				{StageNo: 1, IsLopt: true},
				{StageNo: 1, Str: testutil.Ptr("c")},
				// parse argument with single quote
				{StageNo: 1, IsArg: true},
				{StageNo: 1, QuoteTypeSignal: argtables.SingleQuote},
				{StageNo: 1, Str: testutil.Ptr("awsawsaws1")},
				// parse argument with no quote
				{StageNo: 1, IsArg: true},
				{StageNo: 1, QuoteTypeSignal: argtables.NoQuote},
				{StageNo: 1, Str: testutil.Ptr("awsawsaws2")},
				// parse single quote modifier
				{StageNo: 1, QuoteTypeSignal: argtables.SingleQuote},
				// parse no-quote modifier
				{StageNo: 1, QuoteTypeSignal: argtables.NoQuote},
				// parse service option
				{StageNo: 1, IsSvc: true},
				{StageNo: 1, Str: testutil.Ptr("a3api")},
				// parse option e
				{StageNo: 1, IsOpt: true},
				{StageNo: 1, Str: testutil.Ptr("e")},
				// parse value for option e
				{StageNo: 1, IsValue: true},
				{StageNo: 1, QuoteTypeSignal: argtables.SingleQuote},
				{StageNo: 1, Str: testutil.Ptr("eeee")},
				// parse option f
				{StageNo: 1, IsOpt: true},
				{StageNo: 1, Str: testutil.Ptr("f")},
				// parse value for option f
				{StageNo: 1, IsValue: true},
				{StageNo: 1, QuoteTypeSignal: argtables.SingleQuote},
				{StageNo: 1, Str: testutil.Ptr("ffff")},
				// parse argument svcsvcsvc1
				{StageNo: 1, IsArg: true},
				{StageNo: 1, QuoteTypeSignal: argtables.SingleQuote},
				{StageNo: 1, Str: testutil.Ptr("svcsvcsvc1")},
				// parse argument svcsvcsvc2
				{StageNo: 1, IsArg: true},
				{StageNo: 1, QuoteTypeSignal: argtables.SingleQuote},
				{StageNo: 1, Str: testutil.Ptr("svcsvcsvc2")},
				// parse action option
				{StageNo: 1, IsAct: true},
				{StageNo: 1, Str: testutil.Ptr("list-objects")},
				// parse long option s
				{StageNo: 1, IsLopt: true},
				{StageNo: 1, Str: testutil.Ptr("s")},
				// parse value for long option s
				{StageNo: 1, IsValue: true},
				{StageNo: 1, QuoteTypeSignal: argtables.SingleQuote},
				{StageNo: 1, Str: testutil.Ptr("sss")},
				// parse long option t
				{StageNo: 1, IsLopt: true},
				{StageNo: 1, Str: testutil.Ptr("t")},
				// parse value for long option t
				{StageNo: 1, IsValue: true},
				{StageNo: 1, QuoteTypeSignal: argtables.NoQuote},
				{StageNo: 1, Str: testutil.Ptr("ttt")},
				// parse raw argument 1
				{StageNo: 1, IsArg: true},
				{StageNo: 1, Str: testutil.Ptr("agagagaga1")},
				// parse raw argument 2
				{StageNo: 1, IsArg: true},
				{StageNo: 1, Str: testutil.Ptr("agagagaga2")},

				// parse stage definition for sed
				{StageNo: 2, IsStage: true},
				{StageNo: 2, Str: testutil.Ptr("sed")},
				// parse command sed
				{StageNo: 2, IsCmd: true},
				{StageNo: 2, Str: testutil.Ptr("sed")},
				// parse option e for sed
				{StageNo: 2, IsOpt: true},
				{StageNo: 2, Str: testutil.Ptr("e")},
				// parse argument expression for sed
				{StageNo: 2, IsArg: true},
				{StageNo: 2, Str: testutil.Ptr("/aa/bb/")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldArgs := os.Args
			defer func() { os.Args = oldArgs }()
			os.Args = tt.input

			inputArgs := arglist.Gen()
			got := argtables.GenArgTable(inputArgs)
			for i := range tt.want {
				tt.want[i].No = i + 1
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
