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
	DOLLAR
	STRAIGHTBRACEOPEN
	STRAIGHTBRACECLOSE
	PARENOPEN
	PARENCLOSE
	BACKTICK
	GREATERTHAN
	BANG
	EMPTYLINE
	EOF
)

var TOKEN_LOOKUP_MAP = map[uint]string{
	DOLLAR:             "DOLLAR",
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
	BANG:               "BANG",
	EMPTYLINE:          "EMPTYLINE",
	EOF:                "EOF",
}

var TOKEN_SYMBOL_MAP = map[uint]rune{
	HASH:               '#',
	DOLLAR:             '$',
	UNDERSCORE:         '_',
	STAR:               '*',
	NEWLINE:            '\n',
	DASH:               '-',
	STRAIGHTBRACEOPEN:  '[',
	STRAIGHTBRACECLOSE: ']',
	PARENOPEN:          '(',
	PARENCLOSE:         ')',
	GREATERTHAN:        '>',
	BACKTICK:           '`',
	BANG:               '!',
	EMPTYLINE:          '\n',
	EOF:                0,
}
