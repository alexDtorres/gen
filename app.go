package gen

type Generator interface {
	Generate() error
}

type App struct {
	Generator
}

func (a App) Generate() error {
	pkg := Install()
	pre := pkg.Pre

	pre.Check()
	if err := pkg.Apply(pkg.Eval()); err != nil {
		return err
	}
	// doLoadingAnimation()
	return nil
}

/*
func doLoadingAnimation() {
	fmt.Println("running generator...")
}
*/