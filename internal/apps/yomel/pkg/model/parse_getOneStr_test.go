package model

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtabledtos"
	"github.com/puutaro/yomel/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func Test_getOneStr(t *testing.T) {
	// Tiny helpers to minimize structural boilerplate
	tStr := func(s string) argtabledtos.ArgTableDto { return argtabledtos.ArgTableDto{Str: &s} }
	tCmd := func() argtabledtos.ArgTableDto { return argtabledtos.ArgTableDto{IsCmd: true} }
	tStage := func() argtabledtos.ArgTableDto { return argtabledtos.ArgTableDto{IsStage: true} }

	tests := []struct {
		name           string
		nextStartIndex int
		input          []argtabledtos.ArgTableDto
		isCheckFn      func(argtabledtos.ArgTableDto) bool
		want           *string
	}{
		{
			name:           "should return string pointer when target flag is matched and followed by a string",
			nextStartIndex: 0,
			input: []argtabledtos.ArgTableDto{
				tCmd(), tStr("aws"),
			},
			isCheckFn: func(a argtabledtos.ArgTableDto) bool { return a.IsCmd },
			want:      testutil.Ptr("aws"),
		},
		{
			name:           "should skip elements before nextStartIndex",
			nextStartIndex: 2,
			input: []argtabledtos.ArgTableDto{
				tStage(), tStr("skipped-stage"),
				tStage(), tStr("target-stage"),
			},
			isCheckFn: func(a argtabledtos.ArgTableDto) bool { return a.IsStage },
			want:      testutil.Ptr("target-stage"),
		},
		{
			name:           "should return nil when matched flag is at the end of the slice",
			nextStartIndex: 0,
			input: []argtabledtos.ArgTableDto{
				tCmd(),
			},
			isCheckFn: func(a argtabledtos.ArgTableDto) bool { return a.IsCmd },
			want:      nil,
		},
		{
			name:           "should return nil when the next element's Str field is nil",
			nextStartIndex: 0,
			input: []argtabledtos.ArgTableDto{
				tCmd(),
				tStage(), // Str field is nil
			},
			isCheckFn: func(a argtabledtos.ArgTableDto) bool { return a.IsCmd },
			want:      nil,
		},
		{
			name:           "should return nil when target flag does not exist in the slice",
			nextStartIndex: 0,
			input: []argtabledtos.ArgTableDto{
				tCmd(), tStr("aws"),
			},
			isCheckFn: func(a argtabledtos.ArgTableDto) bool { return a.IsStage },
			want:      nil,
		},
		{
			name:           "should return nil immediately when nextStartIndex exceeds input slice length",
			nextStartIndex: 3,
			input: []argtabledtos.ArgTableDto{
				tCmd(), tStr("aws"),
			},
			isCheckFn: func(a argtabledtos.ArgTableDto) bool { return a.IsCmd },
			want:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getOneStr(tt.nextStartIndex, tt.input, tt.isCheckFn)
			assert.Equal(t, tt.want, got)
		})
	}
}
