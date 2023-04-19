package core

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
func FlagCombinationSensible() {
	for _, f := range cli.OPTIONS {
		if len(f.Requires) == 0 {
			continue
		}
		if cli.GetFlag(cli.ARGUMENTS, f.Name) && !cli.GetFlag(cli.ARGUMENTS, f.Requires) {
			logger.LWarn(fmt.Sprintf("flag '--%s' requires flag '--%s' to be set, otherwise it has no effect.", f.Name, f.Requires))
		}
	}
}

// watches for changes in a file, recompiles the file if a change occurs, can be exited via <C-c>
func WatchForChanges(fileName string, executor func(string)) {
	executor(fileName)
	logger.LInfo("watching for changes...")

	initialStat, err := os.Stat(fileName)
	if err != nil {
		logger.LError("failed to watch for changes: " + err.Error())
	}

	i := 0
	for {
		stat, err := os.Stat(fileName)
		if err != nil {
			logger.L("test")
			logger.LError("failed to watch for changes: " + err.Error())
		}

		if stat.Size() != initialStat.Size() || stat.ModTime() != stat.ModTime() {
			initialStat = stat
			i++
			fmt.Print(logger.ANSI_CLEAR)
			logger.LInfo("detected change, recompiling... (" + fmt.Sprint(i) + ")")
			executor(fileName)
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func Run(fileName string) {
	start := time.Now()

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

	logger.LInfo("compiled '" + fileName + "', took: " + time.Since(start).String())

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
