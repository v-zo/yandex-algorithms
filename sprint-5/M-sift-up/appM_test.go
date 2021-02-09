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
			Result{[]int{0, 12, 6, 8, 3, 15, 7}, 5},
			Result{[]int{0, 15, 12, 8, 3, 6, 7}, 1},
		},
		{
			Result{[]int{0, 50, 45, 20, 17, 14, 10, 6, 5, 4, 33}, 10},
			Result{[]int{0, 50, 45, 20, 17, 33, 10, 6, 5, 4, 14}, 5},
		},
		{
			Result{[]int{0, 1}, 1},
			Result{[]int{0, 1}, 1},
		},
	}

	for _, cs := range cases {
		mutatedHeap := append([]int{}, cs.inp.heap...)
		newPos := siftUp(mutatedHeap, cs.inp.pos)

		if !reflect.DeepEqual(mutatedHeap, cs.out.heap) || newPos != cs.out.pos {
			t.Errorf("\n- case:\n%v\n- got: \n%v\n- want: \n%v", cs.inp, Result{mutatedHeap, newPos}, cs.out)
		}
	}
}
