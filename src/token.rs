#[derive(Debug, PartialEq, Eq)]
pub enum TokenKind {
    Hash,
    Underscore,
    Star,
    Newline,
    Dash,
    StraightBraceOpen,
    StraightBraceClose,
    ParenOpen,
    ParenClose,
    BackTick,
    Text,
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
