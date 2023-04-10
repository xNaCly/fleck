package main

import (
	"log"
	"os"

	"github.com/xnacly/fleck/cli"
	"github.com/xnacly/fleck/preprocessor"
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

	if ARGUMENTS.ShellMacroEnabled {
		log.Println("warning: 'shell-macro-enabled' flag specified, this can harm your operating system and make it vulnerable for attack, proceed at your own digression")
	}

	fileName := ARGUMENTS.InputFile

	preprocessor.Process(fileName)

	s := scanner.New(fileName + ".fleck")
	s.Lex()
}
