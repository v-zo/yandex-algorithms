package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	cases := map[string]string{
		"7\n1 2 3 1 2 3 4\n3": "1 2 3",
		"6\n1 1 1 2 2 3\n1":   "1",
		"1\n4\n1":             "4",
		"5\n4 4 3 2 1 0\n3":   "4 0 1",
	}

	for k, v := range cases {
		sr := strings.NewReader(k)
		reader := bufio.NewReader(sr)
		var wr strings.Builder
		writer := bufio.NewWriter(&wr)

		Solve(reader, writer)

		res := strings.Trim(wr.String(), "\n")
		if v != res {
			t.Errorf("\ncase:\n%s\n got: %s\nwant: %s", k, res, v)
		}
	}
}
