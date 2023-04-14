package cli

// INFO: flags pkg is weird, the file always has to be at the end of the arguments list,
// due to the fact that the flag pkg won't recognize options after an option it does not recognize
// :(

import (
	"flag"
	"fmt"
)

// register program options to the flag pkg, parse them, return arguments struct
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

// prints the help with all options available
func PrintShortHelp() {
	fmt.Println(`Usage:
    fleck [Options] file

Options:`)
	for _, v := range OPTIONS {
		fmt.Printf("\t--%-20s\t\t%s\n", v.Name, v.Description)
	}
	fmt.Println("")
}

// returns the value of an option, if option not found / set returns false,
// otherwise returns the options value
func GetFlag(a Arguments, name string) bool {
	v, ok := a.Flags[name]
	if !ok {
		return false
	}
	return *v
}
