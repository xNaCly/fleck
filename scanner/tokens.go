package scanner

type TokenKind uint

type Token struct {
	Pos   uint
	Kind  TokenKind
	Line  uint
	Value string
}

const (
	HASH = iota + 1
	UNDERSCORE
	STAR
	NEWLINE
	DASH
	STRAIGHTBRACEOPEN
	STRAIGHTBRACECLOSE
	PARENOPEN
	PARENCLOSE
	BACKTICK
	TEXT
	EMPTYLINE
)

var TOKEN_LOOKUP_MAP = map[TokenKind]string{
	HASH:               "HASH",
	UNDERSCORE:         "UNDERSCORE",
	STAR:               "STAR",
	NEWLINE:            "NEWLINE",
	DASH:               "DASH",
	STRAIGHTBRACEOPEN:  "STRAIGHTBRACEOPEN",
	STRAIGHTBRACECLOSE: "STRAIGHTBRACECLOSE",
	PARENOPEN:          "PARENOPEN",
	PARENCLOSE:         "PARENCLOSE",
	BACKTICK:           "BACKTICK",
	TEXT:               "TEXT",
}
