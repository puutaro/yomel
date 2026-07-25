package sh

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/model"
	"github.com/puutaro/yomel/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

// Test_Gen verifies that Gen correctly converts Yomel structures into execution-ready YomelInfo slices.
func Test_Gen(t *testing.T) {
	tests := []struct {
		name  string
		input model.Yomel
		want  []YomelInfo
	}{
		{
			name: "should generate single stage YomelInfo with command and arguments correctly",
			input: model.Yomel{
				Ctrl: model.Control{
					IsLog:        testutil.Ptr(true),
					LogFilter:    "global-filter",
					ErrLogFilter: "global-err-filter",
				},
				Stages: []model.Stage{
					{
						No:        1,
						Desc:      "echo-stage",
						Cmd:       "echo",
						CmdOpArgs: []string{"hello"},
					},
				},
			},
			want: []YomelInfo{
				{
					No:           1,
					Desc:         "echo-stage",
					IsLog:        true,
					LogFilter:    "global-filter",
					ErrLogFilter: "global-err-filter",
					CmdStrs:      "echo \\\n hello",
				},
			},
		},
		{
			name: "should override global filters and log settings with stage-specific configurations",
			input: model.Yomel{
				Ctrl: model.Control{
					IsLog:        testutil.Ptr(false),
					LogFilter:    "global-filter",
					ErrLogFilter: "global-err-filter",
				},
				Stages: []model.Stage{
					{
						No:           1,
						Desc:         "complex-stage",
						Cmd:          "aws",
						CmdOpArgs:    []string{"--profile prod", `"cmd-arg"`},
						Svc:          "s3",
						SvcOpArgs:    []string{"-r", "--svc-lop 'val'", "svc-arg"},
						Act:          "cp",
						ActOpArgs:    []string{"-v", "--recursive", "act-arg"},
						IsLog:        testutil.Ptr(true),
						LogFilter:    "stage-filter",
						ErrLogFilter: "stage-err-filter",
					},
				},
			},
			want: []YomelInfo{
				{
					No:           1,
					Desc:         "complex-stage",
					IsLog:        true,
					LogFilter:    "stage-filter",
					ErrLogFilter: "stage-err-filter",
					CmdStrs:      "aws \\\n --profile prod \\\n \"cmd-arg\" \\\n s3 \\\n -r \\\n --svc-lop 'val' \\\n svc-arg \\\n cp \\\n -v \\\n --recursive \\\n act-arg",
				},
			},
		},
		{
			name: "should return empty slice when stages are empty",
			input: model.Yomel{
				Ctrl:   model.Control{},
				Stages: []model.Stage{},
			},
			want: []YomelInfo{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Gen(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
