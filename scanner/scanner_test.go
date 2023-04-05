package scanner

import (
	"testing"
)

// TODO: add tests for each token:
//	- HASH
//	- UNDERSCORE
//	- STAR
//	- NEWLINE
//	- DASH
//	- STRAIGHTBRACEOPEN
//	- STRAIGHTBRACECLOSE
//	- PARENOPEN
//	- PARENCLOSE
//	- BACKTICK
//	- GREATERTHAN
//	- TEXT
//	- EMPTYLINE

func BenchmarkReadme(b *testing.B) {
	s := NewScanner("../test.md")
	s.Parse()
}
