package rpc

import (
	"fmt"

	"github.com/designsbysm/timber/v2"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func Client() (*grpc.ClientConn, error) {
	address := viper.GetString("rpc.address")
	tls := viper.GetBool("rpc.tls")

	opts := grpc.WithInsecure()
	if tls {
		certFile := viper.GetString("ssl.ca")
		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			return nil, fmt.Errorf("RPC: %s", err.Error())
		}

		opts = grpc.WithTransportCredentials(creds)
	}

	security := ""
	if tls {
		security = " (TLS)"
	}
	timber.Info(fmt.Sprintf("RPC: sending on %s%s", address, security))

	return grpc.Dial(address, opts)
}
