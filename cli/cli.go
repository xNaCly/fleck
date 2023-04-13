package cli

import (
	"flag"
	"fmt"
)

func ParseCli() Arguments {
	resMap := make(map[string]*bool)
	for _, f := range OPTIONS {
		resMap[f.Name] = flag.Bool(f.Name, f.Default, f.Description)
	}

	flag.Parse()
	inputFile := flag.Arg(0)

	return Arguments{
		InputFile: inputFile,
		Flags:     resMap,
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
	v, ok := a.Flags[name]
	if !ok {
		return false
	}
	return *v
}
