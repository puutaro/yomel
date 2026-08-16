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

Telemetry and Filter Options:
  title "<pipeline_title>"    Specify a title for the overall pipeline
  --no-live-stdout            Suppress real-time streaming of standard output (stdout)
  --no-live-stderr            Suppress real-time streaming of standard error (stderr)
  --log                       Activate the internal logging system
  --gen                       Outputs the total pipeline command (dry-run)
  --direct                    Executes the shell pipeline directly without internal logging decoration
  --log-filter "<shell_command>"    Attaches an asynchronous log interceptor for standard output (stdout)
  --err-log-filter "<shell_command>" Attaches an asynchronous log interceptor for standard error (stderr)

Structural Stage Elements:
  stage "<stage_name>"        Initializes a new execution boundary (pipeline stage)
  -cmd "<binary>"             Specifies the main executable or binary command
  -svc "<service_name>"       Declares a sub-service or second-level command hierarchy
  -act "<action_name>"        Declares the operation, verb, or action to be performed

Color Control Options:
  --color "<color_code>"          Specifies the foreground text color for the command body and logs
  --bg-color "<color_code>"       Specifies the background color for the section header/panels
  --comment-color "<color_code>"  Customizes the color used for comments inside the pipeline commands
  --title-color "<color_code>"    Sets the global title and header text color (Global control option only)
  --title-bg-color "<color_code>" Sets the background color for the global title and overall pipeline log headers
  --title-comment-color "<color_code>" Sets the comment color specifically for the total pipeline command section

Option and Argument Value Modifiers:
  --opt[PascalCase] "<flag>"  Generates a short-style option flag with an optional PascalCase description suffix
  --lop[PascalCase] "<flag>"  Generates a long-style option flag with an optional PascalCase description suffix
  --val[PascalCase]           Declares a value associated with the preceding option (must be followed by --s or --n)
  --arg[PascalCase]           Appends a standalone, positional argument to the tail end (must be followed by --s or --n)
  --single, -s                Indicate single-quoted value or argument
  --no-quote, -n              Indicate unquoted value or argument

Environment Variables and TOML Configuration:
  YOMEL_LIGHT_COLOR_MODE      Controls whether to enable the light color theme mode
  YOMEL_ENABLE_TEE            Controls whether terminal tee streaming characteristics are enabled
  YOMEL_TOML_PATH             Specifies the custom file path to the external yomel.toml configuration file`

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
