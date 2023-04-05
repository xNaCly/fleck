package scanner

import (
	"testing"
)

func BenchmarkReadme(b *testing.B) {
	s := NewScanner("../README.md")
	s.Parse()
}

func TestHeadings(t *testing.T) {
	s := NewScanner("./tests/markdown.md")
	s.Parse()
	tokens := s.Tokens()
	expectedTokens := []uint{
		HASH,
		TEXT,
		NEWLINE,
		HASH,
		HASH,
		TEXT,
		NEWLINE,
		HASH,
		HASH,
		HASH,
		TEXT,
		NEWLINE,
		HASH,
		HASH,
		HASH,
		HASH,
		TEXT,
		NEWLINE,
		HASH,
		HASH,
		HASH,
		HASH,
		HASH,
		TEXT,
		NEWLINE,
		HASH,
		HASH,
		HASH,
		HASH,
		HASH,
		HASH,
		TEXT,
		NEWLINE,
		GREATERTHAN,
		TEXT,
		NEWLINE,
		STAR,
		STAR,
		TEXT,
		STAR,
		STAR,
		TEXT,
		UNDERSCORE,
		TEXT,
		UNDERSCORE,
		NEWLINE,
		DASH,
		DASH,
		DASH,
		NEWLINE,
		STRAIGHTBRACEOPEN,
		TEXT,
		STRAIGHTBRACECLOSE,
		PARENOPEN,
		TEXT,
		PARENCLOSE,
		NEWLINE,
		BANG,
		STRAIGHTBRACEOPEN,
		TEXT,
		STRAIGHTBRACECLOSE,
		PARENOPEN,
		TEXT,
		PARENCLOSE,
		NEWLINE,
		BACKTICK,
		TEXT,
		BACKTICK,
		NEWLINE,
	}
	if len(tokens) != len(expectedTokens) {
		t.Errorf("expected %d tokens, got: %d", len(expectedTokens), len(tokens))
	}
	for i, token := range tokens {
		if expectedTokens[i] != token.Kind {
			t.Errorf("expected %d [%s], got %d [%s] for token %d",
				expectedTokens[i], TOKEN_LOOKUP_MAP[expectedTokens[i]], token.Kind, TOKEN_LOOKUP_MAP[token.Kind], i,
			)
		}
	}
}
