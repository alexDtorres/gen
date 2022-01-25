package gen

import "os/exec"

func ExampleAppGenerator() {
	rss := [][]string{
		{"go", "run", "./app", "gen", "testdata"},
		{"go", "run", "./testdata"},
	}
	for _, rs := range rss {
		cmd := exec.Cmd{Args: rs}
		cmd.Run()
	}
	// Output: testdata/main.go
}

