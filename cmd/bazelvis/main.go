package main

import (
	"github.com/jamesma100/bazelvis/pkg/runner"
	"github.com/jamesma100/bazelvis/pkg/parser"
	"os"
  "fmt"
)


func main() {
  arg := os.Args[1]
  fmt.Printf("Received arg: %s\n", arg)
  contents := runner.GetGraph(arg)
  m := parser.BuildMap(contents)
  fmt.Printf("Total length of m: %d\n", len(m))
}
