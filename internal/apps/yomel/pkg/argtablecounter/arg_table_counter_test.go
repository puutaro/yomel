package argtablecounter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_IncrementStageNo verifies that IncrementStageNo returns 1 when isStage is true, and 0 when false.
func Test_IncrementStageNo(t *testing.T) {
	tests := []struct {
		name    string
		isStage bool
		want    int
	}{
		{
			name:    "should return 1 when isStage is true",
			isStage: true,
			want:    1,
		},
		{
			name:    "should return 0 when isStage is false",
			isStage: false,
			want:    0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IncrementStageNo(tt.isStage)
			assert.Equal(t, tt.want, got)
		})
	}
}
