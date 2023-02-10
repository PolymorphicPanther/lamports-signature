package keys

import (
	"errors"
	"math/big"
)

type Int256 struct{ *big.Int }

func CreateInt256(b []byte) (*Int256, error) {
	if len(b) != 32 {
		return nil, errors.New("expected slice of length 24")
	}

	return &Int256{
		Int: (&big.Int{}).SetBytes(b),
	}, nil
}
