package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	cases := map[string]string{
		"3\n2 3 4\n2":   "1",
		"3\n1 3 1\n1":   "0",
		"3\n1 3 5\n3":   "4",
		"4\n1 3 5 6\n6": "5",
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
