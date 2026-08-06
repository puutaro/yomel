// Write direct above line for Comment on code
package info

import (
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
	"github.com/stretchr/testify/assert"
)

func Test_GetVersion(t *testing.T) {
	var info YomelInfo
	toml.Decode(YomelInfoRaw, &info)
	version := info.Yomel.Version
	tests := []struct {
		name      string
		argTables []argtables.ArgTableDto
		want      *string
	}{
		{
			name: "should return version string when IsVersion is true",
			argTables: []argtables.ArgTableDto{
				{StageNo: 0, IsVersion: true},
			},
			want: &version,
		},
		{
			name: "should return nil when IsVersion is false",
			argTables: []argtables.ArgTableDto{
				{StageNo: 0, IsVersion: false},
			},
			want: nil,
		},
		{
			name:      "should return nil when argTables is empty",
			argTables: []argtables.ArgTableDto{},
			want:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetVersion(tt.argTables)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
