package main

import (
	"flag"

	"github.com/designsbysm/collatzrpc/collatzpb"
	"github.com/designsbysm/collatzrpc/rpc"
	"github.com/designsbysm/timber/v2"
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
		if err := rpc.Server(); err != nil {
			panic(err)
		}
	} else if clientMode {
		connection, err := rpc.Client()
		if err != nil {
			panic(err)
		}
		defer connection.Close()

		client := collatzpb.NewCollatzServiceClient(connection)

		seed(client, 27)
		seed(client, 54)
		seed(client, 1024)
		seed(client, 1246749)

		timber.Info("RPC: complete")
	}
}
