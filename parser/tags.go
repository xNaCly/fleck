package parser

import (
	"fmt"
	"strings"
)

type Tag interface {
	String() string
}

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

type Heading struct {
	lvl  uint
	text string
}

func (p Heading) String() string {
	text := strings.TrimSpace(p.text)
	return fmt.Sprintf("<h%d id=\"%s\">%s</h%d>", p.lvl, text, text, p.lvl)
}

type Quote struct {
	children []Tag
}

func (p Quote) String() string {
	b := strings.Builder{}
	b.WriteString("<blockquote>")
	for _, c := range p.children {
		b.WriteString(c.String())
	}
	b.WriteString("</blockquote>")
	return b.String()
}

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

type TodoListItem struct {
	children []Tag
}

func (p TodoListItem) String() string {
	b := strings.Builder{}
	b.WriteString("<li>")
	b.WriteString("<input disabled=\"\">")
	for _, c := range p.children {
		b.WriteString(c.String())
	}
	b.WriteString("</li>")
	return b.String()
}

type CodeBlock struct {
	language string
	text     string
}

func (p CodeBlock) String() string {
	return fmt.Sprintf("<pre class=\"%s\"><code>%s</code></pre>", p.language, p.text)
}

type CodeInline struct {
	text string
}

func (p CodeInline) String() string {
	return "<code>" + p.text + "</code>"
}

type Bold struct {
	text string
}

func (p Bold) String() string {
	return "<strong>" + p.text + "</strong>"
}

type Italic struct {
	text string
}

func (p Italic) String() string {
	return "<em>" + p.text + "</em>"
}

type Image struct {
	alt string
	src string
}

func (p Image) String() string {
	return "<img src=\"" + p.src + "\" alt=\"" + p.alt + "\">"
}

type Link struct {
	href  string
	title string
}

func (p Link) String() string {
	return "<a href=\"" + p.href + "\">" + p.title + "</a>"
}

type Ruler struct{}

func (p Ruler) String() string {
	return "<hr>"
}
