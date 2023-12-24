package utils

import (
	"fmt"
	"math/big"
)

type Signature struct {
	R *big.Int
	S *big.Int
}

func (signature *Signature) String() string {
	return fmt.Sprintf("%x%x", signature.R, signature.S)
}
