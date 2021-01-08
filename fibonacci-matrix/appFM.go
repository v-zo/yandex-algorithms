package main

import "fmt"

func main() {
	//A := [][]int{[]int{1,2},[]int{3,4}}
	//B := [][]int{[]int{2,0},[]int{1,2}}
	//fmt.Println(product(A,B))

	//fmt.Println(powInt(2, 8))

	fmt.Println(fibonacciMatrix(3))
}

func powInt(x, y int) (p int) {
	p = 1

	for y != 0 {
		if y%2 != 0 {
			p *= x
		}

		x *= x
		y /= 2
	}

	return
}

func createUnitMatrix(size int) (M [][]int) {
	M = NewMatrix(size)

	for i := 0; i < size; i++ {
		M[i][i] = 1
	}

	return
}

func matrixPower(A [][]int, n int) (p [][]int) {
	p = createUnitMatrix(len(A))

	for n != 0 {
		if n%2 != 0 {
			p = product(A, A)
		}

		A = product(A, A)
		n /= 2
	}

	return
}

func fibonacciMatrix(n int) int {
	F := [][]int{{1, 1}, {0, 1}}

	return matrixPower(F, n)[0][0]
}

func product(A, B [][]int) (prod [][]int) {
	n := len(A)

	prod = NewMatrix(n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				prod[i][j] += A[i][k] * B[k][j]
			}
		}
	}

	return
}

func NewMatrix(n int) [][]int {
	M := make([][]int, n)
	for i := range M {
		M[i] = make([]int, n)
	}

	return M
}

func fibonacci(n, p, p0 int) int {
	for i := 0; i < n; i++ {
		p, p0 = p0, p
		p = p + p0
	}

	return p
}
