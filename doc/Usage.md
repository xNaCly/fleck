# Usage

## Overview

`fleck` is called like so:

```bash
# fleck [options] file
fleck README.md
```

This command creates a file called README.html, which includes the generated html output.

## Help

Calling `fleck` with the `--help`flag will produce the following output:

```text
Usage:
    fleck [Options] file

Options:
        --help                          prints the help page, exists
        --watch                         watches for changes in the specified page, recompiles the file on change
        --syntax                        enables syntax highlighting for code blocks in the output
        --live-preview                  same as watch, serves the compiled html, reloads tab if change occured
        --debug                         enables debug logs
        --version                       prints version and build information, exists
        --no-prefix                     hide the informational comments generated in the output html
        --no-template                   write html output to a file without the default html scaffolding
        --silent                        disables info logs, reduces output significantly
        --toc                           generates a table of contents at the top of the output file, includes headings 1,2,3
        --toc-full                      generates a full toc, includes headings 1,2,3,4,5,6
        --keep-temp                     keeps fleck from removing temporary files, used for debug purposes
        --preprocessor-enabled          enables the preprocessor to replace macros, decreases performance
        --shell-macro-enabled           enables the dangerous '@shell{command}' macro, which allows the preprocessor to run any command on your system

Online documentation: https://github.com/xnacly/fleck
```

See [Macros](./Macros.md) for more information about the preprocessor and macros.
