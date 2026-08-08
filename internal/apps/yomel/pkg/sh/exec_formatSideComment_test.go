// Test_formatSideComment verifies that formatSideComment correctly aligns side comments across multiple lines.
package sh

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_formatSideComment(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect string
	}{
		{
			name:   "should align side comments correctly when multiple lines contain sideCommentBlank",
			input:  "echo 'hello'SIED_COmMetN_BAlNk`#first comment`\necho 'world'SIED_COmMetN_BAlNk`#second comment`",
			expect: "echo 'hello'    `#first comment`\necho 'world'    `#second comment`",
		},
		{
			name:   "should return original string when no sideCommentBlank is present",
			input:  "echo 'hello'\necho 'world'",
			expect: "echo 'hello'\necho 'world'",
		},
		{
			name:   "should handle empty string correctly",
			input:  "",
			expect: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatSideComment(tt.input)
			assert.Equal(t, tt.expect, got)
		})
	}
}
