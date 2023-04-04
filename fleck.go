package main

import (
	"log"
	"os"

	"github.com/xnacly/fleck/scanner"
)

const VERSION = "v0.1.0"

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("not enough arguments, specify a markdown file")
	}
	s := scanner.NewScanner(os.Args[1])
	s.Parse()
}
