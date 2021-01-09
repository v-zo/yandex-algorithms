package main

import (
	"math/big"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go writeBenchData(fibonacciBigInt, "data.txt", &wg)
	go writeBenchData(fibonacciMatrix, "dataM.txt", &wg)
	wg.Wait()
}

func fibonacciBigInt(n int) *big.Int {
	return findFib(n, NewBigInt(1), NewBigInt(0))
}

func fibonacciMatrix(n int) *big.Int {
	if n == 0 || n == 1 {
		return NewBigInt(1)
	}

	F := FromInt([][]int{{1, 1}, {1, 0}})

	return matrixPower(F, n)[0][0]
}

func findFib(f int, p, p0 *big.Int) *big.Int {
	for i := 0; i < f; i++ {
		p, p0 = p0, p
		p.Add(p, p0)
	}

	return p
}
