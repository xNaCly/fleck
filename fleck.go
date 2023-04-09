package main

import (
	"log"
	"os"

	"github.com/xnacly/fleck/cli"
	"github.com/xnacly/fleck/scanner"
)

var ARGUMENTS cli.Arguments

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("not enough arguments, specify a markdown file")
	}

	ARGUMENTS = cli.ParseCli()
	log.Println(ARGUMENTS)
	s := scanner.New(os.Args[1])
	s.Lex()
}
