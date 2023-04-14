package generator

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/xnacly/fleck/cli"
	"github.com/xnacly/fleck/logger"
	"github.com/xnacly/fleck/parser"
)

const FLECK_PREFIX = `<!-- This file was generated using the fleck markdown to html compiler (https://github.com/xnacly/fleck) -->
<!-- If you found a bug in the generated html, please create a bug report here: https://github.com/xnacly/fleck/issues/new -->
<!-- fleck was invoked as follows:`

const DEFAULT_TEMPLATE = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="generator" content="fleck" />
    <title>@FLECK_TITLE</title>
    <style>
    :root {
        --gray: #d0d7de;
        --light-gray: #f2f1f1;
        --lighter-gray: #f8f8f8;
    }
    body {
        font-family: sans-serif;
        margin: 0;
        padding: 2rem;
        background: #fff;
        display: flex;
        align-items: center;
        justify-content: center;
    }
    .main {
        max-width: 50%
        margin: 0 auto;
    }
    img {
        display: block;
        border-radius: 0.5rem;
        max-height: 40rem;
    }
    code:not(pre > code){
        background: var(--light-gray);
        padding: 0.2rem;
        border-radius: 0.2rem;
    }
    pre {
        background: var(--light-gray);
        padding: 0.5rem;
        border-radius: 0.2rem;
        overflow-y: auto;
    }
    h1, h2 {
        padding-bottom: 0.5rem;
        border-bottom: 1px solid var(--gray);
    }
    blockquote {
        border-left: 0.25rem solid var(--gray);
        background: var(--lighter-gray);
        padding: 0.25rem;
        padding-top: 0.5rem;
        padding-bottom: 0.5rem;
        padding-right: 2rem;
        margin: 0;
        margin-top: 0.25rem;
        margin-bottom: 0.25rem;
        padding-left: 0.5rem;
        border-top-right-radius: 0.2rem;
        border-bottom-right-radius: 0.2rem;
    }
    hr {
        height: 0.15rem;
        padding: 0;
        margin: 0;
        margin-top: 0.5rem;
        margin-bottom: 0.5rem;
        background: var(--gray);
        border: 0;
    }
    </style>
  </head>
  <body>
    <div class="main">
        @FLECK_CONTENT 
    </div>
  </body>
</html>
`

// write html to a file
func WritePlain(fileName string, result []parser.Tag, toc string) {
	writeStart := time.Now()
	name := strings.Split(fileName, ".")[0] + ".html"
	out, err := os.Create(name)
	writer := bufio.NewWriter(out)

	if err != nil {
		logger.LError("failed to open file: " + err.Error())
	}

	if len(toc) != 0 {
		writer.WriteString(toc)
	}

	for _, e := range result {
		if cli.GetFlag(cli.ARGUMENTS, "minify") {
			writer.WriteString(e.String())
		} else {
			writer.WriteString(e.String() + "\n")
		}
	}

	writer.Flush()
	logger.LInfo("wrote generated html to '" + name + "', took: " + time.Since(writeStart).String())
}

// write html to a file using a template
func WriteTemplate(fileName string, result []parser.Tag, toc string) {
	writeStart := time.Now()
	file := strings.Split(fileName, ".")[0]
	writer := strings.Builder{}

	if !cli.GetFlag(cli.ARGUMENTS, "no-prefix") {
		writer.WriteString(FLECK_PREFIX)
		writer.WriteString("fleck ")
		for _, opt := range cli.OPTIONS {
			val := cli.GetFlag(cli.ARGUMENTS, opt.Name)
			if !val {
				continue
			}
			writer.WriteString(fmt.Sprintf("--%s ", opt.Name))
		}
		writer.WriteString(cli.ARGUMENTS.InputFile)
		writer.WriteString("-->")
		writer.WriteString("\n")
	}

	out, err := os.Create(file + ".html")
	out.Write([]byte(writer.String()))
	writer.Reset()

	if len(toc) != 0 {
		writer.WriteString(toc)
	}
	for _, e := range result {
		if cli.GetFlag(cli.ARGUMENTS, "minify") {
			writer.WriteString(e.String())
		} else {
			writer.WriteString(e.String() + "\n")
		}
	}

	if err != nil {
		logger.LError("failed to create file: " + err.Error())
	}

	res := strings.Replace(DEFAULT_TEMPLATE, "@FLECK_TITLE", file, 1)
	res = strings.Replace(res, "@FLECK_CONTENT", writer.String(), 1)

	out.Write([]byte(res))

	logger.LInfo("wrote generated html to '" + file + ".html' using the default template, took: " + time.Since(writeStart).String())
}
