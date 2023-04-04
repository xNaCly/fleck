package scanner

import (
	"testing"
)

func BenchmarkReadme(b *testing.B) {
	s := NewScanner("../test.md")
	s.Parse()
}
