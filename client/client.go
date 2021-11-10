package client

import (
	"context"

	"github.com/designsbysm/collatzrpc/collatzpb"
	"github.com/designsbysm/timber/v2"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func Go() error {
	port := viper.GetString("rpc.port")
	protocol := viper.GetString("rpc.protocol")

	opts := grpc.WithInsecure()
	if protocol == "HTTPS" {
		cert := viper.GetString("ssl.ca")
		creds, err := credentials.NewClientTLSFromFile(cert, "")
		if err != nil {
			return err
		}

		opts = grpc.WithTransportCredentials(creds)
	}

	cc, err := grpc.Dial(port, opts)
	if err != nil {
		return err
	}
	defer cc.Close()

	c := collatzpb.NewCollatzServiceClient(cc)

	req := collatzpb.SeedRequest{
		Value: 27,
	}

	res, err := c.Seed(context.Background(), &req)
	if err != nil {
		return err
	}

	timber.Debug(res.Path)

	return nil
}
