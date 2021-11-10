package main

import (
	"context"

	"github.com/designsbysm/collatzrpc/collatzpb"
	"github.com/designsbysm/timber/v2"
)

func seed(client collatzpb.CollatzServiceClient, value int64) {

	req := collatzpb.SeedRequest{
		Value: value,
	}

	res, err := client.Seed(context.Background(), &req)
	if err != nil {
		panic(err)
	}

	timber.Debug(res.Path)
}
