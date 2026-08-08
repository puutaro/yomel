package sh

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/domain"
	"github.com/puutaro/yomel/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func Test_Gen(t *testing.T) {
	tests := []struct {
		name  string
		input domain.Yomel
		want  YomelInfo
	}{
		{
			name: "should convert Yomel structure with control flags, title, and stages into YomelInfo correctly",
			input: domain.Yomel{
				Ctrl: domain.Control{
					IsLiveStdout: true,
					IsLiveStderr: false,
					IsDirect:     true,
					IsGen:        false,
					Title:        "pipeline-title",
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
			want: YomelInfo{
				IsLiveStdout: true,
				IsLiveStdErr: false,
				IsDirect:     true,
				IsGen:        false,
				Title:        "pipeline-title",
				StageInfos: []StageInfo{
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
		},
		{
			name: "should convert Yomel structure with empty title and multiple stages including service and action options correctly",
			input: domain.Yomel{
				Ctrl: domain.Control{
					IsLiveStdout: false,
					IsLiveStderr: true,
					IsDirect:     false,
					IsGen:        true,
					Title:        "",
				},
				Stages: []domain.Stage{
					{
						No:        1,
						Desc:      "stage-1",
						Cmd:       "aws",
						CmdOpArgs: []string{"s3"},
						Svc:       "bucket",
						SvcOpArgs: []string{"ls"},
						Act:       "run",
						ActOpArgs: []string{"--recursive"},
					},
				},
			},
			want: YomelInfo{
				IsLiveStdout: false,
				IsLiveStdErr: true,
				IsDirect:     false,
				IsGen:        true,
				Title:        "",
				StageInfos: []StageInfo{
					{
						No:                 1,
						Desc:               "stage-1",
						IsLog:              false,
						LogFilter:          "",
						ErrLogFilter:       "",
						CmdStrs:            "aws \\\n s3 \\\n bucket \\\n ls \\\n run \\\n --recursive",
						CmdStrsWithComment: "aws \\\n bucket \\\n run",
					},
				},
			},
		},
		{
			name: "should handle empty stages and default control settings correctly",
			input: domain.Yomel{
				Ctrl: domain.Control{
					IsLiveStdout: false,
					IsLiveStderr: true,
					IsDirect:     false,
					IsGen:        true,
					Title:        "",
				},
				Stages: []domain.Stage{},
			},
			want: YomelInfo{
				IsLiveStdout: false,
				IsLiveStdErr: true,
				IsDirect:     false,
				IsGen:        true,
				Title:        "",
				StageInfos:   []StageInfo{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Gen(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
