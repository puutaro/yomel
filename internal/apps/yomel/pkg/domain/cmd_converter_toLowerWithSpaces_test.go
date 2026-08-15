package domain

import (
	"testing"
)

func Test_toLowerWithSpaces(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "empty string",
			arg:  "",
			want: "",
		},
		{
			name: "all lowercase",
			arg:  "abc",
			want: "abc",
		},
		{
			name: "single uppercase",
			arg:  "A",
			want: "a",
		},
		{
			name: "camel case",
			arg:  "ToLowerWithSpaces",
			want: "to lower with spaces",
		},
		{
			name: "consecutive uppercase letters",
			arg:  "HTTPServer",
			want: "HTTP server",
		},
		{
			name: "mixed with numbers and symbols",
			arg:  "Test123ABC",
			want: "test123AB c",
		},
		{
			name: "leading and trailing spaces",
			arg:  "  TestString  ",
			want: "   test string  ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toLowerWithSpaces(tt.arg); got != tt.want {
				t.Errorf("toLowerWithSpaces() = %q, want %q", got, tt.want)
			}
		})
	}
}
