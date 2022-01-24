package gen

import (
	"errors"
	l "log"
	"os"
)

type AppErr error

func Apply(path string, sess *l.Logger) AppErr {
	sess.Println()

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

	sess.Printf("fpath = %q", fpath)

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
