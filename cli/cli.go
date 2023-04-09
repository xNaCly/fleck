package cli

import (
	"flag"
	"fmt"
)

type Arguments struct {
	ShellMacroEnabled bool
	InputFile         string
}

func ParseCli() Arguments {
	shellMacroEnabled := flag.Bool("shell-macro-enabled", false, "enables the @shell{command} macro")

	flag.Parse()
	inputFile := flag.Arg(0)

	return Arguments{
		InputFile:         inputFile,
		ShellMacroEnabled: *shellMacroEnabled,
	}
}

func PrintShortHelp() {
	fmt.Println(`
Usage:
    fleck [OPTIONS] file
    `)
}
