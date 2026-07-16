package info

import (
	"errors"
	"fmt"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/parser"
)

const detail = `yomel - YAML-like shellscript pipeline runner

Usage:
  yomel [global flags] stage <stage-name> [stage flags...] [stage <stage-name> [stage flags...] ...]

Global Flags:
  --help, -h               Show this help message and exit.
  --version, -v            Show version information and exit.
  --log                    Enable execution logging (displays constructed shell commands).
  --log-filter <cmd>       Specify a command to filter standard output logs (e.g., "grep 'Exception'").
  --err-log-filter <cmd>   Specify a command to filter standard error logs.

Stage Flags (Can be specified after each "stage" keyword):
  -cmd, --cmd <string>     Base command to run in this stage (e.g., "aws", "curl", "sed").
  -svc, --svc <string>     First sub-command or logical block (e.g., "s3", "logs").
  -act, --act <string>     Second sub-command or action (e.g., "cp", "filter").

  --lop "<option>"         Appends a long option (--key). Must be followed by a value modifier (--val --s or --val --n).
                           Example: --lop "region" --val --s "us-east-1" -> --region 'us-east-1'

  -sop, --opt "<option>"   Appends a short option (-key). Must be followed by a value modifier (--val --s or --val --n).
                           Example: --opt "f" --val --n "Dockerfile" -> -f Dockerfile

  --val                    Value modifier flag accompanying --lop or --opt.
    --val --s "<string>"   Emits value enclosed in single quotes.
    --val --n "<string>"   Emits raw value without quotes (useful for numbers or shell variables).

  --arg                    Appends a standalone positional argument at the end.
    --arg --s "<string>"   Appends a single-quoted positional argument.
    --arg --n "<string>"   Appends an unquoted positional argument.

Examples:
  1. Retrieve logs from S3, extract them, and grep for errors:
     yomel \
       stage "download" \
       -cmd "aws" \
       -svc "s3" \
       -act "cp" \
       --arg --s "s3://my-bucket/logs.tar.gz" \
       --arg --n "-" \
       stage "extract" \
       -cmd "tar" \
       --opt "x" \
       --val --n "z" \
       --opt "O" \
       --val --n "-" \
       stage "search" \
       -cmd "grep" \
       --arg --s "ERROR"

  2. Run with global logging enabled:
     yomel \
        --log \
        stage "list" \
        -cmd "ls" \
        --opt "l" \
        --val \
        --n "/var/log"`

func GetHelp(ctrl parser.Control) (*string, error) {
	if !ctrl.IsHelp {
		return nil, nil
	}
	var info YomelInfo
	if _, err := toml.Decode(YomelInfoRaw, &info); err != nil {
		return nil, fmt.Errorf("failed to parse yomel.toml: %v\n", err)
	}
	description := info.Yomel.Description
	if description == "" {
		return nil, errors.New("unknown")
	}
	helpCon := strings.Join(
		[]string{description, "", detail},
		"\n",
	)
	return &helpCon, nil
}
