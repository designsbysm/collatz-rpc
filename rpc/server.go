package rpc

import (
	"fmt"
	"net"
	"os"
	"os/signal"

	"github.com/designsbysm/collatzrpc/collatzpb"
	"github.com/designsbysm/timber/v2"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

type server struct {
	collatzpb.UnimplementedCollatzServiceServer
}

func Server() error {
	address := viper.GetString("rpc.address")
	tls := viper.GetBool("rpc.tls")

	// options
	opts := []grpc.ServerOption{}
	if tls {
		certFile := viper.GetString("ssl.cert")
		keyFile := viper.GetString("ssl.key")

		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			return fmt.Errorf("RPC: %s", err.Error())
		}

		opts = append(opts, grpc.Creds(creds))
	}

	// create
	s := grpc.NewServer(opts...)
	collatzpb.RegisterCollatzServiceServer(s, &server{})

	if viper.GetBool("rpc.reflection") {
		reflection.Register(s)
	}

	// run
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("RPC: %s", err.Error())
	}

	go func() {
		if err := s.Serve(listener); err != nil {
			panic(fmt.Errorf("RPC: %s", err.Error()))
		}
	}()

	// notify
	security := ""
	if tls {
		security = " (TLS)"
	}
	timber.Info(fmt.Sprintf("RPC: listening on %s%s", address, security))

	// wait for ^c
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	// close
	s.Stop()
	listener.Close()
	timber.Info("RPC: closed")

	return nil
}
