package parser

import (
	"fmt"
	"strings"

	"github.com/xnacly/fleck/cli"
	"github.com/xnacly/fleck/logger"
	"github.com/xnacly/fleck/scanner"
)

type Parser struct {
	tokens   []scanner.Token
	tags     []Tag
	current  int
	headings []Heading
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
	if p.check(scanner.BACKTICK) {
		return p.code()
	} else if p.check(scanner.HASH) && (p.prev().Kind == scanner.NEWLINE || p.prev().Kind == 0) {
		return p.heading()
	} else {
		// TODO: currently skips everything except headings, fix that, this here is a temp messuare to keep the program from endless looping
		p.advance()
		return nil
	}
}

func (p *Parser) code() Tag {
	p.advance()
	if p.check(scanner.TEXT) {
		// skip the `
		p.advance()
		if p.check(scanner.BACKTICK) {
			return CodeInline{
				text: p.prev().Value,
			}
		}
	} else if p.check(scanner.BACKTICK) {
		p.advance()
		if !p.check(scanner.BACKTICK) {
			return CodeInline{
				text: "",
			}
		}
		p.advance()
		language := p.peek().Value
		// skip lang definition
		p.advance()
		// skip newline
		p.advance()

		b := strings.Builder{}
		for !p.check(scanner.BACKTICK) {
			if p.check(scanner.TEXT) {
				b.WriteString(p.peek().Value)
			} else {
				b.WriteRune(scanner.TOKEN_SYMBOL_MAP[p.peek().Kind])
			}
			p.advance()
		}

		// skip the three ```
		p.advance()
		p.advance()
		p.advance()

		return CodeBlock{
			language: language,
			text:     b.String(),
		}
	}
	return nil
}

func (p *Parser) paragraph() Tag {
	return Paragraph{}
}

// TODO: find a way to parse the rest of the tokens as well, not just write them to the heading
func (p *Parser) heading() Tag {
	var lvl uint = 0
	children := make([]scanner.Token, 0)

	for p.check(scanner.HASH) {
		lvl++
		p.advance()
	}

	for !p.check(scanner.NEWLINE) {
		children = append(children, p.peek())
		p.advance()
	}

	// skip the newline
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
	heading := Heading{
		lvl:  lvl,
		text: b.String(),
	}

	if cli.GetFlag(cli.ARGUMENTS, "toc") {
		p.headings = append(p.headings, heading)
	}
	return heading
}

func (p *Parser) GenerateToc() string {
	headingMap := map[uint]uint{
		1: 0,
		2: 0,
		3: 0,
	}

	b := strings.Builder{}
	b.WriteString("<h3>Table of contents</h3>")
	b.WriteString("<ul>")
	for _, v := range p.headings {
		// TODO: make this a -toc-level=x flag
		// TODO: switch over levels, indent subheadings using <ul> in <ul>
		if v.lvl > 3 {
			continue
		}
		headingMap[v.lvl]++
		b.WriteString(fmt.Sprintf("<li><a href=\"#%s\">%d.%d.%d</a>: %s</li>", v.text, headingMap[1], headingMap[2], headingMap[3], v.text))
	}
	b.WriteString("</ul>")
	return b.String()
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
