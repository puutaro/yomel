//go:generate sh -c "cd ../.. && ./gen_yomel_info.sh"
package main

import (
	"fmt"
	"os"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/arg_list"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtablesvalid"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/domain"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/env"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/info"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/model"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/modelvalid"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/sh"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/toml"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/tomlvalid"
)

const (
	normalExitSignal = 0
	errorExitSignal  = 1
)

func main() {

	inputArgs := arg_list.Gen()
	helpConByDefault, helpErrByDefault := info.GetHelpByDefault(inputArgs)
	if helpErrByDefault != nil {
		fmt.Fprintf(os.Stderr, "%s\n", helpErrByDefault)
		os.Exit(errorExitSignal)
	}
	if helpConByDefault != nil {
		fmt.Fprintf(os.Stdout, "%s\n", *helpConByDefault)
		os.Exit(normalExitSignal)
	}
	argTables := argtables.GenArgTable(inputArgs)
	helpConByOption, helpErrByOption := info.GetHelpByOption(argTables)
	if helpErrByOption != nil {
		fmt.Fprintf(os.Stderr, "%s\n", helpErrByOption)
		os.Exit(errorExitSignal)
	}
	if helpConByOption != nil {
		fmt.Fprintf(os.Stdout, "%s\n", *helpConByOption)
		os.Exit(normalExitSignal)
	}
	version, versionErr := info.GetVersion(argTables)
	if versionErr != nil {
		fmt.Fprintf(os.Stderr, "%s\n", versionErr)
		os.Exit(errorExitSignal)
	}
	if version != nil {
		fmt.Fprintf(os.Stdout, "%s\n", *version)
		os.Exit(normalExitSignal)
	}
	if argtablesValidErr := argtablesvalid.ArgTableValidate(argTables); argtablesValidErr != nil {
		fmt.Fprintf(os.Stderr, "%s\n", argtablesValidErr)
		os.Exit(errorExitSignal)
	}
	ctrl, stageModels := model.Parse(argTables)
	if modelValidErr := modelvalid.ModelValidate(stageModels); modelValidErr != nil {
		fmt.Fprintf(os.Stderr, "%s\n", modelValidErr)
		os.Exit(errorExitSignal)
		return
	}
	if ctrlModelValidErr := modelvalid.CtrlModeValidate(ctrl, len(stageModels)); ctrlModelValidErr != nil {
		fmt.Fprintf(os.Stderr, "%s\n", ctrlModelValidErr)
		os.Exit(errorExitSignal)
		return
	}
	yomelEnv := env.ReadEnd(ctrl.IsDirect, ctrl.IsGen)
	yomelToml := toml.ReadToml(yomelEnv.YomelTomlPath)
	if yomelToml.Color.EnableLightMode == int(env.Light) {
		yomelEnv.IsLightColorMode = true
	}
	if yomelToml.Stream.EnableTee == int(env.TeeOn) {
		yomelEnv.IsTerminal = true
	}
	if tomlErr := tomlvalid.TomlValidate(
		yomelToml,
		yomelEnv.YomelTomlPath,
	); tomlErr != nil {
		fmt.Fprintf(os.Stderr, "%s\n", tomlErr)
		os.Exit(errorExitSignal)
		return
	}
	yomel := domain.Convert(
		ctrl,
		stageModels,
		yomelToml,
		yomelEnv.IsTerminal,
	)
	yomelInfo := sh.Gen(yomel)
	finishStatus := sh.Exec(yomelInfo, yomelEnv)
	os.Exit(finishStatus)

}
