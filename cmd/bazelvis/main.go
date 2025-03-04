package main

import (
	"github.com/jamesma100/bazelvis/pkg/parser"
	"github.com/jamesma100/bazelvis/pkg/runner"
	"github.com/jamesma100/bazelvis/pkg/ui"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Fatal("Invalid argument count")
	}
	arg := args[1]
	if !strings.HasPrefix(arg, "//") {
		if !strings.HasPrefix(arg, "@") {
			arg = "//" + arg
		}
	}
	colSplt := strings.Split(arg, ":")
	if len(colSplt) == 1 {
		slashSplt := strings.Split(arg, "/")
		target := slashSplt[len(slashSplt)-1]
		arg += ":"
		arg += target
	}
	contents := runner.GetGraph(arg)
	m := parser.BuildMap(contents)

	ui.StartUI(arg, m)

}
