package core

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/xnacly/fleck/cli"
	"github.com/xnacly/fleck/generator"
	"github.com/xnacly/fleck/logger"
	"github.com/xnacly/fleck/parser"
	"github.com/xnacly/fleck/preprocessor"
	"github.com/xnacly/fleck/scanner"
)

// alerts the user if a flag depends on a different flag to have an effect
func FlagCombinationSensible() {
	for _, f := range cli.OPTIONS {
		if len(f.Requires) == 0 {
			continue
		}
		if cli.GetFlag(cli.ARGUMENTS, f.Name) && !cli.GetFlag(cli.ARGUMENTS, f.Requires) {
			logger.LWarn(fmt.Sprintf("flag '--%s' requires flag '--%s' to be set, otherwise it has no effect.", f.Name, f.Requires))
		}
	}
}

// spawns a tab in the default browser on the targeted operating system with the href pointing to `url`
//
// - windows: `cmd /c start`
//
// - linux and bsd's: `xdg-open`
//
// - darwin: `open`
//
// taken from: https://gist.github.com/sevkin/9798d67b2cb9d07cb05f89f14ba682f8
func spawnDefaultBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default:
		cmd = "xdg-open"
	}

	return exec.Command(cmd, append(args, url)...).Start()
}

// live preview allows the user to edit the markdown source and view the effects of the modification in the browser via reload on change
//
// 1. compiles the source (server)
//
// 2. serves the html output (server)
//
// 3. injects a websocket script into the html (server)
//
// 4. starts a websocket connection (server)
//
// 5. opens a tab in the default browser pointing to the html (server)
//
// 6. establishes a websocket connection (client)
//
// 7. waits for changes, notifys the client via websocket when the source changes (server)
//
// 8. reloads if websocket event go (client)
func LivePreview(fileName string) {
	// TODO: implement hmr (rewrite the body instead of reloading the window)

	var upgrader = websocket.Upgrader{}
	var conn *websocket.Conn = nil

	// TODO: make this a flag, --port x
	port := "12345"

	logger.LInfo("starting live preview")

	file := strings.TrimSuffix(fileName, filepath.Ext(fileName))

	http.HandleFunc("/content", func(w http.ResponseWriter, r *http.Request) {
		connection, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			logger.LError("upgrade failed")
			return
		}

		conn = connection
	})

	go WatchForChanges(fileName, func(s string) {
		Run(s)
		if conn != nil {
			err := conn.WriteMessage(websocket.TextMessage, []byte("content_changed"))

			if err != nil {
				logger.LError("failed to send message to websocket: " + err.Error())
			}
		}
	})

	// FIXED: images are missing, serve whole dir
	http.Handle("/", http.FileServer(http.Dir(filepath.Dir(fileName))))

	logger.LInfo("listening on http://localhost:" + port + "/" + file + ".html")
	spawnDefaultBrowser("http://localhost:" + port + "/" + file + ".html")
	http.ListenAndServe(":"+port, nil)
}

// watches for changes in a file, recompiles the file if a change occurs, can be exited via <C-c>
func WatchForChanges(fileName string, executor func(string)) {
	executor(fileName)
	logger.LInfo("watching for changes...")

	iStat, err := os.Stat(fileName)
	if err != nil {
		logger.LError("failed to watch for changes: " + err.Error())
	}

	i := 0
	for {
		stat, err := os.Stat(fileName)
		if err != nil {
			logger.LError("failed to watch for changes: " + err.Error())
		}

		if stat.Size() != iStat.Size() || stat.ModTime() != iStat.ModTime() {
			iStat = stat
			i++
			fmt.Print(logger.ANSI_CLEAR)
			logger.LInfo("detected change, recompiling... (" + fmt.Sprint(i) + ")")
			executor(fileName)
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func Run(fileName string) {
	start := time.Now()

	if cli.GetFlag(cli.ARGUMENTS, "preprocessor-enabled") {
		logger.LInfo("preprocessor enabled, starting...")
		preprocessor.Process(cli.ARGUMENTS, fileName)
		fileName = fileName + ".fleck"
	}

	logger.LDebug("starting scanner")
	lexerStart := time.Now()
	s := scanner.New(fileName)
	tokens := s.Lex()
	logger.LDebug("lexed " + fmt.Sprint(len(tokens)) + " token, took " + time.Since(lexerStart).String())
	if logger.DEBUG {
		s.PrintTokens()
	}

	logger.LDebug("starting parser")
	parserStart := time.Now()
	p := parser.New(tokens)
	result := p.Parse()
	logger.LDebug("parsed " + fmt.Sprint(len(result)) + " items, took " + time.Since(parserStart).String())
	logger.LDebug("parsed tags:", result)

	var toc string
	if cli.GetFlag(cli.ARGUMENTS, "toc") {
		logger.LDebug("generating toc...")
		toc = p.GenerateToc()
	}

	if cli.GetFlag(cli.ARGUMENTS, "no-template") {
		// TODO: see generator/generator.go:L168
		generator.WritePlain(fileName, result, toc)
	} else {
		generator.WriteTemplate(fileName, result, toc)
	}

	logger.LInfo("compiled '" + fileName + "', took: " + time.Since(start).String())

	defer func() {
		if cli.GetFlag(cli.ARGUMENTS, "preprocessor-enabled") {
			if cli.GetFlag(cli.ARGUMENTS, "keep-temp") {
				return
			}
			logger.LDebug("cleanup, removing: '" + fileName + "'")
			err := os.Remove(fileName)
			if err != nil {
				logger.LWarn(err.Error())
			}
		}
	}()
}
