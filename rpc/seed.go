package rpc

import (
	"context"
	"fmt"

	"github.com/designsbysm/collatzrpc/collatz"
	"github.com/designsbysm/collatzrpc/collatzpb"
	"github.com/designsbysm/timber/v2"
)

func (*server) Seed(ctx context.Context, in *collatzpb.SeedRequest) (*collatzpb.SeedResponse, error) {
	seed := in.GetValue()
	timber.Debug(fmt.Sprintf("RPC: Seed %d", seed))

	return collatz.Hailstone(seed)
}
