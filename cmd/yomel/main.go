//go:generate sh -c "cd ../.. && ./gen_yomel_info.sh"
package main

import (
	"fmt"
	"os"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/arglist"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtabledtos"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtabledtosvalid"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtablevalid"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/domain"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/info"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/model"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/modelvalid"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/sh"
)

const (
	normalExitSignal = 0
	errorExitSignal  = 1
)

func main() {

	inputArgs := arglist.Gen()
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
	if argTableDtosValidErr := argtablevalid.ArgTableValid(argTables); argTableDtosValidErr != nil {
		fmt.Fprintf(os.Stderr, "%s\n", argTableDtosValidErr)
		os.Exit(errorExitSignal)
	}
	argTableDtos := argtabledtos.GenArgTableDto(argTables)
	if argTableValidateErr := argtabledtosvalid.ArgTableValidate(argTableDtos); argTableValidateErr != nil {
		fmt.Fprintf(os.Stderr, "%s\n", argTableValidateErr)
		os.Exit(errorExitSignal)
		return
	}
	ctrl, stageModels := model.Parse(argTableDtos)
	if modelValidErr := modelvalid.ModelValidate(stageModels); modelValidErr != nil {
		fmt.Fprintf(os.Stderr, "%s\n", modelValidErr)
		os.Exit(errorExitSignal)
		return
	}
	yomel := domain.Convert(ctrl, stageModels)
	yomelInfo := sh.Gen(yomel)
	execErr := sh.Exec(yomelInfo)
	if execErr != nil {
		os.Exit(errorExitSignal)
	}

}
