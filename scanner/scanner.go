package scanner

import (
	"bufio"
	"fmt"
	"html"
	"log"
	"os"
	"strings"
)

type Scanner struct {
	scan    *bufio.Scanner
	isAtEnd bool    // indicates if EOF is hit
	curLine []rune  // contains the characters of the current line
	curChar rune    // contains the character at the linePos of curLine
	linePos uint    // indicates the scanners position on curLine
	line    uint    // indicates the scanners position in the file
	tokens  []Token // holds scanned token
}

// Returns a new instance of scanner.Scanner.
// To do so, it opens the file, creates a bufio.Scanner with it, scans the first line and assigns all values in the scanner.Scanner struct accordingly
func New(fileName string) Scanner {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("couldn't open file: '" + err.Error() + "'")
	}
	scan := bufio.NewScanner(file)
	scan.Scan()

	firstLine := []rune(scan.Text())

	return Scanner{
		scan:    scan,
		curLine: firstLine,
		curChar: firstLine[0],
		line:    0,
		linePos: 0,
	}
}

// creates a scanner.Token struct with kind, position, value, line and appends it to the scanner.Scanner.tokens array
func (s *Scanner) addToken(kind uint, value string) {
	pos := s.linePos
	if len(value) != 0 {
		// correct text start position
		pos = s.linePos - uint(len(value))
		value = html.EscapeString(value)
	}
	s.tokens = append(s.tokens, Token{
		Pos:   pos,
		Kind:  kind,
		Value: value,
		Line:  s.line,
	})
}

// prints all tokens currently scanned
func (s *Scanner) PrintTokens() {
	for _, token := range s.tokens {
		PrintToken(token)
	}
}

// performs a lookup for the token.Kind and prints the token data
func PrintToken(token Token) {
	fmt.Printf("[ '%s' | %d | %d | '%s' ]\n",
		TOKEN_LOOKUP_MAP[token.Kind],
		token.Pos,
		token.Line,
		token.Value,
	)
}

// increments s.linePos by one and assigns the next char to s.curChar, if at end sets s.curChar to '\n'
func (s *Scanner) advance() {
	if s.linePos+1 >= uint(len(s.curLine)) {
		s.curChar = '\n'
		s.linePos++
		return
	}

	s.linePos++
	s.curChar = s.curLine[s.linePos]
}

// increments s.line by one,
// assigns the next line to s.curChar and assigns the next char to s.curChar,
// if the current line is empty,
// adds an EMPTYLINE token using s.addToken() and tries to iterate to the next line.
func (s *Scanner) advanceLine() {
	ok := s.scan.Scan()

	if s.scan.Err() != nil || !ok {
		s.isAtEnd = true
		return
	}

	s.curLine = []rune(s.scan.Text())
	s.line++
	s.linePos = 0
	for len(s.curLine) == 0 && ok {
		s.addToken(EMPTYLINE, "")
		ok = s.scan.Scan()
		s.curLine = []rune(s.scan.Text())
		s.line++
	}
	if !ok {
		s.isAtEnd = true
		return
	}
	s.curChar = s.curLine[s.linePos]
}

// parses the file given to the Scanner line by line
func (s *Scanner) Lex() []Token {
	for !s.isAtEnd {
		var tokenKind uint
		switch s.curChar {
		case '!':
			tokenKind = BANG
		case '#':
			tokenKind = HASH
		case '>':
			tokenKind = GREATERTHAN
		case '_':
			tokenKind = UNDERSCORE
		case '*':
			tokenKind = STAR
		case '\n':
			s.addToken(NEWLINE, "")
			s.advanceLine()
			// already added token, skip rest of the loop
			continue
		case '-':
			tokenKind = DASH
		case '[':
			tokenKind = STRAIGHTBRACEOPEN
		case ']':
			tokenKind = STRAIGHTBRACECLOSE
		case '(':
			tokenKind = PARENOPEN
		case ')':
			tokenKind = PARENCLOSE
		case '`':
			tokenKind = BACKTICK
		default:
			// PERF: always use a stringbuilder instead of concatenating strings by hand
			var res strings.Builder

			// PERF: pre growing the builder to the remaining line length makes the mem alloc more precise
			res.Grow(len(s.curLine) - int(s.linePos))
		out:
			// PERF: breaking out of the outer loop is faster than doing an ok check with a hashmap,
			// also uses less memory and should allow the compiler to create a lookup table
			for {
				switch s.curChar {
				case '\n', '!', '#', '_', '*', '-', '[', ']', '(', ')', '`', '>':
					break out
				}

				res.WriteRune(s.curChar)
				s.advance()
			}

			// skip empty texts
			if res.Len() != 0 {
				s.addToken(TEXT, res.String())
			}

			// INFO: this skips adding the text again
			continue
		}

		s.addToken(tokenKind, "")
		s.advance()
	}
	s.addToken(EOF, "")
	return s.tokens
}
