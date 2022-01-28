package main

import "os/exec"

func ExampleRunner() {
	rss := [][]string{
		{"go", "run", "drivers/app/main.go", "gen", "testdata"},
		{"go", "run", "testdata/main.go"},
	}
	for _, rs := range rss {
		cmd := exec.Cmd{Args: rs}
		cmd.Run()
	}
	// Output: testdata/main.go
}

