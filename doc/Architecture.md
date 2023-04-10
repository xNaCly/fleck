# Architecture

## Procedure:

1. Take file as argument: `fleck test.md`
2. Preprocessor takes over and replaces macros, writes the result to a temporary `test.md.fleck` file
3. Lexer takes over and transforms the contents of the `test.md.fleck` file into a list of token
4. Parser takes over and transforms the list of token into an abstract syntax tree
5. Code generator takes over and transforms the ast into a test.html file

## Fleck's sequences in depth

### Preprocessor

### Lexer

### Parser

### Code generator
