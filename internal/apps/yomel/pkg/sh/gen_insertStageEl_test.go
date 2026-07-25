package sh

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_stageCommand_insertStageEl verifies that insertStageEl correctly appends strings with the given prefix.
func Test_stageCommand_insertStageEl(t *testing.T) {
	tests := []struct {
		name       string
		initial    stageCommand
		insertStrs []string
		prefix     string
		want       stageCommand
	}{
		{
			name:       "should append strings with prefix when insertStrs is not empty",
			initial:    stageCommand("base"),
			insertStrs: []string{"arg1", "arg2"},
			prefix:     " ",
			want:       stageCommand("base arg1 \\\n arg2"),
		},
		{
			name:       "should do nothing when insertStrs is empty",
			initial:    stageCommand("base"),
			insertStrs: []string{},
			prefix:     " ",
			want:       stageCommand("base"),
		},
		{
			name:       "should append single string with prefix when insertStrs has one element",
			initial:    stageCommand("echo"),
			insertStrs: []string{"hello"},
			prefix:     " \\\n ",
			want:       stageCommand("echo \\\n hello"),
		},
		{
			name:       "should append multiple elements correctly with complex prefix",
			initial:    stageCommand("aws"),
			insertStrs: []string{"s3", "cp"},
			prefix:     " \\\n ",
			want:       stageCommand("aws \\\n s3 \\\n cp"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := tt.initial
			sc.insertStageEl(tt.insertStrs, tt.prefix)
			assert.Equal(t, tt.want, sc)
		})
	}
}
