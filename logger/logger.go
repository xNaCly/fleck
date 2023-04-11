// copied from https://github.com/xNaCly/tiny-interpreter/blob/master/logger/logger.go
package logger

import (
	"log"

	"github.com/xnacly/fleck/cli"
)

const (
	ANSI_RESET  = "\033[0m"
	ANSI_RED    = "\033[91m"
	ANSI_YELLOW = "\033[93m"
	ANSI_BLUE   = "\033[36m"
)

// prefixes s with 'info', prints result
func LInfo(s string) {
	if !cli.GetFlag(cli.ARGUMENTS, "silent") {
		log.Printf("%sinfo%s: %s\n", ANSI_BLUE, ANSI_RESET, s)
	}
}

// prefixes s with 'warn', prints result
func LWarn(s string) {
	log.Printf("%swarn%s: %s\n", ANSI_YELLOW, ANSI_RESET, s)
}

// prefixes s with 'error', calls log.Fatalln, prints result, exits with error code 1
func LError(s string) {
	log.Fatalf("%serror%s: %s\n", ANSI_RED, ANSI_RESET, s)
}

// simple call to the log.Println function, only here to keep things isolated and consistent
func L(v ...any) {
	log.Println(v...)
}
