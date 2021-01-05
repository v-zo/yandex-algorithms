package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	cases := map[string]string{
		"8\n10\n0 0 0 1 3 3 5 10\n4 4 5 7 7 7 8 9 9 10": "5",
		"2\n2\n1 2\n3 4":   "2.5",
		"2\n1\n1 3\n2":     "2",
		"1\n1\n0\n0":       "0",
		"1\n1\n0\n1":       "0.5",
		"1\n2\n0\n1 2":     "1",
		"1\n2\n0\n0 0":     "0",
		"2\n3\n0 0\n0 0 2": "0",
		"2\n3\n0 1\n0 0 2": "0",
	}

	for k, v := range cases {
		sr := strings.NewReader(k)
		reader := bufio.NewReader(sr)
		var wr strings.Builder
		writer := bufio.NewWriter(&wr)

		Solve(reader, writer)

		res := strings.Trim(wr.String(), "\n")
		if v != res {
			t.Errorf("\ncase:\n%s\n got: \n%s\nwant: \n%s", k, res, v)
		}
	}
}
