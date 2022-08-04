package main

import (
	"flag"
	"os"
)

var omitNewLine = flag.Bool("n", false, "Não emitir o caractere de nova linha do final")

const (
	Space   = " "
	NewLine = "\n"
)

func main() {
	flag.Parse()
	var s string = ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += Space
		}

		s += flag.Arg(i)
	}

	if !*omitNewLine {
		s += NewLine
	}

	os.Stdout.WriteString(s)
}
