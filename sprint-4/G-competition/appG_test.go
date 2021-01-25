package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	cases := map[string]string{
		"2\n0 1\n":                  "2",
		"3\n0 1 0\n":                "2",
		"10\n0 0 1 0 1 1 1 0 0 0\n": "8",
		"10\n0 1 0 0 0 0 1 0 0 0\n": "2",
		"10\n0 1 0 0 0 0 1 1 0 0\n": "4",
		"2\n0 0\n":                  "0",
		"0\n":                       "0",
	}

	for k, v := range cases {
		sr := strings.NewReader(k)
		reader := bufio.NewReader(sr)
		var wr strings.Builder
		writer := bufio.NewWriter(&wr)

		Solve(reader, writer)

		writer.Flush()

		res := strings.Trim(wr.String(), "\n")
		if v != res {
			t.Errorf("\n- case:\n%s- got: \n%s\n- want: \n%s", k, res, v)
		}
	}
}
