use crate::token::Token;

pub fn transform(tokens: Vec<Token>) -> String {
    for token in tokens {
        dbg!(token);
    }
    return String::new();
}
