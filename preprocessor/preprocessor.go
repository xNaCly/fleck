package preprocessor

// possible errors:
// - empty @include => warning
// - @include cycle => error and abort
// - empty @today => warning
// - empty @shell => warning

import (
	"bufio"
	"log"
	"os"
	"time"
)

func Process(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln("couldn't open file: '" + err.Error() + "'")
	}
	s := bufio.NewScanner(f)

	ok := s.Scan()
	for ok {
		s.Text()
		// process the line here, replace found macros
		ok = s.Scan()
	}

	// write to filename+".fleck" here
}

func todayMacro(format string) string {
	return time.Now().Format(format)
}
