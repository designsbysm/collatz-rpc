package rpc

import (
	"fmt"
	"net"

	"github.com/designsbysm/collatzrpc/collatzpb"
	"github.com/designsbysm/timber/v2"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
	collatzpb.UnimplementedCollatzServiceServer
}

func Server() error {
	address := viper.GetString("rpc.address")
	tls := viper.GetBool("rpc.tls")

	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("RPC: %s", err.Error())
	}

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

	s := grpc.NewServer(opts...)
	collatzpb.RegisterCollatzServiceServer(s, &server{})

	security := ""
	if tls {
		security = " (TLS)"
	}
	timber.Info(fmt.Sprintf("RPC: listening on %s%s", address, security))

	if err := s.Serve(listener); err != nil {
		return fmt.Errorf("RPC: %s", err.Error())
	}

	timber.Info("RPC: closing")

	return nil
}
