package cli

import (
	"flag"
	"fmt"
	"os"
)

// register program options to the flag pkg, parse them, return arguments struct
func ParseCli() *Arguments {
	optMap := make(map[string]*bool)
	argMap := make(map[string]*string)

	for _, f := range OPTIONS {
		optMap[f.Name] = flag.Bool(f.Name, f.Default, f.Description)
	}

	for _, f := range ARGS {
		switch any(f.Default).(type) {
		case string:
			argMap[f.Name] = flag.String(f.Name, f.Default.(string), f.Description)
		}
	}

	flag.Parse()

	return &Arguments{
		Files: flag.Args(),
		Args:  argMap,
		Flags: optMap,
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
	fmt.Printf("\t%-20s\t%-10s\t%-20s\t%-20s\n\n", "Name", "Default", "Requires", "Description")
	for _, v := range OPTIONS {
		fmt.Printf("\t--%-20s\t%-10t\t%-20s\t%-20s\n", v.Name, v.Default, v.Requires, v.Description)
	}
	for _, v := range ARGS {
		fmt.Printf("\t--%-20s\t%-10v\t%-20s\t%-20s\n", v.Name, v.Default, v.Requires, v.Description)
	}

	fmt.Println("\nOnline documentation: https://github.com/xnacly/fleck")
}

func PrintVersion(version, buildAt, buildBy string) {
	fmt.Printf("fleck: [ver='%s'][buildAt='%s'][buildBy='%s']\n", version, buildAt, buildBy)
	os.Exit(0)
}

// returns the value of an option, if option not found / set returns false,
// otherwise returns the options value
func (a *Arguments) GetFlag(name string) bool {
	v, ok := a.Flags[name]
	if !ok {
		return false
	}
	return *v
}

// returns the value of an option, if option not found / set returns false,
// otherwise returns the options value
func (a *Arguments) GetArg(name string) string {
	v, ok := a.Args[name]
	if !ok {
		return ""
	}
	return *v
}
