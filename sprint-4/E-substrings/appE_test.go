package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	cases := map[string]int{
		"abcabcbb": 3,
		"bbbbb":    1,
		"acbc":     3,
		"aabxcx":   3,
		"abccdef":  3,
	}

	for k, v := range cases {
		res := sol(k)
		if v != res {
			t.Errorf("\ncase:\n%s\n got: \n%v\nwant: \n%v", k, res, v)
		}
	}
}
