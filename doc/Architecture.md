# Architecture

## Procedure:

1. Take file as argument: `fleck test.md`
2. if enabled the preprocessor takes over and replaces macros, writes the result to a temporary `test.md.fleck` file
3. Lexer takes over and transforms the contents of the `test.md.fleck` file into a list of token
4. Parser takes over and transforms the list of token into an abstract syntax tree
5. Code generator takes over and transforms the ast into a test.html file

## Fleck's sequences in depth

### Preprocessor

The preprocessor iterates over every charachter of every line of the input file, looking for an '@' and a macro name after it. If it finds a macro it knows, it replaces /expands the macro with its contents.

The preprocessor has to be enabled via the `--preprocessor-enabled` flag, this is due to performance reasons, see [e021abc](https://github.com/xNaCly/fleck/commit/e021abcf227c80b45e6834876695a7a4205936e0)

### Lexer

The [lexer](https://en.wikipedia.org/wiki/Lexical_analysis) iterates over every character of every line in the input file.
It transforms characters into tokens and returns them, read more [here](https://xnacly.me/posts/2023/lexer-markdown/)

### Parser

The parser creates an [ast](https://en.wikipedia.org/wiki/Abstract_syntax_tree) and passes the result to the code generator.

### Code generator
