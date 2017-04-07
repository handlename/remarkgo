package main

import (
	"flag"
	"fmt"
	"os"

	remark "github.com/handlename/remarkgo"
)

func main() {
	var (
		addr string
	)

	flag.StringVar(&addr, "l", "localhost:8080", "listen addr and port.")
	flag.Parse()

	server, err := remark.NewServer(addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to start server: %s", err.Error())
		os.Exit(1)
	}

	fmt.Printf("listen on http://%s\n", addr)

	if server.Serve() != nil {
		fmt.Fprintf(os.Stderr, "failed to server slideshow: %s", err.Error())
		os.Exit(1)
	}

	return
}
