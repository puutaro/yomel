package info

import (
	"errors"
	"fmt"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtablecounter"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
)

const detail = `Usage:
  yomel [flags] stage [desc] [cmd options/arguments...] [service options/arguments...] [action options/arguments...]

Meta Flags:
  --version          Print version information
  --help             Print help information
  title               Specify a title for the output report or the generated job.
  --no-live-stdout      Disable live output of stdout.
  --no-live-stderr      Disable live output of stderr.
  --gen              Output total pipeline command
  --direct           Exec shell directly (simple exec pipe shell without log)

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
  --opt[PascalCase]  Specify a short option key with an optional Alphanumeric PascalCase description suffix
  --lop[PascalCase]  Specify a long option key with an optional Alphanumeric PascalCase description suffix
  --val[PascalCase]  Specify an option value with an optional Alphanumeric PascalCase description suffix
  --arg[PascalCase]  Specify a positional argument with an optional Alphanumeric PascalCase description suffix
  --single, -s       Indicate single-quoted value or argument
  --no-quote, -n     Indicate unquoted value or argument

Examples:
  1. Retrieve logs from S3, extract them, and grep for errors:
     yomel \
       stage "download" \
       -cmd "aws" \
       -svc "s3" \
       -act "cp" \
       --argS3Path --s "s3://my-bucket/logs.tar.gz" \
       --argDest --n "-" \
       stage "extract" \
       -cmd "tar" \
       --optX \
       --valCompressType --n "z" \
       --optO \
       --valDest --n "-" \
       stage "search" \
       -cmd "grep" \
       --argPattern --s "ERROR"

  2. Run with global logging enabled:
     yomel \
        --log \
        --log-filter "head -10" \
        stage "list" \
        -cmd "ls" \
        --optL \
        --valTargetDir --n "/var/log"

  3. Run with logging disabled partly:
     yomel \
        --log \
        stage "list" \
        -cmd "ls" \
        --optL \
        --valTargetDir --n "/var/log" \
        --no-log \
        stage "replace newline to space" \
        -cmd "tr" \
        --argFrom --s '\n' \
        --argTo --s ' ' \
        --log-filter "head -1" \
        stage "add prefix" \
        -cmd sed \
        --argPattern --s 's/^/$HOME/'`

func GetHelpByDefault(argList []string) (*string, error) {
	if len(argList) > 0 {
		return nil, nil
	}
	return execGetHelp()
}
func GetHelpByOption(argTables []argtables.ArgTable) (*string, error) {
	stageNo := 0
	for _, argTable := range argTables {
		stageNo += argtablecounter.IncrementStageNo(argTable.IsStage)
		if stageNo > 0 {
			return nil, nil
		}
		if argTable.IsHelp {
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
