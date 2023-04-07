package scanner

type Token struct {
	Pos   uint
	Kind  uint
	Line  uint
	Value string
}

const (
	TEXT = iota + 1
	HASH
	UNDERSCORE
	STAR
	NEWLINE
	DASH
	STRAIGHTBRACEOPEN
	STRAIGHTBRACECLOSE
	PARENOPEN
	PARENCLOSE
	BACKTICK
	GREATERTHAN
	BANG
	QUESTIONMARK
	INCLUDE
)

var TOKEN_LOOKUP_MAP = map[uint]string{
	HASH:               "HASH",
	UNDERSCORE:         "UNDERSCORE",
	STAR:               "STAR",
	NEWLINE:            "NEWLINE",
	DASH:               "DASH",
	STRAIGHTBRACEOPEN:  "STRAIGHTBRACEOPEN",
	STRAIGHTBRACECLOSE: "STRAIGHTBRACECLOSE",
	PARENOPEN:          "PARENOPEN",
	PARENCLOSE:         "PARENCLOSE",
	GREATERTHAN:        "GREATERTHAN",
	BACKTICK:           "BACKTICK",
	QUESTIONMARK:       "QUESTIONMARK",
	INCLUDE:            "INCLUDE",
	TEXT:               "TEXT",
	BANG:               "BANG",
}
