package main

import (
	"flag"

	"github.com/designsbysm/collatzrpc/client"
	"github.com/designsbysm/collatzrpc/server"
)

func main() {
	var clientMode bool
	var serverMode bool

	flag.BoolVar(&clientMode, "client", false, "client mode")
	flag.BoolVar(&serverMode, "server", false, "server mode")
	flag.Parse()

	if err := config(); err != nil {
		panic(err)
	}

	if err := logger(); err != nil {
		panic(err)
	}

	if serverMode {
		if err := server.Go(); err != nil {
			panic(err)
		}
	} else if clientMode {
		if err := client.Go(); err != nil {
			panic(err)
		}
	}
}
