# fleck - Markdown-parser

- [Documentation](./doc/Readme.md)

_Fleck_ **is** intended for outputting standalone html. Fleck is german for mark. Fleck is as minimal as possible and requires no dependencies.

> **Warning**
>
> Fleck is not production or release ready, please come back once i have a release candidate ready

## Install

## Usage

Fleck will return an error if no input file is specified. It is called as follows:

```bash
# fleck [options] file
fleck README.md
```

This command creates a file called README.html, which includes the generated html output.

## Supported markdown features

> Fleck implements the markdown format as proposed by John Gruber and Aron Swartz, as defined [here](https://daringfireball.net/projects/markdown/syntax).

Features:

- [ ] Headers
- [ ] Blockquotes
- [ ] unordered Lists
- [ ] to-do lists
- [ ] Code blocks
- [ ] Code inline
- [ ] Bold
- [ ] Italic
- [ ] Image
- [ ] Link
- [ ] horizontal ruler
- [ ] Table
- [ ] inline html (probably never supported)

### Macros:

Extensions / macros for Markdown implemented with fleck support the following syntax and features:

#### Include other markdown files

```markdown
## Test.md:

@include{test.md}
```

The above includes the whole content of the test.md. Similar to a preprocessor in c.

#### Include the current date:

```markdown
Today is @today{2006-01-02}.
```

`@today` gets replaced with the current date, according to the format specified in its argument, here it would result in `Today is 2023-04-08`.
The `@today` macro accepts go format strings, read more [here](https://www.digitalocean.com/community/tutorials/how-to-use-dates-and-times-in-go).

#### Include command output:

> **Warning**
> This macro is very dangerous and needs to be enabled via the `--shell-macro-enabled` flag, like so:
>
> ```bash
> fleck --shell-macro-enabled test.md
> ```
>
> **The preprocessor will execute any command specified as a parameter, this includes each and every command available on your system.**

```
Author: @shell{whoami}
```

`@shell` is replaced with the output of the command specified in the argument, here it would result in `Author: teo`
