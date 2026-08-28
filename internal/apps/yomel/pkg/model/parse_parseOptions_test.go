// Package model provides parsing and data modeling for yomel arguments.
package model

import (
	"reflect"
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
)

// TestParseOptions tests the parseOptions function with various argument tables and conditions.
func TestParseOptions(t *testing.T) {
	tests := []struct {
		name              string
		nextStartIndex    int
		curStageArgTables []argtables.ArgTable
		isTargetMainArg   func(t argtables.ArgTable) bool
		isNextMainArg     func(t argtables.ArgTable) bool
		isTargetOpt       func(argtables.ArgTable) bool
		expectedOptParams []OptParam
	}{
		{
			name:              "Empty argument tables",
			nextStartIndex:    0,
			curStageArgTables: []argtables.ArgTable{},
			isTargetMainArg: func(t argtables.ArgTable) bool {
				return t.IsCmd
			},
			isNextMainArg: func(t argtables.ArgTable) bool {
				return t.IsStage
			},
			isTargetOpt: func(t argtables.ArgTable) bool {
				return t.IsOpt
			},
			expectedOptParams: nil,
		},
		{
			name:           "Parse single option without value parameter",
			nextStartIndex: 0,
			curStageArgTables: []argtables.ArgTable{
				{IsCmd: true, StageNo: 1},
				{Str: strPtr("echo"), StageNo: 1},
				{IsOpt: true, Comment: "TestOpt", StageNo: 1},
				{Str: strPtr("foo"), StageNo: 1},
			},
			isTargetMainArg: func(t argtables.ArgTable) bool {
				return t.IsCmd
			},
			isNextMainArg: func(t argtables.ArgTable) bool {
				return t.IsStage
			},
			isTargetOpt: func(t argtables.ArgTable) bool {
				return t.IsOpt
			},
			expectedOptParams: []OptParam{
				{
					Index:   3,
					OptStr:  "foo",
					Comment: "TestOpt",
					Param:   ParamType{},
				},
			},
		},
		{
			name:           "Parse option with value parameter and double quote",
			nextStartIndex: 0,
			curStageArgTables: []argtables.ArgTable{
				{IsCmd: true, StageNo: 1},
				{Str: strPtr("echo"), StageNo: 1},
				{IsOpt: true, Comment: "TestOpt", StageNo: 1},
				{Str: strPtr("foo"), StageNo: 1},
				{IsValue: true, StageNo: 1},
				{QuoteTypeSignal: argtables.DoubleQuote, Str: strPtr("bar_val"), StageNo: 1},
			},
			isTargetMainArg: func(t argtables.ArgTable) bool {
				return t.IsCmd
			},
			isNextMainArg: func(t argtables.ArgTable) bool {
				return t.IsStage
			},
			isTargetOpt: func(t argtables.ArgTable) bool {
				return t.IsOpt
			},
			expectedOptParams: []OptParam{
				{
					Index:   3,
					OptStr:  "foo",
					Comment: "TestOpt",
					Param: ParamType{
						Str:       strPtr("bar_val"),
						QuoteType: argtables.DoubleQuote,
						Comment:   "",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotOptParams []OptParam
			appendFn := func(op OptParam) {
				gotOptParams = append(gotOptParams, op)
			}

			parseOptions(
				tt.nextStartIndex,
				tt.curStageArgTables,
				tt.isTargetMainArg,
				tt.isNextMainArg,
				tt.isTargetOpt,
				appendFn,
			)

			if len(gotOptParams) == 0 && len(tt.expectedOptParams) == 0 {
				return
			}

			if !reflect.DeepEqual(gotOptParams, tt.expectedOptParams) {
				t.Errorf("parseOptions() got = %v, want %v", gotOptParams, tt.expectedOptParams)
			}
		})
	}
}

// strPtr is a helper function to return a pointer to a string literal.
func strPtr(s string) *string {
	return &s
}
