package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/xnacly/fleck/cli"
	"github.com/xnacly/fleck/logger"
	"github.com/xnacly/fleck/parser"
	"github.com/xnacly/fleck/preprocessor"
	"github.com/xnacly/fleck/scanner"
)

func main() {
	start := time.Now()
	cli.ARGUMENTS = cli.ParseCli()
	if len(cli.ARGUMENTS.InputFile) == 0 {
		cli.PrintShortHelp()
		logger.LError("not enough arguments, specify an input file")
	}

	fileName := cli.ARGUMENTS.InputFile

	if cli.GetFlag(cli.ARGUMENTS, "shell-macro-enabled") && cli.GetFlag(cli.ARGUMENTS, "preprocessor-enabled") {
		logger.LWarn("'shell-macro-enabled' flag specified, this can harm your operating system and make it vulnerable for attack, proceed at your own digression")
	}

	if cli.GetFlag(cli.ARGUMENTS, "preprocessor-enabled") {
		logger.LInfo("preprocessor enabled, starting...")
		preprocessor.Process(cli.ARGUMENTS, fileName)
		fileName = fileName + ".fleck"
	}

	lexerStart := time.Now()
	s := scanner.New(fileName)
	tokens := s.Lex()
	logger.LInfo("lexed " + fmt.Sprint(len(tokens)) + " token, took " + time.Since(lexerStart).String())

	parserStart := time.Now()
	p := parser.New(tokens)
	result := p.Parse()
	logger.LInfo("parsed " + fmt.Sprint(len(result)) + " items, took " + time.Since(parserStart).String())

	writeStart := time.Now()
	name := strings.Split(fileName, ".")[0] + ".html"
	out, err := os.Create(name)
	writer := bufio.NewWriter(out)

	if err != nil {
		logger.LError("failed to open file: " + err.Error())
	}

	if cli.GetFlag(cli.ARGUMENTS, "toc") {
		writer.WriteString(p.GenerateToc())
	}

	for _, e := range result {
		writer.WriteString(e.String() + "\n")
	}

	writer.Flush()
	logger.LInfo("wrote generated html to '" + name + "', took: " + time.Since(writeStart).String())

	defer func() {
		if cli.GetFlag(cli.ARGUMENTS, "preprocessor-enabled") {
			if cli.GetFlag(cli.ARGUMENTS, "keep-temp") {
				return
			}
			logger.LInfo("cleanup, removing: '" + fileName + "'")
			err := os.Remove(fileName)
			if err != nil {
				logger.LWarn(err.Error())
			}
		}
		logger.LInfo("did everything, took: " + time.Since(start).String())
	}()
}
