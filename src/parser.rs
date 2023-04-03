use std::fs;

use crate::token::{self, Token};

/// holds variables used to parse the markdown file
pub struct Parser {
    pub current_char: char,
    /// whole file in a string
    pub input: String,
    /// which pos the parser is at in the whole file
    pub pos: usize,
    /// which line the parser is at
    pub line: usize,
    /// which pos the parser is at in the current line
    pub line_pos: usize,
    /// notate the previous character
    pub last_char: char,
}

impl Parser {
    /// get a new instance of Parser, reads the file with the given file_path and its first character
    pub fn new(file_path: &str) -> Parser {
        let input = fs::read_to_string(file_path).expect("could not read file");
        let current_char = input
            .chars()
            .nth(0)
            .expect("could not get first character of input");
        Parser {
            input,
            current_char,
            last_char: '\0',
            pos: 0,
            line: 0,
            line_pos: 0,
        }
    }

    /// returns true if the parser is a the end of the file
    fn at_end(&self) -> bool {
        if self.current_char == '\0' || self.pos >= self.input.len() {
            return true;
        }
        return false;
    }

    /// returns the next character without advancing or '\0' if at the end of the file
    fn peek(&self) -> char {
        self.input.chars().nth(self.pos + 1).unwrap_or('\0')
    }

    /// returns true if the next character is equal to character
    fn peek_equals(&self, character: char) -> bool {
        self.peek() == character
    }

    /// advances to the next character in the input, increments position & line_postion by one
    fn advance(&mut self) {
        if !self.at_end() && self.pos + 1 <= self.input.len() {
            self.pos += 1;
            self.last_char = self.current_char;
            self.current_char = self.input.chars().nth(self.pos).unwrap_or('\0');
            self.line_pos += 1;
        } else {
            self.current_char = '\0';
        }
    }

    /// wrapper for creating a token, assigns line, line_pos and pos
    fn create_token(&mut self, token_kind: token::TokenKind, token_value: String) -> Token {
        Token {
            line: self.line,
            line_pos: self.line_pos - token_value.len(),
            pos: self.pos,
            kind: token_kind,
            content: token_value,
        }
    }

    /// wrapper for creating a paragraph
    fn create_paragraph(&mut self, text: &str) -> Token {
        self.create_token(token::TokenKind::Paragraph, String::from(text))
    }

    /// parses input, returns a vector of tokens
    pub fn parse(&mut self) -> Vec<Token> {
        let mut res: Vec<Token> = vec![];
        let mut last_paragraph = String::new();
        while !self.at_end() && self.current_char != '\0' {
            let mut token_value = String::new();
            let mut token_kind = token::TokenKind::Unknown;

            match self.current_char {
                '\n' => {
                    if !last_paragraph.is_empty() {
                        res.push(self.create_paragraph(&last_paragraph));
                        last_paragraph = String::new();
                    }
                    self.line += 1;
                    self.line_pos = 0;
                    self.advance();
                    continue;
                }
                '\r' => {
                    self.advance();
                    continue;
                }
                '#' => {
                    if self.last_char != '\n' {
                        last_paragraph.push(self.current_char);
                        self.advance();
                        continue;
                    }
                    if !(self.line_pos == 0 || self.line_pos == 1) {
                        last_paragraph.push(self.current_char);
                        self.advance();
                        continue;
                    }
                    if !last_paragraph.is_empty() {
                        res.push(self.create_paragraph(&last_paragraph));
                        last_paragraph = String::new();
                    }

                    // skip over '#' with a counter:
                    let mut heading_id = 1;
                    self.advance();
                    while self.current_char == '#' {
                        self.advance();
                        heading_id += 1;
                    }
                    // consume last #
                    self.advance();

                    while self.current_char != '\n' {
                        token_value.push(self.current_char);
                        self.advance();
                    }

                    token_kind = match heading_id {
                        1 => token::TokenKind::Heading1,
                        2 => token::TokenKind::Heading2,
                        3 => token::TokenKind::Heading3,
                        4 => token::TokenKind::Heading4,
                        5 => token::TokenKind::Heading5,
                        6 => token::TokenKind::Heading6,
                        _ => token::TokenKind::Paragraph,
                    };
                }
                '`' => {
                    if !last_paragraph.is_empty() {
                        res.push(self.create_paragraph(&last_paragraph));
                        last_paragraph = String::new();
                    }
                    if self.peek_equals('`') {
                        self.advance();
                        if self.peek_equals('`') {
                            let mut code_lang = String::new();
                            self.advance();
                            self.advance();
                            while self.current_char != '\n' {
                                code_lang.push(self.current_char);
                                self.advance();
                            }

                            while self.current_char != '`' {
                                token_value.push(self.current_char);
                                self.advance();
                            }

                            token_kind = token::TokenKind::CodeBlock(code_lang);
                        }
                    } else {
                        self.advance();
                        while self.current_char != '`' {
                            token_value.push(self.current_char);
                            self.advance();
                        }
                        token_kind = token::TokenKind::CodeInline;
                        // consume `
                        self.advance();
                    }
                }
                '_' => {
                    if !last_paragraph.is_empty() {
                        res.push(self.create_paragraph(&last_paragraph));
                        last_paragraph = String::new();
                    }
                    // consume opening '_'
                    self.advance();
                    token_kind = token::TokenKind::Italic;
                    while self.current_char != '_' {
                        token_value.push(self.current_char);
                        self.advance();
                    }
                    // consume closing '_'
                    self.advance();
                }
                '-' => {
                    if self.last_char != '\n' {
                        last_paragraph.push(self.current_char);
                        self.advance();
                        continue;
                    }
                    let mut minus_amount = 1;
                    self.advance();

                    while self.current_char == '-' {
                        minus_amount += 1;
                        self.advance();
                    }

                    match minus_amount {
                        x if x >= 3 => {
                            if !last_paragraph.is_empty() {
                                res.push(self.create_paragraph(&last_paragraph));
                                last_paragraph = String::new();
                            }
                            res.push(self.create_token(token::TokenKind::Ruler, String::new()));
                            self.advance();
                            continue;
                        }
                        1 => {
                            self.advance();
                            match self.current_char {
                                // match todo list
                                '[' => {
                                    if !last_paragraph.is_empty() {
                                        res.push(self.create_paragraph(&last_paragraph));
                                        last_paragraph = String::new();
                                    }
                                    // skip [
                                    self.advance();
                                    // if current char is x, list is done, otherwise not done
                                    token_kind =
                                        token::TokenKind::CheckListItem(self.current_char == 'x');
                                    // advance to ]
                                    self.advance();
                                    // advance to line content
                                    self.advance();
                                    while self.current_char != '\n' {
                                        token_value.push(self.current_char);
                                        self.advance();
                                    }
                                    res.push(self.create_token(token_kind, token_value));
                                    continue;
                                }
                                // match unordered list
                                _ => {
                                    if !last_paragraph.is_empty() {
                                        res.push(self.create_paragraph(&last_paragraph));
                                        last_paragraph = String::new();
                                    }
                                    while self.current_char != '\n' {
                                        token_value.push(self.current_char);
                                        self.advance();
                                    }
                                    res.push(
                                        self.create_token(token::TokenKind::Listitem, token_value),
                                    );
                                    continue;
                                }
                            }
                        }
                        _ => {
                            continue;
                        }
                    }
                }
                '*' => {
                    if self.peek_equals('*') {
                        if !last_paragraph.is_empty() {
                            res.push(self.create_paragraph(&last_paragraph));
                            last_paragraph = String::new();
                        }
                        token_kind = token::TokenKind::Bold;
                        // consume opening '*'
                        self.advance();

                        // check for horizontal ruler
                        if self.peek_equals('*') {
                            if !last_paragraph.is_empty() {
                                res.push(self.create_paragraph(&last_paragraph));
                                last_paragraph = String::new();
                            }
                            res.push(self.create_token(token::TokenKind::Ruler, String::new()));
                            self.advance();
                            continue;
                        }

                        // consume second opening '*'
                        self.advance();
                        while self.current_char != '*' {
                            token_value.push(self.current_char);
                            self.advance();
                        }
                        // consume closing '*'
                        self.advance();
                    }
                }
                _ => {
                    last_paragraph.push(self.current_char);
                    self.advance();
                    continue;
                }
            }

            if token_kind != token::TokenKind::Unknown {
                res.push(self.create_token(token_kind, token_value))
            }
            self.advance();
        }
        res
    }
}
