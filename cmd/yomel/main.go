//go:generate sh -c "cd ../.. && ./gen_yomel_info.sh"
package main

import (
	"fmt"
	"os"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/args"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/info"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/parser"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/sh"
)

const (
	normalExitSignal = 0
	errorExitSignal  = 1
)

func main() {

	argTables := args.GenArgTable()
	yomel := parser.Parse(argTables)
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
