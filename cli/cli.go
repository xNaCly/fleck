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

var OPTIONS []Flag = []Flag{
	{
		"shell-macro-enabled",
		false,
		"enables the dangerous '@shell{command}' macro, which allows the preprocessor to run any command on your system",
	},
	{
		"preprocessor-enabled",
		false,
		"enables the preprocessor to replace macros, decreases performance",
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
	fmt.Println(`
Usage:
    fleck [OPTIONS] file

Options:`)
	for _, v := range OPTIONS {
		fmt.Printf("\t--%s: %s\n", v.Name, v.Description)
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
