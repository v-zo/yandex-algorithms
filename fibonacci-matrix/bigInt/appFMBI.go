package main

import (
	"fmt"
	"math/big"
	"testing"
)

const (
	from = 200
	to   = 2200
)

func main() {
	delta := (to - from) / 10
	var (
		data, dataM []int64
	)

	for n := from; n < to; n += delta {
		fmt.Printf("\033[2K\r%s%d...", "running fibonacciBigInt bench, n=", n)
		res := testing.Benchmark(func(b *testing.B) {
			BenchFib(b.N, n, fibonacciBigInt)
		})
		data = append(data, res.NsPerOp())

		fmt.Printf("\033[2K\r%s%d...", "running fibonacciMatrix bench, n=", n)
		resM := testing.Benchmark(func(b *testing.B) {
			BenchFib(b.N, n, fibonacciMatrix)
		})
		dataM = append(dataM, resM.NsPerOp())
	}

	fmt.Printf("\033[2K\r")
	fmt.Println(data)
	fmt.Println(dataM)
}

func BenchFib(N, p int, fn func(int) *big.Int) {
	for i := 0; i < N; i++ {
		fn(p)
	}
}

func newBigInt(n int) *big.Int {
	return new(big.Int).SetUint64(uint64(n))
}

func findFib(n int, p, p0 *big.Int) *big.Int {
	for i := 0; i < n; i++ {
		p, p0 = p0, p
		p.Add(p, p0)
	}

	return p
}

func fibonacciBigInt(n int) *big.Int {
	return findFib(n, newBigInt(1), newBigInt(0))
}

type MatrixBigInt struct {
	value [][]*big.Int
}

func NewMatrixBigInt(n int) MatrixBigInt {
	M := make([][]*big.Int, n)
	for i := range M {
		M[i] = make([]*big.Int, n)
		for j := range M {
			M[i][j] = newBigInt(0)
		}
	}

	return MatrixBigInt{M}
}

func FromInt(A [][]int) [][]*big.Int {
	n := len(A)
	M := make([][]*big.Int, n)
	for i := range M {
		M[i] = make([]*big.Int, n)
		for j := range A {
			M[i][j] = newBigInt(A[i][j])
		}
	}

	return M
}

func createUnitMatrix(size int) [][]*big.Int {
	M := NewMatrixBigInt(size)

	for i := 0; i < size; i++ {
		M.value[i][i] = newBigInt(1)
	}

	return M.value
}

func product(A, B [][]*big.Int) [][]*big.Int {
	n := len(A)

	prod := NewMatrixBigInt(n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				mult := newBigInt(n)
				mult.Mul(A[i][k], B[k][j])
				prod.value[i][j] = prod.value[i][j].Add(prod.value[i][j], mult)
			}
		}
	}

	return prod.value
}

func matrixPower(A [][]*big.Int, n int) [][]*big.Int {
	if n == 1 {
		return A
	}

	p := createUnitMatrix(len(A))

	if n == 0 {
		return p
	}

	for n != 0 {
		if n%2 != 0 {
			p = product(p, A)
		}

		A = product(A, A)
		n /= 2
	}

	return p
}

func fibonacciMatrix(n int) *big.Int {
	if n == 0 || n == 1 {
		return newBigInt(1)
	}

	F := FromInt([][]int{{1, 1}, {1, 0}})

	return matrixPower(F, n)[0][0]
}
