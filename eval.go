package gen

import (
	"log"
	"os"
)

func Eval(fmt *log.Logger) string {
	spacing := ""
	arg := ""

	for i, expr := range os.Args {
		fmt.Println()
		if i == 0 {
			fmt.Println(spacing, "environment:")
			spacing = "\t"
		}
		if i == 1 {
			fmt.Println(spacing, "symbols:")
			spacing = spacing + "\t"
		}
		if i == 2 {
			arg = expr
		}
		fmt.Println(spacing, expr)
	}
	fmt.Println()

	return arg
}

