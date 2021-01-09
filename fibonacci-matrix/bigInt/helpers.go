package main

import "math/big"

func NewBigInt(n int) *big.Int {
	return new(big.Int).SetUint64(uint64(n))
}
