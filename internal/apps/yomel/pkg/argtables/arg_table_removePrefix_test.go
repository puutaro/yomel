package argtables

import (
	"testing"

	"github.com/puutaro/yomel/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func Test_removePrefix(t *testing.T) {
	tests := []struct {
		name   string
		str    string
		prefix string
		want   *string
	}{
		{
			name:   "should remove prefix when string starts with prefix",
			str:    "--opt",
			prefix: "--",
			want:   testutil.Ptr("opt"),
		},
		{
			name:   "should remove prefix for single character prefix",
			str:    "-cmd",
			prefix: "-",
			want:   testutil.Ptr("cmd"),
		},
		{
			name:   "should return original string without changes when prefix does not match",
			str:    "opt",
			prefix: "--",
			want:   testutil.Ptr("opt"),
		},
		{
			name:   "should handle empty string correctly",
			str:    "",
			prefix: "--",
			want:   testutil.Ptr(""),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removePrefix(tt.str, tt.prefix)
			assert.Equal(t, tt.want, got)
		})
	}
}
