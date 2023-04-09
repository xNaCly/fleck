package main

import (
	"log"
	"os"

	"github.com/xnacly/fleck/cli"
	"github.com/xnacly/fleck/scanner"
)

var ARGUMENTS cli.Arguments

func main() {
	ARGUMENTS = cli.ParseCli()
	if len(ARGUMENTS.InputFile) == 0 {
		log.Println("not enough arguments, specify an input file")
		cli.PrintShortHelp()
		os.Exit(1)
	}

	s := scanner.New(ARGUMENTS.InputFile)
	s.Lex()
}
