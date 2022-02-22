package gen

import "os/exec"

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
		cmd.Run()
  }
}

func clean() {
  aa := []string{"rm", "-rf", "tmp"}
  cmd := exec.Cmd{Args: aa}
  cmd.Run()
}

