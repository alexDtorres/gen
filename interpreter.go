package gen

import (
	"errors"
	"fmt"
	"os"
	s "strings"
)

type Meta interface {
	Eval() string
	Apply(string) AppErr
}

type AppErr error

// func Apply(path string, sess *l.Logger) AppErr {
func Apply(path string) AppErr {
	// sess.Println()

	// An error here indicates a directory
	// does not already exist at the input path.
	_, err := os.ReadDir(path)
	if err == nil {
		// TODO: Implement overwrite action.
		return errors.New("directory already exists at provided path")
	}

	// TODO: Implement write action.
	// mkdir ?path
	if err = os.Mkdir(path, 0b1111101101); err != nil {
		return err.(AppErr)
	}

	fpath := path + "/main.go"
	if err = mkFile(fpath); err != nil {
		return err
	}

	// sess.Printf("fpath = %q", fpath)

	return nil
}

func mkFile(fpath string) AppErr {
	file, err := os.Create(fpath)
	if err != nil {
		return err.(AppErr)
	}
	defer file.Close()
	writeLines(file, fpath)
	return nil
}

func writeLines(fp *os.File, name string) {
	pkgDecl := "package main"
	fp.WriteString(pkgDecl)
	fp.WriteString("\n")
	fp.WriteString("\n")
	imports := "import \"fmt\""
	fp.WriteString(imports)
	fp.WriteString("\n")
	fp.WriteString("\n")
	fcn := "func main() {"
	fp.WriteString(fcn)
	fp.WriteString("\n")
	body := "\tfmt.Println(\"" + name + "\")"
	fp.WriteString(body)
	fp.WriteString("\n")
	fp.WriteString("}")
	fp.WriteString("\n")
}

// func Eval(fmt *log.Logger) string {
func Eval() string {
	arg := ""
	// spacing := ""
	for i, expr := range os.Args {
		/*
		fmt.Println()
		if i == 0 {
			fmt.Println(spacing, "environment:")
			spacing = "\t"
		}
		if i == 1 {
			fmt.Println(spacing, "symbols:")
			spacing = spacing + "\t"
		}
		*/
		if i == 2 {
			arg = expr
		}
		// fmt.Println(spacing, expr)
	}
	// fmt.Println()
	return arg
}
	
// 
type Checker interface {
	Check()
}

// Specifies number of arguments and provides a logger.
type Need struct {
	N int
}

// Checks for the amount of arguments specified by receiver.
func (n Need) Check() {
	if N := len(os.Args); N < n.N {
		usage()
	}
	gen := os.Args[1]
	if gen != "gen" {
		usage()
	}
}

func usage() {
	fmt.Println("Usage:", "app", "gen", "<name>")
}

// Describes a receiver of the Pluralize function.
type Pluralizer interface {
	Pluralize(string) string
}

// Specifies a number of things to be pluralized.
type Plurals struct {
	Number int
	Noun string
}

// Returns a grammatical phrase given a number of things.
func (pp Plurals) Pluralize() string {
	// prims
	form := pp.Noun
	n := pp.Number
	// ops
	if n > 1 {
		// router
		chars := s.Split(form, "")
		// interactor
		chars = append(chars, "s")
		// builder
		form = s.Join(chars, "")
	}

	// combine
	form = fmt.Sprintln(n, form)

	// abstract
	return form
}
