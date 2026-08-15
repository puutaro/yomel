// Test for ConvertFixColorStr to cover all color code switch cases and the default fallback.
package color

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ConvertFixColorStr(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect string
	}{
		{
			name:   "should convert darkBlueCode to hex",
			input:  darkBlueCode,
			expect: "#001d9e",
		},
		{
			name:   "should convert blueCode to hex",
			input:  blueCode,
			expect: "#0026ff",
		},
		{
			name:   "should convert lightBlueCode to hex",
			input:  lightBlueCode,
			expect: "#26c0fc",
		},
		{
			name:   "should convert darkRedCode to hex",
			input:  darkRedCode,
			expect: "#850101",
		},
		{
			name:   "should convert redCode to hex",
			input:  redCode,
			expect: "#ff0000",
		},
		{
			name:   "should convert lightRedCode to hex",
			input:  lightRedCode,
			expect: "#fa4d4d",
		},
		{
			name:   "should convert darkGreenCode to hex",
			input:  darkGreenCode,
			expect: "#014702",
		},
		{
			name:   "should convert greenCode to hex",
			input:  greenCode,
			expect: "#11b812",
		},
		{
			name:   "should convert lightGreenCode to hex",
			input:  lightGreenCode,
			expect: "#43fa46",
		},
		{
			name:   "should convert yellowCode to hex",
			input:  yellowCode,
			expect: "#fff200",
		},
		{
			name:   "should convert darkAzureCode to hex",
			input:  darkAzureCode,
			expect: "#1c4d44",
		},
		{
			name:   "should convert azureCode to hex",
			input:  azureCode,
			expect: "#21ebc6",
		},
		{
			name:   "should convert lightAzureCode to hex",
			input:  lightAzureCode,
			expect: "#67e8eb",
		},
		{
			name:   "should convert darkBrownCode to hex",
			input:  darkBrownCode,
			expect: "#572b07",
		},
		{
			name:   "should convert brownCode to hex",
			input:  brownCode,
			expect: "#b35c15",
		},
		{
			name:   "should convert lightBrownCode to hex",
			input:  lightBrownCode,
			expect: "#f26e27",
		},
		{
			name:   "should convert black to hex",
			input:  black,
			expect: "#000000",
		},
		{
			name:   "should convert white to hex",
			input:  white,
			expect: "#ffffff",
		},
		{
			name:   "should convert darkGrayCode to hex",
			input:  darkGrayCode,
			expect: "#424242",
		},
		{
			name:   "should convert gray to hex",
			input:  gray,
			expect: "#808080",
		},
		{
			name:   "should convert lightGrayCode to hex",
			input:  lightGrayCode,
			expect: "#dbdbdb",
		},
		{
			name:   "should return raw input when unknown color string is passed (default case)",
			input:  "unknown-color-string",
			expect: "unknown-color-string",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertFixColorStr(tt.input)
			assert.Equal(t, tt.expect, got)
		})
	}
}
