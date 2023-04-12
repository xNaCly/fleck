package scanner

import (
	"testing"
)

func BenchmarkReadme(b *testing.B) {
	s := New("../README.md")
	s.Lex()
}

func TestHeadings(t *testing.T) {
	s := New("./tests/markdown.md")
	tokens := s.Lex()
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
		UNDERSCORE,
		TEXT,
		UNDERSCORE,
		TEXT,
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
		EOF,
	}

	s.PrintTokens()
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
