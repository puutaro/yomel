// internal/apps/yomel/pkg/color/color_tool_ConvertStrListToColorStr_test.go
package color

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_ConvertStrListToColorStr verifies that ConvertStrListToColorStr correctly handles empty strings, various color mode constants, and raw fallback strings.
func Test_ConvertStrListToColorStr(t *testing.T) {
	tests := []struct {
		name   string
		strSrc string
		want   []string // 候補となるカラーコード群、または固定の戻り値
	}{
		{
			name:   "should return empty string when strSrc is empty",
			strSrc: "",
			want:   []string{""},
		},
		{
			name:   "should return light color code when strSrc is LightRnd",
			strSrc: LightRnd,
			want: []string{
				lightAzureCode,
				lightBlueCode,
				lightBrownCode,
				lightGrayCode,
				lightGreenCode,
				lightRedCode,
				yellowCode,
			},
		},
		{
			name:   "should return dark color code when strSrc is DarkRnd",
			strSrc: DarkRnd,
			want: []string{
				darkAzureCode,
				darkBlueCode,
				darkBrownCode,
				darkGrayCode,
				darkGreenCode,
				darkRedCode,
			},
		},
		{
			name:   "should return center color code when strSrc is CenterRnd",
			strSrc: CenterRnd,
			want: []string{
				azureCode,
				blueCode,
				brownCode,
				gray,
				redCode,
			},
		},
		{
			name:   "should return rnd color code when strSrc is Rnd",
			strSrc: Rnd,
			want: []string{
				lightAzureCode,
				lightBlueCode,
				lightBrownCode,
				lightGrayCode,
				lightGreenCode,
				lightRedCode,
				darkAzureCode,
				darkBlueCode,
				darkBrownCode,
				darkGrayCode,
				darkGreenCode,
				darkRedCode,
				azureCode,
				blueCode,
				brownCode,
				gray,
				redCode,
			},
		},
		{
			name:   "should return raw string when default case is matched",
			strSrc: "custom-raw-color",
			want:   []string{"custom-raw-color"},
		},
		{
			name:   "should handle string list split by AndOperator correctly",
			strSrc: LightRnd + AndOperator + "custom-raw-color",
			want: []string{
				lightAzureCode,
				lightBlueCode,
				lightBrownCode,
				lightGrayCode,
				lightGreenCode,
				lightRedCode,
				yellowCode,
				"custom-raw-color",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertStrListToColorStr(tt.strSrc)
			if tt.strSrc == "" || tt.strSrc == "custom-raw-color" {
				assert.Equal(t, tt.want[0], got)
			} else {
				assert.Contains(t, tt.want, got)
			}
		})
	}
}
