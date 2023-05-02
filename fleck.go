//go:build !bare

package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/xnacly/fleck/cli"
	"github.com/xnacly/fleck/core"
	"github.com/xnacly/fleck/generator"
	"github.com/xnacly/fleck/logger"
)

// supplied by the build process
var VERSION = ""
var BUILD_AT = ""
var BUILD_BY = ""

// TODO: only rebuild if the file changed, md5 hash?
func main() {
	log.SetOutput(os.Stdout)
	generator.VERSION = VERSION
	cli.ARGUMENTS = *cli.ParseCli()

	if cli.ARGUMENTS.GetFlag("version") {
		cli.PrintVersion(VERSION, BUILD_AT, BUILD_BY)
	}

	if cli.ARGUMENTS.GetFlag("help") {
		cli.PrintLongHelp()
		os.Exit(0)
	}

	if len(cli.ARGUMENTS.Files) == 0 {
		cli.PrintShortHelp()
		logger.LError("not enough arguments, specify an input file")
	}

	core.FlagCombinationSensible()

	logger.DEBUG = cli.ARGUMENTS.GetFlag("debug")
	logger.SILENT = cli.ARGUMENTS.GetFlag("silent")

	if logger.DEBUG {
		fmt.Println(cli.ARGUMENTS.String())
	}

	logger.LDebug("arguments: ", cli.ARGUMENTS.String())

	var wg sync.WaitGroup
	for _, file := range cli.ARGUMENTS.Files {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			if cli.ARGUMENTS.GetFlag("shell-macro-enabled") && cli.ARGUMENTS.GetFlag("preprocessor-enabled") {
				logger.LWarn("'shell-macro-enabled' flag specified, this can harm your operating system and make it vulnerable for attack, proceed at your own digression")
			}

			s, err := os.Stat(file)

			if err != nil {
				logger.LError("failed to stat the file")
			}

			if s.Size() == 0 {
				logger.LWarn("file is empty, exiting.")
				os.Exit(0)
			}

			if cli.ARGUMENTS.GetFlag("live-preview") {
				// TODO:
				if len(cli.ARGUMENTS.Files) > 1 {
					logger.LError("live preview currently only supported for a single file")
				}
				core.LivePreview(file)
			} else if cli.ARGUMENTS.GetFlag("watch") {
				core.WatchForChanges(file, core.Run)
			} else {
				core.Run(file)
			}
		}(file)
	}
	wg.Wait()
}
