# Usage

## Overview

`fleck` is called like so:

```bash
# fleck [options] file
fleck README.md
```

This command creates a file called README.html, which includes the generated html output.

## Help

Calling `fleck` without any arguments results in an error and displays the help page:

```text
Usage:
    fleck [Options] file

Options:
        --help                          prints the help page, exists
        --watch                         watches for changes in the specified page, recompiles the file on change
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

2023/04/18 12:23:50 error: not enough arguments, specify an input file
exit status 1
```

## Options

`fleck` accepts several options:

| Option                   | description                                                              | default value |
| ------------------------ | ------------------------------------------------------------------------ | ------------- |
| `--help`                 | prints the help page, exists                                             | false         |
| `--watch`                | watches for changes in the specified page, recompiles the file on change | false         |
| `--debug`                | enables debug logs                                                       | false         |
| `--version`              | prints version and build information, exists                             | false         |
| `--no-prefix`            | hide the informational comments generated in the output html             | false         |
| `--no-template`          | write html output to a file without the default html scaffolding         | false         |
| `--silent`               | disables all info logs, keeps warnings and errors                        | false         |
| `--toc`                  | generates a table of contents                                            | false         |
| `--toc-full`             | generates a full table of contents, includes headings 1,2,3,4,5,6        | false         |
| `--keep-temp`            | stops fleck from removing temporary files                                | false         |
| `--preprocessor-enabled` | enables the preprocessor and therefore macro expansion                   | false         |
| `--shell-macro-enabled`  | enables the `@shell` macro                                               | false         |

See [Macros](./Macros.md) for more information about the preprocessor and macros.
