package main

import "github.com/alexDtorres/gen"

// app gen <your_app>
func main() {
	pkg := gen.Install()
	pre := pkg.Pre

	pre.Check()
	if err := pkg.Apply(pkg.Eval()); err != nil {
		panic(err)
	}
}

// initial = `ls`
// final   = concatenate(initial, `ls | grep ${name}`)
