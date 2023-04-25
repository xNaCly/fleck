package parser

import (
	"fmt"
	"strings"

	"github.com/xnacly/fleck/cli"
	"github.com/xnacly/fleck/scanner"
)

type Parser struct {
	tokens   []scanner.Token // lexed tokens (input)
	tags     []Tag           // parsed html tags (output)
	current  int             // indicates the current pos in the tokens array
	headings []Heading       // helper array for the toc, if enabled
}

// creates a new instance of Parser
func New(tokens []scanner.Token) Parser {
	return Parser{
		tokens: tokens,
		tags:   []Tag{},
	}
}

// parses the tokens passed to parser.New(), returns an array of HTML tags
func (p *Parser) Parse() []Tag {
	for !p.isAtEnd() {
		tag := p.tag()
		if tag != nil {
			p.tags = append(p.tags, tag)
		}
	}
	return p.tags
}

// dispatches to other methods, returns a TAG
func (p *Parser) tag() Tag {
	if p.check(scanner.GREATERTHAN) {
		return p.quote()
	} else if p.check(scanner.DASH) {
		return p.list()
	} else if p.check(scanner.BANG) {
		return p.img()
	} else if p.check(scanner.BACKTICK) {
		return p.code(false)
	} else if p.check(scanner.HASH) && (p.prev().Kind == scanner.EMPTYLINE || p.prev().Kind == 0) {
		return p.heading()
	} else {
		return p.paragraph()
	}
}

// parses all lists, unordered, ordered, checked
func (p *Parser) list() Tag {
	// TODO: parse checkmark unordered list

	// skip the first
	p.advance()

	if p.check(scanner.DASH) {
		p.advance()
		if p.check(scanner.DASH) {
			return Ruler{}
		}
	}

	// BUG: this is a mess, try to fix this
	children := make([]Tag, 0)
	curLine := make([]Tag, 0)

	// paragraph should only contain inline code, italic and bold or text
	for !p.check(scanner.EMPTYLINE) && !p.isAtEnd() {
		// this is the next li
		// TODO: no nesting supported, maybe implement that
		if p.check(scanner.DASH) && (p.prev().Kind == scanner.NEWLINE || p.prev().Kind == scanner.EMPTYLINE) {
			if len(curLine) != 0 {
				children = append(children, ListItem{
					children: curLine,
				})
				curLine = make([]Tag, 0)
			}
			p.advance()
		}

		switch p.peek().Kind {
		case scanner.STRAIGHTBRACEOPEN:
			curLine = append(curLine, p.link())
		case scanner.BANG:
			// INFO: p.img automatically skips the new line
			curLine = append(curLine, p.img())
		case scanner.BACKTICK:
			curLine = append(curLine, p.code(false))
		case scanner.NEWLINE:
			if len(curLine) != 0 {
				curLine = append(curLine, Br{})
			}
			p.advance()
		case scanner.STAR, scanner.UNDERSCORE:
			curLine = append(curLine, p.emphasis())
		case scanner.TEXT:
			curLine = append(curLine, Text{content: p.peek().Value})
			p.advance()
		default:
			curLine = append(curLine, Text{content: string(scanner.TOKEN_SYMBOL_MAP[p.peek().Kind])})
			p.advance()
		}
	}

	if len(curLine) != 0 {
		children = append(children, ListItem{
			children: curLine,
		})
		curLine = make([]Tag, 0)
	}

	if len(children) == 0 {
		return nil
	}

	return List{children: children}
}

// parses blockquotes
func (p *Parser) quote() Tag {
	// skip the >
	p.advance()

	children := make([]Tag, 0)

	for !p.check(scanner.EMPTYLINE) && !p.isAtEnd() {
		switch p.peek().Kind {
		case scanner.GREATERTHAN:
			p.advance()
			continue
		case scanner.BANG:
			children = append(children, p.img())
		case scanner.NEWLINE:
			if p.prev().Kind == scanner.GREATERTHAN {
				children = append(children, Br{})
			}
			p.advance()
		case scanner.HASH:
			children = append(children, p.heading())
		case scanner.STRAIGHTBRACEOPEN:
			children = append(children, p.link())
		case scanner.BACKTICK:
			children = append(children, p.code(true))
		case scanner.STAR, scanner.UNDERSCORE:
			children = append(children, p.emphasis())
		case scanner.TEXT:
			children = append(children, Text{content: p.peek().Value})
			p.advance()
		default:
			children = append(children, Text{content: string(scanner.TOKEN_SYMBOL_MAP[p.peek().Kind])})
			p.advance()
		}
	}

	return Quote{
		children: children,
	}
}

// parses images
func (p *Parser) img() Tag {
	p.advance()

	if !p.check(scanner.STRAIGHTBRACEOPEN) {
		p.advance()
		return Text{content: "!"}
	}

	p.advance()
	b := strings.Builder{}
	for !p.check(scanner.STRAIGHTBRACECLOSE) && !p.check(scanner.NEWLINE) {
		if p.check(scanner.TEXT) {
			b.WriteString(p.peek().Value)
		} else {
			b.WriteRune(scanner.TOKEN_SYMBOL_MAP[p.peek().Kind])
		}
		p.advance()
	}

	if p.check(scanner.NEWLINE) {
		p.advance()
		return Text{content: "[" + b.String()}
	}

	alt := b.String()
	b.Reset()

	// skip the [
	p.advance()

	if !p.check(scanner.PARENOPEN) {
		return Text{content: "[" + alt + "]"}
	}

	// skip the opening brace
	p.advance()

	for !p.check(scanner.PARENCLOSE) && !p.check(scanner.NEWLINE) {
		if p.check(scanner.TEXT) {
			b.WriteString(p.peek().Value)
		} else {
			b.WriteRune(scanner.TOKEN_SYMBOL_MAP[p.peek().Kind])
		}
		p.advance()
	}

	if p.check(scanner.NEWLINE) {
		p.advance()
		return Text{content: "[" + alt + "](" + b.String()}
	}

	// skip the closing brace
	p.advance()
	// INFO:  skip the newline, fixes a bug which resulted in fleck not parsing two consecutive images
	p.advance()

	return Image{
		alt: alt,
		src: b.String(),
	}
}

// parses anchors / links
func (p *Parser) link() Tag {
	p.advance()

	b := strings.Builder{}
	for !p.check(scanner.STRAIGHTBRACECLOSE) && !p.check(scanner.NEWLINE) {
		if p.check(scanner.TEXT) {
			b.WriteString(p.peek().Value)
		} else {
			b.WriteRune(scanner.TOKEN_SYMBOL_MAP[p.peek().Kind])
		}
		p.advance()
	}

	if p.check(scanner.NEWLINE) {
		p.advance()
		return Text{content: "[" + b.String()}
	}

	title := b.String()
	b.Reset()

	// skip the [
	p.advance()

	if !p.check(scanner.PARENOPEN) {
		return Text{content: "[" + title + "]"}
	}

	// skip the opening brace
	p.advance()

	for !p.check(scanner.PARENCLOSE) && !p.check(scanner.NEWLINE) {
		if p.check(scanner.TEXT) {
			b.WriteString(p.peek().Value)
		} else {
			b.WriteRune(scanner.TOKEN_SYMBOL_MAP[p.peek().Kind])
		}
		p.advance()
	}

	if p.check(scanner.NEWLINE) {
		p.advance()
		return Text{content: "[" + title + "](" + b.String()}
	}

	// skip the closing brace
	p.advance()

	return Link{
		href:  b.String(),
		title: title,
	}
}

// parses bold and italic elements
func (p *Parser) emphasis() Tag {
	kind := p.peek().Kind
	// skip current symbol
	p.advance()

	// executes if next symbol is also kind, such as: ** or __
	if p.check(kind) {
		// if two symbols ** or __ follow immediately, return them as text
		p.advance()
		b := strings.Builder{}
		for !p.check(kind) && !p.check(scanner.NEWLINE) {
			if p.check(scanner.TEXT) {
				b.WriteString(p.peek().Value)
			} else {
				b.WriteRune(scanner.TOKEN_SYMBOL_MAP[p.peek().Kind])
			}
			p.advance()
		}

		if p.prev().Kind == scanner.NEWLINE {
			p.advance()
			return Text{content: b.String()}
		}

		// skip closing symbols
		p.advance()
		p.advance()

		return Bold{
			text: b.String(),
		}
	} else {
		// return both symbols
		if p.check(kind) {
			// also skip the closing symbol
			p.advance()
			return Text{content: string(scanner.TOKEN_SYMBOL_MAP[p.peek().Kind])}
		}

		b := strings.Builder{}
		for !p.check(kind) && !p.check(scanner.NEWLINE) {
			if p.check(scanner.TEXT) {
				b.WriteString(p.peek().Value)
			} else {
				b.WriteRune(scanner.TOKEN_SYMBOL_MAP[p.peek().Kind])
			}
			p.advance()
		}

		// skip the closing symbol
		p.advance()

		if p.prev().Kind == scanner.NEWLINE {
			return Text{content: b.String()}
		}

		return Italic{
			text: b.String(),
		}
	}
}

// parses code blocks and inline code elements
func (p *Parser) code(quoteContext bool) Tag {
	// FIXED: inline code elements containing dashes (-) are not parsed correctly
	// BUG: if the first item on a line is a inline code element, the rest of the line is detected as a paragraph, but excluding the code element at the beginning
	// BUG: if no language or type is specified the parser assumes the next line to be the content
	p.advance()
	if p.check(scanner.BACKTICK) {
		// code block:
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

		// FIXED: if encountered ` ends the codeblock
		for {
			if p.check(scanner.BACKTICK) {
				p.advance()
				if p.check(scanner.BACKTICK) {
					p.advance()
					if p.check(scanner.BACKTICK) {
						p.advance()
						// skip the \n
						p.advance()
						break
					} else {
						b.WriteString("``")
						continue
					}
				} else {
					b.WriteRune('`')
					continue
				}
			}
			if quoteContext && (p.prev().Kind == scanner.NEWLINE || p.prev().Kind == scanner.EMPTYLINE) {
				// skips the > and the space at the start of the line in a quoted context
				p.advance()
				if p.peek().Value == " " {
					p.advance()
				}
				continue
			}
			if p.check(scanner.TEXT) {
				b.WriteString(p.peek().Value)
			} else {
				b.WriteRune(scanner.TOKEN_SYMBOL_MAP[p.peek().Kind])
			}

			p.advance()
		}

		return CodeBlock{
			language: language,
			text:     b.String(),
		}
	} else {
		return Text{content: "`"}
	}
}

// parses a paragraph, a paragraph ends with an EMPTYLINE
func (p *Parser) paragraph() Tag {
	children := make([]Tag, 0)
	// paragraph should only contain inline code, italic and bold or text
	for !p.check(scanner.EMPTYLINE) && !p.isAtEnd() {
		switch p.peek().Kind {
		case scanner.STRAIGHTBRACEOPEN:
			children = append(children, p.link())
		case scanner.BACKTICK:
			// inline code:
			b := strings.Builder{}
			for !p.check(scanner.BACKTICK) && !p.check(scanner.NEWLINE) {
				if p.check(scanner.TEXT) {
					b.WriteString(p.peek().Value)
				} else {
					b.WriteRune(scanner.TOKEN_SYMBOL_MAP[p.peek().Kind])
				}
				p.advance()
			}
			// skip the `
			p.advance()

			children = append(children, CodeInline{text: b.String()})
		case scanner.STAR, scanner.UNDERSCORE:
			children = append(children, p.emphasis())
		case scanner.TEXT:
			children = append(children, Text{content: p.peek().Value})
			p.advance()
		default:
			children = append(children, Text{content: string(scanner.TOKEN_SYMBOL_MAP[p.peek().Kind])})
			p.advance()
		}
	}

	// skip the newline
	p.advance()
	if len(children) == 0 {
		return nil
	}
	return Paragraph{children: children}
}

// parses headings
func (p *Parser) heading() Tag {
	var lvl uint = 0
	children := make([]scanner.Token, 0)

	for p.check(scanner.HASH) {
		lvl++
		p.advance()
	}

	for !p.check(scanner.NEWLINE) && !p.check(scanner.EMPTYLINE) {
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

	if cli.ARGUMENTS.GetFlag("toc") {
		p.headings = append(p.headings, heading)
	}
	return heading
}

// generates a toc, but only if '--toc' is specified, by default only generates the toc out of headings from h1 to h3.
// If the user specifies the '--toc-full' option the h4,h5 and h6 headings are considered.
func (p *Parser) GenerateToc() string {
	headingMap := map[uint]uint{
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
	}

	b := strings.Builder{}
	b.WriteString("<h3>Table of contents</h3>")
	b.WriteString("<ul id=\"toc\">")
	tocFull := cli.ARGUMENTS.GetFlag("toc-full")
	for _, v := range p.headings {
		if !tocFull && v.lvl > 3 {
			continue
		}

		headingMap[v.lvl]++
		switch v.lvl {
		case 1:
			headingMap[2] = 0
			headingMap[3] = 0
			headingMap[4] = 0
			headingMap[5] = 0
			headingMap[6] = 0
			b.WriteString(fmt.Sprintf("<li><a class=\"toc-h1\" href=\"#%s\">%d</a>: %s</li>", strings.TrimSpace(v.text), headingMap[1], v.text))
		case 2:
			headingMap[3] = 0
			headingMap[4] = 0
			headingMap[5] = 0
			headingMap[6] = 0
			b.WriteString(fmt.Sprintf("<li><a class=\"toc-h2\" href=\"#%s\">%d.%d</a>: %s</li>",
				strings.TrimSpace(v.text),
				headingMap[1],
				headingMap[2],
				v.text),
			)
		case 3:
			headingMap[4] = 0
			headingMap[5] = 0
			headingMap[6] = 0
			b.WriteString(fmt.Sprintf("<li><a class=\"toc-h3\" href=\"#%s\">%d.%d.%d</a>: %s</li>",
				strings.TrimSpace(v.text),
				headingMap[1],
				headingMap[2],
				headingMap[3],
				v.text),
			)
		case 4:
			headingMap[5] = 0
			headingMap[6] = 0
			b.WriteString(fmt.Sprintf("<li><a class=\"toc-h4\" href=\"#%s\">%d.%d.%d.%d</a>: %s</li>",
				strings.TrimSpace(v.text),
				headingMap[1],
				headingMap[2],
				headingMap[3],
				headingMap[4],
				v.text),
			)
		case 5:
			headingMap[6] = 0
			b.WriteString(fmt.Sprintf("<li><a class=\"toc-h5\" href=\"#%s\">%d.%d.%d.%d.%d</a>: %s</li>",
				strings.TrimSpace(v.text),
				headingMap[1],
				headingMap[2],
				headingMap[3],
				headingMap[4],
				headingMap[5],
				v.text))
		case 6:
			b.WriteString(fmt.Sprintf("<li><a class=\"toc-h6\" href=\"#%s\">%d.%d.%d.%d.%d.%d</a>: %s</li>",
				strings.TrimSpace(v.text),
				headingMap[1],
				headingMap[2],
				headingMap[3],
				headingMap[4],
				headingMap[5],
				headingMap[6],
				v.text),
			)
		}
	}
	b.WriteString("</ul>")
	return b.String()
}

// checks if types match with the token.Kind of the next tokens, if yes advance and return true, if not return false
func (p *Parser) match(types ...uint) bool {
	for _, t := range types {
		if p.check(t) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) lookahead(types ...uint) bool {
	i := p.current
	for _, t := range types {
		if p.isAtEnd() {
			return false
		}
		if p.tokens[i].Kind != t {
			return false
		} else {
			i++
		}
	}
	return true
}

// checks if the current token.Kind maches the specified kind, returns false if at end or kind aren't the same
func (p *Parser) check(kind uint) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Kind == kind
}

// move to next token
func (p *Parser) advance() {
	if !p.isAtEnd() {
		p.current++
	}
}

// check if current token is EOF
func (p *Parser) isAtEnd() bool {
	return p.peek().Kind == scanner.EOF
}

// get current token
func (p *Parser) peek() scanner.Token {
	return p.tokens[p.current]
}

// get last token
func (p *Parser) prev() scanner.Token {
	if p.current == 0 {
		return scanner.Token{
			Kind: 0,
		}
	}
	return p.tokens[p.current-1]
}
