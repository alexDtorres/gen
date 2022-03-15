package gen_test

import (
	"testing"

	"github.com/alexDtorres/gen"
)

func TestAppGenerate(t *testing.T) {
	app := gen.App{}
	if e := app.Generate(); e != nil {
		panic(e)
	}
	t.Log("get here")
}

/*
func appGenTrialA() {
  tt := [][]string{
		{"app", "gen", "tmp"},
		{"go", "run", "tmp/main.go"},
	}
	for i, t := range tt {
		cmd := exec.Cmd{
			Path: t[0],
			Args: t[1:],
		}
		if i != 0 {
			cmd.Stdin = os.Stdin
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			cmd.Run()
		} else {
			cmd.Run()
		}
  }
}
*/

/*
func clean() {
  aa := []string{"trash", "tmp"}
  cmd := exec.Cmd{Args: aa}
  cmd.Run()
}
*/
