package cli

import (
	"flag"
)

type Arguments struct {
	ShellMacroEnabled bool
}

func ParseCli() Arguments {
	shellMacroEnabled := flag.Bool("shell-macro-enabled", false, "enables the @shell{command} macro")

	flag.Parse()

	return Arguments{
		ShellMacroEnabled: *shellMacroEnabled,
	}
}
