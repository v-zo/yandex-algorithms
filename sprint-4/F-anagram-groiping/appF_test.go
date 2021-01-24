package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	cases := map[string]string{
		"6\ntan eat tea ate nat bat": "0 4\n1 2 3\n5",
		"6\nx z y a s m":             "0\n1\n2\n3\n4\n5",
		"2\nrjvew mkzph":             "0\n1",
		"2\naxd bxc":                 "0\n1",
	}

	for k, v := range cases {
		sr := strings.NewReader(k)
		reader := bufio.NewReader(sr)
		var wr strings.Builder
		writer := bufio.NewWriter(&wr)

		solveProblem(reader, writer)

		res := strings.Trim(wr.String(), "\n")
		if v != res {
			t.Errorf("\ncase:\n%s\n got: \n%s\nwant: \n%s", k, res, v)
		}
	}
}
