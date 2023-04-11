package parser

import (
	"fmt"
	"strings"

	"github.com/xnacly/fleck/logger"
	"github.com/xnacly/fleck/scanner"
)

type Parser struct {
	tokens  []scanner.Token
	tags    []Tag
	current int
}

func New(tokens []scanner.Token) Parser {
	return Parser{
		tokens: tokens,
		tags:   []Tag{},
	}
}

func (p *Parser) Parse() []Tag {
	for !p.isAtEnd() {
		tag := p.tag()
		if tag != nil {
			p.tags = append(p.tags, tag)
		}
	}
	return p.tags
}

func (p *Parser) tag() Tag {
	// parse headings just before paragraphs at the end, everything else before
	if p.peek().Kind == scanner.HASH && (p.prev().Kind == scanner.NEWLINE || p.prev().Kind == 0) {
		return p.heading()
	} else {
		// TODO: currently skips everything except headings, fix that, this here is a temp messuare to keep the program from endless looping
		p.advance()
		return nil
	}
}

func (p *Parser) paragraph() Tag {
	return Paragraph{}
}

func (p *Parser) heading() Tag {
	var lvl uint = 0
	children := make([]scanner.Token, 0)

	for p.peek().Kind == scanner.HASH {
		lvl++
		p.advance()
	}

	for !p.check(scanner.NEWLINE) {
		children = append(children, p.peek())
		p.advance()
	}

	p.advance()

	b := strings.Builder{}

	if lvl < 7 {
		// too many levels down
		for _, c := range children {
			if c.Kind == scanner.TEXT {
				b.WriteString(c.Value)
				continue
			}
			b.WriteRune(scanner.TOKEN_SYMBOL_MAP[c.Kind])
		}
	}

	return Heading{
		lvl:  lvl,
		text: b.String(),
	}
}

func (p *Parser) match(types ...uint) bool {
	for _, t := range types {
		if p.check(t) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) consume(kind uint, msg string) {
	if p.check(kind) {
		p.advance()
		return
	}
	p.error(fmt.Sprintf("expected '%s', got '%s'", scanner.TOKEN_LOOKUP_MAP[kind], scanner.TOKEN_LOOKUP_MAP[p.peek().Kind]))
}

func (p *Parser) check(kind uint) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Kind == kind
}

func (p *Parser) advance() {
	if !p.isAtEnd() {
		p.current++
	}
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Kind == scanner.EOF
}

func (p *Parser) peek() scanner.Token {
	return p.tokens[p.current]
}

func (p *Parser) prev() scanner.Token {
	if p.current == 0 {
		return scanner.Token{
			Kind: 0,
		}
	}
	return p.tokens[p.current-1]
}

func (p *Parser) error(msg string) {
	t := p.peek()
	logger.LError(fmt.Sprintf("line: %d, pos: %d, Error at: %s: %s", t.Line, t.Pos, t.Value, msg))
}
