package model

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
	"github.com/puutaro/yomel/internal/pkg/testutil"
)

func TestParseOptionsAndArgs(t *testing.T) {
	// Prepare test arg tables simulating command line inputs for a stage
	argTables := []argtables.ArgTable{
		{
			No:      1,
			StageNo: 1,
			IsStage: true,
			Str:     testutil.Ptr("FirstStage"),
		},
		{
			No:      2,
			StageNo: 1,
			IsCmd:   true,
			Str:     testutil.Ptr("echo"),
		},
		{
			No:      3,
			StageNo: 1,
			IsOpt:   true,
			Comment: "OptA",
		},
		{
			No:      4,
			StageNo: 1,
			Str:     testutil.Ptr("-n"),
		},
		{
			No:      5,
			StageNo: 1,
			IsValue: true,
			Comment: "ValueA",
		},
		{
			No:              6,
			StageNo:         1,
			QuoteTypeSignal: argtables.DoubleQuote,
			Str:             testutil.Ptr("hello"),
		},
		{
			No:      7,
			StageNo: 1,
			IsArg:   true,
			Comment: "ArgA",
		},
		{
			No:              8,
			StageNo:         1,
			QuoteTypeSignal: argtables.DoubleQuote,
			Str:             testutil.Ptr("world"),
		},
	}

	// Test parseOptions for short options (-o)
	t.Run("parseOptions with short option and value", func(t *testing.T) {
		var opts []OptParam
		parseOptions(
			0,
			argTables,
			func(t argtables.ArgTable) bool { return t.IsCmd },
			func(t argtables.ArgTable) bool { return t.IsOpt },
			func(p OptParam) {
				opts = append(opts, p)
			},
		)

		if len(opts) != 1 {
			iteratorErr := len(opts)
			t.Fatalf("expected 1 option parsed, got %d", iteratorErr)
		}

		opt := opts[0]
		if opt.OptStr != "-n" {
			t.Errorf("expected OptStr to be '-n', got '%s'", opt.OptStr)
		}
		if opt.Comment != "OptA" {
			t.Errorf("expected Comment to be 'OptA', got '%s'", opt.Comment)
		}
		if opt.Param.Str == nil || *opt.Param.Str != "hello" {
			t.Errorf("expected Param.Str to be 'hello', got %v", opt.Param.Str)
		}
		if opt.Param.Comment != "ValueA" {
			t.Errorf("expected Param.Comment to be 'ValueA', got '%s'", opt.Param.Comment)
		}
	})

	// Test parseArg for positional arguments (-a)
	t.Run("parseArg with positional argument", func(t *testing.T) {
		var args []ArgParam
		parseArg(
			0,
			argTables,
			func(t argtables.ArgTable) bool { return t.IsCmd },
			func(ind int, p ParamType) {
				args = append(args, ArgParam{
					Index: ind,
					Param: p,
				})
			},
		)

		if len(args) != 1 {
			argLenErr := len(args)
			t.Fatalf("expected 1 argument parsed, got %d", argLenErr)
		}

		arg := args[0]
		if arg.Param.Str == nil || *arg.Param.Str != "world" {
			t.Errorf("expected Param.Str to be 'world', got %v", arg.Param.Str)
		}
		if arg.Param.Comment != "ArgA" {
			t.Errorf("expected Param.Comment to be 'ArgA', got '%s'", arg.Param.Comment)
		}
	})
}
