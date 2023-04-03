#[derive(Debug, PartialEq, Eq)]
pub enum TokenKind {
    Unknown,
    /// <p>paragraph</p> / paragraph
    Paragraph,
    /// <h1>heading</h1> / # heading
    Heading1,
    /// <h2>heading</h2> / ## heading
    Heading2,
    /// <h3>heading</h4> / ### heading
    Heading3,
    /// <h4>heading</h4> / #### heading
    Heading4,
    /// <h5>heading</h5> / ##### heading
    Heading5,
    /// <h6>heading</h6> / ###### heading
    Heading6,
    /// <blockquote>quote</blockquote> / > quote
    Quote,
    /// <li>list element</li> / - list element
    Listitem,
    /// <code>code</code> / `code`
    CodeInline,
    /// <pre><code>console.log("test")</code></pre> / ```js\nconsole.log("test")```
    /// value in codeblock includes the lang of the code block
    CodeBlock(String),
    /// <hr> / ---
    Ruler,
    /// <b>text</b> / **text**
    Bold,
    /// <i>text</i> / _text_
    Italic,
    /// <img src="path"/> / ![image](path)
    Image,
    /// <a href="path">name</a> / [name](path)
    /// value in link includes the link to point to
    Link(String),
}

#[derive(Debug)]
pub struct Token {
    /// current line
    pub line: usize,
    /// current position on current line
    pub line_pos: usize,
    /// current position in file
    pub pos: usize,
    /// type of token
    pub kind: TokenKind,
    /// value / content of token
    pub content: String,
}
