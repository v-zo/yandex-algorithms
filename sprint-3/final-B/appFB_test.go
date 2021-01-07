package main

import (
	"reflect"
	"testing"
)

//func TestSolution(t *testing.T) {
//	cases := map[string]string{
//		"5 8 9 5 6 8 4 3 2 9 8 1":     "-1",
//	}
//
//	for k, v := range cases {
//		sr := strings.NewReader(k)
//		reader := bufio.NewReader(sr)
//		var wr strings.Builder
//		writer := bufio.NewWriter(&wr)
//
//		Solve(reader, writer)
//
//		res := strings.Trim(wr.String(), "\n")
//		if v != res {
//			t.Errorf("\ncase:\n%s\n got: %s\nwant: %s", k, res, v)
//		}
//	}
//}

type Case struct {
	input    []int
	expected []int
}

func TestQuickSort(t *testing.T) {
	cases := []Case{
		{input: []int{3, 2, 1}, expected: []int{1, 2, 3}},
		{input: []int{4, 3, 2, 1}, expected: []int{1, 2, 3, 4}},
		{input: []int{4, 3, 3, 1}, expected: []int{1, 3, 3, 4}},
		{input: []int{3, 2, 4, 1}, expected: []int{1, 2, 3, 4}},
		{input: []int{3, 5, 2, 4, 1}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{0, 0, 1, 0, 0, 1}, expected: []int{0, 0, 0, 0, 1, 1}},
	}

	for _, v := range cases {
		cp := append([]int{}, v.input...)
		var lb Interface = Leaderboard{cp}
		quickSort(lb, 0, lb.Len()-1)

		if !reflect.DeepEqual(v.expected, cp) {
			t.Errorf("\ncase:\n%v\n got: %v\nwant: %v", v.input, cp, v.expected)
		}
	}
}
