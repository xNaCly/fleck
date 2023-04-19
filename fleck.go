//go:build !bare

package main

import (
	"os"

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
	generator.VERSION = VERSION
	cli.ARGUMENTS = cli.ParseCli()

	if cli.GetFlag(cli.ARGUMENTS, "version") {
		cli.PrintVersion(VERSION, BUILD_AT, BUILD_BY)
	}

	if cli.GetFlag(cli.ARGUMENTS, "help") {
		cli.PrintLongHelp()
		os.Exit(0)
	}

	if len(cli.ARGUMENTS.InputFile) == 0 {
		cli.PrintShortHelp()
		logger.LError("not enough arguments, specify an input file")
	}

	core.FlagCombinationSensible()

	logger.DEBUG = cli.GetFlag(cli.ARGUMENTS, "debug")
	logger.SILENT = cli.GetFlag(cli.ARGUMENTS, "silent")

	logger.LDebug("arguments: ", cli.ARGUMENTS.String())

	fileName := cli.ARGUMENTS.InputFile

	if cli.GetFlag(cli.ARGUMENTS, "shell-macro-enabled") && cli.GetFlag(cli.ARGUMENTS, "preprocessor-enabled") {
		logger.LWarn("'shell-macro-enabled' flag specified, this can harm your operating system and make it vulnerable for attack, proceed at your own digression")
	}
	if cli.GetFlag(cli.ARGUMENTS, "live-preview") {
		logger.LError("not implemented yet, watch out for the next release!")
		// TODO:
		// core.LivePreview(fileName)
	} else if cli.GetFlag(cli.ARGUMENTS, "watch") {
		core.WatchForChanges(fileName, core.Run)
	} else {
		core.Run(fileName)
	}
}
