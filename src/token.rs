#[derive(Debug, PartialEq, Eq)]
pub enum TokenKind {
    Unknown,
    Paragraph,
    Heading1,
    Heading2,
    Heading3,
    Heading4,
    Heading5,
    Heading6,
    Quote,
    Listitem,
    CodeInline,
    /// value in codeblock includes the lang of the code block
    CodeBlock(String),
    Ruler,
    Bold,
    Italic,
    Image,
    /// value in link includes the link to point to
    Link(String),
}

#[derive(Debug)]
pub struct Token {
    pub line: usize,
    pub line_pos: usize,
    pub pos: usize,
    pub kind: TokenKind,
    pub content: String,
}
