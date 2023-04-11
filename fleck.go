package main

import (
	"os"

	"github.com/xnacly/fleck/cli"
	"github.com/xnacly/fleck/logger"
	"github.com/xnacly/fleck/preprocessor"
	"github.com/xnacly/fleck/scanner"
)

func main() {
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

	s := scanner.New(fileName)
	s.Lex()

	defer func() {
		if cli.GetFlag(cli.ARGUMENTS, "keep-temp") {
			return
		}
		logger.LInfo("cleanup, removing: '" + fileName + "'")
		err := os.Remove(fileName)
		if err != nil {
			logger.LWarn(err.Error())
		}
	}()
}
