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
		argTables []argtables.ArgTable
		want      *string
		wantErr   bool
	}{
		{
			name: "should return version string when IsVersion is true",
			argTables: []argtables.ArgTable{
				{StageNo: 0, IsVersion: true},
			},
			want:    &version,
			wantErr: false,
		},
		{
			name: "should return nil when IsVersion is false",
			argTables: []argtables.ArgTable{
				{StageNo: 0, IsVersion: false},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name:      "should return nil when argTables is empty",
			argTables: []argtables.ArgTable{},
			want:      nil,
			wantErr:   false,
		},
		{
			name: "should handle multiple argTables where IsVersion is true",
			argTables: []argtables.ArgTable{
				{StageNo: 0, IsVersion: false},
				{StageNo: 0, IsVersion: true},
			},
			want:    &version,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetVersion(tt.argTables)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
