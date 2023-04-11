package preprocessor

import (
	"bufio"
	"os"
	"os/exec"
	"strings"
	"time"
	"unicode"

	"github.com/xnacly/fleck/cli"
	"github.com/xnacly/fleck/logger"
)

// used for extracting macros from strings
// loops trough the character in the line, finds everything after the @ and until the }
func findMacroMatches(line string) (macroName, macroArgument, combined string, foundSomething bool) {
	foundMacro := strings.Builder{}
	for i := 0; i < len(line); i++ {
		if i+1 >= len(line) {
			break
		}
		if line[i] == '@' && !unicode.IsSpace(rune(line[i+1])) {
			i++
			for len(line) > i && line[i] != '}' {
				foundMacro.WriteByte(line[i])
				i++
			}
		}
	}
	str := foundMacro.String()
	macro := strings.Split(str, "{")
	if foundMacro.Len() == 0 || len(macro) < 2 {
		return "", "", "", false
	}
	macroName = macro[0]
	macroArgument = macro[1]
	combined = "@" + str + "}"
	foundSomething = true
	return
}

// processes the file, replaced and expands macros
func Process(a cli.Arguments, filename string) {
	start := time.Now()
	in, err := os.Open(filename)
	if err != nil {
		logger.LError("couldn't open file: '" + err.Error() + "'")
	}

	out, err := os.Create(filename + ".fleck")
	if err != nil {
		logger.LError("couldn't open file: '" + err.Error() + "'")
	}

	defer func() {
		in.Close()
		out.Close()
		logger.LInfo("preprocessor finished, took " + time.Since(start).String())
	}()

	sIn := bufio.NewScanner(in)
	sOut := bufio.NewWriter(out)

	lines := make([]string, 0)

	ok := sIn.Scan()
	for ok {
		lines = append(lines, sIn.Text())
		ok = sIn.Scan()
	}
	for _, line := range lines {
		if len(line) != 0 {
			name, arg, combined, ok := findMacroMatches(line)
			if ok && len(arg) != 0 {
				switch name {
				case "today":
					line = strings.ReplaceAll(line, combined, time.Now().Format(arg))
				case "include":
					// nested file includes are not supported, maybe in the future? idk
					f, err := os.ReadFile(arg)
					if err != nil {
						logger.LError("Couldn't open '" + arg + "'" + err.Error())
					}
					sOut.WriteString(string(f))
					continue
				case "shell":
					if cli.GetFlag(a, "shell-macro-enabled") {
						logger.LWarn("found @shell macro: '" + combined + "', executing '" + arg + "'")
						out, err := exec.Command(arg).Output()
						if err != nil {
							logger.LWarn("failed to execute: '" + arg + "' " + err.Error())
							out = []byte(err.Error())
						}
						logger.LInfo("executed '" + arg + "' command, result:")
						line = strings.ReplaceAll(line, combined, string(out))
					} else {
						logger.LInfo("found '@shell' macro, but shell macros are disabled, use '--shell-macro-enabled' to enable this macro")
					}
				}
			}
		}
		sOut.WriteString(line + "\n")
	}
	sOut.Flush()
}
