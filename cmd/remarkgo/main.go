package main

import (
	"flag"
	"fmt"
	"os"

	remark "github.com/handlename/remarkgo"
)

func main() {
	var (
		addr    string
		src     string
		cssPath string
	)

	// TODO: read default value from constants in remarkgo.
	flag.StringVar(&src, "s", "index.md", "src markdown file.")
	flag.StringVar(&addr, "l", "localhost:8080", "listen addr and port.")
	flag.StringVar(&cssPath, "c", "", "path to custom.css")
	flag.Parse()

	var err error
	options := []remark.ServerOption{}

	if src == "" {
		handleError(fmt.Errorf("src path required"), "src path must not be empty")
	}

	options = append(options, remark.ServerOptionSrcPath(src))

	if cssPath != "" {
		options = append(options, remark.ServerOptionCustomCSSPath(cssPath))
	}

	server, err := remark.NewServer(addr, options...)
	handleError(err, "failed to start server")

	fmt.Printf("listen on http://%s\n", addr)
	handleError(server.Serve(), "failed to server slideshow")

	return
}

func handleError(err error, msg string) {
	if err == nil {
		return
	}

	fmt.Fprintf(os.Stderr, "%s: %s", msg, err.Error())
	os.Exit(1)
}
