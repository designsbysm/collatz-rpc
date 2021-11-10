package collatz

import (
	"errors"
)

func stone(seed int64) (int64, error) {
	var value int64

	if seed%2 == 0 { //even
		value = seed / 2
	} else { //odd
		value = (seed * 3) + 1

		if value <= seed {
			return 0, errors.New("value overflow")
		}
	}

	return value, nil
}
