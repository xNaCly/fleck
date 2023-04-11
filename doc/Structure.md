# Structure

Fleck's directory structure is intended to separate logic as much as possible.
Each folder / go sub module contains its own encapsulated logic.
Exported logic should only be called in the `fleck.go` file at the root of the project.

## Directories

### cli

The `cli` modules main objective is to handle parsing of command line arguments and to provide a well thought through interface for flecks end user.

It contains the `Arguments` structure, which in itself contains all the flags and arguments fleck can receive and the `ParseCli` function which returns the structure.

`cli` also contains the function which is used to print the help message (`PrintShortHelp`).

### doc

The doc directory contains all written documentation for fleck, such as [architecture](./Architecture.md) and [structure](#)

### logger

The `logger` module contains logging helpers with colors.

### parser

The `parser` module contains all logic regarding the generation of the abstract syntax tree out of the tokens the scanner / lexer created.

### preprocessor

The `preprocessor` module handles macro expansion.

### scanner

The `scanner` converts a stream of characters into a stream of tokens, an extensive write up of the inner workings of the lexer can be found [here](https://xnacly.me/posts/2023/lexer-markdown/).
