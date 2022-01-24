package main

import "github.com/alexDtorres/gen"

// app gen ?name --basic
func main() {
	pkg := gen.Install()
	Apply := pkg.Apply
	Eval := pkg.Eval
	pre := pkg.Pre
	sess := pkg.Sess

	pre.Check()

	name := Eval(sess)
	if err := Apply(name, sess); err != nil {
		sess.Fatal(err)
	}
}

// initial = `ls`
// final   = concatenate(initial, `ls | grep ${name}`)
