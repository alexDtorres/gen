package gen

import (
  "os"
  "os/exec"
)

func ExampleRunner() {
  appGenTrialA()
	// Output: tmp/main.go
	clean()
}

func appGenTrialA() {
  tt := [][]string{
		{"app", "gen", "tmp"},
		{"go", "run", "tmp/main.go"},
	}
	for _, t := range tt {
		cmd := exec.Cmd{Args: t}
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Run()
  }
}

func clean() {
  aa := []string{"trash", "tmp"}
  cmd := exec.Cmd{Args: aa}
  cmd.Run()
}

