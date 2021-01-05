package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	cases := map[string]string{
		"4\n7 8\n7 8\n2 3\n6 10":           "2 3\n6 10",
		"4\n2 3\n5 6\n3 4\n3 4":            "2 4\n5 6",
		"6\n1 3\n3 5\n4 6\n5 6\n2 4\n7 10": "1 6\n7 10",
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
