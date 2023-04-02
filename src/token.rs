#[derive(Debug)]
pub enum TokenKind {
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
    CodeBlock,
    Ruler,
    Emphasis,
    Image,
}

#[derive(Debug)]
pub struct Token {
    pub pos: usize,
    pub kind: TokenKind,
    pub content: String,
}
