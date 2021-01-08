package main

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

type MCase struct {
	A [][]int
	B [][]int
	E [][]int
}

func TestProduct(t *testing.T) {
	cases := []MCase{
		{[][]int{{1, 2}, {3, 4}}, [][]int{{2, 0}, {1, 2}}, [][]int{{4, 4}, {10, 8}}},
		{[][]int{{1, 1}, {1, 2}}, [][]int{{1, 1}, {1, 2}}, [][]int{{2, 3}, {3, 5}}},
	}

	for _, mCase := range cases {
		R := product(mCase.A, mCase.B)

		if !reflect.DeepEqual(mCase.E, R) {
			t.Errorf("\ncase:\n%v x %v\n got: %v\nwant: %v", mCase.A, mCase.B, R, mCase.E)
		}
	}
}

type MPCase struct {
	A [][]int
	n int
	E [][]int
}

func TestMatrixPower(t *testing.T) {
	cases := []MPCase{
		{[][]int{{1, 2}, {3, 4}}, 2, [][]int{{7, 10}, {15, 22}}},
		{[][]int{{3, 4}, {5, 2}}, 5, [][]int{{9323, 7484}, {9355, 7452}}},
		{[][]int{{1, 2}, {3, 4}}, 0, [][]int{{1, 0}, {0, 1}}},
		{[][]int{{1, 2}, {3, 4}}, 1, [][]int{{1, 2}, {3, 4}}},
		{[][]int{{1, 1}, {1, 0}}, 5, [][]int{{8, 5}, {5, 3}}},
	}

	for _, mCase := range cases {
		R := matrixPower(mCase.A, mCase.n)

		if !reflect.DeepEqual(mCase.E, R) {
			t.Errorf("\ncase:\n%v^%v\n got: %v\nwant: %v", mCase.A, mCase.n, R, mCase.E)
		}
	}
}

func TestPowInt(t *testing.T) {
	cases := [][]int{
		{3, 3},
		{2, 2},
		{2, 8},
		{2, 5},
		{3, 8},
		{3, 4},
		{2, 7},
	}

	for _, inp := range cases {
		x, y := inp[0], inp[1]
		r := powInt(x, y)
		exp := int(math.Pow(float64(x), float64(y)))

		if r != exp {
			t.Errorf("\ncase:\n%v^%v\n got: %v\nwant: %v", x, y, r, exp)
		}
	}
}

func TestFibonacci(t *testing.T) {
	cases := [][]int{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{6, 13},
	}

	for _, inp := range cases {
		n, exp := inp[0], inp[1]
		r := fibonacci(n, 1, 0)

		if r != exp {
			t.Errorf("\ncase:\n%d\n got: %d\nwant: %d", n, r, exp)
		}
	}
}

func TestFibonacciMatrix(t *testing.T) {
	for i := 0; i <= 20; i++ {
		r := fibonacciMatrix(i)
		exp := fibonacci(i, 1, 0)
		fmt.Println(r, exp)

		if r != exp {
			t.Errorf("\ncase:\n%d\n got: %d\nwant: %d", i, r, exp)
		}
	}
}

var from = 50
var to = 92

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for n := from; n <= to; n++ {
			fibonacci(n, 1, 0)
		}
	}
}

func BenchmarkFibonacciMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for n := from; n <= to; n++ {
			fibonacciMatrix(n)
		}
	}
}
