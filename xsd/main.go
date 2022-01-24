package main

import (
	"bytes"

	"github.com/alecthomas/kong"
)

var CLI struct {
	Gen struct {
		Write bool   `short:"w" help:"Write results to file instead of printing to stdout."`
		Path  string `arg help:"Path to template XML file." type:"path"`
	} `cmd help:"Generate an XSD file from a template XML file."`
}

func main() {
	cxt := kong.Parse(&CLI)
	out := cxt.Stdout
	args := cxt.Args
	for _, arg := range args {
		buf := bytes.NewBufferString(arg)
		byt := buf.Bytes()
		out.Write(byt)
	}
}
