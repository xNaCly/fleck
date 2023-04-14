//go:build bare

package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"

	"github.com/xnacly/fleck/parser"
	"github.com/xnacly/fleck/scanner"
)

func main() {
	start := time.Now()
	log.Println("fleck - barebones")
	args := os.Args

	if len(args) < 2 {
		log.Fatalln("not enough arguments, specify an input file")
	}

	fileName := args[1]

	s := scanner.New(fileName)
	tokens := s.Lex()

	p := parser.New(tokens)
	result := p.Parse()

	name := strings.Split(fileName, ".")[0] + ".html"
	out, err := os.Create(name)
	writer := bufio.NewWriter(out)

	if err != nil {
		log.Fatalln("failed to open file: " + err.Error())
	}

	for _, e := range result {
		writer.WriteString(e.String())
	}

	writer.Flush()

	log.Println("done, took: ", time.Since(start))
}
