# Macros:

A macro in the context of fleck is a snippet starting with `@`.
The preprocessor (if enabled) will check if the macro is known and if so replaces / expands the macro with the corresponding value.

```text
macro name
 |
 v
@today{2006-01-02}
       ^
       |
       macro argument
```

> **Info**
>
> Macros are expanded via fleck's preprocessor. For performance reasons the preprocessor is disabled by default, to enable the preprocessor and macros, supply fleck with the `--preprocessor-enabled` flag:
>
> ```bash
> fleck --preprocessor-enabled test.md
> ```

#### Include other markdown files

```markdown
## Test.md:

@include{test.md}
```

The above includes the whole content of the test.md.

> **Warning**
>
> `fleck` does currently not support import nesting, this means the root file `fleck` processes can use the `@include` macro, any file already included can not.

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
> fleck --preprocessor-enabled --shell-macro-enabled test.md
> ```
>
> **The preprocessor will execute any command specified as a parameter, this includes each and every command available on your system, such as removing whole directories or stealing data from your system.**

```
Author: @shell{whoami}
```

`@shell` is replaced with the output of the command specified in the argument, here it would result in `Author: teo`
