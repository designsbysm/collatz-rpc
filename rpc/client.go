package rpc

import (
	"fmt"

	"github.com/designsbysm/timber/v2"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func Client() (*grpc.ClientConn, error) {
	port := viper.GetString("rpc.port")
	protocol := viper.GetString("rpc.protocol")

	opts := grpc.WithInsecure()
	if protocol == "HTTPS" {
		certFile := viper.GetString("ssl.ca")
		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			return nil, fmt.Errorf("RPC: %s", err.Error())
		}

		opts = grpc.WithTransportCredentials(creds)
	}

	timber.Info(fmt.Sprintf("RPC: sending %s on %s", protocol, port))

	return grpc.Dial(port, opts)
}
