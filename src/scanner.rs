use std::fs;

use crate::token::{self, Token};

/// holds variables used to parse the markdown file
pub struct Scanner {
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

impl Scanner {
    /// get a new instance of Parser, reads the file with the given file_path and its first character
    pub fn new(file_path: &str) -> Scanner {
        let input = fs::read_to_string(file_path).expect("could not read file");
        let current_char = input
            .chars()
            .nth(0)
            .expect("could not get first character of input");
        Scanner {
            input,
            current_char,
            last_char: '\0',
            pos: 0,
            line: 0,
            line_pos: 0,
        }
    }

    /// returns true if the parser is a the end of the input
    fn at_end(&self) -> bool {
        if self.current_char == '\0' || self.pos >= self.input.len() {
            return true;
        }
        return false;
    }

    /// returns the next character without advancing or '\0' if at the end of the input
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

    /// parses input, returns a vector of tokens
    pub fn parse(&mut self) -> Vec<Token> {
        let mut res: Vec<Token> = vec![];
        while !self.at_end() && self.current_char != '\0' {
            let mut token_value = String::new();
            let mut token_kind = token::TokenKind::Text;

            match self.current_char {
                _ => {}
            }

            self.advance();
        }
        res
    }
}
