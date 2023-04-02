use std::fs;

use crate::token::{self, Token};

pub struct Parser {
    current_char: char,
    input: String,
    pos: usize,
    line: usize,
    line_pos: usize,
}

impl Parser {
    pub fn new(file_path: &str) -> Parser {
        let input = fs::read_to_string(file_path).expect("could not read file");
        let current_char = input
            .chars()
            .nth(1)
            .expect("could not get first character of input");
        Parser {
            input,
            current_char,
            pos: 0,
            line: 0,
            line_pos: 0,
        }
    }

    fn at_end(&self) -> bool {
        if self.current_char == '\0' || self.pos >= self.input.len() {
            return true;
        }
        return false;
    }

    fn peek(&self) -> char {
        self.input.chars().nth(self.pos + 1).unwrap_or('\0')
    }

    fn peek_equals(&self, character: char) -> bool {
        self.peek() == character
    }

    fn advance(&mut self) {
        if !self.at_end() && self.pos + 1 <= self.input.len() {
            self.pos += 1;
            self.current_char = self.input.chars().nth(self.pos).unwrap_or('\0');
            self.line_pos += 1;
        }
    }

    pub fn parse(&mut self) -> Vec<Token> {
        let mut res: Vec<Token> = vec![];
        while !self.at_end() && self.current_char != '\0' {
            let mut token_value = String::new();
            let mut token_kind = token::TokenKind::Paragraph;

            match self.current_char {
                '#' => {
                    // skip over '#' with a counter:
                    let mut heading_id = 1;
                    while self.peek_equals('#') {
                        heading_id += 1;
                        self.advance();
                    }
                    while self.peek() != '\n' {
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
                    }
                }
                _ => {}
            }

            res.push(Token {
                pos: self.pos,
                kind: token_kind,
                content: token_value.to_owned(),
            });
            self.advance();
        }
        res
    }
}
