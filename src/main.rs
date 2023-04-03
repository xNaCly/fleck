#![allow(dead_code)]

const VERSION: &str = "v0.1.0";

mod parser;
mod token;
mod transformer;

fn main() {
    println!("fleck - {}", VERSION);

    let args = std::env::args().collect::<Vec<String>>();
    if args.len() < 2 {
        panic!("not enough arguments");
    }

    let file_name = args.get(1).expect("not enough arguments");
    let mut p = parser::Parser::new(&file_name);
    let tokens = p.parse();
    println!("{}", transformer::transform(tokens));
}
