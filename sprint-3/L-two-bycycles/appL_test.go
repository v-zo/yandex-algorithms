package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	cases := map[string]string{
		"6\n1 2 4 4 6 8\n3":  "3 5",
		"6\n1 2 1 1 6 8\n3":  "5 6",
		"2\n1 2\n1":          "1 2",
		"2\n1 2\n2":          "2 -1",
		"6\n1 2 4 4 4 4\n3":  "3 -1",
		"6\n1 2 4 4 4 4\n10": "-1 -1",
		"1\n1\n1":            "1 -1",
		"1\n1\n2":            "-1 -1",
		"1\n2\n1":            "1 1",
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
