//go:generate sh -c "cd ../.. && ./gen_yomel_info.sh"
package main

import (
	"fmt"
	"os"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/arglist"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtable"
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
	argTables := argtable.GenArgTable(inputArgs)

	if preValidateErr := argtablevalid.ArgTableValidate(argTables); preValidateErr != nil {
		fmt.Fprintf(os.Stderr, "%s\n", preValidateErr)
		os.Exit(errorExitSignal)
		return
	}
	ctrl, stageModels := model.Parse(argTables)

	if modelValidErr := modelvalid.ModelValidate(stageModels); modelValidErr != nil {
		fmt.Fprintf(os.Stderr, "%s\n", modelValidErr)
		os.Exit(errorExitSignal)
		return
	}
	yomel := domain.Convert(ctrl, stageModels)
	helpCon, helpErr := info.GetHelp(yomel.Ctrl)
	if helpErr != nil {
		fmt.Fprintf(os.Stderr, "%s\n", helpErr)
		os.Exit(errorExitSignal)
	}
	if helpCon != nil {
		fmt.Fprintf(os.Stdout, "%s\n", *helpCon)
		os.Exit(normalExitSignal)
	}
	version, versionErr := info.GetVersion(yomel.Ctrl)
	if versionErr != nil {
		fmt.Fprintf(os.Stderr, "%s\n", versionErr)
		os.Exit(errorExitSignal)
	}
	if version != nil {
		fmt.Fprintf(os.Stdout, "%s\n", *version)
		os.Exit(normalExitSignal)
	}
	chainStr := sh.Gen(yomel)
	sh.Exec(chainStr)

}
