package scanner

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
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
func NewScanner(fileName string) Scanner {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("couldn't open file", err)
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
	}
	s.tokens = append(s.tokens, Token{
		Pos:   pos,
		Kind:  kind,
		Value: value,
		Line:  s.line,
	})
}

// getter for scanner.Scanner.tokens
func (s *Scanner) Tokens() []Token {
	return s.tokens
}

// performs a lookup for the Token.Kind value for each token and prints the token values
func (s *Scanner) PrintTokens() {
	for _, token := range s.tokens {
		fmt.Printf("[ '%s' | %d | %d | '%s' ]\n",
			TOKEN_LOOKUP_MAP[token.Kind],
			token.Pos,
			token.Line,
			token.Value,
		)
	}
}

// increments s.linePos by one and assigns the next char to s.curChar
func (s *Scanner) advance() {
	if s.linePos+1 >= uint(len(s.curLine)) {
		s.curChar = '\n'
		s.linePos++
		return
	}

	s.linePos++
	s.curChar = s.curLine[s.linePos]
}

// increments s.line by one, assigns the next line to s.curChar and assigns the next char to s.curChar
func (s *Scanner) advanceLine() {
	ok := s.scan.Scan()

	if s.scan.Err() != nil || !ok {
		s.isAtEnd = true
		return
	}

	s.curLine = []rune(s.scan.Text())
	s.line++
	s.linePos = 0
	for len(s.curLine) == 0 {
		s.scan.Scan()
		s.curLine = []rune(s.scan.Text())
		s.line++
	}
	s.curChar = s.curLine[s.linePos]
}

// parses the file given to the Scanner line by line
func (s *Scanner) Parse() {
	startTime := time.Now()
	for !s.isAtEnd {
		var tokenKind uint
		var tokenVal string
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
			// PERF:
			// this is slow due to the rune -> string conversion and the check if the string made up of the rune contains the chars, also a constant string in a while loop
			// 1.4ms, 1.2k lines, 4.5k token

			// for !strings.ContainsAny(string(s.curChar), "\n#_*-[]()`>") {
			// 	res.WriteRune(s.curChar)
			// 	s.advance()
			// }

			// PERF: option no1:
			// 1.5ms, 1.2k lines, 4.5k token

			// for {
			// 	var isSpecial bool
			// 	for _, c := range SPECIAL_CHARS {
			// 		if c == s.curChar {
			// 			isSpecial = true
			// 			break
			// 		}
			// 	}

			// 	if isSpecial {
			// 		break
			// 	}

			// 	res.WriteRune(s.curChar)
			// 	s.advance()
			// }

			// PERF: option no2:
			// 1.3ms, 1.2k lines, 4.5k token

			// for {
			// 	if _, ok := SPECIAL_CHARS_MAP[s.curChar]; ok {
			// 		break
			// 	}

			// 	res.WriteRune(s.curChar)
			// 	s.advance()
			// }

			// PERF: option no3:
			// 1.0ms, 1.2k lines, 4.5k token
			var res strings.Builder
		out:
			for {
				switch s.curChar {
				case '\n', '!', '#', '_', '*', '-', '[', ']', '(', ')', '`', '>':
					break out
				}

				res.WriteRune(s.curChar)
				s.advance()
			}

			s.addToken(TEXT, res.String())

			// PERF: performing this here instead of in the next loop interation decreases execution time by around 0.1ms
			if s.curChar == '\n' {
				s.addToken(NEWLINE, "")
				s.advanceLine()
			}

			continue
		}

		// PERF: instead of adding a token in each switch, do so here, this decreases runtime to around 1.0-1.15ms
		s.addToken(tokenKind, tokenVal)
		s.advance()
	}
	log.Printf("lexed %d token, took %s\n", len(s.tokens), time.Since(startTime).String())
}
