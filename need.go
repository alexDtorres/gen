package gen

import (
	"log"
	"os"
	s "strings"
)

type Meta interface {
	Eval(*log.Logger) string
	Apply(string, *log.Logger) AppErr
}

type Checker interface {
	Check()
}

type Need struct {
	N   int
	Log *log.Logger
}

func (n Need) Check() {
	proc := n.Log
	if N := len(os.Args); N < n.N {
		// prims
		form := "arg"

		// ops
		if N != 1 {
			// router
			chars := s.Split(form, "")
			// interactor
			chars = append(chars, "s")
			// builder
			form = s.Join(chars, "")
		}

		// combine
		form = "\n" + "got %v " + form + " ; need at least %v" + "\n"

		// abstract
		proc.Printf(form, N, n.N)
		proc.Fatalln("usage:", "gen", "app", "<name>")
	}
}
