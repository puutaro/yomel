package args_test

import (
	"os"
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/args"
	"github.com/stretchr/testify/assert"
)

func ptr(s string) *string { return &s }

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
				{StageNo: 0, Str: ptr("grep log")},
				// --err-log-filter "grep err_log"
				{StageNo: 0, IsErrLogFilter: true},
				{StageNo: 0, Str: ptr("grep err_log")},

				// stage "test1"
				{StageNo: 1, IsStage: true},
				{StageNo: 1, Str: ptr("test1")},
				// --log-filter "grep \"log aws1\""
				{StageNo: 1, IsLogFilter: true},
				{StageNo: 1, Str: ptr("grep \"log aws1\"")},
				// -cmd "aws"
				{StageNo: 1, IsCmd: true},
				{StageNo: 1, Str: ptr("aws")},
				// --opt "a"
				{StageNo: 1, IsOpt: true},
				{StageNo: 1, Str: ptr("a")},
				// --val --s "aaa"
				{StageNo: 1, IsValue: true},
				{StageNo: 1, QuoteTypeSignal: args.SingleQuote},
				{StageNo: 1, Str: ptr("aaa")},
				// --opt "b"
				{StageNo: 1, IsOpt: true},
				{StageNo: 1, Str: ptr("b")},
				// --val --s "bbb"
				{StageNo: 1, IsValue: true},
				{StageNo: 1, QuoteTypeSignal: args.SingleQuote},
				{StageNo: 1, Str: ptr("bbb")},
				// --lop "c"
				{StageNo: 1, IsLopt: true},
				{StageNo: 1, Str: ptr("c")},
				// --arg --s "awsawsaws1"
				{StageNo: 1, IsArg: true},
				{StageNo: 1, QuoteTypeSignal: args.SingleQuote},
				{StageNo: 1, Str: ptr("awsawsaws1")},
				// --arg --s "awsawsaws2"
				{StageNo: 1, IsArg: true},
				{StageNo: 1, QuoteTypeSignal: args.SingleQuote},
				{StageNo: 1, Str: ptr("awsawsaws2")},
				// -svc "a3api"
				{StageNo: 1, IsSvc: true},
				{StageNo: 1, Str: ptr("a3api")},
				// --opt "e"
				{StageNo: 1, IsOpt: true},
				{StageNo: 1, Str: ptr("e")},
				// --val --s "eeee"
				{StageNo: 1, IsValue: true},
				{StageNo: 1, QuoteTypeSignal: args.SingleQuote},
				{StageNo: 1, Str: ptr("eeee")},
				// --opt "f"
				{StageNo: 1, IsOpt: true},
				{StageNo: 1, Str: ptr("f")},
				// --val --s "ffff"
				{StageNo: 1, IsValue: true},
				{StageNo: 1, QuoteTypeSignal: args.SingleQuote},
				{StageNo: 1, Str: ptr("ffff")},
				// --arg --s "svcsvcsvc1"
				{StageNo: 1, IsArg: true},
				{StageNo: 1, QuoteTypeSignal: args.SingleQuote},
				{StageNo: 1, Str: ptr("svcsvcsvc1")},
				// --arg --s "svcsvcsvc2"
				{StageNo: 1, IsArg: true},
				{StageNo: 1, QuoteTypeSignal: args.SingleQuote},
				{StageNo: 1, Str: ptr("svcsvcsvc2")},
				// -act "list-objects"
				{StageNo: 1, IsAct: true},
				{StageNo: 1, Str: ptr("list-objects")},
				// --lop "s"
				{StageNo: 1, IsLopt: true},
				{StageNo: 1, Str: ptr("s")},
				// --val --s "sss"
				{StageNo: 1, IsValue: true},
				{StageNo: 1, QuoteTypeSignal: args.SingleQuote},
				{StageNo: 1, Str: ptr("sss")},
				// --lop "t"
				{StageNo: 1, IsLopt: true},
				{StageNo: 1, Str: ptr("t")},
				// --val --n "ttt"
				{StageNo: 1, IsValue: true},
				{StageNo: 1, QuoteTypeSignal: args.NoQuote},
				{StageNo: 1, Str: ptr("ttt")},
				// --arg "agagagaga1"
				{StageNo: 1, IsArg: true},
				{StageNo: 1, Str: ptr("agagagaga1")},
				// --arg "agagagaga2"
				{StageNo: 1, IsArg: true},
				{StageNo: 1, Str: ptr("agagagaga2")},

				// stage "sed"
				{StageNo: 2, IsStage: true},
				{StageNo: 2, Str: ptr("sed")},
				// -cmd "sed"
				{StageNo: 2, IsCmd: true},
				{StageNo: 2, Str: ptr("sed")},
				// --opt "e"
				{StageNo: 2, IsOpt: true},
				{StageNo: 2, Str: ptr("e")},
				// --arg "/aa/bb/"
				{StageNo: 2, IsArg: true},
				{StageNo: 2, Str: ptr("/aa/bb/")},
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
