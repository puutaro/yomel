// internal/apps/yomel/pkg/sh/gen_stage_info_test.go
package sh

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/domain"
	"github.com/puutaro/yomel/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

// Test_genStageInfo verifies that genStageInfo correctly converts Yomel domain models into StageInfo slices across various configurations.
func Test_genStageInfo(t *testing.T) {
	tests := []struct {
		name  string
		input domain.Yomel
		want  []StageInfo
	}{
		{
			name: "should convert single stage Yomel model into StageInfo correctly",
			input: domain.Yomel{
				Ctrl: domain.Control{
					IsLog:        testutil.Ptr(true),
					LogFilter:    "global-filter",
					ErrLogFilter: "global-err-filter",
				},
				Stages: []domain.Stage{
					{
						No:        1,
						Desc:      "stage-desc",
						Cmd:       "echo",
						CmdOpArgs: []string{"hello"},
					},
				},
			},
			want: []StageInfo{
				{
					No:                 1,
					Desc:               "stage-desc",
					IsLog:              true,
					LogFilter:          "global-filter",
					ErrLogFilter:       "global-err-filter",
					CmdStrs:            "echo \\\n hello",
					CmdStrsWithComment: "echo",
				},
			},
		},
		{
			name: "should handle service, action, and comment options correctly across stages",
			input: domain.Yomel{
				Ctrl: domain.Control{
					IsLog: testutil.Ptr(false),
				},
				Stages: []domain.Stage{
					{
						No:                   1,
						Desc:                 "stage-1",
						Cmd:                  "aws",
						CmdOpArgsWithComment: []string{"s3"},
						Svc:                  "bucket",
						Act:                  "ls",
					},
				},
			},
			want: []StageInfo{
				{
					No:                 1,
					Desc:               "stage-1",
					IsLog:              false,
					LogFilter:          "",
					ErrLogFilter:       "",
					CmdStrs:            "aws \\\n bucket \\\n ls",
					CmdStrsWithComment: "aws \\\n s3 \\\n bucket \\\n ls",
				},
			},
		},
		{
			name: "should handle multiple stages with mixed control overrides and filters",
			input: domain.Yomel{
				Ctrl: domain.Control{
					IsLog:        testutil.Ptr(true),
					LogFilter:    "default-log",
					ErrLogFilter: "default-err",
				},
				Stages: []domain.Stage{
					{
						No:           1,
						Desc:         "first-stage",
						Cmd:          "kubectl",
						CmdOpArgs:    []string{"get", "pods"},
						IsLog:        testutil.Ptr(false),
						LogFilter:    "override-log",
						ErrLogFilter: "override-err",
					},
					{
						No:   2,
						Desc: "second-stage",
						Cmd:  "date",
					},
				},
			},
			want: []StageInfo{
				{
					No:                 1,
					Desc:               "first-stage",
					IsLog:              false,
					LogFilter:          "override-log",
					ErrLogFilter:       "override-err",
					CmdStrs:            "kubectl \\\n get \\\n pods",
					CmdStrsWithComment: "kubectl",
				},
				{
					No:                 2,
					Desc:               "second-stage",
					IsLog:              true,
					LogFilter:          "default-log",
					ErrLogFilter:       "default-err",
					CmdStrs:            "date",
					CmdStrsWithComment: "date",
				},
			},
		},
		{
			name: "should return empty slice when stages are empty",
			input: domain.Yomel{
				Ctrl:   domain.Control{},
				Stages: []domain.Stage{},
			},
			want: []StageInfo{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := genStageInfo(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
