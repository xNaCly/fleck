package parser

import (
	"fmt"
	"html"
	"strconv"
	"strings"
)

// generic interface, allows us to return nil and all the implementing structs in the parser
type Tag interface {
	String() string
}

type Br struct{}

func (p Br) String() string {
	return "</br>"
}

// <p></p> html paragraph
type Paragraph struct {
	children []Tag
}

func (p Paragraph) String() string {
	b := strings.Builder{}
	b.WriteString("<p>")
	for _, c := range p.children {
		b.WriteString(c.String())
	}
	b.WriteString("</p>")
	return b.String()
}

// contains plaintext
type Text struct {
	content string
}

func (p Text) String() string {
	return p.content
}

// any of h1,h2,h3,h4,h5,h6, suffix is denoted using the lvl field
type Heading struct {
	lvl  uint
	text string
}

func (p Heading) String() string {
	text := strings.TrimSpace(p.text)
	return fmt.Sprintf("<h%d id=\"%s\">%s</h%d>", p.lvl, text, text, p.lvl)
}

// <blockquote></blockquote>, can contain all the other elements
type Quote struct {
	children []Tag
}

func (p Quote) String() string {
	b := strings.Builder{}
	var callout bool
	var cType string
	for _, c := range p.children {
		if !callout {
			switch c.(type) {
			case Bold:
				t := c.(Bold)
				switch strings.ToLower(t.text) {
				case "warning":
					t.className = "warning"
				case "info":
					t.className = "info"
				case "danger":
					t.className = "danger"
				case "note":
					t.className = "note"
				default:
					b.WriteString(t.String())
					continue
				}
				b.WriteString(t.String())
				// used to stop checking for callouts
				callout = true
				cType = t.className
				continue
			}
		}

		b.WriteString(c.String())
	}
	prefix := "<blockquote>"

	if callout {
		prefix = "<blockquote class=\"callout blockquote-" + cType + "\">"
	}

	return prefix + b.String() + "</blockquote>"
}

// <ul></ul>, contains ListItem
type List struct {
	children []Tag
}

func (p List) String() string {
	b := strings.Builder{}
	b.WriteString("<ul>")
	for _, c := range p.children {
		b.WriteString(c.String())
	}
	b.WriteString("</ul>")
	return b.String()
}

// <li></li>, can contain almost everything else
type ListItem struct {
	children []Tag
}

func (p ListItem) String() string {
	b := strings.Builder{}
	b.WriteString("<li>")
	for _, c := range p.children {
		b.WriteString(c.String())
	}
	b.WriteString("</li>")
	return b.String()
}

// listitem but with a prefixed disabled checkmark
type TodoListItem struct {
	checked  bool
	children []Tag
}

func (p TodoListItem) String() string {
	b := strings.Builder{}
	b.WriteString("<li>")
	b.WriteString("<input type=\"checkbox\" disabled=\"true\" checked=\"" + strconv.FormatBool(p.checked) + "\">")
	for _, c := range p.children {
		b.WriteString(c.String())
	}
	b.WriteString("</li>")
	return b.String()
}

// <pre><code></code></pre>, contains plaintext and whitespaces, which MUST to be respec
type CodeBlock struct {
	language string
	text     string
}

func (p CodeBlock) String() string {
	return fmt.Sprintf("<pre ><code class=\"language-%s\">%s</code></pre>", p.language, html.EscapeString(p.text))
}

// <code></code>, inline code element, contains plain text
type CodeInline struct {
	text string
}

func (p CodeInline) String() string {
	return "<code>" + p.text + "</code>"
}

// <strong></strong>, bold text
type Bold struct {
	className string
	text      string
}

func (p Bold) String() string {
	return fmt.Sprintf("<strong class=\"%s\">%s</strong>", p.className, p.text)
}

// <em></em>, italic text
type Italic struct {
	text string
}

func (p Italic) String() string {
	return "<em>" + p.text + "</em>"
}

type StrikeThrough struct {
	text string
}

func (p StrikeThrough) String() string {
	return "<s>" + p.text + "</s>"
}

type Highlight struct {
	text string
}

func (p Highlight) String() string {
	return "<mark>" + p.text + "</mark>"
}

// <img src="" alt="">, image with alt and src
type Image struct {
	alt string
	src string
}

func (p Image) String() string {
	return "<img src=\"" + p.src + "\" alt=\"" + p.alt + "\">"
}

// <a href=""></a>, anchor with href and title
type Link struct {
	href  string
	title string
}

func (p Link) String() string {
	return "<a href=\"" + p.href + "\">" + p.title + "</a>"
}

// <hr>
type Ruler struct{}

func (p Ruler) String() string {
	return "<hr>"
}
