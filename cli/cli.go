package cli

import (
	"flag"
	"fmt"
)

type Flag struct {
	Name        string
	Default     bool
	Description string
}

type Arguments struct {
	flags     map[string]*bool
	InputFile string
}

var ARGUMENTS Arguments

var OPTIONS []Flag = []Flag{
	{
		"silent",
		false,
		"disables info logs, reduces output significantly",
	},
	{
		"toc",
		false,
		"generates a table of contents at the top of the output with links to the headings",
	},
	{
		"keep-temp",
		false,
		"keeps fleck from removing temporary files, used for debug purposes",
	},
	{
		"preprocessor-enabled",
		false,
		"enables the preprocessor to replace macros, decreases performance",
	},
	{
		"shell-macro-enabled",
		false,
		"enables the dangerous '@shell{command}' macro, which allows the preprocessor to run any command on your system",
	},
}

func ParseCli() Arguments {
	resMap := make(map[string]*bool)
	for _, f := range OPTIONS {
		resMap[f.Name] = flag.Bool(f.Name, f.Default, f.Description)
	}

	flag.Parse()
	inputFile := flag.Arg(0)

	return Arguments{
		InputFile: inputFile,
		flags:     resMap,
	}
}

func PrintShortHelp() {
	fmt.Println(`Usage:
    fleck [Options] file

Options:`)
	for _, v := range OPTIONS {
		fmt.Printf("\t--%-20s\t\t%s\n", v.Name, v.Description)
	}
	fmt.Println("")
}

func GetFlag(a Arguments, name string) bool {
	v, ok := a.flags[name]
	if !ok {
		return false
	}
	return *v
}
