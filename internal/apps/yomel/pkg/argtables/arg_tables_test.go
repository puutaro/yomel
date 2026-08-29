package argtables_test

import (
	"os"
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/arg_list"
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
				"///",
				"test cmd title",
				"--no-live-stdout",
				"--no-live-stderr",
				"--log",
				"--no-log",
				"--version",
				"--help",
				"--direct",
				"--gen",
				"--log-filter", "grep log",
				"--err-log-filter", "grep err_log",
				"--color", "green",
				"--bg-color", "darkGreen",
				"--comment-color", "gray",
				"--title-color", "green",
				"--title-bg-color", "darkGreen",
				"--title-comment-color", "gray",
				"//", "test1",
				"--log-filter", "grep \"log aws1\"",
				"-c", "aws",
				"-oVertualOpt", "a",
				"-vVertualValue", "--s", "aaa",
				"-o", "b",
				"-v", "--n", "bbb",
				"--oVertualLoption", "c",
				"-aVertualArg", "--s", "awsawsaws1",
				"-a", "--n", "awsawsaws2",
				"--single",
				"--no-quote",
				"-a", "agagagaga1",
				"-a", "agagagaga2",

				"//", "sed",
				"-c", "sed",
				"-o", "e",
				"-a", "/aa/bb/",
			},
			want: []argtables.ArgTable{
				{StageNo: 0, IsTitle: true},
				{StageNo: 0, Str: testutil.Ptr("test cmd title")},
				{StageNo: 0, IsNoLiveStdout: true},
				{StageNo: 0, IsNoLiveStderr: true},
				// parse --log option
				{StageNo: 0, IsLog: true},
				// parse --no-log option
				{StageNo: 0, IsNoLog: true},
				// parse --version option
				{StageNo: 0, IsVersion: true},
				// parse --help option
				{StageNo: 0, IsHelp: true},
				{StageNo: 0, IsDirect: true},
				{StageNo: 0, IsGen: true},
				// parse --log-filter option with pattern
				{StageNo: 0, IsLogFilter: true},
				{StageNo: 0, Str: testutil.Ptr("grep log")},
				// parse --err-log-filter option with pattern
				{StageNo: 0, IsErrLogFilter: true},
				{StageNo: 0, Str: testutil.Ptr("grep err_log")},
				{StageNo: 0, IsColor: true},
				{StageNo: 0, Str: testutil.Ptr("green")},
				{StageNo: 0, IsBgColor: true},
				{StageNo: 0, Str: testutil.Ptr("darkGreen")},
				{StageNo: 0, IsCommentColor: true},
				{StageNo: 0, Str: testutil.Ptr("gray")},
				{StageNo: 0, IsTitleColor: true},
				{StageNo: 0, Str: testutil.Ptr("green")},
				{StageNo: 0, IsTitleBgColor: true},
				{StageNo: 0, Str: testutil.Ptr("darkGreen")},
				{StageNo: 0, IsTitleCommentColor: true},
				{StageNo: 0, Str: testutil.Ptr("gray")},

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
				{StageNo: 1, IsOpt: true, Comment: "VertualOpt"},
				{StageNo: 1, Str: testutil.Ptr("a")},
				// parse value with single quote signal
				{StageNo: 1, IsValue: true, Comment: "VertualValue"},
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
				{StageNo: 1, IsLopt: true, Comment: "VertualLoption"},
				{StageNo: 1, Str: testutil.Ptr("c")},
				// parse argument with single quote
				{StageNo: 1, IsArg: true, Comment: "VertualArg"},
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
		{
			name: "should handle unknown options and explicit short quote signals correctly",
			input: []string{
				"yomel",
				"--unknown-flag",
				"--n",
				"--s",
				"-n",
				"-s",
			},
			want: []argtables.ArgTable{
				{StageNo: 0, UnknownOption: "--unknown-flag"},
				{StageNo: 0, QuoteTypeSignal: argtables.NoQuote},
				{StageNo: 0, QuoteTypeSignal: argtables.SingleQuote},
				{StageNo: 0, UnknownOption: "-n"},
				{StageNo: 0, UnknownOption: "-s"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldArgs := os.Args
			defer func() { os.Args = oldArgs }()
			os.Args = tt.input

			inputArgs := arg_list.Gen()
			got := argtables.GenArgTable(inputArgs)
			for i := range tt.want {
				tt.want[i].No = i + 1
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
