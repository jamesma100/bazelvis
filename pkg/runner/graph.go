package runner

import "os/exec"
import "fmt"

func GetGraph(arg string) string {
  cmd := exec.Command("./bazel", "query", "--notool_deps", "deps(" + arg + ")", "--output", "graph")
  fmt.Printf("Running command: %s\n", cmd)
  stdout, err := cmd.Output()
  if err != nil {
    fmt.Println(err)
  }
  return string(stdout)
}
