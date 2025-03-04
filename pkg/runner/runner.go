package runner

import "os/exec"
import "log"

func GetGraph(arg string) string {
	cmd := exec.Command("bazel", "query", "--notool_deps", "deps("+arg+")", "--output", "graph", "--graph:node_limit", "-1")
	log.Printf("Running command: %s\n", cmd)
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(stdout)
}
