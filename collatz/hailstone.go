package collatz

import "github.com/designsbysm/collatzrpc/collatzpb"

func Hailstone(seed int64) (*collatzpb.SeedResponse, error) {
	var result collatzpb.SeedResponse
	var err error

	result.Path = append(result.Path, seed)

	for seed > 1 {
		seed, err = stone(seed)
		if err != nil {
			return &result, err
		}

		result.Path = append(result.Path, seed)
	}

	return &result, nil
}
