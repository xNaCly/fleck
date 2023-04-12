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
        --minify                        minifies the html output
        --toc-full                      generates a full toc, includes headings 1,2,3,4,5,6
        --silent                        disables info logs, reduces output significantly
        --toc                           generates a table of contents at the top of the output file, includes headings 1,2,3
        --keep-temp                     keeps fleck from removing temporary files, used for debug purposes
        --preprocessor-enabled          enables the preprocessor to replace macros, decreases performance
        --shell-macro-enabled           enables the dangerous '@shell{command}' macro, which allows the preprocessor to run any command on your system

2023/04/12 10:13:30 error: not enough arguments, specify an input file
exit status 1
```

## Options

`fleck` accepts several options:

| Option                   | description                                                       | default value |
| ------------------------ | ----------------------------------------------------------------- | ------------- |
| `--minify`               | minifies the html output                                          | false         |
| `--toc-full`             | generates a full table of contents, includes headings 1,2,3,4,5,6 | false         |
| `--silent`               | disables all info logs, keeps warnings and errors                 | false         |
| `--toc`                  | generates a table of contents                                     | false         |
| `--keep-temp`            | stops fleck from removing temporary files                         | false         |
| `--preprocessor-enabled` | enables the preprocessor and therefore macro expansion            | false         |
| `--shell-macro-enabled`  | enables the `@shell` macro                                        | false         |

See [Macros](./Macros.md) for more information about the preprocessor and macros.
