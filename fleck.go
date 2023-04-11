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

	if cli.GetFlag(ARGUMENTS, "shell-macro-enabled") {
		log.Println("warning: 'shell-macro-enabled' flag specified, this can harm your operating system and make it vulnerable for attack, proceed at your own digression")
	}

	fileName := ARGUMENTS.InputFile

	if cli.GetFlag(ARGUMENTS, "preprocessor-enabled") {
		preprocessor.Process(fileName)
		fileName = fileName + ".fleck"
	}

	s := scanner.New(fileName)
	s.Lex()
}
