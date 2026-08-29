package model

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
	"github.com/puutaro/yomel/internal/pkg/testutil"
)

// TestParseArg tests the parseArg function with various quote types and structures.
func TestParseArg(t *testing.T) {
	tests := []struct {
		name           string
		nextStartIndex int
		argTables      []argtables.ArgTable
		isTargetMain   func(t argtables.ArgTable) bool
		expectedArgs   []ArgParam
		expectedIndex  int
	}{
		{
			name:           "Parse single argument with double quote",
			nextStartIndex: 0,
			argTables: []argtables.ArgTable{
				{IsCmd: true},
				{Str: testutil.Ptr("echo")},
				{IsArg: true, Comment: "Arg1"},
				{Str: testutil.Ptr("hello")},
			},
			isTargetMain: func(t argtables.ArgTable) bool { return t.IsCmd },
			expectedArgs: []ArgParam{
				{
					Index: 3,
					Param: ParamType{
						Str:       testutil.Ptr("hello"),
						QuoteType: argtables.DoubleQuote,
						Comment:   "Arg1",
					},
				},
			},
			expectedIndex: 3,
		},
		{
			name:           "Parse multiple arguments with different quote types",
			nextStartIndex: 0,
			argTables: []argtables.ArgTable{
				{IsCmd: true},
				{Str: testutil.Ptr("command")},
				{IsArg: true, Comment: "Arg1"},
				{QuoteTypeSignal: argtables.SingleQuote},
				{Str: testutil.Ptr("single-val")},
				{IsArg: true, Comment: "Arg2"},
				{QuoteTypeSignal: argtables.NoQuote},
				{Str: testutil.Ptr("no-quote-val")},
			},
			isTargetMain: func(t argtables.ArgTable) bool { return t.IsCmd },
			expectedArgs: []ArgParam{
				{
					Index: 4,
					Param: ParamType{
						Str:       testutil.Ptr("single-val"),
						QuoteType: argtables.SingleQuote,
						Comment:   "Arg1",
					},
				},
				{
					Index: 7,
					Param: ParamType{
						Str:       testutil.Ptr("no-quote-val"),
						QuoteType: argtables.NoQuote,
						Comment:   "Arg2",
					},
				},
			},
			expectedIndex: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var actualArgs []ArgParam
			appendFn := func(ind int, p ParamType) {
				actualArgs = append(actualArgs, ArgParam{
					Index: ind,
					Param: p,
				})
			}

			nextIndex := parseArg(tt.nextStartIndex, tt.argTables, tt.isTargetMain, appendFn)

			if nextIndex != tt.expectedIndex {
				t.Errorf("parseArg() index = %d, expected %d", nextIndex, tt.expectedIndex)
			}

			if len(actualArgs) != len(tt.expectedArgs) {
				t.Fatalf("parseArg() got %d args, expected %d", len(actualArgs), len(tt.expectedArgs))
			}

			for i, actual := range actualArgs {
				expected := tt.expectedArgs[i]
				if actual.Index != expected.Index {
					t.Errorf("arg[%d] index = %d, expected %d", i, actual.Index, expected.Index)
				}
				if (actual.Param.Str == nil) != (expected.Param.Str == nil) ||
					(actual.Param.Str != nil && *actual.Param.Str != *expected.Param.Str) {
					strActual := "<nil>"
					if actual.Param.Str != nil {
						strActual = *actual.Param.Str
					}
					strExpected := "<nil>"
					if expected.Param.Str != nil {
						strExpected = *expected.Param.Str
					}
					t.Errorf("arg[%d] Str = %s, expected %s", i, strActual, strExpected)
				}
				if actual.Param.QuoteType != expected.Param.QuoteType {
					t.Errorf("arg[%d] QuoteType = %v, expected %v", i, actual.Param.QuoteType, expected.Param.QuoteType)
				}
				if actual.Param.Comment != expected.Param.Comment {
					t.Errorf("arg[%d] Comment = %s, expected %s", i, actual.Param.Comment, expected.Param.Comment)
				}
			}
		})
	}
}
