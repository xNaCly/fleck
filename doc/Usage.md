# Usage

## Overview

`fleck` is called like so:

```bash
# fleck [options] file
fleck README.md
```

This command creates a file called README.html, which includes the generated html output.

## Help

Calling `fleck` with the `--help` flag will produce the following output:

```text
Usage:
    fleck [Options] file

Options:
        Name                    Default         Requires                Description         

        --help                  false                                   prints the help page, exists
        --watch                 false                                   watches for changes in the specified page, recompiles the file on change
        --syntax                false                                   enables syntax highlighting for code blocks in the output
        --live-preview          false                                   same as watch, serves the compiled html, reloads tab if change occured
        --debug                 false                                   enables debug logs  
        --version               false                                   prints version and build information, exists
        --no-prefix             false                                   hide the informational comments generated in the output html
        --no-template           false                                   write html output to a file without the default html scaffolding
        --silent                false                                   disables info logs, reduces output significantly
        --toc                   false                                   generates a table of contents at the top of the output file, includes headings 1,2,3
        --toc-full              false           toc                     generates a full toc, includes headings 1,2,3,4,5,6
        --keep-temp             false           preprocessor-enabled    keeps fleck from removing temporary files, used for debug purposes
        --preprocessor-enabled  false                                   enables the preprocessor to replace macros, decreases performance
        --shell-macro-enabled   false           preprocessor-enabled    enables the dangerous '@shell{command}' macro
        --port                  12345           live-preview            specify the port for '--live-preview' to be served on

Online documentation: https://github.com/xnacly/fleck
```

See [Macros](./Macros.md) for more information about the preprocessor and macros.

## Command line option reference

### `--help`

Prints the help page, containing the name of all available options, their default values, which other option they require and a description.

A short help is printed if fleck is called without any arguments:

```text
$ fleck
Usage:
    fleck [Options] file

    Run 'fleck --help' for an in depth help page
    
2023/04/24 10:14:35 error: not enough arguments, specify an input file
exit status 1
```

### `--watch`

The watch option makes fleck watch for changes in the specified source file:

```text
$ fleck --watch README.md
2023/04/24 10:16:13 info: compiled 'README.md', took: 262.228µs
2023/04/24 10:16:13 info: watching for changes...
```

Fleck checks every 100ms if a change occured.
When this happens the screen is cleared and fleck tells the user how many times it already recompiled the source.:

```text
2023/04/24 10:17:34 info: detected change, recompiling... (1)
2023/04/24 10:17:34 info: compiled 'README.md', took: 500.561µs
2023/04/24 10:17:34 info: detected change, recompiling... (2)
2023/04/24 10:17:34 info: compiled 'README.md', took: 720.964µs
```

Fleck checks if the source file changed by comparing its last modification time and its size with the information gathered in the previous iteration.

### `--syntax`

The `--syntax` flag instructs fleck to inject three assets into the generated template.
The first is the [prism][https://prismjs.com/] default css file. The second is the prism javascript source and the third is the language autoloader, which detects used languages in the generated html and automatically loads the corresponding themes.

```text
$ fleck --syntax README.md
2023/04/24 10:23:18 info: compiled 'README.md', took: 186.067µs
```

A code block in the resulting html, looks like the following:

![syntax-highlighting](./assets/syntax-highlighting.png)

### `--live-preview`
### `--debug`
### `--version`
### `--no-prefix`
### `--no-template`
### `--silent`
### `--toc`
### `--toc-full`
### `--keep-temp`
### `--preprocessor-enabled`
### `--shell-macro-enabled`
### `--port`
