package main

import (
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
		R := product(FromInt(mCase.A), FromInt(mCase.B))

		if !reflect.DeepEqual(FromInt(mCase.E), R) {
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
		R := matrixPower(FromInt(mCase.A), mCase.n)

		if !reflect.DeepEqual(FromInt(mCase.E), R) {
			t.Errorf("\ncase:\n%v^%v\n got: %v\nwant: %v", mCase.A, mCase.n, R, mCase.E)
		}
	}
}

func TestFibonacciBigInt(t *testing.T) {
	cases := []int{1, 1, 2, 3, 5, 8, 13, 21, 34}

	for n, exp := range cases {
		r := fibonacciBigInt(n)

		if r.Cmp(newBigInt(exp)) != 0 {
			t.Errorf("\ncase: %d\n got: %d\nwant: %d", n, r, exp)
		}
	}
}

func TestFibonacciMatrix(t *testing.T) {
	for i := 0; i <= 7; i++ {
		r := fibonacciMatrix(i)
		exp := fibonacciBigInt(i)

		if r.Cmp(exp) != 0 {
			t.Errorf("\ncase:\n%d\n got: %d\nwant: %d", i, r, exp)
		}
	}
}

var from = 20000
var to = 20000

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for n := from; n <= to; n++ {
			fibonacciBigInt(n)
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
