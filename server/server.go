package server

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

func Go() error {
	port := viper.GetString("rpc.port")
	protocol := viper.GetString("rpc.protocol")

	listener, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("RPC: %s", err.Error())
	}

	timber.Info(fmt.Sprintf("RPC: serving %s on %s", protocol, port))

	opts := []grpc.ServerOption{}
	if protocol == "HTTPS" {
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

	if err := s.Serve(listener); err != nil {
		return fmt.Errorf("RPC: %s", err.Error())
	}

	timber.Info("RPC: closing")

	return nil
}
