package main

import (
	"github.com/jamesma100/bazelvis/pkg/parser"
	"github.com/jamesma100/bazelvis/pkg/runner"
	"github.com/jamesma100/bazelvis/pkg/ui"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Fatal("Invalid argument count")
	}
	arg := args[1]
	contents := runner.GetGraph(arg)
	m := parser.BuildMap(contents)

	ui.StartUI(arg, m)

}
