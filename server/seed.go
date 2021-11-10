package server

import (
	"context"

	"github.com/designsbysm/collatzrpc/collatzpb"
	"github.com/designsbysm/collatzrpc/server/collatz"
)

func (*server) Seed(ctx context.Context, in *collatzpb.SeedRequest) (*collatzpb.SeedResponse, error) {
	seed := in.GetValue()
	return collatz.Hailstone(seed)
}
