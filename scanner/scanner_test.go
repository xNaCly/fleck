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
		EMPTYLINE,
		HASH,
		HASH,
		TEXT,
		NEWLINE,
		EMPTYLINE,
		HASH,
		HASH,
		HASH,
		TEXT,
		NEWLINE,
		EMPTYLINE,
		HASH,
		HASH,
		HASH,
		HASH,
		TEXT,
		NEWLINE,
		EMPTYLINE,
		HASH,
		HASH,
		HASH,
		HASH,
		HASH,
		TEXT,
		NEWLINE,
		EMPTYLINE,
		HASH,
		HASH,
		HASH,
		HASH,
		HASH,
		HASH,
		TEXT,
		NEWLINE,
		EMPTYLINE,
		GREATERTHAN,
		TEXT,
		NEWLINE,
		EMPTYLINE,
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
		EMPTYLINE,
		DASH,
		DASH,
		DASH,
		NEWLINE,
		EMPTYLINE,
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
		EMPTYLINE,
		BACKTICK,
		TEXT,
		BACKTICK,
		NEWLINE,
		EOF,
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
