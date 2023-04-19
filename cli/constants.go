package cli

import (
	"fmt"
	"strings"
)

var ARGUMENTS Arguments

type Flag struct {
	Name        string
	Default     bool
	Description string
	Requires    string // other flag this flag requires
}

type Arguments struct {
	Flags     map[string]*bool
	InputFile string
}

func (a *Arguments) String() string {
	b := strings.Builder{}
	b.WriteString("\n{\n\tInputFile: '" + a.InputFile + "', \n\tFlags: [")
	for k, v := range a.Flags {
		b.WriteString(fmt.Sprintf("\n\t\t--%s: '%t', ", k, *v))
	}
	b.WriteString("\n\t]\n}")
	return b.String()
}

var OPTIONS []Flag = []Flag{
	{
		"help",
		false,
		"prints the help page, exists",
		"",
	},
	{
		"watch",
		false,
		"watches for changes in the specified page, recompiles the file on change",
		"",
	},
	{
		"syntax",
		false,
		"enables syntax highlighting for code blocks in the output",
		"",
	},
	{
		"live-preview",
		false,
		"same as watch, spawns a default browser tab, reloads html if change occured",
		"",
	},
	{
		"debug",
		false,
		"enables debug logs",
		"",
	},
	{
		"version",
		false,
		"prints version and build information, exists",
		"",
	},
	{
		"no-prefix",
		false,
		"hide the informational comments generated in the output html",
		"",
	},
	{
		"no-template",
		false,
		"write html output to a file without the default html scaffolding",
		"",
	},
	{
		"silent",
		false,
		"disables info logs, reduces output significantly",
		"",
	},
	{
		"toc",
		false,
		"generates a table of contents at the top of the output file, includes headings 1,2,3",
		"",
	},
	{
		"toc-full",
		false,
		"generates a full toc, includes headings 1,2,3,4,5,6",
		"toc",
	},
	{
		"keep-temp",
		false,
		"keeps fleck from removing temporary files, used for debug purposes",
		"preprocessor-enabled",
	},
	{
		"preprocessor-enabled",
		false,
		"enables the preprocessor to replace macros, decreases performance",
		"",
	},
	{
		"shell-macro-enabled",
		false,
		"enables the dangerous '@shell{command}' macro, which allows the preprocessor to run any command on your system",
		"preprocessor-enabled",
	},
}
