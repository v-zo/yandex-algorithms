package main

import (
	"reflect"
	"testing"
)

type Result struct {
	heap []int
	pos  int
}

type Case struct {
	inp Result
	out Result
}

func TestSolution(t *testing.T) {
	cases := []Case{
		{
			Result{[]int{0, 12, 1, 8, 3, 4, 7}, 2},
			Result{[]int{0, 12, 4, 8, 3, 1, 7}, 5},
		},
		{
			Result{[]int{0, 14, 50, 20, 17, 45, 10, 6, 5, 4}, 1},
			Result{[]int{0, 50, 45, 20, 17, 14, 10, 6, 5, 4}, 5},
		},
		{
			Result{[]int{0, 1}, 1},
			Result{[]int{0, 1}, 1},
		},
		{
			Result{[]int{0, 10, 2, 6, 8}, 2},
			Result{[]int{0, 10, 8, 6, 2}, 4},
		},
	}

	for _, cs := range cases {
		mutatedHeap := append([]int{}, cs.inp.heap...)
		newPos := siftDown(mutatedHeap, cs.inp.pos)

		if !reflect.DeepEqual(mutatedHeap, cs.out.heap) || newPos != cs.out.pos {
			t.Errorf("\n- case:\n%v\n- got: \n%v\n- want: \n%v", cs.inp, Result{mutatedHeap, newPos}, cs.out)
		}
	}
}
