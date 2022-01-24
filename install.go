package gen

import "log"

type Deps struct {
	Sess  *log.Logger
	Pre   Checker
	Eval  func(*log.Logger) string
	Apply func(string, *log.Logger) AppErr
}

func Install() Deps {
	d := Deps{}
	l := log.Default()
	n := 3

	d.Apply = Apply
	d.Eval = Eval
	d.Pre = Need{
		N:   n,
		Log: l,
	}
	d.Sess = l

	return d
}
