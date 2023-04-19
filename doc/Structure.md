# Structure

Fleck's directory structure is intended to separate logic as much as possible.
Each folder / go sub module contains its own encapsulated logic.
Exported logic should only be called in the `fleck.go` file at the root of the project.

## Directories

### cli

The `cli` modules main objective is to handle parsing of command line arguments and to provide a well thought through interface for flecks end user.

It contains the `Arguments` structure, which in itself contains all the flags and arguments fleck can receive and the `ParseCli` function which returns the structure.

`cli` also contains the functions used to print the help message (`PrintShortHelp` & `PrintLongHelp`).

### core

The `core` module contains functions necessary for flecks execution, such as the check if the provided flags are sensible and no flag is set which requires an other flag to have an effect on the output.
It also contains the `Run` function and the `Watch` function, which are both used to wrap the calls to the preprocessor, lexer, parser and generator.

### doc

The doc directory contains all written documentation for fleck, such as [architecture](./Architecture.md) and [structure](#)

### generator

The generator module includes the default template fleck will write its output to and is responsible for setting the outputs document title, styling and content.

### logger

The `logger` module contains logging helpers with colors.

### parser

The `parser` module contains all logic regarding the generation of the abstract syntax tree out of the tokens the scanner / lexer created.

### preprocessor

The `preprocessor` module handles macro expansion.

### scanner

The `scanner` converts a stream of characters into a stream of tokens, an extensive write up of the inner workings of the lexer can be found [here](https://xnacly.me/posts/2023/lexer-markdown/).
