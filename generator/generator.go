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

const DEFAULT_TEMPLATE = `<!DOCTYPE html>
<!-- this file was generated using the fleck markdown to html compiler (https://github.com/xnacly/fleck)-->
<!-------- fleck arguments -------->
@FLECK_ARGUMENTS 
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="generator" content="fleck" />
    <title>@FLECK_TITLE</title>
    <style>
    :root {
        --gray: #d0d7de;
        --light-gray: #f2f1f1;
    }
    body {
        font-family: sans-serif;
        padding-left: 2rem;
        padding-right: 2rem;
        background: #fff;
        display: flex;
        align-items: center;
        justify-content: center;
    }
    img {
        border-radius: 0.5rem;
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
    }
    h1, h2 {
        padding-bottom: 0.5rem;
        border-bottom: 1px solid var(--gray);
    }
    </style>
  </head>
  <body>
    <div>
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

	out, err := os.Create(file + ".html")

	if err != nil {
		logger.LError("failed to create file: " + err.Error())
	}

	res := strings.Replace(DEFAULT_TEMPLATE, "@FLECK_TITLE", file, 1)
	res = strings.Replace(res, "@FLECK_CONTENT", writer.String(), 1)

	writer.Reset()
	writer.WriteString("<!-- source file='" + cli.ARGUMENTS.InputFile + "'-->\n")
	for key, value := range cli.ARGUMENTS.Flags {
		writer.WriteString(fmt.Sprintf("<!-- cli.Arguments.Flags[%s]='%v' -->\n", key, *value))
	}

	res = strings.Replace(res, "@FLECK_ARGUMENTS", writer.String(), 1)

	out.Write([]byte(res))

	logger.LInfo("wrote generated html to '" + file + ".html' using the default template, took: " + time.Since(writeStart).String())
}
