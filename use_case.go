package gen

type Deps struct {
	Apply func(string) AppErr
	Eval  func() string
	Pre   Checker

	// Apply func(string, *log.Logger) AppErr
	// Eval  func(*log.Logger) string
}

func Install() Deps {
	d := Deps{}
	n := 3

	d.Apply = Apply
	d.Eval = Eval
	d.Pre = Need{
		N:   n,
	}

	return d
}
