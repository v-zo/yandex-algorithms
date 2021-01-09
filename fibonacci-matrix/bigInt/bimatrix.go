package main

import "math/big"

func newMatrixBigInt(n int) (M [][]*big.Int) {
	M = make([][]*big.Int, n)
	for i := range M {
		M[i] = make([]*big.Int, n)
		for j := range M {
			M[i][j] = NewBigInt(0)
		}
	}

	return
}

func FromInt(A [][]int) [][]*big.Int {
	n := len(A)
	M := make([][]*big.Int, n)
	for i := range M {
		M[i] = make([]*big.Int, n)
		for j := range A {
			M[i][j] = NewBigInt(A[i][j])
		}
	}

	return M
}

func createUnitMatrix(size int) [][]*big.Int {
	M := newMatrixBigInt(size)

	for i := 0; i < size; i++ {
		M[i][i] = NewBigInt(1)
	}

	return M
}

func product(A, B [][]*big.Int) [][]*big.Int {
	n := len(A)

	prod := newMatrixBigInt(n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				mult := NewBigInt(n)
				mult.Mul(A[i][k], B[k][j])
				prod[i][j] = prod[i][j].Add(prod[i][j], mult)
			}
		}
	}

	return prod
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
