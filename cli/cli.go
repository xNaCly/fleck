package cli

// INFO: flags pkg is weird, the file always has to be at the end of the arguments list,
// due to the fact that the flag pkg won't recognize options after an option it does not recognize
// :(

import (
	"flag"
	"fmt"
	"os"
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

// prints a help page with only the usage
func PrintShortHelp() {
	fmt.Println(`Usage:
    fleck [Options] file

    Run 'fleck --help' for an in depth help page
    `)
}

// prints the help with all options available
func PrintLongHelp() {
	fmt.Println(`Usage:
    fleck [Options] file

Options:`)
	for _, v := range OPTIONS {
		fmt.Printf("\t--%-20s\t\t%s\n", v.Name, v.Description)
	}
	fmt.Println("\nOnline documentation: https://github.com/xnacly/fleck")
}

func PrintVersion(version, buildAt, buildBy string) {
	fmt.Printf("fleck: [ver='%s'][buildAt='%s'][buildBy='%s']\n", version, buildAt, buildBy)
	os.Exit(0)
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
