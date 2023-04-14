package main

import (
	"fmt"
	"os"
	"time"

	"github.com/xnacly/fleck/cli"
	"github.com/xnacly/fleck/generator"
	"github.com/xnacly/fleck/logger"
	"github.com/xnacly/fleck/parser"
	"github.com/xnacly/fleck/preprocessor"
	"github.com/xnacly/fleck/scanner"
)

// alerts the user if a flag depends on a different flag to have an effect
func flagCombinationSensible() {
	for _, f := range cli.OPTIONS {
		if len(f.Requires) == 0 {
			continue
		}
		if cli.GetFlag(cli.ARGUMENTS, f.Name) && !cli.GetFlag(cli.ARGUMENTS, f.Requires) {
			logger.LWarn(fmt.Sprintf("flag '--%s' requires flag '--%s' to be set, otherwise it has no effect.", f.Name, f.Requires))
		}
	}
}

// TODO: add a watcher mode
// TODO: only rebuild if the file changed, md5 hash?
func main() {
	start := time.Now()

	cli.ARGUMENTS = cli.ParseCli()
	if len(cli.ARGUMENTS.InputFile) == 0 {
		cli.PrintShortHelp()
		logger.LError("not enough arguments, specify an input file")
	}

	flagCombinationSensible()

	logger.DEBUG = cli.GetFlag(cli.ARGUMENTS, "debug")
	logger.SILENT = cli.GetFlag(cli.ARGUMENTS, "silent")

	logger.LDebug("arguments: ", cli.ARGUMENTS.String())

	fileName := cli.ARGUMENTS.InputFile

	if cli.GetFlag(cli.ARGUMENTS, "shell-macro-enabled") && cli.GetFlag(cli.ARGUMENTS, "preprocessor-enabled") {
		logger.LWarn("'shell-macro-enabled' flag specified, this can harm your operating system and make it vulnerable for attack, proceed at your own digression")
	}

	if cli.GetFlag(cli.ARGUMENTS, "preprocessor-enabled") {
		logger.LInfo("preprocessor enabled, starting...")
		preprocessor.Process(cli.ARGUMENTS, fileName)
		fileName = fileName + ".fleck"
	}

	logger.LDebug("starting scanner")
	lexerStart := time.Now()
	s := scanner.New(fileName)
	tokens := s.Lex()
	logger.LDebug("lexed " + fmt.Sprint(len(tokens)) + " token, took " + time.Since(lexerStart).String())
	if logger.DEBUG {
		s.PrintTokens()
	}

	logger.LDebug("starting parser")
	parserStart := time.Now()
	p := parser.New(tokens)
	result := p.Parse()
	logger.LDebug("parsed " + fmt.Sprint(len(result)) + " items, took " + time.Since(parserStart).String())
	logger.LDebug("parsed tags:", result)

	var toc string
	if cli.GetFlag(cli.ARGUMENTS, "toc") {
		logger.LDebug("generating toc...")
		toc = p.GenerateToc()
	}

	if cli.GetFlag(cli.ARGUMENTS, "no-template") {
		generator.WritePlain(fileName, result, toc)
	} else {
		generator.WriteTemplate(fileName, result, toc)
	}

	logger.LDebug("did everything, took: " + time.Since(start).String())

	defer func() {
		if cli.GetFlag(cli.ARGUMENTS, "preprocessor-enabled") {
			if cli.GetFlag(cli.ARGUMENTS, "keep-temp") {
				return
			}
			logger.LDebug("cleanup, removing: '" + fileName + "'")
			err := os.Remove(fileName)
			if err != nil {
				logger.LWarn(err.Error())
			}
		}
	}()
}
