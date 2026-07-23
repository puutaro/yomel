package args_test

import (
	"os"
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/args"
	"github.com/puutaro/yomel/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func Test_GenArgTable(t *testing.T) {
	tests := []struct {
		name  string   // テストケースの名前
		input []string // 入力する引数
		want  []args.ArgTable
	}{
		{
			name: "make argTable from args",
			input: []string{
				"yomoel",
				"--log",
				"--log-filter", "grep log",
				"--err-log-filter", "grep err_log",
				"stage", "test1",
				"--log-filter", "grep \"log aws1\"",
				"-cmd", "aws",
				"--opt", "a",
				"--val", "--s", "aaa",
				"--opt", "b",
				"--val", "--s", "bbb",
				"--lop", "c",
				"--arg", "--s", "awsawsaws1",
				"--arg", "--s", "awsawsaws2",
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
			want: []args.ArgTable{
				// --log
				{StageNo: 0, IsLog: true},
				// --log-filter "grep log"
				{StageNo: 0, IsLogFilter: true},
				{StageNo: 0, Str: testutil.Ptr("grep log")},
				// --err-log-filter "grep err_log"
				{StageNo: 0, IsErrLogFilter: true},
				{StageNo: 0, Str: testutil.Ptr("grep err_log")},

				// stage "test1"
				{StageNo: 1, IsStage: true},
				{StageNo: 1, Str: testutil.Ptr("test1")},
				// --log-filter "grep \"log aws1\""
				{StageNo: 1, IsLogFilter: true},
				{StageNo: 1, Str: testutil.Ptr("grep \"log aws1\"")},
				// -cmd "aws"
				{StageNo: 1, IsCmd: true},
				{StageNo: 1, Str: testutil.Ptr("aws")},
				// --opt "a"
				{StageNo: 1, IsOpt: true},
				{StageNo: 1, Str: testutil.Ptr("a")},
				// --val --s "aaa"
				{StageNo: 1, IsValue: true},
				{StageNo: 1, QuoteTypeSignal: args.SingleQuote},
				{StageNo: 1, Str: testutil.Ptr("aaa")},
				// --opt "b"
				{StageNo: 1, IsOpt: true},
				{StageNo: 1, Str: testutil.Ptr("b")},
				// --val --s "bbb"
				{StageNo: 1, IsValue: true},
				{StageNo: 1, QuoteTypeSignal: args.SingleQuote},
				{StageNo: 1, Str: testutil.Ptr("bbb")},
				// --lop "c"
				{StageNo: 1, IsLopt: true},
				{StageNo: 1, Str: testutil.Ptr("c")},
				// --arg --s "awsawsaws1"
				{StageNo: 1, IsArg: true},
				{StageNo: 1, QuoteTypeSignal: args.SingleQuote},
				{StageNo: 1, Str: testutil.Ptr("awsawsaws1")},
				// --arg --s "awsawsaws2"
				{StageNo: 1, IsArg: true},
				{StageNo: 1, QuoteTypeSignal: args.SingleQuote},
				{StageNo: 1, Str: testutil.Ptr("awsawsaws2")},
				// -svc "a3api"
				{StageNo: 1, IsSvc: true},
				{StageNo: 1, Str: testutil.Ptr("a3api")},
				// --opt "e"
				{StageNo: 1, IsOpt: true},
				{StageNo: 1, Str: testutil.Ptr("e")},
				// --val --s "eeee"
				{StageNo: 1, IsValue: true},
				{StageNo: 1, QuoteTypeSignal: args.SingleQuote},
				{StageNo: 1, Str: testutil.Ptr("eeee")},
				// --opt "f"
				{StageNo: 1, IsOpt: true},
				{StageNo: 1, Str: testutil.Ptr("f")},
				// --val --s "ffff"
				{StageNo: 1, IsValue: true},
				{StageNo: 1, QuoteTypeSignal: args.SingleQuote},
				{StageNo: 1, Str: testutil.Ptr("ffff")},
				// --arg --s "svcsvcsvc1"
				{StageNo: 1, IsArg: true},
				{StageNo: 1, QuoteTypeSignal: args.SingleQuote},
				{StageNo: 1, Str: testutil.Ptr("svcsvcsvc1")},
				// --arg --s "svcsvcsvc2"
				{StageNo: 1, IsArg: true},
				{StageNo: 1, QuoteTypeSignal: args.SingleQuote},
				{StageNo: 1, Str: testutil.Ptr("svcsvcsvc2")},
				// -act "list-objects"
				{StageNo: 1, IsAct: true},
				{StageNo: 1, Str: testutil.Ptr("list-objects")},
				// --lop "s"
				{StageNo: 1, IsLopt: true},
				{StageNo: 1, Str: testutil.Ptr("s")},
				// --val --s "sss"
				{StageNo: 1, IsValue: true},
				{StageNo: 1, QuoteTypeSignal: args.SingleQuote},
				{StageNo: 1, Str: testutil.Ptr("sss")},
				// --lop "t"
				{StageNo: 1, IsLopt: true},
				{StageNo: 1, Str: testutil.Ptr("t")},
				// --val --n "ttt"
				{StageNo: 1, IsValue: true},
				{StageNo: 1, QuoteTypeSignal: args.NoQuote},
				{StageNo: 1, Str: testutil.Ptr("ttt")},
				// --arg "agagagaga1"
				{StageNo: 1, IsArg: true},
				{StageNo: 1, Str: testutil.Ptr("agagagaga1")},
				// --arg "agagagaga2"
				{StageNo: 1, IsArg: true},
				{StageNo: 1, Str: testutil.Ptr("agagagaga2")},

				// stage "sed"
				{StageNo: 2, IsStage: true},
				{StageNo: 2, Str: testutil.Ptr("sed")},
				// -cmd "sed"
				{StageNo: 2, IsCmd: true},
				{StageNo: 2, Str: testutil.Ptr("sed")},
				// --opt "e"
				{StageNo: 2, IsOpt: true},
				{StageNo: 2, Str: testutil.Ptr("e")},
				// --arg "/aa/bb/"
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

			got := args.GenArgTable()
			for i := range tt.want {
				tt.want[i].No = i + 1
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
