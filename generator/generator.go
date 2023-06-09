package generator

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/xnacly/fleck/cli"
	"github.com/xnacly/fleck/logger"
	"github.com/xnacly/fleck/parser"
)

var VERSION = ""

const FLECK_PREFIX = `<!-- This file was generated using the fleck markdown to html compiler (https://github.com/xnacly/fleck) -->
<!-- If you found a bug in the generated html, please create a bug report here: https://github.com/xnacly/fleck/issues/new -->
<!-- fleck was invoked as follows:`

const DEFAULT_TEMPLATE = `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8" /><meta name="generator" content="fleck@FLECK_VERSION" /><title>@FLECK_TITLE</title><style>
:root {
    --gray: #d0d7de;
    --light-gray: #f2f1f1;
    --lighter-gray: #f3f2f2; 
    --info: rgb(2, 122, 255);
    --info-lighter: rgba(2, 122, 255, 0.1);
    --warning: rgb(224, 220, 0);
    --warning-lighter: rgba(224, 222, 113, 0.2);
    --danger: rgb(251, 70, 76);
    --danger-lighter: rgba(251, 70, 76, 0.1);
    --note: rgb(83, 223, 221);
    --note-lighter: rgb(83, 223, 221, 0.2); 
}
* {
    box-sizing: border-box;
}
body {
    font-family: -apple-system,BlinkMacSystemFont,"Segoe UI","Noto Sans",Helvetica,Arial,sans-serif,"Apple Color Emoji","Segoe UI Emoji";
    font-size: 16px;
    line-height: 1.5;
    word-wrap: break-word;
    margin: 0;
    padding: 2rem;
    background: #fff;
    display: flex;
    align-items: center;
    justify-content: center;
}
.main {
    min-width: 50%;
    max-width: 50%;
    margin: 0 auto;
}
@media (max-width: 1250px) {
    .main {
        max-width: 80%;
        min-width: 80%;
    }
}
@media (max-width: 600px) {
    .main {
        max-width: 100%;
        min-width: 100%;
    }
}
img {
    display: block;
    border-radius: 0.5rem;
    max-width: 80%;
}
code:not(pre > code){
    background: var(--light-gray);
    padding: 0.2rem;
    border-radius: 0.2rem;
    color: black;
    font-size: 85%;
    font-family: ui-monospace,SFMono-Regular,SF Mono,Menlo,Consolas,Liberation Mono,monospace;
}
pre {
    background: var(--light-gray);
    font-family: ui-monospace,SFMono-Regular,SF Mono,Menlo,Consolas,Liberation Mono,monospace;
    color: black;
    padding: 0.5rem;
    border-radius: 0.2rem;
    overflow-y: auto;
    font-size: 85%;
    line-height: 1.45;
    border-radius: 6px;
}
h1, h2 {
    padding-bottom: 0.5rem;
    border-bottom: 1px solid var(--gray);
    font-weight: 600;
}
h3, h4, h5, h6{
    font-weight: 600;
}
h6{
    color: var(--gray);
}
blockquote {
    border-left: 0.25rem solid #d0d7de;
    color: #656d76;
    padding: 0.25rem;
    padding-right: 2rem;
    padding-left: 1.25rem;
    margin: 0;
    margin-top: 1rem;
    margin-bottom: 1rem;
    border-top-right-radius: 0.2rem;
    border-bottom-right-radius: 0.2rem;
}
blockquote .warning {
    color: var(--warning);
}
blockquote .info {
    color: var(--info);
}
blockquote .danger {
    color: var(--danger);
}
blockquote .note {
    color: var(--note);
}
.callout {
    padding: 1rem;
    border-radius: 0.25rem;
    color: black;
}
.blockquote-info {
    border: 0;
    background-color: var(--info-lighter);
}
.blockquote-danger {
    border: 0;
    background-color: var(--danger-lighter);
}
.blockquote-warning {
    border-radius: 0.5rem;
    border: 0;
    background-color: var(--warning-lighter);
}
.blockquote-note {
    border: 0;
    background-color: var(--note-lighter);
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
a, a:visited, a:active {
    color: var(--info);
    text-decoration: none;
}
a:hover {
    text-decoration: underline;
}
#toc {
    list-style: none;
}
#toc .toc-h2 {
    margin-left: 0.5rem;
}
#toc .toc-h3 {
    margin-left: 1rem;
}
#toc .toc-h4 {
    margin-left: 1.5rem;
}
#toc .toc-h5 {
    margin-left: 1.75rem;
}
#toc .toc-h6 {
    margin-left: 2rem;
}
strong {
    font-weight: 600;
}
</style>`

// write html to a file, writes the prefix with the compilation flags contained before writing the parsed html if '--no-prefix' is not specified.
func WritePlain(fileName string, result []parser.Tag, toc string) {
	// TODO: pass writer to WritePlain, this allows for writing plain to file or write plain to string
	writeStart := time.Now()
	file := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	out, err := os.Create(file + ".html")
	writer := bufio.NewWriter(out)

	if !cli.ARGUMENTS.GetFlag("no-prefix") {
		writer.WriteString(FLECK_PREFIX)
		writer.WriteString("fleck ")
		for _, opt := range cli.OPTIONS {
			val := cli.ARGUMENTS.GetFlag(opt.Name)
			if !val {
				continue
			}
			writer.WriteString(fmt.Sprintf("--%s ", opt.Name))
		}
		writer.WriteString(fileName)
		writer.WriteString("-->")
		writer.WriteString("\n")
	}

	if err != nil {
		logger.LError("failed to open file: " + err.Error())
	}

	if len(toc) != 0 {
		writer.WriteString(toc)
	}

	for _, e := range result {
		writer.WriteString(e.String())
	}

	writer.Flush()
	logger.LDebug("wrote generated html to '" + file + "', took: " + time.Since(writeStart).String())
}

// write html to a file using a template, writes the prefix with the compilation flags contained before writing the parsed html if '--no-prefix' is not specified.
// Replaces @FLECK_TITLE in the template with the input filename without extension. Replaces @FLECK_CONTENT with the parsed markdown.
func WriteTemplate(fileName string, result []parser.Tag, toc string) {
	// TODO: support --template="file.fleckplate"

	writeStart := time.Now()
	file := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	writer := strings.Builder{}

	if !cli.ARGUMENTS.GetFlag("no-prefix") {
		writer.WriteString(FLECK_PREFIX)
		writer.WriteString("fleck ")
		for _, opt := range cli.OPTIONS {
			val := cli.ARGUMENTS.GetFlag(opt.Name)
			if !val {
				continue
			}
			writer.WriteString(fmt.Sprintf("--%s ", opt.Name))
		}
		writer.WriteString(file)
		writer.WriteString("-->")
		writer.WriteString("\n")
	}

	out, err := os.Create(file + ".html")
	if err != nil {
		logger.LError("failed to create file: " + err.Error())
	}

	writer.WriteString(
		strings.Replace(
			strings.Replace(DEFAULT_TEMPLATE, "@FLECK_TITLE", file, 1),
			"@FLECK_VERSION",
			VERSION,
			1,
		),
	)

	syntax := cli.ARGUMENTS.GetFlag("syntax")
	if syntax {
		writer.WriteString(`<link href="https://cdn.jsdelivr.net/npm/prismjs/themes/prism.min.css" rel="stylesheet">`)
	}

	if cli.ARGUMENTS.GetFlag("math") {
		writer.WriteString(`<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/katex@0.16.7/dist/katex.min.css" integrity="sha384-3UiQGuEI4TTMaFmGIZumfRPtfKQ3trwQE2JgosJxCnGmQpL/lJdjpcHkaaFwHlcI" crossorigin="anonymous"> <script defer src="https://cdn.jsdelivr.net/npm/katex@0.16.7/dist/katex.min.js" integrity="sha384-G0zcxDFp5LWZtDuRMnBkk3EphCK1lhEf4UEyEM693ka574TZGwo4IWwS6QLzM/2t" crossorigin="anonymous"></script> <script defer src="https://cdn.jsdelivr.net/npm/katex@0.16.7/dist/contrib/auto-render.min.js" integrity="sha384-+VBxd3r6XgURycqtZ117nYw44OOcIax56Z4dCRWbxyPt0Koah1uHoK0o4+/RRE05" crossorigin="anonymous" onload="renderMathInElement(document.body, {delimiters: [{left: '$$', right: '$$', display: true}, {left:'$', right:'$', display: false}]});"></script>`)
	}

	writer.WriteString("</head><body><div class='main'>")

	if len(toc) != 0 {
		writer.WriteString(toc)
	}

	for _, e := range result {
		writer.WriteString(e.String())
	}

	if syntax {
		writer.WriteString(`</div><script src="https://cdn.jsdelivr.net/npm/prismjs/prism.min.js"></script><script src="https://cdn.jsdelivr.net/npm/prismjs/plugins/autoloader/prism-autoloader.min.js"></script></body></html>`)
	} else {
		writer.WriteString("</div></body></html>")
	}

	livepreview := cli.ARGUMENTS.GetFlag("live-preview")
	if livepreview {
		writer.WriteString(`
		<script>
			const socket = new WebSocket("ws://localhost:12345/content");
			socket.onopen = function (event) {
				console.log("[Fleck] Connected to live preview server");
			};
			socket.onmessage = function (event) {
				if (event.data == "content_changed") {
					console.log("[Fleck] Received content change..");
					window.location.reload();
				}
			};
		</script>
		`)
	}

	out.Write([]byte(writer.String()))

	logger.LDebug("wrote generated html to '" + file + ".html' using the default template, took: " + time.Since(writeStart).String())
}
