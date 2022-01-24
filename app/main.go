package main

import (
	"log"

	"github.com/alexDtorres/gen"
)

// app gen ?name --basic
func main() {
	pkg := install()
	Apply := pkg.Apply
	Eval := pkg.Eval
	pre := pkg.Pre
	sess := pkg.Sess
	pre.Check()

	name := Eval(sess)
	sess.Println("?name = ", name)
	if err := Apply(name, sess); err != nil {
		sess.Fatal(err)
	}

	sess.Printf("want = %q \n", "")
	sess.Println()
}

// initial = `ls`
// final   = concatenate(initial, `ls | grep ${name}`)

type deps struct {
	Sess  *log.Logger
	Pre   gen.Checker
	Eval  func(*log.Logger) string
	Apply func(string, *log.Logger) gen.AppErr
}

func install() deps {
	d := deps{}
	l := log.Default()
	n := 3

	d.Apply = gen.Apply
	d.Eval = gen.Eval
	d.Pre = gen.Need{
		N:   n,
		Log: l,
	}
	d.Sess = l

	return d
}
