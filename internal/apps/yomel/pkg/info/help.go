package info

import (
	"errors"
	"fmt"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtablecounter"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtabledtos"
)

const detail = `Usage:
  yomel [flags] stage [desc] [cmd options/arguments...] [service options/arguments...] [action options/arguments...]

Meta Flags:
  --version          Print version information
  --help             Print help information
  --gen            outout total pipeline cmd 
  --direct            Exec shell by direct (simple exec pipe shell without log) 

General Flags:
  --log              Enable stdout logging for pipeline execution
  --no-log           Disable stdout logging for pipeline execution
  --log-filter       Filter stdout logs using shell commands
  --err-log-filter   Filter stderr logs using shell commands

Stage Parameters:
  stage              Define a new pipeline stage with a description
  -cmd               Specify the command to execute
  -svc               Specify the service name
  -act               Specify the action name
  --opt              Specify a short or long option key
  --lop              Specify a long option key
  --val              Specify an option value
  --arg              Specify a positional argument
  --single, -s       Indicate single-quoted value or argument
  --no-quote, -n       Indicate unquoted value or argument

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
        --log-filter "head -10" \
        stage "list" \
        -cmd "ls" \
        --opt "l" \
        --val --n "/var/log"

  3. Run with logging disable partly:
     yomel \
        --log \
        stage "list" \
        -cmd "ls" \
        --opt "l" \
        --val --n "/var/log"
        --no-log \
        stage "replace newline to space" \
        -cmd "tr" \
        --arg --s '\n' \
        --arg --s ' ' \
        --log-filter "head -1"
        stage "add prefix" \
        -cmd sed \
        --arg -s 's/^/$HOME/'`

func GetHelpByDefault(argList []string) (*string, error) {
	if len(argList) > 0 {
		return nil, nil
	}
	return execGetHelp()
}
func GetHelpByOption(argTablesDto []argtabledtos.ArgTableDto) (*string, error) {
	stageNo := 0
	for _, argTableDto := range argTablesDto {
		stageNo += argtablecounter.IncrementStageNo(argTableDto.IsStage)
		if stageNo > 0 {
			return nil, nil
		}
		if argTableDto.IsHelp {
			return execGetHelp()
		}
	}
	return nil, nil
}

func execGetHelp() (*string, error) {
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
