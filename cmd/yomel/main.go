package main

import (
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/args"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/sh"
)

func main() {
	argTables := args.GenArgTable()
	// pp.Printf("argTables stages: %+v\n", argTables)
	chain := sh.Parse(argTables)
	// pp.Printf("stages stages: %+v\n", chain)
	chainStr := sh.Gen(chain)
	// pp.Printf("chain \n%v\n", chainStr)
	sh.Exec(chainStr)

}
