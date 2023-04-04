package scanner

type Token struct {
	Pos   uint
	Kind  uint
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
	GREATERTHAN
	TEXT
	EMPTYLINE
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
	TEXT:               "TEXT",
}
