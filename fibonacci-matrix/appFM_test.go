package main

import (
	"math"
	"reflect"
	"testing"
)

func TestProduct(t *testing.T) {
	A := [][]int{[]int{1, 2}, []int{3, 4}}
	B := [][]int{[]int{2, 0}, []int{1, 2}}
	E := [][]int{[]int{4, 4}, []int{10, 8}}

	R := product(A, B)

	if !reflect.DeepEqual(E, R) {
		t.Errorf("\ncase:\n%v x %v\n got: %v\nwant: %v", A, B, R, E)
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
